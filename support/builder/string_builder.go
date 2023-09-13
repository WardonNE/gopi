package builder

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type StringBuilder struct {
	strings.Builder
}

func NewStringBuilder[S ~string | ~rune | ~[]byte](str S) StringBuilder {
	var builder strings.Builder
	builder.WriteString(string(str))
	return StringBuilder{
		Builder: builder,
	}
}

func (sb StringBuilder) WriteInt(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteUint(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteInt8(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteUint8(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteInt16(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}
func (sb StringBuilder) WriteUint16(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteInt32(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteUint32(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteInt64(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteUint64(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

func (sb StringBuilder) WriteStringBuilder(builder StringBuilder) (int, error) {
	return sb.WriteString(builder.String())
}

func (sb StringBuilder) WriteBuffer(buf bytes.Buffer) (int, error) {
	return sb.Write(buf.Bytes())
}

func (sb StringBuilder) WriteBool(value bool) (int, error) {
	return sb.WriteString(strconv.FormatBool(value))
}

func (sb StringBuilder) TrimSpace() {
	str := strings.TrimSpace(sb.String())
	sb.Reset()
	sb.WriteString(str)
}
