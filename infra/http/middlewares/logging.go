package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/rellyson/car-sales-api/infra/utils"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	colorizedLog := utils.NewColorizedLogger(color.FgCyan)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("[%v] %v %v - UserAgent: %v, RemoteAddr: %v",
			time.Now().Format(time.RFC1123Z), r.Method, r.RequestURI, r.UserAgent(), r.RemoteAddr)

		colorizedLog.Println(message)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
