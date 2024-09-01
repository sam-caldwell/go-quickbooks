package main

import (
	"flag"
	"github.com/sam-caldwell/ansi"
	quickbooks "github.com/sam-caldwell/go-quickbooks"
)

func main() {
	var err error
	var client *quickbooks.Client
	var account quickbooks.Account

	// Define command-line flags
	accountDescriptorFile := flag.String("file", "", "YAML descriptor file")
	accessToken := flag.String("accessToken", "", "AccessToken")
	companyId := flag.String("companyId", "", "CompanyId")
	developmentMode := flag.Bool("developmentMode", false, "Use for development only")
	flag.Parse()

	// Validate required flags
	if *accountDescriptorFile == "" || *accessToken == "" || *companyId == "" {
		flag.Usage()
		ansi.Red().Println("Missing required flags: file, accessToken, companyId").Fatal(1).Reset()
		return
	}

	// Deserialize the account information from the YAML file
	if err = account.DeserializeYAML(*accountDescriptorFile); err != nil {
		ansi.Red().Println(err.Error()).Fatal(1).Reset()
		return
	}

	// Create a new QuickBooks client
	if client, err = quickbooks.NewClient(*companyId, *accessToken, !*developmentMode); err != nil {
		ansi.Red().Printf("Error creating QBO client: %v\n", err).Fatal(1).Reset()
		return
	}

	// Create the account in QuickBooks
	if _, err = account.Create(client); err != nil {
		ansi.Red().Printf("Error creating account: %v\n", err).Fatal(1).Reset()
		return
	}

	// Success message
	ansi.Green().Println("Account created successfully!").Reset()
}
