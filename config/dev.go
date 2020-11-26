// +build dev

package config

import (
	"log"

	commonLog "github.com/domarcio/bexs/src/infra/log"
)

const (
	Env                  string = "dev"
	RouteStorageFilePath string = "data/storage/routes.csv"
	Logfile              string = "data/log/app.log"
)

var (
	LogService commonLog.Logger = commonLog.NewLogfile(Logfile, "[BEXS] ", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
)
