VERSION := 20.10
PREINSTALLED_IMAGE := ubuntu-$(VERSION)-preinstalled-server-arm64+raspi.img
DOWNLAOD_URL := http://cdimage.ubuntu.com/releases/$(VERSION)/release/$(PREINSTALLED_IMAGE).xz

HOST =? $(error HOST must be set)

default: write

gen/$(HOST):
	go run ./generator --host $(HOST)

image/$(PREINSTALLED_IMAGE):
	curl $(DOWNLAOD_URL) -o - | unxz > $@

image/$(HOST): image/$(PREINSTALLED_IMAGE) gen/$(HOST)
	touch $@
	echo "TODO: customized boot image" > $@

write: image/$(HOST)
	@echo "TODO: write image to disk"

.PHONY: write
