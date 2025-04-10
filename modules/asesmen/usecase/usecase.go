package usecase

import (
	"time"
	antreanEntity "vicore_hrd/modules/antrean/entity"
	"vicore_hrd/modules/asesmen"
	"vicore_hrd/modules/asesmen/dto"
	"vicore_hrd/modules/asesmen/entity"
	hrdEntity "vicore_hrd/modules/hrd/entity"
	resumemedis "vicore_hrd/modules/resume_medis"
	resumeEntity "vicore_hrd/modules/resume_medis/entity"

	"github.com/sirupsen/logrus"
)

type AsesmenUseCase struct {
	Logging           *logrus.Logger
	AsesmenMapper     entity.AsesmenMapper
	AsesmenRepository entity.AsesmenRepository
	ResumeUseCase     resumeEntity.ResumeMedisUseCase
	ResumeRepository  resumeEntity.ResumeMedisRepository
	AntreanRepository antreanEntity.AntreanRepository
	HrdEntityRepo     hrdEntity.VicoreHRDRepository
}

func NewAsesmenUseCase(asesmenRepo entity.AsesmenRepository, logging *logrus.Logger, asesmenMapper entity.AsesmenMapper, resumeUsecase resumeEntity.ResumeMedisUseCase, resumeRepository resumeEntity.ResumeMedisRepository, antreanRepository antreanEntity.AntreanRepository, HrdEntityRepo hrdEntity.VicoreHRDRepository) entity.AsesmenUseCase {
	return &AsesmenUseCase{
		Logging:           logging,
		AsesmenMapper:     asesmenMapper,
		AsesmenRepository: asesmenRepo,
		ResumeUseCase:     resumeUsecase,
		ResumeRepository:  resumeRepository,
		AntreanRepository: antreanRepository,
		HrdEntityRepo:     HrdEntityRepo,
	}
}

func (iu *AsesmenUseCase) OnSaveCPPTSOAPUseCase(modulID string, person string, userID string, req dto.InsertCPPTSOAP) (message string, err error) {

	var save = asesmen.DCPPT{
		InsertDttm:   time.Now(),
		UpdDttm:      time.Now(),
		Noreg:        req.Noreg,
		Subjektif:    req.Sujektif,
		Objektif:     req.Objektif,
		Asesmen:      req.Asesmen,
		Plan:         req.Plan,
		InstruksiPpa: req.InstruksiPpa,
		InsertUserId: userID,
		InsertPc:     "WEB",
		Kelompok:     person,
		KdBagian:     req.KdBagian,
		Pelayanan:    req.Pelayanan,
		Tanggal:      time.Now(),
		Dpjp:         req.Dpjp,
	}

	_, er11 := iu.AsesmenRepository.OnSaveCPPTRepository(save)

	if er11 != nil {
		return "Data gagal disimpan", er11
	}

	return "Data berhasil disipan", nil

}

func (iu *AsesmenUseCase) OnUpdateCPPTSOAPUseCae(modulID string, person string, userID string, req dto.OnUpdateCPPTSoapRes) (message string, err error) {
	return message, nil
}

func (iu *AsesmenUseCase) OnSaveCPPTSBARUseCase(modulID string, person string, userID string, req dto.InsertCPPTSBAR) (message string, err error) {

	var save = asesmen.DCPPT{
		InsertDttm:    time.Now(),
		UpdDttm:       time.Now(),
		Noreg:         req.NoReg,
		Situation:     req.Situation,
		Background:    req.Background,
		Asesmen:       req.Asesmen,
		Recomendation: req.Recomendation,
		InsertUserId:  userID,
		InsertPc:      "WEB",
		Kelompok:      person,
		KdBagian:      req.KdBagian,
		Pelayanan:     req.Pelayanan,
		Tanggal:       time.Now(),
		Dpjp:          req.Dpjp,
	}

	_, er11 := iu.AsesmenRepository.OnSaveCPPTRepository(save)

	if er11 != nil {
		return "Data gagal disimpan", er11
	}

	return "Data berhasil disipan", nil
}

func (iu *AsesmenUseCase) OnGetReportCPPTUseCase(ID string) (res []dto.ResponseCPPT, err error) {
	data, errr := iu.AsesmenRepository.OnGetAsesmenCPPTRepository(ID)

	if errr != nil {
		return make([]dto.ResponseCPPT, 0), errr
	}

	mapper := iu.AsesmenMapper.ToResponDataCPPT(data)
	return mapper, nil
}

