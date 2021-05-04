package app

import "github.com/dikanp/controllers/users"

func mapUrls() {
	router.GET("/ping", users.Tes)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users", users.GetAllUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}