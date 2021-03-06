// Code generated by ros-gen-go.
// source: Int8MultiArray.msg
// DO NOT EDIT!
package std_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgInt8MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt8MultiArray) Text() string {
	return t.text
}

func (t *_MsgInt8MultiArray) Name() string {
	return t.name
}

func (t *_MsgInt8MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt8MultiArray) NewMessage() ros.Message {
	m := new(Int8MultiArray)

	return m
}

var (
	MsgInt8MultiArray = &_MsgInt8MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
int8[]            data          # array of data

`,
		"std_msgs/Int8MultiArray",
		"d7c1af35a1b4781bbe79e03dd94b7c13",
	}
)

type Int8MultiArray struct {
	Layout MultiArrayLayout
	Data   []int8
}

func (m *Int8MultiArray) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Data)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Data {
		if err = ros.SerializeMessageField(w, "int8", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *Int8MultiArray) Deserialize(r io.Reader) (err error) {
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
		m.Data = make([]int8, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "int8", &m.Data[i]); err != nil {
				return err
			}
		}
	}

	return
}
