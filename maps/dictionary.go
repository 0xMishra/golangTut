package maps

type (
	Dictionary    map[string]string
	DictionaryErr string
)

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr(
		"cannot perform operation on word because it does not exist",
	)
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(s string) (string, error) {
	def, ok := d[s]
	if ok {
		return def, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = def
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
