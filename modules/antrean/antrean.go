package antrean

import "vicore_hrd/modules/lib"

type (
	AntrianPoliIGD struct {
		Nama         string
		Kodedr       string
		Dokter       string
		Id           string
		Noreg        string
		NoBook       string
		NoAntrian    string
		Jeniskelamin string
		Umurth       int
		RegType      string
		Status       string
		Tgllahir     string
	}

	DepengkajianKeperawatan struct {
		KdBagian   string `json:"kd_bagian"`
		Noreg      string
		Usia       string
		Pelayanan  string
		Anamnesa   string
		CaraMasuk  string
		KPelayanan lib.KPelayanan `gorm:"foreignKey:KdBagian" json:"bagian"`
	}

	DAsesmenDokter struct {
		Noreg      string
		Pelayanan  string         `json:"pelayanan"`
		KdBagian   string         `json:"kd_bagian"`
		KPelayanan lib.KPelayanan `gorm:"foreignKey:KdBagian" json:"bagian"`
	}

	KbangsalKasur struct {
		Kodebangsal string
		Kamar       string
		Kasur       string
		Id          string
		Noreg       string
		Kodedr      string
		Nama        string
		Sex         string
		Umur        string
		Tgllahir    string
		Ket         string
		Dokter      string
		Kelas       string
	}

	PasienRANAP struct {
		Noreg   string
		Tanggal string
	}

	DRegisterPasien struct {
		Jam        string
		Tanggal    string
		Id         string
		Noreg      string
		Nama       string
		Kunjungan  string
		Keterangan string
		Pelayaan   string
		Bagian     string
		KdBag      string
	}
)

func (KbangsalKasur) TableName() string {
	return "his.kbangsalkasur"
}

func (DepengkajianKeperawatan) TableName() string {
	return "vicore_rme.dpengkajian_keperawatan"
}

func (DAsesmenDokter) TableName() string {
	return "vicore_rme.dcppt_soap_dokter"
}
