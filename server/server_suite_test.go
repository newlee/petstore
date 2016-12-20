package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/labstack/echo"
	"strings"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

func requestWithBody(method, path string, e *echo.Echo, body string) (int, string) {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

func request(method, path string, e *echo.Echo) (int, string) {
	return requestWithBody(method, path, e, "")
}