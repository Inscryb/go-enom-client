package response

type Delete struct {
	Response

	DomainDeleted bool `xml:"deletedomain>domaindeleted" json:"domain_deleted"`
}
