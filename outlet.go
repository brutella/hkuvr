package hkuvr

import (
	"fmt"
	"github.com/brutella/uvr"
)

//
// type Entity struct {
//     Accessory interface{}
//     SubIndex  uint8
//     Type      uint8
// }
//
// func NewBoolEntity(name string, subIndex uint8, value bool, valueType uint8) *Entity {
//     info := InfoForAccessoryName(name)
//     outlet := accessory.NewOutlet(info)
//     outlet.SetOn(value)
//
//     // On state is readonly because we don't support write of CAN bus yet
//     outlet.Outlet.On.Permissions = characteristic.PermsRead()
//
//     return &Entity{
//         Accessory: outlet,
//         SubIndex:  subIndex,
//         Type:      valueType,
//     }
// }
//
// func NewFloatEntity(name string, subIndex uint8, value float64, valueType uint8) *Entity {
//     info := InfoForAccessoryName(name)
//     sensor := accessory.NewTemperatureSensor(info, 0, 0, 9999, 0.1)
//     sensor.SetTemperature(value)
//
//     return &Entity{
//         Accessory: sensor,
//         SubIndex:  subIndex,
//         Type:      valueType,
//     }
// }
//
// func HomeKitAccessoryForEntity(e *Entity) *accessory.Accessory {
//     if outlet, ok := e.Accessory.(*accessory.Outlet); ok == true {
//         return outlet.Accessory
//     } else if thermometer, ok := e.Accessory.(*accessory.Thermometer); ok == true {
//         return thermometer.Accessory
//     }
//
//     return nil
// }
//
func StringToBool(str string) (bool, error) {
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

//
// // func NewOutletWithName(name string) *accessory.Outlet {
// //     info := InfoForAccessoryName(name)
// //     outlet := accessory.NewOutlet(info)
// //
// //     // On state is readonly because we don't support write of CAN bus yet
// //     outlet.Outlet.On.Permissions = characteristic.PermsRead()
// //
// //     return outlet
// // }
// //
// // func NewInletWithName(name string) *accessory.Thermometer {
// //     info := InfoForAccessoryName(name)
// //     return accessory.NewTemperatureSensor(info, 0, 0, 9999, 0.1)
// // }
