package server_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/labstack/echo"
	. "github.com/newlee/petstore/server"
)

var e *echo.Echo

var _ = Describe("Server", func() {
	BeforeEach(func() {
		e = echo.New()
		Router(e);
	})

	Context("Pets", func() {
		It("should return empty when have not any pets", func() {
			whenGetAllPets().isEmpty()
		})

		It("should return husky after created.", func() {
			withHusky()
			whenGetAllPets().haveHusky()
		})
	})

	Context("Order", func() {
		BeforeEach(func() {
			withHusky()
		})

		It("should get order by telephone after created", func() {
			withOrder()
			whenGetPetByTelephone().isOrderForHusky()
		})

		It("should update order when exist", func() {

		})

		It("should get 404 when telephone is error", func() {
			c := whenGetPetByErrorTelephone()
			thenGetNotFound(c)
		})
	})
})
