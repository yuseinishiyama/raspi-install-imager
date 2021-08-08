package generate

type UserData struct {
	Host          string
	User          string
	MasterAddress string
	PublicKeys    []string
}

func (n UserData) Name() string {
	return "user-data"
}

func (n UserData) Template() string {
	return `#cloud-config

hostname: {{ .Host }}

# On first boot, set the (default) ubuntu user's password to "ubuntu" and
# expire user passwords
chpasswd:
  expire: true
  list:
  - ubuntu:ubuntu

# Disable password authentication with the SSH daemon
ssh_pwauth: false

users:
- default
- name: {{ .User }}
  # groups assigned to the ubuntu default user
  groups: [adm, dialout, cdrom, floppy, sudo, audio, dip, video, plugdev, netdev, lxd]
  sudo: ALL=(ALL) NOPASSWD:ALL
  ssh_import_id: gh:{{ .User }}

runcmd:
{{- if .MasterAddress }}
  - curl -sfL https://get.k3s.io | K3S_URL=https://{{ .MasterAddress }}:6443 K3S_TOKEN=token sh -
{{ else }}
  - curl -sfL https://get.k3s.io | K3S_TOKEN=token sh -
{{- end }}
`
}
