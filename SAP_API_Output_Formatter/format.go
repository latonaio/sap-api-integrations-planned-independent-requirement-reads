package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-planned-independent-requirement-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToPlannedIndependentRequirement(raw []byte, l *logger.Logger) *PlannedIndependentRequirement {
	pm := &responses.PlannedIndependentRequirement{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PlannedIndependentRequirement{
		Product                      data.Product,
		Plant                        data.Plant,
		MRPArea                      data.MRPArea,
		PlndIndepRqmtType            data.PlndIndepRqmtType,
		PlndIndepRqmtVersion         data.PlndIndepRqmtVersion,
		RequirementPlan              data.RequirementPlan,
		RequirementSegment           data.RequirementSegment,
		PlndIndepRqmtInternalID      data.PlndIndepRqmtInternalID,
		UnitOfMeasure                data.UnitOfMeasure,
		PlndIndepRqmtPeriod          data.PlndIndepRqmtPeriod,
		PlndIndepRqmtPeriodStartDate data.PlndIndepRqmtPeriodStartDate,
		PlannedQuantity              data.PlannedQuantity,
		LastChangeDate               data.LastChangeDate,
	}
}
