package entity

import (
	"vicore_hrd/modules/hrd"
	"vicore_hrd/modules/hrd/dto"
)

// MAPPER VICORE HRD
type VicoreHRDMapper interface{}

// REPOSITORY VICORE HRD
type VicoreHRDRepository interface {
	FindHRDByEmailRepository(EmailStr string) (res hrd.Kemployee, err error)
	OnFindPelayananRepository(KDPelayanan string) (res hrd.KPelayanan, err error)
	OnFindAllDataKaryawanRepository() (res []hrd.Kemployee, err error)
	OnGetDataDokterRepo(NoReg string) (res hrd.Dokter, err error)
}

// USECASE VICORE HRD
type VicoreHRDUseCase interface {
	OnLoginUserByEmailAndPasswordUseCase(Email string, Password string, baseURL string) (user dto.ResponseDataUser, message string, err error)
}
