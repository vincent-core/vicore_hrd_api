package repository

import (
	"errors"
	"fmt"
	"vicore_hrd/modules/triase"
	"vicore_hrd/modules/triase/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type triaseRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewTriaseRepository(db *gorm.DB, logging *logrus.Logger) entity.TriaseRepository {
	return &triaseRepository{
		DB:      db,
		Logging: logging,
	}
}

func (sr *triaseRepository) GetDVitalSignGDRepository(noreg string) (res triase.DVitalSign, err error) {
	query := `SELECT insert_dttm, kategori, ket_person, noreg, tb, td, bb,  nadi, suhu, spo2, pernafasan FROM vicore_rme.dvital_sign where kd_bagian = ? AND noreg=? AND ket_person=?  LIMIT 1`

	result := sr.DB.Raw(query, "IGD001", noreg, "dokter").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) GetDVitalSignPONEKRepository(noreg string) (res triase.DVitalSign, err error) {
	query := `SELECT insert_dttm, kategori, ket_person, noreg, tb, td, bb,  nadi, suhu, spo2, pernafasan FROM vicore_rme.dvital_sign where kd_bagian = ? AND noreg=? AND ket_person=?  LIMIT 1`

	result := sr.DB.Raw(query, "PONEK", noreg, "dokter").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (kb *triaseRepository) OnGetAsesmenTriaseIGDRepository(noReg string, userID string, kdBagian string, pelayanan string) (res triase.AsesmenTriaseIGD, err error) {

	result := kb.DB.Model(&res).Where("noreg = ? AND insert_user_id=? AND kd_bagian=? AND pelayanan=?", noReg, userID, kdBagian, pelayanan).Scan(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (sr *triaseRepository) OnGetPemfisikTriaseRepo(NoReg string) (res triase.TriaseDPemFisik, err error) {
	query := `SELECT noreg, kesadaran, gcs_e as e, gcs_v as v, gcs_m as m, jalan_nafas, pernafasan, sirkulasi,  kd_bagian, akral , pupil, refleks FROM vicore_rme.dpem_fisik WHERE noreg=?  AND kd_bagian=? AND ket_person=? LIMIT 1`

	result := sr.DB.Raw(query, NoReg, "IGD001", "Dokter").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) OnGetPemfisikTriasePONEKRepo(NoReg string) (res triase.TriaseDPemFisik, err error) {
	query := `SELECT noreg, kesadaran, gcs_e as e, gcs_v as v, gcs_m as m, jalan_nafas, pernafasan, sirkulasi,  kd_bagian, akral , pupil, refleks FROM vicore_rme.dpem_fisik WHERE noreg=?  AND kd_bagian=? AND ket_person=? LIMIT 1`

	result := sr.DB.Raw(query, NoReg, "IGD001", "Dokter").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) TriaseAsesmenRepo(NoReg string) (res triase.TriaseAsesmen, err error) {
	query := `Select kd_bagian, asesmed_keluh_utama as keluhan_utama, insert_user_id as user_id , kp.namadokter as nama_dokter From vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON dp.insert_user_id = kp.iddokter WHERE noreg=? and kd_bagian=? AND pelayanan=? LIMIT 1`

	result := sr.DB.Raw(query, NoReg, "IGD001", "rajal").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) TriaseAsesmenRepoKebidanan(NoReg string, KdBagian string) (res triase.TriaseAsesmen, err error) {
	query := `Select kd_bagian, asesmed_keluh_utama as keluhan_utama, insert_user_id as user_id , kp.namadokter as nama_dokter From vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON dp.insert_user_id = kp.iddokter WHERE noreg=? and kd_bagian=? AND pelayanan=? LIMIT 1`

	result := sr.DB.Raw(query, NoReg, KdBagian, "rajal").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) GetAsesmenIGD() (res []triase.IgdAsesmen, err error) {
	query := `SELECT insert_dttm, kd_bagian, asesmed_diagP, noreg from vicore_rme.dcppt_soap_dokter WHERE asesmed_diagP =""`

	result := sr.DB.Raw(query).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

//

func (sr *triaseRepository) GetNamaDokter(Noreg string) (res triase.DokterNama, err error) {
	query := `SELECT dokter FROM his.dpoliugd6 WHERE noreg=?`

	result := sr.DB.Raw(query, Noreg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) GetNamaDokterDiagnosa(Noreg string) (res triase.DokterNama, err error) {
	query := `SELECT noreg, kodediagnosa FROM his.dpolidalam2 LIMIT 1 where noreg=?`

	result := sr.DB.Raw(query, Noreg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) GetDiagnosa(Noreg string) (res triase.DokterDiagnosa, err error) {
	query := `SELECT noreg, kodediagnosa FROM his.dpolidalam2 where noreg=? LIMIT 1 `

	result := sr.DB.Raw(query, Noreg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) UpdateAsesmedKonsulKe(Dokter string, Noreg string) (res triase.IgdAsesmen, err error) {
	query := `UPDATE vicore_rme.dcppt_soap_dokter SET asesmed_diagP=? WHERE noreg=? AND kd_bagian="IGD001" AND pelayanan="rajal"`

	result := sr.DB.Raw(query, Dokter, Noreg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) UpDateDiagnosa(Diagnosa string, Noreg string, KdBagian string) (res triase.IgdAsesmen, err error) {
	query := `UPDATE vicore_rme.dcppt_soap_dokter SET asesmed_diagP=? WHERE noreg=? AND kd_bagian=? `

	result := sr.DB.Raw(query, Diagnosa, Noreg, KdBagian).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) OnGetAsesmenNyeriRepo(NoReg string) (res triase.AsesmenUlangNyeri, err error) {
	query := `SELECT  asesmen,  noreg, kd_bagian, metode, skor_nyeri from vicore_rme.dasesmen_ulang_nyeri WHERE noreg=? AND kd_bagian=? AND asesmen=? LIMIT 1`

	result := sr.DB.Raw(query, NoReg, "IGD001", "AWAL").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) OnGetAsesmenNyeriPONEKRepo(NoReg string) (res triase.AsesmenUlangNyeri, err error) {
	query := `SELECT  asesmen,  noreg, kd_bagian, metode, skor_nyeri from vicore_rme.dasesmen_ulang_nyeri WHERE noreg=? AND kd_bagian=? AND asesmen=? LIMIT 1`

	result := sr.DB.Raw(query, NoReg, "PONEK", "AWAL").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) OnGetAsesmenKeperawatanRepo(NoReg string) (res triase.AsesmenKeperawatan, err error) {
	query := `SELECT noreg, kd_bagian, cara_masuk, riwayat_alergi, keluhan_utama, nyeri 
				FROM vicore_rme.dpengkajian_keperawatan WHERE noreg=? AND kd_bagian='IGD001'  LIMIT 1`

	result := sr.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *triaseRepository) OnGetAsesmenKeperawatanPONEKRepo(NoReg string) (res triase.AsesmenKeperawatan, err error) {
	query := `SELECT noreg, kd_bagian, cara_masuk, riwayat_alergi, keluhan_utama, nyeri 
				FROM vicore_rme.dpengkajian_keperawatan WHERE noreg=? AND kd_bagian='PONEK'  LIMIT 1`

	result := sr.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (kb *triaseRepository) OnGetReportAsesmenTriaseIGDRepository(noReg string) (res triase.AsesmenTriaseIGD, err error) {
	result := kb.DB.Model(&triase.AsesmenTriaseIGD{}).Where(triase.AsesmenTriaseIGD{
		Noreg: noReg, KdBagian: "IGD001", Pelayanan: "rajal",
	}).First(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (kb *triaseRepository) OnGetReportAsesmenTriasePONEKRepository(noReg string) (res triase.AsesmenTriaseIGD, err error) {
	result := kb.DB.Model(&triase.AsesmenTriaseIGD{}).Where(triase.AsesmenTriaseIGD{
		Noreg: noReg, KdBagian: "PONEK", Pelayanan: "rajal",
	}).First(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (kb *triaseRepository) GetAsesmenDokterTriaseRepo(noReg string) (res triase.TriaseModel, err error) {
	query := `SELECT insert_dttm, noreg, kd_bagian, tgl_masuk, jam_masuk, pelayanan from vicore_rme.dcppt_soap_dokter where noreg=? AND kd_bagian=? LIMIT 1`

	result := kb.DB.Raw(query, noReg, "IGD001").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (kb *triaseRepository) GetAsesmenDokterTriasPONEKeRepo(noReg string) (res triase.TriaseModel, err error) {
	query := `SELECT insert_dttm, noreg, kd_bagian, tgl_masuk, jam_masuk, pelayanan from vicore_rme.dcppt_soap_dokter where noreg=? AND kd_bagian=? LIMIT 1`

	result := kb.DB.Raw(query, noReg, "PONEK").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

// ON GET REPORT TRIASE IGD DOKTER
func (kb *triaseRepository) OnGetReportTriaseIGDDokterRepository(noReg string) (res triase.TriaseIGDDokter, err error) {
	result := kb.DB.Model(&res).Where("noreg = ? AND kd_bagian=? AND pelayanan=?", noReg, "IGD001", "rajal").Find(&res)

	if result.Error != nil {
		return res, result.Error
	}
	return res, nil
}

// ON GET REPORT TRIASE IGD DOKTER
func (kb *triaseRepository) OnGetReportTriasePonekRepository(noReg string) (res triase.TriaseIGDDokter, err error) {
	result := kb.DB.Model(&res).Where("noreg = ? AND kd_bagian=? AND pelayanan=?", noReg, "PONEK", "rajal").Find(&res)

	if result.Error != nil {
		return res, result.Error
	}
	return res, nil
}

func (kb *triaseRepository) OnGetSkalaTriaseRepo(NoReg string) (res triase.Triase, err error) {
	query := `SELECT insert_user_id, noreg, kd_bagian, skala_triase_igd  FROM vicore_rme.dpem_fisik  WHERE noreg=? AND kd_bagian=? AND skala_triase_igd != ""  LIMIT 1`

	result := kb.DB.Raw(query, NoReg, "IGD001").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (kb *triaseRepository) OnGetSkalaTriasePONEKRepo(NoReg string) (res triase.Triase, err error) {
	query := `SELECT insert_user_id, noreg, kd_bagian, skala_triase_igd  FROM vicore_rme.dpem_fisik  WHERE noreg=? AND kd_bagian=? AND skala_triase_igd != ""  LIMIT 1`

	result := kb.DB.Raw(query, NoReg, "PONEK").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}
