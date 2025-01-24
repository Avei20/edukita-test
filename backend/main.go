package main

import "backend/internal/app"

func main() {
	server := app.InitHttp()
	server.InitRouteAndServe()
}
