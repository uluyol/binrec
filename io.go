package binrec

import (
	"encoding/binary"
	"fmt"
	"io"
)

// This file contains helper io functions compatible with
// Java protobuf's writeDelimitedTo and mergeDelimitedFrom methods.

// WriteDelimitedTo writes a buffer to an io.Writer in a varint-delimited format.
func WriteDelimitedTo(w io.Writer, data []byte) error {
	var buf [binary.MaxVarintLen64]byte

	n := binary.PutUvarint(buf[:], uint64(len(data)))

	concat := make([]byte, len(data)+n)
	copy(concat, buf[:n])
	copy(concat[n:], data)

	_, err := w.Write(concat)
	return err
}

type Reader interface {
	io.ByteReader
	io.Reader
}

// ReadDelimitedFrom reads a buffer from a Reader in a varint-delimited format.
// bufio.Reader may be used to construct a Reader.
func ReadDelimitedFrom(r Reader) ([]byte, error) {
	dlen, err := binary.ReadUvarint(r)
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil, err
		}
		return nil, fmt.Errorf("unable to read length: %v", err)
	}
	buf := make([]byte, dlen)
	_, err = io.ReadFull(r, buf)
	return buf, err
}
