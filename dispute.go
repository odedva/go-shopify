package goshopify

import (
	"fmt"
	"time"
)

const disputeBasePath = "shopify_payments/disputes"

// const disputeResourceName = "shopyfi_payments"

// DisputeService is an interface for interfacing with the dispute endpoints
// of the Shopify API.
// See: https://shopify.dev/docs/admin-api/rest/reference/shopify_payments/dispute
type DisputeService interface {
	List(interface{}) ([]Dispute, error)
	Get(int64, interface{}) (*Dispute, error)
}

type DisputeListOptions struct {
	SinceID     int64     `url:"since_id,omitempty"`
	LastID      int64     `url:"last_id,omitempty"`
	Status      string    `url:"status,omitempty"`
	InitiatedAt time.Time `url:"initiated_at,omitempty"`
}

// DisputeServiceOp handles communication with the product related methods of
// the Shopify API.
type DisputeServiceOp struct {
	client *Client
}

// Dispute represents a Shopify dispute
type Dispute struct {
	ID                int64      `json:"id"`
	OrderId           int64      `json:"order_id"`
	Type              string     `json:"type"`
	Amount            string     `json:"amount"`
	Currency          string     `json:"currency"`
	Reason            string     `json:"reason"`
	NetworkReasonCode string     `json:"network_reason_code"`
	Status            string     `json:"status"`
	EvidenceDueBy     *time.Time `json:"evidence_due_by,omitempty"`
	EvidenceSentOn    *time.Time `json:"evidence_sent_on,omitempty"`
	FinalizedOn       string     `json:"finalized_on"`
	InitiatedAt       *time.Time `json:"initiated_at,omitempty"`
}

// Represents the result from the dispute/X.json endpoint
type DisputeResource struct {
	Dispute *Dispute `json:"dispute"`
}

// Represents the result from the Disputs.json endpoint
type DisputesResource struct {
	Disputes []Dispute `json:"disputes"`
}

// List disputes
func (s *DisputeServiceOp) List(options interface{}) ([]Dispute, error) {
	path := fmt.Sprintf("%s.json", disputeBasePath)
	resource := new(DisputesResource)
	err := s.client.Get(path, resource, options)
	return resource.Disputes, err
}

// Get dispute
func (s *DisputeServiceOp) Get(disputeID int64, options interface{}) (*Dispute, error) {
	path := fmt.Sprintf("%s/%v.json", disputeBasePath, disputeID)
	resource := new(DisputeResource)
	err := s.client.Get(path, resource, options)
	return resource.Dispute, err
}
