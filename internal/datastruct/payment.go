package datastruct

type Notification struct {
	Bill    Bill   `json:"bill,omitempty"`
	Version string `json:"version,omitempty"`
}

type Bill struct {
	SiteID             string       `json:"siteId,omitempty"`
	BillID             string       `json:"billId,omitempty"`
	Amount             Amount       `json:"amount,omitempty"`
	Status             Status       `json:"status,omitempty"`
	Customer           Customer     `json:"customer,omitempty"`
	CustomFields       CustomFields `json:"customFields,omitempty"`
	Comment            string       `json:"comment,omitempty"`
	CreationDateTime   string       `json:"creationDateTime,omitempty"`
	ExpirationDateTime string       `json:"expirationDateTime,omitempty"`
}

type Payment struct {
	SiteID             string       `json:"siteId,omitempty"`
	BillID             string       `json:"billId,omitempty"`
	Amount             Amount       `json:"amount,omitempty"`
	Status             *Status      `json:"status,omitempty"`
	Customer           Customer     `json:"customer,omitempty"`
	CustomFields       CustomFields `json:"customFields,omitempty"`
	Comment            string       `json:"comment,omitempty"`
	CreationDateTime   string       `json:"creationDateTime,omitempty"`
	ExpirationDateTime string       `json:"expirationDateTime,omitempty"`
	PayURL             string       `json:"payUrl,omitempty"`
}

type Amount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value"`
}

type Status struct {
	Value           string `json:"value,omitempty"`
	ChangedDateTime string `json:"changedDateTime,omitempty"`
}

type Customer struct {
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
	Account string `json:"account,omitempty"`
}

type CustomFields struct {
	PaySourcesFilter string `json:"paySourcesFilter,omitempty"`
	ThemeCode        string `json:"themeCode,omitempty"`
	YourParam1       string `json:"yourParam1,omitempty"`
	YourParam2       string `json:"yourParam2,omitempty"`
}
