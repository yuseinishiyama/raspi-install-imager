package generate

type config struct {
	Hosts map[string]struct {
		Addresses []string
		Master    bool
	}
	Shared struct {
		User        string
		PublicKeys  []string `yaml:"ssh_public_keys"`
		Gateway4    string
		Nameservers Nameservers
	}
}
