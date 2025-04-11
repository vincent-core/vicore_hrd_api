package dto

type (
	DataResponseRegisterPasien struct {
		Tanggal    string `json:"tanggal"`
		Id         string `json:"no_rm"`
		Noreg      string `json:"noreg"`
		Nama       string `json:"nama"`
		Kunjungan  string `json:"kunjungan"`
		Keterangan string `json:"keterangan"`
		Pelayaan   string `json:"pelayanan"`
		Bagian     string `json:"bagian"`
		KdBag      string `json:"kd_bagian"`
	}
)
