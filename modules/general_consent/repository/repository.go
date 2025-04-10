package repository

import (
	"errors"
	"fmt"
	generalconsent "vicore_hrd/modules/general_consent"
	"vicore_hrd/modules/general_consent/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type generalConsenstRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewGeneralConsentRepository(db *gorm.DB, logging *logrus.Logger) entity.GeneralConsentRepository {
	return &generalConsenstRepository{
		DB:      db,
		Logging: logging,
	}
}

func (ig *generalConsenstRepository) OnGetNamaDokterRepo(NoReg string) (res generalconsent.DokterAsesmen, err error) {
	query := `SELECT insert_user_id as user_id, asesmed_konsul_ke as konsul_ke, kp.namadokter as nama_dokter  FROM vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON kp.iddokter=dp.insert_user_id WHERE noreg=? LIMIT 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *generalConsenstRepository) OnGetPasienGeneralConsentRepository(ID string) (res generalconsent.Pasien, err error) {
	query := `SELECT nik as nik, id as  nomor_rekam_medis, firstname as nama_pasien, lastname, jeniskelamin, tgllahir as tanggal_lahir, hp as no_hp, alamat  from his.dprofilpasien WHERE id=? limit 1`

	result := ig.DB.Raw(query, ID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *generalConsenstRepository) OnGetPengkajianKeperawatanRepository(NoReg string, KDBagian string) (res generalconsent.DataPengkajianKeperawatan, er error) {

	query := `SELECT kd_bagian, usia, pelayanan, noreg from vicore_rme.dpengkajian_keperawatan  where noreg=? AND kd_bagian=? limit 1`

	result := ig.DB.Raw(query, NoReg, KDBagian).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil

}

func (ig *generalConsenstRepository) OnGetDataDRegisterRepository(NoReg string) (res generalconsent.DataRegister, err error) {
	query := `SELECT id, noreg, nama from rekam.dregister where noreg=? limit 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *generalConsenstRepository) OnGetPengkajianKeperawatanForGeneralConsent(NoReg string, KDBagian string) (res generalconsent.PengkajianKeperawatan, err error) {

	result := ig.DB.Model(&res).Where(&generalconsent.PengkajianKeperawatan{
		Noreg: NoReg, KdBagian: KDBagian,
	}).Preload("Perawat").First(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (ig *generalConsenstRepository) OnGetPengkajianKeperawatanCodeRepository(NoReg string, KDBagian string) (res generalconsent.Pengkajian, err error) {
	query := `SELECT pk.noreg as noreg, pk.usia as usia,  pk.pelayanan, pk.kd_bagian, kp.namaperawat as nama_perawat 
	from vicore_rme.dpengkajian_keperawatan as pk INNER JOIN his.kperawat as kp on pk.user_id=kp.idperawat
	WHERE pk.noreg=? AND pk.kd_bagian=? limit 1`

	result := ig.DB.Raw(query, NoReg, KDBagian).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *generalConsenstRepository) OnGetGeneralConsentRAJALRepository(noRM string, kdBagian string) (res generalconsent.DGeneralConsent, err error) {
	result := ig.DB.Model(&res).Where(&generalconsent.DGeneralConsent{
		KdBagian: kdBagian, NoRM: noRM, Pelayanan: "RAJAL",
	}).Find(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (ig *generalConsenstRepository) OnGetDataGeneralConsentRepo(noRM string) (res generalconsent.DGeneralConsent, err error) {
	result := ig.DB.Model(&res).Where(&generalconsent.DGeneralConsent{
		NoRM: noRM}).Find(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (ig *generalConsenstRepository) OnGetDataGeneralConsentRepoV2(NoReg string, Pelayanan string, ID string) (res generalconsent.DGeneralConsent, err error) {
	result := ig.DB.Model(&res).Where(&generalconsent.DGeneralConsent{Noreg: NoReg, Pelayanan: Pelayanan, NoRM: ID}).Find(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}
