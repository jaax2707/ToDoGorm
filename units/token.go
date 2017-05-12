package units

import (
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"net/http"
)

type CacheMiddleware struct {
	cache *cache.Cache
}

func NewCacheMiddleware(Cache *cache.Cache) *CacheMiddleware {
	return &CacheMiddleware{cache: Cache}
}

func (x *CacheMiddleware) CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Request().Header.Get("Authorization")
		_, exist := x.cache.Get(t)
		if !exist {
			return c.NoContent(http.StatusUnauthorized)
		}
		return next(c)
	}
}
