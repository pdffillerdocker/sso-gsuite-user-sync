// Copyright (c) 2020, Amazon.com, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

// UserEmail represents a user email address
type UserEmail struct {
	Value   string `json:"value"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

// UserAddress represents address values of users
type UserAddress struct {
	Type string `json:"type"`
}

// User represents a User in AWS SSO
type User struct {
	ID       string   `json:"id,omitempty"`
	Schemas  []string `json:"schemas"`
	Username string   `json:"userName"`
	Name     struct {
		FamilyName string `json:"familyName"`
		GivenName  string `json:"givenName"`
	} `json:"name"`
	DisplayName string        `json:"displayName"`
	Active      bool          `json:"active"`
	Emails      []UserEmail   `json:"emails"`
	Addresses   []UserAddress `json:"addresses"`
}

// UserFilterResults represents filtered results when we search for
// users or List all users
type UserFilterResults struct {
	Schemas      []string `json:"schemas"`
	TotalResults int      `json:"totalResults"`
	ItemsPerPage int      `json:"itemsPerPage"`
	StartIndex   int      `json:"startIndex"`
	Resources    []User   `json:"Resources"`
}
