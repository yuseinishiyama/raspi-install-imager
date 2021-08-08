package generate

type NetworkConfig struct {
	Addresses   []string
	Gateway4    string
	Nameservers Nameservers
}

type Nameservers struct {
	Addresses []string
	Search    []string
}

func (n NetworkConfig) Name() string {
	return "network-config"
}

func (n NetworkConfig) Template() string {
	return `version: 2
ethernets:
  eth0:
    dhcp4: false
    addresses:
    {{- range $i, $address := .Addresses }}
    - {{ $address }}
    {{- end }}
    gateway4: {{ .Gateway4 }}
    nameservers:
      addresses:
      {{- range $i, $address := .Nameservers.Addresses }}
      - {{ $address }}
      {{- end }}
      search:
      {{- range $i, $search := .Nameservers.Search }}
      - {{ $search }}
      {{- end }}
`
}
