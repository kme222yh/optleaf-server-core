package mysql


import (
    "testing"
)


func TestCreateDsn(t *testing.T) {
    testCases := [...] string {
        "root:@tcp(0.0.0.0:3306)/mysql?parseTime=true",
    }

    for i := 0; i < len(testCases); i++ {
        val := testCases[i]
        result := createDsn()
        if(result != val) {
            t.Errorf("createDsn() = \"%v\", want %v", result, val)
        }
    }
}
