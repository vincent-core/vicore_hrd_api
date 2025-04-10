package mapper

import (
	"fmt"
	"strings"
	"vicore_hrd/app/rest"
	resumemedis "vicore_hrd/modules/resume_medis"
	"vicore_hrd/modules/triase"
	"vicore_hrd/modules/triase/dto"
	"vicore_hrd/modules/triase/entity"

	"github.com/sirupsen/logrus"
)

type triaseMapper struct {
	logging *logrus.Logger
}

func NewTriaseMapper(logging *logrus.Logger) entity.TriaseMapper {
	return &triaseMapper{
		logging: logging,
	}
}

func (tm *triaseMapper) ToTriaseMapper(DRegister resumemedis.DRegisterPasien, DProfilePasien resumemedis.DProfilePasien, fisik triase.TriaseDPemFisik, vitalSign triase.DVitalSign, triaseAsesmen triase.TriaseAsesmen, nyeri triase.AsesmenUlangNyeri, asesmen triase.AsesmenKeperawatan, BaseURL string, triaseDokter dto.ResposeTriaseIGDDokter, triaseModel triase.TriaseModel, asesmedTriase triase.AsesmenTriaseIGD, triase triase.Triase) (res dto.ResponseTriase) {
	var imageNyeriList []dto.ImageNyeri

	// Add elements to the slice
	imageNyeriList = append(imageNyeriList, dto.ImageNyeri{
		Skor:     0,
		ImageURL: BaseURL + "/app/images/nyeri/1.png",
	})

	imageNyeriList = append(imageNyeriList, dto.ImageNyeri{
		Skor:     1,
		ImageURL: BaseURL + "/app/images/nyeri/2.png",
	})

	imageNyeriList = append(imageNyeriList, dto.ImageNyeri{
		Skor:     2,
		ImageURL: BaseURL + "/app/images/nyeri/3.png",
	})

	imageNyeriList = append(imageNyeriList, dto.ImageNyeri{
		Skor:     3,
		ImageURL: BaseURL + "/app/images/nyeri/4.png",
	})

	imageNyeriList = append(imageNyeriList, dto.ImageNyeri{
		Skor:     4,
		ImageURL: BaseURL + "/app/images/nyeri/5.png",
	})

	imageNyeriList = append(imageNyeriList, dto.ImageNyeri{
		Skor:     5,
		ImageURL: BaseURL + "/app/images/nyeri/6.png",
	})

	tglMasuk, _ := rest.UbahTanggalIndo(triaseModel.InsertDttm)

	return dto.ResponseTriase{
		Ruangan:         "Instalasi Gawat Darurat",
		TanggalMasuk:    tglMasuk,
		JamPemeriksaan:  triaseModel.JamMasuk + " WIB",
		JamKedatangan:   DRegister.JamMasuk + " WIB",
		NamaPasien:      DProfilePasien.Firstname,
		TanggalLahir:    DProfilePasien.Tgllahir.Format("2006-01-02"),
		NomorRM:         DProfilePasien.Id,
		Triase:          tm.ToMapperVitalSignTriase(fisik, vitalSign),
		KeluhanUtama:    asesmen.KeluhanUtama,
		PetugasTriase:   triaseAsesmen.NamaDokter,
		Nyeri:           asesmen.Nyeri,
		AlasaDatang:     toAlasanDatang(asesmedTriase.AseskepAlasanMasuk),
		StatusAlergi:    asesmen.RiwayatAlergi,
		ImageNyerSource: imageNyeriList,
		Pernafasan:      vitalSign.Pernafasan + " kali per menit",
		SkorNyeri:       nyeri.SkorNyeri,
		Kesadaran:       strings.ToUpper(fisik.Kesadaran),
		PenyebabCedera:  toPenyebabCedera(asesmedTriase.AseskepPenyebabCedera),
		GanguanPerilaku: triaseDokter.GangguanPerilaku,
		StatusKehamilan: triaseDokter.AseskepKehamilan,
		JalanNafas:      fisik.JalanNafas,
		Sirkulasi:       fisik.Sirkulasi,
		SkalaTriase:     toSkalaTriase(triase.SkalaTriaseIgd, nyeri.SkorNyeri),
	}
}

