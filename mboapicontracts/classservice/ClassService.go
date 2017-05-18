package mboapicontracts

import (
	"encoding/xml"
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://clients.mindbodyonline.com/api/0_5"

// NewClient_x0020_ServiceSoap creates an initializes a Client_x0020_ServiceSoap.
func NewClient_x0020_ServiceSoap(cli *soap.Client) Client_x0020_ServiceSoap {
	return &client_x0020_ServiceSoap{cli}
}

// Client_x0020_ServiceSoap was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Client_x0020_ServiceSoap interface {
	// Adds an arrival record for the given client.
	AddArrival(α *AddArrival) (β *AddArrivalResponse, err error)

	// Adds a formula note to a client.
	AddClientFormulaNote(α *AddClientFormulaNote) (β *AddClientFormulaNoteResponse, err error)

	// Adds or updates information for a list of clients.
	AddOrUpdateClients(α *AddOrUpdateClients) (β *AddOrUpdateClientsResponse, err error)

	// Add or update client contact logs.
	AddOrUpdateContactLogs(α *AddOrUpdateContactLogs) (β *AddOrUpdateContactLogsResponse, err error)

	// Deletes a formula note to a client.
	DeleteClientFormulaNote(α *DeleteClientFormulaNote) (β *DeleteClientFormulaNoteResponse, err error)

	// Gets the active membership for a given client.
	GetActiveClientMemberships(α *GetActiveClientMemberships) (β *GetActiveClientMembershipsResponse, err error)

	// Gets account balances for the given clients.
	GetClientAccountBalances(α *GetClientAccountBalances) (β *GetClientAccountBalancesResponse, err error)

	// Get contact logs for a client.
	GetClientContactLogs(α *GetClientContactLogs) (β *GetClientContactLogsResponse, err error)

	// Gets a list of contracts for a given client.
	GetClientContracts(α *GetClientContracts) (β *GetClientContractsResponse, err error)

	// Gets a list of client formula notes.
	GetClientFormulaNotes(α *GetClientFormulaNotes) (β *GetClientFormulaNotesResponse, err error)

	// Gets a list of currently available client indexes.
	GetClientIndexes(α *GetClientIndexes) (β *GetClientIndexesResponse, err error)

	// Get purchases for a client.
	GetClientPurchases(α *GetClientPurchases) (β *GetClientPurchasesResponse, err error)

	// Gets a list of clients.
	GetClientReferralTypes(α *GetClientReferralTypes) (β *GetClientReferralTypesResponse, err error)

	// Get visits for a client.
	GetClientSchedule(α *GetClientSchedule) (β *GetClientScheduleResponse, err error)

	// Gets a client service for a given client.
	GetClientServices(α *GetClientServices) (β *GetClientServicesResponse, err error)

	// Get visits for a client.
	GetClientVisits(α *GetClientVisits) (β *GetClientVisitsResponse, err error)

	// Gets a list of clients.
	GetClients(α *GetClients) (β *GetClientsResponse, err error)

	// Get contact log types for a client.
	GetContactLogTypes(α *GetContactLogTypes) (β *GetContactLogTypesResponse, err error)

	// Gets a list of currently available client indexes.
	GetCustomClientFields(α *GetCustomClientFields) (β *GetCustomClientFieldsResponse, err error)

	// Updates a client service for a given client.
	GetRequiredClientFields(α *GetRequiredClientFields) (β *GetRequiredClientFieldsResponse, err error)

	// Sends the user a new password.
	SendUserNewPassword(α *SendUserNewPassword) (β *SendUserNewPasswordResponse, err error)

	// Updates a client service for a given client.
	UpdateClientServices(α *UpdateClientServices) (β *UpdateClientServicesResponse, err error)

	// Upload a client document.
	UploadClientDocument(α *UploadClientDocument) (β *UploadClientDocumentResponse, err error)

	// Upload a client photo.
	UploadClientPhoto(α *UploadClientPhoto) (β *UploadClientPhotoResponse, err error)

	// Validates a username and password. This method returns the associated
	// clients record and a session GUID on success.
	ValidateLogin(α *ValidateLogin) (β *ValidateLoginResponse, err error)
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

// AddArrival was auto-generated from WSDL.
type AddArrival struct {
	XMLName xml.Name           `xml:"http://clients.mindbodyonline.com/api/0_5 AddArrival" json:"-" yaml:"-"`
	Request *AddArrivalRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// AddArrivalRequest was auto-generated from WSDL.
type AddArrivalRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	LocationID        int                `xml:"LocationID" json:"LocationID" yaml:"LocationID"`
}

// AddArrivalResponse was auto-generated from WSDL.
type AddArrivalResponse struct {
	AddArrivalResult *AddArrivalResult `xml:"AddArrivalResult,omitempty" json:"AddArrivalResult,omitempty" yaml:"AddArrivalResult,omitempty"`
}

// AddArrivalResult was auto-generated from WSDL.
type AddArrivalResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ArrivalAdded     bool           `xml:"ArrivalAdded" json:"ArrivalAdded" yaml:"ArrivalAdded"`
	ClientService    *ClientService `xml:"ClientService,omitempty" json:"ClientService,omitempty" yaml:"ClientService,omitempty"`
}

// AddClientFormulaNote was auto-generated from WSDL.
type AddClientFormulaNote struct {
	XMLName xml.Name                     `xml:"http://clients.mindbodyonline.com/api/0_5 AddClientFormulaNote" json:"-" yaml:"-"`
	Request *AddClientFormulaNoteRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// AddClientFormulaNoteRequest was auto-generated from WSDL.
type AddClientFormulaNoteRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	AppointmentID     int64              `xml:"AppointmentID,omitempty" json:"AppointmentID,omitempty" yaml:"AppointmentID,omitempty"`
	Note              string             `xml:"Note,omitempty" json:"Note,omitempty" yaml:"Note,omitempty"`
}

// AddClientFormulaNoteResponse was auto-generated from WSDL.
type AddClientFormulaNoteResponse struct {
	AddClientFormulaNoteResult *AddClientFormulaNoteResult `xml:"AddClientFormulaNoteResult,omitempty" json:"AddClientFormulaNoteResult,omitempty" yaml:"AddClientFormulaNoteResult,omitempty"`
}

// AddClientFormulaNoteResult was auto-generated from WSDL.
type AddClientFormulaNoteResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	FormulaNote      *FormulaNote   `xml:"FormulaNote,omitempty" json:"FormulaNote,omitempty" yaml:"FormulaNote,omitempty"`
}

