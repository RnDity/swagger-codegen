/* 
 * Swagger Petstore *_/ ' \" =end \\r\\n \\n \\r
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  *_/ ' \" =end       
 *
 * OpenAPI spec version: 1.0.0 *_/ ' \" =end \\r\\n \\n \\r
 * Contact: apiteam@swagger.io *_/ ' \" =end \\r\\n \\n \\r
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	)

type FakeApi struct {
	Configuration Configuration
}

func NewFakeApi() *FakeApi {
	configuration := NewConfiguration()
	return &FakeApi{
		Configuration: *configuration,
	}
}

func NewFakeApiWithBasePath(basePath string) *FakeApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &FakeApi{
		Configuration: *configuration,
	}
}

/**
 * To test code injection *_/ &#39; \&quot; &#x3D;end \\r\\n \\n \\r
 *
 * @param testCodeInjectEndRnNR To test code injection *_/ &#39; \&quot; &#x3D;end \\r\\n \\n \\r
 * @return void
 */
func (a FakeApi) TestCodeInjectEndRnNR(testCodeInjectEndRnNR string) (*APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/fake"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "*_/ '  =end       ",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"*_/ '  =end       ",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	formParams["testCodeInjectEndRnNR"] = testCodeInjectEndRnNR

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return NewAPIResponse(httpResponse.RawResponse), err
	}

	return NewAPIResponse(httpResponse.RawResponse), err
}

