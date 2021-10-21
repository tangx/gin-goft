package adaptors

import (
	"testing"

	"github.com/tangx-labs/gin-goft/cmd/demo/models"
)

func Test_NewUser(t *testing.T) {
	gorm := NewGormAdaptor()

	datas := map[int]string{
		1: "zhangsan",
		2: "lisi",
		3: "wangwu",
	}

	gorm.AutoMigrate(&models.User{})

	for id, name := range datas {
		user := models.User{
			UserId:   id,
			UserName: name,
		}

		gorm.Create(&user)

	}
}
