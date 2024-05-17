package core

import (
	"app/addons/logger"
	"app/gosdk"
)

type AdapterService struct {
	AppCtx gosdk.ApplicationContext
	Logger logger.Logger
}
