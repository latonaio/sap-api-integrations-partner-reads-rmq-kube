package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-partner-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToPartnerCollection(raw []byte, l *logger.Logger) ([]PartnerCollection, error) {
	pm := &responses.PartnerCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerCollection := make([]PartnerCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerCollection = append(partnerCollection, PartnerCollection{
			ObjectID:                          data.ObjectID,
			PartnerID:                         data.PartnerID,
			PartnerUUID:                       data.PartnerUUID,
			LifeCycleStatusCode:               data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:           data.LifeCycleStatusCodeText,
			Duns:                              data.Duns,
			LegalFormCode:                     data.LegalFormCode,
			LegalFormCodeText:                 data.LegalFormCodeText,
			PartnerABCClassificationCode:      data.PartnerABCClassificationCode,
			PartnerABCClassificationCodeText:  data.PartnerABCClassificationCodeText,
			IndustrialSectorCode:              data.IndustrialSectorCode,
			IndustrialSectorCodeText:          data.IndustrialSectorCodeText,
			ContactPermissionCode:             data.ContactPermissionCode,
			ContactPermissionCodeText:         data.ContactPermissionCodeText,
			BusinessPartnerFormattedName:      data.BusinessPartnerFormattedName,
			Name:                              data.Name,
			AdditionalName:                    data.AdditionalName,
			AdditionalName2:                   data.AdditionalName2,
			AdditionalName3:                   data.AdditionalName3,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			StateCode:                         data.StateCode,
			StateCodeText:                     data.StateCodeText,
			CareOfName:                        data.CareOfName,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			District:                          data.District,
			City:                              data.City,
			DifferentCity:                     data.DifferentCity,
			StreetPostalCode:                  data.StreetPostalCode,
			County:                            data.County,
			CompanyPostalCode:                 data.CompanyPostalCode,
			POBox:                             data.POBox,
			POBoxPostalCode:                   data.POBoxPostalCode,
			POBoxDeviatingCountryCode:         data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:     data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingStateCode:           data.POBoxDeviatingStateCode,
			POBoxDeviatingStateCodeText:       data.POBoxDeviatingStateCodeText,
			POBoxDeviatingCity:                data.POBoxDeviatingCity,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			Building:                          data.Building,
			Floor:                             data.Floor,
			Room:                              data.Room,
			Phone:                             data.Phone,
			Mobile:                            data.Mobile,
			Fax:                               data.Fax,
			Email:                             data.Email,
			WebSite:                           data.WebSite,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			BestReachedByCode:                 data.BestReachedByCode,
			BestReachedByCodeText:             data.BestReachedByCodeText,
			TotalPoints:                       data.TotalPoints,
			PartnerLevel:                      data.PartnerLevel,
			PartnerLevelText:                  data.PartnerLevelText,
			NormalisedPhone:                   data.NormalisedPhone,
			NormalisedMobile:                  data.NormalisedMobile,
			CreatedOn:                         data.CreatedOn,
			CreatedBy:                         data.CreatedBy,
			CreatedByIdentityUUID:             data.CreatedByIdentityUUID,
			ChangedOn:                         data.ChangedOn,
			ChangedBy:                         data.ChangedBy,
			ChangedByIdentityUUID:             data.ChangedByIdentityUUID,
			EntityLastChangedOn:               data.EntityLastChangedOn,
			ETag:                              data.ETag,
		})
	}

	return partnerCollection, nil
}

