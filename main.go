package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/jinzhu/gorm"
	"anaco-go-backend/common"
	"anaco-go-backend/users"
	"anaco-go-backend/comments"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	comments.AutoMigrate()
}

func main() {

	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default())

	r.Static("/media", "./media")

	v1 := r.Group("/api")

	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))

	comments.CommentsRegister(v1.Group("/comments"))
	comments.CommentRegister(v1.Group("/comment"))

	v1.Use(users.AuthMiddleware(true))
	users.UserLIst(v1.Group("/all-users"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// test 1 to 1
	tx1 := db.Begin()
	userA := users.UserModel{
		Username: "AAAAAAAAAAAAAAAA",
		Email:    "aaaa@g.cn",
		Bio:      "hehddeda",
		Image:    "",
	}
	tx1.Save(&userA)
	tx1.Commit()
	fmt.Println(userA)

	//db.Save(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA ArticleUserModel
	//db.Where(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)

	r.Run() // listen and serve on 0.0.0.0:8080
}
