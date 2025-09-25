package commands

import (
	"fiber-gorm-channel-ecommerce/src/pkg/boot"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed [seeder_name]",
	Short: "seed data ",
	Long:  "seed data to database",
	Run:   seed,
}

func seed(cmd *cobra.Command, args []string) {
	boot.SeedData(args)
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
