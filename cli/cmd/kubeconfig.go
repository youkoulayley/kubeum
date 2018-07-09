package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var outFile string

func init() {
	kubeconfigCmd.PersistentFlags().StringVarP(&outFile, "out-file", "o", "", "send the result to the specified file.")
	rootCmd.AddCommand(kubeconfigCmd)
}

var kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig [USER] [NAMESPACE]",
	Short: "Generate a kubeconfig for a particular user.",
	Long:  `Generate a kubeconfig for a particular user.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("you have to specify an username and a namespace")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		postJson := "{\"name\": \"" + args[0] + "\", \"namespace\": \"" + args[1] + "\"}"
		postData := []byte(postJson)

		postKubeconfig, err := http.Post(apiKubeum+"/users/kubeconfig", "application/json", bytes.NewBuffer(postData))
		if err != nil {
			logrus.Error(err)
		}

		body, err := ioutil.ReadAll(postKubeconfig.Body)
		if err != nil {
			logrus.Error(err.Error())
		}

		if strings.TrimSpace(outFile) != "" {
			file, err := os.Create(outFile)
			if err != nil {
				logrus.Error("Unable to create file : " + err.Error())
			}
			defer file.Close()

			file.Write(body)
		} else {
			fmt.Print(string(body))
		}
	},
}
