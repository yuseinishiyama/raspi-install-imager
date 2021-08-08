# raspi-install-imager

Generate your Raspberry Pi install image out of a YAML file. **This supports only macOS**.

```bash
# Generate cloud-init config files
raspi-install-imager generate --host pi1 -c config.yml -o artifacts

# Download the install image and make necessary modifications
raspi-install-imager image -o artifacts/pi1.img -c artifacts/pi1 -m mnt

# Write it to a disk (e.g. SD card)
raspi-install-imager write -d /dev/disk2 -i artifacts/pi1.img
```

## What

'raspi-install-imager' generates cloud-init configs, embeds them into an official Raspberry Pi install image, and bakes it into a SD card.

## Why

Setting up Raspberry Pi is tedious especially when you have a cluster. Instead, have an install image that does everything for you.

## Features

- [x] Enabling SSH
- [x] Enabling cgroup
- [x] Static IPs
- [x] Custom user
- [x] Hostname
- [x] Kubernetes
