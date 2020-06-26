package accountgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var accountServiceHost string = os.Getenv("ACCOUNT_SERVICE_HOST")

// AccountInfoResponse ...
type AccountInfoResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetAccountInfo ...
func GetAccountInfo(accountID string) (AccountInfoResponse, error) {
	endpoint := fmt.Sprintf("/accounts/%s", accountID)
	url := fmt.Sprintf("%s%s", accountServiceHost, endpoint)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return AccountInfoResponse{}, err
	}

	defer response.Body.Close()

	accountInfo := AccountInfoResponse{}

	if err := json.NewDecoder(response.Body).Decode(&accountInfo); err != nil {
		return AccountInfoResponse{}, err
	}

	return accountInfo, nil
}
