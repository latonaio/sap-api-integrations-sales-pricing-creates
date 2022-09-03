package main

import (
	sap_api_caller "sap-api-integrations-sales-pricing-creates/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-sales-pricing-creates/SAP_API_Input_Reader"
	"sap-api-integrations-sales-pricing-creates/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	pc := sap_api_post_header_setup.NewSAPPostClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		pc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Sales_Pricing_sample.json")
	accepter := getAccepter(inputSDC)
	salesPricingConditionValidity := inputSDC.ConvertToSalesPricingConditionValidity()
	salesPricingConditionRecord := inputSDC.ConvertToSalesPricingConditionRecord()

	caller.AsyncPostSalesPricing(
		salesPricingConditionValidity,
		salesPricingConditionRecord,
		accepter,
	)
}

func getAccepter(sdc sap_api_input_reader.SDC) []string {
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"SalesPricingConditionValidity", "SalesPricingConditionRecord",
		}
	}
	return accepter
}
