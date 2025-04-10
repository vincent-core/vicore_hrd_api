package mapper

import (
	"fmt"
	"os"
	"strings"
	"vicore_hrd/app/rest"
	"vicore_hrd/modules/asesmen"
	"vicore_hrd/modules/asesmen/dto"
	"vicore_hrd/modules/asesmen/entity"
	generalconsent "vicore_hrd/modules/general_consent"
	resumemedis "vicore_hrd/modules/resume_medis"
)

type AsesmenMapper struct {
}

func NewAsesmenMapperImple() entity.AsesmenMapper {
	return &AsesmenMapper{}
}

func (mm *AsesmenMapper) ToMappingDataCPPT(profielPasien resumemedis.DataProfilePasien, data []asesmen.DataCPPT, NoReg string, dataRekam resumemedis.DataDRekamMedis) (res dto.DataReportCPPT) {
	var cppts = []dto.ResponseCPPT{}

	if len(data) == 0 {
		cppts = make([]dto.ResponseCPPT, 0)
	}

	if len(data) > 0 {
		for _, V := range data {
			tgl, _ := rest.UbahTanggalIndoAndTime(V.InsertDttm)

			cppts = append(cppts, dto.ResponseCPPT{
				Tanggal:       tgl,
				ID:            V.IdCppt,
				Keterangan:    toKeterangan(V.Situation, V.Subjektif, V.Background, V.Objektif, V.Plan, V.Asesmen, V.Recomendation),
				InstruksiPPA:  V.InstruksiPpa,
				DPJP:          V.NamaDokterDpjp,
				PemberiAsuhan: V.NamaProfesional,
			})
		}

	}

	return dto.DataReportCPPT{
		ProfilePasien: toProfilePasien(profielPasien, NoReg, dataRekam),
		CPPT:          cppts,
	}
}

func (mm *AsesmenMapper) ToResponDataCPPT(data []asesmen.DataCPPT) (res []dto.ResponseCPPT) {
	var cppts = []dto.ResponseCPPT{}

	if len(data) == 0 {
		return make([]dto.ResponseCPPT, 0)
	}

	if len(data) > 0 {
		for _, V := range data {
			tgl, _ := rest.UbahTanggalIndoAndTime(V.InsertDttm)
			cppts = append(cppts, dto.ResponseCPPT{
				Tanggal:       tgl,
				ID:            V.IdCppt,
				Keterangan:    toKeterangan(V.Situation, V.Subjektif, V.Background, V.Objektif, V.Plan, V.Asesmen, V.Recomendation),
				InstruksiPPA:  V.InstruksiPpa,
				DPJP:          V.NamaDokterDpjp,
				PemberiAsuhan: V.NamaProfesional,
			})
		}
	}

	return cppts
}

func toProfilePasien(profile resumemedis.DataProfilePasien, NoReg string, dataRekam resumemedis.DataDRekamMedis) (res dto.DataProfilePasien) {
	return dto.DataProfilePasien{
		TanggalLahir: rest.FormatTanggal(profile.Tgllahir),
		JenisKelamin: profile.Jeniskelamin,
		NamaPasien:   profile.Firstname,
		NoReg:        NoReg,
		Ruangan:      strings.ToUpper(dataRekam.Bagian),
	}
}

func toKeterangan(Situation string, Subjective string, Background string, Objecktive string, Plan string, Asesmen string, Rekomentation string) (response string) {
	if Situation != "" {
		cppt := fmt.Sprintf("Situation: %s\nBackground: %s\nAsesmen: %s\nRekomendation: %s", Situation, Background, Asesmen, Rekomentation)
		return cppt
	}

	if Subjective != "" {
		cppt := fmt.Sprintf("Subjektif: %s\nObjective: %s\nAsesmen: %s\nPlan:%s\n", Subjective, Objecktive, Asesmen, Plan)

		return cppt
	}

	return response
}

