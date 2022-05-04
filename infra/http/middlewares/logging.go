package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/rellyson/car-sales-api/application/utils"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	logger := utils.NewLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("[%v] %v %v - UserAgent: %v, RemoteAddr: %v",
			time.Now().Format(time.RFC1123Z), r.Method, r.RequestURI, r.UserAgent(), r.RemoteAddr)

		logger.Custom(message, color.FgCyan)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
