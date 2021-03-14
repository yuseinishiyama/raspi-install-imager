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
	hostname  string
	user      string
	publicKey string
	output    string
)

type config map[string]struct {
	IP     string
	Master bool
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
	rootCmd.Flags().StringVarP(&user, "user", "u", "", "admin user name")
	rootCmd.Flags().StringVarP(&publicKey, "publicKey", "k", "", "ssh public key")
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

	host, ok := conf[hostname]
	if !ok {
		log.Fatalf("hostname %s is not defined in %s", hostname, ConfigFile)
	}

	outputDir := path.Join(output, hostname)

	networkConfig := networkConfig{IP: host.IP}
	if err := networkConfig.generate(outputDir); err != nil {
		log.Fatalf("failed to render network-config. %v", err)
	}

	userData := userData{User: user, PublicKey: publicKey}
	if err := userData.generate(outputDir); err != nil {
		log.Fatalf("failed to render user-data. %v", err)
	}
}
