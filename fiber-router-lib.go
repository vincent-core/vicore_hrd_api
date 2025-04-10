package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// SERVICE IN HERE
func (service *Service) FiberRoutingPelayananListen(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Get("pelayanan", service.LibHandler.GetAllPelayananFiberHandler)
	apiV1.Get("rekam-medis", service.LibHandler.GetDataRekamMedisFiberHandler)
}
