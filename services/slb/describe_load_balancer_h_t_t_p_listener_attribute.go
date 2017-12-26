
package slb

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

func (client *Client) DescribeLoadBalancerHTTPListenerAttribute(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error) {
response = CreateDescribeLoadBalancerHTTPListenerAttributeResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithChan(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (<-chan *DescribeLoadBalancerHTTPListenerAttributeResponse, <-chan error) {
responseChan := make(chan *DescribeLoadBalancerHTTPListenerAttributeResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithCallback(request *DescribeLoadBalancerHTTPListenerAttributeRequest, callback func(response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribeLoadBalancerHTTPListenerAttributeResponse
var err error
defer close(result)
response, err = client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
*requests.RpcRequest
                Tags  string `position:"Query" name:"Tags"`
                ListenerPort  string `position:"Query" name:"ListenerPort"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                AccessKeyId  string `position:"Query" name:"access_key_id"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                LoadBalancerId  string `position:"Query" name:"LoadBalancerId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            ListenerPort     int `json:"ListenerPort" xml:"ListenerPort"`
            BackendServerPort     int `json:"BackendServerPort" xml:"BackendServerPort"`
            Bandwidth     int `json:"Bandwidth" xml:"Bandwidth"`
            Status     string `json:"Status" xml:"Status"`
            SecurityStatus     string `json:"SecurityStatus" xml:"SecurityStatus"`
            XForwardedFor     string `json:"XForwardedFor" xml:"XForwardedFor"`
            Scheduler     string `json:"Scheduler" xml:"Scheduler"`
            StickySession     string `json:"StickySession" xml:"StickySession"`
            StickySessionType     string `json:"StickySessionType" xml:"StickySessionType"`
            CookieTimeout     int `json:"CookieTimeout" xml:"CookieTimeout"`
            Cookie     string `json:"Cookie" xml:"Cookie"`
            HealthCheck     string `json:"HealthCheck" xml:"HealthCheck"`
            HealthCheckDomain     string `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
            HealthCheckURI     string `json:"HealthCheckURI" xml:"HealthCheckURI"`
            HealthyThreshold     int `json:"HealthyThreshold" xml:"HealthyThreshold"`
            UnhealthyThreshold     int `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
            HealthCheckTimeout     int `json:"HealthCheckTimeout" xml:"HealthCheckTimeout"`
            HealthCheckInterval     int `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
            HealthCheckConnectPort     int `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
            HealthCheckHttpCode     string `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
            MaxConnection     int `json:"MaxConnection" xml:"MaxConnection"`
            VServerGroupId     string `json:"VServerGroupId" xml:"VServerGroupId"`
            Gzip     string `json:"Gzip" xml:"Gzip"`
            XForwardedForSLBIP     string `json:"XForwardedFor_SLBIP" xml:"XForwardedFor_SLBIP"`
            XForwardedForSLBID     string `json:"XForwardedFor_SLBID" xml:"XForwardedFor_SLBID"`
            XForwardedForProto     string `json:"XForwardedFor_proto" xml:"XForwardedFor_proto"`
}

func CreateDescribeLoadBalancerHTTPListenerAttributeRequest() (request *DescribeLoadBalancerHTTPListenerAttributeRequest) {
request = &DescribeLoadBalancerHTTPListenerAttributeRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerHTTPListenerAttribute", "", "")
return
}

func CreateDescribeLoadBalancerHTTPListenerAttributeResponse() (response *DescribeLoadBalancerHTTPListenerAttributeResponse) {
response = &DescribeLoadBalancerHTTPListenerAttributeResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

