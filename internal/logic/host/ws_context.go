package user

import (
	"bytes"
	"context"
	"devops-super/internal/dao"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"devops-super/utility/util"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

type wsContext struct {
	host            *entity.Host
	session         *ssh.Session
	sessionFile     *os.File
	sessionFilepath string
	startTime       *gtime.Time
	lastReadTime    *gtime.Time
	lastPingTime    *gtime.Time
	request         *ghttp.Request
	ws              *ghttp.WebSocket
	writeBuffer     bytes.Buffer
	readBuffer      bytes.Buffer
	hasInput        bool
	ctx             context.Context
}

func (s *wsContext) checkTimeout() {
	var handleClose = func() {
		if s.session != nil {
			s.session.Close()
		}
		s.ws.Close()
	}

	for {
		// 超过 5 分钟没有接收到指令则断开连接
		if gtime.Now().Sub(s.lastReadTime).Seconds() > 300 {
			handleClose()
			break
		}
		// 超过 10 秒没有接收到心跳则断开连接
		if gtime.Now().Sub(s.lastPingTime).Seconds() > 10 {
			handleClose()
			break
		}
		glog.Infof(s.ctx, "检查 websocket 连接是否超时, lastPingTime: %v, lastReadTime: %v", s.lastPingTime, s.lastReadTime)
		time.Sleep(time.Second * 10)
	}
}

func (s *wsContext) isSaveSession() bool {
	return s.host.SaveSession
}

// Read 接收 Web 终端的命令
func (s *wsContext) Read(p []byte) (n int, err error) {
	s.lastPingTime = gtime.Now()
	s.lastReadTime = gtime.Now()
	type xtermMessage struct {
		MsgType string `json:"type"`
		Input   string `json:"input"`
		Rows    uint16 `json:"rows"`
		Cols    uint16 `json:"cols"`
	}

	var xtermMsg xtermMessage
	err = s.ws.ReadJSON(&xtermMsg)
	if err != nil {
		return 0, err
	}
	if xtermMsg.MsgType == "input" {
		if cmdStr := strings.TrimSpace(s.readBuffer.String()); xtermMsg.Input == "\r" && len(cmdStr) > 0 {
			s.hasInput = true
			glog.Infof(s.ctx, "输入命令：%s", cmdStr)
			s.readBuffer.Reset()
		}
		if !util.InSlice[string]([]string{"\t", "\r", ""}, xtermMsg.Input) {
			s.readBuffer.Write([]byte(xtermMsg.Input))
		}
		copy(p, fmt.Sprintf("%s", xtermMsg.Input))
		return len(xtermMsg.Input), nil
	} else if xtermMsg.MsgType == "resize" {
		glog.Infof(s.ctx, "resize: cols=%d, rows=%d", xtermMsg.Cols, xtermMsg.Rows)
		// 改变终端大小
		if err = s.session.WindowChange(int(xtermMsg.Rows), int(xtermMsg.Cols)); err != nil {
			glog.Errorf(s.ctx, "改变终端大小失败: %s", err.Error())
		}
	} else if xtermMsg.MsgType == "close" {
		glog.Infof(s.ctx, "关闭 Web 终端")
		if err := s.session.Close(); err != nil {
			glog.Errorf(s.ctx, "关闭 Web 终端失败: %s", err.Error())
		}
		return 0, io.EOF
	}
	return 0, nil
}

// Write 响应到 Web 终端
func (s *wsContext) Write(p []byte) (n int, err error) {
	msgBytes := p

	if s.isSaveSession() {
		if _, err := s.sessionFile.Write(p); err != nil {
			glog.Errorf(s.ctx, "写入会话记录到文件失败, err: %v", err)
		}
	}
	if !utf8.Valid(msgBytes) {
		s.writeBuffer.Write(msgBytes)
		return len(p), nil
	} else {
		if s.writeBuffer.Len() > 0 {
			s.writeBuffer.Write(msgBytes)
			msgBytes = s.writeBuffer.Bytes()
			s.writeBuffer.Reset()
		}
	}

	err = s.ws.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	return len(p), nil
}

func (s *wsContext) saveSession() {
	if s.isSaveSession() && !s.hasInput {
		if err := util.EnsureFileNotExists(s.sessionFilepath); err != nil {
			glog.Errorf(s.ctx, "删除空会话文件失败: %v", err)
		}
	}

	if !s.isSaveSession() || !s.hasInput {
		return
	}
	glog.Infof(s.ctx, "保存会话文件")

	session := new(entity.HostTerminalSession)
	session.HostId = s.host.Id
	session.HostAddr = s.host.HostAddr
	session.HostName = s.host.Name
	session.OperatorName = service.CurrentUser(s.ctx).Username
	session.StartTime = s.startTime
	session.OperatorId = service.CurrentUser(s.ctx).UserId
	session.OperatorRealName = service.CurrentUser(s.ctx).RealName
	session.Filepath = s.sessionFilepath

	if _, err := dao.HostTerminalSession.Ctx(s.ctx).Insert(session); err != nil {
		glog.Errorf(s.ctx, "保存会话数据到数据库失败: %s", err.Error())
		return
	}
}
