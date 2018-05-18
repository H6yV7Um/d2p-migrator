package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/reexec"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Config of pouch-migrator
type Config struct {
	DockerPkg    string
	PouchPkgPath string
	MigrateAll   bool
	Debug        bool
}

var cfg = &Config{}

func main() {
	if reexec.Init() {
		return
	}

	var cmdServe = &cobra.Command{
		Use:          "pouch-migrator",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCmd()
		},
	}

	setupFlags(cmdServe)
	parseFlags(cmdServe, os.Args[1:])

	if err := cmdServe.Execute(); err != nil {
		fmt.Printf("failed to execute pouch-migrator: %v", err)
		os.Exit(1)
	}

}

// initLog initializes log Level and log format of daemon.
func initLog() {
	if cfg.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Infof("start daemon at debug level")
	}

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	}
	logrus.SetFormatter(formatter)
}

func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()

	flagSet.StringVar(&cfg.DockerPkg, "docker-pkg", "docker", "Specify docker package name")
	flagSet.StringVar(&cfg.PouchPkgPath, "pouch-pkg-path", "pouch", "Specify pouch package file path")
	flagSet.BoolVar(&cfg.MigrateAll, "migrate-all", false, "If true, do all migration things, otherwise, just prepare data for migration")
	flagSet.BoolVarP(&cfg.Debug, "debug", "D", false, "DEBUG mode log level")
}

func parseFlags(cmd *cobra.Command, flags []string) {
	err := cmd.Flags().Parse(flags)
	if err == nil || err == pflag.ErrHelp {
		return
	}

	cmd.SetOutput(os.Stderr)
	cmd.Usage()
}

// runCmd prepares configs, setups essential details and runs pouch-migrator
func runCmd() error {
	// initialize log
	initLog()

	migrator, err := NewPouchMigrator(cfg.DockerPkg, cfg.PouchPkgPath, cfg.Debug)
	if err != nil {
		logrus.Errorf("failed to new pouch migrator: %v\n", err)
		os.Exit(1)
	}

	defer migrator.Cleanup()

	if err := migrator.PreMigrate(); err != nil {
		logrus.Errorf("failed to execute PreMigrage: %v", err)
		return err
	}

	if !cfg.MigrateAll {
		logrus.Infof("Just prepare data:\n * Pull Image \n * Prepare snapshots \n * Set disk quota \n * Convert docker container meta json file to pouch container meta json file ")
		return nil
	}

	// If migration failed, revert it.
	needRevert := false
	defer func() {
		if needRevert {
			logrus.Info("Migration has failed, start revert...")

			if err := migrator.RevertMigration(); err != nil {
				logrus.Errorf("failed to revert migration: %v", err)
				return
			}

			logrus.Info("Revert migration done!\n")
		}
	}()

	if err := migrator.Migrate(); err != nil {
		logrus.Infof("failed to migrate: %v\n", err)
		needRevert = true
		return err
	}

	// If PostMigrate failed, should handle by manual.
	if err := migrator.PostMigrate(); err != nil {
		logrus.Errorf("PostMigrate error: %v, need handle by manual!!!", err)
		return err
	}

	return nil
}