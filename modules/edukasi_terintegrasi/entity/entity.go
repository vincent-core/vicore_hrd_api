package entity

import (
	edukasiterintegrasi "vicore_hrd/modules/edukasi_terintegrasi"
	"vicore_hrd/modules/edukasi_terintegrasi/dto"
	"vicore_hrd/modules/hrd"
)

type EdukasiTerintegrasiRepository interface {
	OnGetEdukasiTerintegrasiByNoRMRepository(NoRM string) (res []edukasiterintegrasi.DedukasiTerintegrasi, err error)
	OnGetEdukasiTerintegrasiRepository(NORM string) (res edukasiterintegrasi.DedukasiTerintegrasi, err error)
	OnSaveEdukasiTerintegrasiRepository(data edukasiterintegrasi.DedukasiTerintegrasi) (res edukasiterintegrasi.DedukasiTerintegrasi, err error)
	OnUpdateEdukasiTerintegrasiRepository(ID int, data edukasiterintegrasi.DedukasiTerintegrasi) (res edukasiterintegrasi.DedukasiTerintegrasi, err error)
}

type EdukasiTerintegrasiMapper interface {
	ToMappingEdukasiTerintegrasi(data []edukasiterintegrasi.DedukasiTerintegrasi) (res []dto.ResponseEdukasiTerintegrasi)
	ToMappingPemberiInformasi(data []hrd.Kemployee) (res []dto.ResponsePemberiInformasi)
}

type EdukasiTerinterasiUseCase interface {
	OnGetEdukasiTerintegrasiUseCase(NORM string) (res []dto.ResponseEdukasiTerintegrasi, err error)
	OnSaveEdukasiTerintegrasiUseCase(UserID string, KDBagian string, req dto.ReqEdukasiTerintegrasi) (message string, er error)
	OnUpdateEdukasiTerintegrasiUseCase(req dto.ReqOnUpdateEdukasiTerintegrasi) (message string, err error)
}
