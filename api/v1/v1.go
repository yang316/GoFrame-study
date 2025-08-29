// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"gf_study/api/v1/user"
)

type IV1User interface {
	Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error)
	Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error)
}
