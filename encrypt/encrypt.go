/*  Usage

// 文字列のハッシュ化
hash, err := encrypt.hash(plainText)
if(err != nil){
    // error
}


// ハッシュとプレインテキストの検証
err := encrypt.VerifyHash(hash, plainText)
if(err != nil){
    // error
}


*/


package encrypt


import (
    "unsafe"
    "golang.org/x/crypto/bcrypt"
    "github.com/kme222yh/optleaf-server-core/env"
)


func init(){
    var err error
    hashCost, err = env.GetAsInt("encrypt.hashCost", "12")
    if err != nil {
        hashCost = 12
	}
}


var (
    hashCost int
)


func Hash(plainText string) (string, error) {
    hashByte, err := bcrypt.GenerateFromPassword([]byte(plainText), hashCost)
    hash := *(*string)(unsafe.Pointer(&hashByte))
    return hash, err
}
func VerifyHash(hash string, plainText string) error {
    hashByte := []byte(hash)
    err := bcrypt.CompareHashAndPassword(hashByte,[]byte(plainText))
    return err
}
