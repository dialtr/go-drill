package drill

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/zeroshade/go-drill/internal/rpc/proto/exec/rpc"
	"github.com/zeroshade/go-drill/internal/rpc/proto/exec/shared"
)

var deadbeef = []byte{0xDE, 0xAD, 0xBE, 0xEF}

func TestMakePrefixedMessage(t *testing.T) {
	out := makePrefixedMessage(deadbeef)
	val, nb := binary.Uvarint(out)

	assert.EqualValues(t, len(deadbeef), val)
	assert.Len(t, out, len(deadbeef)+nb)
}

func TestMakePrefixedNil(t *testing.T) {
	assert.Nil(t, makePrefixedMessage(nil))
}

func TestReadPrefixedSimple(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen32)
	nb := binary.PutUvarint(buf, 4)

	out, err := readPrefixed(bytes.NewReader(append(buf[:nb], deadbeef...)))
	assert.NoError(t, err)
	assert.EqualValues(t, deadbeef, out)
}

func TestReadPrefixedEof(t *testing.T) {
	buf := &bytes.Reader{}
	out, err := readPrefixed(buf)
	assert.Nil(t, out)
	assert.Same(t, io.ErrUnexpectedEOF, err)
}

func TestReadPrefixedShortRead(t *testing.T) {
	buf := []byte{0x01}
	out, err := readPrefixed(bytes.NewReader(buf))
	assert.Nil(t, out)
	assert.Same(t, io.ErrUnexpectedEOF, err)
}

func TestReadPrefixedEmpty(t *testing.T) {
	buf := []byte{0, 0, 0, 0, 0}
	out, err := readPrefixed(bytes.NewBuffer(buf))
	assert.Nil(t, out)
	assert.Same(t, errInvalidResponse, err)
}

func TestReadPrefixedNotEnough(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen32)
	nb := binary.PutUvarint(buf, 6)

	out, err := readPrefixed(bytes.NewReader(append(buf[:nb], deadbeef...)))
	assert.Nil(t, out)
	assert.Same(t, io.ErrUnexpectedEOF, err)
}

func TestReadPrefixedRaw(t *testing.T) {
	msg := &rpc.CompleteRpcMessage{
		Header: &rpc.RpcHeader{
			Mode:           rpc.RpcMode_PING.Enum(),
			CoordinationId: proto.Int(5),
			RpcType:        proto.Int(5),
		},
		ProtobufBody: deadbeef,
	}

	data, _ := proto.Marshal(msg)
	buf := make([]byte, binary.MaxVarintLen32)
	nb := binary.PutUvarint(buf, uint64(len(data)))

	out, err := readPrefixedRaw(bytes.NewReader(append(buf[:nb], data...)))
	assert.NoError(t, err)

	assert.Equal(t, msg.Header.GetMode(), out.Header.GetMode())
	assert.Equal(t, msg.Header.GetCoordinationId(), out.Header.GetCoordinationId())
	assert.Equal(t, msg.Header.GetRpcType(), out.Header.GetRpcType())
	assert.Equal(t, msg.GetProtobufBody(), out.GetProtobufBody())
}

func TestReadPrefixedRawErr(t *testing.T) {
	buf := &bytes.Reader{}
	out, err := readPrefixedRaw(buf)
	assert.Nil(t, out)
	assert.Error(t, err)
}

func TestReadPrefixedMessage(t *testing.T) {
	qid := &shared.QueryId{Part1: proto.Int64(12345), Part2: proto.Int64(98765)}
	encoded, _ := proto.Marshal(qid)
	msg := &rpc.CompleteRpcMessage{
		Header: &rpc.RpcHeader{
			Mode:           rpc.RpcMode_PING.Enum(),
			CoordinationId: proto.Int(5),
			RpcType:        proto.Int(5),
		},
		ProtobufBody: encoded,
	}

	data, _ := proto.Marshal(msg)
	buf := make([]byte, binary.MaxVarintLen32)
	nb := binary.PutUvarint(buf, uint64(len(data)))

	out := &shared.QueryId{}
	hdr, err := readPrefixedMessage(bytes.NewReader(append(buf[:nb], data...)), out)
	assert.NoError(t, err)

	assert.Equal(t, msg.Header.GetMode(), hdr.GetMode())
	assert.Equal(t, msg.Header.GetCoordinationId(), hdr.GetCoordinationId())
	assert.Equal(t, msg.Header.GetRpcType(), hdr.GetRpcType())
	assert.Equal(t, qid.GetPart1(), out.GetPart1())
	assert.Equal(t, qid.GetPart2(), out.GetPart2())
}

func TestReadPrefixedMessageErr(t *testing.T) {
	buf := &bytes.Reader{}
	out, err := readPrefixedMessage(buf, nil)
	assert.Nil(t, out)
	assert.Error(t, err)
}