package cmd

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/tshop/helpers"
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
		addresses, err := Client.Address.List(cmd.Context())
		if err != nil {
			fmt.Println("Error getting the addresses")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		helpers.PrintAddresses(addresses.Data)
	},
}

var createAddressCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create address",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := Client.Address.New(cmd.Context(), terminal.AddressNewParams{
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

		helpers.Section("Address created", func() {
			Address.ID = res.Data
			helpers.PrintAddresses([]terminal.Address{*Address})
		})
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

func init() {
	addressCmd.AddCommand(listAddressesCmd)
	addressCmd.AddCommand(createAddressCmd)
	createAddressCmd.Flags().StringVarP(&Address.Name, "name", "n", "", "Name")
	createAddressCmd.Flags().StringVarP(&Address.Country, "country", "c", "", "Country")
	createAddressCmd.Flags().StringVarP(&Address.Province, "province", "p", "", "Province")
	createAddressCmd.Flags().StringVarP(&Address.City, "city", "y", "", "City")
	createAddressCmd.Flags().StringVarP(&Address.Zip, "zip", "z", "", "Zip")
	createAddressCmd.Flags().StringVarP(&Address.Street1, "street1", "s", "", "Street1")
	createAddressCmd.Flags().StringVarP(&Address.Street2, "street2", "t", "", "Street2")
	createAddressCmd.Flags().StringVarP(&Address.Phone, "phone", "o", "", "Phone")
	createAddressCmd.MarkFlagRequired("name")
	createAddressCmd.MarkFlagRequired("country")
	createAddressCmd.MarkFlagRequired("province")
	createAddressCmd.MarkFlagRequired("city")
	createAddressCmd.MarkFlagRequired("zip")
	createAddressCmd.MarkFlagRequired("street1")
	createAddressCmd.MarkFlagRequired("phone")
	addressCmd.AddCommand(deleteAddressCmd)
	rootCmd.AddCommand(addressCmd)
}
