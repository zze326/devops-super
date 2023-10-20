// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CiEnvDao is the data access object for table ci_env.
type CiEnvDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns CiEnvColumns // columns contains all the column names of Table for convenient usage.
}

// CiEnvColumns defines and stores column names for table ci_env.
type CiEnvColumns struct {
	Id                string //
	Name              string // 环境名称
	Image             string // 镜像
	SecretName        string // Kubernetes Secret 名称，拉取镜像使用
	PersistenceConfig string // 持久化配置
	UpdatedAt         string // 更新时间
}

// ciEnvColumns holds the columns for table ci_env.
var ciEnvColumns = CiEnvColumns{
	Id:                "id",
	Name:              "name",
	Image:             "image",
	SecretName:        "secret_name",
	PersistenceConfig: "persistence_config",
	UpdatedAt:         "updated_at",
}

// NewCiEnvDao creates and returns a new DAO object for table data access.
func NewCiEnvDao() *CiEnvDao {
	return &CiEnvDao{
		group:   "default",
		table:   "ci_env",
		columns: ciEnvColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CiEnvDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CiEnvDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CiEnvDao) Columns() CiEnvColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CiEnvDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CiEnvDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CiEnvDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
