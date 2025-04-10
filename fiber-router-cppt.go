package main

import (
	"vicore_hrd/app/rest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// ROUTING FOR CPPT
func (service *Service) FiberRuterCPPTV1(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Post("cppt-soap", rest.JWTProtected(), service.AsemsenHandler.OnSaveCPPSOAPPasienFiberHandler)
	apiV1.Post("cppt-sbar", rest.JWTProtected(), service.AsemsenHandler.OnSaveCPPTSBARPasienFiberHandler)
	apiV1.Put("cppt-soap", rest.JWTProtected(), service.AsemsenHandler.OnUpdateCPPTSOAPFiberHandler)
	apiV1.Put("cppt-sbar", rest.JWTProtected(), service.AsemsenHandler.OnUpdateCPPTSBARFiberHandler)
	apiV1.Get("cppt/:id", service.AsemsenHandler.OnGetCPPTPasienFiberHandler)
	apiV1.Get("report-cppt/:noreg", service.AsemsenHandler.OnGetReportCPPTPasienFiberHandler)
}
