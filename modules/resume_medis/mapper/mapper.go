package mapper

import (
	"strings"
	"vicore_hrd/app/rest"
	"vicore_hrd/modules/antrean"
	resumemedis "vicore_hrd/modules/resume_medis"
	"vicore_hrd/modules/resume_medis/dto"
	"vicore_hrd/modules/resume_medis/entity"
)

type ResumeMedisMapper struct {
}

func NewResumeMedisMapper() entity.ResumeMapper {
	return &ResumeMedisMapper{}
}

func (im *ResumeMedisMapper) ToMappingResumeMedis(dregister resumemedis.DRegisterPasien, pasien resumemedis.DProfilePasien, bangsal resumemedis.DBangsalRep, drekap resumemedis.Drekap, keluarObat []resumemedis.DataKeluarObat1, labor []resumemedis.ResHasilLaborTableLama, radiologi []resumemedis.RegHasilRadiologiTabelLama, fisio []resumemedis.RegHasilRadiologiTabelLama, gizi []resumemedis.RegHasilRadiologiTabelLama, dpemfisik resumemedis.PemFisik, diagnosa []resumemedis.DiagnosaResponse, riwayat resumemedis.RiwayatPenyakit, tindakan []resumemedis.TindakanResponse, diagnosa1 []resumemedis.DiagnosaResponse, diagnosaDariRekam []resumemedis.DiagnosaResponse, diagnosaDokter []resumemedis.DiagnosaResponse) (res dto.ResponseResumeMedis) {

	var labors = []resumemedis.ResHasilLaborTableLama{}
	var radiologis = []resumemedis.RegHasilRadiologiTabelLama{}
	var fisios = []resumemedis.RegHasilRadiologiTabelLama{}
	var gizis = []resumemedis.RegHasilRadiologiTabelLama{}
	var diagnosas = []resumemedis.DiagnosaResponse{}
	var tindakans = []resumemedis.TindakanResponse{}

	if len(labor) == 0 {
		labors = make([]resumemedis.ResHasilLaborTableLama, 0)
	}

	if len(labor) > 0 {
		labors = labor
	}

	if len(radiologi) == 0 {
		radiologis = make([]resumemedis.RegHasilRadiologiTabelLama, 0)
	}

	if len(radiologi) > 0 {
		radiologis = radiologi
	}

	if len(fisio) == 0 {
		fisios = make([]resumemedis.RegHasilRadiologiTabelLama, 0)
	}

	if len(fisio) > 0 {
		fisios = fisio
	}

	if len(gizi) == 0 {
		gizis = make([]resumemedis.RegHasilRadiologiTabelLama, 0)
	}

	if len(gizi) > 0 {
		gizis = gizi
	}

	if len(tindakan) == 0 {
		tindakans = make([]resumemedis.TindakanResponse, 0)
	}

	if len(tindakan) > 0 {
		tindakans = tindakan
	}

	if len(diagnosa) > 0 {
		diagnosas = diagnosa
	}

	if len(diagnosa) == 0 {
		if len(diagnosa1) > 0 {
			diagnosas = diagnosa1
		}

		if len(diagnosa1) == 0 && len(diagnosa) == 0 {
			if len(diagnosaDariRekam) > 0 {
				diagnosas = diagnosaDariRekam
			}
		}

		if len(diagnosa1) == 0 && len(diagnosa) == 0 && len(diagnosaDariRekam) == 0 {
			if len(diagnosaDokter) > 0 {
				diagnosas = diagnosaDokter
			}
		}

	}
	// []resumemedis.DiagnosaResponse
	if len(diagnosa1) == 0 && len(diagnosa) == 0 && len(diagnosaDariRekam) == 0 && len(diagnosaDokter) == 0 {
		diagnosas = make([]resumemedis.DiagnosaResponse, 0)
	}

	return dto.ResponseResumeMedis{
		TanggalLahir:       pasien.Tgllahir.Format("2006-01-02"),
		NomorRekamMedis:    dregister.Id,
		NoRegister:         dregister.Noreg,
		NamaPasien:         dregister.Nama,
		Alamat:             pasien.Alamat,
		TanggalMasuk:       dregister.TglMasuk.Format("2006-01-02"),
		JamMasuk:           dregister.JamMasuk + " WIB",
		RiwayatPenyakit:    riwayat.RiwayatSekarang,
		JenisKelamin:       pasien.Jeniskelamin,
		TanggalKeluar:      bangsal.Tanggal.Format("2006-01-02"),
		JamKeluar:          bangsal.Jam + " WIB",
		KeadaanWaktuKeluar: bangsal.Keterangan,
		DokterMerawat:      bangsal.NamaDokter,
		Ruang:              drekap.Asal,
		Pelayanan:          drekap.Pelayanan,
		Kelas:              strings.ToUpper(drekap.Kelas),
		DataObat:           im.ToMappingDataObat(keluarObat),
		Labor:              labors,
		Radiologi:          radiologis,
		Fiso:               fisios,
		Gizi:               gizis,
		PemeriksaanFisik:   im.ToMappingPemeriksaanFisik(dpemfisik),
		Diagnosa:           diagnosas,
		Tindakan:           tindakans,
	}

}

