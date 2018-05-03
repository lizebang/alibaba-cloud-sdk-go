package green

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

// TextScan invokes the green.TextScan API synchronously
// api document: https://help.aliyun.com/api/green/textscan.html
func (client *Client) TextScan(request *TextScanRequest) (response *TextScanResponse, err error) {
	response = CreateTextScanResponse()
	err = client.DoAction(request, response)
	return
}

// TextScanWithChan invokes the green.TextScan API asynchronously
// api document: https://help.aliyun.com/api/green/textscan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TextScanWithChan(request *TextScanRequest) (<-chan *TextScanResponse, <-chan error) {
	responseChan := make(chan *TextScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TextScan(request)
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

// TextScanWithCallback invokes the green.TextScan API asynchronously
// api document: https://help.aliyun.com/api/green/textscan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TextScanWithCallback(request *TextScanRequest, callback func(response *TextScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TextScanResponse
		var err error
		defer close(result)
		response, err = client.TextScan(request)
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

// TextScanRequest is the request struct for api TextScan
type TextScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// TextScanResponse is the response struct for api TextScan
type TextScanResponse struct {
	*responses.BaseResponse
}

// CreateTextScanRequest creates a request to invoke TextScan API
func CreateTextScanRequest() (request *TextScanRequest) {
	request = &TextScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-25", "TextScan", "/green/text/scan", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTextScanResponse creates a response to parse from TextScan response
func CreateTextScanResponse() (response *TextScanResponse) {
	response = &TextScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}