package qoveryapi

import (
	"context"

	"github.com/qovery/qovery-client-go"

	"github.com/qovery/terraform-provider-qovery/internal/domain/apierrors"
	"github.com/qovery/terraform-provider-qovery/internal/domain/variable"
)

// Ensure environmentEnvironmentVariablesQoveryAPI defined types fully satisfy the variable.Repository interface.
var _ variable.Repository = environmentEnvironmentVariablesQoveryAPI{}

// environmentEnvironmentVariablesQoveryAPI implements the interface variable.Repository.
type environmentEnvironmentVariablesQoveryAPI struct {
	client *qovery.APIClient
}

// newEnvironmentEnvironmentVariablesQoveryAPI return a new instance of a variable.Repository that uses Qovery's API.
func newEnvironmentEnvironmentVariablesQoveryAPI(client *qovery.APIClient) (variable.Repository, error) {
	if client == nil {
		return nil, ErrInvalidQoveryAPIClient
	}

	return &environmentEnvironmentVariablesQoveryAPI{
		client: client,
	}, nil
}

// Create calls Qovery's API to create an environment variable for an environment using the given environmentID and request.
func (p environmentEnvironmentVariablesQoveryAPI) Create(ctx context.Context, environmentID string, request variable.UpsertRequest) (*variable.Variable, error) {
	v, resp, err := p.client.EnvironmentVariableApi.
		CreateEnvironmentEnvironmentVariable(ctx, environmentID).
		EnvironmentVariableRequest(newQoveryEnvironmentVariableRequestFromDomain(request)).
		Execute()
	if err != nil || resp.StatusCode >= 400 {
		return nil, apierrors.NewCreateApiError(apierrors.ApiResourceEnvironmentEnvironmentVariable, request.Key, resp, err)
	}

	return newDomainVariableFromQovery(v)
}

// List calls Qovery's API to retrieve an environment variables from an environment using the given environmentID and variableID.
func (p environmentEnvironmentVariablesQoveryAPI) List(ctx context.Context, environmentID string) (variable.Variables, error) {
	vars, resp, err := p.client.EnvironmentVariableApi.
		ListEnvironmentEnvironmentVariable(ctx, environmentID).
		Execute()
	if err != nil || resp.StatusCode >= 400 {
		return nil, apierrors.NewReadApiError(apierrors.ApiResourceEnvironmentEnvironmentVariable, environmentID, resp, err)
	}

	return newDomainVariablesFromQovery(vars)
}

// Update calls Qovery's API to update an environment variable from an environment using the given environmentID, credentialsID and request.
func (p environmentEnvironmentVariablesQoveryAPI) Update(ctx context.Context, environmentID string, credentialsID string, request variable.UpsertRequest) (*variable.Variable, error) {
	v, resp, err := p.client.EnvironmentVariableApi.
		EditEnvironmentEnvironmentVariable(ctx, environmentID, credentialsID).
		EnvironmentVariableEditRequest(newQoveryEnvironmentVariableEditRequestFromDomain(request)).
		Execute()
	if err != nil || resp.StatusCode >= 400 {
		return nil, apierrors.NewUpdateApiError(apierrors.ApiResourceEnvironmentEnvironmentVariable, credentialsID, resp, err)
	}

	return newDomainVariableFromQovery(v)
}

// Delete calls Qovery's API to delete an environment variable from an environment using the given environmentID and credentialsID.
func (p environmentEnvironmentVariablesQoveryAPI) Delete(ctx context.Context, environmentID string, credentialsID string) error {
	resp, err := p.client.EnvironmentVariableApi.
		DeleteEnvironmentEnvironmentVariable(ctx, environmentID, credentialsID).
		Execute()
	if err != nil || resp.StatusCode >= 300 {
		return apierrors.NewDeleteApiError(apierrors.ApiResourceEnvironmentEnvironmentVariable, credentialsID, resp, err)
	}

	return nil
}
