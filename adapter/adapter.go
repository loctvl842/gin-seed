package adapter

import (
	"app/adapter/core"
	router_v1 "app/adapter/routers/v1"
	"app/addons"
	"app/addons/logger"
	"app/addons/server"
	"app/gosdk"

	_ "ariga.io/atlas-provider-gorm/gormschema"
)

type adapter struct {
	service *core.AdapterService
}

func NewAdapter(appCtx gosdk.ApplicationContext) *adapter {
	return &adapter{
		service: &core.AdapterService{
			AppCtx: appCtx,
			Logger: logger.GetCurrent().GetLogger("adapter"),
		},
	}
}

func (s *adapter) Start() {
	gs := s.service.AppCtx.MustGet(string(addons.GinServerPrefix)).(*server.GinServer)

	// Organize the routes by version
	router_v1.RegisterRoutes(gs.Group("v1"), s.service)

	if err := gs.ListenAndServe(); err != nil {
		s.service.Logger.Fatalln(err)
	}
}
