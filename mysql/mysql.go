/*  Usage

// gorm.DB インスタンス取得
db := mysql.Use()

// マイグレーション
mysql.Migrate(model1, model2, ...)

*/


package mysql

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/kme222yh/optleaf-server-core/env"
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
    user := env.Get("db.user", "root")
    password := env.Get("db.password", "")
    protocol := env.Get("db.protocol", "tcp")
    host := env.Get("db.host", "0.0.0.0")
    port := env.Get("db.port", "3306")
    name := env.Get("db.name", "mysql")
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
