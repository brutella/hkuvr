package hkuvr

import (
	"github.com/brutella/hc/accessory"
)

type UVR1611 struct {
	*accessory.Accessory
}

func NewUVR1611() *UVR1611 {
	acc := UVR1611{}

	info := accessory.Info{
		Name:         "UVR",
		Model:        "1611",
		Manufacturer: "TA",
	}
	acc.Accessory = accessory.New(info, accessory.TypeBridge)

	return &acc
}