func (iu *AsesmenUseCase) OnUpdateCPPTSOAPUseCase(UserID string, req dto.OnUpdateCPPTSoapRes) (message string, err error) {

	data := asesmen.DCPPT{
		UpdDttm:      time.Now(),
		Subjektif:    req.Subjektif,
		Objektif:     req.Objektif,
		Asesmen:      req.Asesmen,
		Plan:         req.Plan,
		InstruksiPpa: req.InstruksiPPA,
	}

	_, err11 := iu.AsesmenRepository.OnUpdateCPPTByIDRepository(req.ID, data)

	if err11 != nil {
		return "Data tidak dapat diubah", err11
	}

	return "Data berhasil diubah", nil
}

func (iu *AsesmenUseCase) OnUpdateCPPSBARUseCase(UserID string, req dto.OnUpdateSBARRes) (message string, err error) {

	data := asesmen.DCPPT{
		UpdDttm:       time.Now(),
		Situation:     req.Situation,
		Background:    req.Background,
		Asesmen:       req.Asesmen,
		Recomendation: req.Recomendation,
		InstruksiPpa:  req.InstruksiPPA,
	}

	_, err11 := iu.AsesmenRepository.OnUpdateCPPTByIDRepository(req.ID, data)

	if err11 != nil {
		return "Data tidak dapat diubah", err11
	}

	return "Data berhasil diubah", nil
}

func (iu *AsesmenUseCase) OnReportAsesmenIGDDokterUseCase(NoReg string, BaseURL string) (res dto.ReportAsesmenDokterIGD) {
	pasienProfile, _ := iu.AsesmenRepository.OnGetDataRegisterPasienRepository(NoReg)
	pasien, _ := iu.ResumeRepository.OnGetDataProfilePasienRepository(pasienProfile.Id)

	asesmenIGD, _ := iu.AsesmenRepository.OnGetAsesmenDokterIGDRepository(NoReg)
	diagnosa, _ := iu.AsesmenRepository.GetDiagnosaRepositoryReportIGD(NoReg)
	labor, _ := iu.ResumeUseCase.HistoryLaboratoriumUsecaseV2(NoReg)
	radiologi, _ := iu.ResumeUseCase.HistoryRadiologiUsecaseV2(NoReg)
	fisio, _ := iu.ResumeUseCase.HistoryFisioterapiUsecaseV2(NoReg)
	gizi, _ := iu.ResumeUseCase.HistoryGiziUsecaseV2(NoReg)

	fisik, _ := iu.AsesmenRepository.OnGetPemeriksaanFisikAsesmenDokterIGD(NoReg)
	obat, _ := iu.AsesmenRepository.OnGetInstruksiObatRepository(NoReg, asesmenIGD.Bagian)
	keluarObat, _ := iu.OnGetAllDataAsesmenDataInstruksiObatUseCase(obat)
	vitalSign, _ := iu.AsesmenRepository.OnGetVitalSignRepository(NoReg, asesmenIGD.Bagian)
	keluarga, _ := iu.AsesmenRepository.OnGetRiwayatPenyakitKeluarga(pasien.Id)

	mapper := iu.AsesmenMapper.ToMappingAsesmenIGDDokter(asesmenIGD, diagnosa, labor, radiologi, fisio, gizi, fisik, keluarObat, vitalSign, keluarga, BaseURL, pasien)
	return mapper
}