// AddOrUpdateClients was auto-generated from WSDL.
type AddOrUpdateClients struct {
	XMLName xml.Name                   `xml:"http://clients.mindbodyonline.com/api/0_5 AddOrUpdateClients" json:"-" yaml:"-"`
	Request *AddOrUpdateClientsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// AddOrUpdateClientsRequest was auto-generated from WSDL.
type AddOrUpdateClientsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	UpdateAction      string             `xml:"UpdateAction,omitempty" json:"UpdateAction,omitempty" yaml:"UpdateAction,omitempty"`
	Test              bool               `xml:"Test,omitempty" json:"Test,omitempty" yaml:"Test,omitempty"`
	Clients           *ArrayOfClient     `xml:"Clients,omitempty" json:"Clients,omitempty" yaml:"Clients,omitempty"`
	SendEmail         bool               `xml:"SendEmail,omitempty" json:"SendEmail,omitempty" yaml:"SendEmail,omitempty"`
}

// AddOrUpdateClientsResponse was auto-generated from WSDL.
type AddOrUpdateClientsResponse struct {
	AddOrUpdateClientsResult *AddOrUpdateClientsResult `xml:"AddOrUpdateClientsResult,omitempty" json:"AddOrUpdateClientsResult,omitempty" yaml:"AddOrUpdateClientsResult,omitempty"`
}

// AddOrUpdateClientsResult was auto-generated from WSDL.
type AddOrUpdateClientsResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Clients          *ArrayOfClient `xml:"Clients,omitempty" json:"Clients,omitempty" yaml:"Clients,omitempty"`
}

// AddOrUpdateContactLogs was auto-generated from WSDL.
type AddOrUpdateContactLogs struct {
	XMLName xml.Name                       `xml:"http://clients.mindbodyonline.com/api/0_5 AddOrUpdateContactLogs" json:"-" yaml:"-"`
	Request *AddOrUpdateContactLogsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// AddOrUpdateContactLogsRequest was auto-generated from WSDL.
type AddOrUpdateContactLogsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	UpdateAction      string             `xml:"UpdateAction,omitempty" json:"UpdateAction,omitempty" yaml:"UpdateAction,omitempty"`
	Test              bool               `xml:"Test,omitempty" json:"Test,omitempty" yaml:"Test,omitempty"`
	ContactLogs       *ArrayOfContactLog `xml:"ContactLogs,omitempty" json:"ContactLogs,omitempty" yaml:"ContactLogs,omitempty"`
}

// AddOrUpdateContactLogsResponse was auto-generated from WSDL.
type AddOrUpdateContactLogsResponse struct {
	AddOrUpdateContactLogsResult *AddOrUpdateContactLogsResult `xml:"AddOrUpdateContactLogsResult,omitempty" json:"AddOrUpdateContactLogsResult,omitempty" yaml:"AddOrUpdateContactLogsResult,omitempty"`
}

