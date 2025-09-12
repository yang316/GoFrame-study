package user

import (
	"gf_study/internal/model/entity"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" summary:"登录" tags:"User" `
	Username string `v:"required|length:6,16#请输入账号|账号长度为:min到:max位" dc:"用户账号" json:"username"`
	Password string `v:"required|length:6,16#请输入密码|密码长度为:min到:max位" dc:"用户密码" json:"password"`
}

type LoginRes struct {
	Info *entity.GfUsers `json:"info"`
}

type RegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" summary:"注册" tags:"User" `
	Username string `v:"required|length:6,16#请输入账号|账号长度为:min到:max位" dc:"用户账号" json:"username"`
	Password string `v:"required|length:6,16#请输入密码|密码长度为:min到:max位" dc:"用户密码" json:"password"`
	Nickname string `v:"required|length:2,10#请输入昵称|昵称长度为:min到:max位" dc:"用户昵称" json:"nickname"`
}

type RegisterRes struct {
	Info *entity.GfUsers `json:"info"`
}

type ThirdLoginReq struct {
	g.Meta   `path:"/user/thirdLogin" method:"post" summary:"第三方登录" tags:"User"`
	Platform string `v:"required" dc:"第三方平台" json:"platform"`
	Code     string `v:"required" dc:"第三方平台Code" json:"code"`
}

type ThirdLoginRes struct {
	Info *entity.GfUsers `json:"info"`
}

type ProfileReq struct {
	g.Meta   `path:"/user/profile" method:"put" summary:"修改个人信息" tags:"User"`
	Nickname string `v:"required|length:2,10" dc:"昵称" json:"nickname"`
	//Password string `v:"required|length:6,16" dc:"密码" json:"password"`
	Avatar string `v:"required"  dc:"头像" json:"avatar"`
}
type ProfileRes struct {
	Result bool `dc:"修改结果" json:"result"`
}

type UploadReq struct {
	g.Meta `path:"/upload" method:"post" mime:"multipart/form-data" summary:"上传文件" tags:"uploads"`
	File   *ghttp.UploadFile `v:"required" p:"file" type:"file" dc:"上传文件"`
}
type UploadRes struct {
	Url      string `json:"url"`
	FileName string `json:"fileName"`
}