func (iu *AsesmenUseCase) OnReportAsesmenIGDDokterRawatInapUseCase(NoReg string, BaseURL string) (res dto.ReportAsesmenDokterIGD) {
	pasienProfile, _ := iu.AsesmenRepository.OnGetDataRegisterPasienRepository(NoReg)
	pasien, _ := iu.ResumeRepository.OnGetDataProfilePasienRepository(pasienProfile.Id)

	// TODO:
	asesmenIGD, _ := iu.AsesmenRepository.OnGetAsesmenDokterRawatInapRepository(NoReg)
	diagnosa, _ := iu.AsesmenRepository.GetDiagnosaRanapRepo(NoReg, asesmenIGD.Bagian)

	labor, _ := iu.ResumeUseCase.HistoryLaboratoriumUsecaseV2(NoReg)
	radiologi, _ := iu.ResumeUseCase.HistoryRadiologiUsecaseV2(NoReg)
	fisio, _ := iu.ResumeUseCase.HistoryFisioterapiUsecaseV2(NoReg)
	gizi, _ := iu.ResumeUseCase.HistoryGiziUsecaseV2(NoReg)

	fisik, _ := iu.AsesmenRepository.OnGetPemeriksaanFisikAsesmenDokterRanap(NoReg, asesmenIGD.Bagian)
	obat, _ := iu.AsesmenRepository.OnGetInstruksiObatRepository(NoReg, asesmenIGD.Bagian)
	keluarObat, _ := iu.OnGetAllDataAsesmenDataInstruksiObatUseCase(obat)
	vitalSign, _ := iu.AsesmenRepository.OnGetVitalSignRepository(NoReg, asesmenIGD.Bagian)
	keluarga, _ := iu.AsesmenRepository.OnGetRiwayatPenyakitKeluarga(pasien.Id)
	// OnFindPelayananRepository
	pelayanan, _ := iu.HrdEntityRepo.OnFindPelayananRepository(asesmenIGD.Bagian)

	mapper := iu.AsesmenMapper.ToMappingAsesmenDokterRawatInap(asesmenIGD, diagnosa, labor, radiologi, fisio, gizi, fisik, keluarObat, vitalSign, keluarga, BaseURL, pasien, pelayanan.Bagian)
	return mapper
}

func (iu *AsesmenUseCase) OnReportAsesmenPONEKDokterUseCase(NoReg string, BaseURL string) (res dto.ReportAsesmenDokterIGD) {

	pasienProfile, _ := iu.AsesmenRepository.OnGetDataRegisterPasienRepository(NoReg)
	pasien, _ := iu.ResumeRepository.OnGetDataProfilePasienRepository(pasienProfile.Id)

	asesmenIGD, _ := iu.AsesmenRepository.OnGetAsesmenDokterPONEKRepository(NoReg)
	diagnosa, _ := iu.AsesmenRepository.GetDiagnosaRepositoryReportPONEK(NoReg)

	labor, _ := iu.ResumeUseCase.HistoryLaboratoriumUsecaseV2(NoReg)
	radiologi, _ := iu.ResumeUseCase.HistoryRadiologiUsecaseV2(NoReg)
	fisio, _ := iu.ResumeUseCase.HistoryFisioterapiUsecaseV2(NoReg)
	gizi, _ := iu.ResumeUseCase.HistoryGiziUsecaseV2(NoReg)

	fisik, _ := iu.AsesmenRepository.OnGetPemeriksaanFisikAsesmenDokterIGD(NoReg)
	obat, _ := iu.AsesmenRepository.OnGetInstruksiObatRepository(NoReg, "PONEK")
	keluarObat, _ := iu.OnGetAllDataAsesmenDataInstruksiObatUseCase(obat)
	vitalSign, _ := iu.AsesmenRepository.OnGetVitalSignRepository(NoReg, "PONEK")
	keluarga, _ := iu.AsesmenRepository.OnGetRiwayatPenyakitKeluarga(pasien.Id)

	mapper := iu.AsesmenMapper.ToMappingAsesmenIGDDokter(asesmenIGD, diagnosa, labor, radiologi, fisio, gizi, fisik, keluarObat, vitalSign, keluarga, BaseURL, pasien)
	return mapper
}

