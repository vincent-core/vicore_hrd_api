package lib

import "time"

type (
	KPelayanan struct {
		InsertDttm    time.Time `gorm:"column:insert_dttm;not null;default:'0000-00-00 00:00:00'"`
		UpdDttm       time.Time `gorm:"column:upd_dttm;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
		KdBag         string    `gorm:"column:kd_bag;size:100;not null;default:'';primaryKey"`
		Bagian        string    `gorm:"column:bagian;size:100;not null;default:''"`
		Pelayanan     string    `gorm:"column:pelayanan;type:enum('POLIKLINIK','IGD','PENUNJANG MEDIK','KAMAR BEDAH','KAMAR BERSALIN','ICU','RECOVERY','NURSE STATION','INSTALASI FARMASI','REKAM MEDIK','LOGISTIK FARMASI','UNIT O2','CSSD','MANAGER','RUMAH TANGGA','KEPERAWATAN','PPI','PMKP','MARKETING','AKUNTANSI','BILLING KASIR','DAPUR','LAIN-LAIN','ADMISSION','HRD','LOGISTIK UMUM','CUSTOMER SERVICES','CONSOLE','DISPLAY','LOGISTIK INVENTARIS','SATUSEHAT');not null;default:'POLIKLINIK'"`
		MapingBPJS    string    `gorm:"column:maping_bpjs;size:100;not null;default:''"`
		AsesmenActive bool      `gorm:"column:asesmen_active;default:null"`
		IdSS          string    `gorm:"column:id_SS;size:300;default:''"`
		IdentifierSS  string    `gorm:"column:identifier_SS;size:300;default:''"`
		LocationSS    string    `gorm:"column:location_SS;size:300;default:''"`
	}

	DRekamMedis struct {
		NamaRm  string `json:"nama_rm"`
		KodeRM  string `json:"kode_rm"`
		LinkUrl string `json:"link_url"`
	}
)

func (KPelayanan) TableName() string {
	return "vicore_lib.kpelayanan"
}
