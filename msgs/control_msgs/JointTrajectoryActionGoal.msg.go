// Code generated by ros-gen-go.
// source: JointTrajectoryActionGoal.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/actionlib_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgJointTrajectoryActionGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJointTrajectoryActionGoal) Text() string {
	return t.text
}

func (t *_MsgJointTrajectoryActionGoal) Name() string {
	return t.name
}

func (t *_MsgJointTrajectoryActionGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJointTrajectoryActionGoal) NewMessage() ros.Message {
	m := new(JointTrajectoryActionGoal)

	return m
}

var (
	MsgJointTrajectoryActionGoal = &_MsgJointTrajectoryActionGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalID goal_id
JointTrajectoryGoal goal
`,
		"control_msgs/JointTrajectoryActionGoal",
		"a99e83ef6185f9fdd7693efe99623a86",
	}
)

type JointTrajectoryActionGoal struct {
	Header std_msgs.Header
	GoalId actionlib_msgs.GoalID
	Goal   JointTrajectoryGoal
}

func (m *JointTrajectoryActionGoal) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalID", &m.GoalId); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "JointTrajectoryGoal", &m.Goal); err != nil {
		return err
	}

	return
}

func (m *JointTrajectoryActionGoal) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// GoalId
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalID", &m.GoalId); err != nil {
		return err
	}

	// Goal
	if err = ros.DeserializeMessageField(r, "JointTrajectoryGoal", &m.Goal); err != nil {
		return err
	}

	return
}
