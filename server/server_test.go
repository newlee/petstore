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

	Context("Products", func() {
		It("should return empty when have not any product", func() {
			c, b := request("GET", "/products", e)
			Expect(c).To(Equal(http.StatusOK))
			Expect(b).To(Equal("[]"))
		})

		It("should return product after created.", func() {
			c, _ := requestWithBody("POST", "/products", e, `{"name":"Husky"}`)
			Expect(c).To(Equal(http.StatusCreated))

			c, b := request("GET", "/products", e)
			products := make([]Product, 0)
			json.Unmarshal([]byte(b),&products)
			Expect(len(products)).To(Equal(1))
			Expect(products[0].Name).To(Equal("Husky"))
		})
	})
})
