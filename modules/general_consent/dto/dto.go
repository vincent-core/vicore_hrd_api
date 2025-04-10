package dto

import (
	generalconsent "vicore_hrd/modules/general_consent"
)

type (
	ToResponseGeneralConsentRAJAL struct {
		Tanggal        string                        `json:"tanggal"`
		Pelayanan      string                        `json:"pelayanan"`
		Pasien         generalconsent.Pasien         `json:"pasien"`
		PenangungJawab generalconsent.PenangungJawab `json:"penanggung_jawab"`
		Petugas        string                        `json:"petugas"`
		PJawabNama     string                        `json:"pjawab_nama"`
		TtdPJawab      string                        `json:"ttd_pjawab"`
		Pewenang       string                        `json:"pewenang"`
	}

	ToResponseGeneralConsentRANAP struct {
		Tanggal        string                        `json:"tanggal"`
		Pelayanan      string                        `json:"pelayanan"`
		Pasien         generalconsent.Pasien         `json:"pasien"`
		PenangungJawab generalconsent.PenangungJawab `json:"penanggung_jawab"`
		Petugas        string                        `json:"petugas"`
		TtdPasien      string                        `json:"ttd_pasien"`
		Privasi        string                        `json:"privasi"`
		HubunganPasien string                        `json:"hubungan"`
		NamaDokterJaga string                        `json:"dokter_jaga"`
		KonsulKe       string                        `json:"konsul_ke"`
	}

	OnReportGeneralConsent struct {
		NORM      string `json:"no_rm" validate:"required"`
		KdBagian  string `json:"kd_bagian" validate:"omitempty"`
		Pelayanan string `json:"pelayanan" validate:"omitempty"`
		NoReg     string `json:"no_reg" validate:"omitempty"`
		Usia      string `json:"usia" validate:"omitempty"`
	}
)
