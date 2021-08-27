package delphix
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
//AbstractToolkitRead - Retrieve the specified AbstractToolkit object.
//
//r - AbstractToolkit Object Reference
//return AbstractToolkit
func (c *Client) AbstractToolkitRead(r string) (AbstractToolkit, error) {
var err error
var result AbstractToolkit
obj, err := c.FindObjectByReference("/toolkit", r)
result = obj.(AbstractToolkit)
return result, err
}
//AbstractToolkitList - Lists installed toolkits.
//
//sourceEnvironment ==>
//   referenceTo: /delphix-source-environment.json
//   description: Restricts list to include only toolkits that are valid for the given source environment.
//   type: string
//   format: objectReference
//return []AbstractToolkit
func (c *Client) AbstractToolkitList(sourceEnvironment string) ([]AbstractToolkit, error) {
var err error
var result []AbstractToolkit
return result, err
}
//AbstractToolkitDelete - Delete the specified AbstractToolkit object.
//
//r - AbstractToolkit Object Reference
func (c *Client) AbstractToolkitDelete(r string) ( error) {
var err error
return err
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
//toDate ==>
//   description: End date for the search. Only actions on or before this date are included.
//   type: string
//   format: date
//pageOffset ==>
//   description: Offset within event list, in units of pageSize chunks.
//   type: integer
//parentAction ==>
//   mapsTo: parentAction
//   description: Limit actions to those with this parent action.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-action.json
//user ==>
//   description: Limit actions to those initiated by this user.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user.json
//   mapsTo: user
//rootActionOnly ==>
//   type: boolean
//   description: Limit actions to those without a parent action.
//sortBy ==>
//   enum: [reference state title details startTime endTime failureDescription parentAction workSource workSourceName workSourcePrincipal]
//   description: Search results are sorted by the field provided.
//   type: string
//searchText ==>
//   description: Limit search results to only include alerts that have searchText string in reference, state, title, details, failureDescription, parentAction, workSource, workSourceName or workSourcePrincipal.
//   type: string
//fromDate ==>
//   description: Start date for the search. Only actions on or after this date are included.
//   type: string
//   format: date
//pageSize ==>
//   description: Limit the number of events returned.
//   type: integer
//   default: 25
//state ==>
//   description: Limit actions to those in the specified state.
//   type: string
//   enum: [EXECUTING WAITING COMPLETED FAILED CANCELED]
//   mapsTo: state
//ignoreActionTypes ==>
//   description: Ignore actions with an action type in this list.
//   type: array
//   items: map[type:string]
//ascending ==>
//   type: boolean
//   description: True if results are to be returned in ascending order.
//return []*ActionStruct
func (c *Client) ActionList(toDate string,pageOffset *int,parentAction string,user string,rootActionOnly *bool,sortBy string,searchText string,fromDate string,pageSize *int,state string,ignoreActionTypes []string,ascending *bool) ([]*ActionStruct, error) {
var err error
var result []*ActionStruct
return result, err
}
//AdvancedSettingsInfoRead - Retrieve the specified AdvancedSettingsInfo object.
//
//r - AdvancedSettingsInfo Object Reference
//return *AdvancedSettingsInfoStruct
func (c *Client) AdvancedSettingsInfoRead(r string) (*AdvancedSettingsInfoStruct, error) {
var err error
var result *AdvancedSettingsInfoStruct
obj, err := c.FindObjectByReference("/system/advancedSettings", r)
result = obj.(*AdvancedSettingsInfoStruct)
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
//target ==>
//   referenceTo: /delphix-user-object.json
//   mapsTo: target
//   description: Limit alerts to those affecting a particular object on the system.
//   type: string
//   format: objectReference
//toDate ==>
//   description: End date to use for the search.
//   type: string
//   format: date
//   mapsTo: timestamp
//   inequalityType: STRICT
//searchText ==>
//   description: Limit search results to only include alerts that have searchText string in eventTitle, eventDescription, eventResponse, eventAction, or severity.
//   type: string
//maxTotal ==>
//   description: The upper bound for calculation of total alert count.
//   type: integer
//ascending ==>
//   type: boolean
//   description: True if results are to be returned in ascending order.
//sortBy ==>
//   type: string
//   enum: [event eventTitle eventDescription eventResponse eventAction eventCommandOutput eventSeverity target targetName timestamp]
//   description: Search results are sorted by the field provided.
//fromDate ==>
//   description: Start date to use for the search.
//   type: string
//   format: date
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//pageSize ==>
//   description: Limit the number of alerts returned.
//   type: integer
//   default: 25
//   minimum: 0
//pageOffset ==>
//   type: integer
//   description: Offset within alert list, in units of pageSize chunks.
//return []*AlertStruct
func (c *Client) AlertList(target string,toDate string,searchText string,maxTotal *int,ascending *bool,sortBy string,fromDate string,pageSize *int,pageOffset *int) ([]*AlertStruct, error) {
var err error
var result []*AlertStruct
return result, err
}
//AlertProfileDelete - Delete the specified AlertProfile object.
//
//r - AlertProfile Object Reference
func (c *Client) AlertProfileDelete(r string) ( error) {
var err error
return err
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
//AlertProfileCreate - Create a new AlertProfile object.
//
//payload *AlertProfileStruct object
//return string
func (c *Client) AlertProfileCreate(payload *AlertProfileStruct) (string, error) {
var err error
var result string
return result, err
}
//AlertProfileUpdate - Update the specified AlertProfile object.
//
//r - AlertProfile Object Reference
//payload *AlertProfileStruct object
func (c *Client) AlertProfileUpdate(r string,payload *AlertProfileStruct) ( error) {
var err error
return err
}
//AlertProfileList - List AlertProfile objects on the system.
//
//return []*AlertProfileStruct
func (c *Client) AlertProfileList() ([]*AlertProfileStruct, error) {
var err error
var result []*AlertProfileStruct
return result, err
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
//AuthorizationCreate - Create a new Authorization object.
//
//payload *AuthorizationStruct object
//return string
func (c *Client) AuthorizationCreate(payload *AuthorizationStruct) (string, error) {
var err error
var result string
return result, err
}
//AuthorizationDelete - Delete the specified Authorization object.
//
//r - Authorization Object Reference
func (c *Client) AuthorizationDelete(r string) ( error) {
var err error
return err
}
//AuthorizationList - Lists authorizations granted in the system.
//
//user ==>
//   description: Lists permissions granted to the specified user.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user.json
//   required: false
//   mapsTo: user
//target ==>
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   required: false
//   mapsTo: target
//   description: Lists the permissions granted on the specified object.
//   type: string
//effective ==>
//   required: false
//   description: Also return inherited authorizations.
//   type: string
//return []*AuthorizationStruct
func (c *Client) AuthorizationList(user string,target string,effective string) ([]*AuthorizationStruct, error) {
var err error
var result []*AuthorizationStruct
return result, err
}
//AuthorizationConfigUpdate - Update the specified AuthorizationConfig object.
//
//r - AuthorizationConfig Object Reference
//payload *AuthorizationConfigStruct object
func (c *Client) AuthorizationConfigUpdate(r string,payload *AuthorizationConfigStruct) ( error) {
var err error
return err
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
//CaCertificateDelete - Delete the specified CaCertificate object.
//
//r - CaCertificate Object Reference
func (c *Client) CaCertificateDelete(r string) ( error) {
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
//CertificateConfigUpdate - Update the specified CertificateConfig object.
//
//r - CertificateConfig Object Reference
//payload *CertificateConfigStruct object
func (c *Client) CertificateConfigUpdate(r string,payload *CertificateConfigStruct) ( error) {
var err error
return err
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
//CertificateSigningRequestCreate - Create a new CertificateSigningRequest object.
//
//payload *CertificateSigningRequestCreateParametersStruct object
//return string
func (c *Client) CertificateSigningRequestCreate(payload *CertificateSigningRequestCreateParametersStruct) (string, error) {
var err error
var result string
return result, err
}
//CertificateSigningRequestDelete - Delete the specified CertificateSigningRequest object.
//
//r - CertificateSigningRequest Object Reference
func (c *Client) CertificateSigningRequestDelete(r string) ( error) {
var err error
return err
}
//CertificateSigningRequestList - List CertificateSigningRequest objects on the system.
//
//return []*CertificateSigningRequestStruct
func (c *Client) CertificateSigningRequestList() ([]*CertificateSigningRequestStruct, error) {
var err error
var result []*CertificateSigningRequestStruct
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
//CipherSuiteList - Lists all of the supported cipher suites.
//
//httpsDefault ==>
//   description: If true, returns just the default list of ciphers for HTTPS.
//   type: boolean
//return []*CipherSuiteStruct
func (c *Client) CipherSuiteList(httpsDefault *bool) ([]*CipherSuiteStruct, error) {
var err error
var result []*CipherSuiteStruct
return result, err
}
//CloudStatusRead - Retrieve the specified CloudStatus object.
//
//r - CloudStatus Object Reference
//return *CloudStatusStruct
func (c *Client) CloudStatusRead(r string) (*CloudStatusStruct, error) {
var err error
var result *CloudStatusStruct
obj, err := c.FindObjectByReference("/service/cloud", r)
result = obj.(*CloudStatusStruct)
return result, err
}
//CloudStatusUpdate - Update the specified CloudStatus object.
//
//r - CloudStatus Object Reference
//payload *CloudStatusStruct object
func (c *Client) CloudStatusUpdate(r string,payload *CloudStatusStruct) ( error) {
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
//ContainerList - Returns a list of databases on the system or within a group.
//
//noJSContainerDataSource ==>
//   description: Restrict databases to those which are not part of a Self-Service data container. This option is mutually exclusive with the "noJSDataSource" option.
//   type: boolean
//validForSecureReplication ==>
//   description: Restrict listing to include only datasets that can be securely replicated.
//   type: boolean
//provisionContainer ==>
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: provisionContainer
//   description: Restrict databases to those provisioned from the specified container. This option is mutually exclusive with the "group" option.
//   type: string
//group ==>
//   description: Restrict databases to those within the specified group. This option is mutually exclusive with the "provisionContainer" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-group.json
//   mapsTo: group
//noJSDataSource ==>
//   description: Restrict databases to those which are not part of a Self-Service data layout (data template or data container). This option is mutually exclusive with the "noJSContainerDataSource" option.
//   type: boolean
//return []Container
func (c *Client) ContainerList(noJSContainerDataSource *bool,validForSecureReplication *bool,provisionContainer string,group string,noJSDataSource *bool) ([]Container, error) {
var err error
var result []Container
return result, err
}
//ContainerDelete - Delete the specified Container object.
//
//r - Container Object Reference
//payload *DeleteParametersStruct object
func (c *Client) ContainerDelete(r string,payload *DeleteParametersStruct) ( error) {
var err error
return err
}
//ContainerUpdate - Update the specified Container object.
//
//r - Container Object Reference
//payload Container object
func (c *Client) ContainerUpdate(r string,payload Container) ( error) {
var err error
return err
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
func (c *Client) ContainerUtilizationList(fromDate string,toDate string,samplingInterval float64) ([]*ContainerUtilizationStruct, error) {
var err error
var result []*ContainerUtilizationStruct
return result, err
}
//CurrentConsumerCapacityDataList - Lists consumers in the system.
//
//return []*CurrentConsumerCapacityDataStruct
func (c *Client) CurrentConsumerCapacityDataList() ([]*CurrentConsumerCapacityDataStruct, error) {
var err error
var result []*CurrentConsumerCapacityDataStruct
return result, err
}
//CurrentGroupCapacityDataList - Lists capacity data for groups in the system.
//
//return []*CurrentGroupCapacityDataStruct
func (c *Client) CurrentGroupCapacityDataList() ([]*CurrentGroupCapacityDataStruct, error) {
var err error
var result []*CurrentGroupCapacityDataStruct
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
//DNSConfigUpdate - Update the specified DNSConfig object.
//
//r - DNSConfig Object Reference
//payload *DNSConfigStruct object
func (c *Client) DNSConfigUpdate(r string,payload *DNSConfigStruct) ( error) {
var err error
return err
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
//DatabaseTemplateUpdate - Update the specified DatabaseTemplate object.
//
//r - DatabaseTemplate Object Reference
//payload *DatabaseTemplateStruct object
func (c *Client) DatabaseTemplateUpdate(r string,payload *DatabaseTemplateStruct) ( error) {
var err error
return err
}
//DatabaseTemplateDelete - Delete the specified DatabaseTemplate object.
//
//r - DatabaseTemplate Object Reference
func (c *Client) DatabaseTemplateDelete(r string) ( error) {
var err error
return err
}
//DatabaseTemplateList - List DatabaseTemplate objects on the system.
//
//return []*DatabaseTemplateStruct
func (c *Client) DatabaseTemplateList() ([]*DatabaseTemplateStruct, error) {
var err error
var result []*DatabaseTemplateStruct
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
//EndEntityCertificateList - List EndEntityCertificate objects on the system.
//
//return []*EndEntityCertificateStruct
func (c *Client) EndEntityCertificateList() ([]*EndEntityCertificateStruct, error) {
var err error
var result []*EndEntityCertificateStruct
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
//EnvironmentUserDelete - Delete the specified EnvironmentUser object.
//
//r - EnvironmentUser Object Reference
//payload *DeleteParametersStruct object
func (c *Client) EnvironmentUserDelete(r string,payload *DeleteParametersStruct) ( error) {
var err error
return err
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
func (c *Client) EnvironmentUserUpdate(r string,payload *EnvironmentUserStruct) ( error) {
var err error
return err
}
//EnvironmentUserList - Returns the list of all environment users in the system.
//
//environment ==>
//   format: objectReference
//   referenceTo: /delphix-source-environment.json
//   mapsTo: environment
//   type: string
//   description: Limit results to users within the given environment.
//return []*EnvironmentUserStruct
func (c *Client) EnvironmentUserList(environment string) ([]*EnvironmentUserStruct, error) {
var err error
var result []*EnvironmentUserStruct
return result, err
}
//FaultList - Returns the list of all the faults that match the given criteria.
//
//pageOffset ==>
//   type: integer
//   description: Offset within fault list, in units of pageSize chunks.
//maxTotal ==>
//   description: The upper bound for calculation of total fault count.
//   type: integer
//ascending ==>
//   type: boolean
//   description: True if results are to be returned in ascending order.
//searchText ==>
//   description: Limit search results to only include faults that have searchText string in severity, title, targetName.
//   type: string
//target ==>
//   description: The reference to the Delphix user-visible object associated with the fault.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   mapsTo: target
//severity ==>
//   description: The impact of the fault on the system: CRITICAL or WARNING.
//   type: string
//   enum: [CRITICAL WARNING]
//   mapsTo: severity
//status ==>
//   description: The status of the fault: ACTIVE, RESOLVED or IGNORED.
//   type: string
//   enum: [ACTIVE RESOLVED IGNORED]
//   mapsTo: status
//fromDate ==>
//   description: Start date to use for the search.
//   type: string
//   format: date
//   mapsTo: dateDiagnosed
//   inequalityType: NON-STRICT
//targets ==>
//   type: array
//   items: map[format:objectReference referenceTo:/delphix-user-object.json type:string]
//   description: The list of references to Delphix user-visible objects associated with the faults.
//   mapsTo: target
//   minItems: 1
//toDate ==>
//   mapsTo: dateDiagnosed
//   inequalityType: STRICT
//   description: End date to use for the search.
//   type: string
//   format: date
//pageSize ==>
//   minimum: 0
//   description: Limit the number of faults returned.
//   type: integer
//   default: 25
//sortBy ==>
//   type: string
//   enum: [severity dateDiagnosed title targetName]
//   description: Search results are sorted by the field provided.
//return []*FaultStruct
func (c *Client) FaultList(pageOffset *int,maxTotal *int,ascending *bool,searchText string,target string,severity string,status string,fromDate string,targets []string,toDate string,pageSize *int,sortBy string) ([]*FaultStruct, error) {
var err error
var result []*FaultStruct
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
//targets ==>
//   minItems: 1
//   type: array
//   items: map[format:objectReference referenceTo:/delphix-user-object.json type:string]
//   description: The references to the Delphix user-visible objects associated with the fault effects.
//   mapsTo: target
//rootCause ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-fault.json
//   mapsTo: rootCause
//   description: The reference to the fault which is the root cause of the fault effect.
//bundleID ==>
//   description: A unique dot delimited identifier associated with the fault effect.
//   type: string
//   mapsTo: bundleID
//return []*FaultEffectStruct
func (c *Client) FaultEffectList(severity string,target string,targets []string,rootCause string,bundleID string) ([]*FaultEffectStruct, error) {
var err error
var result []*FaultEffectStruct
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
func (c *Client) GlobalLinkingSettingsUpdate(r string,payload *GlobalLinkingSettingsStruct) ( error) {
var err error
return err
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
//GroupUpdate - Update the specified Group object.
//
//r - Group Object Reference
//payload *GroupStruct object
func (c *Client) GroupUpdate(r string,payload *GroupStruct) ( error) {
var err error
return err
}
//GroupDelete - Delete the specified Group object.
//
//r - Group Object Reference
func (c *Client) GroupDelete(r string) ( error) {
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
//GroupList - List Group objects on the system.
//
//return []*GroupStruct
func (c *Client) GroupList() ([]*GroupStruct, error) {
var err error
var result []*GroupStruct
return result, err
}
//HeldSpaceCapacityDataList - Lists HeldSpace in the system.
//
//return []*HeldSpaceCapacityDataStruct
func (c *Client) HeldSpaceCapacityDataList() ([]*HeldSpaceCapacityDataStruct, error) {
var err error
var result []*HeldSpaceCapacityDataStruct
return result, err
}
//HistoricalConsumerCapacityDataList - Lists consumers in the system.
//
//container ==>
//   format: objectReference
//   referenceTo: /delphix-container.json
//   description: The container for which to list data.
//   type: string
//startDate ==>
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//   type: string
//   format: date
//   description: Earliest date for which to list data.
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
func (c *Client) HistoricalConsumerCapacityDataList(container string,startDate string,endDate string,resolution *int) ([]*HistoricalConsumerCapacityDataStruct, error) {
var err error
var result []*HistoricalConsumerCapacityDataStruct
return result, err
}
//HistoricalGroupCapacityDataList - Lists historical capacity data for a particular group.
//
//group ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-group.json
//   description: The group to list data for.
//startDate ==>
//   description: Earliest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//   type: string
//   format: date
//endDate ==>
//   inequalityType: NON-STRICT
//   type: string
//   format: date
//   description: Latest date for which to list data.
//   mapsTo: timestamp
//resolution ==>
//   type: integer
//   minimum: 0
//   description: The time range each datapoint should represent, measured in seconds. This parameter is only meaningful if a group is specified.
//return []*HistoricalGroupCapacityDataStruct
func (c *Client) HistoricalGroupCapacityDataList(group string,startDate string,endDate string,resolution *int) ([]*HistoricalGroupCapacityDataStruct, error) {
var err error
var result []*HistoricalGroupCapacityDataStruct
return result, err
}
//HistoricalSystemCapacityDataList - Lists historical system capacity data.
//
//startDate ==>
//   type: string
//   format: date
//   description: Earliest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//endDate ==>
//   format: date
//   description: Latest date for which to list data.
//   mapsTo: timestamp
//   inequalityType: NON-STRICT
//   type: string
//resolution ==>
//   minimum: 0
//   description: The time range each datapoint should represent, measured in seconds.
//   type: integer
//return []*HistoricalSystemCapacityDataStruct
func (c *Client) HistoricalSystemCapacityDataList(startDate string,endDate string,resolution *int) ([]*HistoricalSystemCapacityDataStruct, error) {
var err error
var result []*HistoricalSystemCapacityDataStruct
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
func (c *Client) HostUpdate(r string,payload Host) ( error) {
var err error
return err
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
func (c *Client) HostPrivilegeElevationProfileUpdate(r string,payload *HostPrivilegeElevationProfileStruct) ( error) {
var err error
return err
}
//HostPrivilegeElevationProfileDelete - Delete the specified HostPrivilegeElevationProfile object.
//
//r - HostPrivilegeElevationProfile Object Reference
func (c *Client) HostPrivilegeElevationProfileDelete(r string) ( error) {
var err error
return err
}
//HostPrivilegeElevationProfileList - List HostPrivilegeElevationProfile objects on the system.
//
//return []*HostPrivilegeElevationProfileStruct
func (c *Client) HostPrivilegeElevationProfileList() ([]*HostPrivilegeElevationProfileStruct, error) {
var err error
var result []*HostPrivilegeElevationProfileStruct
return result, err
}
//HostPrivilegeElevationProfileScriptList - List HostPrivilegeElevationProfileScript objects on the system.
//
//return []*HostPrivilegeElevationProfileScriptStruct
func (c *Client) HostPrivilegeElevationProfileScriptList() ([]*HostPrivilegeElevationProfileScriptStruct, error) {
var err error
var result []*HostPrivilegeElevationProfileScriptStruct
return result, err
}
//HostPrivilegeElevationProfileScriptUpdate - Update the specified HostPrivilegeElevationProfileScript object.
//
//r - HostPrivilegeElevationProfileScript Object Reference
//payload *HostPrivilegeElevationProfileScriptStruct object
func (c *Client) HostPrivilegeElevationProfileScriptUpdate(r string,payload *HostPrivilegeElevationProfileScriptStruct) ( error) {
var err error
return err
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
func (c *Client) HostPrivilegeElevationProfileScriptDelete(r string) ( error) {
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
func (c *Client) HostPrivilegeElevationSettingsUpdate(r string,payload *HostPrivilegeElevationSettingsStruct) ( error) {
var err error
return err
}
//HttpConnectorConfigUpdate - Update the specified HttpConnectorConfig object.
//
//r - HttpConnectorConfig Object Reference
//payload *HttpConnectorConfigStruct object
func (c *Client) HttpConnectorConfigUpdate(r string,payload *HttpConnectorConfigStruct) ( error) {
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
//IscsiInitiatorRead - Retrieve the specified IscsiInitiator object.
//
//r - IscsiInitiator Object Reference
//return *IscsiInitiatorStruct
func (c *Client) IscsiInitiatorRead(r string) (*IscsiInitiatorStruct, error) {
var err error
var result *IscsiInitiatorStruct
obj, err := c.FindObjectByReference("/storage/iscsi/initiator", r)
result = obj.(*IscsiInitiatorStruct)
return result, err
}
//IscsiTargetRead - Retrieve the specified IscsiTarget object.
//
//r - IscsiTarget Object Reference
//return *IscsiTargetStruct
func (c *Client) IscsiTargetRead(r string) (*IscsiTargetStruct, error) {
var err error
var result *IscsiTargetStruct
obj, err := c.FindObjectByReference("/storage/iscsi/target", r)
result = obj.(*IscsiTargetStruct)
return result, err
}
//IscsiTargetCreate - Create a new IscsiTarget object.
//
//payload *IscsiTargetStruct object
//return string
func (c *Client) IscsiTargetCreate(payload *IscsiTargetStruct) (string, error) {
var err error
var result string
return result, err
}
//IscsiTargetUpdate - Update the specified IscsiTarget object.
//
//r - IscsiTarget Object Reference
//payload *IscsiTargetStruct object
func (c *Client) IscsiTargetUpdate(r string,payload *IscsiTargetStruct) ( error) {
var err error
return err
}
//IscsiTargetList - List IscsiTarget objects on the system.
//
//return []*IscsiTargetStruct
func (c *Client) IscsiTargetList() ([]*IscsiTargetStruct, error) {
var err error
var result []*IscsiTargetStruct
return result, err
}
//IscsiTargetDelete - Delete the specified IscsiTarget object.
//
//r - IscsiTarget Object Reference
func (c *Client) IscsiTargetDelete(r string) ( error) {
var err error
return err
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
//JSBookmarkCreate - Create a new JSBookmark object.
//
//payload *JSBookmarkCreateParametersStruct object
//return string
func (c *Client) JSBookmarkCreate(payload *JSBookmarkCreateParametersStruct) (string, error) {
var err error
var result string
return result, err
}
//JSBookmarkUpdate - Update the specified JSBookmark object.
//
//r - JSBookmark Object Reference
//payload *JSBookmarkStruct object
func (c *Client) JSBookmarkUpdate(r string,payload *JSBookmarkStruct) ( error) {
var err error
return err
}
//JSBookmarkDelete - Delete the specified JSBookmark object.
//
//r - JSBookmark Object Reference
func (c *Client) JSBookmarkDelete(r string) ( error) {
var err error
return err
}
//JSBookmarkList - Lists the Self-Service bookmarks in the system.
//
//container ==>
//   referenceTo: /delphix-js-data-container.json
//   mapsTo: container
//   description: List all usable bookmarks accessible to the specified data container. This option is mutually exclusive with all other options.
//   type: string
//   format: objectReference
//template ==>
//   description: List all usable bookmarks created in the specified data template. This option is mutually exclusive with all other options.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-template.json
//   mapsTo: template
//return []*JSBookmarkStruct
func (c *Client) JSBookmarkList(container string,template string) ([]*JSBookmarkStruct, error) {
var err error
var result []*JSBookmarkStruct
return result, err
}
//JSBranchDelete - Delete the specified JSBranch object.
//
//r - JSBranch Object Reference
func (c *Client) JSBranchDelete(r string) ( error) {
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
//JSBranchUpdate - Update the specified JSBranch object.
//
//r - JSBranch Object Reference
//payload *JSBranchStruct object
func (c *Client) JSBranchUpdate(r string,payload *JSBranchStruct) ( error) {
var err error
return err
}
//JSBranchList - Lists the Self-Service branches in the system.
//
//dataLayout ==>
//   description: List branches belonging to the given data layout.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-layout.json
//   mapsTo: dataLayout
//return []*JSBranchStruct
func (c *Client) JSBranchList(dataLayout string) ([]*JSBranchStruct, error) {
var err error
var result []*JSBranchStruct
return result, err
}
//JSConfigUpdate - Update the specified JSConfig object.
//
//r - JSConfig Object Reference
//payload *JSConfigStruct object
func (c *Client) JSConfigUpdate(r string,payload *JSConfigStruct) ( error) {
var err error
return err
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
//   referenceTo: /delphix-named-user-object.json
//   mapsTo: usageObject
//   description: Restrict usage data to that related to a specific object.
//   type: string
//   format: objectReference
//return []*JSDailyOperationDurationStruct
func (c *Client) JSDailyOperationDurationList(usageObject string) ([]*JSDailyOperationDurationStruct, error) {
var err error
var result []*JSDailyOperationDurationStruct
return result, err
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
//JSDataContainerUpdate - Update the specified JSDataContainer object.
//
//r - JSDataContainer Object Reference
//payload *JSDataContainerStruct object
func (c *Client) JSDataContainerUpdate(r string,payload *JSDataContainerStruct) ( error) {
var err error
return err
}
//JSDataContainerDelete - Delete this data container.
//
//r - JSDataContainer Object Reference
//payload *JSDataContainerDeleteParametersStruct object
func (c *Client) JSDataContainerDelete(r string,payload *JSDataContainerDeleteParametersStruct) ( error) {
var err error
return err
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
//   mapsTo: template
//   description: Restrict data containers to those provisioned from the specified template. This option is mutually exclusive with the "owner" and "independentOnly" options.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-template.json
//independentOnly ==>
//   description: Restrict data containers to independent data containers that do not have templates. This option is mutually exclusive with the "template" and "owner" options.
//   type: boolean
//return []*JSDataContainerStruct
func (c *Client) JSDataContainerList(owner string,template string,independentOnly *bool) ([]*JSDataContainerStruct, error) {
var err error
var result []*JSDataContainerStruct
return result, err
}
//JSDataSourceList - Lists the Self-Service data sources in the system.
//
//dataLayout ==>
//   description: List the sources associated with the given data layout reference.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-layout.json
//   mapsTo: dataLayout
//container ==>
//   description: List the source associated with the given container reference.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//return []*JSDataSourceStruct
func (c *Client) JSDataSourceList(dataLayout string,container string) ([]*JSDataSourceStruct, error) {
var err error
var result []*JSDataSourceStruct
return result, err
}
//JSDataSourceUpdate - Update the specified JSDataSource object.
//
//r - JSDataSource Object Reference
//payload *JSDataSourceStruct object
func (c *Client) JSDataSourceUpdate(r string,payload *JSDataSourceStruct) ( error) {
var err error
return err
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
//JSDataTemplateList - List the data templates defined in the system.
//
//return []*JSDataTemplateStruct
func (c *Client) JSDataTemplateList() ([]*JSDataTemplateStruct, error) {
var err error
var result []*JSDataTemplateStruct
return result, err
}
//JSDataTemplateDelete - Delete the specified JSDataTemplate object.
//
//r - JSDataTemplate Object Reference
func (c *Client) JSDataTemplateDelete(r string) ( error) {
var err error
return err
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
func (c *Client) JSDataTemplateUpdate(r string,payload *JSDataTemplateStruct) ( error) {
var err error
return err
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
//JSOperationList - Lists the Self-Service action history for a data layout.
//
//dataStartTime ==>
//   description: Operations with "dataTime" after this value will be returned. Used with "dataEndTime" to return a set of operations between two dates.
//   type: string
//   format: date
//dataEndTime ==>
//   description: Operations with "dataTime" before this value will be returned. Used with "dataStartTime" to return a set of operations between two dates.
//   type: string
//   format: date
//dataLayout ==>
//   mapsTo: dataLayout
//   description: Limit operations to the specific data layout. This option is mutually exclusive with the "branch" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-data-layout.json
//branch ==>
//   description: Limit operations to the specified branch. This option is mutually exclusive with the "dataLayout" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-js-branch.json
//   mapsTo: branch
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
//   type: integer
//   default: 0
//   description: The suggested maximum number of visible operations after "dataTime" that should be returned. If there are not sufficient events after additional events before may be returned.
//return []*JSOperationStruct
func (c *Client) JSOperationList(dataStartTime string,dataEndTime string,dataLayout string,branch string,dataTime string,beforeCount *int,afterCount *int) ([]*JSOperationStruct, error) {
var err error
var result []*JSOperationStruct
return result, err
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
return result, err
}
//JobList - Returns a list of jobs in the system. Jobs are listed in start
// time order.
//
//jobState ==>
//   mapsTo: jobState
//   description: Limit jobs to those in the specified job state.
//   type: string
//   enum: [RUNNING SUSPENDED CANCELED COMPLETED FAILED]
//pageSize ==>
//   description: Limit the number of jobs returned.
//   type: integer
//   default: 25
//   minimum: 0
//pageOffset ==>
//   type: integer
//   description: Page offset within job list.
//addEvents ==>
//   default: false
//   description: Whether to include the job events in each job.
//   type: boolean
//jobType ==>
//   description: Limit jobs to those with the specified job type.
//   type: string
//   mapsTo: actionType
//target ==>
//   description: Limit jobs to those affecting a particular object on the system. The target is the object reference for the target in question.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   mapsTo: target
//fromDate ==>
//   description: Filters out jobs older than this date.
//   format: date
//   type: string
//   mapsTo: updateTime
//   inequalityType: NON-STRICT
//toDate ==>
//   description: Filters out jobs newer than this date.
//   format: date
//   type: string
//   mapsTo: startTime
//   inequalityType: NON-STRICT
//searchText ==>
//   description: Limit search results to only include jobs that contain the searchText.
//   type: string
//maxTotal ==>
//   description: The upper bound for calculation of total job count.
//   type: integer
//return []*JobStruct
func (c *Client) JobList(jobState string,pageSize *int,pageOffset *int,addEvents *bool,jobType string,target string,fromDate string,toDate string,searchText string,maxTotal *int) ([]*JobStruct, error) {
var err error
var result []*JobStruct
return result, err
}
//JobUpdate - Update the specified Job object.
//
//r - Job Object Reference
//payload *JobStruct object
func (c *Client) JobUpdate(r string,payload *JobStruct) ( error) {
var err error
return err
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
//KerberosConfigUpdate - Update the specified KerberosConfig object.
//
//r - KerberosConfig Object Reference
//payload *KerberosConfigStruct object
func (c *Client) KerberosConfigUpdate(r string,payload *KerberosConfigStruct) ( error) {
var err error
return err
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
//LdapServerUpdate - Update the specified LdapServer object.
//
//r - LdapServer Object Reference
//payload *LdapServerStruct object
func (c *Client) LdapServerUpdate(r string,payload *LdapServerStruct) ( error) {
var err error
return err
}
//LdapServerList - List LdapServer objects on the system.
//
//return []*LdapServerStruct
func (c *Client) LdapServerList() ([]*LdapServerStruct, error) {
var err error
var result []*LdapServerStruct
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
//LdapServerDelete - Delete the specified LdapServer object.
//
//r - LdapServer Object Reference
func (c *Client) LdapServerDelete(r string) ( error) {
var err error
return err
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
//LocaleSettingsUpdate - Update the specified LocaleSettings object.
//
//r - LocaleSettings Object Reference
//payload *LocaleSettingsStruct object
func (c *Client) LocaleSettingsUpdate(r string,payload *LocaleSettingsStruct) ( error) {
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
//   type: string
//   format: objectReference
//   referenceTo: /delphix-mssql-availability-group.json
//   description: A reference to the SQL Server Availability Group this listener belongs to.
//return []*MSSqlAvailabilityGroupListenerStruct
func (c *Client) MSSqlAvailabilityGroupListenerList(availabilitygroup string) ([]*MSSqlAvailabilityGroupListenerStruct, error) {
var err error
var result []*MSSqlAvailabilityGroupListenerStruct
return result, err
}
//MSSqlClusterInstanceList - Returns a list of instances filtered by SQL Server Availability
// Group.
//
//availabilitygroup ==>
//   format: objectReference
//   referenceTo: /delphix-mssql-availability-group.json
//   description: A reference to the SQL Server Availability Group this instance belongs to.
//   type: string
//return []*MSSqlClusterInstanceStruct
func (c *Client) MSSqlClusterInstanceList(availabilitygroup string) ([]*MSSqlClusterInstanceStruct, error) {
var err error
var result []*MSSqlClusterInstanceStruct
return result, err
}
//MSSqlFailoverClusterInstanceList - Returns a list of instances filtered by SQL Server Failover
// Cluster.
//
//repository ==>
//   referenceTo: /delphix-mssql-failover-cluster-repository.json
//   description: A reference to the SQL Server Failover Cluster repository this instance belongs to.
//   type: string
//   format: objectReference
//return []*MSSqlFailoverClusterInstanceStruct
func (c *Client) MSSqlFailoverClusterInstanceList(repository string) ([]*MSSqlFailoverClusterInstanceStruct, error) {
var err error
var result []*MSSqlFailoverClusterInstanceStruct
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
return result, err
}
//MaskingJobDelete - Delete the specified MaskingJob object.
//
//r - MaskingJob Object Reference
func (c *Client) MaskingJobDelete(r string) ( error) {
var err error
return err
}
//MaskingJobUpdate - Update the specified MaskingJob object.
//
//r - MaskingJob Object Reference
//payload *MaskingJobStruct object
func (c *Client) MaskingJobUpdate(r string,payload *MaskingJobStruct) ( error) {
var err error
return err
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
//   description: List only the Masking Jobs that are associated with the provided container.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   required: false
//return []*MaskingJobStruct
func (c *Client) MaskingJobList(container string) ([]*MaskingJobStruct, error) {
var err error
var result []*MaskingJobStruct
return result, err
}
//MaskingServiceConfigList - Returns a list of all Masking Jobs on the system.
//
//return []*MaskingServiceConfigStruct
func (c *Client) MaskingServiceConfigList() ([]*MaskingServiceConfigStruct, error) {
var err error
var result []*MaskingServiceConfigStruct
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
//MaskingServiceConfigUpdate - Update the specified MaskingServiceConfig object.
//
//r - MaskingServiceConfig Object Reference
//payload *MaskingServiceConfigStruct object
func (c *Client) MaskingServiceConfigUpdate(r string,payload *MaskingServiceConfigStruct) ( error) {
var err error
return err
}
//MinimalPhoneHomeUpdate - Update the specified MinimalPhoneHome object.
//
//r - MinimalPhoneHome Object Reference
//payload *MinimalPhoneHomeStruct object
func (c *Client) MinimalPhoneHomeUpdate(r string,payload *MinimalPhoneHomeStruct) ( error) {
var err error
return err
}
//MinimalPhoneHomeRead - Retrieve the specified MinimalPhoneHome object.
//
//r - MinimalPhoneHome Object Reference
//return *MinimalPhoneHomeStruct
func (c *Client) MinimalPhoneHomeRead(r string) (*MinimalPhoneHomeStruct, error) {
var err error
var result *MinimalPhoneHomeStruct
obj, err := c.FindObjectByReference("/service/minimalPhonehome", r)
result = obj.(*MinimalPhoneHomeStruct)
return result, err
}
//NamespaceList - List Namespace objects on the system.
//
//return []*NamespaceStruct
func (c *Client) NamespaceList() ([]*NamespaceStruct, error) {
var err error
var result []*NamespaceStruct
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
func (c *Client) NamespaceUpdate(r string,payload *NamespaceStruct) ( error) {
var err error
return err
}
//NamespaceDelete - Delete the specified Namespace object.
//
//r - Namespace Object Reference
func (c *Client) NamespaceDelete(r string) ( error) {
var err error
return err
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
//NetworkDSPTestList - Returns the list of previously executed tests.
//
//ascending ==>
//   description: True if results are to be returned in ascending order.
//   type: boolean
//return []*NetworkDSPTestStruct
func (c *Client) NetworkDSPTestList(ascending *bool) ([]*NetworkDSPTestStruct, error) {
var err error
var result []*NetworkDSPTestStruct
return result, err
}
//NetworkDSPTestDelete - Delete the specified NetworkDSPTest object.
//
//r - NetworkDSPTest Object Reference
func (c *Client) NetworkDSPTestDelete(r string) ( error) {
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
//NetworkInterfaceUpdate - Update the specified NetworkInterface object.
//
//r - NetworkInterface Object Reference
//payload *NetworkInterfaceStruct object
func (c *Client) NetworkInterfaceUpdate(r string,payload *NetworkInterfaceStruct) ( error) {
var err error
return err
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
func (c *Client) NetworkLatencyTestDelete(r string) ( error) {
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
//NetworkLatencyTestList - Returns the list of previously executed tests.
//
//ascending ==>
//   type: boolean
//   description: True if results are to be returned in ascending order.
//return []*NetworkLatencyTestStruct
func (c *Client) NetworkLatencyTestList(ascending *bool) ([]*NetworkLatencyTestStruct, error) {
var err error
var result []*NetworkLatencyTestStruct
return result, err
}
//NetworkRouteList - Lists entries in the routing table.
//
//return []*NetworkRouteStruct
func (c *Client) NetworkRouteList() ([]*NetworkRouteStruct, error) {
var err error
var result []*NetworkRouteStruct
return result, err
}
//NetworkThroughputTestDelete - Delete the specified NetworkThroughputTest object.
//
//r - NetworkThroughputTest Object Reference
func (c *Client) NetworkThroughputTestDelete(r string) ( error) {
var err error
return err
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
//ascending ==>
//   description: True if results are to be returned in ascending order.
//   type: boolean
//return []*NetworkThroughputTestStruct
func (c *Client) NetworkThroughputTestList(ascending *bool) ([]*NetworkThroughputTestStruct, error) {
var err error
var result []*NetworkThroughputTestStruct
return result, err
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
//NfsConfigRead - Retrieve the specified NfsConfig object.
//
//r - NfsConfig Object Reference
//return *NfsConfigStruct
func (c *Client) NfsConfigRead(r string) (*NfsConfigStruct, error) {
var err error
var result *NfsConfigStruct
obj, err := c.FindObjectByReference("/service/nfs", r)
result = obj.(*NfsConfigStruct)
return result, err
}
//NfsConfigUpdate - Update the specified NfsConfig object.
//
//r - NfsConfig Object Reference
//payload *NfsConfigStruct object
func (c *Client) NfsConfigUpdate(r string,payload *NfsConfigStruct) ( error) {
var err error
return err
}
//NotificationList - Returns a list of pending notifications for the current session.
//
//channel ==>
//   type: string
//   description: Client-specific ID to specify an independent channel.
//timeout ==>
//   type: string
//   description: Timeout, in milliseconds, to wait for one or more responses.
//max ==>
//   type: string
//   description: Maximum number of entries to return at once.
//return []Notification
func (c *Client) NotificationList(channel string,timeout string,max string) ([]Notification, error) {
var err error
var result []Notification
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
//OperationTemplateUpdate - Update the specified OperationTemplate object.
//
//r - OperationTemplate Object Reference
//payload *OperationTemplateStruct object
func (c *Client) OperationTemplateUpdate(r string,payload *OperationTemplateStruct) ( error) {
var err error
return err
}
//OperationTemplateList - Lists templates available in the Delphix Engine.
//
//return []*OperationTemplateStruct
func (c *Client) OperationTemplateList() ([]*OperationTemplateStruct, error) {
var err error
var result []*OperationTemplateStruct
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
//OperationTemplateDelete - Delete the specified OperationTemplate object.
//
//r - OperationTemplate Object Reference
func (c *Client) OperationTemplateDelete(r string) ( error) {
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
//OracleClusterNodeList - Returns a list of host nodes filtered by cluster.
//
//cluster ==>
//   referenceTo: /delphix-oracle-cluster.json
//   mapsTo: cluster
//   type: string
//   format: objectReference
//   description: The cluster to filter by.
//return []*OracleClusterNodeStruct
func (c *Client) OracleClusterNodeList(cluster string) ([]*OracleClusterNodeStruct, error) {
var err error
var result []*OracleClusterNodeStruct
return result, err
}
//OracleClusterNodeUpdate - Update the specified OracleClusterNode object.
//
//r - OracleClusterNode Object Reference
//payload *OracleClusterNodeStruct object
func (c *Client) OracleClusterNodeUpdate(r string,payload *OracleClusterNodeStruct) ( error) {
var err error
return err
}
//OracleClusterNodeDelete - Delete the specified OracleClusterNode object.
//
//r - OracleClusterNode Object Reference
func (c *Client) OracleClusterNodeDelete(r string) ( error) {
var err error
return err
}
//OracleListenerDelete - Delete the specified OracleListener object.
//
//r - OracleListener Object Reference
func (c *Client) OracleListenerDelete(r string) ( error) {
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
func (c *Client) OracleListenerList(typeOracleListener string,environment string) ([]OracleListener, error) {
var err error
var result []OracleListener
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
//OracleListenerUpdate - Update the specified OracleListener object.
//
//r - OracleListener Object Reference
//payload OracleListener object
func (c *Client) OracleListenerUpdate(r string,payload OracleListener) ( error) {
var err error
return err
}
//OracleTimeflowLogList - Returns a list of fetched or missing Oracle logs for a database,
// TimeFlow or snapshot. The logs are returned in ascending order by
// TimeFlow, SCN.
//
//toScn ==>
//   description: Return logs with SCNs less than or equal to this value.
//   type: string
//pageSize ==>
//   description: Limit the number of logs returned.
//   type: integer
//   default: 25
//   minimum: 0
//snapshot ==>
//   description: Return logs for the specified snapshot up to the next snapshot. This option is mutually exclusive with the "TimeFlow" and "database" options.
//   format: objectReference
//   referenceTo: /delphix-oracle-snapshot.json
//   type: string
//fromDate ==>
//   description: Return logs created after this date.
//   format: date
//   type: string
//toDate ==>
//   type: string
//   description: Return logs created before than this date.
//   format: date
//fromScn ==>
//   description: Return logs with SCNs greater than or equal to this value.
//   type: string
//pageOffset ==>
//   description: Page offset within log list, in units of pageSize chunks.
//   type: integer
//missing ==>
//   description: Only return the missing logs.
//   type: boolean
//database ==>
//   type: string
//   description: Return logs on all TimeFlows associated with the container. This option is mutually exclusive with the "TimeFlow" and "snapshot" options.
//   format: objectReference
//   referenceTo: /delphix-oracle-db-container.json
//timeflow ==>
//   mapsTo: timeflow
//   type: string
//   description: Return logs in the specified TimeFlow. This option is mutually exclusive with the "snapshot" and "database" options.
//   format: objectReference
//   referenceTo: /delphix-oracle-timeflow.json
//return []OracleTimeflowLog
func (c *Client) OracleTimeflowLogList(toScn string,pageSize *int,snapshot string,fromDate string,toDate string,fromScn string,pageOffset *int,missing *bool,database string,timeflow string) ([]OracleTimeflowLog, error) {
var err error
var result []OracleTimeflowLog
return result, err
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
//PasswordPolicyUpdate - Update the specified PasswordPolicy object.
//
//r - PasswordPolicy Object Reference
//payload *PasswordPolicyStruct object
func (c *Client) PasswordPolicyUpdate(r string,payload *PasswordPolicyStruct) ( error) {
var err error
return err
}
//PasswordPolicyDelete - Delete the specified PasswordPolicy object.
//
//r - PasswordPolicy Object Reference
func (c *Client) PasswordPolicyDelete(r string) ( error) {
var err error
return err
}
//PasswordPolicyList - Lists password policies in the system.
//
//return []*PasswordPolicyStruct
func (c *Client) PasswordPolicyList() ([]*PasswordPolicyStruct, error) {
var err error
var result []*PasswordPolicyStruct
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
//PasswordVaultUpdate - Update the specified PasswordVault object.
//
//r - PasswordVault Object Reference
//payload PasswordVault object
func (c *Client) PasswordVaultUpdate(r string,payload PasswordVault) ( error) {
var err error
return err
}
//PasswordVaultDelete - Delete the specified PasswordVault object.
//
//r - PasswordVault Object Reference
func (c *Client) PasswordVaultDelete(r string) ( error) {
var err error
return err
}
//PasswordVaultList - List PasswordVault objects on the system.
//
//return []PasswordVault
func (c *Client) PasswordVaultList() ([]PasswordVault, error) {
var err error
var result []PasswordVault
return result, err
}
//PasswordVaultCreate - Create a new PasswordVault object.
//
//payload PasswordVault object
//return string
func (c *Client) PasswordVaultCreate(payload PasswordVault) (string, error) {
var err error
var result string
return result, err
}
//PasswordVaultRead - Retrieve the specified PasswordVault object.
//
//r - PasswordVault Object Reference
//return PasswordVault
func (c *Client) PasswordVaultRead(r string) (PasswordVault, error) {
var err error
var result PasswordVault
obj, err := c.FindObjectByReference("/service/passwordVault", r)
result = obj.(PasswordVault)
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
//PermissionList - Lists permissions available in the system.
//
//return []*PermissionStruct
func (c *Client) PermissionList() ([]*PermissionStruct, error) {
var err error
var result []*PermissionStruct
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
func (c *Client) PhoneHomeServiceUpdate(r string,payload *PhoneHomeServiceStruct) ( error) {
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
func (c *Client) PolicyUpdate(r string,payload Policy) ( error) {
var err error
return err
}
//PolicyDelete - Delete the specified Policy object.
//
//r - Policy Object Reference
func (c *Client) PolicyDelete(r string) ( error) {
var err error
return err
}
//PolicyList - Returns a list of policies in the domain.
//
//target ==>
//   format: objectReference
//   referenceTo: /delphix-user-object.json
//   description: Limit policies to those affecting a particular object on the system.
//   type: string
//effective ==>
//   description: Whether to include effective policies for the target.
//   type: string
//typePolicy ==>
//   description: Limit policies to those of the given type.
//   type: string
//return []Policy
func (c *Client) PolicyList(target string,effective string,typePolicy string) ([]Policy, error) {
var err error
var result []Policy
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
//ProxyServiceUpdate - Update the specified ProxyService object.
//
//r - ProxyService Object Reference
//payload *ProxyServiceStruct object
func (c *Client) ProxyServiceUpdate(r string,payload *ProxyServiceStruct) ( error) {
var err error
return err
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
func (c *Client) RegistrationStatusUpdate(r string,payload *RegistrationStatusStruct) ( error) {
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
//ReplicationSpecList - List ReplicationSpec objects on the system.
//
//return []*ReplicationSpecStruct
func (c *Client) ReplicationSpecList() ([]*ReplicationSpecStruct, error) {
var err error
var result []*ReplicationSpecStruct
return result, err
}
//ReplicationSpecUpdate - Update the specified ReplicationSpec object.
//
//r - ReplicationSpec Object Reference
//payload *ReplicationSpecStruct object
func (c *Client) ReplicationSpecUpdate(r string,payload *ReplicationSpecStruct) ( error) {
var err error
return err
}
//ReplicationSpecDelete - Delete the specified ReplicationSpec object.
//
//r - ReplicationSpec Object Reference
func (c *Client) ReplicationSpecDelete(r string) ( error) {
var err error
return err
}
//ReplicationTargetStateList - List ReplicationTargetState objects on the system.
//
//return []*ReplicationTargetStateStruct
func (c *Client) ReplicationTargetStateList() ([]*ReplicationTargetStateStruct, error) {
var err error
var result []*ReplicationTargetStateStruct
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
//RoleUpdate - Update the specified Role object.
//
//r - Role Object Reference
//payload *RoleStruct object
func (c *Client) RoleUpdate(r string,payload *RoleStruct) ( error) {
var err error
return err
}
//RoleDelete - Delete the specified Role object.
//
//r - Role Object Reference
func (c *Client) RoleDelete(r string) ( error) {
var err error
return err
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
//RoleList - Lists roles available in the system.
//
//return []*RoleStruct
func (c *Client) RoleList() ([]*RoleStruct, error) {
var err error
var result []*RoleStruct
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
func (c *Client) SMTPConfigUpdate(r string,payload *SMTPConfigStruct) ( error) {
var err error
return err
}
//SNMPV2ConfigUpdate - Update the specified SNMPV2Config object.
//
//r - SNMPV2Config Object Reference
//payload *SNMPV2ConfigStruct object
func (c *Client) SNMPV2ConfigUpdate(r string,payload *SNMPV2ConfigStruct) ( error) {
var err error
return err
}
//SNMPV2ConfigRead - Retrieve the specified SNMPV2Config object.
//
//r - SNMPV2Config Object Reference
//return *SNMPV2ConfigStruct
func (c *Client) SNMPV2ConfigRead(r string) (*SNMPV2ConfigStruct, error) {
var err error
var result *SNMPV2ConfigStruct
obj, err := c.FindObjectByReference("/service/snmp/v2", r)
result = obj.(*SNMPV2ConfigStruct)
return result, err
}
//SNMPV2ManagerRead - Retrieve the specified SNMPV2Manager object.
//
//r - SNMPV2Manager Object Reference
//return *SNMPV2ManagerStruct
func (c *Client) SNMPV2ManagerRead(r string) (*SNMPV2ManagerStruct, error) {
var err error
var result *SNMPV2ManagerStruct
obj, err := c.FindObjectByReference("/service/snmp/v2/manager", r)
result = obj.(*SNMPV2ManagerStruct)
return result, err
}
//SNMPV2ManagerList - Lists SNMP managers in the system.
//
//return []*SNMPV2ManagerStruct
func (c *Client) SNMPV2ManagerList() ([]*SNMPV2ManagerStruct, error) {
var err error
var result []*SNMPV2ManagerStruct
return result, err
}
//SNMPV2ManagerCreate - Create a new SNMPV2Manager object.
//
//payload *SNMPV2ManagerStruct object
//return string
func (c *Client) SNMPV2ManagerCreate(payload *SNMPV2ManagerStruct) (string, error) {
var err error
var result string
return result, err
}
//SNMPV2ManagerDelete - Delete the specified SNMPV2Manager object.
//
//r - SNMPV2Manager Object Reference
func (c *Client) SNMPV2ManagerDelete(r string) ( error) {
var err error
return err
}
//SNMPV2ManagerUpdate - Update the specified SNMPV2Manager object.
//
//r - SNMPV2Manager Object Reference
//payload *SNMPV2ManagerStruct object
func (c *Client) SNMPV2ManagerUpdate(r string,payload *SNMPV2ManagerStruct) ( error) {
var err error
return err
}
//SNMPV3ConfigUpdate - Update the specified SNMPV3Config object.
//
//r - SNMPV3Config Object Reference
//payload *SNMPV3ConfigStruct object
func (c *Client) SNMPV3ConfigUpdate(r string,payload *SNMPV3ConfigStruct) ( error) {
var err error
return err
}
//SNMPV3ConfigRead - Retrieve the specified SNMPV3Config object.
//
//r - SNMPV3Config Object Reference
//return *SNMPV3ConfigStruct
func (c *Client) SNMPV3ConfigRead(r string) (*SNMPV3ConfigStruct, error) {
var err error
var result *SNMPV3ConfigStruct
obj, err := c.FindObjectByReference("/service/snmp/v3", r)
result = obj.(*SNMPV3ConfigStruct)
return result, err
}
//SNMPV3ManagerUpdate - Update the specified SNMPV3Manager object.
//
//r - SNMPV3Manager Object Reference
//payload *SNMPV3ManagerStruct object
func (c *Client) SNMPV3ManagerUpdate(r string,payload *SNMPV3ManagerStruct) ( error) {
var err error
return err
}
//SNMPV3ManagerRead - Retrieve the specified SNMPV3Manager object.
//
//r - SNMPV3Manager Object Reference
//return *SNMPV3ManagerStruct
func (c *Client) SNMPV3ManagerRead(r string) (*SNMPV3ManagerStruct, error) {
var err error
var result *SNMPV3ManagerStruct
obj, err := c.FindObjectByReference("/service/snmp/v3/manager", r)
result = obj.(*SNMPV3ManagerStruct)
return result, err
}
//SNMPV3ManagerCreate - Create a new SNMPV3Manager object.
//
//payload *SNMPV3ManagerStruct object
//return string
func (c *Client) SNMPV3ManagerCreate(payload *SNMPV3ManagerStruct) (string, error) {
var err error
var result string
return result, err
}
//SNMPV3ManagerList - Lists SNMP managers in the system.
//
//return []*SNMPV3ManagerStruct
func (c *Client) SNMPV3ManagerList() ([]*SNMPV3ManagerStruct, error) {
var err error
var result []*SNMPV3ManagerStruct
return result, err
}
//SNMPV3ManagerDelete - Delete the specified SNMPV3Manager object.
//
//r - SNMPV3Manager Object Reference
func (c *Client) SNMPV3ManagerDelete(r string) ( error) {
var err error
return err
}
//SNMPV3USMUserConfigRead - Retrieve the specified SNMPV3USMUserConfig object.
//
//r - SNMPV3USMUserConfig Object Reference
//return *SNMPV3USMUserConfigStruct
func (c *Client) SNMPV3USMUserConfigRead(r string) (*SNMPV3USMUserConfigStruct, error) {
var err error
var result *SNMPV3USMUserConfigStruct
obj, err := c.FindObjectByReference("/service/snmp/v3/usm", r)
result = obj.(*SNMPV3USMUserConfigStruct)
return result, err
}
//SNMPV3USMUserConfigUpdate - Update the specified SNMPV3USMUserConfig object.
//
//r - SNMPV3USMUserConfig Object Reference
//payload *SNMPV3USMUserConfigStruct object
func (c *Client) SNMPV3USMUserConfigUpdate(r string,payload *SNMPV3USMUserConfigStruct) ( error) {
var err error
return err
}
//SNMPV3USMUserConfigList - Lists SNMP User Security Model users.
//
//return []*SNMPV3USMUserConfigStruct
func (c *Client) SNMPV3USMUserConfigList() ([]*SNMPV3USMUserConfigStruct, error) {
var err error
var result []*SNMPV3USMUserConfigStruct
return result, err
}
//SNMPV3USMUserConfigCreate - Create a new SNMPV3USMUserConfig object.
//
//payload *SNMPV3USMUserConfigStruct object
//return string
func (c *Client) SNMPV3USMUserConfigCreate(payload *SNMPV3USMUserConfigStruct) (string, error) {
var err error
var result string
return result, err
}
//SNMPV3USMUserConfigDelete - Delete the specified SNMPV3USMUserConfig object.
//
//r - SNMPV3USMUserConfig Object Reference
func (c *Client) SNMPV3USMUserConfigDelete(r string) ( error) {
var err error
return err
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
func (c *Client) SecurityConfigUpdate(r string,payload *SecurityConfigStruct) ( error) {
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
return result, err
}
//SnapshotCapacityDataList - Lists capacity metrics for all snapshots in the system sorted by
// snapshot space usage decreasing.
//
//container ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   description: The container to list snapshot data for.
//   mapsTo: container
//namespace ==>
//   mapsTo: namespace
//   type: [string null]
//   format: objectReference
//   description: The namespace to list snapshot data for. If null, will limit returned snapshots to the default namespace.
//pageSize ==>
//   type: integer
//   minimum: 0
//   description: Limit the number of entries returned.
//pageOffset ==>
//   description: Offset within list, in units of pageSize chunks.
//   type: integer
//   minimum: 0
//return []*SnapshotCapacityDataStruct
func (c *Client) SnapshotCapacityDataList(container string,namespace string,pageSize *int,pageOffset *int) ([]*SnapshotCapacityDataStruct, error) {
var err error
var result []*SnapshotCapacityDataStruct
return result, err
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
//SourceUpdate - Update the specified Source object.
//
//r - Source Object Reference
//payload Source object
func (c *Client) SourceUpdate(r string,payload Source) ( error) {
var err error
return err
}
//SourceList - Lists sources on the system.
//
//repository ==>
//   referenceTo: /delphix-source-repository.json
//   mapsTo: config.repository.reference
//   description: List sources associated with the given source repository reference.
//   type: string
//   format: objectReference
//environment ==>
//   mapsTo: config.repository.environment.reference
//   description: List sources associated with the given source environment reference.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-source-environment.json
//includeHosts ==>
//   type: boolean
//   description: Whether to include the list of hosts for each source in the response.
//database ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//   description: List visible sources associated with the given container reference. Visible sources are of type LINKED or VIRTUAL.
//config ==>
//   description: List visible sources associated with the given sourceconfig reference. Visible sources are of type LINKED or VIRTUAL.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-source-config.json
//   mapsTo: config
//allSources ==>
//   description: List all sources associated with the given source container reference.
//   type: boolean
//return []Source
func (c *Client) SourceList(repository string,environment string,includeHosts *bool,database string,config string,allSources *bool) ([]Source, error) {
var err error
var result []Source
return result, err
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
//   referenceTo: /delphix-source-environment.json
//   mapsTo: repository.environment.reference
//   type: string
//   description: Restrict source configs to those belonging to the specified environment. This option is mutually exclusive with all other options.
//   format: objectReference
//cdbConfig ==>
//   type: string
//   description: Restrict PDB configs to those belonging to the specified CDB source config.
//   format: objectReference
//   referenceTo: /delphix-oracle-db-config.json
//pdbConfigOnly ==>
//   type: boolean
//   description: Restrict source configs to be Oracle PDB configs only.
//return []SourceConfig
func (c *Client) SourceConfigList(repository string,environment string,cdbConfig string,pdbConfigOnly *bool) ([]SourceConfig, error) {
var err error
var result []SourceConfig
return result, err
}
//SourceConfigDelete - Delete the specified SourceConfig object.
//
//r - SourceConfig Object Reference
func (c *Client) SourceConfigDelete(r string) ( error) {
var err error
return err
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
func (c *Client) SourceConfigUpdate(r string,payload SourceConfig) ( error) {
var err error
return err
}
//SourceEnvironmentDelete - Delete the specified SourceEnvironment object.
//
//r - SourceEnvironment Object Reference
func (c *Client) SourceEnvironmentDelete(r string) ( error) {
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
return result, err
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
func (c *Client) SourceEnvironmentUpdate(r string,payload SourceEnvironment) ( error) {
var err error
return err
}
//SourceRepositoryDelete - Delete the specified SourceRepository object.
//
//r - SourceRepository Object Reference
func (c *Client) SourceRepositoryDelete(r string) ( error) {
var err error
return err
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
return result, err
}
//SourceRepositoryUpdate - Update the specified SourceRepository object.
//
//r - SourceRepository Object Reference
//payload SourceRepository object
func (c *Client) SourceRepositoryUpdate(r string,payload SourceRepository) ( error) {
var err error
return err
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
//SourceRepositoryTemplateList - Returns the list of repository templates.
//
//return []*SourceRepositoryTemplateStruct
func (c *Client) SourceRepositoryTemplateList() ([]*SourceRepositoryTemplateStruct, error) {
var err error
var result []*SourceRepositoryTemplateStruct
return result, err
}
//SourceRepositoryTemplateDelete - Delete the specified SourceRepositoryTemplate object.
//
//r - SourceRepositoryTemplate Object Reference
func (c *Client) SourceRepositoryTemplateDelete(r string) ( error) {
var err error
return err
}
//SourceRepositoryTemplateUpdate - Update the specified SourceRepositoryTemplate object.
//
//r - SourceRepositoryTemplate Object Reference
//payload *SourceRepositoryTemplateStruct object
func (c *Client) SourceRepositoryTemplateUpdate(r string,payload *SourceRepositoryTemplateStruct) ( error) {
var err error
return err
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
//SourceRepositoryTemplateCreate - Create a new SourceRepositoryTemplate object.
//
//payload *SourceRepositoryTemplateStruct object
//return string
func (c *Client) SourceRepositoryTemplateCreate(payload *SourceRepositoryTemplateStruct) (string, error) {
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
return result, err
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
//SplunkHecConfigDelete - Delete the specified SplunkHecConfig object.
//
//r - SplunkHecConfig Object Reference
func (c *Client) SplunkHecConfigDelete(r string) ( error) {
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
//SplunkHecConfigUpdate - Update the specified SplunkHecConfig object.
//
//r - SplunkHecConfig Object Reference
//payload *SplunkHecConfigStruct object
func (c *Client) SplunkHecConfigUpdate(r string,payload *SplunkHecConfigStruct) ( error) {
var err error
return err
}
//SsoConfigRead - Retrieve the specified SsoConfig object.
//
//r - SsoConfig Object Reference
//return *SsoConfigStruct
func (c *Client) SsoConfigRead(r string) (*SsoConfigStruct, error) {
var err error
var result *SsoConfigStruct
obj, err := c.FindObjectByReference("/service/sso", r)
result = obj.(*SsoConfigStruct)
return result, err
}
//SsoConfigUpdate - Update the specified SsoConfig object.
//
//r - SsoConfig Object Reference
//payload *SsoConfigStruct object
func (c *Client) SsoConfigUpdate(r string,payload *SsoConfigStruct) ( error) {
var err error
return err
}
//StaticHostAddressCreate - Create a new StaticHostAddress object.
//
//payload *StaticHostAddressStruct object
//return string
func (c *Client) StaticHostAddressCreate(payload *StaticHostAddressStruct) (string, error) {
var err error
var result string
return result, err
}
//StaticHostAddressUpdate - Update the specified StaticHostAddress object.
//
//r - StaticHostAddress Object Reference
//payload *StaticHostAddressStruct object
func (c *Client) StaticHostAddressUpdate(r string,payload *StaticHostAddressStruct) ( error) {
var err error
return err
}
//StaticHostAddressRead - Retrieve the specified StaticHostAddress object.
//
//r - StaticHostAddress Object Reference
//return *StaticHostAddressStruct
func (c *Client) StaticHostAddressRead(r string) (*StaticHostAddressStruct, error) {
var err error
var result *StaticHostAddressStruct
obj, err := c.FindObjectByReference("/service/host/address", r)
result = obj.(*StaticHostAddressStruct)
return result, err
}
//StaticHostAddressList - List StaticHostAddress objects on the system.
//
//return []*StaticHostAddressStruct
func (c *Client) StaticHostAddressList() ([]*StaticHostAddressStruct, error) {
var err error
var result []*StaticHostAddressStruct
return result, err
}
//StaticHostAddressDelete - Delete the specified StaticHostAddress object.
//
//r - StaticHostAddress Object Reference
func (c *Client) StaticHostAddressDelete(r string) ( error) {
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
//StatisticSliceDelete - Delete the specified StatisticSlice object.
//
//r - StatisticSlice Object Reference
func (c *Client) StatisticSliceDelete(r string) ( error) {
var err error
return err
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
//StatisticSliceList - Returns a list of statistics in the system.
//
//return []*StatisticSliceStruct
func (c *Client) StatisticSliceList() ([]*StatisticSliceStruct, error) {
var err error
var result []*StatisticSliceStruct
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
//StorageTestCreate - Create a new StorageTest object.
//
//payload *StorageTestParametersStruct object
//return string
func (c *Client) StorageTestCreate(payload *StorageTestParametersStruct) (string, error) {
var err error
var result string
return result, err
}
//StorageTestDelete - Delete the specified StorageTest object.
//
//r - StorageTest Object Reference
func (c *Client) StorageTestDelete(r string) ( error) {
var err error
return err
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
//StorageTestList - Returns the list of previously executed tests.
//
//return []*StorageTestStruct
func (c *Client) StorageTestList() ([]*StorageTestStruct, error) {
var err error
var result []*StorageTestStruct
return result, err
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
func (c *Client) SupportAccessStateUpdate(r string,payload *SupportAccessStateStruct) ( error) {
var err error
return err
}
//SupportBundleConfigurationRead - Retrieve the specified SupportBundleConfiguration object.
//
//r - SupportBundleConfiguration Object Reference
//return *SupportBundleConfigurationStruct
func (c *Client) SupportBundleConfigurationRead(r string) (*SupportBundleConfigurationStruct, error) {
var err error
var result *SupportBundleConfigurationStruct
obj, err := c.FindObjectByReference("/service/support/bundle", r)
result = obj.(*SupportBundleConfigurationStruct)
return result, err
}
//SupportBundleConfigurationUpdate - Update the specified SupportBundleConfiguration object.
//
//r - SupportBundleConfiguration Object Reference
//payload *SupportBundleConfigurationStruct object
func (c *Client) SupportBundleConfigurationUpdate(r string,payload *SupportBundleConfigurationStruct) ( error) {
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
func (c *Client) SyslogConfigUpdate(r string,payload *SyslogConfigStruct) ( error) {
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
//SystemInfoUpdate - Update the specified SystemInfo object.
//
//r - SystemInfo Object Reference
//payload *SystemInfoStruct object
func (c *Client) SystemInfoUpdate(r string,payload *SystemInfoStruct) ( error) {
var err error
return err
}
//SystemPackageList - List the packages that can be changed via the web services.
//
//return []*SystemPackageStruct
func (c *Client) SystemPackageList() ([]*SystemPackageStruct, error) {
var err error
var result []*SystemPackageStruct
return result, err
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
//SystemPackageUpdate - Update the specified SystemPackage object.
//
//r - SystemPackage Object Reference
//payload *SystemPackageStruct object
func (c *Client) SystemPackageUpdate(r string,payload *SystemPackageStruct) ( error) {
var err error
return err
}
//SystemStatusList - List SystemStatus objects on the system.
//
//return []*SystemStatusStruct
func (c *Client) SystemStatusList() ([]*SystemStatusStruct, error) {
var err error
var result []*SystemStatusStruct
return result, err
}
//SystemStatusRead - Retrieve the specified SystemStatus object.
//
//r - SystemStatus Object Reference
//return *SystemStatusStruct
func (c *Client) SystemStatusRead(r string) (*SystemStatusStruct, error) {
var err error
var result *SystemStatusStruct
obj, err := c.FindObjectByReference("/system/status", r)
result = obj.(*SystemStatusStruct)
return result, err
}
//SystemVersionList - List SystemVersion objects on the system.
//
//return []*SystemVersionStruct
func (c *Client) SystemVersionList() ([]*SystemVersionStruct, error) {
var err error
var result []*SystemVersionStruct
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
func (c *Client) SystemVersionDelete(r string) ( error) {
var err error
return err
}
//ThemeList - Returns the list of themes.
//
//pageOffset ==>
//   description: Offset within result list, in units of pageSize chunks.
//   type: integer
//   default: 0
//active ==>
//   description: Decides whether or not to only return active theme.
//   type: boolean
//   default: false
//   mapsTo: active
//pageSize ==>
//   description: Limit the number of check results returned.
//   type: integer
//   default: 25
//   minimum: 0
//return []*ThemeStruct
func (c *Client) ThemeList(pageOffset *int,active *bool,pageSize *int) ([]*ThemeStruct, error) {
var err error
var result []*ThemeStruct
return result, err
}
//ThemeRead - Retrieve the specified Theme object.
//
//r - Theme Object Reference
//return *ThemeStruct
func (c *Client) ThemeRead(r string) (*ThemeStruct, error) {
var err error
var result *ThemeStruct
obj, err := c.FindObjectByReference("/theme", r)
result = obj.(*ThemeStruct)
return result, err
}
//ThemeUpdate - Update the specified Theme object.
//
//r - Theme Object Reference
//payload *ThemeStruct object
func (c *Client) ThemeUpdate(r string,payload *ThemeStruct) ( error) {
var err error
return err
}
//ThemeCreate - Create a new Theme object.
//
//payload *ThemeStruct object
//return string
func (c *Client) ThemeCreate(payload *ThemeStruct) (string, error) {
var err error
var result string
return result, err
}
//ThemeDelete - Delete the specified Theme object.
//
//r - Theme Object Reference
func (c *Client) ThemeDelete(r string) ( error) {
var err error
return err
}
//TimeConfigUpdate - Update the specified TimeConfig object.
//
//r - TimeConfig Object Reference
//payload *TimeConfigStruct object
func (c *Client) TimeConfigUpdate(r string,payload *TimeConfigStruct) ( error) {
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
return result, err
}
//TimeflowUpdate - Update the specified Timeflow object.
//
//r - Timeflow Object Reference
//payload Timeflow object
func (c *Client) TimeflowUpdate(r string,payload Timeflow) ( error) {
var err error
return err
}
//TimeflowList - List Timeflow objects on the system.
//
//database ==>
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//   type: string
//   description: List only TimeFlows within this database.
//return []Timeflow
func (c *Client) TimeflowList(database string) ([]Timeflow, error) {
var err error
var result []Timeflow
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
//TimeflowBookmarkCreate - Create a new TimeflowBookmark object.
//
//payload *TimeflowBookmarkCreateParametersStruct object
//return string
func (c *Client) TimeflowBookmarkCreate(payload *TimeflowBookmarkCreateParametersStruct) (string, error) {
var err error
var result string
return result, err
}
//TimeflowBookmarkList - Returns a list of all TimeFlow bookmarks.
//
//database ==>
//   description: Filter results based on specified database.
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: timeflow.container
//   type: string
//return []*TimeflowBookmarkStruct
func (c *Client) TimeflowBookmarkList(database string) ([]*TimeflowBookmarkStruct, error) {
var err error
var result []*TimeflowBookmarkStruct
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
func (c *Client) TimeflowBookmarkDelete(r string) ( error) {
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
//traverseTimeflows ==>
//   description: Whether to restrict snapshots to those in the current TimeFlow and in parent TimeFlows older than the branch point. This option is only used with the "database" option. The default behavior is false, i.e. show all snapshots.
//   type: boolean
//missingNonLoggedDataOnly ==>
//   description: Whether to restrict snapshots to those missing nologging changes. The defaultbehavior is salse.
//   type: boolean
//database ==>
//   description: Restrict snapshots to those within a TimeFlow of the specified database. This option is mutually exclusive with the "TimeFlow" option.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//timeflow ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-timeflow.json
//   mapsTo: timeflow
//   description: Restrict snapshots to those within the specified TimeFlow. This option is mutually exclusive with the "database" option.
//fromDate ==>
//   inequalityType: NON-STRICT
//   description: Start date to use for filtering out results.
//   format: date
//   type: string
//   mapsTo: latestChangePoint.timestamp
//pageSize ==>
//   description: Limit the number of snapshots returned.
//   type: integer
//   minimum: 0
//pageOffset ==>
//   description: Offset within TimeFlow snapshots, in units of pageSize chunks. The pageOffset query parameter is only supported when either a 'TimeFlow' or 'database' query parameter is also set.
//   type: integer
//toDate ==>
//   mapsTo: latestChangePoint.timestamp
//   inequalityType: STRICT
//   description: End date to use for filtering out results.
//   format: date
//   type: string
//return []TimeflowSnapshot
func (c *Client) TimeflowSnapshotList(traverseTimeflows *bool,missingNonLoggedDataOnly *bool,database string,timeflow string,fromDate string,pageSize *int,pageOffset *int,toDate string) ([]TimeflowSnapshot, error) {
var err error
var result []TimeflowSnapshot
return result, err
}
//TimeflowSnapshotDelete - Delete the specified TimeflowSnapshot object.
//
//r - TimeflowSnapshot Object Reference
func (c *Client) TimeflowSnapshotDelete(r string) ( error) {
var err error
return err
}
//TimeflowSnapshotUpdate - Update the specified TimeflowSnapshot object.
//
//r - TimeflowSnapshot Object Reference
//payload TimeflowSnapshot object
func (c *Client) TimeflowSnapshotUpdate(r string,payload TimeflowSnapshot) ( error) {
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
//   format: objectReference
//   referenceTo: /delphix-container.json
//   mapsTo: container
//   description: List the transformations that have been created against the provided container.
//   type: string
//return []*TransformationStruct
func (c *Client) TransformationList(container string,parentContainer string) ([]*TransformationStruct, error) {
var err error
var result []*TransformationStruct
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
func (c *Client) TransformationUpdate(r string,payload *TransformationStruct) ( error) {
var err error
return err
}
//UpgradeCheckResultList - Returns the list of all the check results that match the given
// criteria.
//
//sortBy ==>
//   enum: [SEVERITY TITLE VERSION STATUS]
//   description: Search results are sorted by the field provided.
//   type: string
//ascending ==>
//   type: boolean
//   description: True if results are to be returned in ascending order.
//version ==>
//   type: string
//   format: objectReference
//   referenceTo: /delphix-upgrade-version.json
//   mapsTo: version
//   description: The reference to the Delphix version associated with this verification check.
//status ==>
//   mapsTo: status
//   type: string
//   description: The status of the upgrade check result.
//   enum: [ACTIVE IGNORED RESOLVED]
//pageSize ==>
//   type: integer
//   default: 25
//   minimum: 0
//   description: Limit the number of check results returned.
//pageOffset ==>
//   description: Offset within result list, in units of pageSize chunks.
//   type: integer
//return []*UpgradeCheckResultStruct
func (c *Client) UpgradeCheckResultList(sortBy string,ascending *bool,version string,status string,pageSize *int,pageOffset *int) ([]*UpgradeCheckResultStruct, error) {
var err error
var result []*UpgradeCheckResultStruct
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
//UpgradeVerificationReportList - List UpgradeVerificationReport objects on the system.
//
//version ==>
//   description: The reference to the Delphix version associated with this verification report.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-upgrade-version.json
//   mapsTo: version
//pageSize ==>
//   description: Limit the number of report results returned.
//   type: integer
//   default: 25
//   minimum: 0
//pageOffset ==>
//   description: Offset within result list, in units of pageSize chunks.
//   type: integer
//sortBy ==>
//   description: Search results are sorted by the field provided.
//   type: string
//   enum: [VERSION VERIFICATION_VERSION]
//ascending ==>
//   description: True if results are to be returned in ascending order.
//   type: boolean
//return []*UpgradeVerificationReportStruct
func (c *Client) UpgradeVerificationReportList(version string,pageSize *int,pageOffset *int,sortBy string,ascending *bool) ([]*UpgradeVerificationReportStruct, error) {
var err error
var result []*UpgradeVerificationReportStruct
return result, err
}
//UpgradeVerificationReportRead - Retrieve the specified UpgradeVerificationReport object.
//
//r - UpgradeVerificationReport Object Reference
//return *UpgradeVerificationReportStruct
func (c *Client) UpgradeVerificationReportRead(r string) (*UpgradeVerificationReportStruct, error) {
var err error
var result *UpgradeVerificationReportStruct
obj, err := c.FindObjectByReference("/system/verification/reports", r)
result = obj.(*UpgradeVerificationReportStruct)
return result, err
}
//UpgradeVerificationStepsList - List UpgradeVerificationSteps objects on the system.
//
//pageSize ==>
//   description: Limit the number of step results returned.
//   type: integer
//   default: 25
//   minimum: 0
//pageOffset ==>
//   description: Offset within result list, in units of pageSize chunks.
//   type: integer
//sortBy ==>
//   description: Search results are sorted by the field provided.
//   type: string
//   enum: [RUN_STATUS CLASS_NAME START_TIMESTAMP DURATION]
//ascending ==>
//   description: True if results are to be returned in ascending order.
//   type: boolean
//version ==>
//   description: The reference to the Delphix version associated with this verification steps.
//   type: string
//   format: objectReference
//   referenceTo: /delphix-upgrade-version.json
//runStatus ==>
//   type: string
//   description: Result status of the step.
//   enum: [SUCCESS FAILURE SKIPPED]
//   mapsTo: runStatus
//return []*UpgradeVerificationStepsStruct
func (c *Client) UpgradeVerificationStepsList(pageSize *int,pageOffset *int,sortBy string,ascending *bool,version string,runStatus string) ([]*UpgradeVerificationStepsStruct, error) {
var err error
var result []*UpgradeVerificationStepsStruct
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
func (c *Client) UserUpdate(r string,payload *UserStruct) ( error) {
var err error
return err
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
//   description: DOMAIN user type: DOMAIN_ADMIN, STANDARD_USER, or SELFSERVICE_ONLY.
//   type: string
//   enum: [DOMAIN_ADMIN STANDARD_USER SELFSERVICE_ONLY]
//   required: false
//return []*UserStruct
func (c *Client) UserList(typeUser string,domainUserType string) ([]*UserStruct, error) {
var err error
var result []*UserStruct
return result, err
}
//UserDelete - Delete the specified User object.
//
//r - User Object Reference
func (c *Client) UserDelete(r string) ( error) {
var err error
return err
}
//UserInterfaceConfigUpdate - Update the specified UserInterfaceConfig object.
//
//r - UserInterfaceConfig Object Reference
//payload *UserInterfaceConfigStruct object
func (c *Client) UserInterfaceConfigUpdate(r string,payload *UserInterfaceConfigStruct) ( error) {
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
//UserPathStorageList - Returns the list of all the check results that match the given
// criteria.
//
//pathtype ==>
//   mapsTo: pathtype
//   description: Type of path stored.
//   type: string
//   enum: [UPGRADE_STAGING_LOC]
//return []*UserPathStorageStruct
func (c *Client) UserPathStorageList(pathtype string) ([]*UserPathStorageStruct, error) {
var err error
var result []*UserPathStorageStruct
return result, err
}
//UserPathStorageRead - Retrieve the specified UserPathStorage object.
//
//r - UserPathStorage Object Reference
//return *UserPathStorageStruct
func (c *Client) UserPathStorageRead(r string) (*UserPathStorageStruct, error) {
var err error
var result *UserPathStorageStruct
obj, err := c.FindObjectByReference("/service/userPaths", r)
result = obj.(*UserPathStorageStruct)
return result, err
}
//UserPathStorageUpdate - Update the specified UserPathStorage object.
//
//r - UserPathStorage Object Reference
//payload *UserPathStorageStruct object
func (c *Client) UserPathStorageUpdate(r string,payload *UserPathStorageStruct) ( error) {
var err error
return err
}
//UserPathStorageCreate - Create a new UserPathStorage object.
//
//payload *UserPathStorageStruct object
//return string
func (c *Client) UserPathStorageCreate(payload *UserPathStorageStruct) (string, error) {
var err error
var result string
return result, err
}
//UserPathStorageDelete - Delete the specified UserPathStorage object.
//
//r - UserPathStorage Object Reference
func (c *Client) UserPathStorageDelete(r string) ( error) {
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
//   type: string
//   format: objectReference
//   referenceTo: /delphix-windows-cluster.json
//   description: A reference to the Windows cluster environment this node belongs to.
//   mapsTo: cluster
//return []*WindowsClusterNodeStruct
func (c *Client) WindowsClusterNodeList(cluster string) ([]*WindowsClusterNodeStruct, error) {
var err error
var result []*WindowsClusterNodeStruct
return result, err
}
