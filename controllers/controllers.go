package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"../models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		user := r.Form["username"]
		pass := r.Form["password"]
		fmt.Println("***********************************************", user)
		if models.Login(user[0], pass[0]) {
			fmt.Println("***************Dang nhap thanh cong")
		} else {
			fmt.Println("***************Dang nhap that bai")
		}
		fmt.Println("***********************************************", user)
	}
}

// func Login(user, pass string) bool {
// 	if models.Login(user, pass) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func Register(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
//     for _, item := range models.people {
// 		if item.user == params["user"] && item.pass == params["pass"] {
// 			fmt.Println("thanh cong")
// 		} else {
// 			fmt.Println("that bai")
// 		}
// }

// func GetPeople(c *gin.Context) {
// 	models.GetPeople(c)
// }

// func GetAllPeople(c *gin.Context) {
// 	models.GetAllPeople(c)
// }

// func CreatePeople(c *gin.Context) {
// 	models.CreatePeople(c)
// }

// func UpdatePeople(c *gin.Context) {
// 	models.UpdatePeople(c)
// }

// func DeletePeople(c *gin.Context) {
// 	models.DeletePeople(c)
// }
