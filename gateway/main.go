package main

import (
	app "github.com/yunzhong/gateway/driver"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	a := app.New(app.Before(func() error {
		return nil
	}))

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
