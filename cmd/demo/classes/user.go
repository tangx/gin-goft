package classes

import (
	"github.com/tangx-labs/gin-goft/cmd/demo/adaptors"
	"github.com/tangx-labs/gin-goft/cmd/demo/models"
	"github.com/tangx-labs/gin-goft/httpx"
)

type GetUserByID struct {
	httpx.MethodGet
	UserID string `uri:"id"`

	DBA *adaptors.GormAdaptor
}

func (user *GetUserByID) Path() string {
	return "/users/:id"
}

func (user *GetUserByID) Handler() (interface{}, error) {
	um := &models.User{}

	user.DBA.DB.Where("user_id=?", user.UserID).First(um)
	return um, nil
}
