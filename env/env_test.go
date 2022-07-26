// テストやる時は同ディレクトリに.env作れ
/* .env

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=optleaf
DB_USERNAME=root
DB_PASSWORD=root

ENCRYPT_HASHCOST=12

DB_URL=${DB_HOST}:${DB_PORT}

*/

package env

import (
	"testing"

	"github.com/kme222yh/optleaf-server-core/errors"
)

func TestGet(t *testing.T) {
	vals := [...]struct {
		key        string
		defaultVal string
		assumedVal string
	}{
		{"DB_USERNAME", "root", "root"},
		{"DB_USERNAME", "user", "root"},
		{"DB_PASSWORD", "root", "root"},
		{"DB_PASSWORD", "", "root"},
		{"DB_PROTOCOL", "", ""},
		{"DB_URL", "", "127.0.0.1:3306"},
	}

	for i := 0; i < len(vals); i++ {
		val := vals[i]
		result := Get(val.key, val.defaultVal)
		if result != val.assumedVal {
			t.Errorf("env.Get(\"%v\", \"%v\") = \"%v\", want %v", val.key, val.defaultVal, result, val.assumedVal)
		}
	}
}

func TestGetAsInt(t *testing.T) {
	testCases := [...]struct {
		key        string
		defaultVal string

		assumedVal error
	}{
		{"db.aaaaa", "10", nil}, // 未定義
		{"db.aaaaa", "", errors.New("")},
		{"ENCRYPT_HASHCOST", "12", nil},
	}

	for i := 0; i < len(testCases); i++ {
		val := testCases[i]
		_, err := GetAsInt(val.key, val.defaultVal)

		if val.assumedVal == nil {
			if err != nil {
				t.Errorf("user.GetAsInt(%v, %v) = (some error), want nil", val.key, val.defaultVal)
			}
		} else {
			if err == nil {
				t.Errorf("user.GetAsInt(%v, %v) = nil, want (some error)", val.key, val.defaultVal)
			}
		}
	}
}
