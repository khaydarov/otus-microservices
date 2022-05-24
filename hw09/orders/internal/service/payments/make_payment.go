package payments

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// MakePayment sends request to payment service to make payment
func MakePayment(ctx context.Context, orderId string, amount int) error {
	endpoint := fmt.Sprintf("%s/makePayment", os.Getenv("PAYMENTS_HOST"))
	data := map[string]interface{}{
		"order_id": orderId,
		"amount": amount,
	}

	body, _ := json.Marshal(data)

	request, _ := http.NewRequest( "POST", endpoint, bytes.NewBuffer(body))
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
