package asesmen

import (
	"time"
	"vicore_hrd/modules/his"
	"vicore_hrd/modules/lib"
)

type (
	AsesmenDokter struct {
		InsertDttm        time.Time             `json:"insert_dttm"`         // Timestamp of insertion
		Noreg             string                `json:"noreg"`               // Registration number
		AsesmedKeluhUtama string                `json:"asesmed_keluh_utama"` // Main complaint in medical assessment
		AsesmedRwytSkrg   string                `json:"asesmed_rwyt_skrg"`   // Current medical history
		AsesmedRwytDahulu string                `json:"asesmed_rwyt_dahulu"` // Past medical history
		Pelayanan         string                `json:"pelayanan"`           // Service provided
		KdBagian          string                `json:"kd_bagian"`           // Department code
		KeteranganPerson  string                `json:"keterangan_person"`   // Person's description
		InsertUserID      string                `json:"insert_user_id"`
		Dokter            his.KTaripDokterModel `gorm:"foreignKey:InsertUserID" json:"dokter"`
	}

	AsesmenDokterPengantarRawatInap struct {
		Waktu        string
		Person       string
		UserId       string
		KdBagian     string
		NamaDokter   string
		KeluhanUtama string
		NamaBagian   string
		KonsulKe     string
	}

	PengkajianKeperawatan struct {
		IDPengkajian             int                   `gorm:"column:id_pengkajian;primaryKey;autoIncrement"`
		InsertDttm               time.Time             `gorm:"column:insert_dttm;default:0000-00-00 00:00:00"`
		KdBagian                 string                `gorm:"column:kd_bagian;size:50"`
		UserID                   string                `gorm:"column:user_id;size:50"`
		Noreg                    string                `gorm:"column:noreg;size:225"`
		Usia                     string                `gorm:"column:usia;type:enum('DEWASA','ANAK');default:DEWASA"`
		Pelayanan                string                `gorm:"column:pelayanan;type:enum('RAJAL','RANAP')"`
		Anamnesa                 string                `gorm:"column:anamnesa;size:225"`
		CaraMasuk                string                `gorm:"column:cara_masuk;size:225"`
		AsalPasien               string                `gorm:"column:asal_pasien;size:225"`
		KeluhanUtama             string                `gorm:"column:keluhan_utama;size:225"`
		RiwayatPenyakitSekarang  string                `gorm:"column:riwayat_penyakit_sekarang;type:text"`
		RiwayatPenyakitDahulu    string                `gorm:"column:riwayat_penyakit_dahulu;type:text"`
		RiwayatPenyakitKeluarga  string                `gorm:"column:riwayat_penyakit_keluarga;type:text"`
		RiwayatPengobatanDirumah string                `gorm:"column:riwayat_pengobatan_dirumah;type:enum('TIDAK ADA','ADA')"`
		RiwayatAlergi            string                `gorm:"column:riwayat_alergi;type:enum('TIDAK ADA','ADA')"`
		ReaksiAlergi             string                `gorm:"column:reaksi_alergi;size:225"`
		Nyeri                    string                `gorm:"column:nyeri;type:enum('TIDAK ADA','ADA')"`
		RiwayatKehamilan         string                `gorm:"column:riwayat_kehamilan;"`
		TindakLanjut             string                `gorm:"column:tindak_lanjut;"`
		AlasanRujukan            string                `gorm:"column:alasan_rujukan;"`
		KeteranganRujuk          string                `gorm:"column:keterangan_rujuk;"`
		KondisiPasien            string                `gorm:"column:kondisi_pasien;"`
		Indikasi                 string                `gorm:"column:indikasi;"`
		TransportasiPulang       string                `gorm:"column:transportasi_pulang;"`
		DPJP                     string                `gorm:"column:dpjp;"`
		Ruangan                  string                `gorm:"column:ruangan;"`
		PendidikanPasien         string                `gorm:"column:pendidikan_pasien;"`
		DischargePlanning        string                `gorm:"column:discharge_planning;"`
		KPelayanan               lib.KPelayanan        `gorm:"foreignKey:Ruangan" json:"bagian"`
		KDokter                  his.KTaripDokterModel `gorm:"foreignKey:DPJP" json:"dokter"`
		Perawat                  his.UserPerawatModel  `gorm:"foreignKey:UserID" json:"user"`
	}

	// CPPT PASIEN
	DCPPT struct {
		IDCppt        int `gorm:"primaryKey;autoIncrement"`
		InsertDttm    time.Time
		UpdDttm       time.Time
		InsertUserId  string
		InsertPc      string
		Kelompok      string
		Pelayanan     string
		KdBagian      string
		Tanggal       time.Time
		Noreg         string
		Dpjp          string
		Subjektif     string
		Objektif      string
		Asesmen       string
		Plan          string
		Situation     string
		Background    string
		Recomendation string
		InstruksiPpa  string
		PpaFingerTtd  string
		PpaFingerTgl  time.Time
		PpaFingerJam  string
		VerifikasiTtd string
		VerifikasiTgl time.Time
		VerifikasiJam string
	}

	DataCPPT struct {
		IdCppt          int
		InsertDttm      string
		InsertUserId    string
		Kelompok        string
		Pelayanan       string
		KdBagian        string
		Noreg           string
		Dpjp            string
		Subjektif       string
		Situation       string
		Objektif        string
		Asesmen         string
		Plan            string
		Background      string
		Recomendation   string
		InstruksiPpa    string
		NamaDokterDpjp  string
		NamaProfesional string
		Namabagian      string
		Id              string
	}

	AsesmenDokterIGD struct {
		Waktu            string
		Bagian           string
		Noreg            string
		UserId           string
		Dokter           string
		KeluhanUtama     string
		RiwayatSekarang  string
		RiwayatDahulu    string
		ImageLokalis     string
		CaraKeluar       string
		CaraKeluarDetail string
		Prognosis        string
		KonsulKe         string
		Terapi           string
	}

	DiagnosaResponse struct {
		Diagnosa    string `json:"diagnosa"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Table       string `json:"table"`
	}

	DApotikKeluarObat struct {
		Ket          string
		Jam          string
		StatusPasien string
		Id           string
		Noreg        string
		NoKeluar     string
	}

	DataPasienResgister struct {
		Id    string
		Noreg string
	}

	DPemfisikModel struct {
		NoReg     string
		KdBagian  string
		GcsE      string
		GcsM      string
		GcsV      string
		Kesadaran string
	}
	DVitalSign struct {
		KdBagian   string
		NoReg      string
		Suhu       string
		Pernafasan string
		Nadi       string
		Td         string
		Bb         string
		Spo2       string
	}

	RiwayatPenyakit struct {
		Id       string
		Kelompok string
		Alergi   string
	}

	// TABLE LAPORAN OPERASI
	DPenLabBedah struct {
		Nomor          string
		DrOperator     string
		User           string
		Cdttm          string
		Udttm          string
		Id             string
		Noreg          string
		Nama           string
		Gender         string
		Tgllahir       string
		Usia           string
		Alamat         string
		Tanggungan     string
		Tglbedah       string
		Mulai          string
		Akhir          string
		Lama           string
		Uraian         string
		Klasifikasi    string
		KlasifiksiLuka string
		Jenis          string
		AdaJaringan    string
		Patologi       string
	}

	Bedah struct {
		Mac             string
		Gender          string
		Cdttm           string
		Nomor           string
		Id              string
		Noreg           string
		Nama            string
		Sex             string
		Tgllahir        string
		Tglbedah        string
		Mulai           string
		Akhir           string
		Lama            string
		Operator        string
		AsistenSatu     string
		AsistenDua      string
		Anastesi        string
		Penata          string
		Istrumen        string
		Klasifikasi     string
		KlasifikasiLuka string
		Jenis           string
		AdaJaringan     string
		Patologi        string
		Uraian          string
		Kode            string
		Diagnosa        string
		Ket             string
	}
	DPenLab3 struct {
		Nomor string
		Nama  string
		Ket   string
	}

	Dpenlab2 struct {
		Jenis      string
		Kode       string
		Diagnosa   string
		Ket        string
		Keterangan string
	}
)

func (PengkajianKeperawatan) TableName() string {
	return "vicore_rme.dpengkajian_keperawatan"
}

func (AsesmenDokter) TableName() string {
	return "vicore_rme.dcppt_soap_dokter"
}

func (DCPPT) TableName() string {
	return "vicore_rme.dcppt"
}

func (DPenLabBedah) TableName() string {
	return "kamop.dlap_bedah"
}
