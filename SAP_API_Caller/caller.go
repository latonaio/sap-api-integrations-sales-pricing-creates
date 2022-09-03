package sap_api_caller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-sales-pricing-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-sales-pricing-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	postClient      *sap_api_post_header_setup.SAPPostClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, postClient *sap_api_post_header_setup.SAPPostClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		postClient:      postClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostSalesPricing(
	salesPricingConditionValidity *requests.SalesPricingConditionValidity,
	salesPricingConditionRecord *requests.SalesPricingConditionRecord,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "SalesPricingConditionValidity":
			func() {
				c.SalesPricingConditionValidity(salesPricingConditionValidity)
				wg.Done()
			}()
		case "SalesPricingConditionRecord":
			func() {
				c.SalesPricingConditionRecord(salesPricingConditionRecord)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) SalesPricingConditionValidity(salesPricingConditionValidity *requests.SalesPricingConditionValidity) {
	outputDataSalesPricingConditionValidity, err := c.callSalesPricingSrvAPIRequirementSalesPricingConditionValidity("A_SlsPrcgCndnRecdValidity", salesPricingConditionValidity)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataSalesPricingConditionValidity)
}

func (c *SAPAPICaller) callSalesPricingSrvAPIRequirementSalesPricingConditionValidity(api string, salesPricingConditionValidity *requests.SalesPricingConditionValidity) (*sap_api_output_formatter.SalesPricingConditionValidity, error) {
	body, err := json.Marshal(salesPricingConditionValidity)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_SALES_PRICING_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToSalesPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesPricingConditionRecord(salesPricingConditionRecord *requests.SalesPricingConditionRecord) {
	url := fmt.Sprintf("A_SlsPrcgCndnRecdValidity('%s')/to_SalesPricingConditionRecord", salesPricingConditionRecord.ConditionRecord)
	outputDataToConditionRecord, err := c.callSalesPricingSrvAPIRequirementToConditionRecord(url, salesPricingConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataToConditionRecord)
}

func (c *SAPAPICaller) callSalesPricingSrvAPIRequirementToConditionRecord(api string, toConditionRecord *requests.SalesPricingConditionRecord) (*sap_api_output_formatter.SalesPricingConditionRecord, error) {
	body, err := json.Marshal(toConditionRecord)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_SALES_PRICING_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToSalesPricingConditionRecord(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
