package pkge

type Package struct {
	TrackNumber             string        `json:"track_number"`
	CreatedAt               string        `json:"created_at"`
	LastTrackingDate        string        `json:"last_tracking_date,omitempty"`
	Origin                  string        `json:"origin,omitempty"`
	Destination             string        `json:"destination,omitempty"`
	LastStatus              string        `json:"last_status"`
	Status                  int           `json:"status"`
	Checkpoints             []Checkpoint  `json:"checkpoints"`
	LastStatusDate          string        `json:"last_status_date"`
	EstDeliveryDateFrom     string        `json:"est_delivery_date_from,omitempty"`
	EstDeliveryDateTo       string        `json:"est_delivery_date_to,omitempty"`
	ExtraTrackNumbers       []string      `json:"extra_track_numbers"`
	Hash                    string        `json:"hash"`
	ConsolidatedTrackNumber string        `json:"consolidated_track_number,omitempty"`
	ConsolidationDate       string        `json:"consolidation_date,omitempty"`
	DestinationCountryCode  string        `json:"destination_country_code"`
	Updating                bool          `json:"updating"`
	DaysOnWay               int           `json:"days_on_way"`
	Weight                  string        `json:"weight,omitempty"`
	ExtraInfo               []interface{} `json:"extra_info"`
	Info                    []interface{} `json:"info"`
	CouriersIDS             []int         `json:"couriers_ids"`
	CourierID               int           `json:"courier_id"`
	Couriers                []Courier     `json:"couriers"`
	Courier                 Courier       `json:"courier"`
}

type PackageResponse struct {
	StatusCode int     `json:"code"`
	Payload    Package `json:"payload"`
}

type PackageList struct {
	StatusCode int       `json:"code"`
	Payload    []Package `json:"payload"`
}

type Checkpoint struct {
	// Unique checkpoint identifier. May be null in cases
	// where it is a standard placeholder checkpoint, when
	// no checkpoint has been received from any delivery service.
	ID string `json:"id"`

	// Date and time of parcel addition. ISO 8601
	Date string `json:"date"`

	// Checkpoint description
	Title string `json:"title"`

	// Parcel location at the time of receiving the checkpoint from the delivery service
	Location string `json:"location,omitempty"`

	// Geographical latitude of the parcel location at the time of receiving
	// the checkpoint from the delivery service
	Latitude float32 `json:"latitude,omitempty"`

	// Geographical longitude of the parcel location at the time of receiving
	// the checkpoint from the delivery service
	Longitude float32 `json:"longitude,omitempty"`

	// Delivery service identifier delivering the parcel at the moment
	CourierID int `json:"courierID"`

	// Delivery service delivering the parcel at the moment
	Courier Courier `json:"courier"`
}

type Courier struct {
	// Digital identifier of the delivery service
	ID int `json:"id"`

	// Alphabetic identifier of the delivery service
	Slug string `json:"slug"`

	// Name of the delivery service
	Name string `json:"name"`

	// URL of the delivery service logo
	Logo string `json:"logo"`

	// URL of the delivery service website
	WebsiteLink string `json:"website_link"`

	// List of additional fields whose values need to be
	// transmitted when adding a parcel for tracking through
	// this delivery service
	ExtraFields []ExtraFields `json:"extra_fields"`
}

type CouriersResponse struct {
	StatusCode int       `json:"code"`
	Couriers   []Courier `json:"payload"`
}

type ExtraFields struct {
	// The name of the POST parameter in which the value needs
	// to be transmitted when adding a parcel. Changes depending
	// on the delivery service and the purpose of the field.
	// The actual list of additional fields for each delivery service
	// can be obtained from the API list of delivery services.
	Name string `json:"name"`

	// Type of additional field
	Type string `json:"type"`

	// Description of the required additional information that must
	// be transmitted in this field.
	Placeholder string `json:"placeholder"`

	// Regular expression used to check the mandatory completion of
	// this field to obtain information from the delivery service.
	// If the tracking number matches this regular expression, the field is mandatory.
	FieldRegexp string `json:"field_regexp"`

	// List of available values with their descriptions.
	Values interface{} `json:"values"`
}

type Response struct {
	StatusCode int         `json:"code"`
	Payload    interface{} `json:"payload"`
}
