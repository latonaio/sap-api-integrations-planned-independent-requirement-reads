package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-planned-independent-requirement-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
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

func (c *SAPAPICaller) AsyncGetPlannedIndependentRequirement(product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	func() {
		c.Header(product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string) {
	headerData, err := c.callPlannedIndependentRequirementSrvAPIRequirementHeader("PlannedIndepRqmt", product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	itemData, err := c.callToPlndIndepRqmtItem(headerData[0].ToPlndIndepRqmtItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)
}

func (c *SAPAPICaller) callPlannedIndependentRequirementSrvAPIRequirementHeader(api, product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PLND_INDEP_RQMT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPlndIndepRqmtItem(url string) (*sap_api_output_formatter.ToPlndIndepRqmtItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPlndIndepRqmtItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Product eq '%s' and Plant eq '%s' and MRPArea eq '%s' and PlndIndepRqmtType eq '%s' and PlndIndepRqmtVersion eq '%s' and RequirementPlan eq '%s' and RequirementSegment eq '%s'", product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment))
	req.URL.RawQuery = params.Encode()
}
