// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Vm is the golang structure for table vm.
type Vm struct {
	Id          int         `json:"id"          ` //
	Uuid        string      `json:"uuid"        ` // 主机uuid
	Name        string      `json:"name"        ` //
	Hostname    string      `json:"hostname"    ` //
	Az          string      `json:"az"          ` // 可用区
	OsType      string      `json:"osType"      ` // 操作系统类型
	OsInfo      string      `json:"osInfo"      ` // 操作系统信息
	Hosttype    string      `json:"hosttype"    ` // 主机类型
	Networktype string      `json:"networktype" ` // 网络类型
	PrivateIp   string      `json:"privateIp"   ` //
	PublicIp    string      `json:"publicIp"    ` //
	Created     *gtime.Time `json:"created"     ` //
	Updated     *gtime.Time `json:"updated"     ` //
	Creater     string      `json:"creater"     ` // 创建人
}
