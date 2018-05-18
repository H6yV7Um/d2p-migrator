// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageInfo An object containing all details of an image at API side
// swagger:model ImageInfo

type ImageInfo struct {

	// the CPU architecture.
	Architecture string `json:"Architecture,omitempty"`

	// config
	Config *ContainerConfig `json:"Config,omitempty"`

	// time of image creation.
	CreatedAt string `json:"CreatedAt,omitempty"`

	// ID of an image.
	ID string `json:"Id,omitempty"`

	// the name of the operating system.
	Os string `json:"Os,omitempty"`

	// repository with digest.
	RepoDigests []string `json:"RepoDigests"`

	// repository with tag.
	RepoTags []string `json:"RepoTags"`

	// root f s
	RootFS *ImageInfoRootFS `json:"RootFS,omitempty"`

	// size of image's taking disk space.
	Size int64 `json:"Size,omitempty"`
}

/* polymorph ImageInfo Architecture false */

/* polymorph ImageInfo Config false */

/* polymorph ImageInfo CreatedAt false */

/* polymorph ImageInfo Id false */

/* polymorph ImageInfo Os false */

/* polymorph ImageInfo RepoDigests false */

/* polymorph ImageInfo RepoTags false */

/* polymorph ImageInfo RootFS false */

/* polymorph ImageInfo Size false */

// Validate validates this image info
func (m *ImageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepoDigests(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepoTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRootFS(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageInfo) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {

		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			}
			return err
		}
	}

	return nil
}

func (m *ImageInfo) validateRepoDigests(formats strfmt.Registry) error {

	if swag.IsZero(m.RepoDigests) { // not required
		return nil
	}

	return nil
}

func (m *ImageInfo) validateRepoTags(formats strfmt.Registry) error {

	if swag.IsZero(m.RepoTags) { // not required
		return nil
	}

	return nil
}

func (m *ImageInfo) validateRootFS(formats strfmt.Registry) error {

	if swag.IsZero(m.RootFS) { // not required
		return nil
	}

	if m.RootFS != nil {

		if err := m.RootFS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RootFS")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageInfo) UnmarshalBinary(b []byte) error {
	var res ImageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ImageInfoRootFS the rootfs key references the layer content addresses used by the image.
// swagger:model ImageInfoRootFS

type ImageInfoRootFS struct {

	// the base layer content hash.
	BaseLayer string `json:"BaseLayer,omitempty"`

	// an array of layer content hashes
	Layers []string `json:"Layers"`

	// type of the rootfs
	// Required: true
	Type string `json:"Type"`
}

/* polymorph ImageInfoRootFS BaseLayer false */

/* polymorph ImageInfoRootFS Layers false */

/* polymorph ImageInfoRootFS Type false */

// Validate validates this image info root f s
func (m *ImageInfoRootFS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageInfoRootFS) validateLayers(formats strfmt.Registry) error {

	if swag.IsZero(m.Layers) { // not required
		return nil
	}

	return nil
}

func (m *ImageInfoRootFS) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("RootFS"+"."+"Type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageInfoRootFS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageInfoRootFS) UnmarshalBinary(b []byte) error {
	var res ImageInfoRootFS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}