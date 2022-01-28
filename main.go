package main

import (
	sap_api_caller "sap-api-integrations-planned-independent-requirement-reads/SAP_API_Caller"
	"sap-api-integrations-planned-independent-requirement-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Planned_Independent_Requirement_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header",
		}
	}

	caller.AsyncGetPlannedIndependentRequirement(
		inoutSDC.PlannedIndependentRequirement.Product,
		inoutSDC.PlannedIndependentRequirement.Plant,
		inoutSDC.PlannedIndependentRequirement.MRPArea,
		inoutSDC.PlannedIndependentRequirement.PlndIndepRqmtType,
		inoutSDC.PlannedIndependentRequirement.PlndIndepRqmtVersion,
		inoutSDC.PlannedIndependentRequirement.RequirementPlan,
		inoutSDC.PlannedIndependentRequirement.RequirementSegment,
		accepter,
	)
}
