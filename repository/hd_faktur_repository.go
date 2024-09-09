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

		ketNoTelepon := "VALID"
		ketNoHp := "VALID"
		vEmail := "VALID"
		if !(strings.HasPrefix(value["NO_TELEPON"].(string), "08") || strings.HasPrefix(value["NO_TELEPON"].(string), "628")) {
			ketNoTelepon = "TIDAK VALID"
		}
		if !(strings.HasPrefix(value["NO_HP"].(string), "08") || strings.HasPrefix(value["NO_HP"].(string), "628")) {
			ketNoHp = "TIDAK VALID"
		}
		if len(value["E_MAIL"].(string)) < 18 {
			vEmail = "TIDAK VALID"
		}

		hdFaktur := entity.HDFaktur2Full{
			NoMsn:        value["NO_MSN"].(string),
			NmCustomer11: value["NM1_MOHON"].(string),
			TglLahir2:    value["TGL_LAHIR"].(string),
			KetNoTelp1:   ketNoTelepon,
			KetNoHp1:     ketNoHp,
			VEmail:       vEmail,
			KAgama2:      masterKode["agama"][value["KODE_AGAMA"].(string)],
			KKodeKerja1:  masterKode["kerja"][value["KODE_KERJA"].(string)],
			KJnsJual:     masterKode["jual"][value["JNS_JUAL"].(string)],
			KJnsBeli:     masterKode["beli"][value["JNS_BELI"].(string)],
			KKodeDidik1:  masterKode["pendidikan"][value["KODE_DIDIK"].(string)],
			KKeluarBln1:  masterKode["pengeluaran"][value["KELUAR_BLN"].(string)],
		}

		hdFakturs = append(hdFakturs, hdFaktur)
	}
	tx := cdbConf.Begin()
	if err := tx.Create(&hdFakturs).Error; err != nil {
		tx.Rollback()
		log.Fatalf("error in batch insert: %v", err)
	}
	tx.Commit()

}
