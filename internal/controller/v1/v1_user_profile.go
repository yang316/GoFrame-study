package v1

import (
	"context"
	"gf_study/api/v1/user"
	"gf_study/internal/dao"
	"gf_study/internal/model/do"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerUser) Profile(ctx context.Context, req *user.ProfileReq) (res *user.ProfileRes, err error) {

	result, err := dao.GfUsers.Ctx(ctx).Data(do.GfUsers{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	}).Where("nickname", req.Nickname).Update()
	if err != nil {
		return nil, err
	}

	// 2. 检查 result 是否为 nil
	if result == nil {
		return nil, gerror.New("更新操作未返回有效结果")
	}

	// 3. 获取影响行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	// 4. 检查是否成功更新
	if rowsAffected < 1 {
		return nil, gerror.New("修改失败")
	}

	// 5. 返回成功结果
	return &user.ProfileRes{
		Result: true,
	}, nil
}
