package util

import (
	"bytes"
	"encoding/gob"
)

func EncodeEventMsg(msg any) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(msg)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func DecodeEventMsg(data []byte, m any) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
