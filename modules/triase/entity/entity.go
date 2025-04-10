package entity

import (
	resumemedis "vicore_hrd/modules/resume_medis"
	"vicore_hrd/modules/triase"
	"vicore_hrd/modules/triase/dto"
)

type TriaseMapper interface {
	// * Triase Report
	ToTriaseMapper(DRegister resumemedis.DRegisterPasien, DProfilePasien resumemedis.DProfilePasien, fisik triase.TriaseDPemFisik, vitalSign triase.DVitalSign, triaseAsesmen triase.TriaseAsesmen, nyeri triase.AsesmenUlangNyeri, asesmenNyeri triase.AsesmenKeperawatan, BaseURL string, triaseDokter dto.ResposeTriaseIGDDokter, triaseModel triase.TriaseModel, asesmed triase.AsesmenTriaseIGD, triase triase.Triase) (res dto.ResponseTriase)
	ToResponseTriaseIGDDokter(data triase.TriaseIGDDokter, data1 triase.AsesmenTriaseIGD) (res dto.ResposeTriaseIGDDokter)
}

type TriaseRepository interface {
	GetDVitalSignPONEKRepository(noreg string) (res triase.DVitalSign, err error)
	OnGetPemfisikTriasePONEKRepo(NoReg string) (res triase.TriaseDPemFisik, err error)
	GetDiagnosa(Noreg string) (res triase.DokterDiagnosa, err error)

	GetAsesmenIGD() (res []triase.IgdAsesmen, err error)
	GetNamaDokter(Noreg string) (res triase.DokterNama, err error)
	UpdateAsesmedKonsulKe(Dokter string, Noreg string) (res triase.IgdAsesmen, err error)
	UpDateDiagnosa(Diagnosa string, Noreg string, KdBagian string) (res triase.IgdAsesmen, err error)

	OnGetReportAsesmenTriasePONEKRepository(noReg string) (res triase.AsesmenTriaseIGD, err error)
	OnGetReportTriasePonekRepository(noReg string) (res triase.TriaseIGDDokter, err error)
	OnGetAsesmenTriaseIGDRepository(noReg string, userID string, kdBagian string, pelayanan string) (res triase.AsesmenTriaseIGD, err error)
	GetDVitalSignGDRepository(noreg string) (res triase.DVitalSign, err error)
	OnGetPemfisikTriaseRepo(NoReg string) (res triase.TriaseDPemFisik, err error)
	TriaseAsesmenRepo(NoReg string) (res triase.TriaseAsesmen, err error)
	TriaseAsesmenRepoKebidanan(NoReg string, KdBagian string) (res triase.TriaseAsesmen, err error)
	OnGetAsesmenNyeriRepo(NoReg string) (res triase.AsesmenUlangNyeri, err error)
	OnGetAsesmenNyeriPONEKRepo(NoReg string) (res triase.AsesmenUlangNyeri, err error)

	OnGetSkalaTriasePONEKRepo(NoReg string) (res triase.Triase, err error)

	OnGetAsesmenKeperawatanRepo(NoReg string) (res triase.AsesmenKeperawatan, err error)
	OnGetAsesmenKeperawatanPONEKRepo(NoReg string) (res triase.AsesmenKeperawatan, err error)
	GetAsesmenDokterTriasPONEKeRepo(noReg string) (res triase.TriaseModel, err error)
	OnGetReportAsesmenTriaseIGDRepository(noReg string) (res triase.AsesmenTriaseIGD, err error)
	OnGetReportTriaseIGDDokterRepository(noReg string) (res triase.TriaseIGDDokter, err error)
	OnGetSkalaTriaseRepo(NoReg string) (res triase.Triase, err error)
	GetAsesmenDokterTriaseRepo(noReg string) (res triase.TriaseModel, err error)
}

type TriaseUseCase interface {
	// TRIASE PONEK
	OnGetReportTriasePonekUseCase(Noreg string, BaseURL string) (res dto.ResponseTriase, err error)
	// END TRIASE PONEK
	// * TRIASE REPORT
	OnGetReportTriaseUseCase(Noreg string, BaseURL string) (res dto.ResponseTriase, err error)
}
