
package nas

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

func (client *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
response = CreateDescribeMountTargetsResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeMountTargetsWithChan(request *DescribeMountTargetsRequest) (<-chan *DescribeMountTargetsResponse, <-chan error) {
responseChan := make(chan *DescribeMountTargetsResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeMountTargets(request)
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

func (client *Client) DescribeMountTargetsWithCallback(request *DescribeMountTargetsRequest, callback func(response *DescribeMountTargetsResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeMountTargetsResponse
var err error
defer close(result)
response, err = client.DescribeMountTargets(request)
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

type DescribeMountTargetsRequest struct {
*requests.RpcRequest
                PageSize  string `position:"Query" name:"PageSize"`
                MountTargetDomain  string `position:"Query" name:"MountTargetDomain"`
                PageNumber  string `position:"Query" name:"PageNumber"`
                FileSystemId  string `position:"Query" name:"FileSystemId"`
}


type DescribeMountTargetsResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
                MountTargets struct {
                    MountTarget []struct {
            MountTargetDomain     string `json:"MountTargetDomain" xml:"MountTargetDomain"`
            NetworkType     string `json:"NetworkType" xml:"NetworkType"`
            VpcId     string `json:"VpcId" xml:"VpcId"`
            VswId     string `json:"VswId" xml:"VswId"`
            AccessGroup     string `json:"AccessGroup" xml:"AccessGroup"`
            Status     string `json:"Status" xml:"Status"`
                    }   `json:"MountTarget" xml:"MountTarget"`
                } `json:"MountTargets" xml:"MountTargets"`
}

func CreateDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
request = &DescribeMountTargetsRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("NAS", "2017-06-26", "DescribeMountTargets", "", "")
return
}

func CreateDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
response = &DescribeMountTargetsResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

