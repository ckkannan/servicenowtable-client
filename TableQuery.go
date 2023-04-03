package servicenowtable_client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type resultRows struct {
	Result []map[string]interface{} `json:"result"`
}

func (c *Client) GetRows() ([]map[string]interface{}, error) {
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
	// JSONMap := make(map[string]interface{})

	// fmt.Println(url)
	var JSONMap resultRows
	// r := make(map[string]interface{})

	err = json.Unmarshal([]byte(body), &JSONMap)
	if err != nil {
		fmt.Println("Failed Json")
		return nil, err
	}
	returnRows := make([]map[string]interface{}, len(JSONMap.Result))
	for k := range JSONMap.Result {
		returnRows[k] = make(map[string]interface{})
		for k1, v1 := range JSONMap.Result[k] {

			returnRows[k][k1] = v1
		}
		// 	append(returnRows, datarows.Result[i])
		// fmt.Println(datarows.Result[i]["sys_id"])
	}
	return returnRows, nil

}
