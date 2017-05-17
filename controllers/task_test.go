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

type TaskTesting struct {
	task     TaskTest
	ID       string
	expected int
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
		//Printers: []httpexpect.Printer{
		//	httpexpect.NewDebugPrinter(s.T(), false),
		//},
	})
}

func (s *ExampleTestSuiteTask) TestPostTask() {

	s.handler.POST("/put", s.ctrl.PostTask)
	s.handler.PATCH("/:id", s.ctrl.DeleteTask)

	tasks := []TaskTesting{
		TaskTesting{
			TaskTest{
				"test",
			},
			"1",
			http.StatusOK,
		},
		TaskTesting{
			TaskTest{
				"test2",
			},
			"2",
			http.StatusOK,
		},
		TaskTesting{
			TaskTest{
				"",
			},
			"3",
			http.StatusBadRequest,
		},
	}

	for _, us := range tasks {
		s.expect.POST("/put").WithJSON(us.task).Expect().Status(us.expected)
	}
	for _, us := range tasks {
		s.expect.PATCH("/" + us.ID).Expect().Status(us.expected)
	}

}

func TestExampleTestSuiteTask(t *testing.T) {
	suite.Run(t, new(ExampleTestSuiteTask))
}

func (s *ExampleTestSuiteTask) TearDownSuite() {

}
