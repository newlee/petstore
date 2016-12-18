package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/labstack/echo"
	. "github.com/newlee/petstore/server"
	"net/http"
)

var _ = Describe("Server", func() {
	Context("Home", func() {
		It("should ok", func() {
			e := echo.New()
			Router(e);
			c, _ := request("GET", "/", e)
			Expect(c, http.StatusOK)
		})
	})
})
