
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

func (client *Client) DescribeDomainRealTimeQpsData(request *DescribeDomainRealTimeQpsDataRequest) (response *DescribeDomainRealTimeQpsDataResponse, err error) {
response = CreateDescribeDomainRealTimeQpsDataResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeDomainRealTimeQpsDataWithChan(request *DescribeDomainRealTimeQpsDataRequest) (<-chan *DescribeDomainRealTimeQpsDataResponse, <-chan error) {
responseChan := make(chan *DescribeDomainRealTimeQpsDataResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeDomainRealTimeQpsData(request)
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

func (client *Client) DescribeDomainRealTimeQpsDataWithCallback(request *DescribeDomainRealTimeQpsDataRequest, callback func(response *DescribeDomainRealTimeQpsDataResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeDomainRealTimeQpsDataResponse
var err error
defer close(result)
response, err = client.DescribeDomainRealTimeQpsData(request)
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

type DescribeDomainRealTimeQpsDataRequest struct {
*requests.RpcRequest
                EndTime  string `position:"Query" name:"EndTime"`
                Version  string `position:"Query" name:"Version"`
                DomainName  string `position:"Query" name:"DomainName"`
                StartTime  string `position:"Query" name:"StartTime"`
                IspNameEn  string `position:"Query" name:"IspNameEn"`
                LocationNameEn  string `position:"Query" name:"LocationNameEn"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                SecurityToken  string `position:"Query" name:"SecurityToken"`
}


type DescribeDomainRealTimeQpsDataResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
                Data struct {
                    QpsModel []struct {
            Qps     float64 `json:"Qps" xml:"Qps"`
            TimeStamp     string `json:"TimeStamp" xml:"TimeStamp"`
                    }   `json:"QpsModel" xml:"QpsModel"`
                } `json:"Data" xml:"Data"`
}

func CreateDescribeDomainRealTimeQpsDataRequest() (request *DescribeDomainRealTimeQpsDataRequest) {
request = &DescribeDomainRealTimeQpsDataRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeQpsData", "", "")
return
}

func CreateDescribeDomainRealTimeQpsDataResponse() (response *DescribeDomainRealTimeQpsDataResponse) {
response = &DescribeDomainRealTimeQpsDataResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

