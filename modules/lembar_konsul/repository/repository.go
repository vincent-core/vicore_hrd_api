package repository

import (
	"errors"
	"fmt"
	lembarkonsul "vicore_hrd/modules/lembar_konsul"
	entity "vicore_hrd/modules/lembar_konsul/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type lembarKonsulRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewLembarKonsulRepository(db *gorm.DB, logging *logrus.Logger) entity.LembarKonsulRepository {
	return &lembarKonsulRepository{
		DB:      db,
		Logging: logging,
	}
}

func (ig *lembarKonsulRepository) OnGetDataRegisterRepository(NoReg string) (res lembarkonsul.DRegister, err error) {
	query := `SELECT id, noreg, nama from rekam.dregister where noreg=? limit 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *lembarKonsulRepository) OnGetDataKonsulenPasienRepo(NoReg string) (res lembarkonsul.DKonsulPasien, err error) {
	query := `SELECT * FROM vicore_rme.dkonsul_pasien WHERE noreg=? limit 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}
func (ig *lembarKonsulRepository) OnGetListDataKonsulenPasienRepo(NoReg string) (res []lembarkonsul.DKonsulPasien, err error) {
	query := `SELECT * FROM vicore_rme.dkonsul_pasien WHERE noreg=? LIMIT 5`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (ig *lembarKonsulRepository) OnGetDataKonsulPasienRepository(NoReg string) (res lembarkonsul.DKonsulPasien, err error) {
	query := `
	SELECT insert_dttm, jenis_konsul, dokter_konsul, kp.namadokter as dokter_konsul, iktisar_klinik, kd_bagian, noreg, konsul_ke, kp1.namadokter as dokter_konsul_ke FROM vicore_rme.dkonsul_pasien as dp INNER JOIN his.ktaripdokter as kp ON kp.iddokter = dp.dokter_konsul  INNER JOIN his.ktaripdokter as kp1 ON kp1.iddokter = dp.konsul_ke where noreg=? limit 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (ig *lembarKonsulRepository) OnGetCPPTKonsulenRepository(NoReg string, UserID string) (res lembarkonsul.CpptKonsulen, err error) {
	query := `SELECT insert_dttm, ppa, insert_user_id, kd_bagian, noreg, subjektif, objektif, asesmen, plan FROM vicore_rme.dcppt WHERE noreg=? AND ppa=?  AND insert_user_id=? limit 1`

	result := ig.DB.Raw(query, NoReg, "konsulen", UserID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}
