// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppDao is the data access object for table app.
type AppDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns AppColumns // columns contains all the column names of Table for convenient usage.
}

// AppColumns defines and stores column names for table app.
type AppColumns struct {
	Id       string //
	Appid    string //
	ApiKey   string //
	SecKey   string //
	Owner    string //
	Name     string // 应用名
	Status   string // 1启用 0 禁用
	OwnerUid string // 拥有者uid
	Created  string //
	Updated  string //
}

//  appColumns holds the columns for table app.
var appColumns = AppColumns{
	Id:       "id",
	Appid:    "appid",
	ApiKey:   "api_key",
	SecKey:   "sec_key",
	Owner:    "owner",
	Name:     "name",
	Status:   "status",
	OwnerUid: "owner_uid",
	Created:  "created",
	Updated:  "updated",
}

// NewAppDao creates and returns a new DAO object for table data access.
func NewAppDao() *AppDao {
	return &AppDao{
		group:   "default",
		table:   "app",
		columns: appColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppDao) Columns() AppColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
