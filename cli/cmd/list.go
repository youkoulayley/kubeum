package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/youkoulayley/kubeum/api/models"
	"net/http"
	"os"
	"text/tabwriter"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users of your Kubernetes Cluster.",
	Long:  `The short description is quite exhaustive.`,
	Run: func(cmd *cobra.Command, args []string) {
		var users models.Users

		getUsers, err := http.Get(apiKubeum + "/users")
		if err != nil {
			logrus.Error(getUsers)
		}

		json.NewDecoder(getUsers.Body).Decode(&users)

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tNAMESPACE")

		for _, user := range users {
			line := user.Name + "\t" + user.Namespace

			fmt.Fprintln(w, line)
		}
		fmt.Fprintln(w)
		w.Flush()
	},
}
