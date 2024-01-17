package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/aws/smithy-go/ptr"

	"github.com/raito-io/sdk/internal"
	"github.com/raito-io/sdk/internal/schema"
	"github.com/raito-io/sdk/types"
)

type AccessProviderClient struct {
	client graphql.Client
}

func NewAccessProviderClient(client graphql.Client) AccessProviderClient {
	return AccessProviderClient{
		client: client,
	}
}

// CreateAccessProvider creates a new AccessProvider in Raito Cloud.
// They valid AccessProvider is returned if the creation is successful.
// Otherwise, an error is returned
func (a *AccessProviderClient) CreateAccessProvider(ctx context.Context, ap types.AccessProviderInput) (*types.AccessProvider, error) {
	result, err := schema.CreateAccessProvider(ctx, a.client, ap)
	if err != nil {
		return nil, NewErrClient(err)
	}

	switch response := result.CreateAccessProvider.(type) {
	case *schema.CreateAccessProviderCreateAccessProvider:
		return &response.AccessProvider, nil
	case *schema.CreateAccessProviderCreateAccessProviderPermissionDeniedError:
		return nil, NewErrPermissionDenied("createAccessProvider", response.Message)
	default:
		return nil, fmt.Errorf("unexpected response type: %T", result.CreateAccessProvider)
	}
}

// UpdateAccessProvider updates an existing AccessProvider in Raito Cloud.
// The updated AccessProvider is returned if the update is successful.
// Otherwise, an error is returned.
func (a *AccessProviderClient) UpdateAccessProvider(ctx context.Context, id string, ap schema.AccessProviderInput) (*types.AccessProvider, error) {
	result, err := schema.UpdateAccessProvider(ctx, a.client, id, ap)
	if err != nil {
		return nil, NewErrClient(err)
	}

	switch response := result.UpdateAccessProvider.(type) {
	case *schema.UpdateAccessProviderUpdateAccessProvider:
		return &response.AccessProvider, nil
	case *schema.UpdateAccessProviderUpdateAccessProviderPermissionDeniedError:
		return nil, NewErrPermissionDenied("updateAccessProvider", response.Message)
	default:
		return nil, fmt.Errorf("unexpected response type: %T", result.UpdateAccessProvider)
	}
}

// DeleteAccessProvider deletes an existing AccessProvider in Raito Cloud.
// If the deletion is successful, nil is returned.
// Otherwise, an error is returned.
func (a *AccessProviderClient) DeleteAccessProvider(ctx context.Context, id string) error {
	result, err := schema.DeleteAccessProvider(ctx, a.client, id)
	if err != nil {
		return NewErrClient(err)
	}

	switch response := result.DeleteAccessProvider.(type) {
	case *schema.DeleteAccessProviderDeleteAccessProvider:
		return nil
	case *schema.DeleteAccessProviderDeleteAccessProviderPermissionDeniedError:
		return NewErrPermissionDenied("deleteAccessProvider", response.Message)
	case *schema.DeleteAccessProviderDeleteAccessProviderNotFoundError:
		return NewErrNotFound(id, "accessProvider", response.Message)
	default:
		return fmt.Errorf("unexpected response type: %T", result.DeleteAccessProvider)
	}
}

// GetAccessProvider returns a specific AccessProvider
func (a *AccessProviderClient) GetAccessProvider(ctx context.Context, id string) (*types.AccessProvider, error) {
	result, err := schema.GetAccessProvider(ctx, a.client, id)
	if err != nil {
		return nil, NewErrClient(err)
	}

	switch ap := result.AccessProvider.(type) {
	case *schema.GetAccessProviderAccessProvider:
		return &ap.AccessProvider, nil
	case *schema.GetAccessProviderAccessProviderNotFoundError:
		return nil, NewErrNotFound(id, "accessProvider", ap.Message)
	case *schema.GetAccessProviderAccessProviderPermissionDeniedError:
		return nil, NewErrPermissionDenied("getAccessProvider", ap.Message)
	default:
		return nil, fmt.Errorf("unexpected response type: %T", result.AccessProvider)
	}
}

type AccessProviderListOptions struct {
	order  []types.AccessProviderOrderByInput
	filter *types.AccessProviderFilterInput
}

