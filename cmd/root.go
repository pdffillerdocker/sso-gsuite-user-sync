// Copyright (c) 2020, Amazon.com, Inc.
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

package cmd

import (
	"context"
	"fmt"
    "os"

    "github.com/pdffillerdocker/sso-gsuite-user-sync/internal"
    "github.com/pdffillerdocker/sso-gsuite-user-sync/internal/config"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	commit  = "first"
	date    = "Dec 3, 2020"
	builtBy = "daimon"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Version: "dev",
	Use:     "sso-gsuite-user-sync",
	Short:   "SSO Sync, making AWS SSO be populated automagically",
	Long: `A command line tool to enable you to synchronise your Google
Apps (G-Suite) users to AWS Single Sign-on (AWS SSO)
Complete documentation is available at https://github.com/pdffillerdocker/sso-gsuite-user-sync`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := internal.DoSync(ctx, cfg)
		if err != nil {
			return err
		}

		return nil
	},
}

// Execute is the entry point of the command. If we are
// running inside of AWS Lambda, we use the Lambda
// execution path.
func Execute() {
	if cfg.IsLambda {
		lambda.Start(rootCmd.Execute)
	} else {
	    if err := rootCmd.Execute(); err != nil {
		    log.Fatal(err)
	    }
    }
}

func init() {
	// init config
	cfg = config.New()
    cfg.IsLambda = len(os.Getenv("AWS_LAMBDA_FUNCTION_NAME")) > 0

	// initialize cobra
	cobra.OnInitialize(initConfig)
	addFlags(rootCmd, cfg)

	rootCmd.SetVersionTemplate(fmt.Sprintf("%s, commit %s, built at %s by %s\n", version, commit, date, builtBy))

	// silence on the root cmd
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// allow to read in from environment
	viper.SetEnvPrefix("ssosync")
	viper.AutomaticEnv()

	for _, e := range []string{"secret_bucket", "secret_bucket_region", "google_admin", "google_credentials", "scim_access_token", "scim_endpoint", "log_level", "log_format", "ignore_users", "debug"} {
		if err := viper.BindEnv(e); err != nil {
			log.Fatalf(errors.Wrap(err, "cannot bind environment variable").Error())
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(errors.Wrap(err, "cannot unmarshal config").Error())
	}

	// config logger
	logConfig(cfg)
}

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	//rootCmd.PersistentFlags().StringVarP(&cfg.GoogleCredentials, "google-credentials", "a", config.DefaultGoogleCredentials, "set the credentials for Google")

	rootCmd.PersistentFlags().BoolVarP(&cfg.Debug, "debug", "d", config.DefaultDebug, "Enable verbose / debug logging")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogFormat, "log-format", "", config.DefaultLogFormat, "log format")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogLevel, "log-level", "", config.DefaultLogLevel, "log level")

	rootCmd.Flags().StringVarP(&cfg.SecretBucket, "secret_bucket", "b", "", "Secret Bucket name")
	rootCmd.Flags().StringVarP(&cfg.SecretBucketRegion, "secret_bucket_region", "r", "", "Secret Bucket region")
	rootCmd.Flags().StringVarP(&cfg.SCIMAccessToken, "access-token", "t", "", "The Path to file with SCIM Access Token")
	rootCmd.Flags().StringVarP(&cfg.SCIMEndpoint, "endpoint", "e", "", "SCIM Endpoint")
	rootCmd.Flags().StringVarP(&cfg.GoogleCredentials, "google-credentials", "c", "", "The path to file with credentials for Google")
	rootCmd.Flags().StringVarP(&cfg.GoogleAdmin, "google-admin", "u", "", "Google Admin Email")
	rootCmd.Flags().StringSliceVar(&cfg.IgnoreUsers, "ignore-users", []string{}, "ignores these users")

	rootCmd.Flags().StringVarP(&cfg.LogLevel, "log-level", "", config.DefaultLogLevel, "log level")
    rootCmd.Flags().StringVarP(&cfg.LogFormat, "log-format", "", config.DefaultLogFormat, "log format")
    rootCmd.Flags().BoolVarP(&cfg.Debug, "debug", "d", config.DefaultDebug, "Enable verbose / debug logging")
}

func logConfig(cfg *config.Config) {
	// reset log format
	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if cfg.Debug {
		cfg.LogLevel = "debug"
	}

	// set the configured log level
	if level, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(level)
	}
}
