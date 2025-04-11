package dto

type (
	OnGetResumeMedis struct {
		Noreg string `json:"noreg" validate:"required"`
	}

	ResponseCariPasien struct {
		Norm         string `json:"no_rm"`
		Tglproses    string `json:"tanggal"`
		Pelayanan    string `json:"pelayanan"`
		Noreg        string `json:"noreg"`
		Nama         string `json:"nama"`
		JenisKelamin string `json:"jenis_kelamin"`
		Tgllahir     string `json:"tanggal_lahir"`
	}
)
