// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GfGoods is the golang structure of table gf_goods for DAO operations like Where/Data.
type GfGoods struct {
	g.Meta        `orm:"table:gf_goods, do:true"`
	Id            interface{} //
	GoodsName     interface{} // 商品名称
	GoodsDesc     interface{} // 商品描述
	GoodsPrice    interface{} // 商品价格
	GoodsStock    interface{} // 商品库存
	GoodsType     interface{} // 商品类型
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 更新时间
}
