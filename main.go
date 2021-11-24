package main

import (
    sap_api_caller "sap-api-integrations-planned-independent-requirement-reads/SAP_API_Caller"
    "sap-api-integrations-planned-independent-requirement-reads/file_reader"

    "github.com/latonaio/golang-logging-library/logger"
)

func main() {
    l := logger.NewLogger()
    fr := file_reader.NewFileReader()
    inoutSDC := fr.ReadSDC("./Inputs//SDC_Planned_Independent_Requirement_sample.json")
    caller := sap_api_caller.NewSAPAPICaller(
        "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
    )

    caller.AsyncGetPlannedIndependentRequirement(
        inoutSDC.Product.Plant.MRPArea.PlndIndepRqmtType.PlndIndepRqmtVersion.RequirementPlan.RequirementSegment,
    )
}