// Hardened the HTTP security headers
// based on https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html
// and on https://helmetjs.github.io/#reference

package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

var headersToSet = map[string]string{
	"X-Frame-Options":            "DENY",
	"X-Content-Type-Options":     "nosniff",
	"X-XSS-Protection":           "0",
	"Referrer-Policy":            "strict-origin-when-cross-origin",
	"Content-Type":               "application/json; charset=UTF-8",
	"Cross-Origin-Opener-Policy": "no-referrer",
	"Server":                     "ENSSATV Software Technology",
	"Origin-Agent-Cluster":       "?1",
}

func HandleSecurity(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		for k, v := range headersToSet {
			c.Header(k, v)
		}
		c.Next()
	}
}
