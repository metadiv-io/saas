package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

func New() string {
	id, err := gonanoid.Generate("2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ", 21)
	if err != nil {
		panic(err)
	}
	return id
}

func PrefixNew(prefix string) string {
	id, err := gonanoid.Generate("2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ", 21)
	if err != nil {
		panic(err)
	}
	return prefix + id
}
