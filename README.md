# hkuvr1611

This project is an implementation of a HomeKit bridge for an [UVR1611][uvr1611] device. It uses the [uvr][uvr] library to read data from the CAN bus and the [HomeControl][hc] library to communicate with HomeKit.

[uvr1611]: http://www.ta.co.at/en/products/uvr1611/
[hc]: https://github.com/brutella/hc
[uvr]: https://github.com/brutella/uvr

## Install

    cd $GOPATH/src

    # Clone project
    git clone https://github.com/brutella/hkuvr && cd hkuvr

    # Install dependencies
    go get

## Build

Build `hkuvrd.go` using `go build daemon/hkuvrd.go`. If you're building for the Reaspberry Pi you builds for ARM

     GOOS=linux GOARCH=arm GOARM=6 go build daemon/hkuvrd.go

## Run

Execute the daemon with `./hkuvrd`. The daemon will connect to the UVR1611 using the CAN bus on the `can0` interface.
You have to setup the interface as described in [can](https://github.com/brutella/can).

## Pair

The accessory can be paired with any HomeKit app like [Home][home] (which runs on iPhone, iPad and Apple Watch) using the pin `001-02-003`.

[home]: http://selfcoded.com/home
[home-getting-started]: http://selfcoded.com/home/getting-started

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

hkuvr is available under a non-commercial license. See the LICENSE file for more info.