package enom

import (
	"encoding/xml"
	"github.com/Inscryb/go-enom-client/response"
)

// CreateTransferOrder Transfer domains into an account
func (s Session) CreateTransferOrder(domain string, authInfo string) (response.CreateTransferOrder, error) {
	resp := response.CreateTransferOrder{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("TP_CreateOrder")
	cmd.AddParam("SLD1", sld)
	cmd.AddParam("TLD1", tld)
	cmd.AddParam("AuthInfo1", authInfo)
	cmd.AddParam("DomainCount", "1")
	cmd.AddParam("OrderType", "Autoverification")

	client := Client{&s}
	data, err := client.DoRequest(cmd)
	if err != nil {
		return resp, err
	}

	if err = xml.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

// TransferOrderStatuses Get a list of orders including Closed for the last 6 months
func (s Session) TransferOrderStatuses() (response.TransferOrderStatuses, error) {
	resp := response.TransferOrderStatuses{}

	cmd := s.CreateCommand("TP_GetOrderStatuses")

	client := Client{&s}
	data, err := client.DoRequest(cmd)
	if err != nil {
		return resp, err
	}

	if err = xml.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

// TransferOrderReview Retrieve information on a transfer order
func (s Session) TransferOrderReview(id string) (response.TransferOrderReview, error) {
	resp := response.TransferOrderReview{}

	cmd := s.CreateCommand("TP_GetOrderReview")
	cmd.AddParam("TransferOrderID", id)

	client := Client{&s}
	data, err := client.DoRequest(cmd)
	if err != nil {
		return resp, err
	}

	if err = xml.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

// TransferOrderDetail Get information for a single domain on a transfer order
func (s Session) TransferOrderDetail(id string) (response.TransferOrderDetail, error) {
	resp := response.TransferOrderDetail{}

	cmd := s.CreateCommand("TP_GetOrderDetail")
	cmd.AddParam("TransferOrderDetailID", id)

	client := Client{&s}
	data, err := client.DoRequest(cmd)
	if err != nil {
		return resp, err
	}

	if err = xml.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
