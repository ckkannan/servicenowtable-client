package servicenowtable_client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetRows() (map[string]interface{}, error) {
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
	fmt.Println(url)
	var datarows map[string]interface{}
	err = json.Unmarshal([]byte(body), &datarows)
	if err != nil {
		fmt.Println("Failed Json")
		return nil, err
	}
	return datarows, nil

}
