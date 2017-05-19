package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event event
// swagger:model Event
type Event struct {

	// Address of the event venue
	Address string `json:"address,omitempty"`

	// district
	District *District `json:"district,omitempty"`

	// When the event ends
	EndDate strfmt.Date `json:"end_date,omitempty"`

	// Event short code, as provided by FIRST
	// Required: true
	EventCode *string `json:"event_code"`

	// An integer that represents the event type as a constant.
	EventType int64 `json:"event_type,omitempty"`

	// A human readable string that defines the event type.
	EventTypeString string `json:"event_type_string,omitempty"`

	// URL for the venue on Google Maps
	GmapsURL string `json:"gmaps_url,omitempty"`

	// TBA event key with the format yyyy[EVENT_CODE], where yyyy is the year, and EVENT_CODE is the event code of the event.
	// Required: true
	Key *string `json:"key"`

	// Short name of the venue
	LocationName string `json:"location_name,omitempty"`

	// Official name of event on record either provided by FIRST or organizers of offseason event.
	// Required: true
	Name *string `json:"name"`

	// Same as name but doesn't include event specifiers, such as 'Regional' or 'District'. May be null.
	ShortName string `json:"short_name,omitempty"`

	// When the event starts
	StartDate strfmt.Date `json:"start_date,omitempty"`

	// Timezone name
	Timezone string `json:"timezone,omitempty"`

	// If the event has webcast data associated with it, this contains JSON data of the streams
	Webcasts string `json:"webcasts,omitempty"`

	// The event's website, if any.
	Website string `json:"website,omitempty"`

	// Week of the season the event occurs on
	Week int64 `json:"week,omitempty"`

	// Year the event data is for.
	// Required: true
	Year *int64 `json:"year"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistrict(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateYear(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateDistrict(formats strfmt.Registry) error {

	if swag.IsZero(m.District) { // not required
		return nil
	}

	if m.District != nil {

		if err := m.District.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("district")
			}
			return err
		}
	}

	return nil
}

func (m *Event) validateEventCode(formats strfmt.Registry) error {

	if err := validate.Required("event_code", "body", m.EventCode); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateYear(formats strfmt.Registry) error {

	if err := validate.Required("year", "body", m.Year); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}