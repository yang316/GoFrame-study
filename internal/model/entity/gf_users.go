// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GfUsers is the golang structure for table gf_users.
type GfUsers struct {
	Id            uint        `json:"id"            orm:"id"            description:""`             //
	Username      string      `json:"username"      orm:"username"      description:"账号"`           // 账号
	Nickname      string      `json:"nickname"      orm:"nickname"      description:"昵称"`           // 昵称
	Password      string      `json:"password"      orm:"password"      description:"密码"`           // 密码
	Status        uint        `json:"status"        orm:"status"        description:"状态:0=正常,1=禁用"` // 状态:0=正常,1=禁用
	RegistTime    *gtime.Time `json:"registTime"    orm:"registTime"    description:"注册时间"`         // 注册时间
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"lastLoginTime" description:"上次登录时间"`       // 上次登录时间
	OpenId        string      `json:"openId"        orm:"openId"        description:"第三方openId"`    // 第三方openId
}
