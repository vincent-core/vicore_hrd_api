package usecase

import (
	"vicore_hrd/modules/general_consent/dto"
	"vicore_hrd/modules/general_consent/entity"

	"github.com/sirupsen/logrus"
)

type generalUsecase struct {
	logging       *logrus.Logger
	generalRepo   entity.GeneralConsentRepository
	generalMapper entity.GeneralConsentMapper
}

func NewGeneralUseCase(logging *logrus.Logger, generalRepo entity.GeneralConsentRepository, generalMapper entity.GeneralConsentMapper) entity.GeneralConsentUseCase {
	return &generalUsecase{
		logging:       logging,
		generalRepo:   generalRepo,
		generalMapper: generalMapper,
	}
}

func (eu *generalUsecase) OnGetGeneralConsentRAJALUseCase(Noreg string, KDBagian string) (res dto.ToResponseGeneralConsentRAJAL) {
	pengkajian, _ := eu.generalRepo.OnGetPengkajianKeperawatanCodeRepository(Noreg, KDBagian)
	dregister, _ := eu.generalRepo.OnGetDataDRegisterRepository(Noreg)
	pasien, _ := eu.generalRepo.OnGetPasienGeneralConsentRepository(dregister.Id)
	general, _ := eu.generalRepo.OnGetDataGeneralConsentRepo(dregister.Id)

	mapper := eu.generalMapper.TOMappingGeneralConsent(pasien, pengkajian, general)
	return mapper
}

// OnGetGeneralConsentRAJALUseCase
func (eu *generalUsecase) OnGetGeneralConsentRAJALUseCaseV2(Noreg string, ID string) (res dto.ToResponseGeneralConsentRAJAL) {
	pengkajian, _ := eu.generalRepo.OnGetPengkajianKeperawatanCodeRepository(Noreg, "IGD001")
	dregister, _ := eu.generalRepo.OnGetDataDRegisterRepository(Noreg)
	pasien, _ := eu.generalRepo.OnGetPasienGeneralConsentRepository(dregister.Id)
	general, _ := eu.generalRepo.OnGetDataGeneralConsentRepoV2(Noreg, "RAJAL", ID)

	mapper := eu.generalMapper.TOMappingGeneralConsentRAJAL(pasien, pengkajian, general)
	return mapper
}

func (eu *generalUsecase) OnGetGeneralConsentRANAPUseCase(Noreg string, KDBagian string) (res dto.ToResponseGeneralConsentRANAP) {
	pengkajian, _ := eu.generalRepo.OnGetPengkajianKeperawatanCodeRepository(Noreg, KDBagian)

	dregister, _ := eu.generalRepo.OnGetDataDRegisterRepository(Noreg)
	pasien, _ := eu.generalRepo.OnGetPasienGeneralConsentRepository(dregister.Id)
	general, _ := eu.generalRepo.OnGetDataGeneralConsentRepo(dregister.Id)
	dokter, _ := eu.generalRepo.OnGetNamaDokterRepo(Noreg)

	mapper := eu.generalMapper.TOMappingGeneralConsentRANAP(pasien, pengkajian, general, dokter)
	return mapper
}
