package main

import (
	"flag"
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/go-quickbooks"
)

func main() {
	var err error
	var client *quickbooks.Client

	// Define command-line flags
	accountId := flag.String("id", "", "Account ID to delete")
	accessToken := flag.String("accessToken", "", "AccessToken")
	companyId := flag.String("companyId", "", "CompanyId")
	developmentMode := flag.Bool("developmentMode", false, "Use for development only")
	flag.Parse()

	// Validate required flags
	if *accountId == "" || *accessToken == "" || *companyId == "" {
		flag.Usage()
		ansi.Red().Println("Missing required flags: id, accessToken, companyId").Fatal(1)
		return
	}

	// Create a new QuickBooks client
	if client, err = quickbooks.NewClient(*companyId, *accessToken, !*developmentMode); err != nil {
		ansi.Red().Printf("Error creating QBO client: %v\n", err).Fatal(1)
		return
	}

	// Delete the account in QuickBooks
	var account quickbooks.Account
	if err = account.Delete(client, *accountId); err != nil {
		ansi.Red().Printf("Error deleting account: %v\n", err).Fatal(1)
		return
	}

	// Success message
	ansi.Green().Printf("Account ID %s deleted successfully!\n", *accountId).Reset()
}
