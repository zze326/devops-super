package host_terminal_session

import (
	"bytes"
	"context"
	"devops-super/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"io"
	"os"
	"time"
	"unicode/utf8"
)

func (s *sHostTerminalSession) Replay(ctx context.Context, id int) error {
	glog.Infof(ctx, "建立回放 Web 终端记录 Websocket 连接")
	defer func() {
		if err := recover(); err != nil {
			glog.Error(ctx, err)
		}
	}()
	eSession, err := s.Get(ctx, &do.HostTerminalSession{Id: id})
	if err != nil {
		return err
	}

	file, err := os.Open(eSession.Filepath)
	if err != nil {
		return err
	}

	type msgObj struct {
		Type     string  `json:"type"`
		Speed    float32 `json:"speed"`
		Progress float32 `json:"progress"`
	}

	type respObj struct {
		Total int64  `json:"total"`
		Sent  int64  `json:"sent"`
		Data  string `json:"data"`
		Clear bool   `json:"clear"`
	}

	var (
		closeChan    = make(chan bool)
		stopChan     = make(chan bool)
		resumeChan   = make(chan bool)
		speedChan    = make(chan float32)
		progressChan = make(chan float32)
		writeBuffer  bytes.Buffer
	)

	defer func() {
		defer close(closeChan)
		defer close(stopChan)
		defer close(resumeChan)
		defer close(speedChan)
		defer close(progressChan)
	}()

	ws, err := g.RequestFromCtx(ctx).WebSocket()
	if err != nil {
		return err
	}
	defer ws.Close()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				glog.Error(ctx, err)
			}
		}()
		for {
			msg := new(msgObj)
			if err := ws.ReadJSON(&msg); err != nil {
				glog.Errorf(ctx, "接收消息失败: %v", err)
				break
			}
			if msg.Type == "close" {
				closeChan <- true // 当退出循环时，向通道发送信号
				break
			}
			if msg.Type == "pause" {
				stopChan <- true
			}
			if msg.Type == "continue" {
				resumeChan <- true
			}
			if msg.Type == "speed" {
				if msg.Speed > 0 {
					glog.Infof(ctx, "speed: %v", msg.Speed)
					speedChan <- msg.Speed
				}
			}
			if msg.Type == "progress" {
				glog.Infof(ctx, "progress: %v", msg.Progress)
				progressChan <- msg.Progress
			}
		}
	}()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	glog.Infof(ctx, "stat.Size(): %d", stat.Size())

	var (
		total            = stat.Size()
		sent     int64   = 0
		speed    float32 = 1
		oneSend  int64   = 24
		msgBytes         = make([]byte, oneSend)
	)
loop:
	for {
		select {
		case <-closeChan:
			break loop // 关闭循环
		case <-stopChan:
			<-resumeChan
		case speed = <-speedChan:
		case progress := <-progressChan:
			sendPoint := int64(float32(total) * progress)
			tmpMsgBytes := make([]byte, sendPoint)
			if _, err := file.Seek(0, 0); err != nil {
				glog.Error(ctx, err)
				break loop
			}

			if _, err := io.ReadFull(file, tmpMsgBytes); err != nil {
				glog.Error(ctx, err)
				break loop
			}

			if err := ws.WriteJSON(respObj{
				Total: total,
				Sent:  sent,
				Data:  string(tmpMsgBytes),
				Clear: true,
			}); err != nil {
				glog.Error(ctx, err)
				break loop
			}
			//allMsgBytes = fileBytes[sendPoint:]
			sent = sendPoint
		default:
			sendLen, err := io.ReadFull(file, msgBytes)
			if err == io.EOF {
				break loop
			} else if err != nil && err != io.ErrUnexpectedEOF {
				return err
			}

			sent += int64(sendLen)

			if !utf8.Valid(msgBytes) {
				writeBuffer.Write(msgBytes)
				continue loop
			} else {
				if writeBuffer.Len() > 0 {
					writeBuffer.Write(msgBytes)
					if !utf8.Valid(writeBuffer.Bytes()) {
						continue loop
					}
					msgBytes = writeBuffer.Bytes()
					writeBuffer.Reset()
				}
			}

			if err := ws.WriteJSON(respObj{
				Total: total,
				Sent:  sent,
				Data:  string(msgBytes[:sendLen]),
			}); err != nil {
				glog.Error(ctx, err)
				break loop
			}

			time.Sleep(time.Duration((100 / speed) * float32(time.Millisecond)))
		}
	}
	glog.Info(ctx, "关闭回放 Web 终端记录 Websocket 连接")
	return nil
}
