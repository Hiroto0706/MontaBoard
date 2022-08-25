package main

import "monta-channel/app/controllers"

func main() {
	// threads, _ := models.GetPopularThreadsLimitFive()
	// fmt.Println(threads)
	controllers.StartMainServer()
}
