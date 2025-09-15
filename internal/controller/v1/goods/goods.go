package goods

import (
	"context"
	v1 "gf_study/api/v1"
	"gf_study/api/v1/goods"
	"gf_study/internal/dao"
	"gf_study/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// ControllerGoods is the controller for goods-related operations
type ControllerGoods struct{}

func NewGoods() v1.IV1Goods {
	return &ControllerGoods{}
}

func (c *ControllerGoods) CreateGoods(ctx context.Context, req *goods.CreateGoodsReq) (res *goods.CreateGoodsRes, err error) {
	var goodsData do.GfGoods

	goodsData.GoodsName = req.GoodsName
	goodsData.GoodsDesc = req.GoodsDesc
	goodsData.GoodsPrice = gconv.Float32(req.GoodsPrice)
	goodsData.GoodsStock = gconv.Uint64(req.GoodsStock)
	goodsData.GoodsType = gconv.Byte(req.GoodsType)
	goodsData.CreateTime = gtime.Now()
	goodsData.UpdateTime = gtime.Now()
	// 插入数据
	_, err = dao.GfGoods.Ctx(ctx).Data(goodsData).Insert()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "创建商品失败")
	}

	return &goods.CreateGoodsRes{
		GoodsId: gconv.Int(goodsData.Id),
	}, nil

}

