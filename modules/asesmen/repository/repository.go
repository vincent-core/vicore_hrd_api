package repository

import (
	"errors"
	"fmt"
	"vicore_hrd/modules/asesmen"
	"vicore_hrd/modules/asesmen/dto"
	"vicore_hrd/modules/asesmen/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type asesmenRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewAsesmenRepository(db *gorm.DB, logging *logrus.Logger) entity.AsesmenRepository {
	return &asesmenRepository{
		DB:      db,
		Logging: logging,
	}
}

// * PENGKAJIAN REPOSITORY
func (ar *asesmenRepository) OnGetPengkajianKeperawatanRepository(kdBagian string, pelayanan string, noReg string) (res asesmen.PengkajianKeperawatan, err error) {
	result := ar.DB.Model(&res).Where(&asesmen.PengkajianKeperawatan{
		KdBagian: kdBagian, Pelayanan: pelayanan, Noreg: noReg,
	}).Preload("Perawat").Find(&res)

	if result.Error != nil {
		return res, result.Error
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetAsesmenDokterRANAPRepository(noReg string, kdBagian string) (res asesmen.AsesmenDokter, err error) {
	results := ig.DB.Select("insert_dttm, noreg, asesmed_keluh_utama, asesmed_rwyt_skrg, asesmed_rwyt_dahulu, pelayanan, kd_bagian, keterangan_person, insert_user_id").Where(asesmen.AsesmenDokter{
		Noreg: noReg, Pelayanan: "ranap", KdBagian: kdBagian,
	}).Preload("Dokter").Find(&res)

	if results.Error != nil {
		return res, results.Error
	}

	return res, nil
}

func (ig *asesmenRepository) OnGeAsesmenDokterRepository(noReg string, kdBagian string) (res asesmen.AsesmenDokter, err error) {
	results := ig.DB.Select("insert_dttm, noreg, asesmed_keluh_utama, asesmed_rwyt_skrg, asesmed_rwyt_dahulu, pelayanan, kd_bagian, keterangan_person,insert_user_id").Where(asesmen.AsesmenDokter{
		Noreg: noReg, Pelayanan: "rajal", KdBagian: kdBagian,
	}).Preload("Dokter").Find(&res)

	if results.Error != nil {
		return res, results.Error
	}

	return res, nil
}

func (ar *asesmenRepository) OnUpdateCPPTByIDRepository(IDCppt int, data asesmen.DCPPT) (res asesmen.DCPPT, err error) {
	result := ar.DB.Where(asesmen.DCPPT{
		IDCppt: IDCppt,
	}).Updates(&data).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal disimpan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnSaveCPPTRepository(data asesmen.DCPPT) (res asesmen.DCPPT, err error) {

	result := ig.DB.Create(&data).Scan(&res)

	if result.Error != nil {
		return data, result.Error
	}
	return res, nil
}

func (ig *asesmenRepository) OnGetAsesmenCPPTRepository(NoRM string) (res []asesmen.DataCPPT, err error) {
	query := `SELECT dsp.id_cppt, dsp.insert_dttm, dsp.insert_user_id, dsp.kelompok, dsp.pelayanan, dsp.kd_bagian, dsp.noreg, dsp.dpjp, dsp.subjektif, dsp.situation, dsp.objektif, dsp.asesmen, dsp.plan, dsp.background, dsp.recomendation, dsp.instruksi_ppa, ke.namadokter AS nama_dokter_dpjp, ke2.nama AS nama_profesional, kp.bagian AS namabagian, dp.id FROM vicore_rme.dcppt AS dsp LEFT JOIN rekam.dregister AS dr ON dsp.noreg = dr.noreg 
		LEFT JOIN his.dprofilpasien AS dp ON dr.id = dp.id
		LEFT JOIN his.ktaripdokter AS ke ON dsp.dpjp = ke.iddokter
		LEFT JOIN mutiara.pengajar AS ke2 ON dsp.insert_user_id = ke2.id
		LEFT JOIN vicore_lib.kpelayanan AS kp ON dsp.kd_bagian = kp.kd_bag
	where dp.id =?
	ORDER BY dsp.insert_dttm DESC`

	result := ig.DB.Raw(query, NoRM).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetAsesmenCPPByNoRegTRepository(NoReg string) (res []asesmen.DataCPPT, err error) {
	query := `SELECT dsp.id_cppt, dsp.insert_dttm, dsp.insert_user_id, dsp.kelompok, dsp.pelayanan, dsp.kd_bagian, dsp.noreg, dsp.dpjp, dsp.subjektif, dsp.situation, dsp.objektif, dsp.asesmen, dsp.plan, dsp.background, dsp.recomendation, dsp.instruksi_ppa, ke.namadokter AS nama_dokter_dpjp, ke2.nama AS nama_profesional, kp.bagian AS namabagian, dp.id FROM vicore_rme.dcppt AS dsp LEFT JOIN rekam.dregister AS dr ON dsp.noreg = dr.noreg 
		LEFT JOIN his.dprofilpasien AS dp ON dr.id = dp.id
		LEFT JOIN his.ktaripdokter AS ke ON dsp.dpjp = ke.iddokter
		LEFT JOIN mutiara.pengajar AS ke2 ON dsp.insert_user_id = ke2.id
		LEFT JOIN vicore_lib.kpelayanan AS kp ON dsp.kd_bagian = kp.kd_bag
	where dr.noreg=? ORDER BY dsp.insert_dttm DESC`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

// GET ASESMEN DOKTER IGD
func (ig *asesmenRepository) OnGetAsesmenDokterIGDRepository(NoReg string) (res asesmen.AsesmenDokterIGD, err error) {
	query := `SELECT insert_dttm as waktu, kd_bagian as bagian, asesmen_prognosis  as prognosis ,  keterangan_person as person, noreg, insert_user_id as user_id, kp.namadokter as dokter, asesmed_keluh_utama as keluhan_utama, asesmed_rwyt_skrg as riwayat_sekarang, asesmed_rwyt_dahulu as riwayat_dahulu, asesmed_lokalis_image as image_lokalis, asesmed_konsul_ke as konsul_ke,  asesmed_terapi as terapi FROM vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON dp.insert_user_id = kp.iddokter  where kd_bagian="IGD001" and noreg =? LIMIT 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetAsesmenDokter(NoReg string) (res asesmen.AsesmenDokterIGD, err error) {
	query := `SELECT insert_dttm as waktu,dp.pelayanan as pelayanan, kd_bagian as bagian, asesmen_prognosis  as prognosis ,  keterangan_person as person, noreg, insert_user_id as user_id, kp.namadokter as dokter, asesmed_keluh_utama as keluhan_utama, asesmed_rwyt_skrg as riwayat_sekarang, asesmed_rwyt_dahulu as riwayat_dahulu, asesmed_lokalis_image as image_lokalis, asesmed_konsul_ke as konsul_ke,  asesmed_terapi as terapi FROM vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON dp.insert_user_id = kp.iddokter where dp.pelayanan=? and noreg=? LIMIT 1`

	result := ig.DB.Raw(query, "rajal", NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetAsesmenDokterRawatInapRepository(NoReg string) (res asesmen.AsesmenDokterIGD, err error) {
	query := `SELECT insert_dttm as waktu, kd_bagian as bagian, asesmen_prognosis  as prognosis ,  keterangan_person as person, noreg, insert_user_id as user_id, kp.namadokter as dokter, asesmed_keluh_utama as keluhan_utama, asesmed_rwyt_skrg as riwayat_sekarang, asesmed_rwyt_dahulu as riwayat_dahulu, asesmed_lokalis_image as image_lokalis,  asesmed_konsul_ke as konsul_ke,  asesmed_terapi as terapi FROM vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON dp.insert_user_id = kp.iddokter WHERE noreg =?  AND pelayanan="ranap" LIMIT 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetAsesmenDokterPONEKRepository(NoReg string) (res asesmen.AsesmenDokterIGD, err error) {
	query := `SELECT insert_dttm as waktu, kd_bagian as bagian,  cara_keluar, cara_keluar_detail,  ks.namadokter as konsul_ke,   asesmen_prognosis  as prognosis ,  keterangan_person as person, noreg, insert_user_id as user_id, kp.namadokter as dokter, asesmed_keluh_utama as keluhan_utama, asesmed_rwyt_skrg as riwayat_sekarang, asesmed_rwyt_dahulu as riwayat_dahulu, asesmed_lokalis_image as image_lokalis FROM vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON dp.insert_user_id = kp.iddokter LEFT JOIN his.ktaripdokter as ks on dp.konsul_dokter = ks.iddokter  where kd_bagian="PONEK" and noreg =? LIMIT 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *asesmenRepository) GetDiagnosaRepositoryReportIGD(NoReg string) (res []asesmen.DiagnosaResponse, err error) {
	var diag []asesmen.DiagnosaResponse

	type Diagnoses struct {
		P    string `json:"p"`
		Desp string `json:"desp"`
		S1   string `json:"s1"`
		Des1 string `json:"des1"`
		S2   string `json:"s2"`
		Des2 string `json:"des2"`
		S3   string `json:"s3"`
		Des3 string `json:"des3"`
		S4   string `json:"s4"`
		Des4 string `json:"des4"`
		S5   string `json:"s5"`
		Des5 string `json:"des5"`
		S6   string `json:"s6"`
		Des6 string `json:"des6"`
	}

	data := Diagnoses{}

	query := `
			SELECT asesmed_diagP AS p, b.description as desp, 
			asesmed_diagS1 AS s1, b1.description as des1, 
			asesmed_diagS2 AS s2, b2.description as des2,
			asesmed_diagS3 AS s3, b3.description as des3,
			asesmed_diagS4 AS s4, b4.description as des4,
			asesmed_diagS5 AS s5, b5.description as des5,
			asesmed_diagS6 AS s6, b6.description as des6
			FROM vicore_rme.dcppt_soap_dokter AS a 
			LEFT JOIN vicore_lib.k_icd10 AS b ON a.asesmed_diagP=b.code2 
			LEFT JOIN vicore_lib.k_icd10 AS b1 ON a.asesmed_diagS1=b1.code2
			LEFT JOIN vicore_lib.k_icd10 AS b2 ON a.asesmed_diagS2=b2.code2
			LEFT JOIN vicore_lib.k_icd10 AS b3 ON a.asesmed_diagS3=b3.code2
			LEFT JOIN vicore_lib.k_icd10 AS b4 ON a.asesmed_diagS4=b4.code2
			LEFT JOIN vicore_lib.k_icd10 AS b5 ON a.asesmed_diagS5=b5.code2
			LEFT JOIN vicore_lib.k_icd10 AS b6 ON a.asesmed_diagS6=b6.code2
			WHERE a.noreg=? AND a.kd_bagian=? LIMIT 1;
		`

	result := sr.DB.Raw(query, NoReg, "IGD001").Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	return diag, nil
}

func (sr *asesmenRepository) OnGetDiagnosa(NoReg string) (res []asesmen.DiagnosaResponse, err error) {
	var diag []asesmen.DiagnosaResponse

	type Diagnoses struct {
		P    string `json:"p"`
		Desp string `json:"desp"`
		S1   string `json:"s1"`
		Des1 string `json:"des1"`
		S2   string `json:"s2"`
		Des2 string `json:"des2"`
		S3   string `json:"s3"`
		Des3 string `json:"des3"`
		S4   string `json:"s4"`
		Des4 string `json:"des4"`
		S5   string `json:"s5"`
		Des5 string `json:"des5"`
		S6   string `json:"s6"`
		Des6 string `json:"des6"`
	}

	data := Diagnoses{}

	query := `
			SELECT asesmed_diagP AS p, b.description as desp, 
			asesmed_diagS1 AS s1, b1.description as des1, 
			asesmed_diagS2 AS s2, b2.description as des2,
			asesmed_diagS3 AS s3, b3.description as des3,
			asesmed_diagS4 AS s4, b4.description as des4,
			asesmed_diagS5 AS s5, b5.description as des5,
			asesmed_diagS6 AS s6, b6.description as des6
			FROM vicore_rme.dcppt_soap_dokter AS a 
			LEFT JOIN vicore_lib.k_icd10 AS b ON a.asesmed_diagP=b.code2 
			LEFT JOIN vicore_lib.k_icd10 AS b1 ON a.asesmed_diagS1=b1.code2
			LEFT JOIN vicore_lib.k_icd10 AS b2 ON a.asesmed_diagS2=b2.code2
			LEFT JOIN vicore_lib.k_icd10 AS b3 ON a.asesmed_diagS3=b3.code2
			LEFT JOIN vicore_lib.k_icd10 AS b4 ON a.asesmed_diagS4=b4.code2
			LEFT JOIN vicore_lib.k_icd10 AS b5 ON a.asesmed_diagS5=b5.code2
			LEFT JOIN vicore_lib.k_icd10 AS b6 ON a.asesmed_diagS6=b6.code2
			WHERE a.noreg=?  LIMIT 1;
		`

	result := sr.DB.Raw(query, NoReg).Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	return diag, nil
}

func (sr *asesmenRepository) GetDiagnosaRanapRepo(NoReg string, Bagian string) (res []asesmen.DiagnosaResponse, err error) {
	var diag []asesmen.DiagnosaResponse

	type Diagnoses struct {
		P    string `json:"p"`
		Desp string `json:"desp"`
		S1   string `json:"s1"`
		Des1 string `json:"des1"`
		S2   string `json:"s2"`
		Des2 string `json:"des2"`
		S3   string `json:"s3"`
		Des3 string `json:"des3"`
		S4   string `json:"s4"`
		Des4 string `json:"des4"`
		S5   string `json:"s5"`
		Des5 string `json:"des5"`
		S6   string `json:"s6"`
		Des6 string `json:"des6"`
	}

	data := Diagnoses{}

	query := `
			SELECT asesmed_diagP AS p, b.description as desp, 
			asesmed_diagS1 AS s1, b1.description as des1, 
			asesmed_diagS2 AS s2, b2.description as des2,
			asesmed_diagS3 AS s3, b3.description as des3,
			asesmed_diagS4 AS s4, b4.description as des4,
			asesmed_diagS5 AS s5, b5.description as des5,
			asesmed_diagS6 AS s6, b6.description as des6
			FROM vicore_rme.dcppt_soap_dokter AS a 
			LEFT JOIN vicore_lib.k_icd10 AS b ON a.asesmed_diagP=b.code2 
			LEFT JOIN vicore_lib.k_icd10 AS b1 ON a.asesmed_diagS1=b1.code2
			LEFT JOIN vicore_lib.k_icd10 AS b2 ON a.asesmed_diagS2=b2.code2
			LEFT JOIN vicore_lib.k_icd10 AS b3 ON a.asesmed_diagS3=b3.code2
			LEFT JOIN vicore_lib.k_icd10 AS b4 ON a.asesmed_diagS4=b4.code2
			LEFT JOIN vicore_lib.k_icd10 AS b5 ON a.asesmed_diagS5=b5.code2
			LEFT JOIN vicore_lib.k_icd10 AS b6 ON a.asesmed_diagS6=b6.code2
			WHERE a.noreg=? AND a.kd_bagian=? LIMIT 1;
		`

	result := sr.DB.Raw(query, NoReg, Bagian).Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	return diag, nil
}

func (sr *asesmenRepository) GetDiagnosaRepositoryReportPONEK(NoReg string) (res []asesmen.DiagnosaResponse, err error) {
	var diag []asesmen.DiagnosaResponse

	type Diagnoses struct {
		P    string `json:"p"`
		Desp string `json:"desp"`
		S1   string `json:"s1"`
		Des1 string `json:"des1"`
		S2   string `json:"s2"`
		Des2 string `json:"des2"`
		S3   string `json:"s3"`
		Des3 string `json:"des3"`
		S4   string `json:"s4"`
		Des4 string `json:"des4"`
		S5   string `json:"s5"`
		Des5 string `json:"des5"`
		S6   string `json:"s6"`
		Des6 string `json:"des6"`
	}

	data := Diagnoses{}

	query := `
			SELECT asesmed_diagP AS p, b.description as desp, 
			asesmed_diagS1 AS s1, b1.description as des1, 
			asesmed_diagS2 AS s2, b2.description as des2,
			asesmed_diagS3 AS s3, b3.description as des3,
			asesmed_diagS4 AS s4, b4.description as des4,
			asesmed_diagS5 AS s5, b5.description as des5,
			asesmed_diagS6 AS s6, b6.description as des6
			FROM vicore_rme.dcppt_soap_dokter AS a 
			LEFT JOIN vicore_lib.k_icd10 AS b ON a.asesmed_diagP=b.code2 
			LEFT JOIN vicore_lib.k_icd10 AS b1 ON a.asesmed_diagS1=b1.code2
			LEFT JOIN vicore_lib.k_icd10 AS b2 ON a.asesmed_diagS2=b2.code2
			LEFT JOIN vicore_lib.k_icd10 AS b3 ON a.asesmed_diagS3=b3.code2
			LEFT JOIN vicore_lib.k_icd10 AS b4 ON a.asesmed_diagS4=b4.code2
			LEFT JOIN vicore_lib.k_icd10 AS b5 ON a.asesmed_diagS5=b5.code2
			LEFT JOIN vicore_lib.k_icd10 AS b6 ON a.asesmed_diagS6=b6.code2
			WHERE a.noreg=? AND a.kd_bagian=? LIMIT 1;
		`

	result := sr.DB.Raw(query, NoReg, "PONEK").Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	return diag, nil
}

func (sr *asesmenRepository) OnGetDiagnosaByNoRegANDKdBagian(NoReg string, kdBagian string) (res []asesmen.DiagnosaResponse, err error) {
	var diag []asesmen.DiagnosaResponse

	type Diagnoses struct {
		P    string `json:"p"`
		Desp string `json:"desp"`
		S1   string `json:"s1"`
		Des1 string `json:"des1"`
		S2   string `json:"s2"`
		Des2 string `json:"des2"`
		S3   string `json:"s3"`
		Des3 string `json:"des3"`
		S4   string `json:"s4"`
		Des4 string `json:"des4"`
		S5   string `json:"s5"`
		Des5 string `json:"des5"`
		S6   string `json:"s6"`
		Des6 string `json:"des6"`
	}

	data := Diagnoses{}

	query := `
			SELECT asesmed_diagP AS p, b.description as desp, 
			asesmed_diagS1 AS s1, b1.description as des1, 
			asesmed_diagS2 AS s2, b2.description as des2,
			asesmed_diagS3 AS s3, b3.description as des3,
			asesmed_diagS4 AS s4, b4.description as des4,
			asesmed_diagS5 AS s5, b5.description as des5,
			asesmed_diagS6 AS s6, b6.description as des6
			FROM vicore_rme.dcppt_soap_dokter AS a 
			LEFT JOIN vicore_lib.k_icd10 AS b ON a.asesmed_diagP=b.code2 
			LEFT JOIN vicore_lib.k_icd10 AS b1 ON a.asesmed_diagS1=b1.code2
			LEFT JOIN vicore_lib.k_icd10 AS b2 ON a.asesmed_diagS2=b2.code2
			LEFT JOIN vicore_lib.k_icd10 AS b3 ON a.asesmed_diagS3=b3.code2
			LEFT JOIN vicore_lib.k_icd10 AS b4 ON a.asesmed_diagS4=b4.code2
			LEFT JOIN vicore_lib.k_icd10 AS b5 ON a.asesmed_diagS5=b5.code2
			LEFT JOIN vicore_lib.k_icd10 AS b6 ON a.asesmed_diagS6=b6.code2
			WHERE a.noreg=? AND a.kd_bagian=? LIMIT 1;
		`

	result := sr.DB.Raw(query, NoReg, kdBagian).Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	return diag, nil
}

func (sr *asesmenRepository) OnGetDiagnosaByNoRegANDKdBagianPelayananRANAP(NoReg string, kdBagian string) (res []asesmen.DiagnosaResponse, err error) {
	var diag []asesmen.DiagnosaResponse

	type Diagnoses struct {
		P    string `json:"p"`
		Desp string `json:"desp"`
		S1   string `json:"s1"`
		Des1 string `json:"des1"`
		S2   string `json:"s2"`
		Des2 string `json:"des2"`
		S3   string `json:"s3"`
		Des3 string `json:"des3"`
		S4   string `json:"s4"`
		Des4 string `json:"des4"`
		S5   string `json:"s5"`
		Des5 string `json:"des5"`
		S6   string `json:"s6"`
		Des6 string `json:"des6"`
	}

	data := Diagnoses{}

	query := `
			SELECT asesmed_diagP AS p, b.description as desp, 
			asesmed_diagS1 AS s1, b1.description as des1, 
			asesmed_diagS2 AS s2, b2.description as des2,
			asesmed_diagS3 AS s3, b3.description as des3,
			asesmed_diagS4 AS s4, b4.description as des4,
			asesmed_diagS5 AS s5, b5.description as des5,
			asesmed_diagS6 AS s6, b6.description as des6
			FROM vicore_rme.dcppt_soap_dokter AS a 
			LEFT JOIN vicore_lib.k_icd10 AS b ON a.asesmed_diagP=b.code2 
			LEFT JOIN vicore_lib.k_icd10 AS b1 ON a.asesmed_diagS1=b1.code2
			LEFT JOIN vicore_lib.k_icd10 AS b2 ON a.asesmed_diagS2=b2.code2
			LEFT JOIN vicore_lib.k_icd10 AS b3 ON a.asesmed_diagS3=b3.code2
			LEFT JOIN vicore_lib.k_icd10 AS b4 ON a.asesmed_diagS4=b4.code2
			LEFT JOIN vicore_lib.k_icd10 AS b5 ON a.asesmed_diagS5=b5.code2
			LEFT JOIN vicore_lib.k_icd10 AS b6 ON a.asesmed_diagS6=b6.code2
			WHERE a.noreg=? AND a.kd_bagian=? AND a.pelayanan=?  LIMIT 1;
		`

	result := sr.DB.Raw(query, NoReg, kdBagian, "ranap").Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	if len(diag) == 0 {
		return make([]asesmen.DiagnosaResponse, 0), nil
	}

	return diag, nil
}

func (sr *asesmenRepository) GetDiagnosaRanap(NoReg string) (res []asesmen.DiagnosaResponse, err error) {
	var diag []asesmen.DiagnosaResponse

	type Diagnoses struct {
		P    string `json:"p"`
		Desp string `json:"desp"`
		S1   string `json:"s1"`
		Des1 string `json:"des1"`
		S2   string `json:"s2"`
		Des2 string `json:"des2"`
		S3   string `json:"s3"`
		Des3 string `json:"des3"`
		S4   string `json:"s4"`
		Des4 string `json:"des4"`
		S5   string `json:"s5"`
		Des5 string `json:"des5"`
		S6   string `json:"s6"`
		Des6 string `json:"des6"`
	}

	data := Diagnoses{}

	query := `
			SELECT asesmed_diagP AS p, b.description as desp, 
			asesmed_diagS1 AS s1, b1.description as des1, 
			asesmed_diagS2 AS s2, b2.description as des2,
			asesmed_diagS3 AS s3, b3.description as des3,
			asesmed_diagS4 AS s4, b4.description as des4,
			asesmed_diagS5 AS s5, b5.description as des5,
			asesmed_diagS6 AS s6, b6.description as des6
			FROM vicore_rme.dcppt_soap_dokter AS a 
			LEFT JOIN vicore_lib.k_icd10 AS b ON a.asesmed_diagP=b.code2 
			LEFT JOIN vicore_lib.k_icd10 AS b1 ON a.asesmed_diagS1=b1.code2
			LEFT JOIN vicore_lib.k_icd10 AS b2 ON a.asesmed_diagS2=b2.code2
			LEFT JOIN vicore_lib.k_icd10 AS b3 ON a.asesmed_diagS3=b3.code2
			LEFT JOIN vicore_lib.k_icd10 AS b4 ON a.asesmed_diagS4=b4.code2
			LEFT JOIN vicore_lib.k_icd10 AS b5 ON a.asesmed_diagS5=b5.code2
			LEFT JOIN vicore_lib.k_icd10 AS b6 ON a.asesmed_diagS6=b6.code2
			WHERE a.noreg=?  LIMIT 1;
		`

	result := sr.DB.Raw(query, NoReg).Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, asesmen.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	if len(diag) == 0 {
		return make([]asesmen.DiagnosaResponse, 0), nil
	}

	return diag, nil
}

func (ig *asesmenRepository) OnGetDataPasienKeluarObatRepository(NoReg string) (res []asesmen.DApotikKeluarObat, err error) {
	query := "SELECT jam, statuspasien as status_pasien, id, noreg , nokeluar as no_keluar  from his.dapotikkeluarobat  where noreg=? AND ket='UGD' ORDER BY jam desc"
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetDataRegisterPasienRepository(NoReg string) (res asesmen.DataPasienResgister, err error) {

	query := "SELECT id, noreg  from rekam.dregister where noreg=? limit 1"

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

// /
func (ig *asesmenRepository) OnGetAsesmenDokterRepository(NoReg string) (res asesmen.AsesmenDokterPengantarRawatInap, err error) {
	query := "SELECT dp.insert_dttm as waktu,  keterangan_person as person, insert_user_id as user_id, kd_bagian, kp.namadokter as nama_dokter, asesmed_keluh_utama as keluhan_utama, kl.bagian as nama_bagian, asesmed_konsul_ke as konsul_ke from vicore_rme.dcppt_soap_dokter as dp INNER JOIN his.ktaripdokter as kp ON dp.insert_user_id = kp.iddokter INNER JOIN vicore_lib.kpelayanan as kl on dp.kd_bagian = kl.kd_bag WHERE dp.noreg=? limit 1"

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetDPemFisikPengatarRawatInapRepository(NoReg string, KDBagian string) (res asesmen.DPemfisikModel, err error) {
	query := "SELECT  noreg, kd_bagian, gcs_e, kesadaran, gcs_m, gcs_v from vicore_rme.dpem_fisik WHERE noreg=? AND kd_bagian=? limit 1;"

	result := ig.DB.Raw(query, NoReg, KDBagian).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetVitalSignRepository(NoReg string, KDBagian string) (res asesmen.DVitalSign, err error) {
	query := "SELECT kd_bagian, noreg, suhu, pernafasan, nadi, td, tb, bb, spo2  from vicore_rme.dvital_sign WHERE noreg=? AND kd_bagian=?;"

	result := ig.DB.Raw(query, NoReg, KDBagian).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetInstruksiObatRepository(NoReg string, KetBagians string) (res []asesmen.DApotikKeluarObat, err error) {
	var ketBagian string

	switch KetBagians {
	case "PERI":
		ketBagian = "Perinatologi"
	case "POL009":
		ketBagian = "Haemodialisa"
	case "POL027":
		ketBagian = "Maria"
	case "POL019":
		ketBagian = "Poliklinik Anak"
	case "POL011":
		ketBagian = "Poliklinik Bedah"
	case "OK0001":
		ketBagian = "Poliklinik Bedah"
	case "Klinik Bedah":
		ketBagian = "Poliklinik Bedah"
	case "PONEK":
		ketBagian = "UGD"
	case "IGD001":
		ketBagian = "UGD"
	case "POL020":
		ketBagian = "Poliklinik Internist"
	case "POL006":
		ketBagian = "Poliklinik Syaraf"
	case "POL024":
		ketBagian = "Poliklinik Mata"
	case "LUKA":
		ketBagian = "Lukas"
	case "MEIN":
		ketBagian = "Meinalda"
	default:
		ketBagian = ""
	}

	query := "SELECT ket, jam, statuspasien as status_pasien, id, noreg , nokeluar as no_keluar  from his.dapotikkeluarobat  where noreg=? AND ket=? ORDER BY jam desc"
	result := ig.DB.Raw(query, NoReg, ketBagian).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetInstruksiObatRepo(NoReg string) (res []asesmen.DApotikKeluarObat, err error) {

	query := "SELECT ket, jam, statuspasien as status_pasien, id, noreg , nokeluar as no_keluar  from his.dapotikkeluarobat  where noreg=? ORDER BY tglkeluar ASC LIMIT 1"
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetPemeriksaanFisikAsesmenDokterIGD(NoReg string) (res dto.PemeriksanFisikAwalMedis, err error) {
	query := "SELECT  noreg, ket_person as person, kesadaran, gcs_e as e, gcs_v as v, gcs_m as m, kd_bagian as bagian , dada,  tht, telinga, hidung, tenggorokan, mulut, leher, jantung, paru, perut, kepala, mata, hati, abdomen as kelenjar_getah_bening, limpa, ginjal, alat_kelamin, anggota_gerak, refleks, otot as kekuatan_otot, kulit, rt_vt as rt_vt from vicore_rme.dpem_fisik WHERE noreg=? AND kd_bagian='IGD001' limit 1;"
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetPemeriksaanFisikAsesmenDokterRanap(NoReg string, Bagian string) (res dto.PemeriksanFisikAwalMedis, err error) {
	query := "SELECT  noreg, ket_person as person, kesadaran, gcs_e as e, gcs_v as v, gcs_m as m, kd_bagian as bagian , dada,  tht, telinga, hidung, tenggorokan, mulut, leher, jantung, paru, perut, kepala, mata, hati, abdomen as kelenjar_getah_bening, limpa, ginjal, alat_kelamin, anggota_gerak, refleks, otot as kekuatan_otot, kulit, rt_vt as rt_vt from vicore_rme.dpem_fisik WHERE noreg=? AND kd_bagian=? limit 1;"
	result := ig.DB.Raw(query, NoReg, Bagian).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *asesmenRepository) OnGetRiwayatPenyakitKeluarga(NoRM string) (res []asesmen.RiwayatPenyakit, err error) {
	query := "SELECT  id, kelompok , alergi  from vicore_rme.dalergi where id=? limit 10;"
	result := ig.DB.Raw(query, NoRM).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil

}

func (ig *asesmenRepository) OnGetOperasiPasienRepository(NoReg string) (res asesmen.Bedah, err error) {
	query := "SELECT * from kamop.dlap_bedah WHERE noreg=? limit 1"
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (ig *asesmenRepository) OnGetLaporanOperasiByNoRMRepository(ID string) (res asesmen.Bedah, err error) {
	query := "select * from kamop.dlap_bedah WHERE id=? limit 1"
	result := ig.DB.Raw(query, ID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (ig *asesmenRepository) OnGetBedah3Repository(NomorBedah string) (res []asesmen.DPenLab3, err error) {
	query := "SELECT nomor, nama,ket FROM kamop.dlap_bedah3 WHERE nomor=?"
	result := ig.DB.Raw(query, NomorBedah).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak itemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (ig *asesmenRepository) OnGetDiagnosaRepository(NomorBedah string) (res []asesmen.Dpenlab2, err error) {
	query := "SELECT jenis_diag as jenis, kode, diagnosa, ket, keterangan FROM kamop.dlap_bedah2 WHERE nomor=?"
	result := ig.DB.Raw(query, NomorBedah).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak itemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}
