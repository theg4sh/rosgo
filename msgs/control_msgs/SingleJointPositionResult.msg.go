// Code generated by ros-gen-go.
// source: SingleJointPositionResult.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgSingleJointPositionResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSingleJointPositionResult) Text() string {
	return t.text
}

func (t *_MsgSingleJointPositionResult) Name() string {
	return t.name
}

func (t *_MsgSingleJointPositionResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSingleJointPositionResult) NewMessage() ros.Message {
	m := new(SingleJointPositionResult)

	return m
}

var (
	MsgSingleJointPositionResult = &_MsgSingleJointPositionResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
`,
		"control_msgs/SingleJointPositionResult",
		"d41d8cd98f00b204e9800998ecf8427e",
	}
)

type SingleJointPositionResult struct {
}

func (m *SingleJointPositionResult) Serialize(w io.Writer) (err error) {
	return
}

func (m *SingleJointPositionResult) Deserialize(r io.Reader) (err error) {
	return
}
