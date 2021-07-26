package image

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ulikunitz/xz"

	"github.com/yuseinishiyama/raspi-image-installer/pkg/utils"
)

type image struct {
	version    string
	config     string
	cache      string
	mountpoint string
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
	cmd.Flags().StringVarP(&image.config, "config", "c", "", "path to config file dir")
	cmd.Flags().StringVarP(&image.cache, "cache", "C", ".cache", "directory to keep the original iamge")
	cmd.Flags().StringVarP(&image.mountpoint, "mountpoint", "m", "mnt", "mountpoint of disk image")
	cmd.Flags().StringVarP(&image.output, "output", "o", "", "location of generated disk image")

	return cmd
}

func (i *image) Execute() {
	if len(i.output) == 0 {
		log.Fatal("--output must be specified")
	}

	if err := i.createCache(); err != nil {
		log.Fatal(err)
	}
	if err := i.copyImage(); err != nil {
		log.Fatal(err)
	}
	if err := i.updateImage(); err != nil {
		log.Fatal(err)
	}
}

func (i *image) createCache() error {
	if _, err := os.Stat(i.imageCachePath()); os.IsNotExist(err) {
		if err := os.MkdirAll(i.cache, os.ModePerm); err != nil {
			return err
		}

		imageCache, err := os.Create(i.imageCachePath())
		if err != nil {
			return err
		}
		defer imageCache.Close()

		resp, err := http.Get(i.downloadURL())
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		unarchived, err := xz.NewReader(resp.Body)
		if err != nil {
			return err
		}

		_, err = io.Copy(imageCache, unarchived)
		return err
	}

	return nil
}

func (i *image) copyImage() error {
	if _, err := os.Stat(i.output); os.IsNotExist(err) {
		return utils.SimplyCopy(i.output, i.imageCachePath())
	}

	return nil
}

// updateImage makes necessary modifications to original image. Must be idempotent
func (i *image) updateImage() error {
	if err := i.mount(); err != nil {
		return err
	}
	defer i.unmount()

	if err := i.enableSSH(); err != nil {
		return err
	}

	if err := i.enableCgroup(); err != nil {
		return err
	}

	if err := i.copyFiles(); err != nil {
		return err
	}

	return nil
}

func (i *image) enableSSH() error {
	ssh, err := os.Create(filepath.Join(i.mountpoint, "ssh"))
	if err != nil {
		return err
	}
	defer ssh.Close()
	return nil
}

func (i *image) enableCgroup() error {
	cmdFile := filepath.Join(i.mountpoint, "cmdline.txt")
	cmd, err := ioutil.ReadFile(cmdFile)
	if err != nil {
		return err
	}
	flags := "cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory"
	if strings.Contains(string(cmd), flags) {
		return nil
	}
	cmdString := strings.ReplaceAll(string(cmd), "\n", fmt.Sprintf(" %s\n", flags))
	return ioutil.WriteFile(cmdFile, []byte(cmdString), os.ModePerm)
}

func (i *image) copyFiles() error {
	return filepath.Walk(i.config, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		dst := filepath.Join(i.mountpoint, filepath.Base(path))
		err := utils.SimplyCopy(dst, path)
		if err != nil {
			return err
		}
		return nil
	})
}

func (i *image) mount() error {
	out, err := exec.Command("hdiutil", "attach", "-mountpoint", i.mountpoint, i.output).CombinedOutput()
	if err != nil {
		return fmt.Errorf(string(out))
	}
	return nil
}

func (i *image) unmount() error {
	out, err := exec.Command("hdiutil", "detach", i.mountpoint).CombinedOutput()
	if err != nil {
		return fmt.Errorf(string(out))
	}
	return nil
}

func (i *image) imageCachePath() string {
	filename := filepath.Base(i.downloadURL())
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	return filepath.Join(i.cache, filename)
}

func (i *image) downloadURL() string {
	return fmt.Sprintf("http://cdimage.ubuntu.com/releases/%s/release/ubuntu-%s-preinstalled-server-arm64+raspi.img.xz", i.version, i.version)
}
