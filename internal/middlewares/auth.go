package middlewares

import (
	"SpotSync/internal/auth"
	"SpotSync/internal/httpresponse"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(jwtService auth.JWTService, roles ...string  ) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			// extract token from authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized,httpresponse.Error{
					Success: false,
					Code:    http.StatusUnauthorized,
					Message: "unauthorized",
					Details: "Missing authorization header",
				})
			}

			// check bearer scheme
			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized,httpresponse.Error{
					Success: false,
					Code:    http.StatusUnauthorized,
					Message: "unauthorized",
					Details: "invalid authorization header format",
				})
			}

			tokenString := parts[1]

			// validate token

			claims, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized,httpresponse.Error{
					Success: false,
					Code:    http.StatusUnauthorized,
					Message: "unauthorized",
					Details: "invalid or expired token",
				})
			}

		if len(roles) > 0 {
		allowed := false

		for _, r := range roles {
		if claims.Role == r {
			allowed = true
			break
		}
			}

		if !allowed {
		return c.JSON(http.StatusForbidden,httpresponse.Error{
			Success: false,
			Code:    http.StatusForbidden,
			Message: "forbidden",
			Details: "you do not have permission to access this resource",
		})
		}
		}

			// store user info in context for handlers
			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("user_name", claims.Name)
			c.Set("user_role", claims.Role)

			return next(c)
		}
	}
}