func (mm *AsesmenMapper) ToMappingDataPengantarRawatInap(pasien resumemedis.DProfilePasien, diagnosa []asesmen.DiagnosaResponse, asesmenDokter asesmen.AsesmenDokterPengantarRawatInap, fisik asesmen.DPemfisikModel, dvitalSign asesmen.DVitalSign, keluarObat []resumemedis.DataKeluarObat1, diagnosaByKdBagian []asesmen.DiagnosaResponse, asemenIGd asesmen.AsesmenDokterIGD) (res dto.ReportPengantarRawatInap) {
	var diagnosas = []asesmen.DiagnosaResponse{}

	if len(diagnosa) == 0 {
		diagnosas = make([]asesmen.DiagnosaResponse, 0)
	}

	if diagnosa == nil {
		diagnosas = make([]asesmen.DiagnosaResponse, 0)
	}

	if len(diagnosaByKdBagian) == 0 {
		diagnosas = make([]asesmen.DiagnosaResponse, 0)
	}

	if diagnosaByKdBagian == nil {
		diagnosas = make([]asesmen.DiagnosaResponse, 0)
	}

	if len(diagnosaByKdBagian) > 0 {
		diagnosas = diagnosaByKdBagian
	}

	if len(diagnosa) > 0 {
		diagnosas = diagnosa
	}

	tanggals, _ := rest.UbahTanggalIndo(asesmenDokter.Waktu)

	return dto.ReportPengantarRawatInap{
		NamaPasien:           pasien.Firstname,
		JenisKelamin:         pasien.Jeniskelamin,
		Alamat:               pasien.Alamat,
		NomorRM:              pasien.Id,
		Mohon:                "Perawatan",
		TanggalLahir:         pasien.Tgllahir.Format("2006-01-02"),
		Diagnosa:             diagnosas,
		KeluhanUtama:         asesmenDokter.KeluhanUtama,
		Tanggal:              tanggals,
		Bagian:               asesmenDokter.NamaBagian,
		NamaDPJP:             toNamaDPJP(asemenIGd.KonsulKe, asesmenDokter.NamaDokter, asesmenDokter.KdBagian),
		DokterPenangungJawab: asemenIGd.Dokter,
		PemeriksaanFisik:     toMappingDPemfisik(fisik, dvitalSign),
		InstruksiObat:        mm.ToMappingInstruksiObat(keluarObat),
		InstruksiNarasi:      asemenIGd.Terapi,
	}
}

func toNamaDPJP(value1 string, value2 string, bagian string) (res string) {
	if bagian == "PONEK" {
		if value1 == "" && value2 == "" {
			return ""
		}

		if value1 == "" && value2 != "" {
			return value2
		}

		if value1 != "" && value2 == "" {
			return value1
		}
	}

	if bagian != "PONEK" {

		if value1 == "" && value2 == "" {
			return ""
		}

		if value1 == "" && value2 != "" {
			return value2
		}

		if value1 != "" && value2 == "" {
			return value1
		}

		return value1
	}

	return ""
}

