// Code generated by ros-gen-go.
// source: Waypoint.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

const (
	Waypoint_FRAME_GLOBAL         uint8 = 0
	Waypoint_FRAME_LOCAL_NED      uint8 = 1
	Waypoint_FRAME_MISSION        uint8 = 2
	Waypoint_FRAME_GLOBAL_REL_ALT uint8 = 3
	Waypoint_FRAME_LOCAL_ENU      uint8 = 4
)

type _MsgWaypoint struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypoint) Text() string {
	return t.text
}

func (t *_MsgWaypoint) Name() string {
	return t.name
}

func (t *_MsgWaypoint) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypoint) NewMessage() ros.Message {
	m := new(Waypoint)

	return m
}

var (
	MsgWaypoint = &_MsgWaypoint{
		`# Waypoint.msg
#
# ROS representation of MAVLink MISSION_ITEM
# See mavlink documentation



# see enum MAV_FRAME
uint8 frame
uint8 FRAME_GLOBAL = 0
uint8 FRAME_LOCAL_NED = 1
uint8 FRAME_MISSION = 2
uint8 FRAME_GLOBAL_REL_ALT = 3
uint8 FRAME_LOCAL_ENU = 4

# see enum MAV_CMD and CommandCode.msg
uint16 command

bool is_current
bool autocontinue
# meaning of this params described in enum MAV_CMD
float32 param1
float32 param2
float32 param3
float32 param4
float64 x_lat
float64 y_long
float64 z_alt
`,
		"mavros_msgs/Waypoint",
		"7a0d2b53dcd6b7aff0aa748703e85e92",
	}
)

type Waypoint struct {
	Frame        uint8
	Command      uint16
	IsCurrent    bool
	Autocontinue bool
	Param1       float32
	Param2       float32
	Param3       float32
	Param4       float32
	XLat         float64
	YLong        float64
	ZAlt         float64
}

func (m *Waypoint) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint8", &m.Frame); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.Command); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "bool", &m.IsCurrent); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "bool", &m.Autocontinue); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Param1); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Param2); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Param3); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Param4); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.XLat); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.YLong); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.ZAlt); err != nil {
		return err
	}

	return
}

func (m *Waypoint) Deserialize(r io.Reader) (err error) {
	// Frame
	if err = ros.DeserializeMessageField(r, "uint8", &m.Frame); err != nil {
		return err
	}

	// Command
	if err = ros.DeserializeMessageField(r, "uint16", &m.Command); err != nil {
		return err
	}

	// IsCurrent
	if err = ros.DeserializeMessageField(r, "bool", &m.IsCurrent); err != nil {
		return err
	}

	// Autocontinue
	if err = ros.DeserializeMessageField(r, "bool", &m.Autocontinue); err != nil {
		return err
	}

	// Param1
	if err = ros.DeserializeMessageField(r, "float32", &m.Param1); err != nil {
		return err
	}

	// Param2
	if err = ros.DeserializeMessageField(r, "float32", &m.Param2); err != nil {
		return err
	}

	// Param3
	if err = ros.DeserializeMessageField(r, "float32", &m.Param3); err != nil {
		return err
	}

	// Param4
	if err = ros.DeserializeMessageField(r, "float32", &m.Param4); err != nil {
		return err
	}

	// XLat
	if err = ros.DeserializeMessageField(r, "float64", &m.XLat); err != nil {
		return err
	}

	// YLong
	if err = ros.DeserializeMessageField(r, "float64", &m.YLong); err != nil {
		return err
	}

	// ZAlt
	if err = ros.DeserializeMessageField(r, "float64", &m.ZAlt); err != nil {
		return err
	}

	return
}
