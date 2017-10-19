// Code generated by ros-gen-go.
// source: MultiDOFJointTrajectoryPoint.msg
// DO NOT EDIT!
package trajectory_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/geometry_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgMultiDOFJointTrajectoryPoint struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiDOFJointTrajectoryPoint) Text() string {
	return t.text
}

func (t *_MsgMultiDOFJointTrajectoryPoint) Name() string {
	return t.name
}

func (t *_MsgMultiDOFJointTrajectoryPoint) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiDOFJointTrajectoryPoint) NewMessage() ros.Message {
	m := new(MultiDOFJointTrajectoryPoint)

	return m
}

var (
	MsgMultiDOFJointTrajectoryPoint = &_MsgMultiDOFJointTrajectoryPoint{
		`# Each multi-dof joint can specify a transform (up to 6 DOF)
geometry_msgs/Transform[] transforms

# There can be a velocity specified for the origin of the joint 
geometry_msgs/Twist[] velocities

# There can be an acceleration specified for the origin of the joint 
geometry_msgs/Twist[] accelerations

duration time_from_start
`,
		"trajectory_msgs/MultiDOFJointTrajectoryPoint",
		"3ebe08d1abd5b65862d50e09430db776",
	}
)

type MultiDOFJointTrajectoryPoint struct {
	Transforms    []geometry_msgs.Transform
	Velocities    []geometry_msgs.Twist
	Accelerations []geometry_msgs.Twist
	TimeFromStart ros.Duration
}

func (m *MultiDOFJointTrajectoryPoint) Serialize(w io.Writer) (err error) {
	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Transforms)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Transforms {
		if err = ros.SerializeMessageField(w, "geometry_msgs/Transform", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Velocities)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Velocities {
		if err = ros.SerializeMessageField(w, "geometry_msgs/Twist", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Accelerations)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Accelerations {
		if err = ros.SerializeMessageField(w, "geometry_msgs/Twist", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "duration", &m.TimeFromStart); err != nil {
		return err
	}

	return
}

func (m *MultiDOFJointTrajectoryPoint) Deserialize(r io.Reader) (err error) {
	// Transforms
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Transforms: %s", err)
		}
		m.Transforms = make([]geometry_msgs.Transform, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "geometry_msgs/Transform", &m.Transforms[i]); err != nil {
				return err
			}
		}
	}

	// Velocities
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Velocities: %s", err)
		}
		m.Velocities = make([]geometry_msgs.Twist, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "geometry_msgs/Twist", &m.Velocities[i]); err != nil {
				return err
			}
		}
	}

	// Accelerations
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Accelerations: %s", err)
		}
		m.Accelerations = make([]geometry_msgs.Twist, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "geometry_msgs/Twist", &m.Accelerations[i]); err != nil {
				return err
			}
		}
	}

	// TimeFromStart
	if err = ros.DeserializeMessageField(r, "duration", &m.TimeFromStart); err != nil {
		return err
	}

	return
}
