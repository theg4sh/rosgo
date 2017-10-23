// Code generated by ros-gen-go.
// source: WaypointPush.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvWaypointPush struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvWaypointPush) Name() string                  { return t.name }
func (t *_SrvWaypointPush) MD5Sum() string                { return t.md5sum }
func (t *_SrvWaypointPush) Text() string                  { return t.text }
func (t *_SrvWaypointPush) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvWaypointPush) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvWaypointPush) NewService() ros.Service {
	return new(WaypointPush)
}

var (
	SrvWaypointPush = &_SrvWaypointPush{
		"mavros_msgs/WaypointPush",
		"5b2bca74e830798e1579b58cafd50527",
		`# Send waypoints to device
#
#  :start_index: will define a partial waypoint update. Set to 0 for full update
#
# Returns success status and transfered count

uint16 start_index
mavros_msgs/Waypoint[] waypoints
---
bool success
uint32 wp_transfered
`,
		MsgWaypointPushRequest,
		MsgWaypointPushResponse,
	}
)

type WaypointPush struct {
	Request  WaypointPushRequest
	Response WaypointPushResponse
}

func (s *WaypointPush) ReqMessage() ros.Message { return &s.Request }
func (s *WaypointPush) ResMessage() ros.Message { return &s.Response }

// WaypointPushRequest

type _MsgWaypointPushRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointPushRequest) Text() string {
	return t.text
}

func (t *_MsgWaypointPushRequest) Name() string {
	return t.name
}

func (t *_MsgWaypointPushRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointPushRequest) NewMessage() ros.Message {
	m := new(WaypointPushRequest)

	return m
}

var (
	MsgWaypointPushRequest = &_MsgWaypointPushRequest{
		`# Send waypoints to device
#
#  :start_index: will define a partial waypoint update. Set to 0 for full update
#
# Returns success status and transfered count

uint16 start_index
mavros_msgs/Waypoint[] waypoints
`,
		"mavros_msgs/WaypointPushRequest",
		"5b2bca74e830798e1579b58cafd50527",
	}
)

type WaypointPushRequest struct {
	StartIndex uint16
	Waypoints  []Waypoint
}

func (m *WaypointPushRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint16", &m.StartIndex); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Waypoints)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Waypoints {
		if err = ros.SerializeMessageField(w, "mavros_msgs/Waypoint", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *WaypointPushRequest) Deserialize(r io.Reader) (err error) {
	// StartIndex
	if err = ros.DeserializeMessageField(r, "uint16", &m.StartIndex); err != nil {
		return err
	}

	// Waypoints
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Waypoints: %s", err)
		}
		m.Waypoints = make([]Waypoint, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "mavros_msgs/Waypoint", &m.Waypoints[i]); err != nil {
				return err
			}
		}
	}

	return
}

// WaypointPushResponse

type _MsgWaypointPushResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointPushResponse) Text() string {
	return t.text
}

func (t *_MsgWaypointPushResponse) Name() string {
	return t.name
}

func (t *_MsgWaypointPushResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointPushResponse) NewMessage() ros.Message {
	m := new(WaypointPushResponse)

	return m
}

var (
	MsgWaypointPushResponse = &_MsgWaypointPushResponse{
		`
bool success
uint32 wp_transfered
`,
		"mavros_msgs/WaypointPushResponse",
		"5b2bca74e830798e1579b58cafd50527",
	}
)

type WaypointPushResponse struct {
	Success      bool
	WpTransfered uint32
}

func (m *WaypointPushResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.WpTransfered); err != nil {
		return err
	}

	return
}

func (m *WaypointPushResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// WpTransfered
	if err = ros.DeserializeMessageField(r, "uint32", &m.WpTransfered); err != nil {
		return err
	}

	return
}