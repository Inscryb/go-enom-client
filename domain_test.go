package enom

import (
	"fmt"
	"github.com/Inscryb/go-enom-client/requests"
	"testing"
)

var (
	DomainTestApiUrl         = "https://resellertest.enom.com"
	DomainTestApiResellerID  = "resellid"
	DomainTestApiResellerPwd = "resellpw"
)

func TestDomainCheck(t *testing.T) {
	session := CreateSession(DomainTestApiUrl, DomainTestApiResellerID, DomainTestApiResellerPwd)

	domainCheck, err := session.DomainCheck("inscryb.com")
	if err != nil {
		t.Error(err)
	}

	if domainCheck.IsAvailable() {
		t.Errorf("DomainCheck = %v; want false", domainCheck)
	}

	domainCheck, err = session.DomainCheck(fmt.Sprint("inscryb-", RandomString(15), ".com"))
	if err != nil {
		t.Error(err)
	}
	if !domainCheck.IsAvailable() {
		t.Errorf("DomainCheck = %v; want true", domainCheck)
	}

	domainCheck, err = session.DomainCheck("triplecservicesllc.com")
	if err != nil {
		t.Error(err)
	}

	if domainCheck.IsAvailable() {
		t.Errorf("DomainCheck = %v; want false", domainCheck)
	}
}

func TestDomainNameSpinner(t *testing.T) {
	session := CreateSession(DomainTestApiUrl, DomainTestApiResellerID, DomainTestApiResellerPwd)

	req := requests.DomainNameSpinnerRequest{Domain: "inscryb.com"}
	domainSpinner, err := session.DomainNameSpinner(req)
	if err != nil {
		t.Error(err)
	}

	if domainSpinner.ErrCount > 0 {
		t.Errorf("domain spinner contains errors")
	}
}

func TestDomainPurchase(t *testing.T) {
	session := CreateSession(DomainTestApiUrl, DomainTestApiResellerID, DomainTestApiResellerPwd)

	domainPurchase, err := session.DomainPurchase(fmt.Sprint("inscryb-", RandomString(15), ".com"))
	if err != nil {
		t.Error(err)
	}

	if domainPurchase.ResponseCode != 200 {
		println("Error Reason -----")
		println(domainPurchase.ResponseText)
		t.Errorf("DomainPurchase = %v; want 200", domainPurchase.ResponseCode)
	}
}
