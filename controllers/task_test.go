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

	s.handler.POST("/", s.ctrl.PostTask)

	users := []TaskTesting{
		TaskTesting{
			TaskTest{
				"test",
			},
			"1",
			http.StatusCreated,
		},
		TaskTesting{
			TaskTest{
				"",
			},
			"2",
			http.StatusBadRequest,
		},
	}

	for _, us := range users {
		s.expect.POST("/").WithJSON(us.task).Expect().Status(us.expected)
	}
}

func (s *ExampleTestSuiteTask) TestDeleteTask() {
	s.handler.PATCH("/:id", s.ctrl.DeleteTask)

	tasks := []TaskTesting{
		TaskTesting{
			task: TaskTest{
				"task1",
			},
			ID:       "3",
			expected: http.StatusBadRequest,
		},
		TaskTesting{
			task: TaskTest{
				"task2",
			},
			ID:       "7",
			expected: http.StatusOK,
		},
	}

	for _, t := range tasks {
		s.expect.PATCH("/" + t.ID).Expect().Status(t.expected)
	}
}

func TestExampleTestSuiteTask(t *testing.T) {
	suite.Run(t, new(ExampleTestSuiteTask))
}

func (s *ExampleTestSuiteTask) TearDownSuite() {

}
