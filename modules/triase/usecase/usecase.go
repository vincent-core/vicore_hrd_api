package usecase

import (
	resumeEntity "vicore_hrd/modules/resume_medis/entity"
	"vicore_hrd/modules/triase/dto"
	"vicore_hrd/modules/triase/entity"

	"github.com/sirupsen/logrus"
)

type triaseUsecase struct {
	logging          *logrus.Logger
	triaseMapper     entity.TriaseMapper
	resumeRepository resumeEntity.ResumeMedisRepository
	triaseRepository entity.TriaseRepository
}

func NewTriaseUsecase(logging *logrus.Logger, triaseMapper entity.TriaseMapper, resumeRepository resumeEntity.ResumeMedisRepository, triaseRepo entity.TriaseRepository) entity.TriaseUseCase {
	return &triaseUsecase{
		logging:          logging,
		triaseMapper:     triaseMapper,
		triaseRepository: triaseRepo,
		resumeRepository: resumeRepository,
	}
}

func (tu *triaseUsecase) OnGetReportTriaseUseCase(Noreg string, BaseURL string) (res dto.ResponseTriase, err error) {

	triaseData, _ := tu.triaseRepository.OnGetReportTriaseIGDDokterRepository(Noreg)
	asesmenData, _ := tu.triaseRepository.OnGetReportAsesmenTriaseIGDRepository(Noreg)

	mapperData := tu.triaseMapper.ToResponseTriaseIGDDokter(triaseData, asesmenData)

	dataRegister, _ := tu.resumeRepository.GetDataTanggalMasukPasienRepository(Noreg)
	profile, _ := tu.resumeRepository.OnGetDataProfilePasienRespository(dataRegister.Id)
	vital, _ := tu.triaseRepository.GetDVitalSignGDRepository(Noreg)
	fisik, _ := tu.triaseRepository.OnGetPemfisikTriaseRepo(Noreg)

	triaseAsesmen, _ := tu.triaseRepository.TriaseAsesmenRepo(Noreg)
	asesmen, _ := tu.triaseRepository.OnGetAsesmenNyeriRepo(Noreg)
	pengkajian, _ := tu.triaseRepository.OnGetAsesmenKeperawatanRepo(Noreg)
	triaseModel, _ := tu.triaseRepository.GetAsesmenDokterTriaseRepo(Noreg)
	triase, _ := tu.triaseRepository.OnGetSkalaTriaseRepo(Noreg)

	// MAPPING TRIEASE
	mapper := tu.triaseMapper.ToTriaseMapper(dataRegister, profile, fisik, vital, triaseAsesmen, asesmen, pengkajian, BaseURL, mapperData, triaseModel, asesmenData, triase)
	return mapper, nil
}

func (tu *triaseUsecase) OnGetReportTriasePonekUseCase(Noreg string, BaseURL string) (res dto.ResponseTriase, err error) {

	triaseData, _ := tu.triaseRepository.OnGetReportTriasePonekRepository(Noreg)
	asesmenData, _ := tu.triaseRepository.OnGetReportAsesmenTriasePONEKRepository(Noreg)

	mapperData := tu.triaseMapper.ToResponseTriaseIGDDokter(triaseData, asesmenData)

	dataRegister, _ := tu.resumeRepository.GetDataTanggalMasukPasienRepository(Noreg)
	profile, _ := tu.resumeRepository.OnGetDataProfilePasienRespository(dataRegister.Id)

	vital, _ := tu.triaseRepository.GetDVitalSignPONEKRepository(Noreg)
	fisik, _ := tu.triaseRepository.OnGetPemfisikTriasePONEKRepo(Noreg)

	triaseAsesmen, _ := tu.triaseRepository.TriaseAsesmenRepoKebidanan(Noreg, "PONEK")
	asesmen, _ := tu.triaseRepository.OnGetAsesmenNyeriPONEKRepo(Noreg)
	pengkajian, _ := tu.triaseRepository.OnGetAsesmenKeperawatanPONEKRepo(Noreg)

	triaseModel, _ := tu.triaseRepository.GetAsesmenDokterTriasPONEKeRepo(Noreg)

	triase, _ := tu.triaseRepository.OnGetSkalaTriasePONEKRepo(Noreg)

	// MAPPING TRIEASE
	mapper := tu.triaseMapper.ToTriaseMapper(dataRegister, profile, fisik, vital, triaseAsesmen, asesmen, pengkajian, BaseURL, mapperData, triaseModel, asesmenData, triase)

	return mapper, nil
}
