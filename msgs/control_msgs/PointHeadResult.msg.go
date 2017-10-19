// Code generated by ros-gen-go.
// source: PointHeadResult.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgPointHeadResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointHeadResult) Text() string {
	return t.text
}

func (t *_MsgPointHeadResult) Name() string {
	return t.name
}

func (t *_MsgPointHeadResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointHeadResult) NewMessage() ros.Message {
	m := new(PointHeadResult)

	return m
}

var (
	MsgPointHeadResult = &_MsgPointHeadResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
`,
		"control_msgs/PointHeadResult",
		"d41d8cd98f00b204e9800998ecf8427e",
	}
)

type PointHeadResult struct {
}

func (m *PointHeadResult) Serialize(w io.Writer) (err error) {
	return
}

func (m *PointHeadResult) Deserialize(r io.Reader) (err error) {
	return
}
