package main

import (
	"fmt"
	"monta-channel/app/controllers"
	"monta-channel/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
