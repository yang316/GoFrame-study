package v1

import (
	"context"
	"gf_study/internal/dao"
	"gf_study/internal/model/do"
	"gf_study/internal/model/entity"

	"golang.org/x/crypto/bcrypt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"

	"gf_study/api/v1/user"
)

func (c *ControllerUser) Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error) {

	// 检查用户是否已存在
	isExist, err := dao.GfUsers.Ctx(ctx).Where("username", req.Username).Exist()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "检查用户是否存在失败")
	}
	if isExist {
		return nil, gerror.Newf("用户名 %s 已存在", req.Username)
	}

	// 密码加密处理（实际项目中必须添加）
	encryptedPassword, err := EncryptPassword(req.Password)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "密码加密失败")
	}

	// 插入新用户
	result, err := dao.GfUsers.Ctx(ctx).Insert(do.GfUsers{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: encryptedPassword, // 使用加密后的密码
		RegistTime: gtime.Now(),
		LastLoginTime: gtime.Now(),
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "用户注册失败")
	}

	// 获取插入的用户ID
	userId, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "获取用户ID失败")
	}

	// 查询刚创建的用户信息
	userEntity := &entity.GfUsers{
		Id:       uint(userId),
		Username: req.Username,
		Nickname: req.Nickname,
	}


	// 构建返回结果（注意：不要返回原始密码）
	return &user.RegisterRes{
		Info: userEntity,
	}, nil
}
// 加密
func EncryptPassword(password string) (string, error) {
	// 使用 bcrypt 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
// 解密密码
func DecryptPassword(encryptedPassword string, password string) (string, error) {

	// 使用 bcrypt 解密密码
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		return "", err
	}
	return password, nil
}