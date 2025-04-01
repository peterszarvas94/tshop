package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
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
