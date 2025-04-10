package triase

type (
	DVitalSign struct {
		InsertDttm string
		Kategori   string `json:"kategori"`
		KetPerson  string `json:"ket_person"`
		Noreg      string `json:"noreg"`
		Tb         string `json:"tb"`
		Td         string `json:"td"`
		Bb         string `json:"bb"`
		Nadi       string
		Suhu       string
		Spo2       string
		Pernafasan string
	}

	TriaseDPemFisik struct {
		Noreg      string
		KdBagian   string
		Akral      string
		Pupil      string
		Refleks    string
		Kesadaran  string
		JalanNafas string
		Sirkulasi  string
		E          string
		V          string
		M          string
	}

	TriaseAsesmen struct {
		KdBagian     string
		KeluhanUtama string
		UserId       string
		NamaDokter   string
	}

	IgdAsesmen struct {
		InsertDttm      string
		AsesmedKonsulKe string
		Noreg           string
		KdBagian        string
	}

	DokterNama struct {
		Dokter string
	}

	DokterDiagnosa struct {
		Noreg        string
		Kodediagnosa string
	}

	// insert_dttm, asesmed_konsul_ke, noreg, kd_bagian from vicore_rme.dcppt_soap_dokter WHERE kd_bagian="IGD001" AND asesmed_konsul_ke="dr. Surijanto Muskita"

	AsesmenUlangNyeri struct {
		Noreg     string
		KdBagian  string
		Metode    string
		SkorNyeri int
	}

	AsesmenKeperawatan struct {
		Noreg         string
		KdBagian      string
		CaraMasuk     string
		RiwayatAlergi string
		KeluhanUtama  string
		Nyeri         string
	}

	TriaseModel struct {
		InsertDttm string
		Noreg      string
		KdBagian   string
		TglMasuk   string
		JamMasuk   string
		Pelayanan  string
	}

	AsesmenTriaseIGD struct {
		InsertDttm              string
		InsertPc                string
		InsertUserId            string
		TglMasuk                string
		KeteranganPerson        string
		Pelayanan               string
		KdBagian                string
		Noreg                   string
		AseskepGangguanPerilaku string
		AseskepKehamilanDjj     string
		AseskepAlasanMasuk      string
		AseskepCaraMasuk        string
		AseskepPenyebabCedera   string
		AseskepKehamilan        string
		AseskepKehamilanGravida string
		AseskepKehamilanPara    string
		AseskepKehamilanAbortus string
		AseskepKehamilanHpht    string
		AseskepKehamilanTtp     string
	}

	Triase struct {
		InsertUserId   string
		Noreg          string
		SkalaTriaseIgd string
	}

	TriaseIGDDokter struct {
		InsertDttm     string
		InsertDevice   string
		InsertUserId   string
		KetPerson      string
		Pelayanan      string
		Kategori       string
		KdBagian       string
		Noreg          string
		SkalaNyeri     int
		SkalaNyeriP    string
		SkalaNyeriQ    string
		SkalaNyeriR    string
		SkalaNyeriS    string
		SkalaNyeriT    string
		FlaccWajah     int
		FlaccKaki      int
		FlaccAktifitas int
		FlaccMenangis  int
		FlaccBersuara  int
		FlaccTotal     int
		SkalaTriase    int
		SkalaTriaseIgd string
	}
)

func (AsesmenTriaseIGD) TableName() string {
	return "vicore_rme.dcppt_soap_pasien"
}

func (TriaseIGDDokter) TableName() string {
	return "vicore_rme.dpem_fisik"
}
