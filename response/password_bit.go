package response

type PasswordBit struct {
	Response

	DomainPassword string `xml:"DomainPassword" json:"domain_password"`
	PasswordSet    int    `xml:"password-set"   json:"password_set"`
}
