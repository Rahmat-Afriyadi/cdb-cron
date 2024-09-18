package repository

import (
	"cron/config"
	"cron/entity"
)

func NewMstTelp1() []entity.MstTelp1 {
	localDb2 := config.NewLocal2DB()
	datas := []entity.MstTelp1{}

	localDb2.Table("mst_telp1").Find(&datas)

	return datas

}

func NewMstHp1() []entity.MstHp1 {
	localDb2 := config.NewLocal2DB()
	datas := []entity.MstHp1{}

	localDb2.Table("mst_hp1").Find(&datas)

	return datas

}
