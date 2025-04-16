package main

import (
	"vicore_hrd/app/rest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// SERVICES FOR ANTREAN
func (service *Service) FiberAntreanRouterListen(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Post("login", service.HRDHandler.LoginByEmailAndPasswordFiberHandler)
	apiV1.Post("antrean-pasien", rest.JWTProtected(), service.AntreanHandler.OnGetAntreanFiberHandler)
	apiV1.Post("dashboard", service.AntreanHandler.OnDashboardFiberHandler)
	apiV1.Post("register-pasien", service.AntreanHandler.OnGetDataRegistrasiPasienFiberHandler)

	apiV1.Post("pasien-pulang-by-date", service.AntreanHandler.OnGetPasienPulangByDateFiberHandler)
}
