package mpower

import (
	"errors"
)

// The Store holds the store information and ised to define the store data for mpower transaction
type Store struct {
	Name          string `json:"name"`
	Tagline       string `json:"tagline"`
	PhoneNumber   string `json:"phone"`
	PostalAddress string `json:"postal_address"`
	LogoURL       string `json:"logo_url"`
}

// Get - gets a value from the struct by using its field name
//
// Example.
//    key := newStore.Get("Name")
func (store *Store) Get(fieldName string) string {
	return get(store, fieldName)
}

// NewStore - returns a new store object
//
// Example.
//    newSetup := mpower.NewStore(map[string]string{
//        "name":          "Awesome Store",
//        "tagline":       "Easy shopping",
//        "phoneNumber":   "0272271893",
//        "postalAddress": "P.0. Box MP555, Accra",
//        "logoURL":       "http://www.awesomestore.com.gh/logo.png",
//    })
func NewStore(storeInfo interface{}) (error, *Store) {
	var store *Store
	var err error
	switch storeInfo.(type) {
	case string:
		name, _ := storeInfo.(string)
		store = &Store{Name: name}
		err = nil
	case map[string]string:
		valueOfStore, _ := storeInfo.(map[string]string)
		if _, ok := valueOfStore["name"]; !ok {
			err = errors.New("Provide a name field with value")
		} else {
			store = &Store{
				Name: valueOfStore["name"],
			}

			var val string
			var ok bool
			if val, ok = valueOfStore["tagline"]; ok {
				store.Tagline = val
			}

			if val, ok = valueOfStore["phoneNumber"]; ok {
				store.PhoneNumber = val
			}

			if val, ok = valueOfStore["postalAddress"]; ok {
				store.PostalAddress = val
			}

			if val, ok = valueOfStore["logoURL"]; ok {
				store.LogoURL = val
			}
		}
	}

	if err != nil {
		return err, nil
	}

	return nil, store
}
