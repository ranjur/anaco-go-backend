package users

import (
	"gopkg.in/gin-gonic/gin.v1"

	"anaco-go-backend/common"
)

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     string `json:"image"`
	Following bool    `json:"following"`
}

type UserSerializer struct {
	c *gin.Context

}

type ListUserSerializer struct {
	c *gin.Context
	UserModel

}

type UsersSerializer struct {
	C        *gin.Context
	Users []UserModel
}

type UserResponse struct {
	ID        uint    `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    string `json:"image"`
	Token    string  `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	myUserModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		ID: myUserModel.ID,
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}

type ListUserResponse struct {
	ID        uint    `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    string `json:"image"`
	CommentsCount uint  `json:"commentsCount"`
}

func (s *ListUserSerializer) Response() ListUserResponse {
	user := ListUserResponse{
		ID: s.ID,
		Username: s.Username,
		Email:    s.Email,
		Bio:      s.Bio,
		Image:    s.Image,
		CommentsCount:    s.commentsCount(),
	}
	return user
}

func (s *UsersSerializer) Response() [] ListUserResponse {
	response := [] ListUserResponse{}
	for _, user := range s.Users {
		serializer := ListUserSerializer{s.C, user}
		response = append(response, serializer.Response())
	}
	return response
}