// AddOrUpdateContactLogsResult was auto-generated from WSDL.
type AddOrUpdateContactLogsResult struct {
	Status           StatusCode         `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string             `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel     `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ContactLogs      *ArrayOfContactLog `xml:"ContactLogs,omitempty" json:"ContactLogs,omitempty" yaml:"ContactLogs,omitempty"`
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

// ArrayOfClient was auto-generated from WSDL.
type ArrayOfClient struct {
	Client []*Client `xml:"Client,omitempty" json:"Client,omitempty" yaml:"Client,omitempty"`
}

// ArrayOfClientContract was auto-generated from WSDL.
type ArrayOfClientContract struct {
	ClientContract []*ClientContract `xml:"ClientContract,omitempty" json:"ClientContract,omitempty" yaml:"ClientContract,omitempty"`
}

// ArrayOfClientIndex was auto-generated from WSDL.
type ArrayOfClientIndex struct {
	ClientIndex []*ClientIndex `xml:"ClientIndex,omitempty" json:"ClientIndex,omitempty" yaml:"ClientIndex,omitempty"`
}

// ArrayOfClientIndexValue was auto-generated from WSDL.
type ArrayOfClientIndexValue struct {
	ClientIndexValue []*ClientIndexValue `xml:"ClientIndexValue,omitempty" json:"ClientIndexValue,omitempty" yaml:"ClientIndexValue,omitempty"`
}

// ArrayOfClientMembership was auto-generated from WSDL.
type ArrayOfClientMembership struct {
	ClientMembership []*ClientMembership `xml:"ClientMembership,omitempty" json:"ClientMembership,omitempty" yaml:"ClientMembership,omitempty"`
}

// ArrayOfClientRelationship was auto-generated from WSDL.
type ArrayOfClientRelationship struct {
	ClientRelationship []*ClientRelationship `xml:"ClientRelationship,omitempty" json:"ClientRelationship,omitempty" yaml:"ClientRelationship,omitempty"`
}

// ArrayOfClientService was auto-generated from WSDL.
type ArrayOfClientService struct {
	ClientService []*ClientService `xml:"ClientService,omitempty" json:"ClientService,omitempty" yaml:"ClientService,omitempty"`
}

// ArrayOfContactLog was auto-generated from WSDL.
type ArrayOfContactLog struct {
	ContactLog []*ContactLog `xml:"ContactLog,omitempty" json:"ContactLog,omitempty" yaml:"ContactLog,omitempty"`
}

// ArrayOfContactLogComment was auto-generated from WSDL.
type ArrayOfContactLogComment struct {
	ContactLogComment []*ContactLogComment `xml:"ContactLogComment,omitempty" json:"ContactLogComment,omitempty" yaml:"ContactLogComment,omitempty"`
}

// ArrayOfContactLogSubtype was auto-generated from WSDL.
type ArrayOfContactLogSubtype struct {
	ContactLogSubtype []*ContactLogSubtype `xml:"ContactLogSubtype,omitempty" json:"ContactLogSubtype,omitempty" yaml:"ContactLogSubtype,omitempty"`
}

// ArrayOfContactLogType was auto-generated from WSDL.
type ArrayOfContactLogType struct {
	ContactLogType []*ContactLogType `xml:"ContactLogType,omitempty" json:"ContactLogType,omitempty" yaml:"ContactLogType,omitempty"`
}

// ArrayOfCustomClientField was auto-generated from WSDL.
type ArrayOfCustomClientField struct {
	CustomClientField []*CustomClientField `xml:"CustomClientField,omitempty" json:"CustomClientField,omitempty" yaml:"CustomClientField,omitempty"`
}

// ArrayOfFormulaNote was auto-generated from WSDL.
type ArrayOfFormulaNote struct {
	FormulaNote []*FormulaNote `xml:"FormulaNote,omitempty" json:"FormulaNote,omitempty" yaml:"FormulaNote,omitempty"`
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

// ArrayOfPayment was auto-generated from WSDL.
type ArrayOfPayment struct {
	Payment []*Payment `xml:"Payment,omitempty" json:"Payment,omitempty" yaml:"Payment,omitempty"`
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

// ArrayOfSaleItem was auto-generated from WSDL.
type ArrayOfSaleItem struct {
	SaleItem []*SaleItem `xml:"SaleItem,omitempty" json:"SaleItem,omitempty" yaml:"SaleItem,omitempty"`
}

// ArrayOfSalesRep was auto-generated from WSDL.
type ArrayOfSalesRep struct {
	SalesRep []*SalesRep `xml:"SalesRep,omitempty" json:"SalesRep,omitempty" yaml:"SalesRep,omitempty"`
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

// ClientContract was auto-generated from WSDL.
type ClientContract struct {
	Site          *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages      *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute       string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode     string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	AgreementDate DateTime       `xml:"AgreementDate" json:"AgreementDate" yaml:"AgreementDate"`
	StartDate     DateTime       `xml:"StartDate" json:"StartDate" yaml:"StartDate"`
	EndDate       DateTime       `xml:"EndDate" json:"EndDate" yaml:"EndDate"`
	ContractName  string         `xml:"ContractName,omitempty" json:"ContractName,omitempty" yaml:"ContractName,omitempty"`
	Action        ActionCode     `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID            int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
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

// ClientMembership was auto-generated from WSDL.
type ClientMembership struct {
	Site                *Site            `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages            *ArrayOfString   `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute             string           `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode           string           `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Current             bool             `xml:"Current" json:"Current" yaml:"Current"`
	Count               int              `xml:"Count" json:"Count" yaml:"Count"`
	Remaining           int              `xml:"Remaining" json:"Remaining" yaml:"Remaining"`
	Action              ActionCode       `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
	ID                  int64            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name                string           `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	PaymentDate         DateTime         `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`
	ActiveDate          DateTime         `xml:"ActiveDate,omitempty" json:"ActiveDate,omitempty" yaml:"ActiveDate,omitempty"`
	ExpirationDate      DateTime         `xml:"ExpirationDate,omitempty" json:"ExpirationDate,omitempty" yaml:"ExpirationDate,omitempty"`
	Program             *Program         `xml:"Program,omitempty" json:"Program,omitempty" yaml:"Program,omitempty"`
	RestrictedLocations *ArrayOfLocation `xml:"RestrictedLocations,omitempty" json:"RestrictedLocations,omitempty" yaml:"RestrictedLocations,omitempty"`
	IconCode            string           `xml:"IconCode,omitempty" json:"IconCode,omitempty" yaml:"IconCode,omitempty"`
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

// ClientSendUserNewPasswordRequest was auto-generated from WSDL.
type ClientSendUserNewPasswordRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	UserEmail         string             `xml:"UserEmail,omitempty" json:"UserEmail,omitempty" yaml:"UserEmail,omitempty"`
	UserFirstName     string             `xml:"UserFirstName,omitempty" json:"UserFirstName,omitempty" yaml:"UserFirstName,omitempty"`
	UserLastName      string             `xml:"UserLastName,omitempty" json:"UserLastName,omitempty" yaml:"UserLastName,omitempty"`
}

// ClientSendUserNewPasswordResult was auto-generated from WSDL.
type ClientSendUserNewPasswordResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
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

// ContactLog was auto-generated from WSDL.
type ContactLog struct {
	Site              *Site                     `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages          *ArrayOfString            `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute           string                    `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode         string                    `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID                int64                     `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	CreatedBy         *Staff                    `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty" yaml:"CreatedBy,omitempty"`
	Client            *Client                   `xml:"Client,omitempty" json:"Client,omitempty" yaml:"Client,omitempty"`
	CreatedDateTime   DateTime                  `xml:"CreatedDateTime,omitempty" json:"CreatedDateTime,omitempty" yaml:"CreatedDateTime,omitempty"`
	FollowupByDate    DateTime                  `xml:"FollowupByDate,omitempty" json:"FollowupByDate,omitempty" yaml:"FollowupByDate,omitempty"`
	ContactName       string                    `xml:"ContactName,omitempty" json:"ContactName,omitempty" yaml:"ContactName,omitempty"`
	Text              string                    `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
	AssignedTo        *Staff                    `xml:"AssignedTo,omitempty" json:"AssignedTo,omitempty" yaml:"AssignedTo,omitempty"`
	ContactMethod     string                    `xml:"ContactMethod,omitempty" json:"ContactMethod,omitempty" yaml:"ContactMethod,omitempty"`
	IsSystemGenerated bool                      `xml:"IsSystemGenerated,omitempty" json:"IsSystemGenerated,omitempty" yaml:"IsSystemGenerated,omitempty"`
	Comments          *ArrayOfContactLogComment `xml:"Comments,omitempty" json:"Comments,omitempty" yaml:"Comments,omitempty"`
	Types             *ArrayOfContactLogType    `xml:"Types,omitempty" json:"Types,omitempty" yaml:"Types,omitempty"`
	Action            ActionCode                `xml:"Action,omitempty" json:"Action,omitempty" yaml:"Action,omitempty"`
}

// ContactLogComment was auto-generated from WSDL.
type ContactLogComment struct {
	Site            *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages        *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute         string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode       string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID              int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Text            string         `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`
	CreatedBy       *Staff         `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty" yaml:"CreatedBy,omitempty"`
	CreatedDateTime DateTime       `xml:"CreatedDateTime,omitempty" json:"CreatedDateTime,omitempty" yaml:"CreatedDateTime,omitempty"`
}

// ContactLogSubtype was auto-generated from WSDL.
type ContactLogSubtype struct {
	Site      *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID        int            `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name      string         `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

// ContactLogType was auto-generated from WSDL.
type ContactLogType struct {
	Site      *Site                     `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages  *ArrayOfString            `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute   string                    `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode string                    `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID        int                       `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	Name      string                    `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Subtypes  *ArrayOfContactLogSubtype `xml:"Subtypes,omitempty" json:"Subtypes,omitempty" yaml:"Subtypes,omitempty"`
}

// CustomClientField was auto-generated from WSDL.
type CustomClientField struct {
	ID       int    `xml:"ID" json:"ID" yaml:"ID"`
	DataType string `xml:"DataType,omitempty" json:"DataType,omitempty" yaml:"DataType,omitempty"`
	Name     string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	Value    string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// DeleteCientFormulaNoteRequest was auto-generated from WSDL.
type DeleteCientFormulaNoteRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	FormulaNoteID     int64              `xml:"FormulaNoteID" json:"FormulaNoteID" yaml:"FormulaNoteID"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
}

// DeleteClientFormulaNote was auto-generated from WSDL.
type DeleteClientFormulaNote struct {
	XMLName xml.Name                       `xml:"http://clients.mindbodyonline.com/api/0_5 DeleteClientFormulaNote" json:"-" yaml:"-"`
	Request *DeleteCientFormulaNoteRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// DeleteClientFormulaNoteResponse was auto-generated from WSDL.
type DeleteClientFormulaNoteResponse struct {
	DeleteClientFormulaNoteResult *DeleteClientFormulaNoteResult `xml:"DeleteClientFormulaNoteResult,omitempty" json:"DeleteClientFormulaNoteResult,omitempty" yaml:"DeleteClientFormulaNoteResult,omitempty"`
}

// DeleteClientFormulaNoteResult was auto-generated from WSDL.
type DeleteClientFormulaNoteResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	FormulaNote      *FormulaNote   `xml:"FormulaNote,omitempty" json:"FormulaNote,omitempty" yaml:"FormulaNote,omitempty"`
}

// FormulaNote was auto-generated from WSDL.
type FormulaNote struct {
	Site          *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages      *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute       string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode     string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ID            int64          `xml:"ID,omitempty" json:"ID,omitempty" yaml:"ID,omitempty"`
	ClientID      string         `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	Note          string         `xml:"Note,omitempty" json:"Note,omitempty" yaml:"Note,omitempty"`
	EntryDateTime DateTime       `xml:"EntryDateTime" json:"EntryDateTime" yaml:"EntryDateTime"`
	AppointmentID int64          `xml:"AppointmentID,omitempty" json:"AppointmentID,omitempty" yaml:"AppointmentID,omitempty"`
}

// GetActiveClientMemberships was auto-generated from WSDL.
type GetActiveClientMemberships struct {
	XMLName xml.Name                           `xml:"http://clients.mindbodyonline.com/api/0_5 GetActiveClientMemberships" json:"-" yaml:"-"`
	Request *GetActiveClientMembershipsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetActiveClientMembershipsRequest was auto-generated from WSDL.
type GetActiveClientMembershipsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	LocationID        int                `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// GetActiveClientMembershipsResponse was auto-generated from WSDL.
type GetActiveClientMembershipsResponse struct {
	GetActiveClientMembershipsResult *GetActiveClientMembershipsResult `xml:"GetActiveClientMembershipsResult,omitempty" json:"GetActiveClientMembershipsResult,omitempty" yaml:"GetActiveClientMembershipsResult,omitempty"`
}

// GetActiveClientMembershipsResult was auto-generated from WSDL.
type GetActiveClientMembershipsResult struct {
	Status            StatusCode               `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode         int                      `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message           string                   `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail         XMLDetailLevel           `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount       int                      `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex  int                      `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount    int                      `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ClientMemberships *ArrayOfClientMembership `xml:"ClientMemberships,omitempty" json:"ClientMemberships,omitempty" yaml:"ClientMemberships,omitempty"`
}

// GetClientAccountBalances was auto-generated from WSDL.
type GetClientAccountBalances struct {
	XMLName xml.Name                         `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientAccountBalances" json:"-" yaml:"-"`
	Request *GetClientAccountBalancesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientAccountBalancesRequest was auto-generated from WSDL.
type GetClientAccountBalancesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientIDs         *ArrayOfString     `xml:"ClientIDs,omitempty" json:"ClientIDs,omitempty" yaml:"ClientIDs,omitempty"`
	BalanceDate       DateTime           `xml:"BalanceDate,omitempty" json:"BalanceDate,omitempty" yaml:"BalanceDate,omitempty"`
	ClassID           int64              `xml:"ClassID,omitempty" json:"ClassID,omitempty" yaml:"ClassID,omitempty"`
}

// GetClientAccountBalancesResponse was auto-generated from WSDL.
type GetClientAccountBalancesResponse struct {
	GetClientAccountBalancesResult *GetClientAccountBalancesResult `xml:"GetClientAccountBalancesResult,omitempty" json:"GetClientAccountBalancesResult,omitempty" yaml:"GetClientAccountBalancesResult,omitempty"`
}

// GetClientAccountBalancesResult was auto-generated from WSDL.
type GetClientAccountBalancesResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Clients          *ArrayOfClient `xml:"Clients,omitempty" json:"Clients,omitempty" yaml:"Clients,omitempty"`
}

// GetClientContactLogs was auto-generated from WSDL.
type GetClientContactLogs struct {
	XMLName xml.Name                     `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientContactLogs" json:"-" yaml:"-"`
	Request *GetClientContactLogsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientContactLogsRequest was auto-generated from WSDL.
type GetClientContactLogsRequest struct {
	SourceCredentials   *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials     *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail           XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize            int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex    int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields              *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID            string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	StartDate           DateTime           `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate             DateTime           `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
	StaffIDs            *ArrayOfLong       `xml:"StaffIDs,omitempty" json:"StaffIDs,omitempty" yaml:"StaffIDs,omitempty"`
	ShowSystemGenerated bool               `xml:"ShowSystemGenerated,omitempty" json:"ShowSystemGenerated,omitempty" yaml:"ShowSystemGenerated,omitempty"`
	TypeIDs             *ArrayOfInt        `xml:"TypeIDs,omitempty" json:"TypeIDs,omitempty" yaml:"TypeIDs,omitempty"`
	SubtypeIDs          *ArrayOfInt        `xml:"SubtypeIDs,omitempty" json:"SubtypeIDs,omitempty" yaml:"SubtypeIDs,omitempty"`
}

// GetClientContactLogsResponse was auto-generated from WSDL.
type GetClientContactLogsResponse struct {
	GetClientContactLogsResult *GetClientContactLogsResult `xml:"GetClientContactLogsResult,omitempty" json:"GetClientContactLogsResult,omitempty" yaml:"GetClientContactLogsResult,omitempty"`
}

// GetClientContactLogsResult was auto-generated from WSDL.
type GetClientContactLogsResult struct {
	Status           StatusCode         `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string             `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel     `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ContactLogs      *ArrayOfContactLog `xml:"ContactLogs,omitempty" json:"ContactLogs,omitempty" yaml:"ContactLogs,omitempty"`
}

// GetClientContracts was auto-generated from WSDL.
type GetClientContracts struct {
	XMLName xml.Name                   `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientContracts" json:"-" yaml:"-"`
	Request *GetClientContractsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientContractsRequest was auto-generated from WSDL.
type GetClientContractsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
}

// GetClientContractsResponse was auto-generated from WSDL.
type GetClientContractsResponse struct {
	GetClientContractsResult *GetClientContractsResult `xml:"GetClientContractsResult,omitempty" json:"GetClientContractsResult,omitempty" yaml:"GetClientContractsResult,omitempty"`
}

// GetClientContractsResult was auto-generated from WSDL.
type GetClientContractsResult struct {
	Status           StatusCode             `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                    `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                 `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel         `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                    `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                    `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                    `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Contracts        *ArrayOfClientContract `xml:"Contracts,omitempty" json:"Contracts,omitempty" yaml:"Contracts,omitempty"`
}

// GetClientFormulaNotes was auto-generated from WSDL.
type GetClientFormulaNotes struct {
	XMLName xml.Name                      `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientFormulaNotes" json:"-" yaml:"-"`
	Request *GetClientFormulaNotesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientFormulaNotesRequest was auto-generated from WSDL.
type GetClientFormulaNotesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	AppointmentID     int64              `xml:"AppointmentID,omitempty" json:"AppointmentID,omitempty" yaml:"AppointmentID,omitempty"`
}

// GetClientFormulaNotesResponse was auto-generated from WSDL.
type GetClientFormulaNotesResponse struct {
	GetClientFormulaNotesResult *GetClientFormulaNotesResult `xml:"GetClientFormulaNotesResult,omitempty" json:"GetClientFormulaNotesResult,omitempty" yaml:"GetClientFormulaNotesResult,omitempty"`
}

// GetClientFormulaNotesResult was auto-generated from WSDL.
type GetClientFormulaNotesResult struct {
	Status           StatusCode          `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                 `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string              `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel      `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                 `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                 `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                 `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	FormulaNotes     *ArrayOfFormulaNote `xml:"FormulaNotes,omitempty" json:"FormulaNotes,omitempty" yaml:"FormulaNotes,omitempty"`
}

// GetClientIndexes was auto-generated from WSDL.
type GetClientIndexes struct {
	XMLName xml.Name                 `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientIndexes" json:"-" yaml:"-"`
	Request *GetClientIndexesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientIndexesRequest was auto-generated from WSDL.
type GetClientIndexesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	RequiredOnly      bool               `xml:"RequiredOnly,omitempty" json:"RequiredOnly,omitempty" yaml:"RequiredOnly,omitempty"`
}

// GetClientIndexesResponse was auto-generated from WSDL.
type GetClientIndexesResponse struct {
	GetClientIndexesResult *GetClientIndexesResult `xml:"GetClientIndexesResult,omitempty" json:"GetClientIndexesResult,omitempty" yaml:"GetClientIndexesResult,omitempty"`
}

// GetClientIndexesResult was auto-generated from WSDL.
type GetClientIndexesResult struct {
	Status           StatusCode          `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                 `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string              `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel      `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                 `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                 `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                 `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ClientIndexes    *ArrayOfClientIndex `xml:"ClientIndexes,omitempty" json:"ClientIndexes,omitempty" yaml:"ClientIndexes,omitempty"`
}

// GetClientPurchases was auto-generated from WSDL.
type GetClientPurchases struct {
	XMLName xml.Name                   `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientPurchases" json:"-" yaml:"-"`
	Request *GetClientPurchasesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientPurchasesRequest was auto-generated from WSDL.
type GetClientPurchasesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	StartDate         DateTime           `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate           DateTime           `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
	SaleID            int                `xml:"SaleID,omitempty" json:"SaleID,omitempty" yaml:"SaleID,omitempty"`
}

// GetClientPurchasesResponse was auto-generated from WSDL.
type GetClientPurchasesResponse struct {
	GetClientPurchasesResult *GetClientPurchasesResult `xml:"GetClientPurchasesResult,omitempty" json:"GetClientPurchasesResult,omitempty" yaml:"GetClientPurchasesResult,omitempty"`
}

// GetClientPurchasesResult was auto-generated from WSDL.
type GetClientPurchasesResult struct {
	Status           StatusCode       `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int              `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string           `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel   `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int              `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int              `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int              `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Purchases        *ArrayOfSaleItem `xml:"Purchases,omitempty" json:"Purchases,omitempty" yaml:"Purchases,omitempty"`
}

// GetClientReferralTypes was auto-generated from WSDL.
type GetClientReferralTypes struct {
	XMLName xml.Name                       `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientReferralTypes" json:"-" yaml:"-"`
	Request *GetClientReferralTypesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientReferralTypesRequest was auto-generated from WSDL.
type GetClientReferralTypesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	IncludeInactive   bool               `xml:"IncludeInactive,omitempty" json:"IncludeInactive,omitempty" yaml:"IncludeInactive,omitempty"`
}

// GetClientReferralTypesResponse was auto-generated from WSDL.
type GetClientReferralTypesResponse struct {
	GetClientReferralTypesResult *GetClientReferralTypesResult `xml:"GetClientReferralTypesResult,omitempty" json:"GetClientReferralTypesResult,omitempty" yaml:"GetClientReferralTypesResult,omitempty"`
}

// GetClientReferralTypesResult was auto-generated from WSDL.
type GetClientReferralTypesResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ReferralTypes    *ArrayOfString `xml:"ReferralTypes,omitempty" json:"ReferralTypes,omitempty" yaml:"ReferralTypes,omitempty"`
}

// GetClientSchedule was auto-generated from WSDL.
type GetClientSchedule struct {
	XMLName xml.Name                  `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientSchedule" json:"-" yaml:"-"`
	Request *GetClientScheduleRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientScheduleRequest was auto-generated from WSDL.
type GetClientScheduleRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	StartDate         DateTime           `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate           DateTime           `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
}

// GetClientScheduleResponse was auto-generated from WSDL.
type GetClientScheduleResponse struct {
	GetClientScheduleResult *GetClientScheduleResult `xml:"GetClientScheduleResult,omitempty" json:"GetClientScheduleResult,omitempty" yaml:"GetClientScheduleResult,omitempty"`
}

// GetClientScheduleResult was auto-generated from WSDL.
type GetClientScheduleResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Visits           *ArrayOfVisit  `xml:"Visits,omitempty" json:"Visits,omitempty" yaml:"Visits,omitempty"`
}

