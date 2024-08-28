package main

import (
	"context"
	"fmt"
	"github.com/dstgo/lobby/pkg/cfgx"
	app "github.com/dstgo/lobby/server"
	"github.com/dstgo/lobby/server/conf"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"strings"
)

var (
	Version    string
	BuildTime  string
	ConfigFile string
)

var rootCmd = &cobra.Command{
	Use:          "lobby [commands] [-flags]",
	Short:        "lobby is the web server of wendy panel, responsible for managing nodes that from any machine.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// load config file
		appConf := conf.App{
			Version:   Version,
			BuildTime: BuildTime,
		}

		if err := cfgx.LoadConfigAndMapTo(ConfigFile, &appConf); err != nil {
			return err
		}

		// print banner
		if err := app.PrintBanner(os.Stderr); err != nil {
			return err
		}

		appConf.Log.Prompt = "[lobby]"
		// initialize app logger
		logger, err := app.NewLogger(appConf.Log)
		if err != nil {
			return err
		}
		defer logger.Close()

		// set it to the default logger
		slog.SetDefault(logger.Slog())
		slog.Info(fmt.Sprintf("logging in level: %s", strings.ToLower(appConf.Log.Level.String())))

		// this is the root context for the whole program
		rootCtx := context.Background()

		// initialize app
		server, err := app.NewApp(rootCtx, &appConf)
		if err != nil {
			return err
		}

		// run the server
		return server.Spin()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "f", "conf.yaml", "server configuration file")
	rootCmd.AddCommand(versionCmd)
}

func main() {
	rootCmd.Execute()
}
