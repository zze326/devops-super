// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HostTerminalSession is the golang structure of table host_terminal_session for DAO operations like Where/Data.
type HostTerminalSession struct {
	g.Meta           `orm:"table:host_terminal_session, do:true"`
	Id               interface{} //
	HostId           interface{} // 主机 ID
	HostAddr         interface{} // 主机名或IP
	HostName         interface{} // 主机名
	OperatorId       interface{} // 操作人 ID
	OperatorName     interface{} // 操作人用户名
	OperatorRealName interface{} // 操作人真实姓名
	StartTime        *gtime.Time // 会话开始时间
	Filepath         interface{} // 会话文件路径
	UpdatedAt        *gtime.Time //
}
