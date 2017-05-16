package controllers

import (
	"github.com/gavv/httpexpect"
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

type ExampleTestSuiteAuth struct {
	suite.Suite
	handler *echo.Echo
	expect  *httpexpect.Expect
	ctrl    *Auth
}

type UserTest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *ExampleTestSuiteAuth) SetupTest() {

	s.handler = echo.New()

	s.ctrl = NewAuth(access.NewAuthAccessMock(), cache.New(10*time.Minute, 10*time.Minute))

	s.expect = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(s.handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(s.T()),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(s.T(), true),
		},
	})
}

func (s *ExampleTestSuiteAuth) TestRegisterSuccess() {
	// status Created
	s.handler.POST("/", s.ctrl.Register)
	s.expect.POST("/").WithJSON(UserTest{"test11", "1111"}).
		Expect().
		Status(http.StatusCreated)
}

func (s *ExampleTestSuiteAuth) TestRegisterFailed() {
	// status MethodNotAllowed
	s.handler.POST("/", s.ctrl.Register)
	s.expect.POST("/").WithJSON(UserTest{"test", "1111"}).
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func (s *ExampleTestSuiteAuth) TestLoginSuccess() {
	// status OK
	s.handler.POST("/", s.ctrl.Login)
	s.expect.POST("/").WithJSON(UserTest{"test", "1111"}).
		Expect().
		Status(http.StatusOK)
}

func (s *ExampleTestSuiteAuth) TestLoginFailed() {
	// status Unauthorized
	s.handler.POST("/", s.ctrl.Login)
	s.expect.POST("/").WithJSON(UserTest{"test11", "1111"}).
		Expect().
		Status(http.StatusUnauthorized)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuiteAuth))
}
func (s *ExampleTestSuiteAuth) TearDownSuite() {

}
