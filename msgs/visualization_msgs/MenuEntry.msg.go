// Code generated by ros-gen-go.
// source: MenuEntry.msg
// DO NOT EDIT!
package visualization_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgMenuEntry struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMenuEntry) Text() string {
	return t.text
}

func (t *_MsgMenuEntry) Name() string {
	return t.name
}

func (t *_MsgMenuEntry) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMenuEntry) NewMessage() ros.Message {
	m := new(MenuEntry)

	return m
}

var (
	MsgMenuEntry = &_MsgMenuEntry{
		`# MenuEntry message.

# Each InteractiveMarker message has an array of MenuEntry messages.
# A collection of MenuEntries together describe a
# menu/submenu/subsubmenu/etc tree, though they are stored in a flat
# array.  The tree structure is represented by giving each menu entry
# an ID number and a "parent_id" field.  Top-level entries are the
# ones with parent_id = 0.  Menu entries are ordered within their
# level the same way they are ordered in the containing array.  Parent
# entries must appear before their children.

# Example:
# - id = 3
#   parent_id = 0
#   title = "fun"
# - id = 2
#   parent_id = 0
#   title = "robot"
# - id = 4
#   parent_id = 2
#   title = "pr2"
# - id = 5
#   parent_id = 2
#   title = "turtle"
#
# Gives a menu tree like this:
#  - fun
#  - robot
#    - pr2
#    - turtle

# ID is a number for each menu entry.  Must be unique within the
# control, and should never be 0.
uint32 id

# ID of the parent of this menu entry, if it is a submenu.  If this
# menu entry is a top-level entry, set parent_id to 0.
uint32 parent_id

# menu / entry title
string title

# Arguments to command indicated by command_type (below)
string command

# Command_type stores the type of response desired when this menu
# entry is clicked.
# FEEDBACK: send an InteractiveMarkerFeedback message with menu_entry_id set to this entry's id.
# ROSRUN: execute "rosrun" with arguments given in the command field (above).
# ROSLAUNCH: execute "roslaunch" with arguments given in the command field (above).
uint8 FEEDBACK=0
uint8 ROSRUN=1
uint8 ROSLAUNCH=2
uint8 command_type
`,
		"visualization_msgs/MenuEntry",
		"b90ec63024573de83b57aa93eb39be2d",
	}
)

type MenuEntry struct {
	Id          uint32
	ParentId    uint32
	Title       string
	Command     string
	CommandType uint8
}

func (m *MenuEntry) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint32", &m.Id); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.ParentId); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Title); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Command); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.CommandType); err != nil {
		return err
	}

	return
}

func (m *MenuEntry) Deserialize(r io.Reader) (err error) {
	// Id
	if err = ros.DeserializeMessageField(r, "uint32", &m.Id); err != nil {
		return err
	}

	// ParentId
	if err = ros.DeserializeMessageField(r, "uint32", &m.ParentId); err != nil {
		return err
	}

	// Title
	if err = ros.DeserializeMessageField(r, "string", &m.Title); err != nil {
		return err
	}

	// Command
	if err = ros.DeserializeMessageField(r, "string", &m.Command); err != nil {
		return err
	}

	// CommandType
	if err = ros.DeserializeMessageField(r, "uint8", &m.CommandType); err != nil {
		return err
	}

	return
}
