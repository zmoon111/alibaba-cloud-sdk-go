package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeDomainHttpsData(request *DescribeDomainHttpsDataRequest) (response *DescribeDomainHttpsDataResponse, err error) {
	response = CreateDescribeDomainHttpsDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainHttpsDataWithChan(request *DescribeDomainHttpsDataRequest) (<-chan *DescribeDomainHttpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainHttpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainHttpsData(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeDomainHttpsDataWithCallback(request *DescribeDomainHttpsDataRequest, callback func(response *DescribeDomainHttpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainHttpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainHttpsData(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeDomainHttpsDataRequest struct {
	*requests.RpcRequest
	DomainType    string `position:"Query" name:"DomainType"`
	FixTimeGap    string `position:"Query" name:"FixTimeGap"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainNames   string `position:"Query" name:"DomainNames"`
	Action        string `position:"Query" name:"Action"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	Cls           string `position:"Query" name:"Cls"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
}

type DescribeDomainHttpsDataResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId"`
	DomainNames          string `json:"DomainNames"`
	DataInterval         string `json:"DataInterval"`
	HttpsStatisticsInfos []struct {
		Time               string  `json:"Time"`
		L1HttpsBps         float64 `json:"L1HttpsBps"`
		L1HttpsInnerBps    float64 `json:"L1HttpsInnerBps"`
		L1HttpsOutBps      float64 `json:"L1HttpsOutBps"`
		L1HttpsQps         int64   `json:"L1HttpsQps"`
		L1HttpsInnerQps    int64   `json:"L1HttpsInnerQps"`
		L1HttpsOutQps      int64   `json:"L1HttpsOutQps"`
		L1HttpsTtraf       int64   `json:"L1HttpsTtraf"`
		HttpsSrcBps        int64   `json:"HttpsSrcBps"`
		HttpsSrcTraf       int64   `json:"HttpsSrcTraf"`
		L1HttpsInnerTraf   int64   `json:"L1HttpsInnerTraf"`
		L1HttpsOutTraf     int64   `json:"L1HttpsOutTraf"`
		HttpsByteHitRate   float64 `json:"HttpsByteHitRate"`
		HttpsReqHitRate    float64 `json:"HttpsReqHitRate"`
		L1HttpsHitRate     float64 `json:"L1HttpsHitRate"`
		L1HttpsInnerAcc    float64 `json:"L1HttpsInner_acc"`
		L1HttpsOutAcc      float64 `json:"L1HttpsOut_acc"`
		L1HttpsTacc        float64 `json:"L1HttpsTacc"`
		L1DyHttpsBps       float64 `json:"L1DyHttpsBps"`
		L1DyHttpsInnerBps  float64 `json:"L1DyHttpsInnerBps"`
		L1DyHttpsOutBps    float64 `json:"L1DyHttpsOutBps"`
		L1StHttpsBps       float64 `json:"L1StHttpsBps"`
		L1StHttpsInnerBps  float64 `json:"L1StHttpsInnerBps"`
		L1StHttpsOutBps    float64 `json:"L1StHttpsOutBps"`
		L1DyHttpsTraf      float64 `json:"l1DyHttpsTraf"`
		L1DyHttpsInnerTraf float64 `json:"L1DyHttpsInnerTraf"`
		L1DyHttpsOutTraf   float64 `json:"L1DyHttpsOutTraf"`
		L1StHttpsTraf      float64 `json:"L1StHttpsTraf"`
		L1StHttpsInnerTraf float64 `json:"L1StHttpsInnerTraf"`
		L1StHttpsOutTraf   float64 `json:"L1StHttpsOutTraf"`
		L1DyHttpsQps       float64 `json:"L1DyHttpsQps"`
		L1DyHttpsInnerQps  float64 `json:"L1DyHttpsInnerQps"`
		L1DyHttpsOutQps    float64 `json:"L1DyHttpsOutQps"`
		L1StHttpsQps       float64 `json:"L1StHttpsQps"`
		L1StHttpsInnerQps  float64 `json:"L1StHttpsInnerQps"`
		L1StHttpsOutQps    float64 `json:"L1StHttpsOutQps"`
		L1DyHttpsAcc       float64 `json:"L1DyHttpsAcc"`
		L1DyHttpsInnerAcc  float64 `json:"L1DyHttpsInnerAcc"`
		L1DyHttpsOutAcc    float64 `json:"L1DyHttpsOutAcc"`
		L1StHttpsAcc       float64 `json:"L1StHttpsAcc"`
		L1StHttpsInnerAcc  float64 `json:"L1StHttpsInnerAcc"`
		L1StHttpsOutAcc    float64 `json:"L1StHttpsOutAcc"`
	} `json:"HttpsStatisticsInfos"`
}

func CreateDescribeDomainHttpsDataRequest() (request *DescribeDomainHttpsDataRequest) {
	request = &DescribeDomainHttpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainHttpsData", "", "")
	return
}

func CreateDescribeDomainHttpsDataResponse() (response *DescribeDomainHttpsDataResponse) {
	response = &DescribeDomainHttpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}