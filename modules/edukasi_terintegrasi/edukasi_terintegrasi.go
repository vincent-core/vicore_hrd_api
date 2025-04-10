package edukasiterintegrasi

type (
	DedukasiTerintegrasi struct {
		IdEdukasi         int    `gorm:"primaryKey;autoIncrement;column:id_edukasi"`
		InsertDttm        string `gorm:"column:insert_dttm;default:0000-00-00 00:00:00"`
		InsertUser        string `gorm:"column:insert_user;size:50"`
		NoRm              string `gorm:"column:no_rm;size:50"`
		NoReg             string `gorm:"column:no_reg;size:225"`
		KdBagian          string `gorm:"column:kd_bagian;size:50"`
		Informasi         string `gorm:"column:informasi;type:text"`
		Metode            string `gorm:"column:metode;type:text"`
		PemberiInformasi  string `gorm:"column:pemberi_informasi;size:225"`
		PenerimaInformasi string `gorm:"column:penerima_informasi;size:225"`
		Evaluasi          string `gorm:"column:evaluasi;type:text"`
	}
)

func (DedukasiTerintegrasi) TableName() string {
	return "vicore_rme.dedukasi_terintegrasi"
}
