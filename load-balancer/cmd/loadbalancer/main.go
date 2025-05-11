// main.go
package main

import (
	"encoding/json"
	"load-balancer/internal/app"g"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync/atomic"
)

const (
	configPath = "./config/exampleconfig.json"
)

func main() {
	app.Run(configPath)
}





