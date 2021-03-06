// Code generated by ros-gen-go.
// source: HilControls.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgHilControls struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgHilControls) Text() string {
	return t.text
}

func (t *_MsgHilControls) Name() string {
	return t.name
}

func (t *_MsgHilControls) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgHilControls) NewMessage() ros.Message {
	m := new(HilControls)

	return m
}

var (
	MsgHilControls = &_MsgHilControls{
		`# HilControls.msg
#
# ROS representation of MAVLink HIL_CONTROLS
# (deprecated, use HIL_ACTUATOR_CONTROLS instead)
# See mavlink message documentation here:
# https://pixhawk.ethz.ch/mavlink/#HIL_CONTROLS

std_msgs/Header header
float32 roll_ailerons
float32 pitch_elevator
float32 yaw_rudder
float32 throttle
float32 aux1
float32 aux2
float32 aux3
float32 aux4
uint8 mode
uint8 nav_mode
`,
		"mavros_msgs/HilControls",
		"698148349c3a2e5720afcae2d934acca",
	}
)

type HilControls struct {
	Header        std_msgs.Header
	RollAilerons  float32
	PitchElevator float32
	YawRudder     float32
	Throttle      float32
	Aux1          float32
	Aux2          float32
	Aux3          float32
	Aux4          float32
	Mode          uint8
	NavMode       uint8
}

func (m *HilControls) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.RollAilerons); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.PitchElevator); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.YawRudder); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Throttle); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Aux1); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Aux2); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Aux3); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Aux4); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Mode); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.NavMode); err != nil {
		return err
	}

	return
}

func (m *HilControls) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// RollAilerons
	if err = ros.DeserializeMessageField(r, "float32", &m.RollAilerons); err != nil {
		return err
	}

	// PitchElevator
	if err = ros.DeserializeMessageField(r, "float32", &m.PitchElevator); err != nil {
		return err
	}

	// YawRudder
	if err = ros.DeserializeMessageField(r, "float32", &m.YawRudder); err != nil {
		return err
	}

	// Throttle
	if err = ros.DeserializeMessageField(r, "float32", &m.Throttle); err != nil {
		return err
	}

	// Aux1
	if err = ros.DeserializeMessageField(r, "float32", &m.Aux1); err != nil {
		return err
	}

	// Aux2
	if err = ros.DeserializeMessageField(r, "float32", &m.Aux2); err != nil {
		return err
	}

	// Aux3
	if err = ros.DeserializeMessageField(r, "float32", &m.Aux3); err != nil {
		return err
	}

	// Aux4
	if err = ros.DeserializeMessageField(r, "float32", &m.Aux4); err != nil {
		return err
	}

	// Mode
	if err = ros.DeserializeMessageField(r, "uint8", &m.Mode); err != nil {
		return err
	}

	// NavMode
	if err = ros.DeserializeMessageField(r, "uint8", &m.NavMode); err != nil {
		return err
	}

	return
}
