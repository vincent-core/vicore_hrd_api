package repository

import (
	"errors"
	"fmt"
	"vicore_hrd/modules/antrean"
)

func (lu *antreanRepository) OnGetDataRegisterPasienByID(ID string) (res []antrean.DRegisterPasien, err error) {
	query := "SELECT tanggal,  jam, id, noreg, nama, kunjungan, keterangan from rekam.dregister where id=? order by tanggal desc limit 15"
	result := lu.DB.Raw(query, ID).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (lu *antreanRepository) OnGetDRekamPasienRANAPRepo(ID string) (res []antrean.PasienRANAP, err error) {
	query := "SELECT rekam.drm_1.noreg, CONCAT(rekam.dregister.tanggal,' ',rekam.dregister.jam) AS tanggal from rekam.drm_1  LEFT JOIN rekam.dregister ON rekam.drm_1.noreg=rekam.dregister.noreg  WHERE rekam.drm_1.id=? AND rekam.drm_1.pelayanan=? ORDER BY rekam.drm_1.noreg DESC"
	result := lu.DB.Raw(query, ID, "Rawat Inap").Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}
	return res, nil
}

func (lu *antreanRepository) OnGetSingleRegisterDataPasienRepository(NoReg string) (res antrean.DRegisterPasien, err error) {
	query := "SELECT tanggal,  jam, id, noreg, nama, kunjungan, keterangan from rekam.dregister where noreg=? LIMIT 1"

	result := lu.DB.Raw(query, NoReg).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak ditemukan", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}
