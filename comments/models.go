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
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}