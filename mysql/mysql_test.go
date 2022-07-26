/*

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=optleaf
DB_USERNAME=root
DB_PASSWORD=root

*/

package mysql

import (
	"testing"
)

func TestCreateDsn(t *testing.T) {
	testCases := [...]string{
		"root:root@tcp(127.0.0.1:3306)/optleaf?parseTime=true",
	}

	for i := 0; i < len(testCases); i++ {
		val := testCases[i]
		result := createDsn()
		if result != val {
			t.Errorf("createDsn() = \"%v\", want %v", result, val)
		}
	}
}
