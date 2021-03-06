// Code generated by ros-gen-go.
// source: GeoPoseStamped.msg
// DO NOT EDIT!
package geographic_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgGeoPoseStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGeoPoseStamped) Text() string {
	return t.text
}

func (t *_MsgGeoPoseStamped) Name() string {
	return t.name
}

func (t *_MsgGeoPoseStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGeoPoseStamped) NewMessage() ros.Message {
	m := new(GeoPoseStamped)

	return m
}

var (
	MsgGeoPoseStamped = &_MsgGeoPoseStamped{
		`Header header
geographic_msgs/GeoPose pose
`,
		"geographic_msgs/GeoPoseStamped",
		"cc409c8ed6064d8a846fa207bf3fba6b",
	}
)

type GeoPoseStamped struct {
	Header std_msgs.Header
	Pose   GeoPose
}

func (m *GeoPoseStamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geographic_msgs/GeoPose", &m.Pose); err != nil {
		return err
	}

	return
}

func (m *GeoPoseStamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Pose
	if err = ros.DeserializeMessageField(r, "geographic_msgs/GeoPose", &m.Pose); err != nil {
		return err
	}

	return
}
