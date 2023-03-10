package routes

import (
	"net/http"
	"time"

	g "service/gateway/global"
	httpHandler "service/gateway/handlers/http"
	m "service/gateway/middleware"

	"service/pkg/router"
)

func basicMiddlewares(next http.Handler, methods ...string) http.Handler {
	return m.Translator(m.Panic(m.Timeout(time.Duration(g.CFG.Timeout), m.ConcurrentLimiter(g.CFG.MaxConcurrentRequests, m.Json(m.Cors(m.Method(next, methods...)))))))
}

func HTTP(mux *router.Router) {
	// /api
	{
		mux.Handle("/api/me/", basicMiddlewares(httpHandler.Me, "GET"))
		// /api/auth
		{
			mux.Handle("/api/auth/register/", basicMiddlewares(httpHandler.Register, "POST"))
			mux.Handle("/api/auth/login/", basicMiddlewares(httpHandler.Login, "POST"))
			mux.Handle("/api/auth/refresh/", basicMiddlewares(httpHandler.Refresh, "POST"))
		}
		mux.Handle("/api/.+/", basicMiddlewares(httpHandler.Hi, "GET"))
	}

	// Not found
	mux.Handle("/", basicMiddlewares(httpHandler.NotFound))
	mux.Handle("/.*/", basicMiddlewares(httpHandler.NotFound))
}
