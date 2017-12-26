
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

func (client *Client) DescribeLiveSpecificDomainMapping(request *DescribeLiveSpecificDomainMappingRequest) (response *DescribeLiveSpecificDomainMappingResponse, err error) {
response = CreateDescribeLiveSpecificDomainMappingResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeLiveSpecificDomainMappingWithChan(request *DescribeLiveSpecificDomainMappingRequest) (<-chan *DescribeLiveSpecificDomainMappingResponse, <-chan error) {
responseChan := make(chan *DescribeLiveSpecificDomainMappingResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeLiveSpecificDomainMapping(request)
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

func (client *Client) DescribeLiveSpecificDomainMappingWithCallback(request *DescribeLiveSpecificDomainMappingRequest, callback func(response *DescribeLiveSpecificDomainMappingResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeLiveSpecificDomainMappingResponse
var err error
defer close(result)
response, err = client.DescribeLiveSpecificDomainMapping(request)
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

type DescribeLiveSpecificDomainMappingRequest struct {
*requests.RpcRequest
                PullDomain  string `position:"Query" name:"PullDomain"`
                PushDomain  string `position:"Query" name:"PushDomain"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
}


type DescribeLiveSpecificDomainMappingResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                DomainMappingModels struct {
                    DomainMappingModel []struct {
            PushDomain     string `json:"PushDomain" xml:"PushDomain"`
            PullDomain     string `json:"PullDomain" xml:"PullDomain"`
                    }   `json:"DomainMappingModel" xml:"DomainMappingModel"`
                } `json:"DomainMappingModels" xml:"DomainMappingModels"`
}

func CreateDescribeLiveSpecificDomainMappingRequest() (request *DescribeLiveSpecificDomainMappingRequest) {
request = &DescribeLiveSpecificDomainMappingRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveSpecificDomainMapping", "", "")
return
}

func CreateDescribeLiveSpecificDomainMappingResponse() (response *DescribeLiveSpecificDomainMappingResponse) {
response = &DescribeLiveSpecificDomainMappingResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

