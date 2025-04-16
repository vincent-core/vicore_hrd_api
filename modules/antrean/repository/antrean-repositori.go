package repository

import (
	"errors"
	"fmt"
	"vicore_hrd/modules/antrean"
	"vicore_hrd/modules/antrean/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type antreanRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewAntreanRepository(db *gorm.DB, logging *logrus.Logger) entity.AntreanRepository {
	return &antreanRepository{
		DB:      db,
		Logging: logging,
	}
}

func (lu *antreanRepository) GetAntrianUGD() (res []antrean.AntrianPoliIGD, err error) {
	query := "SELECT nama, kodedr, a.id AS id, noreg, no_book, a.reg_type no_antrian, umurth, status, tgllahir, b.jeniskelamin, c.namadokter AS dokter from his.antrianpoliugd AS a LEFT JOIN his.dprofilpasien AS b ON  a.id=b.id LEFT JOIN his.ktaripdokter AS c ON c.iddokter=a.kodedr"

	result := lu.DB.Raw(query).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *antreanRepository) OnGetDataPasienPulangByDateRepository(date string) (res []antrean.DRegisterPasien, err error) {
	query := `SELECT a.nama, a.id, a.noreg, a.bagian, b.tgllahir FROM rekam.dregister AS a 
LEFT JOIN his.dprofilpasien AS b ON a.id=b.id 
WHERE a.tanggal = ? AND a.bagian = 'Poli UGD'`
	result := lu.DB.Raw(query, date).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *antreanRepository) GetAntrianIGDDokterUmumRepository(KodeDokter string, dateFrom string, dateTo string) (res []antrean.AntrianPoliIGD, err error) {
	query := "SELECT nama, kodedr, a.id AS id, noreg, no_book, a.reg_type no_antrian, umurth, status, tgllahir, b.jeniskelamin, c.namadokter AS dokter from his.antrianpoliugd AS a LEFT JOIN his.dprofilpasien AS b ON  a.id=b.id LEFT JOIN his.ktaripdokter AS c ON c.iddokter=a.kodedr WHERE a.kodedr=? OR a.kodedr='' OR a.kodedr='NONE' AND a.jam >= ? AND a.jam < ? "

	result := lu.DB.Raw(query, KodeDokter, dateFrom, dateTo).Scan(&res)
	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal disimpan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (lu *antreanRepository) GetPasienBangsalForDokter(kodeBangsal string, kodeDokter string, dateFrom string, dateTo string) (res []antrean.KbangsalKasur, err error) {
	query := `SELECT a.kodebangsal, a.kamar, a.kasur, a.id, b.tgllahir AS tgllahir, a.noreg, a.kodedr, b.firstname AS nama, a.sex, a.ket, ds.namadokter AS dokter,  a.kasur AS kasur, a.kamar AS kamar 
FROM his.kbangsalkasur AS a 
INNER JOIN his.dprofilpasien AS b ON b.id=a.id 
LEFT JOIN his.ktaripdokter AS ds ON ds.iddokter=a.kodedr 
WHERE kodebangsal=? AND a.id !="" AND a.jam >= ? AND a.jam < ? `
	result := lu.DB.Raw(query, kodeBangsal, dateFrom, dateTo).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *antreanRepository) GetPasienBangsal(kodeBangsal string, dateFrom string, dateTo string) (res []antrean.KbangsalKasur, err error) {
	query := `SELECT a.kodebangsal, a.kamar, a.kasur, b.tgllahir AS tgllahir, a.id, a.noreg, a.kodedr, b.firstname as nama, b.jeniskelamin AS sex, a.ket, ds.namadokter AS dokter FROM his.kbangsalkasur AS a INNER JOIN his.dprofilpasien AS b ON b.id=a.id LEFT JOIN his.ktaripdokter AS ds ON ds.iddokter=a.kodedr WHERE kodebangsal=? AND a.id !="" AND a.jam >= ? AND a.jam < ?`

	result := lu.DB.Raw(query, kodeBangsal, dateFrom, dateTo).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *antreanRepository) OnGetPengkajianPerawatRepo(NoReg string) (res antrean.DepengkajianKeperawatan, err error) {
	result := lu.DB.Model(&res).Select("kd_bagian, noreg, usia, pelayanan, anamnesa, cara_masuk").Where(&antrean.DepengkajianKeperawatan{
		Noreg: NoReg,
	}).Preload("KPelayanan").Find(&res)

	if result.Error != nil {
		return res, result.Error
	}

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *antreanRepository) OnGetPengkajianDokterRepo(NoReg string) (res antrean.DAsesmenDokter, err error) {
	result := lu.DB.Model(&res).Select("kd_bagian, noreg,  pelayanan").Where(&antrean.DAsesmenDokter{
		Noreg: NoReg,
	}).Order("insert_dttm DESC").Preload("KPelayanan").Limit(1).Find(&res)

	if result.Error != nil {
		return res, result.Error
	}

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}
