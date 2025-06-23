package kongsy

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start(to string, port, limit, interval int) error {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(LimitByRealIP(limit, time.Duration(interval)*time.Second))
	router.Use(middleware.Logger)

	remote, err := url.Parse(to)
	if err != nil {
		return fmt.Errorf("failed to parse remote URL: %w", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		r.Host = remote.Host
		proxy.ServeHTTP(w, r)
	})

	slog.Info(
		fmt.Sprintf("Starting server on :%d", port),
		slog.String("to", to),
		slog.Int("limit", limit),
		slog.Int("interval", interval),
	)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		router,
	)
}
