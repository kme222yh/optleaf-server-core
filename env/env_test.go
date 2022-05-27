// テストやる時は同ディレクトリに.env作れ
/* .env

gin.mode=debug

api.port=:8911

db.user=root
db.password=
db.protocol=null
db.host=127.0.0.1
db.port=3306
db.name=optleaf

jwt.validitPeriod=30

*/


package env


import (
    "testing"
    "github.com/kme222yh/optleaf-server-core/errors"
)


func TestGet(t *testing.T) {
    vals := [...] struct {
        key string
        defaultVal string
        assumedVal string
    } {
        {"gin.mode",    "release",  "debug"},
        {"gin.hage",    "",         ""},
        {"gin.hage",    "hage",     "hage"},
        {"db.user",     "root",     "root"},
        {"db.user",     "user",     "root"},
        {"db.password", "secret",   "secret"},
        {"db.password", "",         ""},
        {"db.protocol", "tcp",      ""},
        {"db.protocol", "",         ""},
    }


    for i := 0; i < len(vals); i++ {
        val := vals[i]
        result := Get(val.key, val.defaultVal)
        if(result != val.assumedVal) {
            t.Errorf("env.Get(\"%v\", \"%v\") = \"%v\", want %v", val.key, val.defaultVal, result, val.assumedVal)
        }
    }
}



func TestGetAsInt(t *testing.T) {
    testCases := [...] struct {
        key string
        defaultVal string

        assumedVal error
    } {
        {"jwt.validitPeriod",   "20",   nil},               // 数値指定
        {"jwt.validitPeriod",   "30",   nil},
        {"jwt.validitPeriod",   "",     nil},
        {"jwt.validitPeriod",   "hoge", nil},
        {"db.password",         "10",   nil},               // 未指定
        {"db.password",         "hoge", errors.New("")},
        {"db.password",         "",     errors.New("")},
        {"db.user",             "10",   errors.New("")},    // 文字指定
        {"db.user",             "",     errors.New("")},
        {"db.user",             "hoge", errors.New("")},
        {"db.aaaaa",            "10",   nil},               // 未定義
        {"db.aaaaa",            "",     errors.New("")},
        {"db.aaaaa",            "hoge", errors.New("")},
    }


    for i := 0; i < len(testCases); i++ {
        val := testCases[i]
        _, err := GetAsInt(val.key, val.defaultVal)

        if(val.assumedVal == nil){
            if(err != nil){
                t.Errorf("user.GetAsInt(%v, %v) = (some error), want nil", val.key, val.defaultVal)
            }
        } else {
            if(err == nil){
                t.Errorf("user.GetAsInt(%v, %v) = nil, want (some error)", val.key, val.defaultVal)
            }
        }
    }
}
