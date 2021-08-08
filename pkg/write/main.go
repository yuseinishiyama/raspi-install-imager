package write

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

type write struct {
	disk  string
	image string
}

func Command() *cobra.Command {
	write := write{}

	cmd := &cobra.Command{
		Use:   "write",
		Short: "write image to device",
		Run: func(cmd *cobra.Command, args []string) {
			write.Execute()
		},
	}

	cmd.Flags().StringVarP(&write.disk, "disk", "d", "", "disk to write to")
	cmd.Flags().StringVarP(&write.image, "image", "i", "", "image to write")

	return cmd
}

func (w *write) Execute() {
	if err := w.validate(); err != nil {
		log.Fatal(err)
	}

	exec.Command("diskutil", "unmountDisk", w.disk).Run()

	rawName := filepath.Join(filepath.Dir(w.disk), "r" + filepath.Base(w.disk))
	log.Printf("copying %s to %s", w.image, rawName)
	out, err := exec.Command(
		"sudo",
		"dd",
		"bs=1M",
		fmt.Sprintf("if=%s", w.image),
		fmt.Sprintf("of=%s", rawName),
	).CombinedOutput()

	if err != nil {
		log.Fatalf("failed to write %s to %s, error: %s", w.image, w.disk, string(out))
	}

	exec.Command("sync").Run()

	log.Printf("ejecting %s", w.disk)
	exec.Command("diskutil", "eject", w.disk).Run()
}

func (w *write) validate() error {
	if len(w.disk) == 0 {
		return fmt.Errorf("--disk must be specified")
	}

	if len(w.image) == 0 {
		return fmt.Errorf("--image must be specified")
	}

	out, err := exec.Command("diskutil", "list").CombinedOutput()
	if err != nil {
		return fmt.Errorf(string(out))
	}

	r := regexp.MustCompile(fmt.Sprintf("(?m)^%s \\(external, physical\\)", w.disk))
	if !r.MatchString(string(out)) {
		return fmt.Errorf("%s can't be found as an external disk", w.disk)
	}

	if _, err := os.Stat(w.image); os.IsNotExist(err) {
		return fmt.Errorf("%s can't be found", w.image)
	}

	return nil
}
