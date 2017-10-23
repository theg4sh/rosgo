// Code generated by ros-gen-go.
// source: WaypointPull.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvWaypointPull struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvWaypointPull) Name() string                  { return t.name }
func (t *_SrvWaypointPull) MD5Sum() string                { return t.md5sum }
func (t *_SrvWaypointPull) Text() string                  { return t.text }
func (t *_SrvWaypointPull) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvWaypointPull) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvWaypointPull) NewService() ros.Service {
	return new(WaypointPull)
}

var (
	SrvWaypointPull = &_SrvWaypointPull{
		"mavros_msgs/WaypointPull",
		"a8d9ecef8fb37028d2db2a9aa4ed7e79",
		`# Requests waypoints from device
#
# Returns success status and received count

---
bool success
uint32 wp_received
`,
		MsgWaypointPullRequest,
		MsgWaypointPullResponse,
	}
)

type WaypointPull struct {
	Request  WaypointPullRequest
	Response WaypointPullResponse
}

func (s *WaypointPull) ReqMessage() ros.Message { return &s.Request }
func (s *WaypointPull) ResMessage() ros.Message { return &s.Response }

// WaypointPullRequest

type _MsgWaypointPullRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointPullRequest) Text() string {
	return t.text
}

func (t *_MsgWaypointPullRequest) Name() string {
	return t.name
}

func (t *_MsgWaypointPullRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointPullRequest) NewMessage() ros.Message {
	m := new(WaypointPullRequest)

	return m
}

var (
	MsgWaypointPullRequest = &_MsgWaypointPullRequest{
		`# Requests waypoints from device
#
# Returns success status and received count

`,
		"mavros_msgs/WaypointPullRequest",
		"a8d9ecef8fb37028d2db2a9aa4ed7e79",
	}
)

type WaypointPullRequest struct {
}

func (m *WaypointPullRequest) Serialize(w io.Writer) (err error) {
	return
}

func (m *WaypointPullRequest) Deserialize(r io.Reader) (err error) {
	return
}

// WaypointPullResponse

type _MsgWaypointPullResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointPullResponse) Text() string {
	return t.text
}

func (t *_MsgWaypointPullResponse) Name() string {
	return t.name
}

func (t *_MsgWaypointPullResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointPullResponse) NewMessage() ros.Message {
	m := new(WaypointPullResponse)

	return m
}

var (
	MsgWaypointPullResponse = &_MsgWaypointPullResponse{
		`
bool success
uint32 wp_received
`,
		"mavros_msgs/WaypointPullResponse",
		"a8d9ecef8fb37028d2db2a9aa4ed7e79",
	}
)

type WaypointPullResponse struct {
	Success    bool
	WpReceived uint32
}

func (m *WaypointPullResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.WpReceived); err != nil {
		return err
	}

	return
}

func (m *WaypointPullResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// WpReceived
	if err = ros.DeserializeMessageField(r, "uint32", &m.WpReceived); err != nil {
		return err
	}

	return
}
