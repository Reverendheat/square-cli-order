package square

import "time"

type CatalogV1Id struct {
	CatalogV1ID string `json:"catalog_v1_id"`
	LocationID  string `json:"location_id"`
}

type LocationOverrides struct {
	LocationID     string `json:"location_id"`
	TrackInventory bool   `json:"track_inventory"`
	SoldOut        bool   `json:"sold_out"`
}

type PriceMoney struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type ItemVariationData struct {
	ItemID            string              `json:"item_id"`
	Name              string              `json:"name"`
	Ordinal           int                 `json:"ordinal"`
	PricingType       string              `json:"pricing_type"`
	PriceMoney        PriceMoney          `json:"price_money"`
	LocationOverrides []LocationOverrides `json:"location_overrides"`
	Sellable          bool                `json:"sellable"`
	Stockable         bool                `json:"stockable"`
}

type Variation struct {
	Type                  string            `json:"type"`
	ID                    string            `json:"id"`
	UpdatedAt             time.Time         `json:"updated_at"`
	CreatedAt             time.Time         `json:"created_at"`
	Version               int64             `json:"version"`
	IsDeleted             bool              `json:"is_deleted"`
	CatalogV1Ids          []CatalogV1Id     `json:"catalog_v1_ids"`
	PresentAtAllLocations bool              `json:"present_at_all_locations"`
	PresentAtLocationIds  []string          `json:"present_at_location_ids"`
	ItemVariationData     ItemVariationData `json:"item_variation_data"`
}

type ItemData struct {
	Name                             string             `json:"name"`
	Abbreviation                     string             `json:"abbreviation"`
	LabelColor                       string             `json:"label_color"`
	Visibility                       string             `json:"visibility"`
	CategoryID                       string             `json:"category_id"`
	Variations                       []Variation        `json:"variations"`
	ProductType                      string             `json:"product_type"`
	ModifierListInfo                 []ModifierListInfo `json:"modifier_list_info"`
	SkipModifierScreen               bool               `json:"skip_modifier_screen"`
	EcomAvailable                    bool               `json:"ecom_available"`
	EcomVisibility                   string             `json:"ecom_visibility"`
	PickupFulfillmentPreferencesID   string             `json:"pickup_fulfillment_preferences_id"`
	DeliveryFulfillmentPreferencesID string             `json:"delivery_fulfillment_preferences_id"`
}

type ImageData struct {
	URL string `json:"url"`
}

type ModifierData struct {
	Name           string     `json:"name"`
	PriceMoney     PriceMoney `json:"price_money"`
	OnByDefault    bool       `json:"on_by_default"`
	Ordinal        int        `json:"ordinal"`
	ModifierListID string     `json:"modifier_list_id"`
}

type Modifier struct {
	Type                  string       `json:"type"`
	ID                    string       `json:"id"`
	UpdatedAt             time.Time    `json:"updated_at"`
	CreatedAt             time.Time    `json:"created_at"`
	Version               int64        `json:"version"`
	IsDeleted             bool         `json:"is_deleted"`
	PresentAtAllLocations bool         `json:"present_at_all_locations"`
	ModifierData          ModifierData `json:"modifier_data"`
}

type ModifierListData struct {
	Name          string     `json:"name"`
	SelectionType string     `json:"selection_type"`
	Modifiers     []Modifier `json:"modifiers"`
}

type CategoryData struct {
	Name string `json:"name"`
}

type TaxData struct {
	Name                   string `json:"name"`
	CalculationPhase       string `json:"calculation_phase"`
	InclusionType          string `json:"inclusion_type"`
	Percentage             string `json:"percentage"`
	AppliesToCustomAmounts bool   `json:"applies_to_custom_amounts"`
	Enabled                bool   `json:"enabled"`
	TaxTypeID              string `json:"tax_type_id"`
	TaxTypeName            string `json:"tax_type_name"`
}

