package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kktv/pkg/comments"
)

// createCommentHandler 新增短評的
func createCommentHandler(c *gin.Context) {
	var req comments.CommentReq
	var srv = comments.New()

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 在這裡添加將新短評保存到數據庫的邏輯
	res, err := srv.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回成功響應
	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully!", "comment": res})
}

func main() {
	r := gin.Default()

	// 定義一個 POST 路由
	r.POST("/comments", createCommentHandler)

	// 啟動服務器
	r.Run() // 預設在 localhost:8080
}
