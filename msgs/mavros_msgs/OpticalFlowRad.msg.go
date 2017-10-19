// Code generated by ros-gen-go.
// source: OpticalFlowRad.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgOpticalFlowRad struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgOpticalFlowRad) Text() string {
	return t.text
}

func (t *_MsgOpticalFlowRad) Name() string {
	return t.name
}

func (t *_MsgOpticalFlowRad) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgOpticalFlowRad) NewMessage() ros.Message {
	m := new(OpticalFlowRad)

	return m
}

var (
	MsgOpticalFlowRad = &_MsgOpticalFlowRad{
		`# OPTICAL_FLOW_RAD message data

std_msgs/Header header

uint32 integration_time_us
float32 integrated_x
float32 integrated_y
float32 integrated_xgyro
float32 integrated_ygyro
float32 integrated_zgyro
int16 temperature
uint8 quality
uint32 time_delta_distance_us
float32 distance
`,
		"mavros_msgs/OpticalFlowRad",
		"65d93e03c6188c7ee30415b2a39ad40d",
	}
)

type OpticalFlowRad struct {
	Header              std_msgs.Header
	IntegrationTimeUs   uint32
	IntegratedX         float32
	IntegratedY         float32
	IntegratedXgyro     float32
	IntegratedYgyro     float32
	IntegratedZgyro     float32
	Temperature         int16
	Quality             uint8
	TimeDeltaDistanceUs uint32
	Distance            float32
}

func (m *OpticalFlowRad) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.IntegrationTimeUs); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.IntegratedX); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.IntegratedY); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.IntegratedXgyro); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.IntegratedYgyro); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.IntegratedZgyro); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int16", &m.Temperature); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Quality); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.TimeDeltaDistanceUs); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Distance); err != nil {
		return err
	}

	return
}

func (m *OpticalFlowRad) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// IntegrationTimeUs
	if err = ros.DeserializeMessageField(r, "uint32", &m.IntegrationTimeUs); err != nil {
		return err
	}

	// IntegratedX
	if err = ros.DeserializeMessageField(r, "float32", &m.IntegratedX); err != nil {
		return err
	}

	// IntegratedY
	if err = ros.DeserializeMessageField(r, "float32", &m.IntegratedY); err != nil {
		return err
	}

	// IntegratedXgyro
	if err = ros.DeserializeMessageField(r, "float32", &m.IntegratedXgyro); err != nil {
		return err
	}

	// IntegratedYgyro
	if err = ros.DeserializeMessageField(r, "float32", &m.IntegratedYgyro); err != nil {
		return err
	}

	// IntegratedZgyro
	if err = ros.DeserializeMessageField(r, "float32", &m.IntegratedZgyro); err != nil {
		return err
	}

	// Temperature
	if err = ros.DeserializeMessageField(r, "int16", &m.Temperature); err != nil {
		return err
	}

	// Quality
	if err = ros.DeserializeMessageField(r, "uint8", &m.Quality); err != nil {
		return err
	}

	// TimeDeltaDistanceUs
	if err = ros.DeserializeMessageField(r, "uint32", &m.TimeDeltaDistanceUs); err != nil {
		return err
	}

	// Distance
	if err = ros.DeserializeMessageField(r, "float32", &m.Distance); err != nil {
		return err
	}

	return
}
