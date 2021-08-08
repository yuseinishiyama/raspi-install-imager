package generate

type UserData struct {
	User       string
	PublicKeys []string
}

func (n UserData) Name() string {
	return "user-data"
}

func (n UserData) Template() string {
	return `#cloud-config

# On first boot, set the admin user's password that must change
chpasswd:
  expire: true
  list:
  - {{ .User }}:{{ .User }}

# Enable password authentication with the SSH daemon
ssh_pwauth: true

users:
- name: {{ .User }}
  # groups assigned to the ubuntu default user
  groups: [adm, dialout, cdrom, floppy, sudo, audio, dip, video, plugdev, netdev, lxd]
  sudo: ALL=(ALL) NOPASSWD:ALL
  ssh_import_id:
  - gh:{{ .User }}
`
}
