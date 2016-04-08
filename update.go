package hkuvr

import (
	"fmt"
	"github.com/brutella/uvr"
)

func UpdateOutletValue(svc *Outlet, i uint8, c *uvr.Client) error {
	outlet := uvr.NewOutlet(i)
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
			svc.OutletInUse.SetValue(true)
		} else {
			svc.OutletInUse.SetValue(false)
			return fmt.Errorf("Outlet state of unknown type %v", value)
		}
	} else {
		return err
	}

	return nil
}

func UpdateOutletName(svc *Outlet, i uint8, c *uvr.Client) error {
	outlet := uvr.NewOutlet(i)
	if value, err := c.Read(outlet.Description); err == nil {
		if str, ok := value.(string); ok == true && len(str) > 0 {
			svc.Name.SetValue(str)
		} else {
			return fmt.Errorf("Outlet description of unknown type %v", value)
		}
	} else {
		return err
	}

	return nil
}

func UpdateTemperatureSensorName(svc *TemperatureSensor, i uint8, c *uvr.Client) error {
	inlet := uvr.NewInlet(i)
	if value, err := c.Read(inlet.Description); err == nil {
		if str, ok := value.(string); ok == true && len(str) > 0 {
			svc.Name.SetValue(str)
		} else {
			return fmt.Errorf("Outlet description of unknown type %v", value)
		}
	} else {
		return err
	}

	return nil
}

func UpdateTemperatureSensorValue(svc *TemperatureSensor, i uint8, c *uvr.Client) error {
	inlet := uvr.NewInlet(i)
	if value, err := c.Read(inlet.Value); err == nil {
		if float, ok := value.(float32); ok == true {
			svc.CurrentTemperature.SetValue(float64(float))
		} else {
			return fmt.Errorf("Inlet state of unknown type %v", value)
		}
	} else {
		return err
	}

	return nil
}
