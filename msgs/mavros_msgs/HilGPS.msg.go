// Code generated by ros-gen-go.
// source: HilGPS.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/geographic_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgHilGPS struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgHilGPS) Text() string {
	return t.text
}

func (t *_MsgHilGPS) Name() string {
	return t.name
}

func (t *_MsgHilGPS) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgHilGPS) NewMessage() ros.Message {
	m := new(HilGPS)

	return m
}

var (
	MsgHilGPS = &_MsgHilGPS{
		`# HilControls.msg
#
# ROS representation of MAVLink HIL_GPS
# See mavlink message documentation here:
# https://pixhawk.ethz.ch/mavlink/#HIL_GPS

std_msgs/Header header
uint8 fix_type
geographic_msgs/GeoPoint geo
uint16 eph
uint16 epv
uint16 vel
int16 vn
int16 ve
int16 vd
uint16 cog
uint8 satellites_visible
`,
		"mavros_msgs/HilGPS",
		"313b3baf2319db196fa18376a4900a7b",
	}
)

type HilGPS struct {
	Header            std_msgs.Header
	FixType           uint8
	Geo               geographic_msgs.GeoPoint
	Eph               uint16
	Epv               uint16
	Vel               uint16
	Vn                int16
	Ve                int16
	Vd                int16
	Cog               uint16
	SatellitesVisible uint8
}

func (m *HilGPS) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.FixType); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geographic_msgs/GeoPoint", &m.Geo); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.Eph); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.Epv); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.Vel); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int16", &m.Vn); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int16", &m.Ve); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int16", &m.Vd); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.Cog); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.SatellitesVisible); err != nil {
		return err
	}

	return
}

func (m *HilGPS) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// FixType
	if err = ros.DeserializeMessageField(r, "uint8", &m.FixType); err != nil {
		return err
	}

	// Geo
	if err = ros.DeserializeMessageField(r, "geographic_msgs/GeoPoint", &m.Geo); err != nil {
		return err
	}

	// Eph
	if err = ros.DeserializeMessageField(r, "uint16", &m.Eph); err != nil {
		return err
	}

	// Epv
	if err = ros.DeserializeMessageField(r, "uint16", &m.Epv); err != nil {
		return err
	}

	// Vel
	if err = ros.DeserializeMessageField(r, "uint16", &m.Vel); err != nil {
		return err
	}

	// Vn
	if err = ros.DeserializeMessageField(r, "int16", &m.Vn); err != nil {
		return err
	}

	// Ve
	if err = ros.DeserializeMessageField(r, "int16", &m.Ve); err != nil {
		return err
	}

	// Vd
	if err = ros.DeserializeMessageField(r, "int16", &m.Vd); err != nil {
		return err
	}

	// Cog
	if err = ros.DeserializeMessageField(r, "uint16", &m.Cog); err != nil {
		return err
	}

	// SatellitesVisible
	if err = ros.DeserializeMessageField(r, "uint8", &m.SatellitesVisible); err != nil {
		return err
	}

	return
}
