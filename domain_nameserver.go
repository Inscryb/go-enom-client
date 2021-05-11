package enom

import (
	"encoding/xml"
	"github.com/Inscryb/go-enom-client/response"
)

// RegisterNameserver Register a domain name server
func (s Session) RegisterNameserver(domain string, name string, ip string) (response.RegisterNameserver, error) {
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

// UpdateNameserver Change the IP address of a name server in the Registry’s records
func (s Session) UpdateNameserver(domain string, name string, oldIp string, newIp string) (response.UpdateNameserver, error) {
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

// CheckNSStatus Retrieve information about a name server
func (s Session) CheckNSStatus(domain string, host string) (response.CheckNSStatus, error) {
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
