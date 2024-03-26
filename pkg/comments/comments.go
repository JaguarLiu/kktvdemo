package comments

import (
	"fmt"
	"time"
)

// Comment Service
type CommentService interface {
	Create(comment Comment)
}

type CommentSrv struct{}

// Comment API Request Payload
type CommentReq struct {
	UserID  string `json:"userId"`
	MovieID string `json:"movieId"`
	Content string `json:"content"`
	Rating  int    `json:"rating"`
}

// Comment entity
type Comment struct {
	ID        string
	UserID    string
	MovieID   string
	Content   string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Reports   int
}

func New() *CommentSrv {
	return &CommentSrv{}
}

// 創建一條新的短評
func (srv *CommentSrv) Create(req CommentReq) (*Comment, error) {
	// 驗證輸入
	if req.Content == "" || len(req.Content) > 20 || req.Rating < 1 || req.Rating > 5 {
		return nil, fmt.Errorf("invalid input")
	}
	comment := Comment{}
	// 假設 generateUniqueID
	commentID := generateUniqueID()
	currentTime := time.Now()

	comment.ID = commentID
	comment.UserID = req.UserID
	comment.MovieID = req.MovieID
	comment.Content = req.Content
	comment.Rating = req.Rating
	comment.CreatedAt = currentTime
	comment.UpdatedAt = currentTime
	comment.Reports = 0

	// 儲存短評到數據庫
	// saveCommentToDatabase(comment)

	return &comment, nil
}

// generateUniqueID 生成唯一標識符
func generateUniqueID() string {
	// 實現細節省略
	return "uniqueID123"
}
