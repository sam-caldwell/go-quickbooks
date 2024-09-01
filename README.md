Quickbooks API package
======================

## Description

An API Package for interacting with Quickbooks along with some
utility commands.

## Usage

### `qb-account-create`

This command will read the account details from account.yaml, authenticate with QuickBooks using the provided
access token, and create the account in the specified QuickBooks company. If developmentMode is set to true,
it will use the sandbox environment.

```shell
qb-account-create -file=account.yaml -accessToken=yourAccessToken -companyId=yourCompanyId -developmentMode=true
```

### `qb-account-describe`

```shell
qb-account-describe -id=${ACCT_ID} -accessToken=yourAccessToken -companyId=yourCompanyId -developmentMode=true
```

### `qb-account-list`

```shell
qb-account-list -offset ${offset} -limit ${limit} -accessToken=yourAccessToken -companyId=yourCompanyId \
                -developmentMode=true
```

### `qb-account-delete`

```shell
qb-account-create -id=${ACCT_ID} -accessToken=yourAccessToken -companyId=yourCompanyId -developmentMode=true
```