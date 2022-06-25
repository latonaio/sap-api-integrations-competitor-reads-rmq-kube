package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-competitor-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCompetitorCollection(raw []byte, l *logger.Logger) ([]CompetitorCollection, error) {
	pm := &responses.CompetitorCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CompetitorCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	competitorCollection := make([]CompetitorCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		competitorCollection = append(competitorCollection, CompetitorCollection{
			ObjectID:                          data.ObjectID,
			CompetitorID:                      data.CompetitorID,
			CompetitorUUID:                    data.CompetitorUUID,
			StatusCode:                        data.StatusCode,
			StatusCodeText:                    data.StatusCodeText,
			ClassificationCode:                data.ClassificationCode,
			ClassificationCodeText:            data.ClassificationCodeText,
			BusinessPartnerFormattedName:      data.BusinessPartnerFormattedName,
			Name:                              data.Name,
			AdditionalName:                    data.AdditionalName,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			RegionCode:                        data.RegionCode,
			RegionCodeText:                    data.RegionCodeText,
			CareOfName:                        data.CareOfName,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			City:                              data.City,
			AdditionalCityName:                data.AdditionalCityName,
			District:                          data.District,
			County:                            data.County,
			CompanyPostalCode:                 data.CompanyPostalCode,
			StreetPostalCode:                  data.StreetPostalCode,
			POBoxPostalCode:                   data.POBoxPostalCode,
			POBox:                             data.POBox,
			POBoxDeviatingCountryCode:         data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:     data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingCity:                data.POBoxDeviatingCity,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			TaxJurisdictionCode:               data.TaxJurisdictionCode,
			TaxJurisdictionCodeText:           data.TaxJurisdictionCodeText,
			POBoxDeviatingStateCode:           data.POBoxDeviatingStateCode,
			POBoxDeviatingStateCodeText:       data.POBoxDeviatingStateCodeText,
			Phone:                             data.Phone,
			Fax:                               data.Fax,
			Email:                             data.Email,
			WebSite:                           data.WebSite,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			BestReachedByCode:                 data.BestReachedByCode,
			BestReachedByCodeText:             data.BestReachedByCodeText,
			NormalisedPhone:                   data.NormalisedPhone,
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

	return competitorCollection, nil
}
