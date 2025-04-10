package mapper

import (
	"vicore_hrd/app/rest"
	generalconsent "vicore_hrd/modules/general_consent"
	"vicore_hrd/modules/general_consent/dto"
	"vicore_hrd/modules/general_consent/entity"
)

type GeneralMapper struct {
}

func NewGeneralConsentMapper() entity.GeneralConsentMapper {
	return &GeneralMapper{}
}

func (mm *GeneralMapper) TOMappingGeneralConsent(data generalconsent.Pasien, pengkajian generalconsent.Pengkajian, general generalconsent.DGeneralConsent) (res dto.ToResponseGeneralConsentRAJAL) {
	tgl, _ := rest.UbahTanggalIndo(data.TanggalLahir)

	var pasien = generalconsent.Pasien{}

	pasien.TanggalLahir = tgl
	pasien.Alamat = data.Alamat
	pasien.Nik = data.Nik
	pasien.NoHp = data.NoHp
	pasien.NomorRekamMedis = data.NomorRekamMedis
	pasien.NamaPasien = data.NamaPasien

	return dto.ToResponseGeneralConsentRAJAL{
		Pasien:         pasien,
		Pelayanan:      pengkajian.Pelayanan,
		Petugas:        pengkajian.NamaPerawat,
		PenangungJawab: mm.ToMappingGeneral(general, pasien),
		Tanggal:        rest.FormatTanggalWaktu(general.InsertDttm),
	}

}

func (mm *GeneralMapper) TOMappingGeneralConsentRAJAL(data generalconsent.Pasien, pengkajian generalconsent.Pengkajian, general generalconsent.DGeneralConsent) (res dto.ToResponseGeneralConsentRAJAL) {
	tgl, _ := rest.UbahTanggalIndo(data.TanggalLahir)

	var pasien = generalconsent.Pasien{}

	pasien.TanggalLahir = tgl
	pasien.Alamat = data.Alamat
	pasien.Nik = data.Nik
	pasien.NoHp = data.NoHp
	pasien.NomorRekamMedis = data.NomorRekamMedis
	pasien.NamaPasien = data.NamaPasien

	return dto.ToResponseGeneralConsentRAJAL{
		Pasien:         pasien,
		Pelayanan:      "INSTALASI GAWAT DARURAT",
		Petugas:        pengkajian.NamaPerawat,
		PenangungJawab: mm.ToMappingGeneral(general, pasien),
		Tanggal:        rest.FormatTanggalWaktu(general.InsertDttm),
		PJawabNama:     general.PJawabNama,
		TtdPJawab:      general.TtdPjawab,
		Pewenang:       general.Pewenang,
	}

}

func (mm *GeneralMapper) TOMappingGeneralConsentRANAP(data generalconsent.Pasien, pengkajian generalconsent.Pengkajian, general generalconsent.DGeneralConsent, dokterJaga generalconsent.DokterAsesmen) (res dto.ToResponseGeneralConsentRANAP) {
	tgl, _ := rest.UbahTanggalIndo(data.TanggalLahir)

	var pasien = generalconsent.Pasien{}

	pasien.TanggalLahir = tgl
	pasien.Alamat = data.Alamat
	pasien.Nik = data.Nik
	pasien.NoHp = data.NoHp
	pasien.NomorRekamMedis = data.NomorRekamMedis
	pasien.NamaPasien = data.NamaPasien

	return dto.ToResponseGeneralConsentRANAP{
		Pasien:         pasien,
		Pelayanan:      "RANAP",
		Petugas:        pengkajian.NamaPerawat,
		PenangungJawab: mm.ToMappingGeneral(general, pasien),
		Tanggal:        rest.FormatTanggalWaktu(general.InsertDttm),
		TtdPasien:      general.TtdPjawab,
		Privasi:        general.Privasi,
		HubunganPasien: general.HubDenganPasien,
		NamaDokterJaga: dokterJaga.NamaDokter,
		KonsulKe:       dokterJaga.KonsulKe,
	}

}

func (mm *GeneralMapper) ToMappingGeneral(general generalconsent.DGeneralConsent, pasien generalconsent.Pasien) (res generalconsent.PenangungJawab) {
	return generalconsent.PenangungJawab{
		Nama:         toAutoFill(general.PJawabNama, pasien.NamaPasien),
		TanggalLahir: toAutoFill(general.PJawabTglLahir, pasien.TanggalLahir),
		Alamat:       toAutoFill(general.PJawabAlamat, pasien.Alamat),
		NoHP:         toAutoFill(general.PJawabNoHP, pasien.NoHp),
	}
}

func toAutoFill(value string, value2 string) (res string) {
	if value == "" {
		return value2
	} else {
		return value
	}
}
