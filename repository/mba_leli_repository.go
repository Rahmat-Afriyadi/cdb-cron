package repository

import (
	"cron/config"
	"cron/entity"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

func SaveTrWmsFaktur2(hdFakturs []entity.MbaLeliFull) {
	localDb := config.NewLocalDB()
	if len(hdFakturs) > 0 {
		tx := localDb.Begin()
		batchSize := 500
		for i := 0; i < len(hdFakturs); i += batchSize {
			end := i + batchSize
			if end > len(hdFakturs) {
				end = len(hdFakturs)
			}

			batch := hdFakturs[i:end]
			if err := tx.Table("tr_wms_faktur2").Save(&batch).Error; err != nil {
				tx.Rollback()
				fmt.Println("Error:", err)
				return
			}
		}
		tx.Commit()
	}
}
func InsertDataMbaLeli(mohonFaktur []map[string]interface{}) {
	// cdbConf := config.NewCDBWebDB()
	localDb2 := config.NewLocal2DB()
	deleteTx := localDb2.Begin()
	result := deleteTx.Exec("delete from data_mba_leli")
	if result.Error != nil {
		deleteTx.Rollback()
		fmt.Println("Error:", result.Error)
		return
	}
	deleteTx.Commit()

	hdFakturs := []entity.MbaLeliFull{}

	for _, value := range mohonFaktur {
		hdFaktur := entity.MbaLeliFull{
			NoMsn:        value["NO_MSN"].(string),
			NoRgk:        value["NO_RGK"].(string),
			NmMtr:        value["NM_MTR"].(string),
			KdDlr:        value["NO_DLRP"].(string),
			NmDlr:        value["NM_DLR"].(string),
			NmCustomer11: value["NM1_MOHON"].(string),
			JnsKlm1:      value["JNS_KLM"].(string),
			NoHp1:        value["NO_HP"].(string),
			Alamat11:     value["AL1_MOHON"].(string),
			Kel1:         value["KEL_MOHON"].(string),
			Kec1:         value["KEC_MOHON"].(string),
			Kota1:        value["KOTA_MOHON"].(string),
			Kodepos1:     value["KODE_POS"].(string),
			TglMohon:     value["TGL_MOHON"].(string),
			KdUser:       "99",
			KdCustomer:   "1",
			StsCetak3:    0,
			KdClient:     "1",
			// JnsKlm2:             value[""].(string),
			// StsKirim:            value[""].(string),
			// StsCetak:            value[""].(string),
			// StsSource:           value[""].(string),
			// StsAsuransi:         value[""].(string),
			// StsTertanggung:      value[""].(string),
			// StsKawin:            value[""].(string),
			// StsLengkap:          value[""].(string),
			// Agama2:              value[""].(string),
			// KodeDidik2:          value[""].(string),
			// KodeKerja2:          value[""].(string),
			// KeluarBln2:          value[""].(string),
			// StsSms:              value[""].(string),
			// StsCetak3:           value[""].(string),
			TujuanPakai1: value["TUJU_PAK1"].(string),
			Email1:       value["E_MAIL"].(string),
			Agama1:       value["KODE_AGAMA"].(string),
			SediaDihub:   value["SEDIA_DIHUB"].(string),
			JnsMotor:     value["JNS_MOTOR"].(string),
			SmDibeli:     value["SM_DIBELI"].(string),
			PropMohon:    value["PROP_MOHON"].(string),
			Hobby3:       value["HOBI"].(string),
			NoKk:         value["NO_KK"].(string),
			TglFaktur:    value["TGL_FAKTUR"].(string),
			AktifJual:    value["AKTIF_JUAL"].(string),
		}
		if value["NM2_MOHON"] != nil {
			hdFaktur.NmCustomer12 = value["NM2_MOHON"].(string)
		}
		if value["TGL_LAHIR"] != nil {
			tglLahir1, err1 := time.Parse("2006-01-02", value["TGL_LAHIR"].(string))
			if err1 == nil {
				hdFaktur.TglLahir1 = &tglLahir1
			}
		}
		if value["NO_TELEPON"] != nil {
			hdFaktur.NoTelp1 = value["NO_TELEPON"].(string)
		}
		if value["NO_LEAS"] != nil {
			hdFaktur.NoLeas = value["NO_LEAS"].(string)
		}
		if value["JML_ANGSURAN"] != nil {
			hdFaktur.Angsuran = value["JML_ANGSURAN"].(string)
		}
		if value["NO_HP"] != nil {
			hdFaktur.NoHp1 = value["NO_HP"].(string)
		}
		if value["RT_MOHON"] != nil {
			hdFaktur.Rt1 = value["RT_MOHON"].(string)
		}
		if value["RW_MOHON"] != nil {
			hdFaktur.Rw1 = value["RW_MOHON"].(string)
		}
		if value["AKUNFACEBOOK"] != nil {
			hdFaktur.SmFacebook1 = value["AKUNFACEBOOK"].(string)
		}
		if value["AKUNINSTAGRAM"] != nil {
			hdFaktur.SmInstagram1 = value["AKUNINSTAGRAM"].(string)
		}
		if value["AKUNTWITTER"] != nil {
			hdFaktur.SmTwitter1 = value["AKUNTWITTER"].(string)
		}
		if value["AKUNYOUTUBE"] != nil {
			hdFaktur.SmYoutube1 = value["AKUNYOUTUBE"].(string)
		}
		if value["AL1_SRT"] != nil {
			hdFaktur.AlamatSrt11 = value["AL1_SRT"].(string)
		}
		if value["AL2_MOHON"] != nil {
			hdFaktur.Alamat12 = value["AL2_MOHON"].(string)
		}
		if value["AL2_SRT"] != nil {
			hdFaktur.AlamatSrt12 = value["AL2_SRT"].(string)
		}
		if value["BDN_USAHA"] != nil {
			hdFaktur.BdnUsaha = value["BDN_USAHA"].(string)
		}
		if value["POS_SRT"] != nil {
			hdFaktur.KodeposSrt1 = value["POS_SRT"].(string)
		}
		if value["KODE_KERJA"] != nil {
			hdFaktur.KodeKerja = value["KODE_KERJA"].(string)
		}
		if value["KELUAR_BLN"] != nil {
			hdFaktur.KeluarBln = value["KELUAR_BLN"].(string)
		}
		if value["KODE_DIDIK"] != nil {
			hdFaktur.KodeDidik = value["KODE_DIDIK"].(string)
		}
		if value["MOTOR_HIR"] != nil {
			hdFaktur.MotorHir = value["MOTOR_HIR"].(string)
		}
		if value["JNS_BELI"] != nil {
			hdFaktur.JnsBeli = value["JNS_BELI"].(string)
		}
		if value["JNS_JUAL"] != nil {
			hdFaktur.JnsJual = value["JNS_JUAL"].(string)
		}
		if value["ID_SALES_ASLI"] != nil {
			hdFaktur.IdSales = value["ID_SALES_ASLI"].(string)
		}
		if value["NM_SALES_ASLI"] != nil {
			hdFaktur.NmSales1 = value["NM_SALES_ASLI"].(string)
		}
		if value["NO_MTR"] != nil {
			hdFaktur.NoMtr = value["NO_MTR"].(string)
		}
		if value["KEC_SRT"] != nil {
			hdFaktur.KecSrt1 = value["KEC_SRT"].(string)
		}
		if value["KOTA_SRT"] != nil {
			hdFaktur.KotaSrt1 = value["KOTA_SRT"].(string)
		}
		if value["PRSH_NAMA"] != nil {
			hdFaktur.KerjaDiFkt = value["PRSH_NAMA"].(string)
		}
		if value["PRSH_ALAMAT"] != nil {
			hdFaktur.AlamatKtrFkt = value["PRSH_ALAMAT"].(string)
		}
		if value["PRSH_KEC"] != nil {
			hdFaktur.KecKtrFkt = value["PRSH_KEC"].(string)
		}
		if value["PRSH_KOTA"] != nil {
			hdFaktur.KotaKtrFkt = value["PRSH_KOTA"].(string)
		}
		if value["PRSH_PROP"] != nil {
			hdFaktur.PropKtrFkt = value["PRSH_PROP"].(string)
		}
		if value["NO_NPWP"] != nil {
			hdFaktur.NoNpwp = value["NO_NPWP"].(string)
		}
		if value["NO_KTPNPWP"] != nil {
			if len(value["NO_KTPNPWP"].(string)) > 16 {
				hdFaktur.NoKtpNpwp = value["NO_KTPNPWP"].(string)[:16]
			} else {
				hdFaktur.NoKtpNpwp = value["NO_KTPNPWP"].(string)
			}
		}
		if value["DP"] != nil {
			if floatVal, ok := value["DP"].(float64); ok {
				hdFaktur.DpMtr = floatVal
			} else if stringVal, ok := value["DP"].(string); ok {
				floatVal, _ := strconv.ParseFloat(stringVal, 64)
				hdFaktur.DpMtr = floatVal

			}
		}
		if value["CICILAN"] != nil {
			hdFaktur.CicilanMtr = value["CICILAN"].(float64)
		}
		if value["STS_RUMAH"] != nil {
			hdFaktur.StsKepemilikanRumah = value["STS_RUMAH"].(string)
		}
		if value["STS_HP"] != nil {
			hdFaktur.StsHpPraPasca = value["STS_HP"].(string)
		}
		if value["NAMA_PNJJWB"] != nil {
			hdFaktur.PicPerusahaan = value["NAMA_PNJJWB"].(string)
		}
		hdFaktur.TglAkhirTenor = nil

		hdFakturs = append(hdFakturs, hdFaktur)
	}
	if len(hdFakturs) > 0 {
		tx := localDb2.Begin()
		batchSize := 1000
		for i := 0; i < len(hdFakturs); i += batchSize {
			end := i + batchSize
			if end > len(hdFakturs) {
				end = len(hdFakturs)
			}

			batch := hdFakturs[i:end]
			if err := tx.Table("data_mba_leli").Create(&batch).Error; err != nil {
				tx.Rollback()
				fmt.Println("Error:", err)
				return
			}
		}
		tx.Commit()
	}
	fmt.Println("Done yaa...")

}

func GroupSetOfCheck() {
	localDb2 := config.NewLocal2DB()

	var result0 *sql.Rows
	sqlStatement := "SELECT no_telp1, count(*) FROM data_mba_leli group by no_telp1 having count(*) >1"
	result0, err := localDb2.Raw(sqlStatement).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer result0.Close()
	for result0.Next() {
		var no_telp1 string
		var count int
		if err := result0.Scan(&no_telp1, &count); err != nil {
			log.Fatal(err)
		}
		testSave := localDb2.Table("mst_telp1").Save(&entity.MstTelp1{NoTelp1: no_telp1})
		if testSave.Error != nil {
			fmt.Println("disini errornya yaa ", testSave.Error)
		}
	}

	sqlStatement = "SELECT no_hp1, count(*) FROM data_mba_leli group by no_hp1 having count(*) >1"
	result0, err = localDb2.Raw(sqlStatement).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer result0.Close()
	for result0.Next() {
		var no_hp1 string
		var count int
		if err := result0.Scan(&no_hp1, &count); err != nil {
			log.Fatal(err)
		}
		testSave := localDb2.Table("mst_hp1").Save(&entity.MstHp1{NoHp1: no_hp1})
		if testSave.Error != nil {
			fmt.Println("disini errornya yaa ", testSave.Error)
		}
	}

	queries := []struct {
		Query  string
		Params []interface{}
	}{
		{"UPDATE data_mba_leli set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran month) where tgl_akhir_tenor is null and angsuran not in ('','0','N')", []interface{}{}},
		{"UPDATE data_mba_leli SET kode_kerja2 = ? WHERE kode_kerja2 IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET kode_kerja = ? WHERE kode_kerja IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET kec_ktr_fkt = ? WHERE kec_ktr_fkt IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET hobby2 = ? WHERE hobby2 IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET kec_ktr = ? WHERE kec_ktr IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET tujuan_pakai2 = ? WHERE tujuan_pakai2 IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET no_leas2 = ? WHERE no_leas2 IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET id_merk_mtr_seblm = ? WHERE id_merk_mtr_seblm IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET id_jns_mtr_seblm = ? WHERE id_jns_mtr_seblm IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET hobby3 = ? WHERE hobby3 IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET ket_no_telp1 = ? WHERE ket_no_telp1 IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET ket_no_telp2 = ? WHERE ket_no_telp2 IS NULL", []interface{}{""}},
		{"UPDATE data_mba_leli SET ket_no_hp1 = ? WHERE ket_no_hp1 IS NULL", []interface{}{""}},
	}
	// Execute each query
	tx := localDb2.Begin()
	for _, q := range queries {
		result := tx.Exec(q.Query, q.Params...)
		if result.Error != nil {
			tx.Rollback()
			fmt.Println("Error:", result.Error)
			return
		}
	}
	tx.Commit()

	queryCTelp := "update db_wkm2.data_mba_leli set ket_no_telp1='C' where ket_no_telp1 = '' and no_telp1<>'' and no_telp1 in ("
	mstTelp := NewMstTelp1()
	for _, value := range mstTelp {
		queryCTelp += "'" + value.NoTelp1 + "',"
	}
	queryCTelp = queryCTelp[:len(queryCTelp)-1]
	queryCTelp += ");"

	tx1 := localDb2.Begin()
	result := tx1.Exec(queryCTelp)
	if result.Error != nil {
		tx1.Rollback()
		fmt.Println("Error:", result.Error)
		return
	}
	tx1.Commit()

	queryCHp := "update db_wkm2.data_mba_leli set ket_no_hp1='C' where ket_no_hp1 = '' and no_hp1<>'' and no_hp1 in ("
	mstHp := NewMstHp1()
	for _, value := range mstHp {
		queryCHp += "'" + value.NoHp1 + "',"
	}
	queryCHp = queryCHp[:len(queryCHp)-1]
	queryCHp += ");"
	tx2 := localDb2.Begin()
	result = tx2.Exec(queryCHp)
	if result.Error != nil {
		tx2.Rollback()
		fmt.Println("Error:", result.Error)
		return
	}
	tx2.Commit()

	numbers3X := entity.GetX3HP()
	query3X := "update db_wkm2.data_mba_leli set ket_no_hp1='3X' where no_hp1 like ?;"
	tx3 := localDb2.Begin()
	for _, value := range numbers3X {
		result := tx3.Exec(query3X, "%"+value)
		if result.Error != nil {
			tx3.Rollback()
			fmt.Println("Error:", result.Error)
			return
		}
	}
	tx3.Commit()

	query3X = "update db_wkm2.data_mba_leli set ket_no_telp1='3X' where no_telp1 like ?;"
	tx4 := localDb2.Begin()
	for _, value := range numbers3X {
		result := tx4.Exec(query3X, "%"+value)
		if result.Error != nil {
			tx4.Rollback()
			fmt.Println("Error:", result.Error)
			return
		}
	}
	tx4.Commit()

	query3X = "update db_wkm2.data_mba_leli set ket_no_hp1='3X' where no_hp1 like ?;"
	tx5 := localDb2.Begin()
	for _, value := range numbers3X {
		result := tx5.Exec(query3X, "%"+value)
		if result.Error != nil {
			tx5.Rollback()
			fmt.Println("Error:", result.Error)
			return
		}
	}
	tx5.Commit()

	query4X := entity.Get4xTelp()
	tx6 := localDb2.Begin()
	result = tx6.Exec(query4X)
	if result.Error != nil {
		tx6.Rollback()
		fmt.Println("Error:", result.Error)
		return
	}
	tx6.Commit()

	query4X = entity.Get4XHp()
	tx7 := localDb2.Begin()
	result = tx7.Exec(query4X)
	if result.Error != nil {
		tx7.Rollback()
		fmt.Println("Error:", result.Error)
		return
	}
	tx7.Commit()

	query5X := entity.Get5xTelp()
	tx8 := localDb2.Begin()
	result = tx8.Exec(query5X)
	if result.Error != nil {
		tx8.Rollback()
		fmt.Println("Error:", result.Error)
		return
	}
	tx8.Commit()

	query5X = entity.Get5xHp()
	tx9 := localDb2.Begin()
	result = tx9.Exec(query5X)
	if result.Error != nil {
		tx9.Rollback()
		fmt.Println("Error:", result.Error)
		return
	}
	tx9.Commit()

}
