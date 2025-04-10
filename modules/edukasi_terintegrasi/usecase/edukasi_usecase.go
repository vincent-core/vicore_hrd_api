package usecase

import (
	"time"
	edukasiterintegrasi "vicore_hrd/modules/edukasi_terintegrasi"
	"vicore_hrd/modules/edukasi_terintegrasi/dto"
	"vicore_hrd/modules/edukasi_terintegrasi/entity"

	"github.com/sirupsen/logrus"
)

type edukasiUseCase struct {
	Logging           *logrus.Logger
	EdukasiRepository entity.EdukasiTerintegrasiRepository
	EdukasiMapper     entity.EdukasiTerintegrasiMapper
}

func NewEdukasiTerintegrasiUseCase(logging *logrus.Logger, EdukasiMapper entity.EdukasiTerintegrasiMapper, EdukasiRepository entity.EdukasiTerintegrasiRepository) entity.EdukasiTerinterasiUseCase {
	return &edukasiUseCase{
		Logging:           logging,
		EdukasiMapper:     EdukasiMapper,
		EdukasiRepository: EdukasiRepository,
	}
}

// GET USECASE
func (eu *edukasiUseCase) OnGetEdukasiTerintegrasiUseCase(NORM string) (res []dto.ResponseEdukasiTerintegrasi, err error) {
	edukasi, er11 := eu.EdukasiRepository.OnGetEdukasiTerintegrasiByNoRMRepository(NORM)

	if er11 != nil {
		return make([]dto.ResponseEdukasiTerintegrasi, 0), er11
	}

	if len(edukasi) == 0 {
		return make([]dto.ResponseEdukasiTerintegrasi, 0), nil
	}

	if len(edukasi) > 0 {
		edukasi := eu.EdukasiMapper.ToMappingEdukasiTerintegrasi(edukasi)

		return edukasi, nil
	}
	return make([]dto.ResponseEdukasiTerintegrasi, 0), nil
}

func (eu *edukasiUseCase) OnUpdateEdukasiTerintegrasiUseCase(req dto.ReqOnUpdateEdukasiTerintegrasi) (message string, err error) {
	var data = edukasiterintegrasi.DedukasiTerintegrasi{
		Informasi:         req.Informasi,
		PemberiInformasi:  req.PemberiInformasi,
		PenerimaInformasi: req.PenerimaInformasi,
		Metode:            req.Metode,
		Evaluasi:          req.Evaluasi,
	}

	_, err1 := eu.EdukasiRepository.OnUpdateEdukasiTerintegrasiRepository(req.ID, data)

	if err1 != nil {
		return "Data gagal diubah", err
	}

	return "Data berhasil diubah", nil
}

func (eu *edukasiUseCase) OnSaveEdukasiTerintegrasiUseCase(UserID string, KDBagian string, req dto.ReqEdukasiTerintegrasi) (message string, er error) {
	var data = edukasiterintegrasi.DedukasiTerintegrasi{
		InsertDttm:        time.Now().Format("2006-01-02 15:04:05"),
		InsertUser:        UserID,
		PemberiInformasi:  req.PemberiInformasi,
		Informasi:         req.Informasi,
		PenerimaInformasi: req.PenerimaInformasi,
		Evaluasi:          req.Evaluasi,
		Metode:            req.Metode,
		KdBagian:          KDBagian,
		NoRm:              req.NORM,
		NoReg:             req.NoReg,
	}

	_, er11 := eu.EdukasiRepository.OnSaveEdukasiTerintegrasiRepository(data)

	if er11 != nil {
		return "Data gagal disimpan", er11
	}

	return "data berhasil disimpan", nil
}
