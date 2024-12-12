package middleware

import (
	"net/http"

	"ecom/internal/apicollector"
)

type ApiCollectorMiddleware struct {
	apiCollector apicollector.ApiCollector
}

func NewApiCollectorMiddleware(apiCollector apicollector.ApiCollector) *ApiCollectorMiddleware {
	return &ApiCollectorMiddleware{
		apiCollector: apiCollector,
	}
}

func (m *ApiCollectorMiddleware) CollectInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiPath := r.Method + " " + r.URL.Path
		m.apiCollector.AddApiCall(apiPath)
		next.ServeHTTP(w, r)
	})
}
