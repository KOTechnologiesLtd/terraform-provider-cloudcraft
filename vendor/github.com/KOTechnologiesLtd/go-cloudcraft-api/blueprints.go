package cloudcraft

import (
	"fmt"
)

type DataDetails struct {
	Grid    *string `json:"grid,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *int    `json:"version,omitempty"`
}

// BluePrintInfo struct represents the data of a BluePrint.
type BluePrint struct {
	ID         *string      `json:"id,omitempty"`
	LastUserID *string      `json:"LastUserId,omitempty"`
	CreatorID  *string      `json:"CreatorId,omitempty"`
	Data       *DataDetails `json:"data,omitempty"`
	CreatedAt  *string      `json:"createdAt,omitempty"`
	UpdatedAt  *string      `json:"updatedAt,omitempty"`
}

// BluePrintCreate creates a BluePrint.
func (client *Client) BluePrintCreate(bp *BluePrint) error {
	//log.Printf("[DEBUG] BluePrint Name %s", *bp.Data.Name)
	//log.Printf("[DEBUG] BluePrint Grid %s", *bp.Data.Grid)
	return client.RequestResponse("POST", "/blueprint", bp, &bp)
}

// BluePrintUpdate updates the BluePrint.
func (client *Client) BluePrintUpdate(bp *BluePrint) error {
	return client.RequestResponse("PUT", fmt.Sprintf("/blueprint/%v", *bp.ID), bp, &bp)
}

// BluePrintDelete deletes the BluePrint
func (client *Client) BluePrintDelete(bp *BluePrint) error {
	return client.RequestResponse("DELETE", fmt.Sprintf("/blueprint/%v", *bp.ID), nil, nil)
}

// BluePrintGet retrieves the BluePrint
func (client *Client) BluePrintGet(bp *BluePrint) error {
	return client.RequestResponse("GET", fmt.Sprintf("/blueprint/%v", *bp.ID), bp, &bp)
}