func (im *ResumeMedisMapper) ToMappingDataObat(keluarObat []resumemedis.DataKeluarObat1) (res []dto.DataKeluarObat) {

	if len(keluarObat) == 0 {
		return make([]dto.DataKeluarObat, 0)
	} else {
		for _, v := range keluarObat {
			res = append(res, dto.DataKeluarObat{
				TglKeluar: v.TglKeluar.Format("2006-01-02"),
				NamaObat:  v.NamaObat,
				Jumlah:    v.Jumlah,
			})
			return res
		}
	}

	return res
}

func (im *ResumeMedisMapper) ToMappingPemeriksaanFisik(dpemfisik resumemedis.PemFisik) (res dto.ResponsePemerikssanFisik) {
	return dto.ResponsePemerikssanFisik{
		Tb:         dpemfisik.Tb + " Cm",
		Td:         dpemfisik.Td + " Cm",
		Bb:         dpemfisik.Bb + " Kg",
		Nadi:       dpemfisik.Nadi + " per menit",
		Suhu:       dpemfisik.Suhu + "°C",
		Spo2:       dpemfisik.Spo2 + "%",
		Pernafasan: dpemfisik.Pernafasan + "per menit",
	}
}

func (im *ResumeMedisMapper) ToMappingCariDataPasienPulang(data []resumemedis.CariDataPasienPulang, jk resumemedis.DataJenis) (res []dto.ResponseCariPasien) {

	if len(data) == 0 {
		return make([]dto.ResponseCariPasien, 0)
	}

	if len(data) > 0 {
		for _, v := range data {
			tglIndo, _ := rest.UbahTanggalIndo(v.Tglproses)
			res = append(res, dto.ResponseCariPasien{
				Tglproses:    tglIndo,
				Norm:         v.Norm,
				Pelayanan:    v.Pelayanan,
				Noreg:        v.Noreg,
				Nama:         v.Nama,
				JenisKelamin: jk.JenisKelamin,
				Tgllahir:     v.Tgllahir,
			})

		}
		return res

	}
	return res
}

func (im *ResumeMedisMapper) ToMappingCariPasienDRegister(data []antrean.DRegisterPasien, pasien resumemedis.DataProfilePasien) (res []dto.ResponseCariPasien) {

	if len(data) == 0 {
		return make([]dto.ResponseCariPasien, 0)
	}

	if len(data) > 0 {
		for _, v := range data {
			tglIndo, _ := rest.UbahTanggalIndo(v.Tanggal)
			res = append(res, dto.ResponseCariPasien{
				Tglproses:    tglIndo,
				Norm:         pasien.Id,
				Pelayanan:    v.Keterangan,
				Noreg:        v.Noreg,
				Nama:         v.Nama,
				JenisKelamin: pasien.Jeniskelamin,
			})

		}
		return res

	}
	return res
}

func (im *ResumeMedisMapper) TOMappingDataCPPPasienByNoReg() (res dto.DataResponseCPPT, err error) {
	return res, nil
}

// ]antrean.DRegisterPasien
