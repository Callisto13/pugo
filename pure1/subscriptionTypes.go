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

// type Subscription interface{}

// type Subscription struct {
// 	ID   string `json:"id,omitempty"`
// 	Name string `json:"name,omitempty"`
// }

// SubscriptionAsset type describes the subscription-asset object returned by the API
type SubscriptionAsset struct {
	ID              string            `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	InstallAddress  map[string]string `json:"install_address,omitempty"`
	ActivationDate  int               `json:"activation_date,omitempty"`
	EndOfLifeDate   int               `json:"end_of_life_date,omitempty"`
	Version         string            `json:"version,omitempty"`
	Model           string            `json:"model,omitempty"`
	ChassisSerialNo string            `json:"chassis_serial_number,omitempty"`
	Space           Space
	Subscription    ResourceData `json:"subscription,omitempty"`
	License         ResourceData `json:"license,omitempty"`
	AsOf            int          `json:"_as_of,omitempty"`
}

type Space struct {
	Effective     MetricData `json:"effective_usage,omitempty"`
	UsedRatio     MetricData `json:"used_ratio,omitempty"`
	Capacity      MetricData `json:"capacity,omitempty"`
	DataReduction MetricData `json:"data_reduction,omitempty"`
}

type MetricData struct {
	Data   float64      `json:"data,omitempty"`
	Metric ResourceData `json:"metric,omitempty"`
	Unit   string       `json:"unit,omitempty"`
}

type ResourceData struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ResourceType string `json:"metrics,omitempty"`
}

// SubscriptionLicense type describes the subscription-license object returned by the API
type SubscriptionLicense struct {
	ID                 string                 `json:"id,omitempty"`
	Name               string                 `json:"name,omitempty"`
	AverageOnDemand    MetricData             `json:"average_on_demand,omitempty"`
	StartDate          int                    `json:"start_date,omitempty"`
	ExpirationDate     int                    `json:"expiration_date,omitempty"`
	LastUpdatedDate    int                    `json:"last_updated_date,omitempty"`
	MarketplacePartner map[string]interface{} `json:"marketplace_partner,omitempty"`
	PowerAvg           interface{}            `json:"power_average,omitempty"`
	PreRatio           MetricData             `json:"pre_ratio,omitempty"`
	QuarterOnDemand    MetricData             `json:"quarter_on_demand,omitempty"`
	Reservation        MetricData             `json:"reservation,omitempty"`
	Resources          []Resource             `json:"resources,omitempty"`
	ServiceTier        string                 `json:"service_tier,omitempty"`
	SiteAddress        map[string]string      `json:"site_address,omitempty"`
	Subscription       ResourceData           `json:"subscription,omitempty"`
	Usage              MetricData             `json:"usage,omitempty"`
	AsOf               int                    `json:"_as_of,omitempty"`
}

type Resource struct {
	ActivationTime int        `json:"activation_time,omitempty"`
	FQDN           string     `json:"fqdn,omitempty"`
	ID             string     `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	ResourceType   string     `json:"resource_type,omitempty"`
	Usage          MetricData `json:"usage,omitempty"`
}
