package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	ID   uint   `json:"id"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

var people []Person
var db *gorm.DB
var err error

func Login(user, pass string) bool {
	OpenConnection()
	AutoMig()
	var customer Person
	err = db.Where("user = ?", user).First(&customer).Error
	CloseConnection()
	if err != nil || customer.Pass != pass {
		return false
	} else {
		return true
	}
}

func OpenConnection() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
}

func AutoMig() {
	db.AutoMigrate(&Person{})
}

func CloseConnection() {
	db.Close()
}

// func Register(user string, pass string) {
// 	// user := c.Params.ByName("user")
// 	// pass := c.Params.ByName("pass")

// 	var customer Person
// 	err1 := db.Where("user = ?", user).First(&customer).Error

// 	if err1 != nil || customer.Pass != pass {
// 		fmt.Println("Dang nhap that bai")
// 	} else {
// 		fmt.Println("Dang nhap thanh cong")
// 	}
// }

// func GetPeople(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var people Person
// 	if err := db.Where("id = ?", id).First(&people).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, people)
// 	}

// }

// func GetAllPeople(c *gin.Context) {
// 	var people []Person
// 	if err := db.Find(&people).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, people)
// 	}
// }

// func CreatePeople(c *gin.Context) {
// 	var people Person
// 	c.BindJSON(&people)

// 	db.Create(&people)
// 	c.JSON(200, people)
// }

// func UpdatePeople(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var people Person
// 	if err := db.Where("id = ?", id).First(&people).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.BindJSON(&people)
// 		db.Save(&people)
// 		c.JSON(200, people)
// 	}
// }

// func DeletePeople(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var people Person
// 	d := db.Where("id = ?", id).Delete(&people)
// 	fmt.Println(d)
// 	c.JSON(200, gin.H{"id #" + id: "deleted"})
// }
