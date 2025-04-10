package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

func (service *Service) FiberRoutingResumeListen(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Get("index", service.Resumehandler.OnIndex)
	apiV1.Get("resume-medis/:noreg", service.Resumehandler.OnGetDataResumeMedisFiberHandler)
	apiV1.Get("cari-pasien-pulang/:norm", service.Resumehandler.OnGetPasienFiberHandler)
}
