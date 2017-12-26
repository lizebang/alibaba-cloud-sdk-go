
package dm

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

func (client *Client) DeleteReceiver(request *DeleteReceiverRequest) (response *DeleteReceiverResponse, err error) {
response = CreateDeleteReceiverResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DeleteReceiverWithChan(request *DeleteReceiverRequest) (<-chan *DeleteReceiverResponse, <-chan error) {
responseChan := make(chan *DeleteReceiverResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DeleteReceiver(request)
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

func (client *Client) DeleteReceiverWithCallback(request *DeleteReceiverRequest, callback func(response *DeleteReceiverResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DeleteReceiverResponse
var err error
defer close(result)
response, err = client.DeleteReceiver(request)
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

type DeleteReceiverRequest struct {
*requests.RpcRequest
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                ReceiverId  string `position:"Query" name:"ReceiverId"`
}


type DeleteReceiverResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteReceiverRequest() (request *DeleteReceiverRequest) {
request = &DeleteReceiverRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Dm", "2015-11-23", "DeleteReceiver", "", "")
return
}

func CreateDeleteReceiverResponse() (response *DeleteReceiverResponse) {
response = &DeleteReceiverResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

