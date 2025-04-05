package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
)

var addressCmd = &cobra.Command{
	Use:     "address",
	Short:   "Manage addresses",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(0),
}

var listAddressesCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all addresses",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		addresses, err := Client.Address.List(context.Background())
		if err != nil {
			fmt.Println("Error getting the addresses")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		printAddresses(addresses.Data)
	},
}

var createAddressCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create address",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Address.New(cmd.Context(), terminal.AddressNewParams{
			Name:     terminal.F(Address.Name),
			Country:  terminal.F(Address.Country),
			Province: terminal.F(Address.Province),
			City:     terminal.F(Address.City),
			Zip:      terminal.F(Address.Zip),
			Street1:  terminal.F(Address.Street1),
			Street2:  terminal.F(Address.Street2),
			Phone:    terminal.F(Address.Phone),
		})

		if err != nil {
			fmt.Println("Error creating the address")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Address created")
	},
}

var deleteAddressCmd = &cobra.Command{
	Use:     "delete [id]",
	Short:   "Delete address",
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := Client.Address.Delete(cmd.Context(), args[0])
		if err != nil {
			fmt.Println("Can not delete address")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Address deleted")
	},
}

func printAddresses(addresses []terminal.Address) {

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "ID\tName\tCountry\tProvince\tCity\tZip\tStreet1\tStreet2")
	fmt.Fprintln(w, "--\t----\t-------\t--------\t----\t---\t-------\t-------")
	for _, address := range addresses {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", address.ID, address.Name, address.Country, address.Province, address.City, address.Zip, address.Street1, address.Street2)
	}
	w.Flush()
}
