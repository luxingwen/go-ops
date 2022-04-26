// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Plugin is the golang structure for table plugin.
type Plugin struct {
	Id          int         `json:"id"          ` //
	Uuid        string      `json:"uuid"        ` // 插件uuid
	Name        string      `json:"name"        ` // 插件名
	PackageName string      `json:"packageName" ` // 包名
	Os          string      `json:"os"          ` // 操作系统
	Arch        string      `json:"arch"        ` // 架构
	Md5         string      `json:"md5"         ` // 包md5名称
	Creater     string      `json:"creater"     ` // 创建人
	Updater     string      `json:"updater"     ` // 更新人
	Created     *gtime.Time `json:"created"     ` //
	Updated     *gtime.Time `json:"updated"     ` //
}
