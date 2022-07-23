package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-partner-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetPartner(partnerID, countryCode, partnerProgram, dimensionStatus, partnerContactID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "PartnerCollection":
			func() {
				c.PartnerCollection(partnerID)
				wg.Done()
			}()
		case "PartnerAddressCollection":
			func() {
				c.PartnerAddressCollection(countryCode)
				wg.Done()
			}()
		case "PartnerProgramsCollection":
			func() {
				c.PartnerProgramsCollection(partnerProgram)
				wg.Done()
			}()
		case "PartnerProductDimensions":
			func() {
				c.PartnerProductDimensions(dimensionStatus)
				wg.Done()
			}()
		case "PartnerContactCollection":
			func() {
				c.PartnerContactCollection(partnerContactID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) PartnerCollection(partnerID string) {
	data, err := c.callPartnerSrvAPIRequirementPartnerCollection("PartnerCollectionData", partnerID)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "PartnerCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPartnerSrvAPIRequirementPartnerCollection(api, partnerID string) ([]sap_api_output_formatter.PartnerCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerCollection(req, partnerID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerAddressCollection(countryCode string) {
	data, err := c.callPartnerSrvAPIRequirementPartnerAddressCollection("PartnerAddressCollectionData", countryCode)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "PartnerAddressCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPartnerSrvAPIRequirementPartnerAddressCollection(api, countryCode string) ([]sap_api_output_formatter.PartnerAddressCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerAddressCollection(req, countryCode)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerAddressCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerProgramsCollection(partnerProgram string) {
	data, err := c.callPartnerSrvAPIRequirementPartnerProgramsCollection("PartnerProgramsCollectionData", partnerProgram)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "PartnerProgramsCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPartnerSrvAPIRequirementPartnerProgramsCollection(api, partnerProgram string) ([]sap_api_output_formatter.PartnerProgramsCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerProgramsCollection(req, partnerProgram)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerProgramsCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerProductDimensions(dimensionStatus string) {
	data, err := c.callPartnerSrvAPIRequirementPartnerProductDimensions("PartnerProductDimensionsData", dimensionStatus)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "PartnerProductDimensionsData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPartnerSrvAPIRequirementPartnerProductDimensions(api, dimensionStatus string) ([]sap_api_output_formatter.PartnerProductDimensions, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerProductDimensions(req, dimensionStatus)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerProductDimensions(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerContactCollection(partnerContactID string) {
	data, err := c.callPartnerSrvAPIRequirementPartnerContactCollection("PartnerContactCollectionData", partnerContactID)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "PartnerContactCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callPartnerSrvAPIRequirementPartnerContactCollection(api, partnerContactID string) ([]sap_api_output_formatter.PartnerContactCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerContactCollection(req, partnerContactID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerContactCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithPartnerCollection(req *http.Request, partnerID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PartnerID eq '%s'", partnerID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPartnerAddressCollection(req *http.Request, countryCode string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CountryCode eq '%s'", countryCode))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPartnerProgramsCollection(req *http.Request, partnerProgram string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PartnerProgram eq '%s'", partnerProgram))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPartnerProductDimensions(req *http.Request, dimensionStatus string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("DimensionStatus eq '%s'", dimensionStatus))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPartnerContactCollection(req *http.Request, partnerContactID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PartnerContactID eq '%s'", partnerContactID))
	req.URL.RawQuery = params.Encode()
}
