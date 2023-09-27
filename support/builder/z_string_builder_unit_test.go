package builder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStringBuilder(t *testing.T) {
	builder := NewStringBuilder[string]("string")
	assert.Equal(t, "string", builder.String())

	builder = NewStringBuilder[rune]('s')
	assert.Equal(t, "s", builder.String())

	builder = NewStringBuilder[[]byte]([]byte("string"))
	assert.Equal(t, "string", builder.String())
}

func TestWriteInt(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteInt(-123)
	assert.Equal(t, "-123", builder.String())
}

func TestWriteUint(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteUint(123)
	assert.Equal(t, "123", builder.String())
}

func TestWriteInt8(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteInt8(-1)
	assert.Equal(t, "-1", builder.String())
}

func TestWriteUint8(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteUint8(1)
	assert.Equal(t, "1", builder.String())
}

func TestWriteInt16(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteInt16(-1)
	assert.Equal(t, "-1", builder.String())
}

func TestWriteUint16(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteUint16(1)
	assert.Equal(t, "1", builder.String())
}

func TestWriteInt32(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteInt32(-123)
	assert.Equal(t, "-123", builder.String())
}

func TestWriteUint32(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteUint32(123)
	assert.Equal(t, "123", builder.String())
}

func TestWriteInt64(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteInt64(-123)
	assert.Equal(t, "-123", builder.String())
}

func TestWriteUint64(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteUint64(123)
	assert.Equal(t, "123", builder.String())
}

func TestWriteStringBuilder(t *testing.T) {
	builder1 := NewStringBuilder("builder_1")
	builder2 := NewStringBuilder("builder_2")
	builder1.WriteStringBuilder(builder2)
	assert.Equal(t, "builder_1builder_2", builder1.String())
}

func TestWriteBuffer(t *testing.T) {
	builder := NewStringBuilder("")
	buffer := bytes.NewBuffer([]byte("buffer"))
	builder.WriteBuffer(buffer)
	assert.Equal(t, "buffer", builder.String())
}

func TestWriteBool(t *testing.T) {
	builder := NewStringBuilder("")
	builder.WriteBool(true)
	builder.WriteBool(false)
	assert.Equal(t, "truefalse", builder.String())
}

func TestTrimSpace(t *testing.T) {
	builder := NewStringBuilder("   123   ")
	builder.TrimSpace()
	assert.Equal(t, "123", builder.String())
}
