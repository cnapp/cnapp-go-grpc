// Copyright (C) 2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	goflag "flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
	// init glog to get its flags
	_ "github.com/golang/glog"

	pkgcmd "github.com/cnapp/cnapp-go-grpc/pkg/cmd"
)

var (
	cliName           = "cnappd"
	helpMessage       = "cnappd - The cnapp server"
	completionExample = `
               # Load the cnappd completion code for bash into the current shell
	       source <(cnappd completion bash)

               # Write bash completion code to a file and source if from .bash_profile
	       cnappd completion bash > ~/.cnapp/completion.bash.inc
	       printf "\n# cnappd shell completion\nsource '$HOME/.cnapp/completion.bash.inc'\n" >> $HOME/.bash_profile
               source $HOME/.bash_profile

	       # Load the cnappd completion code for zsh[1] into the current shell
	       source <(cnappd completion zsh)`
)

func init() {
	// Tell gRPC not to log to console.
	grpclog.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags))
}

// NewcnappdCommand creates the `cnappd` command and its nested children.
func NewcnappdCommand(out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   cliName,
		Short: "command-line tool to setup cnapp server",
		Long:  `cnappd is a command-line tool to setup a cnapp server.`,
	}
	rootCmd.AddCommand(
		newRunCmd(out),
		pkgcmd.NewVersionCmd(out, helpMessage),
		pkgcmd.NewCompletionCommand(out, completionExample),
	)
	cobra.EnablePrefixMatching = true

	// add glog flags
	rootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	// https://github.com/kubernetes/dns/pull/27/files
	goflag.CommandLine.Parse([]string{})

	return rootCmd
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := NewcnappdCommand(os.Stdout)
	if err := cmd.Execute(); err != nil {
		fmt.Println(pkgcmd.RedOut(err))
		os.Exit(1)
	}
}
