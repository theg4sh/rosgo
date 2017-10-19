// Code generated by ros-gen-go.
// source: Joy.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgJoy struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJoy) Text() string {
	return t.text
}

func (t *_MsgJoy) Name() string {
	return t.name
}

func (t *_MsgJoy) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJoy) NewMessage() ros.Message {
	m := new(Joy)

	return m
}

var (
	MsgJoy = &_MsgJoy{
		`# Reports the state of a joysticks axes and buttons.
Header header           # timestamp in the header is the time the data is received from the joystick
float32[] axes          # the axes measurements from a joystick
int32[] buttons         # the buttons measurements from a joystick 
`,
		"sensor_msgs/Joy",
		"5a9ea5f83505693b71e785041e67a8bb",
	}
)

type Joy struct {
	Header  std_msgs.Header
	Axes    []float32
	Buttons []int32
}

func (m *Joy) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Axes)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Axes {
		if err = ros.SerializeMessageField(w, "float32", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Buttons)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Buttons {
		if err = ros.SerializeMessageField(w, "int32", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *Joy) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Axes
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Axes: %s", err)
		}
		m.Axes = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float32", &m.Axes[i]); err != nil {
				return err
			}
		}
	}

	// Buttons
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Buttons: %s", err)
		}
		m.Buttons = make([]int32, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "int32", &m.Buttons[i]); err != nil {
				return err
			}
		}
	}

	return
}