func ConvertToPartnerAddressCollection(raw []byte, l *logger.Logger) ([]PartnerAddressCollection, error) {
	pm := &responses.PartnerAddressCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerAddressCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerAddressCollection := make([]PartnerAddressCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerAddressCollection = append(partnerAddressCollection, PartnerAddressCollection{
			ObjectID:                          data.ObjectID,
			ParentObjectID:                    data.ParentObjectID,
			PartnerID:                         data.PartnerID,
			MainIndicator:                     data.MainIndicator,
			ShipTo:                            data.ShipTo,
			DefaultShipTo:                     data.DefaultShipTo,
			BillTo:                            data.BillTo,
			DefaultBillTo:                     data.DefaultBillTo,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			StateCode:                         data.StateCode,
			StateCodeText:                     data.StateCodeText,
			CareOfName:                        data.CareOfName,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			District:                          data.District,
			City:                              data.City,
			DifferentCity:                     data.DifferentCity,
			StreetPostalCode:                  data.StreetPostalCode,
			County:                            data.County,
			CompanyPostalCode:                 data.CompanyPostalCode,
			POBoxIndicator:                    data.POBoxIndicator,
			POBox:                             data.POBox,
			POBoxPostalCode:                   data.POBoxPostalCode,
			POBoxDeviatingCountryCode:         data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:     data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingStateCode:           data.POBoxDeviatingStateCode,
			POBoxDeviatingStateCodeText:       data.POBoxDeviatingStateCodeText,
			POBoxDeviatingCity:                data.POBoxDeviatingCity,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			Building:                          data.Building,
			Floor:                             data.Floor,
			Room:                              data.Room,
			Phone:                             data.Phone,
			Mobile:                            data.Mobile,
			Fax:                               data.Fax,
			Email:                             data.Email,
			WebSite:                           data.WebSite,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			BestReachedByCode:                 data.BestReachedByCode,
			BestReachedByCodeText:             data.BestReachedByCodeText,
			NormalisedPhone:                   data.NormalisedPhone,
			NormalisedMobile:                  data.NormalisedMobile,
			ETag:                              data.ETag,
		})
	}

	return partnerAddressCollection, nil
}

func ConvertToPartnerProgramsCollection(raw []byte, l *logger.Logger) ([]PartnerProgramsCollection, error) {
	pm := &responses.PartnerProgramsCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerProgramsCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerProgramsCollection := make([]PartnerProgramsCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerProgramsCollection = append(partnerProgramsCollection, PartnerProgramsCollection{
			ObjectID:           data.ObjectID,
			ParentObjectID:     data.ParentObjectID,
			PartnerProgram:     data.PartnerProgram,
			PartnerProgramText: data.PartnerProgramText,
			MembershipID:       data.MembershipID,
			PartnerType:        data.PartnerType,
			PartnerTypeText:    data.PartnerTypeText,
			Status:             data.Status,
			StatusText:         data.StatusText,
			AgreementStartDate: data.AgreementStartDate,
			AgreementEndDate:   data.AgreementEndDate,
		})
	}

	return partnerProgramsCollection, nil
}

func ConvertToPartnerProductDimensions(raw []byte, l *logger.Logger) ([]PartnerProductDimensions, error) {
	pm := &responses.PartnerProductDimensions{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerProductDimensions. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerProductDimensions := make([]PartnerProductDimensions, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerProductDimensions = append(partnerProductDimensions, PartnerProductDimensions{
			ObjectID:            data.ObjectID,
			ParentObjectID:      data.ParentObjectID,
			DimensionStatus:     data.DimensionStatus,
			DimensionStatusText: data.DimensionStatusText,
			ProductID:           data.ProductID,
			StartDate:           data.StartDate,
			EndDate:             data.EndDate,
		})
	}

	return partnerProductDimensions, nil
}

func ConvertToPartnerContactCollection(raw []byte, l *logger.Logger) ([]PartnerContactCollection, error) {
	pm := &responses.PartnerContactCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerContactCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerContactCollection := make([]PartnerContactCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerContactCollection = append(partnerContactCollection, PartnerContactCollection{
			ObjectID:            data.ObjectID,
			ParentObjectID:      data.ParentObjectID,
			PartnerID:           data.PartnerID,
			PartnerContactID:    data.PartnerContactID,
			MainIndicator:       data.MainIndicator,
			DepartmentCode:      data.DepartmentCode,
			DepartmentCodeText:  data.DepartmentCodeText,
			FunctionCode:        data.FunctionCode,
			FunctionCodeText:    data.FunctionCodeText,
			VIPReasonCode:       data.VIPReasonCode,
			VIPReasonCodeText:   data.VIPReasonCodeText,
			EntityLastChangedOn: data.EntityLastChangedOn,
			ETag:                data.ETag,
		})
	}

	return partnerContactCollection, nil
}
