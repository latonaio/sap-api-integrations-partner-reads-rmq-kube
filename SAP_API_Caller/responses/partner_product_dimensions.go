package responses

type PartnerProductDimensions struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ObjectID            string `json:"ObjectID"`
			ParentObjectID      string `json:"ParentObjectID"`
			DimensionStatus     string `json:"DimensionStatus"`
			DimensionStatusText string `json:"DimensionStatusText"`
			ProductID           string `json:"ProductID"`
			StartDate           string `json:"StartDate"`
			EndDate             string `json:"EndDate"`
			PartnerPrograms     struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PartnerPrograms"`
		} `json:"results"`
	} `json:"d"`
}