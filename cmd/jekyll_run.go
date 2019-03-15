package cmd

import (
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tmSpiderSegmentFaultCmd represents the taskrm command
var jekyllServeCmd = &cobra.Command{
	Use:   "jekyllb",
	Short: "bundle execu jekyll $1",
	Long:  `alias bundle exec jekyll in golang`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		techMojoJekyllRun("build")
	},
}

func init() {
	rootCmd.AddCommand(jekyllServeCmd)
}

func techMojoJekyllRun(buildOrServe string) {
	jekyllDir := viper.GetString("tech_mojotv_cn.srcDir")

	thisCmd := exec.Command("bundle", "exec", "jekyll", buildOrServe)
	thisCmd.Dir = jekyllDir
	o, err := thisCmd.CombinedOutput()
	logrus.Info(string(o))
	if err != nil {
		logrus.Error(err)
	}
}
