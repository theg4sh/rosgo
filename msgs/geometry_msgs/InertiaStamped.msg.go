// Code generated by ros-gen-go.
// source: InertiaStamped.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgInertiaStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInertiaStamped) Text() string {
	return t.text
}

func (t *_MsgInertiaStamped) Name() string {
	return t.name
}

func (t *_MsgInertiaStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInertiaStamped) NewMessage() ros.Message {
	m := new(InertiaStamped)

	return m
}

var (
	MsgInertiaStamped = &_MsgInertiaStamped{
		`Header header
Inertia inertia
`,
		"geometry_msgs/InertiaStamped",
		"ddee48caeab5a966c5e8d166654a9ac7",
	}
)

type InertiaStamped struct {
	Header  std_msgs.Header
	Inertia Inertia
}

func (m *InertiaStamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Inertia", &m.Inertia); err != nil {
		return err
	}

	return
}

func (m *InertiaStamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Inertia
	if err = ros.DeserializeMessageField(r, "Inertia", &m.Inertia); err != nil {
		return err
	}

	return
}
