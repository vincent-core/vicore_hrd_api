package usecase

import (
	"log"
	"vicore_hrd/app/rest"
	"vicore_hrd/modules/asesmen"
	asesmenEntity "vicore_hrd/modules/asesmen/entity"
	lembarkonsul "vicore_hrd/modules/lembar_konsul"
	"vicore_hrd/modules/lembar_konsul/dto"
	entity "vicore_hrd/modules/lembar_konsul/entity"
	resumeEntity "vicore_hrd/modules/resume_medis/entity"

	hrdEntity "vicore_hrd/modules/hrd/entity"

	"github.com/sirupsen/logrus"
)

type lembarKonsulUseCase struct {
	logging                *logrus.Logger
	lembarKonsulMapper     entity.LembarKonsulMapper
	lembarKonsulRepository entity.LembarKonsulRepository
	resumeMedisRepository  resumeEntity.ResumeMedisRepository
	hrdRepository          hrdEntity.VicoreHRDRepository
	asesmenRepository      asesmenEntity.AsesmenRepository
}

func NewLembarKonsulUseCase(logging *logrus.Logger, lembarRepository entity.LembarKonsulRepository, lembarKonsulMapper entity.LembarKonsulMapper, resumeMedisRepository resumeEntity.ResumeMedisRepository, hrdRepository hrdEntity.VicoreHRDRepository, asesmenRepository asesmenEntity.AsesmenRepository) entity.LembarKonsulUseCase {
	return &lembarKonsulUseCase{
		logging:                logging,
		lembarKonsulRepository: lembarRepository,
		lembarKonsulMapper:     lembarKonsulMapper,
		resumeMedisRepository:  resumeMedisRepository,
		hrdRepository:          hrdRepository,
		asesmenRepository:      asesmenRepository,
	}
}

func (tu *lembarKonsulUseCase) OnGetReportLembarKonsulUseCase(Noreg string, BaseURL string) (res dto.DataReponseLembarKonsule, err error) {
	// GET ASESEMEN DOKTER
	dataRegister, _ := tu.lembarKonsulRepository.OnGetDataRegisterRepository(Noreg)
	profilePasien, _ := tu.resumeMedisRepository.OnGetDataProfilePasienRespository(dataRegister.Id)
	konsulPasien, _ := tu.lembarKonsulRepository.OnGetDataKonsulenPasienRepo(Noreg)
	dataDRekam, _ := tu.resumeMedisRepository.OnGetDataDRekamRepository(Noreg)

	// GET DATA DOKTER
	dokter, _ := tu.hrdRepository.OnGetDataDokterRepo(konsulPasien.DokterKonsul)
	// GET RUANGAN
	pelayanan, _ := tu.hrdRepository.OnFindPelayananRepository(konsulPasien.KdBagian)
	asesmen, _ := tu.asesmenRepository.GetDiagnosaRanap(Noreg)
	// asesmen, _ := tu.asesmenRepository.OnGetDiagnosaByNoRegANDKdBagianPelayananRANAP(Noreg, pelayanan.KdBag)

	cpptKonsulen := lembarkonsul.CpptKonsulen{}

	if konsulPasien.KonsulKe != "" {
		daata, _ := tu.lembarKonsulRepository.OnGetCPPTKonsulenRepository(Noreg, konsulPasien.KonsulKe[0:5])
		cpptKonsulen = daata
	}

	if konsulPasien.KonsulKe == "" {
		cpptKonsulen = lembarkonsul.CpptKonsulen{}
	}

	// GET CPPT KONSULEN

	mapper := tu.lembarKonsulMapper.ToMappingLembarKonsul(profilePasien, dataRegister, konsulPasien, dataDRekam, dokter, pelayanan, asesmen, cpptKonsulen)
	return mapper, nil
}

func (tu *lembarKonsulUseCase) OnGetReportLembarKonsulUseCaseV2(Noreg string, BaseURL string) (res dto.DataReponseLembarKonsuleV2, err error) {
	// GET ASESEMEN DOKTER
	dataRegister, _ := tu.lembarKonsulRepository.OnGetDataRegisterRepository(Noreg)
	profilePasien, _ := tu.resumeMedisRepository.OnGetDataProfilePasienRespository(dataRegister.Id)
	konsulPasien, _ := tu.lembarKonsulRepository.OnGetListDataKonsulenPasienRepo(Noreg)

	var DokterKonsulens []dto.KonsulanDokter
	var AsesmenDokter []asesmen.DiagnosaResponse
	var DataCPPTKonsulen []lembarkonsul.CpptKonsulen

	if len(konsulPasien) == 0 {
		DokterKonsulens = make([]dto.KonsulanDokter, 0)
	}

	if len(konsulPasien) > 0 {
		for i := 0; i < len(konsulPasien); i++ {
			dokterKonsul, _ := tu.hrdRepository.OnGetDataDokterRepo(konsulPasien[i].DokterKonsul)
			bagian, _ := tu.hrdRepository.OnFindPelayananRepository(konsulPasien[i].KdBagian)
			tglKonsul, _ := rest.UbahTanggalIndoAndTime(konsulPasien[i].InsertDttm)

			DokterKonsulens = append(DokterKonsulens, dto.KonsulanDokter{
				DokterKonsul:        konsulPasien[i].KonsulKe,
				DokterMemintaKonsul: dokterKonsul.Namadokter,
				IktisarKlinik:       konsulPasien[i].IktisarKlinik,
				JenisKonsultasi:     konsulPasien[i].JenisKonsul,
				Spesialisasi:        dokterKonsul.Spesialisasi,
				Tanggal:             tglKonsul,
				KonsuleKe:           i + 1,
				Ruangan:             bagian.Bagian,
			})

			if konsulPasien[i].KonsulKe != "" {
				data, err := tu.lembarKonsulRepository.OnGetCPPTKonsulenRepository(Noreg, konsulPasien[i].KonsulKe[0:5])
				namaDokter, _ := tu.hrdRepository.OnGetDataDokterRepo(konsulPasien[i].KonsulKe[0:5])

				if err != nil {
					log.Println("Error fetching CPPT Konsulen:", err)
				} else {
					DataCPPTKonsulen = append(DataCPPTKonsulen, lembarkonsul.CpptKonsulen{
						Ppa:          data.Ppa,
						InsertDttm:   data.InsertDttm,
						InsertUserId: namaDokter.Namadokter,
						KdBagian:     data.KdBagian,
						Subjektif:    data.Subjektif,
						Objektif:     data.Objektif,
						Asesmen:      data.Asesmen,
						Plan:         data.Plan,
					})
				}

			}

		}

		asesmen, _ := tu.asesmenRepository.GetDiagnosaRanap(Noreg)

		// GetDiagnosaRanap(NoReg string) (res []asesmen.DiagnosaResponse, err error)

		AsesmenDokter = append(AsesmenDokter, asesmen...)

	}

	// GET CPPT KONSULEN
	if AsesmenDokter == nil {
		AsesmenDokter = make([]asesmen.DiagnosaResponse, 0)
	}

	if len(AsesmenDokter) == 0 {
		AsesmenDokter = make([]asesmen.DiagnosaResponse, 0)
	}

	mapper := tu.lembarKonsulMapper.ToMappingLembarKonsulV2(profilePasien, AsesmenDokter, DokterKonsulens, Noreg, DataCPPTKonsulen)
	return mapper, nil
}
