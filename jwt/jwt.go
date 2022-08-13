package jwt

/*	interface

token := jwt.New().For("access").ValidityDays(2).Create(claims)
if err := token.Error(); err != nil {
	return err
}
tokenString := token.String()



token := jwt.New().For("access").Verify(tokenString)
if err := token.Error(); err != nil {
	return err
}
email := token.ParseString("email")
id := token.ParseInt("id")

*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/kme222yh/optleaf-server-core/env"
)

type JwtToken struct {
	subject      string
	validityDays int
	tokenString  string
	context      *jwt_go.Token
	claims       jwt_go.MapClaims
	err          error
}

var (
	appKey string
)

func init() {
	appKey = env.Get("APP_KEY", "")
	if appKey == "" {
		panic("APP_KEY does not exist in .env")
	}
}

func New() *JwtToken {
	token := new(JwtToken)
	token.subject = env.Get("JWT_DEFAULT_SUBJECT", "sklruntaiobvaio4b6tiua4o2")
	var err error
	token.validityDays, err = env.GetAsInt("JWT_DEFAULT_VALIDITYDAYS", "1")
	if err != nil {
		panic("JWT setting error.")
	}
	return token
}

func (t *JwtToken) For(subject string) *JwtToken {
	t.subject = subject
	return t
}

func (t *JwtToken) ValidityDays(validityDays int) *JwtToken {
	t.validityDays = validityDays
	return t
}

func (t *JwtToken) Create(claims_struct interface{}) *JwtToken {
	claims, _ := structToJsonTagMap(claims_struct)
	claims["sub"] = t.subject
	claims["exp"] = time.Now().AddDate(0, 0, t.validityDays).Unix()
	t.context = jwt_go.NewWithClaims(jwt_go.SigningMethodHS256, jwt_go.MapClaims(claims))
	t.tokenString, t.err = t.context.SignedString([]byte(appKey))
	return t
}

func (t *JwtToken) String() string {
	return t.tokenString
}

func (t *JwtToken) Verify(tokenString string) *JwtToken {
	t.context, t.err = parse(tokenString)
	if t.err != nil {
		return t
	}
	var ok bool
	t.claims, ok = t.context.Claims.(jwt_go.MapClaims)
	if t.claims["sub"].(string) != t.subject {
		t.err = fmt.Errorf("token's subject mismatch.")
	} else if ok == false || t.context.Valid == false {
		t.err = fmt.Errorf("The token has some problems.")
	}
	return t
}

func (t *JwtToken) Error() error {
	return t.err
}

func (t *JwtToken) ParseInt(key string) (int, error) {
	val, ok := t.claims[key].(float64)
	if !ok {
		return 0, fmt.Errorf("Failed to parse token.")
	}
	return int(val), nil
}

func (t *JwtToken) ParseString(key string) (string, error) {
	val, ok := t.claims[key].(string)
	if !ok {
		return "", fmt.Errorf("Failed to parse token.")
	}
	return val, nil
}

func parse(tokenString string) (*jwt_go.Token, error) {
	return jwt_go.Parse(tokenString, func(token *jwt_go.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt_go.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(appKey), nil
	})
}

// https://zenn.dev/torkralle/articles/4e4d0f703d8122
func structToJsonTagMap(data interface{}) (map[string]interface{}, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	// JSONの整形
	err = json.Indent(out, jsonStr, "", "    ")
	if err != nil {
		return nil, err
	}
	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(out.String()), &mapData); err != nil {
		return nil, err
	}
	return mapData, err
}
