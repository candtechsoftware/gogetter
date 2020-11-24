package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var (
	// Flag varaible buckets
	domain   string
	userName string

	makeProjectCmd = &cobra.Command{
		Use:   "mk",
		Short: "Creates a project",
		Long:  "Creates a project and a git repo for a go project",
		Run: func(cmd *cobra.Command, args []string) {
			username, _ := cmd.Flags().GetString("user")
			domain, _ := cmd.Flags().GetString("domain")
			git, _ := cmd.Flags().GetBool("git")
			projectName := args[0]
			module := fmt.Sprintf("%s/%s/%s", domain, username, projectName)

			err := createProject(module, projectName, git)
			if err != nil {
				log.Fatal("Errror in creating project ", err)
			}
		},
	}
)

// Helper function to create a git repository in the current working directory
func createGitRepo() {
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Git Command failed with %s \n", err)
	}
}

func createGoModule(module string) {
	cmd := exec.Command("go", "mod", "init", module)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed in creating a go module command: %s", err)
	}
}

// Creates  project
func createProject(module, projectName string, isRepo bool) error {

	_, err := os.Stat(projectName)

	if os.IsNotExist(err) {
		if errDir := os.MkdirAll(projectName, 0755); errDir != nil {
			return errDir
		}
	}

	os.Chdir("./" + projectName)

	if isRepo {
		createGitRepo()
	}
	createGoModule(module)
	return nil
}

func init() {
	rootCmd.AddCommand(makeProjectCmd)
	makeProjectCmd.Flags().StringVarP(&userName, "user", "u", "", "User name")
	makeProjectCmd.Flags().StringVarP(&domain, "domain", "d", "github.com", "Base Domain name")
	makeProjectCmd.Flags().Bool("git", true, "Create a git repo")
}
