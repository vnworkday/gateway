package main

import (
	app "github.com/vnworkday/gateway/cmd/gateway"
	_ "github.com/vnworkday/gateway/docs"
)

// @title			VN Workday Gateway API
// @version		0.1.0
// @description	This is the API documentation for VN Workday Gateway API.
// @BasePath		/api/v1
// @schemes		http https
// @host			localhost:3000
// @produces		application/json
// @contact.name	Duy Nguyen
// @contact.email	ntduy.cs@gmail.com
// @contact.url	https://github.com/vnworkday/gateway
func main() {
	app.Start()
}
