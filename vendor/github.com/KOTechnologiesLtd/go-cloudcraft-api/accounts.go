package cloudcraft

import (
	"fmt"
)

// AccountIntegrationAws struct represents the data of an Account.
type AccountIntegrationAws struct {
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	RoleArn    *string `json:"roleArn,omitempty"`
	ExternalID *string `json:"externalId,omitempty"`
	CreatedAt  *string `json:"createdAt,omitempty"`
	UpdatedAt  *string `json:"updatedAt,omitempty"`
	CreatorID  *string `json:"CreatorId,omitempty"`
}

// AccountsAwsInfoResp contains data of all Account.
type AccountsAwsInfoResp struct {
	AccountsAwsInfo []AccountIntegrationAws `json:"accounts,omitempty"`
}

// AccountIntegrationsAllAws get and return the Aws Accounts.
func (client *Client) AccountIntegrationsAllAws() ([]AccountIntegrationAws, error) {
	var out AccountsAwsInfoResp
	err := client.RequestResponse("GET", "/aws/account", nil, &out)
	{
		if err != nil {
			return nil, err
		}
	}
	if len(out.AccountsAwsInfo) == 0 {
		return nil, fmt.Errorf("Cloudcraft AWS Account Integrations not found or there are none")
	}
	return out.AccountsAwsInfo, nil
}

// AccountIntegrationAws get and return the Aws Account Information.
func (client *Client) AccountIntegrationAws(AwsAccountId string) (AccountIntegrationAws, error) {
	accounts, err := client.AccountIntegrationsAllAws()
	{
		if err != nil {
			return AccountIntegrationAws{}, err
		}
	}
	//log.Printf("accounts%v - length%d", accounts, len(accounts))
	AccountIntegrationAwsItem := AccountIntegrationAws{}
	for i := range accounts {

		if *accounts[i].ID == AwsAccountId {
			AccountIntegrationAwsItem = accounts[i]
			//ID := *accounts[i].ID
			//Name := *accounts[i].Name
			//log.Printf("ID %s, Name %s", ID, Name)
			return AccountIntegrationAwsItem, nil
		}
	}
	if (AccountIntegrationAws{}) == AccountIntegrationAwsItem {
		return AccountIntegrationAwsItem, fmt.Errorf("Cloudcraft AWS Account Integration not found")
	}
	return AccountIntegrationAwsItem, nil
}

// AccountIntegrationAwsCreate updates AWS Account Integration.
func (client *Client) AccountIntegrationAwsCreate(acc *AccountIntegrationAws) error {

	return client.RequestResponse("POST", "/aws/account", acc, &acc)
}

// AccountIntegrationAwsUpdate updates AWS Account Integrations.
func (client *Client) AccountIntegrationAwsUpdate(acc *AccountIntegrationAws) error {
	return client.RequestResponse("PUT", fmt.Sprintf("/aws/account/%v", *acc.ID), acc, &acc)
}

// AccountAwsDelAccountIntegrationAwsDeleteete updates AWS Account Integrations.
func (client *Client) AccountIntegrationAwsDelete(acc *AccountIntegrationAws) error {
	return client.RequestResponse("DELETE", fmt.Sprintf("/aws/account/%v", *acc.ID), nil, nil)
}
