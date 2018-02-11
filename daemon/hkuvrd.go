package main

import (
	"github.com/brutella/can"
	"github.com/brutella/hkuvr"
	"github.com/brutella/uvr"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"

	"log"
	"os"
	"strings"
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
var transport hc.Transport

func main() {
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
	// 4. Setup IP transport and publish bridge
	// 5. Wait n seconds
	// 6. Read values from CAN bus
	// 7. Go to 5
	go func() {
		// 1.
		if err := connect(); err != nil {
			log.Fatal(err)
		}

		var bridge *accessory.Accessory = hkuvr.NewUVR1611().Accessory

		// 2.
		var objects = setupUVR(bridge)

		// 4.
		if t, err := hc.NewIPTransport(hc.Config{}, bridge); err != nil {
			log.Fatal(err)
		} else {
			transport = t
		}

		go transport.Start()

		for {
			// 6.
			updateObjectValues(objects)

			// 5.
			<-time.After(time.Second * 10)
		}
	}()

	hc.OnTermination(func() {

		if transport != nil {
			<-transport.Stop()
		}

		if client != nil {
			client.Disconnect(serverID)
		}

		bus.Disconnect()
		os.Exit(1)
	})

	bus.ConnectAndPublish()
}

func collectObjects() []hkuvr.Object {
	objects := []hkuvr.Object{}

	var desc interface{}
	var val interface{}
	var err error

	for i := uint8(1); i <= MaxOutletsCount; i++ {
		out := uvr.NewOutlet(i)

		if desc, err = client.Read(out.Description); err != nil {
			log.Fatal(err)
		}

		str := strings.TrimSpace(desc.(string))

		if strings.HasSuffix(str, uvr.DescriptionUnused) {
			log.Println("[INFO] Ignore outlet", i)
			continue
		}

		if val, err = client.Read(out.State); err != nil {
			log.Fatal(err)
		}

		if obj, err := hkuvr.NewObject(val, str, i); err != nil {
			log.Fatal(err)
		} else {
			objects = append(objects, obj)
		}
	}

	for i := uint8(1); i <= MaxInletsCount; i++ {
		in := uvr.NewInlet(i)

		if desc, err = client.Read(in.Description); err != nil {
			log.Fatal(err)
		}

		str := desc.(string)

		if strings.HasSuffix(str, uvr.DescriptionUnused) {
			log.Println("[INFO] Ignore inlet", i)
			continue
		}

		if val, err = client.Read(in.Value); err != nil {
			log.Fatal(err)
		}

		if obj, err := hkuvr.NewObject(val, str, i); err != nil {
			log.Fatal(err)
		} else {
			objects = append(objects, obj)
		}
	}

	return objects
}

// Reads the values of all in-/outlets to determine the accessory type (Outlet or Thermometer)
func setupUVR(acc *accessory.Accessory) []hkuvr.Object {
	var objects = collectObjects()

	for _, obj := range objects {
		acc.AddService(obj.Service())
	}

	return objects
}

// Updates the value of all entities
func updateObjectValues(objects []hkuvr.Object) {
	for _, obj := range objects {
		if err := obj.Update(client); err != nil {
			log.Fatal(err)
		}
	}
}

// Creates a new client and connects to the CAN bus
func connect() error {
	client = uvr.NewClient(clientID, bus)
	return client.Connect(serverID)
}
