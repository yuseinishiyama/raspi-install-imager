# raspi-install-imager

Generate your Raspberry Pi install image out of YAML

_NOTE: This project is WIP and works only partially._

## What

'raspi-install-imager' generates cloud-init configs, embeds them into an official Raspberry Pi install image, and bakes it into a SD card.

## Why

Setting up Raspberry Pi is tedious especially when you have a cluster. Instead, have an install image that does everything for you.

## Features

- [x] Enabling SSH
- [x] Enabling cgroup
- [x] Static IPs
- [x] Custom default user
- [ ] Hostname
- [ ] HDMI settings
- [ ] Kubernetes
