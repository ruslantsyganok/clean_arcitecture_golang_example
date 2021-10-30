package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"zen_api/internal/datastruct"
)

type PaymentService interface {
	NewTransaction(userID, courseID int64) error
}

type paymentService struct {
	secret string
}

func NewPaymentService(secret string) PaymentService {
	return &paymentService{secret: secret}
}

func (p *paymentService) NewTransaction(userID, courseID int64) error {
	log.Println("started transaction")

	// get course price by courseID
	coursePrice := int64(1)

	// UNIQUE TRANSACTION ID
	id := uuid.NewV4()
	transactionId := id.String()

	url := "https://api.qiwi.com/partner/bill/v1/bills/" + transactionId
	expirationDateTime := time.Now().Add(time.Hour)

	paymentReq := datastruct.Payment{
		ExpirationDateTime: expirationDateTime.Format("2006-01-02T15:04:05-07:00"),
		Amount: datastruct.Amount{
			Currency: "RUB",
			Value:    strconv.Itoa(int(coursePrice)),
		},
		CustomFields: datastruct.CustomFields{
			YourParam1: strconv.Itoa(int(userID)),
			YourParam2: strconv.Itoa(int(courseID)),
		},
	}

	bs, err := json.Marshal(paymentReq)
	if err != nil {
		return fmt.Errorf("cannot marshal payment request: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bs))
	if err != nil {
		return fmt.Errorf("cannot make a request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+p.secret)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	log.Println("request: ", req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot make a HTTP dial: %v", err)
	}

	bs, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read from a body: %v", err)
	}

	log.Println("response: ", string(bs))

	var paymentInfo datastruct.Payment
	err = json.Unmarshal(bs, &paymentInfo)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %v", err)
	}
	log.Println("finished transaction")

	return nil
	//http.Redirect(w, r, paymentInfo.PayURL, http.StatusSeeOther)
}
