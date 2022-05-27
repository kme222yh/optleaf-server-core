package encrypt


import (
    "testing"
    "github.com/kme222yh/optleaf-server-core/errors"
)


func TestHash(t *testing.T) {
    testCases := [...] struct {
        plainText string
        assumedErr error
    } {
        {"hoge", nil},
        {"", nil},
        {" ", nil},
        {"pq234jpr1fj234prj1cvp234jr1opj234rp1", nil},
    }

    for i := 0; i < len(testCases); i++ {
        val := testCases[i]
        _, err := Hash(val.plainText)
        if(val.assumedErr == nil){
            if(err != nil){
                t.Errorf("encrypt.Hash(\"%v\") returns (some error), want nil", val.plainText)
            }
        } else {
            if(err == nil){
                t.Errorf("encrypt.Hash(\"%v\") returns nil, want (some error)", val.plainText)
            }
        }
    }
}


func TestVerifyHash(t *testing.T) {
    testCases := [...] struct {
        plainText string
        compareText string
        assumedErr error
    } {
        {"hoge", "hoge", nil},
        {"hoge", "hage", errors.New("")},
        {"pq234jpr1fj234prj1cvp234jr1opj234rp1", "pq234jpr1fj234prj1cvp234jr1opj234rp1a", errors.New("")},
        {"", "", nil},
        {"", " ", errors.New("")},
    }

    for i := 0; i < len(testCases); i++ {
        val := testCases[i]
        hash, err := Hash(val.plainText)
        err = VerifyHash(hash, val.compareText)
        if(val.assumedErr == nil){
            if(err != nil){
                t.Errorf("encrypt.VerifyHash(hash(\"%v\"), \"%v\") returns (some error), want nil", val.plainText, val.plainText)
            }
        } else {
            if(err == nil){
                t.Errorf("encrypt.VerifyHash(hash(\"%v\"), \"%v\") returns nil, want (some error)", val.plainText, val.plainText)
            }
        }
    }
}