type RelatedObject struct {
	Type                  string           `json:"type"`
	ID                    string           `json:"id"`
	UpdatedAt             time.Time        `json:"updated_at"`
	CreatedAt             time.Time        `json:"created_at"`
	Version               int64            `json:"version"`
	IsDeleted             bool             `json:"is_deleted"`
	CatalogV1Ids          []CatalogV1Id    `json:"catalog_v1_ids,omitempty"`
	PresentAtAllLocations bool             `json:"present_at_all_locations"`
	CategoryData          CategoryData     `json:"category_data,omitempty"`
	TaxData               TaxData          `json:"tax_data,omitempty"`
	ModifierListData      ModifierListData `json:"modifier_list_data,omitempty"`
	PresentAtLocationIds  []string         `json:"present_at_location_ids,omitempty"`
	ImageData             ImageData        `json:"image_data,omitempty"`
}

type CatalogObject struct {
	Type                  string        `json:"type"`
	ID                    string        `json:"id"`
	UpdatedAt             time.Time     `json:"updated_at"`
	CreatedAt             time.Time     `json:"created_at"`
	Version               int64         `json:"version"`
	IsDeleted             bool          `json:"is_deleted"`
	CatalogV1Ids          []CatalogV1Id `json:"catalog_v1_ids,omitempty"`
	PresentAtAllLocations bool          `json:"present_at_all_locations"`
	PresentAtLocationIds  []string      `json:"present_at_location_ids,omitempty"`
	ItemData              ItemData      `json:"item_data"`
}

type ModifierListInfo struct {
	ModifierListID       string `json:"modifier_list_id"`
	MinSelectedModifiers int    `json:"min_selected_modifiers"`
	MaxSelectedModifiers int    `json:"max_selected_modifiers"`
	Enabled              bool   `json:"enabled"`
}
type ItemOption struct {
	ItemOptionID string `json:"item_option_id"`
}

type CatalogItem struct {
	Type                  string    `json:"type"`
	ID                    string    `json:"id"`
	UpdatedAt             time.Time `json:"updated_at"`
	CreatedAt             time.Time `json:"created_at"`
	Version               int64     `json:"version"`
	IsDeleted             bool      `json:"is_deleted"`
	PresentAtAllLocations bool      `json:"present_at_all_locations"`
	ItemData              ItemData  `json:"item_data"`
}

type SearchCatalogItemsResponse struct {
	Items []CatalogItem `json:"items"`
}

type CategoryObjectResponse struct {
	Objects []CategoryObject `json:"objects"`
}

type CategoryObject struct {
	Type         string    `json:"type"`
	ID           string    `json:"id"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	Version      int64     `json:"version"`
	IsDeleted    bool      `json:"is_deleted"`
	CatalogV1Ids []struct {
		CatalogV1ID string `json:"catalog_v1_id"`
		LocationID  string `json:"location_id"`
	} `json:"catalog_v1_ids"`
	PresentAtAllLocations bool `json:"present_at_all_locations"`
	CategoryData          struct {
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		IsTopLevel   bool   `json:"is_top_level"`
	} `json:"category_data"`
}

type SearchCatalogObjectsResponse struct {
	Objects        []CatalogObject `json:"objects"`
	RelatedObjects []RelatedObject `json:"related_objects"`
	LatestTime     time.Time       `json:"latest_time"`
}

type BatchRetrieveCatalogObjectsRequest struct {
	ObjectIds []string `json:"object_ids"`
}

type BatchRetrieveCatalogObjectsResponse struct {
	Objects []CatalogObject `json:"objects"`
}

type ModifierListResponse struct {
	Objects []ModifierList `json:"objects"`
}

type VariationResponse struct {
	Objects []Variation `json:"objects"`
}

type ModifierList struct {
	Type                  string           `json:"type"`
	ID                    string           `json:"id"`
	UpdatedAt             time.Time        `json:"updated_at"`
	CreatedAt             time.Time        `json:"created_at"`
	Version               int64            `json:"version"`
	IsDeleted             bool             `json:"is_deleted"`
	PresentAtAllLocations bool             `json:"present_at_all_locations"`
	ModifierListData      ModifierListData `json:"modifier_list_data"`
}

type CatalogItemOptions struct {
	Modifiers []ModifierList
	Options   []Variation
}
