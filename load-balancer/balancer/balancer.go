package balancer

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Backend struct {
	URL   *url.URL
	Alive bool
	mu    sync.RWMutex
}

type LoadBalancer struct {
	backends []*Backend
	current  int
	mu       sync.Mutex
}

func NewLoadBalancer(backendUrls []string) *LoadBalancer {
	var backends []*Backend
	for _, rawURL := range backendUrls {
		parsedURL, _ := url.Parse(rawURL)
		backends = append(backends, &Backend{
			URL:   parsedURL,
			Alive: true,
		})
	}
	return &LoadBalancer{
		backends: backends,
	}
}

func (lb *LoadBalancer) HealthCheck(ctx context.Context, rate time.Duration) {
	ticker := time.NewTicker(rate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, server := range lb.backends {
				resp, err := http.Get(server.URL.String())
				if err != nil || resp.StatusCode >= 500 {
					server.SetAlive(false)
				} else {
					server.SetAlive(true)
				}
				if resp != nil {
					resp.Body.Close()
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (lb *LoadBalancer) NextBackend() *url.URL {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	n := len(lb.backends)
	for i := 0; i < n; i++ {
		lb.current = (lb.current + 1) % n
		b := lb.backends[lb.current]
		if b.IsAlive() {
			return b.URL
		}
	}

	return lb.backends[0].URL
}

func (b *Backend) IsAlive() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Alive
}

func (b *Backend) SetAlive(alive bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Alive = alive
}
