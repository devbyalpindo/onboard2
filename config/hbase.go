package config

import (
	_ "github.com/apache/calcite-avatica-go/v5"
	"github.com/jinzhu/gorm"
	"log"
)

func ConnectHbase() (*gorm.DB, error) {
	connectionString := "http://localhost:8765/onboarding"
	hbase, err := gorm.Open("avatica", connectionString)

	if err != nil {
		log.Println("Error connect to HBASE : ", err.Error())
		return nil, err
	}

	log.Print("HBASE connection success")
	return hbase, nil
}
