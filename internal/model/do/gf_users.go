// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GfUsers is the golang structure of table gf_users for DAO operations like Where/Data.
type GfUsers struct {
	g.Meta        `orm:"table:gf_users, do:true"`
	Id            interface{} //
	Username      interface{} // 账号
	Nickname      interface{} // 昵称
	Password      interface{} // 密码
	Status        interface{} // 状态:0=正常,1=禁用
	RegistTime    *gtime.Time // 注册时间
	LastLoginTime *gtime.Time // 上次登录时间
	OpenId        interface{} // 第三方openId
}
