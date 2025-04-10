package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// SERVICE IN FOR LOGIN
func (service *Service) FiberRoutingUserAndListen(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Post("login", service.HRDHandler.LoginByEmailAndPasswordFiberHandler)
}
