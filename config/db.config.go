package config

import (
	"fmt"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewOracleDB() *gorm.DB {
	options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"SSL":                "false",
	}
	url := oracle.BuildUrl("192.168.0.2", 1521, "wmsbase", "nuel", "nuel", options)
	dialector := oracle.New(oracle.Config{
		DSN:                       url,
		IgnoreCase:                false,
		NamingCaseSensitive:       true,
		VarcharSizeIsCharLength:   true,
		RowNumberAliasForOracle11: "ROW_NUM",
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     false,
		CreateBatchSize: 50,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func NewCDBWebDB() *gorm.DB {
	dsn := "root2:root2@tcp(192.168.12.171:3306)/cdb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     false,
		CreateBatchSize: 50,
	})
	if err != nil {
		fmt.Println("Masuk sini gk guys ", err)
		panic(err)
	}
	return db
}
