package generate

type UserData struct {
	User       string
	PublicKeys []string
}

func (n UserData) Name() string {
	return "user-data"
}

func (n UserData) Template() string {
	return `# On first boot, set the admin user's password that must change
chpasswd:
  expire: true
  list:
  - {{ .User }}:{{ .User }}

# Enable password authentication with the SSH daemon
ssh_pwauth: true

users:
- name: {{ .User }}
  sudo: ALL=(ALL) NOPASSWD:ALL
  ssh_authorized_keys:
  {{- range $i, $publicKey := .PublicKeys }}
  - {{ $publicKey }}
  {{- end }}
`
}
