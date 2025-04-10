package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// SERVICE IN FOR LOGIN
func (service *Service) FiberTriaseRoutingListen(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Get("report-triase/:noreg", service.TriaseHandler.OnGetDataTriaseFiberHandler)
	apiV1.Get("report-triase-ponek/:noreg", service.TriaseHandler.OnGetDataTriasePonekFiberHadler)
	// OnChangedAsesmendIGDHandler
	apiV1.Get("asesmen-igd", service.TriaseHandler.OnChangedAsesmendIGDHandler)
}
