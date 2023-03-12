package servicenowtable_client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type rowOrg struct {
	Sys_id      string `json:"sys_id"`
	To_adgroup  string `json:"to_adgroup"`
	To_org_name string `json:"to_org_name"`
	To_org_type string `json:"to_org_type"`
}

func (c *Client) GetOrgRows() ([]rowOrg, error) {
	url := fmt.Sprintf("%s/api/now/table/%s?sysparm_query=%s&sysparm_display_value=true", c.sn_url, c.Table, c.Query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed NewRequest")
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		fmt.Println("Failed dorequest url", url)
		return nil, err
	}
	// fmt.Println(url)
	rowOrgs := []rowOrg{}
	err = json.Unmarshal([]byte(body), &rowOrgs)
	if err != nil {
		fmt.Println("Failed Json")
		return nil, err
	}
	return rowOrgs, nil

}
