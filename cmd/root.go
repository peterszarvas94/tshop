package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
)

var rootCmd = &cobra.Command{
	Use:   "tshop",
	Short: "CLI for terminal.shop",
	Args:  cobra.ExactArgs(0),
}

var Client *terminal.Client

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error running the command")
		fmt.Println(err.Error())
		os.Exit(1)
		os.Exit(1)
	}
}

// these stucts hold the flags
var User = &terminal.ProfileUser{
	Name:  "",
	Email: "",
}

var Address = &terminal.Address{
	Name:     "",
	Country:  "",
	Province: "",
	City:     "",
	Zip:      "",
	Street1:  "",
	Street2:  "",
	Phone:    "",
}

func init() {
	// version
	rootCmd.AddCommand(versionCmd)

	// product
	productCmd.AddCommand(listProductsCmd)
	productCmd.AddCommand(describeProductCmd)
	rootCmd.AddCommand(productCmd)

	// profil
	profilCmd.AddCommand(profilInfoCmd)
	profilCmd.AddCommand(profilUpdateCmd)
	profilUpdateCmd.Flags().StringVarP(&User.Name, "name", "n", "", "Name")
	profilUpdateCmd.Flags().StringVarP(&User.Email, "email", "e", "", "Email")
	rootCmd.AddCommand(profilCmd)

	// address
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

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		token, ok := os.LookupEnv("TERMINAL_TOKEN")
		if !ok || token == "" {
			fmt.Printf("Environment variable \"TERMINAL_TOKEN\" is missing\n")
			os.Exit(1)
		}

		env, ok := os.LookupEnv("TERMINAL_ENV")
		if !ok || env == "" {
			fmt.Printf("Environment variable \"TERMINAL_ENV\" is missing\n")
			os.Exit(1)
		}

		var urlOption option.RequestOption
		switch env {
		case ("dev"):
			urlOption = option.WithEnvironmentDev()
			break
		case ("prod"):
			urlOption = option.WithEnvironmentProduction()
			break
		default:
			fmt.Println("Invalid environment variable \"TERMINAL_ENV\", should be \"dev\" or \"prod\"")
			os.Exit(1)
		}

		Client = terminal.NewClient(
			urlOption,
			option.WithBearerToken(token),
		)
	}
}
