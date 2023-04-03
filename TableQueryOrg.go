package servicenowtable_client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type result struct {
	Results []rowOrg `json:"result"`
}

type rowOrg struct {
	Sys_id      string `json:"sys_id"`
	To_adgroup  string `json:"to_adgroup"`
	To_org_name string `json:"to_org_name"`
	To_org_type string `json:"to_org_type"`
}

func (c *Client) GetOrgRows() ([]rowOrg, error) {
	url := fmt.Sprintf("%s/api/now/table/%s?sysparm_query=%s&sysparm_display_value=true&sysparm_fields=%s", c.sn_url, c.Table, c.Query, c.Fields)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// fmt.Println("Failed NewRequest")
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		// fmt.Println("Failed dorequest url", url)
		return nil, err
	}
	var r result

	//rowOrgs := []rowOrg{}
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		// fmt.Println("Failed Json", err.Error())
		return nil, err
	}
	// fmt.Println(r.Results[0].To_adgroup)
	// fmt.Printf("datarows: %v", r.Results)
	return r.Results, nil

}
