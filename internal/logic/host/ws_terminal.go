package user

import (
	"context"
	"devops-super/internal/consts"
	"devops-super/internal/model/entity"
	"devops-super/utility/util"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"path"
	"time"
)

func (s *sHost) WsTerminal(ctx context.Context, in *entity.Host) error {
	wsCtx := &wsContext{
		ctx:          ctx,
		request:      g.RequestFromCtx(ctx),
		lastReadTime: gtime.Now(),
		lastPingTime: gtime.Now(),
		host:         in,
	}
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

	wsCtx.session = session
	wsCtx.startTime = gtime.Now()
	defer wsCtx.saveSession()
	go wsCtx.checkTimeout()

	if wsCtx.ws, err = wsCtx.request.WebSocket(); err != nil {
		return err
	}
	defer wsCtx.ws.Close()

	fd := int(os.Stdin.Fd())
	session.Stdout = wsCtx
	session.Stderr = wsCtx
	session.Stdin = wsCtx
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

	if wsCtx.isSaveSession() {
		wsCtx.sessionFilepath = path.Join(g.Cfg().MustGet(ctx, "host.terminal.sessionFileDir", consts.HOST_TERMINAL_SESSION_SAVE_DIRECTORY).String(), fmt.Sprintf("%d", in.Id), fmt.Sprintf("%d.sessionb", wsCtx.startTime.UnixMicro()))
		wsCtx.sessionFile, err = util.OpenOrCreateFile(wsCtx.sessionFilepath)
		if err != nil {
			glog.Errorf(ctx, "创建会话文件失败: %v", err)
		} else {
			defer wsCtx.sessionFile.Close()
		}
	}

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
		Timeout:         8 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", in.HostAddr, in.Port), config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
