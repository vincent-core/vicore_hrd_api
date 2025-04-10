package dto

type (
	InsertCPPTSOAP struct {
		Noreg        string `json:"noreg" `
		Kelompok     string `json:"kelompok" validate:"required"`
		Pelayanan    string `json:"pelayanan" validate:"required"`
		KdBagian     string `json:"kd_bagian" validate:"required"`
		Dpjp         string `json:"dpjp" validate:"required"`
		Sujektif     string `json:"subjektif" validate:"required"`
		Objektif     string `json:"objektif" validate:"required"`
		Asesmen      string `json:"asesmen" validate:"required"`
		Plan         string `json:"plan" validate:"required"`
		InstruksiPpa string `json:"instruksi_ppa" validate:"omitempty"`
	}

	InsertCPPTSBAR struct {
		NoReg         string `json:"no_reg" validate:"required"`
		KdBagian      string `json:"kd_bagian" validate:"required"`
		Kelompok      string `json:"kelompok" validate:"required"`
		Situation     string `json:"situation" validate:"required"`
		Asesmen       string `json:"asesmen" validate:"required"`
		Background    string `json:"background" validate:"required"`
		Recomendation string `json:"recomendation" validate:"required"`
		Ppa           string `json:"ppa" validate:"omitempty"`
		Pelayanan     string `json:"pelayanan" validate:"omitempty"`
		Dpjp          string `json:"dpjp" validate:"omitempty"`
	}

	ResponseCPPT struct {
		ID            int    `json:"id"`
		Tanggal       string `json:"tanggal"`
		Keterangan    string `json:"cppt"`
		InstruksiPPA  string `json:"instruksi_ppa"`
		DPJP          string `json:"dpjp"`
		PemberiAsuhan string `json:"pemberi_asuhan"`
	}

	DataReportCPPT struct {
		ProfilePasien DataProfilePasien `json:"profil_pasien"`
		CPPT          []ResponseCPPT    `json:"cppt"`
	}

	//====
	// AMBIL DARI RKEM DRM SATU
	DataProfilePasien struct {
		NoRm         string `json:"no_rm"`
		TanggalLahir string `json:"tgl_lahir"`
		JenisKelamin string `json:"jenis_kelamin"`
		NamaPasien   string `json:"nama"`
		Ruangan      string `json:"ruangan"`
		NoReg        string `json:"no_reg"`
	}

	OnUpdateCPPTSoapRes struct {
		ID           int    `json:"id"`
		Subjektif    string `json:"subjektif" validate:"required"`
		Objektif     string `json:"objektif" validate:"required"`
		Asesmen      string `json:"asesmen" validate:"required"`
		Plan         string `json:"plan" validate:"required"`
		InstruksiPPA string `json:"instruksi_ppa" validate:"required"`
	}

	OnUpdateSBARRes struct {
		ID            int    `json:"id" validate:"required"`
		Situation     string `json:"situation" validate:"required"`
		Background    string `json:"background" validate:"required"`
		Asesmen       string `json:"asesmen" validate:"required"`
		Recomendation string `json:"recomendation" validate:"required"`
		InstruksiPPA  string `json:"instruksi_ppa" validate:"required"`
	}

	InstruksiObat struct {
		TglKeluar string `json:"tgl_keluar"`
		NamaObat  string `json:"nama_obat"`
		Jumlah    int    `json:"jumlah"`
	}
)
