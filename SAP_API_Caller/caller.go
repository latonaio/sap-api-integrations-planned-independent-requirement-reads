package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}


func (c *SAPAPICaller) AsyncGetPlannedIndependentRequirement(Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		c.PlannedIndependentRequirement(Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) PlannedIndependentRequirement(Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment string) {
	res, err := c.callPlannedIndependentRequirementSrvAPIRequirement("PlannedIndepRqmt(Product='{Product}',Plant='{Plant}',MRPArea='{MRPArea}',PlndIndepRqmtType='{PlndIndepRqmtType}',PlndIndepRqmtVersion='{PlndIndepRqmtVersion}',RequirementPlan='{RequirementPlan}',RequirementSegment='{RequirementSegment}')/to_PlndIndepRqmtItem", Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callPlannedIndependentRequirementSrvAPIRequirement(api, Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PLND_INDEP_RQMT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment")
	params.Add("$filter", fmt.Sprintf("Product eq '%s' and Plant eq '%s' and MRPArea eq '%s' and PlndIndepRqmtType eq '%s' and PlndIndepRqmtVersion eq '%s' and RequirementPlan eq '%s' and RequirementSegment eq '%s'", Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}