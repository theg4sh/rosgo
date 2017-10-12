// Code generated by ros-gen-go.
// source: HomePosition.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/geographic_msgs"
	"github.com/cnord/rosgo/msgs/geometry_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgHomePosition struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgHomePosition) Text() string {
	return t.text
}

func (t *_MsgHomePosition) Name() string {
	return t.name
}

func (t *_MsgHomePosition) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgHomePosition) NewMessage() ros.Message {
	m := new(HomePosition)

	return m
}

var (
	MsgHomePosition = &_MsgHomePosition{
		`# MAVLink message: HOME_POSITION
# http://mavlink.org/messages/common#HOME_POSITION

std_msgs/Header header

geographic_msgs/GeoPoint geo # geodetic coordinates in WGS-84 datum

geometry_msgs/Point position	# local position
geometry_msgs/Quaternion orientation	# XXX: verify field name (q[4])
geometry_msgs/Vector3 approach	# position of the end of approach vector
`,
		"mavros_msgs/HomePosition",
		"c1167922de8c97acdb0ec714c1dba774",
	}
)

type HomePosition struct {
	Header      std_msgs.Header
	Geo         geographic_msgs.GeoPoint
	Position    geometry_msgs.Point
	Orientation geometry_msgs.Quaternion
	Approach    geometry_msgs.Vector3
}

func (m *HomePosition) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geographic_msgs/GeoPoint", &m.Geo); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Point", &m.Position); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Quaternion", &m.Orientation); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.Approach); err != nil {
		return err
	}

	return
}

func (m *HomePosition) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// Geo
	if err = ros.DeserializeMessageField(r, "geographic_msgs/GeoPoint", &m.Geo); err != nil {
		return err
	}

	// Position
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Point", &m.Position); err != nil {
		return err
	}

	// Orientation
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Quaternion", &m.Orientation); err != nil {
		return err
	}

	// Approach
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.Approach); err != nil {
		return err
	}

	return
}
