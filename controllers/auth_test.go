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

type UserTesting struct {
	userTest UserTest
	expected int
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

func (s *ExampleTestSuiteAuth) TestLogin() {
	// status OK

	s.handler.POST("/", s.ctrl.Login)

	users := []UserTesting{
		UserTesting{
			UserTest{
				"test",
				"1111",
			},
			http.StatusOK,
		},
		UserTesting{
			UserTest{
				"test555",
				"test111125",
			},
			http.StatusUnauthorized,
		},
	}

	for _, us := range users {
		s.expect.POST("/").WithJSON(us.userTest).Expect().Status(us.expected)
	}
}

func (s *ExampleTestSuiteAuth) TestRegister() {
	// status MethodNotAllowed
	s.handler.POST("/", s.ctrl.Register)

	users := []UserTesting{
		UserTesting{
			UserTest{
				"test",
				"1111",
			},
			http.StatusMethodNotAllowed,
		},
		UserTesting{
			UserTest{
				"test555",
				"test111125",
			},
			http.StatusCreated,
		},
	}

	for _, us := range users {
		s.expect.POST("/").WithJSON(us.userTest).Expect().Status(us.expected)
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuiteAuth))
}

func (s *ExampleTestSuiteAuth) TearDownSuite() {

}
