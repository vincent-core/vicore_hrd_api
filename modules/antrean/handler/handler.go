package handler

import (
	"fmt"
	"net/http"
	"strings"
	"vicore_hrd/app/rest"
	"vicore_hrd/modules/antrean"
	"vicore_hrd/modules/antrean/dto"
	"vicore_hrd/modules/antrean/entity"
	asesmenEntity "vicore_hrd/modules/asesmen/entity"
	"vicore_hrd/pkg/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AntreanHandler struct {
	AntreanRepository entity.AntreanRepository
	AntreanUseCase    entity.AntreanUseCase
	AntreanMapper     entity.AntreanMapper
	Logging           *logrus.Logger
	AsesmenRepository asesmenEntity.AsesmenRepository
}

func (hh *AntreanHandler) OnDashboardFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.Dashboard)
	errs := c.BodyParser(&payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak dapat diproses", http.StatusCreated)
		hh.Logging.Error(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	validate := validator.New()

	if errs := validate.Struct(payload); errs != nil {
		errors := helper.FormatValidationError(errs)
		message := fmt.Sprintf("Error %s, Data tidak dapat disimpan", strings.Join(errors, "\n"))
		response := helper.APIResponse(message, http.StatusBadRequest, errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	data, er11 := hh.AntreanUseCase.OnDashboardUseCase(payload.KDBagian)

	if er11 != nil {
		response := helper.APIResponseFailure("data gagal didapat", http.StatusCreated)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AntreanHandler) OnGetAntreanFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.GetAntranPasien)
	errs := c.BodyParser(&payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak dapat diproses", http.StatusCreated)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	validate := validator.New()

	if errs := validate.Struct(payload); errs != nil {
		errors := helper.FormatValidationError(errs)
		message := fmt.Sprintf("Error %s, Data tidak dapat disimpan", strings.Join(errors, "\n"))
		response := helper.APIResponse(message, http.StatusBadRequest, errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	person := c.Locals("person").(string)
	userID := c.Locals("userID").(string)

	data, message, er11 := hh.AntreanUseCase.OnGetAntrianIGDUseCase(payload.KDBagian, person, userID)

	if er11 != nil {
		response := helper.APIResponseFailure(message, http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse(message, http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AntreanHandler) OnGetDataRegistrasiPasienFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.OnGetDataRegisterByID)
	errs := c.BodyParser(&payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak dapat diproses", http.StatusBadRequest)
		hh.Logging.Error(response)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	if errs := validate.Struct(payload); errs != nil {
		errors := helper.FormatValidationError(errs)
		message := fmt.Sprintf("Error %s, Data tidak dapat disimpan", strings.Join(errors, "\n"))
		response := helper.APIResponse(message, http.StatusBadRequest, errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var dregisterPasien = []antrean.DRegisterPasien{}

	data, _ := hh.AntreanRepository.OnGetDataRegisterPasienByID(payload.ID)

	if len(data) == 0 {
		response := helper.APIResponse("OK", http.StatusOK, make([]antrean.DRegisterPasien, 0))
		return c.Status(fiber.StatusOK).JSON(response)
	}

	for i := 0; i <= len(data)-1; i++ {
		asesmen, _ := hh.AntreanRepository.OnGetPengkajianDokterRepo(data[i].Noreg)

		tglIndo, _ := rest.UbahTanggalIndo(data[i].Tanggal)

		dregisterPasien = append(dregisterPasien, antrean.DRegisterPasien{
			Tanggal:    tglIndo + " " + data[i].Jam + " WIB",
			Id:         data[i].Id,
			Noreg:      data[i].Noreg,
			Nama:       data[i].Nama,
			Kunjungan:  data[i].Kunjungan,
			Pelayaan:   asesmen.Pelayanan,
			Bagian:     asesmen.KPelayanan.Bagian,
			Keterangan: data[i].Keterangan,
			KdBag:      data[i].KdBag,
		})
	}

	mapper := hh.AntreanMapper.ToMappingDataDResiterPasien(dregisterPasien)

	response := helper.APIResponse("OK", http.StatusOK, mapper)
	return c.Status(fiber.StatusOK).JSON(response)
}
