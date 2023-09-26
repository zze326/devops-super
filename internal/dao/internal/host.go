// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HostDao is the data access object for table host.
type HostDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns HostColumns // columns contains all the column names of Table for convenient usage.
}

// HostColumns defines and stores column names for table host.
type HostColumns struct {
	Id          string //
	Name        string // 名称
	HostAddr    string // 主机名或IP
	Port        string // 端口
	Username    string // 用户名
	Password    string // 密码
	PrivateKey  string // 私钥
	UseKey      string // 是否使用公钥连接
	Desc        string // 描述
	SaveSession string // 是否保存会话
	UpdatedAt   string // 更新时间
	HostGroupId string // 主机组 id
}

// hostColumns holds the columns for table host.
var hostColumns = HostColumns{
	Id:          "id",
	Name:        "name",
	HostAddr:    "host_addr",
	Port:        "port",
	Username:    "username",
	Password:    "password",
	PrivateKey:  "private_key",
	UseKey:      "use_key",
	Desc:        "desc",
	SaveSession: "save_session",
	UpdatedAt:   "updated_at",
	HostGroupId: "host_group_id",
}

// NewHostDao creates and returns a new DAO object for table data access.
func NewHostDao() *HostDao {
	return &HostDao{
		group:   "default",
		table:   "host",
		columns: hostColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HostDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HostDao) Columns() HostColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HostDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
