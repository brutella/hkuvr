GO ?= go

rpi:
	GOOS=linux GOARCH=arm GOARM=6 $(GO) build daemon/hkuvrd.go