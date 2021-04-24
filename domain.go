package enom

import (
	"encoding/xml"
	"github.com/Inscryb/go-enom-client/requests"
	"github.com/Inscryb/go-enom-client/response"
	"strings"
)

func ParseDomain(domain string) (string, string) {
	fragments := strings.Split(domain, ".")
	slice := fragments[0:2]

	return slice[0], slice[1]
}

// DomainCheck Check the availability of a domain name
func (s Session) DomainCheck(domain string) (response.DomainCheck, error) {
	resp := response.DomainCheck{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("check")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)
	cmd.AddParam("Version", "2")
	cmd.AddParam("IncludePrice", "true")

	client := Client{&s}
	data, err := client.DoRequest(cmd)
	if err != nil {
		return resp, err
	}

	type rDomain struct {
		Domain response.DomainCheck `xml:"Domain"`
	}
	type domains struct {
		response.Response
		Domains []rDomain `xml:"Domains"`
	}

	r := domains{}
	if err = xml.Unmarshal(data, &r); err != nil {
		return resp, err
	}

	return r.Domains[0].Domain, nil
}

// DomainNameSpinner List related domain names
func (s Session) DomainNameSpinner(request requests.DomainNameSpinnerRequest) (response.SpinnerResponse, error) {
	resp := response.SpinnerResponse{}

	sld, tld := ParseDomain(request.Domain)

	cmd := s.CreateCommand("NameSpinner")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)
	cmd.AddParam("tldlist", request.TLDList)

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

// DomainPurchase purchase a domain name
func (s Session) DomainPurchase(domain string) (response.DomainPurchase, error) {
	resp := response.DomainPurchase{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("purchase")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

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

// DomainCreateTransferOrder Transfer domains into an account
func (s Session) DomainCreateTransferOrder(domain string, authInfo string) (response.CreateTransferOrder, error) {
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

// DomainWhoisContact Get Whois contact information for a domain name
func (s Session) DomainWhoisContact(domain string) (response.WhoisContact, error) {
	resp := response.WhoisContact{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("GetWhoisContact")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

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

// DomainRenew Extend (renew) the registration period for a domain name
func (s Session) DomainRenew(domain string, period string) (response.Extend, error) {
	resp := response.Extend{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("Extend")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)
	cmd.AddParam("NumYears", period)

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

// DomainDelete delete domain with in 5 days of purchase (Will not working on all accounts)
func (s Session) DomainDelete(domain string) (response.Delete, error) {
	resp := response.Delete{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("DeleteRegistration")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

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

// DomainSetRegLock Set the registrar lock for a domain name
func (s Session) DomainSetRegLock(domain string, enable bool) (response.DomainInfo, error) {
	resp := response.DomainInfo{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("SetRegLock")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

	var value string
	if enable {
		value = "1"
	} else {
		value = "0"
	}

	cmd.AddParam("UnlockRegistrar", value)

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

// DomainInfo Get information about a single domain name
func (s Session) DomainInfo(domain string) (response.DomainInfo, error) {
	resp := response.DomainInfo{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("GetDomainInfo")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

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

// DomainContacts Get all contact data for a domain name
func (s Session) DomainContacts(domain string) (response.Contacts, error) {
	resp := response.Contacts{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("GetContacts")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

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

// DomainSetContact Update contact information for a domain name
func (s Session) DomainSetContact(domain string, contactData map[string]string) (response.Contacts, error) {
	resp := response.Contacts{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("Contacts")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

	for key, val := range contactData {
		cmd.AddParam(key, val)
	}

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

// DomainRegisterNameserver Register a domain name server
func (s Session) DomainRegisterNameserver(domain string, name string, ip string) (response.RegisterNameserver, error) {
	resp := response.RegisterNameserver{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("RegisterNameserver")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

	cmd.AddParam("Add", "true")
	cmd.AddParam("NsName", name)
	cmd.AddParam("Ip", ip)

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

// DomainUpdateNameserver Change the IP address of a name server in the Registry’s records
func (s Session) DomainUpdateNameserver(domain string, name string, oldIp string, newIp string) (response.UpdateNameserver, error) {
	resp := response.UpdateNameserver{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("UpdateNameserver")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

	cmd.AddParam("NS", name)
	cmd.AddParam("OldIp", oldIp)
	cmd.AddParam("NewIp", newIp)

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

// DomainCheckNSStatus Retrieve information about a name server
func (s Session) DomainCheckNSStatus(domain string, host string) (response.CheckNSStatus, error) {
	resp := response.CheckNSStatus{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("CheckNSStatus")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

	cmd.AddParam("CheckNSName", host)

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

// DomainTransferOrderStatuses Get a list of orders including Closed for the last 6 months
func (s Session) DomainTransferOrderStatuses() (response.TransferOrderStatuses, error) {
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

// DomainTransferOrderReview Retrieve information on a transfer order
func (s Session) DomainTransferOrderReview(id string) (response.TransferOrderReview, error) {
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

// DomainTransferOrderDetail Get information for a single domain on a transfer order
func (s Session) DomainTransferOrderDetail(id string) (response.TransferOrderDetail, error) {
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
