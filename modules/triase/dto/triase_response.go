package dto

type (
	ResponseTriase struct {
		NamaPasien      string                   `json:"nama_pasien"`
		Ruangan         string                   `json:"ruangan"`
		TanggalLahir    string                   `json:"tgl_lahir"`
		NomorRM         string                   `json:"no_rm"`
		TanggalMasuk    string                   `json:"tgl_masuk"`
		JamKedatangan   string                   `json:"jam_kedatangan"`
		JamPemeriksaan  string                   `json:"jam_pemeriksaan"`
		AlasaDatang     string                   `json:"alasan_datang"`
		PenyebabCedera  string                   `json:"penyebab_cedera"`
		KeluhanUtama    string                   `json:"keluhan_utama"`
		Triase          ResponseTandaVitalTriase `json:"triase"`
		StatusAlergi    string                   `json:"status_alergi"`
		GanguanPerilaku string                   `json:"gangguan_perilaku"`
		StatusKehamilan string                   `json:"status_kehamilan"`
		JalanNafas      string                   `json:"jalan_nafas"`
		Pernafasan      string                   `json:"pernafasan"`
		Sirkulasi       string                   `json:"sirkulasi"`
		Kesadaran       string                   `json:"kesadaran"`
		Nyeri           string                   `json:"nyeri"`
		PetugasTriase   string                   `json:"petugas_triase"`
		SkorNyeri       int                      `json:"skor_nyeri"`
		ImageNyerSource []ImageNyeri             `json:"gambar_nyeri"`
		SkalaTriase     string                   `json:"skala_triase"`
	}

	ResponseTandaVitalTriase struct {
		GCS           string `json:"gcs"`
		TD            string `json:"td"`
		Nadi          string `json:"nadi"`
		Pupil         string `json:"pupil"`
		Pernafasan    string `json:"pernafasan"`
		Suhu          string `json:"suhu"`
		RefleksCahaya string `json:"refleks"`
		SPO2          string `json:"spo2"`
		Akral         string `json:"akral"`
	}

	ImageNyeri struct {
		Skor     int    `json:"skor"`
		ImageURL string `json:"image_url"`
	}

	ResposeTriaseIGDDokter struct {
		Jam                     string `json:"jam"`
		TanggalMasuk            string `json:"tanggal_masuk"`
		UserID                  string `json:"user_id"`
		AseskepKehamilanDjj     string `json:"ddj"`
		AseskepAlasanMasuk      string `json:"alasan_masuk"`
		AseskepCaraMasuk        string `json:"cara_masuk"`
		AseskepPenyebabCedera   string `json:"penyebab_cedera"`
		AseskepKehamilan        string `json:"kehamilan"`
		AseskepKehamilanGravida string `json:"gravida"`
		AseskepKehamilanPara    string `json:"para"`
		AseskepKehamilanAbortus string `json:"abortus"`
		AseskepKehamilanHpht    string `json:"hpht"`
		AseskepKehamilanTtp     string `json:"ttp"`
		GangguanPerilaku        string `json:"gangguan_perilaku"`
		SkalaNyeri              int    `json:"skala_nyeri"`
		SkalaNyeriP             string `json:"nyeri_p"`
		SkalaNyeriQ             string `json:"nyeri_q"`
		SkalaNyeriR             string `json:"nyeri_r"`
		SkalaNyeriS             string `json:"nyeri_s"`
		SkalaNyeriT             string `json:"nyeri_t"`
		SkalaTriase             string `json:"skala_triase"`
		FlaccWajah              int    `json:"flag_wajah"`
		FlaccKaki               int    `json:"flag_kaki"`
		FlaccAktifitas          int    `json:"flag_aktifitas"`
		FlaccMenangis           int    `json:"flag_menangis"`
		FlaccBersuara           int    `json:"flag_bersuara"`
		FlaccTotal              int    `json:"flag_total"`
	}
)
