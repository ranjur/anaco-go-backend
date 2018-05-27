package comments

import (
	"gopkg.in/gin-gonic/gin.v1"
	"anaco-go-backend/users"
)

type CommentSerializer struct {
	C *gin.Context
	CommentModel
}

type CommentsSerializer struct {
	C        *gin.Context
	Comments []CommentModel
}

type CommentResponse struct {
	ID        uint                  `json:"id"`
	Body      string                `json:"body"`
	ToUserID  uint            `json:"to_user_id"`
	CreatedAt string                `json:"createdAt"`
	UpdatedAt string                `json:"updatedAt"`
	IsLiked bool                `json:"isLiked"`
}

func (s *CommentSerializer) Response() CommentResponse {
	myUserModel := s.C.MustGet("my_user_model").(users.UserModel)
	response := CommentResponse{
		ID:        s.ID,
		Body:      s.Body,
		ToUserID:    s.ToUserID,
		CreatedAt: s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		IsLiked: s.isLiked(myUserModel),
	}
	return response
}

func (s *CommentsSerializer) Response() []CommentResponse {
	response := []CommentResponse{}
	for _, comment := range s.Comments {
		serializer := CommentSerializer{s.C, comment}
		response = append(response, serializer.Response())
	}
	return response
}
