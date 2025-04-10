package mapper

import (
	"strings"
	"vicore_hrd/modules/antrean"
	"vicore_hrd/modules/antrean/dto"
	"vicore_hrd/modules/antrean/entity"

	"github.com/sirupsen/logrus"
)

type AntreanMapper struct {
	Logging *logrus.Logger
}

func NewAntreanMapperImple(logging *logrus.Logger) entity.AntreanMapper {
	return &AntreanMapper{
		Logging: logging,
	}
}

func (mm *AntreanMapper) ToMappingDataDResiterPasien(data []antrean.DRegisterPasien) (res []dto.DataResponseRegisterPasien) {
	for _, V := range data {
		res = append(res, dto.DataResponseRegisterPasien{
			Tanggal:    V.Tanggal,
			Id:         V.Id,
			Noreg:      V.Noreg,
			Nama:       V.Nama,
			Kunjungan:  V.Kunjungan,
			Keterangan: V.Keterangan,
			Pelayaan:   strings.ToUpper(V.Pelayaan),
			Bagian:     V.Bagian,
		})
	}
	return res
}

func (mm *AntreanMapper) ToMappingPasienRANAP(data []antrean.PasienRANAP) (res []dto.DataResponseRegisterPasien) {
	for _, V := range data {

		res = append(res, dto.DataResponseRegisterPasien{
			Tanggal:    V.Tanggal,
			Id:         "-",
			Noreg:      V.Noreg,
			Nama:       "",
			Kunjungan:  "-",
			Keterangan: "RANAP",
		})
	}
	return res
}
