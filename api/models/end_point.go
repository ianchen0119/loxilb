// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndPoint end point
//
// swagger:model EndPoint
type EndPoint struct {

	// Host Description
	Description string `json:"description,omitempty"`

	// Host name in CIDR
	HostName string `json:"hostName,omitempty"`

	// Number of inactive retries
	InactiveReTries int64 `json:"inactiveReTries,omitempty"`

	// How frequently to probe in seconds
	ProbeDuration int64 `json:"probeDuration,omitempty"`

	// The l4port to probe on
	ProbePort int64 `json:"probePort,omitempty"`

	// URI for http/https probes
	ProbeReq string `json:"probeReq,omitempty"`

	// Response for http/https probes
	ProbeResp string `json:"probeResp,omitempty"`

	// Type of probe used
	ProbeType string `json:"probeType,omitempty"`
}

// Validate validates this end point
func (m *EndPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this end point based on context it is used
func (m *EndPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndPoint) UnmarshalBinary(b []byte) error {
	var res EndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
