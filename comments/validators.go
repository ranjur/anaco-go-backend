package comments

import (
	"gopkg.in/gin-gonic/gin.v1"
	"anaco-go-backend/users"
	"anaco-go-backend/common"
)

type CommentModelValidator struct {
	Comment struct {
		Body string `form:"body" json:"body" binding:"required,max=2048"`
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
	} `json:"comment"`
	commentModel CommentModel `json:"-"`
}

func NewCommentModelValidator() CommentModelValidator {
	return CommentModelValidator{}
}

func (s *CommentModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.commentModel.Body = s.Comment.Body
	s.commentModel.ByUser = myUserModel
	return nil
}
