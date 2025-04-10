package entity

import (
	"vicore_hrd/modules/asesmen"
	"vicore_hrd/modules/hrd"
	lembarkonsul "vicore_hrd/modules/lembar_konsul"
	"vicore_hrd/modules/lembar_konsul/dto"
	resumemedis "vicore_hrd/modules/resume_medis"
)

type LembarKonsulRepository interface {
	OnGetDataRegisterRepository(NoReg string) (res lembarkonsul.DRegister, err error)
	OnGetDataKonsulPasienRepository(NoReg string) (res lembarkonsul.DKonsulPasien, err error)
	OnGetDataKonsulenPasienRepo(NoReg string) (res lembarkonsul.DKonsulPasien, err error)
	OnGetCPPTKonsulenRepository(NoReg string, UserID string) (res lembarkonsul.CpptKonsulen, err error)
	OnGetListDataKonsulenPasienRepo(NoReg string) (res []lembarkonsul.DKonsulPasien, err error)
}

type LembarKonsulUseCase interface {
	OnGetReportLembarKonsulUseCase(Noreg string, BaseURL string) (res dto.DataReponseLembarKonsule, err error)
	OnGetReportLembarKonsulUseCaseV2(Noreg string, BaseURL string) (res dto.DataReponseLembarKonsuleV2, err error)
}

type LembarKonsulMapper interface {
	ToMappingLembarKonsul(profilePasien resumemedis.DProfilePasien, dregister lembarkonsul.DRegister, konsul lembarkonsul.DKonsulPasien, dRekam resumemedis.DataDRekamMedis, Dokter hrd.Dokter, Pelayanan hrd.KPelayanan, asesmen []asesmen.DiagnosaResponse, cpptKonsulen lembarkonsul.CpptKonsulen) (res dto.DataReponseLembarKonsule)
	ToMappingLembarKonsulV2(profilePasien resumemedis.DProfilePasien, asesmen []asesmen.DiagnosaResponse, DokterKonsulens []dto.KonsulanDokter, NoReg string, DataCPPTKonsulen []lembarkonsul.CpptKonsulen) (res dto.DataReponseLembarKonsuleV2)
}
