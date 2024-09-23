package entity

type MbaLeliFull struct {
	NoMsn               string  `gorm:"primary_key;column:no_msn"`
	NoRgk               string  `gorm:"column:no_rgk"`
	KdDlr               string  `gorm:"column:kd_dlr"`
	NmDlr               string  `gorm:"column:nm_dlr"`
	NmCustomer11        string  `gorm:"column:nm_customer11"`
	NmCustomer12        string  `gorm:"column:nm_customer12"`
	TglLahir1           string  `gorm:"column:tgl_lahir1"`
	JnsKlm1             string  `gorm:"column:jns_klm1"`
	NoTelp1             string  `gorm:"column:no_telp1"`
	NoHp1               string  `gorm:"column:no_hp1"`
	Alamat11            string  `gorm:"column:alamat11"`
	Alamat12            string  `gorm:"column:alamat12"`
	Rt1                 string  `gorm:"column:rt1"`
	Rw1                 string  `gorm:"column:rw1"`
	Kel1                string  `gorm:"column:kel1"`
	Kec1                string  `gorm:"column:kec1"`
	Kota1               string  `gorm:"column:kota1"`
	Kodepos1            string  `gorm:"column:kodepos1"`
	StsAlamat           string  `gorm:"column:sts_alamat"`
	AlamatSrt11         string  `gorm:"column:alamat_srt11"`
	AlamatSrt12         string  `gorm:"column:alamat_srt12"`
	KelSrt1             string  `gorm:"column:kel_srt1"`
	KecSrt1             string  `gorm:"column:kec_srt1"`
	KotaSrt1            string  `gorm:"column:kota_srt1"`
	KodeposSrt1         string  `gorm:"column:kodepos_srt1"`
	PicPerusahaan       string  `gorm:"column:pic_perusahaan"`
	NoMtr               string  `gorm:"column:no_mtr"`
	NmMtr               string  `gorm:"column:nm_mtr"`
	KodeKerja           string  `gorm:"column:kode_kerja"`
	KodeDidik           string  `gorm:"column:kode_didik"`
	KeluarBln           string  `gorm:"column:keluar_bln"`
	MotorHir            string  `gorm:"column:motor_hir"`
	JnsBeli             string  `gorm:"column:jns_beli"`
	BdnUsaha            string  `gorm:"column:bdn_usaha"`
	JnsJual             string  `gorm:"column:jns_jual"`
	TglMohon            string  `gorm:"column:tgl_mohon"`
	KdUser              string  `gorm:"column:kd_user"`
	KdCustomer          string  `gorm:"column:kd_customer"`
	IdSales             string  `gorm:"column:id_sales"`
	NmSales1            string  `gorm:"column:nm_sales1"`
	TujuanPakai1        string  `gorm:"column:tujuan_pakai1"`
	JnsKlm2             string  `gorm:"column:jns_klm2"`
	StsKirim            string  `gorm:"column:sts_kirim"`
	StsCetak            string  `gorm:"column:sts_cetak"`
	StsSource           string  `gorm:"column:sts_source"`
	StsAsuransi         string  `gorm:"column:sts_asuransi"`
	StsTertanggung      string  `gorm:"column:sts_tertanggung"`
	StsKawin            string  `gorm:"column:sts_kawin"`
	StsLengkap          string  `gorm:"column:sts_lengkap"`
	Agama2              string  `gorm:"column:agama2"`
	KodeDidik2          string  `gorm:"column:kode_didik2"`
	KodeKerja2          string  `gorm:"column:kode_kerja2"`
	TujuanPakai2        string  `gorm:"column:tujuan_pakai2"`
	KeluarBln2          string  `gorm:"column:keluar_bln2"`
	StsSms              string  `gorm:"column:sts_sms"`
	StsCetak3           int     `gorm:"column:sts_cetak3"`
	NoKtpNpwp           string  `gorm:"column:no_ktpnpwp"`
	DpMtr               float64 `gorm:"column:dp_mtr"`
	CicilanMtr          float64 `gorm:"column:cicilan_mtr"`
	StsKepemilikanRumah string  `gorm:"column:sts_kepemilikan_rumah"`
	StsHpPraPasca       string  `gorm:"column:sts_hp_pra_pasca"`
	KdClient            string  `gorm:"column:kd_client"`
	Email1              string  `gorm:"column:email1"`
	NoLeas              string  `gorm:"column:no_leas"`
	Angsuran            string  `gorm:"column:angsuran"`
	Agama1              string  `gorm:"column:agama1"`
	SediaDihub          string  `gorm:"column:sedia_dihub"`
	JnsMotor            string  `gorm:"column:jns_motor"`
	SmDibeli            string  `gorm:"column:sm_dibeli"`
	PropMohon           string  `gorm:"column:prop_mohon"`
	SmsNo               string  `gorm:"column:sms_no"`
	TglAkhirTenor       *string `gorm:"column:tgl_akhir_tenor"`
	SmFacebook1         string  `gorm:"column:sm_facebook1"`
	SmInstagram1        string  `gorm:"column:sm_instagram1"`
	SmTwitter1          string  `gorm:"column:sm_twitter1"`
	SmYoutube1          string  `gorm:"column:sm_youtube1"`
	Hobby3              string  `gorm:"column:hobby3"`
	NoKk                string  `gorm:"column:no_kk"`
	TglFaktur           string  `gorm:"column:tgl_faktur"`
	KerjaDiFkt          string  `gorm:"column:kerja_di_fkt"`
	AlamatKtrFkt        string  `gorm:"column:alamat_ktr_fkt"`
	KecKtrFkt           string  `gorm:"column:kec_ktr_fkt"`
	KotaKtrFkt          string  `gorm:"column:kota_ktr_fkt"`
	PropKtrFkt          string  `gorm:"column:prop_ktr_fkt"`
	NoNpwp              string  `gorm:"column:no_npwp"`
	AktifJual           string  `gorm:"column:aktif_jual"`
	AktifJual2          string  `gorm:"default:'';column:aktif_jual2"`
	KetAktifJual2       string  `gorm:"default:'';column:ket_aktif_jual2"`
	KdCard              string  `gorm:"default:'';column:kd_card"`
	NamaKtp             string  `gorm:"default:'';column:nama_ktp"`
	KdDlr2              string  `gorm:"default:'';column:kd_dlr2"`
	Hobby1              string  `gorm:"default:'';column:hobby1"`
	NmKerja2            string  `gorm:"default:'';column:nm_kerja2"`
	AlasanTdkDownload   string  `gorm:"default:'';column:alasan_tdk_download"`

	IdMerkMtrSeblm        string  `gorm:"default:'';column:id_merk_mtr_seblm"`
	IdJnsMtrSeblm         string  `gorm:"default:'';column:id_jns_mtr_seblm"`
	AlasanTdkRenewal2     string  `gorm:"default:'';column:alasan_tdk_renewal2"`
	TerimaKartu           string  `gorm:"default:'';column:terima_kartu"`
	NmLeas2               string  `gorm:"default:'';column:nm_leas2"`
	StsKonfirmasi         string  `gorm:"default:'';column:sts_konfirmasi"`
	AlasanVoidKonfirmasi  string  `gorm:"default:'';column:alasan_void_konfirmasi"`
	StsRo                 string  `gorm:"default:'';column:sts_ro"`
	StsBawaKartu          string  `gorm:"default:'';column:sts_bawa_kartu"`
	StsBawaKartu2         string  `gorm:"default:'';column:sts_bawa_kartu2"`
	StsKartu              string  `gorm:"default:'';column:sts_kartu"`
	BlnCetakKartu         string  `gorm:"default:'';column:bln_cetak_kartu"`
	StsRenewal2           string  `gorm:"default:'';column:sts_renewal2"`
	AlasanBelumBayar2     string  `gorm:"default:'';column:alasan_belum_bayar2"`
	DownloadMa            string  `gorm:"default:'';column:download_ma"`
	KartuDiterimaSesuai   string  `gorm:"default:'';column:kartu_diterima_sesuai"`
	SmFacebook            string  `gorm:"default:'';column:sm_facebook"`
	SmInstagram           string  `gorm:"default:'';column:sm_instagram"`
	SmTwitter             string  `gorm:"default:'';column:sm_twitter"`
	SmYoutube             string  `gorm:"default:'';column:sm_youtube"`
	SalesPeopleInfoKpbDig string  `gorm:"default:'';column:sales_people_info_kpb_dig"`
	DeliveryManInfoKpbDig string  `gorm:"default:'';column:delivery_man_info_kpb_dig"`
	TypeMtrSeblm          string  `gorm:"default:'';column:type_mtr_seblm"`
	MerkMtrSeblm          string  `gorm:"default:'';column:merk_mtr_seblm"`
	SegmenMtrSebelm       string  `gorm:"default:'';column:segmen_mtr_sebelm"`
	KdDlrSeblm            string  `gorm:"default:'';column:kd_dlr_seblm"`
	StsKdDlrSeblm         string  `gorm:"default:'';column:sts_kd_dlr_seblm"`
	DpBayarKredit         float64 `gorm:"default:0;column:dp_bayar_kredit"`
	SumberDanaCash        string  `gorm:"default:'';column:sumber_dana_cash"`
	SumberDanaCashLainnya string  `gorm:"default:'';column:sumber_dana_cash_lainnya"`
	AlasanSumberDanaCash  string  `gorm:"default:'';column:alasan_sumber_dana_cash"`
	KdAktivitasJualR      string  `gorm:"default:'';column:kd_aktivitas_jual_r"`
	JmlCallRenewal        float64 `gorm:"default:0;column:jml_call_renewal"`
}
