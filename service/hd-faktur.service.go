package service

import (
	"cron/config"
	"fmt"
	"time"
)

func GetDataMohonFaktur() {

	datas := []map[string]interface{}{}
	oracleDB := config.NewOracleDB()
	now := time.Now()
	query := fmt.Sprintf("select a.NO_MSN,a.NM1_MOHON,TO_CHAR(A.TGL_LAHIR,'YYYY-MM-DD'),a.NO_TELEPON,a.NO_HP,a.KODE_AGAMA,a.E_MAIL,a.KODE_KERJA,a.AKUNFACEBOOK,a.AKUNTWITTER,a.AKUNINSTAGRAM,a.AKUNYOUTUBE,a.AL1_MOHON,TO_CHAR(A.TGL_FAKTUR,'YYYY-MM-DD'),a.NO_DLRP,a.NM_SALES,d.NM_MTR,a.JNS_JUAL,  from  maindealer.mohonfaktur a, maindealer.ppblistdetail b, maindealer.typemotorahm c, maindealer.typemotor d where a.no_msn=b.no_msn and b.kd_mdl=c.kd_mdl and c.no_mtr=d.no_mtr and a.tgl_faktur>=to_date('%s 00:00:00','ddmmyyyy hh24:mi:ss')  and a.tgl_faktur<=to_date('%s 23:59:59','ddmmyyyy hh24:mi:ss')", now.Format("02012006"), now.Format("02012006"))

	oracleDB.Raw(query).Scan(&datas)
}

func UpdateDataMohonFaktur() {
	insertQueryBegin := fmt.Sprintf("INSERT INTO cdb.hd_faktur2023 (no_msn,nm_customer11,tgl_lahir2,no_telp1,ket_no_telp1,no_hp1,ket_no_hp1,agama2,k_agama2,email1,v_email,kode_kerja,k_kode_kerja1,sm_facebook1,sm_twitter1,sm_instagram1,sm_youtube1,alamat11,tgl_faktur,kd_dlr,nm_sales1,sts_valid,nm_mtr,jns_jual,k_jns_jual,jns_beli,k_jns_beli,no_leas,k_no_leas,no_rgk,no_ktpnpwp,kel1,kec1,kodepos1,id_sales,no_kk,kode_didik,k_kode_didik1,keluar_bln,k_keluar_bln1,kerja_di11,alamat_ktr11,tgl_mohon) VALUES %s", "test")
	// insertQuery := fmt.Sprintf("('JME1E1091661','SUHARNO','1984-09-25','089512893458','VALID','089512893458','VALID','1','ISLAM','N@GMAIL.COM','TIDAK VALID','11','Lain-lain','','','','','KP. KARANG MULYA RT 003 RW 001','2024-08-26','055','AFIFAH HUSNA','TIDAK VALID','BEAT SPORTY CBS','2','Kredit','I','INDIVIDUAL','06','ADIRA - ADIRA DINAMIKA MULTI FINANCE','JME115RK091355','3309172509840001','KARANG MULYA','KARANG TENGAH','15157','583594','3671121711150012','4','SLTA/SMA','5','Rp 2,500,001 - Rp 4,000,000','','','2024-08-16');")
	insertQuery := fmt.Sprintf("('JME1E1091661','SUHARNO','1984-09-25','089512893458','VALID','089512893458','VALID','1','ISLAM','N@GMAIL.COM','TIDAK VALID','11','Lain-lain','','','','','KP. KARANG MULYA RT 003 RW 001','2024-08-26','055','AFIFAH HUSNA','TIDAK VALID','BEAT SPORTY CBS','2','Kredit','I','INDIVIDUAL','06','ADIRA - ADIRA DINAMIKA MULTI FINANCE','JME115RK091355','3309172509840001','KARANG MULYA','KARANG TENGAH','15157','583594','3671121711150012','4','SLTA/SMA','5','Rp 2,500,001 - Rp 4,000,000','','','2024-08-16'); %s ", "test")

	fmt.Println("ini data ", insertQuery, insertQueryBegin)
}

// {JM04E1872694 08979620183 08979620183 ardidwicahyono@gmail.com 2b}
