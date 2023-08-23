package main

import "github.com/KampungBudaya/Kampung-Budaya-2023-BE/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
