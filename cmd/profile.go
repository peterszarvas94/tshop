package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var profilCmd = &cobra.Command{
	Use:     "profile",
	Short:   "Manage profile",
	Aliases: []string{"pr"},
	Args:    cobra.ExactArgs(0),
}

var profilInfoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Get profile info",
	Aliases: []string{"i"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		profil, err := Client.Profile.Me(context.Background())
		if err != nil {
			panic(err.Error())
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "Name\tEmail")
		fmt.Fprintln(w, "----\t-----")
		fmt.Fprintf(w, "%s\t%s\n", profil.Data.User.Name, profil.Data.User.Email)
		w.Flush()
	},
}

var profilUpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update profil info",
	Aliases: []string{"u"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" || email == "" {
			cmd.Usage()
			os.Exit(1)
		}

		profil, err := Client.Profile.Update(cmd.Context(), terminal.ProfileUpdateParams{
			Name:  terminal.F(name),
			Email: terminal.F(email),
		})
		if err != nil {
			panic(err.Error())
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "Successfully modified profil")
		fmt.Fprintln(w, "Name\tEmail")
		fmt.Fprintln(w, "----\t-----")
		fmt.Fprintf(w, "%s\t%s\n", profil.Data.User.Name, profil.Data.User.Email)
		w.Flush()
	},
}
