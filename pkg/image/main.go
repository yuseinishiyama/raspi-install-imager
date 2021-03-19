package image

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type image struct {
	version    string
	files      []string
	cache      string
	mountpoint string
	unmount    bool
	output     string
}

func Command() *cobra.Command {
	image := image{}

	cmd := &cobra.Command{
		Use:   "image",
		Short: "downloads and edits install image",
		Run: func(cmd *cobra.Command, args []string) {
			image.Execute()
		},
	}

	cmd.Flags().StringVarP(&image.version, "version", "v", "20.10", "ubuntu version")
	cmd.Flags().StringSliceVarP(&image.files, "file", "f", []string{}, "files to add")
	cmd.Flags().StringVarP(&image.cache, "cache", "c", ".cache", "directory to keep the original iamge")
	cmd.Flags().StringVarP(&image.mountpoint, "mountpoint", "m", "mnt", "mountpoint of disk image")
	cmd.Flags().BoolVarP(&image.unmount, "unmount", "u", true, "unmount when finish")
	cmd.Flags().StringVarP(&image.output, "output", "o", "image", "location of generated disk image")

	return cmd
}

func (i *image) Execute() {

	out, err := os.Create(i.output)
	if err != nil {
		log.Fatalf("failed to create %s. %v", i.output, err)
	}

	defer out.Close()

	url := downloadURL(i.version)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to get %s. %v", url, err)
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
}

func downloadURL(version string) string {
	return fmt.Sprintf("http://cdimage.ubuntu.com/releases/%s/release/ubuntu-%s-preinstalled-server-arm64+raspi.img.xz", version, version)
}
