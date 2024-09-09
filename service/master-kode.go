package service

import (
	"cron/config"
)

func GetDataMasterKode() map[string]map[string]string {

	localDb := config.NewLocalDB()
	data := map[string]map[string]string{}
	rawItems := []map[string]interface{}{}
	localDb.Raw("select kode_kerja2, nm_kerja from mst_kerja ").Scan(&rawItems)
	items := map[string]string{}
	for _, value := range rawItems {
		items[value["kode_kerja2"].(string)] = value["nm_kerja"].(string)
	}
	data["kerja"] = items

	items = map[string]string{}
	rawItems = []map[string]interface{}{}
	localDb.Raw("select kd_pendidikan, nm_pendidikan from mst_pendidikan ").Scan(&rawItems)
	for _, value := range rawItems {
		items[value["kd_pendidikan"].(string)] = value["nm_pendidikan"].(string)
	}
	data["pendidikan"] = items

	items = map[string]string{}
	rawItems = []map[string]interface{}{}
	localDb.Raw("select kd_pengeluaran, pengeluaran from mst_pengeluaran ").Scan(&rawItems)
	for _, value := range rawItems {
		items[value["kd_pengeluaran"].(string)] = value["pengeluaran"].(string)
	}
	data["pengeluaran"] = items

	items = map[string]string{}
	rawItems = []map[string]interface{}{}
	localDb.Raw("select no_leas2, nm_leasing from mst_leasing ").Scan(&rawItems)
	for _, value := range rawItems {
		items[value["no_leas2"].(string)] = value["nm_leasing"].(string)
	}
	data["leasing"] = items

	items = map[string]string{}
	rawItems = []map[string]interface{}{}
	localDb.Raw("select kd_agama, agama from mst_agama ").Scan(&rawItems)
	for _, value := range rawItems {
		items[value["kd_agama"].(string)] = value["agama"].(string)
	}
	data["agama"] = items

	items = map[string]string{}
	rawItems = []map[string]interface{}{}
	localDb.Raw("select kd_jual, nm_jual from mst_jns_jual ").Scan(&rawItems)
	for _, value := range rawItems {
		items[value["kd_jual"].(string)] = value["nm_jual"].(string)
	}
	data["jual"] = items

	items = map[string]string{}
	rawItems = []map[string]interface{}{}
	localDb.Raw("select kd_beli, nm_beli from mst_jns_beli ").Scan(&rawItems)
	for _, value := range rawItems {
		items[value["kd_beli"].(string)] = value["nm_beli"].(string)
	}
	data["beli"] = items

	return data
}
