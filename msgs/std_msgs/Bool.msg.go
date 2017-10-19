// Code generated by ros-gen-go.
// source: Bool.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgBool struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgBool) Text() string {
	return t.text
}

func (t *_MsgBool) Name() string {
	return t.name
}

func (t *_MsgBool) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgBool) NewMessage() ros.Message {
	m := new(Bool)

	return m
}

var (
	MsgBool = &_MsgBool{
		`bool data`,
		"std_msgs/Bool",
		"8b94c1b53db61fb6aed406028ad6332a",
	}
)

type Bool struct {
	Data bool
}

func (m *Bool) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Data); err != nil {
		return err
	}

	return
}

func (m *Bool) Deserialize(r io.Reader) (err error) {
	// Data
	if err = ros.DeserializeMessageField(r, "bool", &m.Data); err != nil {
		return err
	}

	return
}
