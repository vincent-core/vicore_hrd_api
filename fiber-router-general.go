package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// SERVICE IN FOR LOGIN
func (service *Service) FiberRoutingGeneralConsent(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Get("general-consent/:noreg/:kd_bagian", service.GeneralHandler.OnGetGeneralConsentFiberHandler)
	apiV1.Get("general-consent-ranap/:noreg/:kd_bagian", service.GeneralHandler.OnGetGeneralConsentRanapFiberHandler)
	apiV1.Get("general-consent-rajal/:noreg/:id", service.GeneralHandler.OnGetGeneralConsentRajalFiberHandler)
}
