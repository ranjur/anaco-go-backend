package comments

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	"anaco-go-backend/users"
	"anaco-go-backend/common"
)

type CommentModel struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	ByUser   users.UserModel
	ByUserID uint
	ToUser   users.UserModel
	ToUserID uint
	Body      string `gorm:"size:2048"`
}

type CommentLikeModel struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	User    users.UserModel
	UserID  uint
	Comment   CommentModel
	CommentID uint
}

func GetUserComments(userModel users.UserModel) ([]CommentModel, error) {
	db := common.GetDB()
	var models []CommentModel
	err := db.Where("to_user_id=?", userModel.ID).Order("updated_at desc").Find(&models).Error
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

func FindOneComment(condition interface{}) (CommentModel, error) {
	db := common.GetDB()
	var model CommentModel
	err := db.Where(condition).First(&model).Error
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