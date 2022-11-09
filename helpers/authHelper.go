package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserToUid(c *gin.Context, userId string) (err error){
	uid := c.GetString("uid")
	err = nil

	if uid != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}