func (mm *AsesmenMapper) ToMappingAsesmenIGDDokter(igdDokter asesmen.AsesmenDokterIGD, diagnosa []asesmen.DiagnosaResponse, Labor []resumemedis.ResHasilLaborTableLama, radiologi []resumemedis.RegHasilRadiologiTabelLama, fisio []resumemedis.RegHasilRadiologiTabelLama, gizi []resumemedis.RegHasilRadiologiTabelLama, fisik dto.PemeriksanFisikAwalMedis, keluarObat []resumemedis.DataKeluarObat1, vitalSign asesmen.DVitalSign, keluarga []asesmen.RiwayatPenyakit, BaseURL string, pasien resumemedis.DataProfilePasien) (res dto.ReportAsesmenDokterIGD) {

	var diagnosas []asesmen.DiagnosaResponse
	var labors = []resumemedis.ResHasilLaborTableLama{}
	var radiologis = []resumemedis.RegHasilRadiologiTabelLama{}
	var fisios = []resumemedis.RegHasilRadiologiTabelLama{}
	var gizis = []resumemedis.RegHasilRadiologiTabelLama{}

	if len(diagnosa) == 0 {
		diagnosas = make([]asesmen.DiagnosaResponse, 0)
	}

	if len(diagnosa) > 0 {
		diagnosas = diagnosa
	}

	if len(Labor) == 0 {
		labors = make([]resumemedis.ResHasilLaborTableLama, 0)
	}

	if len(Labor) > 0 {
		labors = Labor
	}

	if len(radiologi) == 0 {
		radiologis = make([]resumemedis.RegHasilRadiologiTabelLama, 0)
	}

	if len(radiologi) > 0 {
		radiologis = radiologi
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

	if len(diagnosa) > 0 {
		diagnosas = diagnosa
	}

	tglIndo, _ := rest.UbahTanggalIndo(pasien.Tgllahir)

	pasiens := dto.DataProfilePasien{
		NoRm:         pasien.Id,
		TanggalLahir: tglIndo,
		JenisKelamin: pasien.Jeniskelamin,
		NamaPasien:   pasien.Firstname,
		Ruangan:      "INSTALASI GAWAT DARURAT",
		NoReg:        igdDokter.Noreg,
	}

	tglIndoAses, _ := rest.UbahTanggalIndoAndTime(igdDokter.Waktu)
	return dto.ReportAsesmenDokterIGD{
		Tanggal:          tglIndoAses,
		KeluhanUtama:     igdDokter.KeluhanUtama,
		PenyakitSekarang: igdDokter.RiwayatSekarang,
		PenyakitDahulu:   toStrip(igdDokter.RiwayatDahulu),
		Prognosis:        toPrognosis(igdDokter.Prognosis),
		Dokter:           igdDokter.Dokter,
		ImageLokalis:     toImageLokalis(igdDokter.ImageLokalis, BaseURL),
		Diagnosa:         diagnosas,
		Labor:            labors,
		Radiologi:        radiologis,
		Fiso:             fisios,
		Gizi:             gizis,
		PemeriksaanFisik: mm.ToResponseFisik(fisik),
		Planning:         mm.ToMappingInstruksiObat(keluarObat),
		VitalSign:        mm.ToMappingTandaVital(fisik, vitalSign),
		PenyakitKeluarga: toRiwayatPenyakitKeluargaString(keluarga),
		ProfilePasien:    pasiens,
		KonsulKe:         igdDokter.KonsulKe,
		Terapi:           igdDokter.Terapi,
		CaraKeluar:       igdDokter.CaraKeluar,
		CaraKeluarDetail: igdDokter.CaraKeluarDetail,
	}
}

func (mm *AsesmenMapper) ToMappingAsesmenDokterRawatInap(igdDokter asesmen.AsesmenDokterIGD, diagnosa []asesmen.DiagnosaResponse, Labor []resumemedis.ResHasilLaborTableLama, radiologi []resumemedis.RegHasilRadiologiTabelLama, fisio []resumemedis.RegHasilRadiologiTabelLama, gizi []resumemedis.RegHasilRadiologiTabelLama, fisik dto.PemeriksanFisikAwalMedis, keluarObat []resumemedis.DataKeluarObat1, vitalSign asesmen.DVitalSign, keluarga []asesmen.RiwayatPenyakit, BaseURL string, pasien resumemedis.DataProfilePasien, Bagian string) (res dto.ReportAsesmenDokterIGD) {

	var diagnosas []asesmen.DiagnosaResponse
	var labors = []resumemedis.ResHasilLaborTableLama{}
	var radiologis = []resumemedis.RegHasilRadiologiTabelLama{}
	var fisios = []resumemedis.RegHasilRadiologiTabelLama{}
	var gizis = []resumemedis.RegHasilRadiologiTabelLama{}

	if len(diagnosa) == 0 {
		diagnosas = make([]asesmen.DiagnosaResponse, 0)
	}

	if len(diagnosa) > 0 {
		diagnosas = diagnosa
	}

	if len(Labor) == 0 {
		labors = make([]resumemedis.ResHasilLaborTableLama, 0)
	}

	if len(Labor) > 0 {
		labors = Labor
	}

	if len(radiologi) == 0 {
		radiologis = make([]resumemedis.RegHasilRadiologiTabelLama, 0)
	}

	if len(radiologi) > 0 {
		radiologis = radiologi
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

	if len(diagnosa) > 0 {
		diagnosas = diagnosa
	}

	tglIndo, _ := rest.UbahTanggalIndo(pasien.Tgllahir)

	pasiens := dto.DataProfilePasien{
		NoRm:         pasien.Id,
		TanggalLahir: tglIndo,
		JenisKelamin: pasien.Jeniskelamin,
		NamaPasien:   pasien.Firstname,
		Ruangan:      Bagian,
		NoReg:        igdDokter.Noreg,
	}

	tglIndoAses, _ := rest.UbahTanggalIndoAndTime(igdDokter.Waktu)
	return dto.ReportAsesmenDokterIGD{
		Tanggal:          tglIndoAses,
		KeluhanUtama:     igdDokter.KeluhanUtama,
		PenyakitSekarang: igdDokter.RiwayatSekarang,
		PenyakitDahulu:   toStrip(igdDokter.RiwayatDahulu),
		Prognosis:        toStrip(igdDokter.Prognosis),
		Dokter:           igdDokter.Dokter,
		ImageLokalis:     toImageLokalis(igdDokter.ImageLokalis, BaseURL),
		Diagnosa:         diagnosas,
		Labor:            labors,
		Radiologi:        radiologis,
		Fiso:             fisios,
		Gizi:             gizis,
		PemeriksaanFisik: mm.ToResponseFisik(fisik),
		Planning:         mm.ToMappingInstruksiObat(keluarObat),
		VitalSign:        mm.ToMappingTandaVital(fisik, vitalSign),
		PenyakitKeluarga: toRiwayatPenyakitKeluargaString(keluarga),
		ProfilePasien:    pasiens,
		KonsulKe:         igdDokter.KonsulKe,
		Terapi:           igdDokter.Terapi,
	}
}

func toMappingDPemfisik(fisik asesmen.DPemfisikModel, vitalSign asesmen.DVitalSign) (res dto.PemeriksaanFisik) {
	return dto.PemeriksaanFisik{
		E:            fisik.GcsE,
		M:            fisik.GcsM,
		V:            fisik.GcsV,
		TekananDarah: vitalSign.Td + " mmHg",
		Hr:           vitalSign.Nadi + " per menit ",
		RR:           vitalSign.Pernafasan + " per menit",
		Sens:         strings.ToUpper(fisik.Kesadaran),
		Temp:         vitalSign.Suhu + " °C",
	}
}

func (mm *AsesmenMapper) ToMappingInstruksiObat(dataObat []resumemedis.DataKeluarObat1) (res []dto.InstruksiObat) {
	var obats = []dto.InstruksiObat{}

	if len(dataObat) > 0 {
		for _, V := range dataObat {
			obats = append(obats, dto.InstruksiObat{
				NamaObat:  V.NamaObat,
				Jumlah:    V.Jumlah,
				TglKeluar: V.TglKeluar.Format("2006-01-02"),
			},
			)
		}
	}

	if len(dataObat) == 0 {
		obats = make([]dto.InstruksiObat, 0)
	}

	if dataObat == nil {
		obats = make([]dto.InstruksiObat, 0)
	}

	return obats
}

func (mm *AsesmenMapper) ToMappingTandaVital(fisik dto.PemeriksanFisikAwalMedis, vitalSign asesmen.DVitalSign) (res dto.ResponseTandaVital) {
	return dto.ResponseTandaVital{
		GCS:        fmt.Sprintf("E %s M %s V %s", fisik.E, fisik.M, fisik.V),
		TD:         vitalSign.Td,
		Nadi:       vitalSign.Nadi + " per menit ",
		Suhu:       vitalSign.Suhu + " °C",
		Kesadaran:  strings.ToUpper(fisik.Kesadaran),
		Pernafasan: vitalSign.Pernafasan + " per menit",
		SPO2:       vitalSign.Spo2 + " %",
	}
}

func (mm *AsesmenMapper) ToResponseFisik(data dto.PemeriksanFisikAwalMedis) (res dto.ResponsePemfisik) {
	return dto.ResponsePemfisik{
		Kepala:              toDBN(data.Kepala),
		Mata:                toDBN(data.Mata),
		THT:                 toDBN(data.THT),
		Mulut:               toDBN(data.Mulut),
		Leher:               toDBN(data.Leher),
		Dada:                toDBN(data.Dada),
		Jantung:             toDBN(data.Jantung),
		Paru:                toDBN(data.Paru),
		Perut:               toDBN(data.Perut),
		Hati:                toDBN(data.Hati),
		Limpa:               toDBN(data.Limpa),
		Ginjal:              toDBN(data.Ginjal),
		ALatKelamin:         toDBN(data.ALatKelamin),
		AnggotaGerak:        toDBN(data.AnggotaGerak),
		Refleks:             toDBN(data.Refleks),
		KekuatanOtot:        toDBN(data.KekuatanOtot),
		Kulit:               toDBN(data.Kulit),
		KelenjarGetahBening: toDBN(data.KelenjarGetahBening),
		RtVt:                toDBN(data.RtVt),
	}
}

func toDBN(value string) (res string) {
	if value == "" {
		return "DBN"
	} else {
		return value
	}
}

func toRiwayatPenyakitKeluargaString(keluarga []asesmen.RiwayatPenyakit) (value string) {
	if len(keluarga) > 0 {
		var namaPenyakit []string

		for _, penyakit := range keluarga {
			namaPenyakit = append(namaPenyakit, penyakit.Alergi)
		}

		namaPenyakitString := strings.Join(namaPenyakit, ", ")

		return namaPenyakitString
	} else {
		return "-"
	}
}

func toStrip(value string) (res string) {
	if value == "" {
		return "-"
	} else {
		return value
	}
}

func toPrognosis(value string) (res string) {
	if value == "" {
		return "Baik"
	} else {
		return value
	}
}

func toImageLokalis(images string, BaseURL string) (res string) {
	if images == "" {
		return BaseURL + "/app/images/public/lokalis_default.png"
	} else {
		return os.Getenv("IMAGE_LOKALIS") + images
	}
}

func (mm *AsesmenMapper) ToMappingDataLaporanBedah(NoReg string, data asesmen.Bedah, pasien generalconsent.Pasien, bedah3 []asesmen.DPenLab3, bedah2 []asesmen.Dpenlab2) (res dto.ResponseLaporanBedah) {

	tglLahir, _ := rest.UbahTanggalIndo(pasien.TanggalLahir)
	tglOperasi, _ := rest.UbahTanggalIndo(data.Cdttm)

	var pasiens = dto.Pasien{
		Noreg:        NoReg,
		Nama:         pasien.NamaPasien,
		TanggalLahir: tglLahir,
		NoRm:         pasien.NomorRekamMedis,
		JenisKelamin: data.Gender,
	}

	return dto.ResponseLaporanBedah{
		TanggalOpersi:          tglOperasi,
		ProfilPasien:           pasiens,
		JamOperasiDimulai:      data.Mulai + " WIB",
		JamOperasiSelesai:      data.Akhir + " WIB",
		LamaOperasiBerlangsung: data.Lama + " WIB",
		Klasifiksai:            toKlasifikasi(data.Klasifikasi),
		KlasifikasiLuka:        toKlafikasiLuka(data.KlasifikasiLuka),
		UraianOperasi:          data.Uraian,
		JenisOperasi:           toJenisOperasi(data.Jenis),
		JenisJaringan:          strings.ToUpper(data.Patologi),
		PengirimanJaringan:     toPengirimanJaringan(data.AdaJaringan),
		NamaAsisten:            toAsisten(bedah3),
		NamaInstrumen:          toInstrumen(bedah3),
		NamaAhli:               toAhliBedah(bedah3),
		NamaPerawatAnastesi:    toPenataAnastesi(bedah3),
		Tindakan:               toTindakanOperasi(bedah2),
		DiagnosaPre:            toDiagnosaPre(bedah2),
		DiagnosaPost:           toDiagnosaPost(bedah2),
		NamaAhliAnastesi:       toAnastesi(bedah3),
		Tanggal:                tglOperasi,
	}
}

func toDiagnosaPre(data []asesmen.Dpenlab2) (res []dto.TindakanOperasi) {

	var tindakans = []dto.TindakanOperasi{}

	if len(data) == 0 {
		tindakans = make([]dto.TindakanOperasi, 0)
	}

	if data == nil {
		tindakans = make([]dto.TindakanOperasi, 0)
	}

	for _, V := range data {
		if V.Keterangan == "icd10" && V.Jenis == "pre" {
			tindakans = append(tindakans, dto.TindakanOperasi{
				Keterangan: V.Ket,
				Kode:       V.Kode,
				Deskripsi:  V.Diagnosa,
			})
		}
	}

	return tindakans
}

func toDiagnosaPost(data []asesmen.Dpenlab2) (res []dto.TindakanOperasi) {
	var tindakans = []dto.TindakanOperasi{}

	if len(data) == 0 {
		tindakans = make([]dto.TindakanOperasi, 0)
	}

	if data == nil {
		tindakans = make([]dto.TindakanOperasi, 0)
	}

	for _, V := range data {
		fmt.Println(V.Keterangan)
		if V.Keterangan == "icd10" && V.Jenis == "post" {
			tindakans = append(tindakans, dto.TindakanOperasi{
				Keterangan: V.Ket,
				Kode:       V.Kode,
				Deskripsi:  V.Diagnosa,
			})
		}

	}

	return tindakans
}

func toTindakanOperasi(data []asesmen.Dpenlab2) (res []dto.TindakanOperasi) {
	var datas = []dto.TindakanOperasi{}

	if len(data) == 0 {
		datas = make([]dto.TindakanOperasi, 0)
	}

	if data == nil {
		datas = make([]dto.TindakanOperasi, 0)
	}
	for _, V := range data {
		if V.Keterangan == "icd9" {
			datas = append(datas, dto.TindakanOperasi{
				Keterangan: V.Ket,
				Kode:       V.Kode,
				Deskripsi:  V.Diagnosa,
			})
		}

	}

	return datas
}

func toAsisten(data []asesmen.DPenLab3) (res []dto.Asisten) {

	var datas = []dto.Asisten{}

	if len(data) == 0 {
		datas = make([]dto.Asisten, 0)
	}

	if data == nil {
		datas = make([]dto.Asisten, 0)
	}

	for _, V := range data {
		if V.Ket == "asisten1" {
			datas = append(datas, dto.Asisten{
				Nama: V.Nama,
				Ket:  strings.ToUpper(V.Ket),
			})
		}

		if V.Ket == "asisten2" {
			datas = append(datas, dto.Asisten{
				Nama: V.Nama,
				Ket:  strings.ToUpper(V.Ket),
			})
		}

	}

	return datas
}

func toInstrumen(data []asesmen.DPenLab3) (res []dto.Istrumen) {
	var datas = []dto.Istrumen{}

	if len(data) == 0 {
		datas = make([]dto.Istrumen, 0)
	}

	if data == nil {
		datas = make([]dto.Istrumen, 0)
	}

	for _, V := range data {
		if V.Ket == "instrumen" {
			datas = append(datas, dto.Istrumen{
				Nama: V.Nama,
			})
		} else {
			datas = make([]dto.Istrumen, 0)
		}

	}

	return datas
}

func toAhliBedah(data []asesmen.DPenLab3) (res []dto.Penata) {

	var datas = []dto.Penata{}
	if data == nil {
		datas = make([]dto.Penata, 0)
	}

	if len(data) == 0 {
		datas = make([]dto.Penata, 0)
	}
	for _, V := range data {
		if V.Ket == "operator" {
			datas = append(datas, dto.Penata{
				Nama: V.Nama,
			})
		}

	}

	return datas
}

func toPenataAnastesi(data []asesmen.DPenLab3) (res []dto.Penata) {
	var datas = []dto.Penata{}

	if data == nil {
		datas = make([]dto.Penata, 0)
	}

	if len(data) == 0 {
		datas = make([]dto.Penata, 0)
	}

	for _, V := range data {
		if V.Ket == "penata" {
			datas = append(datas, dto.Penata{
				Nama: V.Nama,
			})
		}

	}

	return datas
}

func toAnastesi(data []asesmen.DPenLab3) (res []dto.Penata) {
	var datas = []dto.Penata{}

	if data == nil {
		datas = make([]dto.Penata, 0)
	}

	if len(data) == 0 {
		datas = make([]dto.Penata, 0)
	}

	for _, V := range data {
		if V.Ket == "anestesi" {
			datas = append(datas, dto.Penata{
				Nama: V.Nama,
			})
		}

	}

	return datas
}

func toKlafikasiLuka(value string) (res []dto.KlasifikasiLuka) {
	klasifikasiLuka := []dto.KlasifikasiLuka{}

	switch value {
	case "BERSIH":
		klasifikasiLuka = append(klasifikasiLuka,
			dto.KlasifikasiLuka{
				Nama:     "BERSIH",
				IsActive: true,
			},
			dto.KlasifikasiLuka{
				Nama:     "BERSIH TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "KOTOR ATAU DENGAN INFEKSI",
				IsActive: false,
			},
		)
	case "BERSIH TERKONTAMINASI":
		klasifikasiLuka = append(klasifikasiLuka,
			dto.KlasifikasiLuka{
				Nama:     "BERSIH",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "BERSIH TERCEMAR",
				IsActive: true,
			},
			dto.KlasifikasiLuka{
				Nama:     "TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "KOTOR ATAU DENGAN INFEKSI",
				IsActive: false,
			},
		)
	case "KOTOR":
		klasifikasiLuka = append(klasifikasiLuka,
			dto.KlasifikasiLuka{
				Nama:     "BERSIH",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "BERSIH TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "KOTOR ATAU DENGAN INFEKSI",
				IsActive: true,
			},
		)
	case "TERCEMAR":
		klasifikasiLuka = append(klasifikasiLuka,
			dto.KlasifikasiLuka{
				Nama:     "BERSIH",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "BERSIH TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "TERCEMAR",
				IsActive: true,
			},
			dto.KlasifikasiLuka{
				Nama:     "KOTOR ATAU DENGAN INFEKSI",
				IsActive: false,
			},
		)

	default:
		klasifikasiLuka = append(klasifikasiLuka,
			dto.KlasifikasiLuka{
				Nama:     "BERSIH",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "BERSIH TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "TERCEMAR",
				IsActive: false,
			},
			dto.KlasifikasiLuka{
				Nama:     "KOTOR ATAU DENGAN INFEKSI",
				IsActive: false,
			},
		)

	}

	return klasifikasiLuka

}

func toKlasifikasi(value string) (res []dto.Klasifiksai) {

	klasifikasis := []dto.Klasifiksai{}

	switch value {
	case "EMERGENCY":
		klasifikasis = append(klasifikasis, dto.Klasifiksai{
			Nama:     "EMERGENCY",
			IsActive: true,
		})

		klasifikasis = append(klasifikasis, dto.Klasifiksai{
			Nama:     "ELECTIVE",
			IsActive: false,
		})
	case "ELECTIVE":

		klasifikasis = append(klasifikasis, dto.Klasifiksai{
			Nama:     "ELECTIVE",
			IsActive: true,
		})

		klasifikasis = append(klasifikasis, dto.Klasifiksai{
			Nama:     "EMERGENCY",
			IsActive: false,
		})

	default:
		klasifikasis = append(klasifikasis, dto.Klasifiksai{
			Nama:     "EMERGENCY",
			IsActive: false,
		})

		klasifikasis = append(klasifikasis, dto.Klasifiksai{
			Nama:     "ELECTIVE",
			IsActive: false,
		})

	}

	return klasifikasis
}

func toPengirimanJaringan(value string) (res []dto.PengirimanJaringan) {
	pengirimanJaringan := []dto.PengirimanJaringan{}

	switch value {
	case "true":
		pengirimanJaringan = append(pengirimanJaringan,
			dto.PengirimanJaringan{
				Nama:     "YA",
				IsActive: true,
			},
			dto.PengirimanJaringan{
				Nama:     "TIDAK",
				IsActive: false,
			},
		)
	case "false":
		pengirimanJaringan = append(pengirimanJaringan,
			dto.PengirimanJaringan{
				Nama:     "YA",
				IsActive: false,
			},
			dto.PengirimanJaringan{
				Nama:     "TIDAK",
				IsActive: true,
			},
		)
	default:
		pengirimanJaringan = append(pengirimanJaringan,
			dto.PengirimanJaringan{
				Nama:     "YA",
				IsActive: false,
			},
			dto.PengirimanJaringan{
				Nama:     "TIDAK",
				IsActive: false,
			})

	}

	return pengirimanJaringan
}

func toJenisOperasi(value string) (res []dto.JenisOperasi) {
	jenisOperasi := []dto.JenisOperasi{
		{
			Nama:     "CANGIH",
			IsActive: false,
		},
		{
			Nama:     "KHUSUS",
			IsActive: false,
		},
		{
			Nama:     "BESAR",
			IsActive: false,
		},
		{
			Nama:     "SEDANG",
			IsActive: false,
		},
		{
			Nama:     "KECIL",
			IsActive: false,
		},
	}

	switch value {
	case "SEDANG":
		{
			return []dto.JenisOperasi{
				{
					Nama:     "CANGIH",
					IsActive: false,
				},
				{
					Nama:     "KHUSUS",
					IsActive: false,
				},
				{
					Nama:     "BESAR",
					IsActive: false,
				},
				{
					Nama:     "SEDANG",
					IsActive: true,
				},
				{
					Nama:     "KECIL",
					IsActive: false,
				},
			}
		}
	case "KECIL":
		{
			return []dto.JenisOperasi{
				{
					Nama:     "CANGIH",
					IsActive: false,
				},
				{
					Nama:     "KHUSUS",
					IsActive: false,
				},
				{
					Nama:     "BESAR",
					IsActive: false,
				},
				{
					Nama:     "SEDANG",
					IsActive: false,
				},
				{
					Nama:     "KECIL",
					IsActive: true,
				},
			}
		}
	case "KHUSUS":
		return []dto.JenisOperasi{
			{
				Nama:     "CANGIH",
				IsActive: false,
			},
			{
				Nama:     "KHUSUS",
				IsActive: true,
			},
			{
				Nama:     "BESAR",
				IsActive: false,
			},
			{
				Nama:     "SEDANG",
				IsActive: false,
			},
			{
				Nama:     "KECIL",
				IsActive: false,
			},
		}
	case "BESAR":
		return []dto.JenisOperasi{
			{
				Nama:     "CANGIH",
				IsActive: false,
			},
			{
				Nama:     "KHUSUS",
				IsActive: false,
			},
			{
				Nama:     "BESAR",
				IsActive: true,
			},
			{
				Nama:     "SEDANG",
				IsActive: false,
			},
			{
				Nama:     "KECIL",
				IsActive: false,
			},
		}
	case "CANGIH":
		return []dto.JenisOperasi{
			{
				Nama:     "CANGIH",
				IsActive: true,
			},
			{
				Nama:     "KHUSUS",
				IsActive: false,
			},
			{
				Nama:     "BESAR",
				IsActive: false,
			},
			{
				Nama:     "SEDANG",
				IsActive: false,
			},
			{
				Nama:     "KECIL",
				IsActive: false,
			},
		}
	default:
		return jenisOperasi
	}

}
