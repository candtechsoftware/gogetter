package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var (
	installCmd = &cobra.Command{
		Use:   "get",
		Short: "Go Gets multiple packages",
		Long:  "Go Gets multiple packages from cli or from a requirments.txt file",
		Run: func(cmd *cobra.Command, args []string) {
			getPackages(args)
		},
	}
)

func getPackages(packages []string) error {
	for _, p := range packages {
		err := goGet(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func goGet(p string) error {
	cmd := exec.Command("go", "get", p)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}
