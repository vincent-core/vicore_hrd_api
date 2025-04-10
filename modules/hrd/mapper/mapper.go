package mapper

import (
	"vicore_hrd/modules/hrd/entity"
)

type hrdMapperImple struct {
}

func NewHRDMapperImple() entity.VicoreHRDMapper {
	return &hrdMapperImple{}
}