// WithAccessProviderListOrder can be used to specify the order of the returned AccessProviders.
func WithAccessProviderListOrder(input ...types.AccessProviderOrderByInput) func(options *AccessProviderListOptions) {
	return func(options *AccessProviderListOptions) {
		options.order = append(options.order, input...)
	}
}

// WithAccessProviderListFilter can be used to filter the returned AccessProviders.
func WithAccessProviderListFilter(input *types.AccessProviderFilterInput) func(options *AccessProviderListOptions) {
	return func(options *AccessProviderListOptions) {
		options.filter = input
	}
}

// ListAccessProviders returns a list of AccessProviders in Raito Cloud.
// The order of the list can be specified with WithAccessProviderListOrder.
// A filter can be specified with WithAccessProviderListFilter.
// A channel is returned that can be used to receive the list of AccessProviders.
// To close the channel ensure to cancel the context.
func (a *AccessProviderClient) ListAccessProviders(ctx context.Context, ops ...func(*AccessProviderListOptions)) <-chan types.ListItem[types.AccessProvider] {
	options := AccessProviderListOptions{}
	for _, op := range ops {
		op(&options)
	}

	loadPageFn := func(ctx context.Context, cursor *string) (*schema.PageInfo, []schema.AccessProviderPageEdgesEdge, error) {
		output, err := schema.ListAccessProviders(ctx, a.client, cursor, ptr.Int(25), options.filter, options.order)
		if err != nil {
			return nil, nil, NewErrClient(err)
		}

		switch page := output.AccessProviders.(type) {
		case *schema.ListAccessProvidersAccessProvidersPagedResult:
			return &page.PageInfo.PageInfo, page.Edges, nil
		case *schema.ListAccessProvidersAccessProvidersPermissionDeniedError:
			return nil, nil, NewErrPermissionDenied("listAccessProviders", page.Message)
		default:
			return nil, nil, errors.New("unreachable")
		}
	}

	edgeFn := func(edge *schema.AccessProviderPageEdgesEdge) (*string, *schema.AccessProvider, error) {
		cursor := edge.Cursor

		if edge.Node == nil {
			return cursor, nil, nil
		}

		listItem := (*edge.Node).(*schema.AccessProviderPageEdgesEdgeNodeAccessProvider)

		return cursor, &listItem.AccessProvider, nil
	}

	return internal.PaginationExecutor(ctx, loadPageFn, edgeFn)
}

type AccessProviderWhoListOptions struct {
	order []types.AccessProviderWhoOrderByInput
}

// WithAccessProviderWhoListOrder can be used to specify the order of the returned AccessProviderWhoList
func WithAccessProviderWhoListOrder(input ...schema.AccessProviderWhoOrderByInput) func(options *AccessProviderWhoListOptions) {
	return func(options *AccessProviderWhoListOptions) {
		options.order = append(options.order, input...)
	}
}

// GetAccessProviderWhoList returns all who items of an AccessProvider in Raito Cloud.
// The order of the list can be specified with WithAccessProviderWhoListOrder.
// A channel is returned that can be used to receive the list of AccessProviderWhoListItem.
// To close the channel ensure to cancel the context.
func (a *AccessProviderClient) GetAccessProviderWhoList(ctx context.Context, id string, ops ...func(*AccessProviderWhoListOptions)) <-chan types.ListItem[types.AccessProviderWhoListItem] { //nolint:dupl
	options := AccessProviderWhoListOptions{}
	for _, op := range ops {
		op(&options)
	}

	loadPageFn := func(ctx context.Context, cursor *string) (*types.PageInfo, []types.AccessProviderWhoListEdgesEdge, error) {
		output, err := schema.GetAccessProviderWhoList(ctx, a.client, id, cursor, ptr.Int(25), nil, options.order)
		if err != nil {
			return nil, nil, NewErrClient(err)
		}

		switch ap := output.AccessProvider.(type) {
		case *schema.GetAccessProviderWhoListAccessProvider:
			switch whoList := ap.WhoList.(type) {
			case *schema.GetAccessProviderWhoListAccessProviderWhoListPagedResult:
				return &whoList.PageInfo.PageInfo, whoList.Edges, nil
			case *schema.GetAccessProviderWhoListAccessProviderWhoListPermissionDeniedError:
				return nil, nil, NewErrPermissionDenied("accessProviderWhoList", whoList.Message)
			}
		case *schema.GetAccessProviderWhoListAccessProviderNotFoundError:
			return nil, nil, NewErrNotFound("AccessProvider", id, ap.Message)
		case *schema.GetAccessProviderWhoListAccessProviderPermissionDeniedError:
			return nil, nil, NewErrPermissionDenied("accessProvider", ap.Message)
		default:
			return nil, nil, fmt.Errorf("unexpected type '%T': %w", ap, ErrUnknownType)
		}

		return nil, nil, errors.New("unreachable")
	}

	edgeFn := func(edge *types.AccessProviderWhoListEdgesEdge) (*string, *schema.AccessProviderWhoListItem, error) {
		cursor := edge.Cursor

		if edge.Node == nil {
			return cursor, nil, nil
		}

		listItem := (*edge.Node).(*types.AccessProviderWhoListEdgesEdgeNodeAccessWhoItem)

		return cursor, &listItem.AccessProviderWhoListItem, nil
	}

	return internal.PaginationExecutor(ctx, loadPageFn, edgeFn)
}

