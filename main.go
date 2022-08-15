package main

import (
	"monta-channel/app/controllers"
)

func main() {
	// fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUser(2)
	// fmt.Println(user)

	// models.CreateThread("毎日腹筋1000回した結果wwwwwwwww")

	// user.CreateContent("結論口臭い", 7)
	// fmt.Println(user)

	// contents, _ := models.GetContentsByThreadID(7)
	// for _, v := range contents {
	// 	v.Content = "修正版 口臭い"
	// 	v.UpdateContent()
	// }

	// contents, _ = models.GetContentsByThreadID(7)
	// for _, v := range contents {
	// 	fmt.Println(v)
	// }

	// fmt.Println("====================================")

	// con, _ := models.GetContent(17)
	// fmt.Println(con)

	// con.DeleteContent()

	// t, _ := models.GetThread(1)
	// fmt.Println(t)

	// t.DeleteThread()

	// contents, _ = models.GetContentsByThreadID(1)
	// for _, v := range contents {
	// 	fmt.Println(v)
	// }
}
