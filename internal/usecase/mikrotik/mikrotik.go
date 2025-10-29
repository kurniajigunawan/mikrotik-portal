package mikrotik

import (
	"context"
	"log"

	"github.com/aidapedia/go-routeros/driver"
	"github.com/aidapedia/go-routeros/model"
	"github.com/aidapedia/go-routeros/module"
)

func (u *Usecase) ResetSession(ctx context.Context, username string) error {
	active, errs := driver.New(u.routerBuilder, module.HotspotActiveModule)
	if errs != nil {
		log.Fatal(errs)
	}
	activeRes, err := active.Print(context.Background(), model.PrintRequest{
		Where: []model.Where{
			{
				Field:    "user",
				Value:    username,
				Operator: model.OperatorEqual,
			},
		},
	})
	if err != nil {
		return err
	}

	if len(activeRes) == 0 {
		return err
	}

	for _, record := range activeRes {
		req, ok := record.(*module.HotspotActiveData)
		if !ok {
			return err
		}
		err := active.Remove(context.Background(), req.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
