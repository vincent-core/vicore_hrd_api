package entity

import (
	"vicore_hrd/modules/antrean"
	"vicore_hrd/modules/antrean/dto"
)

type AntreanMapper interface {
	ToMappingDataDResiterPasien(data []antrean.DRegisterPasien) (res []dto.DataResponseRegisterPasien)
	ToMappingPasienRANAP(data []antrean.PasienRANAP) (res []dto.DataResponseRegisterPasien)
}

type AntreanRepository interface {
	OnGetPengkajianDokterRepo(NoReg string) (res antrean.DAsesmenDokter, err error)
	GetAntrianUGD() (res []antrean.AntrianPoliIGD, err error)
	GetAntrianIGDDokterUmumRepository(KodeDokter string, dateFrom string, dateTo string) (res []antrean.AntrianPoliIGD, err error)
	GetPasienBangsalForDokter(kodeBangsal string, kodeDokter string, dateFrom string, dateTo string) (res []antrean.KbangsalKasur, err error)
	GetPasienBangsal(kodeBangsal string, dateFrom string, dateTo string) (res []antrean.KbangsalKasur, err error)
	OnGetDataRegisterPasienByID(ID string) (res []antrean.DRegisterPasien, err error)
	OnGetDataPasienPulangByDateRepository(date string) (res []antrean.DRegisterPasien, err error)
	OnGetDRekamPasienRANAPRepo(ID string) (res []antrean.PasienRANAP, err error)

	// DATA SINGLE DREGISTER
	OnGetSingleRegisterDataPasienRepository(NoReg string) (res antrean.DRegisterPasien, err error)
	OnGetPengkajianPerawatRepo(NoReg string) (res antrean.DepengkajianKeperawatan, err error)
}

type AntreanUseCase interface {
	OnGetAntrianIGDUseCase(req dto.GetAntranPasien, person string, userID string) (res []dto.AntrianPasien, message string, err error)
	OnGetPasienPulangByDateUseCase(date string) (res []dto.AntrianPasien, err error)
	OnDashboardUseCase(modulID string) (res dto.ResponseDashboard, err error)
}