// GetClientServices was auto-generated from WSDL.
type GetClientServices struct {
	XMLName xml.Name                  `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientServices" json:"-" yaml:"-"`
	Request *GetClientServicesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientServicesRequest was auto-generated from WSDL.
type GetClientServicesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	ClassID           int                `xml:"ClassID" json:"ClassID" yaml:"ClassID"`
	ProgramIDs        *ArrayOfInt        `xml:"ProgramIDs,omitempty" json:"ProgramIDs,omitempty" yaml:"ProgramIDs,omitempty"`
	SessionTypeIDs    *ArrayOfInt        `xml:"SessionTypeIDs,omitempty" json:"SessionTypeIDs,omitempty" yaml:"SessionTypeIDs,omitempty"`
	LocationIDs       *ArrayOfInt        `xml:"LocationIDs,omitempty" json:"LocationIDs,omitempty" yaml:"LocationIDs,omitempty"`
	VisitCount        int                `xml:"VisitCount,omitempty" json:"VisitCount,omitempty" yaml:"VisitCount,omitempty"`
	StartDate         DateTime           `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate           DateTime           `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
	ShowActiveOnly    bool               `xml:"ShowActiveOnly,omitempty" json:"ShowActiveOnly,omitempty" yaml:"ShowActiveOnly,omitempty"`
}

