package entity

import (
	"vicore_hrd/modules/asesmen"
	"vicore_hrd/modules/asesmen/dto"
	generalconsent "vicore_hrd/modules/general_consent"
	resumemedis "vicore_hrd/modules/resume_medis"
)

type AsesmenMapper interface {

	// DATA MAPPING CPPT
	ToResponDataCPPT(data []asesmen.DataCPPT) (res []dto.ResponseCPPT)
	ToMappingDataCPPT(profielPasien resumemedis.DataProfilePasien, data []asesmen.DataCPPT, NoReg string, dataRekam resumemedis.DataDRekamMedis) (res dto.DataReportCPPT)
	// END DATA CPPT

	// ToMappingCPPT(data []asesmen.DataCPPT, profile resumemedis.DataDRekamMedis) (res []dto.DataReportCPPT)
	ToMappingAsesmenIGDDokter(igdDokter asesmen.AsesmenDokterIGD, diagnosa []asesmen.DiagnosaResponse, Labor []resumemedis.ResHasilLaborTableLama, radiologi []resumemedis.RegHasilRadiologiTabelLama, fisio []resumemedis.RegHasilRadiologiTabelLama, gizi []resumemedis.RegHasilRadiologiTabelLama, fisik dto.PemeriksanFisikAwalMedis, keluarObat []resumemedis.DataKeluarObat1, vitalSign asesmen.DVitalSign, riwayatKeluarga []asesmen.RiwayatPenyakit, BaseURL string, pasien resumemedis.DataProfilePasien) (res dto.ReportAsesmenDokterIGD)

	// PENGANTAR RAWAT INAP
	ToMappingDataPengantarRawatInap(pasien resumemedis.DProfilePasien, diagnosa []asesmen.DiagnosaResponse, asesmenDokter asesmen.AsesmenDokterPengantarRawatInap, fisik asesmen.DPemfisikModel, dvitalSign asesmen.DVitalSign, keluarObat []resumemedis.DataKeluarObat1, diagnosaByKdBagian []asesmen.DiagnosaResponse, asesmenIGD asesmen.AsesmenDokterIGD) (res dto.ReportPengantarRawatInap)
	ToMappingInstruksiObat(dataObat []resumemedis.DataKeluarObat1) (res []dto.InstruksiObat)
	ToMappingTandaVital(fisik dto.PemeriksanFisikAwalMedis, vitalSign asesmen.DVitalSign) (res dto.ResponseTandaVital)

	//====== LAPORAN OPERASI
	ToMappingDataLaporanBedah(NoReg string, data asesmen.Bedah, pasien generalconsent.Pasien, bedah3 []asesmen.DPenLab3, bedah2 []asesmen.Dpenlab2) (res dto.ResponseLaporanBedah)
	ToMappingAsesmenDokterRawatInap(igdDokter asesmen.AsesmenDokterIGD, diagnosa []asesmen.DiagnosaResponse, Labor []resumemedis.ResHasilLaborTableLama, radiologi []resumemedis.RegHasilRadiologiTabelLama, fisio []resumemedis.RegHasilRadiologiTabelLama, gizi []resumemedis.RegHasilRadiologiTabelLama, fisik dto.PemeriksanFisikAwalMedis, keluarObat []resumemedis.DataKeluarObat1, vitalSign asesmen.DVitalSign, keluarga []asesmen.RiwayatPenyakit, BaseURL string, pasien resumemedis.DataProfilePasien, Bagian string) (res dto.ReportAsesmenDokterIGD)
}

type AsesmenUseCase interface {
	OnReportAsesmenPONEKDokterUseCase(NoReg string, BaseURL string) (res dto.ReportAsesmenDokterIGD)
	OnSaveCPPTSOAPUseCase(modulID string, person string, userID string, req dto.InsertCPPTSOAP) (message string, err error)
	OnSaveCPPTSBARUseCase(modulID string, person string, userID string, req dto.InsertCPPTSBAR) (message string, err error)
	OnGetReportCPPTUseCase(ID string) (ToMappingAsesmenIGDDokter []dto.ResponseCPPT, err error)
	// OnGetReportCPPByNoRegTUseCase(NoReg string) (res []dto.ResponseCPPT, err error)

	// USECASE UDPATE CPPT
	OnUpdateCPPTSOAPUseCase(UserID string, req dto.OnUpdateCPPTSoapRes) (message string, err error)
	OnUpdateCPPSBARUseCase(UserID string, req dto.OnUpdateSBARRes) (message string, err error)

	// * ASESMEN IGD

	OnReportAsesmenIGDDokterUseCase(NoReg string, BaseURL string) (res dto.ReportAsesmenDokterIGD)

	// * REPORT PENGANTAR RAWAT INAP
	OnReportPengantarRawatInapUseCase(NoReg string) (res dto.ReportPengantarRawatInap, err error)
	OnReportPengantarRawatInapUseCaseV2(NoReg string, KdBagian string) (res dto.ReportPengantarRawatInap, err error)

	// =========== END DATA
	OnGetReportCPPByNoRegTUseCase(NoReg string) (res dto.DataReportCPPT, err error)
	OnReportAsesmenIGDDokterRawatInapUseCase(NoReg string, BaseURL string) (res dto.ReportAsesmenDokterIGD)
}

