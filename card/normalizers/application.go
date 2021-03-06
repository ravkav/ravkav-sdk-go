package normalizers

import (
	"encoding/hex"
	"github.com/ybaruchel/ravkav-sdk-go/card/parsers"
	"github.com/ybaruchel/ravkav-sdk-go/contracts"
	"strconv"
)

type applicationNormalizer struct{}

func NewApplicationNormalizer(_ contracts.CardOutput) contracts.Normalizer {
	return &applicationNormalizer{}
}

func (n *applicationNormalizer) Normalize(record contracts.Record, recordIndex int) (map[string]interface{}, error) {
	bytes := record.Bytes()
	bytes = bytes[19 : bytes[18]+19]

	var bytesLong []byte = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for i := range bytes {
		bytesLong[(8-len(bytes))+i] = bytes[i]
	}

	return map[string]interface{}{
		"cardNumber": strconv.Itoa(parsers.Hex2Int(hex.EncodeToString(bytesLong))),
	}, nil
}
