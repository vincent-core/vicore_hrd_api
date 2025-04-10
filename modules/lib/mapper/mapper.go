package mapper

import (
	"vicore_hrd/modules/lib"
	"vicore_hrd/modules/lib/dto"
	"vicore_hrd/modules/lib/entity"
)

type LibMapper struct {
}

func NewLibMapperImple() entity.LibMapper {
	return &LibMapper{}
}

func (im *LibMapper) ToMappingPelayanan(data []lib.KPelayanan) (res []dto.ResponsePelayanan) {

	if len(data) == 0 {
		return make([]dto.ResponsePelayanan, 0)
	}

	if len(data) > 0 {
		for _, V := range data {
			res = append(res, dto.ResponsePelayanan{
				KDBagian: V.KdBag,
				Bagian:   V.Bagian,
			})
		}

		return res
	}

	return res
}
