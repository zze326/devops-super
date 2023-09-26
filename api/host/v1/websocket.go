package v1

import "github.com/gogf/gf/v2/frame/g"

type WsTerminalReq struct {
	g.Meta `method:"get" path:"/host/{id}/terminal" summary:"连接终端" tags:"主机"`
	Id     int `v:"required" path:"id"`
}

type WsTerminalRes struct{}

type WsSftpFileManagerReq struct {
	g.Meta `method:"get" path:"/host/{id}/sftp-file-manager" summary:"连接 SFTP 文件管理器" tags:"主机"`
	Id     int `v:"required" path:"id"`
}

type WsSftpFileManagerRes struct{}
