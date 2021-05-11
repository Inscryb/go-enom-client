package enom

import (
	"github.com/Inscryb/go-enom-client/requests"
	"testing"
)

func TestGetHosts(t *testing.T) {
	session := CreateSession(DomainTestApiUrl, DomainTestApiResellerID, DomainTestApiResellerPwd)

	resp, err := session.GetHosts("resellerdocs.com")
	if err != nil {
		t.Error(err)
	}

	if resp.ErrCount > 0 {
		t.Errorf("domain get hosts contains errors")
	}
}

func TestSetHosts(t *testing.T) {
	session := CreateSession(DomainTestApiUrl, DomainTestApiResellerID, DomainTestApiResellerPwd)

	req := requests.DomainSetHostRequest{
		Domain: "inscryb.com",
		Hosts: []requests.DomainSetHost{
			{HostName: "@", RecordType: "A", Address: "66.150.5.189", MXPref: 0},
			{HostName: "photos", RecordType: "CNAME", Address: "photos.msn.com.", MXPref: 0},
			{HostName: "yahoo", RecordType: "URL", Address: "204.71.200.72", MXPref: 0},
			{HostName: "msn", RecordType: "FRAME", Address: "http://www.msn.com", MXPref: 0},
		},
	}
	resp, err := session.SetHosts(req)
	if err != nil {
		t.Error(err)
	}

	if resp.ErrCount > 0 {
		t.Errorf("domain get hosts contains errors")
	}
}
