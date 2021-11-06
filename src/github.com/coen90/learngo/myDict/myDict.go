package myDict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errorNotFound = errors.New("Not found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	// if err == errorNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	return nil
}
