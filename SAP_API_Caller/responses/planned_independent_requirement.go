package responses

type PlannedIndependentRequirement struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
