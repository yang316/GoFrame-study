// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"gf_study/api/v1/goods"
	"gf_study/api/v1/user"
)

type IV1User interface {
	Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error)
	Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error)
	ThirdLogin(ctx context.Context,req *user.ThirdLoginReq)(res *user.ThirdLoginRes,err error)
	Profile(ctx context.Context,req *user.ProfileReq)(res *user.ProfileRes,err error)
	Upload(ctx context.Context,req *user.UploadReq)(res *user.UploadRes,err error)
}

type IV1Goods interface {
	CreateGoods(ctx context.Context, req *goods.CreateGoodsReq) (res *goods.CreateGoodsRes, err error)
}

