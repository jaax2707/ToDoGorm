package units

import (
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"net/http"
)

// AuthorizationMiddleware represents struct of cache
type AuthorizationMiddleware struct {
	cache *cache.Cache
}

// NewCacheMiddleware return AuthorizationMiddleware Object
func NewCacheMiddleware(Cache *cache.Cache) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{cache: Cache}
}

// CheckToken represents middleware,
// which check token in request header and compare it if is not expired in cache,
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
