// Code generated by ros-gen-go.
// source: RouteNetwork.msg
// DO NOT EDIT!
package geographic_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/msgs/uuid_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgRouteNetwork struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgRouteNetwork) Text() string {
	return t.text
}

func (t *_MsgRouteNetwork) Name() string {
	return t.name
}

func (t *_MsgRouteNetwork) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgRouteNetwork) NewMessage() ros.Message {
	m := new(RouteNetwork)

	return m
}

var (
	MsgRouteNetwork = &_MsgRouteNetwork{
		`# Geographic map route network.
#
# A directed graph of WayPoint nodes and RouteSegment edges.  This
# information is extracted from the more-detailed contents of a
# GeographicMap.  A RouteNetwork contains only the way points and
# route segments of interest for path planning.

Header          header

uuid_msgs/UniqueID id    # This route network identifier
BoundingBox     bounds   # 2D bounding box for network

WayPoint[]      points   # Way points in this network
RouteSegment[]  segments # Directed edges of this network

KeyValue[]      props    # Network key/value properties
`,
		"geographic_msgs/RouteNetwork",
		"fd717c0a34a7c954deed32c6847f30a8",
	}
)

type RouteNetwork struct {
	Header   std_msgs.Header
	Id       uuid_msgs.UniqueID
	Bounds   BoundingBox
	Points   []WayPoint
	Segments []RouteSegment
	Props    []KeyValue
}

func (m *RouteNetwork) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uuid_msgs/UniqueID", &m.Id); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "BoundingBox", &m.Bounds); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Points)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Points {
		if err = ros.SerializeMessageField(w, "WayPoint", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Segments)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Segments {
		if err = ros.SerializeMessageField(w, "RouteSegment", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Props)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Props {
		if err = ros.SerializeMessageField(w, "KeyValue", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *RouteNetwork) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Id
	if err = ros.DeserializeMessageField(r, "uuid_msgs/UniqueID", &m.Id); err != nil {
		return err
	}

	// Bounds
	if err = ros.DeserializeMessageField(r, "BoundingBox", &m.Bounds); err != nil {
		return err
	}

	// Points
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Points: %s", err)
		}
		m.Points = make([]WayPoint, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "WayPoint", &m.Points[i]); err != nil {
				return err
			}
		}
	}

	// Segments
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Segments: %s", err)
		}
		m.Segments = make([]RouteSegment, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "RouteSegment", &m.Segments[i]); err != nil {
				return err
			}
		}
	}

	// Props
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Props: %s", err)
		}
		m.Props = make([]KeyValue, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "KeyValue", &m.Props[i]); err != nil {
				return err
			}
		}
	}

	return
}
