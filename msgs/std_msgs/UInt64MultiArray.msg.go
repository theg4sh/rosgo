// Code generated by ros-gen-go.
// source: UInt64MultiArray.msg
// DO NOT EDIT!
package std_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgUInt64MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt64MultiArray) Text() string {
	return t.text
}

func (t *_MsgUInt64MultiArray) Name() string {
	return t.name
}

func (t *_MsgUInt64MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt64MultiArray) NewMessage() ros.Message {
	m := new(UInt64MultiArray)

	return m
}

var (
	MsgUInt64MultiArray = &_MsgUInt64MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
uint64[]          data          # array of data

`,
		"std_msgs/UInt64MultiArray",
		"6088f127afb1d6c72927aa1247e945af",
	}
)

type UInt64MultiArray struct {
	Layout MultiArrayLayout
	Data   []uint64
}

func (m *UInt64MultiArray) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Data)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Data {
		if err = ros.SerializeMessageField(w, "uint64", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *UInt64MultiArray) Deserialize(r io.Reader) (err error) {
	// Layout
	if err = ros.DeserializeMessageField(r, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Data
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Data: %s", err)
		}
		m.Data = make([]uint64, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "uint64", &m.Data[i]); err != nil {
				return err
			}
		}
	}

	return
}
