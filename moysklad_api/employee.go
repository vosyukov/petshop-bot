package moysklad_api

import (
	"encoding/json"
)

type Employee struct {
	Meta struct {
		Href         string `json:"href"`
		MetadataHref string `json:"metadataHref"`
		Type         string `json:"type"`
		MediaType    string `json:"mediaType"`
		UUIDHref     string `json:"uuidHref"`
	} `json:"meta"`
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Owner     struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
			UUIDHref     string `json:"uuidHref"`
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
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	Created      string `json:"created"`
	UID          string `json:"uid"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	FullName     string `json:"fullName"`
	ShortFio     string `json:"shortFio"`
}

func GetEmployee(id string) (Employee, error) {
	var data, err = SendHttpGetRequest("https://online.moysklad.ru/api/remap/1.2/entity/employee/"+id, "")

	var red Employee

	json.Unmarshal([]byte(data), &red)

	return red, err
}
