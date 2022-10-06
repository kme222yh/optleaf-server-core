/*
interface

testtp.Router(r)

testtp.New("/hoge").Get().AssertStatusCode(t, 200)
testtp.New("/hoge").Post().AssertStatusCode(t, 400)

params := struct {
	email `json:"email"`,
	name `json:"name"`,
} {
	"test@hoge.com"
	"hogehoge"
}
testtp.New("/create/json").Post(params).AssertStatusCode(t, 201)

params := struct {
	email `url:"email"`,
	name `url:"name"`,
} {
	"test@hoge.com"
	"hogehoge"
}
testtp.New("/create/form").PostForm(params).AssertStatusCode(t, 201)

testtp.New("/security/hoge").Bearer(token).Get().AssertStatusCode(t, 200).Decode(&params)


*/

package testtp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
)

var (
	router *gin.Engine
)

type testInstance struct {
	url    string
	bearer string
	w      *httptest.ResponseRecorder
}

func Router(r *gin.Engine) {
	router = r
}

func New(url string) *testInstance {
	m := new(testInstance)
	m.url = url
	m.w = httptest.NewRecorder()
	return m
}

func (m testInstance) Get() *testInstance {
	req, _ := http.NewRequest(http.MethodGet, m.url, nil)
	if m.bearer != "" {
		req.Header.Set("Authorization", "Bearer "+m.bearer)
	}
	router.ServeHTTP(m.w, req)
	return &m
}

func (m testInstance) Delete() *testInstance {
	req, _ := http.NewRequest(http.MethodDelete, m.url, nil)
	if m.bearer != "" {
		req.Header.Set("Authorization", "Bearer "+m.bearer)
	}
	router.ServeHTTP(m.w, req)
	return &m
}

func (m testInstance) Post(params interface{}) *testInstance {
	var req *http.Request
	if params != nil {
		jsonStr, _ := json.Marshal(params)
		req, _ = http.NewRequest(http.MethodPost, m.url, bytes.NewBuffer([]byte(jsonStr)))
	} else {
		req, _ = http.NewRequest(http.MethodPost, m.url, nil)
	}
	req.Header.Set("Content-Type", "application/json")
	if m.bearer != "" {
		req.Header.Set("Authorization", "Bearer "+m.bearer)
	}
	router.ServeHTTP(m.w, req)
	return &m
}

func (m testInstance) Put(params interface{}) *testInstance {
	var req *http.Request
	if params != nil {
		jsonStr, _ := json.Marshal(params)
		req, _ = http.NewRequest(http.MethodPut, m.url, bytes.NewBuffer([]byte(jsonStr)))
	} else {
		req, _ = http.NewRequest(http.MethodPut, m.url, nil)
	}
	req.Header.Set("Content-Type", "application/json")
	if m.bearer != "" {
		req.Header.Set("Authorization", "Bearer "+m.bearer)
	}
	router.ServeHTTP(m.w, req)
	return &m
}

func (m testInstance) PostForm(params interface{}) *testInstance {
	var req *http.Request
	if params != nil {
		v, _ := query.Values(params)
		req, _ = http.NewRequest(http.MethodPost, m.url, strings.NewReader(v.Encode()))
	} else {
		req, _ = http.NewRequest(http.MethodPost, m.url, nil)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if m.bearer != "" {
		req.Header.Set("Authorization", "Bearer "+m.bearer)
	}
	router.ServeHTTP(m.w, req)
	return &m
}

func (m testInstance) Bearer(bearer string) *testInstance {
	m.bearer = bearer
	return &m
}

func (m testInstance) AssertStatusCode(t *testing.T, statusCode int) *testInstance {
	assert.Equal(t, statusCode, m.w.Code)
	return &m
}

func (m testInstance) Decode(params any) *testInstance {
	decoder := json.NewDecoder(m.w.Body)
	decoder.Decode(&params)
	return &m
}
