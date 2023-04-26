package business

import (
	"fmt"
	"log"

	"fiberpgx1/model"
)

type Admin struct {
	Pg *model.Adapter
}

type CreateArticleIn struct {
	model.ArticleModel
}

type CreateArticleOut struct {
	Err error
}

func (a *Admin) CreateArticle(in CreateArticleIn) (out CreateArticleOut) {
	// validasi
	if in.Title == "" {
		out.Err = fmt.Errorf("title is required")
		return
	}
	err := in.ArticleModel.Insert(a.Pg)
	if err != nil {
		out.Err = fmt.Errorf("failed to save article")
		log.Println(err)
		return
	}
	return
}
