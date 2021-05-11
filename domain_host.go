package enom

import (
	"encoding/xml"
	"errors"
	"github.com/Inscryb/go-enom-client/requests"
	"github.com/Inscryb/go-enom-client/response"
	"strconv"
)

// GetHosts Get host records for a domain name.
func (s Session) GetHosts(domain string) (response.GetHosts, error) {
	resp := response.GetHosts{}

	sld, tld := ParseDomain(domain)

	cmd := s.CreateCommand("gethosts")
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

// SetHosts Set host records for a domain name
func (s Session) SetHosts(req requests.DomainSetHostRequest) (response.Response, error) {
	resp := response.Response{}

	sld, tld := ParseDomain(req.Domain)

	cmd := s.CreateCommand("sethosts")
	cmd.AddParam("sld", sld)
	cmd.AddParam("tld", tld)

	if len(req.Hosts) > 50 {
		return resp, errors.New("can not have more then 50 host records")
	}

	for i, host := range req.Hosts {
		num := i + 1
		cmd.AddParam("HostName"+strconv.Itoa(num), host.HostName)
		cmd.AddParam("RecordType"+strconv.Itoa(num), host.RecordType)
		cmd.AddParam("Address"+strconv.Itoa(num), host.Address)
		cmd.AddParam("MXPref"+strconv.Itoa(num), strconv.Itoa(host.MXPref))
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
