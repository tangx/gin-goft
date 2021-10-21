package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx-labs/gin-goft/cmd/demo/adaptors"
	"github.com/tangx-labs/gin-goft/cmd/demo/models"
	"github.com/tangx-labs/gin-goft/httpx"
	"github.com/tangx/ginbinder"
)

type GetUserByID struct {
	httpx.MethodGet
	DBA *adaptors.GormAdaptor `ginbinder:"-"`
}
type GetUserByIDParams struct {
	UserID string `uri:"id"`
}

func (user *GetUserByID) Path() string {
	return "/users/:id"
}

func (user *GetUserByID) Handler(c *gin.Context) (interface{}, error) {
	um := &models.User{}

	params := &GetUserByIDParams{}
	if err := ginbinder.ShouldBindRequest(c, params); err != nil {
		return nil, err
	}

	user.DBA.DB.Where("user_id=?", params.UserID).First(um)
	return um, nil
}
