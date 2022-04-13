package goshopify

import (
	"fmt"
	"time"
)

const InventoryLevelsBasePath = "inventory_levels"

// InventoryLevelService is an interface for interacting with the
// inventory levels endpoints of the Shopify API
// See https://shopify.dev/api/admin-rest/2022-04/resources/inventorylevel
type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	Adjust(InventoryLevel) (*InventoryLevel, error)
	Connect(InventoryLevel) (*InventoryLevel, error)
	Set(InventoryLevel) (*InventoryLevel, error)
}

// InventoryLevelServiceOp is the default implementation of the InventoryLevelService interface
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory level
type InventoryLevel struct {
	InventoryItemId     int64      `json:"inventory_item_id"`
	LocationId          int64      `json:"location_id"`
	Available           int        `json:"available,omitempty"`
	AvailableAdjustment int        `json:"available_adjustment"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	AdminGraphqlAPIID   string     `json:"admin_graphql_api_id,omitempty"`
}

// InventoryLevelResource is used for handling single level requests and responses
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_level"`
}

// InventoryLevelsResource is used for handling multiple level responsees
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

// List inventory levels
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", InventoryLevelsBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryLevels, err
}

func (s *InventoryLevelServiceOp) Adjust(item InventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/adjust.json", InventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, item, resource)
	return resource.InventoryLevel, err
}

func (s *InventoryLevelServiceOp) Connect(item InventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/connect.json", InventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, item, resource)
	return resource.InventoryLevel, err
}

func (s *InventoryLevelServiceOp) Set(item InventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/set.json", InventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, item, resource)
	return resource.InventoryLevel, err
}
