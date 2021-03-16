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
	cmd.Flags().StringVarP(&gen.output, "output", "o", "gen", "root output directory")

	return cmd
}

func (g *generate) Execute() {
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

	outputDir := path.Join(g.output, g.hostname)

	networkConfig := NetworkConfig{
		Addresses:   host.Addresses,
		Gateway4:    conf.Shared.Gateway4,
		Nameservers: conf.Shared.Nameservers,
	}

	userData := UserData{
		User:       conf.Shared.User,
		PublicKeys: conf.Shared.PublicKeys,
	}

	for _, template := range []templating{networkConfig, userData} {
		if err := Generate(template, outputDir); err != nil {
			log.Fatalf("failed to render user-data. %v", err)
		}
	}
}
