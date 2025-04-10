package dto

type (
	ResponseEdukasiTerintegrasi struct {
		ID                int    `json:"id"`
		Tanggal           string `json:"tanggal"`
		Informasi         string `json:"informasi"`
		Metode            string `json:"metode"`
		PemberiInformasi  string `json:"pemberi_informasi"`
		PenerimaInformasi string `json:"penerima_informasi"`
		Evaluasi          string `json:"evaluasi"`
	}

	ReqEdukasiTerintegrasi struct {
		Informasi         string `json:"informasi" validate:"required"`
		Metode            string `json:"metode" validate:"required"`
		PemberiInformasi  string `json:"pemberi_informasi" validate:"required"`
		PenerimaInformasi string `json:"penerima_informasi" validate:"required"`
		Evaluasi          string `json:"evaluasi" validate:"required"`
		NORM              string `json:"no_rm" validate:"required"`
		NoReg             string `json:"no_reg" validate:"required"`
	}

	ReqOnUpdateEdukasiTerintegrasi struct {
		ID                int    `json:"id"`
		Informasi         string `json:"informasi" validate:"required"`
		Metode            string `json:"metode" validate:"required"`
		PemberiInformasi  string `json:"pemberi_informasi" validate:"required"`
		PenerimaInformasi string `json:"penerima_informasi" validate:"required"`
		Evaluasi          string `json:"evaluasi" validate:"required"`
	}

	ResponsePemberiInformasi struct {
		Nama         string `json:"nama"`
		JenisKelamin string `json:"jenis_kelamin"`
	}
)
