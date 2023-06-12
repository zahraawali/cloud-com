/*
Copyright © 2023 Zahraa Wali <kareemzahraa286@gmail.com>
*/

package cmd

import (
	"reporter/app"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This command used for start application",
	RunE: func(cmd *cobra.Command, args []string) error {
		dsn, err := cmd.Flags().GetString("dsn")
		if err != nil {
			return err
		}
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			return err
		}

		authService, err := cmd.Flags().GetString("auth")
		if err != nil {
			return err
		}

		return app.Starter(dsn, addr, authService)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String("dsn", "postgres://user:user@127.0.0.1/test?sslmode=disable", "database connection string")
	serveCmd.Flags().String("addr", ":8081", "application host and port")
	serveCmd.Flags().String("auth", "http://127.0.0.1:8082", "authorization services address")
}
