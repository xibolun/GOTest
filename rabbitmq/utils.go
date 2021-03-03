package rabbitmq

import (
	"bytes"
	"github.com/damonchen/rubymarshal"
)

func RubyMarshal(o interface{}) (data []byte, err error) {
	w := bytes.NewBuffer([]byte{})
	enc := rbmarshal.NewEncoder(w)
	err = enc.Encode(o)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func UnRubyMarshal(data []byte, v interface{}) error {
	return rbmarshal.NewDecoder(bytes.NewReader(data)).Decode(v)
}
