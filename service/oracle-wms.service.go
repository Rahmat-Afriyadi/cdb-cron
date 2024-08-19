package service

import (
	"cron/config"
	"fmt"
)

type Result struct {
	NO_MSN   string
	NM_CUS   string
	NO_HP1   string
	NO_TELP1 string
}

func OracleWMS() {
	oracleDB := config.NewOracleDB()

	var data []Result
	oracleDB.Raw("SELECT NO_MSN, NM_CUS, NO_HP1, NO_TELP1 FROM MAINDEALER.MOHONFAKTUR_WKM_TDKVALID where NO_TELP1 is null").Scan(&data)
	fmt.Println("ini result ", data[0].NO_HP1 != "")

}
