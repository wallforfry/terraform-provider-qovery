package port

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Ports []Port

const (
	DefaultProtocol = ProtocolHTTP
	MinPort         = 1
	MaxPort         = 65535
)

var (
	// ErrInvalidPort is the error return if a Port is invalid.
	ErrInvalidPort = errors.New("invalid port")
	// ErrInvalidPorts is the error return if a Ports is invalid.
	ErrInvalidPorts = errors.New("invalid ports")
	// ErrInvalidPortIDParam is the error return if the port id param is invalid.
	ErrInvalidPortIDParam = errors.New("invalid port id params")
	// ErrInvalidInternalPortParam is returned if the internal port param is invalid.
	ErrInvalidInternalPortParam = errors.New("invalid internal port param")
	// ErrInvalidExternalPortParam is returned if the external port param is invalid.
	ErrInvalidExternalPortParam = errors.New("invalid external port param")
	// ErrInvalidNameParam is returned if the name param is invalid.
	ErrInvalidNameParam = errors.New("invalid name param")
	// ErrInvalidProtocolParam is returned if the protocol param is invalid.
	ErrInvalidProtocolParam = errors.New("invalid protocol param")
	// ErrInvalidUpsertRequest is returned if the create request is invalid.
	ErrInvalidUpsertRequest = errors.New("invalid port upsert request")
)

// Validate returns an error to tell whether the Ports' domain model is valid or not.
func (ss Ports) Validate() error {
	for _, it := range ss {
		if err := it.Validate(); err != nil {
			return errors.Wrap(err, ErrInvalidPorts.Error())
		}
	}

	return nil
}

// IsValid returns a bool to tell whether the Ports' domain model is valid or not.
func (ss Ports) IsValid() bool {
	return ss.Validate() == nil
}

type Port struct {
	ID                 uuid.UUID
	InternalPort       int32 `validate:"required"`
	PubliclyAccessible bool

	Protocol     *Protocol
	Name         *string
	ExternalPort *int32
}

// Validate returns an error to tell whether the Port domain model is valid or not.
func (s Port) Validate() error {
	if err := validator.New().Struct(s); err != nil {
		return errors.Wrap(err, ErrInvalidPort.Error())
	}

	return nil
}

// IsValid returns a bool to tell whether the Port domain model is valid or not.
func (s Port) IsValid() bool {
	return s.Validate() == nil
}

// NewPortParams represents the arguments needed to create a Port.
type NewPortParams struct {
	PortID             string
	InternalPort       int32
	PubliclyAccessible bool
	Protocol           string

	Name         *string
	ExternalPort *int32
}

// NewPort returns a new instance of a Port domain model.
func NewPort(params NewPortParams) (*Port, error) {
	portUUID, err := uuid.Parse(params.PortID)
	if err != nil {
		return nil, errors.Wrap(err, ErrInvalidPortIDParam.Error())
	}

	portProtocol, err := NewProtocolFromString(params.Protocol)
	if err != nil {
		return nil, errors.Wrap(err, ErrInvalidProtocolParam.Error())
	}

	if params.InternalPort < 0 {
		return nil, ErrInvalidInternalPortParam
	}

	if params.ExternalPort != nil && *params.ExternalPort < 0 {
		return nil, ErrInvalidExternalPortParam
	}

	if params.Name != nil && *params.Name == "" {
		return nil, ErrInvalidNameParam
	}

	v := &Port{
		ID:                 portUUID,
		InternalPort:       params.InternalPort,
		PubliclyAccessible: params.PubliclyAccessible,
		Protocol:           portProtocol,
		Name:               params.Name,
		ExternalPort:       params.ExternalPort,
	}

	if err := v.Validate(); err != nil {
		return nil, err
	}

	return v, nil
}

// UpsertRequest represents the parameters needed to create & update a Variable.
type UpsertRequest struct {
	InternalPort       int32 `validate:"required"`
	PubliclyAccessible bool  `validate:"required"`

	Protocol     *string
	Name         *string
	ExternalPort *int32
}

// Validate returns an error to tell whether the UpsertRequest is valid or not.
func (r UpsertRequest) Validate() error {
	if err := validator.New().Struct(r); err != nil {
		return errors.Wrap(err, ErrInvalidUpsertRequest.Error())
	}

	return nil
}

// IsValid returns a bool to tell whether the UpsertRequest is valid or not.
func (r UpsertRequest) IsValid() bool {
	return r.Validate() == nil
}
