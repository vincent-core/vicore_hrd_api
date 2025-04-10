package dto

type (
	ResponseLaporanBedah struct {
		ProfilPasien           Pasien               `json:"profil_pasien"`
		NamaAhli               []Penata             `json:"ahli_bedah"`
		NamaAsisten            []Asisten            `json:"asisten"`
		NamaInstrumen          []Istrumen           `json:"istrumen"`
		NamaAhliAnastesi       []Penata             `json:"ahli_anastesi"`
		NamaPerawatAnastesi    []Penata             `json:"perawat_anastesi"`
		TanggalOpersi          string               `json:"tanggal_operasi"`
		JamOperasiDimulai      string               `json:"jam_operasi_dimulai"`
		JamOperasiSelesai      string               `json:"jam_operasi_selesai"`
		LamaOperasiBerlangsung string               `json:"lama_operasi_berlangsung"`
		Klasifiksai            []Klasifiksai        `json:"klasifikasi"`
		JenisOperasi           []JenisOperasi       `json:"jenis_operasi"`
		KlasifikasiLuka        []KlasifikasiLuka    `json:"klasifikasi_luka"`
		PengirimanJaringan     []PengirimanJaringan `json:"pengiriman_jaringan"`
		JenisJaringan          string               `json:"jenis_jaringan"`
		UraianOperasi          string               `json:"uraian_operasi"`
		Tindakan               []TindakanOperasi    `json:"tindakan_operasi"`
		DiagnosaPre            []TindakanOperasi    `json:"diagnosa_pre"`
		DiagnosaPost           []TindakanOperasi    `json:"diagnosa_post"`
		Tanggal                string               `json:"tanggal"`
	}

	TindakanOperasi struct {
		Keterangan string `json:"keterangan"`
		Kode       string `json:"kode"`
		Deskripsi  string `json:"deskripsi"`
	}

	Pasien struct {
		Noreg        string `json:"no_reg"`
		Nama         string `json:"nama"`
		TanggalLahir string `json:"tgl_lahir"`
		NoRm         string `json:"no_rm"`
		JenisKelamin string `json:"jenis_kelamin"`
	}

	Asisten struct {
		Nama string `json:"nama"`
		Ket  string `json:"ket"`
	}

	Istrumen struct {
		Nama string `json:"nama"`
	}

	Penata struct {
		Nama string `json:"nama"`
	}

	Perawat struct {
		Nama string `json:"nama"`
	}

	Klasifiksai struct {
		Nama     string `json:"nama"`
		IsActive bool   `json:"is_active"`
	}

	JenisOperasi struct {
		Nama     string `json:"nama"`
		IsActive bool   `json:"is_active"`
	}

	KlasifikasiLuka struct {
		Nama     string `json:"nama"`
		IsActive bool   `json:"is_active"`
	}

	PengirimanJaringan struct {
		Nama     string `json:"nama"`
		IsActive bool   `json:"is_active"`
	}
)
