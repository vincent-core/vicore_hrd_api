package main

import (
	"vicore_hrd/app/rest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// SERVICES FOR ANTREAN
func (service *Service) FiberEdukasiTeringrasi(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV1.Get("edukasi-terintegrasi/:id", service.EdukasiHandler.GetEdukasiTerintegrasiFiberHandler)
	apiV1.Get("pemberi-informasi", service.EdukasiHandler.OnGetPemberiInformasiTerintegrasiFiberHandler)
	apiV1.Post("edukasi-terintegrasi", rest.JWTProtected(), service.EdukasiHandler.OnSaveEdukasiTerintegrasiFiberHadler)
	apiV1.Put("edukasi-terintegrasi", rest.JWTProtected(), service.EdukasiHandler.OnUpdateEdukasiTerintegrasiFiberHandler)
}
