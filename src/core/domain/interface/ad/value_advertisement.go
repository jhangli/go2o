/**
 * Copyright 2015 @ S1N1 Team.
 * name : value_advertisement
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package ad

type ValueAdvertisement struct {
	// 编号
	Id int `db:"id" auto:"yes" pk:"yes"`

	PartnerId int `db:"partner_id"`
	// 名称
	Name string `db:"name"`

	// 是否内部
	IsInternal int `db:"is_internal"`

	// 广告类型
	Type int `db:"type_id"`

	// 是否启用
	Enabled int `db:"enabled"`

	// 修改时间
	UpdateTime int64 `db:"update_time"`
}
