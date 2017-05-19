package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Award award
// swagger:model Award
type Award struct {

	// An integer that represents the award type as a constant.
	// Required: true
	AwardType *int64 `json:"award_type"`

	// The event_key of the event the award was won at.
	// Required: true
	EventKey *string `json:"event_key"`

	// The name of the award as provided by FIRST. May vary for the same award type.
	// Required: true
	Name *string `json:"name"`

	// recipient list
	RecipientList []*AwardRecipientListItems0 `json:"recipient_list"`

	// The year this award was won.
	// Required: true
	Year *int64 `json:"year"`
}

// Validate validates this award
func (m *Award) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwardType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipientList(formats); err != nil {
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

func (m *Award) validateAwardType(formats strfmt.Registry) error {

	if err := validate.Required("award_type", "body", m.AwardType); err != nil {
		return err
	}

	return nil
}

func (m *Award) validateEventKey(formats strfmt.Registry) error {

	if err := validate.Required("event_key", "body", m.EventKey); err != nil {
		return err
	}

	return nil
}

func (m *Award) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Award) validateRecipientList(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipientList) { // not required
		return nil
	}

	for i := 0; i < len(m.RecipientList); i++ {

		if swag.IsZero(m.RecipientList[i]) { // not required
			continue
		}

		if m.RecipientList[i] != nil {

			if err := m.RecipientList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipient_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Award) validateYear(formats strfmt.Registry) error {

	if err := validate.Required("year", "body", m.Year); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Award) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Award) UnmarshalBinary(b []byte) error {
	var res Award
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwardRecipientListItems0 AwardRecipient
// swagger:model AwardRecipientListItems0
type AwardRecipientListItems0 struct {

	// Name of the winning person
	Awardee string `json:"awardee,omitempty"`

	// Key of the winning team
	TeamKey string `json:"team_key,omitempty"`
}

// Validate validates this award recipient list items0
func (m *AwardRecipientListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwardRecipientListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwardRecipientListItems0) UnmarshalBinary(b []byte) error {
	var res AwardRecipientListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}