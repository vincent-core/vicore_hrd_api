package his

type (
	KTaripDokterModel struct {
		Iddokter     string `gorm:"primaryKey;column:iddokter" json:"id_dokter"`
		Namadokter   string
		Alamat       string
		Jeniskelamin string
		Pendidikan   string
		Statusdokter string
		Spesialisasi string
	}

	UserPerawatModel struct {
		Idperawat     string `gorm:"primaryKey;column:idperawat" json:"id_perawat"`
		Namaperawat   string `json:"nama"`
		Alamat        string `json:"alamat"`
		Jeniskelamin  string `json:"jenis_kelamin"`
		Statusperawat string `json:"status"`
	}
)

func (KTaripDokterModel) TableName() string {
	return "his.ktaripdokter"
}

func (UserPerawatModel) TableName() string {
	return "his.kperawat"
}
