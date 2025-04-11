package resumemedis

import "time"

type (
	DRegisterPasien struct {
		TglMasuk time.Time
		Id       string
		Noreg    string
		Nama     string
		JamMasuk string
	}

	DProfilePasien struct {
		Id           string
		Jeniskelamin string
		Firstname    string
		Tgllahir     time.Time
		Alamat       string
	}

	DBangsalRep struct {
		NamaDokter string
		Bagian     string
		Tanggal    time.Time
		Jam        string
		Noreg      string
		Nama       string
		Keterangan string
	}

	AsesmenRanap struct {
		KdBagian  string
		Pelayanan string
		Noreg     string
	}

	DiagnosaICD struct {
		S1          string
		Code        string
		Code2       string
		Description string
	}

	DRekamSatu struct {
		Noreg string
		IcdIn string
		IcdOu string
	}

	Drekap struct {
		TglKeluar string
		Id        string
		Noreg     string
		Nama      string
		Umur      string
		Alamat    string
		Kelas     string
		Asal      string
		Pelayanan string
	}

	DApotikKeluarObat struct {
		Jam          string
		StatusPasien string
		Id           string
		Noreg        string
		NoKeluar     string
	}

	DataKeluarObat1 struct {
		TglKeluar time.Time `json:"tgl_keluar"`
		NoKeluar  string    `json:"no_keluar"`
		NoAmbil   string    `json:"no_ambil"`
		NamaObat  string    `json:"nama_obat"`
		Jumlah    int       `json:"jumlah"`
	}

	DPenLab struct {
		Jaminput string
		Noreg    string
		Id       string
	}

	ResHasilLaborTableLama struct {
		Tanggal           string              `json:"tanggal"`
		NamaKelompok      string              `json:"nama_kelompok"`
		DPemeriksaanLabor []DPemeriksaanLabor `json:"penlab"`
	}

	DPenmedPemeriksaan struct {
		NamaKelompok string
	}

	DPemeriksaanLabor struct {
		Pemeriksaan string `json:"pemeriksaan_deskripsi"`
		Normal      string `json:"normal"`
		Satuan      string `json:"satuan"`
		Hasil       string `json:"hasil"`
	}

	RegHasilRadiologiTabelLama struct {
		Tanggal              string                 `json:"tanggal"`
		NamaKelompok         string                 `json:"nama_kelompok"`
		DHasilRadiologiOldDB []DHasilRadiologiOldDB `json:"radiologi"`
	}

	DHasilRadiologiOldDB struct {
		Pemeriksaan string `json:"pemeriksaan_deskripsi"`
		Uraian      string `json:"uraian"`
		Hasil       string `json:"hasil"`
	}

	DHasilRadiologiV2 struct {
		Jaminput string
		Bagian   string
	}

	PemFisik struct {
		Kategori   string `json:"kategori"`
		KetPerson  string `json:"ket_person"`
		Noreg      string `json:"noreg"`
		Tb         string `json:"tb"`
		Td         string `json:"td"`
		Bb         string `json:"bb"`
		Nadi       string `json:"nadi"`
		Suhu       string `json:"suhu"`
		Spo2       string `json:"spo2"`
		Pernafasan string `json:"penafasan"`
	}

	DiagnosaResponse struct {
		Diagnosa    string `json:"diagnosa"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Table       string `json:"table"`
	}

	TindakanResponse struct {
		Code2       string `json:"kode2"`
		Description string `json:"description"`
	}

	RiwayatPenyakit struct {
		Noreg           string
		KdBagian        string
		RiwayatSekarang string
		RiwayatDahulu   string
		RiwayatObat     string
	}

	CariDataPasienPulang struct {
		Norm      string `json:"no_rm"`
		Tglproses string `json:"tanggal"`
		Pelayanan string `json:"pelayanan"`
		Bagian    string `json:"bagian"`
		Noreg     string `json:"noreg"`
		Nama      string `json:"nama"`
		Tgllahir  string `json:"tanggal_lahir"`
	}
	DataJenis struct {
		Id           string `json:"id"`
		JenisKelamin string `json:"jenis_kelamin"`
		Nama         string `json:"nama"`
	}

	DataDRekamMedis struct {
		Pelayanan  string
		Bagian     string
		Tahun      string
		Bulan      string
		Tglproses  string
		Jamproses  string
		Id         string
		Noreg      string
		Nama       string
		Alamat     string
		Sex        string
		Iddokter   string
		Dokter     string
		Tglmasuk   string
		Tglkeluar  string
		IcdIn      string
		IcdOu      string
		DiagnosaIn string
		DignosaOu  string
		Golongan   string
		Kunjungan  string
		KeaKeluar  string
		KetKeluar  string
	}

	DataProfilePasien struct {
		Nik          string
		Nokapst      string
		Id           string
		Firstname    string
		Lastname     string
		Agama        string
		Jeniskelamin string
		Tempatlahir  string
		Tgllahir     string
		Alamat       string
		Suku         string
		Kelurahan    string
		Kecamatan    string
		Kotamadya    string
		Kabupaten    string
		Propinsi     string
		Negara       string
		Telp         string
		Hp           string
		Namaayah     string
		Namaibu      string
	}
)

func (CariDataPasienPulang) TableName() string {
	return "rekam.drm_1"
}

func (DataDRekamMedis) TableName() string {
	return "rekam.drm_1"
}

func (DataProfilePasien) TableName() string {
	return "his.dprofilpasien"
}
