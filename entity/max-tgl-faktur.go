package entity

import "time"

type MaxTglFaktur struct {
	TglFaktur time.Time `gorm:"column:tgl_faktur"`
}
