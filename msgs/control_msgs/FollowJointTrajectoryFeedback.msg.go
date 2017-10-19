// Code generated by ros-gen-go.
// source: FollowJointTrajectoryFeedback.msg
// DO NOT EDIT!
package control_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/msgs/trajectory_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgFollowJointTrajectoryFeedback struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFollowJointTrajectoryFeedback) Text() string {
	return t.text
}

func (t *_MsgFollowJointTrajectoryFeedback) Name() string {
	return t.name
}

func (t *_MsgFollowJointTrajectoryFeedback) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFollowJointTrajectoryFeedback) NewMessage() ros.Message {
	m := new(FollowJointTrajectoryFeedback)

	return m
}

var (
	MsgFollowJointTrajectoryFeedback = &_MsgFollowJointTrajectoryFeedback{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
Header header
string[] joint_names
trajectory_msgs/JointTrajectoryPoint desired
trajectory_msgs/JointTrajectoryPoint actual
trajectory_msgs/JointTrajectoryPoint error

`,
		"control_msgs/FollowJointTrajectoryFeedback",
		"10817c60c2486ef6b33e97dcd87f4474",
	}
)

type FollowJointTrajectoryFeedback struct {
	Header     std_msgs.Header
	JointNames []string
	Desired    trajectory_msgs.JointTrajectoryPoint
	Actual     trajectory_msgs.JointTrajectoryPoint
	Error      trajectory_msgs.JointTrajectoryPoint
}

func (m *FollowJointTrajectoryFeedback) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.JointNames)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.JointNames {
		if err = ros.SerializeMessageField(w, "string", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "trajectory_msgs/JointTrajectoryPoint", &m.Desired); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "trajectory_msgs/JointTrajectoryPoint", &m.Actual); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "trajectory_msgs/JointTrajectoryPoint", &m.Error); err != nil {
		return err
	}

	return
}

func (m *FollowJointTrajectoryFeedback) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// JointNames
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for JointNames: %s", err)
		}
		m.JointNames = make([]string, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "string", &m.JointNames[i]); err != nil {
				return err
			}
		}
	}

	// Desired
	if err = ros.DeserializeMessageField(r, "trajectory_msgs/JointTrajectoryPoint", &m.Desired); err != nil {
		return err
	}

	// Actual
	if err = ros.DeserializeMessageField(r, "trajectory_msgs/JointTrajectoryPoint", &m.Actual); err != nil {
		return err
	}

	// Error
	if err = ros.DeserializeMessageField(r, "trajectory_msgs/JointTrajectoryPoint", &m.Error); err != nil {
		return err
	}

	return
}
