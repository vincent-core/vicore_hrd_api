package handler

import (
	"net/http"
	entity "vicore_hrd/modules/lembar_konsul/entity"
	"vicore_hrd/pkg/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LembarKonsulhandler struct {
	Logging          *logrus.Logger
	LembarKonsulRepo entity.LembarKonsulRepository
	LembarUseCase    entity.LembarKonsulUseCase
}

func (hh *LembarKonsulhandler) OnGetReportLembarKonsulFiberHandler(c *fiber.Ctx) error {

	var NoReg = c.Params("noreg")
	var BaseURL = c.BaseURL()

	data, _ := hh.LembarUseCase.OnGetReportLembarKonsulUseCase(NoReg, BaseURL)

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *LembarKonsulhandler) OnGetReportLembarKonsulV2FiberHandler(c *fiber.Ctx) error {

	var NoReg = c.Params("noreg")
	var BaseURL = c.BaseURL()

	data, _ := hh.LembarUseCase.OnGetReportLembarKonsulUseCaseV2(NoReg, BaseURL)

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}
