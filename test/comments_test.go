package comments

import (
	"kktv/pkg/comments"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentService_Create_ValidInput(t *testing.T) {
	srv := comments.New()
	comment := comments.CommentReq{
		UserID:  "user123",
		MovieID: "movie123",
		Content: "Good movie",
		Rating:  5,
	}

	createdComment, err := srv.Create(comment)
	assert.Nil(t, err)
	assert.NotNil(t, createdComment)
	assert.NotEmpty(t, createdComment.ID)
	assert.Equal(t, comment.UserID, createdComment.UserID)
	assert.Equal(t, comment.MovieID, createdComment.MovieID)
	assert.Equal(t, comment.Content, createdComment.Content)
	assert.Equal(t, comment.Rating, createdComment.Rating)
	assert.Equal(t, createdComment.CreatedAt, createdComment.UpdatedAt)
	assert.Equal(t, 0, createdComment.Reports)
}

func TestCommentService_Create_InvalidInput(t *testing.T) {
	tests := []struct {
		name    string
		comment comments.CommentReq
	}{
		{"Empty Content", comments.CommentReq{UserID: "user123", MovieID: "movie123", Content: "", Rating: 5}},
		{"Content Too Long", comments.CommentReq{UserID: "user123", MovieID: "movie123", Content: "This is a very long content that exceeds the maximum allowed length of twenty characters.", Rating: 5}},
		{"Rating Too Low", comments.CommentReq{UserID: "user123", MovieID: "movie123", Content: "Good movie", Rating: 0}},
		{"Rating Too High", comments.CommentReq{UserID: "user123", MovieID: "movie123", Content: "Good movie", Rating: 6}},
	}

	srv := comments.New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.Create(tt.comment)
			assert.NotNil(t, err)
		})
	}
}
