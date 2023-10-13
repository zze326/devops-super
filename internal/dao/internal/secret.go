// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SecretDao is the data access object for table secret.
type SecretDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SecretColumns // columns contains all the column names of Table for convenient usage.
}

// SecretColumns defines and stores column names for table secret.
type SecretColumns struct {
	Id        string //
	Name      string // 名称
	Type      string // 类型:1-git认证,2-Kubernetes config
	Content   string // 认证配置内容
	UpdatedAt string // 更新时间
}

// secretColumns holds the columns for table secret.
var secretColumns = SecretColumns{
	Id:        "id",
	Name:      "name",
	Type:      "type",
	Content:   "content",
	UpdatedAt: "updated_at",
}

// NewSecretDao creates and returns a new DAO object for table data access.
func NewSecretDao() *SecretDao {
	return &SecretDao{
		group:   "default",
		table:   "secret",
		columns: secretColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SecretDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SecretDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SecretDao) Columns() SecretColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SecretDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SecretDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SecretDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