func (iu *AsesmenUseCase) OnReportPengantarRawatInapUseCase(NoReg string) (res dto.ReportPengantarRawatInap, err error) {
	// GET ASESMEN DOKTER
	pasienRegister, _ := iu.AsesmenRepository.OnGetDataRegisterPasienRepository(NoReg)
	profilePasien, _ := iu.ResumeRepository.OnGetDataProfilePasienRespository(pasienRegister.Id)
	diagnosa, _ := iu.AsesmenRepository.GetDiagnosaRepositoryReportIGD(NoReg)
	asesmenDokter, _ := iu.AsesmenRepository.OnGetAsesmenDokterRepository(NoReg)
	fisik, _ := iu.AsesmenRepository.OnGetDPemFisikPengatarRawatInapRepository(NoReg, asesmenDokter.KdBagian)
	dvitalSign, _ := iu.AsesmenRepository.OnGetVitalSignRepository(NoReg, asesmenDokter.KdBagian)
	obat, _ := iu.AsesmenRepository.OnGetInstruksiObatRepository(NoReg, asesmenDokter.KdBagian)
	// obat, _ := iu.AsesmenRepository.OnGetInstruksiObatRepo(NoReg)

	keluarObat, _ := iu.OnGetAllDataAsesmenDataInstruksiObatUseCase(obat)
	diagnosaByKdBagian, _ := iu.AsesmenRepository.OnGetDiagnosaByNoRegANDKdBagian(NoReg, asesmenDokter.KdBagian)

	asesmenIGD, _ := iu.AsesmenRepository.OnGetAsesmenDokter(NoReg)

	// MAPPING DATA
	mapper := iu.AsesmenMapper.ToMappingDataPengantarRawatInap(profilePasien, diagnosa, asesmenDokter, fisik, dvitalSign, keluarObat, diagnosaByKdBagian, asesmenIGD)

	return mapper, nil
}

func (iu *AsesmenUseCase) OnReportPengantarRawatInapUseCaseV2(NoReg string, KdBagian string) (res dto.ReportPengantarRawatInap, err error) {
	// GET ASESMEN DOKTER
	pasienRegister, _ := iu.AsesmenRepository.OnGetDataRegisterPasienRepository(NoReg)
	profilePasien, _ := iu.ResumeRepository.OnGetDataProfilePasienRespository(pasienRegister.Id)
	diagnosa, _ := iu.AsesmenRepository.GetDiagnosaRepositoryReportIGD(NoReg)
	asesmenDokter, _ := iu.AsesmenRepository.OnGetAsesmenDokterRepository(NoReg)
	fisik, _ := iu.AsesmenRepository.OnGetDPemFisikPengatarRawatInapRepository(NoReg, asesmenDokter.KdBagian)
	dvitalSign, _ := iu.AsesmenRepository.OnGetVitalSignRepository(NoReg, asesmenDokter.KdBagian)
	obat, _ := iu.AsesmenRepository.OnGetInstruksiObatRepository(NoReg, asesmenDokter.KdBagian)
	keluarObat, _ := iu.OnGetAllDataAsesmenDataInstruksiObatUseCase(obat)
	diagnosaByKdBagian, _ := iu.AsesmenRepository.OnGetDiagnosaByNoRegANDKdBagian(NoReg, asesmenDokter.KdBagian)

	asesmenIGD, _ := iu.AsesmenRepository.OnGetAsesmenDokter(NoReg)

	// MAPPING DATA
	mapper := iu.AsesmenMapper.ToMappingDataPengantarRawatInap(profilePasien, diagnosa, asesmenDokter, fisik, dvitalSign, keluarObat, diagnosaByKdBagian, asesmenIGD)

	return mapper, nil
}

func (iu *AsesmenUseCase) OnGetAllDataAsesmenDataInstruksiObatUseCase(obat []asesmen.DApotikKeluarObat) (res []resumemedis.DataKeluarObat1, err error) {

	var dataObat = []resumemedis.DataKeluarObat1{}

	for _, v := range obat {
		keluarObat, err11 := iu.ResumeRepository.OnGetDataKeluarObat1Repository(v.NoKeluar)

		if err11 != nil {
			return res, err11
		}

		dataObat = append(dataObat, keluarObat...)

	}

	return dataObat, nil
}

func (iu *AsesmenUseCase) OnGetReportCPPByNoRegTUseCase(NoReg string) (res dto.DataReportCPPT, err error) {
	// PERTAMA GET DATA DARI DREGISTER
	singlePasien, _ := iu.AntreanRepository.OnGetSingleRegisterDataPasienRepository(NoReg)
	profilePasien, _ := iu.ResumeRepository.OnGetDataProfilePasienRepository(singlePasien.Id)
	data, errr := iu.AsesmenRepository.OnGetAsesmenCPPByNoRegTRepository(NoReg)
	dataRekam, _ := iu.ResumeRepository.OnGetDataDRekamRepository(NoReg)

	if errr != nil {
		return dto.DataReportCPPT{}, errr
	}

	mapper := iu.AsesmenMapper.ToMappingDataCPPT(profilePasien, data, NoReg, dataRekam)

	return mapper, nil
}
