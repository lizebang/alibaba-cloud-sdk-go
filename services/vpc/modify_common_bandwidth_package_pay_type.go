
package vpc

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

func (client *Client) ModifyCommonBandwidthPackagePayType(request *ModifyCommonBandwidthPackagePayTypeRequest) (response *ModifyCommonBandwidthPackagePayTypeResponse, err error) {
response = CreateModifyCommonBandwidthPackagePayTypeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ModifyCommonBandwidthPackagePayTypeWithChan(request *ModifyCommonBandwidthPackagePayTypeRequest) (<-chan *ModifyCommonBandwidthPackagePayTypeResponse, <-chan error) {
responseChan := make(chan *ModifyCommonBandwidthPackagePayTypeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ModifyCommonBandwidthPackagePayType(request)
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

func (client *Client) ModifyCommonBandwidthPackagePayTypeWithCallback(request *ModifyCommonBandwidthPackagePayTypeRequest, callback func(response *ModifyCommonBandwidthPackagePayTypeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ModifyCommonBandwidthPackagePayTypeResponse
var err error
defer close(result)
response, err = client.ModifyCommonBandwidthPackagePayType(request)
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

type ModifyCommonBandwidthPackagePayTypeRequest struct {
*requests.RpcRequest
                ResourceUid  string `position:"Query" name:"ResourceUid"`
                KbpsBandwidth  string `position:"Query" name:"KbpsBandwidth"`
                ResourceBid  string `position:"Query" name:"ResourceBid"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                BandwidthPackageId  string `position:"Query" name:"BandwidthPackageId"`
                Bandwidth  string `position:"Query" name:"Bandwidth"`
                Duration  string `position:"Query" name:"Duration"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                AutoPay  string `position:"Query" name:"AutoPay"`
                PayType  string `position:"Query" name:"PayType"`
                PricingCycle  string `position:"Query" name:"PricingCycle"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
}


type ModifyCommonBandwidthPackagePayTypeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            OrderId     int64 `json:"OrderId" xml:"OrderId"`
            Code     string `json:"Code" xml:"Code"`
            Message     string `json:"Message" xml:"Message"`
}

func CreateModifyCommonBandwidthPackagePayTypeRequest() (request *ModifyCommonBandwidthPackagePayTypeRequest) {
request = &ModifyCommonBandwidthPackagePayTypeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackagePayType", "", "")
return
}

func CreateModifyCommonBandwidthPackagePayTypeResponse() (response *ModifyCommonBandwidthPackagePayTypeResponse) {
response = &ModifyCommonBandwidthPackagePayTypeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

