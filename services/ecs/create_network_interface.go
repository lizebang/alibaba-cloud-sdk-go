
package ecs

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

func (client *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
response = CreateCreateNetworkInterfaceResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) CreateNetworkInterfaceWithChan(request *CreateNetworkInterfaceRequest) (<-chan *CreateNetworkInterfaceResponse, <-chan error) {
responseChan := make(chan *CreateNetworkInterfaceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.CreateNetworkInterface(request)
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

func (client *Client) CreateNetworkInterfaceWithCallback(request *CreateNetworkInterfaceRequest, callback func(response *CreateNetworkInterfaceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *CreateNetworkInterfaceResponse
var err error
defer close(result)
response, err = client.CreateNetworkInterface(request)
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

type CreateNetworkInterfaceRequest struct {
*requests.RpcRequest
                NetworkInterfaceName  string `position:"Query" name:"NetworkInterfaceName"`
                PrimaryIpAddress  string `position:"Query" name:"PrimaryIpAddress"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                Description  string `position:"Query" name:"Description"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                VSwitchId  string `position:"Query" name:"VSwitchId"`
                ClientToken  string `position:"Query" name:"ClientToken"`
                SecurityGroupId  string `position:"Query" name:"SecurityGroupId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type CreateNetworkInterfaceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            NetworkInterfaceId     string `json:"NetworkInterfaceId" xml:"NetworkInterfaceId"`
}

func CreateCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
request = &CreateNetworkInterfaceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Ecs", "2014-05-26", "CreateNetworkInterface", "", "")
return
}

func CreateCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
response = &CreateNetworkInterfaceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

