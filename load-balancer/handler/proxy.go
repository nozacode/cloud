package handler

import (
	"load-balancer/balancer"
	"log"
	"net/http"
	"net/http/httputil"
)

func NewProxyHandler(lb *balancer.LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		targetURL := lb.NextBackend()

		if targetURL == nil {
			http.Error(w, "Нет доступных серверов", http.StatusServiceUnavailable)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		log.Printf("Проксируем запрос на %s", targetURL.String())

		proxy.ServeHTTP(w, r)
	}
}
