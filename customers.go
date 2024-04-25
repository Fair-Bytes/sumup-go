// Code generated by `gogenitor`. DO NOT EDIT.
package sumup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Address is Profile's personal address information.
type Address struct {
	// City name from the address.
	City *string `json:"city,omitempty"`
	// Two letter country code formatted according to [ISO3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// First line of the address with details of the street name and number.
	Line1 *string `json:"line1,omitempty"`
	// Second line of the address with details of the building, unit, apartment, and floor numbers.
	Line2 *string `json:"line2,omitempty"`
	// Postal code from the address.
	PostalCode *string `json:"postal_code,omitempty"`
	// State name or abbreviation from the address.
	State *string `json:"state,omitempty"`
}

// Customer is the type definition for a Customer.
type Customer struct {
	// Unique ID of the customer.
	CustomerId string `json:"customer_id"`
	// Personal details for the customer.
	PersonalDetails *PersonalDetails `json:"personal_details,omitempty"`
}

// PaymentInstrumentCard is Details of the payment card that is saved as a payment instrument.
type PaymentInstrumentCard struct {
	// Indicates whether the payment instrument is active and can be used for payments. To deactivate it, send a `DELETE` request to the resource endpoint.
	Active bool `json:"active"`
	// __Required when payment type is `card`.__ Details of the payment card.
	Card Card `json:"card"`
	// Unique token identifying the saved payment card for a customer.
	Token string `json:"token"`
	// Type of the payment instrument.
	Type PaymentInstrumentCardType `json:"type"`
}

// Type of the payment instrument.
type PaymentInstrumentCardType string

const (
	PaymentInstrumentCardTypeCard PaymentInstrumentCardType = "card"
)

// PaymentInstrumentResponse is Payment Instrument Response
type PaymentInstrumentResponse struct {
	// Indicates whether the payment instrument is active and can be used for payments. To deactivate it, send a `DELETE` request to the resource endpoint.
	Active *bool `json:"active,omitempty"`
	// Details of the payment card.
	Card *PaymentInstrumentResponseCard `json:"card,omitempty"`
	// Creation date of payment instrument. Response format expressed according to [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) code.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Created mandate
	Mandate *MandateResponse `json:"mandate,omitempty"`
	// Unique token identifying the saved payment card for a customer.
	Token *string `json:"token,omitempty"`
	// Type of the payment instrument.
	Type *PaymentInstrumentResponseType `json:"type,omitempty"`
}

// PaymentInstrumentResponseCard is Details of the payment card.
type PaymentInstrumentResponseCard struct {
	// Last 4 digits of the payment card number.
	Last4Digits *string `json:"last_4_digits,omitempty"`
	// Issuing card network of the payment card.
	Type *PaymentInstrumentResponseCardType `json:"type,omitempty"`
}

// Issuing card network of the payment card.
type PaymentInstrumentResponseCardType string

const (
	PaymentInstrumentResponseCardTypeAmex         PaymentInstrumentResponseCardType = "AMEX"
	PaymentInstrumentResponseCardTypeCup          PaymentInstrumentResponseCardType = "CUP"
	PaymentInstrumentResponseCardTypeDiners       PaymentInstrumentResponseCardType = "DINERS"
	PaymentInstrumentResponseCardTypeDiscover     PaymentInstrumentResponseCardType = "DISCOVER"
	PaymentInstrumentResponseCardTypeElo          PaymentInstrumentResponseCardType = "ELO"
	PaymentInstrumentResponseCardTypeElv          PaymentInstrumentResponseCardType = "ELV"
	PaymentInstrumentResponseCardTypeHipercard    PaymentInstrumentResponseCardType = "HIPERCARD"
	PaymentInstrumentResponseCardTypeJcb          PaymentInstrumentResponseCardType = "JCB"
	PaymentInstrumentResponseCardTypeMaestro      PaymentInstrumentResponseCardType = "MAESTRO"
	PaymentInstrumentResponseCardTypeMastercard   PaymentInstrumentResponseCardType = "MASTERCARD"
	PaymentInstrumentResponseCardTypeUnknown      PaymentInstrumentResponseCardType = "UNKNOWN"
	PaymentInstrumentResponseCardTypeVisa         PaymentInstrumentResponseCardType = "VISA"
	PaymentInstrumentResponseCardTypeVisaElectron PaymentInstrumentResponseCardType = "VISA_ELECTRON"
	PaymentInstrumentResponseCardTypeVisaVpay     PaymentInstrumentResponseCardType = "VISA_VPAY"
)

// Type of the payment instrument.
type PaymentInstrumentResponseType string

const (
	PaymentInstrumentResponseTypeCard PaymentInstrumentResponseType = "card"
)

// PersonalDetails is Personal details for the customer.
type PersonalDetails struct {
	// Profile's personal address information.
	Address *Address `json:"address,omitempty"`
	// Date of birth of the customer.
	Birthdate *time.Time `json:"birthdate,omitempty"`
	// Email address of the customer.
	Email *string `json:"email,omitempty"`
	// First name of the customer.
	FirstName *string `json:"first_name,omitempty"`
	// Last name of the customer.
	LastName *string `json:"last_name,omitempty"`
	// Phone number of the customer.
	Phone *string `json:"phone,omitempty"`
}

// CreateCustomer request body.
type CreateCustomerBody struct {
	// Unique ID of the customer.
	CustomerId string `json:"customer_id"`
	// Personal details for the customer.
	PersonalDetails *PersonalDetails `json:"personal_details,omitempty"`
}