// GetClientServicesResponse was auto-generated from WSDL.
type GetClientServicesResponse struct {
	GetClientServicesResult *GetClientServicesResult `xml:"GetClientServicesResult,omitempty" json:"GetClientServicesResult,omitempty" yaml:"GetClientServicesResult,omitempty"`
}

// GetClientServicesResult was auto-generated from WSDL.
type GetClientServicesResult struct {
	Status           StatusCode            `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                   `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel        `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                   `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                   `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                   `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ClientServices   *ArrayOfClientService `xml:"ClientServices,omitempty" json:"ClientServices,omitempty" yaml:"ClientServices,omitempty"`
}

// GetClientVisits was auto-generated from WSDL.
type GetClientVisits struct {
	XMLName xml.Name                `xml:"http://clients.mindbodyonline.com/api/0_5 GetClientVisits" json:"-" yaml:"-"`
	Request *GetClientVisitsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientVisitsRequest was auto-generated from WSDL.
type GetClientVisitsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	StartDate         DateTime           `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`
	EndDate           DateTime           `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
	UnpaidsOnly       bool               `xml:"UnpaidsOnly" json:"UnpaidsOnly" yaml:"UnpaidsOnly"`
}

// GetClientVisitsResponse was auto-generated from WSDL.
type GetClientVisitsResponse struct {
	GetClientVisitsResult *GetClientVisitsResult `xml:"GetClientVisitsResult,omitempty" json:"GetClientVisitsResult,omitempty" yaml:"GetClientVisitsResult,omitempty"`
}

// GetClientVisitsResult was auto-generated from WSDL.
type GetClientVisitsResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Visits           *ArrayOfVisit  `xml:"Visits,omitempty" json:"Visits,omitempty" yaml:"Visits,omitempty"`
}

// GetClients was auto-generated from WSDL.
type GetClients struct {
	XMLName xml.Name           `xml:"http://clients.mindbodyonline.com/api/0_5 GetClients" json:"-" yaml:"-"`
	Request *GetClientsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetClientsRequest was auto-generated from WSDL.
type GetClientsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientIDs         *ArrayOfString     `xml:"ClientIDs,omitempty" json:"ClientIDs,omitempty" yaml:"ClientIDs,omitempty"`
	SearchText        string             `xml:"SearchText,omitempty" json:"SearchText,omitempty" yaml:"SearchText,omitempty"`
	IsProspect        bool               `xml:"IsProspect,omitempty" json:"IsProspect,omitempty" yaml:"IsProspect,omitempty"`
}

// GetClientsResponse was auto-generated from WSDL.
type GetClientsResponse struct {
	GetClientsResult *GetClientsResult `xml:"GetClientsResult,omitempty" json:"GetClientsResult,omitempty" yaml:"GetClientsResult,omitempty"`
}

// GetClientsResult was auto-generated from WSDL.
type GetClientsResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Clients          *ArrayOfClient `xml:"Clients,omitempty" json:"Clients,omitempty" yaml:"Clients,omitempty"`
}

// GetContactLogTypes was auto-generated from WSDL.
type GetContactLogTypes struct {
	XMLName xml.Name                   `xml:"http://clients.mindbodyonline.com/api/0_5 GetContactLogTypes" json:"-" yaml:"-"`
	Request *GetContactLogTypesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetContactLogTypesRequest was auto-generated from WSDL.
type GetContactLogTypesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetContactLogTypesResponse was auto-generated from WSDL.
type GetContactLogTypesResponse struct {
	GetContactLogTypesResult *GetContactLogTypesResult `xml:"GetContactLogTypesResult,omitempty" json:"GetContactLogTypesResult,omitempty" yaml:"GetContactLogTypesResult,omitempty"`
}

// GetContactLogTypesResult was auto-generated from WSDL.
type GetContactLogTypesResult struct {
	Status           StatusCode             `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                    `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                 `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel         `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                    `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                    `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                    `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ContatctLogTypes *ArrayOfContactLogType `xml:"ContatctLogTypes,omitempty" json:"ContatctLogTypes,omitempty" yaml:"ContatctLogTypes,omitempty"`
}

