package render

import (
	"context"
	"fmt"
)

func (u *Usecase) GetPage(ctx context.Context, page string) (GetPageResponse, error) {
	data, ok := u.pages[page]
	if !ok {
		return GetPageResponse{}, fmt.Errorf("page %s not found", page)
	}
	return GetPageResponse{
		Page: page,
		Data: data,
	}, nil
}
