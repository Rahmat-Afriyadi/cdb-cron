package repository

import (
	"cron/config"
	"cron/entity"
	"cron/service"
	"log"
	"strings"
)

func UpdateData(mohonFaktur []map[string]interface{}) {
	cdbConf := config.NewCDBWebDB()
	masterKode := service.GetDataMasterKode()

	hdFakturs := []entity.HDFaktur2Full{}
	for _, value := range mohonFaktur {

		stsValid := "TIDAK VALID"
		ketNoTelepon := "TIDAK VALID"
		ketNoHp := "TIDAK VALID"
		vEmail := "TIDAK VALID"
		if value["NO_TELEPON"] != nil {
			if strings.HasPrefix(value["NO_TELEPON"].(string), "08") || strings.HasPrefix(value["NO_TELEPON"].(string), "628") {
				ketNoTelepon = "VALID"
			}
		}
		if value["NO_HP"] != nil {
			if strings.HasPrefix(value["NO_HP"].(string), "08") || strings.HasPrefix(value["NO_HP"].(string), "628") {
				ketNoHp = "VALID"
			}
		}
		if value["E_MAIL"] != nil {
			if len(value["E_MAIL"].(string)) > 17 {
				vEmail = "VALID"
			}
		}
		if ketNoTelepon == "VALID" && ketNoHp == "VALID" && vEmail == "VALID" {
			stsValid = "VALID"
		}

		hdFaktur := entity.HDFaktur2Full{
			NoMsn:        value["NO_MSN"].(string),
			NmCustomer11: value["NM1_MOHON"].(string),
			TglLahir2:    value["TGL_LAHIR"].(string),
			KetNoTelp1:   ketNoTelepon,
			KetNoHp1:     ketNoHp,
			Agama2:       value["KODE_AGAMA"].(string),
			KAgama2:      masterKode["agama"][value["KODE_AGAMA"].(string)],
			Email1:       value["E_MAIL"].(string),
			VEmail:       vEmail,
			KodeKerja:    value["KODE_KERJA"].(string),
			KKodeKerja1:  masterKode["kerja"][value["KODE_KERJA"].(string)],
			Alamat11:     value["AL1_MOHON"].(string),
			TglFaktur:    value["TGL_FAKTUR"].(string),
			KdDlr:        value["NO_DLRP"].(string),
			NmSales1:     value["NM_SALES"].(string),
			StsValid:     stsValid,
			NmMtr:        value["NM_MTR"].(string),
			JnsJual:      value["JNS_JUAL"].(string),
			KJnsJual:     masterKode["jual"][value["JNS_JUAL"].(string)],
			JnsBeli:      value["JNS_BELI"].(string),
			KJnsBeli:     masterKode["beli"][value["JNS_BELI"].(string)],
			NoLeas:       value["NO_LEAS"].(string),
			KNoLeas:      masterKode["leasing"][value["NO_LEAS"].(string)],
			NoRgk:        value["NO_RGK"].(string),
			NoKtpnpwp:    value["NO_KTPNPWP"].(string),
			Kel1:         value["KEL_MOHON"].(string),
			Kec1:         value["KEC_MOHON"].(string),
			Kodepos1:     value["KODE_POS"].(string),
			IdSales:      value["ID_SALES_ASLI"].(string),
			NoKk:         value["NO_KK"].(string),
			KodeDidik:    value["KODE_DIDIK"].(string),
			KKodeDidik1:  masterKode["pendidikan"][value["KODE_DIDIK"].(string)],
			KeluarBln:    value["KELUAR_BLN"].(string),
			KKeluarBln1:  masterKode["pengeluaran"][value["KELUAR_BLN"].(string)],
			TglMohon:     value["TGL_MOHON"].(string),
		}

		if value["NO_TELEPON"] != nil {
			hdFaktur.NoTelp1 = value["NO_TELEPON"].(string)
		}
		if value["NO_HP"] != nil {
			hdFaktur.NoHp1 = value["NO_HP"].(string)
		}
		if value["AKUNFACEBOOK"] != nil {
			hdFaktur.SmFacebook1 = value["AKUNFACEBOOK"].(string)
		}
		if value["AKUNTWITTER"] != nil {
			hdFaktur.SmTwitter1 = value["AKUNTWITTER"].(string)
		}
		if value["AKUNINSTAGRAM"] != nil {
			hdFaktur.SmInstagram1 = value["AKUNINSTAGRAM"].(string)
		}
		if value["AKUNYOUTUBE"] != nil {
			hdFaktur.SmYoutube1 = value["AKUNYOUTUBE"].(string)
		}
		if value["PRSH_NAMA"] != nil {
			hdFaktur.KerjaDi11 = value["PRSH_NAMA"].(string)
		}
		if value["PRSH_ALAMAT"] != nil {
			hdFaktur.AlamatKtr11 = value["PRSH_ALAMAT"].(string)
		}

		hdFakturs = append(hdFakturs, hdFaktur)
	}
	tx := cdbConf.Begin()
	if err := tx.Save(&hdFakturs).Error; err != nil {
		tx.Rollback()
		log.Fatalf("error in batch insert: %v", err)
	}
	tx.Commit()

}
