package cmd

import (
	"diagon_ally/settings"
	"diagon_ally/server"
	"github.com/spf13/cobra"
	"log"

	"github.com/rjeczalik/notify"
)

var userSettings *settings.Settings

var rootCmd = &cobra.Command{
	Use:   "diagon",
	Short: "Diagon-ally is a tool to help create and insert diagrams in notes",
	Long: `Diagon-ally watches a directory of template images 
	        (svg only for now, will extend to latex), and exports them to 
          destination images (png by default) whenever the template is edited.
          It also helps create new diagrams by copying over a base template, and
          helps insert them into your notes by filling an insertion template
          and placing it in your clipboard to insert.`,
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
}

func initConfig() {
	var err error
	userSettings, err = settings.GetSettings()
	if err!=nil {
		log.Fatalln(err)
	}
}
