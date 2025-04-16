package midtrans

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/models"
)

func CreateOrder(orderId int, amount int64) (res string) {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVERKEY")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENTKEY")

	// midtrans.ServerKey = "SB-Mid-server-yuXIcN-TJok6vvxxhot7C4Vl"
	// midtrans.ClientKey = "SB-Mid-client-GWXKp9r1IWAcmHGp"

	midtrans.Environment = midtrans.Sandbox

	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(orderId),
			GrossAmt: amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, _ := snap.CreateTransaction(req)
	fmt.Println("Response :", snapResp)

	res = snapResp.RedirectURL
	return
}

func CheckMidtransPaymentStatus(orderID string) (*models.PaymentMidtrans, error) {
	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/status", orderID)
	req, _ := http.NewRequest("GET", url, nil)

	// Replace with your Midtrans server key
	serverKey := os.Getenv("MIDTRANS_SERVERKEY")
	auth := base64.StdEncoding.EncodeToString([]byte(serverKey + ":"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+auth)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	fmt.Println("Midtrans response:", response)

	getString := func(key string) (string, error) {
		if val, ok := response[key].(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("field %s not found or not a string", key)
	}

	transactionID, err := getString("transaction_id")
	if err != nil {
		return nil, err
	}

	currency, err := getString("currency")
	if err != nil {
		return nil, err
	}

	paymentType, err := getString("payment_type")
	if err != nil {
		return nil, err
	}

	transactionStatus, err := getString("transaction_status")
	if err != nil {
		return nil, err
	}

	statusMessage, err := getString("status_message")
	if err != nil {
		return nil, err
	}

	merchantID, err := getString("merchant_id")
	if err != nil {
		return nil, err
	}

	grossAmount, err := strconv.ParseFloat(response["gross_amount"].(string), 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse gross_amount: %w", err)
	}

	transactionTimeStr := response["transaction_time"].(string)
	transactionTime, err := time.Parse("2006-01-02 15:04:05", transactionTimeStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse transaction_time: %w", err)
	}

	payMT := models.PaymentMidtrans{
		TransactionId:     transactionID,
		OrderId:           orderID,
		GrossAmount:       int64(grossAmount),
		Currency:          currency,
		PaymentType:       paymentType,
		TransactionStatus: transactionStatus,
		StatusMessage:     statusMessage,
		MerchantId:        merchantID,
		TransactionTime:   transactionTime,
	}

	return &payMT, nil
}
