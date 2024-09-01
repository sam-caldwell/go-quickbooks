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
	offset := flag.Int("offset", 0, "Pagination offset")
	limit := flag.Int("limit", 10, "Number of accounts to list")
	accessToken := flag.String("accessToken", "", "AccessToken")
	companyId := flag.String("companyId", "", "CompanyId")
	developmentMode := flag.Bool("developmentMode", false, "Use for development only")
	flag.Parse()

	// Validate required flags
	if *accessToken == "" || *companyId == "" {
		flag.Usage()
		ansi.Red().Println("Missing required flags: accessToken, companyId").Fatal(1)
		return
	}

	// Create a new QuickBooks client
	if client, err = quickbooks.NewClient(*companyId, *accessToken, !*developmentMode); err != nil {
		ansi.Red().Printf("Error creating QBO client: %v\n", err).Fatal(1)
		return
	}

	// List the accounts in QuickBooks
	var account quickbooks.Account
	accountSet, err := account.List(client, *offset, *limit)
	if err != nil {
		ansi.Red().Printf("Error listing accounts: %v\n", err).Fatal(1)
		return
	}

	// Pretty-print each account in the list
	ansi.Green()
	for _, account := range accountSet.Accounts {
		out, err := account.PrettyPrint()
		if err != nil {
			ansi.Red().
				Printf("Error generating pretty JSON for account: %v\n", err).
				Fatal(1)
			return
		}
		ansi.Print(string(out)).LF()
	}
	ansi.Reset()
}
