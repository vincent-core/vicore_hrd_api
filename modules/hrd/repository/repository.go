package repository

import (
	"errors"
	"fmt"
	"vicore_hrd/modules/hrd"
	"vicore_hrd/modules/hrd/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type hrdRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewHisRepository(db *gorm.DB, logging *logrus.Logger) entity.VicoreHRDRepository {
	return &hrdRepository{
		DB:      db,
		Logging: logging,
	}
}

func (lu *hrdRepository) OnFindAllDataKaryawanRepository() (res []hrd.Kemployee, err error) {
	result := lu.DB.Where(hrd.Kemployee{}).Find(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *hrdRepository) FindHRDByEmailRepository(EmailStr string) (res hrd.Kemployee, err error) {
	result := lu.DB.Where(hrd.Kemployee{
		Email: EmailStr,
	}).First(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *hrdRepository) OnGetDataDokterRepo(IDDokter string) (res hrd.Dokter, err error) {
	query := `SELECT nik, iddokter, jeniskelamin, namadokter, spesialisasi FROM his.ktaripdokter where iddokter=? LIMIT 1`

	result := lu.DB.Raw(query, IDDokter).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *hrdRepository) OnFindPelayananRepository(KDPelayanan string) (res hrd.KPelayanan, err error) {
	result := lu.DB.Where(hrd.KPelayanan{KdBag: KDPelayanan}).Order("no_urut ASC").First(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}
