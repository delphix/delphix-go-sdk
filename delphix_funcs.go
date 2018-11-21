package delphix

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

//APISessionCreate - Create a new APISession object.
//
//payload *APISessionStruct object
//return *APISessionStruct
func (c *Client) APISessionCreate(payload *APISessionStruct) (*APISessionStruct, error) {
	var err error
	var result *APISessionStruct
	return result, err
}

//APISessionRead - Returns the settings of the current session, if one has been
// started.
//
//r - APISession Object Reference
//return *APISessionStruct
func (c *Client) APISessionRead(r string) (*APISessionStruct, error) {
	var err error
	var result *APISessionStruct
	obj, err := c.FindObjectByReference("/session", r)
	result = obj.(*APISessionStruct)
	return result, err
}

//ActionRead - Retrieve the specified Action object.
//
//r - Action Object Reference
//return *ActionStruct
func (c *Client) ActionRead(r string) (*ActionStruct, error) {
	var err error
	var result *ActionStruct
	obj, err := c.FindObjectByReference("/action", r)
	result = obj.(*ActionStruct)
	return result, err
}

//ActionList - Retrieve an historical log of actions.
//
//searchText ==>
//   description: Limit search results to only include alerts that have searchText string in reference, state, title, details, failureDescription, parentAction, workSource or workSourceName.
//   type: string
//fromDate ==>
//   description: Start date for the search. Only actions on or after this date are included.
//   type: string
//   format: date
//toDate ==>
//   description: End date for the search. Only actions on or before this date are included.
//   type: string
//   format: date
//pageSize ==>
//   description: Limit the number of events returned.
//   type: integer
//   default: 25
//user ==>
//   referenceTo: /delphix-user.json
//   mapsTo: user
//   description: Limit actions to those initiated by this user.
//   type: string
//   format: objectReference
//ascending ==>
//   description: True if results are to be returned in ascending order.
//   type: boolean
//sortBy ==>
//   description: Search results are sorted by the field provided.
//   type: string
//   enum: [reference state title details startTime endTime failureDescription parentAction workSource workSourceName]
//pageOffset ==>
//   description: Offset within event list, in units of pageSize chunks.
//   type: integer
//state ==>
//   description: Limit actions to those in the specified state.
//   type: string
//   enum: [EXECUTING WAITING COMPLETED FAILED CANCELED]
//   mapsTo: state
//ignoreActionTypes ==>
//   type: array
//   items: map[type:string]
//   description: Ignore actions with an action type in this list.
//parentAction ==>
//   referenceTo: /delphix-action.json
//   mapsTo: parentAction
//   description: Limit actions to those with this parent action.
//   type: string
//   format: objectReference
//rootActionOnly ==>
//   description: Limit actions to those without a parent action.
//   type: boolean
//return []*ActionStruct
func (c *Client) ActionList(searchText string, fromDate string, toDate string, pageSize *int, user string, ascending *bool, sortBy string, pageOffset *int, state string, ignoreActionTypes []string, parentAction string, rootActionOnly *bool) ([]*ActionStruct, error) {
	// var err error
	// var result []*ActionStruct
	// obj, err := c.executeListAndReturnResults("/action")
	// for _, i := range obj {
	// 	result = append(result, i.(*ActionStruct))
	// }
	// return result, err
	var err error
	var result []*ActionStruct
	obj, err := c.executeListAndReturnResults("/action")
	for _, i := range obj.Result.([]interface{}) {
		objMap := i.(map[string]interface{})
		switch objMap["type"].(string) {
		}
		result = append(result, i.(*ActionStruct))
	}
	return result, err
}

//AlertRead - Retrieve the specified Alert object.
//
//r - Alert Object Reference
//return *AlertStruct
func (c *Client) AlertRead(r string) (*AlertStruct, error) {
	var err error
	var result *AlertStruct
	obj, err := c.FindObjectByReference("/alert", r)
	result = obj.(*AlertStruct)
	return result, err
}

//AlertList - Returns a list of alerts on the system.
//
//sortBy ==>
//   description: Search results are sorted by the field provided.
//   type: string
//   enum: [event eventTitle eventDescription eventResponse eventAction eventCommandOutput eventSeverity target targetName timestamp]
//fromDate ==>
//   format: date
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//   description: Start date to use for the search.
//   type: string
//toDate ==>
//   type: string
//   format: date
//   mapsTo: timestamp
//   inequalityType: STRICT
//   description: End date to use for the search.
//pageSize ==>
//   description: Limit the number of alerts returned.
//   type: integer
//   default: 25
//   minimum: 0
//maxTotal ==>
//   description: The upper bound for calculation of total alert count.
//   type: integer
//target ==>
//   referenceTo: /delphix-user-object.json
//   mapsTo: target
//   description: Limit alerts to those affecting a particular object on the system.
//   type: string
//   format: objectReference
//pageOffset ==>
//   description: Offset within alert list, in units of pageSize chunks.
//   type: integer
//searchText ==>
//   description: Limit search results to only include alerts that have searchText string in eventTitle, eventDescription, eventResponse, eventAction, or severity.
//   type: string
//ascending ==>
//   description: True if results are to be returned in ascending order.
//   type: boolean
//return []*AlertStruct
func (c *Client) AlertList(sortBy string, fromDate string, toDate string, pageSize *int, maxTotal *int, target string, pageOffset *int, searchText string, ascending *bool) ([]*AlertStruct, error) {
	var err error
	var result []*AlertStruct
	obj, err := c.executeListAndReturnResults("/alert")
	for _, i := range obj {
		result = append(result, i.(*AlertStruct))
	}
	return result, err
}

