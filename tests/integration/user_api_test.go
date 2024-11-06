package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tanvirislam06/go-api-test/internal/user"
	"github.com/tanvirislam06/go-api-test/pkg/routes"
)

var _ = Describe("User API Integration Tests", func() {

	var server *httptest.Server

	BeforeEach(func() {
		repository := user.NewRepository()
		service := user.NewUserService(repository)
		handler := user.NewHandler(*service)

		route := routes.SetupRouter(handler)
		server = httptest.NewServer(route)
	})

	AfterEach(func() {
		server.Close()
	})

	It("should return a valid respone", func() {
		resp, err := http.Get(server.URL + "/user?id=1")
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		var userReponse user.User

		json.NewDecoder(resp.Body).Decode(&userReponse)
		Expect(userReponse.ID).To(Equal(1))
		Expect(userReponse.Name).To(Equal("John"))
		Expect(userReponse.Email).To(Equal("john@example.com"))
	})

	It("should return an invalid respone", func() {
		resp, err := http.Get(server.URL + "/user?id=3")
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
	})
})
