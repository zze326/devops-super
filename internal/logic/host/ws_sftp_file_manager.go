package user

import (
	"bytes"
	"context"
	"devops-super/internal/model/entity"
	"devops-super/utility/util"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/pkg/sftp"
	"math"
	"os"
	"path"
	"strings"
)

func (s *sHost) WsSftpFileManager(ctx context.Context, in *entity.Host) (err error) {
	wsCtx := &wsContext{
		ctx:          ctx,
		request:      g.RequestFromCtx(ctx),
		lastPingTime: gtime.Now(),
		lastReadTime: gtime.Now(),
	}
	if wsCtx.ws, err = wsCtx.request.WebSocket(); err != nil {
		return err
	}

	sftpClient, err := s.SftpClient(in)
	if err != nil {
		return err
	}
	defer func() {
		if err := sftpClient.Close(); err != nil {
			glog.Error(ctx, err)
		}
	}()

	type fileOperate struct {
		Type            string `json:"type"`
		Path            string `json:"path"`
		Filename        string `json:"filename"`
		Data            []byte `json:"data"`
		ChunkStart      int64  `json:"chunk_start"`
		ChunkEnd        int64  `json:"chunk_end"`
		TotalSize       int64  `json:"total_size"`
		ShowHiddenFiles bool   `json:"show_hidden_files"`
	}

	type wsResp struct {
		Type      string `json:"type"`
		Success   bool   `json:"success"`
		Data      any    `json:"data"`
		Path      string `json:"path"`
		TotalSize int64  `json:"total_size"`
		ChunkEnd  int64  `json:"chunk_end"`
		Msg       string `json:"msg"`
	}

	type fileinfo struct {
		Name    string `json:"name"`
		Mode    string `json:"mode"`
		Size    int64  `json:"size"`
		ModTime string `json:"mtime"`
		AbsPath string `json:"abs_path"`
		IsDir   bool   `json:"is_dir"`
	}

	go wsCtx.checkTimeout()
	var msgBytesBuffer bytes.Buffer
	for {
		// 接受字节
		//var msgBytes []byte
		_, msgBytes, err := wsCtx.ws.ReadMessage()
		if err != nil {
			glog.Errorf(ctx, "websocket message receive error: %s", err.Error())
			break
		}
		wsCtx.lastReadTime = gtime.Now()

		msg := new(fileOperate)
		if msgBytes[0] != '{' {
			msgBytesBuffer.Write(msgBytes)
			if msgBytes[len(msgBytes)-1] != '}' {
				continue
			}
			if json.Unmarshal(msgBytesBuffer.Bytes(), &msg) != nil {
				continue
			}
			msgBytesBuffer.Reset()
		} else {
			if msgBytes[len(msgBytes)-1] != '}' {
				msgBytesBuffer.Write(msgBytes)
				continue
			}
			if err := json.Unmarshal(msgBytes, &msg); err != nil {
				msgBytesBuffer.Write(msgBytes)
				if json.Unmarshal(msgBytesBuffer.Bytes(), &msg) != nil {
					continue
				}
				msgBytesBuffer.Reset()
			}
		}

		if msg.Type == "exit" {
			break
		}
		switch msg.Type {
		// 心跳
		case "ping":
			wsCtx.lastPingTime = gtime.Now()
		// 列出目录下的文件
		case "list":
			files, err := sftpClient.ReadDir(msg.Path)
			if err != nil {
				glog.Errorf(ctx, "read dir error: %s", err.Error())
				break
			}
			var fileinfos []*fileinfo
			for _, file := range files {
				if !msg.ShowHiddenFiles && strings.HasPrefix(file.Name(), ".") {
					continue
				}
				fileinfos = append(fileinfos, &fileinfo{
					Name:    file.Name(),
					Mode:    util.FileMode(file.Mode()),
					Size:    file.Size(),
					ModTime: file.ModTime().Format("2006-01-02 15:04:05"),
					AbsPath: path.Join(msg.Path, file.Name()),
					IsDir:   file.IsDir(),
				})
			}

			if err := wsCtx.ws.WriteJSON(&wsResp{
				Type: "listData",
				Data: fileinfos,
				Path: msg.Path,
			}); err != nil {
				glog.Errorf(ctx, "websocket message send error: %s", err.Error())
				break
			} // 列出目录下的文件
		// 上传文件
		case "uploadFileChunk":
			file, err := sftpClient.OpenFile(path.Join(msg.Path, msg.Filename), os.O_CREATE|os.O_WRONLY)
			if err != nil {
				if err2 := wsCtx.ws.WriteJSON(&wsResp{
					Type:    "uploadFileChunk",
					Success: false,
					Msg:     err.Error(),
				}); err2 != nil {
					glog.Errorf(ctx, "websocket message send error: %s", err2.Error())
					break
				}
				continue
			}
			defer file.Close()

			_, err = file.WriteAt(msg.Data, msg.ChunkStart)
			if err != nil {
				glog.Errorf(ctx, "写入文件时出错：", err)
				break
			}

			if msg.ChunkEnd >= msg.TotalSize {
				if err := wsCtx.ws.WriteJSON(&wsResp{
					Type:    "uploadFileChunk",
					Success: true,
					Path:    msg.Path,
				}); err != nil {
					glog.Errorf(ctx, "websocket message send error: %s", err.Error())
					break
				}
			} else {
				glog.Infof(ctx, "正在上传文件: %s 到 %s, 进度: %.0f%%, 偏移量: %d", msg.Filename, msg.Path, math.Round(float64(msg.ChunkEnd)/float64(msg.TotalSize)*100), msg.ChunkEnd)
				if err := wsCtx.ws.WriteJSON(&wsResp{
					Type:      "uploadingFileChunk",
					Success:   true,
					ChunkEnd:  msg.ChunkEnd,
					Path:      msg.Path,
					TotalSize: msg.TotalSize,
				}); err != nil {
					glog.Errorf(ctx, "websocket message send error: %s", err.Error())
					break
				}
			}
		// 删除文件
		case "delete":
			if err := sftpClient.Remove(msg.Path); err != nil {
				if err2 := wsCtx.ws.WriteJSON(&wsResp{
					Type:    "delete",
					Success: false,
					Msg:     err.Error(),
				}); err2 != nil {
					glog.Errorf(ctx, "websocket message send error: %s", err2.Error())
					break
				}
				glog.Errorf(ctx, "删除文件失败：", err)
				continue
			}
			if err2 := wsCtx.ws.WriteJSON(&wsResp{
				Type:    "delete",
				Success: true,
			}); err2 != nil {
				glog.Errorf(ctx, "websocket message send error: %s", err2.Error())
				break
			}
		}
	}
	glog.Infof(ctx, "断开 SFTP 文件管理器 Websocket 连接")
	return nil
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
