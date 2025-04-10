package handler

import (
	"fmt"
	"net/http"
	"strings"
	"vicore_hrd/modules/hrd/dto"
	"vicore_hrd/modules/hrd/entity"
	"vicore_hrd/pkg/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HRDHandler struct {
	HRDRepository entity.VicoreHRDRepository
	HRDUseCase    entity.VicoreHRDUseCase
	HRDMapper     entity.VicoreHRDMapper
	Logging       *logrus.Logger
}

func (hh *HRDHandler) LoginByEmailAndPasswordFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.ReqLoginApp)
	errs := c.BodyParser(&payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak dapat diproses", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	if errs := validate.Struct(payload); errs != nil {
		errors := helper.FormatValidationError(errs)
		message := fmt.Sprintf("Error %s, Data tidak dapat disimpan", strings.Join(errors, "\n"))
		response := helper.APIResponse(message, http.StatusBadRequest, errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	baseURL := c.BaseURL()
	user, message, er11 := hh.HRDUseCase.OnLoginUserByEmailAndPasswordUseCase(payload.Email, payload.Password, baseURL)

	if er11 != nil {
		response := helper.APIResponseFailure(message, http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse(message, http.StatusOK, user)
	return c.Status(fiber.StatusOK).JSON(response)
}
