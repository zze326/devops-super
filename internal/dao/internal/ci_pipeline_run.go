// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CiPipelineRunDao is the data access object for table ci_pipeline_run.
type CiPipelineRunDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns CiPipelineRunColumns // columns contains all the column names of Table for convenient usage.
}

// CiPipelineRunColumns defines and stores column names for table ci_pipeline_run.
type CiPipelineRunColumns struct {
	Id         string //
	PipelineId string // 流水线 id
	PodName    string // Pod 名称
	Namespace  string // 名称空间
	Status     string // 状态:0-运行中,1:成功,2:失败,3:取消
	UpdatedAt  string // 更新时间
	CreatedAt  string // 创建时间
}

// ciPipelineRunColumns holds the columns for table ci_pipeline_run.
var ciPipelineRunColumns = CiPipelineRunColumns{
	Id:         "id",
	PipelineId: "pipeline_id",
	PodName:    "pod_name",
	Namespace:  "namespace",
	Status:     "status",
	UpdatedAt:  "updated_at",
	CreatedAt:  "created_at",
}

// NewCiPipelineRunDao creates and returns a new DAO object for table data access.
func NewCiPipelineRunDao() *CiPipelineRunDao {
	return &CiPipelineRunDao{
		group:   "default",
		table:   "ci_pipeline_run",
		columns: ciPipelineRunColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CiPipelineRunDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CiPipelineRunDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CiPipelineRunDao) Columns() CiPipelineRunColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CiPipelineRunDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CiPipelineRunDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CiPipelineRunDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
