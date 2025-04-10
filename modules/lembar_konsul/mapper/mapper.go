package mapper

import (
	"fmt"
	"strings"
	"time"
	"vicore_hrd/app/rest"
	"vicore_hrd/modules/asesmen"
	asesmenDTO "vicore_hrd/modules/asesmen/dto"
	"vicore_hrd/modules/hrd"
	lembarKonsul "vicore_hrd/modules/lembar_konsul"
	"vicore_hrd/modules/lembar_konsul/dto"
	entity "vicore_hrd/modules/lembar_konsul/entity"
	resumemedis "vicore_hrd/modules/resume_medis"
)

type lembarKonsulMapper struct {
}

func NewLembarKonsulMapperImpl() entity.LembarKonsulMapper {
	return &lembarKonsulMapper{}
}

func (tm *lembarKonsulMapper) ToMappingLembarKonsul(profilePasien resumemedis.DProfilePasien, dregister lembarKonsul.DRegister, konsulPasien lembarKonsul.DKonsulPasien, dRekam resumemedis.DataDRekamMedis, Dokter hrd.Dokter, Pelayanan hrd.KPelayanan, asesmen []asesmen.DiagnosaResponse, cppKonsulen lembarKonsul.CpptKonsulen) (res dto.DataReponseLembarKonsule) {

	birthDate, _ := time.Parse("2006-01-02", profilePasien.Tgllahir.Format("2006-01-02"))

	now := time.Now()
	age := now.Sub(birthDate)
	years := int(age.Hours() / 24 / 365)
	months := int(age.Hours()/24/30) % 12
	days := int(age.Hours()/24) % 30
	namaPasien := profilePasien.Firstname
	tglLahir := profilePasien.Tgllahir.Format("2006-01-02")
	jenisKelamin := profilePasien.Jeniskelamin
	tglKonsul, _ := rest.UbahTanggalIndoAndTime(konsulPasien.InsertDttm)
	tglJawabanKonsul, _ := rest.UbahTanggalIndoAndTime(cppKonsulen.InsertDttm)

	return dto.DataReponseLembarKonsule{
		JawabanKonsulen: dto.JawabanKonsul{
			Tanggal:  tglJawabanKonsul,
			Terapi:   cppKonsulen.Plan,
			Penemuan: cppKonsulen.Subjektif,
			Anjuran:  cppKonsulen.Asesmen,
		},
		LembarKonsul: dto.ReponseLembarKonsul{
			NamaPasien:            profilePasien.Firstname,
			TanggalLahir:          profilePasien.Tgllahir.Format("2006-01-02"),
			NomorRekamMedis:       profilePasien.Id,
			NoReg:                 dregister.Noreg,
			Umur:                  fmt.Sprintf("%d tahun, %d bulan, %d hari", years, months, days),
			MohonKonulstasiPasien: profilePasien.Firstname,
			JenisKonsultasi:       konsulPasien.JenisKonsul,
			Tanggal:               tglKonsul,
			IktisarKlinik:         konsulPasien.IktisarKlinik,
			DokterMemintaKonsul:   Dokter.Namadokter,
			Dokter:                konsulPasien.KonsulKe,
			Ruangan:               strings.ToUpper(Pelayanan.Bagian),
			DiagnosaKerja:         asesmen,
		},
		ProfilePasien: asesmenDTO.DataProfilePasien{
			NamaPasien:   namaPasien,
			TanggalLahir: tglLahir,
			JenisKelamin: jenisKelamin,
			NoRm:         profilePasien.Id,
			NoReg:        dregister.Noreg,
			Ruangan:      strings.ToUpper(dRekam.Bagian),
		},
	}

}

func (tm *lembarKonsulMapper) ToMappingLembarKonsulV2(profilePasien resumemedis.DProfilePasien, asesmen []asesmen.DiagnosaResponse, DokterKonsulens []dto.KonsulanDokter, NoReg string, DataCPPTKonsulen []lembarKonsul.CpptKonsulen) (res dto.DataReponseLembarKonsuleV2) {

	birthDate, _ := time.Parse("2006-01-02", profilePasien.Tgllahir.Format("2006-01-02"))

	now := time.Now()
	age := now.Sub(birthDate)
	years := int(age.Hours() / 24 / 365)
	months := int(age.Hours()/24/30) % 12
	days := int(age.Hours()/24) % 30
	var DokterDPJP = ""
	var Bagian = ""

	if len(DokterKonsulens) > 0 {
		DokterDPJP = DokterKonsulens[0].DokterMemintaKonsul
		Bagian = DokterKonsulens[0].Ruangan
	}

	if len(DokterKonsulens) == 0 {
		DokterDPJP = ""
		Bagian = ""
	}

	tglLahir := profilePasien.Tgllahir.Format("2006-01-02")
	jenisKelamin := profilePasien.Jeniskelamin

	return dto.DataReponseLembarKonsuleV2{
		LembarKonsul: dto.ReponseLembarKonsulV2{
			NamaPasien:            profilePasien.Firstname,
			TanggalLahir:          profilePasien.Tgllahir.Format("2006-01-02"),
			NomorRekamMedis:       profilePasien.Id,
			NoReg:                 NoReg,
			DokterDpjp:            DokterDPJP,
			Ruangan:               Bagian,
			Umur:                  fmt.Sprintf("%d tahun, %d bulan, %d hari", years, months, days),
			MohonKonulstasiPasien: profilePasien.Firstname,
			DiagnosaKerja:         asesmen,
			KonsulanDokter:        DokterKonsulens,
		},
		ProfilePasien: asesmenDTO.DataProfilePasien{
			NoRm:         profilePasien.Id,
			NamaPasien:   profilePasien.Firstname,
			TanggalLahir: tglLahir,
			JenisKelamin: jenisKelamin,
			NoReg:        NoReg,
			Ruangan:      Bagian,
		},
		JawabanKonsulen: toJawabanKonsul(DataCPPTKonsulen, asesmen),
	}

}

func toJawabanKonsul(DataCPPTKonsulen []lembarKonsul.CpptKonsulen, _ []asesmen.DiagnosaResponse) (res []dto.JawabanKonsul) {

	if len(DataCPPTKonsulen) > 0 {
		for i := range DataCPPTKonsulen {
			tglKonsul, _ := rest.UbahTanggalIndoAndTime(DataCPPTKonsulen[i].InsertDttm)

			res = append(res, dto.JawabanKonsul{
				KonsulKe:   i + 1,
				Tanggal:    tglKonsul,
				Terapi:     DataCPPTKonsulen[i].Plan,
				Penemuan:   DataCPPTKonsulen[i].Subjektif,
				Anjuran:    DataCPPTKonsulen[i].Asesmen,
				NamaDokter: DataCPPTKonsulen[i].InsertUserId,
			})
		}

		return res

	}

	if len(DataCPPTKonsulen) == 0 {
		res = make([]dto.JawabanKonsul, 0)

		return res
	}

	return res
}
