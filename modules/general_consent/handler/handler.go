package handler

import (
	"net/http"
	"vicore_hrd/modules/general_consent/entity"
	"vicore_hrd/pkg/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type GeneralHandler struct {
	Logging        *logrus.Logger
	GeneralUsecase entity.GeneralConsentUseCase
	GeneralRepo    entity.GeneralConsentRepository
	GeneralMapper  entity.GeneralConsentMapper
}

func (hh *GeneralHandler) OnGetGeneralConsentFiberHandler(c *fiber.Ctx) error {

	var NoReg = c.Params("noreg")
	var Bagian = c.Params("kd_bagian")

	general := hh.GeneralUsecase.OnGetGeneralConsentRAJALUseCase(NoReg, Bagian)

	response := helper.APIResponse("OK", http.StatusOK, general)
	return c.Status(fiber.StatusOK).JSON(response)

}

func (hh *GeneralHandler) OnGetGeneralConsentRajalFiberHandler(c *fiber.Ctx) error {

	var NoReg = c.Params("noreg")
	var ID = c.Params("id")

	general := hh.GeneralUsecase.OnGetGeneralConsentRAJALUseCaseV2(NoReg, ID)

	response := helper.APIResponse("OK", http.StatusOK, general)
	return c.Status(fiber.StatusOK).JSON(response)

}

func (hh *GeneralHandler) OnGetGeneralConsentRanapFiberHandler(c *fiber.Ctx) error {

	var NoReg = c.Params("noreg")
	var Bagian = c.Params("kd_bagian")

	general := hh.GeneralUsecase.OnGetGeneralConsentRANAPUseCase(NoReg, Bagian)

	response := helper.APIResponse("OK", http.StatusOK, general)
	return c.Status(fiber.StatusOK).JSON(response)

}
