package main

import (
	"monta-channel/app/controllers"
)

func main() {
	// threads, _ := models.GetThreadsByCategory(1)
	// log.Println(threads)

	controllers.StartMainServer()
}
