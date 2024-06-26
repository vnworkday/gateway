package main

import (
	"flag"
	"os"

	app "github.com/vnworkday/gateway/cmd/gateway"
	_ "github.com/vnworkday/gateway/docs"
)

// @title						VN Workday Gateway API
// @version					0.1.0
// @description				This is the API documentation for VN Workday Gateway API.
// @BasePath					/api/v1
// @schemes					http https
// @host						http://localhost:3000
// @produces					application/json
// @contact.name				Duy Nguyen
// @contact.email				ntduy.cs@gmail.com
// @contact.url				https://github.com/vnworkday/gateway
// @license.name				MIT
// @license.url				https://opensource.org/licenses/MIT
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						Authorization
// @description				Please provide a valid JWT token in the header.
// @securityDefinitions.apiKey	ApiKey
// @in							header
// @name						x-api-key
// @description				Please provide a valid API key in the header.
func main() {
	var profile string

	flag.StringVar(&profile, "profile", "local", "Profile to run the application")
	flag.Parse()

	err := os.Setenv("PROFILE", profile)
	if err != nil {
		panic(err)
	}

	app.Start()
}
