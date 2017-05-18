package staffservice

import (
	"encoding/xml"
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://clients.mindbodyonline.com/api/0_5"

// NewStaff_x0020_ServiceSoap creates an initializes a Staff_x0020_ServiceSoap.
func NewStaff_x0020_ServiceSoap(cli *soap.Client) Staff_x0020_ServiceSoap {
	return &staff_x0020_ServiceSoap{cli}
}

// Staff_x0020_ServiceSoap was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Staff_x0020_ServiceSoap interface {
	// Add or update staff.
	AddOrUpdateStaff(α *AddOrUpdateStaff) (β *AddOrUpdateStaffResponse, err error)

	// Gets a list of sales reps based on the requested rep IDs
	GetSalesReps(α *GetSalesReps) (β *GetSalesRepsResponse, err error)

	// Gets a list of staff members.
	GetStaff(α *GetStaff) (β *GetStaffResponse, err error)

	// Gets a staff member's image URL if it exists.
	GetStaffImgURL(α *GetStaffImgURL) (β *GetStaffImgURLResponse, err error)

	// Gets a list of staff permissions based on the given staff member.
	GetStaffPermissions(α *GetStaffPermissions) (β *GetStaffPermissionsResponse, err error)

	// Validates a username and password. This method returns the staff
	// on success.
	ValidateStaffLogin(α *ValidateStaffLogin) (β *ValidateStaffLoginResponse, err error)
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

// StaffFilter was auto-generated from WSDL.
type StaffFilter string

// Validate validates StaffFilter.
func (v StaffFilter) Validate() bool {
	for _, vv := range []string{
		"StaffViewable",
		"AppointmentInstructor",
		"ClassInstructor",
		"Male",
		"Female",
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

// AddOrUpdateStaff was auto-generated from WSDL.
type AddOrUpdateStaff struct {
	XMLName xml.Name                 `xml:"http://clients.mindbodyonline.com/api/0_5 AddOrUpdateStaff" json:"-" yaml:"-"`
	Request *AddOrUpdateStaffRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// AddOrUpdateStaffRequest was auto-generated from WSDL.
type AddOrUpdateStaffRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	UpdateAction      string             `xml:"UpdateAction,omitempty" json:"UpdateAction,omitempty" yaml:"UpdateAction,omitempty"`
	Test              bool               `xml:"Test,omitempty" json:"Test,omitempty" yaml:"Test,omitempty"`
	Staff             *ArrayOfStaff      `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
}

// AddOrUpdateStaffResponse was auto-generated from WSDL.
type AddOrUpdateStaffResponse struct {
	AddOrUpdateStaffResult *AddOrUpdateStaffResult `xml:"AddOrUpdateStaffResult,omitempty" json:"AddOrUpdateStaffResult,omitempty" yaml:"AddOrUpdateStaffResult,omitempty"`
}

// AddOrUpdateStaffResult was auto-generated from WSDL.
type AddOrUpdateStaffResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Staff            *ArrayOfStaff  `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
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

// ArrayOfPermission was auto-generated from WSDL.
type ArrayOfPermission struct {
	Permission []*Permission `xml:"Permission,omitempty" json:"Permission,omitempty" yaml:"Permission,omitempty"`
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

// ArrayOfSalesRep was auto-generated from WSDL.
type ArrayOfSalesRep struct {
	SalesRep []*SalesRep `xml:"SalesRep,omitempty" json:"SalesRep,omitempty" yaml:"SalesRep,omitempty"`
}

// ArrayOfStaff was auto-generated from WSDL.
type ArrayOfStaff struct {
	Staff []*Staff `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
}

// ArrayOfStaffFilter was auto-generated from WSDL.
type ArrayOfStaffFilter struct {
	StaffFilter []StaffFilter `xml:"StaffFilter,omitempty" json:"StaffFilter,omitempty" yaml:"StaffFilter,omitempty"`
}

// ArrayOfString was auto-generated from WSDL.
type ArrayOfString struct {
	String []string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// ArrayOfUnavailability was auto-generated from WSDL.
type ArrayOfUnavailability struct {
	Unavailability []*Unavailability `xml:"Unavailability,omitempty" json:"Unavailability,omitempty" yaml:"Unavailability,omitempty"`
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

// CustomClientField was auto-generated from WSDL.
type CustomClientField struct {
	ID       int    `xml:"ID" json:"ID" yaml:"ID"`
	DataType string `xml:"DataType,omitempty" json:"DataType,omitempty" yaml:"DataType,omitempty"`
	Name     string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Value    string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// GetSalesReps was auto-generated from WSDL.
type GetSalesReps struct {
	XMLName xml.Name             `xml:"http://clients.mindbodyonline.com/api/0_5 GetSalesReps" json:"-" yaml:"-"`
	Request *GetSalesRepsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetSalesRepsRequest was auto-generated from WSDL.
type GetSalesRepsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	SalesRepNumbers   *ArrayOfInt        `xml:"SalesRepNumbers,omitempty" json:"SalesRepNumbers,omitempty" yaml:"SalesRepNumbers,omitempty"`
	ShowActiveOnly    bool               `xml:"ShowActiveOnly,omitempty" json:"ShowActiveOnly,omitempty" yaml:"ShowActiveOnly,omitempty"`
}

// GetSalesRepsResponse was auto-generated from WSDL.
type GetSalesRepsResponse struct {
	GetSalesRepsResult *GetSalesRepsResult `xml:"GetSalesRepsResult,omitempty" json:"GetSalesRepsResult,omitempty" yaml:"GetSalesRepsResult,omitempty"`
}

// GetSalesRepsResult was auto-generated from WSDL.
type GetSalesRepsResult struct {
	Status           StatusCode       `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int              `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string           `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel   `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int              `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int              `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int              `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	SalesReps        *ArrayOfSalesRep `xml:"SalesReps,omitempty" json:"SalesReps,omitempty" yaml:"SalesReps,omitempty"`
}

// GetStaff was auto-generated from WSDL.
type GetStaff struct {
	XMLName xml.Name         `xml:"http://clients.mindbodyonline.com/api/0_5 GetStaff" json:"-" yaml:"-"`
	Request *GetStaffRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetStaffImgURL was auto-generated from WSDL.
type GetStaffImgURL struct {
	XMLName xml.Name               `xml:"http://clients.mindbodyonline.com/api/0_5 GetStaffImgURL" json:"-" yaml:"-"`
	Request *GetStaffImgURLRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetStaffImgURLRequest was auto-generated from WSDL.
type GetStaffImgURLRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	StaffID           int64              `xml:"StaffID" json:"StaffID" yaml:"StaffID"`
}

// GetStaffImgURLResponse was auto-generated from WSDL.
type GetStaffImgURLResponse struct {
	GetStaffImgURLResult *GetStaffImgURLResult `xml:"GetStaffImgURLResult,omitempty" json:"GetStaffImgURLResult,omitempty" yaml:"GetStaffImgURLResult,omitempty"`
}

// GetStaffImgURLResult was auto-generated from WSDL.
type GetStaffImgURLResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ImageURL         string         `xml:"ImageURL,omitempty" json:"ImageURL,omitempty" yaml:"ImageURL,omitempty"`
	MobileImageURL   string         `xml:"MobileImageURL,omitempty" json:"MobileImageURL,omitempty" yaml:"MobileImageURL,omitempty"`
}

// GetStaffPermissions was auto-generated from WSDL.
type GetStaffPermissions struct {
	XMLName xml.Name                    `xml:"http://clients.mindbodyonline.com/api/0_5 GetStaffPermissions" json:"-" yaml:"-"`
	Request *GetStaffPermissionsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetStaffPermissionsRequest was auto-generated from WSDL.
type GetStaffPermissionsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	StaffID           int64              `xml:"StaffID" json:"StaffID" yaml:"StaffID"`
}

// GetStaffPermissionsResponse was auto-generated from WSDL.
type GetStaffPermissionsResponse struct {
	GetStaffPermissionsResult *GetStaffPermissionsResult `xml:"GetStaffPermissionsResult,omitempty" json:"GetStaffPermissionsResult,omitempty" yaml:"GetStaffPermissionsResult,omitempty"`
}

// GetStaffPermissionsResult was auto-generated from WSDL.
type GetStaffPermissionsResult struct {
	Status           StatusCode         `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string             `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel     `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Permissions      *ArrayOfPermission `xml:"Permissions,omitempty" json:"Permissions,omitempty" yaml:"Permissions,omitempty"`
}

// GetStaffRequest was auto-generated from WSDL.
type GetStaffRequest struct {
	SourceCredentials *SourceCredentials  `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials    `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel      `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                 `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                 `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString      `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	StaffIDs          *ArrayOfLong        `xml:"StaffIDs,omitempty" json:"StaffIDs,omitempty" yaml:"StaffIDs,omitempty"`
	StaffCredentials  *StaffCredentials   `xml:"StaffCredentials,omitempty" json:"StaffCredentials,omitempty" yaml:"StaffCredentials,omitempty"`
	Filters           *ArrayOfStaffFilter `xml:"Filters,omitempty" json:"Filters,omitempty" yaml:"Filters,omitempty"`
	SessionTypeID     int                 `xml:"SessionTypeID,omitempty" json:"SessionTypeID,omitempty" yaml:"SessionTypeID,omitempty"`
	StartDateTime     DateTime            `xml:"StartDateTime,omitempty" json:"StartDateTime,omitempty" yaml:"StartDateTime,omitempty"`
	LocationID        int                 `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// GetStaffResponse was auto-generated from WSDL.
type GetStaffResponse struct {
	GetStaffResult *GetStaffResult `xml:"GetStaffResult,omitempty" json:"GetStaffResult,omitempty" yaml:"GetStaffResult,omitempty"`
}

// GetStaffResult was auto-generated from WSDL.
type GetStaffResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	StaffMembers     *ArrayOfStaff  `xml:"StaffMembers,omitempty" json:"StaffMembers,omitempty" yaml:"StaffMembers,omitempty"`
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

// Permission was auto-generated from WSDL.
type Permission struct {
	DisplayName string `xml:"DisplayName,omitempty" json:"DisplayName,omitempty" yaml:"DisplayName,omitempty"`
	Name        string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Value       string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
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

// StaffCredentials was auto-generated from WSDL.
type StaffCredentials struct {
	Username string      `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password string      `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	SiteIDs  *ArrayOfInt `xml:"SiteIDs,omitempty" json:"SiteIDs,omitempty" yaml:"SiteIDs,omitempty"`
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

// UserCredentials was auto-generated from WSDL.
type UserCredentials struct {
	Username   string      `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password   string      `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	SiteIDs    *ArrayOfInt `xml:"SiteIDs,omitempty" json:"SiteIDs,omitempty" yaml:"SiteIDs,omitempty"`
	LocationID int         `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// ValidateLoginRequest was auto-generated from WSDL.
type ValidateLoginRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	Username          string             `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password          string             `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
}

// ValidateLoginResult was auto-generated from WSDL.
type ValidateLoginResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	GUID             string         `xml:"GUID,omitempty" json:"GUID,omitempty" yaml:"GUID,omitempty"`
	Client           *Client        `xml:"Client,omitempty" json:"Client,omitempty" yaml:"Client,omitempty"`
	Staff            *Staff         `xml:"Staff,omitempty" json:"Staff,omitempty" yaml:"Staff,omitempty"`
}

// ValidateStaffLogin was auto-generated from WSDL.
type ValidateStaffLogin struct {
	XMLName xml.Name              `xml:"http://clients.mindbodyonline.com/api/0_5 ValidateStaffLogin" json:"-" yaml:"-"`
	Request *ValidateLoginRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// ValidateStaffLoginResponse was auto-generated from WSDL.
type ValidateStaffLoginResponse struct {
	ValidateStaffLoginResult *ValidateLoginResult `xml:"ValidateStaffLoginResult,omitempty" json:"ValidateStaffLoginResult,omitempty" yaml:"ValidateStaffLoginResult,omitempty"`
}

// staff_x0020_ServiceSoap implements the Staff_x0020_ServiceSoap interface.
type staff_x0020_ServiceSoap struct {
	cli *soap.Client
}

// Add or update staff.
func (p *staff_x0020_ServiceSoap) AddOrUpdateStaff(α *AddOrUpdateStaff) (β *AddOrUpdateStaffResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M AddOrUpdateStaffResponse `xml:"AddOrUpdateStaffResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of sales reps based on the requested rep IDs
func (p *staff_x0020_ServiceSoap) GetSalesReps(α *GetSalesReps) (β *GetSalesRepsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetSalesRepsResponse `xml:"GetSalesRepsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of staff members.
func (p *staff_x0020_ServiceSoap) GetStaff(α *GetStaff) (β *GetStaffResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetStaffResponse `xml:"GetStaffResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a staff member's image URL if it exists.
func (p *staff_x0020_ServiceSoap) GetStaffImgURL(α *GetStaffImgURL) (β *GetStaffImgURLResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetStaffImgURLResponse `xml:"GetStaffImgURLResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of staff permissions based on the given staff member.
func (p *staff_x0020_ServiceSoap) GetStaffPermissions(α *GetStaffPermissions) (β *GetStaffPermissionsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetStaffPermissionsResponse `xml:"GetStaffPermissionsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Validates a username and password. This method returns the staff
// on success.
func (p *staff_x0020_ServiceSoap) ValidateStaffLogin(α *ValidateStaffLogin) (β *ValidateStaffLoginResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ValidateStaffLoginResponse `xml:"ValidateStaffLoginResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
