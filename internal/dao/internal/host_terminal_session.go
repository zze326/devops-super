// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HostTerminalSessionDao is the data access object for table host_terminal_session.
type HostTerminalSessionDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns HostTerminalSessionColumns // columns contains all the column names of Table for convenient usage.
}

// HostTerminalSessionColumns defines and stores column names for table host_terminal_session.
type HostTerminalSessionColumns struct {
	Id               string //
	HostId           string // 主机 ID
	HostAddr         string // 主机名或IP
	HostName         string // 主机名
	OperatorId       string // 操作人 ID
	OperatorName     string // 操作人用户名
	OperatorRealName string // 操作人真实姓名
	StartTime        string // 会话开始时间
	Filepath         string // 会话文件路径
	UpdatedAt        string //
}

// hostTerminalSessionColumns holds the columns for table host_terminal_session.
var hostTerminalSessionColumns = HostTerminalSessionColumns{
	Id:               "id",
	HostId:           "host_id",
	HostAddr:         "host_addr",
	HostName:         "host_name",
	OperatorId:       "operator_id",
	OperatorName:     "operator_name",
	OperatorRealName: "operator_real_name",
	StartTime:        "start_time",
	Filepath:         "filepath",
	UpdatedAt:        "updated_at",
}

// NewHostTerminalSessionDao creates and returns a new DAO object for table data access.
func NewHostTerminalSessionDao() *HostTerminalSessionDao {
	return &HostTerminalSessionDao{
		group:   "default",
		table:   "host_terminal_session",
		columns: hostTerminalSessionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HostTerminalSessionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HostTerminalSessionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HostTerminalSessionDao) Columns() HostTerminalSessionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HostTerminalSessionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HostTerminalSessionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HostTerminalSessionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
