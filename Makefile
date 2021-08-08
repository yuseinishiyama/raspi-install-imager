TARGET_DISK ?= /dev/disk2
HOSTS := pi1 pi2 pi3
WRITE_TARGETS := $(patsubst %,%.write,$(HOSTS))

$(HOSTS):
	go run ./cmd/raspi-install-imager generate --host $@ -c samples/config.yml -o artifacts
	go run ./cmd/raspi-install-imager image -o artifacts/$@.img -c artifacts/$@ -m mnt

%.write: %
	go run ./cmd/raspi-install-imager write -d $(TARGET_DISK) -i artifacts/$<.img
