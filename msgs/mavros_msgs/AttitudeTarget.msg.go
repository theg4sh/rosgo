// Code generated by ros-gen-go.
// source: AttitudeTarget.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/geometry_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgAttitudeTarget struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAttitudeTarget) Text() string {
	return t.text
}

func (t *_MsgAttitudeTarget) Name() string {
	return t.name
}

func (t *_MsgAttitudeTarget) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAttitudeTarget) NewMessage() ros.Message {
	m := new(AttitudeTarget)

	return m
}

var (
	MsgAttitudeTarget = &_MsgAttitudeTarget{
		`# Message for SET_ATTITUDE_TARGET
#
# Some complex system requires all feautures that mavlink
# message provide. See issue #402, #418.

std_msgs/Header header

uint8 type_mask
uint8 IGNORE_ROLL_RATE = 1	# body_rate.x
uint8 IGNORE_PITCH_RATE = 2	# body_rate.y
uint8 IGNORE_YAW_RATE = 4	# body_rate.z
uint8 IGNORE_THRUST = 64
uint8 IGNORE_ATTITUDE = 128	# orientation field

geometry_msgs/Quaternion orientation
geometry_msgs/Vector3 body_rate
float32 thrust
`,
		"mavros_msgs/AttitudeTarget",
		"456f3af666b22ccd0222ea053a86b548",
	}
)

type AttitudeTarget struct {
	Header      std_msgs.Header
	TypeMask    uint8
	Orientation geometry_msgs.Quaternion
	BodyRate    geometry_msgs.Vector3
	Thrust      float32
}

func (m *AttitudeTarget) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.TypeMask); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Quaternion", &m.Orientation); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.BodyRate); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Thrust); err != nil {
		return err
	}

	return
}

func (m *AttitudeTarget) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// TypeMask
	if err = ros.DeserializeMessageField(r, "uint8", &m.TypeMask); err != nil {
		return err
	}

	// Orientation
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Quaternion", &m.Orientation); err != nil {
		return err
	}

	// BodyRate
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.BodyRate); err != nil {
		return err
	}

	// Thrust
	if err = ros.DeserializeMessageField(r, "float32", &m.Thrust); err != nil {
		return err
	}

	return
}
