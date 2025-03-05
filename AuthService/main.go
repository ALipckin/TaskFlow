package main

import (
	"AuthService/controllers"
	"AuthService/initializers"
	"AuthService/middleware"
	"fmt"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.SignUp)
	http.HandleFunc("/validate", controllers.Validate)

	http.Handle("/getUserInfo", middleware.RequireAuth(http.HandlerFunc(controllers.UserInfoHandler)))

	fmt.Println("Server started on :8081")
	http.ListenAndServe(":8081", nil)
}
