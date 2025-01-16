// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRepositoryCommitAnalyticsAuthorsRequest wrapper for the ListRepositoryCommitAnalyticsAuthors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListRepositoryCommitAnalyticsAuthors.go.html to see an example of how to use ListRepositoryCommitAnalyticsAuthorsRequest.
type ListRepositoryCommitAnalyticsAuthorsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRepositoryCommitAnalyticsAuthorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort by value is supported for this parameter. Default order for author name is ascending.
	SortBy ListRepositoryCommitAnalyticsAuthorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoryCommitAnalyticsAuthorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoryCommitAnalyticsAuthorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoryCommitAnalyticsAuthorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoryCommitAnalyticsAuthorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRepositoryCommitAnalyticsAuthorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRepositoryCommitAnalyticsAuthorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRepositoryCommitAnalyticsAuthorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryCommitAnalyticsAuthorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRepositoryCommitAnalyticsAuthorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRepositoryCommitAnalyticsAuthorsResponse wrapper for the ListRepositoryCommitAnalyticsAuthors operation
type ListRepositoryCommitAnalyticsAuthorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CommitAnalyticsAuthorCollection instances
	CommitAnalyticsAuthorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRepositoryCommitAnalyticsAuthorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoryCommitAnalyticsAuthorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoryCommitAnalyticsAuthorsSortOrderEnum Enum with underlying type: string
type ListRepositoryCommitAnalyticsAuthorsSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoryCommitAnalyticsAuthorsSortOrderEnum
const (
	ListRepositoryCommitAnalyticsAuthorsSortOrderAsc  ListRepositoryCommitAnalyticsAuthorsSortOrderEnum = "ASC"
	ListRepositoryCommitAnalyticsAuthorsSortOrderDesc ListRepositoryCommitAnalyticsAuthorsSortOrderEnum = "DESC"
)

var mappingListRepositoryCommitAnalyticsAuthorsSortOrderEnum = map[string]ListRepositoryCommitAnalyticsAuthorsSortOrderEnum{
	"ASC":  ListRepositoryCommitAnalyticsAuthorsSortOrderAsc,
	"DESC": ListRepositoryCommitAnalyticsAuthorsSortOrderDesc,
}

var mappingListRepositoryCommitAnalyticsAuthorsSortOrderEnumLowerCase = map[string]ListRepositoryCommitAnalyticsAuthorsSortOrderEnum{
	"asc":  ListRepositoryCommitAnalyticsAuthorsSortOrderAsc,
	"desc": ListRepositoryCommitAnalyticsAuthorsSortOrderDesc,
}

// GetListRepositoryCommitAnalyticsAuthorsSortOrderEnumValues Enumerates the set of values for ListRepositoryCommitAnalyticsAuthorsSortOrderEnum
func GetListRepositoryCommitAnalyticsAuthorsSortOrderEnumValues() []ListRepositoryCommitAnalyticsAuthorsSortOrderEnum {
	values := make([]ListRepositoryCommitAnalyticsAuthorsSortOrderEnum, 0)
	for _, v := range mappingListRepositoryCommitAnalyticsAuthorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryCommitAnalyticsAuthorsSortOrderEnumStringValues Enumerates the set of values in String for ListRepositoryCommitAnalyticsAuthorsSortOrderEnum
func GetListRepositoryCommitAnalyticsAuthorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRepositoryCommitAnalyticsAuthorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryCommitAnalyticsAuthorsSortOrderEnum(val string) (ListRepositoryCommitAnalyticsAuthorsSortOrderEnum, bool) {
	enum, ok := mappingListRepositoryCommitAnalyticsAuthorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRepositoryCommitAnalyticsAuthorsSortByEnum Enum with underlying type: string
type ListRepositoryCommitAnalyticsAuthorsSortByEnum string

// Set of constants representing the allowable values for ListRepositoryCommitAnalyticsAuthorsSortByEnum
const (
	ListRepositoryCommitAnalyticsAuthorsSortByAuthorname ListRepositoryCommitAnalyticsAuthorsSortByEnum = "authorName"
)

var mappingListRepositoryCommitAnalyticsAuthorsSortByEnum = map[string]ListRepositoryCommitAnalyticsAuthorsSortByEnum{
	"authorName": ListRepositoryCommitAnalyticsAuthorsSortByAuthorname,
}

var mappingListRepositoryCommitAnalyticsAuthorsSortByEnumLowerCase = map[string]ListRepositoryCommitAnalyticsAuthorsSortByEnum{
	"authorname": ListRepositoryCommitAnalyticsAuthorsSortByAuthorname,
}

// GetListRepositoryCommitAnalyticsAuthorsSortByEnumValues Enumerates the set of values for ListRepositoryCommitAnalyticsAuthorsSortByEnum
func GetListRepositoryCommitAnalyticsAuthorsSortByEnumValues() []ListRepositoryCommitAnalyticsAuthorsSortByEnum {
	values := make([]ListRepositoryCommitAnalyticsAuthorsSortByEnum, 0)
	for _, v := range mappingListRepositoryCommitAnalyticsAuthorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryCommitAnalyticsAuthorsSortByEnumStringValues Enumerates the set of values in String for ListRepositoryCommitAnalyticsAuthorsSortByEnum
func GetListRepositoryCommitAnalyticsAuthorsSortByEnumStringValues() []string {
	return []string{
		"authorName",
	}
}

// GetMappingListRepositoryCommitAnalyticsAuthorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryCommitAnalyticsAuthorsSortByEnum(val string) (ListRepositoryCommitAnalyticsAuthorsSortByEnum, bool) {
	enum, ok := mappingListRepositoryCommitAnalyticsAuthorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
