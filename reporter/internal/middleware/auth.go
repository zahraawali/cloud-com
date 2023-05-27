package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(authServiceBase string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the token is present in the request header
		token := c.GetHeader("Authorization")

		// make a request to validate token
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/user/validate", authServiceBase), nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorSchema{
				Code:    http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}
		req.Header.Add("Authorization", token)

		cli := &http.Client{}
		resp, err := cli.Do(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorSchema{
				Code:    http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}

		if resp != nil && resp.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorSchema{
				Code:    http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}

		c.Next()
	}
}

type ErrorSchema struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
