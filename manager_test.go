package goconf

import (
	"fmt"
	"testing"
)

func TestNewConfigManager(t *testing.T) {
	cm := DefaultConfigManager()

	type PhoneNumber struct {
		Type   string `json:"type"`
		Number string `json:"number"`
	}
type Address struct {
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postalCode"`
}
	type West struct {
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	Sex         string        `json:"sex"`
	Age         int64         `json:"age"`
	Address     Address       `json:"address"`
	PhoneNumber []PhoneNumber `json:"phoneNumber"`
}
	v := &West{}
	cm.Load("./testdata/conf.json", &v)
	fmt.Printf("%+v\n", GetOrDefault(v, 10, "Age"))
	fmt.Printf("%+v\n", GetOrDefault(v, "fuck", "aaaaaa"))

}
