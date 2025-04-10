package dto

type (
	GetAntranPasien struct {
		KDBagian string `json:"kd_bagian" validate:"required"`
	}

	OnGetDataRegisterByID struct {
		ID string `json:"no_rm" validate:"required"`
	}

	Dashboard struct {
		KDBagian string `json:"kd_bagian" validate:"required"`
	}

	AntrianPasien struct {
		Tgllahir       string `json:"tgl_lahir"`
		NoAntrean      string `json:"no_antrean"`
		JenisKelamin   string `json:"jenis_kelamin"`
		Debitur        string `json:"debitur"`
		KodeDebitur    string `json:"kd_debitur"`
		Noreg          string `json:"no_reg"`
		Mrn            string `json:"no_rm"`
		Keterangan     string `json:"keterangan"`
		NamaPasien     string `json:"nama_pasien"`
		KdBag          string `json:"kd_bagian"`
		Bagian         string `json:"bagian"`
		Pelayanan      string `json:"pelayanan"`
		NamaDokter     string `json:"nama_dokter"`
		KdDokter       string `json:"kd_dokter"`
		Kamar          string `json:"kamar"`
		Kasur          string `json:"kasur"`
		AsesmenDokter  string `json:"asesmen_dokter"`
		AsesmenPerawat string `json:"asesmen_perawat"`
	}

	ResponseDashboard struct {
		Jumlah int `json:"jumlah_pasien"`
	}
)
