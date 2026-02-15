package deliverd

// courier type from 'couriers' table
type Courier struct {
	CourierID       int
	NameFirst       string
	NameMiddle      string
	NameLast        string
	AddressFirst    string
	AddressSecond   string
	AddressPostcode string
	AddressCounty   string
	AddressCountry  string
	Email           string
	PhoneNumber     string
	Status          CourierStatus
	DateBirth       int
	DateJoined      int
	DateLeft        int
}

// drops from 'drops', these are parcels on a route.
type Drops struct {
	DropID    int
	PackageID int
	RouteID   int
	Order     int
}

// keys from 'keys' table, used for authentication and authorization
type Keys struct {
	KeyID       string // 20 max
	KeyType     APIKeyType
	KeyOwner    int
	DateCreated int
	Status      APIKeyStatus
}

// packages
type Packages struct {
	PackageID           int
	PartnerID           int
	Description         string
	DateCreated         int
	DestinationFirst    string
	DestinationSecond   string
	DestinationPostcode string
	DestinationCounty   string
	DestinationCountry  string
	RecipientNameFirst  string
	RecipientNameLast   string
	RecipientEmail      string
	RecipientPhone      string
	Classification      int
	Status              UpdateID
}

// partners
type Partner struct {
	PartnersID      int
	CompanyName     string
	CompanyID       string
	AddressFirst    string
	AddressSecond   string
	AddressPostcode string
	AddressCounty   string
	AddressCountry  string
	Status          PartnerStatus
	DateCreation    int
	RepNameFirst    string
	RepNameLast     string
	RepEmail        string
	RepPhone        string
}

type PartnerStatus int

const (
	// partner is active and can create deliveries
	PARTNER_ACTIVE PartnerStatus = iota
	// partner is inactive and cannot create deliveries
	PARTNER_INACTIVE
)

// route (a courier's collection of drops)
type Route struct {
	RouteID       int
	CourierID     int
	Status        RouteStatus
	TimeCreated   int
	TimeCompleted int
}

// updates for a drops status
type Update struct {
	UpdateID        int
	PackageID       int
	UpdateType      UpdateID
	AddressFirst    string
	AddressSecond   string
	AddressPostcode string
	AddressCounty   string
	AddressCountry  string
	Freeform        string
	UpdateDate      int
}

// status of courier
type CourierStatus int

const (
	// courier is active and can be assigned deliveries
	COURIER_ACTIVE CourierStatus = iota
	// courier is inactive and cannot be assigned deliveries
	COURIER_INACTIVE
	// courier is on leave and cannot be assigned deliveries
	COURIER_ON_LEAVE
)

// for fields in 'updates' and 'packages'
type UpdateID int

const (
	// the sender has informed network of new delivery
	UPDATE_INFORMED_SENDER UpdateID = iota
	// the delivery has arrived at distribution
	UPDATE_RECEIVED_DISTRIBUTION
	// the delivery is ready to be picked up by delivery
	UPDATE_DELIVERY_READ
	// the delivery is out for delivery
	UPDATE_DELIVERY_OUT
	// the delivery is done
	UPDATE_DELIVERY_DONE
	// couldn't deliver
	UPDATE_DELIVERY_FAIL
	// package being returned to sender
	UPDATE_RETURN_SENDER
	// delivery is being delayed
	UPDATE_DELIVERY_DELAY
)

// type of key
type APIKeyType int

const (
	// key used by partner
	KEY_PARTNER APIKeyType = iota
	// key used by courier
	KEY_COURIER
	// key used by admin
	KEY_ADMIN
)

// status of api key
type APIKeyStatus int

const (
	// key is active and can be used
	STATUS_ACTIVE APIKeyStatus = iota
	// key is inactive and cannot be used
	STATUS_INACTIVE
)

// status of a route
type RouteStatus int

const (
	// route is active and can be used
	ROUTE_ACTIVE RouteStatus = iota
	// route is in progress
	ROUTE_IN_PROGRESS
	// route is completed
	ROUTE_COMPLETED
	// route is cancelled
	ROUTE_CANCELLED
)

type Classification struct {
	ClassificationID int
	Name             string
	Description      string
	ExpectedTime     int
	RequiredTime     int
	PriceGBP         float64
}
