/*  Usage

.envの中について
keyの値が「何もない」場合はnullを指定しろ

// 環境変数取り出し
env.Get("hogehoge", "default") => string

env.Get()がdefualt valueを要求するので、.envファイルが存在しなくても動くようにする→した


os.Getenv(key)はkeyが無効であれば""を返す
ソース：https://go.dev/play/p/HOIUGERr2SM


int型で欲しい時は
env.GetAsInt("hogehoge", "default") => int
数字で指定してるやつだけに使ってね、じゃないとエラー吐く

*/


package env


import (
    "os"
    "strconv"

    // ↓で.env読み込み、無ければ無視
    _ "github.com/joho/godotenv/autoload"
)


func Get(key string, defaultVal string) string {
    val := os.Getenv(key)
    if(val == "") {
        return defaultVal
    } else if(val == "null") {
        return ""
    } else {
        return val
    }
}


func GetAsInt(key string, defaultVal string) (int, error) {
    val := Get(key, defaultVal)
    numVal, err := strconv.Atoi(val)
    return numVal, err
}
