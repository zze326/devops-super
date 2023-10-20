// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CiPipelineDao is the data access object for table ci_pipeline.
type CiPipelineDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CiPipelineColumns // columns contains all the column names of Table for convenient usage.
}

// CiPipelineColumns defines and stores column names for table ci_pipeline.
type CiPipelineColumns struct {
	Id                  string //
	Name                string // 名称
	KubernetesConfigId  string // 关联的 Kubernetes Config id
	KubernetesNamespace string // Pod 所在命名空间
	Params              string // 构建参数
	Config              string // 配置
	Desc                string // 描述
	UpdatedAt           string // 更新时间
}

// ciPipelineColumns holds the columns for table ci_pipeline.
var ciPipelineColumns = CiPipelineColumns{
	Id:                  "id",
	Name:                "name",
	KubernetesConfigId:  "kubernetes_config_id",
	KubernetesNamespace: "kubernetes_namespace",
	Params:              "params",
	Config:              "config",
	Desc:                "desc",
	UpdatedAt:           "updated_at",
}

// NewCiPipelineDao creates and returns a new DAO object for table data access.
func NewCiPipelineDao() *CiPipelineDao {
	return &CiPipelineDao{
		group:   "default",
		table:   "ci_pipeline",
		columns: ciPipelineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CiPipelineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CiPipelineDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CiPipelineDao) Columns() CiPipelineColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CiPipelineDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CiPipelineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CiPipelineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
