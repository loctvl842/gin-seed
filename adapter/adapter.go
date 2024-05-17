package adapter

import (
	router_v1 "app/adapter/routers/v1"
	"app/addons"
	"app/addons/logger"
	"app/addons/server"
	"app/gosdk"
)

type adapter struct {
	appCtx gosdk.ApplicationContext
	logger logger.Logger
}

func NewAdapter(appCtx gosdk.ApplicationContext) *adapter {
	return &adapter{
		appCtx: appCtx,
		logger: logger.GetCurrent().GetLogger("adapter"),
	}
}

func (s *adapter) Start() {
	gs := s.appCtx.MustGet(string(addons.GinServerPrefix)).(*server.GinServer)

	router_v1.RegisterRoutes(gs.Group("v1"))

	if err := gs.ListenAndServe(); err != nil {
		s.logger.Fatalln(err)
	}
}
