package config

import (
	"fmt"
	"os"
	"strconv"

	oracle "github.com/godoes/gorm-oracle"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewOracleDB() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"SSL":                "false",
	}
	port, _ := strconv.Atoi(os.Getenv("ORACLE_PORT"))
	url := oracle.BuildUrl(os.Getenv("ORACLE_HOST"), port, os.Getenv("ORACLE_MAIN"), os.Getenv("ORACLE_USERNAME"), os.Getenv("ORACLE_PASS"), options)
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

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	dsn := os.Getenv("DB_CDB")
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

func NewLocalDB() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	dsn := os.Getenv("DB_LOCAL")
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
