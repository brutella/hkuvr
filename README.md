# hkuvr

hkuvr is a HomeKit bridge for an [UVR1611][uvr1611] device. It uses [uvr][uvr] to read data from the CAN bus and the [hc][hc] to communicate with HomeKit.

[uvr1611]: http://www.ta.co.at/en/products/uvr1611/
[hc]: https://github.com/brutella/hc
[uvr]: https://github.com/brutella/uvr

## Build

Build `hkuvr.go` using `go build cmd/hkuvrd.go`. If you're building for the Reaspberry Pi you builds for ARM

     GOOS=linux GOARCH=arm GOARM=6 go build cmd/hkuvr.go

## Run

Execute the daemon with `./hkuvrd`. The daemon will connect to the UVR1611 using the CAN bus on the `can0` interface.
You have to setup the interface as described in [can](https://github.com/brutella/can).

## Pair

The accessory can be paired with any HomeKit app (eg [Home 3][home]) using the pin `001-02-003`.

[home]: https://hochgatterer.me/home

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

hkuvr is available under a non-commercial license. See the LICENSE file for more info.