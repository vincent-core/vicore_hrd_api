package repository

import (
	"errors"
	"fmt"
	"vicore_hrd/modules/lib"
	"vicore_hrd/modules/lib/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type libRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewLibRepository(db *gorm.DB, logging *logrus.Logger) entity.LibRepository {
	return &libRepository{
		DB:      db,
		Logging: logging,
	}
}

func (lu *libRepository) FindAllPelayananRepository() (res []lib.KPelayanan, err error) {
	result := lu.DB.Where(lib.KPelayanan{AsesmenActive: true}).Order("no_urut ASC").Find(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *libRepository) OnGetDataRekamMedis() (res []lib.DRekamMedis, err error) {
	query := "SELECT nama_rm, kode_rm, link_url FROM vicore_rme.drekam_medis limit 100"

	result := lu.DB.Raw(query).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil

}
