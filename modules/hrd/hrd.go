package hrd

import (
	"time"
)

type (

	//  GET KTARIPDKONTER

	Dokter struct {
		Nik          string
		IdDokter     string
		Jeniskelamin string
		Spesialisasi string
		Namadokter   string
	}
	// DATA TABLE HRD
	Kemployee struct {
		IDK              string    `gorm:"column:idk;primaryKey" json:"idk"`
		KeteranganPerson string    `gorm:"column:keterangan_person" json:"keterangan_person" gorm:"default:dokter"`
		Kelas            string    `gorm:"column:kelas" json:"kelas" gorm:"default:KELAS 3"`
		Plafon           float64   `gorm:"column:plafon" json:"plafon" gorm:"default:0"`
		Password         string    `gorm:"column:password" json:"password"`
		Ktp              string    `gorm:"column:ktp" json:"ktp"`
		AlamatSekarang   string    `gorm:"column:alamat_sekarang" json:"alamat_sekarang"`
		Photo            string    `gorm:"column:photo" json:"photo"`
		Nomor            string    `gorm:"column:nomor" json:"nomor"`
		Bagian           string    `gorm:"column:bagian" json:"bagian"`
		TglMasuk         time.Time `gorm:"column:tglmasuk" json:"tglmasuk" gorm:"default:0000-00-00"`
		Keterangan       string    `gorm:"column:keterangan" json:"keterangan"`
		Nama             string    `gorm:"column:nama" json:"nama"`
		JenisKelamin     string    `gorm:"column:jeniskelamin" json:"jeniskelamin"`
		TglLahir         time.Time `gorm:"column:tgllahir" json:"tgllahir" gorm:"default:0000-00-00"`
		TempatLahir      string    `gorm:"column:tempatlahir" json:"tempatlahir"`
		Usia             int       `gorm:"column:usia" json:"usia" gorm:"default:0"`
		Agama            string    `gorm:"column:agama" json:"agama"`
		Alamat           string    `gorm:"column:alamat" json:"alamat"`
		DinasMalam       string    `gorm:"column:dinasmalam" json:"dinasmalam" gorm:"default:'Tidak'"`
		Pendidikan       string    `gorm:"column:pendidikan" json:"pendidikan"`
		Jurusan          string    `gorm:"column:jurusan" json:"jurusan"`
		ThnLulus         int       `gorm:"column:thnlulus" json:"thnlulus" gorm:"default:0"`
		NoIjazah         string    `gorm:"column:noijazah" json:"noijazah"`
		Str              string    `gorm:"column:str" json:"str"`
		TerbitStr        time.Time `gorm:"column:terbitstr" json:"terbitstr" gorm:"default:0000-00-00"`
		BerlakuStr       time.Time `gorm:"column:berlakustr" json:"berlakustr" gorm:"default:0000-00-00"`
		SikSipa          string    `gorm:"column:siksipa" json:"siksipa"`
		BerlakuSikSipa   time.Time `gorm:"column:berlakusiksipa" json:"berlakusiksipa" gorm:"default:0000-00-00"`
		Verifikasi       string    `gorm:"column:verifikasi" json:"verifikasi"`
		Status           string    `gorm:"column:status" json:"status"`
		Anak             string    `gorm:"column:anak" json:"anak"`
		Telp             string    `gorm:"column:telp" json:"telp"`
		HP               string    `gorm:"column:hp" json:"hp"`
		Fax              string    `gorm:"column:fax" json:"fax"`
		Email            string    `gorm:"column:email" json:"email"`
		NamaInstansi     string    `gorm:"column:namainstansi" json:"namainstansi"`
		AlamatInstansi   string    `gorm:"column:alamatinstansi" json:"alamatinstansi"`
		Posisi           string    `gorm:"column:posisi" json:"posisi"`
		KodePos          string    `gorm:"column:kodepos" json:"kodepos"`
		Univ             string    `gorm:"column:univ" json:"univ"`
		Kota             string    `gorm:"column:kota" json:"kota"`
		Negara           string    `gorm:"column:negara" json:"negara"`
		Bank             string    `gorm:"column:bank" json:"bank"`
		Account          string    `gorm:"column:account" json:"account"`
		AtasNama         string    `gorm:"column:atasnama" json:"atasnama"`
		Cabang           string    `gorm:"column:cabang" json:"cabang"`
		Office           string    `gorm:"column:office" json:"office" gorm:"default:false"`
		Peringatan       int       `gorm:"column:peringatan" json:"peringatan" gorm:"default:0"`
		Teguran          int       `gorm:"column:teguran" json:"teguran" gorm:"default:0"`
		Berkala          string    `gorm:"column:berkala" json:"berkala"`
		KodePelayanan    string    `gorm:"column:kode_pelayanan"`
	}

	KPelayanan struct {
		KdBag     string `gorm:"primaryKey;column:kd_bag" json:"kd_bag"`
		Bagian    string `json:"bagian"`
		Pelayanan string `json:"pelayanan"`
		NoUrut    int    `column:"no_urut"`
	}
)

func (Kemployee) TableName() string {
	return "vicore_hrd.kemployee"
}

func (KPelayanan) TableName() string {
	return "vicore_lib.kpelayanan"
}
