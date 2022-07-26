/*  Usage

// gorm.DB インスタンス取得
db := mysql.Use()

// マイグレーション
mysql.Migrate(model1, model2, ...)

*/

package mysql

import (
	"fmt"

	"github.com/kme222yh/optleaf-server-core/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB接続
func init() {
	dsn := createDsn()
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// DB接続情報を作る
func createDsn() string {
	user := env.Get("DB_USERNAME", "root")
	password := env.Get("DB_PASSWORD", "")
	protocol := env.Get("DB_PROTOCOL", "tcp")
	host := env.Get("DB_HOST", "0.0.0.0")
	port := env.Get("DB_PORT", "3306")
	name := env.Get("DB_DATABASE", "mysql")
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", user, password, protocol, host, port, name)
}

var (
	db  *gorm.DB
	err error
)

func Use() *gorm.DB {
	return db
}

func Migrate(models ...interface{}) {
	db.AutoMigrate(models...)
}
