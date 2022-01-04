# sap-api-integrations-planned-independent-requirement-reads  
sap-api-integrations-planned-independent-requirement-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 計画独立所要量 データを取得するマイクロサービスです。    
sap-api-integrations-planned-independent-requirement-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-planned-independent-requirement-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PLND_INDEP_RQMT_SRV_0001/overview

## 動作環境
sap-api-integrations-planned-independent-requirement-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-planned-independent-requirement-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、 "PlannedIndependentRequirement"が指定されています。    
  
```
	"api_schema": "/PlannedIndepRqmt",
	"accepter": ["Header"],
	"material_code": "FG-FL-MV-V00",
	"plant": "1710",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "/PlannedIndepRqmt",
	"accepter": ["All"],
	"material_code": "FG-FL-MV-V00",
	"plant": "1710",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetProductMaster(product, plant, mrpArea, valuationArea, productSalesOrg, productDistributionChnl string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "PlannedIndependentRequirement":
			func() {
				c.PlannedIndependentRequirement(Product, Plant, MRPArea, PlndIndepRqmtType, PlndIndepRqmtVersion, RequirementPlan, RequirementSegment)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 計画独立所要量 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"Update_mc" ～ "PlndIndepRqmtLastChgdDateTime" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-planned-independent-requirement-reads/SAP_API_Caller/caller.go#L60",
	"function": "sap-api-integrations-planned-independent-requirement-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"Update_mc": true,
			"Product": "FG-FL-MV-V00",
			"Plant": "1710",
			"MRPArea": "1710",
			"PlndIndepRqmtType": "VSF",
			"PlndIndepRqmtVersion": "00",
			"RequirementPlan": "",
			"RequirementSegment": "",
			"PlndIndepRqmtPeriod": "201711",
			"PeriodType": "M",
			"PlndIndepRqmtPeriodStartDate": "/Date(1509494400000)/",
			"PlndIndepRqmtInternalID": "101",
			"WorkingDayDate": "/Date(1509494400000)/",
			"PlannedQuantity": "0",
			"WithdrawalQuantity": "0",
			"UnitOfMeasure": "PC",
			"LastChangedByUser": "CB9980000078",
			"LastChangeDate": "/Date(1511827200000)/",
			"PlndIndepRqmtLastChgdDateTime": ""
		},
		{
			"Update_mc": true,
			"Product": "FG-FL-MV-V00",
			"Plant": "1710",
			"MRPArea": "1710",
			"PlndIndepRqmtType": "VSF",
			"PlndIndepRqmtVersion": "00",
			"RequirementPlan": "",
			"RequirementSegment": "",
			"PlndIndepRqmtPeriod": "201712",
			"PeriodType": "M",
			"PlndIndepRqmtPeriodStartDate": "/Date(1512086400000)/",
			"PlndIndepRqmtInternalID": "101",
			"WorkingDayDate": "/Date(1512086400000)/",
			"PlannedQuantity": "0",
			"WithdrawalQuantity": "0",
			"UnitOfMeasure": "PC",
			"LastChangedByUser": "CB9980000078",
			"LastChangeDate": "/Date(1511827200000)/",
			"PlndIndepRqmtLastChgdDateTime": ""
		}
	],
	"time": "2021-12-17T17:22:46.68264+09:00"
}


```


