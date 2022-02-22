// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SecurityContext SecurityContext holds security configuration that will be applied to a container.
// Some fields are present in both SecurityContext and PodSecurityContext.  When both
// are set, the values in SecurityContext take precedence.
//
// swagger:model v1SecurityContext
type V1SecurityContext struct {

	// AllowPrivilegeEscalation controls whether a process can gain more
	// privileges than its parent process. This bool directly controls if
	// the no_new_privs flag will be set on the container process.
	// AllowPrivilegeEscalation is true always when the container is:
	// 1) run as Privileged
	// 2) has CAP_SYS_ADMIN
	// +optional
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`

	// The capabilities to add/drop when running containers.
	// Defaults to the default set of capabilities granted by the container runtime.
	// +optional
	Capabilities *V1Capabilities `json:"capabilities,omitempty"`

	// Run container in privileged mode.
	// Processes in privileged containers are essentially equivalent to root on the host.
	// Defaults to false.
	// +optional
	Privileged bool `json:"privileged,omitempty"`

	// procMount denotes the type of proc mount to use for the containers.
	// The default is DefaultProcMount which uses the container runtime defaults for
	// readonly paths and masked paths.
	// This requires the ProcMountType feature flag to be enabled.
	// +optional
	ProcMount string `json:"procMount,omitempty"`

	// Whether this container has a read-only root filesystem.
	// Default is false.
	// +optional
	ReadOnlyRootFilesystem bool `json:"readOnlyRootFilesystem,omitempty"`

	// The GID to run the entrypoint of the container process.
	// Uses runtime default if unset.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsGroup string `json:"runAsGroup,omitempty"`

	// Indicates that the container must run as a non-root user.
	// If true, the Kubelet will validate the image at runtime to ensure that it
	// does not run as UID 0 (root) and fail to start the container if it does.
	// If unset or false, no such validation will be performed.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`

	// The UID to run the entrypoint of the container process.
	// Defaults to user specified in image metadata if unspecified.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsUser string `json:"runAsUser,omitempty"`

	// The SELinux context to be applied to the container.
	// If unspecified, the container runtime will allocate a random SELinux context for each
	// container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	SeLinuxOptions *V1SELinuxOptions `json:"seLinuxOptions,omitempty"`

	// The seccomp options to use by this container. If seccomp options are
	// provided at both the pod & container level, the container options
	// override the pod options.
	// +optional
	SeccompProfile *V1SeccompProfile `json:"seccompProfile,omitempty"`

	// The Windows specific settings applied to all containers.
	// If unspecified, the options from the PodSecurityContext will be used.
	// If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	WindowsOptions *V1WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}

// Validate validates this v1 security context
func (m *V1SecurityContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeLinuxOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeccompProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindowsOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SecurityContext) validateCapabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if m.Capabilities != nil {
		if err := m.Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *V1SecurityContext) validateSeLinuxOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.SeLinuxOptions) { // not required
		return nil
	}

	if m.SeLinuxOptions != nil {
		if err := m.SeLinuxOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seLinuxOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seLinuxOptions")
			}
			return err
		}
	}

	return nil
}

func (m *V1SecurityContext) validateSeccompProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.SeccompProfile) { // not required
		return nil
	}

	if m.SeccompProfile != nil {
		if err := m.SeccompProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seccompProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seccompProfile")
			}
			return err
		}
	}

	return nil
}

func (m *V1SecurityContext) validateWindowsOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.WindowsOptions) { // not required
		return nil
	}

	if m.WindowsOptions != nil {
		if err := m.WindowsOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windowsOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("windowsOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 security context based on the context it is used
func (m *V1SecurityContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeLinuxOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeccompProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWindowsOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SecurityContext) contextValidateCapabilities(ctx context.Context, formats strfmt.Registry) error {

	if m.Capabilities != nil {
		if err := m.Capabilities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *V1SecurityContext) contextValidateSeLinuxOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.SeLinuxOptions != nil {
		if err := m.SeLinuxOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seLinuxOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seLinuxOptions")
			}
			return err
		}
	}

	return nil
}

func (m *V1SecurityContext) contextValidateSeccompProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.SeccompProfile != nil {
		if err := m.SeccompProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seccompProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("seccompProfile")
			}
			return err
		}
	}

	return nil
}

func (m *V1SecurityContext) contextValidateWindowsOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.WindowsOptions != nil {
		if err := m.WindowsOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windowsOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("windowsOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SecurityContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SecurityContext) UnmarshalBinary(b []byte) error {
	var res V1SecurityContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
