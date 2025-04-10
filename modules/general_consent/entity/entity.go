package entity

import (
	generalconsent "vicore_hrd/modules/general_consent"
	"vicore_hrd/modules/general_consent/dto"
)

type GeneralConsentRepository interface {
	OnGetPasienGeneralConsentRepository(ID string) (res generalconsent.Pasien, err error)
	OnGetPengkajianKeperawatanRepository(NoReg string, KDBagian string) (res generalconsent.DataPengkajianKeperawatan, er error)
	OnGetDataDRegisterRepository(NoReg string) (res generalconsent.DataRegister, err error)
	OnGetPengkajianKeperawatanForGeneralConsent(NoReg string, KDBagian string) (res generalconsent.PengkajianKeperawatan, err error)

	// GENERAL CONSENT
	OnGetPengkajianKeperawatanCodeRepository(NoReg string, KDBagian string) (res generalconsent.Pengkajian, err error)
	OnGetGeneralConsentRAJALRepository(noRM string, kdBagian string) (res generalconsent.DGeneralConsent, err error)
	OnGetDataGeneralConsentRepo(noRM string) (res generalconsent.DGeneralConsent, err error)
	OnGetNamaDokterRepo(NoReg string) (res generalconsent.DokterAsesmen, err error)
	OnGetDataGeneralConsentRepoV2(NoReg string, Pelayanan string, ID string) (res generalconsent.DGeneralConsent, err error)
}

type GeneralConsentUseCase interface {
	OnGetGeneralConsentRAJALUseCase(Noreg string, KDBagian string) (res dto.ToResponseGeneralConsentRAJAL)
	OnGetGeneralConsentRANAPUseCase(Noreg string, KDBagian string) (res dto.ToResponseGeneralConsentRANAP)
	OnGetGeneralConsentRAJALUseCaseV2(Noreg string, ID string) (res dto.ToResponseGeneralConsentRAJAL)
}

type GeneralConsentMapper interface {
	TOMappingGeneralConsent(data generalconsent.Pasien, pengkajian generalconsent.Pengkajian, general generalconsent.DGeneralConsent) (res dto.ToResponseGeneralConsentRAJAL)
	TOMappingGeneralConsentRANAP(data generalconsent.Pasien, pengkajian generalconsent.Pengkajian, general generalconsent.DGeneralConsent, dokter generalconsent.DokterAsesmen) (res dto.ToResponseGeneralConsentRANAP)
	TOMappingGeneralConsentRAJAL(data generalconsent.Pasien, pengkajian generalconsent.Pengkajian, general generalconsent.DGeneralConsent) (res dto.ToResponseGeneralConsentRAJAL)
}
