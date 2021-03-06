// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FacilityUpdateRequest facility update request
//
// swagger:model FacilityUpdateRequest
type FacilityUpdateRequest []*FacilityUpdateRequestItems0

// Validate validates this facility update request
func (m FacilityUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// FacilityUpdateRequestItems0 facility update request items0
//
// swagger:model FacilityUpdateRequestItems0
type FacilityUpdateRequestItems0 struct {

	// Address
	Address *Address `json:"address,omitempty"`

	// Category
	Category string `json:"category,omitempty"`

	// Contact number
	Contact string `json:"contact,omitempty"`

	// Facility Email
	Email string `json:"email,omitempty"`

	// Facility Name
	FacilityName string `json:"facilityName,omitempty"`

	// Geo Location
	GeoLocation string `json:"geoLocation,omitempty"`

	// Operating hours end of day
	OperatingHourEnd string `json:"operatingHourEnd,omitempty"`

	// Operating hours start of day
	OperatingHourStart string `json:"operatingHourStart,omitempty"`

	// osid
	Osid string `json:"osid,omitempty"`

	// programs
	Programs []*FacilityUpdateRequestItems0ProgramsItems0 `json:"programs"`

	// status
	Status string `json:"status,omitempty"`

	// Website URL
	WebsiteURL string `json:"websiteUrl,omitempty"`
}

// Validate validates this facility update request items0
func (m *FacilityUpdateRequestItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrograms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacilityUpdateRequestItems0) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *FacilityUpdateRequestItems0) validatePrograms(formats strfmt.Registry) error {

	if swag.IsZero(m.Programs) { // not required
		return nil
	}

	for i := 0; i < len(m.Programs); i++ {
		if swag.IsZero(m.Programs[i]) { // not required
			continue
		}

		if m.Programs[i] != nil {
			if err := m.Programs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("programs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacilityUpdateRequestItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityUpdateRequestItems0) UnmarshalBinary(b []byte) error {
	var res FacilityUpdateRequestItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FacilityUpdateRequestItems0ProgramsItems0 facility update request items0 programs items0
//
// swagger:model FacilityUpdateRequestItems0ProgramsItems0
type FacilityUpdateRequestItems0ProgramsItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// rate
	Rate float64 `json:"rate,omitempty"`

	// schedule
	Schedule *FacilityUpdateRequestItems0ProgramsItems0Schedule `json:"schedule,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this facility update request items0 programs items0
func (m *FacilityUpdateRequestItems0ProgramsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacilityUpdateRequestItems0ProgramsItems0) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacilityUpdateRequestItems0ProgramsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityUpdateRequestItems0ProgramsItems0) UnmarshalBinary(b []byte) error {
	var res FacilityUpdateRequestItems0ProgramsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FacilityUpdateRequestItems0ProgramsItems0Schedule facility update request items0 programs items0 schedule
//
// swagger:model FacilityUpdateRequestItems0ProgramsItems0Schedule
type FacilityUpdateRequestItems0ProgramsItems0Schedule struct {

	// days
	Days []string `json:"days"`

	// end time
	EndTime string `json:"endTime,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`
}

// Validate validates this facility update request items0 programs items0 schedule
func (m *FacilityUpdateRequestItems0ProgramsItems0Schedule) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FacilityUpdateRequestItems0ProgramsItems0Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityUpdateRequestItems0ProgramsItems0Schedule) UnmarshalBinary(b []byte) error {
	var res FacilityUpdateRequestItems0ProgramsItems0Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
