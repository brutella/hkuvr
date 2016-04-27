package hkuvr

import (
	"github.com/brutella/hc/service"
	"github.com/brutella/uvr"

	"fmt"
)

type Object interface {
	Update(*uvr.Client) error

	Service() *service.Service

	SubIndex() uint8
}

func NewObject(val interface{}, name string, idx uint8) (Object, error) {
	if str, ok := val.(string); ok == true {
		if v, err := stringToBool(str); err == nil {
			svc := NewOutlet(idx)
			svc.On.SetValue(v)
			svc.Name.SetValue(name)

			return svc, nil
		}
	} else if v, ok := val.(float32); ok == true {
		svc := NewTemperatureSensor(idx)
		svc.CurrentTemperature.SetValue(float64(v))
		svc.Name.SetValue(name)

		return svc, nil
	}

	return nil, fmt.Errorf("Cannot create object: Value: %v, name: %s, index: %d", val, name, idx)
}
