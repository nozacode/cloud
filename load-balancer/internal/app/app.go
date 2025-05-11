package app

type Config struct {
	Host string
}

func Run(cfg string) {
	// сборка и запуск приложения

	// init repo

	// init logger

	// init
}


type LoadBalancer struct {
	backends []*url.URL
	index    uint32 // для round-robin
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

type LoadBalancer struct {
	backends []*url.URL
	index    uint32 // для round-robin
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

type LoadBalancer struct {
	backends []*url.URL
	index    uint32 // для round-robin
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

type LoadBalancer struct {
	backends []*url.URL
	index    uint32 // для round-robin
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func newLoadBalancer(backendUrls []string) *LoadBalancer {
	var backends []*url.URL
	for _, backend := range backendUrls {
		parsedUrl, err := url.Parse(backend)
		if err != nil {
			log.Fatalf("Невозможно распарсить URL бэкенда %s: %v", backend, err)
		}
		backends = append(backends, parsedUrl)
	}

	return &LoadBalancer{
		backends: backends,
	}
}

func (lb *LoadBalancer) getNextBackend() *url.URL {
	index := atomic.AddUint32(&lb.index, 1)
	return lb.backends[int(index)%len(lb.backends)]
}

func (lb *LoadBalancer) handler(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextBackend()
	log.Printf("Перенаправляем запрос на: %s", target)

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		log.Printf("Ошибка при проксировании на %s: %v", target, err)
		http.Error(w, "Бэкенд недоступен", http.StatusBadGateway)
	}

	proxy.ServeHTTP(w, r)
}

// func main() {
// 	config, err := loadConfig("config.json")
// 	if err != nil {
// 		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
// 	}

// 	lb := newLoadBalancer(config.Backends)

// 	http.HandleFunc("/", lb.handler)

// 	log.Printf("Балансировщик запущен на порту %s...", config.Port)
// 	err = http.ListenAndServe(":"+config.Port, nil)
// 	if err != nil {
// 		log.Fatalf("Ошибка запуска сервера: %v", err)
// 	}
// }
