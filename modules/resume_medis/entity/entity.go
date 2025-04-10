package entity

import (
	"vicore_hrd/modules/antrean"
	resumemedis "vicore_hrd/modules/resume_medis"
	"vicore_hrd/modules/resume_medis/dto"
)

type ResumeMedisRepository interface {
	OnGetKDiagnosaRepository(Code string) (res resumemedis.DiagnosaICD, err error)
	GetDiagnosaDariDRekamRepository(NoReg string) (res resumemedis.DRekamSatu, err error)
	OnGetAsesmenRanapRepository(NoReg string) (res resumemedis.AsesmenRanap, err error)
	GetPemeriksaaanFisikRanapRespository(noreg string) (res resumemedis.PemFisik, err error)
	OnGetDBilingRekapByNoregRepository(NoReg string) (res resumemedis.Drekap, err error)
	GetDataTanggalMasukPasienRepository(NoReg string) (res resumemedis.DRegisterPasien, err error)
	OnGetDataProfilePasienRespository(ID string) (res resumemedis.DProfilePasien, err error)
	OnGetPasienKeluarRepository(NoReg string) (res resumemedis.DBangsalRep, err error)
	OnGetDBilingRekapRepository(ID string) (res resumemedis.Drekap, err error)
	OnGetDataPasienKeluarObatRepository(NoReg string) (res resumemedis.DApotikKeluarObat, err error)
	OnGetDataKeluarObat1Repository(NoKeluar string) (res []resumemedis.DataKeluarObat1, err error)
	GetPenlabTabelLamaRepository(noReg string) (res []resumemedis.DPenLab, err error)
	GetPenLabKelompokTabelLamaRepository(jamInput string, noReg string) (res []resumemedis.DPenmedPemeriksaan, err error)
	GetPenLabPemeriksaanTabelLamaRepository(jamInput string, noReg string, kelompok string) (res []resumemedis.DPemeriksaanLabor, err error)
	GetHasilRadiologiRepositoryV2(noReg string) (res []resumemedis.DHasilRadiologiV2, err error)
	GetDetailHasilRadiologiOldDB(bagian string, noReg string, jamInput string) (res []resumemedis.DHasilRadiologiOldDB, err error)
	GetHasilFisioterapiRepositoryV2(noReg string) (res []resumemedis.DHasilRadiologiV2, err error)
	GetHasilGiziRepositoryV2(noReg string) (res []resumemedis.DHasilRadiologiV2, err error)
	GetPemeriksaaanFisikRespository(noreg string) (res resumemedis.PemFisik, err error)
	GetDiagnosaRepositoryBangsal(noreg string, kdBagian string) (res []resumemedis.DiagnosaResponse, err error)
	GetDiagnosa(noreg string) (res []resumemedis.DiagnosaResponse, err error)
	OnGetRiwayatPenyakitRepository(NoReg string, KDBagian string) (res resumemedis.RiwayatPenyakit, err error)
	OnGetTindakanICD9RepositoryBangsalDokter(noReg string, kdBagian string) (res []resumemedis.TindakanResponse, err error)

	// CARI PASIEN PULANG
	CariPasienPulangRepository(Cari string) (res []resumemedis.CariDataPasienPulang, err error)
	CariPasienPulangRanapRepository(Cari string) (res []resumemedis.CariDataPasienPulang, err error)
	OnGetJenisKelamin(NOID string) (res resumemedis.DataJenis, err error)
	OnGetProfilePasien(ID string) (res resumemedis.DataJenis, err error)

	// DATA DARI DREKAM || PASIEN YANG DI DAPAT SAAT PULANG
	OnGetDataDRekamRepository(NoReg string) (res resumemedis.DataDRekamMedis, err error)
	OnGetDataProfilePasienRepository(ID string) (res resumemedis.DataProfilePasien, err error)

	// DATA PROFILE PASIEN
	GetDiagnosaDokterByNoreg(NoReg string) (res []resumemedis.DiagnosaResponse, err error)
}

type ResumeMedisUseCase interface {
	OnReportdataRingkasanPulangUseCase(NoReq string) (res dto.ResponseResumeMedis)
	HistoryLaboratoriumUsecaseV2(noReg string) (res []resumemedis.ResHasilLaborTableLama, err error)
	HistoryRadiologiUsecaseV2(noReg string) (res []resumemedis.RegHasilRadiologiTabelLama, err error)
	HistoryFisioterapiUsecaseV2(noReg string) (res []resumemedis.RegHasilRadiologiTabelLama, err error)
	HistoryGiziUsecaseV2(noReg string) (res []resumemedis.RegHasilRadiologiTabelLama, err error)
}

type ResumeMapper interface {
	ToMappingCariPasienDRegister(data []antrean.DRegisterPasien, pasien resumemedis.DataProfilePasien) (res []dto.ResponseCariPasien)
	ToMappingResumeMedis(dregister resumemedis.DRegisterPasien, pasien resumemedis.DProfilePasien, bangsal resumemedis.DBangsalRep, drekap resumemedis.Drekap, keluarObat []resumemedis.DataKeluarObat1, labor []resumemedis.ResHasilLaborTableLama, radiologi []resumemedis.RegHasilRadiologiTabelLama, fisio []resumemedis.RegHasilRadiologiTabelLama, gizi []resumemedis.RegHasilRadiologiTabelLama, pemFisik resumemedis.PemFisik, diagnosa []resumemedis.DiagnosaResponse, riwayat resumemedis.RiwayatPenyakit, tindakan []resumemedis.TindakanResponse, diagnosa1 []resumemedis.DiagnosaResponse, diagnosaDariRekam []resumemedis.DiagnosaResponse, diagnosaDokter []resumemedis.DiagnosaResponse) (res dto.ResponseResumeMedis)
	ToMappingCariDataPasienPulang(data []resumemedis.CariDataPasienPulang, jK resumemedis.DataJenis) (res []dto.ResponseCariPasien)
}
