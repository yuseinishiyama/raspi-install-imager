package generate

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type generate struct {
	configPath string
	hostname   string
	output     string
}

func Command() *cobra.Command {
	gen := generate{}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generates cloud-init configs",
		Run: func(cmd *cobra.Command, args []string) {
			gen.Execute()
		},
	}

	cmd.Flags().StringVarP(&gen.configPath, "config", "c", "config.yml", "configuration file")
	cmd.Flags().StringVar(&gen.hostname, "host", "", "host name")
	cmd.Flags().StringVarP(&gen.output, "output", "o", "", "root output directory")

	return cmd
}

func (g *generate) Execute() {
	if len(g.output) == 0 {
		log.Fatal("--output must be specified")
	}

	bytes, err := ioutil.ReadFile(g.configPath)
	if err != nil {
		log.Fatalf("failed to read %s. %v", g.configPath, err)
	}

	conf := config{}
	if err := yaml.Unmarshal(bytes, &conf); err != nil {
		log.Fatalf("failed to unmarshal %s. %v", g.configPath, err)
	}

	host, ok := conf.Hosts[g.hostname]
	if !ok {
		log.Fatalf("hostname %q is not defined in %s", g.hostname, g.configPath)
	}

	var masterAddr string
	for _, e := range conf.Hosts {
		if e.Master {
			masterAddr = e.Address
		}
	}

	// unset master address if it's the master
	if masterAddr == host.Address {
		masterAddr = ""
	}

	outputDir := path.Join(g.output, g.hostname)

	networkConfig := NetworkConfig{
		Address:      host.Address,
		PrefixLength: conf.Shared.PrefixLength,
		Gateway4:     conf.Shared.Gateway4,
		Nameserver:   conf.Shared.Nameserver,
	}

	userData := UserData{
		Host:          g.hostname,
		User:          conf.Shared.User,
		MasterAddress: masterAddr,
		PublicKeys:    conf.Shared.PublicKeys,
	}

	for _, template := range []templating{networkConfig, userData} {
		if err := Generate(template, outputDir); err != nil {
			log.Fatalf("failed to render user-data. %v", err)
		}
	}
}
