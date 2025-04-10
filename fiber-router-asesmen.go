package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

func (service *Service) FiberRoutingAsesmenListen(Logging *logrus.Logger, Router fiber.Router, Store *session.Store) {
	apiV1 := Router.Group("/app/v1/")
	apiV2 := Router.Group("/app/v2/")
	apiV1.Get("asesmen-dokter-igd/:noreg", service.AsemsenHandler.OnReportAsesmenIGDDokterHandler)
	apiV1.Get("pengantar-rawat-inap/:noreg", service.AsemsenHandler.OnReportPengantarRawatInapFiberHandler)
	apiV2.Get("pengantar-rawat-inap/:noreg/:kd_bagian", service.AsemsenHandler.OnReportPengantarRawatInapFiberV2Handler)
	apiV1.Get("laporan-operasi/:noreg", service.AsemsenHandler.OnGetLaporanOperasiFiberHandler)
	apiV2.Get("laporan-operasi/:norm", service.AsemsenHandler.OnGetLaporanOperasiByNORMFiberHandler)
	apiV1.Get("asesmen-dokter-ponek/:noreg", service.AsemsenHandler.OnReportAsesmenPONEKDokterHandler)
	apiV1.Get("asesmen-dokter-rawat-inap/:noreg", service.AsemsenHandler.OnReportAsesmenIGDDokterRawatInapHandler)

}
