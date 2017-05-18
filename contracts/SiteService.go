package contracts

import (
	"encoding/xml"
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://clients.mindbodyonline.com/api/0_5"

// NewSite_x0020_ServiceSoap creates an initializes a Site_x0020_ServiceSoap.
func NewSite_x0020_ServiceSoap(cli *soap.Client) Site_x0020_ServiceSoap {
	return &site_x0020_ServiceSoap{cli}
}

// Site_x0020_ServiceSoap was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Site_x0020_ServiceSoap interface {
	// Gets an activation code.
	GetActivationCode(α *GetActivationCode) (β *GetActivationCodeResponse, err error)

	// Gets a list of prospect stages for a site.
	GetGenders(α *GetGenders) (β *GetGendersResponse, err error)

	// Gets a list of locations.
	GetLocations(α *GetLocations) (β *GetLocationsResponse, err error)

	// Gets a list of active mobile providers.
	GetMobileProviders(α *GetMobileProviders) (β *GetMobileProvidersResponse, err error)

	// Gets a list of programs.
	GetPrograms(α *GetPrograms) (β *GetProgramsResponse, err error)

	// Gets a list of prospect stages for a site.
	GetProspectStages(α *GetProspectStages) (β *GetProspectStagesResponse, err error)

	// Gets a list of relationships.
	GetRelationships(α *GetRelationships) (β *GetRelationshipsResponse, err error)

	// Gets all resources schedule.
	GetResourceSchedule(α *GetResourceSchedule) (β *GetResourceScheduleResponse, err error)

	// Gets a list of resources.
	GetResources(α *GetResources) (β *GetResourcesResponse, err error)

	// Gets a list of session types.
	GetSessionTypes(α *GetSessionTypes) (β *GetSessionTypesResponse, err error)

	// Gets a list of sites.
	GetSites(α *GetSites) (β *GetSitesResponse, err error)

	// Reserves a resource.
	ReserveResource(α *ReserveResource) (β *ReserveResourceResponse, err error)
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

// ArrayOfGenderOption was auto-generated from WSDL.
type ArrayOfGenderOption struct {
	GenderOption []*GenderOption `xml:"GenderOption,omitempty" json:"GenderOption,omitempty" yaml:"GenderOption,omitempty"`
}

// ArrayOfInt was auto-generated from WSDL.
type ArrayOfInt struct {
	Int []int `xml:"int,omitempty" json:"int,omitempty" yaml:"int,omitempty"`
}

// ArrayOfLocation was auto-generated from WSDL.
type ArrayOfLocation struct {
	Location []*Location `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
}

// ArrayOfMobileProvider was auto-generated from WSDL.
type ArrayOfMobileProvider struct {
	MobileProvider []*MobileProvider `xml:"MobileProvider,omitempty" json:"MobileProvider,omitempty" yaml:"MobileProvider,omitempty"`
}

// ArrayOfProgram was auto-generated from WSDL.
type ArrayOfProgram struct {
	Program []*Program `xml:"Program,omitempty" json:"Program,omitempty" yaml:"Program,omitempty"`
}

// ArrayOfProspectStage was auto-generated from WSDL.
type ArrayOfProspectStage struct {
	ProspectStage []*ProspectStage `xml:"ProspectStage,omitempty" json:"ProspectStage,omitempty" yaml:"ProspectStage,omitempty"`
}

// ArrayOfRelationship was auto-generated from WSDL.
type ArrayOfRelationship struct {
	Relationship []*Relationship `xml:"Relationship,omitempty" json:"Relationship,omitempty" yaml:"Relationship,omitempty"`
}

// ArrayOfResource was auto-generated from WSDL.
type ArrayOfResource struct {
	Resource []*Resource `xml:"Resource,omitempty" json:"Resource,omitempty" yaml:"Resource,omitempty"`
}

// ArrayOfSessionType was auto-generated from WSDL.
type ArrayOfSessionType struct {
	SessionType []*SessionType `xml:"SessionType,omitempty" json:"SessionType,omitempty" yaml:"SessionType,omitempty"`
}

// ArrayOfSite was auto-generated from WSDL.
type ArrayOfSite struct {
	Site []*Site `xml:"Site,omitempty" json:"Site,omitempty" yaml:"Site,omitempty"`
}

// ArrayOfString was auto-generated from WSDL.
type ArrayOfString struct {
	String []string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// GenderOption was auto-generated from WSDL.
type GenderOption struct {
	ID       int    `xml:"ID" json:"ID" yaml:"ID"`
	Name     string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
	IsActive bool   `xml:"IsActive" json:"IsActive" yaml:"IsActive"`
	IsSystem bool   `xml:"IsSystem" json:"IsSystem" yaml:"IsSystem"`
}

// GetActivationCode was auto-generated from WSDL.
type GetActivationCode struct {
	XMLName xml.Name                  `xml:"http://clients.mindbodyonline.com/api/0_5 GetActivationCode" json:"-" yaml:"-"`
	Request *GetActivationCodeRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetActivationCodeRequest was auto-generated from WSDL.
type GetActivationCodeRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetActivationCodeResponse was auto-generated from WSDL.
type GetActivationCodeResponse struct {
	GetActivationCodeResult *GetActivationCodeResult `xml:"GetActivationCodeResult,omitempty" json:"GetActivationCodeResult,omitempty" yaml:"GetActivationCodeResult,omitempty"`
}

// GetActivationCodeResult was auto-generated from WSDL.
type GetActivationCodeResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ActivationCode   string         `xml:"ActivationCode,omitempty" json:"ActivationCode,omitempty" yaml:"ActivationCode,omitempty"`
	ActivationLink   string         `xml:"ActivationLink,omitempty" json:"ActivationLink,omitempty" yaml:"ActivationLink,omitempty"`
}

// GetGenders was auto-generated from WSDL.
type GetGenders struct {
	XMLName xml.Name           `xml:"http://clients.mindbodyonline.com/api/0_5 GetGenders" json:"-" yaml:"-"`
	Request *GetGendersRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetGendersRequest was auto-generated from WSDL.
type GetGendersRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	IncludeInactive   bool               `xml:"IncludeInactive,omitempty" json:"IncludeInactive,omitempty" yaml:"IncludeInactive,omitempty"`
}

// GetGendersResponse was auto-generated from WSDL.
type GetGendersResponse struct {
	GetGendersResult *GetGendersResult `xml:"GetGendersResult,omitempty" json:"GetGendersResult,omitempty" yaml:"GetGendersResult,omitempty"`
}

// GetGendersResult was auto-generated from WSDL.
type GetGendersResult struct {
	Status           StatusCode           `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                  `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string               `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel       `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                  `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                  `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                  `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	GenderOptions    *ArrayOfGenderOption `xml:"GenderOptions,omitempty" json:"GenderOptions,omitempty" yaml:"GenderOptions,omitempty"`
}

// GetLocations was auto-generated from WSDL.
type GetLocations struct {
	XMLName xml.Name             `xml:"http://clients.mindbodyonline.com/api/0_5 GetLocations" json:"-" yaml:"-"`
	Request *GetLocationsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetLocationsRequest was auto-generated from WSDL.
type GetLocationsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetLocationsResponse was auto-generated from WSDL.
type GetLocationsResponse struct {
	GetLocationsResult *GetLocationsResult `xml:"GetLocationsResult,omitempty" json:"GetLocationsResult,omitempty" yaml:"GetLocationsResult,omitempty"`
}

// GetLocationsResult was auto-generated from WSDL.
type GetLocationsResult struct {
	Status           StatusCode       `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int              `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string           `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel   `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int              `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int              `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int              `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Locations        *ArrayOfLocation `xml:"Locations,omitempty" json:"Locations,omitempty" yaml:"Locations,omitempty"`
}

// GetMobileProviders was auto-generated from WSDL.
type GetMobileProviders struct {
	XMLName xml.Name                   `xml:"http://clients.mindbodyonline.com/api/0_5 GetMobileProviders" json:"-" yaml:"-"`
	Request *GetMobileProvidersRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetMobileProvidersRequest was auto-generated from WSDL.
type GetMobileProvidersRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetMobileProvidersResponse was auto-generated from WSDL.
type GetMobileProvidersResponse struct {
	GetMobileProvidersResult *GetMobileProvidersResult `xml:"GetMobileProvidersResult,omitempty" json:"GetMobileProvidersResult,omitempty" yaml:"GetMobileProvidersResult,omitempty"`
}

// GetMobileProvidersResult was auto-generated from WSDL.
type GetMobileProvidersResult struct {
	Status           StatusCode             `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                    `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                 `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel         `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                    `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                    `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                    `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	MobileProviders  *ArrayOfMobileProvider `xml:"MobileProviders,omitempty" json:"MobileProviders,omitempty" yaml:"MobileProviders,omitempty"`
}

// GetPrograms was auto-generated from WSDL.
type GetPrograms struct {
	XMLName xml.Name            `xml:"http://clients.mindbodyonline.com/api/0_5 GetPrograms" json:"-" yaml:"-"`
	Request *GetProgramsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetProgramsRequest was auto-generated from WSDL.
type GetProgramsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ScheduleType      ScheduleType       `xml:"ScheduleType" json:"ScheduleType" yaml:"ScheduleType"`
	OnlineOnly        bool               `xml:"OnlineOnly" json:"OnlineOnly" yaml:"OnlineOnly"`
}

// GetProgramsResponse was auto-generated from WSDL.
type GetProgramsResponse struct {
	GetProgramsResult *GetProgramsResult `xml:"GetProgramsResult,omitempty" json:"GetProgramsResult,omitempty" yaml:"GetProgramsResult,omitempty"`
}

// GetProgramsResult was auto-generated from WSDL.
type GetProgramsResult struct {
	Status           StatusCode      `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int             `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string          `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel  `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int             `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int             `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int             `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Programs         *ArrayOfProgram `xml:"Programs,omitempty" json:"Programs,omitempty" yaml:"Programs,omitempty"`
}

// GetProspectStages was auto-generated from WSDL.
type GetProspectStages struct {
	XMLName xml.Name                  `xml:"http://clients.mindbodyonline.com/api/0_5 GetProspectStages" json:"-" yaml:"-"`
	Request *GetProspectStagesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetProspectStagesRequest was auto-generated from WSDL.
type GetProspectStagesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	IncludeInactive   bool               `xml:"IncludeInactive,omitempty" json:"IncludeInactive,omitempty" yaml:"IncludeInactive,omitempty"`
}

// GetProspectStagesResponse was auto-generated from WSDL.
type GetProspectStagesResponse struct {
	GetProspectStagesResult *GetProspectStagesResult `xml:"GetProspectStagesResult,omitempty" json:"GetProspectStagesResult,omitempty" yaml:"GetProspectStagesResult,omitempty"`
}

// GetProspectStagesResult was auto-generated from WSDL.
type GetProspectStagesResult struct {
	Status           StatusCode            `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                   `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string                `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel        `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                   `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                   `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                   `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	ProspectStages   *ArrayOfProspectStage `xml:"ProspectStages,omitempty" json:"ProspectStages,omitempty" yaml:"ProspectStages,omitempty"`
}

// GetRelationships was auto-generated from WSDL.
type GetRelationships struct {
	XMLName xml.Name                 `xml:"http://clients.mindbodyonline.com/api/0_5 GetRelationships" json:"-" yaml:"-"`
	Request *GetRelationshipsRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetRelationshipsRequest was auto-generated from WSDL.
type GetRelationshipsRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
}

// GetRelationshipsResponse was auto-generated from WSDL.
type GetRelationshipsResponse struct {
	GetRelationshipsResult *GetRelationshipsResult `xml:"GetRelationshipsResult,omitempty" json:"GetRelationshipsResult,omitempty" yaml:"GetRelationshipsResult,omitempty"`
}

// GetRelationshipsResult was auto-generated from WSDL.
type GetRelationshipsResult struct {
	Status           StatusCode           `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                  `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string               `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel       `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                  `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                  `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                  `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Relationships    *ArrayOfRelationship `xml:"Relationships,omitempty" json:"Relationships,omitempty" yaml:"Relationships,omitempty"`
}

// GetResourceSchedule was auto-generated from WSDL.
type GetResourceSchedule struct {
	XMLName xml.Name                    `xml:"http://clients.mindbodyonline.com/api/0_5 GetResourceSchedule" json:"-" yaml:"-"`
	Request *GetResourceScheduleRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetResourceScheduleRequest was auto-generated from WSDL.
type GetResourceScheduleRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	Date              DateTime           `xml:"Date,omitempty" json:"Date,omitempty" yaml:"Date,omitempty"`
}

// GetResourceScheduleResponse was auto-generated from WSDL.
type GetResourceScheduleResponse struct {
	GetResourceScheduleResult *GetResourceScheduleResult `xml:"GetResourceScheduleResult,omitempty" json:"GetResourceScheduleResult,omitempty" yaml:"GetResourceScheduleResult,omitempty"`
}

// GetResourceScheduleResult was auto-generated from WSDL.
type GetResourceScheduleResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Results          *RecordSet     `xml:"Results,omitempty" json:"Results,omitempty" yaml:"Results,omitempty"`
}

// GetResources was auto-generated from WSDL.
type GetResources struct {
	XMLName xml.Name             `xml:"http://clients.mindbodyonline.com/api/0_5 GetResources" json:"-" yaml:"-"`
	Request *GetResourcesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetResourcesRequest was auto-generated from WSDL.
type GetResourcesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	SessionTypeIDs    *ArrayOfInt        `xml:"SessionTypeIDs,omitempty" json:"SessionTypeIDs,omitempty" yaml:"SessionTypeIDs,omitempty"`
	LocationID        int                `xml:"LocationID" json:"LocationID" yaml:"LocationID"`
	StartDateTime     DateTime           `xml:"StartDateTime" json:"StartDateTime" yaml:"StartDateTime"`
	EndDateTime       DateTime           `xml:"EndDateTime" json:"EndDateTime" yaml:"EndDateTime"`
}

// GetResourcesResponse was auto-generated from WSDL.
type GetResourcesResponse struct {
	GetResourcesResult *GetResourcesResult `xml:"GetResourcesResult,omitempty" json:"GetResourcesResult,omitempty" yaml:"GetResourcesResult,omitempty"`
}

// GetResourcesResult was auto-generated from WSDL.
type GetResourcesResult struct {
	Status           StatusCode       `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int              `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string           `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel   `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int              `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int              `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int              `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Resources        *ArrayOfResource `xml:"Resources,omitempty" json:"Resources,omitempty" yaml:"Resources,omitempty"`
}

// GetSessionTypes was auto-generated from WSDL.
type GetSessionTypes struct {
	XMLName xml.Name                `xml:"http://clients.mindbodyonline.com/api/0_5 GetSessionTypes" json:"-" yaml:"-"`
	Request *GetSessionTypesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetSessionTypesRequest was auto-generated from WSDL.
type GetSessionTypesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ProgramIDs        *ArrayOfInt        `xml:"ProgramIDs,omitempty" json:"ProgramIDs,omitempty" yaml:"ProgramIDs,omitempty"`
	OnlineOnly        bool               `xml:"OnlineOnly" json:"OnlineOnly" yaml:"OnlineOnly"`
}

// GetSessionTypesResponse was auto-generated from WSDL.
type GetSessionTypesResponse struct {
	GetSessionTypesResult *GetSessionTypesResult `xml:"GetSessionTypesResult,omitempty" json:"GetSessionTypesResult,omitempty" yaml:"GetSessionTypesResult,omitempty"`
}

// GetSessionTypesResult was auto-generated from WSDL.
type GetSessionTypesResult struct {
	Status           StatusCode          `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int                 `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string              `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel      `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int                 `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int                 `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int                 `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	SessionTypes     *ArrayOfSessionType `xml:"SessionTypes,omitempty" json:"SessionTypes,omitempty" yaml:"SessionTypes,omitempty"`
}

// GetSites was auto-generated from WSDL.
type GetSites struct {
	XMLName xml.Name         `xml:"http://clients.mindbodyonline.com/api/0_5 GetSites" json:"-" yaml:"-"`
	Request *GetSitesRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// GetSitesRequest was auto-generated from WSDL.
type GetSitesRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	SearchText        string             `xml:"SearchText,omitempty" json:"SearchText,omitempty" yaml:"SearchText,omitempty"`
	RelatedSiteID     int                `xml:"RelatedSiteID,omitempty" json:"RelatedSiteID,omitempty" yaml:"RelatedSiteID,omitempty"`
	ShowOnlyTotalWOD  bool               `xml:"ShowOnlyTotalWOD,omitempty" json:"ShowOnlyTotalWOD,omitempty" yaml:"ShowOnlyTotalWOD,omitempty"`
}

// GetSitesResponse was auto-generated from WSDL.
type GetSitesResponse struct {
	GetSitesResult *GetSitesResult `xml:"GetSitesResult,omitempty" json:"GetSitesResult,omitempty" yaml:"GetSitesResult,omitempty"`
}

// GetSitesResult was auto-generated from WSDL.
type GetSitesResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
	Sites            *ArrayOfSite   `xml:"Sites,omitempty" json:"Sites,omitempty" yaml:"Sites,omitempty"`
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

// MobileProvider was auto-generated from WSDL.
type MobileProvider struct {
	ProviderID      int    `xml:"ProviderID" json:"ProviderID" yaml:"ProviderID"`
	ProviderName    string `xml:"ProviderName,omitempty" json:"ProviderName,omitempty" yaml:"ProviderName,omitempty"`
	ProviderAddress string `xml:"ProviderAddress,omitempty" json:"ProviderAddress,omitempty" yaml:"ProviderAddress,omitempty"`
	Active          bool   `xml:"Active" json:"Active" yaml:"Active"`
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

// RecordSet was auto-generated from WSDL.
type RecordSet struct {
	Row []*Row `xml:"Row,omitempty" json:"Row,omitempty" yaml:"Row,omitempty"`
}

// Relationship was auto-generated from WSDL.
type Relationship struct {
	ID                int    `xml:"ID" json:"ID" yaml:"ID"`
	RelationshipName1 string `xml:"RelationshipName1,omitempty" json:"RelationshipName1,omitempty" yaml:"RelationshipName1,omitempty"`
	RelationshipName2 string `xml:"RelationshipName2,omitempty" json:"RelationshipName2,omitempty" yaml:"RelationshipName2,omitempty"`
}

// ReserveResource was auto-generated from WSDL.
type ReserveResource struct {
	XMLName xml.Name                `xml:"http://clients.mindbodyonline.com/api/0_5 ReserveResource" json:"-" yaml:"-"`
	Request *ReserveResourceRequest `xml:"Request,omitempty" json:"Request,omitempty" yaml:"Request,omitempty"`
}

// ReserveResourceRequest was auto-generated from WSDL.
type ReserveResourceRequest struct {
	SourceCredentials *SourceCredentials `xml:"SourceCredentials,omitempty" json:"SourceCredentials,omitempty" yaml:"SourceCredentials,omitempty"`
	UserCredentials   *UserCredentials   `xml:"UserCredentials,omitempty" json:"UserCredentials,omitempty" yaml:"UserCredentials,omitempty"`
	XMLDetail         XMLDetailLevel     `xml:"XMLDetail,omitempty" json:"XMLDetail,omitempty" yaml:"XMLDetail,omitempty"`
	PageSize          int                `xml:"PageSize,omitempty" json:"PageSize,omitempty" yaml:"PageSize,omitempty"`
	CurrentPageIndex  int                `xml:"CurrentPageIndex,omitempty" json:"CurrentPageIndex,omitempty" yaml:"CurrentPageIndex,omitempty"`
	Fields            *ArrayOfString     `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`
	ResourceID        int                `xml:"ResourceID" json:"ResourceID" yaml:"ResourceID"`
	ClientID          int                `xml:"ClientID" json:"ClientID" yaml:"ClientID"`
	StaffID           int                `xml:"StaffID" json:"StaffID" yaml:"StaffID"`
	StartDateTime     DateTime           `xml:"StartDateTime" json:"StartDateTime" yaml:"StartDateTime"`
	EndDateTime       DateTime           `xml:"EndDateTime" json:"EndDateTime" yaml:"EndDateTime"`
	LocationID        int                `xml:"LocationID" json:"LocationID" yaml:"LocationID"`
	ProgramID         int                `xml:"ProgramID" json:"ProgramID" yaml:"ProgramID"`
	Notes             string             `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

// ReserveResourceResponse was auto-generated from WSDL.
type ReserveResourceResponse struct {
	ReserveResourceResult *ReserveResourceResult `xml:"ReserveResourceResult,omitempty" json:"ReserveResourceResult,omitempty" yaml:"ReserveResourceResult,omitempty"`
}

// ReserveResourceResult was auto-generated from WSDL.
type ReserveResourceResult struct {
	Status           StatusCode     `xml:"Status" json:"Status" yaml:"Status"`
	ErrorCode        int            `xml:"ErrorCode" json:"ErrorCode" yaml:"ErrorCode"`
	Message          string         `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
	XMLDetail        XMLDetailLevel `xml:"XMLDetail" json:"XMLDetail" yaml:"XMLDetail"`
	ResultCount      int            `xml:"ResultCount" json:"ResultCount" yaml:"ResultCount"`
	CurrentPageIndex int            `xml:"CurrentPageIndex" json:"CurrentPageIndex" yaml:"CurrentPageIndex"`
	TotalPageCount   int            `xml:"TotalPageCount" json:"TotalPageCount" yaml:"TotalPageCount"`
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

// Row was auto-generated from WSDL.
type Row struct {
	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
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

// UserCredentials was auto-generated from WSDL.
type UserCredentials struct {
	Username   string      `xml:"Username,omitempty" json:"Username,omitempty" yaml:"Username,omitempty"`
	Password   string      `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	SiteIDs    *ArrayOfInt `xml:"SiteIDs,omitempty" json:"SiteIDs,omitempty" yaml:"SiteIDs,omitempty"`
	LocationID int         `xml:"LocationID,omitempty" json:"LocationID,omitempty" yaml:"LocationID,omitempty"`
}

// site_x0020_ServiceSoap implements the Site_x0020_ServiceSoap interface.
type site_x0020_ServiceSoap struct {
	cli *soap.Client
}

// Gets an activation code.
func (p *site_x0020_ServiceSoap) GetActivationCode(α *GetActivationCode) (β *GetActivationCodeResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetActivationCodeResponse `xml:"GetActivationCodeResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of prospect stages for a site.
func (p *site_x0020_ServiceSoap) GetGenders(α *GetGenders) (β *GetGendersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetGendersResponse `xml:"GetGendersResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of locations.
func (p *site_x0020_ServiceSoap) GetLocations(α *GetLocations) (β *GetLocationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetLocationsResponse `xml:"GetLocationsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of active mobile providers.
func (p *site_x0020_ServiceSoap) GetMobileProviders(α *GetMobileProviders) (β *GetMobileProvidersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetMobileProvidersResponse `xml:"GetMobileProvidersResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of programs.
func (p *site_x0020_ServiceSoap) GetPrograms(α *GetPrograms) (β *GetProgramsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetProgramsResponse `xml:"GetProgramsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of prospect stages for a site.
func (p *site_x0020_ServiceSoap) GetProspectStages(α *GetProspectStages) (β *GetProspectStagesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetProspectStagesResponse `xml:"GetProspectStagesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of relationships.
func (p *site_x0020_ServiceSoap) GetRelationships(α *GetRelationships) (β *GetRelationshipsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetRelationshipsResponse `xml:"GetRelationshipsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets all resources schedule.
func (p *site_x0020_ServiceSoap) GetResourceSchedule(α *GetResourceSchedule) (β *GetResourceScheduleResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetResourceScheduleResponse `xml:"GetResourceScheduleResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of resources.
func (p *site_x0020_ServiceSoap) GetResources(α *GetResources) (β *GetResourcesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetResourcesResponse `xml:"GetResourcesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of session types.
func (p *site_x0020_ServiceSoap) GetSessionTypes(α *GetSessionTypes) (β *GetSessionTypesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetSessionTypesResponse `xml:"GetSessionTypesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Gets a list of sites.
func (p *site_x0020_ServiceSoap) GetSites(α *GetSites) (β *GetSitesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetSitesResponse `xml:"GetSitesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Reserves a resource.
func (p *site_x0020_ServiceSoap) ReserveResource(α *ReserveResource) (β *ReserveResourceResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ReserveResourceResponse `xml:"ReserveResourceResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
