package generate

type config struct {
	Hosts map[string]struct {
		Address string
		Master  bool
	}
	Shared struct {
		User         string
		PublicKeys   []string `yaml:"ssh_public_keys"`
		PrefixLength int      `yaml:"prefix_length"`
		Gateway4     string
		Nameserver   Nameserver
	}
}
