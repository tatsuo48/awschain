/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	expect "github.com/Netflix/go-expect"
	"github.com/spf13/cobra"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: fmt.Sprintf("%s -- %s", Version, Revision),
	Use:     "awschain",
	Short:   "awschain is set cuurent AWS* environment varibales to envchain namespace",
	Long: `awschain is set cuurent AWS* environment varibales to envchain namespace
examples and usage of using your application. 

For example:
awschain [envchain NAMESPACE]`,

	RunE: func(cobra_cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires only one envchain NAMESPACE")
		}
		namespace := args[0]
		awsEnvs, err := fetchAwsEnvs()
		if err != nil {
			return err
		}
		c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()
		envchainArgs := append([]string{"--set", namespace}, awsEnvs...)
		cmd := exec.Command("envchain", envchainArgs...)
		cmd.Stdin = c.Tty()
		cmd.Stdout = c.Tty()
		cmd.Stderr = c.Tty()

		go func() {
			c.ExpectEOF()
		}()

		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		for _, key := range awsEnvs {
			time.Sleep(time.Second)
			c.Send(fmt.Sprintf("%s\n", os.Getenv(key)))
		}
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

func fetchAwsEnvs() ([]string, error) {
	awsEnvs := []string{}
	for _, pair := range os.Environ() {
		r := regexp.MustCompile(`AWS`)
		if r.MatchString(pair) {
			awsEnvs = append(awsEnvs, strings.Split(pair, "=")[0])
		}
	}
	if len(awsEnvs) == 0 {
		return nil, errors.New("AWS* environmet variable is not found")
	}
	return awsEnvs, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
