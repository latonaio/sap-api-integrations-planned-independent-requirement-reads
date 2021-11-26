package sap_api_caller

type PlannedIndependentRequirementReads struct {
	ConnectionKey                        string `json:"connection_key"`
	Result                               bool   `json:"result"`
	RedisKey                             string `json:"redis_key"`
	Filepath                             string `json:"filepath"`
	APISchema                            string `json:"api_schema"`
	MaterialCode                         string `json:"material_code"`
	Plant                                string `json:"plant"`
	Deleted                              string `json:"deleted"`
}

type PlannedIndependentRequirement struct {
    PlannedIndependentRequirement string   `json:"PlannedIndependentRequirement"` 
    Product                       string   `json:"Product"`
    Plant                         string   `json:"Plant"`
    MRPArea                       string   `json:"MRPArea"`
    PlndIndepRqmtType             string   `json:"PlndIndepRqmtType"` 
    PlndIndepRqmtVersion          string   `json:"PlndIndepRqmtVersion"`
    RequirementPlan               string   `json:"RequirementPlan"`
    RequirementSegment            string   `json:"RequirementSegment"`
    PlndIndepRqmtInternalID       int      `json:"PlndIndepRqmtInternalID"`
    UnitOfMeasure                 string   `json:"UnitOfMeasure"`
    PlndIndepRqmtPeriod           string   `json:"PlndIndepRqmtPeriod"`
    PlndIndepRqmtPeriodStartDate  string   `json:"PlndIndepRqmtPeriodStartDate"`
    PlannedQuantity               float64  `json:"PlannedQuantity"`
    LastChangeDate                string   `json:"LastChangeDate"`
}
