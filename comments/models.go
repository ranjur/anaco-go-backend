package comments

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	"anaco-go-backend/users"
	"anaco-go-backend/common"
)

type CommentModel struct {
	gorm.Model
	ByUser   users.UserModel
	ByUserID uint
	ToUser   users.UserModel
	ToUserID uint
	Body      string `gorm:"size:2048"`
}

func GetUserComments(UserModel users.UserModel) ([]CommentModel, error) {
	db := common.GetDB()
	var models []CommentModel

	tx := db.Begin()

	tx.Where("to_user_id=?", UserModel.ID).Order("updated_at desc").Find(&models)
	var err = tx.Commit().Error
	return models, err
}


// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&CommentModel{})
	db.AutoMigrate(&CommentLikeModel{})
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneComment(commentID uint) (CommentModel, error) {
	db := common.GetDB()
	var model CommentModel

	tx := db.Begin()

	tx.Where("id=?", commentID).First(&model)
	var err = tx.Commit().Error
	return model, err
}



func (c CommentModel) like(u users.UserModel) error {
	db := common.GetDB()
	var like CommentLikeModel
	err := db.FirstOrCreate(&like, &CommentLikeModel{
		UserID:  u.ID,
		CommentID: c.ID,
	}).Error
	return err
}

func (c CommentModel) dislike(u users.UserModel) error {
	db := common.GetDB()
	err := db.Where(CommentLikeModel{
		UserID:  u.ID,
		CommentID: c.ID,
	}).Delete(CommentLikeModel{}).Error
	return err
}

type CommentLikeModel struct {
	gorm.Model
	User    users.UserModel
	UserID  uint
	Comment   CommentModel
	CommentID uint
}

func (c CommentModel) isLiked(u users.UserModel) bool {
	db := common.GetDB()
	var comment CommentLikeModel
	db.Where(CommentLikeModel{
		UserID:  u.ID,
		CommentID: c.ID,
	}).First(&comment)
	return comment.ID != 0
}