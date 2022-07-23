package responses

type PartnerContactCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                         string `json:"ObjectID"`
			ParentObjectID                   string `json:"ParentObjectID"`
			PartnerID                        string `json:"PartnerID"`
			PartnerContactID                 string `json:"PartnerContactID"`
			MainIndicator                    bool   `json:"MainIndicator"`
			DepartmentCode                   string `json:"DepartmentCode"`
			DepartmentCodeText               string `json:"DepartmentCodeText"`
			FunctionCode                     string `json:"FunctionCode"`
			FunctionCodeText                 string `json:"FunctionCodeText"`
			VIPReasonCode                    string `json:"VIPReasonCode"`
			VIPReasonCodeText                string `json:"VIPReasonCodeText"`
			EntityLastChangedOn              string `json:"EntityLastChangedOn"`
			ETag                             string `json:"ETag"`
			PartnerHasPartnerContactFunction struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerHasPartnerContactFunction"`
		} `json:"results"`
	} `json:"d"`
}