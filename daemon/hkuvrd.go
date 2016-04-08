package main

import (
	"github.com/brutella/can"
	"github.com/brutella/hkuvr"
	"github.com/brutella/log"
	"github.com/brutella/uvr"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/service"

	"fmt"
	dlog "log"
	"os"
	"time"
)

const (
	MaxOutletsCount uint8 = 13
	MaxInletsCount  uint8 = 16
)

var clientID uint8 = 0x10
var serverID uint8 = 0x1
var client *uvr.Client
var bus *can.Bus

// Reference to all in-/outlets
var entities []*hkuvr.Entity

// Reference to UVR1611 bridge
var bridge *accessory.Accessory = hkuvr.NewUVR1611().Accessory
var outlets []*hkuvr.Outlet
var temperatureSensors []*hkuvr.TemperatureSensor

var config = hap.Config{Pin: "00102003"}
var transport hap.Transport

func main() {
	log.Info = false
	log.Verbose = false
	dlog.SetFlags(dlog.LstdFlags | dlog.Lshortfile)

	var err error
	if bus, err = can.NewBusForInterfaceWithName("can0"); err != nil {
		log.Fatal(err)
	}

	// 1. Connect to UVR
	// 2. Read values of all in-/outlets from CAN bus to determine type
	//    The value of an inlet might be
	//    - float (temperature); e.g. 32.5
	//    - string (true/false);, e.g. EIN
	// 3. Read descriptions from CAN bus to update service names
	// 4. Setup IP transport and publish accessories
	// 5. Wait n seconds
	// 6. Read values from CAN bus
	// 7. Go to 5
	go func() {

		// 1.
		if err := connect(); err != nil {
			log.Fatal(err)
		}

		// 2.
		entities = setupUVR(bridge)

		// 3.
		updateEntityNames(entities)

		// 4.
		if t, err := hap.NewIPTransport(config, bridge); err != nil {
			log.Fatal(err)
		} else {
			transport = t
		}

		go transport.Start()

		for {
			// Wait before updating values because `readEntities()` already updates the values
			// 5.
			<-time.After(time.Second * 10)

			// 6.
			updateEntityValues(entities)
		}
	}()

	hap.OnTermination(func() {

		if transport != nil {
			transport.Stop()
		}

		if client != nil {
			client.Disconnect(serverID)
		}

		bus.Disconnect()
		os.Exit(1)
	})

	bus.ConnectAndPublish()
}

// Reads the values of all in-/outlets to determine the accessory type (Outlet or Thermometer)
func setupUVR(acc *accessory.Accessory) []*hkuvr.Entity {
	var entities = []*hkuvr.Entity{}

	for i := uint8(1); i < MaxOutletsCount; i++ {
		out := uvr.NewOutlet(i)

		if value, err := client.Read(out.State); err == nil {
			if svc := serviceForValue(value, fmt.Sprintf("Ausgang %d", i)); svc != nil {
				acc.AddService(svc.(*service.Service))
				entities = append(entities, hkuvr.NewEntity(svc, i))
			}
		} else {
			dlog.Fatal(err, i)
		}
	}

	for i := uint8(1); i < MaxInletsCount; i++ {
		in := uvr.NewInlet(i)

		if value, err := client.Read(in.Value); err == nil {
			if svc := serviceForValue(value, fmt.Sprintf("Eingang %d", i)); svc != nil {
				acc.AddService(svc.(*service.Service))
				entities = append(entities, hkuvr.NewEntity(svc, i))
			}
		} else {
			dlog.Fatal(err, i)
		}
	}

	return entities
}

func serviceForValue(value interface{}, name string) interface{} {
	if str, ok := value.(string); ok == true {
		if v, err := hkuvr.StringToBool(str); err == nil {
			svc := hkuvr.NewOutlet()
			svc.On.SetValue(v)
			svc.Name.SetValue(name)

			return svc
		}
	} else if v, ok := value.(float32); ok == true {
		svc := hkuvr.NewTemperatureSensor()
		svc.CurrentTemperature.SetValue(float64(v))
		svc.Name.SetValue(name)

		return svc
	}

	return nil
}

// Updates the name of all entities
func updateEntityNames(entities []*hkuvr.Entity) {
	for _, e := range entities {
		if outlet := e.Outlet; outlet != nil {
			if err := hkuvr.UpdateOutletName(outlet, e.SubIndex, client); err != nil {
				log.Fatal(err)
			}
		}

		if err := hkuvr.UpdateTemperatureSensorName(e.TemperatureSensor, e.SubIndex, client); err != nil {
			log.Fatal(err)
		}
	}
}

// Updates the value of all entities
func updateEntityValues(entities []*hkuvr.Entity) {
	for _, e := range entities {
		if outlet := e.Outlet; outlet != nil {
			if err := hkuvr.UpdateOutletValue(outlet, e.SubIndex, client); err != nil {
				log.Fatal(err)
			}
		}

		if err := hkuvr.UpdateTemperatureSensorValue(e.TemperatureSensor, e.SubIndex, client); err != nil {
			log.Fatal(err)
		}
	}
}

// Creates a new client and connects to the CAN bus
func connect() error {
	client = uvr.NewClient(clientID, bus)
	return client.Connect(serverID)
}
