package saleservice

import (
	"encoding/xml"
	"reflect"

	"github.com/vacoj/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://clients.mindbodyonline.com/api/0_5"

// NewSale_x0020_ServiceSoap creates an initializes a Sale_x0020_ServiceSoap.
func NewSale_x0020_ServiceSoap(cli *soap.Client) Sale_x0020_ServiceSoap {
	return &sale_x0020_ServiceSoap{cli}
}

// Sale_x0020_ServiceSoap was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Sale_x0020_ServiceSoap interface {
	// Validates and completes a sale by processing all items added
	// to a shopping cart.
	CheckoutShoppingCart(α *CheckoutShoppingCart) (β *CheckoutShoppingCartResponse, err error)

	// Gets a list of card types that the site accepts.
	GetAcceptedCardType(α *GetAcceptedCardType) (β *GetAcceptedCardTypeResponse, err error)

	// Gets a list of Custom Payment Methods.
	GetCustomPaymentMethods(α *GetCustomPaymentMethods) (β *GetCustomPaymentMethodsResponse, err error)

	// Gets a list of packages available for sale.
	GetPackages(α *GetPackages) (β *GetPackagesResponse, err error)

	// Get a list of products available for sale.
	GetProducts(α *GetProducts) (β *GetProductsResponse, err error)

	// Gets a list of sales.
	GetSales(α *GetSales) (β *GetSalesResponse, err error)

	// Gets a list of services available for sale.
	GetServices(α *GetServices) (β *GetServicesResponse, err error)

	// Redeem a Spa Finder Gift Card.
	RedeemSpaFinderWellnessCard(α *RedeemSpaFinderWellnessCard) (β *RedeemSpaFinderWellnessCardResponse, err error)

	// Return a sale used in business mode. This only supports comp
	// payment method.
	ReturnSale(α *ReturnSale) (β *ReturnSaleResponse, err error)

	// Update select products information.
	UpdateProducts(α *UpdateProducts) (β *UpdateProductsResponse, err error)

	// Modify sale date in business mode
	UpdateSaleDate(α *UpdateSaleDate) (β *UpdateSaleDateResponse, err error)

	// Update select services information.
	UpdateServices(α *UpdateServices) (β *UpdateServicesResponse, err error)
}

// DateTime in WSDL format.
type DateTime string

// ActionCode was auto-generated from WSDL.
type ActionCode string

