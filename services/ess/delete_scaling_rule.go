
package ess

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

func (client *Client) DeleteScalingRule(request *DeleteScalingRuleRequest) (response *DeleteScalingRuleResponse, err error) {
response = CreateDeleteScalingRuleResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DeleteScalingRuleWithChan(request *DeleteScalingRuleRequest) (<-chan *DeleteScalingRuleResponse, <-chan error) {
responseChan := make(chan *DeleteScalingRuleResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteScalingRule(request)
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

func (client *Client) DeleteScalingRuleWithCallback(request *DeleteScalingRuleRequest, callback func(response *DeleteScalingRuleResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteScalingRuleResponse
var err error
defer close(result)
response, err = client.DeleteScalingRule(request)
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

type DeleteScalingRuleRequest struct {
*requests.RpcRequest
                ScalingRuleId  string `position:"Query" name:"ScalingRuleId"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DeleteScalingRuleResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteScalingRuleRequest() (request *DeleteScalingRuleRequest) {
request = &DeleteScalingRuleRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingRule", "", "")
return
}

func CreateDeleteScalingRuleResponse() (response *DeleteScalingRuleResponse) {
response = &DeleteScalingRuleResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

