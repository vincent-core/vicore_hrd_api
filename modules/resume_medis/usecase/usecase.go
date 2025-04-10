package usecase

import (
	resumemedis "vicore_hrd/modules/resume_medis"
	"vicore_hrd/modules/resume_medis/dto"
	"vicore_hrd/modules/resume_medis/entity"

	"github.com/sirupsen/logrus"
)

type resumeUseCase struct {
	logging          *logrus.Logger
	resumeMapper     entity.ResumeMapper
	resumeRepository entity.ResumeMedisRepository
}

func NewResumeUseCase(logging *logrus.Logger, resumeMapper entity.ResumeMapper, resumeRepository entity.ResumeMedisRepository) entity.ResumeMedisUseCase {
	return &resumeUseCase{
		logging:          logging,
		resumeMapper:     resumeMapper,
		resumeRepository: resumeRepository,
	}
}

func (ru *resumeUseCase) OnGetDiagnosaDariDRekamUseCase(NoReg string) (res []resumemedis.DiagnosaResponse) {
	// GET
	var dataDiagnosa = []resumemedis.DiagnosaResponse{}

	diagnosa, _ := ru.resumeRepository.GetDiagnosaDariDRekamRepository(NoReg)

	// GET DIAGNOSA
	if diagnosa.Noreg != "" {
		id1, _ := ru.resumeRepository.OnGetKDiagnosaRepository(diagnosa.IcdIn)

		// GABUNG KAN DATA
		if id1.Code2 != "" {
			dataDiagnosa = append(dataDiagnosa, resumemedis.DiagnosaResponse{
				Diagnosa:    id1.Code2,
				Description: id1.Description,
				Type:        "PRIMER",
				Table:       "P",
			})
		} else {
			return make([]resumemedis.DiagnosaResponse, 0)
		}

		return dataDiagnosa
	} else {
		return make([]resumemedis.DiagnosaResponse, 0)
	}
}

func (ru *resumeUseCase) OnReportdataRingkasanPulangUseCase(NoReg string) (res dto.ResponseResumeMedis) {
	// GET DATA REGISTER PASIEN
	dregister, _ := ru.resumeRepository.GetDataTanggalMasukPasienRepository(NoReg)
	pasien, _ := ru.resumeRepository.OnGetDataProfilePasienRespository(dregister.Id)
	bangsal, _ := ru.resumeRepository.OnGetPasienKeluarRepository(NoReg)

	drekap, _ := ru.resumeRepository.OnGetDBilingRekapByNoregRepository(dregister.Noreg)
	keluarObat, _ := ru.resumeRepository.OnGetDataPasienKeluarObatRepository(NoReg)
	keluarObat1, _ := ru.resumeRepository.OnGetDataKeluarObat1Repository(keluarObat.NoKeluar)

	labor, _ := ru.HistoryLaboratoriumUsecaseV2(NoReg)
	radiologi, _ := ru.HistoryRadiologiUsecaseV2(NoReg)
	fisio, _ := ru.HistoryFisioterapiUsecaseV2(NoReg)
	gizi, _ := ru.HistoryGiziUsecaseV2(NoReg)

	var kdRuangan = ""

	if len(drekap.Asal) > 3 {
		kdRuangan = drekap.Asal[:4]
	}

	// GET ASESEMEN DARI BANGSAL
	getAsesemen, _ := ru.resumeRepository.OnGetAsesmenRanapRepository(NoReg)

	//===
	fisik, _ := ru.resumeRepository.GetPemeriksaaanFisikRanapRespository(NoReg)
	diagnosa, _ := ru.resumeRepository.GetDiagnosa(NoReg)
	riwayat, _ := ru.resumeRepository.OnGetRiwayatPenyakitRepository(NoReg, kdRuangan)
	tindakan, _ := ru.resumeRepository.OnGetTindakanICD9RepositoryBangsalDokter(NoReg, kdRuangan)
	diagnosa1, _ := ru.resumeRepository.GetDiagnosaRepositoryBangsal(NoReg, getAsesemen.KdBagian)
	// GetDiagnosa(noreg string) (res []resumemedis.DiagnosaResponse, err error)

	//  get adata
	// OnGetDiagnosaDariDRekamUseCase
	diagnosaDariRekam := ru.OnGetDiagnosaDariDRekamUseCase(NoReg)
	diagnosaDokter, _ := ru.resumeRepository.GetDiagnosaDokterByNoreg(NoReg)

	// MAPPING DATA RESUME MEDIS
	mapper := ru.resumeMapper.ToMappingResumeMedis(dregister, pasien, bangsal, drekap, keluarObat1, labor, radiologi, fisio, gizi, fisik, diagnosa, riwayat, tindakan, diagnosa1, diagnosaDariRekam, diagnosaDokter)
	return mapper
}