//AlertProfileCreate - Create a new AlertProfile object.
//
//payload *AlertProfileStruct object
//return string
func (c *Client) AlertProfileCreate(payload *AlertProfileStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//AlertProfileDelete - Delete the specified AlertProfile object.
//
//r - AlertProfile Object Reference
func (c *Client) AlertProfileDelete(r string) error {
	var err error
	return err
}

//AlertProfileList - List AlertProfile objects on the system.
//
//return []*AlertProfileStruct
func (c *Client) AlertProfileList() ([]*AlertProfileStruct, error) {
	var err error
	var result []*AlertProfileStruct
	obj, err := c.executeListAndReturnResults("/alert/profile")
	for _, i := range obj {
		result = append(result, i.(*AlertProfileStruct))
	}
	return result, err
}

//AlertProfileRead - Retrieve the specified AlertProfile object.
//
//r - AlertProfile Object Reference
//return *AlertProfileStruct
func (c *Client) AlertProfileRead(r string) (*AlertProfileStruct, error) {
	var err error
	var result *AlertProfileStruct
	obj, err := c.FindObjectByReference("/alert/profile", r)
	result = obj.(*AlertProfileStruct)
	return result, err
}

//AlertProfileUpdate - Update the specified AlertProfile object.
//
//r - AlertProfile Object Reference
//payload *AlertProfileStruct object
func (c *Client) AlertProfileUpdate(r string, payload *AlertProfileStruct) error {
	var err error
	return err
}

//AuthorizationRead - Retrieve the specified Authorization object.
//
//r - Authorization Object Reference
//return *AuthorizationStruct
func (c *Client) AuthorizationRead(r string) (*AuthorizationStruct, error) {
	var err error
	var result *AuthorizationStruct
	obj, err := c.FindObjectByReference("/authorization", r)
	result = obj.(*AuthorizationStruct)
	return result, err
}

//AuthorizationDelete - Delete the specified Authorization object.
//
//r - Authorization Object Reference
func (c *Client) AuthorizationDelete(r string) error {
	var err error
	return err
}

//AuthorizationCreate - Create a new Authorization object.
//
//payload *AuthorizationStruct object
//return string
func (c *Client) AuthorizationCreate(payload *AuthorizationStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//AuthorizationList - Lists authorizations granted in the system.
//
//user ==>
//   mapsTo: user
//   description: Lists permissions granted to the specified user.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user.json
//   required: false
//target ==>
//   mapsTo: target
//   description: Lists the permissions granted on the specified object.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   required: false
//effective ==>
//   description: Also return inherited authorizations.
//   type: string
//   required: false
//return []*AuthorizationStruct
func (c *Client) AuthorizationList(user string, target string, effective string) ([]*AuthorizationStruct, error) {
	var err error
	var result []*AuthorizationStruct
	obj, err := c.executeListAndReturnResults("/authorization")
	for _, i := range obj {
		result = append(result, i.(*AuthorizationStruct))
	}
	return result, err
}

//AuthorizationConfigRead - Retrieve the specified AuthorizationConfig object.
//
//r - AuthorizationConfig Object Reference
//return *AuthorizationConfigStruct
func (c *Client) AuthorizationConfigRead(r string) (*AuthorizationConfigStruct, error) {
	var err error
	var result *AuthorizationConfigStruct
	obj, err := c.FindObjectByReference("/authorization/configuration", r)
	result = obj.(*AuthorizationConfigStruct)
	return result, err
}

//AuthorizationConfigUpdate - Update the specified AuthorizationConfig object.
//
//r - AuthorizationConfig Object Reference
//payload *AuthorizationConfigStruct object
func (c *Client) AuthorizationConfigUpdate(r string, payload *AuthorizationConfigStruct) error {
	var err error
	return err
}

//CaCertificateDelete - Delete the specified CaCertificate object.
//
//r - CaCertificate Object Reference
func (c *Client) CaCertificateDelete(r string) error {
	var err error
	return err
}

//CaCertificateList - List CaCertificate objects on the system.
//
//delphixCa ==>
//   description: List the Delphix CA only. If false, display everything except the Delphix CA.
//   type: boolean
//return []*CaCertificateStruct
func (c *Client) CaCertificateList(delphixCa *bool) ([]*CaCertificateStruct, error) {
	var err error
	var result []*CaCertificateStruct
	obj, err := c.executeListAndReturnResults("/service/tls/caCertificate")
	for _, i := range obj {
		result = append(result, i.(*CaCertificateStruct))
	}
	return result, err
}

//CaCertificateCreate - Import a CA certificate in PEM format.
//
//payload *PemCertificateStruct object
//return string
func (c *Client) CaCertificateCreate(payload *PemCertificateStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//CaCertificateRead - Retrieve the specified CaCertificate object.
//
//r - CaCertificate Object Reference
//return *CaCertificateStruct
func (c *Client) CaCertificateRead(r string) (*CaCertificateStruct, error) {
	var err error
	var result *CaCertificateStruct
	obj, err := c.FindObjectByReference("/service/tls/caCertificate", r)
	result = obj.(*CaCertificateStruct)
	return result, err
}

//CertificateConfigRead - Retrieve the specified CertificateConfig object.
//
//r - CertificateConfig Object Reference
//return *CertificateConfigStruct
func (c *Client) CertificateConfigRead(r string) (*CertificateConfigStruct, error) {
	var err error
	var result *CertificateConfigStruct
	obj, err := c.FindObjectByReference("/service/tls", r)
	result = obj.(*CertificateConfigStruct)
	return result, err
}

//CertificateConfigUpdate - Update the specified CertificateConfig object.
//
//r - CertificateConfig Object Reference
//payload *CertificateConfigStruct object
func (c *Client) CertificateConfigUpdate(r string, payload *CertificateConfigStruct) error {
	var err error
	return err
}

//CertificateSigningRequestList - List CertificateSigningRequest objects on the system.
//
//return []*CertificateSigningRequestStruct
func (c *Client) CertificateSigningRequestList() ([]*CertificateSigningRequestStruct, error) {
	var err error
	var result []*CertificateSigningRequestStruct
	obj, err := c.executeListAndReturnResults("/service/tls/csr")
	for _, i := range obj {
		result = append(result, i.(*CertificateSigningRequestStruct))
	}
	return result, err
}

//CertificateSigningRequestCreate - Create a new CertificateSigningRequest object.
//
//payload *CertificateSigningRequestCreateParametersStruct object
//return string
func (c *Client) CertificateSigningRequestCreate(payload *CertificateSigningRequestCreateParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//CertificateSigningRequestRead - Retrieve the specified CertificateSigningRequest object.
//
//r - CertificateSigningRequest Object Reference
//return *CertificateSigningRequestStruct
func (c *Client) CertificateSigningRequestRead(r string) (*CertificateSigningRequestStruct, error) {
	var err error
	var result *CertificateSigningRequestStruct
	obj, err := c.FindObjectByReference("/service/tls/csr", r)
	result = obj.(*CertificateSigningRequestStruct)
	return result, err
}

//CertificateSigningRequestDelete - Delete the specified CertificateSigningRequest object.
//
//r - CertificateSigningRequest Object Reference
func (c *Client) CertificateSigningRequestDelete(r string) error {
	var err error
	return err
}

//ContainerRead - Retrieve the specified Container object.
//
//r - Container Object Reference
//return Container
func (c *Client) ContainerRead(r string) (Container, error) {
	var err error
	var result Container
	obj, err := c.FindObjectByReference("/database", r)
	result = obj.(Container)
	return result, err
}

//ContainerUpdate - Update the specified Container object.
//
//r - Container Object Reference
//payload Container object
func (c *Client) ContainerUpdate(r string, payload Container) error {
	var err error
	return err
}

//ContainerDelete - Delete the specified Container object.
//
//r - Container Object Reference
//payload *DeleteParametersStruct object
func (c *Client) ContainerDelete(r string, payload *DeleteParametersStruct) error {
	var err error
	return err
}

//ContainerList - Returns a list of databases on the system or within a group.
//
//noJSDataSource ==>
//   description: Restrict databases to those which are not part of a Self-Service data layout (data template or data container). This option is mutually exclusive with the "noJSContainerDataSource" option.
//   type: boolean
//noJSContainerDataSource ==>
//   description: Restrict databases to those which are not part of a Self-Service data container. This option is mutually exclusive with the "noJSDataSource" option.
//   type: boolean
//validForSecureReplication ==>
//   description: Restrict listing to include only datasets that can be securely replicated.
//   type: boolean
//provisionContainer ==>
//   description: Restrict databases to those provisioned from the specified container. This option is mutually exclusive with the "group" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: provisionContainer
//group ==>
//   description: Restrict databases to those within the specified group. This option is mutually exclusive with the "provisionContainer" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-group.json
//   mapsTo: group
//return []Container
func (c *Client) ContainerList(noJSDataSource *bool, noJSContainerDataSource *bool, validForSecureReplication *bool, provisionContainer string, group string) ([]Container, error) {
	var err error
	var result []Container
	obj, err := c.executeListAndReturnResults("/database")
	for _, i := range obj {
		result = append(result, i.(Container))
	}
	return result, err
}

//ContainerUtilizationList - Reports the utilization of all containers during a particular
// period of time.
//
//fromDate ==>
//   description: The earliest date for which container utilization statistics will be reported.
//   type: string
//   format: date
//   required: true
//toDate ==>
//   description: The latest date for which container utilization statistics will be reported.
//   type: string
//   format: date
//   required: true
//samplingInterval ==>
//   type: number
//   required: true
//   description: The interval at which data is to be sampled, measured in seconds.
//return []*ContainerUtilizationStruct
func (c *Client) ContainerUtilizationList(fromDate string, toDate string, samplingInterval float64) ([]*ContainerUtilizationStruct, error) {
	var err error
	var result []*ContainerUtilizationStruct
	obj, err := c.executeListAndReturnResults("/database/performanceHistory")
	for _, i := range obj {
		result = append(result, i.(*ContainerUtilizationStruct))
	}
	return result, err
}

//CurrentConsumerCapacityDataList - Lists consumers in the system.
//
//return []*CurrentConsumerCapacityDataStruct
func (c *Client) CurrentConsumerCapacityDataList() ([]*CurrentConsumerCapacityDataStruct, error) {
	var err error
	var result []*CurrentConsumerCapacityDataStruct
	obj, err := c.executeListAndReturnResults("/capacity/consumer")
	for _, i := range obj {
		result = append(result, i.(*CurrentConsumerCapacityDataStruct))
	}
	return result, err
}

//CurrentGroupCapacityDataList - Lists capacity data for groups in the system.
//
//return []*CurrentGroupCapacityDataStruct
func (c *Client) CurrentGroupCapacityDataList() ([]*CurrentGroupCapacityDataStruct, error) {
	var err error
	var result []*CurrentGroupCapacityDataStruct
	obj, err := c.executeListAndReturnResults("/capacity/group")
	for _, i := range obj {
		result = append(result, i.(*CurrentGroupCapacityDataStruct))
	}
	return result, err
}

//CurrentSystemCapacityDataRead - Retrieve the specified CurrentSystemCapacityData object.
//
//r - CurrentSystemCapacityData Object Reference
//return *CurrentSystemCapacityDataStruct
func (c *Client) CurrentSystemCapacityDataRead(r string) (*CurrentSystemCapacityDataStruct, error) {
	var err error
	var result *CurrentSystemCapacityDataStruct
	obj, err := c.FindObjectByReference("/capacity/system", r)
	result = obj.(*CurrentSystemCapacityDataStruct)
	return result, err
}

//DNSConfigUpdate - Update the specified DNSConfig object.
//
//r - DNSConfig Object Reference
//payload *DNSConfigStruct object
func (c *Client) DNSConfigUpdate(r string, payload *DNSConfigStruct) error {
	var err error
	return err
}

//DNSConfigRead - Retrieve the specified DNSConfig object.
//
//r - DNSConfig Object Reference
//return *DNSConfigStruct
func (c *Client) DNSConfigRead(r string) (*DNSConfigStruct, error) {
	var err error
	var result *DNSConfigStruct
	obj, err := c.FindObjectByReference("/service/dns", r)
	result = obj.(*DNSConfigStruct)
	return result, err
}

//DatabaseTemplateRead - Retrieve the specified DatabaseTemplate object.
//
//r - DatabaseTemplate Object Reference
//return *DatabaseTemplateStruct
func (c *Client) DatabaseTemplateRead(r string) (*DatabaseTemplateStruct, error) {
	var err error
	var result *DatabaseTemplateStruct
	obj, err := c.FindObjectByReference("/database/template", r)
	result = obj.(*DatabaseTemplateStruct)
	return result, err
}

//DatabaseTemplateDelete - Delete the specified DatabaseTemplate object.
//
//r - DatabaseTemplate Object Reference
func (c *Client) DatabaseTemplateDelete(r string) error {
	var err error
	return err
}

//DatabaseTemplateUpdate - Update the specified DatabaseTemplate object.
//
//r - DatabaseTemplate Object Reference
//payload *DatabaseTemplateStruct object
func (c *Client) DatabaseTemplateUpdate(r string, payload *DatabaseTemplateStruct) error {
	var err error
	return err
}

//DatabaseTemplateList - List DatabaseTemplate objects on the system.
//
//return []*DatabaseTemplateStruct
func (c *Client) DatabaseTemplateList() ([]*DatabaseTemplateStruct, error) {
	var err error
	var result []*DatabaseTemplateStruct
	obj, err := c.executeListAndReturnResults("/database/template")
	for _, i := range obj {
		result = append(result, i.(*DatabaseTemplateStruct))
	}
	return result, err
}

//DatabaseTemplateCreate - Create a new DatabaseTemplate object.
//
//payload *DatabaseTemplateStruct object
//return string
func (c *Client) DatabaseTemplateCreate(payload *DatabaseTemplateStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//DomainRead - Retrieve the specified Domain object.
//
//r - Domain Object Reference
//return *DomainStruct
func (c *Client) DomainRead(r string) (*DomainStruct, error) {
	var err error
	var result *DomainStruct
	obj, err := c.FindObjectByReference("/domain", r)
	result = obj.(*DomainStruct)
	return result, err
}

//EndEntityCertificateRead - Retrieve the specified EndEntityCertificate object.
//
//r - EndEntityCertificate Object Reference
//return *EndEntityCertificateStruct
func (c *Client) EndEntityCertificateRead(r string) (*EndEntityCertificateStruct, error) {
	var err error
	var result *EndEntityCertificateStruct
	obj, err := c.FindObjectByReference("/service/tls/endEntityCertificate", r)
	result = obj.(*EndEntityCertificateStruct)
	return result, err
}

//EndEntityCertificateList - List EndEntityCertificate objects on the system.
//
//return []*EndEntityCertificateStruct
func (c *Client) EndEntityCertificateList() ([]*EndEntityCertificateStruct, error) {
	var err error
	var result []*EndEntityCertificateStruct
	obj, err := c.executeListAndReturnResults("/service/tls/endEntityCertificate")
	for _, i := range obj {
		result = append(result, i.(*EndEntityCertificateStruct))
	}
	return result, err
}

//EnvironmentUserCreate - Create a new EnvironmentUser object.
//
//payload *EnvironmentUserStruct object
//return string
func (c *Client) EnvironmentUserCreate(payload *EnvironmentUserStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//EnvironmentUserDelete - Delete the specified EnvironmentUser object.
//
//r - EnvironmentUser Object Reference
//payload *DeleteParametersStruct object
func (c *Client) EnvironmentUserDelete(r string, payload *DeleteParametersStruct) error {
	var err error
	return err
}

//EnvironmentUserRead - Retrieve the specified EnvironmentUser object.
//
//r - EnvironmentUser Object Reference
//return *EnvironmentUserStruct
func (c *Client) EnvironmentUserRead(r string) (*EnvironmentUserStruct, error) {
	var err error
	var result *EnvironmentUserStruct
	obj, err := c.FindObjectByReference("/environment/user", r)
	result = obj.(*EnvironmentUserStruct)
	return result, err
}

//EnvironmentUserUpdate - Update the specified EnvironmentUser object.
//
//r - EnvironmentUser Object Reference
//payload *EnvironmentUserStruct object
func (c *Client) EnvironmentUserUpdate(r string, payload *EnvironmentUserStruct) error {
	var err error
	return err
}

//EnvironmentUserList - Returns the list of all environment users in the system.
//
//environment ==>
//   type: string
//   description: Limit results to users within the given environment.
//   format: objectReference
//   referenceTo: /delphix-source-environment.json
//   mapsTo: environment
//return []*EnvironmentUserStruct
func (c *Client) EnvironmentUserList(environment string) ([]*EnvironmentUserStruct, error) {
	var err error
	var result []*EnvironmentUserStruct
	obj, err := c.executeListAndReturnResults("/environment/user")
	for _, i := range obj {
		result = append(result, i.(*EnvironmentUserStruct))
	}
	return result, err
}

//FaultRead - Retrieve the specified Fault object.
//
//r - Fault Object Reference
//return *FaultStruct
func (c *Client) FaultRead(r string) (*FaultStruct, error) {
	var err error
	var result *FaultStruct
	obj, err := c.FindObjectByReference("/fault", r)
	result = obj.(*FaultStruct)
	return result, err
}

//FaultList - Returns the list of all the faults that match the given criteria.
//
//maxTotal ==>
//   description: The upper bound for calculation of total fault count.
//   type: integer
//target ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   mapsTo: target
//   description: The reference to the Delphix user-visible object associated with the fault.
//severity ==>
//   description: The impact of the fault on the system: CRITICAL or WARNING.
//   type: string
//   enum: [CRITICAL WARNING]
//   mapsTo: severity
//status ==>
//   enum: [ACTIVE RESOLVED IGNORED]
//   mapsTo: status
//   description: The status of the fault: ACTIVE, RESOLVED or IGNORED.
//   type: string
//fromDate ==>
//   description: Start date to use for the search.
//   type: string
//   format: date
//   mapsTo: dateDiagnosed
//   inequalityType: NON-STRICT
//toDate ==>
//   description: End date to use for the search.
//   type: string
//   format: date
//   mapsTo: dateDiagnosed
//   inequalityType: STRICT
//pageSize ==>
//   description: Limit the number of faults returned.
//   type: integer
//   default: 25
//   minimum: 0
//pageOffset ==>
//   description: Offset within fault list, in units of pageSize chunks.
//   type: integer
//sortBy ==>
//   description: Search results are sorted by the field provided.
//   type: string
//   enum: [severity dateDiagnosed title targetName]
//ascending ==>
//   description: True if results are to be returned in ascending order.
//   type: boolean
//searchText ==>
//   description: Limit search results to only include faults that have searchText string in severity, title, targetName.
//   type: string
//return []*FaultStruct
func (c *Client) FaultList(maxTotal *int, target string, severity string, status string, fromDate string, toDate string, pageSize *int, pageOffset *int, sortBy string, ascending *bool, searchText string) ([]*FaultStruct, error) {
	var err error
	var result []*FaultStruct
	obj, err := c.executeListAndReturnResults("/fault")
	for _, i := range obj {
		result = append(result, i.(*FaultStruct))
	}
	return result, err
}

//FaultEffectRead - Retrieve the specified FaultEffect object.
//
//r - FaultEffect Object Reference
//return *FaultEffectStruct
func (c *Client) FaultEffectRead(r string) (*FaultEffectStruct, error) {
	var err error
	var result *FaultEffectStruct
	obj, err := c.FindObjectByReference("/fault/effect", r)
	result = obj.(*FaultEffectStruct)
	return result, err
}

//FaultEffectList - Returns the list of all the fault effects that match the given
// criteria.
//
//severity ==>
//   description: The impact of the fault effect on the system: CRITICAL or WARNING.
//   type: string
//   enum: [CRITICAL WARNING]
//   mapsTo: severity
//target ==>
//   description: The reference to the Delphix user-visible object associated with the fault effect.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   mapsTo: target
//rootCause ==>
//   description: The reference to the fault which is the root cause of the fault effect.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-fault.json
//   mapsTo: rootCause
//bundleID ==>
//   description: A unique dot delimited identifier associated with the fault effect.
//   type: string
//   mapsTo: bundleID
//return []*FaultEffectStruct
func (c *Client) FaultEffectList(severity string, target string, rootCause string, bundleID string) ([]*FaultEffectStruct, error) {
	var err error
	var result []*FaultEffectStruct
	obj, err := c.executeListAndReturnResults("/fault/effect")
	for _, i := range obj {
		result = append(result, i.(*FaultEffectStruct))
	}
	return result, err
}

//GlobalLinkingSettingsRead - Retrieve the specified GlobalLinkingSettings object.
//
//r - GlobalLinkingSettings Object Reference
//return *GlobalLinkingSettingsStruct
func (c *Client) GlobalLinkingSettingsRead(r string) (*GlobalLinkingSettingsStruct, error) {
	var err error
	var result *GlobalLinkingSettingsStruct
	obj, err := c.FindObjectByReference("/service/linkingsettings", r)
	result = obj.(*GlobalLinkingSettingsStruct)
	return result, err
}

//GlobalLinkingSettingsUpdate - Update the specified GlobalLinkingSettings object.
//
//r - GlobalLinkingSettings Object Reference
//payload *GlobalLinkingSettingsStruct object
func (c *Client) GlobalLinkingSettingsUpdate(r string, payload *GlobalLinkingSettingsStruct) error {
	var err error
	return err
}

//GroupRead - Retrieve the specified Group object.
//
//r - Group Object Reference
//return *GroupStruct
func (c *Client) GroupRead(r string) (*GroupStruct, error) {
	var err error
	var result *GroupStruct
	obj, err := c.FindObjectByReference("/group", r)
	result = obj.(*GroupStruct)
	return result, err
}

//GroupCreate - Create a new Group object.
//
//payload *GroupStruct object
//return string
func (c *Client) GroupCreate(payload *GroupStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//GroupList - List Group objects on the system.
//
//return []*GroupStruct
func (c *Client) GroupList() ([]*GroupStruct, error) {
	var err error
	var result []*GroupStruct
	obj, err := c.executeListAndReturnResults("/group")
	for _, i := range obj {
		result = append(result, i.(*GroupStruct))
	}
	return result, err
}

//GroupUpdate - Update the specified Group object.
//
//r - Group Object Reference
//payload *GroupStruct object
func (c *Client) GroupUpdate(r string, payload *GroupStruct) error {
	var err error
	return err
}

//GroupDelete - Delete the specified Group object.
//
//r - Group Object Reference
func (c *Client) GroupDelete(r string) error {
	var err error
	return err
}

//HistoricalConsumerCapacityDataList - Lists consumers in the system.
//
//container ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   description: The container for which to list data.
//startDate ==>
//   type: string
//   format: date
//   description: Earliest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//endDate ==>
//   inequalityType: NON-STRICT
//   type: string
//   format: date
//   description: Latest date for which to list data.
//   mapsTo: timestamp
//resolution ==>
//   type: integer
//   minimum: 0
//   description: The time range each datapoint should represent, measured in seconds. This parameter is only meaningful if a container is specified.
//return []*HistoricalConsumerCapacityDataStruct
func (c *Client) HistoricalConsumerCapacityDataList(container string, startDate string, endDate string, resolution *int) ([]*HistoricalConsumerCapacityDataStruct, error) {
	var err error
	var result []*HistoricalConsumerCapacityDataStruct
	obj, err := c.executeListAndReturnResults("/capacity/consumer/historical")
	for _, i := range obj {
		result = append(result, i.(*HistoricalConsumerCapacityDataStruct))
	}
	return result, err
}

//HistoricalGroupCapacityDataList - Lists historical capacity data for a particular group.
//
//group ==>
//   description: The group to list data for.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-group.json
//startDate ==>
//   type: string
//   format: date
//   description: Earliest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//endDate ==>
//   description: Latest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//   type: string
//   format: date
//resolution ==>
//   type: integer
//   minimum: 0
//   description: The time range each datapoint should represent, measured in seconds. This parameter is only meaningful if a group is specified.
//return []*HistoricalGroupCapacityDataStruct
func (c *Client) HistoricalGroupCapacityDataList(group string, startDate string, endDate string, resolution *int) ([]*HistoricalGroupCapacityDataStruct, error) {
	var err error
	var result []*HistoricalGroupCapacityDataStruct
	obj, err := c.executeListAndReturnResults("/capacity/group/historical")
	for _, i := range obj {
		result = append(result, i.(*HistoricalGroupCapacityDataStruct))
	}
	return result, err
}

//HistoricalSystemCapacityDataList - Lists historical system capacity data.
//
//startDate ==>
//   format: date
//   description: Earliest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//   type: string
//endDate ==>
//   type: string
//   format: date
//   description: Latest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//resolution ==>
//   type: integer
//   minimum: 0
//   description: The time range each datapoint should represent, measured in seconds.
//return []*HistoricalSystemCapacityDataStruct
func (c *Client) HistoricalSystemCapacityDataList(startDate string, endDate string, resolution *int) ([]*HistoricalSystemCapacityDataStruct, error) {
	var err error
	var result []*HistoricalSystemCapacityDataStruct
	obj, err := c.executeListAndReturnResults("/capacity/system/historical")
	for _, i := range obj {
		result = append(result, i.(*HistoricalSystemCapacityDataStruct))
	}
	return result, err
}

//HostList - Returns the list of all hosts in the system.
//
//environment ==>
//   type: string
//   description: Include only hosts belonging to the given environment.
//   format: objectReference
//   referenceTo: /delphix-source-environment.json
//return []Host
func (c *Client) HostList(environment string) ([]Host, error) {
	var err error
	var result []Host
	obj, err := c.executeListAndReturnResults("/host")
	for _, i := range obj {
		result = append(result, i.(Host))
	}
	return result, err
}

//HostRead - Retrieve the specified Host object.
//
//r - Host Object Reference
//return Host
func (c *Client) HostRead(r string) (Host, error) {
	var err error
	var result Host
	obj, err := c.FindObjectByReference("/host", r)
	result = obj.(Host)
	return result, err
}

//HostUpdate - Update the specified Host object.
//
//r - Host Object Reference
//payload Host object
func (c *Client) HostUpdate(r string, payload Host) error {
	var err error
	return err
}

//HostPrivilegeElevationProfileList - List HostPrivilegeElevationProfile objects on the system.
//
//return []*HostPrivilegeElevationProfileStruct
func (c *Client) HostPrivilegeElevationProfileList() ([]*HostPrivilegeElevationProfileStruct, error) {
	var err error
	var result []*HostPrivilegeElevationProfileStruct
	obj, err := c.executeListAndReturnResults("/host/privilegeElevation/profile")
	for _, i := range obj {
		result = append(result, i.(*HostPrivilegeElevationProfileStruct))
	}
	return result, err
}

//HostPrivilegeElevationProfileRead - Retrieve the specified HostPrivilegeElevationProfile object.
//
//r - HostPrivilegeElevationProfile Object Reference
//return *HostPrivilegeElevationProfileStruct
func (c *Client) HostPrivilegeElevationProfileRead(r string) (*HostPrivilegeElevationProfileStruct, error) {
	var err error
	var result *HostPrivilegeElevationProfileStruct
	obj, err := c.FindObjectByReference("/host/privilegeElevation/profile", r)
	result = obj.(*HostPrivilegeElevationProfileStruct)
	return result, err
}

//HostPrivilegeElevationProfileCreate - Create a new HostPrivilegeElevationProfile object.
//
//payload *HostPrivilegeElevationProfileStruct object
//return string
func (c *Client) HostPrivilegeElevationProfileCreate(payload *HostPrivilegeElevationProfileStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//HostPrivilegeElevationProfileUpdate - Update the specified HostPrivilegeElevationProfile object.
//
//r - HostPrivilegeElevationProfile Object Reference
//payload *HostPrivilegeElevationProfileStruct object
func (c *Client) HostPrivilegeElevationProfileUpdate(r string, payload *HostPrivilegeElevationProfileStruct) error {
	var err error
	return err
}

//HostPrivilegeElevationProfileDelete - Delete the specified HostPrivilegeElevationProfile object.
//
//r - HostPrivilegeElevationProfile Object Reference
func (c *Client) HostPrivilegeElevationProfileDelete(r string) error {
	var err error
	return err
}

//HostPrivilegeElevationProfileScriptRead - Retrieve the specified HostPrivilegeElevationProfileScript object.
//
//r - HostPrivilegeElevationProfileScript Object Reference
//return *HostPrivilegeElevationProfileScriptStruct
func (c *Client) HostPrivilegeElevationProfileScriptRead(r string) (*HostPrivilegeElevationProfileScriptStruct, error) {
	var err error
	var result *HostPrivilegeElevationProfileScriptStruct
	obj, err := c.FindObjectByReference("/host/privilegeElevation/profileScript", r)
	result = obj.(*HostPrivilegeElevationProfileScriptStruct)
	return result, err
}

//HostPrivilegeElevationProfileScriptUpdate - Update the specified HostPrivilegeElevationProfileScript object.
//
//r - HostPrivilegeElevationProfileScript Object Reference
//payload *HostPrivilegeElevationProfileScriptStruct object
func (c *Client) HostPrivilegeElevationProfileScriptUpdate(r string, payload *HostPrivilegeElevationProfileScriptStruct) error {
	var err error
	return err
}

//HostPrivilegeElevationProfileScriptList - List HostPrivilegeElevationProfileScript objects on the system.
//
//return []*HostPrivilegeElevationProfileScriptStruct
func (c *Client) HostPrivilegeElevationProfileScriptList() ([]*HostPrivilegeElevationProfileScriptStruct, error) {
	var err error
	var result []*HostPrivilegeElevationProfileScriptStruct
	obj, err := c.executeListAndReturnResults("/host/privilegeElevation/profileScript")
	for _, i := range obj {
		result = append(result, i.(*HostPrivilegeElevationProfileScriptStruct))
	}
	return result, err
}

//HostPrivilegeElevationProfileScriptCreate - Create a new HostPrivilegeElevationProfileScript object.
//
//payload *HostPrivilegeElevationProfileScriptStruct object
//return string
func (c *Client) HostPrivilegeElevationProfileScriptCreate(payload *HostPrivilegeElevationProfileScriptStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//HostPrivilegeElevationProfileScriptDelete - Delete the specified HostPrivilegeElevationProfileScript object.
//
//r - HostPrivilegeElevationProfileScript Object Reference
func (c *Client) HostPrivilegeElevationProfileScriptDelete(r string) error {
	var err error
	return err
}

//HostPrivilegeElevationSettingsRead - Retrieve the specified HostPrivilegeElevationSettings object.
//
//r - HostPrivilegeElevationSettings Object Reference
//return *HostPrivilegeElevationSettingsStruct
func (c *Client) HostPrivilegeElevationSettingsRead(r string) (*HostPrivilegeElevationSettingsStruct, error) {
	var err error
	var result *HostPrivilegeElevationSettingsStruct
	obj, err := c.FindObjectByReference("/host/privilegeElevation", r)
	result = obj.(*HostPrivilegeElevationSettingsStruct)
	return result, err
}

//HostPrivilegeElevationSettingsUpdate - Update the specified HostPrivilegeElevationSettings object.
//
//r - HostPrivilegeElevationSettings Object Reference
//payload *HostPrivilegeElevationSettingsStruct object
func (c *Client) HostPrivilegeElevationSettingsUpdate(r string, payload *HostPrivilegeElevationSettingsStruct) error {
	var err error
	return err
}

//HttpConnectorConfigRead - Retrieve the specified HttpConnectorConfig object.
//
//r - HttpConnectorConfig Object Reference
//return *HttpConnectorConfigStruct
func (c *Client) HttpConnectorConfigRead(r string) (*HttpConnectorConfigStruct, error) {
	var err error
	var result *HttpConnectorConfigStruct
	obj, err := c.FindObjectByReference("/service/httpConnector", r)
	result = obj.(*HttpConnectorConfigStruct)
	return result, err
}

//HttpConnectorConfigUpdate - Update the specified HttpConnectorConfig object.
//
//r - HttpConnectorConfig Object Reference
//payload *HttpConnectorConfigStruct object
func (c *Client) HttpConnectorConfigUpdate(r string, payload *HttpConnectorConfigStruct) error {
	var err error
	return err
}

//JSBookmarkList - Lists the Self-Service bookmarks in the system.
//
//container ==>
//   description: List all usable bookmarks accessible to the specified data container. This option is mutually exclusive with all other options.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-container.json
//   mapsTo: container
//template ==>
//   description: List all usable bookmarks created in the specified data template. This option is mutually exclusive with all other options.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-template.json
//   mapsTo: template
//return []*JSBookmarkStruct
func (c *Client) JSBookmarkList(container string, template string) ([]*JSBookmarkStruct, error) {
	var err error
	var result []*JSBookmarkStruct
	obj, err := c.executeListAndReturnResults("/selfservice/bookmark")
	for _, i := range obj {
		result = append(result, i.(*JSBookmarkStruct))
	}
	return result, err
}

//JSBookmarkCreate - Create a new JSBookmark object.
//
//payload *JSBookmarkCreateParametersStruct object
//return string
func (c *Client) JSBookmarkCreate(payload *JSBookmarkCreateParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//JSBookmarkRead - Retrieve the specified JSBookmark object.
//
//r - JSBookmark Object Reference
//return *JSBookmarkStruct
func (c *Client) JSBookmarkRead(r string) (*JSBookmarkStruct, error) {
	var err error
	var result *JSBookmarkStruct
	obj, err := c.FindObjectByReference("/selfservice/bookmark", r)
	result = obj.(*JSBookmarkStruct)
	return result, err
}

//JSBookmarkUpdate - Update the specified JSBookmark object.
//
//r - JSBookmark Object Reference
//payload *JSBookmarkStruct object
func (c *Client) JSBookmarkUpdate(r string, payload *JSBookmarkStruct) error {
	var err error
	return err
}

//JSBookmarkDelete - Delete the specified JSBookmark object.
//
//r - JSBookmark Object Reference
func (c *Client) JSBookmarkDelete(r string) error {
	var err error
	return err
}

//JSBranchCreate - Create a new JSBranch object.
//
//payload *JSBranchCreateParametersStruct object
//return string
func (c *Client) JSBranchCreate(payload *JSBranchCreateParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//JSBranchRead - Retrieve the specified JSBranch object.
//
//r - JSBranch Object Reference
//return *JSBranchStruct
func (c *Client) JSBranchRead(r string) (*JSBranchStruct, error) {
	var err error
	var result *JSBranchStruct
	obj, err := c.FindObjectByReference("/selfservice/branch", r)
	result = obj.(*JSBranchStruct)
	return result, err
}

//JSBranchDelete - Delete the specified JSBranch object.
//
//r - JSBranch Object Reference
func (c *Client) JSBranchDelete(r string) error {
	var err error
	return err
}

//JSBranchUpdate - Update the specified JSBranch object.
//
//r - JSBranch Object Reference
//payload *JSBranchStruct object
func (c *Client) JSBranchUpdate(r string, payload *JSBranchStruct) error {
	var err error
	return err
}

//JSBranchList - Lists the Self-Service branches in the system.
//
//dataLayout ==>
//   format: objectReference
//   referenceTo: /delphix-js-data-layout.json
//   mapsTo: dataLayout
//   description: List branches belonging to the given data layout.
//   type: string
//return []*JSBranchStruct
func (c *Client) JSBranchList(dataLayout string) ([]*JSBranchStruct, error) {
	var err error
	var result []*JSBranchStruct
	obj, err := c.executeListAndReturnResults("/selfservice/branch")
	for _, i := range obj {
		result = append(result, i.(*JSBranchStruct))
	}
	return result, err
}

//JSConfigRead - Retrieve the specified JSConfig object.
//
//r - JSConfig Object Reference
//return *JSConfigStruct
func (c *Client) JSConfigRead(r string) (*JSConfigStruct, error) {
	var err error
	var result *JSConfigStruct
	obj, err := c.FindObjectByReference("/selfservice/config", r)
	result = obj.(*JSConfigStruct)
	return result, err
}

//JSConfigUpdate - Update the specified JSConfig object.
//
//r - JSConfig Object Reference
//payload *JSConfigStruct object
func (c *Client) JSConfigUpdate(r string, payload *JSConfigStruct) error {
	var err error
	return err
}

//JSDailyOperationDurationRead - Retrieve the specified JSDailyOperationDuration object.
//
//r - JSDailyOperationDuration Object Reference
//return *JSDailyOperationDurationStruct
func (c *Client) JSDailyOperationDurationRead(r string) (*JSDailyOperationDurationStruct, error) {
	var err error
	var result *JSDailyOperationDurationStruct
	obj, err := c.FindObjectByReference("/selfservice/usagedata/operationduration", r)
	result = obj.(*JSDailyOperationDurationStruct)
	return result, err
}

//JSDailyOperationDurationList - List the operation durations in the system, optionally restricted
// to those operations related to a single object.
//
//usageObject ==>
//   description: Restrict usage data to that related to a specific object.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-named-user-object.json
//   mapsTo: usageObject
//return []*JSDailyOperationDurationStruct
func (c *Client) JSDailyOperationDurationList(usageObject string) ([]*JSDailyOperationDurationStruct, error) {
	var err error
	var result []*JSDailyOperationDurationStruct
	obj, err := c.executeListAndReturnResults("/selfservice/usagedata/operationduration")
	for _, i := range obj {
		result = append(result, i.(*JSDailyOperationDurationStruct))
	}
	return result, err
}

//JSDataContainerDelete - Delete this data container.
//
//r - JSDataContainer Object Reference
//payload *JSDataContainerDeleteParametersStruct object
func (c *Client) JSDataContainerDelete(r string, payload *JSDataContainerDeleteParametersStruct) error {
	var err error
	return err
}

//JSDataContainerCreate - Create a new JSDataContainer object.
//
//payload JSDataContainerCreateParameters object
//return string
func (c *Client) JSDataContainerCreate(payload JSDataContainerCreateParameters) (string, error) {
	var err error
	var result string
	return result, err
}

//JSDataContainerUpdate - Update the specified JSDataContainer object.
//
//r - JSDataContainer Object Reference
//payload *JSDataContainerStruct object
func (c *Client) JSDataContainerUpdate(r string, payload *JSDataContainerStruct) error {
	var err error
	return err
}

//JSDataContainerRead - Retrieve the specified JSDataContainer object.
//
//r - JSDataContainer Object Reference
//return *JSDataContainerStruct
func (c *Client) JSDataContainerRead(r string) (*JSDataContainerStruct, error) {
	var err error
	var result *JSDataContainerStruct
	obj, err := c.FindObjectByReference("/selfservice/container", r)
	result = obj.(*JSDataContainerStruct)
	return result, err
}

//JSDataContainerList - List the data containers defined in the system.
//
//owner ==>
//   description: Restrict data containers to those belonging to the specified user.This option is mutually exclusive with the "template" and "independentOnly" options.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user.json
//   mapsTo: owner
//template ==>
//   description: Restrict data containers to those provisioned from the specified template. This option is mutually exclusive with the "owner" and "independentOnly" options.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-template.json
//   mapsTo: template
//independentOnly ==>
//   description: Restrict data containers to independent data containers that do not have templates. This option is mutually exclusive with the "template" and "owner" options.
//   type: boolean
//return []*JSDataContainerStruct
func (c *Client) JSDataContainerList(owner string, template string, independentOnly *bool) ([]*JSDataContainerStruct, error) {
	var err error
	var result []*JSDataContainerStruct
	obj, err := c.executeListAndReturnResults("/selfservice/container")
	for _, i := range obj {
		result = append(result, i.(*JSDataContainerStruct))
	}
	return result, err
}

//JSDataSourceRead - Retrieve the specified JSDataSource object.
//
//r - JSDataSource Object Reference
//return *JSDataSourceStruct
func (c *Client) JSDataSourceRead(r string) (*JSDataSourceStruct, error) {
	var err error
	var result *JSDataSourceStruct
	obj, err := c.FindObjectByReference("/selfservice/datasource", r)
	result = obj.(*JSDataSourceStruct)
	return result, err
}

//JSDataSourceUpdate - Update the specified JSDataSource object.
//
//r - JSDataSource Object Reference
//payload *JSDataSourceStruct object
func (c *Client) JSDataSourceUpdate(r string, payload *JSDataSourceStruct) error {
	var err error
	return err
}

//JSDataSourceList - Lists the Self-Service data sources in the system.
//
//dataLayout ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-layout.json
//   mapsTo: dataLayout
//   description: List the sources associated with the given data layout reference.
//container ==>
//   description: List the source associated with the given container reference.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//return []*JSDataSourceStruct
func (c *Client) JSDataSourceList(dataLayout string, container string) ([]*JSDataSourceStruct, error) {
	var err error
	var result []*JSDataSourceStruct
	obj, err := c.executeListAndReturnResults("/selfservice/datasource")
	for _, i := range obj {
		result = append(result, i.(*JSDataSourceStruct))
	}
	return result, err
}

//JSDataTemplateCreate - Create a new JSDataTemplate object.
//
//payload *JSDataTemplateCreateParametersStruct object
//return string
func (c *Client) JSDataTemplateCreate(payload *JSDataTemplateCreateParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//JSDataTemplateDelete - Delete the specified JSDataTemplate object.
//
//r - JSDataTemplate Object Reference
func (c *Client) JSDataTemplateDelete(r string) error {
	var err error
	return err
}

//JSDataTemplateList - List the data templates defined in the system.
//
//return []*JSDataTemplateStruct
func (c *Client) JSDataTemplateList() ([]*JSDataTemplateStruct, error) {
	var err error
	var result []*JSDataTemplateStruct
	obj, err := c.executeListAndReturnResults("/selfservice/template")
	for _, i := range obj {
		result = append(result, i.(*JSDataTemplateStruct))
	}
	return result, err
}

//JSDataTemplateRead - Retrieve the specified JSDataTemplate object.
//
//r - JSDataTemplate Object Reference
//return *JSDataTemplateStruct
func (c *Client) JSDataTemplateRead(r string) (*JSDataTemplateStruct, error) {
	var err error
	var result *JSDataTemplateStruct
	obj, err := c.FindObjectByReference("/selfservice/template", r)
	result = obj.(*JSDataTemplateStruct)
	return result, err
}

//JSDataTemplateUpdate - Update the specified JSDataTemplate object.
//
//r - JSDataTemplate Object Reference
//payload *JSDataTemplateStruct object
func (c *Client) JSDataTemplateUpdate(r string, payload *JSDataTemplateStruct) error {
	var err error
	return err
}

//JSOperationRead - Retrieve the specified JSOperation object.
//
//r - JSOperation Object Reference
//return *JSOperationStruct
func (c *Client) JSOperationRead(r string) (*JSOperationStruct, error) {
	var err error
	var result *JSOperationStruct
	obj, err := c.FindObjectByReference("/selfservice/operation", r)
	result = obj.(*JSOperationStruct)
	return result, err
}

//JSOperationList - Lists the Self-Service action history for a data layout.
//
//dataLayout ==>
//   description: Limit operations to the specific data layout. This option is mutually exclusive with the "branch" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-layout.json
//   mapsTo: dataLayout
//branch ==>
//   referenceTo: /delphix-js-branch.json
//   mapsTo: branch
//   description: Limit operations to the specified branch. This option is mutually exclusive with the "dataLayout" option.
//   type: string
//   format: objectReference
//dataTime ==>
//   description: Limit operations that occurred around the specified "dataTime". "beforeCount" and "afterCount" should specify the number of events to be returned.
//   type: string
//   format: date
//   mapsTo: dataTime
//beforeCount ==>
//   description: The suggested maximum number of visible operations prior to "dataTime" that should be returned. If there are not sufficient events before additional events after may be returned.
//   type: integer
//   default: 0
//afterCount ==>
//   description: The suggested maximum number of visible operations after "dataTime" that should be returned. If there are not sufficient events after additional events before may be returned.
//   type: integer
//   default: 0
//dataStartTime ==>
//   format: date
//   description: Operations with "dataTime" after this value will be returned. Used with "dataEndTime" to return a set of operations between two dates.
//   type: string
//dataEndTime ==>
//   description: Operations with "dataTime" before this value will be returned. Used with "dataStartTime" to return a set of operations between two dates.
//   type: string
//   format: date
//return []*JSOperationStruct
func (c *Client) JSOperationList(dataLayout string, branch string, dataTime string, beforeCount *int, afterCount *int, dataStartTime string, dataEndTime string) ([]*JSOperationStruct, error) {
	var err error
	var result []*JSOperationStruct
	obj, err := c.executeListAndReturnResults("/selfservice/operation")
	for _, i := range obj {
		result = append(result, i.(*JSOperationStruct))
	}
	return result, err
}

//JSWeeklyOperationCountRead - Retrieve the specified JSWeeklyOperationCount object.
//
//r - JSWeeklyOperationCount Object Reference
//return *JSWeeklyOperationCountStruct
func (c *Client) JSWeeklyOperationCountRead(r string) (*JSWeeklyOperationCountStruct, error) {
	var err error
	var result *JSWeeklyOperationCountStruct
	obj, err := c.FindObjectByReference("/selfservice/usagedata/operationcount", r)
	result = obj.(*JSWeeklyOperationCountStruct)
	return result, err
}

//JSWeeklyOperationCountList - List the operation counts in the system, optionally restricted to
// those operations related to a single object.
//
//usageObject ==>
//   description: Restrict usage data to that related to a specific object.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-named-user-object.json
//   mapsTo: usageObject
//return []*JSWeeklyOperationCountStruct
func (c *Client) JSWeeklyOperationCountList(usageObject string) ([]*JSWeeklyOperationCountStruct, error) {
	var err error
	var result []*JSWeeklyOperationCountStruct
	obj, err := c.executeListAndReturnResults("/selfservice/usagedata/operationcount")
	for _, i := range obj {
		result = append(result, i.(*JSWeeklyOperationCountStruct))
	}
	return result, err
}

//JobRead - Retrieve the specified Job object.
//
//r - Job Object Reference
//return *JobStruct
func (c *Client) JobRead(r string) (*JobStruct, error) {
	var err error
	var result *JobStruct
	obj, err := c.FindObjectByReference("/job", r)
	result = obj.(*JobStruct)
	return result, err
}

//JobList - Returns a list of jobs in the system. Jobs are listed in start
// time order.
//
//searchText ==>
//   type: string
//   description: Limit search results to only include jobs that contain the searchText.
//jobType ==>
//   description: Limit jobs to those with the specified job type.
//   type: string
//   mapsTo: actionType
//pageOffset ==>
//   type: integer
//   description: Page offset within job list.
//addEvents ==>
//   description: Whether to include the job events in each job.
//   type: boolean
//   default: false
//toDate ==>
//   inequalityType: NON-STRICT
//   description: Filters out jobs newer than this date.
//   format: date
//   type: string
//   mapsTo: startTime
//pageSize ==>
//   description: Limit the number of jobs returned.
//   type: integer
//   default: 25
//   minimum: 0
//maxTotal ==>
//   description: The upper bound for calculation of total job count.
//   type: integer
//jobState ==>
//   mapsTo: jobState
//   description: Limit jobs to those in the specified job state.
//   type: string
//   enum: [RUNNING SUSPENDED CANCELED COMPLETED FAILED]
//target ==>
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   mapsTo: target
//   description: Limit jobs to those affecting a particular object on the system. The target is the object reference for the target in question.
//   type: string
//fromDate ==>
//   mapsTo: updateTime
//   inequalityType: NON-STRICT
//   description: Filters out jobs older than this date.
//   format: date
//   type: string
//return []*JobStruct
func (c *Client) JobList(searchText string, jobType string, pageOffset *int, addEvents *bool, toDate string, pageSize *int, maxTotal *int, jobState string, target string, fromDate string) ([]*JobStruct, error) {
	var err error
	var result []*JobStruct
	obj, err := c.executeListAndReturnResults("/job")
	for _, i := range obj {
		result = append(result, i.(*JobStruct))
	}
	return result, err
}

//JobUpdate - Update the specified Job object.
//
//r - Job Object Reference
//payload *JobStruct object
func (c *Client) JobUpdate(r string, payload *JobStruct) error {
	var err error
	return err
}

//KerberosConfigUpdate - Update the specified KerberosConfig object.
//
//r - KerberosConfig Object Reference
//payload *KerberosConfigStruct object
func (c *Client) KerberosConfigUpdate(r string, payload *KerberosConfigStruct) error {
	var err error
	return err
}

//KerberosConfigRead - Retrieve the specified KerberosConfig object.
//
//r - KerberosConfig Object Reference
//return *KerberosConfigStruct
func (c *Client) KerberosConfigRead(r string) (*KerberosConfigStruct, error) {
	var err error
	var result *KerberosConfigStruct
	obj, err := c.FindObjectByReference("/service/kerberos", r)
	result = obj.(*KerberosConfigStruct)
	return result, err
}

//LdapInfoRead - Retrieve the specified LdapInfo object.
//
//r - LdapInfo Object Reference
//return *LdapInfoStruct
func (c *Client) LdapInfoRead(r string) (*LdapInfoStruct, error) {
	var err error
	var result *LdapInfoStruct
	obj, err := c.FindObjectByReference("/service/ldap", r)
	result = obj.(*LdapInfoStruct)
	return result, err
}

//LdapServerCreate - Create a new LdapServer object.
//
//payload *LdapServerStruct object
//return string
func (c *Client) LdapServerCreate(payload *LdapServerStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//LdapServerList - List LdapServer objects on the system.
//
//return []*LdapServerStruct
func (c *Client) LdapServerList() ([]*LdapServerStruct, error) {
	var err error
	var result []*LdapServerStruct
	obj, err := c.executeListAndReturnResults("/service/ldap/server")
	for _, i := range obj {
		result = append(result, i.(*LdapServerStruct))
	}
	return result, err
}

//LdapServerRead - Retrieve the specified LdapServer object.
//
//r - LdapServer Object Reference
//return *LdapServerStruct
func (c *Client) LdapServerRead(r string) (*LdapServerStruct, error) {
	var err error
	var result *LdapServerStruct
	obj, err := c.FindObjectByReference("/service/ldap/server", r)
	result = obj.(*LdapServerStruct)
	return result, err
}

//LdapServerDelete - Delete the specified LdapServer object.
//
//r - LdapServer Object Reference
func (c *Client) LdapServerDelete(r string) error {
	var err error
	return err
}

//LdapServerUpdate - Update the specified LdapServer object.
//
//r - LdapServer Object Reference
//payload *LdapServerStruct object
func (c *Client) LdapServerUpdate(r string, payload *LdapServerStruct) error {
	var err error
	return err
}

//LocaleSettingsUpdate - Update the specified LocaleSettings object.
//
//r - LocaleSettings Object Reference
//payload *LocaleSettingsStruct object
func (c *Client) LocaleSettingsUpdate(r string, payload *LocaleSettingsStruct) error {
	var err error
	return err
}

//LocaleSettingsRead - Retrieve the specified LocaleSettings object.
//
//r - LocaleSettings Object Reference
//return *LocaleSettingsStruct
func (c *Client) LocaleSettingsRead(r string) (*LocaleSettingsStruct, error) {
	var err error
	var result *LocaleSettingsStruct
	obj, err := c.FindObjectByReference("/service/locale", r)
	result = obj.(*LocaleSettingsStruct)
	return result, err
}

//MSSqlAvailabilityGroupListenerList - Returns a list of listeners filtered by SQL Server Availability
// Group.
//
//availabilitygroup ==>
//   description: A reference to the SQL Server Availability Group this listener belongs to.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-mssql-availability-group.json
//return []*MSSqlAvailabilityGroupListenerStruct
func (c *Client) MSSqlAvailabilityGroupListenerList(availabilitygroup string) ([]*MSSqlAvailabilityGroupListenerStruct, error) {
	var err error
	var result []*MSSqlAvailabilityGroupListenerStruct
	obj, err := c.executeListAndReturnResults("/environment/windows/availabilitygrouplistener")
	for _, i := range obj {
		result = append(result, i.(*MSSqlAvailabilityGroupListenerStruct))
	}
	return result, err
}

//MSSqlClusterInstanceList - Returns a list of instances filtered by SQL Server Availability
// Group.
//
//availabilitygroup ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-mssql-availability-group.json
//   description: A reference to the SQL Server Availability Group this instance belongs to.
//return []*MSSqlClusterInstanceStruct
func (c *Client) MSSqlClusterInstanceList(availabilitygroup string) ([]*MSSqlClusterInstanceStruct, error) {
	var err error
	var result []*MSSqlClusterInstanceStruct
	obj, err := c.executeListAndReturnResults("/environment/windows/clusterinstance")
	for _, i := range obj {
		result = append(result, i.(*MSSqlClusterInstanceStruct))
	}
	return result, err
}

//MSSqlFailoverClusterInstanceList - Returns a list of instances filtered by SQL Server Failover
// Cluster.
//
//repository ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-mssql-failover-cluster-repository.json
//   description: A reference to the SQL Server Failover Cluster repository this instance belongs to.
//return []*MSSqlFailoverClusterInstanceStruct
func (c *Client) MSSqlFailoverClusterInstanceList(repository string) ([]*MSSqlFailoverClusterInstanceStruct, error) {
	var err error
	var result []*MSSqlFailoverClusterInstanceStruct
	obj, err := c.executeListAndReturnResults("/environment/windows/failoverclusterrepositoryinstance")
	for _, i := range obj {
		result = append(result, i.(*MSSqlFailoverClusterInstanceStruct))
	}
	return result, err
}

//MSSqlFailoverClusterListenerList - Returns a list of listeners filtered by SQL Server Failover
// Cluster.
//
//repository ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-mssql-failover-cluster-repository.json
//   description: A reference to the SQL Server Failover Cluster repository this listener belongs to.
//return []*MSSqlFailoverClusterListenerStruct
func (c *Client) MSSqlFailoverClusterListenerList(repository string) ([]*MSSqlFailoverClusterListenerStruct, error) {
	var err error
	var result []*MSSqlFailoverClusterListenerStruct
	obj, err := c.executeListAndReturnResults("/environment/windows/failoverclusterrepositorylistener")
	for _, i := range obj {
		result = append(result, i.(*MSSqlFailoverClusterListenerStruct))
	}
	return result, err
}

//MaskingJobRead - Retrieve the specified MaskingJob object.
//
//r - MaskingJob Object Reference
//return *MaskingJobStruct
func (c *Client) MaskingJobRead(r string) (*MaskingJobStruct, error) {
	var err error
	var result *MaskingJobStruct
	obj, err := c.FindObjectByReference("/maskingjob", r)
	result = obj.(*MaskingJobStruct)
	return result, err
}

//MaskingJobList - Returns a list of all Masking Jobs on the system.
//
//container ==>
//   required: false
//   description: List only the Masking Jobs that are associated with the provided container.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//return []*MaskingJobStruct
func (c *Client) MaskingJobList(container string) ([]*MaskingJobStruct, error) {
	var err error
	var result []*MaskingJobStruct
	obj, err := c.executeListAndReturnResults("/maskingjob")
	for _, i := range obj {
		result = append(result, i.(*MaskingJobStruct))
	}
	return result, err
}

//MaskingJobUpdate - Update the specified MaskingJob object.
//
//r - MaskingJob Object Reference
//payload *MaskingJobStruct object
func (c *Client) MaskingJobUpdate(r string, payload *MaskingJobStruct) error {
	var err error
	return err
}

//MaskingJobDelete - Delete the specified MaskingJob object.
//
//r - MaskingJob Object Reference
func (c *Client) MaskingJobDelete(r string) error {
	var err error
	return err
}

//MaskingServiceConfigUpdate - Update the specified MaskingServiceConfig object.
//
//r - MaskingServiceConfig Object Reference
//payload *MaskingServiceConfigStruct object
func (c *Client) MaskingServiceConfigUpdate(r string, payload *MaskingServiceConfigStruct) error {
	var err error
	return err
}

//MaskingServiceConfigList - Returns a list of all Masking Jobs on the system.
//
//return []*MaskingServiceConfigStruct
func (c *Client) MaskingServiceConfigList() ([]*MaskingServiceConfigStruct, error) {
	var err error
	var result []*MaskingServiceConfigStruct
	obj, err := c.executeListAndReturnResults("/maskingjob/serviceconfig")
	for _, i := range obj {
		result = append(result, i.(*MaskingServiceConfigStruct))
	}
	return result, err
}

//MaskingServiceConfigRead - Retrieve the specified MaskingServiceConfig object.
//
//r - MaskingServiceConfig Object Reference
//return *MaskingServiceConfigStruct
func (c *Client) MaskingServiceConfigRead(r string) (*MaskingServiceConfigStruct, error) {
	var err error
	var result *MaskingServiceConfigStruct
	obj, err := c.FindObjectByReference("/maskingjob/serviceconfig", r)
	result = obj.(*MaskingServiceConfigStruct)
	return result, err
}

//NamespaceDelete - Delete the specified Namespace object.
//
//r - Namespace Object Reference
func (c *Client) NamespaceDelete(r string) error {
	var err error
	return err
}

//NamespaceList - List Namespace objects on the system.
//
//return []*NamespaceStruct
func (c *Client) NamespaceList() ([]*NamespaceStruct, error) {
	var err error
	var result []*NamespaceStruct
	obj, err := c.executeListAndReturnResults("/namespace")
	for _, i := range obj {
		result = append(result, i.(*NamespaceStruct))
	}
	return result, err
}

//NamespaceRead - Retrieve the specified Namespace object.
//
//r - Namespace Object Reference
//return *NamespaceStruct
func (c *Client) NamespaceRead(r string) (*NamespaceStruct, error) {
	var err error
	var result *NamespaceStruct
	obj, err := c.FindObjectByReference("/namespace", r)
	result = obj.(*NamespaceStruct)
	return result, err
}

//NamespaceUpdate - Update the specified Namespace object.
//
//r - Namespace Object Reference
//payload *NamespaceStruct object
func (c *Client) NamespaceUpdate(r string, payload *NamespaceStruct) error {
	var err error
	return err
}

//NetworkDSPTestDelete - Delete the specified NetworkDSPTest object.
//
//r - NetworkDSPTest Object Reference
func (c *Client) NetworkDSPTestDelete(r string) error {
	var err error
	return err
}

//NetworkDSPTestRead - Retrieve the specified NetworkDSPTest object.
//
//r - NetworkDSPTest Object Reference
//return *NetworkDSPTestStruct
func (c *Client) NetworkDSPTestRead(r string) (*NetworkDSPTestStruct, error) {
	var err error
	var result *NetworkDSPTestStruct
	obj, err := c.FindObjectByReference("/network/test/dsp", r)
	result = obj.(*NetworkDSPTestStruct)
	return result, err
}

//NetworkDSPTestList - Returns the list of previously executed tests.
//
//return []*NetworkDSPTestStruct
func (c *Client) NetworkDSPTestList() ([]*NetworkDSPTestStruct, error) {
	var err error
	var result []*NetworkDSPTestStruct
	obj, err := c.executeListAndReturnResults("/network/test/dsp")
	for _, i := range obj {
		result = append(result, i.(*NetworkDSPTestStruct))
	}
	return result, err
}

//NetworkDSPTestCreate - Create a new NetworkDSPTest object.
//
//payload *NetworkDSPTestParametersStruct object
//return string
func (c *Client) NetworkDSPTestCreate(payload *NetworkDSPTestParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//NetworkInterfaceRead - Retrieve the specified NetworkInterface object.
//
//r - NetworkInterface Object Reference
//return *NetworkInterfaceStruct
func (c *Client) NetworkInterfaceRead(r string) (*NetworkInterfaceStruct, error) {
	var err error
	var result *NetworkInterfaceStruct
	obj, err := c.FindObjectByReference("/network/interface", r)
	result = obj.(*NetworkInterfaceStruct)
	return result, err
}

//NetworkInterfaceList - List NetworkInterface objects on the system.
//
//return []*NetworkInterfaceStruct
func (c *Client) NetworkInterfaceList() ([]*NetworkInterfaceStruct, error) {
	var err error
	var result []*NetworkInterfaceStruct
	obj, err := c.executeListAndReturnResults("/network/interface")
	for _, i := range obj {
		result = append(result, i.(*NetworkInterfaceStruct))
	}
	return result, err
}

//NetworkInterfaceUpdate - Update the specified NetworkInterface object.
//
//r - NetworkInterface Object Reference
//payload *NetworkInterfaceStruct object
func (c *Client) NetworkInterfaceUpdate(r string, payload *NetworkInterfaceStruct) error {
	var err error
	return err
}

//NetworkLatencyTestRead - Retrieve the specified NetworkLatencyTest object.
//
//r - NetworkLatencyTest Object Reference
//return *NetworkLatencyTestStruct
func (c *Client) NetworkLatencyTestRead(r string) (*NetworkLatencyTestStruct, error) {
	var err error
	var result *NetworkLatencyTestStruct
	obj, err := c.FindObjectByReference("/network/test/latency", r)
	result = obj.(*NetworkLatencyTestStruct)
	return result, err
}

//NetworkLatencyTestCreate - Create a new NetworkLatencyTest object.
//
//payload *NetworkLatencyTestParametersStruct object
//return string
func (c *Client) NetworkLatencyTestCreate(payload *NetworkLatencyTestParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//NetworkLatencyTestDelete - Delete the specified NetworkLatencyTest object.
//
//r - NetworkLatencyTest Object Reference
func (c *Client) NetworkLatencyTestDelete(r string) error {
	var err error
	return err
}

//NetworkLatencyTestList - Returns the list of previously executed tests.
//
//return []*NetworkLatencyTestStruct
func (c *Client) NetworkLatencyTestList() ([]*NetworkLatencyTestStruct, error) {
	var err error
	var result []*NetworkLatencyTestStruct
	obj, err := c.executeListAndReturnResults("/network/test/latency")
	for _, i := range obj {
		result = append(result, i.(*NetworkLatencyTestStruct))
	}
	return result, err
}

//NetworkRouteList - Lists entries in the routing table.
//
//return []*NetworkRouteStruct
func (c *Client) NetworkRouteList() ([]*NetworkRouteStruct, error) {
	var err error
	var result []*NetworkRouteStruct
	obj, err := c.executeListAndReturnResults("/network/route")
	for _, i := range obj {
		result = append(result, i.(*NetworkRouteStruct))
	}
	return result, err
}

//NetworkThroughputTestRead - Retrieve the specified NetworkThroughputTest object.
//
//r - NetworkThroughputTest Object Reference
//return *NetworkThroughputTestStruct
func (c *Client) NetworkThroughputTestRead(r string) (*NetworkThroughputTestStruct, error) {
	var err error
	var result *NetworkThroughputTestStruct
	obj, err := c.FindObjectByReference("/network/test/throughput", r)
	result = obj.(*NetworkThroughputTestStruct)
	return result, err
}

//NetworkThroughputTestList - Returns the list of previously executed tests.
//
//return []*NetworkThroughputTestStruct
func (c *Client) NetworkThroughputTestList() ([]*NetworkThroughputTestStruct, error) {
	var err error
	var result []*NetworkThroughputTestStruct
	obj, err := c.executeListAndReturnResults("/network/test/throughput")
	for _, i := range obj {
		result = append(result, i.(*NetworkThroughputTestStruct))
	}
	return result, err
}

//NetworkThroughputTestDelete - Delete the specified NetworkThroughputTest object.
//
//r - NetworkThroughputTest Object Reference
func (c *Client) NetworkThroughputTestDelete(r string) error {
	var err error
	return err
}

//NetworkThroughputTestCreate - Create a new NetworkThroughputTest object.
//
//payload *NetworkThroughputTestParametersStruct object
//return string
func (c *Client) NetworkThroughputTestCreate(payload *NetworkThroughputTestParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//NotificationList - Returns a list of pending notifications for the current session.
//
//timeout ==>
//   type: string
//   description: Timeout, in milliseconds, to wait for one or more responses.
//max ==>
//   description: Maximum number of entries to return at once.
//   type: string
//channel ==>
//   type: string
//   description: Client-specific ID to specify an independent channel.
//return []Notification
func (c *Client) NotificationList(timeout string, max string, channel string) ([]Notification, error) {
	var err error
	var result []Notification
	obj, err := c.executeListAndReturnResults("/notification")
	for _, i := range obj {
		result = append(result, i.(Notification))
	}
	return result, err
}

//OperationTemplateRead - Retrieve the specified OperationTemplate object.
//
//r - OperationTemplate Object Reference
//return *OperationTemplateStruct
func (c *Client) OperationTemplateRead(r string) (*OperationTemplateStruct, error) {
	var err error
	var result *OperationTemplateStruct
	obj, err := c.FindObjectByReference("/source/operationTemplate", r)
	result = obj.(*OperationTemplateStruct)
	return result, err
}

//OperationTemplateUpdate - Update the specified OperationTemplate object.
//
//r - OperationTemplate Object Reference
//payload *OperationTemplateStruct object
func (c *Client) OperationTemplateUpdate(r string, payload *OperationTemplateStruct) error {
	var err error
	return err
}

//OperationTemplateList - Lists templates available in the Delphix Engine.
//
//return []*OperationTemplateStruct
func (c *Client) OperationTemplateList() ([]*OperationTemplateStruct, error) {
	var err error
	var result []*OperationTemplateStruct
	obj, err := c.executeListAndReturnResults("/source/operationTemplate")
	for _, i := range obj {
		result = append(result, i.(*OperationTemplateStruct))
	}
	return result, err
}

//OperationTemplateCreate - Create a new OperationTemplate object.
//
//payload *OperationTemplateStruct object
//return string
func (c *Client) OperationTemplateCreate(payload *OperationTemplateStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//OperationTemplateDelete - Delete the specified OperationTemplate object.
//
//r - OperationTemplate Object Reference
func (c *Client) OperationTemplateDelete(r string) error {
	var err error
	return err
}

//OracleClusterNodeRead - Retrieve the specified OracleClusterNode object.
//
//r - OracleClusterNode Object Reference
//return *OracleClusterNodeStruct
func (c *Client) OracleClusterNodeRead(r string) (*OracleClusterNodeStruct, error) {
	var err error
	var result *OracleClusterNodeStruct
	obj, err := c.FindObjectByReference("/environment/oracle/clusternode", r)
	result = obj.(*OracleClusterNodeStruct)
	return result, err
}

//OracleClusterNodeCreate - Create a new OracleClusterNode object.
//
//payload *OracleClusterNodeCreateParametersStruct object
//return string
func (c *Client) OracleClusterNodeCreate(payload *OracleClusterNodeCreateParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//OracleClusterNodeDelete - Delete the specified OracleClusterNode object.
//
//r - OracleClusterNode Object Reference
func (c *Client) OracleClusterNodeDelete(r string) error {
	var err error
	return err
}

//OracleClusterNodeList - Returns a list of host nodes filtered by cluster.
//
//cluster ==>
//   type: string
//   format: objectReference
//   description: The cluster to filter by.
//   referenceTo: /delphix-oracle-cluster.json
//   mapsTo: cluster
//return []*OracleClusterNodeStruct
func (c *Client) OracleClusterNodeList(cluster string) ([]*OracleClusterNodeStruct, error) {
	var err error
	var result []*OracleClusterNodeStruct
	obj, err := c.executeListAndReturnResults("/environment/oracle/clusternode")
	for _, i := range obj {
		result = append(result, i.(*OracleClusterNodeStruct))
	}
	return result, err
}

//OracleClusterNodeUpdate - Update the specified OracleClusterNode object.
//
//r - OracleClusterNode Object Reference
//payload *OracleClusterNodeStruct object
func (c *Client) OracleClusterNodeUpdate(r string, payload *OracleClusterNodeStruct) error {
	var err error
	return err
}

//OracleListenerUpdate - Update the specified OracleListener object.
//
//r - OracleListener Object Reference
//payload OracleListener object
func (c *Client) OracleListenerUpdate(r string, payload OracleListener) error {
	var err error
	return err
}

//OracleListenerList - Returns a list of listeners within the environment or the system.
//
//typeOracleListener ==>
//   description: Restrict listeners to type.
//   type: string
//   enum: [OracleNodeListener OracleScanListener]
//   required: false
//   mapsTo: type
//environment ==>
//   description: Restrict listeners belonging to the specified environment.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-source-environment.json
//   required: false
//   mapsTo: environment
//return []OracleListener
func (c *Client) OracleListenerList(typeOracleListener string, environment string) ([]OracleListener, error) {
	var err error
	var result []OracleListener
	obj, err := c.executeListAndReturnResults("/environment/oracle/listener")
	for _, i := range obj {
		result = append(result, i.(OracleListener))
	}
	return result, err
}

//OracleListenerRead - Retrieve the specified OracleListener object.
//
//r - OracleListener Object Reference
//return OracleListener
func (c *Client) OracleListenerRead(r string) (OracleListener, error) {
	var err error
	var result OracleListener
	obj, err := c.FindObjectByReference("/environment/oracle/listener", r)
	result = obj.(OracleListener)
	return result, err
}

//OracleListenerCreate - Create a new OracleListener object.
//
//payload OracleListener object
//return string
func (c *Client) OracleListenerCreate(payload OracleListener) (string, error) {
	var err error
	var result string
	return result, err
}

//OracleListenerDelete - Delete the specified OracleListener object.
//
//r - OracleListener Object Reference
func (c *Client) OracleListenerDelete(r string) error {
	var err error
	return err
}

//OracleTimeflowLogList - Returns a list of fetched or missing Oracle logs for a database,
// TimeFlow or snapshot. The logs are returned in ascending order by
// TimeFlow, SCN.
//
//timeflow ==>
//   type: string
//   description: Return logs in the specified TimeFlow. This option is mutually exclusive with the "snapshot" and "database" options.
//   format: objectReference
//   referenceTo: /delphix-oracle-timeflow.json
//   mapsTo: timeflow
//snapshot ==>
//   referenceTo: /delphix-oracle-snapshot.json
//   type: string
//   description: Return logs for the specified snapshot up to the next snapshot. This option is mutually exclusive with the "TimeFlow" and "database" options.
//   format: objectReference
//fromDate ==>
//   description: Return logs created after this date.
//   format: date
//   type: string
//toDate ==>
//   description: Return logs created before than this date.
//   format: date
//   type: string
//toScn ==>
//   description: Return logs with SCNs less than or equal to this value.
//   type: string
//missing ==>
//   type: boolean
//   description: Only return the missing logs.
//database ==>
//   type: string
//   description: Return logs on all TimeFlows associated with the container. This option is mutually exclusive with the "TimeFlow" and "snapshot" options.
//   format: objectReference
//   referenceTo: /delphix-oracle-db-container.json
//fromScn ==>
//   type: string
//   description: Return logs with SCNs greater than or equal to this value.
//pageSize ==>
//   type: integer
//   default: 25
//   minimum: 0
//   description: Limit the number of logs returned.
//pageOffset ==>
//   description: Page offset within log list, in units of pageSize chunks.
//   type: integer
//return []OracleTimeflowLog
func (c *Client) OracleTimeflowLogList(timeflow string, snapshot string, fromDate string, toDate string, toScn string, missing *bool, database string, fromScn string, pageSize *int, pageOffset *int) ([]OracleTimeflowLog, error) {
	var err error
	var result []OracleTimeflowLog
	obj, err := c.executeListAndReturnResults("/timeflow/oracle/log")
	for _, i := range obj {
		result = append(result, i.(OracleTimeflowLog))
	}
	return result, err
}

//PasswordPolicyCreate - Create a new PasswordPolicy object.
//
//payload *PasswordPolicyStruct object
//return string
func (c *Client) PasswordPolicyCreate(payload *PasswordPolicyStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//PasswordPolicyUpdate - Update the specified PasswordPolicy object.
//
//r - PasswordPolicy Object Reference
//payload *PasswordPolicyStruct object
func (c *Client) PasswordPolicyUpdate(r string, payload *PasswordPolicyStruct) error {
	var err error
	return err
}

//PasswordPolicyRead - Retrieve the specified PasswordPolicy object.
//
//r - PasswordPolicy Object Reference
//return *PasswordPolicyStruct
func (c *Client) PasswordPolicyRead(r string) (*PasswordPolicyStruct, error) {
	var err error
	var result *PasswordPolicyStruct
	obj, err := c.FindObjectByReference("/passwordPolicy", r)
	result = obj.(*PasswordPolicyStruct)
	return result, err
}

//PasswordPolicyDelete - Delete the specified PasswordPolicy object.
//
//r - PasswordPolicy Object Reference
func (c *Client) PasswordPolicyDelete(r string) error {
	var err error
	return err
}

//PasswordPolicyList - Lists password policies in the system.
//
//return []*PasswordPolicyStruct
func (c *Client) PasswordPolicyList() ([]*PasswordPolicyStruct, error) {
	var err error
	var result []*PasswordPolicyStruct
	obj, err := c.executeListAndReturnResults("/passwordPolicy")
	for _, i := range obj {
		result = append(result, i.(*PasswordPolicyStruct))
	}
	return result, err
}

//PermissionList - Lists permissions available in the system.
//
//return []*PermissionStruct
func (c *Client) PermissionList() ([]*PermissionStruct, error) {
	var err error
	var result []*PermissionStruct
	obj, err := c.executeListAndReturnResults("/permission")
	for _, i := range obj {
		result = append(result, i.(*PermissionStruct))
	}
	return result, err
}

//PermissionRead - Retrieve the specified Permission object.
//
//r - Permission Object Reference
//return *PermissionStruct
func (c *Client) PermissionRead(r string) (*PermissionStruct, error) {
	var err error
	var result *PermissionStruct
	obj, err := c.FindObjectByReference("/permission", r)
	result = obj.(*PermissionStruct)
	return result, err
}

//PhoneHomeServiceRead - Retrieve the specified PhoneHomeService object.
//
//r - PhoneHomeService Object Reference
//return *PhoneHomeServiceStruct
func (c *Client) PhoneHomeServiceRead(r string) (*PhoneHomeServiceStruct, error) {
	var err error
	var result *PhoneHomeServiceStruct
	obj, err := c.FindObjectByReference("/service/phonehome", r)
	result = obj.(*PhoneHomeServiceStruct)
	return result, err
}

//PhoneHomeServiceUpdate - Update the specified PhoneHomeService object.
//
//r - PhoneHomeService Object Reference
//payload *PhoneHomeServiceStruct object
func (c *Client) PhoneHomeServiceUpdate(r string, payload *PhoneHomeServiceStruct) error {
	var err error
	return err
}

//PolicyList - Returns a list of policies in the domain.
//
//typePolicy ==>
//   type: string
//   description: Limit policies to those of the given type.
//target ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   description: Limit policies to those affecting a particular object on the system.
//effective ==>
//   description: Whether to include effective policies for the target.
//   type: string
//return []Policy
func (c *Client) PolicyList(typePolicy string, target string, effective string) ([]Policy, error) {
	var err error
	var result []Policy
	obj, err := c.executeListAndReturnResults("/policy")
	for _, i := range obj {
		result = append(result, i.(Policy))
	}
	return result, err
}

//PolicyCreate - Create a new Policy object.
//
//payload Policy object
//return string
func (c *Client) PolicyCreate(payload Policy) (string, error) {
	var err error
	var result string
	return result, err
}

//PolicyDelete - Delete the specified Policy object.
//
//r - Policy Object Reference
func (c *Client) PolicyDelete(r string) error {
	var err error
	return err
}

//PolicyRead - Retrieve the specified Policy object.
//
//r - Policy Object Reference
//return Policy
func (c *Client) PolicyRead(r string) (Policy, error) {
	var err error
	var result Policy
	obj, err := c.FindObjectByReference("/policy", r)
	result = obj.(Policy)
	return result, err
}

//PolicyUpdate - Update the specified Policy object.
//
//r - Policy Object Reference
//payload Policy object
func (c *Client) PolicyUpdate(r string, payload Policy) error {
	var err error
	return err
}

//ProxyServiceUpdate - Update the specified ProxyService object.
//
//r - ProxyService Object Reference
//payload *ProxyServiceStruct object
func (c *Client) ProxyServiceUpdate(r string, payload *ProxyServiceStruct) error {
	var err error
	return err
}

//ProxyServiceRead - Retrieve the specified ProxyService object.
//
//r - ProxyService Object Reference
//return *ProxyServiceStruct
func (c *Client) ProxyServiceRead(r string) (*ProxyServiceStruct, error) {
	var err error
	var result *ProxyServiceStruct
	obj, err := c.FindObjectByReference("/service/proxy", r)
	result = obj.(*ProxyServiceStruct)
	return result, err
}

//PublicSystemInfoRead - Retrieve the specified PublicSystemInfo object.
//
//r - PublicSystemInfo Object Reference
//return *PublicSystemInfoStruct
func (c *Client) PublicSystemInfoRead(r string) (*PublicSystemInfoStruct, error) {
	var err error
	var result *PublicSystemInfoStruct
	obj, err := c.FindObjectByReference("/about", r)
	result = obj.(*PublicSystemInfoStruct)
	return result, err
}

//RegistrationInfoRead - Retrieve the specified RegistrationInfo object.
//
//r - RegistrationInfo Object Reference
//return *RegistrationInfoStruct
func (c *Client) RegistrationInfoRead(r string) (*RegistrationInfoStruct, error) {
	var err error
	var result *RegistrationInfoStruct
	obj, err := c.FindObjectByReference("/registration", r)
	result = obj.(*RegistrationInfoStruct)
	return result, err
}

//RegistrationStatusRead - Retrieve the specified RegistrationStatus object.
//
//r - RegistrationStatus Object Reference
//return *RegistrationStatusStruct
func (c *Client) RegistrationStatusRead(r string) (*RegistrationStatusStruct, error) {
	var err error
	var result *RegistrationStatusStruct
	obj, err := c.FindObjectByReference("/registration/status", r)
	result = obj.(*RegistrationStatusStruct)
	return result, err
}

//RegistrationStatusUpdate - Update the specified RegistrationStatus object.
//
//r - RegistrationStatus Object Reference
//payload *RegistrationStatusStruct object
func (c *Client) RegistrationStatusUpdate(r string, payload *RegistrationStatusStruct) error {
	var err error
	return err
}

//ReplicationSourceStateRead - Retrieve the specified ReplicationSourceState object.
//
//r - ReplicationSourceState Object Reference
//return *ReplicationSourceStateStruct
func (c *Client) ReplicationSourceStateRead(r string) (*ReplicationSourceStateStruct, error) {
	var err error
	var result *ReplicationSourceStateStruct
	obj, err := c.FindObjectByReference("/replication/sourcestate", r)
	result = obj.(*ReplicationSourceStateStruct)
	return result, err
}

//ReplicationSourceStateList - List ReplicationSourceState objects on the system.
//
//return []*ReplicationSourceStateStruct
func (c *Client) ReplicationSourceStateList() ([]*ReplicationSourceStateStruct, error) {
	var err error
	var result []*ReplicationSourceStateStruct
	obj, err := c.executeListAndReturnResults("/replication/sourcestate")
	for _, i := range obj {
		result = append(result, i.(*ReplicationSourceStateStruct))
	}
	return result, err
}

//ReplicationSpecCreate - Create a new ReplicationSpec object.
//
//payload *ReplicationSpecStruct object
//return string
func (c *Client) ReplicationSpecCreate(payload *ReplicationSpecStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//ReplicationSpecRead - Retrieve the specified ReplicationSpec object.
//
//r - ReplicationSpec Object Reference
//return *ReplicationSpecStruct
func (c *Client) ReplicationSpecRead(r string) (*ReplicationSpecStruct, error) {
	var err error
	var result *ReplicationSpecStruct
	obj, err := c.FindObjectByReference("/replication/spec", r)
	result = obj.(*ReplicationSpecStruct)
	return result, err
}

//ReplicationSpecUpdate - Update the specified ReplicationSpec object.
//
//r - ReplicationSpec Object Reference
//payload *ReplicationSpecStruct object
func (c *Client) ReplicationSpecUpdate(r string, payload *ReplicationSpecStruct) error {
	var err error
	return err
}

//ReplicationSpecDelete - Delete the specified ReplicationSpec object.
//
//r - ReplicationSpec Object Reference
func (c *Client) ReplicationSpecDelete(r string) error {
	var err error
	return err
}

//ReplicationSpecList - List ReplicationSpec objects on the system.
//
//return []*ReplicationSpecStruct
func (c *Client) ReplicationSpecList() ([]*ReplicationSpecStruct, error) {
	var err error
	var result []*ReplicationSpecStruct
	obj, err := c.executeListAndReturnResults("/replication/spec")
	for _, i := range obj {
		result = append(result, i.(*ReplicationSpecStruct))
	}
	return result, err
}

//ReplicationTargetStateRead - Retrieve the specified ReplicationTargetState object.
//
//r - ReplicationTargetState Object Reference
//return *ReplicationTargetStateStruct
func (c *Client) ReplicationTargetStateRead(r string) (*ReplicationTargetStateStruct, error) {
	var err error
	var result *ReplicationTargetStateStruct
	obj, err := c.FindObjectByReference("/replication/targetstate", r)
	result = obj.(*ReplicationTargetStateStruct)
	return result, err
}

//ReplicationTargetStateList - List ReplicationTargetState objects on the system.
//
//return []*ReplicationTargetStateStruct
func (c *Client) ReplicationTargetStateList() ([]*ReplicationTargetStateStruct, error) {
	var err error
	var result []*ReplicationTargetStateStruct
	obj, err := c.executeListAndReturnResults("/replication/targetstate")
	for _, i := range obj {
		result = append(result, i.(*ReplicationTargetStateStruct))
	}
	return result, err
}

//RoleRead - Retrieve the specified Role object.
//
//r - Role Object Reference
//return *RoleStruct
func (c *Client) RoleRead(r string) (*RoleStruct, error) {
	var err error
	var result *RoleStruct
	obj, err := c.FindObjectByReference("/role", r)
	result = obj.(*RoleStruct)
	return result, err
}

//RoleCreate - Create a new Role object.
//
//payload *RoleStruct object
//return string
func (c *Client) RoleCreate(payload *RoleStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//RoleUpdate - Update the specified Role object.
//
//r - Role Object Reference
//payload *RoleStruct object
func (c *Client) RoleUpdate(r string, payload *RoleStruct) error {
	var err error
	return err
}

//RoleDelete - Delete the specified Role object.
//
//r - Role Object Reference
func (c *Client) RoleDelete(r string) error {
	var err error
	return err
}

//RoleList - Lists roles available in the system.
//
//return []*RoleStruct
func (c *Client) RoleList() ([]*RoleStruct, error) {
	var err error
	var result []*RoleStruct
	obj, err := c.executeListAndReturnResults("/role")
	for _, i := range obj {
		result = append(result, i.(*RoleStruct))
	}
	return result, err
}

//SMTPConfigRead - Retrieve the specified SMTPConfig object.
//
//r - SMTPConfig Object Reference
//return *SMTPConfigStruct
func (c *Client) SMTPConfigRead(r string) (*SMTPConfigStruct, error) {
	var err error
	var result *SMTPConfigStruct
	obj, err := c.FindObjectByReference("/service/smtp", r)
	result = obj.(*SMTPConfigStruct)
	return result, err
}

//SMTPConfigUpdate - Update the specified SMTPConfig object.
//
//r - SMTPConfig Object Reference
//payload *SMTPConfigStruct object
func (c *Client) SMTPConfigUpdate(r string, payload *SMTPConfigStruct) error {
	var err error
	return err
}

//SNMPConfigUpdate - Update the specified SNMPConfig object.
//
//r - SNMPConfig Object Reference
//payload *SNMPConfigStruct object
func (c *Client) SNMPConfigUpdate(r string, payload *SNMPConfigStruct) error {
	var err error
	return err
}

//SNMPConfigRead - Retrieve the specified SNMPConfig object.
//
//r - SNMPConfig Object Reference
//return *SNMPConfigStruct
func (c *Client) SNMPConfigRead(r string) (*SNMPConfigStruct, error) {
	var err error
	var result *SNMPConfigStruct
	obj, err := c.FindObjectByReference("/service/snmp", r)
	result = obj.(*SNMPConfigStruct)
	return result, err
}

//SNMPManagerRead - Retrieve the specified SNMPManager object.
//
//r - SNMPManager Object Reference
//return *SNMPManagerStruct
func (c *Client) SNMPManagerRead(r string) (*SNMPManagerStruct, error) {
	var err error
	var result *SNMPManagerStruct
	obj, err := c.FindObjectByReference("/service/snmp/manager", r)
	result = obj.(*SNMPManagerStruct)
	return result, err
}

//SNMPManagerUpdate - Update the specified SNMPManager object.
//
//r - SNMPManager Object Reference
//payload *SNMPManagerStruct object
func (c *Client) SNMPManagerUpdate(r string, payload *SNMPManagerStruct) error {
	var err error
	return err
}

//SNMPManagerCreate - Create a new SNMPManager object.
//
//payload *SNMPManagerStruct object
//return string
func (c *Client) SNMPManagerCreate(payload *SNMPManagerStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//SNMPManagerDelete - Delete the specified SNMPManager object.
//
//r - SNMPManager Object Reference
func (c *Client) SNMPManagerDelete(r string) error {
	var err error
	return err
}

//SNMPManagerList - Lists SNMP managers in the system.
//
//return []*SNMPManagerStruct
func (c *Client) SNMPManagerList() ([]*SNMPManagerStruct, error) {
	var err error
	var result []*SNMPManagerStruct
	obj, err := c.executeListAndReturnResults("/service/snmp/manager")
	for _, i := range obj {
		result = append(result, i.(*SNMPManagerStruct))
	}
	return result, err
}

//SamlInfoRead - Retrieve the specified SamlInfo object.
//
//r - SamlInfo Object Reference
//return *SamlInfoStruct
func (c *Client) SamlInfoRead(r string) (*SamlInfoStruct, error) {
	var err error
	var result *SamlInfoStruct
	obj, err := c.FindObjectByReference("/service/saml", r)
	result = obj.(*SamlInfoStruct)
	return result, err
}

//SamlServiceProviderUpdate - Update the specified SamlServiceProvider object.
//
//r - SamlServiceProvider Object Reference
//payload *SamlServiceProviderStruct object
func (c *Client) SamlServiceProviderUpdate(r string, payload *SamlServiceProviderStruct) error {
	var err error
	return err
}

//SamlServiceProviderRead - Retrieve the specified SamlServiceProvider object.
//
//r - SamlServiceProvider Object Reference
//return *SamlServiceProviderStruct
func (c *Client) SamlServiceProviderRead(r string) (*SamlServiceProviderStruct, error) {
	var err error
	var result *SamlServiceProviderStruct
	obj, err := c.FindObjectByReference("/service/saml/serviceprovider", r)
	result = obj.(*SamlServiceProviderStruct)
	return result, err
}

//SamlServiceProviderDelete - Delete the specified SamlServiceProvider object.
//
//r - SamlServiceProvider Object Reference
func (c *Client) SamlServiceProviderDelete(r string) error {
	var err error
	return err
}

//SamlServiceProviderCreate - Create a new SamlServiceProvider object.
//
//payload *SamlServiceProviderStruct object
//return string
func (c *Client) SamlServiceProviderCreate(payload *SamlServiceProviderStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//SamlServiceProviderList - List SamlServiceProvider objects on the system.
//
//return []*SamlServiceProviderStruct
func (c *Client) SamlServiceProviderList() ([]*SamlServiceProviderStruct, error) {
	var err error
	var result []*SamlServiceProviderStruct
	obj, err := c.executeListAndReturnResults("/service/saml/serviceprovider")
	for _, i := range obj {
		result = append(result, i.(*SamlServiceProviderStruct))
	}
	return result, err
}

//SchemaRead - Retrieve the specified Schema object.
//
//r - Schema Object Reference
//return *SchemaStruct
func (c *Client) SchemaRead(r string) (*SchemaStruct, error) {
	var err error
	var result *SchemaStruct
	obj, err := c.FindObjectByReference("/service/schema", r)
	result = obj.(*SchemaStruct)
	return result, err
}

//ScrubStatusRead - Retrieve the specified ScrubStatus object.
//
//r - ScrubStatus Object Reference
//return *ScrubStatusStruct
func (c *Client) ScrubStatusRead(r string) (*ScrubStatusStruct, error) {
	var err error
	var result *ScrubStatusStruct
	obj, err := c.FindObjectByReference("/storage/scrub", r)
	result = obj.(*ScrubStatusStruct)
	return result, err
}

//SecurityConfigRead - Retrieve the specified SecurityConfig object.
//
//r - SecurityConfig Object Reference
//return *SecurityConfigStruct
func (c *Client) SecurityConfigRead(r string) (*SecurityConfigStruct, error) {
	var err error
	var result *SecurityConfigStruct
	obj, err := c.FindObjectByReference("/service/security", r)
	result = obj.(*SecurityConfigStruct)
	return result, err
}

//SecurityConfigUpdate - Update the specified SecurityConfig object.
//
//r - SecurityConfig Object Reference
//payload *SecurityConfigStruct object
func (c *Client) SecurityConfigUpdate(r string, payload *SecurityConfigStruct) error {
	var err error
	return err
}

//SerializationPointRead - Retrieve the specified SerializationPoint object.
//
//r - SerializationPoint Object Reference
//return *SerializationPointStruct
func (c *Client) SerializationPointRead(r string) (*SerializationPointStruct, error) {
	var err error
	var result *SerializationPointStruct
	obj, err := c.FindObjectByReference("/replication/serializationpoint", r)
	result = obj.(*SerializationPointStruct)
	return result, err
}

//SerializationPointList - List SerializationPoint objects on the system.
//
//return []*SerializationPointStruct
func (c *Client) SerializationPointList() ([]*SerializationPointStruct, error) {
	var err error
	var result []*SerializationPointStruct
	obj, err := c.executeListAndReturnResults("/replication/serializationpoint")
	for _, i := range obj {
		result = append(result, i.(*SerializationPointStruct))
	}
	return result, err
}

//SnapshotCapacityDataList - Lists capacity metrics for all snapshots in the syste sorted by
// snapshot space usage decreasing.
//
//container ==>
//   description: The container to list snapshot data for.
//   mapsTo: container
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//pageSize ==>
//   description: Limit the number of entries returned.
//   type: integer
//   minimum: 0
//pageOffset ==>
//   description: Offset within list, in units of pageSize chunks.
//   type: integer
//   minimum: 0
//return []*SnapshotCapacityDataStruct
func (c *Client) SnapshotCapacityDataList(container string, pageSize *int, pageOffset *int) ([]*SnapshotCapacityDataStruct, error) {
	var err error
	var result []*SnapshotCapacityDataStruct
	obj, err := c.executeListAndReturnResults("/capacity/snapshot")
	for _, i := range obj {
		result = append(result, i.(*SnapshotCapacityDataStruct))
	}
	return result, err
}

//SourceList - Lists sources on the system.
//
//config ==>
//   mapsTo: config
//   description: List visible sources associated with the given sourceconfig reference. Visible sources are of type LINKED or VIRTUAL.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-source-config.json
//allSources ==>
//   description: List all sources associated with the given source container reference.
//   type: boolean
//repository ==>
//   mapsTo: config.repository.reference
//   description: List sources associated with the given source repository reference.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-source-repository.json
//environment ==>
//   format: objectReference
//   referenceTo: /delphix-source-environment.json
//   mapsTo: config.repository.environment.reference
//   description: List sources associated with the given source environment reference.
//   type: string
//includeHosts ==>
//   description: Whether to include the list of hosts for each source in the response.
//   type: boolean
//database ==>
//   description: List visible sources associated with the given container reference. Visible sources are of type LINKED or VIRTUAL.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//return []Source
func (c *Client) SourceList(config string, allSources *bool, repository string, environment string, includeHosts *bool, database string) ([]Source, error) {
	var err error
	var result []Source
	obj, err := c.executeListAndReturnResults("/source")
	for _, i := range obj {
		result = append(result, i.(Source))
	}
	return result, err
}

//SourceUpdate - Update the specified Source object.
//
//r - Source Object Reference
//payload Source object
func (c *Client) SourceUpdate(r string, payload Source) error {
	var err error
	return err
}

//SourceRead - Retrieve the specified Source object.
//
//r - Source Object Reference
//return Source
func (c *Client) SourceRead(r string) (Source, error) {
	var err error
	var result Source
	obj, err := c.FindObjectByReference("/source", r)
	result = obj.(Source)
	return result, err
}

//SourceConfigList - Returns a list of source configs within the repository or the
// environment.
//
//repository ==>
//   referenceTo: /delphix-source-repository.json
//   mapsTo: repository
//   type: string
//   description: Restrict source configs to those belonging to the specified repository. This option is mutually exclusive with all other options.
//   format: objectReference
//environment ==>
//   format: objectReference
//   referenceTo: /delphix-source-environment.json
//   mapsTo: repository.environment.reference
//   type: string
//   description: Restrict source configs to those belonging to the specified environment. This option is mutually exclusive with all other options.
//cdbConfig ==>
//   format: objectReference
//   referenceTo: /delphix-oracle-db-config.json
//   type: string
//   description: Restrict PDB configs to those belonging to the specified CDB source config.
//pdbConfigOnly ==>
//   description: Restrict source configs to be Oracle PDB configs only.
//   type: boolean
//return []SourceConfig
func (c *Client) SourceConfigList(repository string, environment string, cdbConfig string, pdbConfigOnly *bool) ([]SourceConfig, error) {
	var err error
	var result []SourceConfig
	obj, err := c.executeListAndReturnResults("/sourceconfig")
	for _, i := range obj {
		result = append(result, i.(SourceConfig))
	}
	return result, err
}

//SourceConfigDelete - Delete the specified SourceConfig object.
//
//r - SourceConfig Object Reference
func (c *Client) SourceConfigDelete(r string) error {
	var err error
	return err
}

//SourceConfigCreate - Create a new SourceConfig object.
//
//payload SourceConfig object
//return string
func (c *Client) SourceConfigCreate(payload SourceConfig) (string, error) {
	var err error
	var result string
	return result, err
}

//SourceConfigRead - Retrieve the specified SourceConfig object.
//
//r - SourceConfig Object Reference
//return SourceConfig
func (c *Client) SourceConfigRead(r string) (SourceConfig, error) {
	var err error
	var result SourceConfig
	obj, err := c.FindObjectByReference("/sourceconfig", r)
	result = obj.(SourceConfig)
	return result, err
}

//SourceConfigUpdate - Update the specified SourceConfig object.
//
//r - SourceConfig Object Reference
//payload SourceConfig object
func (c *Client) SourceConfigUpdate(r string, payload SourceConfig) error {
	var err error
	return err
}

//SourceEnvironmentRead - Retrieve the specified SourceEnvironment object.
//
//r - SourceEnvironment Object Reference
//return SourceEnvironment
func (c *Client) SourceEnvironmentRead(r string) (SourceEnvironment, error) {
	var err error
	var result SourceEnvironment
	obj, err := c.FindObjectByReference("/environment", r)
	result = obj.(SourceEnvironment)
	return result, err
}

//SourceEnvironmentUpdate - Update the specified SourceEnvironment object.
//
//r - SourceEnvironment Object Reference
//payload SourceEnvironment object
func (c *Client) SourceEnvironmentUpdate(r string, payload SourceEnvironment) error {
	var err error
	return err
}

//SourceEnvironmentCreate - Create a new SourceEnvironment object.
//
//payload SourceEnvironmentCreateParameters object
//return string
func (c *Client) SourceEnvironmentCreate(payload SourceEnvironmentCreateParameters) (string, error) {
	var err error
	var result string
	return result, err
}

//SourceEnvironmentList - Returns the list of all source environments.
//
//typeSourceEnvironment ==>
//   type: string
//   description: Filter the results based on the type of environment.
//   enum: [WindowsHostEnvironment UnixHostEnvironment ASEUnixHostEnvironment OracleCluster]
//return []SourceEnvironment
func (c *Client) SourceEnvironmentList(typeSourceEnvironment string) ([]SourceEnvironment, error) {
	var err error
	var result []SourceEnvironment
	obj, err := c.executeListAndReturnResults("/environment")
	for _, i := range obj {
		objType := i.(map[string]interface{})["type"].(string)
		fmt.Printf("Type: %v\n", objType)
		result = append(result, i.(SourceEnvironment))
	}
	return result, err
}

//SourceEnvironmentDelete - Delete the specified SourceEnvironment object.
//
//r - SourceEnvironment Object Reference
func (c *Client) SourceEnvironmentDelete(r string) error {
	var err error
	return err
}

//SourceRepositoryDelete - Delete the specified SourceRepository object.
//
//r - SourceRepository Object Reference
func (c *Client) SourceRepositoryDelete(r string) error {
	var err error
	return err
}

//SourceRepositoryUpdate - Update the specified SourceRepository object.
//
//r - SourceRepository Object Reference
//payload SourceRepository object
func (c *Client) SourceRepositoryUpdate(r string, payload SourceRepository) error {
	var err error
	return err
}

//SourceRepositoryList - Returns a list of source repositories within the environment or
// the system.
//
//environment ==>
//   type: string
//   format: objectReference
//   description: Restrict source repositories belong to the specified environment.
//   referenceTo: /delphix-source-environment.json
//   mapsTo: environment
//return []SourceRepository
func (c *Client) SourceRepositoryList(environment string) ([]SourceRepository, error) {
	var err error
	var result []SourceRepository
	obj, err := c.executeListAndReturnResults("/repository")
	for _, i := range obj {
		result = append(result, i.(SourceRepository))
	}
	return result, err
}

//SourceRepositoryCreate - Create a new SourceRepository object.
//
//payload SourceRepository object
//return string
func (c *Client) SourceRepositoryCreate(payload SourceRepository) (string, error) {
	var err error
	var result string
	return result, err
}

//SourceRepositoryRead - Retrieve the specified SourceRepository object.
//
//r - SourceRepository Object Reference
//return SourceRepository
func (c *Client) SourceRepositoryRead(r string) (SourceRepository, error) {
	var err error
	var result SourceRepository
	obj, err := c.FindObjectByReference("/repository", r)
	result = obj.(SourceRepository)
	return result, err
}

//SourceRepositoryTemplateList - Returns the list of repository templates.
//
//return []*SourceRepositoryTemplateStruct
func (c *Client) SourceRepositoryTemplateList() ([]*SourceRepositoryTemplateStruct, error) {
	var err error
	var result []*SourceRepositoryTemplateStruct
	obj, err := c.executeListAndReturnResults("/repository/template")
	for _, i := range obj {
		result = append(result, i.(*SourceRepositoryTemplateStruct))
	}
	return result, err
}

//SourceRepositoryTemplateCreate - Create a new SourceRepositoryTemplate object.
//
//payload *SourceRepositoryTemplateStruct object
//return string
func (c *Client) SourceRepositoryTemplateCreate(payload *SourceRepositoryTemplateStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//SourceRepositoryTemplateRead - Retrieve the specified SourceRepositoryTemplate object.
//
//r - SourceRepositoryTemplate Object Reference
//return *SourceRepositoryTemplateStruct
func (c *Client) SourceRepositoryTemplateRead(r string) (*SourceRepositoryTemplateStruct, error) {
	var err error
	var result *SourceRepositoryTemplateStruct
	obj, err := c.FindObjectByReference("/repository/template", r)
	result = obj.(*SourceRepositoryTemplateStruct)
	return result, err
}

//SourceRepositoryTemplateDelete - Delete the specified SourceRepositoryTemplate object.
//
//r - SourceRepositoryTemplate Object Reference
func (c *Client) SourceRepositoryTemplateDelete(r string) error {
	var err error
	return err
}

//SourceRepositoryTemplateUpdate - Update the specified SourceRepositoryTemplate object.
//
//r - SourceRepositoryTemplate Object Reference
//payload *SourceRepositoryTemplateStruct object
func (c *Client) SourceRepositoryTemplateUpdate(r string, payload *SourceRepositoryTemplateStruct) error {
	var err error
	return err
}

//SplunkHecConfigRead - Retrieve the specified SplunkHecConfig object.
//
//r - SplunkHecConfig Object Reference
//return *SplunkHecConfigStruct
func (c *Client) SplunkHecConfigRead(r string) (*SplunkHecConfigStruct, error) {
	var err error
	var result *SplunkHecConfigStruct
	obj, err := c.FindObjectByReference("/service/insight/splunkHec", r)
	result = obj.(*SplunkHecConfigStruct)
	return result, err
}

//SplunkHecConfigDelete - Delete the specified SplunkHecConfig object.
//
//r - SplunkHecConfig Object Reference
func (c *Client) SplunkHecConfigDelete(r string) error {
	var err error
	return err
}

//SplunkHecConfigCreate - Create a new SplunkHecConfig object.
//
//payload *SplunkHecConfigStruct object
//return string
func (c *Client) SplunkHecConfigCreate(payload *SplunkHecConfigStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//SplunkHecConfigList - List SplunkHecConfig objects on the system.
//
//return []*SplunkHecConfigStruct
func (c *Client) SplunkHecConfigList() ([]*SplunkHecConfigStruct, error) {
	var err error
	var result []*SplunkHecConfigStruct
	obj, err := c.executeListAndReturnResults("/service/insight/splunkHec")
	for _, i := range obj {
		result = append(result, i.(*SplunkHecConfigStruct))
	}
	return result, err
}

//SplunkHecConfigUpdate - Update the specified SplunkHecConfig object.
//
//r - SplunkHecConfig Object Reference
//payload *SplunkHecConfigStruct object
func (c *Client) SplunkHecConfigUpdate(r string, payload *SplunkHecConfigStruct) error {
	var err error
	return err
}

//StatisticRead - Retrieve the specified Statistic object.
//
//r - Statistic Object Reference
//return *StatisticStruct
func (c *Client) StatisticRead(r string) (*StatisticStruct, error) {
	var err error
	var result *StatisticStruct
	obj, err := c.FindObjectByReference("/analytics/statistic", r)
	result = obj.(*StatisticStruct)
	return result, err
}

//StatisticList - Returns a list of statistics in the system.
//
//return []*StatisticStruct
func (c *Client) StatisticList() ([]*StatisticStruct, error) {
	var err error
	var result []*StatisticStruct
	obj, err := c.executeListAndReturnResults("/analytics/statistic")
	for _, i := range obj {
		result = append(result, i.(*StatisticStruct))
	}
	return result, err
}

//StatisticSliceRead - Retrieve the specified StatisticSlice object.
//
//r - StatisticSlice Object Reference
//return *StatisticSliceStruct
func (c *Client) StatisticSliceRead(r string) (*StatisticSliceStruct, error) {
	var err error
	var result *StatisticSliceStruct
	obj, err := c.FindObjectByReference("/analytics", r)
	result = obj.(*StatisticSliceStruct)
	return result, err
}

//StatisticSliceCreate - Create a new StatisticSlice object.
//
//payload *StatisticSliceStruct object
//return string
func (c *Client) StatisticSliceCreate(payload *StatisticSliceStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//StatisticSliceDelete - Delete the specified StatisticSlice object.
//
//r - StatisticSlice Object Reference
func (c *Client) StatisticSliceDelete(r string) error {
	var err error
	return err
}

//StatisticSliceList - Returns a list of statistics in the system.
//
//return []*StatisticSliceStruct
func (c *Client) StatisticSliceList() ([]*StatisticSliceStruct, error) {
	var err error
	var result []*StatisticSliceStruct
	obj, err := c.executeListAndReturnResults("/analytics")
	for _, i := range obj {
		result = append(result, i.(*StatisticSliceStruct))
	}
	return result, err
}

//StorageDeviceRead - Retrieve the specified StorageDevice object.
//
//r - StorageDevice Object Reference
//return *StorageDeviceStruct
func (c *Client) StorageDeviceRead(r string) (*StorageDeviceStruct, error) {
	var err error
	var result *StorageDeviceStruct
	obj, err := c.FindObjectByReference("/storage/device", r)
	result = obj.(*StorageDeviceStruct)
	return result, err
}

//StorageDeviceList - List StorageDevice objects on the system.
//
//return []*StorageDeviceStruct
func (c *Client) StorageDeviceList() ([]*StorageDeviceStruct, error) {
	var err error
	var result []*StorageDeviceStruct
	obj, err := c.executeListAndReturnResults("/storage/device")
	for _, i := range obj {
		result = append(result, i.(*StorageDeviceStruct))
	}
	return result, err
}

//StorageDeviceRemovalStatusRead - Retrieve the specified StorageDeviceRemovalStatus object.
//
//r - StorageDeviceRemovalStatus Object Reference
//return *StorageDeviceRemovalStatusStruct
func (c *Client) StorageDeviceRemovalStatusRead(r string) (*StorageDeviceRemovalStatusStruct, error) {
	var err error
	var result *StorageDeviceRemovalStatusStruct
	obj, err := c.FindObjectByReference("/storage/remove", r)
	result = obj.(*StorageDeviceRemovalStatusStruct)
	return result, err
}

//StorageTestList - Returns the list of previously executed tests.
//
//return []*StorageTestStruct
func (c *Client) StorageTestList() ([]*StorageTestStruct, error) {
	var err error
	var result []*StorageTestStruct
	obj, err := c.executeListAndReturnResults("/storage/test")
	for _, i := range obj {
		result = append(result, i.(*StorageTestStruct))
	}
	return result, err
}

//StorageTestCreate - Create a new StorageTest object.
//
//payload *StorageTestParametersStruct object
//return string
func (c *Client) StorageTestCreate(payload *StorageTestParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//StorageTestRead - Retrieve the specified StorageTest object.
//
//r - StorageTest Object Reference
//return *StorageTestStruct
func (c *Client) StorageTestRead(r string) (*StorageTestStruct, error) {
	var err error
	var result *StorageTestStruct
	obj, err := c.FindObjectByReference("/storage/test", r)
	result = obj.(*StorageTestStruct)
	return result, err
}

//StorageTestDelete - Delete the specified StorageTest object.
//
//r - StorageTest Object Reference
func (c *Client) StorageTestDelete(r string) error {
	var err error
	return err
}

//SupportAccessStateRead - Retrieve the specified SupportAccessState object.
//
//r - SupportAccessState Object Reference
//return *SupportAccessStateStruct
func (c *Client) SupportAccessStateRead(r string) (*SupportAccessStateStruct, error) {
	var err error
	var result *SupportAccessStateStruct
	obj, err := c.FindObjectByReference("/service/support/access", r)
	result = obj.(*SupportAccessStateStruct)
	return result, err
}

//SupportAccessStateUpdate - Update the specified SupportAccessState object.
//
//r - SupportAccessState Object Reference
//payload *SupportAccessStateStruct object
func (c *Client) SupportAccessStateUpdate(r string, payload *SupportAccessStateStruct) error {
	var err error
	return err
}

//SyslogConfigRead - Retrieve the specified SyslogConfig object.
//
//r - SyslogConfig Object Reference
//return *SyslogConfigStruct
func (c *Client) SyslogConfigRead(r string) (*SyslogConfigStruct, error) {
	var err error
	var result *SyslogConfigStruct
	obj, err := c.FindObjectByReference("/service/syslog", r)
	result = obj.(*SyslogConfigStruct)
	return result, err
}

//SyslogConfigUpdate - Update the specified SyslogConfig object.
//
//r - SyslogConfig Object Reference
//payload *SyslogConfigStruct object
func (c *Client) SyslogConfigUpdate(r string, payload *SyslogConfigStruct) error {
	var err error
	return err
}

//SystemInfoUpdate - Update the specified SystemInfo object.
//
//r - SystemInfo Object Reference
//payload *SystemInfoStruct object
func (c *Client) SystemInfoUpdate(r string, payload *SystemInfoStruct) error {
	var err error
	return err
}

//SystemInfoRead - Retrieve the specified SystemInfo object.
//
//r - SystemInfo Object Reference
//return *SystemInfoStruct
func (c *Client) SystemInfoRead(r string) (*SystemInfoStruct, error) {
	var err error
	var result *SystemInfoStruct
	obj, err := c.FindObjectByReference("/system", r)
	result = obj.(*SystemInfoStruct)
	return result, err
}

//SystemPackageList - List the packages that can be changed via the web services.
//
//return []*SystemPackageStruct
func (c *Client) SystemPackageList() ([]*SystemPackageStruct, error) {
	var err error
	var result []*SystemPackageStruct
	obj, err := c.executeListAndReturnResults("/system/package")
	for _, i := range obj {
		result = append(result, i.(*SystemPackageStruct))
	}
	return result, err
}

//SystemPackageUpdate - Update the specified SystemPackage object.
//
//r - SystemPackage Object Reference
//payload *SystemPackageStruct object
func (c *Client) SystemPackageUpdate(r string, payload *SystemPackageStruct) error {
	var err error
	return err
}

//SystemPackageRead - Retrieve the specified SystemPackage object.
//
//r - SystemPackage Object Reference
//return *SystemPackageStruct
func (c *Client) SystemPackageRead(r string) (*SystemPackageStruct, error) {
	var err error
	var result *SystemPackageStruct
	obj, err := c.FindObjectByReference("/system/package", r)
	result = obj.(*SystemPackageStruct)
	return result, err
}

//SystemVersionList - List SystemVersion objects on the system.
//
//return []*SystemVersionStruct
func (c *Client) SystemVersionList() ([]*SystemVersionStruct, error) {
	var err error
	var result []*SystemVersionStruct
	obj, err := c.executeListAndReturnResults("/system/version")
	for _, i := range obj {
		result = append(result, i.(*SystemVersionStruct))
	}
	return result, err
}

//SystemVersionRead - Retrieve the specified SystemVersion object.
//
//r - SystemVersion Object Reference
//return *SystemVersionStruct
func (c *Client) SystemVersionRead(r string) (*SystemVersionStruct, error) {
	var err error
	var result *SystemVersionStruct
	obj, err := c.FindObjectByReference("/system/version", r)
	result = obj.(*SystemVersionStruct)
	return result, err
}

//SystemVersionDelete - Delete the specified SystemVersion object.
//
//r - SystemVersion Object Reference
func (c *Client) SystemVersionDelete(r string) error {
	var err error
	return err
}

//TimeConfigRead - Retrieve the specified TimeConfig object.
//
//r - TimeConfig Object Reference
//return *TimeConfigStruct
func (c *Client) TimeConfigRead(r string) (*TimeConfigStruct, error) {
	var err error
	var result *TimeConfigStruct
	obj, err := c.FindObjectByReference("/service/time", r)
	result = obj.(*TimeConfigStruct)
	return result, err
}

//TimeConfigUpdate - Update the specified TimeConfig object.
//
//r - TimeConfig Object Reference
//payload *TimeConfigStruct object
func (c *Client) TimeConfigUpdate(r string, payload *TimeConfigStruct) error {
	var err error
	return err
}

//TimeZoneList - Lists all of the supported time zones.
//
//date ==>
//   description: The offset of this time zone from UTC is calculated at this date. If Daylight Saving Time is in effect at the specified date, the offset value is adjusted with the amount of daylight saving. If no date is specified then a default time of now is used.
//   type: string
//   format: date
//return []*TimeZoneStruct
func (c *Client) TimeZoneList(date string) ([]*TimeZoneStruct, error) {
	var err error
	var result []*TimeZoneStruct
	obj, err := c.executeListAndReturnResults("/timezone")
	for _, i := range obj {
		result = append(result, i.(*TimeZoneStruct))
	}
	return result, err
}

//TimeflowRead - Retrieve the specified Timeflow object.
//
//r - Timeflow Object Reference
//return Timeflow
func (c *Client) TimeflowRead(r string) (Timeflow, error) {
	var err error
	var result Timeflow
	obj, err := c.FindObjectByReference("/timeflow", r)
	result = obj.(Timeflow)
	return result, err
}

//TimeflowUpdate - Update the specified Timeflow object.
//
//r - Timeflow Object Reference
//payload Timeflow object
func (c *Client) TimeflowUpdate(r string, payload Timeflow) error {
	var err error
	return err
}

//TimeflowList - List Timeflow objects on the system.
//
//database ==>
//   type: string
//   description: List only TimeFlows within this database.
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//return []Timeflow
func (c *Client) TimeflowList(database string) ([]Timeflow, error) {
	var err error
	var result []Timeflow
	obj, err := c.executeListAndReturnResults("/timeflow")
	for _, i := range obj {
		result = append(result, i.(Timeflow))
	}
	return result, err
}

//TimeflowBookmarkList - Returns a list of all TimeFlow bookmarks.
//
//database ==>
//   mapsTo: timeflow.container
//   type: string
//   description: Filter results based on specified database.
//   format: objectReference
//   referenceTo: /delphix-container.json
//return []*TimeflowBookmarkStruct
func (c *Client) TimeflowBookmarkList(database string) ([]*TimeflowBookmarkStruct, error) {
	var err error
	var result []*TimeflowBookmarkStruct
	obj, err := c.executeListAndReturnResults("/timeflow/bookmark")
	for _, i := range obj {
		result = append(result, i.(*TimeflowBookmarkStruct))
	}
	return result, err
}

//TimeflowBookmarkCreate - Create a new TimeflowBookmark object.
//
//payload *TimeflowBookmarkCreateParametersStruct object
//return string
func (c *Client) TimeflowBookmarkCreate(payload *TimeflowBookmarkCreateParametersStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//TimeflowBookmarkRead - Retrieve the specified TimeflowBookmark object.
//
//r - TimeflowBookmark Object Reference
//return *TimeflowBookmarkStruct
func (c *Client) TimeflowBookmarkRead(r string) (*TimeflowBookmarkStruct, error) {
	var err error
	var result *TimeflowBookmarkStruct
	obj, err := c.FindObjectByReference("/timeflow/bookmark", r)
	result = obj.(*TimeflowBookmarkStruct)
	return result, err
}

//TimeflowBookmarkDelete - Delete the specified TimeflowBookmark object.
//
//r - TimeflowBookmark Object Reference
func (c *Client) TimeflowBookmarkDelete(r string) error {
	var err error
	return err
}

//TimeflowSnapshotRead - Retrieve the specified TimeflowSnapshot object.
//
//r - TimeflowSnapshot Object Reference
//return TimeflowSnapshot
func (c *Client) TimeflowSnapshotRead(r string) (TimeflowSnapshot, error) {
	var err error
	var result TimeflowSnapshot
	obj, err := c.FindObjectByReference("/snapshot", r)
	result = obj.(TimeflowSnapshot)
	return result, err
}

//TimeflowSnapshotList - Returns a list of snapshots on the system or within a particular
// object. By default, all snapshots within the domain are listed.
//
//fromDate ==>
//   description: Start date to use for filtering out results.
//   format: date
//   type: string
//   mapsTo: latestChangePoint.timestamp
//   inequalityType: NON-STRICT
//pageSize ==>
//   type: integer
//   minimum: 0
//   description: Limit the number of snapshots returned.
//pageOffset ==>
//   description: Offset within TimeFlow snapshots, in units of pageSize chunks. The pageOffset query parameter is only supported when either a 'TimeFlow' or 'database' query parameter is also set.
//   type: integer
//toDate ==>
//   type: string
//   mapsTo: latestChangePoint.timestamp
//   inequalityType: STRICT
//   description: End date to use for filtering out results.
//   format: date
//traverseTimeflows ==>
//   description: Whether to restrict snapshots to those in the current TimeFlow and in parent TimeFlows older than the branch point. This option is only used with the "database" option. The default behavior is false, i.e. show all snapshots.
//   type: boolean
//missingNonLoggedDataOnly ==>
//   description: Whether to restrict snapshots to those missing nologging changes. The defaultbehavior is salse.
//   type: boolean
//database ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//   description: Restrict snapshots to those within a TimeFlow of the specified database. This option is mutually exclusive with the "TimeFlow" option.
//timeflow ==>
//   description: Restrict snapshots to those within the specified TimeFlow. This option is mutually exclusive with the "database" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-timeflow.json
//   mapsTo: timeflow
//return []TimeflowSnapshot
func (c *Client) TimeflowSnapshotList(fromDate string, pageSize *int, pageOffset *int, toDate string, traverseTimeflows *bool, missingNonLoggedDataOnly *bool, database string, timeflow string) ([]TimeflowSnapshot, error) {
	var err error
	var result []TimeflowSnapshot
	obj, err := c.executeListAndReturnResults("/snapshot")
	for _, i := range obj {
		result = append(result, i.(TimeflowSnapshot))
	}
	return result, err
}

//TimeflowSnapshotUpdate - Update the specified TimeflowSnapshot object.
//
//r - TimeflowSnapshot Object Reference
//payload TimeflowSnapshot object
func (c *Client) TimeflowSnapshotUpdate(r string, payload TimeflowSnapshot) error {
	var err error
	return err
}

//TimeflowSnapshotDelete - Delete the specified TimeflowSnapshot object.
//
//r - TimeflowSnapshot Object Reference
func (c *Client) TimeflowSnapshotDelete(r string) error {
	var err error
	return err
}

//ToolkitList - Lists installed toolkits.
//
//sourceEnvironment ==>
//   referenceTo: /delphix-source-environment.json
//   description: Restricts list to include only toolkits that are valid for the given source environment.
//   type: string
//   format: objectReference
//return []*ToolkitStruct
func (c *Client) ToolkitList(sourceEnvironment string) ([]*ToolkitStruct, error) {
	var err error
	var result []*ToolkitStruct
	obj, err := c.executeListAndReturnResults("/toolkit")
	for _, i := range obj {
		result = append(result, i.(*ToolkitStruct))
	}
	return result, err
}

//ToolkitRead - Retrieve the specified Toolkit object.
//
//r - Toolkit Object Reference
//return *ToolkitStruct
func (c *Client) ToolkitRead(r string) (*ToolkitStruct, error) {
	var err error
	var result *ToolkitStruct
	obj, err := c.FindObjectByReference("/toolkit", r)
	result = obj.(*ToolkitStruct)
	return result, err
}

//ToolkitDelete - Delete the specified Toolkit object.
//
//r - Toolkit Object Reference
func (c *Client) ToolkitDelete(r string) error {
	var err error
	return err
}

//TransformationList - Returns a list of all transformations on the system.
//
//container ==>
//   referenceTo: /delphix-container.json
//   mapsTo: container
//   description: Return the transformation responsible for the given container reference.
//   type: string
//   format: objectReference
//parentContainer ==>
//   description: List the transformations that have been created against the provided container.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//return []*TransformationStruct
func (c *Client) TransformationList(container string, parentContainer string) ([]*TransformationStruct, error) {
	var err error
	var result []*TransformationStruct
	obj, err := c.executeListAndReturnResults("/transformation")
	for _, i := range obj {
		result = append(result, i.(*TransformationStruct))
	}
	return result, err
}

//TransformationRead - Retrieve the specified Transformation object.
//
//r - Transformation Object Reference
//return *TransformationStruct
func (c *Client) TransformationRead(r string) (*TransformationStruct, error) {
	var err error
	var result *TransformationStruct
	obj, err := c.FindObjectByReference("/transformation", r)
	result = obj.(*TransformationStruct)
	return result, err
}

//TransformationUpdate - Update the specified Transformation object.
//
//r - Transformation Object Reference
//payload *TransformationStruct object
func (c *Client) TransformationUpdate(r string, payload *TransformationStruct) error {
	var err error
	return err
}

//UpgradeCheckResultList - List UpgradeCheckResult objects on the system.
//
//return []*UpgradeCheckResultStruct
func (c *Client) UpgradeCheckResultList() ([]*UpgradeCheckResultStruct, error) {
	var err error
	var result []*UpgradeCheckResultStruct
	obj, err := c.executeListAndReturnResults("/system/upgradeCheckResult")
	for _, i := range obj {
		result = append(result, i.(*UpgradeCheckResultStruct))
	}
	return result, err
}

//UpgradeCheckResultRead - Retrieve the specified UpgradeCheckResult object.
//
//r - UpgradeCheckResult Object Reference
//return *UpgradeCheckResultStruct
func (c *Client) UpgradeCheckResultRead(r string) (*UpgradeCheckResultStruct, error) {
	var err error
	var result *UpgradeCheckResultStruct
	obj, err := c.FindObjectByReference("/system/upgradeCheckResult", r)
	result = obj.(*UpgradeCheckResultStruct)
	return result, err
}

//UserRead - Retrieve the specified User object.
//
//r - User Object Reference
//return *UserStruct
func (c *Client) UserRead(r string) (*UserStruct, error) {
	var err error
	var result *UserStruct
	obj, err := c.FindObjectByReference("/user", r)
	result = obj.(*UserStruct)
	return result, err
}

//UserList - Lists users in the system.
//
//typeUser ==>
//   description: User type SYSTEM or DOMAIN.
//   type: string
//   default: DOMAIN
//   enum: [SYSTEM DOMAIN]
//   mapsTo: userType
//domainUserType ==>
//   type: string
//   enum: [DELPHIX_ADMIN STANDARD_USER SELFSERVICE_ONLY]
//   required: false
//   description: DOMAIN user type: DELPHIX_ADMIN, STANDARD_USER, or SELFSERVICE_ONLY.
//return []*UserStruct
func (c *Client) UserList(typeUser string, domainUserType string) ([]*UserStruct, error) {
	var err error
	var result []*UserStruct
	obj, err := c.executeListAndReturnResults("/user")
	for _, i := range obj {
		result = append(result, i.(*UserStruct))
	}
	return result, err
}

//UserCreate - Create a new User object.
//
//payload *UserStruct object
//return string
func (c *Client) UserCreate(payload *UserStruct) (string, error) {
	var err error
	var result string
	return result, err
}

//UserUpdate - Update the specified User object.
//
//r - User Object Reference
//payload *UserStruct object
func (c *Client) UserUpdate(r string, payload *UserStruct) error {
	var err error
	return err
}

//UserDelete - Delete the specified User object.
//
//r - User Object Reference
func (c *Client) UserDelete(r string) error {
	var err error
	return err
}

//UserInterfaceConfigRead - Retrieve the specified UserInterfaceConfig object.
//
//r - UserInterfaceConfig Object Reference
//return *UserInterfaceConfigStruct
func (c *Client) UserInterfaceConfigRead(r string) (*UserInterfaceConfigStruct, error) {
	var err error
	var result *UserInterfaceConfigStruct
	obj, err := c.FindObjectByReference("/service/userInterface", r)
	result = obj.(*UserInterfaceConfigStruct)
	return result, err
}

//UserInterfaceConfigUpdate - Update the specified UserInterfaceConfig object.
//
//r - UserInterfaceConfig Object Reference
//payload *UserInterfaceConfigStruct object
func (c *Client) UserInterfaceConfigUpdate(r string, payload *UserInterfaceConfigStruct) error {
	var err error
	return err
}

//WindowsClusterNodeRead - Retrieve the specified WindowsClusterNode object.
//
//r - WindowsClusterNode Object Reference
//return *WindowsClusterNodeStruct
func (c *Client) WindowsClusterNodeRead(r string) (*WindowsClusterNodeStruct, error) {
	var err error
	var result *WindowsClusterNodeStruct
	obj, err := c.FindObjectByReference("/environment/windows/clusternode", r)
	result = obj.(*WindowsClusterNodeStruct)
	return result, err
}

//WindowsClusterNodeList - Returns a list of nodes filtered by cluster.
//
//cluster ==>
//   format: objectReference
//   referenceTo: /delphix-windows-cluster.json
//   description: A reference to the Windows cluster environment this node belongs to.
//   mapsTo: cluster
//   type: string
//return []*WindowsClusterNodeStruct
func (c *Client) WindowsClusterNodeList(cluster string) ([]*WindowsClusterNodeStruct, error) {
	var err error
	var result []*WindowsClusterNodeStruct
	obj, err := c.executeListAndReturnResults("/environment/windows/clusternode")
	for _, i := range obj {
		result = append(result, i.(*WindowsClusterNodeStruct))
	}
	return result, err
}

//X509CertificateRead - Retrieve the specified X509Certificate object.
//
//r - X509Certificate Object Reference
//return *X509CertificateStruct
func (c *Client) X509CertificateRead(r string) (*X509CertificateStruct, error) {
	var err error
	var result *X509CertificateStruct
	obj, err := c.FindObjectByReference("/service/certificate", r)
	result = obj.(*X509CertificateStruct)
	return result, err
}

//X509CertificateList - List X509Certificate objects on the system.
//
//return []*X509CertificateStruct
func (c *Client) X509CertificateList() ([]*X509CertificateStruct, error) {
	var err error
	var result []*X509CertificateStruct
	obj, err := c.executeListAndReturnResults("/service/certificate")
	for _, i := range obj {
		result = append(result, i.(*X509CertificateStruct))
	}
	return result, err
}

//X509CertificateDelete - Delete the specified X509Certificate object.
//
//r - X509Certificate Object Reference
func (c *Client) X509CertificateDelete(r string) error {
	var err error
	return err
}

func (c *Client) SourceEnvironmentListTest(typeSourceEnvironment string) ([]SourceEnvironment, error) {
	var err error
	var result []SourceEnvironment
	obj, err := c.executeListAndCastAndReturnResults("/environment")
	fmt.Printf("TESTING: %+v\n\n", obj)
	for _, i := range obj.Result.([]interface{}) {
		objMap := i.(map[string]interface{})
		switch objMap["type"].(string) {
		case "UnixHostEnvironment":
			var arrayItem UnixHostEnvironmentStruct
			mapstructure.Decode(objMap, &arrayItem)
			result = append(result, arrayItem)
			return result, err
		}
	}
	return result, err
}
