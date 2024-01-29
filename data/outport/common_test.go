package outport

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subrahamanyam341/andes-xyz/core"
	"github.com/subrahamanyam341/andes-xyz/core/mock"
	"github.com/subrahamanyam341/andes-xyz/data"
	"github.com/subrahamanyam341/andes-xyz/data/block"
)

func TestGetHeaderBytesAndType(t *testing.T) {
	t.Parallel()

	headerBytes, headerType, err := GetHeaderBytesAndType(nil, nil)
	require.Nil(t, headerBytes)
	require.Empty(t, headerType)
	require.Equal(t, core.ErrNilMarshalizer, err)

	marshaller := &mock.MarshalizerMock{}
	headerBytes, headerType, err = GetHeaderBytesAndType(marshaller, nil)
	require.Nil(t, headerBytes)
	require.Empty(t, headerType)
	require.Equal(t, errInvalidHeaderType, err)

	var header data.HeaderHandler

	header = &block.Header{}
	headerBytes, headerType, err = GetHeaderBytesAndType(marshaller, header)
	expectedHeaderBytes, _ := marshaller.Marshal(header)
	require.Equal(t, expectedHeaderBytes, headerBytes)
	require.Equal(t, core.ShardHeaderV1, headerType)
	require.Nil(t, err)

	header = &block.HeaderV2{}
	headerBytes, headerType, err = GetHeaderBytesAndType(marshaller, header)
	expectedHeaderBytes, _ = marshaller.Marshal(header)
	require.Equal(t, expectedHeaderBytes, headerBytes)
	require.Equal(t, core.ShardHeaderV2, headerType)
	require.Nil(t, err)

	header = &block.MetaBlock{}
	headerBytes, headerType, err = GetHeaderBytesAndType(marshaller, header)
	expectedHeaderBytes, _ = marshaller.Marshal(header)
	require.Equal(t, expectedHeaderBytes, headerBytes)
	require.Equal(t, core.MetaHeader, headerType)
	require.Nil(t, err)
}

func TestGetBody(t *testing.T) {
	t.Parallel()

	receivedBody, err := GetBody(nil)
	require.Nil(t, receivedBody)
	require.Equal(t, errNilBodyHandler, err)

	var body data.BodyHandler = &block.Body{}
	receivedBody, err = GetBody(body)
	require.Nil(t, err)
	require.Equal(t, body, receivedBody)
}

func TestConvertPubKeys(t *testing.T) {
	t.Parallel()

	validatorsPubKeys := map[uint32][][]byte{
		0:                     {[]byte("pubKey1"), []byte("pubKey2")},
		core.MetachainShardId: {[]byte("pubKey3")},
	}

	ret := ConvertPubKeys(validatorsPubKeys)
	require.Equal(t, map[uint32]*PubKeys{
		0:                     {Keys: validatorsPubKeys[0]},
		core.MetachainShardId: {Keys: validatorsPubKeys[core.MetachainShardId]},
	}, ret)
}
