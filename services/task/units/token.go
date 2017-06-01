package units

import (
	pb "github.com/jaax2707/ToDoGorm/services/auth/protobuf"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"net/http"
)

// Client represents struct of grpc TokenClient interface
type Client struct {
	authClient pb.TokenClient
}

// NewClient represents func witch initialize Client struct
func NewClient(tokClient pb.TokenClient) *Client {
	return &Client{authClient: tokClient}
}

// CheckToken represents middleware,
// which check token in request header and compare it if is not expired in cache,
func (client *Client) CheckToken(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		t := c.Request().Header.Get("Authorization")
		_, err := client.authClient.GetToken(context.Background(), &pb.Request{Token: t})

		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}
