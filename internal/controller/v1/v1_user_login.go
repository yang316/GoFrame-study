package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"gf_study/api/v1/user"
	"gf_study/internal/dao"
	"gf_study/internal/model/do"
	"gf_study/internal/model/entity"
)

func (c *ControllerUser) Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	var userData do.GfUsers

	//查看用户名是否注册
	err = dao.GfUsers.Ctx(ctx).Where("username", req.Username).Scan(&userData)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "用户信息有误或未注册")
	}
	if gconv.Uint(userData.Id) == 0 {
		return nil, gerror.New("用户名未注册")
	}
	//查看密码是否正确
	// 解密查看密码
	checkPassword, err := DecryptPassword(gconv.String(userData.Password), req.Password)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "密码解密失败")
	}
	if req.Password != checkPassword {
		return nil, gerror.New("密码错误")
	}
	userEntity := &entity.GfUsers{
		Id:       gconv.Uint(userData.Id),
		Username: gconv.String(userData.Username),
		Nickname: gconv.String(userData.Nickname),
		
	}

	// 构建返回结果（注意：不要返回原始密码）
	return &user.LoginRes{
		Info: userEntity,
	}, nil
}
