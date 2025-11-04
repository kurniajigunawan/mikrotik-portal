package mikrotik

import (
	"context"
	"errors"
	"net/http"

	"github.com/aidapedia/go-routeros/driver"
	"github.com/aidapedia/go-routeros/model"
	"github.com/aidapedia/go-routeros/module"

	gerr "github.com/aidapedia/gdk/error"
	ghttp "github.com/aidapedia/gdk/http"
)

func (u *Usecase) ResetSession(ctx context.Context, username string) error {
	active, err := driver.New(u.routerBuilder, module.HotspotActiveModule)
	if err != nil {
		return gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
	}
	activeRes, err := active.Print(ctx, model.PrintRequest{
		Where: []model.Where{
			{
				Field:    "user",
				Value:    username,
				Operator: model.OperatorEqual,
			},
		},
	})
	if err != nil {
		return gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
	}

	if len(activeRes) == 0 {
		// To prevent user to abuse reset session
		return nil
	}

	for _, record := range activeRes {
		req, ok := record.(*module.HotspotActiveData)
		if !ok {
			return gerr.NewWithMetadata(errors.New("failed to cast record to HotspotActiveData"), ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
		}
		err := active.Remove(ctx, req.ID)
		if err != nil {
			return gerr.NewWithMetadata(err, ghttp.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
		}
	}
	return nil
}
