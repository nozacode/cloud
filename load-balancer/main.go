package main

import (
	"context"
	"load-balancer/balancer"
	"load-balancer/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}
	lb := balancer.NewLoadBalancer(backends)

	go lb.HealthCheck(ctx, 10*time.Second)

	http.HandleFunc("/", handler.NewProxyHandler(lb))

	log.Println("Load Balancer запущен на порту :8080")

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	sigReceived := <-signalChan
	log.Printf("Получен сигнал: %v, завершение работы...", sigReceived)

	cancel()

	log.Println("Балансировщик успешно завершил работу.")
}
