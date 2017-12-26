
package mts

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

func (client *Client) ListMedia(request *ListMediaRequest) (response *ListMediaResponse, err error) {
response = CreateListMediaResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) ListMediaWithChan(request *ListMediaRequest) (<-chan *ListMediaResponse, <-chan error) {
responseChan := make(chan *ListMediaResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.ListMedia(request)
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

func (client *Client) ListMediaWithCallback(request *ListMediaRequest, callback func(response *ListMediaResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *ListMediaResponse
var err error
defer close(result)
response, err = client.ListMedia(request)
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

type ListMediaRequest struct {
*requests.RpcRequest
                To  string `position:"Query" name:"To"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                From  string `position:"Query" name:"From"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                MaximumPageSize  string `position:"Query" name:"MaximumPageSize"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                NextPageToken  string `position:"Query" name:"NextPageToken"`
}


type ListMediaResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            NextPageToken     string `json:"NextPageToken" xml:"NextPageToken"`
                MediaList struct {
                    Media []struct {
            MediaId     string `json:"MediaId" xml:"MediaId"`
            Title     string `json:"Title" xml:"Title"`
            Description     string `json:"Description" xml:"Description"`
            CoverURL     string `json:"CoverURL" xml:"CoverURL"`
            CateId     int64 `json:"CateId" xml:"CateId"`
            Duration     string `json:"Duration" xml:"Duration"`
            Format     string `json:"Format" xml:"Format"`
            Size     string `json:"Size" xml:"Size"`
            Bitrate     string `json:"Bitrate" xml:"Bitrate"`
            Width     string `json:"Width" xml:"Width"`
            Height     string `json:"Height" xml:"Height"`
            Fps     string `json:"Fps" xml:"Fps"`
            PublishState     string `json:"PublishState" xml:"PublishState"`
            CreationTime     string `json:"CreationTime" xml:"CreationTime"`
                Tags struct {
                Tag []    string `json:"Tag" xml:"Tag"`
                } `json:"Tags" xml:"Tags"`
                RunIdList struct {
                RunId []    string `json:"RunId" xml:"RunId"`
                } `json:"RunIdList" xml:"RunIdList"`
            File struct {
            URL     string `json:"URL" xml:"URL"`
            State     string `json:"State" xml:"State"`
            }  `json:"File" xml:"File"`
                    }   `json:"Media" xml:"Media"`
                } `json:"MediaList" xml:"MediaList"`
}

func CreateListMediaRequest() (request *ListMediaRequest) {
request = &ListMediaRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Mts", "2014-06-18", "ListMedia", "", "")
return
}

func CreateListMediaResponse() (response *ListMediaResponse) {
response = &ListMediaResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

