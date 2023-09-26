package user

import (
	"context"
	"devops-super/internal/model/entity"
	"devops-super/utility/util"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func (s *sHost) WsTerminal(ctx context.Context, in *entity.Host) error {
	s.ctx = ctx
	s.request = g.RequestFromCtx(ctx)

	sshClient, err := s.SshClient(in)
	if err != nil {
		return err
	}
	defer func() {
		if err := sshClient.Close(); err != nil {
			glog.Error(ctx, err)
		}
	}()

	session, err := sshClient.NewSession()
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	defer func() {
		if err := session.Close(); err != nil && err != io.EOF {
			glog.Error(ctx, err)
		}
	}()

	s.session = session
	s.startTime = gtime.Now()
	if s.ws, err = s.request.WebSocket(); err != nil {
		return err
	}
	defer func() {
		if err := s.ws.Close(); err != nil {
			glog.Error(ctx, err)
		}
	}()

	fd := int(os.Stdin.Fd())
	session.Stdout = s
	session.Stderr = s
	session.Stdin = s
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	termWidth, termHeight, err := terminal.GetSize(fd)
	err = session.RequestPty("xterm", termHeight, termWidth, modes)
	if err != nil {
		glog.Error(ctx, err)
		return err
	}

	err = session.Shell()
	if err != nil {
		glog.Error(ctx, err)
		return err
	}

	go func() {
		for {
			// 超过 60 分钟没有接收到指令则断开连接
			if gtime.Now().Sub(s.lastReadTime).Seconds() > 3600 {
				if err := session.Close(); err != nil {
					glog.Error(ctx, err)
				}
				break
			}
			time.Sleep(time.Second * 10)
		}
	}()

	err = session.Wait()
	if err != nil {
		if _, ok := err.(*ssh.ExitError); !ok {
			glog.Error(ctx, err)
			return err
		}
	}
	glog.Infof(ctx, "关闭 Web 终端 Websocket 连接")
	return nil
}

// Read 接收 Web 终端的命令
func (s *sHost) Read(p []byte) (n int, err error) {
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
func (s *sHost) Write(p []byte) (n int, err error) {
	msgBytes := p

	//if s.isSaveSession {
	//	if _, err := s.sessionFile.Write(p); err != nil {
	//		g.Logger.Errorf("写入会话记录到文件失败, err: %v", err)
	//	}
	//}
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

func (s *sHost) SshClient(in *entity.Host) (*ssh.Client, error) {
	var authMethods []ssh.AuthMethod
	if in.UseKey {
		signer, err := ssh.ParsePrivateKey([]byte(in.PrivateKey))
		if err != nil {
			return nil, err
		}
		authMethods = append(authMethods, ssh.PublicKeys(signer))
	} else {
		authMethods = append(authMethods, ssh.Password(in.Password))
	}

	config := &ssh.ClientConfig{
		User:            in.Username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         15 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", in.HostAddr, in.Port), config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *sHost) SftpClient(in *entity.Host) (*sftp.Client, error) {
	client, err := s.SshClient(in)
	if err != nil {
		return nil, err
	}
	// 创建 SFTP 客户端
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return nil, err
	}
	return sftpClient, nil
}