func toSkalaTriase(value string, skalaNyeri int) (res string) {
	var data = ""

	if value == "" {
		if skalaNyeri >= 9 {
			data = "Resusitasi"
		}

		data = "Emergency / Gawat Darurat"
	} else {
		data = value
	}

	return data
}

func toAlasanDatang(alasan string) (res string) {
	if alasan == "" {
		return "Penyakit"
	} else {
		return alasan
	}
}

func toPenyebabCedera(cedera string) (res string) {
	if cedera == "" {
		return "-"
	} else {
		return cedera
	}
}

func (tm *triaseMapper) ToMapperVitalSignTriase(fisik triase.TriaseDPemFisik, vitalSign triase.DVitalSign) (res dto.ResponseTandaVitalTriase) {
	return dto.ResponseTandaVitalTriase{
		Akral:         fisik.Akral,
		Pupil:         fisik.Pupil,
		RefleksCahaya: fisik.Refleks,
		Nadi:          vitalSign.Nadi + " mmHg",
		TD:            vitalSign.Td,
		Pernafasan:    vitalSign.Pernafasan + " kali per meni",
		Suhu:          vitalSign.Suhu + " °C",
		SPO2:          vitalSign.Spo2 + " %",
		GCS:           fmt.Sprintf("E %s V %s M %s", fisik.E, fisik.V, fisik.M),
	}
}

func ToImageNyeri(SkorNyeri int, BaseURL string) (res string) {
	switch SkorNyeri {
	case 0:
		return ""
	default:
		return "1"
	}

}

func (a *triaseMapper) ToResponseTriaseIGDDokter(data triase.TriaseIGDDokter, data1 triase.AsesmenTriaseIGD) (res dto.ResposeTriaseIGDDokter) {

	var tanggal = ""

	if len(data1.TglMasuk) > 9 {
		tanggal = data1.TglMasuk[0:10]
	} else {
		tanggal = data1.TglMasuk
	}

	return dto.ResposeTriaseIGDDokter{
		Jam:                     data1.InsertDttm,
		TanggalMasuk:            tanggal,
		UserID:                  data1.InsertUserId,
		AseskepKehamilanDjj:     data1.AseskepKehamilanDjj,
		AseskepAlasanMasuk:      data1.AseskepAlasanMasuk,
		AseskepCaraMasuk:        data1.AseskepCaraMasuk,
		AseskepPenyebabCedera:   data1.AseskepPenyebabCedera,
		AseskepKehamilan:        data1.AseskepKehamilan,
		AseskepKehamilanGravida: data1.AseskepKehamilanGravida,
		AseskepKehamilanPara:    data1.AseskepKehamilanPara,
		AseskepKehamilanAbortus: data1.AseskepKehamilanAbortus,
		AseskepKehamilanHpht:    data1.AseskepKehamilanHpht,
		AseskepKehamilanTtp:     data1.AseskepKehamilanTtp,
		GangguanPerilaku:        data1.AseskepGangguanPerilaku,
		SkalaNyeri:              data.SkalaNyeri,
		SkalaNyeriP:             data.SkalaNyeriP,
		SkalaNyeriQ:             data.SkalaNyeriQ,
		SkalaNyeriR:             data.SkalaNyeriR,
		SkalaNyeriS:             data.SkalaNyeriS,
		SkalaNyeriT:             data.SkalaNyeriT,
		SkalaTriase:             data.SkalaTriaseIgd,
		FlaccWajah:              data.FlaccWajah,
		FlaccKaki:               data.FlaccKaki,
		FlaccAktifitas:          data.FlaccAktifitas,
		FlaccMenangis:           data.FlaccMenangis,
		FlaccBersuara:           data.FlaccBersuara,
		FlaccTotal:              data.FlaccTotal,
	}
}
