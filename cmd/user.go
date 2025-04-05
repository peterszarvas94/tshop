package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

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

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "Name\tEmail")
		fmt.Fprintln(w, "----\t-----")
		fmt.Fprintf(w, "%s\t%s\n", user.Data.User.Name, user.Data.User.Email)
		w.Flush()
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

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "Successfully modified user")
		fmt.Fprintln(w, "Name\tEmail")
		fmt.Fprintln(w, "----\t-----")
		fmt.Fprintf(w, "%s\t%s\n", user.Data.User.Name, user.Data.User.Email)
		w.Flush()
	},
}
