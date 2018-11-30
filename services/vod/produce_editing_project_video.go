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

// ProduceEditingProjectVideo invokes the vod.ProduceEditingProjectVideo API synchronously
// api document: https://help.aliyun.com/api/vod/produceeditingprojectvideo.html
func (client *Client) ProduceEditingProjectVideo(request *ProduceEditingProjectVideoRequest) (response *ProduceEditingProjectVideoResponse, err error) {
	response = CreateProduceEditingProjectVideoResponse()
	err = client.DoAction(request, response)
	return
}

// ProduceEditingProjectVideoWithChan invokes the vod.ProduceEditingProjectVideo API asynchronously
// api document: https://help.aliyun.com/api/vod/produceeditingprojectvideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProduceEditingProjectVideoWithChan(request *ProduceEditingProjectVideoRequest) (<-chan *ProduceEditingProjectVideoResponse, <-chan error) {
	responseChan := make(chan *ProduceEditingProjectVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProduceEditingProjectVideo(request)
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

// ProduceEditingProjectVideoWithCallback invokes the vod.ProduceEditingProjectVideo API asynchronously
// api document: https://help.aliyun.com/api/vod/produceeditingprojectvideo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProduceEditingProjectVideoWithCallback(request *ProduceEditingProjectVideoRequest, callback func(response *ProduceEditingProjectVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProduceEditingProjectVideoResponse
		var err error
		defer close(result)
		response, err = client.ProduceEditingProjectVideo(request)
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

// ProduceEditingProjectVideoRequest is the request struct for api ProduceEditingProjectVideo
type ProduceEditingProjectVideoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MediaMetadata        string           `position:"Query" name:"MediaMetadata"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Title                string           `position:"Query" name:"Title"`
	CoverURL             string           `position:"Query" name:"CoverURL"`
	UserData             string           `position:"Query" name:"UserData"`
	Timeline             string           `position:"Query" name:"Timeline"`
	ProduceConfig        string           `position:"Query" name:"ProduceConfig"`
	ProjectId            string           `position:"Query" name:"ProjectId"`
}

// ProduceEditingProjectVideoResponse is the response struct for api ProduceEditingProjectVideo
type ProduceEditingProjectVideoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MediaId   string `json:"MediaId" xml:"MediaId"`
	ProjectId string `json:"ProjectId" xml:"ProjectId"`
}

// CreateProduceEditingProjectVideoRequest creates a request to invoke ProduceEditingProjectVideo API
func CreateProduceEditingProjectVideoRequest() (request *ProduceEditingProjectVideoRequest) {
	request = &ProduceEditingProjectVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ProduceEditingProjectVideo", "", "")
	return
}

// CreateProduceEditingProjectVideoResponse creates a response to parse from ProduceEditingProjectVideo response
func CreateProduceEditingProjectVideoResponse() (response *ProduceEditingProjectVideoResponse) {
	response = &ProduceEditingProjectVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
