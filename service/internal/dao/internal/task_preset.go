// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskPresetDao is the data access object for table task_preset.
type TaskPresetDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TaskPresetColumns // columns contains all the column names of Table for convenient usage.
}

// TaskPresetColumns defines and stores column names for table task_preset.
type TaskPresetColumns struct {
	Id      string //
	Uuid    string //
	Name    string // 任务名
	Type    string // 任务类型
	Content string // 任务内容
	Peers   string // 主机列表
	Creater string // 创建人
	Created string // 创建时间
	Updated string // 更新时间
}

//  taskPresetColumns holds the columns for table task_preset.
var taskPresetColumns = TaskPresetColumns{
	Id:      "id",
	Uuid:    "uuid",
	Name:    "name",
	Type:    "type",
	Content: "content",
	Peers:   "peers",
	Creater: "creater",
	Created: "created",
	Updated: "updated",
}

// NewTaskPresetDao creates and returns a new DAO object for table data access.
func NewTaskPresetDao() *TaskPresetDao {
	return &TaskPresetDao{
		group:   "default",
		table:   "task_preset",
		columns: taskPresetColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskPresetDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskPresetDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskPresetDao) Columns() TaskPresetColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskPresetDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskPresetDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskPresetDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
