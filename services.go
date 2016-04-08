package hkuvr

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type Entity struct {
	Outlet            *Outlet
	TemperatureSensor *TemperatureSensor
	SubIndex          uint8
}

func NewEntity(service interface{}, subIndex uint8) *Entity {
	if svc, ok := service.(*Outlet); ok == true {
		return &Entity{Outlet: svc, SubIndex: subIndex}
	}

	return &Entity{TemperatureSensor: service.(*TemperatureSensor), SubIndex: subIndex}
}

type Outlet struct {
	*service.Outlet

	Name *characteristic.Name
}

func NewOutlet() *Outlet {
	svc := Outlet{}

	svc.Outlet = service.NewOutlet()

	svc.Name = characteristic.NewName()
	svc.Name.SetValue("Unbenannt")
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}

type TemperatureSensor struct {
	*service.TemperatureSensor

	Name *characteristic.Name
}

func NewTemperatureSensor() *TemperatureSensor {
	svc := TemperatureSensor{}

	svc.TemperatureSensor = service.NewTemperatureSensor()

	svc.Name = characteristic.NewName()
	svc.Name.SetValue("Unbenannt")
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
