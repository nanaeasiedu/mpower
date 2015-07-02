package mpowergo

import (
	"errors"
)

type Store struct {
	Name          string
	Tagline       string
	PhoneNumber   string
	PostalAddress string
	LogoUrl       string
}

func (store *Store) Get(fieldName string) string {
	return get(store, fieldName)
}

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

			if val, ok = valueOfStore["logoUrl"]; ok {
				store.LogoUrl = val
			}
		}
	}

	if err != nil {
		return err, nil
	}

	return nil, store
}
