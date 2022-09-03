package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         string `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey                 string `json:"connection_key"`
	Result                        bool   `json:"result"`
	RedisKey                      string `json:"redis_key"`
	Filepath                      string `json:"filepath"`
	SalesPricingConditionValidity struct {
		ConditionRecord             string  `json:"ConditionRecord"`
		ConditionValidityEndDate    string  `json:"ConditionValidityEndDate"`
		ConditionValidityStartDate  *string `json:"ConditionValidityStartDate"`
		ConditionApplication        *string `json:"ConditionApplication"`
		ConditionType               *string `json:"ConditionType"`
		ConditionReleaseStatus      *string `json:"ConditionReleaseStatus"`
		SalesDocument               *string `json:"SalesDocument"`
		SalesDocumentItem           *string `json:"SalesDocumentItem"`
		ConditionContract           *string `json:"ConditionContract"`
		CustomerGroup               *string `json:"CustomerGroup"`
		CustomerPriceGroup          *string `json:"CustomerPriceGroup"`
		MaterialPricingGroup        *string `json:"MaterialPricingGroup"`
		SoldToParty                 *string `json:"SoldToParty"`
		BPForSoldToParty            *string `json:"BPForSoldToParty"`
		Customer                    *string `json:"Customer"`
		BPForCustomer               *string `json:"BPForCustomer"`
		PayerParty                  *string `json:"PayerParty"`
		BPForPayerParty             *string `json:"BPForPayerParty"`
		ShipToParty                 *string `json:"ShipToParty"`
		BPForShipToParty            *string `json:"BPForShipToParty"`
		Supplier                    *string `json:"Supplier"`
		BPForSupplier               *string `json:"BPForSupplier"`
		MaterialGroup               *string `json:"MaterialGroup"`
		Material                    *string `json:"Material"`
		PriceListType               *string `json:"PriceListType"`
		CustomerTaxClassification1  *string `json:"CustomerTaxClassification1"`
		ProductTaxClassification1   *string `json:"ProductTaxClassification1"`
		SDDocument                  *string `json:"SDDocument"`
		ReferenceSDDocument         *string `json:"ReferenceSDDocument"`
		ReferenceSDDocumentItem     *string `json:"ReferenceSDDocumentItem"`
		SalesOffice                 *string `json:"SalesOffice"`
		SalesGroup                  *string `json:"SalesGroup"`
		SalesOrganization           *string `json:"SalesOrganization"`
		DistributionChannel         *string `json:"DistributionChannel"`
		TransactionCurrency         *string `json:"TransactionCurrency"`
		ConditionProcessingStatus   *string `json:"ConditionProcessingStatus"`
		PricingDate                 *string `json:"PricingDate"`
		ConditionScaleBasisValue    *string `json:"ConditionScaleBasisValue"`
		TaxCode                     *string `json:"TaxCode"`
		ServiceDocument             *string `json:"ServiceDocument"`
		ServiceDocumentItem         *string `json:"ServiceDocumentItem"`
		CustomerConditionGroup      *string `json:"CustomerConditionGroup"`
		SalesPricingConditionRecord struct {
			ConditionRecord              string  `json:"ConditionRecord"`
			ConditionSequentialNumber    string  `json:"ConditionSequentialNumber"`
			ConditionTable               string  `json:"ConditionTable"`
			ConditionApplication         string  `json:"ConditionApplication"`
			ConditionType                string  `json:"ConditionType"`
			ConditionValidityEndDate     string  `json:"ConditionValidityEndDate"`
			ConditionValidityStartDate   *string `json:"ConditionValidityStartDate"`
			CreationDate                 *string `json:"CreationDate"`
			PricingScaleType             *string `json:"PricingScaleType"`
			PricingScaleBasis            *string `json:"PricingScaleBasis"`
			ConditionScaleQuantity       *string `json:"ConditionScaleQuantity"`
			ConditionScaleQuantityUnit   *string `json:"ConditionScaleQuantityUnit"`
			ConditionScaleAmount         *string `json:"ConditionScaleAmount"`
			ConditionScaleAmountCurrency *string `json:"ConditionScaleAmountCurrency"`
			ConditionCalculationType     *string `json:"ConditionCalculationType"`
			ConditionRateValue           *string `json:"ConditionRateValue"`
			ConditionRateValueUnit       *string `json:"ConditionRateValueUnit"`
			ConditionQuantity            *string `json:"ConditionQuantity"`
			ConditionQuantityUnit        *string `json:"ConditionQuantityUnit"`
			BaseUnit                     *string `json:"BaseUnit"`
			ConditionIsDeleted           *bool   `json:"ConditionIsDeleted"`
			PaymentTerms                 *string `json:"PaymentTerms"`
			IncrementalScale             *string `json:"IncrementalScale"`
			PricingScaleLine             *string `json:"PricingScaleLine"`
			ConditionReleaseStatus       *string `json:"ConditionReleaseStatus"`
		} `json:"SalesPricingConditionRecord"`
	} `json:"SalesPricingConditionValidity"`
	APISchema       string   `json:"api_schema"`
	Accepter        []string `json:"accepter"`
	ConditionRecord string   `json:"condition_record"`
	Deleted         bool     `json:"deleted"`
}
