package shipment

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// CancelCourierReservation sends request to shipment service to cancel courier reservation
func CancelCourierReservation(orderId string) error {
	endpoint := fmt.Sprintf("%s/cancelCourierReservation", os.Getenv("SHIPMENT_HOST"))
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
		return errors.New("internal service error")
	}

	return nil
}
