package service

import (
	"cron/config"
	"cron/entity"
	"fmt"
	"time"
)

func GetDataMohonFaktur() []map[string]interface{} {

	datas := []map[string]interface{}{}
	oracleDB := config.NewOracleDB()
	now := time.Now()
	from := now.AddDate(0, 0, -3)
	if now.Weekday() == 1 {
		from = from.AddDate(0, 0, -2)
	}
	query := fmt.Sprintf("select a.NO_MSN,a.NM1_MOHON,TO_CHAR(A.TGL_LAHIR,'YYYY-MM-DD') TGL_LAHIR,a.NO_TELEPON,a.NO_HP,a.KODE_AGAMA,a.E_MAIL,a.KODE_KERJA,a.AKUNFACEBOOK,a.AKUNTWITTER,a.AKUNINSTAGRAM,a.AKUNYOUTUBE,a.AL1_MOHON,TO_CHAR(A.TGL_FAKTUR,'YYYY-MM-DD') TGL_FAKTUR,a.NO_DLRP,a.NM_SALES,d.NM_MTR,a.JNS_JUAL,	a.JNS_BELI,a.NO_LEAS,a.NO_RGK,a.NO_KTPNPWP,a.KEL_MOHON,a.KEC_MOHON,a.KODE_POS,a.ID_SALES_ASLI,a.NO_KK,a.KODE_DIDIK,a.KELUAR_BLN,a.PRSH_NAMA,a.PRSH_ALAMAT,TO_CHAR(A.TGL_MOHON,'YYYY-MM-DD') TGL_MOHON from  maindealer.mohonfaktur a, maindealer.ppblistdetail b, maindealer.typemotorahm c, maindealer.typemotor d where a.no_msn=b.no_msn and b.kd_mdl=c.kd_mdl	and c.no_mtr=d.no_mtr and a.tgl_faktur>=to_date('%s 00:00:00','ddmmyyyy hh24:mi:ss')  and a.tgl_faktur<=to_date('%s 23:59:59','ddmmyyyy hh24:mi:ss')", from.Format("02012006"), now.Format("02012006"))

	result := oracleDB.Raw(query).Scan(&datas)
	if result.Error != nil {
		fmt.Println("errornya disini yaa ", result.Error)
	}
	return datas
}

func GetDataMohonFakturMbaLeli() []map[string]interface{} {

	datas := []map[string]interface{}{}
	oracleDB := config.NewOracleDB()
	now := time.Now()
	query := fmt.Sprintf("select a.NO_MSN,a.NO_RGK,a.NO_DLRP,E.NM_DLR,a.NM1_MOHON,a.NM2_MOHON,TO_CHAR(a.TGL_LAHIR,'YYYY-MM-DD') TGL_LAHIR ,a.JNS_KLM,a.NO_TELEPON, a.NO_HP,a.AL1_MOHON,a.AL2_MOHON,a.RT_MOHON,a.RW_MOHON,a.KEL_MOHON,a.KEC_MOHON,a.KOTA_MOHON,a.KODE_POS, a.AL_SRT,a.AL1_SRT,a.AL2_SRT,a.AL3_SRT,a.KEC_SRT,a.KOTA_SRT,a.pos_srt, a.NAMA_PNJJWB,d.no_mtr, d.nm_mtr,a.kode_kerja, a.kode_didik, a.keluar_bln, a.motor_hir, a.jns_beli, a.bdn_usaha, a.jns_jual, TO_CHAR(a.tgl_mohon,'YYYY-MM-DD') TGL_MOHON,a.id_sales_asli, a.nm_sales_asli, a.no_ktpnpwp, a.dp, a.cicilan, a.sts_rumah, a.sts_hp, a.e_mail, a.no_leas,a.jml_angsuran,a.KODE_AGAMA,a.SEDIA_DIHUB,a.JNS_MOTOR,a.SM_DIBELI,a.PROP_MOHON,a.TUJU_PAK1,a.AKUNFACEBOOK,a.AKUNINSTAGRAM,a.AKUNTWITTER,a.AKUNYOUTUBE,a.HOBI, a.NO_KK,TO_CHAR(a.tgl_faktur,'YYYY-MM-DD') TGL_FAKTUR,a.PRSH_NAMA,a.PRSH_ALAMAT,a.PRSH_KEC,a.PRSH_KOTA,a.PRSH_PROP,a.NO_NPWP,a.AKTIF_JUAL from maindealer.mohonfaktur a, maindealer.ppblistdetail b, maindealer.typemotorahm c, maindealer.typemotor d, maindealer.dealer e where a.no_msn=b.no_msn and b.kd_mdl=c.kd_mdl and c.no_mtr=d.no_mtr and a.no_dlrp=e.no_dlr and a.tgl_faktur>=to_date('%s 00:00:00','ddmmyyyy hh24:mi:ss') and a.tgl_faktur<=to_date('%s 23:59:59','ddmmyyyy hh24:mi:ss') and SUBSTR(a.AL1_MOHON,1,5)<>'PUTIH'", GetMaxTglFakturTrWmsFaktur2().TglFaktur.Format("02012006"), now.Format("02012006"))

	result := oracleDB.Raw(query).Scan(&datas)
	if result.Error != nil {
		fmt.Println("errornya disini yaa ", result.Error)
	}
	return datas
}

func GetDataMbaLeli() []entity.MbaLeliFull {
	localDb2 := config.NewLocal2DB()
	datas := []entity.MbaLeliFull{}
	query := "select * from data_mba_leli"
	result := localDb2.Raw(query).Scan(&datas)
	if result.Error != nil {
		fmt.Println("errornya disini yaa ", result.Error)
	}
	return datas

}

func GetMaxTglFakturTrWmsFaktur2() entity.MaxTglFaktur {
	localDb := config.NewLocalDB()
	data := entity.MaxTglFaktur{}
	query := "select max(tgl_faktur) tgl_faktur from tr_wms_faktur2"
	result := localDb.Raw(query).Scan(&data)
	if result.Error != nil {
		fmt.Println("errornya disini yaa ", result.Error)
	}
	return data
}
func GetMaxTglFakturHdFaktur2() entity.MaxTglFaktur {
	cdbConf := config.NewCDBWebDB()
	data := entity.MaxTglFaktur{}
	query := "select max(tgl_faktur) hd_faktur2 from tr_wms_faktur2"
	result := cdbConf.Raw(query).Scan(&data)
	if result.Error != nil {
		fmt.Println("errornya disini yaa ", result.Error)
	}
	return data
}

// {JM04E1872694 08979620183 08979620183 ardidwicahyono@gmail.com 2b}
