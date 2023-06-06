package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yunzhong/gateway/driver"
)

func main() {
	Dirverapp := driver.NewDriver(driver.Before(func() error {
		return nil
	}))
	if err := Dirverapp.Run(); err != nil {
		logrus.Fatal(err)
	}
}
