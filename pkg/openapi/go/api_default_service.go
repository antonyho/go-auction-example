/*
 * Auction Bid Tracker
 *
 * This is an example server for auction bid tracker.
 *
 * API version: 1.0.0
 * Contact: antony.h@riseup.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API. 
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// BidItemById - 
func (s *DefaultApiService) BidItemById(id string, bidding Bidding) (interface{}, error) {
	// TODO - update BidItemById with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'BidItemById' not implemented")
}

// GetWinningBidByItemId - 
func (s *DefaultApiService) GetWinningBidByItemId(id string) (interface{}, error) {
	// TODO - update GetWinningBidByItemId with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetWinningBidByItemId' not implemented")
}

// ListAllBidsByItemId - 
func (s *DefaultApiService) ListAllBidsByItemId(id string) (interface{}, error) {
	// TODO - update ListAllBidsByItemId with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ListAllBidsByItemId' not implemented")
}

// ListAllBidsByUserId - 
func (s *DefaultApiService) ListAllBidsByUserId(id string) (interface{}, error) {
	// TODO - update ListAllBidsByUserId with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ListAllBidsByUserId' not implemented")
}
