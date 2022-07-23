package responses

type PartnerProgramsCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ObjectID                 string `json:"ObjectID"`
			ParentObjectID           string `json:"ParentObjectID"`
			PartnerProgram           string `json:"PartnerProgram"`
			PartnerProgramText       string `json:"PartnerProgramText"`
			MembershipID             string `json:"MembershipID"`
			PartnerType              string `json:"PartnerType"`
			PartnerTypeText          string `json:"PartnerTypeText"`
			Status                   string `json:"Status"`
			StatusText               string `json:"StatusText"`
			AgreementStartDate       string `json:"AgreementStartDate"`
			AgreementEndDate         string `json:"AgreementEndDate"`
			PartnerProductDimensions struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerProductDimensions"`
		} `json:"results"`
	} `json:"d"`
}
