/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package pure1

import (
	"encoding/json"
)

// SubscriptionService type creates a service to expose metrics endpoints
type SubscriptionService struct {
	client *Client
}

// GetSubscriptions returns a list of subscription objects
func (s *SubscriptionService) GetSubscriptions(params map[string]string) ([]interface{}, error) {
	var path = "subscriptions"

	req, err := s.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	r := &pure1Response{}

	_, err = s.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []interface{}{}
	// m := []Subscription{}

	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			var i interface{}
			str, _ := json.Marshal(item)
			json.Unmarshal([]byte(str), &i)
			m = append(m, i)
		}

		if len(m) < r.TotalItems {
			if r.ContinuationToken != nil {
				if params == nil {
					params = map[string]string{"continuation_token": r.ContinuationToken.(string)}
				} else {
					params["continuation_token"] = r.ContinuationToken.(string)
				}
				req, err := s.client.NewRequest("GET", path, params, nil)
				if err != nil {
					return nil, err
				}

				_, err = s.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return m, err
}

// GetSubscriptionAssets returns a list of subscription objects
func (s *SubscriptionService) GetSubscriptionAssets(params map[string]string) ([]SubscriptionAsset, error) {
	var path = "subscription-assets"

	req, err := s.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	r := &pure1Response{}

	_, err = s.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []SubscriptionAsset{}

	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := SubscriptionAsset{}
			str, _ := json.Marshal(item)
			json.Unmarshal([]byte(str), &i)
			m = append(m, i)
		}

		if len(m) < r.TotalItems {
			if r.ContinuationToken != nil {
				if params == nil {
					params = map[string]string{"continuation_token": r.ContinuationToken.(string)}
				} else {
					params["continuation_token"] = r.ContinuationToken.(string)
				}
				req, err := s.client.NewRequest("GET", path, params, nil)
				if err != nil {
					return nil, err
				}

				_, err = s.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return m, err
}

// GetSubscriptionLicenses returns a list of subscription objects
func (s *SubscriptionService) GetSubscriptionLicenses(params map[string]string) ([]SubscriptionLicense, error) {
	var path = "subscription-licenses"

	req, err := s.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	r := &pure1Response{}

	_, err = s.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []SubscriptionLicense{}

	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			var i SubscriptionLicense
			str, _ := json.Marshal(item)
			json.Unmarshal([]byte(str), &i)
			m = append(m, i)
		}

		if len(m) < r.TotalItems {
			if r.ContinuationToken != nil {
				if params == nil {
					params = map[string]string{"continuation_token": r.ContinuationToken.(string)}
				} else {
					params["continuation_token"] = r.ContinuationToken.(string)
				}
				req, err := s.client.NewRequest("GET", path, params, nil)
				if err != nil {
					return nil, err
				}

				_, err = s.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return m, err
}
