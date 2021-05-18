package util

import "github.com/gin-gonic/gin"
import "github.com/unknwon/com"
import "github.com/whh881114/go_learning_scripts/gin-blog/pkg/setting"

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()

	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
