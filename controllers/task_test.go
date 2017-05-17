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

type ExampleTestSuiteTask struct {
	suite.Suite
	handler *echo.Echo
	expect  *httpexpect.Expect
	ctrl    *Task
}

type TaskTest struct {
	Name string `json:"name"`
}

func (s *ExampleTestSuiteTask) SetupTest() {

	s.handler = echo.New()

	s.ctrl = NewTask(access.NewTaskAccessMock(), cache.New(10*time.Minute, 10*time.Minute))

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

type TaskTesting struct {
	taskTest TaskTest
	expected int
}

func (s *ExampleTestSuiteTask) TestGetAll() {
	// status OK

	s.handler.GET("/", s.ctrl.GetAll)

	users := []TaskTesting{
		TaskTesting{
			TaskTest{
				"test",
			},
			http.StatusOK,
		},
	}

	for _, us := range users {
		s.expect.GET("/").Expect().Status(us.expected)
	}
}

func (s *ExampleTestSuiteTask) TestPostTask() {
	// status MethodNotAllowed
	s.handler.POST("/", s.ctrl.PostTask)

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
				"test222",
				"2222",
			},
			http.StatusCreated,
		},
	}

	for _, us := range users {
		s.expect.POST("/").WithJSON(us.userTest).Expect().Status(us.expected)
	}
}

//func (s *ExampleTestSuiteTask) TestDeleteTask() {
//	// status MethodNotAllowed
//	s.handler.POST("/", s.ctrl.DeleteTask)
//
//	users := []UserTesting{
//		UserTesting{
//			UserTest{
//				"test",
//				"1111",
//			},
//			http.StatusMethodNotAllowed,
//		},
//		UserTesting{
//			UserTest{
//				"test222",
//				"2222",
//			},
//			http.StatusCreated,
//		},
//	}
//
//	for _, us := range users {
//		s.expect.POST("/").WithJSON(us.userTest).Expect().Status(us.expected)
//	}
//}

func TestExampleTestSuiteTask(t *testing.T) {
	suite.Run(t, new(ExampleTestSuiteTask))
}

func (s *ExampleTestSuiteTask) TearDownSuite() {

}
