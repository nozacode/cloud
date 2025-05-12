package balancer

import (
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type LoadBalancer struct {
	Backends []*Backend
	curr     int
	mutex    sync.Mutex
}

// NewRoundRobinBalancer инициализирует бэкенды
func NewLoadBalancer(backendUrls []string) *LoadBalancer {
	var backends []*Backend

	for _, rawUrl := range backendUrls {
		parsedUrl, err := url.Parse(rawUrl)
		if err == nil {
			backends = append(backends, &Backend{
				URL:   parsedUrl,
				Alive: true, // по умолчанию все живы
			})
		}
	}

	return &LoadBalancer{
		Backends: backends,
	}
}

// GetNextBackend реализует Round Robin
func (lb *LoadBalancer) GetNextBackend() *Backend {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	backendsCount := len(lb.Backends)
	for i := 0; i < backendsCount; i++ {
		backend := lb.Backends[lb.curr]
		lb.curr = (lb.curr + 1) % backendsCount
		if backend.IsAlive() {
			return backend
		}
	}
	return nil
}

// HealthCheck проверяет живы ли бэкенды
/*Эта функция проверяет, доступен ли каждый бэкенд.
Если нет ответа или ошибка — он считается "мертвым".*/
func (lb *LoadBalancer) HealthCheck(interval time.Duration) {
	for {
		for _, backend := range lb.Backends {
			go func(b *Backend) {
				resp, err := http.Get(b.URL.String())
				if err != nil || resp.StatusCode >= 500 {
					b.SetAlive(false)
					log.Printf("[HealthCheck] %s ➜ DEAD", b.URL)
					return
				}
				b.SetAlive(true)
				log.Printf("[HealthCheck] %s ➜ ALIVE", b.URL)
			}(backend)
		}
		time.Sleep(interval)
	}
}
