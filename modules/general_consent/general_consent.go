package generalconsent

import "time"

type (
	Pasien struct {
		NamaPasien      string `json:"nama_pasien"`
		Nik             string `json:"nik"`
		TanggalLahir    string `json:"tgl_lahir"`
		NomorRekamMedis string `json:"no_rm"`
		Alamat          string `json:"alamat"`
		NoHp            string `json:"no_hp"`
	}

	DokterAsesmen struct {
		UserId     string
		NamaDokter string
		KonsulKe   string
	}

	PenangungJawab struct {
		Nama         string `json:"nama"`
		TanggalLahir string `json:"tgl_lahir"`
		Alamat       string `json:"alamat"`
		NoHP         string `json:"no_hp"`
	}

	PengkajianKeperawatan struct {
		IDPengkajian             int              `gorm:"column:id_pengkajian;primaryKey;autoIncrement"`
		InsertDttm               time.Time        `gorm:"column:insert_dttm;default:0000-00-00 00:00:00"`
		KdBagian                 string           `gorm:"column:kd_bagian;size:50"`
		UserID                   string           `gorm:"column:user_id;size:50"`
		Noreg                    string           `gorm:"column:noreg;size:225"`
		Usia                     string           `gorm:"column:usia;type:enum('DEWASA','ANAK');default:DEWASA"`
		Pelayanan                string           `gorm:"column:pelayanan;type:enum('RAJAL','RANAP')"`
		Anamnesa                 string           `gorm:"column:anamnesa;size:225"`
		CaraMasuk                string           `gorm:"column:cara_masuk;size:225"`
		AsalPasien               string           `gorm:"column:asal_pasien;size:225"`
		KeluhanUtama             string           `gorm:"column:keluhan_utama;size:225"`
		RiwayatPenyakitSekarang  string           `gorm:"column:riwayat_penyakit_sekarang;type:text"`
		RiwayatPenyakitDahulu    string           `gorm:"column:riwayat_penyakit_dahulu;type:text"`
		RiwayatPenyakitKeluarga  string           `gorm:"column:riwayat_penyakit_keluarga;type:text"`
		RiwayatPengobatanDirumah string           `gorm:"column:riwayat_pengobatan_dirumah;type:enum('TIDAK ADA','ADA')"`
		RiwayatAlergi            string           `gorm:"column:riwayat_alergi;type:enum('TIDAK ADA','ADA')"`
		ReaksiAlergi             string           `gorm:"column:reaksi_alergi;size:225"`
		Nyeri                    string           `gorm:"column:nyeri;type:enum('TIDAK ADA','ADA')"`
		RiwayatKehamilan         string           `gorm:"column:riwayat_kehamilan;"`
		TindakLanjut             string           `gorm:"column:tindak_lanjut;"`
		AlasanRujukan            string           `gorm:"column:alasan_rujukan;"`
		KeteranganRujuk          string           `gorm:"column:keterangan_rujuk;"`
		KondisiPasien            string           `gorm:"column:kondisi_pasien;"`
		Indikasi                 string           `gorm:"column:indikasi;"`
		TransportasiPulang       string           `gorm:"column:transportasi_pulang;"`
		DPJP                     string           `gorm:"column:dpjp;"`
		Ruangan                  string           `gorm:"column:ruangan;"`
		PendidikanPasien         string           `gorm:"column:pendidikan_pasien;"`
		DischargePlanning        string           `gorm:"column:discharge_planning;"`
		Perawat                  UserPerawatModel `gorm:"foreignKey:UserID" json:"user"`
	}

	Pengkajian struct {
		Noreg       string
		KdBagian    string
		Usia        string
		Pelayanan   string
		NamaPerawat string
	}

	DataPengkajianKeperawatan struct {
		KdBagian  string
		Usia      string
		Pelayanan string
		Noreg     string
	}

	UserPerawatModel struct {
		Idperawat   string `gorm:"primaryKey;column:idperawat" json:"id_perawat"`
		Namaperawat string `json:"nama"`
	}

	DataRegister struct {
		Id    string
		Noreg string
		Nama  string
	}

	DGeneralConsent struct {
		IDGeneral       int    `gorm:"column:id_general;primaryKey;autoIncrement"`
		InsertDttm      string `gorm:"column:insert_dttm;not null;default:'0000-00-00 00:00:00'"`
		UserID          string `gorm:"column:user_id;size:225;not null;default:''"`
		NoRM            string `gorm:"column:no_rm;size:50;not null;default:''"`
		KdBagian        string `gorm:"column:kd_bagian;size:225;default:''"`
		Pelayanan       string `gorm:"column:pelayanan;type:enum('RAJAL', 'RANAP');default:null"`
		PJawabNama      string `gorm:"column:pjawab_nama;size:225;not null;default:''"`
		PJawabAlamat    string `gorm:"column:pjawab_alamat;size:225;not null;default:''"`
		PJawabTglLahir  string `gorm:"column:pjawab_tgllahir;size:225;not null;default:''"`
		PJawabNoHP      string `gorm:"column:pjawab_nohp;size:225;not null;default:''"`
		Pewenang        string `gorm:"column:pewenang;type:text"`
		HubDenganPasien string `gorm:"column:hub_pasien;type:text" json:"hub_pasien"`
		TtdPjawab       string `gorm:"column:ttd_pjawab;type:text" json:"ttd_pjawab"`
		Privasi         string `gorm:"column:privasi;type:text" json:"privasi"`
		Noreg           string `gorm:"column:noreg;type:text" json:"noreg"`
		Ruangan         string `gorm:"column:ruangan;type:text" json:"ruangan"`
		Kelas           string `gorm:"column:kelas;type:text" json:"kelas"`
	}
)

func (PengkajianKeperawatan) TableName() string {
	return "vicore_rme.dpengkajian_keperawatan"
}

func (DGeneralConsent) TableName() string {
	return "vicore_rme.dgeneral_consent"
}

func (UserPerawatModel) TableName() string {
	return "his.kperawat"
}
