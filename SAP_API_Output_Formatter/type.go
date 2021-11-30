package sap_api_output_formatter

type PlannedIndependentRequirementReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant"`
	Deleted       bool   `json:"deleted"`
}

type PlannedIndependentRequirement struct {
	PlannedIndependentRequirement string `json:"PlannedIndependentRequirement"`
	Product                       string `json:"Product"`
	Plant                         string `json:"Plant"`
	MRPArea                       string `json:"MRPArea"`
	PlndIndepRqmtType             string `json:"PlndIndepRqmtType"`
	PlndIndepRqmtVersion          string `json:"PlndIndepRqmtVersion"`
	RequirementPlan               string `json:"RequirementPlan"`
	RequirementSegment            string `json:"RequirementSegment"`
	PlndIndepRqmtInternalID       string `json:"PlndIndepRqmtInternalID"`
	UnitOfMeasure                 string `json:"UnitOfMeasure"`
	PlndIndepRqmtPeriod           string `json:"PlndIndepRqmtPeriod"`
	PlndIndepRqmtPeriodStartDate  string `json:"PlndIndepRqmtPeriodStartDate"`
	PlannedQuantity               string `json:"PlannedQuantity"`
	LastChangeDate                string `json:"LastChangeDate"`
}
