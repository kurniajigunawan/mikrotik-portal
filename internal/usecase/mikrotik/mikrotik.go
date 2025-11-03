package mikrotik

import (
	"context"
	"errors"
	"net/http"

	"github.com/aidapedia/go-routeros/driver"
	"github.com/aidapedia/go-routeros/model"
	"github.com/aidapedia/go-routeros/module"

	gErr "github.com/aidapedia/gdk/error"

	pkgLog "github.com/kurniajigunawan/mikrotik-portal/pkg/log"
)

func (u *Usecase) ResetSession(ctx context.Context, username string) error {
	active, err := driver.New(u.routerBuilder, module.HotspotActiveModule)
	if err != nil {
		return gErr.NewWithMetadata(err, pkgLog.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
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
		return gErr.NewWithMetadata(err, pkgLog.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
	}

	if len(activeRes) == 0 {
		// To prevent user to abuse reset session
		return nil
	}

	for _, record := range activeRes {
		req, ok := record.(*module.HotspotActiveData)
		if !ok {
			return gErr.NewWithMetadata(errors.New("failed to cast record to HotspotActiveData"), pkgLog.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
		}
		err := active.Remove(ctx, req.ID)
		if err != nil {
			return gErr.NewWithMetadata(err, pkgLog.Metadata(http.StatusInternalServerError, "Internal Server Error. Please try again later."))
		}
	}
	return nil
}
