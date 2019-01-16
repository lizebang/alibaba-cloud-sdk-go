package vod

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

// DescribeVodDomainRealTimeReqHitRateData invokes the vod.DescribeVodDomainRealTimeReqHitRateData API synchronously
// api document: https://help.aliyun.com/api/vod/describevoddomainrealtimereqhitratedata.html
func (client *Client) DescribeVodDomainRealTimeReqHitRateData(request *DescribeVodDomainRealTimeReqHitRateDataRequest) (response *DescribeVodDomainRealTimeReqHitRateDataResponse, err error) {
	response = CreateDescribeVodDomainRealTimeReqHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainRealTimeReqHitRateDataWithChan invokes the vod.DescribeVodDomainRealTimeReqHitRateData API asynchronously
// api document: https://help.aliyun.com/api/vod/describevoddomainrealtimereqhitratedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVodDomainRealTimeReqHitRateDataWithChan(request *DescribeVodDomainRealTimeReqHitRateDataRequest) (<-chan *DescribeVodDomainRealTimeReqHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainRealTimeReqHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainRealTimeReqHitRateData(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeVodDomainRealTimeReqHitRateDataWithCallback invokes the vod.DescribeVodDomainRealTimeReqHitRateData API asynchronously
// api document: https://help.aliyun.com/api/vod/describevoddomainrealtimereqhitratedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVodDomainRealTimeReqHitRateDataWithCallback(request *DescribeVodDomainRealTimeReqHitRateDataRequest, callback func(response *DescribeVodDomainRealTimeReqHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainRealTimeReqHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainRealTimeReqHitRateData(request)
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

// DescribeVodDomainRealTimeReqHitRateDataRequest is the request struct for api DescribeVodDomainRealTimeReqHitRateData
type DescribeVodDomainRealTimeReqHitRateDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVodDomainRealTimeReqHitRateDataResponse is the response struct for api DescribeVodDomainRealTimeReqHitRateData
type DescribeVodDomainRealTimeReqHitRateDataResponse struct {
	*responses.BaseResponse
	RequestId string                                        `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeVodDomainRealTimeReqHitRateData `json:"Data" xml:"Data"`
}

// CreateDescribeVodDomainRealTimeReqHitRateDataRequest creates a request to invoke DescribeVodDomainRealTimeReqHitRateData API
func CreateDescribeVodDomainRealTimeReqHitRateDataRequest() (request *DescribeVodDomainRealTimeReqHitRateDataRequest) {
	request = &DescribeVodDomainRealTimeReqHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainRealTimeReqHitRateData", "vod", "openAPI")
	return
}

// CreateDescribeVodDomainRealTimeReqHitRateDataResponse creates a response to parse from DescribeVodDomainRealTimeReqHitRateData response
func CreateDescribeVodDomainRealTimeReqHitRateDataResponse() (response *DescribeVodDomainRealTimeReqHitRateDataResponse) {
	response = &DescribeVodDomainRealTimeReqHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
