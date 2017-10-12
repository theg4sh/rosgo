// Code generated by ros-gen-go.
// source: UniqueID.msg
// DO NOT EDIT!
package uuid_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgUniqueID struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUniqueID) Text() string {
	return t.text
}

func (t *_MsgUniqueID) Name() string {
	return t.name
}

func (t *_MsgUniqueID) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUniqueID) NewMessage() ros.Message {
	m := new(UniqueID)

	return m
}

var (
	MsgUniqueID = &_MsgUniqueID{
		`# A universally unique identifier (UUID).
#
#  http://en.wikipedia.org/wiki/Universally_unique_identifier
#  http://tools.ietf.org/html/rfc4122.html

uint8[16] uuid
`,
		"uuid_msgs/UniqueID",
		"fec2a93b6f5367ee8112c9c0b41ff310",
	}
)

type UniqueID struct {
	UUID [16]uint8
}

func (m *UniqueID) Serialize(w io.Writer) (err error) {
	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.UUID)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.UUID {
		if err = ros.SerializeMessageField(w, "uint8", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *UniqueID) Deserialize(r io.Reader) (err error) {
	// UUID
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for UUID: %s", err)
		}
		if size > 16 {
			return fmt.Errorf("array size for UUID too large: expected=16, got=%d", size)
		}
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "uint8", &m.UUID[i]); err != nil {
				return err
			}
		}
	}

	return
}
