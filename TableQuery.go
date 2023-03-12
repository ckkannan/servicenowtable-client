package servicenowtable

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DataRow struct {
	TableName string `json:"tablename"`
	Name      string `json:"name"`
}

func (c *Client) GetRows() (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/now/table/%s&sysparam_query=%s&sysparm_display_value=true", c.sn_url, c.Table, c.Query), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var datarows map[string]interface{}
	err = json.Unmarshal([]byte(body), &datarows)
	if err != nil {
		return nil, err
	}
	return datarows, nil

}
