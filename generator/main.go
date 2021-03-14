package main

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const (
	ConfigFile = "config.yml"
)

var (
	hostname string
	output   string
)

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

func main() {
	rootCmd := &cobra.Command{
		Use:   "generator",
		Short: "generate drop-in configurations for Raspberry Pi boot image",
		Run: func(cmd *cobra.Command, args []string) {
			execute()
		},
	}

	rootCmd.Flags().StringVar(&hostname, "host", "", "host name")
	rootCmd.Flags().StringVarP(&output, "output", "o", "gen", "root output directory")

	rootCmd.Execute()
}

func execute() {
	bytes, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("failed to read %s. %v", ConfigFile, err)
	}

	conf := config{}
	if err := yaml.Unmarshal(bytes, &conf); err != nil {
		log.Fatalf("failed to unmarshal %s. %v", ConfigFile, err)
	}

	host, ok := conf.Hosts[hostname]
	if !ok {
		log.Fatalf("hostname %s is not defined in %s", hostname, ConfigFile)
	}

	outputDir := path.Join(output, hostname)

	networkConfig := networkConfig{
		Addresses:   host.Addresses,
		Gateway4:    conf.Shared.Gateway4,
		Nameservers: conf.Shared.Nameservers,
	}
	if err := generate("templates/network-config", networkConfig, outputDir); err != nil {
		log.Fatalf("failed to render network-config. %v", err)
	}

	userData := userData{
		User:       conf.Shared.User,
		PublicKeys: conf.Shared.PublicKeys,
	}
	if err := generate("templates/user-data", userData, outputDir); err != nil {
		log.Fatalf("failed to render user-data. %v", err)
	}
}