// GET HASIL LABORRATORIUM PADA TABLE LAMA
func (ru *resumeUseCase) HistoryLaboratoriumUsecaseV2(noReg string) (res []resumemedis.ResHasilLaborTableLama, err error) {

	// GET DATA DARI TABLE LAMA
	labor, _ := ru.resumeRepository.GetPenlabTabelLamaRepository(noReg)

	var data []resumemedis.ResHasilLaborTableLama

	// GetPenLabKelompokTabelLamaRepository
	if len(labor) > 0 {
		// LAKUKAN QUERY PADA TABLE LABOR LAMA
		for i := 0; i <= len(labor)-1; i++ {
			kel, _ := ru.resumeRepository.GetPenLabKelompokTabelLamaRepository(labor[i].Jaminput, labor[i].Noreg)
			// KETIKA KELOMPOK SUDAH DAPAT LAKUKAN QUERY PEMERIKSAAN
			for a := 0; a <= len(kel)-1; a++ {
				// LAKUKAN QUERY DAPATKAN DETAIL
				// var pemeriksaans []resumemedis.DPemeriksaanLabor

				pemeriksaan, _ := ru.resumeRepository.GetPenLabPemeriksaanTabelLamaRepository(labor[i].Jaminput, labor[i].Noreg, kel[a].NamaKelompok)

				data = append(data, resumemedis.ResHasilLaborTableLama{
					Tanggal:           labor[i].Jaminput,
					NamaKelompok:      kel[a].NamaKelompok,
					DPemeriksaanLabor: pemeriksaan,
				})

			}

		}
	}

	if len(labor) < 1 {
		return data, nil
	} else {

		return data, nil

	}

}

func (ru *resumeUseCase) HistoryLaboratoriumPONEKUsecaseV2(noReg string) (res []resumemedis.ResHasilLaborTableLama, err error) {

	// GET DATA DARI TABLE LAMA
	labor, _ := ru.resumeRepository.GetPenlabTabelLamaRepository(noReg)

	var data []resumemedis.ResHasilLaborTableLama

	// GetPenLabKelompokTabelLamaRepository
	if len(labor) > 0 {
		// LAKUKAN QUERY PADA TABLE LABOR LAMA
		for i := 0; i <= len(labor)-1; i++ {
			kel, _ := ru.resumeRepository.GetPenLabKelompokTabelLamaRepository(labor[i].Jaminput, labor[i].Noreg)
			// KETIKA KELOMPOK SUDAH DAPAT LAKUKAN QUERY PEMERIKSAAN
			for a := 0; a <= len(kel)-1; a++ {
				// LAKUKAN QUERY DAPATKAN DETAIL
				// var pemeriksaans []resumemedis.DPemeriksaanLabor

				pemeriksaan, _ := ru.resumeRepository.GetPenLabPemeriksaanTabelLamaRepository(labor[i].Jaminput, labor[i].Noreg, kel[a].NamaKelompok)

				data = append(data, resumemedis.ResHasilLaborTableLama{
					Tanggal:           labor[i].Jaminput,
					NamaKelompok:      kel[a].NamaKelompok,
					DPemeriksaanLabor: pemeriksaan,
				})

			}

		}
	}

	if len(labor) < 1 {
		return data, nil
	} else {

		return data, nil

	}

}

func (hu *resumeUseCase) HistoryRadiologiUsecaseV2(noReg string) (res []resumemedis.RegHasilRadiologiTabelLama, err error) {
	// GET HASIL RADIOLOGI
	radiologi, _ := hu.resumeRepository.GetHasilRadiologiRepositoryV2(noReg)

	hu.logging.Info(radiologi)

	var data []resumemedis.RegHasilRadiologiTabelLama

	if len(radiologi) > 0 {
		for i := 0; i <= len(radiologi)-1; i++ {
			detailRad, _ := hu.resumeRepository.GetDetailHasilRadiologiOldDB(radiologi[i].Bagian, noReg, radiologi[i].Jaminput)

			data = append(data, resumemedis.RegHasilRadiologiTabelLama{
				Tanggal:              radiologi[i].Jaminput,
				NamaKelompok:         radiologi[i].Bagian,
				DHasilRadiologiOldDB: detailRad,
			})

		}
	}

	if len(radiologi) < 1 {
		return data, nil
	}

	return data, nil
}

func (hu *resumeUseCase) HistoryFisioterapiUsecaseV2(noReg string) (res []resumemedis.RegHasilRadiologiTabelLama, err error) {
	fisio, _ := hu.resumeRepository.GetHasilFisioterapiRepositoryV2(noReg)
	var data []resumemedis.RegHasilRadiologiTabelLama

	if len(fisio) > 0 {
		for i := 0; i <= len(fisio)-1; i++ {
			detailRad, _ := hu.resumeRepository.GetDetailHasilRadiologiOldDB(fisio[i].Bagian, noReg, fisio[i].Jaminput)

			data = append(data, resumemedis.RegHasilRadiologiTabelLama{
				Tanggal:              fisio[i].Jaminput,
				NamaKelompok:         fisio[i].Bagian,
				DHasilRadiologiOldDB: detailRad,
			})

		}
	}

	if len(fisio) < 1 {
		return data, nil
	}

	return data, nil
}

func (hu *resumeUseCase) HistoryGiziUsecaseV2(noReg string) (res []resumemedis.RegHasilRadiologiTabelLama, err error) {
	fisio, _ := hu.resumeRepository.GetHasilGiziRepositoryV2(noReg)

	hu.logging.Info(fisio)

	var data []resumemedis.RegHasilRadiologiTabelLama

	if len(fisio) > 0 {
		for i := 0; i <= len(fisio)-1; i++ {
			detailRad, _ := hu.resumeRepository.GetDetailHasilRadiologiOldDB(fisio[i].Bagian, noReg, fisio[i].Jaminput)
			data = append(data, resumemedis.RegHasilRadiologiTabelLama{
				Tanggal:              fisio[i].Jaminput,
				NamaKelompok:         fisio[i].Bagian,
				DHasilRadiologiOldDB: detailRad,
			})

		}
	}

	if len(fisio) < 1 {
		return data, nil
	}

	return data, nil
}
