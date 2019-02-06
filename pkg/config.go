package config

import "log"

type Application struct {
	Errlog  *log.Logger
	Infolog *log.Logger
}
