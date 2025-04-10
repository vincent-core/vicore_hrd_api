package handler

import (
	"fmt"
	"net/http"
	"strings"
	"vicore_hrd/modules/edukasi_terintegrasi/dto"
	"vicore_hrd/modules/edukasi_terintegrasi/entity"
	hrdEntity "vicore_hrd/modules/hrd/entity"
	"vicore_hrd/pkg/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type EdukasiHandler struct {
	EdukasiRepository entity.EdukasiTerintegrasiRepository
	EdukasiUseCase    entity.EdukasiTerinterasiUseCase
	EdukasiMapper     entity.EdukasiTerintegrasiMapper
	HRDRepository     hrdEntity.VicoreHRDRepository
	Logging           *logrus.Logger
}

func (eh *EdukasiHandler) GetEdukasiTerintegrasiFiberHandler(c *fiber.Ctx) error {
	var ID = c.Params("id")

	eh.Logging.Info("GET EDUKASI TERINTEGRASI")
	eh.Logging.Info(ID)

	edukasi, er11 := eh.EdukasiUseCase.OnGetEdukasiTerintegrasiUseCase(ID)

	if er11 != nil {
		response := helper.APIResponseFailure("Data tidak dapat diproses", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("OK", http.StatusOK, edukasi)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (eh *EdukasiHandler) OnGetPemberiInformasiTerintegrasiFiberHandler(c *fiber.Ctx) error {
	data, er11 := eh.HRDRepository.OnFindAllDataKaryawanRepository()

	if er11 != nil {
		response := helper.APIResponseFailure("Data tidak dapat diproses", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// LAKUKAN MAPPING DATA
	maper := eh.EdukasiMapper.ToMappingPemberiInformasi(data)

	response := helper.APIResponse("OK", http.StatusOK, maper)
	return c.Status(fiber.StatusOK).JSON(response)

}

func (eh *EdukasiHandler) OnSaveEdukasiTerintegrasiFiberHadler(c *fiber.Ctx) error {
	payload := new(dto.ReqEdukasiTerintegrasi)
	errs := c.BodyParser(&payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak dapat diproses", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	if errs := validate.Struct(payload); errs != nil {
		errors := helper.FormatValidationError(errs)
		message := fmt.Sprintf("Error %s, Data tidak dapat disimpan", strings.Join(errors, "\n"))
		response := helper.APIResponse(message, http.StatusAccepted, errors)
		return c.Status(fiber.StatusAccepted).JSON(response)
	}

	userID := c.Locals("userID").(string)
	bagian := c.Locals("modulID").(string)

	message, e11 := eh.EdukasiUseCase.OnSaveEdukasiTerintegrasiUseCase(userID, bagian, *payload)

	if e11 != nil {
		response := helper.APIResponseFailure(message, http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponseFailure(message, http.StatusOK)
	return c.Status(fiber.StatusOK).JSON(response)

}

func (eh *EdukasiHandler) OnUpdateEdukasiTerintegrasiFiberHandler(c *fiber.Ctx) error {

	payload := new(dto.ReqOnUpdateEdukasiTerintegrasi)
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

	message, err11 := eh.EdukasiUseCase.OnUpdateEdukasiTerintegrasiUseCase(*payload)

	if err11 != nil {
		response := helper.APIResponseFailure(message, http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponseFailure(message, http.StatusOK)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (e *EdukasiHandler) GetPemberiInformasiEdukasiTerintegrasiFiberHandler(c *fiber.Ctx) error {
	return nil
}
