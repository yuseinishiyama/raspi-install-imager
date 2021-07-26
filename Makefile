gen-config:
	go run ./cmd/raspi-install-imager generate --host pi1 -c samples/config.yml -o artifacts
	go run ./cmd/raspi-install-imager generate --host pi2 -c samples/config.yml -o artifacts
	go run ./cmd/raspi-install-imager generate --host pi3 -c samples/config.yml -o artifacts

gen-image: gen-config
	go run ./cmd/raspi-install-imager image -o artifacts/pi1.img -c artifacts/pi1 -m mnt
#	go run ./cmd/raspi-install-imager image -o artifacts/pi2.img -c artifacts/pi2 -m mnt
#	go run ./cmd/raspi-install-imager image -o artifacts/pi3.img -c artifacts/pi3 -m mnt

.PHONY: gen-config gen-image
