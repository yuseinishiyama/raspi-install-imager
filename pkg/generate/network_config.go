package generate

type NetworkConfig struct {
	Address      string
	PrefixLength int
	Gateway4     string
	Nameserver   Nameserver
}

type Nameserver struct {
	Address string
	Search  string
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
    - {{ .Address }}/{{ .PrefixLength }}
    gateway4: {{ .Gateway4 }}
    nameservers:
      addresses:
      - {{ .Nameserver.Address }}
      search:
      - {{ .Nameserver.Search }}
`
}
