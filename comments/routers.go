package comments

import (
	"errors"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"anaco-go-backend/users"
	"anaco-go-backend/common"
	"strconv"
)

func CommentsRegister(router *gin.RouterGroup) {
	router.POST("/:username", CommentCreate)
	router.GET("/:username", UserCommentList)

}

func CommentRegister(router *gin.RouterGroup) {
	router.POST("/:comment_id", CreateCommentLike)

}

func CommentCreate(c *gin.Context) {
	commentModelValidator := NewCommentModelValidator()
	if err := commentModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	username := c.Param("username")
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


func UserCommentList(c *gin.Context) {
	username := c.Param("username")
	UserModel, err := users.FindOneUser(&users.UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("Invalid username")))
		return
	}
	commentModels, err := GetUserComments(UserModel)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("Database error")))
		return
	}
	serializer := CommentsSerializer{c, commentModels}
	c.JSON(http.StatusOK, gin.H{"comments": serializer.Response()})
}

func CreateCommentLike(c *gin.Context) {
	commentID := c.Param("comment_id")
	ID, err := strconv.ParseUint(commentID,10, 32)
	thisComment, err := FindOneComment(uint(ID))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", err))
		return
	}
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	err = thisComment.like(myUserModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success","liked": true })
}