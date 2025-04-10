package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"vicore_hrd/app/rest"
	"vicore_hrd/modules/asesmen/dto"
	"vicore_hrd/modules/asesmen/entity"
	GeneralEntity "vicore_hrd/modules/general_consent/entity"
	ResumeEntity "vicore_hrd/modules/resume_medis/entity"
	"vicore_hrd/pkg/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AsesmenHandler struct {
	Logging           *logrus.Logger
	AsesmenRepository entity.AsesmenRepository
	AsesmenMapper     entity.AsesmenMapper
	AsesmenUseCase    entity.AsesmenUseCase
	ResumeRepoitory   ResumeEntity.ResumeMedisRepository
	GeneralRepository GeneralEntity.GeneralConsentRepository
}

func (hh *AsesmenHandler) OnSaveCPPSOAPPasienFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.InsertCPPTSOAP)
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

	modulID := c.Locals("modulID").(string)
	person := c.Locals("person").(string)
	userID := c.Locals("userID").(string)

	message, ee11 := hh.AsesmenUseCase.OnSaveCPPTSOAPUseCase(modulID, person, userID, *payload)

	if ee11 != nil {
		response := helper.APIResponseFailure(message, http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponseFailure(message, http.StatusCreated)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (hh *AsesmenHandler) OnSaveCPPTSBARPasienFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.InsertCPPTSBAR)
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

	modulID := c.Locals("modulID").(string)
	person := c.Locals("person").(string)
	userID := c.Locals("userID").(string)

	message, ee11 := hh.AsesmenUseCase.OnSaveCPPTSBARUseCase(modulID, person, userID, *payload)

	if ee11 != nil {
		response := helper.APIResponseFailure(message, http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponseFailure(message, http.StatusCreated)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (hh *AsesmenHandler) OnGetCPPTPasienFiberHandler(c *fiber.Ctx) error {
	var ID = c.Params("id")

	data, er11 := hh.AsesmenUseCase.OnGetReportCPPTUseCase(ID)

	if er11 != nil {
		response := helper.APIResponse("Error", http.StatusCreated, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnGetCPPTPasienDecriptFiberHandler(c *fiber.Ctx) error {
	var ID = c.Params("id")

	data, er11 := hh.AsesmenUseCase.OnGetReportCPPTUseCase(ID)

	if er11 != nil {
		response := helper.APIResponse("Error", http.StatusBadRequest, data)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	encript, _ := rest.MCEncrypt(c.JSON(data), os.Getenv("SECRET_KEY"))

	response := helper.APIResponse("OK", http.StatusOK, encript)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnUpdateCPPTSOAPFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.OnUpdateCPPTSoapRes)
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

	userID := c.Locals("userID").(string)

	message, er11 := hh.AsesmenUseCase.OnUpdateCPPTSOAPUseCase(userID, *payload)

	if er11 != nil {
		response := helper.APIResponseFailure(message, http.StatusCreated)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponseFailure(message, http.StatusOK)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnUpdateCPPTSBARFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.OnUpdateSBARRes)
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

	userID := c.Locals("userID").(string)

	message, er11 := hh.AsesmenUseCase.OnUpdateCPPSBARUseCase(userID, *payload)

	if er11 != nil {
		response := helper.APIResponseFailure(message, http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponseFailure(message, http.StatusOK)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnReportAsesmenIGDDokterHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")
	var BaseURL = c.BaseURL()

	data := hh.AsesmenUseCase.OnReportAsesmenIGDDokterUseCase(NoReg, BaseURL)

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnReportAsesmenIGDDokterRawatInapHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")
	var BaseURL = c.BaseURL()

	data := hh.AsesmenUseCase.OnReportAsesmenIGDDokterRawatInapUseCase(NoReg, BaseURL)

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnReportAsesmenPONEKDokterHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")
	var BaseURL = c.BaseURL()

	data := hh.AsesmenUseCase.OnReportAsesmenPONEKDokterUseCase(NoReg, BaseURL)

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnReportPengantarRawatInapFiberHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")

	data, err11 := hh.AsesmenUseCase.OnReportPengantarRawatInapUseCase(NoReg)

	if err11 != nil {
		response := helper.APIResponse("Error", http.StatusBadRequest, data)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnReportPengantarRawatInapFiberV2Handler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")
	var KdBagian = c.Params("kd_bagian")

	data, _ := hh.AsesmenUseCase.OnReportPengantarRawatInapUseCaseV2(NoReg, KdBagian)

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnGetReportCPPTPasienFiberHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")

	data, er11 := hh.AsesmenUseCase.OnGetReportCPPByNoRegTUseCase(NoReg)

	if er11 != nil {
		response := helper.APIResponse("Error", http.StatusBadRequest, data)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("OK", http.StatusOK, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

// ========== LAPORAN OPERASI
func (hh *AsesmenHandler) OnGetLaporanOperasiFiberHandler(c *fiber.Ctx) error {
	var NoReg = c.Params("noreg")

	dataOperasi, er11 := hh.AsesmenRepository.OnGetOperasiPasienRepository(NoReg)

	if er11 != nil {
		response := helper.APIResponseFailure("Data tidak ditemukan", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if dataOperasi.Noreg == "" {
		response := helper.APIResponseFailure("Data tidak ditemukan", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// GET PROFILE PASIEN BY ID LAKUKAN MAPPING DATA
	pasien, _ := hh.GeneralRepository.OnGetPasienGeneralConsentRepository(dataOperasi.Id)
	bedah3, _ := hh.AsesmenRepository.OnGetBedah3Repository(dataOperasi.Nomor)
	bedah2, _ := hh.AsesmenRepository.OnGetDiagnosaRepository(dataOperasi.Nomor)
	mapper := hh.AsesmenMapper.ToMappingDataLaporanBedah(NoReg, dataOperasi, pasien, bedah3, bedah2)

	response := helper.APIResponse("OK", http.StatusOK, mapper)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (hh *AsesmenHandler) OnGetLaporanOperasiByNORMFiberHandler(c *fiber.Ctx) error {
	var NORM = c.Params("norm")

	dataOperasi, er11 := hh.AsesmenRepository.OnGetLaporanOperasiByNoRMRepository(NORM)

	if er11 != nil {
		response := helper.APIResponseFailure("Data tidak ditemukan", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if dataOperasi.Noreg == "" {
		response := helper.APIResponseFailure("Data tidak ditemukan", http.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// GET PROFILE PASIEN BY ID LAKUKAN MAPPING DATA
	pasien, _ := hh.GeneralRepository.OnGetPasienGeneralConsentRepository(dataOperasi.Id)
	bedah3, _ := hh.AsesmenRepository.OnGetBedah3Repository(dataOperasi.Nomor)
	bedah2, _ := hh.AsesmenRepository.OnGetDiagnosaRepository(dataOperasi.Nomor)
	mapper := hh.AsesmenMapper.ToMappingDataLaporanBedah(dataOperasi.Noreg, dataOperasi, pasien, bedah3, bedah2)

	response := helper.APIResponse("OK", http.StatusOK, mapper)
	return c.Status(fiber.StatusOK).JSON(response)
}
