// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package internal provides internal functionality for the dbotconf package.
package internal

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dbotconf",
		Short: "Dependabot configuration utility",
		Long:  "dbotconf manages Dependabot configuration for multi-module Go repositories.",
		Example: `
  dbotconf generate > .github/dependabot.yml

  dbotconf verify .github/dependabot.yml`,
	}

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate Dependabot configuration",
		Run:   runGenerate,
	}

	verifyCmd = &cobra.Command{
		Use:   "verify [flags] path",
		Short: "Verify Dependabot configuration is complete",
		Long:  "Ensure Dependabot configuration contains update checks for all modules in the repository.",
		Run:   runVerify,
	}
)

const ignoreFlag = "ignore"

// BuildAndExecute runs the dbotconf command.
func BuildAndExecute() error {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(verifyCmd)

	rootCmd.PersistentFlags().StringSlice(ignoreFlag, nil, "glob patterns to ignore")

	return rootCmd.Execute()
}