// Validate validates ActionCode.
func (v ActionCode) Validate() bool {
	for _, vv := range []string{
		"None",
		"Added",
		"Updated",
		"Failed",
		"Removed",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AppointmentStatus was auto-generated from WSDL.
type AppointmentStatus string

// Validate validates AppointmentStatus.
func (v AppointmentStatus) Validate() bool {
	for _, vv := range []string{
		"Booked",
		"Completed",
		"Confirmed",
		"Arrived",
		"NoShow",
		"Cancelled",
		"LateCancelled",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ScheduleType was auto-generated from WSDL.
type ScheduleType string

// Validate validates ScheduleType.
func (v ScheduleType) Validate() bool {
	for _, vv := range []string{
		"All",
		"DropIn",
		"Enrollment",
		"Appointment",
		"Resource",
		"Media",
		"Arrival",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusCode was auto-generated from WSDL.
type StatusCode string

// Validate validates StatusCode.
func (v StatusCode) Validate() bool {
	for _, vv := range []string{
		"Success",
		"InvalidCredentials",
		"InvalidParameters",
		"InternalException",
		"Unknown",
		"FailedAction",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// XMLDetailLevel was auto-generated from WSDL.
type XMLDetailLevel string

// Validate validates XMLDetailLevel.
func (v XMLDetailLevel) Validate() bool {
	for _, vv := range []string{
		"Bare",
		"Basic",
		"Full",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Appointment was auto-generated from WSDL.
type Appointment struct {
	Site             *Site             `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages         *ArrayOfString    `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute          string            `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode        string            `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	GenderPreference string            `xml:"GenderPreference,omitempty" json:"GenderPreference,omitempty" yaml:"GenderPreference,omitempty"`
	Duration         int               `xml:"Duration,omitempty" json:"Duration,omitempty" yaml:"Duration,omitempty"`
	ProviderID       string            `xml:"ProviderID,omitempty" json:"ProviderID,omitempty" yaml:"ProviderID,omitempty"`
	Action           ActionCode        `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID               int64             `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Status           AppointmentStatus `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	StartDateTime    DateTime          `xml:"StartDateTime,omitempty" json:"StartDateTime,omitempty" yaml:"StartDateTime,omitempty"`
	EndDateTime      DateTime          `xml:"EndDateTime,omitempty" json:"EndDateTime,omitempty" yaml:"EndDateTime,omitempty"`
	Notes            string            `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	StaffRequested   bool              `xml:"StaffRequested,omitempty" json:"StaffRequested,omitempty" yaml:"StaffRequested,omitempty"`
	Program          *Program          `xml:"Program,omitempty" json:"Program,omitempty" yaml:"Program,omitempty"`
	SessionType      *SessionType      `xml:"SessionType,omitempty" json:"SessionType,omitempty" yaml:"SessionType,omitempty"`
	Location         *Location         `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Staff            *Staff            `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
	Client           *Client           `xml:"Client,omitempty" json:"Client,omitempty" yaml:"Client,omitempty"`
	FirstAppointment bool              `xml:"FirstAppointment,omitempty" json:"FirstAppointment,omitempty" yaml:"FirstAppointment,omitempty"`
	ClientService    *ClientService    `xml:"ClientService,omitempty" json:"ClientService,omitempty" yaml:"ClientService,omitempty"`
	Resources        *ArrayOfResource  `xml:"Resources,omitempty" json:"Resources,omitempty" yaml:"Resources,omitempty"`
}

// ArrayOfAppointment was auto-generated from WSDL.
type ArrayOfAppointment struct {
	Appointment []*Appointment `xml:"Appointment,omitempty" json:"Appointment,omitempty" yaml:"Appointment,omitempty"`
}

// ArrayOfAvailability was auto-generated from WSDL.
type ArrayOfAvailability struct {
	Availability []*Availability `xml:"Availability,omitempty" json:"Availability,omitempty" yaml:"Availability,omitempty"`
}

// ArrayOfCartItem was auto-generated from WSDL.
type ArrayOfCartItem struct {
	CartItem []*CartItem `xml:"CartItem,omitempty" json:"CartItem,omitempty" yaml:"CartItem,omitempty"`
}

// ArrayOfClass was auto-generated from WSDL.
type ArrayOfClass struct {
	Class []*Class `xml:"Class,omitempty" json:"Class,omitempty" yaml:"Class,omitempty"`
}

// ArrayOfClassSchedule was auto-generated from WSDL.
type ArrayOfClassSchedule struct {
	ClassSchedule []*ClassSchedule `xml:"ClassSchedule,omitempty" json:"ClassSchedule,omitempty" yaml:"ClassSchedule,omitempty"`
}

// ArrayOfClient was auto-generated from WSDL.
type ArrayOfClient struct {
	Client []*Client `xml:"Client,omitempty" json:"Client,omitempty" yaml:"Client,omitempty"`
}

// ArrayOfClientIndex was auto-generated from WSDL.
type ArrayOfClientIndex struct {
	ClientIndex []*ClientIndex `xml:"ClientIndex,omitempty" json:"ClientIndex,omitempty" yaml:"ClientIndex,omitempty"`
}

// ArrayOfClientIndexValue was auto-generated from WSDL.
type ArrayOfClientIndexValue struct {
	ClientIndexValue []*ClientIndexValue `xml:"ClientIndexValue,omitempty" json:"ClientIndexValue,omitempty" yaml:"ClientIndexValue,omitempty"`
}

// ArrayOfClientRelationship was auto-generated from WSDL.
type ArrayOfClientRelationship struct {
	ClientRelationship []*ClientRelationship `xml:"ClientRelationship,omitempty" json:"ClientRelationship,omitempty" yaml:"ClientRelationship,omitempty"`
}

// ArrayOfCustomClientField was auto-generated from WSDL.
type ArrayOfCustomClientField struct {
	CustomClientField []*CustomClientField `xml:"CustomClientField,omitempty" json:"CustomClientField,omitempty" yaml:"CustomClientField,omitempty"`
}

// ArrayOfCustomPaymentInfo was auto-generated from WSDL.
type ArrayOfCustomPaymentInfo struct {
	CustomPaymentInfo []*CustomPaymentInfo `xml:"CustomPaymentInfo,omitempty" json:"CustomPaymentInfo,omitempty" yaml:"CustomPaymentInfo,omitempty"`
}

// ArrayOfInt was auto-generated from WSDL.
type ArrayOfInt struct {
	Int []int `xml:"int,omitempty" json:"int,omitempty" yaml:"int,omitempty"`
}

// ArrayOfLocation was auto-generated from WSDL.
type ArrayOfLocation struct {
	Location []*Location `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

// ArrayOfLong was auto-generated from WSDL.
type ArrayOfLong struct {
	Long []int64 `xml:"long,omitempty" json:"long,omitempty" yaml:"long,omitempty"`
}

// ArrayOfPackage was auto-generated from WSDL.
type ArrayOfPackage struct {
	Package []*Package `xml:"Package,omitempty" json:"Package,omitempty" yaml:"Package,omitempty"`
}

// ArrayOfPayment was auto-generated from WSDL.
type ArrayOfPayment struct {
	Payment []*Payment `xml:"Payment,omitempty" json:"Payment,omitempty" yaml:"Payment,omitempty"`
}

// ArrayOfPaymentInfo was auto-generated from WSDL.
type ArrayOfPaymentInfo struct {
	PaymentInfo []*PaymentInfo `xml:"PaymentInfo,omitempty" json:"PaymentInfo,omitempty" yaml:"PaymentInfo,omitempty"`
}

// ArrayOfProduct was auto-generated from WSDL.
type ArrayOfProduct struct {
	Product []*Product `xml:"Product,omitempty" json:"Product,omitempty" yaml:"Product,omitempty"`
}

// ArrayOfProgram was auto-generated from WSDL.
type ArrayOfProgram struct {
	Program []*Program `xml:"Program,omitempty" json:"Program,omitempty" yaml:"Program,omitempty"`
}

// ArrayOfProviderIDUpdate was auto-generated from WSDL.
type ArrayOfProviderIDUpdate struct {
	ProviderIDUpdate []*ProviderIDUpdate `xml:"ProviderIDUpdate,omitempty" json:"ProviderIDUpdate,omitempty" yaml:"ProviderIDUpdate,omitempty"`
}

// ArrayOfRep was auto-generated from WSDL.
type ArrayOfRep struct {
	Rep []*Rep `xml:"Rep,omitempty" json:"Rep,omitempty" yaml:"Rep,omitempty"`
}

// ArrayOfResource was auto-generated from WSDL.
type ArrayOfResource struct {
	Resource []*Resource `xml:"Resource,omitempty" json:"Resource,omitempty" yaml:"Resource,omitempty"`
}

// ArrayOfSale was auto-generated from WSDL.
type ArrayOfSale struct {
	Sale []*Sale `xml:"Sale,omitempty" json:"Sale,omitempty" yaml:"Sale,omitempty"`
}

// ArrayOfSalesRep was auto-generated from WSDL.
type ArrayOfSalesRep struct {
	SalesRep []*SalesRep `xml:"SalesRep,omitempty" json:"SalesRep,omitempty" yaml:"SalesRep,omitempty"`
}

// ArrayOfService was auto-generated from WSDL.
type ArrayOfService struct {
	Service []*Service `xml:"Service,omitempty" json:"Service,omitempty" yaml:"Service,omitempty"`
}

// ArrayOfString was auto-generated from WSDL.
type ArrayOfString struct {
	String []string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// ArrayOfUnavailability was auto-generated from WSDL.
type ArrayOfUnavailability struct {
	Unavailability []*Unavailability `xml:"Unavailability,omitempty" json:"Unavailability,omitempty" yaml:"Unavailability,omitempty"`
}

// ArrayOfVisit was auto-generated from WSDL.
type ArrayOfVisit struct {
	Visit []*Visit `xml:"Visit,omitempty" json:"Visit,omitempty" yaml:"Visit,omitempty"`
}

// Availability was auto-generated from WSDL.
type Availability struct {
	Site                *Site           `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages            *ArrayOfString  `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute             string          `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode           string          `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID                  int             `xml:"ID" json:"ID" yaml:"ID"`
	Staff               *Staff          `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
	SessionType         *SessionType    `xml:"SessionType,omitempty" json:"SessionType,omitempty" yaml:"SessionType,omitempty"`
	Programs            *ArrayOfProgram `xml:"Programs,omitempty" json:"Programs,omitempty" yaml:"Programs,omitempty"`
	StartDateTime       DateTime        `xml:"StartDateTime" json:"StartDateTime" yaml:"StartDateTime"`
	EndDateTime         DateTime        `xml:"EndDateTime" json:"EndDateTime" yaml:"EndDateTime"`
	BookableEndDateTime DateTime        `xml:"BookableEndDateTime,omitempty" json:"BookableEndDateTime,omitempty" yaml:"BookableEndDateTime,omitempty"`
	Location            *Location       `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

// CartItem was auto-generated from WSDL.
type CartItem struct {
	Site           *Site               `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages       *ArrayOfString      `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute        string              `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode      string              `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Item           *Item               `xml:"Item,omitempty" json:"Item,omitempty" yaml:"Item,omitempty"`
	DiscountAmount float64             `xml:"DiscountAmount" json:"DiscountAmount" yaml:"DiscountAmount"`
	Appointments   *ArrayOfAppointment `xml:"Appointments,omitempty" json:"Appointments,omitempty" yaml:"Appointments,omitempty"`
	EnrollmentIDs  *ArrayOfInt         `xml:"EnrollmentIDs,omitempty" json:"EnrollmentIDs,omitempty" yaml:"EnrollmentIDs,omitempty"`
	ClassIDs       *ArrayOfInt         `xml:"ClassIDs,omitempty" json:"ClassIDs,omitempty" yaml:"ClassIDs,omitempty"`
	CourseIDs      *ArrayOfLong        `xml:"CourseIDs,omitempty" json:"CourseIDs,omitempty" yaml:"CourseIDs,omitempty"`
	VisitIDs       *ArrayOfLong        `xml:"VisitIDs,omitempty" json:"VisitIDs,omitempty" yaml:"VisitIDs,omitempty"`
	AppointmentIDs *ArrayOfLong        `xml:"AppointmentIDs,omitempty" json:"AppointmentIDs,omitempty" yaml:"AppointmentIDs,omitempty"`
	Action         ActionCode          `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID             int                 `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Quantity       int                 `xml:"Quantity,omitempty" json:"Quantity,omitempty" yaml:"Quantity,omitempty"`
}

// CashInfo was auto-generated from WSDL.
type CashInfo struct {
	Name   string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	Notes  string  `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

// CheckInfo was auto-generated from WSDL.
type CheckInfo struct {
	Name   string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	Notes  string  `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

// CheckoutShoppingCart was auto-generated from WSDL.
type CheckoutShoppingCart struct {
	XMLName xml.Name                     `xml:"http://clients.mindbodyonline.com/api/0_5 CheckoutShoppingCart" json:"-" yaml:"-"`
	Request *CheckoutShoppingCartRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// CheckoutShoppingCartRequest was auto-generated from WSDL.
type CheckoutShoppingCartRequest struct {
	SourceCredentials *SourceCredentials  `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials    `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel      `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                 `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                 `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString      `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	CartID            string              `xml:"CartID,omitempty" json:"CartID,omitempty" yaml:"CartID,omitempty"`
	ClientID          string              `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	Test              bool                `xml:"Test,omitempty" json:"Test,omitempty" yaml:"Test,omitempty"`
	CartItems         *ArrayOfCartItem    `xml:"CartItems,omitempty" json:"CartItems,omitempty" yaml:"CartItems,omitempty"`
	InStore           bool                `xml:"InStore,omitempty" json:"InStore,omitempty" yaml:"InStore,omitempty"`
	PromotionCode     string              `xml:"PromotionCode,omitempty" json:"PromotionCode,omitempty" yaml:"PromotionCode,omitempty"`
	Payments          *ArrayOfPaymentInfo `xml:"Payments,omitempty" json:"Payments,omitempty" yaml:"Payments,omitempty"`
	SendEmail         bool                `xml:"SendEmail,omitempty" json:"SendEmail,omitempty" yaml:"SendEmail,omitempty"`
	LocationID        int                 `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
	Image             []byte              `xml:"Image,omitempty" json:"Image,omitempty" yaml:"Image,omitempty"`
	ImageFileName     string              `xml:"ImageFileName,omitempty" json:"ImageFileName,omitempty" yaml:"ImageFileName,omitempty"`
}

// CheckoutShoppingCartResponse was auto-generated from WSDL.
type CheckoutShoppingCartResponse struct {
	CheckoutShoppingCartResult *CheckoutShoppingCartResult `xml:"CheckoutShoppingCartResult,omitempty" json:"CheckoutShoppingCartResult,omitempty" yaml:"CheckoutShoppingCartResult,omitempty"`
}

// CheckoutShoppingCartResult was auto-generated from WSDL.
type CheckoutShoppingCartResult struct {
	Status           StatusCode            `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                   `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel        `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                   `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                   `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                   `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ShoppingCart     *ShoppingCart         `xml:"ShoppingCart,omitempty" json:"ShoppingCart,omitempty" yaml:"ShoppingCart,omitempty"`
	Classes          *ArrayOfClass         `xml:"Classes,omitempty" json:"Classes,omitempty" yaml:"Classes,omitempty"`
	Appointments     *ArrayOfAppointment   `xml:"Appointments,omitempty" json:"Appointments,omitempty" yaml:"Appointments,omitempty"`
	Enrollments      *ArrayOfClassSchedule `xml:"Enrollments,omitempty" json:"Enrollments,omitempty" yaml:"Enrollments,omitempty"`
}

// Class was auto-generated from WSDL.
type Class struct {
	Site                *Site             `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages            *ArrayOfString    `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute             string            `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode           string            `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ClassScheduleID     int               `xml:"ClassScheduleID,omitempty" json:"ClassScheduleID,omitempty" yaml:"ClassScheduleID,omitempty"`
	Visits              *ArrayOfVisit     `xml:"Visits,omitempty" json:"Visits,omitempty" yaml:"Visits,omitempty"`
	Clients             *ArrayOfClient    `xml:"Clients,omitempty" json:"Clients,omitempty" yaml:"Clients,omitempty"`
	Location            *Location         `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Resource            *Resource         `xml:"Resource,omitempty" json:"Resource,omitempty" yaml:"Resource,omitempty"`
	MaxCapacity         int               `xml:"MaxCapacity,omitempty" json:"MaxCapacity,omitempty" yaml:"MaxCapacity,omitempty"`
	WebCapacity         int               `xml:"WebCapacity,omitempty" json:"WebCapacity,omitempty" yaml:"WebCapacity,omitempty"`
	TotalBooked         int               `xml:"TotalBooked,omitempty" json:"TotalBooked,omitempty" yaml:"TotalBooked,omitempty"`
	TotalBookedWaitlist int               `xml:"TotalBookedWaitlist,omitempty" json:"TotalBookedWaitlist,omitempty" yaml:"TotalBookedWaitlist,omitempty"`
	WebBooked           int               `xml:"WebBooked,omitempty" json:"WebBooked,omitempty" yaml:"WebBooked,omitempty"`
	SemesterID          int               `xml:"SemesterID,omitempty" json:"SemesterID,omitempty" yaml:"SemesterID,omitempty"`
	IsCanceled          bool              `xml:"IsCanceled,omitempty" json:"IsCanceled,omitempty" yaml:"IsCanceled,omitempty"`
	Substitute          bool              `xml:"Substitute,omitempty" json:"Substitute,omitempty" yaml:"Substitute,omitempty"`
	Active              bool              `xml:"Active,omitempty" json:"Active,omitempty" yaml:"Active,omitempty"`
	IsWaitlistAvailable bool              `xml:"IsWaitlistAvailable,omitempty" json:"IsWaitlistAvailable,omitempty" yaml:"IsWaitlistAvailable,omitempty"`
	IsEnrolled          bool              `xml:"IsEnrolled,omitempty" json:"IsEnrolled,omitempty" yaml:"IsEnrolled,omitempty"`
	HideCancel          bool              `xml:"HideCancel,omitempty" json:"HideCancel,omitempty" yaml:"HideCancel,omitempty"`
	Action              ActionCode        `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                  int               `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	IsAvailable         bool              `xml:"IsAvailable,omitempty" json:"IsAvailable,omitempty" yaml:"IsAvailable,omitempty"`
	StartDateTime       DateTime          `xml:"StartDateTime,omitempty" json:"StartDateTime,omitempty" yaml:"StartDateTime,omitempty"`
	EndDateTime         DateTime          `xml:"EndDateTime,omitempty" json:"EndDateTime,omitempty" yaml:"EndDateTime,omitempty"`
	ClassDescription    *ClassDescription `xml:"ClassDescription,omitempty" json:"ClassDescription,omitempty" yaml:"ClassDescription,omitempty"`
	Staff               *Staff            `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
}

// ClassDescription was auto-generated from WSDL.
type ClassDescription struct {
	Site        *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages    *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute     string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode   string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ImageURL    string         `xml:"ImageURL,omitempty" json:"ImageURL,omitempty" yaml:"ImageURL,omitempty"`
	Level       *Level         `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`
	Action      ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID          int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name        string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description string         `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Prereq      string         `xml:"Prereq,omitempty" json:"Prereq,omitempty" yaml:"Prereq,omitempty"`
	Notes       string         `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	LastUpdated DateTime       `xml:"LastUpdated,omitempty" json:"LastUpdated,omitempty" yaml:"LastUpdated,omitempty"`
	Program     *Program       `xml:"Program,omitempty" json:"Program,omitempty" yaml:"Program,omitempty"`
	SessionType *SessionType   `xml:"SessionType,omitempty" json:"SessionType,omitempty" yaml:"SessionType,omitempty"`
	Active      bool           `xml:"Active,omitempty" json:"Active,omitempty" yaml:"Active,omitempty"`
}

// ClassSchedule was auto-generated from WSDL.
type ClassSchedule struct {
	Site                       *Site             `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages                   *ArrayOfString    `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute                    string            `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode                  string            `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Classes                    *ArrayOfClass     `xml:"Classes,omitempty" json:"Classes,omitempty" yaml:"Classes,omitempty"`
	Clients                    *ArrayOfClient    `xml:"Clients,omitempty" json:"Clients,omitempty" yaml:"Clients,omitempty"`
	Course                     *Course           `xml:"Course,omitempty" json:"Course,omitempty" yaml:"Course,omitempty"`
	SemesterID                 int               `xml:"SemesterID,omitempty" json:"SemesterID,omitempty" yaml:"SemesterID,omitempty"`
	IsAvailable                bool              `xml:"IsAvailable,omitempty" json:"IsAvailable,omitempty" yaml:"IsAvailable,omitempty"`
	Action                     ActionCode        `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                         int               `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	ClassDescription           *ClassDescription `xml:"ClassDescription,omitempty" json:"ClassDescription,omitempty" yaml:"ClassDescription,omitempty"`
	DaySunday                  bool              `xml:"DaySunday,omitempty" json:"DaySunday,omitempty" yaml:"DaySunday,omitempty"`
	DayMonday                  bool              `xml:"DayMonday,omitempty" json:"DayMonday,omitempty" yaml:"DayMonday,omitempty"`
	DayTuesday                 bool              `xml:"DayTuesday,omitempty" json:"DayTuesday,omitempty" yaml:"DayTuesday,omitempty"`
	DayWednesday               bool              `xml:"DayWednesday,omitempty" json:"DayWednesday,omitempty" yaml:"DayWednesday,omitempty"`
	DayThursday                bool              `xml:"DayThursday,omitempty" json:"DayThursday,omitempty" yaml:"DayThursday,omitempty"`
	DayFriday                  bool              `xml:"DayFriday,omitempty" json:"DayFriday,omitempty" yaml:"DayFriday,omitempty"`
	DaySaturday                bool              `xml:"DaySaturday,omitempty" json:"DaySaturday,omitempty" yaml:"DaySaturday,omitempty"`
	AllowOpenEnrollment        bool              `xml:"AllowOpenEnrollment,omitempty" json:"AllowOpenEnrollment,omitempty" yaml:"AllowOpenEnrollment,omitempty"`
	AllowDateForwardEnrollment bool              `xml:"AllowDateForwardEnrollment,omitempty" json:"AllowDateForwardEnrollment,omitempty" yaml:"AllowDateForwardEnrollment,omitempty"`
	StartTime                  DateTime          `xml:"StartTime,omitempty" json:"StartTime,omitempty" yaml:"StartTime,omitempty"`
	EndTime                    DateTime          `xml:"EndTime,omitempty" json:"EndTime,omitempty" yaml:"EndTime,omitempty"`
	StartDate                  DateTime          `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate                    DateTime          `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
	Staff                      *Staff            `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
	Location                   *Location         `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

// Client was auto-generated from WSDL.
type Client struct {
	Site                             *Site                      `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages                         *ArrayOfString             `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute                          string                     `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode                        string                     `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	NewID                            string                     `xml:"NewID,omitempty" json:"NewID,omitempty" yaml:"NewID,omitempty"`
	AccountBalance                   float64                    `xml:"AccountBalance,omitempty" json:"AccountBalance,omitempty" yaml:"AccountBalance,omitempty"`
	ClientIndexes                    *ArrayOfClientIndex        `xml:"ClientIndexes,omitempty" json:"ClientIndexes,omitempty" yaml:"ClientIndexes,omitempty"`
	Username                         string                     `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password                         string                     `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	Notes                            string                     `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	MobileProvider                   int                        `xml:"MobileProvider,omitempty" json:"MobileProvider,omitempty" yaml:"MobileProvider,omitempty"`
	ClientCreditCard                 *ClientCreditCard          `xml:"ClientCreditCard,omitempty" json:"ClientCreditCard,omitempty" yaml:"ClientCreditCard,omitempty"`
	LastFormulaNotes                 string                     `xml:"LastFormulaNotes,omitempty" json:"LastFormulaNotes,omitempty" yaml:"LastFormulaNotes,omitempty"`
	AppointmentGenderPreference      string                     `xml:"AppointmentGenderPreference,omitempty" json:"AppointmentGenderPreference,omitempty" yaml:"AppointmentGenderPreference,omitempty"`
	Gender                           string                     `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`
	IsCompany                        bool                       `xml:"IsCompany,omitempty" json:"IsCompany,omitempty" yaml:"IsCompany,omitempty"`
	Inactive                         bool                       `xml:"Inactive,omitempty" json:"Inactive,omitempty" yaml:"Inactive,omitempty"`
	ClientRelationships              *ArrayOfClientRelationship `xml:"ClientRelationships,omitempty" json:"ClientRelationships,omitempty" yaml:"ClientRelationships,omitempty"`
	Reps                             *ArrayOfRep                `xml:"Reps,omitempty" json:"Reps,omitempty" yaml:"Reps,omitempty"`
	SaleReps                         *ArrayOfSalesRep           `xml:"SaleReps,omitempty" json:"SaleReps,omitempty" yaml:"SaleReps,omitempty"`
	CustomClientFields               *ArrayOfCustomClientField  `xml:"CustomClientFields,omitempty" json:"CustomClientFields,omitempty" yaml:"CustomClientFields,omitempty"`
	LiabilityRelease                 bool                       `xml:"LiabilityRelease,omitempty" json:"LiabilityRelease,omitempty" yaml:"LiabilityRelease,omitempty"`
	EmergencyContactInfoName         string                     `xml:"EmergencyContactInfoName,omitempty" json:"EmergencyContactInfoName,omitempty" yaml:"EmergencyContactInfoName,omitempty"`
	EmergencyContactInfoRelationship string                     `xml:"EmergencyContactInfoRelationship,omitempty" json:"EmergencyContactInfoRelationship,omitempty" yaml:"EmergencyContactInfoRelationship,omitempty"`
	EmergencyContactInfoPhone        string                     `xml:"EmergencyContactInfoPhone,omitempty" json:"EmergencyContactInfoPhone,omitempty" yaml:"EmergencyContactInfoPhone,omitempty"`
	EmergencyContactInfoEmail        string                     `xml:"EmergencyContactInfoEmail,omitempty" json:"EmergencyContactInfoEmail,omitempty" yaml:"EmergencyContactInfoEmail,omitempty"`
	PromotionalEmailOptIn            bool                       `xml:"PromotionalEmailOptIn,omitempty" json:"PromotionalEmailOptIn,omitempty" yaml:"PromotionalEmailOptIn,omitempty"`
	CreationDate                     DateTime                   `xml:"CreationDate,omitempty" json:"CreationDate,omitempty" yaml:"CreationDate,omitempty"`
	Liability                        *Liability                 `xml:"Liability,omitempty" json:"Liability,omitempty" yaml:"Liability,omitempty"`
	ProspectStage                    *ProspectStage             `xml:"ProspectStage,omitempty" json:"ProspectStage,omitempty" yaml:"ProspectStage,omitempty"`
	UniqueID                         int64                      `xml:"UniqueID,omitempty" json:"UniqueID,omitempty" yaml:"UniqueID,omitempty"`
	MembershipIcon                   int                        `xml:"MembershipIcon,omitempty" json:"MembershipIcon,omitempty" yaml:"MembershipIcon,omitempty"`
	Action                           ActionCode                 `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                               string                     `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	FirstName                        string                     `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	MiddleName                       string                     `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`
	LastName                         string                     `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
	Email                            string                     `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`
	EmailOptIn                       bool                       `xml:"EmailOptIn,omitempty" json:"EmailOptIn,omitempty" yaml:"EmailOptIn,omitempty"`
	AddressLine1                     string                     `xml:"AddressLine1,omitempty" json:"AddressLine1,omitempty" yaml:"AddressLine1,omitempty"`
	AddressLine2                     string                     `xml:"AddressLine2,omitempty" json:"AddressLine2,omitempty" yaml:"AddressLine2,omitempty"`
	City                             string                     `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	State                            string                     `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	PostalCode                       string                     `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`
	Country                          string                     `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	MobilePhone                      string                     `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`
	HomePhone                        string                     `xml:"HomePhone,omitempty" json:"HomePhone,omitempty" yaml:"HomePhone,omitempty"`
	WorkPhone                        string                     `xml:"WorkPhone,omitempty" json:"WorkPhone,omitempty" yaml:"WorkPhone,omitempty"`
	WorkExtension                    string                     `xml:"WorkExtension,omitempty" json:"WorkExtension,omitempty" yaml:"WorkExtension,omitempty"`
	BirthDate                        DateTime                   `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`
	FirstAppointmentDate             DateTime                   `xml:"FirstAppointmentDate,omitempty" json:"FirstAppointmentDate,omitempty" yaml:"FirstAppointmentDate,omitempty"`
	ReferredBy                       string                     `xml:"ReferredBy,omitempty" json:"ReferredBy,omitempty" yaml:"ReferredBy,omitempty"`
	HomeLocation                     *Location                  `xml:"HomeLocation,omitempty" json:"HomeLocation,omitempty" yaml:"HomeLocation,omitempty"`
	YellowAlert                      string                     `xml:"YellowAlert,omitempty" json:"YellowAlert,omitempty" yaml:"YellowAlert,omitempty"`
	RedAlert                         string                     `xml:"RedAlert,omitempty" json:"RedAlert,omitempty" yaml:"RedAlert,omitempty"`
	PhotoURL                         string                     `xml:"PhotoURL,omitempty" json:"PhotoURL,omitempty" yaml:"PhotoURL,omitempty"`
	IsProspect                       bool                       `xml:"IsProspect,omitempty" json:"IsProspect,omitempty" yaml:"IsProspect,omitempty"`
	Status                           string                     `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`
	ContactMethod                    *Short                     `xml:"ContactMethod,omitempty" json:"ContactMethod,omitempty" yaml:"ContactMethod,omitempty"`
}

// ClientCreditCard was auto-generated from WSDL.
type ClientCreditCard struct {
	CardType   string `xml:"CardType,omitempty" json:"CardType,omitempty" yaml:"CardType,omitempty"`
	LastFour   string `xml:"LastFour,omitempty" json:"LastFour,omitempty" yaml:"LastFour,omitempty"`
	CardNumber string `xml:"CardNumber,omitempty" json:"CardNumber,omitempty" yaml:"CardNumber,omitempty"`
	CardHolder string `xml:"CardHolder,omitempty" json:"CardHolder,omitempty" yaml:"CardHolder,omitempty"`
	ExpMonth   string `xml:"ExpMonth,omitempty" json:"ExpMonth,omitempty" yaml:"ExpMonth,omitempty"`
	ExpYear    string `xml:"ExpYear,omitempty" json:"ExpYear,omitempty" yaml:"ExpYear,omitempty"`
	Address    string `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	City       string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	State      string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	PostalCode string `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`
}

// ClientIndex was auto-generated from WSDL.
type ClientIndex struct {
	Site                 *Site                    `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages             *ArrayOfString           `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute              string                   `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode            string                   `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	RequiredBusinessMode bool                     `xml:"RequiredBusinessMode,omitempty" json:"RequiredBusinessMode,omitempty" yaml:"RequiredBusinessMode,omitempty"`
	RequiredConsumerMode bool                     `xml:"RequiredConsumerMode,omitempty" json:"RequiredConsumerMode,omitempty" yaml:"RequiredConsumerMode,omitempty"`
	Action               ActionCode               `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                   int                      `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name                 string                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Values               *ArrayOfClientIndexValue `xml:"Values,omitempty" json:"Values,omitempty" yaml:"Values,omitempty"`
}

// ClientIndexValue was auto-generated from WSDL.
type ClientIndexValue struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Action    ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID        int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name      string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Active    bool           `xml:"Active,omitempty" json:"Active,omitempty" yaml:"Active,omitempty"`
}

// ClientRelationship was auto-generated from WSDL.
type ClientRelationship struct {
	Site             *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages         *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute          string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode        string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	RelatedClient    *Client        `xml:"RelatedClient,omitempty" json:"RelatedClient,omitempty" yaml:"RelatedClient,omitempty"`
	Relationship     *Relationship  `xml:"Relationship,omitempty" json:"Relationship,omitempty" yaml:"Relationship,omitempty"`
	RelationshipName string         `xml:"RelationshipName,omitempty" json:"RelationshipName,omitempty" yaml:"RelationshipName,omitempty"`
}

// ClientService was auto-generated from WSDL.
type ClientService struct {
	Site           *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages       *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute        string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode      string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Current        bool           `xml:"Current" json:"Current" yaml:"Current"`
	Count          int            `xml:"Count" json:"Count" yaml:"Count"`
	Remaining      int            `xml:"Remaining" json:"Remaining" yaml:"Remaining"`
	Action         ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID             int64          `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name           string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	PaymentDate    DateTime       `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`
	ActiveDate     DateTime       `xml:"ActiveDate,omitempty" json:"ActiveDate,omitempty" yaml:"ActiveDate,omitempty"`
	ExpirationDate DateTime       `xml:"ExpirationDate,omitempty" json:"ExpirationDate,omitempty" yaml:"ExpirationDate,omitempty"`
	Program        *Program       `xml:"Program,omitempty" json:"Program,omitempty" yaml:"Program,omitempty"`
}

// Color was auto-generated from WSDL.
type Color struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Action    ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID        int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name      string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// CompInfo was auto-generated from WSDL.
type CompInfo struct {
	Name   string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
}

// Course was auto-generated from WSDL.
type Course struct {
	ID          int64     `xml:"ID" json:"ID" yaml:"ID"`
	Name        string    `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description string    `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Notes       string    `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	StartDate   DateTime  `xml:"StartDate" json:"StartDate" yaml:"StartDate"`
	EndDate     DateTime  `xml:"EndDate" json:"EndDate" yaml:"EndDate"`
	Location    *Location `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Organizer   *Staff    `xml:"Organizer,omitempty" json:"Organizer,omitempty" yaml:"Organizer,omitempty"`
	Program     *Program  `xml:"Program,omitempty" json:"Program,omitempty" yaml:"Program,omitempty"`
	ImageURL    string    `xml:"ImageURL,omitempty" json:"ImageURL,omitempty" yaml:"ImageURL,omitempty"`
}

// CreditCardInfo was auto-generated from WSDL.
type CreditCardInfo struct {
	Name              string     `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	CVV               string     `xml:"CVV,omitempty" json:"CVV,omitempty" yaml:"CVV,omitempty"`
	Action            ActionCode `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	CreditCardNumber  string     `xml:"CreditCardNumber,omitempty" json:"CreditCardNumber,omitempty" yaml:"CreditCardNumber,omitempty"`
	Amount            float64    `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	ExpMonth          string     `xml:"ExpMonth,omitempty" json:"ExpMonth,omitempty" yaml:"ExpMonth,omitempty"`
	ExpYear           string     `xml:"ExpYear,omitempty" json:"ExpYear,omitempty" yaml:"ExpYear,omitempty"`
	BillingName       string     `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`
	BillingAddress    string     `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`
	BillingCity       string     `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`
	BillingState      string     `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`
	BillingPostalCode string     `xml:"BillingPostalCode,omitempty" json:"BillingPostalCode,omitempty" yaml:"BillingPostalCode,omitempty"`
	SaveInfo          bool       `xml:"SaveInfo,omitempty" json:"SaveInfo,omitempty" yaml:"SaveInfo,omitempty"`
}

// CustomClientField was auto-generated from WSDL.
type CustomClientField struct {
	ID       int    `xml:"ID" json:"ID" yaml:"ID"`
	DataType string `xml:"DataType,omitempty" json:"DataType,omitempty" yaml:"DataType,omitempty"`
	Name     string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Value    string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// CustomPaymentInfo was auto-generated from WSDL.
type CustomPaymentInfo struct {
	Name   string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	ID     int     `xml:"ID" json:"ID" yaml:"ID"`
}

// DebitAccountInfo was auto-generated from WSDL.
type DebitAccountInfo struct {
	Name   string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
}

// EncryptedTrackDataInfo was auto-generated from WSDL.
type EncryptedTrackDataInfo struct {
	Name      string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount    float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	TrackData string  `xml:"TrackData,omitempty" json:"TrackData,omitempty" yaml:"TrackData,omitempty"`
}

// GetAcceptedCardType was auto-generated from WSDL.
type GetAcceptedCardType struct {
	XMLName xml.Name                    `xml:"http://clients.mindbodyonline.com/api/0_5 GetAcceptedCardType" json:"-" yaml:"-"`
	Request *GetAcceptedCardTypeRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetAcceptedCardTypeRequest was auto-generated from WSDL.
type GetAcceptedCardTypeRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetAcceptedCardTypeResponse was auto-generated from WSDL.
type GetAcceptedCardTypeResponse struct {
	GetAcceptedCardTypeResult *GetAcceptedCardTypeResult `xml:"GetAcceptedCardTypeResult,omitempty" json:"GetAcceptedCardTypeResult,omitempty" yaml:"GetAcceptedCardTypeResult,omitempty"`
}

// GetAcceptedCardTypeResult was auto-generated from WSDL.
type GetAcceptedCardTypeResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	CardTypes        *ArrayOfString `xml:"CardTypes,omitempty" json:"CardTypes,omitempty" yaml:"CardTypes,omitempty"`
}

// GetCustomPaymentMethods was auto-generated from WSDL.
type GetCustomPaymentMethods struct {
	XMLName xml.Name                        `xml:"http://clients.mindbodyonline.com/api/0_5 GetCustomPaymentMethods" json:"-" yaml:"-"`
	Request *GetCustomPaymentMethodsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetCustomPaymentMethodsRequest was auto-generated from WSDL.
type GetCustomPaymentMethodsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetCustomPaymentMethodsResponse was auto-generated from WSDL.
type GetCustomPaymentMethodsResponse struct {
	GetCustomPaymentMethodsResult *GetCustomPaymentMethodsResult `xml:"GetCustomPaymentMethodsResult,omitempty" json:"GetCustomPaymentMethodsResult,omitempty" yaml:"GetCustomPaymentMethodsResult,omitempty"`
}

// GetCustomPaymentMethodsResult was auto-generated from WSDL.
type GetCustomPaymentMethodsResult struct {
	Status           StatusCode                `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                       `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                    `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel            `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                       `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                       `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                       `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	PaymentMethods   *ArrayOfCustomPaymentInfo `xml:"PaymentMethods,omitempty" json:"PaymentMethods,omitempty" yaml:"PaymentMethods,omitempty"`
}

// GetPackages was auto-generated from WSDL.
type GetPackages struct {
	XMLName xml.Name            `xml:"http://clients.mindbodyonline.com/api/0_5 GetPackages" json:"-" yaml:"-"`
	Request *GetPackagesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetPackagesRequest was auto-generated from WSDL.
type GetPackagesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	PackageIDs        *ArrayOfInt        `xml:"PackageIDs,omitempty" json:"PackageIDs,omitempty" yaml:"PackageIDs,omitempty"`
	SellOnline        bool               `xml:"SellOnline" json:"SellOnline" yaml:"SellOnline"`
}

// GetPackagesResponse was auto-generated from WSDL.
type GetPackagesResponse struct {
	GetPackagesResult *GetPackagesResult `xml:"GetPackagesResult,omitempty" json:"GetPackagesResult,omitempty" yaml:"GetPackagesResult,omitempty"`
}

// GetPackagesResult was auto-generated from WSDL.
type GetPackagesResult struct {
	Status           StatusCode      `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int             `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string          `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel  `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int             `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int             `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int             `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Packages         *ArrayOfPackage `xml:"Packages,omitempty" json:"Packages,omitempty" yaml:"Packages,omitempty"`
}

// GetProducts was auto-generated from WSDL.
type GetProducts struct {
	XMLName xml.Name            `xml:"http://clients.mindbodyonline.com/api/0_5 GetProducts" json:"-" yaml:"-"`
	Request *GetProductsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetProductsRequest was auto-generated from WSDL.
type GetProductsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ProductIDs        *ArrayOfString     `xml:"ProductIDs,omitempty" json:"ProductIDs,omitempty" yaml:"ProductIDs,omitempty"`
	SearchText        string             `xml:"SearchText,omitempty" json:"SearchText,omitempty" yaml:"SearchText,omitempty"`
	SearchDomain      string             `xml:"SearchDomain,omitempty" json:"SearchDomain,omitempty" yaml:"SearchDomain,omitempty"`
	CategoryIDs       *ArrayOfInt        `xml:"CategoryIDs,omitempty" json:"CategoryIDs,omitempty" yaml:"CategoryIDs,omitempty"`
	SubCategoryIDs    *ArrayOfInt        `xml:"SubCategoryIDs,omitempty" json:"SubCategoryIDs,omitempty" yaml:"SubCategoryIDs,omitempty"`
	SellOnline        bool               `xml:"SellOnline" json:"SellOnline" yaml:"SellOnline"`
	LocationID        int                `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// GetProductsResponse was auto-generated from WSDL.
type GetProductsResponse struct {
	GetProductsResult *GetProductsResult `xml:"GetProductsResult,omitempty" json:"GetProductsResult,omitempty" yaml:"GetProductsResult,omitempty"`
}

// GetProductsResult was auto-generated from WSDL.
type GetProductsResult struct {
	Status           StatusCode      `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int             `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string          `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel  `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int             `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int             `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int             `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Products         *ArrayOfProduct `xml:"Products,omitempty" json:"Products,omitempty" yaml:"Products,omitempty"`
}

// GetSales was auto-generated from WSDL.
type GetSales struct {
	XMLName xml.Name         `xml:"http://clients.mindbodyonline.com/api/0_5 GetSales" json:"-" yaml:"-"`
	Request *GetSalesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetSalesRequest was auto-generated from WSDL.
type GetSalesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	SaleID            int64              `xml:"SaleID,omitempty" json:"SaleID,omitempty" yaml:"SaleID,omitempty"`
	StartSaleDateTime DateTime           `xml:"StartSaleDateTime,omitempty" json:"StartSaleDateTime,omitempty" yaml:"StartSaleDateTime,omitempty"`
	EndSaleDateTime   DateTime           `xml:"EndSaleDateTime,omitempty" json:"EndSaleDateTime,omitempty" yaml:"EndSaleDateTime,omitempty"`
	PaymentMethodID   int                `xml:"PaymentMethodID,omitempty" json:"PaymentMethodID,omitempty" yaml:"PaymentMethodID,omitempty"`
}

// GetSalesResponse was auto-generated from WSDL.
type GetSalesResponse struct {
	GetSalesResult *GetSalesResult `xml:"GetSalesResult,omitempty" json:"GetSalesResult,omitempty" yaml:"GetSalesResult,omitempty"`
}

// GetSalesResult was auto-generated from WSDL.
type GetSalesResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Sales            *ArrayOfSale   `xml:"Sales,omitempty" json:"Sales,omitempty" yaml:"Sales,omitempty"`
}

// GetServices was auto-generated from WSDL.
type GetServices struct {
	XMLName xml.Name            `xml:"http://clients.mindbodyonline.com/api/0_5 GetServices" json:"-" yaml:"-"`
	Request *GetServicesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetServicesRequest was auto-generated from WSDL.
type GetServicesRequest struct {
	SourceCredentials   *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials     *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail           XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize            int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex    int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields              *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ProgramIDs          *ArrayOfInt        `xml:"ProgramIDs,omitempty" json:"ProgramIDs,omitempty" yaml:"ProgramIDs,omitempty"`
	SessionTypeIDs      *ArrayOfInt        `xml:"SessionTypeIDs,omitempty" json:"SessionTypeIDs,omitempty" yaml:"SessionTypeIDs,omitempty"`
	ServiceIDs          *ArrayOfString     `xml:"ServiceIDs,omitempty" json:"ServiceIDs,omitempty" yaml:"ServiceIDs,omitempty"`
	ClassID             int                `xml:"ClassID,omitempty" json:"ClassID,omitempty" yaml:"ClassID,omitempty"`
	ClassScheduleID     int                `xml:"ClassScheduleID,omitempty" json:"ClassScheduleID,omitempty" yaml:"ClassScheduleID,omitempty"`
	SellOnline          bool               `xml:"SellOnline,omitempty" json:"SellOnline,omitempty" yaml:"SellOnline,omitempty"`
	LocationID          int                `xml:"LocationID" json:"LocationID" yaml:"LocationID"`
	HideRelatedPrograms bool               `xml:"HideRelatedPrograms" json:"HideRelatedPrograms" yaml:"HideRelatedPrograms"`
	StaffID             int64              `xml:"StaffID,omitempty" json:"StaffID,omitempty" yaml:"StaffID,omitempty"`
}

// GetServicesResponse was auto-generated from WSDL.
type GetServicesResponse struct {
	GetServicesResult *GetServicesResult `xml:"GetServicesResult,omitempty" json:"GetServicesResult,omitempty" yaml:"GetServicesResult,omitempty"`
}

// GetServicesResult was auto-generated from WSDL.
type GetServicesResult struct {
	Status           StatusCode      `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int             `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string          `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel  `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int             `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int             `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int             `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Services         *ArrayOfService `xml:"Services,omitempty" json:"Services,omitempty" yaml:"Services,omitempty"`
}

// GiftCardInfo was auto-generated from WSDL.
type GiftCardInfo struct {
	Name       string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount     float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	Notes      string  `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	CardNumber string  `xml:"CardNumber,omitempty" json:"CardNumber,omitempty" yaml:"CardNumber,omitempty"`
}

// Item was auto-generated from WSDL.
type Item struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
}

// Level was auto-generated from WSDL.
type Level struct {
	ID          int    `xml:"ID" json:"ID" yaml:"ID"`
	Name        string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// Liability was auto-generated from WSDL.
type Liability struct {
	IsReleased    bool     `xml:"IsReleased" json:"IsReleased" yaml:"IsReleased"`
	AgreementDate DateTime `xml:"AgreementDate,omitempty" json:"AgreementDate,omitempty" yaml:"AgreementDate,omitempty"`
	ReleasedBy    int64    `xml:"ReleasedBy,omitempty" json:"ReleasedBy,omitempty" yaml:"ReleasedBy,omitempty"`
}

// Location was auto-generated from WSDL.
type Location struct {
	Site                *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages            *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute             string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode           string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	BusinessID          int            `xml:"BusinessID,omitempty" json:"BusinessID,omitempty" yaml:"BusinessID,omitempty"`
	SiteID              int            `xml:"SiteID,omitempty" json:"SiteID,omitempty" yaml:"SiteID,omitempty"`
	BusinessDescription string         `xml:"BusinessDescription,omitempty" json:"BusinessDescription,omitempty" yaml:"BusinessDescription,omitempty"`
	AdditionalImageURLs *ArrayOfString `xml:"AdditionalImageURLs,omitempty" json:"AdditionalImageURLs,omitempty" yaml:"AdditionalImageURLs,omitempty"`
	FacilitySquareFeet  int            `xml:"FacilitySquareFeet,omitempty" json:"FacilitySquareFeet,omitempty" yaml:"FacilitySquareFeet,omitempty"`
	TreatmentRooms      int            `xml:"TreatmentRooms,omitempty" json:"TreatmentRooms,omitempty" yaml:"TreatmentRooms,omitempty"`
	ProSpaFinderSite    bool           `xml:"ProSpaFinderSite,omitempty" json:"ProSpaFinderSite,omitempty" yaml:"ProSpaFinderSite,omitempty"`
	HasClasses          bool           `xml:"HasClasses,omitempty" json:"HasClasses,omitempty" yaml:"HasClasses,omitempty"`
	PhoneExtension      string         `xml:"PhoneExtension,omitempty" json:"PhoneExtension,omitempty" yaml:"PhoneExtension,omitempty"`
	Action              ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                  int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name                string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Address             string         `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	Address2            string         `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`
	Tax1                float64        `xml:"Tax1,omitempty" json:"Tax1,omitempty" yaml:"Tax1,omitempty"`
	Tax2                float64        `xml:"Tax2,omitempty" json:"Tax2,omitempty" yaml:"Tax2,omitempty"`
	Tax3                float64        `xml:"Tax3,omitempty" json:"Tax3,omitempty" yaml:"Tax3,omitempty"`
	Tax4                float64        `xml:"Tax4,omitempty" json:"Tax4,omitempty" yaml:"Tax4,omitempty"`
	Tax5                float64        `xml:"Tax5,omitempty" json:"Tax5,omitempty" yaml:"Tax5,omitempty"`
	Phone               string         `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`
	City                string         `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	StateProvCode       string         `xml:"StateProvCode,omitempty" json:"StateProvCode,omitempty" yaml:"StateProvCode,omitempty"`
	PostalCode          string         `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`
	Latitude            float64        `xml:"Latitude,omitempty" json:"Latitude,omitempty" yaml:"Latitude,omitempty"`
	Longitude           float64        `xml:"Longitude,omitempty" json:"Longitude,omitempty" yaml:"Longitude,omitempty"`
	DistanceInMiles     float64        `xml:"DistanceInMiles,omitempty" json:"DistanceInMiles,omitempty" yaml:"DistanceInMiles,omitempty"`
	ImageURL            string         `xml:"ImageURL,omitempty" json:"ImageURL,omitempty" yaml:"ImageURL,omitempty"`
	Description         string         `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	HasSite             bool           `xml:"HasSite,omitempty" json:"HasSite,omitempty" yaml:"HasSite,omitempty"`
	CanBook             bool           `xml:"CanBook,omitempty" json:"CanBook,omitempty" yaml:"CanBook,omitempty"`
}

// MBObject was auto-generated from WSDL.
type MBObject struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
}

// MBRequest was auto-generated from WSDL.
type MBRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// MBResult was auto-generated from WSDL.
type MBResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
}

// Package was auto-generated from WSDL.
type Package struct {
	Site               *Site           `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages           *ArrayOfString  `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute            string          `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode          string          `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID                 int             `xml:"ID" json:"ID" yaml:"ID"`
	Name               string          `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	DiscountPercentage float64         `xml:"DiscountPercentage" json:"DiscountPercentage" yaml:"DiscountPercentage"`
	SellOnline         bool            `xml:"SellOnline" json:"SellOnline" yaml:"SellOnline"`
	Services           *ArrayOfService `xml:"Services,omitempty" json:"Services,omitempty" yaml:"Services,omitempty"`
	Products           *ArrayOfProduct `xml:"Products,omitempty" json:"Products,omitempty" yaml:"Products,omitempty"`
}

// Payment was auto-generated from WSDL.
type Payment struct {
	ID       int64   `xml:"ID" json:"ID" yaml:"ID"`
	Amount   float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	Method   int     `xml:"Method" json:"Method" yaml:"Method"`
	Type     string  `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	LastFour string  `xml:"LastFour,omitempty" json:"LastFour,omitempty" yaml:"LastFour,omitempty"`
	Notes    string  `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

// PaymentInfo was auto-generated from WSDL.
type PaymentInfo struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// Product was auto-generated from WSDL.
type Product struct {
	Site        *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages    *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute     string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode   string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Price       float64        `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
	TaxIncluded float64        `xml:"TaxIncluded,omitempty" json:"TaxIncluded,omitempty" yaml:"TaxIncluded,omitempty"`
	TaxRate     float64        `xml:"TaxRate,omitempty" json:"TaxRate,omitempty" yaml:"TaxRate,omitempty"`
	Action      ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID          string         `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	GroupID     int            `xml:"GroupID,omitempty" json:"GroupID,omitempty" yaml:"GroupID,omitempty"`
	Name        string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	OnlinePrice float64        `xml:"OnlinePrice,omitempty" json:"OnlinePrice,omitempty" yaml:"OnlinePrice,omitempty"`
	ShortDesc   string         `xml:"ShortDesc,omitempty" json:"ShortDesc,omitempty" yaml:"ShortDesc,omitempty"`
	LongDesc    string         `xml:"LongDesc,omitempty" json:"LongDesc,omitempty" yaml:"LongDesc,omitempty"`
	Color       *Color         `xml:"Color,omitempty" json:"Color,omitempty" yaml:"Color,omitempty"`
	Size        *Size          `xml:"Size,omitempty" json:"Size,omitempty" yaml:"Size,omitempty"`
}

// Program was auto-generated from WSDL.
type Program struct {
	Site         *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages     *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute      string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode    string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID           int            `xml:"ID" json:"ID" yaml:"ID"`
	Name         string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	ScheduleType ScheduleType   `xml:"ScheduleType,omitempty" json:"ScheduleType,omitempty" yaml:"ScheduleType,omitempty"`
	CancelOffset int            `xml:"CancelOffset,omitempty" json:"CancelOffset,omitempty" yaml:"CancelOffset,omitempty"`
}

// ProspectStage was auto-generated from WSDL.
type ProspectStage struct {
	ID          int    `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Active      bool   `xml:"Active" json:"Active" yaml:"Active"`
}

// ProviderIDUpdate was auto-generated from WSDL.
type ProviderIDUpdate struct {
	OldProviderID string `xml:"OldProviderID,omitempty" json:"OldProviderID,omitempty" yaml:"OldProviderID,omitempty"`
	NewProviderID string `xml:"NewProviderID,omitempty" json:"NewProviderID,omitempty" yaml:"NewProviderID,omitempty"`
}

// RedeemSpaFinderWellnessCard was auto-generated from WSDL.
type RedeemSpaFinderWellnessCard struct {
	XMLName xml.Name                            `xml:"http://clients.mindbodyonline.com/api/0_5 RedeemSpaFinderWellnessCard" json:"-" yaml:"-"`
	Request *RedeemSpaFinderWellnessCardRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// RedeemSpaFinderWellnessCardRequest was auto-generated from WSDL.
type RedeemSpaFinderWellnessCardRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	CardID            string             `xml:"CardID,omitempty" json:"CardID,omitempty" yaml:"CardID,omitempty"`
	FaceAmount        float64            `xml:"FaceAmount" json:"FaceAmount" yaml:"FaceAmount"`
	Currency          string             `xml:"Currency,omitempty" json:"Currency,omitempty" yaml:"Currency,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	LocationID        int                `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// RedeemSpaFinderWellnessCardResponse was auto-generated from
// WSDL.
type RedeemSpaFinderWellnessCardResponse struct {
	RedeemSpaFinderWellnessCardResult *RedeemSpaFinderWellnessCardResult `xml:"RedeemSpaFinderWellnessCardResult,omitempty" json:"RedeemSpaFinderWellnessCardResult,omitempty" yaml:"RedeemSpaFinderWellnessCardResult,omitempty"`
}

// RedeemSpaFinderWellnessCardResult was auto-generated from WSDL.
type RedeemSpaFinderWellnessCardResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
}

// Relationship was auto-generated from WSDL.
type Relationship struct {
	ID                int    `xml:"ID" json:"ID" yaml:"ID"`
	RelationshipName1 string `xml:"RelationshipName1,omitempty" json:"RelationshipName1,omitempty" yaml:"RelationshipName1,omitempty"`
	RelationshipName2 string `xml:"RelationshipName2,omitempty" json:"RelationshipName2,omitempty" yaml:"RelationshipName2,omitempty"`
}

// Rep was auto-generated from WSDL.
type Rep struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID        int            `xml:"ID" json:"ID" yaml:"ID"`
	Staff     *Staff         `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
}

// Resource was auto-generated from WSDL.
type Resource struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Action    ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID        int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name      string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// ReturnSale was auto-generated from WSDL.
type ReturnSale struct {
	XMLName xml.Name           `xml:"http://clients.mindbodyonline.com/api/0_5 ReturnSale" json:"-" yaml:"-"`
	Request *ReturnSaleRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// ReturnSaleRequest was auto-generated from WSDL.
type ReturnSaleRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	Test              bool               `xml:"Test" json:"Test" yaml:"Test"`
	SaleID            int64              `xml:"SaleID,omitempty" json:"SaleID,omitempty" yaml:"SaleID,omitempty"`
	ReturnReason      string             `xml:"ReturnReason,omitempty" json:"ReturnReason,omitempty" yaml:"ReturnReason,omitempty"`
}

// ReturnSaleResponse was auto-generated from WSDL.
type ReturnSaleResponse struct {
	ReturnSaleResult *ReturnSaleResult `xml:"ReturnSaleResult,omitempty" json:"ReturnSaleResult,omitempty" yaml:"ReturnSaleResult,omitempty"`
}

// ReturnSaleResult was auto-generated from WSDL.
type ReturnSaleResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ReturnSaleID     int64          `xml:"ReturnSaleID,omitempty" json:"ReturnSaleID,omitempty" yaml:"ReturnSaleID,omitempty"`
	TrainerID        int64          `xml:"TrainerID,omitempty" json:"TrainerID,omitempty" yaml:"TrainerID,omitempty"`
	Amount           float64        `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
}

// Sale was auto-generated from WSDL.
type Sale struct {
	ID           int64           `xml:"ID" json:"ID" yaml:"ID"`
	SaleTime     DateTime        `xml:"SaleTime" json:"SaleTime" yaml:"SaleTime"`
	SaleDate     DateTime        `xml:"SaleDate" json:"SaleDate" yaml:"SaleDate"`
	SaleDateTime DateTime        `xml:"SaleDateTime" json:"SaleDateTime" yaml:"SaleDateTime"`
	Location     *Location       `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Payments     *ArrayOfPayment `xml:"Payments,omitempty" json:"Payments,omitempty" yaml:"Payments,omitempty"`
}

// SalesRep was auto-generated from WSDL.
type SalesRep struct {
	Site            *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages        *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute         string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode       string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	SalesRepNumber  int            `xml:"SalesRepNumber,omitempty" json:"SalesRepNumber,omitempty" yaml:"SalesRepNumber,omitempty"`
	ID              int64          `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	FirstName       string         `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	LastName        string         `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
	SalesRepNumbers *ArrayOfInt    `xml:"SalesRepNumbers,omitempty" json:"SalesRepNumbers,omitempty" yaml:"SalesRepNumbers,omitempty"`
}

// ScheduleItem was auto-generated from WSDL.
type ScheduleItem struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
}

// Service was auto-generated from WSDL.
type Service struct {
	Site        *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages    *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute     string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode   string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Price       float64        `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
	OnlinePrice float64        `xml:"OnlinePrice,omitempty" json:"OnlinePrice,omitempty" yaml:"OnlinePrice,omitempty"`
	TaxIncluded float64        `xml:"TaxIncluded,omitempty" json:"TaxIncluded,omitempty" yaml:"TaxIncluded,omitempty"`
	ProgramID   int            `xml:"ProgramID,omitempty" json:"ProgramID,omitempty" yaml:"ProgramID,omitempty"`
	TaxRate     float64        `xml:"TaxRate,omitempty" json:"TaxRate,omitempty" yaml:"TaxRate,omitempty"`
	ProductID   float64        `xml:"ProductID,omitempty" json:"ProductID,omitempty" yaml:"ProductID,omitempty"`
	Action      ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID          string         `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name        string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Count       int            `xml:"Count,omitempty" json:"Count,omitempty" yaml:"Count,omitempty"`
}

// SessionType was auto-generated from WSDL.
type SessionType struct {
	Site              *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages          *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute           string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode         string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	DefaultTimeLength int            `xml:"DefaultTimeLength,omitempty" json:"DefaultTimeLength,omitempty" yaml:"DefaultTimeLength,omitempty"`
	ProgramID         int            `xml:"ProgramID,omitempty" json:"ProgramID,omitempty" yaml:"ProgramID,omitempty"`
	NumDeducted       int            `xml:"NumDeducted,omitempty" json:"NumDeducted,omitempty" yaml:"NumDeducted,omitempty"`
	Action            ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name              string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// ShoppingCart was auto-generated from WSDL.
type ShoppingCart struct {
	Site          *Site            `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages      *ArrayOfString   `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute       string           `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode     string           `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	AuthCode      string           `xml:"AuthCode,omitempty" json:"AuthCode,omitempty" yaml:"AuthCode,omitempty"`
	Action        ActionCode       `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID            string           `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	CartItems     *ArrayOfCartItem `xml:"CartItems,omitempty" json:"CartItems,omitempty" yaml:"CartItems,omitempty"`
	SubTotal      float64          `xml:"SubTotal,omitempty" json:"SubTotal,omitempty" yaml:"SubTotal,omitempty"`
	DiscountTotal float64          `xml:"DiscountTotal,omitempty" json:"DiscountTotal,omitempty" yaml:"DiscountTotal,omitempty"`
	TaxTotal      float64          `xml:"TaxTotal,omitempty" json:"TaxTotal,omitempty" yaml:"TaxTotal,omitempty"`
	GrandTotal    float64          `xml:"GrandTotal,omitempty" json:"GrandTotal,omitempty" yaml:"GrandTotal,omitempty"`
}

// Site was auto-generated from WSDL.
type Site struct {
	ID                     int    `xml:"ID" json:"ID" yaml:"ID"`
	Name                   string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Description            string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	LogoURL                string `xml:"LogoURL,omitempty" json:"LogoURL,omitempty" yaml:"LogoURL,omitempty"`
	PageColor1             string `xml:"PageColor1,omitempty" json:"PageColor1,omitempty" yaml:"PageColor1,omitempty"`
	PageColor2             string `xml:"PageColor2,omitempty" json:"PageColor2,omitempty" yaml:"PageColor2,omitempty"`
	PageColor3             string `xml:"PageColor3,omitempty" json:"PageColor3,omitempty" yaml:"PageColor3,omitempty"`
	PageColor4             string `xml:"PageColor4,omitempty" json:"PageColor4,omitempty" yaml:"PageColor4,omitempty"`
	AcceptsVisa            bool   `xml:"AcceptsVisa,omitempty" json:"AcceptsVisa,omitempty" yaml:"AcceptsVisa,omitempty"`
	AcceptsDiscover        bool   `xml:"AcceptsDiscover,omitempty" json:"AcceptsDiscover,omitempty" yaml:"AcceptsDiscover,omitempty"`
	AcceptsMasterCard      bool   `xml:"AcceptsMasterCard,omitempty" json:"AcceptsMasterCard,omitempty" yaml:"AcceptsMasterCard,omitempty"`
	AcceptsAmericanExpress bool   `xml:"AcceptsAmericanExpress,omitempty" json:"AcceptsAmericanExpress,omitempty" yaml:"AcceptsAmericanExpress,omitempty"`
	ContactEmail           string `xml:"ContactEmail,omitempty" json:"ContactEmail,omitempty" yaml:"ContactEmail,omitempty"`
	ESA                    bool   `xml:"ESA,omitempty" json:"ESA,omitempty" yaml:"ESA,omitempty"`
	TotalWOD               bool   `xml:"TotalWOD,omitempty" json:"TotalWOD,omitempty" yaml:"TotalWOD,omitempty"`
	TaxInclusivePrices     bool   `xml:"TaxInclusivePrices,omitempty" json:"TaxInclusivePrices,omitempty" yaml:"TaxInclusivePrices,omitempty"`
	SMSPackageEnabled      bool   `xml:"SMSPackageEnabled,omitempty" json:"SMSPackageEnabled,omitempty" yaml:"SMSPackageEnabled,omitempty"`
	AllowsDashboardAccess  bool   `xml:"AllowsDashboardAccess,omitempty" json:"AllowsDashboardAccess,omitempty" yaml:"AllowsDashboardAccess,omitempty"`
	PricingLevel           string `xml:"PricingLevel,omitempty" json:"PricingLevel,omitempty" yaml:"PricingLevel,omitempty"`
}

// Size was auto-generated from WSDL.
type Size struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Action    ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID        int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name      string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// SourceCredentials was auto-generated from WSDL.
type SourceCredentials struct {
	SourceName string      `xml:"SourceName,omitempty" json:"SourceName,omitempty" yaml:"SourceName,omitempty"`
	Password   string      `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	SiteIDs    *ArrayOfInt `xml:"SiteIDs,omitempty" json:"SiteIDs,omitempty" yaml:"SiteIDs,omitempty"`
}

// Staff was auto-generated from WSDL.
type Staff struct {
	Site                     *Site                    `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages                 *ArrayOfString           `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute                  string                   `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode                string                   `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Appointments             *ArrayOfAppointment      `xml:"Appointments,omitempty" json:"Appointments,omitempty" yaml:"Appointments,omitempty"`
	Unavailabilities         *ArrayOfUnavailability   `xml:"Unavailabilities,omitempty" json:"Unavailabilities,omitempty" yaml:"Unavailabilities,omitempty"`
	Availabilities           *ArrayOfAvailability     `xml:"Availabilities,omitempty" json:"Availabilities,omitempty" yaml:"Availabilities,omitempty"`
	Email                    string                   `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`
	MobilePhone              string                   `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`
	HomePhone                string                   `xml:"HomePhone,omitempty" json:"HomePhone,omitempty" yaml:"HomePhone,omitempty"`
	WorkPhone                string                   `xml:"WorkPhone,omitempty" json:"WorkPhone,omitempty" yaml:"WorkPhone,omitempty"`
	Address                  string                   `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`
	Address2                 string                   `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`
	City                     string                   `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`
	State                    string                   `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`
	Country                  string                   `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
	PostalCode               string                   `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`
	ForeignZip               string                   `xml:"ForeignZip,omitempty" json:"ForeignZip,omitempty" yaml:"ForeignZip,omitempty"`
	SortOrder                int                      `xml:"SortOrder,omitempty" json:"SortOrder,omitempty" yaml:"SortOrder,omitempty"`
	LoginLocations           *ArrayOfLocation         `xml:"LoginLocations,omitempty" json:"LoginLocations,omitempty" yaml:"LoginLocations,omitempty"`
	MultiLocation            bool                     `xml:"MultiLocation,omitempty" json:"MultiLocation,omitempty" yaml:"MultiLocation,omitempty"`
	AppointmentTrn           bool                     `xml:"AppointmentTrn,omitempty" json:"AppointmentTrn,omitempty" yaml:"AppointmentTrn,omitempty"`
	ReservationTrn           bool                     `xml:"ReservationTrn,omitempty" json:"ReservationTrn,omitempty" yaml:"ReservationTrn,omitempty"`
	IndependentContractor    bool                     `xml:"IndependentContractor,omitempty" json:"IndependentContractor,omitempty" yaml:"IndependentContractor,omitempty"`
	AlwaysAllowDoubleBooking bool                     `xml:"AlwaysAllowDoubleBooking,omitempty" json:"AlwaysAllowDoubleBooking,omitempty" yaml:"AlwaysAllowDoubleBooking,omitempty"`
	UserAccessLevel          string                   `xml:"UserAccessLevel,omitempty" json:"UserAccessLevel,omitempty" yaml:"UserAccessLevel,omitempty"`
	ProviderIDs              *ArrayOfString           `xml:"ProviderIDs,omitempty" json:"ProviderIDs,omitempty" yaml:"ProviderIDs,omitempty"`
	ProviderIDUpdateList     *ArrayOfProviderIDUpdate `xml:"ProviderIDUpdateList,omitempty" json:"ProviderIDUpdateList,omitempty" yaml:"ProviderIDUpdateList,omitempty"`
	Action                   ActionCode               `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                       int64                    `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name                     string                   `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	FirstName                string                   `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`
	LastName                 string                   `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
	ImageURL                 string                   `xml:"ImageURL,omitempty" json:"ImageURL,omitempty" yaml:"ImageURL,omitempty"`
	Bio                      string                   `xml:"Bio,omitempty" json:"Bio,omitempty" yaml:"Bio,omitempty"`
	IsMale                   bool                     `xml:"isMale" json:"isMale" yaml:"isMale"`
}

// StoredCardInfo was auto-generated from WSDL.
type StoredCardInfo struct {
	Name     string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount   float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	LastFour string  `xml:"LastFour,omitempty" json:"LastFour,omitempty" yaml:"LastFour,omitempty"`
}

// Tip was auto-generated from WSDL.
type Tip struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Amount    float64        `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
	StaffID   int64          `xml:"StaffID,omitempty" json:"StaffID,omitempty" yaml:"StaffID,omitempty"`
}

// TrackDataInfo was auto-generated from WSDL.
type TrackDataInfo struct {
	Name      string  `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Amount    float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	TrackData string  `xml:"TrackData,omitempty" json:"TrackData,omitempty" yaml:"TrackData,omitempty"`
}

// Unavailability was auto-generated from WSDL.
type Unavailability struct {
	Site          *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages      *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute       string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode     string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID            int            `xml:"ID" json:"ID" yaml:"ID"`
	StartDateTime DateTime       `xml:"StartDateTime" json:"StartDateTime" yaml:"StartDateTime"`
	EndDateTime   DateTime       `xml:"EndDateTime" json:"EndDateTime" yaml:"EndDateTime"`
	Description   string         `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

// UpdateProducts was auto-generated from WSDL.
type UpdateProducts struct {
	XMLName xml.Name               `xml:"http://clients.mindbodyonline.com/api/0_5 UpdateProducts" json:"-" yaml:"-"`
	Request *UpdateProductsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// UpdateProductsRequest was auto-generated from WSDL.
type UpdateProductsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	Products          *ArrayOfProduct    `xml:"Products,omitempty" json:"Products,omitempty" yaml:"Products,omitempty"`
	Test              bool               `xml:"Test,omitempty" json:"Test,omitempty" yaml:"Test,omitempty"`
}

// UpdateProductsResponse was auto-generated from WSDL.
type UpdateProductsResponse struct {
	UpdateProductsResult *UpdateProductsResult `xml:"UpdateProductsResult,omitempty" json:"UpdateProductsResult,omitempty" yaml:"UpdateProductsResult,omitempty"`
}

// UpdateProductsResult was auto-generated from WSDL.
type UpdateProductsResult struct {
	Status           StatusCode      `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int             `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string          `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel  `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int             `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int             `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int             `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Products         *ArrayOfProduct `xml:"Products,omitempty" json:"Products,omitempty" yaml:"Products,omitempty"`
}

// UpdateSaleDate was auto-generated from WSDL.
type UpdateSaleDate struct {
	XMLName xml.Name               `xml:"http://clients.mindbodyonline.com/api/0_5 UpdateSaleDate" json:"-" yaml:"-"`
	Request *UpdateSaleDateRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// UpdateSaleDateRequest was auto-generated from WSDL.
type UpdateSaleDateRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	SaleID            int64              `xml:"SaleID,omitempty" json:"SaleID,omitempty" yaml:"SaleID,omitempty"`
	SaleDate          DateTime           `xml:"SaleDate,omitempty" json:"SaleDate,omitempty" yaml:"SaleDate,omitempty"`
}

// UpdateSaleDateResponse was auto-generated from WSDL.
type UpdateSaleDateResponse struct {
	UpdateSaleDateResult *UpdateSaleDateResult `xml:"UpdateSaleDateResult,omitempty" json:"UpdateSaleDateResult,omitempty" yaml:"UpdateSaleDateResult,omitempty"`
}

// UpdateSaleDateResult was auto-generated from WSDL.
type UpdateSaleDateResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Sale             *Sale          `xml:"Sale,omitempty" json:"Sale,omitempty" yaml:"Sale,omitempty"`
}

// UpdateServices was auto-generated from WSDL.
type UpdateServices struct {
	XMLName xml.Name               `xml:"http://clients.mindbodyonline.com/api/0_5 UpdateServices" json:"-" yaml:"-"`
	Request *UpdateServicesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// UpdateServicesRequest was auto-generated from WSDL.
type UpdateServicesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	Services          *ArrayOfService    `xml:"Services,omitempty" json:"Services,omitempty" yaml:"Services,omitempty"`
	Test              bool               `xml:"Test,omitempty" json:"Test,omitempty" yaml:"Test,omitempty"`
}

// UpdateServicesResponse was auto-generated from WSDL.
type UpdateServicesResponse struct {
	UpdateServicesResult *UpdateServicesResult `xml:"UpdateServicesResult,omitempty" json:"UpdateServicesResult,omitempty" yaml:"UpdateServicesResult,omitempty"`
}

// UpdateServicesResult was auto-generated from WSDL.
type UpdateServicesResult struct {
	Status           StatusCode      `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int             `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string          `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel  `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int             `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int             `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int             `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Services         *ArrayOfService `xml:"Services,omitempty" json:"Services,omitempty" yaml:"Services,omitempty"`
}

// UserCredentials was auto-generated from WSDL.
type UserCredentials struct {
	Username   string      `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password   string      `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	SiteIDs    *ArrayOfInt `xml:"SiteIDs,omitempty" json:"SiteIDs,omitempty" yaml:"SiteIDs,omitempty"`
	LocationID int         `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// Visit was auto-generated from WSDL.
type Visit struct {
	Site                        *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages                    *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute                     string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode                   string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID                          int64          `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	ClassID                     int            `xml:"ClassID,omitempty" json:"ClassID,omitempty" yaml:"ClassID,omitempty"`
	AppointmentID               int            `xml:"AppointmentID,omitempty" json:"AppointmentID,omitempty" yaml:"AppointmentID,omitempty"`
	AppointmentGenderPreference string         `xml:"AppointmentGenderPreference,omitempty" json:"AppointmentGenderPreference,omitempty" yaml:"AppointmentGenderPreference,omitempty"`
	StartDateTime               DateTime       `xml:"StartDateTime,omitempty" json:"StartDateTime,omitempty" yaml:"StartDateTime,omitempty"`
	LateCancelled               bool           `xml:"LateCancelled,omitempty" json:"LateCancelled,omitempty" yaml:"LateCancelled,omitempty"`
	EndDateTime                 DateTime       `xml:"EndDateTime,omitempty" json:"EndDateTime,omitempty" yaml:"EndDateTime,omitempty"`
	Name                        string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Staff                       *Staff         `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
	Location                    *Location      `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Client                      *Client        `xml:"Client,omitempty" json:"Client,omitempty" yaml:"Client,omitempty"`
	WebSignup                   bool           `xml:"WebSignup,omitempty" json:"WebSignup,omitempty" yaml:"WebSignup,omitempty"`
	Action                      ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	SignedIn                    bool           `xml:"SignedIn,omitempty" json:"SignedIn,omitempty" yaml:"SignedIn,omitempty"`
	AppointmentStatus           string         `xml:"AppointmentStatus,omitempty" json:"AppointmentStatus,omitempty" yaml:"AppointmentStatus,omitempty"`
	MakeUp                      bool           `xml:"MakeUp,omitempty" json:"MakeUp,omitempty" yaml:"MakeUp,omitempty"`
	Service                     *ClientService `xml:"Service,omitempty" json:"Service,omitempty" yaml:"Service,omitempty"`
}

// sale_x0020_ServiceSoap implements the Sale_x0020_ServiceSoap interface.
type sale_x0020_ServiceSoap struct {
	cli *soap.Client
}

// Validates and completes a sale by processing all items added
// to a shopping cart.
func (p *sale_x0020_ServiceSoap) CheckoutShoppingCart(α *CheckoutShoppingCart) (β *CheckoutShoppingCartResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M CheckoutShoppingCartResponse `xml:"CheckoutShoppingCartResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of card types that the site accepts.
func (p *sale_x0020_ServiceSoap) GetAcceptedCardType(α *GetAcceptedCardType) (β *GetAcceptedCardTypeResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetAcceptedCardTypeResponse `xml:"GetAcceptedCardTypeResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of Custom Payment Methods.
func (p *sale_x0020_ServiceSoap) GetCustomPaymentMethods(α *GetCustomPaymentMethods) (β *GetCustomPaymentMethodsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetCustomPaymentMethodsResponse `xml:"GetCustomPaymentMethodsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of packages available for sale.
func (p *sale_x0020_ServiceSoap) GetPackages(α *GetPackages) (β *GetPackagesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetPackagesResponse `xml:"GetPackagesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Get a list of products available for sale.
func (p *sale_x0020_ServiceSoap) GetProducts(α *GetProducts) (β *GetProductsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetProductsResponse `xml:"GetProductsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of sales.
func (p *sale_x0020_ServiceSoap) GetSales(α *GetSales) (β *GetSalesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetSalesResponse `xml:"GetSalesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of services available for sale.
func (p *sale_x0020_ServiceSoap) GetServices(α *GetServices) (β *GetServicesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetServicesResponse `xml:"GetServicesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Redeem a Spa Finder Gift Card.
func (p *sale_x0020_ServiceSoap) RedeemSpaFinderWellnessCard(α *RedeemSpaFinderWellnessCard) (β *RedeemSpaFinderWellnessCardResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M RedeemSpaFinderWellnessCardResponse `xml:"RedeemSpaFinderWellnessCardResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Return a sale used in business mode. This only supports comp
// payment method.
func (p *sale_x0020_ServiceSoap) ReturnSale(α *ReturnSale) (β *ReturnSaleResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ReturnSaleResponse `xml:"ReturnSaleResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Update select products information.
func (p *sale_x0020_ServiceSoap) UpdateProducts(α *UpdateProducts) (β *UpdateProductsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M UpdateProductsResponse `xml:"UpdateProductsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Modify sale date in business mode
func (p *sale_x0020_ServiceSoap) UpdateSaleDate(α *UpdateSaleDate) (β *UpdateSaleDateResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M UpdateSaleDateResponse `xml:"UpdateSaleDateResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Update select services information.
func (p *sale_x0020_ServiceSoap) UpdateServices(α *UpdateServices) (β *UpdateServicesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M UpdateServicesResponse `xml:"UpdateServicesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
