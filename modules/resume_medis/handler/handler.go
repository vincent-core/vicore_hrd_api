package handler

import (
	"net/http"
	antreanEntity "vicore_hrd/modules/antrean/entity"
	resumemedis "vicore_hrd/modules/resume_medis"
	"vicore_hrd/modules/resume_medis/entity"
	"vicore_hrd/pkg/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ResumeHandler struct {
	Logging           *logrus.Logger
	ResumerMapper     entity.ResumeMapper
	ResumeUseCase     entity.ResumeMedisUseCase
	ResumeRepository  entity.ResumeMedisRepository
	AntreanRepository antreanEntity.AntreanRepository
}

func (hh *ResumeHandler) OnGetDataResumeMedisFiberHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")

	data := hh.ResumeUseCase.OnReportdataRingkasanPulangUseCase(NoReg)

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *ResumeHandler) OnIndex(c *fiber.Ctx) error {

	response := helper.APIResponse("OK", http.StatusOK, "data")
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *ResumeHandler) CariPasienPulangFiberHandler(c *fiber.Ctx) error {
	var Cari = c.Params("norm")

	data, err12 := hh.ResumeRepository.CariPasienPulangRanapRepository(Cari)

	if err12 != nil {
		response := helper.APIResponse("data tidak ditemukan", http.StatusBadRequest, make([]resumemedis.CariDataPasienPulang, 0))
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if len(data) == 0 {
		response := helper.APIResponse("data kosong", http.StatusOK, make([]resumemedis.CariDataPasienPulang, 0))
		return c.Status(fiber.StatusOK).JSON(response)
	}

	// OnGetJenisKelamin(NOID string) (res resumemedis.DataJenis, err error)
	jk, _ := hh.ResumeRepository.OnGetJenisKelamin(data[0].Norm)

	mapper := hh.ResumerMapper.ToMappingCariDataPasienPulang(data, jk)

	response := helper.APIResponse("OK", http.StatusOK, mapper)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *ResumeHandler) OnGetPasienFiberHandler(c *fiber.Ctx) error {
	var Cari = c.Params("norm")

	data, err12 := hh.AntreanRepository.OnGetDataRegisterPasienByID(Cari)

	if err12 != nil {
		response := helper.APIResponse("data tidak ditemukan", http.StatusBadRequest, make([]resumemedis.CariDataPasienPulang, 0))
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if len(data) == 0 {
		response := helper.APIResponse("data kosong", http.StatusOK, make([]resumemedis.CariDataPasienPulang, 0))
		return c.Status(fiber.StatusOK).JSON(response)
	}

	profile, _ := hh.ResumeRepository.OnGetDataProfilePasienRepository(Cari)

	mapper := hh.ResumerMapper.ToMappingCariPasienDRegister(data, profile)

	response := helper.APIResponse("OK", http.StatusOK, mapper)
	return c.Status(fiber.StatusOK).JSON(response)

}
