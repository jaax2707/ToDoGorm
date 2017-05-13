package units

import (
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"net/http"
)

type AuthorizationMiddleware struct {
	cache *cache.Cache
}

func NewCacheMiddleware(Cache *cache.Cache) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{cache: Cache}
}

func (m *AuthorizationMiddleware) CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Request().Header.Get("Authorization")
		_, exist := m.cache.Get(t)
		if !exist {
			return c.NoContent(http.StatusUnauthorized)
		}
		return next(c)
	}
}