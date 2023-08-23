package main

import "github.com/KampungBudaya/Kampung-Budaya-2023-BE/app"

// @title Kampung Budaya's API
// @version 1.0

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
