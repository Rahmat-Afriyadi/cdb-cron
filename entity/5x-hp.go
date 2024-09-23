package entity

func Get5xTelp() string {
	return "update db_wkm2.data_mba_leli set ket_no_telp1='5' where no_telp1 in ('0216012044','6012044','08001881700')"
}

func Get5xHp() string {
	return "update db_wkm2.data_mba_leli set ket_no_Hp1='5' where no_hp1 in ('0216012044','6012044','08001881700')"
}
