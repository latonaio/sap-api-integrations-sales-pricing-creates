package requests

type SlsPrcgCndnRecdSuplmnt struct {
	ConditionRecord              string  `json:"ConditionRecord"`
	ConditionSequentialNumber    string  `json:"ConditionSequentialNumber"`
	ConditionTable               string  `json:"ConditionTable"`
	ConditionApplication         string  `json:"ConditionApplication"`
	ConditionType                string  `json:"ConditionType"`
	ConditionValidityEndDate     string  `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate   string  `json:"ConditionValidityStartDate"`
	CreatedByUser                *string `json:"CreatedByUser"`
	CreationDate                 *string `json:"CreationDate"`
	ConditionTextID              *string `json:"ConditionTextID"`
	PricingScaleType             *string `json:"PricingScaleType"`
	PricingScaleBasis            *string `json:"PricingScaleBasis"`
	ConditionScaleQuantity       *string `json:"ConditionScaleQuantity"`
	ConditionScaleQuantityUnit   *string `json:"ConditionScaleQuantityUnit"`
	ConditionScaleAmount         *string `json:"ConditionScaleAmount"`
	ConditionScaleAmountCurrency *string `json:"ConditionScaleAmountCurrency"`
	ConditionCalculationType     *string `json:"ConditionCalculationType"`
	ConditionRateValue           *string `json:"ConditionRateValue"`
	ConditionRateValueUnit       *string `json:"ConditionRateValueUnit"`
	ConditionRateRatio           *string `json:"ConditionRateRatio"`
	ConditionRateRatioUnit       *string `json:"ConditionRateRatioUnit"`
	ConditionRateAmount          *string `json:"ConditionRateAmount"`
	ConditionCurrency            *string `json:"ConditionCurrency"`
	ConditionQuantity            *string `json:"ConditionQuantity"`
	ConditionQuantityUnit        *string `json:"ConditionQuantityUnit"`
	ConditionToBaseQtyNmrtr      *string `json:"ConditionToBaseQtyNmrtr"`
	ConditionToBaseQtyDnmntr     *string `json:"ConditionToBaseQtyDnmntr"`
	BaseUnit                     *string `json:"BaseUnit"`
	ConditionLowerLimit          *string `json:"ConditionLowerLimit"`
	ConditionLowerLimitAmount    *string `json:"ConditionLowerLimitAmount"`
	ConditionLowerLimitRatio     *string `json:"ConditionLowerLimitRatio"`
	ConditionUpperLimit          *string `json:"ConditionUpperLimit"`
	ConditionUpperLimitAmount    *string `json:"ConditionUpperLimitAmount"`
	ConditionUpperLimitRatio     *string `json:"ConditionUpperLimitRatio"`
	ConditionAlternativeCurrency *string `json:"ConditionAlternativeCurrency"`
	ConditionExclusion           *string `json:"ConditionExclusion"`
	ConditionIsDeleted           *bool   `json:"ConditionIsDeleted"`
	AdditionalValueDays          *string `json:"AdditionalValueDays"`
	FixedValueDate               *string `json:"FixedValueDate"`
	PaymentTerms                 *string `json:"PaymentTerms"`
	CndnMaxNumberOfSalesOrders   *string `json:"CndnMaxNumberOfSalesOrders"`
	MinimumConditionBasisValue   *string `json:"MinimumConditionBasisValue"`
	MaximumConditionBasisValue   *string `json:"MaximumConditionBasisValue"`
	MaximumConditionAmount       *string `json:"MaximumConditionAmount"`
	IncrementalScale             *string `json:"IncrementalScale"`
	PricingScaleLine             *string `json:"PricingScaleLine"`
	ConditionReleaseStatus       *string `json:"ConditionReleaseStatus"`
	ETag                         *string `json:"ETag"`
}