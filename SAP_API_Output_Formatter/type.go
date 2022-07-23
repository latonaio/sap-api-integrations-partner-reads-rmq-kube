package sap_api_output_formatter

type Partner struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	PartnerCode   string `json:"partner_code"`
	Deleted       bool   `json:"deleted"`
}

type PartnerCollection struct {
	ObjectID                          string `json:"ObjectID"`
	PartnerID                         string `json:"PartnerID"`
	PartnerUUID                       string `json:"PartnerUUID"`
	LifeCycleStatusCode               string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText           string `json:"LifeCycleStatusCodeText"`
	Duns                              string `json:"DUNS"`
	LegalFormCode                     string `json:"LegalFormCode"`
	LegalFormCodeText                 string `json:"LegalFormCodeText"`
	PartnerABCClassificationCode      string `json:"PartnerABCClassificationCode"`
	PartnerABCClassificationCodeText  string `json:"PartnerABCClassificationCodeText"`
	IndustrialSectorCode              string `json:"IndustrialSectorCode"`
	IndustrialSectorCodeText          string `json:"IndustrialSectorCodeText"`
	ContactPermissionCode             string `json:"ContactPermissionCode"`
	ContactPermissionCodeText         string `json:"ContactPermissionCodeText"`
	BusinessPartnerFormattedName      string `json:"BusinessPartnerFormattedName"`
	Name                              string `json:"Name"`
	AdditionalName                    string `json:"AdditionalName"`
	AdditionalName2                   string `json:"AdditionalName2"`
	AdditionalName3                   string `json:"AdditionalName3"`
	FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
	CountryCode                       string `json:"CountryCode"`
	CountryCodeText                   string `json:"CountryCodeText"`
	StateCode                         string `json:"StateCode"`
	StateCodeText                     string `json:"StateCodeText"`
	CareOfName                        string `json:"CareOfName"`
	AddressLine1                      string `json:"AddressLine1"`
	AddressLine2                      string `json:"AddressLine2"`
	HouseNumber                       string `json:"HouseNumber"`
	Street                            string `json:"Street"`
	AddressLine4                      string `json:"AddressLine4"`
	AddressLine5                      string `json:"AddressLine5"`
	District                          string `json:"District"`
	City                              string `json:"City"`
	DifferentCity                     string `json:"DifferentCity"`
	StreetPostalCode                  string `json:"StreetPostalCode"`
	County                            string `json:"County"`
	CompanyPostalCode                 string `json:"CompanyPostalCode"`
	POBox                             string `json:"POBox"`
	POBoxPostalCode                   string `json:"POBoxPostalCode"`
	POBoxDeviatingCountryCode         string `json:"POBoxDeviatingCountryCode"`
	POBoxDeviatingCountryCodeText     string `json:"POBoxDeviatingCountryCodeText"`
	POBoxDeviatingStateCode           string `json:"POBoxDeviatingStateCode"`
	POBoxDeviatingStateCodeText       string `json:"POBoxDeviatingStateCodeText"`
	POBoxDeviatingCity                string `json:"POBoxDeviatingCity"`
	TimeZoneCode                      string `json:"TimeZoneCode"`
	TimeZoneCodeText                  string `json:"TimeZoneCodeText"`
	Building                          string `json:"Building"`
	Floor                             string `json:"Floor"`
	Room                              string `json:"Room"`
	Phone                             string `json:"Phone"`
	Mobile                            string `json:"Mobile"`
	Fax                               string `json:"Fax"`
	Email                             string `json:"Email"`
	WebSite                           string `json:"WebSite"`
	LanguageCode                      string `json:"LanguageCode"`
	LanguageCodeText                  string `json:"LanguageCodeText"`
	BestReachedByCode                 string `json:"BestReachedByCode"`
	BestReachedByCodeText             string `json:"BestReachedByCodeText"`
	TotalPoints                       string `json:"TotalPoints"`
	PartnerLevel                      string `json:"PartnerLevel"`
	PartnerLevelText                  string `json:"PartnerLevelText"`
	NormalisedPhone                   string `json:"NormalisedPhone"`
	NormalisedMobile                  string `json:"NormalisedMobile"`
	CreatedOn                         string `json:"CreatedOn"`
	CreatedBy                         string `json:"CreatedBy"`
	CreatedByIdentityUUID             string `json:"CreatedByIdentityUUID"`
	ChangedOn                         string `json:"ChangedOn"`
	ChangedBy                         string `json:"ChangedBy"`
	ChangedByIdentityUUID             string `json:"ChangedByIdentityUUID"`
	EntityLastChangedOn               string `json:"EntityLastChangedOn"`
	ETag                              string `json:"ETag"`
}

