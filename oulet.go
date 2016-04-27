package hkuvr

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
	"github.com/brutella/uvr"

	"fmt"
)

type Outlet struct {
	*service.Switch

	Name *characteristic.Name

	subIndex uint8
}

func NewOutlet(idx uint8) *Outlet {
	svc := Outlet{}
	svc.subIndex = idx

	svc.Switch = service.NewSwitch()

	svc.Name = characteristic.NewName()
	svc.Name.SetValue("Unbenannt")
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}

func (svc *Outlet) SubIndex() uint8 {
	return svc.subIndex
}

func (svc *Outlet) Service() *service.Service {
	return svc.Switch.Service
}

func (svc *Outlet) Update(c *uvr.Client) error {
	outlet := uvr.NewOutlet(svc.subIndex)
	if value, err := c.Read(outlet.State); err == nil {
		if str, ok := value.(string); ok == true {
			switch str {
			case uvr.OutletStateOn:
				svc.On.SetValue(true)
			case uvr.OutletStateOff:
				svc.On.SetValue(false)
			default:
				return fmt.Errorf("Outlet state of unknown value %v (%X)", str, str)
			}
		} else {
			return fmt.Errorf("Outlet state of unknown type %v", value)
		}
	} else {
		return err
	}

	return nil
}

func stringToBool(str string) (bool, error) {
	switch str {
	case uvr.OutletStateOn:
		return true, nil
	case uvr.OutletStateOff:
		return false, nil
	default:
		break
	}

	return false, fmt.Errorf("Unknown string value %v (%X)", str, str)
}