// ListPaymentInstrumentsResponse is the type definition for a ListPaymentInstrumentsResponse.
type ListPaymentInstrumentsResponse []PaymentInstrumentResponse

// CreatePaymentInstrument request body.
type CreatePaymentInstrumentBody struct {
	// Indicates whether the payment instrument is active and can be used for payments. To deactivate it, send a `DELETE` request to the resource endpoint.
	Active bool `json:"active"`
	// __Required when payment type is `card`.__ Details of the payment card.
	Card Card `json:"card"`
	// Unique token identifying the saved payment card for a customer.
	Token string `json:"token"`
	// Type of the payment instrument.
	Type CreatePaymentInstrumentBodyType `json:"type"`
}

// Type of the payment instrument.
type CreatePaymentInstrumentBodyType string

const (
	CreatePaymentInstrumentBodyTypeCard CreatePaymentInstrumentBodyType = "card"
)

// UpdateCustomer request body.
type UpdateCustomerBody struct {
	// Personal details for the customer.
	PersonalDetails *PersonalDetails `json:"personal_details,omitempty"`
}

// DeactivatePaymentInstrumentResponse is the type definition for a DeactivatePaymentInstrumentResponse.
type DeactivatePaymentInstrumentResponse struct {
}

type CustomersService service

// Create: Create a customer
// Creates a new saved customer resource which you can later manipulate and save payment instruments to.
func (s *CustomersService) Create(ctx context.Context, body CreateCustomerBody) (*Customer, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, fmt.Errorf("encoding json body request failed: %v", err)
	}

	path := fmt.Sprintf("/v0.1/customers")

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, buf)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("invalid response: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := dec.Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("read error response: %s", err.Error())
		}

		return nil, &apiErr
	}

	var v Customer
	if err := dec.Decode(&v); err != nil {
		return nil, fmt.Errorf("decode response: %s", err.Error())
	}

	return &v, nil
}

// ListPaymentInstruments: List payment instruments
// Lists all payment instrument resources that are saved for an identified customer.
func (s *CustomersService) ListPaymentInstruments(ctx context.Context, customerId string) (*ListPaymentInstrumentsResponse, error) {
	path := fmt.Sprintf("/v0.1/customers/%v/payment-instruments", customerId)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("invalid response: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := dec.Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("read error response: %s", err.Error())
		}

		return nil, &apiErr
	}

	var v ListPaymentInstrumentsResponse
	if err := dec.Decode(&v); err != nil {
		return nil, fmt.Errorf("decode response: %s", err.Error())
	}

	return &v, nil
}

// CreatePaymentInstrument: Create a payment instrument
// Creates and activates a new payment instrument resource by saving a payment card for an identified customer. Implement to improve customer experience by skipping the step of entering payment instrument details.
//
// The token created via this endpoint **can not** be used for recurring payments by merchants operating within the EU. For more information visit our [recurring payments guide](https://developer.sumup.com/docs/recurring-payments/).
func (s *CustomersService) CreatePaymentInstrument(ctx context.Context, customerId string, body CreatePaymentInstrumentBody) (*PaymentInstrumentResponse, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, fmt.Errorf("encoding json body request failed: %v", err)
	}

	path := fmt.Sprintf("/v0.1/customers/%v/payment-instruments", customerId)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, buf)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("invalid response: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := dec.Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("read error response: %s", err.Error())
		}

		return nil, &apiErr
	}

	var v PaymentInstrumentResponse
	if err := dec.Decode(&v); err != nil {
		return nil, fmt.Errorf("decode response: %s", err.Error())
	}

	return &v, nil
}

// Get: Retrieve a customer
// Retrieves an identified saved customer resource through the unique `customer_id` parameter, generated upon customer creation.
func (s *CustomersService) Get(ctx context.Context, customerId string) (*Customer, error) {
	path := fmt.Sprintf("/v0.1/customers/%v", customerId)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("invalid response: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := dec.Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("read error response: %s", err.Error())
		}

		return nil, &apiErr
	}

	var v Customer
	if err := dec.Decode(&v); err != nil {
		return nil, fmt.Errorf("decode response: %s", err.Error())
	}

	return &v, nil
}

// Update: Update a customer
// Updates an identified saved customer resource's personal details.
//
// The request only overwrites the parameters included in the request, all other parameters will remain with their initially assigned values.
func (s *CustomersService) Update(ctx context.Context, customerId string, body UpdateCustomerBody) (*Customer, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, fmt.Errorf("encoding json body request failed: %v", err)
	}

	path := fmt.Sprintf("/v0.1/customers/%v", customerId)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, buf)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("invalid response: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := dec.Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("read error response: %s", err.Error())
		}

		return nil, &apiErr
	}

	var v Customer
	if err := dec.Decode(&v); err != nil {
		return nil, fmt.Errorf("decode response: %s", err.Error())
	}

	return &v, nil
}

// DeactivatePaymentInstrument: Deactivate a payment instrument
// Deactivates an identified card payment instrument resource for a customer.
func (s *CustomersService) DeactivatePaymentInstrument(ctx context.Context, customerId string, token string) (*DeactivatePaymentInstrumentResponse, error) {
	path := fmt.Sprintf("/v0.1/customers/%v/payment-instruments/%v", customerId, token)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("invalid response: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := dec.Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("read error response: %s", err.Error())
		}

		return nil, &apiErr
	}

	var v DeactivatePaymentInstrumentResponse
	if err := dec.Decode(&v); err != nil {
		return nil, fmt.Errorf("decode response: %s", err.Error())
	}

	return &v, nil
}
