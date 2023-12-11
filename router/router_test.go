package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go-test-demo/app"
	"go-test-demo/db"
	"go-test-demo/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	g *gin.Engine
)

func init() {
	g = gin.Default()
	InitRouter(g)
	app.InitRepository(db.DB)
}

func TestInitRouter(t *testing.T) {
	po := url.Values{"name": {"test"}}
	req := httptest.NewRequest(http.MethodPost, "/v1/from", strings.NewReader(po.Encode()))
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	g.ServeHTTP(rec, req)
	s := rec.Body.String()
	assert.Equal(t, rec.Result().StatusCode, 200)
	t.Log(s)
}

func TestUser(t *testing.T) {
	u := &models.User{Name: "test", Age: 18, Phone: "123456"}
	marshal, _ := json.Marshal(u)
	bodys := string(marshal)

	req := httptest.NewRequest(http.MethodPost, "/v1/user", strings.NewReader(bodys))
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	g.ServeHTTP(rec, req)
	assert.Equal(t, rec.Result().StatusCode, http.StatusOK)
	t.Logf(rec.Body.String())
}

func TestTicket(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/ticket/info/1", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	g.ServeHTTP(rec, req)
	assert.Equal(t, rec.Result().StatusCode, http.StatusOK)
	t.Logf(rec.Body.String())
}
