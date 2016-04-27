package hkuvr

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
	"github.com/brutella/uvr"

	"fmt"
)

type TemperatureSensor struct {
	*service.TemperatureSensor

	Name     *characteristic.Name
	subIndex uint8
}

func NewTemperatureSensor(subIndex uint8) *TemperatureSensor {
	svc := TemperatureSensor{}
	svc.subIndex = subIndex
	svc.TemperatureSensor = service.NewTemperatureSensor()
	svc.TemperatureSensor.CurrentTemperature.SetMinValue(-100)
	svc.TemperatureSensor.CurrentTemperature.SetMaxValue(300)

	svc.Name = characteristic.NewName()
	svc.Name.SetValue("Unbenannt")
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}

func (svc *TemperatureSensor) SubIndex() uint8 {
	return svc.subIndex
}

func (svc *TemperatureSensor) Service() *service.Service {
	return svc.TemperatureSensor.Service
}

func (svc *TemperatureSensor) Update(c *uvr.Client) error {
	inlet := uvr.NewInlet(svc.subIndex)
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