type PartnerAddressCollection struct {
	ObjectID                          string `json:"ObjectID"`
	ParentObjectID                    string `json:"ParentObjectID"`
	PartnerID                         string `json:"PartnerID"`
	MainIndicator                     bool   `json:"MainIndicator"`
	ShipTo                            bool   `json:"ShipTo"`
	DefaultShipTo                     bool   `json:"DefaultShipTo"`
	BillTo                            bool   `json:"BillTo"`
	DefaultBillTo                     bool   `json:"DefaultBillTo"`
	FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
	CountryCode                       string `json:"CountryCode"`
	CountryCodeText                   string `json:"CountryCodeText"`
	StateCode                         string `json:"StateCode"`
	StateCodeText                     string `json:"StateCodeText"`
	CareOfName                        string `json:"CareOfName"`
	AddressLine1                      string `json:"AddressLine1"`
	AddressLine2                      string `json:"AddressLine2"`
	HouseNumber                       string `json:"HouseNumber"`
	Street                            string `json:"Street"`
	AddressLine4                      string `json:"AddressLine4"`
	AddressLine5                      string `json:"AddressLine5"`
	District                          string `json:"District"`
	City                              string `json:"City"`
	DifferentCity                     string `json:"DifferentCity"`
	StreetPostalCode                  string `json:"StreetPostalCode"`
	County                            string `json:"County"`
	CompanyPostalCode                 string `json:"CompanyPostalCode"`
	POBoxIndicator                    bool   `json:"POBoxIndicator"`
	POBox                             string `json:"POBox"`
	POBoxPostalCode                   string `json:"POBoxPostalCode"`
	POBoxDeviatingCountryCode         string `json:"POBoxDeviatingCountryCode"`
	POBoxDeviatingCountryCodeText     string `json:"POBoxDeviatingCountryCodeText"`
	POBoxDeviatingStateCode           string `json:"POBoxDeviatingStateCode"`
	POBoxDeviatingStateCodeText       string `json:"POBoxDeviatingStateCodeText"`
	POBoxDeviatingCity                string `json:"POBoxDeviatingCity"`
	TimeZoneCode                      string `json:"TimeZoneCode"`
	TimeZoneCodeText                  string `json:"TimeZoneCodeText"`
	Building                          string `json:"Building"`
	Floor                             string `json:"Floor"`
	Room                              string `json:"Room"`
	Phone                             string `json:"Phone"`
	Mobile                            string `json:"Mobile"`
	Fax                               string `json:"Fax"`
	Email                             string `json:"Email"`
	WebSite                           string `json:"WebSite"`
	LanguageCode                      string `json:"LanguageCode"`
	LanguageCodeText                  string `json:"LanguageCodeText"`
	BestReachedByCode                 string `json:"BestReachedByCode"`
	BestReachedByCodeText             string `json:"BestReachedByCodeText"`
	NormalisedPhone                   string `json:"NormalisedPhone"`
	NormalisedMobile                  string `json:"NormalisedMobile"`
	ETag                              string `json:"ETag"`
}

type PartnerProgramsCollection struct {
	ObjectID           string `json:"ObjectID"`
	ParentObjectID     string `json:"ParentObjectID"`
	PartnerProgram     string `json:"PartnerProgram"`
	PartnerProgramText string `json:"PartnerProgramText"`
	MembershipID       string `json:"MembershipID"`
	PartnerType        string `json:"PartnerType"`
	PartnerTypeText    string `json:"PartnerTypeText"`
	Status             string `json:"Status"`
	StatusText         string `json:"StatusText"`
	AgreementStartDate string `json:"AgreementStartDate"`
	AgreementEndDate   string `json:"AgreementEndDate"`
}

type PartnerProductDimensions struct {
	ObjectID            string `json:"ObjectID"`
	ParentObjectID      string `json:"ParentObjectID"`
	DimensionStatus     string `json:"DimensionStatus"`
	DimensionStatusText string `json:"DimensionStatusText"`
	ProductID           string `json:"ProductID"`
	StartDate           string `json:"StartDate"`
	EndDate             string `json:"EndDate"`
}

type PartnerContactCollection struct {
	ObjectID            string `json:"ObjectID"`
	ParentObjectID      string `json:"ParentObjectID"`
	PartnerID           string `json:"PartnerID"`
	PartnerContactID    string `json:"PartnerContactID"`
	MainIndicator       bool   `json:"MainIndicator"`
	DepartmentCode      string `json:"DepartmentCode"`
	DepartmentCodeText  string `json:"DepartmentCodeText"`
	FunctionCode        string `json:"FunctionCode"`
	FunctionCodeText    string `json:"FunctionCodeText"`
	VIPReasonCode       string `json:"VIPReasonCode"`
	VIPReasonCodeText   string `json:"VIPReasonCodeText"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
	ETag                string `json:"ETag"`
}