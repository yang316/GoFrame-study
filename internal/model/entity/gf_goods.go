// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GfGoods is the golang structure for table gf_goods.
type GfGoods struct {
	Id            uint        `json:"id"            orm:"id"            description:""`                   //
	GoodsName     string      `json:"goodsName"     orm:"goodsName"     description:"商品名称"`             // 商品名称
	GoodsDesc     string      `json:"goodsDesc"     orm:"goodsDesc"     description:"商品描述"`             // 商品描述
	GoodsPrice    float32     `json:"goodsPrice"    orm:"goodsPrice"    description:"商品价格"`             // 商品价格
	GoodsStock    uint64         `json:"goodsStock"    orm:"goodsStock"    description:"商品库存"`             // 商品库存
	GoodsType     byte        `json:"goodsType"     orm:"goodsType"     description:"商品类型"`             // 商品类型
	CreateTime    *gtime.Time `json:"createTime"    orm:"createTime"    description:"创建时间"`             // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"updateTime"    description:"更新时间"`             // 更新时间
}