// GetCustomClientFields was auto-generated from WSDL.
type GetCustomClientFields struct {
	XMLName xml.Name                      `xml:"http://clients.mindbodyonline.com/api/0_5 GetCustomClientFields" json:"-" yaml:"-"`
	Request *GetCustomClientFieldsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetCustomClientFieldsRequest was auto-generated from WSDL.
type GetCustomClientFieldsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetCustomClientFieldsResponse was auto-generated from WSDL.
type GetCustomClientFieldsResponse struct {
	GetCustomClientFieldsResult *GetCustomClientFieldsResult `xml:"GetCustomClientFieldsResult,omitempty" json:"GetCustomClientFieldsResult,omitempty" yaml:"GetCustomClientFieldsResult,omitempty"`
}

// GetCustomClientFieldsResult was auto-generated from WSDL.
type GetCustomClientFieldsResult struct {
	Status             StatusCode                `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode          int                       `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message            string                    `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail          XMLDetailLevel            `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount        int                       `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex   int                       `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount     int                       `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	CustomClientFields *ArrayOfCustomClientField `xml:"CustomClientFields,omitempty" json:"CustomClientFields,omitempty" yaml:"CustomClientFields,omitempty"`
}

// GetRequiredClientFields was auto-generated from WSDL.
type GetRequiredClientFields struct {
	XMLName xml.Name                        `xml:"http://clients.mindbodyonline.com/api/0_5 GetRequiredClientFields" json:"-" yaml:"-"`
	Request *GetRequiredClientFieldsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetRequiredClientFieldsRequest was auto-generated from WSDL.
type GetRequiredClientFieldsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetRequiredClientFieldsResponse was auto-generated from WSDL.
type GetRequiredClientFieldsResponse struct {
	GetRequiredClientFieldsResult *GetRequiredClientFieldsResult `xml:"GetRequiredClientFieldsResult,omitempty" json:"GetRequiredClientFieldsResult,omitempty" yaml:"GetRequiredClientFieldsResult,omitempty"`
}

// GetRequiredClientFieldsResult was auto-generated from WSDL.
type GetRequiredClientFieldsResult struct {
	Status               StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode            int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message              string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail            XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount          int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex     int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount       int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	RequiredClientFields *ArrayOfString `xml:"RequiredClientFields,omitempty" json:"RequiredClientFields,omitempty" yaml:"RequiredClientFields,omitempty"`
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

// Payment was auto-generated from WSDL.
type Payment struct {
	ID       int64   `xml:"ID" json:"ID" yaml:"ID"`
	Amount   float64 `xml:"Amount" json:"Amount" yaml:"Amount"`
	Method   int     `xml:"Method" json:"Method" yaml:"Method"`
	Type     string  `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`
	LastFour string  `xml:"LastFour,omitempty" json:"LastFour,omitempty" yaml:"LastFour,omitempty"`
	Notes    string  `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
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

// Sale was auto-generated from WSDL.
type Sale struct {
	ID           int64           `xml:"ID" json:"ID" yaml:"ID"`
	SaleTime     DateTime        `xml:"SaleTime" json:"SaleTime" yaml:"SaleTime"`
	SaleDate     DateTime        `xml:"SaleDate" json:"SaleDate" yaml:"SaleDate"`
	SaleDateTime DateTime        `xml:"SaleDateTime" json:"SaleDateTime" yaml:"SaleDateTime"`
	Location     *Location       `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	Payments     *ArrayOfPayment `xml:"Payments,omitempty" json:"Payments,omitempty" yaml:"Payments,omitempty"`
}

// SaleItem was auto-generated from WSDL.
type SaleItem struct {
	Site           *Site          `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
	Messages       *ArrayOfString `xml:"Messages,omitempty" json:"Messages,omitempty" yaml:"Messages,omitempty"`
	Execute        string         `xml:"Execute,omitempty" json:"Execute,omitempty" yaml:"Execute,omitempty"`
	ErrorCode      string         `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	Sale           *Sale          `xml:"Sale,omitempty" json:"Sale,omitempty" yaml:"Sale,omitempty"`
	Description    string         `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	AccountPayment bool           `xml:"AccountPayment,omitempty" json:"AccountPayment,omitempty" yaml:"AccountPayment,omitempty"`
	Price          float64        `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`
	AmountPaid     float64        `xml:"AmountPaid,omitempty" json:"AmountPaid,omitempty" yaml:"AmountPaid,omitempty"`
	Discount       float64        `xml:"Discount,omitempty" json:"Discount,omitempty" yaml:"Discount,omitempty"`
	Tax            float64        `xml:"Tax,omitempty" json:"Tax,omitempty" yaml:"Tax,omitempty"`
	Returned       bool           `xml:"Returned,omitempty" json:"Returned,omitempty" yaml:"Returned,omitempty"`
	Quantity       int            `xml:"Quantity,omitempty" json:"Quantity,omitempty" yaml:"Quantity,omitempty"`
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

// SendUserNewPassword was auto-generated from WSDL.
type SendUserNewPassword struct {
	XMLName xml.Name                          `xml:"http://clients.mindbodyonline.com/api/0_5 SendUserNewPassword" json:"-" yaml:"-"`
	Request *ClientSendUserNewPasswordRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// SendUserNewPasswordResponse was auto-generated from WSDL.
type SendUserNewPasswordResponse struct {
	SendUserNewPasswordResult *ClientSendUserNewPasswordResult `xml:"SendUserNewPasswordResult,omitempty" json:"SendUserNewPasswordResult,omitempty" yaml:"SendUserNewPasswordResult,omitempty"`
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

// UpdateClientServices was auto-generated from WSDL.
type UpdateClientServices struct {
	XMLName xml.Name                     `xml:"http://clients.mindbodyonline.com/api/0_5 UpdateClientServices" json:"-" yaml:"-"`
	Request *UpdateClientServicesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// UpdateClientServicesRequest was auto-generated from WSDL.
type UpdateClientServicesRequest struct {
	SourceCredentials *SourceCredentials    `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials      `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel        `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                   `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                   `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString        `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientServices    *ArrayOfClientService `xml:"ClientServices,omitempty" json:"ClientServices,omitempty" yaml:"ClientServices,omitempty"`
	Test              bool                  `xml:"Test,omitempty" json:"Test,omitempty" yaml:"Test,omitempty"`
}

// UpdateClientServicesResponse was auto-generated from WSDL.
type UpdateClientServicesResponse struct {
	UpdateClientServicesResult *UpdateClientServicesResult `xml:"UpdateClientServicesResult,omitempty" json:"UpdateClientServicesResult,omitempty" yaml:"UpdateClientServicesResult,omitempty"`
}

// UpdateClientServicesResult was auto-generated from WSDL.
type UpdateClientServicesResult struct {
	Status           StatusCode            `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                   `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel        `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                   `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                   `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                   `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ClientServices   *ArrayOfClientService `xml:"ClientServices,omitempty" json:"ClientServices,omitempty" yaml:"ClientServices,omitempty"`
}

// UploadClientDocument was auto-generated from WSDL.
type UploadClientDocument struct {
	XMLName xml.Name                     `xml:"http://clients.mindbodyonline.com/api/0_5 UploadClientDocument" json:"-" yaml:"-"`
	Request *UploadClientDocumentRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// UploadClientDocumentRequest was auto-generated from WSDL.
type UploadClientDocumentRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	FileName          string             `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`
	Bytes             []byte             `xml:"Bytes,omitempty" json:"Bytes,omitempty" yaml:"Bytes,omitempty"`
}

// UploadClientDocumentResponse was auto-generated from WSDL.
type UploadClientDocumentResponse struct {
	UploadClientDocumentResult *UploadClientDocumentResult `xml:"UploadClientDocumentResult,omitempty" json:"UploadClientDocumentResult,omitempty" yaml:"UploadClientDocumentResult,omitempty"`
}

// UploadClientDocumentResult was auto-generated from WSDL.
type UploadClientDocumentResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	FileSize         int64          `xml:"FileSize" json:"FileSize" yaml:"FileSize"`
	FileName         string         `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`
}

// UploadClientPhoto was auto-generated from WSDL.
type UploadClientPhoto struct {
	XMLName xml.Name                  `xml:"http://clients.mindbodyonline.com/api/0_5 UploadClientPhoto" json:"-" yaml:"-"`
	Request *UploadClientPhotoRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// UploadClientPhotoRequest was auto-generated from WSDL.
type UploadClientPhotoRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ClientID          string             `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
	Bytes             []byte             `xml:"Bytes,omitempty" json:"Bytes,omitempty" yaml:"Bytes,omitempty"`
}

// UploadClientPhotoResponse was auto-generated from WSDL.
type UploadClientPhotoResponse struct {
	UploadClientPhotoResult *UploadClientPhotoResult `xml:"UploadClientPhotoResult,omitempty" json:"UploadClientPhotoResult,omitempty" yaml:"UploadClientPhotoResult,omitempty"`
}

// UploadClientPhotoResult was auto-generated from WSDL.
type UploadClientPhotoResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	PhotoUrl         string         `xml:"PhotoUrl,omitempty" json:"PhotoUrl,omitempty" yaml:"PhotoUrl,omitempty"`
	ClientID         string         `xml:"ClientID,omitempty" json:"ClientID,omitempty" yaml:"ClientID,omitempty"`
}

// UserCredentials was auto-generated from WSDL.
type UserCredentials struct {
	Username   string      `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password   string      `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	SiteIDs    *ArrayOfInt `xml:"SiteIDs,omitempty" json:"SiteIDs,omitempty" yaml:"SiteIDs,omitempty"`
	LocationID int         `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// ValidateLogin was auto-generated from WSDL.
type ValidateLogin struct {
	XMLName xml.Name              `xml:"http://clients.mindbodyonline.com/api/0_5 ValidateLogin" json:"-" yaml:"-"`
	Request *ValidateLoginRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
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

// ValidateLoginResponse was auto-generated from WSDL.
type ValidateLoginResponse struct {
	ValidateLoginResult *ValidateLoginResult `xml:"ValidateLoginResult,omitempty" json:"ValidateLoginResult,omitempty" yaml:"ValidateLoginResult,omitempty"`
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

// client_x0020_ServiceSoap implements the Client_x0020_ServiceSoap interface.
type client_x0020_ServiceSoap struct {
	cli *soap.Client
}

// Adds an arrival record for the given client.
func (p *client_x0020_ServiceSoap) AddArrival(α *AddArrival) (β *AddArrivalResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M AddArrivalResponse `xml:"AddArrivalResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Adds a formula note to a client.
func (p *client_x0020_ServiceSoap) AddClientFormulaNote(α *AddClientFormulaNote) (β *AddClientFormulaNoteResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M AddClientFormulaNoteResponse `xml:"AddClientFormulaNoteResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Adds or updates information for a list of clients.
func (p *client_x0020_ServiceSoap) AddOrUpdateClients(α *AddOrUpdateClients) (β *AddOrUpdateClientsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M AddOrUpdateClientsResponse `xml:"AddOrUpdateClientsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Add or update client contact logs.
func (p *client_x0020_ServiceSoap) AddOrUpdateContactLogs(α *AddOrUpdateContactLogs) (β *AddOrUpdateContactLogsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M AddOrUpdateContactLogsResponse `xml:"AddOrUpdateContactLogsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Deletes a formula note to a client.
func (p *client_x0020_ServiceSoap) DeleteClientFormulaNote(α *DeleteClientFormulaNote) (β *DeleteClientFormulaNoteResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M DeleteClientFormulaNoteResponse `xml:"DeleteClientFormulaNoteResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets the active membership for a given client.
func (p *client_x0020_ServiceSoap) GetActiveClientMemberships(α *GetActiveClientMemberships) (β *GetActiveClientMembershipsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetActiveClientMembershipsResponse `xml:"GetActiveClientMembershipsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets account balances for the given clients.
func (p *client_x0020_ServiceSoap) GetClientAccountBalances(α *GetClientAccountBalances) (β *GetClientAccountBalancesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientAccountBalancesResponse `xml:"GetClientAccountBalancesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Get contact logs for a client.
func (p *client_x0020_ServiceSoap) GetClientContactLogs(α *GetClientContactLogs) (β *GetClientContactLogsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientContactLogsResponse `xml:"GetClientContactLogsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of contracts for a given client.
func (p *client_x0020_ServiceSoap) GetClientContracts(α *GetClientContracts) (β *GetClientContractsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientContractsResponse `xml:"GetClientContractsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of client formula notes.
func (p *client_x0020_ServiceSoap) GetClientFormulaNotes(α *GetClientFormulaNotes) (β *GetClientFormulaNotesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientFormulaNotesResponse `xml:"GetClientFormulaNotesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of currently available client indexes.
func (p *client_x0020_ServiceSoap) GetClientIndexes(α *GetClientIndexes) (β *GetClientIndexesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientIndexesResponse `xml:"GetClientIndexesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Get purchases for a client.
func (p *client_x0020_ServiceSoap) GetClientPurchases(α *GetClientPurchases) (β *GetClientPurchasesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientPurchasesResponse `xml:"GetClientPurchasesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of clients.
func (p *client_x0020_ServiceSoap) GetClientReferralTypes(α *GetClientReferralTypes) (β *GetClientReferralTypesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientReferralTypesResponse `xml:"GetClientReferralTypesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Get visits for a client.
func (p *client_x0020_ServiceSoap) GetClientSchedule(α *GetClientSchedule) (β *GetClientScheduleResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientScheduleResponse `xml:"GetClientScheduleResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a client service for a given client.
func (p *client_x0020_ServiceSoap) GetClientServices(α *GetClientServices) (β *GetClientServicesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientServicesResponse `xml:"GetClientServicesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Get visits for a client.
func (p *client_x0020_ServiceSoap) GetClientVisits(α *GetClientVisits) (β *GetClientVisitsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientVisitsResponse `xml:"GetClientVisitsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of clients.
func (p *client_x0020_ServiceSoap) GetClients(α *GetClients) (β *GetClientsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetClientsResponse `xml:"GetClientsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Get contact log types for a client.
func (p *client_x0020_ServiceSoap) GetContactLogTypes(α *GetContactLogTypes) (β *GetContactLogTypesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetContactLogTypesResponse `xml:"GetContactLogTypesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of currently available client indexes.
func (p *client_x0020_ServiceSoap) GetCustomClientFields(α *GetCustomClientFields) (β *GetCustomClientFieldsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetCustomClientFieldsResponse `xml:"GetCustomClientFieldsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Updates a client service for a given client.
func (p *client_x0020_ServiceSoap) GetRequiredClientFields(α *GetRequiredClientFields) (β *GetRequiredClientFieldsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetRequiredClientFieldsResponse `xml:"GetRequiredClientFieldsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Sends the user a new password.
func (p *client_x0020_ServiceSoap) SendUserNewPassword(α *SendUserNewPassword) (β *SendUserNewPasswordResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SendUserNewPasswordResponse `xml:"SendUserNewPasswordResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Updates a client service for a given client.
func (p *client_x0020_ServiceSoap) UpdateClientServices(α *UpdateClientServices) (β *UpdateClientServicesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M UpdateClientServicesResponse `xml:"UpdateClientServicesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Upload a client document.
func (p *client_x0020_ServiceSoap) UploadClientDocument(α *UploadClientDocument) (β *UploadClientDocumentResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M UploadClientDocumentResponse `xml:"UploadClientDocumentResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Upload a client photo.
func (p *client_x0020_ServiceSoap) UploadClientPhoto(α *UploadClientPhoto) (β *UploadClientPhotoResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M UploadClientPhotoResponse `xml:"UploadClientPhotoResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Validates a username and password. This method returns the associated
// clients record and a session GUID on success.
func (p *client_x0020_ServiceSoap) ValidateLogin(α *ValidateLogin) (β *ValidateLoginResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ValidateLoginResponse `xml:"ValidateLoginResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
