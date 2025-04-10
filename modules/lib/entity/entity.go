package entity

import (
	"vicore_hrd/modules/lib"
	"vicore_hrd/modules/lib/dto"
)

type LibMapper interface {
	ToMappingPelayanan(data []lib.KPelayanan) (res []dto.ResponsePelayanan)
}

type LibRepository interface {
	FindAllPelayananRepository() (res []lib.KPelayanan, err error)
	OnGetDataRekamMedis() (res []lib.DRekamMedis, err error)
}

type LibUseCase interface{}
