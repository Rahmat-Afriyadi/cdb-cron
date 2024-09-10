package entity

type HDFaktur2Full struct {
	NoMsn        string `gorm:"primary_key;column:no_msn"`
	NmCustomer11 string `gorm:"column:nm_customer11"`
	TglLahir2    string `gorm:"column:tgl_lahir2"`
	NoTelp1      string `gorm:"column:no_telp1"`
	KetNoTelp1   string `gorm:"column:ket_no_telp1"`
	NoHp1        string `gorm:"column:no_hp1"`
	KetNoHp1     string `gorm:"column:ket_no_hp1"`
	Agama2       string `gorm:"column:agama2"`
	KAgama2      string `gorm:"column:k_agama2"`
	Email1       string `gorm:"column:email1"`
	VEmail       string `gorm:"column:v_email"`
	KodeKerja    string `gorm:"column:kode_kerja"`
	KKodeKerja1  string `gorm:"column:k_kode_kerja1"`
	SmFacebook1  string `gorm:"column:sm_facebook1"`
	SmTwitter1   string `gorm:"column:sm_twitter1"`
	SmInstagram1 string `gorm:"column:sm_instagram1"`
	SmYoutube1   string `gorm:"column:sm_youtube1"`
	Alamat11     string `gorm:"column:alamat11"`
	TglFaktur    string `gorm:"column:tgl_faktur"`
	KdDlr        string `gorm:"column:kd_dlr"`
	NmSales1     string `gorm:"column:nm_sales1"`
	StsValid     string `gorm:"column:sts_valid"`
	NmMtr        string `gorm:"column:nm_mtr"`
	JnsJual      string `gorm:"column:jns_jual"`
	KJnsJual     string `gorm:"column:k_jns_jual"`
	JnsBeli      string `gorm:"column:jns_beli"`
	KJnsBeli     string `gorm:"column:k_jns_beli"`
	NoLeas       string `gorm:"column:no_leas"`
	KNoLeas      string `gorm:"column:k_no_leas"`
	NoRgk        string `gorm:"column:no_rgk"`
	NoKtpnpwp    string `gorm:"column:no_ktpnpwp"`
	Kel1         string `gorm:"column:kel1"`
	Kec1         string `gorm:"column:kec1"`
	Kodepos1     string `gorm:"column:kodepos1"`
	IdSales      string `gorm:"column:id_sales"`
	NoKk         string `gorm:"column:no_kk"`
	KodeDidik    string `gorm:"column:kode_didik"`
	KKodeDidik1  string `gorm:"column:k_kode_didik1"`
	KeluarBln    string `gorm:"column:keluar_bln"`
	KKeluarBln1  string `gorm:"column:k_keluar_bln1"`
	KerjaDi11    string `gorm:"column:kerja_di11"`
	AlamatKtr11  string `gorm:"column:alamat_ktr11"`
	TglMohon     string `gorm:"column:tgl_mohon"`
}

// func (HDFaktur2Full) TableName() string {
// 	return "hd_faktur2"
// }
