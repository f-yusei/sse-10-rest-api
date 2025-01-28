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
		GetEnv("DB_PORT", "4000"),
		GetEnv("DB_NAME", "your_db_name"),
	)

	DB, err = gorm.Open(gormsql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&entity.Bell{}, &entity.Store{}, &entity.CallLog{})
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