type AsesmenRepository interface {
	// ranap
	GetDiagnosaRanapRepo(NoReg string, Bagian string) (res []asesmen.DiagnosaResponse, err error)
	OnGetAsesmenDokterRawatInapRepository(NoReg string) (res asesmen.AsesmenDokterIGD, err error)
	OnGetAsesmenDokterPONEKRepository(NoReg string) (res asesmen.AsesmenDokterIGD, err error)
	GetDiagnosaRepositoryReportPONEK(NoReg string) (res []asesmen.DiagnosaResponse, err error)
	OnGetPemeriksaanFisikAsesmenDokterRanap(NoReg string, Bagian string) (res dto.PemeriksanFisikAwalMedis, err error)

	OnGetDiagnosaByNoRegANDKdBagianPelayananRANAP(NoReg string, kdBagian string) (res []asesmen.DiagnosaResponse, err error)
	GetDiagnosaRanap(NoReg string) (res []asesmen.DiagnosaResponse, err error)
	OnGetLaporanOperasiByNoRMRepository(ID string) (res asesmen.Bedah, err error)
	// * PENGKAJIAN
	OnGetPengkajianKeperawatanRepository(kdBagian string, pelayanan string, noReg string) (res asesmen.PengkajianKeperawatan, err error)
	OnGetAsesmenDokterRANAPRepository(noReg string, kdBagian string) (res asesmen.AsesmenDokter, err error)
	OnGeAsesmenDokterRepository(noReg string, kdBagian string) (res asesmen.AsesmenDokter, err error)
	// * CPPT REPOSITORY
	OnUpdateCPPTByIDRepository(IDCppt int, data asesmen.DCPPT) (res asesmen.DCPPT, err error)
	OnSaveCPPTRepository(data asesmen.DCPPT) (res asesmen.DCPPT, err error)
	OnGetAsesmenCPPTRepository(NoRM string) (res []asesmen.DataCPPT, err error)
	OnGetAsesmenCPPByNoRegTRepository(NoReg string) (res []asesmen.DataCPPT, err error)
	// * ASESMEN IGD REPOSITORY
	OnGetAsesmenDokter(NoReg string) (res asesmen.AsesmenDokterIGD, err error)
	OnGetAsesmenDokterIGDRepository(NoReg string) (res asesmen.AsesmenDokterIGD, err error)
	GetDiagnosaRepositoryReportIGD(NoReg string) (res []asesmen.DiagnosaResponse, err error)
	OnGetDiagnosa(NoReg string) (res []asesmen.DiagnosaResponse, err error)
	OnGetDiagnosaByNoRegANDKdBagian(NoReg string, kdBagian string) (res []asesmen.DiagnosaResponse, err error)
	OnGetDataRegisterPasienRepository(NoReg string) (res asesmen.DataPasienResgister, err error)
	OnGetAsesmenDokterRepository(NoReg string) (res asesmen.AsesmenDokterPengantarRawatInap, err error)

	// * REPORT PENGANTAR RAWAT INAP
	OnGetDPemFisikPengatarRawatInapRepository(NoReg string, KDBagian string) (res asesmen.DPemfisikModel, err error)
	OnGetVitalSignRepository(NoReg string, KDBagian string) (res asesmen.DVitalSign, err error)
	OnGetInstruksiObatRepository(NoReg string, KetBagians string) (res []asesmen.DApotikKeluarObat, err error)
	OnGetInstruksiObatRepo(NoReg string) (res []asesmen.DApotikKeluarObat, err error)
	OnGetPemeriksaanFisikAsesmenDokterIGD(NoReg string) (res dto.PemeriksanFisikAwalMedis, err error)
	OnGetRiwayatPenyakitKeluarga(NoRM string) (res []asesmen.RiwayatPenyakit, err error)

	// * LAPORAN OPERASI
	OnGetOperasiPasienRepository(NoReg string) (res asesmen.Bedah, err error)
	OnGetBedah3Repository(NomorBedah string) (res []asesmen.DPenLab3, err error)
	OnGetDiagnosaRepository(NomorBedah string) (res []asesmen.Dpenlab2, err error)
}
