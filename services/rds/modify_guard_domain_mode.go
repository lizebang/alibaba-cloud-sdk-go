
package rds

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

func (client *Client) ModifyGuardDomainMode(request *ModifyGuardDomainModeRequest) (response *ModifyGuardDomainModeResponse, err error) {
response = CreateModifyGuardDomainModeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ModifyGuardDomainModeWithChan(request *ModifyGuardDomainModeRequest) (<-chan *ModifyGuardDomainModeResponse, <-chan error) {
responseChan := make(chan *ModifyGuardDomainModeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyGuardDomainMode(request)
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

func (client *Client) ModifyGuardDomainModeWithCallback(request *ModifyGuardDomainModeRequest, callback func(response *ModifyGuardDomainModeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyGuardDomainModeResponse
var err error
defer close(result)
response, err = client.ModifyGuardDomainMode(request)
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

type ModifyGuardDomainModeRequest struct {
*requests.RpcRequest
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                ReplicaId  string `position:"Query" name:"ReplicaId"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
                DomainMode  string `position:"Query" name:"DomainMode"`
}


type ModifyGuardDomainModeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyGuardDomainModeRequest() (request *ModifyGuardDomainModeRequest) {
request = &ModifyGuardDomainModeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "ModifyGuardDomainMode", "", "")
return
}

func CreateModifyGuardDomainModeResponse() (response *ModifyGuardDomainModeResponse) {
response = &ModifyGuardDomainModeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

