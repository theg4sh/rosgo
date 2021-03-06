// Code generated by ros-gen-go.
// source: AccelWithCovarianceStamped.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgAccelWithCovarianceStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAccelWithCovarianceStamped) Text() string {
	return t.text
}

func (t *_MsgAccelWithCovarianceStamped) Name() string {
	return t.name
}

func (t *_MsgAccelWithCovarianceStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAccelWithCovarianceStamped) NewMessage() ros.Message {
	m := new(AccelWithCovarianceStamped)

	return m
}

var (
	MsgAccelWithCovarianceStamped = &_MsgAccelWithCovarianceStamped{
		`# This represents an estimated accel with reference coordinate frame and timestamp.
Header header
AccelWithCovariance accel
`,
		"geometry_msgs/AccelWithCovarianceStamped",
		"96adb295225031ec8d57fb4251b0a886",
	}
)

type AccelWithCovarianceStamped struct {
	Header std_msgs.Header
	Accel  AccelWithCovariance
}

func (m *AccelWithCovarianceStamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "AccelWithCovariance", &m.Accel); err != nil {
		return err
	}

	return
}

func (m *AccelWithCovarianceStamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Accel
	if err = ros.DeserializeMessageField(r, "AccelWithCovariance", &m.Accel); err != nil {
		return err
	}

	return
}
