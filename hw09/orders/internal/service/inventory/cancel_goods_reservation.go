package inventory

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// CancelGoodsReservation sends request to cancel reservation
func CancelGoodsReservation(orderId string) error {
	endpoint := fmt.Sprintf("%s/cancelGoodsReservation", os.Getenv("INVENTORY_HOST"))
	data := map[string]interface{}{
		"order_id": orderId,
	}

	body, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("internal service while canceling reservation goods")
	}

	return nil
}
