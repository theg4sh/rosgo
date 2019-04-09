// Code generated by ros-gen-go.
// source: Trajectory.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

const (
	Trajectory_MAV_TRAJECTORY_REPRESENTATION_WAYPOINTS uint8 = 0
	Trajectory_MAV_TRAJECTORY_REPRESENTATION_BEZIER    uint8 = 1
)

type _MsgTrajectory struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTrajectory) Text() string {
	return t.text
}

func (t *_MsgTrajectory) Name() string {
	return t.name
}

func (t *_MsgTrajectory) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTrajectory) NewMessage() ros.Message {
	m := new(Trajectory)

	return m
}

var (
	MsgTrajectory = &_MsgTrajectory{
		`# MAVLink message: TRAJECTORY
# https://mavlink.io/en/messages/common.html#TRAJECTORY

std_msgs/Header header

uint8 type # See enum MAV_TRAJECTORY_REPRESENTATION.
uint8 MAV_TRAJECTORY_REPRESENTATION_WAYPOINTS = 0
uint8 MAV_TRAJECTORY_REPRESENTATION_BEZIER = 1

mavros_msgs/PositionTarget point_1
mavros_msgs/PositionTarget point_2
mavros_msgs/PositionTarget point_3
mavros_msgs/PositionTarget point_4
mavros_msgs/PositionTarget point_5

uint8[5] point_valid # States if respective point is valid.

float32[5] time_horizon # if type MAV_TRAJECTORY_REPRESENTATION_BEZIER, it represents the time horizon for each point, otherwise set to NaN
`,
		"mavros_msgs/Trajectory",
		"3d34ec9673348ab7c0bc80373ec76fc8",
	}
)

type Trajectory struct {
	Header      std_msgs.Header
	Type        uint8
	Point1      PositionTarget
	Point2      PositionTarget
	Point3      PositionTarget
	Point4      PositionTarget
	Point5      PositionTarget
	PointValid  [5]uint8
	TimeHorizon [5]float32
}

func (m *Trajectory) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Type); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "mavros_msgs/PositionTarget", &m.Point1); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "mavros_msgs/PositionTarget", &m.Point2); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "mavros_msgs/PositionTarget", &m.Point3); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "mavros_msgs/PositionTarget", &m.Point4); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "mavros_msgs/PositionTarget", &m.Point5); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.PointValid)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.PointValid {
		if err = ros.SerializeMessageField(w, "uint8", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.TimeHorizon)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.TimeHorizon {
		if err = ros.SerializeMessageField(w, "float32", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *Trajectory) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// Type
	if err = ros.DeserializeMessageField(r, "uint8", &m.Type); err != nil {
		return err
	}

	// Point1
	if err = ros.DeserializeMessageField(r, "mavros_msgs/PositionTarget", &m.Point1); err != nil {
		return err
	}

	// Point2
	if err = ros.DeserializeMessageField(r, "mavros_msgs/PositionTarget", &m.Point2); err != nil {
		return err
	}

	// Point3
	if err = ros.DeserializeMessageField(r, "mavros_msgs/PositionTarget", &m.Point3); err != nil {
		return err
	}

	// Point4
	if err = ros.DeserializeMessageField(r, "mavros_msgs/PositionTarget", &m.Point4); err != nil {
		return err
	}

	// Point5
	if err = ros.DeserializeMessageField(r, "mavros_msgs/PositionTarget", &m.Point5); err != nil {
		return err
	}

	// PointValid
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for PointValid: %s", err)
		}
		if size > 5 {
			return fmt.Errorf("array size for PointValid too large: expected=5, got=%d", size)
		}
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "uint8", &m.PointValid[i]); err != nil {
				return err
			}
		}
	}

	// TimeHorizon
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for TimeHorizon: %s", err)
		}
		if size > 5 {
			return fmt.Errorf("array size for TimeHorizon too large: expected=5, got=%d", size)
		}
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float32", &m.TimeHorizon[i]); err != nil {
				return err
			}
		}
	}

	return
}