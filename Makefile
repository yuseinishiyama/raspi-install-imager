VERSION ?= 20.10
PREINSTALLED_IMAGE := ubuntu-$(VERSION)-preinstalled-server-arm64+raspi.img
DOWNLAOD_URL := http://cdimage.ubuntu.com/releases/$(VERSION)/release/$(PREINSTALLED_IMAGE).xz

HOST ?= $(error HOST must be set)
MOUNTPOINT ?= mnt

default: write

gen/$(HOST):
	go run ./generator -c samples/config.yml --host $(HOST)

image/$(PREINSTALLED_IMAGE):
	curl $(DOWNLAOD_URL) -o - | unxz > $@

image/$(HOST).img: image/$(PREINSTALLED_IMAGE) gen/$(HOST)
	cp $< $@
	@echo mounting image...
	hdiutil attach -mountpoint $(MOUNTPOINT) $@ 1>/dev/null
	@echo enabling ssh...
	touch $(MOUNTPOINT)/ssh
	@echo enabling cgroup...
	sed -i "~" '1s/$$/ cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory/' $(MOUNTPOINT)/cmdline.txt
	@echo overwriting cloud-init configs...
	cp -b $(word 2,$^)/* $(MOUNTPOINT)
	@echo unmounting image...
	hdiutil detach $(MOUNTPOINT)

write: image/$(HOST)
	@echo "TODO: write image to disk"

.PHONY: write
