package cmd

import (
	"diagon_ally/server"
	"diagon_ally/settings"
	"log"

	"github.com/spf13/cobra"

	"github.com/rjeczalik/notify"
)

var userSettings *settings.Settings
var (
	onUpdate            string
	source              string
	dest                string
	forceCreateSettings bool
)

var rootCmd = &cobra.Command{
	Use:   "diagon",
	Short: "Diagon-ally is a tool to help create and insert diagrams in notes",
	Long: `Diagon-ally watches a directory of template images, and exports them to 
          destination images whenever a new template is created or an existing template is edited.`,
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
	rootCmd.PersistentFlags().BoolVarP(&forceCreateSettings, "force", "f", false, "Force creation of settings if it does not exist")
}

func initConfig() {
	var err error
	userSettings, err = settings.GetSettings(forceCreateSettings)
	if err != nil {
		log.Fatalln(err)
	}
	userSettings.UpdateFlags(source, dest, onUpdate)
}
