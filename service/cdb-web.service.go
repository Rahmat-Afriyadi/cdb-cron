package service

import (
	"cron/config"
	"time"
)

type HDFaktur2 struct {
	NoMsn     string
	NoTelp    string
	NoHp      string
	Email     string
	KodeKerja string
}

func getLastDayOfMonth(year int, month time.Month) time.Time {
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := nextMonth.Add(-time.Hour * 24)
	return lastDayOfMonth
}

func UpdateMohonFakturWKMTDKValid() {
	lastDay := getLastDayOfMonth(time.Now().Year(), time.Now().AddDate(0, -1, 0).Month())

	oracleDB := config.NewOracleDB()
	cdbDB := config.NewCDBWebDB()

	var datas []HDFaktur2
	cdbDB.Raw("SELECT no_msn NoMsn, no_telp11 NoTelp, no_hp11 NoHp, email2 Email, kode_kerja2 KodeKerja from hd_faktur2 where tgl_faktur >= ? and tgl_faktur <= ? and modified != '00000000' ", lastDay.Format("2006-01")+"-01", lastDay.Format("2006-01-02")).Scan(&datas)
	for _, v := range datas {
		if v.Email != "" {
			oracleDB.Table("MAINDEALER.MOHONFAKTUR_WKM_TDKVALID").Where("NO_MSN = ?", v.NoMsn).Update("EMAIL_CUS", v.Email).Update("TGL_WKM", time.Now())
		}
		if v.NoTelp != "" && v.NoTelp != "N" {
			oracleDB.Table("MAINDEALER.MOHONFAKTUR_WKM_TDKVALID").Where("NO_MSN = ?", v.NoMsn).Update("NO_TELP1", v.NoTelp).Update("TGL_WKM", time.Now())
		}
		if v.NoHp != "" && v.NoHp != "N" {
			oracleDB.Table("MAINDEALER.MOHONFAKTUR_WKM_TDKVALID").Where("NO_MSN = ?", v.NoMsn).Update("NO_HP1", v.NoHp).Update("TGL_WKM", time.Now())
		}
		if v.KodeKerja != "" && v.KodeKerja != "0" {
			oracleDB.Table("MAINDEALER.MOHONFAKTUR_WKM_TDKVALID").Where("NO_MSN = ?", v.NoMsn).Update("KERJA_CUS", v.KodeKerja).Update("TGL_WKM", time.Now())
		}
	}
}

// {JM04E1872694 08979620183 08979620183 ardidwicahyono@gmail.com 2b}
