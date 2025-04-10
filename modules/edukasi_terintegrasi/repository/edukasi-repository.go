package repository

import (
	"errors"
	"fmt"
	edukasiterintegrasi "vicore_hrd/modules/edukasi_terintegrasi"
	"vicore_hrd/modules/edukasi_terintegrasi/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type edukasiTerintegrasiRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewEdukasiTerintegrasi(db *gorm.DB, logging *logrus.Logger) entity.EdukasiTerintegrasiRepository {
	return &edukasiTerintegrasiRepository{
		DB:      db,
		Logging: logging,
	}
}

func (ar *edukasiTerintegrasiRepository) OnGetEdukasiTerintegrasiByNoRMRepository(NoRM string) (res []edukasiterintegrasi.DedukasiTerintegrasi, err error) {
	result := ar.DB.Model(&res).Where(&edukasiterintegrasi.DedukasiTerintegrasi{
		NoRm: NoRM,
	}).Find(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (edu *edukasiTerintegrasiRepository) OnGetEdukasiTerintegrasiRepository(NORM string) (res edukasiterintegrasi.DedukasiTerintegrasi, err error) {
	result := edu.DB.Where(edukasiterintegrasi.DedukasiTerintegrasi{
		NoRm: NORM,
	}).Updates(&res).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal disimpan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (edu *edukasiTerintegrasiRepository) OnUpdateEdukasiTerintegrasiRepository(ID int, data edukasiterintegrasi.DedukasiTerintegrasi) (res edukasiterintegrasi.DedukasiTerintegrasi, err error) {

	result := edu.DB.Where(&edukasiterintegrasi.DedukasiTerintegrasi{
		IdEdukasi: ID,
	}).Updates(&data).Scan(&res)

	if result.Error != nil {
		return data, result.Error
	}

	return res, nil
}

func (edu *edukasiTerintegrasiRepository) OnSaveEdukasiTerintegrasiRepository(data edukasiterintegrasi.DedukasiTerintegrasi) (res edukasiterintegrasi.DedukasiTerintegrasi, err error) {
	result := edu.DB.Create(&data).Scan(&res)

	if result.Error != nil {
		return data, result.Error
	}
	return res, nil
}
