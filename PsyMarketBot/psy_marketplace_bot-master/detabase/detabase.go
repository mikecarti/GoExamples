package detabase

import (
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"

	"psy_bot/lib/e"
)

type User struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

func New(projectKey string) (*base.Base, error) {

	// initialize with project key
	// returns ErrBadProjectKey if project key is invalid
	d, err := deta.New(deta.WithProjectKey(projectKey))
	if err != nil {
		return nil, e.Wrap("failed to init new Deta instance:", err)
	}

	// initialize with base name
	// returns ErrBadBaseName if base name is invalid
	db, err := base.New(d, "base_name")
	if err != nil {
		return nil, e.Wrap("failed to init new Base instance:", err)
	}

	return db, nil
}
