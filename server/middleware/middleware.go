package middleware

import (
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const servicePort = "4040"

// SetServer sets various server requirement
func SetServer(r *mux.Router) *http.Server {
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) // Allowing all origin as of now

	allowedHeaders := handlers.AllowedHeaders([]string{
		"X-Requested-With",
		"X-CSRF-Token",
		"X-Auth-Token",
		"Content-Type",
		"processData",
		"contentType",
		"Origin",
		"Authorization",
		"Accept",
		"Client-Security-Token",
		"Accept-Encoding",
		"timezone",
		"locale",
	})

	allowedMethods := handlers.AllowedMethods([]string{
		"POST",
		"GET",
		"DELETE",
		"PUT",
		"PATCH",
		"OPTIONS"})

	allowCredential := handlers.AllowCredentials()

	s := &http.Server{
		Addr: ":" + servicePort,
		//	Handler: r,
		Handler: handlers.CORS(
			allowedHeaders,
			allowedMethods,
			allowedOrigins,
			allowCredential)(
			context.ClearHandler(
				r,
			)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	return s
}
