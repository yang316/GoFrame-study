package goods

import "github.com/gogf/gf/v2/frame/g"

type CreateGoodsReq struct {
	g.Meta    `path:"/goods/create" method:"post" summary:"创建商品"`
	GoodsName string `json:"goodsName" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsDesc string `json:"goodsDesc" v:"required#商品描述不能为空" dc:"商品描述"`
	GoodsPrice int `json:"goodsPrice" v:"required#商品价格不能为空" dc:"商品价格"`
	GoodsStock int `json:"goodsStock" v:"required#商品库存不能为空" dc:"商品库存"`
	GoodsType int `json:"goodsType" v:"required#商品类型不能为空" dc:"商品类型"`
}

type CreateGoodsRes struct {
	GoodsId int `json:"goodsId"`
}