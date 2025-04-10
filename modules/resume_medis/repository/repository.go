package repository

import (
	"errors"
	"fmt"
	resumemedis "vicore_hrd/modules/resume_medis"
	"vicore_hrd/modules/resume_medis/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type resumeMedisRepository struct {
	DB      *gorm.DB
	Logging *logrus.Logger
}

func NewResumeMedisRespository(db *gorm.DB, logging *logrus.Logger) entity.ResumeMedisRepository {
	return &resumeMedisRepository{
		DB:      db,
		Logging: logging,
	}
}

func (ig *resumeMedisRepository) GetDataTanggalMasukPasienRepository(NoReg string) (res resumemedis.DRegisterPasien, err error) {
	query := "SELECT tanggal as tgl_masuk, jam as jam_masuk, id, noreg, nama  from rekam.dregister where noreg=? limit 1"
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetDataProfilePasienRespository(ID string) (res resumemedis.DProfilePasien, err error) {
	query := "SELECT jeniskelamin, id, firstname, tgllahir, alamat FROM his.dprofilpasien where id =? LIMIT 1"

	result := ig.DB.Raw(query, ID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (ig *resumeMedisRepository) OnGetPasienKeluarRepository(NoReg string) (res resumemedis.DBangsalRep, err error) {
	query := "SELECT namadokter as nama_dokter, bagian, tanggal, jam, noreg, nama, keterangan from his.dbangsal_rep where noreg=? limit 1"
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetAsesmenRanapRepository(NoReg string) (res resumemedis.AsesmenRanap, err error) {
	query := "select kd_bagian, noreg, pelayanan from vicore_rme.dcppt_soap_dokter where noreg=? AND pelayanan=? limit 1"
	result := ig.DB.Raw(query, NoReg, "ranap").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetDBilingRekapRepository(ID string) (res resumemedis.Drekap, err error) {
	query := "SELECT tglkeluar, id, noreg, nama, umur , asal, kelas , pelayanan, alamat from biling.drekap where id=?   ORDER BY tglkeluar DESC limit 1 "
	result := ig.DB.Raw(query, ID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) GetDiagnosaDariDRekamRepository(NoReg string) (res resumemedis.DRekamSatu, err error) {
	query := "select noreg, icd_in, icd_ou from rekam.drm_1  where noreg=? limit 1 "
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetKDiagnosaRepository(Code string) (res resumemedis.DiagnosaICD, err error) {
	query := "select s1, code, code2, description from vicore_lib.k_icd10 where code =? limit 1 "
	result := ig.DB.Raw(query, Code).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetDBilingRekapByNoregRepository(NoReg string) (res resumemedis.Drekap, err error) {
	query := "SELECT tglkeluar, id, noreg, nama, umur , asal, kelas , pelayanan, alamat from biling.drekap where noreg=?   ORDER BY tglkeluar DESC limit 1 "
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetDataPasienKeluarObatRepository(NoReg string) (res resumemedis.DApotikKeluarObat, err error) {
	query := "SELECT jam, statuspasien as status_pasien, id, noreg , nokeluar as no_keluar  from his.dapotikkeluarobat  where noreg=? ORDER BY jam desc  limit 1"
	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetDataKeluarObat1Repository(NoKeluar string) (res []resumemedis.DataKeluarObat1, err error) {

	query := "SELECT tglkeluar as tgl_keluar, nokeluar as no_keluar, noambil as no_ambil, namaobat as nama_obat, jumlah from his.dapotikkeluarobat1 where nokeluar=? ORDER BY tglkeluar DESC"
	result := ig.DB.Raw(query, NoKeluar).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (lu *resumeMedisRepository) GetPenlabTabelLamaRepository(noReg string) (res []resumemedis.DPenLab, err error) {
	query := `SELECT DISTINCT(jaminput), noreg, id FROM his.dpenmlab WHERE noreg=?`

	result := lu.DB.Raw(query, noReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (lu *resumeMedisRepository) GetPenLabKelompokTabelLamaRepository(jamInput string, noReg string) (res []resumemedis.DPenmedPemeriksaan, err error) {

	query := `SELECT DISTINCT(kelompok) AS nama_kelompok FROM his.dpenmlab WHERE jaminput=? AND noreg=?;`

	result := lu.DB.Raw(query, jamInput, noReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil

}

func (lu *resumeMedisRepository) GetPenLabPemeriksaanTabelLamaRepository(jamInput string, noReg string, kelompok string) (res []resumemedis.DPemeriksaanLabor, err error) {

	query := `SELECT dl.pemeriksaan, dl.normal, dl.satuan, dl.hasil FROM his.dpenmlab AS dl LEFT JOIN his.kpemeriksaanlaborat AS kp ON dl.pemeriksaan=kp.pemeriksaan WHERE dl.noreg=? AND dl.jaminput=? AND dl.kelompok=? ORDER BY kp.urut`

	result := lu.DB.Raw(query, noReg, jamInput, kelompok).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil

}

// Get Hasil Radiologi Old Table
func (lu *resumeMedisRepository) GetHasilRadiologiRepositoryV2(noReg string) (res []resumemedis.DHasilRadiologiV2, err error) {
	query := `SELECT DISTINCT(jaminput), bagian FROM his.dpenmlablain WHERE noreg=? AND (bagian='Radiologi' OR bagian='Citiscan' OR bagian='Usg' OR bagian='Ekg')`

	result := lu.DB.Raw(query, noReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

// GET DETAIL HASIL RADIOLOGI
func (lu *resumeMedisRepository) GetDetailHasilRadiologiOldDB(bagian string, noReg string, jamInput string) (res []resumemedis.DHasilRadiologiOldDB, err error) {
	query := `SELECT pemeriksaan,uraian,hasil FROM his.dpenmlablain WHERE bagian=? AND noreg=? AND jaminput=?`

	result := lu.DB.Raw(query, bagian, noReg, jamInput).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (lu *resumeMedisRepository) GetHasilFisioterapiRepositoryV2(noReg string) (res []resumemedis.DHasilRadiologiV2, err error) {
	query := `SELECT DISTINCT(jaminput), bagian FROM his.dpenmlablain WHERE noreg=? AND bagian='Fisiotherapy'`
	result := lu.DB.Raw(query, noReg).Scan(&res)
	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (lu *resumeMedisRepository) GetHasilGiziRepositoryV2(noReg string) (res []resumemedis.DHasilRadiologiV2, err error) {
	query := `SELECT DISTINCT(jaminput), bagian FROM his.dpenmlablain WHERE noreg=? AND bagian='Gizi'`

	result := lu.DB.Raw(query, noReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (sr *resumeMedisRepository) GetPemeriksaaanFisikRespository(noreg string) (res resumemedis.PemFisik, err error) {
	query := `SELECT kategori, ket_person, noreg, tb, td, bb,  nadi, suhu, spo2, pernafasan FROM vicore_rme.dvital_sign where kd_bagian = "IGD001" AND noreg=?  ORDER BY insert_dttm DESC LIMIT 1`

	result := sr.DB.Raw(query, noreg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *resumeMedisRepository) GetPemeriksaaanFisikRanapRespository(noreg string) (res resumemedis.PemFisik, err error) {
	query := `SELECT kategori, ket_person, noreg, tb, td, bb,  nadi, suhu, spo2, pernafasan FROM vicore_rme.dvital_sign where noreg=?  ORDER BY insert_dttm DESC LIMIT 1`

	result := sr.DB.Raw(query, noreg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data gagal diupdate", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (sr *resumeMedisRepository) GetDiagnosaRepositoryBangsal(noreg string, kdBagian string) (res []resumemedis.DiagnosaResponse, err error) {
	var diag []resumemedis.DiagnosaResponse

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

	result := sr.DB.Raw(query, noreg, kdBagian).Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	sr.Logging.Info("GET DATA DIAGNOSA")
	sr.Logging.Info(diag)

	return diag, nil
}

func (sr *resumeMedisRepository) GetDiagnosa(noreg string) (res []resumemedis.DiagnosaResponse, err error) {
	var diag []resumemedis.DiagnosaResponse

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
			WHERE a.noreg=? LIMIT 1;
		`

	result := sr.DB.Raw(query, noreg).Scan(&data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.P) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	return diag, nil
}
func (sr *resumeMedisRepository) GetDiagnosaDokterByNoreg(NoReg string) (res []resumemedis.DiagnosaResponse, err error) {
	var diag []resumemedis.DiagnosaResponse

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
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.P,
			Description: data.Desp,
			Type:        "primer",
			Table:       "P",
		})
	}

	if len(data.S1) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S1,
			Description: data.Des1,
			Type:        "sekunder",
			Table:       "S1",
		})
	}

	if len(data.S2) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S2,
			Description: data.Des2,
			Type:        "sekunder",
			Table:       "S2",
		})
	}

	if len(data.S3) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S3,
			Description: data.Des3,
			Type:        "sekunder",
			Table:       "S3",
		})
	}
	if len(data.S4) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S4,
			Description: data.Des4,
			Type:        "sekunder",
			Table:       "S4",
		})
	}

	if len(data.Des5) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S5,
			Description: data.Des5,
			Type:        "sekunder",
			Table:       "S5",
		})
	}

	if len(data.S6) > 1 {
		diag = append(diag, resumemedis.DiagnosaResponse{
			Diagnosa:    data.S6,
			Description: data.Des6,
			Type:        "sekunder",
			Table:       "S6",
		})
	}

	sr.Logging.Info("GET DATA DIAGNOSA")
	sr.Logging.Info(diag)

	return diag, nil
}

func (sr *resumeMedisRepository) OnGetTindakanICD9RepositoryBangsalDokter(noReg string, kdBagian string) (res []resumemedis.TindakanResponse, err error) {
	var listTindakan []resumemedis.TindakanResponse

	type Tindakan struct {
		Noreg        string
		KdBagian     string
		AsesmenPros1 string `gorm:"column:asesmed_pros1" json:"asesmen_pros1"`
		AsesmenPros2 string `gorm:"column:asesmed_pros2" json:"asesmen_pros2"`
		Des1         string `gorm:"column:des1" json:"des1"`
		Des2         string `gorm:"column:des2" json:"des2"`
	}

	data := Tindakan{}

	query := `
			SELECT noreg, kd_bagian, asesmed_pros1,asesmed_pros2 , lib.Description AS des1 , lib1.Description AS des2 FROM vicore_rme.dcppt_soap_dokter AS dd LEFT JOIN vicore_lib.k_icd9 AS lib ON lib.Code2 = dd.asesmed_pros1
			LEFT JOIN vicore_lib.k_icd9 AS lib1 ON lib1.Code2 = dd.asesmed_pros2
			WHERE noreg=? AND kd_bagian=? LIMIT 1;
			`

	result := sr.DB.Raw(query, noReg, kdBagian).Scan(&data)

	sr.Logging.Info(data)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	if len(data.AsesmenPros1) > 1 {
		listTindakan = append(listTindakan, resumemedis.TindakanResponse{
			Description: data.Des1,
			Code2:       data.AsesmenPros1,
		})
	}

	if len(data.AsesmenPros2) > 1 {
		listTindakan = append(listTindakan, resumemedis.TindakanResponse{
			Description: data.Des2,
			Code2:       data.AsesmenPros2,
		})
	}

	return listTindakan, nil
}

func (ig *resumeMedisRepository) OnGetRiwayatPenyakitRepository(NoReg string, KDBagian string) (res resumemedis.RiwayatPenyakit, err error) {
	query := `SELECT noreg, kd_bagian, asesmed_rwyt_skrg as riwayat_sekarang, asesmed_rwyt_dahulu as riwayat_dahulu, asesmed_rwyt_obat as riwayat_obat from vicore_rme.dcppt_soap_dokter where noreg=? AND kd_bagian="IGD001" LIMIT 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) CariPasienPulangRepository(Cari string) (res []resumemedis.CariDataPasienPulang, err error) {
	query := `SELECT dm.tglproses, dm.id as norm, dm.pelayanan, dm.noreg, dm.nama FROM rekam.drm_1 as dm 
				INNER JOIN rekam.dregister as dr on dr.noreg = dm.noreg WHERE dm.id=? ORDER BY dm.tglproses DESC LIMIT 20`

	result := ig.DB.Raw(query, Cari).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) CariPasienPulangRanapRepository(Cari string) (res []resumemedis.CariDataPasienPulang, err error) {
	query := `SELECT dm.tglproses, dm.id as norm, dm.pelayanan, dm.noreg, dm.nama FROM rekam.drm_1 as dm 
				INNER JOIN rekam.dregister as dr on dr.noreg = dm.noreg WHERE dm.id=? AND dm.pelayanan=? ORDER BY dm.tglproses DESC LIMIT 20`

	result := ig.DB.Raw(query, Cari, "Rawat Inap").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetJenisKelamin(NOID string) (res resumemedis.DataJenis, err error) {
	query := `select  jeniskelamin as jenis_kelamin from his.dprofilpasien where id=?  limit 1`

	result := ig.DB.Raw(query, NOID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetProfilePasien(ID string) (res resumemedis.DataJenis, err error) {
	query := `select id   jeniskelamin as jenis_kelamin, firstname  as nama from his.dprofilpasien where id=?  limit 1`

	result := ig.DB.Raw(query, ID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetDataDRekamRepository(NoReg string) (res resumemedis.DataDRekamMedis, err error) {
	query := `SELECT pelayanan, bagian, tahun, bulan, tglproses, jamproses, id, noreg, nama, alamat, sex , iddokter, dokter, tglmasuk, tglkeluar, icd_in, icd_ou, diagnosa_in, diagnosa_ou, golongan, kunjungan, kea_keluar from rekam.drm_1 where noreg=? limit 1`

	result := ig.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

func (ig *resumeMedisRepository) OnGetDataProfilePasienRepository(ID string) (res resumemedis.DataProfilePasien, err error) {

	query := `select nik, nokapst, id, firstname, lastname, agama, jeniskelamin, tempatlahir, tgllahir,
       alamat, suku , kelurahan, kecamatan, kotamadya, kabupaten, propinsi, negara, telp, hp, namaayah from his.dprofilpasien  where id=? limit 1`

	result := ig.DB.Raw(query, ID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, data gagal didapat", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}
