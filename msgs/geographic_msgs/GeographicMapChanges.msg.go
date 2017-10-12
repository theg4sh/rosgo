// Code generated by ros-gen-go.
// source: GeographicMapChanges.msg
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

type _MsgGeographicMapChanges struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGeographicMapChanges) Text() string {
	return t.text
}

func (t *_MsgGeographicMapChanges) Name() string {
	return t.name
}

func (t *_MsgGeographicMapChanges) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGeographicMapChanges) NewMessage() ros.Message {
	m := new(GeographicMapChanges)

	return m
}

var (
	MsgGeographicMapChanges = &_MsgGeographicMapChanges{
		`# A list of geographic map changes.

Header header                   # stamp specifies time of change
                                # frame_id (normally /map)

GeographicMap diffs             # new and changed points and features
uuid_msgs/UniqueID[] deletes    # deleted map components
`,
		"geographic_msgs/GeographicMapChanges",
		"4fd027f54298203ec12aa1c4b20e6cb8",
	}
)

type GeographicMapChanges struct {
	Header  std_msgs.Header
	Diffs   GeographicMap
	Deletes []uuid_msgs.UniqueID
}

func (m *GeographicMapChanges) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "GeographicMap", &m.Diffs); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Deletes)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Deletes {
		if err = ros.SerializeMessageField(w, "uuid_msgs/UniqueID", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *GeographicMapChanges) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Diffs
	if err = ros.DeserializeMessageField(r, "GeographicMap", &m.Diffs); err != nil {
		return err
	}

	// Deletes
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Deletes: %s", err)
		}
		m.Deletes = make([]uuid_msgs.UniqueID, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "uuid_msgs/UniqueID", &m.Deletes[i]); err != nil {
				return err
			}
		}
	}

	return
}
