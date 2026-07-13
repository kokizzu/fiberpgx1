package business

import (
	"errors"
	"testing"
)

func TestCreateArticleRequiresTitle(t *testing.T) {
	t.Parallel()

	out := (&Admin{}).CreateArticle(CreateArticleIn{})
	if !errors.Is(out.Err, ErrTitleRequired) {
		t.Fatalf("expected ErrTitleRequired, got %v", out.Err)
	}
}
