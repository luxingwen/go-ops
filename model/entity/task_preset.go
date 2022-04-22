// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskPreset is the golang structure for table task_preset.
type TaskPreset struct {
	Id      int         `json:"id"      ` //
	Uuid    string      `json:"uuid"    ` //
	Name    string      `json:"name"    ` // 任务名
	Type    string      `json:"type"    ` // 任务类型
	Content string      `json:"content" ` // 任务内容
	Peers   string      `json:"peers"   ` // 主机列表
	Creater string      `json:"creater" ` // 创建人
	Created *gtime.Time `json:"created" ` // 创建时间
	Updated *gtime.Time `json:"updated" ` // 更新时间
}
