package comments

import (
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"anaco-go-backend/users"
	"anaco-go-backend/common"
)

func CommentsRegister(router *gin.RouterGroup) {
	router.POST("/", CommentCreate)
}

func CommentCreate(c *gin.Context) {
	commentModelValidator := NewCommentModelValidator()
	if err := commentModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	username := c.PostForm("username")
	toUserModel, err := users.FindOneUser(&users.UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("Invalid username")))
		return
	}
	commentModelValidator.commentModel.ToUser = toUserModel

	if err := SaveOne(&commentModelValidator.commentModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := CommentSerializer{c, commentModelValidator.commentModel}
	c.JSON(http.StatusCreated, gin.H{"comments": serializer.Response()})
}
