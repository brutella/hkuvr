package hkuvr

import (
	"github.com/brutella/hc/accessory"
)

func InfoForAccessoryName(name string) accessory.Info {
	info := accessory.Info{
		Name:         name,
		Manufacturer: "TA",
		Model:        "UVR1611",
	}

	return info
}
