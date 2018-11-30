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

// ListVodTemplate invokes the vod.ListVodTemplate API synchronously
// api document: https://help.aliyun.com/api/vod/listvodtemplate.html
func (client *Client) ListVodTemplate(request *ListVodTemplateRequest) (response *ListVodTemplateResponse, err error) {
	response = CreateListVodTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ListVodTemplateWithChan invokes the vod.ListVodTemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/listvodtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVodTemplateWithChan(request *ListVodTemplateRequest) (<-chan *ListVodTemplateResponse, <-chan error) {
	responseChan := make(chan *ListVodTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVodTemplate(request)
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

// ListVodTemplateWithCallback invokes the vod.ListVodTemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/listvodtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVodTemplateWithCallback(request *ListVodTemplateRequest, callback func(response *ListVodTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVodTemplateResponse
		var err error
		defer close(result)
		response, err = client.ListVodTemplate(request)
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

// ListVodTemplateRequest is the request struct for api ListVodTemplate
type ListVodTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TemplateType         string           `position:"Query" name:"TemplateType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ListVodTemplateResponse is the response struct for api ListVodTemplate
type ListVodTemplateResponse struct {
	*responses.BaseResponse
	RequestId           string            `json:"RequestId" xml:"RequestId"`
	VodTemplateInfoList []VodTemplateInfo `json:"VodTemplateInfoList" xml:"VodTemplateInfoList"`
}

// CreateListVodTemplateRequest creates a request to invoke ListVodTemplate API
func CreateListVodTemplateRequest() (request *ListVodTemplateRequest) {
	request = &ListVodTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListVodTemplate", "", "")
	return
}

// CreateListVodTemplateResponse creates a response to parse from ListVodTemplate response
func CreateListVodTemplateResponse() (response *ListVodTemplateResponse) {
	response = &ListVodTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
