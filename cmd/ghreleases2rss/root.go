package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/toozej/ghreleases2rss/internal/ghreleases2rss"
	"github.com/toozej/ghreleases2rss/pkg/man"
	"github.com/toozej/ghreleases2rss/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:              "ghreleases2rss",
	Short:            "Subscribe to GitHub projects' releases in RSS reader",
	Long:             `Subscribe to GitHub repo release feeds in Miniflux`,
	Args:             cobra.ExactArgs(0),
	PersistentPreRun: rootCmdPreRun,
	Run:              ghreleases2rss.Run,
}

func rootCmdPreRun(cmd *cobra.Command, args []string) {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return
	}
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	_, err := maxprocs.Set()
	if err != nil {
		log.Error("Error setting maxprocs: ", err)
	}

	// create rootCmd-level flags
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug-level logging")
	rootCmd.PersistentFlags().BoolP("clearCategoryFeeds", "r", false, "Delete all feeds within category before subscribing to new feeds")
	rootCmd.Flags().StringP("file", "f", "", "Input file with GitHub repo URLs or names (required)")
	rootCmd.Flags().StringP("category", "c", "", "RSS feed category name (optional)")
	_ = rootCmd.MarkFlagRequired("file")

	// add sub-commands
	rootCmd.AddCommand(
		man.NewManCmd(),
		version.Command(),
	)
}
