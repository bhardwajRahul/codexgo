package form

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/errs"
	"github.com/bastean/codexgo/pkg/cmd/server/util/key"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/gin-gonic/gin"
)

func UserUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get(key.UserId)

		if !exists {
			c.Error(errs.MissingKey(key.UserId, "UserUpdate"))
			c.Abort()
			return
		}

		command := new(user.UpdateCommand)

		c.BindJSON(command)

		command.Id = id.(string)

		err := user.Update.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, reply.JSON(true, "account updated", reply.Payload{}))
	}
}
