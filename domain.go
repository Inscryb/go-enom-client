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

// TldList Retrieve a list of the TLDs that you offer.
func (s Session) TldList() (response.TldList, error) {
	resp := response.TldList{}

	cmd := s.CreateCommand("gettldlist")

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

// TldDetails Retrieve TLD characteristics in detail for a specified TLD.
func (s Session) TldDetails(tld string) (response.TldDetails, error) {
	resp := response.TldDetails{}

	cmd := s.CreateCommand("GetTLDDetails")
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

// NameSpinner List related domain names
func (s Session) NameSpinner(request requests.DomainNameSpinnerRequest) (response.SpinnerResponse, error) {
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

// Purchase purchase a domain name
func (s Session) Purchase(domain string) (response.DomainPurchase, error) {
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

// Renew Extend (renew) the registration period for a domain name
func (s Session) Renew(domain string, period string) (response.Extend, error) {
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

// Delete delete domain with in 5 days of purchase (Will not working on all accounts)
func (s Session) Delete(domain string) (response.Delete, error) {
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

// SetRegLock Set the registrar lock for a domain name
func (s Session) SetRegLock(domain string, enable bool) (response.DomainInfo, error) {
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

// Info Get information about a single domain name
func (s Session) Info(domain string) (response.DomainInfo, error) {
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
