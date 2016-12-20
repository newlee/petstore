package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/labstack/echo"
	. "github.com/newlee/petstore/server"
	"net/http"
	"encoding/json"
)

var e *echo.Echo
var orderJson = `{"petName":"Husky","guestName":"David","telephone":"123456","address":"address1"}`

var _ = Describe("Server", func() {
	BeforeEach(func() {
		e = echo.New()
		Router(e);
	})

	Context("Home", func() {
		It("should ok", func() {
			c, _ := request("GET", "/", e)
			Expect(c).To(Equal(http.StatusOK))
		})
	})

	Context("Pets", func() {
		It("should return empty when have not any pets", func() {
			c, b := request("GET", "/pets", e)
			Expect(c).To(Equal(http.StatusOK))
			Expect(b).To(Equal("[]"))
		})

		It("should return husky after created.", func() {
			c, _ := requestWithBody("POST", "/pets", e, `{"name":"Husky"}`)
			Expect(c).To(Equal(http.StatusCreated))

			c, b := request("GET", "/pets", e)
			pets := make([]Pet, 0)
			json.Unmarshal([]byte(b), &pets)
			Expect(len(pets)).To(Equal(1))
			Expect(pets[0].Name).To(Equal("Husky"))
		})
	})

	Context("Order", func() {
		BeforeEach(func() {
			requestWithBody("POST", "/pets", e, `{"name":"Husky"}`)
		})

		It("should get order by telephone after created", func() {
			c, _ := requestWithBody("POST", "/orders", e, orderJson)
			Expect(c).To(Equal(http.StatusCreated))

			c, b := request("GET", "/orders/123456", e)
			Expect(b).To(Equal(orderJson))
		})

		It("should get 404 when telephone is error", func() {
			c, _ := request("GET", "/orders/error", e)
			Expect(c).To(Equal(http.StatusNotFound))
		})
	})
})
