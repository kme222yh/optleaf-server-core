// テストやる時は同ディレクトリに.env作れ
/* .env

app.key=secret

*/

package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummary(t *testing.T) {

	testClaims := struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Hoge  string `json:"hoge"`
	}{
		1,
		"user@test.com",
		"2um49cum2t89304t5ynm2-4jm2-",
	}

	token := New().For("test").ValidityDays(2).Create(testClaims)
	assert.NoError(t, token.Error())
	tokenString := token.String()

	token = New().For("hoge").Verify(tokenString)
	assert.Error(t, token.Error())
	token = New().For("test").Verify(tokenString)
	assert.NoError(t, token.Error())

	id, err := token.ParseInt("id")
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	email, err := token.ParseString("email")
	assert.NoError(t, err)
	assert.Equal(t, "user@test.com", email)

	_, err = token.ParseInt("email")
	assert.Error(t, err)

	_, err = token.ParseString("aaaaaa")
	assert.Error(t, err)

	token = New().For("test").ValidityDays(-1).Create(testClaims)
	assert.NoError(t, token.Error())
	tokenString = token.String()
	token = New().For("test").Verify(tokenString)
	assert.Error(t, token.Error())

	token = New().Create(testClaims)
	assert.NoError(t, token.Error())
	tokenString = token.String()
	token = New().Verify(tokenString)
	assert.NoError(t, token.Error())

}
