package moysklad_api

import (
	"encoding/json"
)

type Retailshift struct {
	Meta struct {
		Href      string `json:"href"`
		Type      string `json:"type"`
		MediaType string `json:"mediaType"`
	} `json:"meta"`
	ID           string `json:"id"`
	AccountID    string `json:"accountId"`
	SyncID       string `json:"syncId"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Owner        struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
		} `json:"meta"`
	} `json:"owner"`
	Shared bool `json:"shared"`
	Group  struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
		} `json:"meta"`
	} `json:"group"`
	Moment       string `json:"moment"`
	Created      string `json:"created"`
	Printed      bool   `json:"printed"`
	Published    bool   `json:"published"`
	VatEnabled   bool   `json:"vatEnabled"`
	VatIncluded  bool   `json:"vatIncluded"`
	Organization struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
	} `json:"organization"`
	Store struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
	} `json:"store"`
	OrganizationAccount struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
	} `json:"organizationAccount"`
	Attributes []struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
		ID    string `json:"id"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"attributes"`
	CloseDate      string `json:"closeDate"`
	ProceedsNoCash int    `json:"proceedsNoCash"`
	ProceedsCash   int    `json:"proceedsCash"`
	ReceivedNoCash int    `json:"receivedNoCash"`
	ReceivedCash   int    `json:"receivedCash"`
	RetailStore    struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
	} `json:"retailStore"`
	Acquire struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
			UUIDHref     string `json:"uuidHref"`
		} `json:"meta"`
	} `json:"acquire"`
	BankPercent     float64 `json:"bankPercent"`
	BankComission   float64 `json:"bankComission"`
	QrBankPercent   float64 `json:"qrBankPercent"`
	QrBankComission float64 `json:"qrBankComission"`
	Cheque          struct {
		Start struct {
			ShiftNumber     string `json:"shiftNumber"`
			FnNumber        string `json:"fnNumber"`
			FiscalDocNumber string `json:"fiscalDocNumber"`
			FiscalDocSign   string `json:"fiscalDocSign"`
			Time            string `json:"time"`
		} `json:"start"`
		End struct {
			ShiftNumber     string `json:"shiftNumber"`
			FnNumber        string `json:"fnNumber"`
			FiscalDocNumber string `json:"fiscalDocNumber"`
			FiscalDocSign   string `json:"fiscalDocSign"`
			Time            string `json:"time"`
			ChequesTotal    int    `json:"chequesTotal"`
			FiscalDocsTotal int    `json:"fiscalDocsTotal"`
		} `json:"end"`
	} `json:"cheque"`
	Operations []struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
			UUIDHref     string `json:"uuidHref"`
		} `json:"meta"`
	} `json:"operations"`
	PaymentOperations []struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
			UUIDHref     string `json:"uuidHref"`
		} `json:"meta"`
		LinkedSum int `json:"linkedSum"`
	} `json:"paymentOperations"`
}

func GetRetailShift(id string) (Retailshift, error) {
	var data, err = SendHttpGetRequest("https://online.moysklad.ru/api/remap/1.2/entity/retailshift/"+id, "")

	var red Retailshift

	json.Unmarshal([]byte(data), &red)

	return red, err
}
