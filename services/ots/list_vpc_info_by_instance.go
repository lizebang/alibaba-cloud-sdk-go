package ots

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

// ListVpcInfoByInstance invokes the ots.ListVpcInfoByInstance API synchronously
// api document: https://help.aliyun.com/api/ots/listvpcinfobyinstance.html
func (client *Client) ListVpcInfoByInstance(request *ListVpcInfoByInstanceRequest) (response *ListVpcInfoByInstanceResponse, err error) {
	response = CreateListVpcInfoByInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ListVpcInfoByInstanceWithChan invokes the ots.ListVpcInfoByInstance API asynchronously
// api document: https://help.aliyun.com/api/ots/listvpcinfobyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVpcInfoByInstanceWithChan(request *ListVpcInfoByInstanceRequest) (<-chan *ListVpcInfoByInstanceResponse, <-chan error) {
	responseChan := make(chan *ListVpcInfoByInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVpcInfoByInstance(request)
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

// ListVpcInfoByInstanceWithCallback invokes the ots.ListVpcInfoByInstance API asynchronously
// api document: https://help.aliyun.com/api/ots/listvpcinfobyinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVpcInfoByInstanceWithCallback(request *ListVpcInfoByInstanceRequest, callback func(response *ListVpcInfoByInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVpcInfoByInstanceResponse
		var err error
		defer close(result)
		response, err = client.ListVpcInfoByInstance(request)
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

// ListVpcInfoByInstanceRequest is the request struct for api ListVpcInfoByInstance
type ListVpcInfoByInstanceRequest struct {
	*requests.RpcRequest
	AccessKeyId     string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceName    string           `position:"Query" name:"InstanceName"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
}

// ListVpcInfoByInstanceResponse is the response struct for api ListVpcInfoByInstance
type ListVpcInfoByInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	InstanceName string                          `json:"InstanceName" xml:"InstanceName"`
	TotalCount   int                             `json:"TotalCount" xml:"TotalCount"`
	PageNum      int                             `json:"PageNum" xml:"PageNum"`
	PageSize     int                             `json:"PageSize" xml:"PageSize"`
	VpcInfos     VpcInfosInListVpcInfoByInstance `json:"VpcInfos" xml:"VpcInfos"`
}

// CreateListVpcInfoByInstanceRequest creates a request to invoke ListVpcInfoByInstance API
func CreateListVpcInfoByInstanceRequest() (request *ListVpcInfoByInstanceRequest) {
	request = &ListVpcInfoByInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ots", "2016-06-20", "ListVpcInfoByInstance", "ots", "openAPI")
	return
}

// CreateListVpcInfoByInstanceResponse creates a response to parse from ListVpcInfoByInstance response
func CreateListVpcInfoByInstanceResponse() (response *ListVpcInfoByInstanceResponse) {
	response = &ListVpcInfoByInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}