package response

type TOText struct {
	OrderTypeDesc string `xml:"ordertypedesc" json:"order_type_description"`
	Tod           string `xml:"tod"           json:"tod"`
}

type TOStatus struct {
	StatusDesc  string `xml:"statusdesc"    json:"status_description"`
	OrderTypeId int    `xml:"ordertypeid"   json:"order_type_id"`
	TOText      TOText `xml:"tot"           json:"to_text"`
}

type TransferOrder struct {
	Id        int      `xml:"transferorderid" json:"id"`
	OrderDate string   `xml:"orderdate"       json:"order_date"`
	StatusId  int      `xml:"statusid"        json:"status_id"`
	TOStatus  TOStatus `xml:"tos"             json:"transfer_order_status"`
}

type TransferOrderStatuses struct {
	Response

	Count         int             `xml:"TransferOrderCount" json:"transfer_order_count"`
	TransferOrder []TransferOrder `xml:"transferorder"      json:"transfer_order"`

	AtLeastOnePendingOrder bool `xml:"FoundAtLeastOnePendingOrder" json:"at_least_one_pending_order"`
}
