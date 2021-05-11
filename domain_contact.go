package enom

import (
	"encoding/xml"
	"github.com/Inscryb/go-enom-client/response"
)

// WhoisContact Get Whois contact information for a domain name
func (s Session) WhoisContact(domain string) (response.WhoisContact, error) {
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

// Contacts Get all contact data for a domain name
func (s Session) Contacts(domain string) (response.Contacts, error) {
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

// SetContact Update contact information for a domain name
func (s Session) SetContact(domain string, contactData map[string]string) (response.Contacts, error) {
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
