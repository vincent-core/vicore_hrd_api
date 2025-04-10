package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

func (service *Service) FiberLembarKonsulListen(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV2 := Router.Group("/app/v2/")
	apiV1.Get("lembar-konsul/:noreg", service.LembarKonsulHandler.OnGetReportLembarKonsulFiberHandler)
	apiV2.Get("lembar-konsul/:noreg", service.LembarKonsulHandler.OnGetReportLembarKonsulV2FiberHandler)
}
