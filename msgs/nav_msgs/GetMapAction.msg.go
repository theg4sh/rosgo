// Code generated by ros-gen-go.
// source: GetMapAction.msg
// DO NOT EDIT!
package nav_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgGetMapAction struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGetMapAction) Text() string {
	return t.text
}

func (t *_MsgGetMapAction) Name() string {
	return t.name
}

func (t *_MsgGetMapAction) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGetMapAction) NewMessage() ros.Message {
	m := new(GetMapAction)

	return m
}

var (
	MsgGetMapAction = &_MsgGetMapAction{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

GetMapActionGoal action_goal
GetMapActionResult action_result
GetMapActionFeedback action_feedback
`,
		"nav_msgs/GetMapAction",
		"e611ad23fbf237c031b7536416dc7cd7",
	}
)

type GetMapAction struct {
	ActionGoal     GetMapActionGoal
	ActionResult   GetMapActionResult
	ActionFeedback GetMapActionFeedback
}

func (m *GetMapAction) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "GetMapActionGoal", &m.ActionGoal); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "GetMapActionResult", &m.ActionResult); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "GetMapActionFeedback", &m.ActionFeedback); err != nil {
		return err
	}

	return
}

func (m *GetMapAction) Deserialize(r io.Reader) (err error) {
	// ActionGoal
	if err = ros.DeserializeMessageField(r, "GetMapActionGoal", &m.ActionGoal); err != nil {
		return err
	}

	// ActionResult
	if err = ros.DeserializeMessageField(r, "GetMapActionResult", &m.ActionResult); err != nil {
		return err
	}

	// ActionFeedback
	if err = ros.DeserializeMessageField(r, "GetMapActionFeedback", &m.ActionFeedback); err != nil {
		return err
	}

	return
}
