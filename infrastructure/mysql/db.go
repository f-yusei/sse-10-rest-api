package infrastructure

import (
	"crypto/tls"
	"fmt"
	"gcp_go_cloud_run/app/infrastructure/mysql/entity"
	"log"
	"os"

	gormsql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func InitDB() {

	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: GetEnv("DB_HOST", "localhost"),
	})
	if err != nil {
		log.Fatalf("failed to register TLS config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&tls=tidb",
		GetEnv("DB_USER", "root"),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "5000"),
		GetEnv("DB_NAME", "your_db_name"),
	)

	DB, err = gorm.Open(gormsql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&entity.Bell{}, &entity.Store{}, &entity.CallLog{})
}

func Seed() {
	// Seed data
	bell := entity.Bell{ID: 1, StoreID: 0, DeviceID: "device_id_1", Status: "active"}
	store := entity.Store{Name: "たこ焼き", Bells: []entity.Bell{bell}}
	DB.Create(&store)
}

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func DebugPrint(v interface{}) {
	fmt.Printf("%+v\n", v)
}
