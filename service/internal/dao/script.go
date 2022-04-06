// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"osp/service/internal/dao/internal"
)

// internalScriptDao is internal type for wrapping internal DAO implements.
type internalScriptDao = *internal.ScriptDao

// scriptDao is the data access object for table script.
// You can define custom methods on it to extend its functionality as you wish.
type scriptDao struct {
	internalScriptDao
}

var (
	// Script is globally public accessible object for table script operations.
	Script = scriptDao{
		internal.NewScriptDao(),
	}
)

// Fill with you ideas below.
