package lembarkonsul

type (
	CpptKonsulen struct {
		Ppa          string
		InsertDttm   string
		InsertUserId string
		KdBagian     string
		Subjektif    string
		Objektif     string
		Asesmen      string
		Plan         string
	}
	LembarKonsul struct {
		Id          int
		Nama        string
		Jenis       string
		Jumlah      int
		Harga       int
		Tanggal     string
		Tanggal_kir string
	}

	DRegister struct {
		Id    string
		Noreg string
		Nama  string
	}

	DKonsulPasien struct {
		DokterKonsul   string
		InsertDttm     string
		JenisKonsul    string
		DokterKonsulKe string
		KonsulKe       string
		IktisarKlinik  string
		KdBagian       string
		Noreg          string
	}
)
