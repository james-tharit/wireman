package utils

import (
	"fmt"
	"math/rand"
	"time"
	"wireman/graph/model"

	"github.com/google/uuid"
)

func CreateMockArticle() *model.Article {
	var literalString = `
Motion without direction is hollow and foolish.

Living life without a clear vision is a painful mistake I wish I had corrected earlier in my life. To fulfill YOUR dreams, your actions need to be aligned with your aspirations.

It’s never too late!

Are you pursuing YOUR goals or someone else’s?
`
	currentTime := time.Now()
	id := uuid.NewString()
	randomNum := rand.Intn(100)
	return &model.Article{
		ID:          id,
		Title:       "Mistake that took me years to understand...",
		Body:        literalString,
		Author:      fmt.Sprintf("Anonymous Giraff%d", randomNum),
		Views:       0,
		Likes:       0,
		Shares:      0,
		ArticleType: model.ArticleStyleGreenViper,
		Created:     currentTime.String(),
		Updated:     currentTime.String(),
	}
}
