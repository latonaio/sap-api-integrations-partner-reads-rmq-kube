package responses

type PartnerAddressCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			Partner                           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"Partner"`
		} `json:"results"`
	} `json:"d"`
}
