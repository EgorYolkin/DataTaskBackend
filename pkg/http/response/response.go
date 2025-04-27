package response

import (
	"DataTask/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

// JSONResponse default API response
// @name JSONResponse
// @Description default JSON-response
type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data" swaggertype:"object"`
	Error   string      `json:"error"`
}

func JSON(
	c *gin.Context,
	status int,
	success bool,
	data interface{},
	err string,
) {
	formatLog := fmt.Sprintf(
		"%d %s %s",
		status,
		c.Request.Method,
		c.Request.URL.Path,
	)

	if err != "" {
		logger.Log.Error(formatLog)
	} else {
		logger.Log.Info(formatLog)
	}

	c.JSON(status, JSONResponse{
		Success: success,
		Data:    data,
		Error:   err,
	})
}
