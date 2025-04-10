package handler

import (
	"net/http"
	"vicore_hrd/modules/triase/entity"
	"vicore_hrd/pkg/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TriaseHandler struct {
	TriaseMapper  entity.TriaseMapper
	TriaseUseCase entity.TriaseUseCase
	TriaseRepo    entity.TriaseRepository
	Logging       *logrus.Logger
}

func (hh *TriaseHandler) OnGetDataTriaseFiberHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")
	var BaseURL = c.BaseURL()

	triase, _ := hh.TriaseUseCase.OnGetReportTriaseUseCase(NoReg, BaseURL)

	response := helper.APIResponse("OK", http.StatusOK, triase)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *TriaseHandler) OnGetDataTriasePonekFiberHadler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")
	var BaseURL = c.BaseURL()

	triase, _ := hh.TriaseUseCase.OnGetReportTriasePonekUseCase(NoReg, BaseURL)

	response := helper.APIResponse("OK", http.StatusOK, triase)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *TriaseHandler) OnChangedAsesmendIGDHandler(c *fiber.Ctx) error {
	data, _ := hh.TriaseRepo.GetAsesmenIGD()

	for _, V := range data {
		// GET ASESMEN
		// GetDiagnosa(Noreg string) (res triase.DokterDiagnosa, err error)
		dataDokter, _ := hh.TriaseRepo.GetDiagnosa(V.Noreg)
		hh.Logging.Info("DIAGNOSA PASIEN")
		hh.Logging.Info(dataDokter.Kodediagnosa)
		update, _ := hh.TriaseRepo.UpDateDiagnosa(dataDokter.Kodediagnosa, V.Noreg, V.KdBagian)
		// UpDateDiagnosa(Diagnosa string, Noreg string, KdBagian string) (res triase.IgdAsesmen, err error)

		// UPDATE NAMA DOKTER
		// updateD, _ := hh.TriaseRepo.UpDateDiagnosa(dataDokter.Dokter, V.Noreg, V.KdBagian)
		// UpDateDiagnosa(Diagnosa string, Noreg string, KdBagian string) (res triase.IgdAsesmen, err error)
		// UpdateAsesmedKonsulKe
		hh.Logging.Info(update)
	}

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)

}
