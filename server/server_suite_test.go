package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/labstack/echo"
	"strings"
	"encoding/json"

	. "github.com/newlee/petstore/server"
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

var huskyOrder = `{"petName":"Husky","guestName":"David","telephone":"123456","address":"address1"}`

var withOrder = func() {
	c, _ := requestWithBody("POST", "/orders", e, huskyOrder)
	Expect(c).To(Equal(http.StatusCreated))
}

var withHusky = func() {
	c, _ := requestWithBody("POST", "/pets", e, `{"name":"Husky"}`)
	Expect(c).To(Equal(http.StatusCreated))
}

type ResponseBody struct {
	body string
}

var whenGetAllPets = func() *ResponseBody {
	c, b := request("GET", "/pets", e)
	Expect(c).To(Equal(http.StatusOK))
	return &ResponseBody{body: b}
}

var whenGetPetByTelephone = func() *ResponseBody {
	c, b := request("GET", "/orders/123456", e)
	Expect(c).To(Equal(http.StatusOK))
	return &ResponseBody{body: b}
}

var whenGetPetByErrorTelephone = func() int {
	c, _ := request("GET", "/orders/error", e)
	return c;
}
var thenGetNotFound = func(c int) {
	Expect(c).To(Equal(http.StatusNotFound))
}
func (b *ResponseBody) isEmpty()  {
	Expect(b.body).To(Equal("[]"))
}

func (b *ResponseBody) haveHusky()  {
	pets := make([]Pet, 0)
	json.Unmarshal([]byte(b.body), &pets)
	Expect(len(pets)).To(Equal(1))
	Expect(pets[0].Name).To(Equal("Husky"))
}

func (b *ResponseBody) isOrderForHusky()  {
	Expect(b.body).To(Equal(huskyOrder))
}