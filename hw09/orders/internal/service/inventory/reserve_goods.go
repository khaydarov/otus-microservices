package inventory

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// ReserveGoods sends request inventory service to reserve goods
func ReserveGoods(orderId string, goodIds []int) ([]int, error) {
	endpoint := fmt.Sprintf("%s/reserveGoods", os.Getenv("INVENTORY_HOST"))
	data := map[string]interface{}{
		"order_id": orderId,
		"good_ids": goodIds,
	}

	body, _ := json.Marshal(data)

	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("internal service while reserving goods")
	}

	return goodIds, nil
}