type AccessProviderWhatListOptions struct {
	order []schema.AccessWhatOrderByInput
}

// WithAccessProviderWhatListOrder can be used to specify the order of the returned AccessProviderWhatList
func WithAccessProviderWhatListOrder(input ...schema.AccessWhatOrderByInput) func(options *AccessProviderWhatListOptions) {
	return func(options *AccessProviderWhatListOptions) {
		options.order = append(options.order, input...)
	}
}

// GetAccessProviderWhatDataObjectList returns all what items of an AccessProvider in Raito Cloud.
// The order of the list can be specified with WithAccessProviderWhatListOrder.
// A channel is returned that can be used to receive the list of AccessProviderWhatDataObjectListItem.
// To close the channel ensure to cancel the context.
func (a *AccessProviderClient) GetAccessProviderWhatDataObjectList(ctx context.Context, id string, ops ...func(*AccessProviderWhatListOptions)) <-chan types.ListItem[types.AccessProviderWhatListItem] { //nolint:dupl
	options := AccessProviderWhatListOptions{}
	for _, op := range ops {
		op(&options)
	}

	loadPageFn := func(ctx context.Context, cursor *string) (*types.PageInfo, []types.AccessProviderWhatListEdgesEdge, error) {
		output, err := schema.GetAccessProviderWhatDataObjectList(ctx, a.client, id, cursor, ptr.Int(25), nil, options.order)
		if err != nil {
			return nil, nil, NewErrClient(err)
		}

		switch ap := output.AccessProvider.(type) {
		case *schema.GetAccessProviderWhatDataObjectListAccessProvider:
			switch whatList := ap.WhatDataObjects.(type) {
			case *schema.GetAccessProviderWhatDataObjectListAccessProviderWhatDataObjectsPagedResult:
				return &whatList.PageInfo.PageInfo, whatList.Edges, nil
			case *schema.GetAccessProviderWhatDataObjectListAccessProviderWhatDataObjectsPermissionDeniedError:
				return nil, nil, NewErrPermissionDenied("accessProviderWhatDataObjectList", whatList.Message)
			}
		case *schema.GetAccessProviderWhatDataObjectListAccessProviderNotFoundError:
			return nil, nil, NewErrNotFound("AccessProvider", id, ap.Message)
		case *schema.GetAccessProviderWhatDataObjectListAccessProviderPermissionDeniedError:
			return nil, nil, NewErrPermissionDenied("accessProvider", ap.Message)
		default:
			return nil, nil, fmt.Errorf("unexpected type '%T': %w", ap, ErrUnknownType)
		}

		return nil, nil, errors.New("unreachable")
	}

	edgeFn := func(edge *types.AccessProviderWhatListEdgesEdge) (*string, *schema.AccessProviderWhatListItem, error) {
		cursor := edge.Cursor

		if edge.Node == nil {
			return cursor, nil, nil
		}

		listItem := (*edge.Node).(*types.AccessProviderWhatListEdgesEdgeNodeAccessWhatItem)

		return cursor, &listItem.AccessProviderWhatListItem, nil
	}

	return internal.PaginationExecutor(ctx, loadPageFn, edgeFn)
}
