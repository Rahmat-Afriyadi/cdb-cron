package repository

import (
	"cron/entity"
	"fmt"
)

func UpdateDataMbaLeli(mohonFaktur []map[string]interface{}) {
	// cdbConf := config.NewCDBWebDB()

	hdFakturs := []entity.MbaLeliFull{}
	for _, value := range mohonFaktur {
		hdFaktur := entity.MbaLeliFull{
			NoMsn:         value["NO_MSN"].(string),
			NoRgk:         value["NO_RGK"].(string),
			KdDlr:         value["NO_DLRP"].(string),
			NmDlr:         value["NM_DLR"].(string),
			NmCustomer11:  value["NM1_MOHON"].(string),
			TglLahir1:     value["TGL_LAHIR"].(string),
			JnsKlm1:       value["JNS_KLM"].(string),
			NoHp1:         value["NO_HP"].(string),
			Alamat11:      value["AL1_MOHON"].(string),
			Kel1:          value["KEL_MOHON"].(string),
			Kec1:          value["KEC_MOHON"].(string),
			Kota1:         value["KOTA_MOHON"].(string),
			Kodepos1:      value["KODE_POS"].(string),
			PicPerusahaan: value["NAMA_PNJJWB"].(string),
			TglMohon:      value["TGL_MOHON"].(string),
			KdUser:        "99",
			KdCustomer:    "1",
			IdSales:       value["id_sales_asli"].(string),
			NmSales1:      value["nm_sales_asli"].(string),
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
			// TujuanPakai2:        value[""].(string),
			// KeluarBln2:          value[""].(string),
			// StsSms:              value[""].(string),
			// StsCetak3:           value[""].(string),
			NoKtpNpwp:           value["no_ktpnpwp"].(string),
			DpMtr:               value["dp"].(string),
			CicilanMtr:          value["cicilan"].(uint64),
			StsKepemilikanRumah: value["sts_rumah"].(string),
			StsHpPraPasca:       value["sts_hp"].(string),
			KdClient:            value[""].(string),
			Email1:              value["E_MAIL"].(string),
			NoLeas:              value["no_leas"].(string),
			Angsuran:            value["jml_angsuran"].(string),
			Agama1:              value["KODE_AGAMA"].(string),
			SediaDihub:          value["SEDIA_DIHUB"].(string),
			JnsMotor:            value["JNS_MOTOR"].(string),
			SmDibeli:            value["SM_DIBELI"].(string),
			PropMohon:           value["PROP_MOHON"].(string),
			TglAkhirTenor:       value[""].(string),
			Hobby3:              value["HOBI"].(string),
			NoKk:                value["NO_KK"].(string),
			TglFaktur:           value["TGL_FAKTUR"].(string),
			AktifJual:           value["AKTIF_JUAL"].(string),
		}
		if value["NM2_MOHON"] != nil {
			hdFaktur.NmCustomer12 = value["NM2_MOHON"].(string)
		}
		if value["NO_TELEPON"] != nil {
			hdFaktur.NoTelp1 = value["NO_TELEPON"].(string)
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
		if value["KELUAR_BULAN"] != nil {
			hdFaktur.KeluarBln = value["KELUAR_BULAN"].(string)
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

		hdFakturs = append(hdFakturs, hdFaktur)
	}
	fmt.Println("ini data yaa ", hdFakturs[0])
	// if len(hdFakturs) > 0 {
	// 	tx := cdbConf.Begin()
	// 	if err := tx.Save(&hdFakturs).Error; err != nil {
	// 		tx.Rollback()
	// 		log.Fatalf("error in batch insert: %v", err)
	// 	}
	// 	tx.Commit()
	// }

}
