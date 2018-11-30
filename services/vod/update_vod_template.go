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

// UpdateVodTemplate invokes the vod.UpdateVodTemplate API synchronously
// api document: https://help.aliyun.com/api/vod/updatevodtemplate.html
func (client *Client) UpdateVodTemplate(request *UpdateVodTemplateRequest) (response *UpdateVodTemplateResponse, err error) {
	response = CreateUpdateVodTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateVodTemplateWithChan invokes the vod.UpdateVodTemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/updatevodtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateVodTemplateWithChan(request *UpdateVodTemplateRequest) (<-chan *UpdateVodTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateVodTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateVodTemplate(request)
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

// UpdateVodTemplateWithCallback invokes the vod.UpdateVodTemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/updatevodtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateVodTemplateWithCallback(request *UpdateVodTemplateRequest, callback func(response *UpdateVodTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateVodTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateVodTemplate(request)
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

// UpdateVodTemplateRequest is the request struct for api UpdateVodTemplate
type UpdateVodTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TemplateConfig       string           `position:"Query" name:"TemplateConfig"`
	Name                 string           `position:"Query" name:"Name"`
	VodTemplateId        string           `position:"Query" name:"VodTemplateId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateVodTemplateResponse is the response struct for api UpdateVodTemplate
type UpdateVodTemplateResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	VodTemplateId string `json:"VodTemplateId" xml:"VodTemplateId"`
}

// CreateUpdateVodTemplateRequest creates a request to invoke UpdateVodTemplate API
func CreateUpdateVodTemplateRequest() (request *UpdateVodTemplateRequest) {
	request = &UpdateVodTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UpdateVodTemplate", "", "")
	return
}

// CreateUpdateVodTemplateResponse creates a response to parse from UpdateVodTemplate response
func CreateUpdateVodTemplateResponse() (response *UpdateVodTemplateResponse) {
	response = &UpdateVodTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
