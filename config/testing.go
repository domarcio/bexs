// +build testing

package config

import commonLog "github.com/domarcio/bexs/src/infra/log"

const (
	Env                  string = "testing"
	RouteStorageFilePath string = "data/storage/testing.routes.csv"
	Logfile              string = "data/log/app.log"
)

var (
	LogService commonLog.Logger = commonLog.NewLogprint()
)
