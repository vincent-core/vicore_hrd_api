package handler

import (
	"net/http"
	"vicore_hrd/modules/lib/entity"
	"vicore_hrd/pkg/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LibHandler struct {
	LibRepository entity.LibRepository
	LibMapper     entity.LibMapper
	Logging       *logrus.Logger
}

func (hh *LibHandler) GetAllPelayananFiberHandler(c *fiber.Ctx) error {
	data, _ := hh.LibRepository.FindAllPelayananRepository()
	mapper := hh.LibMapper.ToMappingPelayanan(data)

	response := helper.APIResponse("OK", http.StatusOK, mapper)
	return c.Status(fiber.StatusOK).JSON(response)

}

func (hh *LibHandler) GetDataRekamMedisFiberHandler(c *fiber.Ctx) error {
	data, er11 := hh.LibRepository.OnGetDataRekamMedis()

	if er11 != nil {
		hh.Logging.Info(er11)
	}

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}
