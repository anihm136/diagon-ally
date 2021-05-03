package cmd

import (
	"diagon_ally/server"
	"diagon_ally/settings"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/rjeczalik/notify"
)

var userSettings *settings.Settings
var (
	onUpdate string
	source   string
	dest     string
)

var rootCmd = &cobra.Command{
	Use:   "diagon",
	Short: "Diagon-ally is a tool to help create and insert diagrams in notes",
	Long: `Diagon-ally watches a directory of template images 
	        (svg only for now, will extend to latex), and exports them to 
          destination images (png by default) whenever the template is edited.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan notify.EventInfo, 1)
		defer notify.Stop(c)
		done := make(chan struct{}, 1)
		log.Printf("Watching %s\n", userSettings.WatchDir)
		server.WatchDir(c, userSettings.WatchDir)
		server.OnUpdate(c, userSettings)
		<-done
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&source, "source_dir", "s", "", "Directory to watch for changes")
	rootCmd.PersistentFlags().StringVarP(&dest, "dest_dir", "d", "", "Directory to export to")
	rootCmd.PersistentFlags().StringVarP(&onUpdate, "on_update", "u", "", "Command to run for export on update")
}

func initConfig() {
	var err error
	userSettings, err = settings.GetSettings()
	if err != nil {
		log.Fatalln(err)
	}
	userSettings.UpdateFlags(source, dest, onUpdate)
}
