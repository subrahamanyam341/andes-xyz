//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/dharitri-org/protobuf/protobuf --gogoslick_out=$GOPATH/src config.proto

package outport

import (
	"github.com/subrahamanyam341/andes-xyz/core"
	"github.com/subrahamanyam341/andes-xyz/core/check"
	"github.com/subrahamanyam341/andes-xyz/data"
	"github.com/subrahamanyam341/andes-xyz/data/block"
	"github.com/subrahamanyam341/andes-xyz/marshal"
)

// GetHeaderBytesAndType returns the marshalled header bytes along with header type, if known
func GetHeaderBytesAndType(marshaller marshal.Marshalizer, headerHandler data.HeaderHandler) ([]byte, core.HeaderType, error) {
	if check.IfNil(marshaller) {
		return nil, "", core.ErrNilMarshalizer
	}

	var headerType core.HeaderType

	switch headerHandler.(type) {
	case *block.HeaderV2:
		headerType = core.ShardHeaderV2
	case *block.MetaBlock:
		headerType = core.MetaHeader
	case *block.Header:
		headerType = core.ShardHeaderV1
	default:
		return nil, "", errInvalidHeaderType
	}

	headerBytes, err := marshaller.Marshal(headerHandler)
	return headerBytes, headerType, err
}

// GetBody converts the BodyHandler interface to Body struct
func GetBody(bodyHandler data.BodyHandler) (*block.Body, error) {
	if check.IfNil(bodyHandler) {
		return nil, errNilBodyHandler
	}

	body, castOk := bodyHandler.(*block.Body)
	if !castOk {
		return nil, errCannotCastBlockBody
	}

	return body, nil
}

// ConvertPubKeys converts a map<shard, validators> into a map<shard, validatorsProtoMessage>
func ConvertPubKeys(validatorsPubKeys map[uint32][][]byte) map[uint32]*PubKeys {
	ret := make(map[uint32]*PubKeys, len(validatorsPubKeys))

	for shard, validators := range validatorsPubKeys {
		ret[shard] = &PubKeys{Keys: validators}
	}

	return ret
}
