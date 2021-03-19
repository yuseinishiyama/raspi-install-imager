package main

import (
	"github.com/spf13/cobra"
	"github.com/yuseinishiyama/raspi-image-installer/pkg/generate"
	"github.com/yuseinishiyama/raspi-image-installer/pkg/image"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "raspi-install-imager",
		Short: "generates Raspberry Pi boot image out of YAML",
	}
	rootCmd.AddCommand(generate.Command())
	rootCmd.AddCommand(image.Command())
	rootCmd.Execute()
}
