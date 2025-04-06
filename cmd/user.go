package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/helpers"
	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var userCmd = &cobra.Command{
	Use:     "user",
	Short:   "Manage user",
	Aliases: []string{"u"},
	Args:    cobra.ExactArgs(0),
}

var userInfoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Get user info",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		user, err := Client.Profile.Me(cmd.Context())
		if err != nil {
			fmt.Println("Error getting the user")
		}

		helpers.PrintUser(user.Data.User)
	},
}

var userUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update user info",
	Aliases: []string{"u"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if User.Name == "" && User.Email == "" {
			cmd.Usage()
			os.Exit(1)
		}

		oldProfil, err := Client.Profile.Me(cmd.Context())
		if err != nil {
			fmt.Println("Error getting the user")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if User.Name == "" {
			User.Name = oldProfil.Data.User.Name
		}

		if User.Email == "" {
			User.Email = oldProfil.Data.User.Email
		}

		user, err := Client.Profile.Update(cmd.Context(), terminal.ProfileUpdateParams{
			Name:  terminal.F(User.Name),
			Email: terminal.F(User.Email),
		})
		if err != nil {
			fmt.Println("Error updating the user")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.Section("User info updated", func() {
			helpers.PrintUser(user.Data.User)
		})
	},
}

func init() {
	userCmd.AddCommand(userInfoCmd)
	userCmd.AddCommand(userUpdateCmd)
	userUpdateCmd.Flags().StringVarP(&User.Name, "name", "n", "", "Name")
	userUpdateCmd.Flags().StringVarP(&User.Email, "email", "e", "", "Email")
	rootCmd.AddCommand(userCmd)
}
