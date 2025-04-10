package mapper

import (
	"vicore_hrd/app/rest"
	edukasiterintegrasi "vicore_hrd/modules/edukasi_terintegrasi"
	"vicore_hrd/modules/edukasi_terintegrasi/dto"
	"vicore_hrd/modules/edukasi_terintegrasi/entity"
	"vicore_hrd/modules/hrd"
)

type edukasiTerintegrasiMapper struct {
}

func NewEdukasiTerintegrasiMapper() entity.EdukasiTerintegrasiMapper {
	return &edukasiTerintegrasiMapper{}
}

func (mm *edukasiTerintegrasiMapper) ToMappingPemberiInformasi(data []hrd.Kemployee) (res []dto.ResponsePemberiInformasi) {
	if len(data) == 0 {
		return make([]dto.ResponsePemberiInformasi, 0)
	}

	if len(data) > 0 {
		for _, V := range data {
			res = append(res, dto.ResponsePemberiInformasi{
				Nama:         V.Nama,
				JenisKelamin: V.JenisKelamin,
			})
		}

		return res
	}

	return make([]dto.ResponsePemberiInformasi, 0)
}

func (mm *edukasiTerintegrasiMapper) ToMappingEdukasiTerintegrasi(data []edukasiterintegrasi.DedukasiTerintegrasi) (res []dto.ResponseEdukasiTerintegrasi) {
	if len(data) > 0 {
		for _, V := range data {
			res = append(res, dto.ResponseEdukasiTerintegrasi{
				ID:                V.IdEdukasi,
				Tanggal:           rest.FormatTanggalWaktu(V.InsertDttm),
				Informasi:         V.Informasi,
				Metode:            V.Metode,
				PemberiInformasi:  V.PemberiInformasi,
				PenerimaInformasi: V.PenerimaInformasi,
				Evaluasi:          V.Evaluasi,
			})
		}

		return res
	}

	return res
}
