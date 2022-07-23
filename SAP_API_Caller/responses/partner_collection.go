package responses

type PartnerCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			PartnerAddress                    struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerAddress"`
			PartnerAttachment struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerAttachment"`
			PartnerHasPartnerContact struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerHasPartnerContact"`
			PartnerPrograms struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerPrograms"`
			PartnerSalesOrganisation struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerSalesOrganisation"`
		} `json:"results"`
	} `json:"d"`
}
