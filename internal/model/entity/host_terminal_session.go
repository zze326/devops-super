// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HostTerminalSession is the golang structure for table host_terminal_session.
type HostTerminalSession struct {
	Id               int         `json:"id"               description:""`        //
	HostId           int         `json:"hostId"           description:"主机 ID"`   // 主机 ID
	HostAddr         string      `json:"hostAddr"         description:"主机名或IP"`  // 主机名或IP
	HostName         string      `json:"hostName"         description:"主机名"`     // 主机名
	OperatorId       int         `json:"operatorId"       description:"操作人 ID"`  // 操作人 ID
	OperatorName     string      `json:"operatorName"     description:"操作人用户名"`  // 操作人用户名
	OperatorRealName string      `json:"operatorRealName" description:"操作人真实姓名"` // 操作人真实姓名
	StartTime        *gtime.Time `json:"startTime"        description:"会话开始时间"`  // 会话开始时间
	Filepath         string      `json:"filepath"         description:"会话文件路径"`  // 会话文件路径
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:""`        //
}
