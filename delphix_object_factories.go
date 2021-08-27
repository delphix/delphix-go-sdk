package delphix

// APIErrorFactory is just a simple function to instantiate the APIErrorStruct
func APIErrorFactory(
	Action string,
	CommandOutput string,
	Details string,
	Diagnoses []*DiagnosisResultStruct,
	Id string,
) APIErrorStruct {
	return APIErrorStruct{
		Type:          "APIError",
		Action:        Action,
		CommandOutput: CommandOutput,
		Details:       Details,
		Diagnoses:     Diagnoses,
		Id:            Id,
	}
}

// APISessionFactory is just a simple function to instantiate the APISessionStruct
func APISessionFactory(
	Client string,
	Locale string,
	Version *APIVersionStruct,
) APISessionStruct {
	return APISessionStruct{
		Type:    "APISession",
		Client:  Client,
		Locale:  Locale,
		Version: Version,
	}
}

// APIVersionFactory is just a simple function to instantiate the APIVersionStruct
func APIVersionFactory(
	Major *int,
	Micro *int,
	Minor *int,
) APIVersionStruct {
	return APIVersionStruct{
		Type:  "APIVersion",
		Major: Major,
		Micro: Micro,
		Minor: Minor,
	}
}

// ASEAttachDataFactory is just a simple function to instantiate the ASEAttachDataStruct
func ASEAttachDataFactory(
	Config string,
	DbCredentials Credential,
	DbUser string,
	DumpCredentials string,
	DumpHistoryFileEnabled *bool,
	ExternalFilePath string,
	LoadBackupPath string,
	LoadLocation string,
	MountBase string,
	Operations *LinkedSourceOperationsStruct,
	SourceHostUser string,
	StagingHostUser string,
	StagingOperations *StagingSourceOperationsStruct,
	StagingPostScript string,
	StagingPreScript string,
	StagingRepository string,
	ValidatedSyncMode string,
) ASEAttachDataStruct {
	return ASEAttachDataStruct{
		Type:                   "ASEAttachData",
		Config:                 Config,
		DbCredentials:          DbCredentials,
		DbUser:                 DbUser,
		DumpCredentials:        DumpCredentials,
		DumpHistoryFileEnabled: DumpHistoryFileEnabled,
		ExternalFilePath:       ExternalFilePath,
		LoadBackupPath:         LoadBackupPath,
		LoadLocation:           LoadLocation,
		MountBase:              MountBase,
		Operations:             Operations,
		SourceHostUser:         SourceHostUser,
		StagingHostUser:        StagingHostUser,
		StagingOperations:      StagingOperations,
		StagingPostScript:      StagingPostScript,
		StagingPreScript:       StagingPreScript,
		StagingRepository:      StagingRepository,
		ValidatedSyncMode:      ValidatedSyncMode,
	}
}

// ASEBackupLocationFactory is just a simple function to instantiate the ASEBackupLocationStruct
func ASEBackupLocationFactory(
	BackupHost string,
	BackupHostUser string,
	BackupServerName string,
) ASEBackupLocationStruct {
	return ASEBackupLocationStruct{
		Type:             "ASEBackupLocation",
		BackupHost:       BackupHost,
		BackupHostUser:   BackupHostUser,
		BackupServerName: BackupServerName,
	}
}

// ASECompatibilityCriteriaFactory is just a simple function to instantiate the ASECompatibilityCriteriaStruct
func ASECompatibilityCriteriaFactory(
	Architecture *int,
	Environment string,
	Os string,
	Processor string,
	StagingEnabled *bool,
) ASECompatibilityCriteriaStruct {
	return ASECompatibilityCriteriaStruct{
		Type:           "ASECompatibilityCriteria",
		Architecture:   Architecture,
		Environment:    Environment,
		Os:             Os,
		Processor:      Processor,
		StagingEnabled: StagingEnabled,
	}
}

// ASECreateTransformationParametersFactory is just a simple function to instantiate the ASECreateTransformationParametersStruct
func ASECreateTransformationParametersFactory(
	Container *ASEDBContainerStruct,
	EnvironmentUser string,
	Operations []SourceOperation,
	Repository string,
) ASECreateTransformationParametersStruct {
	return ASECreateTransformationParametersStruct{
		Type:            "ASECreateTransformationParameters",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Repository:      Repository,
	}
}

// ASEDBConfigConnectivityFactory is just a simple function to instantiate the ASEDBConfigConnectivityStruct
func ASEDBConfigConnectivityFactory(
	Credentials Credential,
	Username string,
) ASEDBConfigConnectivityStruct {
	return ASEDBConfigConnectivityStruct{
		Type:        "ASEDBConfigConnectivity",
		Credentials: Credentials,
		Username:    Username,
	}
}

// ASEDBContainerFactory is just a simple function to instantiate the ASEDBContainerStruct
func ASEDBContainerFactory(
	CreationTime string,
	CurrentTimeflow string,
	Description string,
	Group string,
	Masked *bool,
	Name string,
	Namespace string,
	Os string,
	PreviousTimeflow string,
	Processor string,
	ProvisionContainer string,
	Reference string,
	Runtime *ASEDBContainerRuntimeStruct,
	SourcingPolicy *SourcingPolicyStruct,
	Transformation *bool,
) ASEDBContainerStruct {
	return ASEDBContainerStruct{
		Type:               "ASEDBContainer",
		CreationTime:       CreationTime,
		CurrentTimeflow:    CurrentTimeflow,
		Description:        Description,
		Group:              Group,
		Masked:             Masked,
		Name:               Name,
		Namespace:          Namespace,
		Os:                 Os,
		PreviousTimeflow:   PreviousTimeflow,
		Processor:          Processor,
		ProvisionContainer: ProvisionContainer,
		Reference:          Reference,
		Runtime:            Runtime,
		SourcingPolicy:     SourcingPolicy,
		Transformation:     Transformation,
	}
}

// ASEDBContainerRuntimeFactory is just a simple function to instantiate the ASEDBContainerRuntimeStruct
func ASEDBContainerRuntimeFactory(
	LastRestoredBackupDate string,
	LastRestoredBackupTimeZone string,
	LogSyncActive *bool,
	PreProvisioningStatus *PreProvisioningRuntimeStruct,
) ASEDBContainerRuntimeStruct {
	return ASEDBContainerRuntimeStruct{
		Type:                       "ASEDBContainerRuntime",
		LastRestoredBackupDate:     LastRestoredBackupDate,
		LastRestoredBackupTimeZone: LastRestoredBackupTimeZone,
		LogSyncActive:              LogSyncActive,
		PreProvisioningStatus:      PreProvisioningStatus,
	}
}

// ASEExportParametersFactory is just a simple function to instantiate the ASEExportParametersStruct
func ASEExportParametersFactory(
	ConfigParams map[string]string,
	FileMappingRules string,
	RecoverDatabase *bool,
	SourceConfig ASEDBConfig,
	TimeflowPointParameters TimeflowPointParameters,
) ASEExportParametersStruct {
	return ASEExportParametersStruct{
		Type:                    "ASEExportParameters",
		ConfigParams:            ConfigParams,
		FileMappingRules:        FileMappingRules,
		RecoverDatabase:         RecoverDatabase,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// ASEHostEnvironmentParametersFactory is just a simple function to instantiate the ASEHostEnvironmentParametersStruct
func ASEHostEnvironmentParametersFactory(
	Credentials Credential,
	DbUser string,
) ASEHostEnvironmentParametersStruct {
	return ASEHostEnvironmentParametersStruct{
		Type:        "ASEHostEnvironmentParameters",
		Credentials: Credentials,
		DbUser:      DbUser,
	}
}

// ASEInstanceFactory is just a simple function to instantiate the ASEInstanceStruct
func ASEInstanceFactory(
	Credentials Credential,
	DbUser string,
	Discovered *bool,
	DumpHistoryFile string,
	Environment string,
	InstallationPath string,
	InstanceName string,
	InstanceOwner string,
	IsqlPath string,
	Name string,
	Namespace string,
	PageSize *int,
	Ports []*int,
	ProvisioningEnabled *bool,
	Reference string,
	ServicePrincipalName string,
	Staging *bool,
	Version string,
) ASEInstanceStruct {
	return ASEInstanceStruct{
		Type:                 "ASEInstance",
		Credentials:          Credentials,
		DbUser:               DbUser,
		Discovered:           Discovered,
		DumpHistoryFile:      DumpHistoryFile,
		Environment:          Environment,
		InstallationPath:     InstallationPath,
		InstanceName:         InstanceName,
		InstanceOwner:        InstanceOwner,
		IsqlPath:             IsqlPath,
		Name:                 Name,
		Namespace:            Namespace,
		PageSize:             PageSize,
		Ports:                Ports,
		ProvisioningEnabled:  ProvisioningEnabled,
		Reference:            Reference,
		ServicePrincipalName: ServicePrincipalName,
		Staging:              Staging,
		Version:              Version,
	}
}

// ASELatestBackupSyncParametersFactory is just a simple function to instantiate the ASELatestBackupSyncParametersStruct
func ASELatestBackupSyncParametersFactory(
	DropAndRecreateDevices *bool,
) ASELatestBackupSyncParametersStruct {
	return ASELatestBackupSyncParametersStruct{
		Type:                   "ASELatestBackupSyncParameters",
		DropAndRecreateDevices: DropAndRecreateDevices,
	}
}

// ASELinkDataFactory is just a simple function to instantiate the ASELinkDataStruct
func ASELinkDataFactory(
	Config string,
	DbCredentials Credential,
	DbUser string,
	DumpCredentials string,
	DumpHistoryFileEnabled *bool,
	ExternalFilePath string,
	LoadBackupPath string,
	LoadLocation string,
	MountBase string,
	Operations *LinkedSourceOperationsStruct,
	SourceHostUser string,
	SourcingPolicy *SourcingPolicyStruct,
	StagingHostUser string,
	StagingOperations *StagingSourceOperationsStruct,
	StagingPostScript string,
	StagingPreScript string,
	StagingRepository string,
	SyncParameters ASESyncParameters,
	ValidatedSyncMode string,
) ASELinkDataStruct {
	return ASELinkDataStruct{
		Type:                   "ASELinkData",
		Config:                 Config,
		DbCredentials:          DbCredentials,
		DbUser:                 DbUser,
		DumpCredentials:        DumpCredentials,
		DumpHistoryFileEnabled: DumpHistoryFileEnabled,
		ExternalFilePath:       ExternalFilePath,
		LoadBackupPath:         LoadBackupPath,
		LoadLocation:           LoadLocation,
		MountBase:              MountBase,
		Operations:             Operations,
		SourceHostUser:         SourceHostUser,
		SourcingPolicy:         SourcingPolicy,
		StagingHostUser:        StagingHostUser,
		StagingOperations:      StagingOperations,
		StagingPostScript:      StagingPostScript,
		StagingPreScript:       StagingPreScript,
		StagingRepository:      StagingRepository,
		SyncParameters:         SyncParameters,
		ValidatedSyncMode:      ValidatedSyncMode,
	}
}

// ASELinkedSourceFactory is just a simple function to instantiate the ASELinkedSourceStruct
func ASELinkedSourceFactory(
	Config string,
	Container string,
	Description string,
	DumpCredentials string,
	DumpHistoryFileEnabled *bool,
	ExternalFilePath string,
	Hosts []string,
	Linked *bool,
	LoadBackupPath string,
	LoadLocation string,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	Operations *LinkedSourceOperationsStruct,
	Reference string,
	Runtime *ASESourceRuntimeStruct,
	Staging *bool,
	StagingSource string,
	Status string,
	ValidatedSyncMode string,
	Virtual *bool,
) ASELinkedSourceStruct {
	return ASELinkedSourceStruct{
		Type:                   "ASELinkedSource",
		Config:                 Config,
		Container:              Container,
		Description:            Description,
		DumpCredentials:        DumpCredentials,
		DumpHistoryFileEnabled: DumpHistoryFileEnabled,
		ExternalFilePath:       ExternalFilePath,
		Hosts:                  Hosts,
		Linked:                 Linked,
		LoadBackupPath:         LoadBackupPath,
		LoadLocation:           LoadLocation,
		LogCollectionEnabled:   LogCollectionEnabled,
		Name:                   Name,
		Namespace:              Namespace,
		Operations:             Operations,
		Reference:              Reference,
		Runtime:                Runtime,
		Staging:                Staging,
		StagingSource:          StagingSource,
		Status:                 Status,
		ValidatedSyncMode:      ValidatedSyncMode,
		Virtual:                Virtual,
	}
}

// ASENewBackupSyncParametersFactory is just a simple function to instantiate the ASENewBackupSyncParametersStruct
func ASENewBackupSyncParametersFactory(
	DropAndRecreateDevices *bool,
) ASENewBackupSyncParametersStruct {
	return ASENewBackupSyncParametersStruct{
		Type:                   "ASENewBackupSyncParameters",
		DropAndRecreateDevices: DropAndRecreateDevices,
	}
}

// ASEPlatformParametersFactory is just a simple function to instantiate the ASEPlatformParametersStruct
func ASEPlatformParametersFactory() ASEPlatformParametersStruct {
	return ASEPlatformParametersStruct{
		Type: "ASEPlatformParameters",
	}
}

// ASEProvisionParametersFactory is just a simple function to instantiate the ASEProvisionParametersStruct
func ASEProvisionParametersFactory(
	Container *ASEDBContainerStruct,
	Masked *bool,
	MaskingJob string,
	Source *ASEVirtualSourceStruct,
	SourceConfig ASEDBConfig,
	TimeflowPointParameters TimeflowPointParameters,
	TruncateLogOnCheckpoint *bool,
) ASEProvisionParametersStruct {
	return ASEProvisionParametersStruct{
		Type:                    "ASEProvisionParameters",
		Container:               Container,
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		Source:                  Source,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
		TruncateLogOnCheckpoint: TruncateLogOnCheckpoint,
	}
}

// ASESIConfigFactory is just a simple function to instantiate the ASESIConfigStruct
func ASESIConfigFactory(
	Credentials Credential,
	DatabaseName string,
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	Reference string,
	Repository string,
	User string,
) ASESIConfigStruct {
	return ASESIConfigStruct{
		Type:            "ASESIConfig",
		Credentials:     Credentials,
		DatabaseName:    DatabaseName,
		Discovered:      Discovered,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Namespace:       Namespace,
		Reference:       Reference,
		Repository:      Repository,
		User:            User,
	}
}

// ASESnapshotFactory is just a simple function to instantiate the ASESnapshotStruct
func ASESnapshotFactory(
	Consistency string,
	Container string,
	CreationTime string,
	DbEncryptionKey string,
	FirstChangePoint *ASETimeflowPointStruct,
	LatestChangePoint *ASETimeflowPointStruct,
	MissingNonLoggedData *bool,
	Name string,
	Namespace string,
	Reference string,
	Retention *int,
	Runtime *ASESnapshotRuntimeStruct,
	Temporary *bool,
	Timeflow string,
	Timezone string,
	Version string,
) ASESnapshotStruct {
	return ASESnapshotStruct{
		Type:                 "ASESnapshot",
		Consistency:          Consistency,
		Container:            Container,
		CreationTime:         CreationTime,
		DbEncryptionKey:      DbEncryptionKey,
		FirstChangePoint:     FirstChangePoint,
		LatestChangePoint:    LatestChangePoint,
		MissingNonLoggedData: MissingNonLoggedData,
		Name:                 Name,
		Namespace:            Namespace,
		Reference:            Reference,
		Retention:            Retention,
		Runtime:              Runtime,
		Temporary:            Temporary,
		Timeflow:             Timeflow,
		Timezone:             Timezone,
		Version:              Version,
	}
}

// ASESnapshotRuntimeFactory is just a simple function to instantiate the ASESnapshotRuntimeStruct
func ASESnapshotRuntimeFactory(
	Provisionable *bool,
) ASESnapshotRuntimeStruct {
	return ASESnapshotRuntimeStruct{
		Type:          "ASESnapshotRuntime",
		Provisionable: Provisionable,
	}
}

// ASESourceConnectionInfoFactory is just a simple function to instantiate the ASESourceConnectionInfoStruct
func ASESourceConnectionInfoFactory(
	DatabaseName string,
	Host string,
	JdbcString string,
	Port *int,
	User string,
	Version string,
) ASESourceConnectionInfoStruct {
	return ASESourceConnectionInfoStruct{
		Type:         "ASESourceConnectionInfo",
		DatabaseName: DatabaseName,
		Host:         Host,
		JdbcString:   JdbcString,
		Port:         Port,
		User:         User,
		Version:      Version,
	}
}

// ASESourceRuntimeFactory is just a simple function to instantiate the ASESourceRuntimeStruct
func ASESourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	DatabaseSize float64,
	DurabilityLevel string,
	Enabled string,
	NotAccessibleReason string,
	Status string,
	TruncateLogOnCheckpoint *bool,
) ASESourceRuntimeStruct {
	return ASESourceRuntimeStruct{
		Type:                    "ASESourceRuntime",
		Accessible:              Accessible,
		AccessibleTimestamp:     AccessibleTimestamp,
		DatabaseSize:            DatabaseSize,
		DurabilityLevel:         DurabilityLevel,
		Enabled:                 Enabled,
		NotAccessibleReason:     NotAccessibleReason,
		Status:                  Status,
		TruncateLogOnCheckpoint: TruncateLogOnCheckpoint,
	}
}

// ASESpecificBackupSyncParametersFactory is just a simple function to instantiate the ASESpecificBackupSyncParametersStruct
func ASESpecificBackupSyncParametersFactory(
	BackupFiles []string,
	DropAndRecreateDevices *bool,
) ASESpecificBackupSyncParametersStruct {
	return ASESpecificBackupSyncParametersStruct{
		Type:                   "ASESpecificBackupSyncParameters",
		BackupFiles:            BackupFiles,
		DropAndRecreateDevices: DropAndRecreateDevices,
	}
}

// ASEStagingSourceFactory is just a simple function to instantiate the ASEStagingSourceStruct
func ASEStagingSourceFactory(
	Config string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	Operations *StagingSourceOperationsStruct,
	PostScript string,
	PreScript string,
	Reference string,
	Runtime *ASESourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) ASEStagingSourceStruct {
	return ASEStagingSourceStruct{
		Type:                 "ASEStagingSource",
		Config:               Config,
		Container:            Container,
		Description:          Description,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		MountBase:            MountBase,
		Name:                 Name,
		Namespace:            Namespace,
		Operations:           Operations,
		PostScript:           PostScript,
		PreScript:            PreScript,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		Status:               Status,
		Virtual:              Virtual,
	}
}

// ASETimeflowFactory is just a simple function to instantiate the ASETimeflowStruct
func ASETimeflowFactory(
	Container string,
	CreationType string,
	Name string,
	Namespace string,
	ParentPoint *ASETimeflowPointStruct,
	ParentSnapshot string,
	Reference string,
) ASETimeflowStruct {
	return ASETimeflowStruct{
		Type:           "ASETimeflow",
		Container:      Container,
		CreationType:   CreationType,
		Name:           Name,
		Namespace:      Namespace,
		ParentPoint:    ParentPoint,
		ParentSnapshot: ParentSnapshot,
		Reference:      Reference,
	}
}

// ASETimeflowPointFactory is just a simple function to instantiate the ASETimeflowPointStruct
func ASETimeflowPointFactory(
	Location string,
	Timeflow string,
	Timestamp string,
) ASETimeflowPointStruct {
	return ASETimeflowPointStruct{
		Type:      "ASETimeflowPoint",
		Location:  Location,
		Timeflow:  Timeflow,
		Timestamp: Timestamp,
	}
}

// ASEVirtualSourceFactory is just a simple function to instantiate the ASEVirtualSourceStruct
func ASEVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	Container string,
	Description string,
	FileMappingRules string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	Operations *VirtualSourceOperationsStruct,
	Reference string,
	Runtime *ASESourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) ASEVirtualSourceStruct {
	return ASEVirtualSourceStruct{
		Type:                            "ASEVirtualSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		Container:                       Container,
		Description:                     Description,
		FileMappingRules:                FileMappingRules,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		Operations:                      Operations,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Virtual:                         Virtual,
	}
}

// ActionFactory is just a simple function to instantiate the ActionStruct
func ActionFactory(
	ActionType string,
	Details string,
	EndTime string,
	FailureAction string,
	FailureDescription string,
	FailureMessageCode string,
	Namespace string,
	OriginIp string,
	ParentAction string,
	Reference string,
	StartTime string,
	State string,
	Title string,
	User string,
	UserAgent string,
	WorkSource string,
	WorkSourceName string,
	WorkSourcePrincipal string,
) ActionStruct {
	return ActionStruct{
		Type:                "Action",
		ActionType:          ActionType,
		Details:             Details,
		EndTime:             EndTime,
		FailureAction:       FailureAction,
		FailureDescription:  FailureDescription,
		FailureMessageCode:  FailureMessageCode,
		Namespace:           Namespace,
		OriginIp:            OriginIp,
		ParentAction:        ParentAction,
		Reference:           Reference,
		StartTime:           StartTime,
		State:               State,
		Title:               Title,
		User:                User,
		UserAgent:           UserAgent,
		WorkSource:          WorkSource,
		WorkSourceName:      WorkSourceName,
		WorkSourcePrincipal: WorkSourcePrincipal,
	}
}

// AdvancedSettingsInfoFactory is just a simple function to instantiate the AdvancedSettingsInfoStruct
func AdvancedSettingsInfoFactory() AdvancedSettingsInfoStruct {
	return AdvancedSettingsInfoStruct{
		Type: "AdvancedSettingsInfo",
	}
}

// AlertFactory is just a simple function to instantiate the AlertStruct
func AlertFactory(
	Event string,
	EventAction string,
	EventCommandOutput string,
	EventDescription string,
	EventResponse string,
	EventSeverity string,
	EventTitle string,
	Namespace string,
	Reference string,
	Target string,
	TargetName string,
	TargetObjectType string,
	Timestamp string,
) AlertStruct {
	return AlertStruct{
		Type:               "Alert",
		Event:              Event,
		EventAction:        EventAction,
		EventCommandOutput: EventCommandOutput,
		EventDescription:   EventDescription,
		EventResponse:      EventResponse,
		EventSeverity:      EventSeverity,
		EventTitle:         EventTitle,
		Namespace:          Namespace,
		Reference:          Reference,
		Target:             Target,
		TargetName:         TargetName,
		TargetObjectType:   TargetObjectType,
		Timestamp:          Timestamp,
	}
}

// AlertActionEmailListFactory is just a simple function to instantiate the AlertActionEmailListStruct
func AlertActionEmailListFactory(
	Addresses []string,
	Format string,
) AlertActionEmailListStruct {
	return AlertActionEmailListStruct{
		Type:      "AlertActionEmailList",
		Addresses: Addresses,
		Format:    Format,
	}
}

// AlertActionEmailUserFactory is just a simple function to instantiate the AlertActionEmailUserStruct
func AlertActionEmailUserFactory(
	Format string,
) AlertActionEmailUserStruct {
	return AlertActionEmailUserStruct{
		Type:   "AlertActionEmailUser",
		Format: Format,
	}
}

// AlertProfileFactory is just a simple function to instantiate the AlertProfileStruct
func AlertProfileFactory(
	Actions []AlertAction,
	FilterSpec AlertFilter,
	Namespace string,
	Reference string,
	User string,
) AlertProfileStruct {
	return AlertProfileStruct{
		Type:       "AlertProfile",
		Actions:    Actions,
		FilterSpec: FilterSpec,
		Namespace:  Namespace,
		Reference:  Reference,
		User:       User,
	}
}

// AndFilterFactory is just a simple function to instantiate the AndFilterStruct
func AndFilterFactory(
	SubFilters []AlertFilter,
) AndFilterStruct {
	return AndFilterStruct{
		Type:       "AndFilter",
		SubFilters: SubFilters,
	}
}

// AppDataAdditionalMountPointFactory is just a simple function to instantiate the AppDataAdditionalMountPointStruct
func AppDataAdditionalMountPointFactory(
	Environment string,
	MountPath string,
	SharedPath string,
) AppDataAdditionalMountPointStruct {
	return AppDataAdditionalMountPointStruct{
		Type:        "AppDataAdditionalMountPoint",
		Environment: Environment,
		MountPath:   MountPath,
		SharedPath:  SharedPath,
	}
}

// AppDataCachedMountPointFactory is just a simple function to instantiate the AppDataCachedMountPointStruct
func AppDataCachedMountPointFactory(
	Environment string,
	MountPath string,
	Ordinal *int,
	SharedPath string,
) AppDataCachedMountPointStruct {
	return AppDataCachedMountPointStruct{
		Type:        "AppDataCachedMountPoint",
		Environment: Environment,
		MountPath:   MountPath,
		Ordinal:     Ordinal,
		SharedPath:  SharedPath,
	}
}

// AppDataContainerFactory is just a simple function to instantiate the AppDataContainerStruct
func AppDataContainerFactory(
	CreationTime string,
	CurrentTimeflow string,
	Description string,
	Group string,
	Guid string,
	Masked *bool,
	Name string,
	Namespace string,
	Os string,
	PreviousTimeflow string,
	Processor string,
	ProvisionContainer string,
	Reference string,
	Runtime *AppDataContainerRuntimeStruct,
	SourcingPolicy *SourcingPolicyStruct,
	Toolkit string,
	Transformation *bool,
) AppDataContainerStruct {
	return AppDataContainerStruct{
		Type:               "AppDataContainer",
		CreationTime:       CreationTime,
		CurrentTimeflow:    CurrentTimeflow,
		Description:        Description,
		Group:              Group,
		Guid:               Guid,
		Masked:             Masked,
		Name:               Name,
		Namespace:          Namespace,
		Os:                 Os,
		PreviousTimeflow:   PreviousTimeflow,
		Processor:          Processor,
		ProvisionContainer: ProvisionContainer,
		Reference:          Reference,
		Runtime:            Runtime,
		SourcingPolicy:     SourcingPolicy,
		Toolkit:            Toolkit,
		Transformation:     Transformation,
	}
}

// AppDataContainerRuntimeFactory is just a simple function to instantiate the AppDataContainerRuntimeStruct
func AppDataContainerRuntimeFactory(
	LogSyncActive *bool,
	PreProvisioningStatus *PreProvisioningRuntimeStruct,
) AppDataContainerRuntimeStruct {
	return AppDataContainerRuntimeStruct{
		Type:                  "AppDataContainerRuntime",
		LogSyncActive:         LogSyncActive,
		PreProvisioningStatus: PreProvisioningStatus,
	}
}

// AppDataCreateTransformationParametersFactory is just a simple function to instantiate the AppDataCreateTransformationParametersStruct
func AppDataCreateTransformationParametersFactory(
	Container *AppDataContainerStruct,
	EnvironmentUser string,
	Operations []SourceOperation,
	Payload *JsonStruct,
	Repository string,
) AppDataCreateTransformationParametersStruct {
	return AppDataCreateTransformationParametersStruct{
		Type:            "AppDataCreateTransformationParameters",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Payload:         Payload,
		Repository:      Repository,
	}
}

// AppDataDirectLinkDataFactory is just a simple function to instantiate the AppDataDirectLinkDataStruct
func AppDataDirectLinkDataFactory(
	Config string,
	EnvironmentUser string,
	Excludes []string,
	FollowSymlinks []string,
	Operations *LinkedSourceOperationsStruct,
	Parameters *JsonStruct,
	SourcingPolicy *SourcingPolicyStruct,
	SyncParameters *AppDataSyncParametersStruct,
) AppDataDirectLinkDataStruct {
	return AppDataDirectLinkDataStruct{
		Type:            "AppDataDirectLinkData",
		Config:          Config,
		EnvironmentUser: EnvironmentUser,
		Excludes:        Excludes,
		FollowSymlinks:  FollowSymlinks,
		Operations:      Operations,
		Parameters:      Parameters,
		SourcingPolicy:  SourcingPolicy,
		SyncParameters:  SyncParameters,
	}
}

// AppDataDirectSourceConfigFactory is just a simple function to instantiate the AppDataDirectSourceConfigStruct
func AppDataDirectSourceConfigFactory(
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	Parameters *JsonStruct,
	Path string,
	Reference string,
	Repository string,
	Toolkit string,
) AppDataDirectSourceConfigStruct {
	return AppDataDirectSourceConfigStruct{
		Type:            "AppDataDirectSourceConfig",
		Discovered:      Discovered,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Namespace:       Namespace,
		Parameters:      Parameters,
		Path:            Path,
		Reference:       Reference,
		Repository:      Repository,
		Toolkit:         Toolkit,
	}
}

// AppDataEmptyVFilesCreationParametersFactory is just a simple function to instantiate the AppDataEmptyVFilesCreationParametersStruct
func AppDataEmptyVFilesCreationParametersFactory(
	Container *AppDataContainerStruct,
	Source *AppDataVirtualSourceStruct,
	SourceConfig AppDataSourceConfig,
) AppDataEmptyVFilesCreationParametersStruct {
	return AppDataEmptyVFilesCreationParametersStruct{
		Type:         "AppDataEmptyVFilesCreationParameters",
		Container:    Container,
		Source:       Source,
		SourceConfig: SourceConfig,
	}
}

// AppDataExportParametersFactory is just a simple function to instantiate the AppDataExportParametersStruct
func AppDataExportParametersFactory(
	FilesystemLayout *AppDataFilesystemLayoutStruct,
	SourceConfig AppDataSourceConfig,
	TimeflowPointParameters TimeflowPointParameters,
) AppDataExportParametersStruct {
	return AppDataExportParametersStruct{
		Type:                    "AppDataExportParameters",
		FilesystemLayout:        FilesystemLayout,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// AppDataFilesystemLayoutFactory is just a simple function to instantiate the AppDataFilesystemLayoutStruct
func AppDataFilesystemLayoutFactory(
	TargetDirectory string,
) AppDataFilesystemLayoutStruct {
	return AppDataFilesystemLayoutStruct{
		Type:            "AppDataFilesystemLayout",
		TargetDirectory: TargetDirectory,
	}
}

// AppDataLinkedDirectSourceFactory is just a simple function to instantiate the AppDataLinkedDirectSourceStruct
func AppDataLinkedDirectSourceFactory(
	Config string,
	Container string,
	Description string,
	Excludes []string,
	FollowSymlinks []string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	Operations *LinkedSourceOperationsStruct,
	Parameters *JsonStruct,
	Reference string,
	Runtime *AppDataSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Toolkit string,
	Virtual *bool,
) AppDataLinkedDirectSourceStruct {
	return AppDataLinkedDirectSourceStruct{
		Type:                 "AppDataLinkedDirectSource",
		Config:               Config,
		Container:            Container,
		Description:          Description,
		Excludes:             Excludes,
		FollowSymlinks:       FollowSymlinks,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Namespace:            Namespace,
		Operations:           Operations,
		Parameters:           Parameters,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		Status:               Status,
		Toolkit:              Toolkit,
		Virtual:              Virtual,
	}
}

// AppDataLinkedStagedSourceFactory is just a simple function to instantiate the AppDataLinkedStagedSourceStruct
func AppDataLinkedStagedSourceFactory(
	Config string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	Operations *LinkedSourceOperationsStruct,
	Parameters *JsonStruct,
	Reference string,
	Runtime *AppDataSourceRuntimeStruct,
	Staging *bool,
	StagingEnvironment string,
	StagingEnvironmentUser string,
	StagingMountBase string,
	Status string,
	Toolkit string,
	Virtual *bool,
) AppDataLinkedStagedSourceStruct {
	return AppDataLinkedStagedSourceStruct{
		Type:                   "AppDataLinkedStagedSource",
		Config:                 Config,
		Container:              Container,
		Description:            Description,
		Hosts:                  Hosts,
		Linked:                 Linked,
		LogCollectionEnabled:   LogCollectionEnabled,
		Name:                   Name,
		Namespace:              Namespace,
		Operations:             Operations,
		Parameters:             Parameters,
		Reference:              Reference,
		Runtime:                Runtime,
		Staging:                Staging,
		StagingEnvironment:     StagingEnvironment,
		StagingEnvironmentUser: StagingEnvironmentUser,
		StagingMountBase:       StagingMountBase,
		Status:                 Status,
		Toolkit:                Toolkit,
		Virtual:                Virtual,
	}
}

// AppDataPlatformParametersFactory is just a simple function to instantiate the AppDataPlatformParametersStruct
func AppDataPlatformParametersFactory(
	Payload *JsonStruct,
) AppDataPlatformParametersStruct {
	return AppDataPlatformParametersStruct{
		Type:    "AppDataPlatformParameters",
		Payload: Payload,
	}
}

// AppDataProvisionParametersFactory is just a simple function to instantiate the AppDataProvisionParametersStruct
func AppDataProvisionParametersFactory(
	Container *AppDataContainerStruct,
	Masked *bool,
	MaskingJob string,
	Source *AppDataVirtualSourceStruct,
	SourceConfig *AppDataDirectSourceConfigStruct,
	TimeflowPointParameters TimeflowPointParameters,
) AppDataProvisionParametersStruct {
	return AppDataProvisionParametersStruct{
		Type:                    "AppDataProvisionParameters",
		Container:               Container,
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		Source:                  Source,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// AppDataRepositoryFactory is just a simple function to instantiate the AppDataRepositoryStruct
func AppDataRepositoryFactory(
	Environment string,
	Name string,
	Namespace string,
	Parameters *JsonStruct,
	ProvisioningEnabled *bool,
	Reference string,
	Staging *bool,
	Toolkit string,
	Version string,
) AppDataRepositoryStruct {
	return AppDataRepositoryStruct{
		Type:                "AppDataRepository",
		Environment:         Environment,
		Name:                Name,
		Namespace:           Namespace,
		Parameters:          Parameters,
		ProvisioningEnabled: ProvisioningEnabled,
		Reference:           Reference,
		Staging:             Staging,
		Toolkit:             Toolkit,
		Version:             Version,
	}
}

// AppDataSnapshotFactory is just a simple function to instantiate the AppDataSnapshotStruct
func AppDataSnapshotFactory(
	Consistency string,
	Container string,
	CreationTime string,
	FirstChangePoint *AppDataTimeflowPointStruct,
	LatestChangePoint *AppDataTimeflowPointStruct,
	Metadata *JsonStruct,
	MissingNonLoggedData *bool,
	Name string,
	Namespace string,
	Reference string,
	Retention *int,
	Runtime *AppDataSnapshotRuntimeStruct,
	Temporary *bool,
	Timeflow string,
	Timezone string,
	Toolkit string,
	Version string,
) AppDataSnapshotStruct {
	return AppDataSnapshotStruct{
		Type:                 "AppDataSnapshot",
		Consistency:          Consistency,
		Container:            Container,
		CreationTime:         CreationTime,
		FirstChangePoint:     FirstChangePoint,
		LatestChangePoint:    LatestChangePoint,
		Metadata:             Metadata,
		MissingNonLoggedData: MissingNonLoggedData,
		Name:                 Name,
		Namespace:            Namespace,
		Reference:            Reference,
		Retention:            Retention,
		Runtime:              Runtime,
		Temporary:            Temporary,
		Timeflow:             Timeflow,
		Timezone:             Timezone,
		Toolkit:              Toolkit,
		Version:              Version,
	}
}

// AppDataSnapshotRuntimeFactory is just a simple function to instantiate the AppDataSnapshotRuntimeStruct
func AppDataSnapshotRuntimeFactory(
	Provisionable *bool,
) AppDataSnapshotRuntimeStruct {
	return AppDataSnapshotRuntimeStruct{
		Type:          "AppDataSnapshotRuntime",
		Provisionable: Provisionable,
	}
}

// AppDataSourceConnectionInfoFactory is just a simple function to instantiate the AppDataSourceConnectionInfoStruct
func AppDataSourceConnectionInfoFactory(
	Host string,
	Path string,
	Version string,
) AppDataSourceConnectionInfoStruct {
	return AppDataSourceConnectionInfoStruct{
		Type:    "AppDataSourceConnectionInfo",
		Host:    Host,
		Path:    Path,
		Version: Version,
	}
}

// AppDataSourceRuntimeFactory is just a simple function to instantiate the AppDataSourceRuntimeStruct
func AppDataSourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	DatabaseSize float64,
	Enabled string,
	NotAccessibleReason string,
	Status string,
) AppDataSourceRuntimeStruct {
	return AppDataSourceRuntimeStruct{
		Type:                "AppDataSourceRuntime",
		Accessible:          Accessible,
		AccessibleTimestamp: AccessibleTimestamp,
		DatabaseSize:        DatabaseSize,
		Enabled:             Enabled,
		NotAccessibleReason: NotAccessibleReason,
		Status:              Status,
	}
}

// AppDataStagedLinkDataFactory is just a simple function to instantiate the AppDataStagedLinkDataStruct
func AppDataStagedLinkDataFactory(
	Config string,
	EnvironmentUser string,
	Operations *LinkedSourceOperationsStruct,
	Parameters *JsonStruct,
	SourcingPolicy *SourcingPolicyStruct,
	StagingEnvironment string,
	StagingEnvironmentUser string,
	StagingMountBase string,
	SyncParameters *AppDataSyncParametersStruct,
) AppDataStagedLinkDataStruct {
	return AppDataStagedLinkDataStruct{
		Type:                   "AppDataStagedLinkData",
		Config:                 Config,
		EnvironmentUser:        EnvironmentUser,
		Operations:             Operations,
		Parameters:             Parameters,
		SourcingPolicy:         SourcingPolicy,
		StagingEnvironment:     StagingEnvironment,
		StagingEnvironmentUser: StagingEnvironmentUser,
		StagingMountBase:       StagingMountBase,
		SyncParameters:         SyncParameters,
	}
}

// AppDataStagedSourceConfigFactory is just a simple function to instantiate the AppDataStagedSourceConfigStruct
func AppDataStagedSourceConfigFactory(
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	Parameters *JsonStruct,
	Reference string,
	Repository string,
	Toolkit string,
) AppDataStagedSourceConfigStruct {
	return AppDataStagedSourceConfigStruct{
		Type:            "AppDataStagedSourceConfig",
		Discovered:      Discovered,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Namespace:       Namespace,
		Parameters:      Parameters,
		Reference:       Reference,
		Repository:      Repository,
		Toolkit:         Toolkit,
	}
}

// AppDataSyncParametersFactory is just a simple function to instantiate the AppDataSyncParametersStruct
func AppDataSyncParametersFactory(
	Parameters *JsonStruct,
) AppDataSyncParametersStruct {
	return AppDataSyncParametersStruct{
		Type:       "AppDataSyncParameters",
		Parameters: Parameters,
	}
}

// AppDataTimeflowFactory is just a simple function to instantiate the AppDataTimeflowStruct
func AppDataTimeflowFactory(
	Container string,
	CreationType string,
	Name string,
	Namespace string,
	ParentPoint *AppDataTimeflowPointStruct,
	ParentSnapshot string,
	Reference string,
) AppDataTimeflowStruct {
	return AppDataTimeflowStruct{
		Type:           "AppDataTimeflow",
		Container:      Container,
		CreationType:   CreationType,
		Name:           Name,
		Namespace:      Namespace,
		ParentPoint:    ParentPoint,
		ParentSnapshot: ParentSnapshot,
		Reference:      Reference,
	}
}

// AppDataTimeflowPointFactory is just a simple function to instantiate the AppDataTimeflowPointStruct
func AppDataTimeflowPointFactory(
	Location string,
	Timeflow string,
	Timestamp string,
) AppDataTimeflowPointStruct {
	return AppDataTimeflowPointStruct{
		Type:      "AppDataTimeflowPoint",
		Location:  Location,
		Timeflow:  Timeflow,
		Timestamp: Timestamp,
	}
}

// AppDataVirtualSourceFactory is just a simple function to instantiate the AppDataVirtualSourceStruct
func AppDataVirtualSourceFactory(
	AdditionalMountPoints []*AppDataAdditionalMountPointStruct,
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	Operations *VirtualSourceOperationsStruct,
	Parameters *JsonStruct,
	Reference string,
	Runtime *AppDataSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Toolkit string,
	Virtual *bool,
) AppDataVirtualSourceStruct {
	return AppDataVirtualSourceStruct{
		Type:                            "AppDataVirtualSource",
		AdditionalMountPoints:           AdditionalMountPoints,
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		Container:                       Container,
		Description:                     Description,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		Name:                            Name,
		Namespace:                       Namespace,
		Operations:                      Operations,
		Parameters:                      Parameters,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Toolkit:                         Toolkit,
		Virtual:                         Virtual,
	}
}

// AppDataWindowsTimeflowFactory is just a simple function to instantiate the AppDataWindowsTimeflowStruct
func AppDataWindowsTimeflowFactory(
	Container string,
	CreationType string,
	Name string,
	Namespace string,
	ParentPoint *AppDataTimeflowPointStruct,
	ParentSnapshot string,
	Reference string,
) AppDataWindowsTimeflowStruct {
	return AppDataWindowsTimeflowStruct{
		Type:           "AppDataWindowsTimeflow",
		Container:      Container,
		CreationType:   CreationType,
		Name:           Name,
		Namespace:      Namespace,
		ParentPoint:    ParentPoint,
		ParentSnapshot: ParentSnapshot,
		Reference:      Reference,
	}
}

// ApplyVersionParametersFactory is just a simple function to instantiate the ApplyVersionParametersStruct
func ApplyVersionParametersFactory(
	EnableSourcesOnFailure *bool,
	IgnoreQuiesceSourcesFailures *bool,
	QuiesceSources *bool,
	UpgradeType string,
	Verify *bool,
) ApplyVersionParametersStruct {
	return ApplyVersionParametersStruct{
		Type:                         "ApplyVersionParameters",
		EnableSourcesOnFailure:       EnableSourcesOnFailure,
		IgnoreQuiesceSourcesFailures: IgnoreQuiesceSourcesFailures,
		QuiesceSources:               QuiesceSources,
		UpgradeType:                  UpgradeType,
		Verify:                       Verify,
	}
}

// AttachSourceParametersFactory is just a simple function to instantiate the AttachSourceParametersStruct
func AttachSourceParametersFactory(
	AttachData AttachData,
) AttachSourceParametersStruct {
	return AttachSourceParametersStruct{
		Type:       "AttachSourceParameters",
		AttachData: AttachData,
	}
}

// AuthFilterParametersFactory is just a simple function to instantiate the AuthFilterParametersStruct
func AuthFilterParametersFactory(
	ObjectType string,
	Objects []string,
	Permission string,
) AuthFilterParametersStruct {
	return AuthFilterParametersStruct{
		Type:       "AuthFilterParameters",
		ObjectType: ObjectType,
		Objects:    Objects,
		Permission: Permission,
	}
}

// AuthFilterResultFactory is just a simple function to instantiate the AuthFilterResultStruct
func AuthFilterResultFactory(
	Objects []string,
) AuthFilterResultStruct {
	return AuthFilterResultStruct{
		Type:    "AuthFilterResult",
		Objects: Objects,
	}
}

// AuthGetByPropertiesParametersFactory is just a simple function to instantiate the AuthGetByPropertiesParametersStruct
func AuthGetByPropertiesParametersFactory(
	Role string,
	Target string,
	User string,
) AuthGetByPropertiesParametersStruct {
	return AuthGetByPropertiesParametersStruct{
		Type:   "AuthGetByPropertiesParameters",
		Role:   Role,
		Target: Target,
		User:   User,
	}
}

// AuthorizationFactory is just a simple function to instantiate the AuthorizationStruct
func AuthorizationFactory(
	Name string,
	Namespace string,
	Reference string,
	Role string,
	Target string,
	User string,
) AuthorizationStruct {
	return AuthorizationStruct{
		Type:      "Authorization",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
		Role:      Role,
		Target:    Target,
		User:      User,
	}
}

// AuthorizationConfigFactory is just a simple function to instantiate the AuthorizationConfigStruct
func AuthorizationConfigFactory(
	EnvironmentAndHostAuth *bool,
) AuthorizationConfigStruct {
	return AuthorizationConfigStruct{
		Type:                   "AuthorizationConfig",
		EnvironmentAndHostAuth: EnvironmentAndHostAuth,
	}
}

// AvailableUpgradeTypesFactory is just a simple function to instantiate the AvailableUpgradeTypesStruct
func AvailableUpgradeTypesFactory(
	UpgradeTypes []string,
) AvailableUpgradeTypesStruct {
	return AvailableUpgradeTypesStruct{
		Type:         "AvailableUpgradeTypes",
		UpgradeTypes: UpgradeTypes,
	}
}

// BatchContainerDeleteParametersFactory is just a simple function to instantiate the BatchContainerDeleteParametersStruct
func BatchContainerDeleteParametersFactory(
	Containers []string,
	DeleteParameters *DeleteParametersStruct,
) BatchContainerDeleteParametersStruct {
	return BatchContainerDeleteParametersStruct{
		Type:             "BatchContainerDeleteParameters",
		Containers:       Containers,
		DeleteParameters: DeleteParameters,
	}
}

// BatchContainerRefreshParametersFactory is just a simple function to instantiate the BatchContainerRefreshParametersStruct
func BatchContainerRefreshParametersFactory(
	Containers []string,
	Location string,
) BatchContainerRefreshParametersStruct {
	return BatchContainerRefreshParametersStruct{
		Type:       "BatchContainerRefreshParameters",
		Containers: Containers,
		Location:   Location,
	}
}

// BatchSnapshotDeleteParametersFactory is just a simple function to instantiate the BatchSnapshotDeleteParametersStruct
func BatchSnapshotDeleteParametersFactory(
	Snapshots []string,
) BatchSnapshotDeleteParametersStruct {
	return BatchSnapshotDeleteParametersStruct{
		Type:      "BatchSnapshotDeleteParameters",
		Snapshots: Snapshots,
	}
}

// BooleanEqualConstraintFactory is just a simple function to instantiate the BooleanEqualConstraintStruct
func BooleanEqualConstraintFactory(
	AxisName string,
	Equals *bool,
) BooleanEqualConstraintStruct {
	return BooleanEqualConstraintStruct{
		Type:     "BooleanEqualConstraint",
		AxisName: AxisName,
		Equals:   Equals,
	}
}

// CPUInfoFactory is just a simple function to instantiate the CPUInfoStruct
func CPUInfoFactory(
	Cores *int,
	Speed float64,
) CPUInfoStruct {
	return CPUInfoStruct{
		Type:  "CPUInfo",
		Cores: Cores,
		Speed: Speed,
	}
}

// CaCertificateFactory is just a simple function to instantiate the CaCertificateStruct
func CaCertificateFactory(
	Accepted *bool,
	IsCertificateAuthority *bool,
	IssuedByDN string,
	Issuer string,
	Md5Fingerprint string,
	Name string,
	Namespace string,
	NotAfter string,
	NotBefore string,
	Reference string,
	SerialNumber string,
	Sha1Fingerprint string,
	SubjectAlternativeNames []string,
) CaCertificateStruct {
	return CaCertificateStruct{
		Type:                    "CaCertificate",
		Accepted:                Accepted,
		IsCertificateAuthority:  IsCertificateAuthority,
		IssuedByDN:              IssuedByDN,
		Issuer:                  Issuer,
		Md5Fingerprint:          Md5Fingerprint,
		Name:                    Name,
		Namespace:               Namespace,
		NotAfter:                NotAfter,
		NotBefore:               NotBefore,
		Reference:               Reference,
		SerialNumber:            SerialNumber,
		Sha1Fingerprint:         Sha1Fingerprint,
		SubjectAlternativeNames: SubjectAlternativeNames,
	}
}

// CapacityBreakdownFactory is just a simple function to instantiate the CapacityBreakdownStruct
func CapacityBreakdownFactory(
	ActiveSpace float64,
	ActualSpace float64,
	DescendantSpace float64,
	FallbackIngestedSizeTimestamp string,
	IngestedSize float64,
	LogSpace float64,
	ManualSpace float64,
	PolicySpace float64,
	SyncSpace float64,
	TimeflowUnvirtualizedSpace float64,
	UnownedSnapshotSpace float64,
	UnvirtualizedSpace float64,
) CapacityBreakdownStruct {
	return CapacityBreakdownStruct{
		Type:                          "CapacityBreakdown",
		ActiveSpace:                   ActiveSpace,
		ActualSpace:                   ActualSpace,
		DescendantSpace:               DescendantSpace,
		FallbackIngestedSizeTimestamp: FallbackIngestedSizeTimestamp,
		IngestedSize:                  IngestedSize,
		LogSpace:                      LogSpace,
		ManualSpace:                   ManualSpace,
		PolicySpace:                   PolicySpace,
		SyncSpace:                     SyncSpace,
		TimeflowUnvirtualizedSpace:    TimeflowUnvirtualizedSpace,
		UnownedSnapshotSpace:          UnownedSnapshotSpace,
		UnvirtualizedSpace:            UnvirtualizedSpace,
	}
}

// CapacityDeleteDependenciesParametersFactory is just a simple function to instantiate the CapacityDeleteDependenciesParametersStruct
func CapacityDeleteDependenciesParametersFactory(
	DependencyTree *DeletionDependencyStruct,
) CapacityDeleteDependenciesParametersStruct {
	return CapacityDeleteDependenciesParametersStruct{
		Type:           "CapacityDeleteDependenciesParameters",
		DependencyTree: DependencyTree,
	}
}

// CertificateConfigFactory is just a simple function to instantiate the CertificateConfigStruct
func CertificateConfigFactory(
	CertificateExpirationCriticalThreshold *int,
	CertificateExpirationWarningThreshold *int,
	EnableUserManagedClientAuthForEngineToHostDsp *bool,
	EnableUserManagedClientAuthForNetworkThroughputTests *bool,
	EnableUserManagedClientAuthForReplication *bool,
	EnableUserManagedServerAuthForEngineToHostDsp *bool,
	EnableUserManagedServerAuthForNetworkThroughputTests *bool,
	EnableUserManagedServerAuthForReplication *bool,
	ValidateWindowsConnectorCertificate *bool,
) CertificateConfigStruct {
	return CertificateConfigStruct{
		Type:                                   "CertificateConfig",
		CertificateExpirationCriticalThreshold: CertificateExpirationCriticalThreshold,
		CertificateExpirationWarningThreshold:  CertificateExpirationWarningThreshold,
		EnableUserManagedClientAuthForEngineToHostDsp:        EnableUserManagedClientAuthForEngineToHostDsp,
		EnableUserManagedClientAuthForNetworkThroughputTests: EnableUserManagedClientAuthForNetworkThroughputTests,
		EnableUserManagedClientAuthForReplication:            EnableUserManagedClientAuthForReplication,
		EnableUserManagedServerAuthForEngineToHostDsp:        EnableUserManagedServerAuthForEngineToHostDsp,
		EnableUserManagedServerAuthForNetworkThroughputTests: EnableUserManagedServerAuthForNetworkThroughputTests,
		EnableUserManagedServerAuthForReplication:            EnableUserManagedServerAuthForReplication,
		ValidateWindowsConnectorCertificate:                  ValidateWindowsConnectorCertificate,
	}
}

// CertificateFetchParametersFactory is just a simple function to instantiate the CertificateFetchParametersStruct
func CertificateFetchParametersFactory(
	Host string,
	Port *int,
) CertificateFetchParametersStruct {
	return CertificateFetchParametersStruct{
		Type: "CertificateFetchParameters",
		Host: Host,
		Port: Port,
	}
}

// CertificateSigningRequestFactory is just a simple function to instantiate the CertificateSigningRequestStruct
func CertificateSigningRequestFactory(
	EndEntity EndEntity,
	KeyPair KeyPair,
	Name string,
	Namespace string,
	Reference string,
	RequestInPem string,
	SubjectAlternativeNames []string,
) CertificateSigningRequestStruct {
	return CertificateSigningRequestStruct{
		Type:                    "CertificateSigningRequest",
		EndEntity:               EndEntity,
		KeyPair:                 KeyPair,
		Name:                    Name,
		Namespace:               Namespace,
		Reference:               Reference,
		RequestInPem:            RequestInPem,
		SubjectAlternativeNames: SubjectAlternativeNames,
	}
}

// CertificateSigningRequestCreateParametersFactory is just a simple function to instantiate the CertificateSigningRequestCreateParametersStruct
func CertificateSigningRequestCreateParametersFactory(
	Dname X500DistinguishedName,
	EndEntity EndEntity,
	ForceReplace *bool,
	KeyPair KeyPair,
	SubjectAlternativeNames []string,
) CertificateSigningRequestCreateParametersStruct {
	return CertificateSigningRequestCreateParametersStruct{
		Type:                    "CertificateSigningRequestCreateParameters",
		Dname:                   Dname,
		EndEntity:               EndEntity,
		ForceReplace:            ForceReplace,
		KeyPair:                 KeyPair,
		SubjectAlternativeNames: SubjectAlternativeNames,
	}
}

// CertificateUploadParametersFactory is just a simple function to instantiate the CertificateUploadParametersStruct
func CertificateUploadParametersFactory(
	Alias string,
	Keypass string,
	KeystoreType string,
	Storepass string,
) CertificateUploadParametersStruct {
	return CertificateUploadParametersStruct{
		Type:         "CertificateUploadParameters",
		Alias:        Alias,
		Keypass:      Keypass,
		KeystoreType: KeystoreType,
		Storepass:    Storepass,
	}
}

// ChangeCurrentPasswordPolicyParametersFactory is just a simple function to instantiate the ChangeCurrentPasswordPolicyParametersStruct
func ChangeCurrentPasswordPolicyParametersFactory(
	CurrentPasswordPolicy string,
	UserType string,
) ChangeCurrentPasswordPolicyParametersStruct {
	return ChangeCurrentPasswordPolicyParametersStruct{
		Type:                  "ChangeCurrentPasswordPolicyParameters",
		CurrentPasswordPolicy: CurrentPasswordPolicy,
		UserType:              UserType,
	}
}

// CipherSuiteFactory is just a simple function to instantiate the CipherSuiteStruct
func CipherSuiteFactory(
	Name string,
) CipherSuiteStruct {
	return CipherSuiteStruct{
		Type: "CipherSuite",
		Name: Name,
	}
}

// CloudEnableParametersFactory is just a simple function to instantiate the CloudEnableParametersStruct
func CloudEnableParametersFactory(
	ProxyConfiguration *ProxyConfigurationStruct,
	ProxyMode string,
) CloudEnableParametersStruct {
	return CloudEnableParametersStruct{
		Type:               "CloudEnableParameters",
		ProxyConfiguration: ProxyConfiguration,
		ProxyMode:          ProxyMode,
	}
}

// CloudStatusFactory is just a simple function to instantiate the CloudStatusStruct
func CloudStatusFactory(
	CentralManagementProductName string,
	DelphixDataServicesComponentInfo *DelphixDataServicesComponentInfoStruct,
	DelphixDataServicesComponentStatus string,
	ProxyConfiguration *ProxyConfigurationStruct,
	ProxyMode string,
) CloudStatusStruct {
	return CloudStatusStruct{
		Type:                               "CloudStatus",
		CentralManagementProductName:       CentralManagementProductName,
		DelphixDataServicesComponentInfo:   DelphixDataServicesComponentInfo,
		DelphixDataServicesComponentStatus: DelphixDataServicesComponentStatus,
		ProxyConfiguration:                 ProxyConfiguration,
		ProxyMode:                          ProxyMode,
	}
}

// CommvaultConnectivityParametersFactory is just a simple function to instantiate the CommvaultConnectivityParametersStruct
func CommvaultConnectivityParametersFactory(
	CommserveHostName string,
	Environment string,
	EnvironmentUser string,
	SourceClientName string,
	StagingClientName string,
) CommvaultConnectivityParametersStruct {
	return CommvaultConnectivityParametersStruct{
		Type:              "CommvaultConnectivityParameters",
		CommserveHostName: CommserveHostName,
		Environment:       Environment,
		EnvironmentUser:   EnvironmentUser,
		SourceClientName:  SourceClientName,
		StagingClientName: StagingClientName,
	}
}

// CompatibilityCriteriaFactory is just a simple function to instantiate the CompatibilityCriteriaStruct
func CompatibilityCriteriaFactory(
	Architecture *int,
	Environment string,
	Os string,
	Processor string,
	StagingEnabled *bool,
) CompatibilityCriteriaStruct {
	return CompatibilityCriteriaStruct{
		Type:           "CompatibilityCriteria",
		Architecture:   Architecture,
		Environment:    Environment,
		Os:             Os,
		Processor:      Processor,
		StagingEnabled: StagingEnabled,
	}
}

// CompatibleRepositoriesResultFactory is just a simple function to instantiate the CompatibleRepositoriesResultStruct
func CompatibleRepositoriesResultFactory(
	Criteria *CompatibilityCriteriaStruct,
	Repositories []SourceRepository,
) CompatibleRepositoriesResultStruct {
	return CompatibleRepositoriesResultStruct{
		Type:         "CompatibleRepositoriesResult",
		Criteria:     Criteria,
		Repositories: Repositories,
	}
}

// ConfiguredStorageDeviceFactory is just a simple function to instantiate the ConfiguredStorageDeviceStruct
func ConfiguredStorageDeviceFactory(
	BootDevice *bool,
	Configured *bool,
	ExpandableSize float64,
	Fragmentation string,
	Model string,
	Name string,
	Namespace string,
	Reference string,
	Serial string,
	Size float64,
	UsedSize float64,
	Vendor string,
) ConfiguredStorageDeviceStruct {
	return ConfiguredStorageDeviceStruct{
		Type:           "ConfiguredStorageDevice",
		BootDevice:     BootDevice,
		Configured:     Configured,
		ExpandableSize: ExpandableSize,
		Fragmentation:  Fragmentation,
		Model:          Model,
		Name:           Name,
		Namespace:      Namespace,
		Reference:      Reference,
		Serial:         Serial,
		Size:           Size,
		UsedSize:       UsedSize,
		Vendor:         Vendor,
	}
}

// ConnectorConnectivityFactory is just a simple function to instantiate the ConnectorConnectivityStruct
func ConnectorConnectivityFactory(
	Address string,
	Credentials Credential,
	Port *int,
	Proxy string,
	Username string,
) ConnectorConnectivityStruct {
	return ConnectorConnectivityStruct{
		Type:        "ConnectorConnectivity",
		Address:     Address,
		Credentials: Credentials,
		Port:        Port,
		Proxy:       Proxy,
		Username:    Username,
	}
}

// ContainerUtilizationFactory is just a simple function to instantiate the ContainerUtilizationStruct
func ContainerUtilizationFactory(
	Container string,
	Deleted *bool,
	Hidden *bool,
	Utilization []*ContainerUtilizationIntervalStruct,
) ContainerUtilizationStruct {
	return ContainerUtilizationStruct{
		Type:        "ContainerUtilization",
		Container:   Container,
		Deleted:     Deleted,
		Hidden:      Hidden,
		Utilization: Utilization,
	}
}

// ContainerUtilizationIntervalFactory is just a simple function to instantiate the ContainerUtilizationIntervalStruct
func ContainerUtilizationIntervalFactory(
	AverageThroughput float64,
	Timestamp string,
) ContainerUtilizationIntervalStruct {
	return ContainerUtilizationIntervalStruct{
		Type:              "ContainerUtilizationInterval",
		AverageThroughput: AverageThroughput,
		Timestamp:         Timestamp,
	}
}

// CpuUtilDatapointFactory is just a simple function to instantiate the CpuUtilDatapointStruct
func CpuUtilDatapointFactory(
	Dtrace *int,
	Idle *int,
	Kernel *int,
	Timestamp string,
	User *int,
) CpuUtilDatapointStruct {
	return CpuUtilDatapointStruct{
		Type:      "CpuUtilDatapoint",
		Dtrace:    Dtrace,
		Idle:      Idle,
		Kernel:    Kernel,
		Timestamp: Timestamp,
		User:      User,
	}
}

// CpuUtilDatapointStreamFactory is just a simple function to instantiate the CpuUtilDatapointStreamStruct
func CpuUtilDatapointStreamFactory(
	Cpu *int,
	Datapoints []Datapoint,
) CpuUtilDatapointStreamStruct {
	return CpuUtilDatapointStreamStruct{
		Type:       "CpuUtilDatapointStream",
		Cpu:        Cpu,
		Datapoints: Datapoints,
	}
}

// CreateMaskingJobTransformationParametersFactory is just a simple function to instantiate the CreateMaskingJobTransformationParametersStruct
func CreateMaskingJobTransformationParametersFactory(
	Container Container,
	EnvironmentUser string,
	Repository string,
) CreateMaskingJobTransformationParametersStruct {
	return CreateMaskingJobTransformationParametersStruct{
		Type:            "CreateMaskingJobTransformationParameters",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Repository:      Repository,
	}
}

// CredentialUpdateParametersFactory is just a simple function to instantiate the CredentialUpdateParametersStruct
func CredentialUpdateParametersFactory(
	NewCredential *PasswordCredentialStruct,
	OldCredential *PasswordCredentialStruct,
) CredentialUpdateParametersStruct {
	return CredentialUpdateParametersStruct{
		Type:          "CredentialUpdateParameters",
		NewCredential: NewCredential,
		OldCredential: OldCredential,
	}
}

// CredentialsEnvVarsFactory is just a simple function to instantiate the CredentialsEnvVarsStruct
func CredentialsEnvVarsFactory(
	BaseVarName string,
	Credentials Credential,
) CredentialsEnvVarsStruct {
	return CredentialsEnvVarsStruct{
		Type:        "CredentialsEnvVars",
		BaseVarName: BaseVarName,
		Credentials: Credentials,
	}
}

// CurrentConsumerCapacityDataFactory is just a simple function to instantiate the CurrentConsumerCapacityDataStruct
func CurrentConsumerCapacityDataFactory(
	Breakdown *CapacityBreakdownStruct,
	Container string,
	Group string,
	GroupName string,
	MultiOwner *bool,
	Name string,
	Parent string,
	StorageContainer string,
	Timestamp string,
) CurrentConsumerCapacityDataStruct {
	return CurrentConsumerCapacityDataStruct{
		Type:             "CurrentConsumerCapacityData",
		Breakdown:        Breakdown,
		Container:        Container,
		Group:            Group,
		GroupName:        GroupName,
		MultiOwner:       MultiOwner,
		Name:             Name,
		Parent:           Parent,
		StorageContainer: StorageContainer,
		Timestamp:        Timestamp,
	}
}

// CurrentGroupCapacityDataFactory is just a simple function to instantiate the CurrentGroupCapacityDataStruct
func CurrentGroupCapacityDataFactory(
	Group string,
	Source *CapacityBreakdownStruct,
	Timestamp string,
	Virtual *CapacityBreakdownStruct,
) CurrentGroupCapacityDataStruct {
	return CurrentGroupCapacityDataStruct{
		Type:      "CurrentGroupCapacityData",
		Group:     Group,
		Source:    Source,
		Timestamp: Timestamp,
		Virtual:   Virtual,
	}
}

// CurrentSystemCapacityDataFactory is just a simple function to instantiate the CurrentSystemCapacityDataStruct
func CurrentSystemCapacityDataFactory(
	Source *CapacityBreakdownStruct,
	Timestamp string,
	TotalSpace float64,
	Virtual *CapacityBreakdownStruct,
) CurrentSystemCapacityDataStruct {
	return CurrentSystemCapacityDataStruct{
		Type:       "CurrentSystemCapacityData",
		Source:     Source,
		Timestamp:  Timestamp,
		TotalSpace: TotalSpace,
		Virtual:    Virtual,
	}
}

// CyberArkPasswordVaultFactory is just a simple function to instantiate the CyberArkPasswordVaultStruct
func CyberArkPasswordVaultFactory(
	ApplicationId string,
	ClientCertificate ClientCertificate,
	Host string,
	Name string,
	Namespace string,
	Port *int,
	Reference string,
) CyberArkPasswordVaultStruct {
	return CyberArkPasswordVaultStruct{
		Type:              "CyberArkPasswordVault",
		ApplicationId:     ApplicationId,
		ClientCertificate: ClientCertificate,
		Host:              Host,
		Name:              Name,
		Namespace:         Namespace,
		Port:              Port,
		Reference:         Reference,
	}
}

// CyberArkPasswordVaultTestParametersFactory is just a simple function to instantiate the CyberArkPasswordVaultTestParametersStruct
func CyberArkPasswordVaultTestParametersFactory(
	ApplicationId string,
	ClientCertificate ClientCertificate,
	Host string,
	Port *int,
) CyberArkPasswordVaultTestParametersStruct {
	return CyberArkPasswordVaultTestParametersStruct{
		Type:              "CyberArkPasswordVaultTestParameters",
		ApplicationId:     ApplicationId,
		ClientCertificate: ClientCertificate,
		Host:              Host,
		Port:              Port,
	}
}

// CyberArkVaultCredentialFactory is just a simple function to instantiate the CyberArkVaultCredentialStruct
func CyberArkVaultCredentialFactory(
	QueryString string,
	Vault string,
) CyberArkVaultCredentialStruct {
	return CyberArkVaultCredentialStruct{
		Type:        "CyberArkVaultCredential",
		QueryString: QueryString,
		Vault:       Vault,
	}
}

// DNSConfigFactory is just a simple function to instantiate the DNSConfigStruct
func DNSConfigFactory(
	Domain []string,
	Servers []string,
	Source string,
) DNSConfigStruct {
	return DNSConfigStruct{
		Type:    "DNSConfig",
		Domain:  Domain,
		Servers: Servers,
		Source:  Source,
	}
}

// DSPAutotunerParametersFactory is just a simple function to instantiate the DSPAutotunerParametersStruct
func DSPAutotunerParametersFactory(
	DestinationType string,
	Direction string,
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfoStruct,
	RemoteHost string,
) DSPAutotunerParametersStruct {
	return DSPAutotunerParametersStruct{
		Type:                    "DSPAutotunerParameters",
		DestinationType:         DestinationType,
		Direction:               Direction,
		RemoteDelphixEngineInfo: RemoteDelphixEngineInfo,
		RemoteHost:              RemoteHost,
	}
}

// DSPBestParametersFactory is just a simple function to instantiate the DSPBestParametersStruct
func DSPBestParametersFactory(
	BufferSize *int,
	DestinationType string,
	Namespace string,
	NumConnections *int,
	QueueDepth *int,
	Reference string,
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfoStruct,
	RemoteHost string,
	Throughput float64,
) DSPBestParametersStruct {
	return DSPBestParametersStruct{
		Type:                    "DSPBestParameters",
		BufferSize:              BufferSize,
		DestinationType:         DestinationType,
		Namespace:               Namespace,
		NumConnections:          NumConnections,
		QueueDepth:              QueueDepth,
		Reference:               Reference,
		RemoteDelphixEngineInfo: RemoteDelphixEngineInfo,
		RemoteHost:              RemoteHost,
		Throughput:              Throughput,
	}
}

// DSPOptionsFactory is just a simple function to instantiate the DSPOptionsStruct
func DSPOptionsFactory(
	BandwidthLimit *int,
	Compression *bool,
	Encryption *bool,
	NumConnections *int,
) DSPOptionsStruct {
	return DSPOptionsStruct{
		Type:           "DSPOptions",
		BandwidthLimit: BandwidthLimit,
		Compression:    Compression,
		Encryption:     Encryption,
		NumConnections: NumConnections,
	}
}

// DataResultFactory is just a simple function to instantiate the DataResultStruct
func DataResultFactory(
	Action string,
	Job string,
	Result string,
	Status string,
) DataResultStruct {
	return DataResultStruct{
		Type:   "DataResult",
		Action: Action,
		Job:    Job,
		Result: Result,
		Status: Status,
	}
}

// DatabaseTemplateFactory is just a simple function to instantiate the DatabaseTemplateStruct
func DatabaseTemplateFactory(
	Description string,
	Name string,
	Namespace string,
	Parameters map[string]string,
	Reference string,
	SourceType string,
) DatabaseTemplateStruct {
	return DatabaseTemplateStruct{
		Type:        "DatabaseTemplate",
		Description: Description,
		Name:        Name,
		Namespace:   Namespace,
		Parameters:  Parameters,
		Reference:   Reference,
		SourceType:  SourceType,
	}
}

// DatabaseTemplateConfigFactory is just a simple function to instantiate the DatabaseTemplateConfigStruct
func DatabaseTemplateConfigFactory(
	Reserved []string,
	SourceType string,
) DatabaseTemplateConfigStruct {
	return DatabaseTemplateConfigStruct{
		Type:       "DatabaseTemplateConfig",
		Reserved:   Reserved,
		SourceType: SourceType,
	}
}

// DatapointSetFactory is just a simple function to instantiate the DatapointSetStruct
func DatapointSetFactory(
	DatapointStreams []DatapointStream,
	Overflow *bool,
	Resolution *int,
) DatapointSetStruct {
	return DatapointSetStruct{
		Type:             "DatapointSet",
		DatapointStreams: DatapointStreams,
		Overflow:         Overflow,
		Resolution:       Resolution,
	}
}

// DeleteParametersFactory is just a simple function to instantiate the DeleteParametersStruct
func DeleteParametersFactory(
	Force *bool,
) DeleteParametersStruct {
	return DeleteParametersStruct{
		Type:  "DeleteParameters",
		Force: Force,
	}
}

// DeletionDependencyFactory is just a simple function to instantiate the DeletionDependencyStruct
func DeletionDependencyFactory(
	Dependencies []TypedObject,
	DisplayName string,
	Locked *bool,
	NamespaceName string,
	Prerequisites []*DeletionDependencyPrerequisiteStruct,
	Size float64,
	TargetReference string,
	TargetType string,
) DeletionDependencyStruct {
	return DeletionDependencyStruct{
		Type:            "DeletionDependency",
		Dependencies:    Dependencies,
		DisplayName:     DisplayName,
		Locked:          Locked,
		NamespaceName:   NamespaceName,
		Prerequisites:   Prerequisites,
		Size:            Size,
		TargetReference: TargetReference,
		TargetType:      TargetType,
	}
}

// DeletionDependencyPrerequisiteFactory is just a simple function to instantiate the DeletionDependencyPrerequisiteStruct
func DeletionDependencyPrerequisiteFactory(
	DisplayName string,
	Locked *bool,
	NamespaceName string,
	OperationType string,
	TargetReference string,
	TargetType string,
) DeletionDependencyPrerequisiteStruct {
	return DeletionDependencyPrerequisiteStruct{
		Type:            "DeletionDependencyPrerequisite",
		DisplayName:     DisplayName,
		Locked:          Locked,
		NamespaceName:   NamespaceName,
		OperationType:   OperationType,
		TargetReference: TargetReference,
		TargetType:      TargetType,
	}
}

// DelphixDataServicesComponentInfoFactory is just a simple function to instantiate the DelphixDataServicesComponentInfoStruct
func DelphixDataServicesComponentInfoFactory(
	BuildTimestamp string,
	BuildVersion string,
	Status string,
) DelphixDataServicesComponentInfoStruct {
	return DelphixDataServicesComponentInfoStruct{
		Type:           "DelphixDataServicesComponentInfo",
		BuildTimestamp: BuildTimestamp,
		BuildVersion:   BuildVersion,
		Status:         Status,
	}
}

// DelphixManagedBackupIngestionStrategyFactory is just a simple function to instantiate the DelphixManagedBackupIngestionStrategyStruct
func DelphixManagedBackupIngestionStrategyFactory(
	BackupPolicy string,
	CompressionEnabled *bool,
) DelphixManagedBackupIngestionStrategyStruct {
	return DelphixManagedBackupIngestionStrategyStruct{
		Type:               "DelphixManagedBackupIngestionStrategy",
		BackupPolicy:       BackupPolicy,
		CompressionEnabled: CompressionEnabled,
	}
}

// DetachSourceParametersFactory is just a simple function to instantiate the DetachSourceParametersStruct
func DetachSourceParametersFactory(
	Source string,
) DetachSourceParametersStruct {
	return DetachSourceParametersStruct{
		Type:   "DetachSourceParameters",
		Source: Source,
	}
}

// DiagnosisResultFactory is just a simple function to instantiate the DiagnosisResultStruct
func DiagnosisResultFactory(
	Failure *bool,
	Message string,
	MessageCode string,
) DiagnosisResultStruct {
	return DiagnosisResultStruct{
		Type:        "DiagnosisResult",
		Failure:     Failure,
		Message:     Message,
		MessageCode: MessageCode,
	}
}

// DiskOpsDatapointStreamFactory is just a simple function to instantiate the DiskOpsDatapointStreamStruct
func DiskOpsDatapointStreamFactory(
	Datapoints []Datapoint,
	Device string,
	Error *bool,
	Op string,
) DiskOpsDatapointStreamStruct {
	return DiskOpsDatapointStreamStruct{
		Type:       "DiskOpsDatapointStream",
		Datapoints: Datapoints,
		Device:     Device,
		Error:      Error,
		Op:         Op,
	}
}

// DomainFactory is just a simple function to instantiate the DomainStruct
func DomainFactory(
	Name string,
	Namespace string,
	Reference string,
) DomainStruct {
	return DomainStruct{
		Type:      "Domain",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// DspConnectivityFactory is just a simple function to instantiate the DspConnectivityStruct
func DspConnectivityFactory(
	Host string,
	NumberOfConsecutiveConnections *int,
	NumberOfSeconds *int,
	NumberOfThreads *int,
	User string,
) DspConnectivityStruct {
	return DspConnectivityStruct{
		Type:                           "DspConnectivity",
		Host:                           Host,
		NumberOfConsecutiveConnections: NumberOfConsecutiveConnections,
		NumberOfSeconds:                NumberOfSeconds,
		NumberOfThreads:                NumberOfThreads,
		User:                           User,
	}
}

// DxFsIoQueueOpsDatapointStreamFactory is just a simple function to instantiate the DxFsIoQueueOpsDatapointStreamStruct
func DxFsIoQueueOpsDatapointStreamFactory(
	Datapoints []Datapoint,
	Op string,
	Priority string,
) DxFsIoQueueOpsDatapointStreamStruct {
	return DxFsIoQueueOpsDatapointStreamStruct{
		Type:       "DxFsIoQueueOpsDatapointStream",
		Datapoints: Datapoints,
		Op:         Op,
		Priority:   Priority,
	}
}

// DxFsOpsDatapointStreamFactory is just a simple function to instantiate the DxFsOpsDatapointStreamStruct
func DxFsOpsDatapointStreamFactory(
	Cached *bool,
	Datapoints []Datapoint,
	Op string,
	Path string,
	Sync *bool,
) DxFsOpsDatapointStreamStruct {
	return DxFsOpsDatapointStreamStruct{
		Type:       "DxFsOpsDatapointStream",
		Cached:     Cached,
		Datapoints: Datapoints,
		Op:         Op,
		Path:       Path,
		Sync:       Sync,
	}
}

// EcdsaKeyPairFactory is just a simple function to instantiate the EcdsaKeyPairStruct
func EcdsaKeyPairFactory(
	KeySize *int,
	SignatureAlgorithm string,
) EcdsaKeyPairStruct {
	return EcdsaKeyPairStruct{
		Type:               "EcdsaKeyPair",
		KeySize:            KeySize,
		SignatureAlgorithm: SignatureAlgorithm,
	}
}

// EndEntityCertificateFactory is just a simple function to instantiate the EndEntityCertificateStruct
func EndEntityCertificateFactory(
	EndEntity EndEntity,
	IsCertificateAuthority *bool,
	IssuedByDN string,
	Issuer string,
	Md5Fingerprint string,
	Name string,
	Namespace string,
	NotAfter string,
	NotBefore string,
	Reference string,
	SerialNumber string,
	Sha1Fingerprint string,
	SubjectAlternativeNames []string,
) EndEntityCertificateStruct {
	return EndEntityCertificateStruct{
		Type:                    "EndEntityCertificate",
		EndEntity:               EndEntity,
		IsCertificateAuthority:  IsCertificateAuthority,
		IssuedByDN:              IssuedByDN,
		Issuer:                  Issuer,
		Md5Fingerprint:          Md5Fingerprint,
		Name:                    Name,
		Namespace:               Namespace,
		NotAfter:                NotAfter,
		NotBefore:               NotBefore,
		Reference:               Reference,
		SerialNumber:            SerialNumber,
		Sha1Fingerprint:         Sha1Fingerprint,
		SubjectAlternativeNames: SubjectAlternativeNames,
	}
}

// EndEntityCertificateReplaceChainParametersFactory is just a simple function to instantiate the EndEntityCertificateReplaceChainParametersStruct
func EndEntityCertificateReplaceChainParametersFactory(
	Chain *PemCertificateChainStruct,
	EndEntity EndEntity,
) EndEntityCertificateReplaceChainParametersStruct {
	return EndEntityCertificateReplaceChainParametersStruct{
		Type:      "EndEntityCertificateReplaceChainParameters",
		Chain:     Chain,
		EndEntity: EndEntity,
	}
}

// EndEntityCertificateReplaceKeystoreParametersFactory is just a simple function to instantiate the EndEntityCertificateReplaceKeystoreParametersStruct
func EndEntityCertificateReplaceKeystoreParametersFactory(
	EndEntity EndEntity,
	Token string,
) EndEntityCertificateReplaceKeystoreParametersStruct {
	return EndEntityCertificateReplaceKeystoreParametersStruct{
		Type:      "EndEntityCertificateReplaceKeystoreParameters",
		EndEntity: EndEntity,
		Token:     Token,
	}
}

// EndEntityDspFactory is just a simple function to instantiate the EndEntityDspStruct
func EndEntityDspFactory() EndEntityDspStruct {
	return EndEntityDspStruct{
		Type: "EndEntityDsp",
	}
}

// EndEntityHttpsFactory is just a simple function to instantiate the EndEntityHttpsStruct
func EndEntityHttpsFactory() EndEntityHttpsStruct {
	return EndEntityHttpsStruct{
		Type: "EndEntityHttps",
	}
}

// EngineAggregateIngestedSizeFactory is just a simple function to instantiate the EngineAggregateIngestedSizeStruct
func EngineAggregateIngestedSizeFactory() EngineAggregateIngestedSizeStruct {
	return EngineAggregateIngestedSizeStruct{
		Type: "EngineAggregateIngestedSize",
	}
}

// EnumEqualConstraintFactory is just a simple function to instantiate the EnumEqualConstraintStruct
func EnumEqualConstraintFactory(
	AxisName string,
	Equals string,
) EnumEqualConstraintStruct {
	return EnumEqualConstraintStruct{
		Type:     "EnumEqualConstraint",
		AxisName: AxisName,
		Equals:   Equals,
	}
}

// EnvironmentUserFactory is just a simple function to instantiate the EnvironmentUserStruct
func EnvironmentUserFactory(
	Credential Credential,
	Environment string,
	GroupId *int,
	Name string,
	Namespace string,
	Reference string,
	UserId *int,
) EnvironmentUserStruct {
	return EnvironmentUserStruct{
		Type:        "EnvironmentUser",
		Credential:  Credential,
		Environment: Environment,
		GroupId:     GroupId,
		Name:        Name,
		Namespace:   Namespace,
		Reference:   Reference,
		UserId:      UserId,
	}
}

// ErrorResultFactory is just a simple function to instantiate the ErrorResultStruct
func ErrorResultFactory(
	Error *APIErrorStruct,
	Status string,
) ErrorResultStruct {
	return ErrorResultStruct{
		Type:   "ErrorResult",
		Error:  Error,
		Status: Status,
	}
}

// EventFilterFactory is just a simple function to instantiate the EventFilterStruct
func EventFilterFactory(
	EventTypes []string,
) EventFilterStruct {
	return EventFilterStruct{
		Type:       "EventFilter",
		EventTypes: EventTypes,
	}
}

// ExternalBackupIngestionStrategyFactory is just a simple function to instantiate the ExternalBackupIngestionStrategyStruct
func ExternalBackupIngestionStrategyFactory(
	ValidatedSyncMode string,
) ExternalBackupIngestionStrategyStruct {
	return ExternalBackupIngestionStrategyStruct{
		Type:              "ExternalBackupIngestionStrategy",
		ValidatedSyncMode: ValidatedSyncMode,
	}
}

// FaultFactory is just a simple function to instantiate the FaultStruct
func FaultFactory(
	Action string,
	BundleID string,
	DateDiagnosed string,
	DateResolved string,
	Description string,
	Namespace string,
	Reference string,
	ResolutionComments string,
	Response string,
	Severity string,
	Status string,
	Target string,
	TargetName string,
	TargetObjectType string,
	Title string,
) FaultStruct {
	return FaultStruct{
		Type:               "Fault",
		Action:             Action,
		BundleID:           BundleID,
		DateDiagnosed:      DateDiagnosed,
		DateResolved:       DateResolved,
		Description:        Description,
		Namespace:          Namespace,
		Reference:          Reference,
		ResolutionComments: ResolutionComments,
		Response:           Response,
		Severity:           Severity,
		Status:             Status,
		Target:             Target,
		TargetName:         TargetName,
		TargetObjectType:   TargetObjectType,
		Title:              Title,
	}
}

// FaultEffectFactory is just a simple function to instantiate the FaultEffectStruct
func FaultEffectFactory(
	Action string,
	BundleID string,
	CausedBy string,
	DateDiagnosed string,
	Description string,
	Namespace string,
	Reference string,
	Response string,
	RootCause string,
	Severity string,
	Target string,
	TargetName string,
	Title string,
) FaultEffectStruct {
	return FaultEffectStruct{
		Type:          "FaultEffect",
		Action:        Action,
		BundleID:      BundleID,
		CausedBy:      CausedBy,
		DateDiagnosed: DateDiagnosed,
		Description:   Description,
		Namespace:     Namespace,
		Reference:     Reference,
		Response:      Response,
		RootCause:     RootCause,
		Severity:      Severity,
		Target:        Target,
		TargetName:    TargetName,
		Title:         Title,
	}
}

// FaultResolveParametersFactory is just a simple function to instantiate the FaultResolveParametersStruct
func FaultResolveParametersFactory(
	Comments string,
	Ignore *bool,
) FaultResolveParametersStruct {
	return FaultResolveParametersStruct{
		Type:     "FaultResolveParameters",
		Comments: Comments,
		Ignore:   Ignore,
	}
}

// FeatureFlagParametersFactory is just a simple function to instantiate the FeatureFlagParametersStruct
func FeatureFlagParametersFactory(
	Name string,
) FeatureFlagParametersStruct {
	return FeatureFlagParametersStruct{
		Type: "FeatureFlagParameters",
		Name: Name,
	}
}

// FileDownloadResultFactory is just a simple function to instantiate the FileDownloadResultStruct
func FileDownloadResultFactory(
	Token string,
	Url string,
) FileDownloadResultStruct {
	return FileDownloadResultStruct{
		Type:  "FileDownloadResult",
		Token: Token,
		Url:   Url,
	}
}

// FileMappingParametersFactory is just a simple function to instantiate the FileMappingParametersStruct
func FileMappingParametersFactory(
	MappingRules string,
	TimeflowPointParameters []TimeflowPointParameters,
) FileMappingParametersStruct {
	return FileMappingParametersStruct{
		Type:                    "FileMappingParameters",
		MappingRules:            MappingRules,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// FileMappingResultFactory is just a simple function to instantiate the FileMappingResultStruct
func FileMappingResultFactory(
	MappedFiles map[string]string,
) FileMappingResultStruct {
	return FileMappingResultStruct{
		Type:        "FileMappingResult",
		MappedFiles: MappedFiles,
	}
}

// FileUploadResultFactory is just a simple function to instantiate the FileUploadResultStruct
func FileUploadResultFactory(
	Token string,
	Url string,
) FileUploadResultStruct {
	return FileUploadResultStruct{
		Type:  "FileUploadResult",
		Token: Token,
		Url:   Url,
	}
}

// GetCurrentPasswordPolicyParametersFactory is just a simple function to instantiate the GetCurrentPasswordPolicyParametersStruct
func GetCurrentPasswordPolicyParametersFactory(
	UserType string,
) GetCurrentPasswordPolicyParametersStruct {
	return GetCurrentPasswordPolicyParametersStruct{
		Type:     "GetCurrentPasswordPolicyParameters",
		UserType: UserType,
	}
}

// GlobalLinkingSettingsFactory is just a simple function to instantiate the GlobalLinkingSettingsStruct
func GlobalLinkingSettingsFactory(
	EncryptedLinkingEnabledByDefault *bool,
) GlobalLinkingSettingsStruct {
	return GlobalLinkingSettingsStruct{
		Type:                             "GlobalLinkingSettings",
		EncryptedLinkingEnabledByDefault: EncryptedLinkingEnabledByDefault,
	}
}

// GroupFactory is just a simple function to instantiate the GroupStruct
func GroupFactory(
	Description string,
	Name string,
	Namespace string,
	Reference string,
) GroupStruct {
	return GroupStruct{
		Type:        "Group",
		Description: Description,
		Name:        Name,
		Namespace:   Namespace,
		Reference:   Reference,
	}
}

// HashiCorpAppRoleAuthenticationFactory is just a simple function to instantiate the HashiCorpAppRoleAuthenticationStruct
func HashiCorpAppRoleAuthenticationFactory(
	RoleId string,
	SecretId string,
) HashiCorpAppRoleAuthenticationStruct {
	return HashiCorpAppRoleAuthenticationStruct{
		Type:     "HashiCorpAppRoleAuthentication",
		RoleId:   RoleId,
		SecretId: SecretId,
	}
}

// HashiCorpCertificateAuthenticationFactory is just a simple function to instantiate the HashiCorpCertificateAuthenticationStruct
func HashiCorpCertificateAuthenticationFactory(
	ClientCertificate ClientCertificate,
	RoleName string,
) HashiCorpCertificateAuthenticationStruct {
	return HashiCorpCertificateAuthenticationStruct{
		Type:              "HashiCorpCertificateAuthentication",
		ClientCertificate: ClientCertificate,
		RoleName:          RoleName,
	}
}

// HashiCorpTokenAuthenticationFactory is just a simple function to instantiate the HashiCorpTokenAuthenticationStruct
func HashiCorpTokenAuthenticationFactory(
	Token string,
) HashiCorpTokenAuthenticationStruct {
	return HashiCorpTokenAuthenticationStruct{
		Type:  "HashiCorpTokenAuthentication",
		Token: Token,
	}
}

// HashiCorpVaultFactory is just a simple function to instantiate the HashiCorpVaultStruct
func HashiCorpVaultFactory(
	Authentication HashiCorpAuthentication,
	Host string,
	Name string,
	Namespace string,
	Port *int,
	Reference string,
	VaultNamespace string,
) HashiCorpVaultStruct {
	return HashiCorpVaultStruct{
		Type:           "HashiCorpVault",
		Authentication: Authentication,
		Host:           Host,
		Name:           Name,
		Namespace:      Namespace,
		Port:           Port,
		Reference:      Reference,
		VaultNamespace: VaultNamespace,
	}
}

// HashiCorpVaultCredentialFactory is just a simple function to instantiate the HashiCorpVaultCredentialStruct
func HashiCorpVaultCredentialFactory(
	Engine string,
	Path string,
	SecretKey string,
	UsernameKey string,
	Vault string,
) HashiCorpVaultCredentialStruct {
	return HashiCorpVaultCredentialStruct{
		Type:        "HashiCorpVaultCredential",
		Engine:      Engine,
		Path:        Path,
		SecretKey:   SecretKey,
		UsernameKey: UsernameKey,
		Vault:       Vault,
	}
}

// HashiCorpVaultTestParametersFactory is just a simple function to instantiate the HashiCorpVaultTestParametersStruct
func HashiCorpVaultTestParametersFactory(
	Authentication HashiCorpAuthentication,
	Host string,
	Port *int,
	VaultNamespace string,
) HashiCorpVaultTestParametersStruct {
	return HashiCorpVaultTestParametersStruct{
		Type:           "HashiCorpVaultTestParameters",
		Authentication: Authentication,
		Host:           Host,
		Port:           Port,
		VaultNamespace: VaultNamespace,
	}
}

// HeldSpaceCapacityDataFactory is just a simple function to instantiate the HeldSpaceCapacityDataStruct
func HeldSpaceCapacityDataFactory(
	Breakdown *CapacityBreakdownStruct,
	StorageContainer string,
) HeldSpaceCapacityDataStruct {
	return HeldSpaceCapacityDataStruct{
		Type:             "HeldSpaceCapacityData",
		Breakdown:        Breakdown,
		StorageContainer: StorageContainer,
	}
}

// HistoricalConsumerCapacityDataFactory is just a simple function to instantiate the HistoricalConsumerCapacityDataStruct
func HistoricalConsumerCapacityDataFactory(
	Breakdown *CapacityBreakdownStruct,
	Container string,
	Group string,
	GroupName string,
	Name string,
	Parent string,
	Timestamp string,
) HistoricalConsumerCapacityDataStruct {
	return HistoricalConsumerCapacityDataStruct{
		Type:      "HistoricalConsumerCapacityData",
		Breakdown: Breakdown,
		Container: Container,
		Group:     Group,
		GroupName: GroupName,
		Name:      Name,
		Parent:    Parent,
		Timestamp: Timestamp,
	}
}

// HistoricalGroupCapacityDataFactory is just a simple function to instantiate the HistoricalGroupCapacityDataStruct
func HistoricalGroupCapacityDataFactory(
	Group string,
	Source *CapacityBreakdownStruct,
	Timestamp string,
	Virtual *CapacityBreakdownStruct,
) HistoricalGroupCapacityDataStruct {
	return HistoricalGroupCapacityDataStruct{
		Type:      "HistoricalGroupCapacityData",
		Group:     Group,
		Source:    Source,
		Timestamp: Timestamp,
		Virtual:   Virtual,
	}
}

// HistoricalSystemCapacityDataFactory is just a simple function to instantiate the HistoricalSystemCapacityDataStruct
func HistoricalSystemCapacityDataFactory(
	Source *CapacityBreakdownStruct,
	Timestamp string,
	TotalSpace float64,
	Virtual *CapacityBreakdownStruct,
) HistoricalSystemCapacityDataStruct {
	return HistoricalSystemCapacityDataStruct{
		Type:       "HistoricalSystemCapacityData",
		Source:     Source,
		Timestamp:  Timestamp,
		TotalSpace: TotalSpace,
		Virtual:    Virtual,
	}
}

// HostConfigurationFactory is just a simple function to instantiate the HostConfigurationStruct
func HostConfigurationFactory(
	Discovered *bool,
	LastRefreshed string,
	LastUpdated string,
	Machine *HostMachineStruct,
	OperatingSystem *HostOSStruct,
	PowerShellVersion string,
) HostConfigurationStruct {
	return HostConfigurationStruct{
		Type:              "HostConfiguration",
		Discovered:        Discovered,
		LastRefreshed:     LastRefreshed,
		LastUpdated:       LastUpdated,
		Machine:           Machine,
		OperatingSystem:   OperatingSystem,
		PowerShellVersion: PowerShellVersion,
	}
}

// HostEnvironmentCreateParametersFactory is just a simple function to instantiate the HostEnvironmentCreateParametersStruct
func HostEnvironmentCreateParametersFactory(
	HostEnvironment HostEnvironment,
	HostParameters HostCreateParameters,
	LogCollectionEnabled *bool,
	PrimaryUser *EnvironmentUserStruct,
) HostEnvironmentCreateParametersStruct {
	return HostEnvironmentCreateParametersStruct{
		Type:                 "HostEnvironmentCreateParameters",
		HostEnvironment:      HostEnvironment,
		HostParameters:       HostParameters,
		LogCollectionEnabled: LogCollectionEnabled,
		PrimaryUser:          PrimaryUser,
	}
}

// HostMachineFactory is just a simple function to instantiate the HostMachineStruct
func HostMachineFactory(
	MemorySize float64,
	Platform string,
) HostMachineStruct {
	return HostMachineStruct{
		Type:       "HostMachine",
		MemorySize: MemorySize,
		Platform:   Platform,
	}
}

// HostOSFactory is just a simple function to instantiate the HostOSStruct
func HostOSFactory(
	Distribution string,
	Kernel string,
	Name string,
	Release string,
	Timezone string,
	Version string,
) HostOSStruct {
	return HostOSStruct{
		Type:         "HostOS",
		Distribution: Distribution,
		Kernel:       Kernel,
		Name:         Name,
		Release:      Release,
		Timezone:     Timezone,
		Version:      Version,
	}
}

// HostPrivilegeElevationProfileFactory is just a simple function to instantiate the HostPrivilegeElevationProfileStruct
func HostPrivilegeElevationProfileFactory(
	IsDefault *bool,
	Name string,
	Namespace string,
	Reference string,
	Version string,
) HostPrivilegeElevationProfileStruct {
	return HostPrivilegeElevationProfileStruct{
		Type:      "HostPrivilegeElevationProfile",
		IsDefault: IsDefault,
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
		Version:   Version,
	}
}

// HostPrivilegeElevationProfileScriptFactory is just a simple function to instantiate the HostPrivilegeElevationProfileScriptStruct
func HostPrivilegeElevationProfileScriptFactory(
	Contents string,
	Name string,
	Namespace string,
	Profile string,
	Reference string,
) HostPrivilegeElevationProfileScriptStruct {
	return HostPrivilegeElevationProfileScriptStruct{
		Type:      "HostPrivilegeElevationProfileScript",
		Contents:  Contents,
		Name:      Name,
		Namespace: Namespace,
		Profile:   Profile,
		Reference: Reference,
	}
}

// HostPrivilegeElevationSettingsFactory is just a simple function to instantiate the HostPrivilegeElevationSettingsStruct
func HostPrivilegeElevationSettingsFactory(
	DefaultProfile string,
) HostPrivilegeElevationSettingsStruct {
	return HostPrivilegeElevationSettingsStruct{
		Type:           "HostPrivilegeElevationSettings",
		DefaultProfile: DefaultProfile,
	}
}

// HostRuntimeFactory is just a simple function to instantiate the HostRuntimeStruct
func HostRuntimeFactory(
	Available *bool,
	AvailableTimestamp string,
	NotAvailableReason string,
	TraceRouteInfo *TracerouteInfoStruct,
) HostRuntimeStruct {
	return HostRuntimeStruct{
		Type:               "HostRuntime",
		Available:          Available,
		AvailableTimestamp: AvailableTimestamp,
		NotAvailableReason: NotAvailableReason,
		TraceRouteInfo:     TraceRouteInfo,
	}
}

// HttpConnectorConfigFactory is just a simple function to instantiate the HttpConnectorConfigStruct
func HttpConnectorConfigFactory(
	CipherSuites []*CipherSuiteStruct,
	HttpMode string,
	HttpPort *int,
	HttpsPort *int,
	MaskingLegacyPorts *bool,
	TlsVersions []string,
) HttpConnectorConfigStruct {
	return HttpConnectorConfigStruct{
		Type:               "HttpConnectorConfig",
		CipherSuites:       CipherSuites,
		HttpMode:           HttpMode,
		HttpPort:           HttpPort,
		HttpsPort:          HttpsPort,
		MaskingLegacyPorts: MaskingLegacyPorts,
		TlsVersions:        TlsVersions,
	}
}

// IScsiOpsDatapointStreamFactory is just a simple function to instantiate the IScsiOpsDatapointStreamStruct
func IScsiOpsDatapointStreamFactory(
	Client string,
	Datapoints []Datapoint,
	Op string,
) IScsiOpsDatapointStreamStruct {
	return IScsiOpsDatapointStreamStruct{
		Type:       "IScsiOpsDatapointStream",
		Client:     Client,
		Datapoints: Datapoints,
		Op:         Op,
	}
}

// IntegerEqualConstraintFactory is just a simple function to instantiate the IntegerEqualConstraintStruct
func IntegerEqualConstraintFactory(
	AxisName string,
	Equals *int,
) IntegerEqualConstraintStruct {
	return IntegerEqualConstraintStruct{
		Type:     "IntegerEqualConstraint",
		AxisName: AxisName,
		Equals:   Equals,
	}
}

// IntegerGreaterThanConstraintFactory is just a simple function to instantiate the IntegerGreaterThanConstraintStruct
func IntegerGreaterThanConstraintFactory(
	AxisName string,
	GreaterThan *int,
) IntegerGreaterThanConstraintStruct {
	return IntegerGreaterThanConstraintStruct{
		Type:        "IntegerGreaterThanConstraint",
		AxisName:    AxisName,
		GreaterThan: GreaterThan,
	}
}

// IntegerLessThanConstraintFactory is just a simple function to instantiate the IntegerLessThanConstraintStruct
func IntegerLessThanConstraintFactory(
	AxisName string,
	LessThan *int,
) IntegerLessThanConstraintStruct {
	return IntegerLessThanConstraintStruct{
		Type:     "IntegerLessThanConstraint",
		AxisName: AxisName,
		LessThan: LessThan,
	}
}

// InterfaceAddressFactory is just a simple function to instantiate the InterfaceAddressStruct
func InterfaceAddressFactory(
	Address string,
	AddressType string,
	EnableSSH *bool,
	SessionInUse *bool,
	State string,
) InterfaceAddressStruct {
	return InterfaceAddressStruct{
		Type:         "InterfaceAddress",
		Address:      Address,
		AddressType:  AddressType,
		EnableSSH:    EnableSSH,
		SessionInUse: SessionInUse,
		State:        State,
	}
}

// IoOpsDatapointFactory is just a simple function to instantiate the IoOpsDatapointStruct
func IoOpsDatapointFactory(
	AvgLatency *int,
	Count *int,
	Latency map[string]string,
	Size map[string]string,
	Throughput *int,
	Timestamp string,
) IoOpsDatapointStruct {
	return IoOpsDatapointStruct{
		Type:       "IoOpsDatapoint",
		AvgLatency: AvgLatency,
		Count:      Count,
		Latency:    Latency,
		Size:       Size,
		Throughput: Throughput,
		Timestamp:  Timestamp,
	}
}

// IscsiInitiatorFactory is just a simple function to instantiate the IscsiInitiatorStruct
func IscsiInitiatorFactory(
	InitiatorName string,
) IscsiInitiatorStruct {
	return IscsiInitiatorStruct{
		Type:          "IscsiInitiator",
		InitiatorName: InitiatorName,
	}
}

// IscsiTargetFactory is just a simple function to instantiate the IscsiTargetStruct
func IscsiTargetFactory(
	ChapPassword string,
	ChapPasswordIn string,
	ChapUsername string,
	ChapUsernameIn string,
	Iqn string,
	Name string,
	Namespace string,
	Port *int,
	Portal string,
	Reference string,
	State string,
) IscsiTargetStruct {
	return IscsiTargetStruct{
		Type:           "IscsiTarget",
		ChapPassword:   ChapPassword,
		ChapPasswordIn: ChapPasswordIn,
		ChapUsername:   ChapUsername,
		ChapUsernameIn: ChapUsernameIn,
		Iqn:            Iqn,
		Name:           Name,
		Namespace:      Namespace,
		Port:           Port,
		Portal:         Portal,
		Reference:      Reference,
		State:          State,
	}
}

// IscsiTargetDiscoverParametersFactory is just a simple function to instantiate the IscsiTargetDiscoverParametersStruct
func IscsiTargetDiscoverParametersFactory(
	ChapPassword string,
	ChapPasswordIn string,
	ChapUsername string,
	ChapUsernameIn string,
	Port *int,
	Portal string,
) IscsiTargetDiscoverParametersStruct {
	return IscsiTargetDiscoverParametersStruct{
		Type:           "IscsiTargetDiscoverParameters",
		ChapPassword:   ChapPassword,
		ChapPasswordIn: ChapPasswordIn,
		ChapUsername:   ChapUsername,
		ChapUsernameIn: ChapUsernameIn,
		Port:           Port,
		Portal:         Portal,
	}
}

// JDBCConnectivityFactory is just a simple function to instantiate the JDBCConnectivityStruct
func JDBCConnectivityFactory(
	Password string,
	Url string,
	User string,
) JDBCConnectivityStruct {
	return JDBCConnectivityStruct{
		Type:     "JDBCConnectivity",
		Password: Password,
		Url:      Url,
		User:     User,
	}
}

// JSBookmarkFactory is just a simple function to instantiate the JSBookmarkStruct
func JSBookmarkFactory(
	BookmarkType string,
	Branch string,
	CheckoutCount *int,
	Container string,
	ContainerName string,
	CreationTime string,
	Description string,
	Expiration string,
	Name string,
	Namespace string,
	Reference string,
	Shared *bool,
	Tags []string,
	Template string,
	TemplateName string,
	Timestamp string,
	Usable *bool,
) JSBookmarkStruct {
	return JSBookmarkStruct{
		Type:          "JSBookmark",
		BookmarkType:  BookmarkType,
		Branch:        Branch,
		CheckoutCount: CheckoutCount,
		Container:     Container,
		ContainerName: ContainerName,
		CreationTime:  CreationTime,
		Description:   Description,
		Expiration:    Expiration,
		Name:          Name,
		Namespace:     Namespace,
		Reference:     Reference,
		Shared:        Shared,
		Tags:          Tags,
		Template:      Template,
		TemplateName:  TemplateName,
		Timestamp:     Timestamp,
		Usable:        Usable,
	}
}

// JSBookmarkCheckoutCountFactory is just a simple function to instantiate the JSBookmarkCheckoutCountStruct
func JSBookmarkCheckoutCountFactory(
	Bookmark string,
	CheckoutCount *int,
	Namespace string,
	Reference string,
) JSBookmarkCheckoutCountStruct {
	return JSBookmarkCheckoutCountStruct{
		Type:          "JSBookmarkCheckoutCount",
		Bookmark:      Bookmark,
		CheckoutCount: CheckoutCount,
		Namespace:     Namespace,
		Reference:     Reference,
	}
}

// JSBookmarkCreateParametersFactory is just a simple function to instantiate the JSBookmarkCreateParametersStruct
func JSBookmarkCreateParametersFactory(
	Bookmark *JSBookmarkStruct,
	TimelinePointParameters JSTimelinePointTimeParameters,
) JSBookmarkCreateParametersStruct {
	return JSBookmarkCreateParametersStruct{
		Type:                    "JSBookmarkCreateParameters",
		Bookmark:                Bookmark,
		TimelinePointParameters: TimelinePointParameters,
	}
}

// JSBookmarkDataParentFactory is just a simple function to instantiate the JSBookmarkDataParentStruct
func JSBookmarkDataParentFactory(
	Bookmark string,
	BookmarkName string,
) JSBookmarkDataParentStruct {
	return JSBookmarkDataParentStruct{
		Type:         "JSBookmarkDataParent",
		Bookmark:     Bookmark,
		BookmarkName: BookmarkName,
	}
}

// JSBookmarkTagUsageDataFactory is just a simple function to instantiate the JSBookmarkTagUsageDataStruct
func JSBookmarkTagUsageDataFactory(
	BookmarkTag string,
	Referenced float64,
	Unique float64,
) JSBookmarkTagUsageDataStruct {
	return JSBookmarkTagUsageDataStruct{
		Type:        "JSBookmarkTagUsageData",
		BookmarkTag: BookmarkTag,
		Referenced:  Referenced,
		Unique:      Unique,
	}
}

// JSBookmarkUsageDataFactory is just a simple function to instantiate the JSBookmarkUsageDataStruct
func JSBookmarkUsageDataFactory(
	Bookmark string,
	DataLayout string,
	ExternallyReferenced float64,
	Shared float64,
	Unique float64,
) JSBookmarkUsageDataStruct {
	return JSBookmarkUsageDataStruct{
		Type:                 "JSBookmarkUsageData",
		Bookmark:             Bookmark,
		DataLayout:           DataLayout,
		ExternallyReferenced: ExternallyReferenced,
		Shared:               Shared,
		Unique:               Unique,
	}
}

// JSBranchFactory is just a simple function to instantiate the JSBranchStruct
func JSBranchFactory(
	DataLayout string,
	FirstOperation string,
	LastOperation string,
	Name string,
	Namespace string,
	Reference string,
) JSBranchStruct {
	return JSBranchStruct{
		Type:           "JSBranch",
		DataLayout:     DataLayout,
		FirstOperation: FirstOperation,
		LastOperation:  LastOperation,
		Name:           Name,
		Namespace:      Namespace,
		Reference:      Reference,
	}
}

// JSBranchCreateParametersFactory is just a simple function to instantiate the JSBranchCreateParametersStruct
func JSBranchCreateParametersFactory(
	DataContainer string,
	Name string,
	TimelinePointParameters JSTimelinePointParameters,
) JSBranchCreateParametersStruct {
	return JSBranchCreateParametersStruct{
		Type:                    "JSBranchCreateParameters",
		DataContainer:           DataContainer,
		Name:                    Name,
		TimelinePointParameters: TimelinePointParameters,
	}
}

// JSBranchUsageDataFactory is just a simple function to instantiate the JSBranchUsageDataStruct
func JSBranchUsageDataFactory(
	Branch string,
	DataContainer string,
	SharedOthers float64,
	SharedSelf float64,
	Unique float64,
) JSBranchUsageDataStruct {
	return JSBranchUsageDataStruct{
		Type:          "JSBranchUsageData",
		Branch:        Branch,
		DataContainer: DataContainer,
		SharedOthers:  SharedOthers,
		SharedSelf:    SharedSelf,
		Unique:        Unique,
	}
}

// JSConfigFactory is just a simple function to instantiate the JSConfigStruct
func JSConfigFactory(
	DefaultBookmarkExpiration *int,
	RetryAttempts *int,
) JSConfigStruct {
	return JSConfigStruct{
		Type:                      "JSConfig",
		DefaultBookmarkExpiration: DefaultBookmarkExpiration,
		RetryAttempts:             RetryAttempts,
	}
}

// JSContainerUsageDataFactory is just a simple function to instantiate the JSContainerUsageDataStruct
func JSContainerUsageDataFactory(
	DataContainer string,
	SharedOthers float64,
	SharedSelf float64,
	Unique float64,
	Unvirtualized float64,
) JSContainerUsageDataStruct {
	return JSContainerUsageDataStruct{
		Type:          "JSContainerUsageData",
		DataContainer: DataContainer,
		SharedOthers:  SharedOthers,
		SharedSelf:    SharedSelf,
		Unique:        Unique,
		Unvirtualized: Unvirtualized,
	}
}

// JSDailyOperationDurationFactory is just a simple function to instantiate the JSDailyOperationDurationStruct
func JSDailyOperationDurationFactory(
	DailyAverageDuration *int,
	DailyCount *int,
	DailyMaxDuration *int,
	DailyMinDuration *int,
	Namespace string,
	Operation string,
	Reference string,
	StartDate string,
	UsageObject string,
) JSDailyOperationDurationStruct {
	return JSDailyOperationDurationStruct{
		Type:                 "JSDailyOperationDuration",
		DailyAverageDuration: DailyAverageDuration,
		DailyCount:           DailyCount,
		DailyMaxDuration:     DailyMaxDuration,
		DailyMinDuration:     DailyMinDuration,
		Namespace:            Namespace,
		Operation:            Operation,
		Reference:            Reference,
		StartDate:            StartDate,
		UsageObject:          UsageObject,
	}
}

// JSDataChildFactory is just a simple function to instantiate the JSDataChildStruct
func JSDataChildFactory(
	Branch string,
	BranchName string,
	ContainerName string,
	Operation string,
	UserName string,
) JSDataChildStruct {
	return JSDataChildStruct{
		Type:          "JSDataChild",
		Branch:        Branch,
		BranchName:    BranchName,
		ContainerName: ContainerName,
		Operation:     Operation,
		UserName:      UserName,
	}
}

// JSDataContainerFactory is just a simple function to instantiate the JSDataContainerStruct
func JSDataContainerFactory(
	ActiveBranch string,
	FirstOperation string,
	LastOperation string,
	LastUpdated string,
	LockUserName string,
	LockUserReference string,
	Name string,
	Namespace string,
	Notes string,
	OperationCount *int,
	Owner string,
	Properties map[string]string,
	Reference string,
	State string,
	Template string,
) JSDataContainerStruct {
	return JSDataContainerStruct{
		Type:              "JSDataContainer",
		ActiveBranch:      ActiveBranch,
		FirstOperation:    FirstOperation,
		LastOperation:     LastOperation,
		LastUpdated:       LastUpdated,
		LockUserName:      LockUserName,
		LockUserReference: LockUserReference,
		Name:              Name,
		Namespace:         Namespace,
		Notes:             Notes,
		OperationCount:    OperationCount,
		Owner:             Owner,
		Properties:        Properties,
		Reference:         Reference,
		State:             State,
		Template:          Template,
	}
}

// JSDataContainerActiveBranchParametersFactory is just a simple function to instantiate the JSDataContainerActiveBranchParametersStruct
func JSDataContainerActiveBranchParametersFactory(
	Time string,
) JSDataContainerActiveBranchParametersStruct {
	return JSDataContainerActiveBranchParametersStruct{
		Type: "JSDataContainerActiveBranchParameters",
		Time: Time,
	}
}

// JSDataContainerCreateWithRefreshParametersFactory is just a simple function to instantiate the JSDataContainerCreateWithRefreshParametersStruct
func JSDataContainerCreateWithRefreshParametersFactory(
	DataSources []*JSDataSourceCreateParametersStruct,
	Name string,
	Notes string,
	Owners []string,
	Properties map[string]string,
	Template string,
	TimelinePointParameters JSTimelinePointParameters,
) JSDataContainerCreateWithRefreshParametersStruct {
	return JSDataContainerCreateWithRefreshParametersStruct{
		Type:                    "JSDataContainerCreateWithRefreshParameters",
		DataSources:             DataSources,
		Name:                    Name,
		Notes:                   Notes,
		Owners:                  Owners,
		Properties:              Properties,
		Template:                Template,
		TimelinePointParameters: TimelinePointParameters,
	}
}

// JSDataContainerCreateWithoutRefreshParametersFactory is just a simple function to instantiate the JSDataContainerCreateWithoutRefreshParametersStruct
func JSDataContainerCreateWithoutRefreshParametersFactory(
	DataSources []*JSDataSourceCreateParametersStruct,
	Name string,
	Notes string,
	Owners []string,
	Properties map[string]string,
	Template string,
) JSDataContainerCreateWithoutRefreshParametersStruct {
	return JSDataContainerCreateWithoutRefreshParametersStruct{
		Type:        "JSDataContainerCreateWithoutRefreshParameters",
		DataSources: DataSources,
		Name:        Name,
		Notes:       Notes,
		Owners:      Owners,
		Properties:  Properties,
		Template:    Template,
	}
}

// JSDataContainerDeleteParametersFactory is just a simple function to instantiate the JSDataContainerDeleteParametersStruct
func JSDataContainerDeleteParametersFactory(
	DeleteDataSources *bool,
) JSDataContainerDeleteParametersStruct {
	return JSDataContainerDeleteParametersStruct{
		Type:              "JSDataContainerDeleteParameters",
		DeleteDataSources: DeleteDataSources,
	}
}

// JSDataContainerLockParametersFactory is just a simple function to instantiate the JSDataContainerLockParametersStruct
func JSDataContainerLockParametersFactory(
	LockUser string,
) JSDataContainerLockParametersStruct {
	return JSDataContainerLockParametersStruct{
		Type:     "JSDataContainerLockParameters",
		LockUser: LockUser,
	}
}

// JSDataContainerModifyOwnerParametersFactory is just a simple function to instantiate the JSDataContainerModifyOwnerParametersStruct
func JSDataContainerModifyOwnerParametersFactory(
	Owner string,
) JSDataContainerModifyOwnerParametersStruct {
	return JSDataContainerModifyOwnerParametersStruct{
		Type:  "JSDataContainerModifyOwnerParameters",
		Owner: Owner,
	}
}

// JSDataContainerRefreshParametersFactory is just a simple function to instantiate the JSDataContainerRefreshParametersStruct
func JSDataContainerRefreshParametersFactory(
	ForceOption *bool,
) JSDataContainerRefreshParametersStruct {
	return JSDataContainerRefreshParametersStruct{
		Type:        "JSDataContainerRefreshParameters",
		ForceOption: ForceOption,
	}
}

// JSDataContainerResetParametersFactory is just a simple function to instantiate the JSDataContainerResetParametersStruct
func JSDataContainerResetParametersFactory(
	ForceOption *bool,
) JSDataContainerResetParametersStruct {
	return JSDataContainerResetParametersStruct{
		Type:        "JSDataContainerResetParameters",
		ForceOption: ForceOption,
	}
}

// JSDataContainerRestoreParametersFactory is just a simple function to instantiate the JSDataContainerRestoreParametersStruct
func JSDataContainerRestoreParametersFactory(
	ForceOption *bool,
	TimelinePointParameters JSTimelinePointParameters,
) JSDataContainerRestoreParametersStruct {
	return JSDataContainerRestoreParametersStruct{
		Type:                    "JSDataContainerRestoreParameters",
		ForceOption:             ForceOption,
		TimelinePointParameters: TimelinePointParameters,
	}
}

// JSDataContainerUndoParametersFactory is just a simple function to instantiate the JSDataContainerUndoParametersStruct
func JSDataContainerUndoParametersFactory(
	Operation string,
) JSDataContainerUndoParametersStruct {
	return JSDataContainerUndoParametersStruct{
		Type:      "JSDataContainerUndoParameters",
		Operation: Operation,
	}
}

// JSDataSourceFactory is just a simple function to instantiate the JSDataSourceStruct
func JSDataSourceFactory(
	Container string,
	DataLayout string,
	Description string,
	Enabled *bool,
	Masked *bool,
	Name string,
	Namespace string,
	Priority *int,
	Properties map[string]string,
	Reference string,
	Runtime SourceConnectionInfo,
) JSDataSourceStruct {
	return JSDataSourceStruct{
		Type:        "JSDataSource",
		Container:   Container,
		DataLayout:  DataLayout,
		Description: Description,
		Enabled:     Enabled,
		Masked:      Masked,
		Name:        Name,
		Namespace:   Namespace,
		Priority:    Priority,
		Properties:  Properties,
		Reference:   Reference,
		Runtime:     Runtime,
	}
}

// JSDataSourceCreateParametersFactory is just a simple function to instantiate the JSDataSourceCreateParametersStruct
func JSDataSourceCreateParametersFactory(
	Container string,
	Properties map[string]string,
	Source *JSDataSourceStruct,
) JSDataSourceCreateParametersStruct {
	return JSDataSourceCreateParametersStruct{
		Type:       "JSDataSourceCreateParameters",
		Container:  Container,
		Properties: Properties,
		Source:     Source,
	}
}

// JSDataTemplateFactory is just a simple function to instantiate the JSDataTemplateStruct
func JSDataTemplateFactory(
	ActiveBranch string,
	ConfirmTimeConsumingOperations *bool,
	FirstOperation string,
	LastOperation string,
	LastUpdated string,
	Name string,
	Namespace string,
	Notes string,
	Properties map[string]string,
	Reference string,
) JSDataTemplateStruct {
	return JSDataTemplateStruct{
		Type:                           "JSDataTemplate",
		ActiveBranch:                   ActiveBranch,
		ConfirmTimeConsumingOperations: ConfirmTimeConsumingOperations,
		FirstOperation:                 FirstOperation,
		LastOperation:                  LastOperation,
		LastUpdated:                    LastUpdated,
		Name:                           Name,
		Namespace:                      Namespace,
		Notes:                          Notes,
		Properties:                     Properties,
		Reference:                      Reference,
	}
}

// JSDataTemplateCreateParametersFactory is just a simple function to instantiate the JSDataTemplateCreateParametersStruct
func JSDataTemplateCreateParametersFactory(
	DataSources []*JSDataSourceCreateParametersStruct,
	Name string,
	Notes string,
	Properties map[string]string,
) JSDataTemplateCreateParametersStruct {
	return JSDataTemplateCreateParametersStruct{
		Type:        "JSDataTemplateCreateParameters",
		DataSources: DataSources,
		Name:        Name,
		Notes:       Notes,
		Properties:  Properties,
	}
}

// JSOperationFactory is just a simple function to instantiate the JSOperationStruct
func JSOperationFactory(
	Bookmark string,
	Branch string,
	DataLayout string,
	DataParent JSDataParent,
	DataTime string,
	Description string,
	EndTime string,
	ForceOption *bool,
	Name string,
	Namespace string,
	Operation string,
	Reference string,
	RootAction string,
	StartTime string,
	User string,
) JSOperationStruct {
	return JSOperationStruct{
		Type:        "JSOperation",
		Bookmark:    Bookmark,
		Branch:      Branch,
		DataLayout:  DataLayout,
		DataParent:  DataParent,
		DataTime:    DataTime,
		Description: Description,
		EndTime:     EndTime,
		ForceOption: ForceOption,
		Name:        Name,
		Namespace:   Namespace,
		Operation:   Operation,
		Reference:   Reference,
		RootAction:  RootAction,
		StartTime:   StartTime,
		User:        User,
	}
}

// JSOperationEndpointFactory is just a simple function to instantiate the JSOperationEndpointStruct
func JSOperationEndpointFactory(
	First string,
	Last string,
) JSOperationEndpointStruct {
	return JSOperationEndpointStruct{
		Type:  "JSOperationEndpoint",
		First: First,
		Last:  Last,
	}
}

// JSOperationEndpointBranchParametersFactory is just a simple function to instantiate the JSOperationEndpointBranchParametersStruct
func JSOperationEndpointBranchParametersFactory(
	Branch string,
) JSOperationEndpointBranchParametersStruct {
	return JSOperationEndpointBranchParametersStruct{
		Type:   "JSOperationEndpointBranchParameters",
		Branch: Branch,
	}
}

// JSOperationEndpointDataLayoutParametersFactory is just a simple function to instantiate the JSOperationEndpointDataLayoutParametersStruct
func JSOperationEndpointDataLayoutParametersFactory(
	DataLayout string,
) JSOperationEndpointDataLayoutParametersStruct {
	return JSOperationEndpointDataLayoutParametersStruct{
		Type:       "JSOperationEndpointDataLayoutParameters",
		DataLayout: DataLayout,
	}
}

// JSSourceDataTimestampFactory is just a simple function to instantiate the JSSourceDataTimestampStruct
func JSSourceDataTimestampFactory(
	Branch string,
	Name string,
	Priority *int,
	Source string,
	Timestamp string,
) JSSourceDataTimestampStruct {
	return JSSourceDataTimestampStruct{
		Type:      "JSSourceDataTimestamp",
		Branch:    Branch,
		Name:      Name,
		Priority:  Priority,
		Source:    Source,
		Timestamp: Timestamp,
	}
}

// JSSourceDataTimestampParametersFactory is just a simple function to instantiate the JSSourceDataTimestampParametersStruct
func JSSourceDataTimestampParametersFactory(
	Branch string,
	Time string,
) JSSourceDataTimestampParametersStruct {
	return JSSourceDataTimestampParametersStruct{
		Type:   "JSSourceDataTimestampParameters",
		Branch: Branch,
		Time:   Time,
	}
}

// JSTemplateUsageDataFactory is just a simple function to instantiate the JSTemplateUsageDataStruct
func JSTemplateUsageDataFactory(
	Bookmarks float64,
	Containers float64,
	Template string,
	Total float64,
	Unvirtualized float64,
) JSTemplateUsageDataStruct {
	return JSTemplateUsageDataStruct{
		Type:          "JSTemplateUsageData",
		Bookmarks:     Bookmarks,
		Containers:    Containers,
		Template:      Template,
		Total:         Total,
		Unvirtualized: Unvirtualized,
	}
}

// JSTimelinePointBookmarkInputFactory is just a simple function to instantiate the JSTimelinePointBookmarkInputStruct
func JSTimelinePointBookmarkInputFactory(
	Bookmark string,
) JSTimelinePointBookmarkInputStruct {
	return JSTimelinePointBookmarkInputStruct{
		Type:     "JSTimelinePointBookmarkInput",
		Bookmark: Bookmark,
	}
}

// JSTimelinePointLatestTimeInputFactory is just a simple function to instantiate the JSTimelinePointLatestTimeInputStruct
func JSTimelinePointLatestTimeInputFactory(
	SourceDataLayout string,
) JSTimelinePointLatestTimeInputStruct {
	return JSTimelinePointLatestTimeInputStruct{
		Type:             "JSTimelinePointLatestTimeInput",
		SourceDataLayout: SourceDataLayout,
	}
}

// JSTimelinePointTimeInputFactory is just a simple function to instantiate the JSTimelinePointTimeInputStruct
func JSTimelinePointTimeInputFactory(
	Branch string,
	Time string,
) JSTimelinePointTimeInputStruct {
	return JSTimelinePointTimeInputStruct{
		Type:   "JSTimelinePointTimeInput",
		Branch: Branch,
		Time:   Time,
	}
}

// JSTimestampDataParentFactory is just a simple function to instantiate the JSTimestampDataParentStruct
func JSTimestampDataParentFactory(
	Branch string,
	BranchName string,
	Time string,
) JSTimestampDataParentStruct {
	return JSTimestampDataParentStruct{
		Type:       "JSTimestampDataParent",
		Branch:     Branch,
		BranchName: BranchName,
		Time:       Time,
	}
}

// JSUserUsageDataFactory is just a simple function to instantiate the JSUserUsageDataStruct
func JSUserUsageDataFactory(
	NumContainers *int,
	Total float64,
	User string,
) JSUserUsageDataStruct {
	return JSUserUsageDataStruct{
		Type:          "JSUserUsageData",
		NumContainers: NumContainers,
		Total:         Total,
		User:          User,
	}
}

// JSWeeklyOperationCountFactory is just a simple function to instantiate the JSWeeklyOperationCountStruct
func JSWeeklyOperationCountFactory(
	Namespace string,
	Reference string,
	StartDate string,
	UsageObject string,
	WeeklyCount *int,
	WeeklyDuration *int,
) JSWeeklyOperationCountStruct {
	return JSWeeklyOperationCountStruct{
		Type:           "JSWeeklyOperationCount",
		Namespace:      Namespace,
		Reference:      Reference,
		StartDate:      StartDate,
		UsageObject:    UsageObject,
		WeeklyCount:    WeeklyCount,
		WeeklyDuration: WeeklyDuration,
	}
}

// JobFactory is just a simple function to instantiate the JobStruct
func JobFactory(
	ActionType string,
	CancelReason string,
	Cancelable *bool,
	EmailAddresses []string,
	Events []*JobEventStruct,
	JobState string,
	Name string,
	Namespace string,
	ParentAction string,
	ParentActionState string,
	PercentComplete float64,
	Queued *bool,
	Reference string,
	StartTime string,
	Suspendable *bool,
	Target string,
	TargetName string,
	TargetObjectType string,
	Title string,
	UpdateTime string,
	User string,
) JobStruct {
	return JobStruct{
		Type:              "Job",
		ActionType:        ActionType,
		CancelReason:      CancelReason,
		Cancelable:        Cancelable,
		EmailAddresses:    EmailAddresses,
		Events:            Events,
		JobState:          JobState,
		Name:              Name,
		Namespace:         Namespace,
		ParentAction:      ParentAction,
		ParentActionState: ParentActionState,
		PercentComplete:   PercentComplete,
		Queued:            Queued,
		Reference:         Reference,
		StartTime:         StartTime,
		Suspendable:       Suspendable,
		Target:            Target,
		TargetName:        TargetName,
		TargetObjectType:  TargetObjectType,
		Title:             Title,
		UpdateTime:        UpdateTime,
		User:              User,
	}
}

// JobEventFactory is just a simple function to instantiate the JobEventStruct
func JobEventFactory(
	Diagnoses []*DiagnosisResultStruct,
	EventType string,
	MessageAction string,
	MessageCode string,
	MessageCommandOutput string,
	MessageDetails string,
	PercentComplete float64,
	State string,
	Timestamp string,
) JobEventStruct {
	return JobEventStruct{
		Type:                 "JobEvent",
		Diagnoses:            Diagnoses,
		EventType:            EventType,
		MessageAction:        MessageAction,
		MessageCode:          MessageCode,
		MessageCommandOutput: MessageCommandOutput,
		MessageDetails:       MessageDetails,
		PercentComplete:      PercentComplete,
		State:                State,
		Timestamp:            Timestamp,
	}
}

// KerberosConfigFactory is just a simple function to instantiate the KerberosConfigStruct
func KerberosConfigFactory(
	Enabled *bool,
	Kdcs []*KerberosKDCStruct,
	Keytab string,
	Name string,
	Namespace string,
	Principal string,
	Realm string,
	Reference string,
) KerberosConfigStruct {
	return KerberosConfigStruct{
		Type:      "KerberosConfig",
		Enabled:   Enabled,
		Kdcs:      Kdcs,
		Keytab:    Keytab,
		Name:      Name,
		Namespace: Namespace,
		Principal: Principal,
		Realm:     Realm,
		Reference: Reference,
	}
}

// KerberosCredentialFactory is just a simple function to instantiate the KerberosCredentialStruct
func KerberosCredentialFactory() KerberosCredentialStruct {
	return KerberosCredentialStruct{
		Type: "KerberosCredential",
	}
}

// KerberosKDCFactory is just a simple function to instantiate the KerberosKDCStruct
func KerberosKDCFactory(
	Hostname string,
	Port *int,
) KerberosKDCStruct {
	return KerberosKDCStruct{
		Type:     "KerberosKDC",
		Hostname: Hostname,
		Port:     Port,
	}
}

// KeyPairCredentialFactory is just a simple function to instantiate the KeyPairCredentialStruct
func KeyPairCredentialFactory(
	PrivateKey string,
	PublicKey string,
) KeyPairCredentialStruct {
	return KeyPairCredentialStruct{
		Type:       "KeyPairCredential",
		PrivateKey: PrivateKey,
		PublicKey:  PublicKey,
	}
}

// KeystoreClientCertificateFactory is just a simple function to instantiate the KeystoreClientCertificateStruct
func KeystoreClientCertificateFactory(
	Token string,
) KeystoreClientCertificateStruct {
	return KeystoreClientCertificateStruct{
		Type:  "KeystoreClientCertificate",
		Token: Token,
	}
}

// LdapInfoFactory is just a simple function to instantiate the LdapInfoStruct
func LdapInfoFactory(
	Enabled *bool,
) LdapInfoStruct {
	return LdapInfoStruct{
		Type:    "LdapInfo",
		Enabled: Enabled,
	}
}

// LdapServerFactory is just a simple function to instantiate the LdapServerStruct
func LdapServerFactory(
	AuthMethod string,
	Host string,
	Name string,
	Namespace string,
	Port *int,
	Reference string,
	UseSSL *bool,
) LdapServerStruct {
	return LdapServerStruct{
		Type:       "LdapServer",
		AuthMethod: AuthMethod,
		Host:       Host,
		Name:       Name,
		Namespace:  Namespace,
		Port:       Port,
		Reference:  Reference,
		UseSSL:     UseSSL,
	}
}

// LicenseInfoFactory is just a simple function to instantiate the LicenseInfoStruct
func LicenseInfoFactory() LicenseInfoStruct {
	return LicenseInfoStruct{
		Type: "LicenseInfo",
	}
}

// LinkParametersFactory is just a simple function to instantiate the LinkParametersStruct
func LinkParametersFactory(
	Description string,
	Group string,
	LinkData LinkData,
	Name string,
) LinkParametersStruct {
	return LinkParametersStruct{
		Type:        "LinkParameters",
		Description: Description,
		Group:       Group,
		LinkData:    LinkData,
		Name:        Name,
	}
}

// LinkedSourceOperationsFactory is just a simple function to instantiate the LinkedSourceOperationsStruct
func LinkedSourceOperationsFactory(
	PostSync []SourceOperation,
	PreSync []SourceOperation,
) LinkedSourceOperationsStruct {
	return LinkedSourceOperationsStruct{
		Type:     "LinkedSourceOperations",
		PostSync: PostSync,
		PreSync:  PreSync,
	}
}

// ListResultFactory is just a simple function to instantiate the ListResultStruct
func ListResultFactory(
	Action string,
	Job string,
	Overflow *bool,
	Result string,
	Status string,
	Total *int,
) ListResultStruct {
	return ListResultStruct{
		Type:     "ListResult",
		Action:   Action,
		Job:      Job,
		Overflow: Overflow,
		Result:   Result,
		Status:   Status,
		Total:    Total,
	}
}

// LocaleSettingsFactory is just a simple function to instantiate the LocaleSettingsStruct
func LocaleSettingsFactory(
	Locale string,
	Name string,
	Namespace string,
	Reference string,
) LocaleSettingsStruct {
	return LocaleSettingsStruct{
		Type:      "LocaleSettings",
		Locale:    Locale,
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// LocalizedUpgradeCheckResultFactory is just a simple function to instantiate the LocalizedUpgradeCheckResultStruct
func LocalizedUpgradeCheckResultFactory() LocalizedUpgradeCheckResultStruct {
	return LocalizedUpgradeCheckResultStruct{
		Type: "LocalizedUpgradeCheckResult",
	}
}

// LoginRequestFactory is just a simple function to instantiate the LoginRequestStruct
func LoginRequestFactory(
	KeepAliveMode string,
	Password string,
	Target string,
	Username string,
) LoginRequestStruct {
	return LoginRequestStruct{
		Type:          "LoginRequest",
		KeepAliveMode: KeepAliveMode,
		Password:      Password,
		Target:        Target,
		Username:      Username,
	}
}

// MSSqlAttachDataFactory is just a simple function to instantiate the MSSqlAttachDataStruct
func MSSqlAttachDataFactory(
	EncryptionKey string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
	PptRepository string,
	SourceHostUser string,
	StagingPostScript string,
	StagingPreScript string,
	SyncStrategy MSSqlSyncStrategy,
) MSSqlAttachDataStruct {
	return MSSqlAttachDataStruct{
		Type:              "MSSqlAttachData",
		EncryptionKey:     EncryptionKey,
		Operations:        Operations,
		PptHostUser:       PptHostUser,
		PptRepository:     PptRepository,
		SourceHostUser:    SourceHostUser,
		StagingPostScript: StagingPostScript,
		StagingPreScript:  StagingPreScript,
		SyncStrategy:      SyncStrategy,
	}
}

// MSSqlAvailabilityGroupFactory is just a simple function to instantiate the MSSqlAvailabilityGroupStruct
func MSSqlAvailabilityGroupFactory(
	Environment string,
	FulltextInstalled *bool,
	Instances []MSSqlBaseClusterInstance,
	Listeners []MSSqlBaseClusterListener,
	Name string,
	Namespace string,
	ProvisioningEnabled *bool,
	Reference string,
	Staging *bool,
	Version string,
) MSSqlAvailabilityGroupStruct {
	return MSSqlAvailabilityGroupStruct{
		Type:                "MSSqlAvailabilityGroup",
		Environment:         Environment,
		FulltextInstalled:   FulltextInstalled,
		Instances:           Instances,
		Listeners:           Listeners,
		Name:                Name,
		Namespace:           Namespace,
		ProvisioningEnabled: ProvisioningEnabled,
		Reference:           Reference,
		Staging:             Staging,
		Version:             Version,
	}
}

// MSSqlAvailabilityGroupDBConfigFactory is just a simple function to instantiate the MSSqlAvailabilityGroupDBConfigStruct
func MSSqlAvailabilityGroupDBConfigFactory(
	DatabaseName string,
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	MirroringState string,
	MssqlUser MSSqlUser,
	Name string,
	Namespace string,
	RecoveryModel string,
	Reference string,
	Repository string,
) MSSqlAvailabilityGroupDBConfigStruct {
	return MSSqlAvailabilityGroupDBConfigStruct{
		Type:            "MSSqlAvailabilityGroupDBConfig",
		DatabaseName:    DatabaseName,
		Discovered:      Discovered,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		MirroringState:  MirroringState,
		MssqlUser:       MssqlUser,
		Name:            Name,
		Namespace:       Namespace,
		RecoveryModel:   RecoveryModel,
		Reference:       Reference,
		Repository:      Repository,
	}
}

// MSSqlAvailabilityGroupListenerFactory is just a simple function to instantiate the MSSqlAvailabilityGroupListenerStruct
func MSSqlAvailabilityGroupListenerFactory(
	Address string,
	Name string,
	Port *int,
) MSSqlAvailabilityGroupListenerStruct {
	return MSSqlAvailabilityGroupListenerStruct{
		Type:    "MSSqlAvailabilityGroupListener",
		Address: Address,
		Name:    Name,
		Port:    Port,
	}
}

// MSSqlClusterInstanceFactory is just a simple function to instantiate the MSSqlClusterInstanceStruct
func MSSqlClusterInstanceFactory(
	InstanceOwner string,
	Name string,
	Node string,
	Port *int,
	ServerName string,
	Version string,
) MSSqlClusterInstanceStruct {
	return MSSqlClusterInstanceStruct{
		Type:          "MSSqlClusterInstance",
		InstanceOwner: InstanceOwner,
		Name:          Name,
		Node:          Node,
		Port:          Port,
		ServerName:    ServerName,
		Version:       Version,
	}
}

// MSSqlCommvaultConfigFactory is just a simple function to instantiate the MSSqlCommvaultConfigStruct
func MSSqlCommvaultConfigFactory(
	CommserveHostName string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	SourceClientName string,
	StagingClientName string,
) MSSqlCommvaultConfigStruct {
	return MSSqlCommvaultConfigStruct{
		Type:              "MSSqlCommvaultConfig",
		CommserveHostName: CommserveHostName,
		ConfigParams:      ConfigParams,
		ConfigTemplate:    ConfigTemplate,
		SourceClientName:  SourceClientName,
		StagingClientName: StagingClientName,
	}
}

// MSSqlCompatibilityCriteriaFactory is just a simple function to instantiate the MSSqlCompatibilityCriteriaStruct
func MSSqlCompatibilityCriteriaFactory(
	Architecture *int,
	Environment string,
	Os string,
	Processor string,
	StagingEnabled *bool,
	Version string,
) MSSqlCompatibilityCriteriaStruct {
	return MSSqlCompatibilityCriteriaStruct{
		Type:           "MSSqlCompatibilityCriteria",
		Architecture:   Architecture,
		Environment:    Environment,
		Os:             Os,
		Processor:      Processor,
		StagingEnabled: StagingEnabled,
		Version:        Version,
	}
}

// MSSqlCreateTransformationParametersFactory is just a simple function to instantiate the MSSqlCreateTransformationParametersStruct
func MSSqlCreateTransformationParametersFactory(
	Container *MSSqlDatabaseContainerStruct,
	EnvironmentUser string,
	Operations []SourceOperation,
	Repository string,
) MSSqlCreateTransformationParametersStruct {
	return MSSqlCreateTransformationParametersStruct{
		Type:            "MSSqlCreateTransformationParameters",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Repository:      Repository,
	}
}

// MSSqlDBContainerRuntimeFactory is just a simple function to instantiate the MSSqlDBContainerRuntimeStruct
func MSSqlDBContainerRuntimeFactory(
	LastRestoredBackupSetUUID string,
	LogSyncActive *bool,
	PreProvisioningStatus *PreProvisioningRuntimeStruct,
) MSSqlDBContainerRuntimeStruct {
	return MSSqlDBContainerRuntimeStruct{
		Type:                      "MSSqlDBContainerRuntime",
		LastRestoredBackupSetUUID: LastRestoredBackupSetUUID,
		LogSyncActive:             LogSyncActive,
		PreProvisioningStatus:     PreProvisioningStatus,
	}
}

// MSSqlDatabaseContainerFactory is just a simple function to instantiate the MSSqlDatabaseContainerStruct
func MSSqlDatabaseContainerFactory(
	CreationTime string,
	CurrentTimeflow string,
	Description string,
	Group string,
	Masked *bool,
	Name string,
	Namespace string,
	Os string,
	PreviousTimeflow string,
	Processor string,
	ProvisionContainer string,
	Reference string,
	Runtime *MSSqlDBContainerRuntimeStruct,
	SourcingPolicy *SourcingPolicyStruct,
	Transformation *bool,
) MSSqlDatabaseContainerStruct {
	return MSSqlDatabaseContainerStruct{
		Type:               "MSSqlDatabaseContainer",
		CreationTime:       CreationTime,
		CurrentTimeflow:    CurrentTimeflow,
		Description:        Description,
		Group:              Group,
		Masked:             Masked,
		Name:               Name,
		Namespace:          Namespace,
		Os:                 Os,
		PreviousTimeflow:   PreviousTimeflow,
		Processor:          Processor,
		ProvisionContainer: ProvisionContainer,
		Reference:          Reference,
		Runtime:            Runtime,
		SourcingPolicy:     SourcingPolicy,
		Transformation:     Transformation,
	}
}

// MSSqlDatabaseUserFactory is just a simple function to instantiate the MSSqlDatabaseUserStruct
func MSSqlDatabaseUserFactory(
	Password *PasswordCredentialStruct,
	User string,
) MSSqlDatabaseUserStruct {
	return MSSqlDatabaseUserStruct{
		Type:     "MSSqlDatabaseUser",
		Password: Password,
		User:     User,
	}
}

// MSSqlDelphixManagedSyncStrategyFactory is just a simple function to instantiate the MSSqlDelphixManagedSyncStrategyStruct
func MSSqlDelphixManagedSyncStrategyFactory(
	BackupPolicy string,
	CompressionEnabled *bool,
	Config string,
	MssqlUser MSSqlUser,
) MSSqlDelphixManagedSyncStrategyStruct {
	return MSSqlDelphixManagedSyncStrategyStruct{
		Type:               "MSSqlDelphixManagedSyncStrategy",
		BackupPolicy:       BackupPolicy,
		CompressionEnabled: CompressionEnabled,
		Config:             Config,
		MssqlUser:          MssqlUser,
	}
}

// MSSqlDomainUserFactory is just a simple function to instantiate the MSSqlDomainUserStruct
func MSSqlDomainUserFactory(
	Password Credential,
	User string,
) MSSqlDomainUserStruct {
	return MSSqlDomainUserStruct{
		Type:     "MSSqlDomainUser",
		Password: Password,
		User:     User,
	}
}

// MSSqlEnvironmentUserFactory is just a simple function to instantiate the MSSqlEnvironmentUserStruct
func MSSqlEnvironmentUserFactory(
	User string,
) MSSqlEnvironmentUserStruct {
	return MSSqlEnvironmentUserStruct{
		Type: "MSSqlEnvironmentUser",
		User: User,
	}
}

// MSSqlExistingMostRecentBackupSyncParametersFactory is just a simple function to instantiate the MSSqlExistingMostRecentBackupSyncParametersStruct
func MSSqlExistingMostRecentBackupSyncParametersFactory() MSSqlExistingMostRecentBackupSyncParametersStruct {
	return MSSqlExistingMostRecentBackupSyncParametersStruct{
		Type: "MSSqlExistingMostRecentBackupSyncParameters",
	}
}

// MSSqlExistingSpecificBackupSyncParametersFactory is just a simple function to instantiate the MSSqlExistingSpecificBackupSyncParametersStruct
func MSSqlExistingSpecificBackupSyncParametersFactory(
	BackupUUID string,
) MSSqlExistingSpecificBackupSyncParametersStruct {
	return MSSqlExistingSpecificBackupSyncParametersStruct{
		Type:       "MSSqlExistingSpecificBackupSyncParameters",
		BackupUUID: BackupUUID,
	}
}

// MSSqlExportParametersFactory is just a simple function to instantiate the MSSqlExportParametersStruct
func MSSqlExportParametersFactory(
	ConfigParams map[string]string,
	EnableCdc *bool,
	FileMappingRules string,
	FilesystemLayout *TimeflowFilesystemLayoutStruct,
	RecoverDatabase *bool,
	RecoveryModel string,
	SourceConfig MSSqlDBConfig,
	TimeflowPointParameters TimeflowPointParameters,
) MSSqlExportParametersStruct {
	return MSSqlExportParametersStruct{
		Type:                    "MSSqlExportParameters",
		ConfigParams:            ConfigParams,
		EnableCdc:               EnableCdc,
		FileMappingRules:        FileMappingRules,
		FilesystemLayout:        FilesystemLayout,
		RecoverDatabase:         RecoverDatabase,
		RecoveryModel:           RecoveryModel,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// MSSqlExternalManagedSourceSyncStrategyFactory is just a simple function to instantiate the MSSqlExternalManagedSourceSyncStrategyStruct
func MSSqlExternalManagedSourceSyncStrategyFactory(
	Config string,
	MssqlCommvaultConfig string,
	MssqlNetbackupConfig string,
	MssqlUser MSSqlUser,
	SharedBackupLocations []string,
	ValidatedSyncMode string,
) MSSqlExternalManagedSourceSyncStrategyStruct {
	return MSSqlExternalManagedSourceSyncStrategyStruct{
		Type:                  "MSSqlExternalManagedSourceSyncStrategy",
		Config:                Config,
		MssqlCommvaultConfig:  MssqlCommvaultConfig,
		MssqlNetbackupConfig:  MssqlNetbackupConfig,
		MssqlUser:             MssqlUser,
		SharedBackupLocations: SharedBackupLocations,
		ValidatedSyncMode:     ValidatedSyncMode,
	}
}

// MSSqlFailoverClusterDBConfigFactory is just a simple function to instantiate the MSSqlFailoverClusterDBConfigStruct
func MSSqlFailoverClusterDBConfigFactory(
	DatabaseName string,
	Discovered *bool,
	DriveLetter string,
	EnvironmentUser string,
	LinkingEnabled *bool,
	MirroringState string,
	MssqlUser MSSqlUser,
	Name string,
	Namespace string,
	RecoveryModel string,
	Reference string,
	Repository string,
) MSSqlFailoverClusterDBConfigStruct {
	return MSSqlFailoverClusterDBConfigStruct{
		Type:            "MSSqlFailoverClusterDBConfig",
		DatabaseName:    DatabaseName,
		Discovered:      Discovered,
		DriveLetter:     DriveLetter,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		MirroringState:  MirroringState,
		MssqlUser:       MssqlUser,
		Name:            Name,
		Namespace:       Namespace,
		RecoveryModel:   RecoveryModel,
		Reference:       Reference,
		Repository:      Repository,
	}
}

// MSSqlFailoverClusterDriveLetterFactory is just a simple function to instantiate the MSSqlFailoverClusterDriveLetterStruct
func MSSqlFailoverClusterDriveLetterFactory(
	DriveLetter string,
	Label string,
) MSSqlFailoverClusterDriveLetterStruct {
	return MSSqlFailoverClusterDriveLetterStruct{
		Type:        "MSSqlFailoverClusterDriveLetter",
		DriveLetter: DriveLetter,
		Label:       Label,
	}
}

// MSSqlFailoverClusterInstanceFactory is just a simple function to instantiate the MSSqlFailoverClusterInstanceStruct
func MSSqlFailoverClusterInstanceFactory(
	InstanceOwner string,
	Name string,
	Node string,
	Port *int,
	ServerName string,
	Version string,
) MSSqlFailoverClusterInstanceStruct {
	return MSSqlFailoverClusterInstanceStruct{
		Type:          "MSSqlFailoverClusterInstance",
		InstanceOwner: InstanceOwner,
		Name:          Name,
		Node:          Node,
		Port:          Port,
		ServerName:    ServerName,
		Version:       Version,
	}
}

// MSSqlFailoverClusterListenerFactory is just a simple function to instantiate the MSSqlFailoverClusterListenerStruct
func MSSqlFailoverClusterListenerFactory(
	Address string,
	Name string,
	Port *int,
) MSSqlFailoverClusterListenerStruct {
	return MSSqlFailoverClusterListenerStruct{
		Type:    "MSSqlFailoverClusterListener",
		Address: Address,
		Name:    Name,
		Port:    Port,
	}
}

// MSSqlFailoverClusterRepositoryFactory is just a simple function to instantiate the MSSqlFailoverClusterRepositoryStruct
func MSSqlFailoverClusterRepositoryFactory(
	Drives []*MSSqlFailoverClusterDriveLetterStruct,
	Environment string,
	FulltextInstalled *bool,
	Instances []MSSqlBaseClusterInstance,
	Listeners []MSSqlBaseClusterListener,
	Name string,
	Namespace string,
	ProvisioningEnabled *bool,
	Reference string,
	Staging *bool,
	Version string,
) MSSqlFailoverClusterRepositoryStruct {
	return MSSqlFailoverClusterRepositoryStruct{
		Type:                "MSSqlFailoverClusterRepository",
		Drives:              Drives,
		Environment:         Environment,
		FulltextInstalled:   FulltextInstalled,
		Instances:           Instances,
		Listeners:           Listeners,
		Name:                Name,
		Namespace:           Namespace,
		ProvisioningEnabled: ProvisioningEnabled,
		Reference:           Reference,
		Staging:             Staging,
		Version:             Version,
	}
}

// MSSqlInstanceFactory is just a simple function to instantiate the MSSqlInstanceStruct
func MSSqlInstanceFactory(
	Discovered *bool,
	Environment string,
	FulltextInstalled *bool,
	InstallationPath string,
	InstanceName string,
	InstanceOwner string,
	Name string,
	Namespace string,
	Port *int,
	ProvisioningEnabled *bool,
	Reference string,
	ServerName string,
	Staging *bool,
	Version string,
) MSSqlInstanceStruct {
	return MSSqlInstanceStruct{
		Type:                "MSSqlInstance",
		Discovered:          Discovered,
		Environment:         Environment,
		FulltextInstalled:   FulltextInstalled,
		InstallationPath:    InstallationPath,
		InstanceName:        InstanceName,
		InstanceOwner:       InstanceOwner,
		Name:                Name,
		Namespace:           Namespace,
		Port:                Port,
		ProvisioningEnabled: ProvisioningEnabled,
		Reference:           Reference,
		ServerName:          ServerName,
		Staging:             Staging,
		Version:             Version,
	}
}

// MSSqlLinkDataFactory is just a simple function to instantiate the MSSqlLinkDataStruct
func MSSqlLinkDataFactory(
	EncryptionKey string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
	PptRepository string,
	SourceHostUser string,
	SourcingPolicy *SourcingPolicyStruct,
	StagingPostScript string,
	StagingPreScript string,
	SyncParameters MSSqlSyncParameters,
	SyncStrategy MSSqlSyncStrategy,
) MSSqlLinkDataStruct {
	return MSSqlLinkDataStruct{
		Type:              "MSSqlLinkData",
		EncryptionKey:     EncryptionKey,
		Operations:        Operations,
		PptHostUser:       PptHostUser,
		PptRepository:     PptRepository,
		SourceHostUser:    SourceHostUser,
		SourcingPolicy:    SourcingPolicy,
		StagingPostScript: StagingPostScript,
		StagingPreScript:  StagingPreScript,
		SyncParameters:    SyncParameters,
		SyncStrategy:      SyncStrategy,
	}
}

// MSSqlLinkedSourceFactory is just a simple function to instantiate the MSSqlLinkedSourceStruct
func MSSqlLinkedSourceFactory(
	Config string,
	Container string,
	Description string,
	EncryptionKey string,
	Hosts []string,
	IngestionStrategy IngestionStrategy,
	Linked *bool,
	LogCollectionEnabled *bool,
	MssqlCommvaultConfig string,
	MssqlNetbackupConfig string,
	Name string,
	Namespace string,
	Operations *LinkedSourceOperationsStruct,
	Reference string,
	Runtime *MSSqlSourceRuntimeStruct,
	SharedBackupLocations []string,
	Staging *bool,
	StagingSource string,
	Status string,
	Virtual *bool,
) MSSqlLinkedSourceStruct {
	return MSSqlLinkedSourceStruct{
		Type:                  "MSSqlLinkedSource",
		Config:                Config,
		Container:             Container,
		Description:           Description,
		EncryptionKey:         EncryptionKey,
		Hosts:                 Hosts,
		IngestionStrategy:     IngestionStrategy,
		Linked:                Linked,
		LogCollectionEnabled:  LogCollectionEnabled,
		MssqlCommvaultConfig:  MssqlCommvaultConfig,
		MssqlNetbackupConfig:  MssqlNetbackupConfig,
		Name:                  Name,
		Namespace:             Namespace,
		Operations:            Operations,
		Reference:             Reference,
		Runtime:               Runtime,
		SharedBackupLocations: SharedBackupLocations,
		Staging:               Staging,
		StagingSource:         StagingSource,
		Status:                Status,
		Virtual:               Virtual,
	}
}

// MSSqlLinkedSourceUpgradeParametersFactory is just a simple function to instantiate the MSSqlLinkedSourceUpgradeParametersStruct
func MSSqlLinkedSourceUpgradeParametersFactory(
	PptRepository string,
	SourceConfig MSSqlDBConfig,
) MSSqlLinkedSourceUpgradeParametersStruct {
	return MSSqlLinkedSourceUpgradeParametersStruct{
		Type:          "MSSqlLinkedSourceUpgradeParameters",
		PptRepository: PptRepository,
		SourceConfig:  SourceConfig,
	}
}

// MSSqlNetbackupConfigFactory is just a simple function to instantiate the MSSqlNetbackupConfigStruct
func MSSqlNetbackupConfigFactory(
	ConfigParams map[string]string,
	ConfigTemplate string,
	MasterName string,
	SourceClientName string,
) MSSqlNetbackupConfigStruct {
	return MSSqlNetbackupConfigStruct{
		Type:             "MSSqlNetbackupConfig",
		ConfigParams:     ConfigParams,
		ConfigTemplate:   ConfigTemplate,
		MasterName:       MasterName,
		SourceClientName: SourceClientName,
	}
}

// MSSqlNewCopyOnlyFullBackupSyncParametersFactory is just a simple function to instantiate the MSSqlNewCopyOnlyFullBackupSyncParametersStruct
func MSSqlNewCopyOnlyFullBackupSyncParametersFactory(
	BackupPolicy string,
	CompressionEnabled *bool,
) MSSqlNewCopyOnlyFullBackupSyncParametersStruct {
	return MSSqlNewCopyOnlyFullBackupSyncParametersStruct{
		Type:               "MSSqlNewCopyOnlyFullBackupSyncParameters",
		BackupPolicy:       BackupPolicy,
		CompressionEnabled: CompressionEnabled,
	}
}

// MSSqlPlatformParametersFactory is just a simple function to instantiate the MSSqlPlatformParametersStruct
func MSSqlPlatformParametersFactory() MSSqlPlatformParametersStruct {
	return MSSqlPlatformParametersStruct{
		Type: "MSSqlPlatformParameters",
	}
}

// MSSqlProvisionParametersFactory is just a simple function to instantiate the MSSqlProvisionParametersStruct
func MSSqlProvisionParametersFactory(
	Container *MSSqlDatabaseContainerStruct,
	Masked *bool,
	MaskingJob string,
	Source *MSSqlVirtualSourceStruct,
	SourceConfig MSSqlDBConfig,
	TimeflowPointParameters TimeflowPointParameters,
) MSSqlProvisionParametersStruct {
	return MSSqlProvisionParametersStruct{
		Type:                    "MSSqlProvisionParameters",
		Container:               Container,
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		Source:                  Source,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// MSSqlSIConfigFactory is just a simple function to instantiate the MSSqlSIConfigStruct
func MSSqlSIConfigFactory(
	DatabaseName string,
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	MirroringState string,
	MssqlUser MSSqlUser,
	Name string,
	Namespace string,
	RecoveryModel string,
	Reference string,
	Repository string,
) MSSqlSIConfigStruct {
	return MSSqlSIConfigStruct{
		Type:            "MSSqlSIConfig",
		DatabaseName:    DatabaseName,
		Discovered:      Discovered,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		MirroringState:  MirroringState,
		MssqlUser:       MssqlUser,
		Name:            Name,
		Namespace:       Namespace,
		RecoveryModel:   RecoveryModel,
		Reference:       Reference,
		Repository:      Repository,
	}
}

// MSSqlSnapshotFactory is just a simple function to instantiate the MSSqlSnapshotStruct
func MSSqlSnapshotFactory(
	BackupSetUUID string,
	Consistency string,
	Container string,
	CreationTime string,
	FirstChangePoint *MSSqlTimeflowPointStruct,
	InternalVersion *int,
	LatestChangePoint *MSSqlTimeflowPointStruct,
	MissingNonLoggedData *bool,
	Name string,
	Namespace string,
	Reference string,
	Retention *int,
	Runtime *MSSqlSnapshotRuntimeStruct,
	Temporary *bool,
	Timeflow string,
	Timezone string,
	Version string,
) MSSqlSnapshotStruct {
	return MSSqlSnapshotStruct{
		Type:                 "MSSqlSnapshot",
		BackupSetUUID:        BackupSetUUID,
		Consistency:          Consistency,
		Container:            Container,
		CreationTime:         CreationTime,
		FirstChangePoint:     FirstChangePoint,
		InternalVersion:      InternalVersion,
		LatestChangePoint:    LatestChangePoint,
		MissingNonLoggedData: MissingNonLoggedData,
		Name:                 Name,
		Namespace:            Namespace,
		Reference:            Reference,
		Retention:            Retention,
		Runtime:              Runtime,
		Temporary:            Temporary,
		Timeflow:             Timeflow,
		Timezone:             Timezone,
		Version:              Version,
	}
}

// MSSqlSnapshotRuntimeFactory is just a simple function to instantiate the MSSqlSnapshotRuntimeStruct
func MSSqlSnapshotRuntimeFactory(
	Provisionable *bool,
) MSSqlSnapshotRuntimeStruct {
	return MSSqlSnapshotRuntimeStruct{
		Type:          "MSSqlSnapshotRuntime",
		Provisionable: Provisionable,
	}
}

// MSSqlSourceConfigConnectivityFactory is just a simple function to instantiate the MSSqlSourceConfigConnectivityStruct
func MSSqlSourceConfigConnectivityFactory(
	MssqlUser MSSqlUser,
) MSSqlSourceConfigConnectivityStruct {
	return MSSqlSourceConfigConnectivityStruct{
		Type:      "MSSqlSourceConfigConnectivity",
		MssqlUser: MssqlUser,
	}
}

// MSSqlSourceConnectionInfoFactory is just a simple function to instantiate the MSSqlSourceConnectionInfoStruct
func MSSqlSourceConnectionInfoFactory(
	DatabaseName string,
	Host string,
	InstanceJDBCString string,
	InstanceName string,
	Port *int,
	Version string,
) MSSqlSourceConnectionInfoStruct {
	return MSSqlSourceConnectionInfoStruct{
		Type:               "MSSqlSourceConnectionInfo",
		DatabaseName:       DatabaseName,
		Host:               Host,
		InstanceJDBCString: InstanceJDBCString,
		InstanceName:       InstanceName,
		Port:               Port,
		Version:            Version,
	}
}

// MSSqlSourceRuntimeFactory is just a simple function to instantiate the MSSqlSourceRuntimeStruct
func MSSqlSourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	DatabaseSize float64,
	Enabled string,
	NotAccessibleReason string,
	Status string,
) MSSqlSourceRuntimeStruct {
	return MSSqlSourceRuntimeStruct{
		Type:                "MSSqlSourceRuntime",
		Accessible:          Accessible,
		AccessibleTimestamp: AccessibleTimestamp,
		DatabaseSize:        DatabaseSize,
		Enabled:             Enabled,
		NotAccessibleReason: NotAccessibleReason,
		Status:              Status,
	}
}

// MSSqlStagingSourceFactory is just a simple function to instantiate the MSSqlStagingSourceStruct
func MSSqlStagingSourceFactory(
	Config string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	PostScript string,
	PreScript string,
	Reference string,
	Runtime *MSSqlSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) MSSqlStagingSourceStruct {
	return MSSqlStagingSourceStruct{
		Type:                 "MSSqlStagingSource",
		Config:               Config,
		Container:            Container,
		Description:          Description,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		MountBase:            MountBase,
		Name:                 Name,
		Namespace:            Namespace,
		PostScript:           PostScript,
		PreScript:            PreScript,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		Status:               Status,
		Virtual:              Virtual,
	}
}

// MSSqlTimeflowFactory is just a simple function to instantiate the MSSqlTimeflowStruct
func MSSqlTimeflowFactory(
	Container string,
	CreationType string,
	DatabaseGuid string,
	Name string,
	Namespace string,
	ParentPoint *MSSqlTimeflowPointStruct,
	ParentSnapshot string,
	Reference string,
) MSSqlTimeflowStruct {
	return MSSqlTimeflowStruct{
		Type:           "MSSqlTimeflow",
		Container:      Container,
		CreationType:   CreationType,
		DatabaseGuid:   DatabaseGuid,
		Name:           Name,
		Namespace:      Namespace,
		ParentPoint:    ParentPoint,
		ParentSnapshot: ParentSnapshot,
		Reference:      Reference,
	}
}

// MSSqlTimeflowPointFactory is just a simple function to instantiate the MSSqlTimeflowPointStruct
func MSSqlTimeflowPointFactory(
	Location string,
	Timeflow string,
	Timestamp string,
) MSSqlTimeflowPointStruct {
	return MSSqlTimeflowPointStruct{
		Type:      "MSSqlTimeflowPoint",
		Location:  Location,
		Timeflow:  Timeflow,
		Timestamp: Timestamp,
	}
}

// MSSqlVirtualSourceFactory is just a simple function to instantiate the MSSqlVirtualSourceStruct
func MSSqlVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	Container string,
	Description string,
	EnableCdcOnProvision *bool,
	FileMappingRules string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	Operations *VirtualSourceOperationsStruct,
	PostScript string,
	PreScript string,
	Reference string,
	Runtime *MSSqlSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) MSSqlVirtualSourceStruct {
	return MSSqlVirtualSourceStruct{
		Type:                            "MSSqlVirtualSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		Container:                       Container,
		Description:                     Description,
		EnableCdcOnProvision:            EnableCdcOnProvision,
		FileMappingRules:                FileMappingRules,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		Operations:                      Operations,
		PostScript:                      PostScript,
		PreScript:                       PreScript,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Virtual:                         Virtual,
	}
}

// MaskingJobFactory is just a simple function to instantiate the MaskingJobStruct
func MaskingJobFactory(
	AssociatedContainer string,
	MaskingJobId string,
	Name string,
	Namespace string,
	Reference string,
) MaskingJobStruct {
	return MaskingJobStruct{
		Type:                "MaskingJob",
		AssociatedContainer: AssociatedContainer,
		MaskingJobId:        MaskingJobId,
		Name:                Name,
		Namespace:           Namespace,
		Reference:           Reference,
	}
}

// MaskingServiceConfigFactory is just a simple function to instantiate the MaskingServiceConfigStruct
func MaskingServiceConfigFactory(
	Credentials Credential,
	Name string,
	Namespace string,
	Port *int,
	Reference string,
	Scheme string,
	Server string,
	Username string,
) MaskingServiceConfigStruct {
	return MaskingServiceConfigStruct{
		Type:        "MaskingServiceConfig",
		Credentials: Credentials,
		Name:        Name,
		Namespace:   Namespace,
		Port:        Port,
		Reference:   Reference,
		Scheme:      Scheme,
		Server:      Server,
		Username:    Username,
	}
}

// MigrateCompatibilityParametersFactory is just a simple function to instantiate the MigrateCompatibilityParametersStruct
func MigrateCompatibilityParametersFactory(
	Environment string,
	SourceRepository string,
) MigrateCompatibilityParametersStruct {
	return MigrateCompatibilityParametersStruct{
		Type:             "MigrateCompatibilityParameters",
		Environment:      Environment,
		SourceRepository: SourceRepository,
	}
}

// MinimalPhoneHomeFactory is just a simple function to instantiate the MinimalPhoneHomeStruct
func MinimalPhoneHomeFactory(
	Enabled *bool,
) MinimalPhoneHomeStruct {
	return MinimalPhoneHomeStruct{
		Type:    "MinimalPhoneHome",
		Enabled: Enabled,
	}
}

// MySQLAttachDataFactory is just a simple function to instantiate the MySQLAttachDataStruct
func MySQLAttachDataFactory(
	Config string,
	ConfigParams map[string]string,
	DbCredentials *PasswordCredentialStruct,
	DbUser string,
	Operations *LinkedSourceOperationsStruct,
	StagingHostUser string,
	StagingPort *int,
	StagingRepository string,
) MySQLAttachDataStruct {
	return MySQLAttachDataStruct{
		Type:              "MySQLAttachData",
		Config:            Config,
		ConfigParams:      ConfigParams,
		DbCredentials:     DbCredentials,
		DbUser:            DbUser,
		Operations:        Operations,
		StagingHostUser:   StagingHostUser,
		StagingPort:       StagingPort,
		StagingRepository: StagingRepository,
	}
}

// MySQLBinlogCoordinatesFactory is just a simple function to instantiate the MySQLBinlogCoordinatesStruct
func MySQLBinlogCoordinatesFactory(
	MasterLogName string,
	MasterLogPos *int,
) MySQLBinlogCoordinatesStruct {
	return MySQLBinlogCoordinatesStruct{
		Type:          "MySQLBinlogCoordinates",
		MasterLogName: MasterLogName,
		MasterLogPos:  MasterLogPos,
	}
}

// MySQLCompatibilityCriteriaFactory is just a simple function to instantiate the MySQLCompatibilityCriteriaStruct
func MySQLCompatibilityCriteriaFactory(
	Architecture *int,
	Environment string,
	Os string,
	Processor string,
	StagingEnabled *bool,
	Version *MySQLVersionStruct,
) MySQLCompatibilityCriteriaStruct {
	return MySQLCompatibilityCriteriaStruct{
		Type:           "MySQLCompatibilityCriteria",
		Architecture:   Architecture,
		Environment:    Environment,
		Os:             Os,
		Processor:      Processor,
		StagingEnabled: StagingEnabled,
		Version:        Version,
	}
}

// MySQLCreateTransformationParametersFactory is just a simple function to instantiate the MySQLCreateTransformationParametersStruct
func MySQLCreateTransformationParametersFactory(
	Container *MySQLDatabaseContainerStruct,
	EnvironmentUser string,
	Operations []SourceOperation,
	Port *int,
	Repository string,
) MySQLCreateTransformationParametersStruct {
	return MySQLCreateTransformationParametersStruct{
		Type:            "MySQLCreateTransformationParameters",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Port:            Port,
		Repository:      Repository,
	}
}

// MySQLDBContainerRuntimeFactory is just a simple function to instantiate the MySQLDBContainerRuntimeStruct
func MySQLDBContainerRuntimeFactory(
	LogSyncActive *bool,
	PreProvisioningStatus *PreProvisioningRuntimeStruct,
) MySQLDBContainerRuntimeStruct {
	return MySQLDBContainerRuntimeStruct{
		Type:                  "MySQLDBContainerRuntime",
		LogSyncActive:         LogSyncActive,
		PreProvisioningStatus: PreProvisioningStatus,
	}
}

// MySQLDatabaseContainerFactory is just a simple function to instantiate the MySQLDatabaseContainerStruct
func MySQLDatabaseContainerFactory(
	CreationTime string,
	CurrentTimeflow string,
	Description string,
	Group string,
	Masked *bool,
	Name string,
	Namespace string,
	Os string,
	PreviousTimeflow string,
	Processor string,
	ProvisionContainer string,
	Reference string,
	Runtime *MySQLDBContainerRuntimeStruct,
	SourcingPolicy *SourcingPolicyStruct,
	Transformation *bool,
	Variant string,
) MySQLDatabaseContainerStruct {
	return MySQLDatabaseContainerStruct{
		Type:               "MySQLDatabaseContainer",
		CreationTime:       CreationTime,
		CurrentTimeflow:    CurrentTimeflow,
		Description:        Description,
		Group:              Group,
		Masked:             Masked,
		Name:               Name,
		Namespace:          Namespace,
		Os:                 Os,
		PreviousTimeflow:   PreviousTimeflow,
		Processor:          Processor,
		ProvisionContainer: ProvisionContainer,
		Reference:          Reference,
		Runtime:            Runtime,
		SourcingPolicy:     SourcingPolicy,
		Transformation:     Transformation,
		Variant:            Variant,
	}
}

// MySQLExistingMySQLDumpSyncParametersFactory is just a simple function to instantiate the MySQLExistingMySQLDumpSyncParametersStruct
func MySQLExistingMySQLDumpSyncParametersFactory(
	BackupLocation string,
	ReplicationCoordinates MySQLReplicationCoordinates,
) MySQLExistingMySQLDumpSyncParametersStruct {
	return MySQLExistingMySQLDumpSyncParametersStruct{
		Type:                   "MySQLExistingMySQLDumpSyncParameters",
		BackupLocation:         BackupLocation,
		ReplicationCoordinates: ReplicationCoordinates,
	}
}

// MySQLExportParametersFactory is just a simple function to instantiate the MySQLExportParametersStruct
func MySQLExportParametersFactory(
	ConfigParams map[string]string,
	FileMappingRules string,
	FilesystemLayout *TimeflowFilesystemLayoutStruct,
	RecoverDatabase *bool,
	SourceConfig *MySQLServerConfigStruct,
	TimeflowPointParameters TimeflowPointParameters,
) MySQLExportParametersStruct {
	return MySQLExportParametersStruct{
		Type:                    "MySQLExportParameters",
		ConfigParams:            ConfigParams,
		FileMappingRules:        FileMappingRules,
		FilesystemLayout:        FilesystemLayout,
		RecoverDatabase:         RecoverDatabase,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// MySQLGtidCoordinatesFactory is just a simple function to instantiate the MySQLGtidCoordinatesStruct
func MySQLGtidCoordinatesFactory(
	Gtids []string,
) MySQLGtidCoordinatesStruct {
	return MySQLGtidCoordinatesStruct{
		Type:  "MySQLGtidCoordinates",
		Gtids: Gtids,
	}
}

// MySQLInstallFactory is just a simple function to instantiate the MySQLInstallStruct
func MySQLInstallFactory(
	Discovered *bool,
	Environment string,
	InstallationPath string,
	InternalVersion *MySQLVersionStruct,
	Name string,
	Namespace string,
	ProvisioningEnabled *bool,
	Reference string,
	Staging *bool,
	Version string,
) MySQLInstallStruct {
	return MySQLInstallStruct{
		Type:                "MySQLInstall",
		Discovered:          Discovered,
		Environment:         Environment,
		InstallationPath:    InstallationPath,
		InternalVersion:     InternalVersion,
		Name:                Name,
		Namespace:           Namespace,
		ProvisioningEnabled: ProvisioningEnabled,
		Reference:           Reference,
		Staging:             Staging,
		Version:             Version,
	}
}

// MySQLLinkDataFactory is just a simple function to instantiate the MySQLLinkDataStruct
func MySQLLinkDataFactory(
	Config string,
	ConfigParams map[string]string,
	DbCredentials *PasswordCredentialStruct,
	DbUser string,
	Operations *LinkedSourceOperationsStruct,
	SourcingPolicy *SourcingPolicyStruct,
	StagingHostUser string,
	StagingPort *int,
	StagingRepository string,
	SyncParameters MySQLSyncParameters,
) MySQLLinkDataStruct {
	return MySQLLinkDataStruct{
		Type:              "MySQLLinkData",
		Config:            Config,
		ConfigParams:      ConfigParams,
		DbCredentials:     DbCredentials,
		DbUser:            DbUser,
		Operations:        Operations,
		SourcingPolicy:    SourcingPolicy,
		StagingHostUser:   StagingHostUser,
		StagingPort:       StagingPort,
		StagingRepository: StagingRepository,
		SyncParameters:    SyncParameters,
	}
}

// MySQLLinkedSourceFactory is just a simple function to instantiate the MySQLLinkedSourceStruct
func MySQLLinkedSourceFactory(
	Config string,
	ConfigParams map[string]string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	Operations *LinkedSourceOperationsStruct,
	Reference string,
	Runtime *MySQLSourceRuntimeStruct,
	Staging *bool,
	StagingSource string,
	Status string,
	Virtual *bool,
) MySQLLinkedSourceStruct {
	return MySQLLinkedSourceStruct{
		Type:                 "MySQLLinkedSource",
		Config:               Config,
		ConfigParams:         ConfigParams,
		Container:            Container,
		Description:          Description,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Namespace:            Namespace,
		Operations:           Operations,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		StagingSource:        StagingSource,
		Status:               Status,
		Virtual:              Virtual,
	}
}

// MySQLNewMySQLDumpSyncParametersFactory is just a simple function to instantiate the MySQLNewMySQLDumpSyncParametersStruct
func MySQLNewMySQLDumpSyncParametersFactory() MySQLNewMySQLDumpSyncParametersStruct {
	return MySQLNewMySQLDumpSyncParametersStruct{
		Type: "MySQLNewMySQLDumpSyncParameters",
	}
}

// MySQLPlatformParametersFactory is just a simple function to instantiate the MySQLPlatformParametersStruct
func MySQLPlatformParametersFactory(
	Port *int,
) MySQLPlatformParametersStruct {
	return MySQLPlatformParametersStruct{
		Type: "MySQLPlatformParameters",
		Port: Port,
	}
}

// MySQLProvisionParametersFactory is just a simple function to instantiate the MySQLProvisionParametersStruct
func MySQLProvisionParametersFactory(
	Container *MySQLDatabaseContainerStruct,
	Masked *bool,
	MaskingJob string,
	Source *MySQLVirtualSourceStruct,
	SourceConfig *MySQLServerConfigStruct,
	TimeflowPointParameters TimeflowPointParameters,
) MySQLProvisionParametersStruct {
	return MySQLProvisionParametersStruct{
		Type:                    "MySQLProvisionParameters",
		Container:               Container,
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		Source:                  Source,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// MySQLServerConfigFactory is just a simple function to instantiate the MySQLServerConfigStruct
func MySQLServerConfigFactory(
	Credentials *PasswordCredentialStruct,
	DataDirectory string,
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	Port *int,
	Reference string,
	Repository string,
	User string,
) MySQLServerConfigStruct {
	return MySQLServerConfigStruct{
		Type:            "MySQLServerConfig",
		Credentials:     Credentials,
		DataDirectory:   DataDirectory,
		Discovered:      Discovered,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Namespace:       Namespace,
		Port:            Port,
		Reference:       Reference,
		Repository:      Repository,
		User:            User,
	}
}

// MySQLSnapshotFactory is just a simple function to instantiate the MySQLSnapshotStruct
func MySQLSnapshotFactory(
	Consistency string,
	Container string,
	CreationTime string,
	FirstChangePoint *MySQLTimeflowPointStruct,
	InternalVersion *MySQLVersionStruct,
	LatestChangePoint *MySQLTimeflowPointStruct,
	MissingNonLoggedData *bool,
	Name string,
	Namespace string,
	Reference string,
	Retention *int,
	Runtime *MySQLSnapshotRuntimeStruct,
	Temporary *bool,
	Timeflow string,
	Timezone string,
	Version string,
) MySQLSnapshotStruct {
	return MySQLSnapshotStruct{
		Type:                 "MySQLSnapshot",
		Consistency:          Consistency,
		Container:            Container,
		CreationTime:         CreationTime,
		FirstChangePoint:     FirstChangePoint,
		InternalVersion:      InternalVersion,
		LatestChangePoint:    LatestChangePoint,
		MissingNonLoggedData: MissingNonLoggedData,
		Name:                 Name,
		Namespace:            Namespace,
		Reference:            Reference,
		Retention:            Retention,
		Runtime:              Runtime,
		Temporary:            Temporary,
		Timeflow:             Timeflow,
		Timezone:             Timezone,
		Version:              Version,
	}
}

// MySQLSnapshotRuntimeFactory is just a simple function to instantiate the MySQLSnapshotRuntimeStruct
func MySQLSnapshotRuntimeFactory(
	Provisionable *bool,
) MySQLSnapshotRuntimeStruct {
	return MySQLSnapshotRuntimeStruct{
		Type:          "MySQLSnapshotRuntime",
		Provisionable: Provisionable,
	}
}

// MySQLSourceConnectionInfoFactory is just a simple function to instantiate the MySQLSourceConnectionInfoStruct
func MySQLSourceConnectionInfoFactory(
	DataDirectory string,
	Host string,
	JdbcString string,
	Port *int,
	User string,
	Version string,
) MySQLSourceConnectionInfoStruct {
	return MySQLSourceConnectionInfoStruct{
		Type:          "MySQLSourceConnectionInfo",
		DataDirectory: DataDirectory,
		Host:          Host,
		JdbcString:    JdbcString,
		Port:          Port,
		User:          User,
		Version:       Version,
	}
}

// MySQLSourceRuntimeFactory is just a simple function to instantiate the MySQLSourceRuntimeStruct
func MySQLSourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	DatabaseSize float64,
	Enabled string,
	NotAccessibleReason string,
	Status string,
) MySQLSourceRuntimeStruct {
	return MySQLSourceRuntimeStruct{
		Type:                "MySQLSourceRuntime",
		Accessible:          Accessible,
		AccessibleTimestamp: AccessibleTimestamp,
		DatabaseSize:        DatabaseSize,
		Enabled:             Enabled,
		NotAccessibleReason: NotAccessibleReason,
		Status:              Status,
	}
}

// MySQLStagingSourceFactory is just a simple function to instantiate the MySQLStagingSourceStruct
func MySQLStagingSourceFactory(
	Config string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	PostScript string,
	PreScript string,
	Reference string,
	Runtime *MySQLSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) MySQLStagingSourceStruct {
	return MySQLStagingSourceStruct{
		Type:                 "MySQLStagingSource",
		Config:               Config,
		Container:            Container,
		Description:          Description,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		MountBase:            MountBase,
		Name:                 Name,
		Namespace:            Namespace,
		PostScript:           PostScript,
		PreScript:            PreScript,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		Status:               Status,
		Virtual:              Virtual,
	}
}

// MySQLTimeflowFactory is just a simple function to instantiate the MySQLTimeflowStruct
func MySQLTimeflowFactory(
	Container string,
	CreationType string,
	Name string,
	Namespace string,
	ParentPoint *MySQLTimeflowPointStruct,
	ParentSnapshot string,
	Reference string,
) MySQLTimeflowStruct {
	return MySQLTimeflowStruct{
		Type:           "MySQLTimeflow",
		Container:      Container,
		CreationType:   CreationType,
		Name:           Name,
		Namespace:      Namespace,
		ParentPoint:    ParentPoint,
		ParentSnapshot: ParentSnapshot,
		Reference:      Reference,
	}
}

// MySQLTimeflowPointFactory is just a simple function to instantiate the MySQLTimeflowPointStruct
func MySQLTimeflowPointFactory(
	Location string,
	Timeflow string,
	Timestamp string,
) MySQLTimeflowPointStruct {
	return MySQLTimeflowPointStruct{
		Type:      "MySQLTimeflowPoint",
		Location:  Location,
		Timeflow:  Timeflow,
		Timestamp: Timestamp,
	}
}

// MySQLVersionFactory is just a simple function to instantiate the MySQLVersionStruct
func MySQLVersionFactory(
	Variant string,
	Version string,
) MySQLVersionStruct {
	return MySQLVersionStruct{
		Type:    "MySQLVersion",
		Variant: Variant,
		Version: Version,
	}
}

// MySQLVirtualSourceFactory is just a simple function to instantiate the MySQLVirtualSourceStruct
func MySQLVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	Container string,
	Description string,
	FileMappingRules string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	Operations *VirtualSourceOperationsStruct,
	Reference string,
	Runtime *MySQLSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) MySQLVirtualSourceStruct {
	return MySQLVirtualSourceStruct{
		Type:                            "MySQLVirtualSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		Container:                       Container,
		Description:                     Description,
		FileMappingRules:                FileMappingRules,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		Operations:                      Operations,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Virtual:                         Virtual,
	}
}

// MySQLXtraBackupSyncParametersFactory is just a simple function to instantiate the MySQLXtraBackupSyncParametersStruct
func MySQLXtraBackupSyncParametersFactory(
	BackupLocation string,
	ReplicationCoordinates MySQLReplicationCoordinates,
) MySQLXtraBackupSyncParametersStruct {
	return MySQLXtraBackupSyncParametersStruct{
		Type:                   "MySQLXtraBackupSyncParameters",
		BackupLocation:         BackupLocation,
		ReplicationCoordinates: ReplicationCoordinates,
	}
}

// NTPConfigFactory is just a simple function to instantiate the NTPConfigStruct
func NTPConfigFactory(
	Enabled *bool,
	MulticastAddress string,
	Servers []string,
	UseMulticast *bool,
) NTPConfigStruct {
	return NTPConfigStruct{
		Type:             "NTPConfig",
		Enabled:          Enabled,
		MulticastAddress: MulticastAddress,
		Servers:          Servers,
		UseMulticast:     UseMulticast,
	}
}

// NamedKeyPairCredentialFactory is just a simple function to instantiate the NamedKeyPairCredentialStruct
func NamedKeyPairCredentialFactory(
	PrivateKey string,
	PublicKey string,
	Username string,
) NamedKeyPairCredentialStruct {
	return NamedKeyPairCredentialStruct{
		Type:       "NamedKeyPairCredential",
		PrivateKey: PrivateKey,
		PublicKey:  PublicKey,
		Username:   Username,
	}
}

// NamedPasswordCredentialFactory is just a simple function to instantiate the NamedPasswordCredentialStruct
func NamedPasswordCredentialFactory(
	Password string,
	Username string,
) NamedPasswordCredentialStruct {
	return NamedPasswordCredentialStruct{
		Type:     "NamedPasswordCredential",
		Password: Password,
		Username: Username,
	}
}

// NamespaceFactory is just a simple function to instantiate the NamespaceStruct
func NamespaceFactory(
	Description string,
	FailedOver *bool,
	FailoverReport string,
	Locked *bool,
	Name string,
	Namespace string,
	NamespaceType string,
	Reference string,
	SecureNamespace *bool,
	Source string,
	Tag string,
) NamespaceStruct {
	return NamespaceStruct{
		Type:            "Namespace",
		Description:     Description,
		FailedOver:      FailedOver,
		FailoverReport:  FailoverReport,
		Locked:          Locked,
		Name:            Name,
		Namespace:       Namespace,
		NamespaceType:   NamespaceType,
		Reference:       Reference,
		SecureNamespace: SecureNamespace,
		Source:          Source,
		Tag:             Tag,
	}
}

// NamespaceFailoverParametersFactory is just a simple function to instantiate the NamespaceFailoverParametersStruct
func NamespaceFailoverParametersFactory(
	SmartFailover *bool,
) NamespaceFailoverParametersStruct {
	return NamespaceFailoverParametersStruct{
		Type:          "NamespaceFailoverParameters",
		SmartFailover: SmartFailover,
	}
}

// NetbackupConnectivityParametersFactory is just a simple function to instantiate the NetbackupConnectivityParametersStruct
func NetbackupConnectivityParametersFactory(
	Environment string,
	EnvironmentUser string,
	MasterServerName string,
	SourceClientName string,
) NetbackupConnectivityParametersStruct {
	return NetbackupConnectivityParametersStruct{
		Type:             "NetbackupConnectivityParameters",
		Environment:      Environment,
		EnvironmentUser:  EnvironmentUser,
		MasterServerName: MasterServerName,
		SourceClientName: SourceClientName,
	}
}

// NetworkDSPTestFactory is just a simple function to instantiate the NetworkDSPTestStruct
func NetworkDSPTestFactory(
	EndTime string,
	Name string,
	Namespace string,
	NumConnections *int,
	Parameters *NetworkDSPTestParametersStruct,
	Reference string,
	RemoteAddress string,
	StartTime string,
	State string,
	Throughput float64,
) NetworkDSPTestStruct {
	return NetworkDSPTestStruct{
		Type:           "NetworkDSPTest",
		EndTime:        EndTime,
		Name:           Name,
		Namespace:      Namespace,
		NumConnections: NumConnections,
		Parameters:     Parameters,
		Reference:      Reference,
		RemoteAddress:  RemoteAddress,
		StartTime:      StartTime,
		State:          State,
		Throughput:     Throughput,
	}
}

// NetworkDSPTestParametersFactory is just a simple function to instantiate the NetworkDSPTestParametersStruct
func NetworkDSPTestParametersFactory(
	BlockSize *int,
	Compression *bool,
	DestinationType string,
	Direction string,
	Duration *int,
	Encryption *bool,
	NumConnections *int,
	QueueDepth *int,
	ReceiveSocketBuffer *int,
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfoStruct,
	RemoteHost string,
	SendSocketBuffer *int,
	XportScheduler string,
) NetworkDSPTestParametersStruct {
	return NetworkDSPTestParametersStruct{
		Type:                    "NetworkDSPTestParameters",
		BlockSize:               BlockSize,
		Compression:             Compression,
		DestinationType:         DestinationType,
		Direction:               Direction,
		Duration:                Duration,
		Encryption:              Encryption,
		NumConnections:          NumConnections,
		QueueDepth:              QueueDepth,
		ReceiveSocketBuffer:     ReceiveSocketBuffer,
		RemoteDelphixEngineInfo: RemoteDelphixEngineInfo,
		RemoteHost:              RemoteHost,
		SendSocketBuffer:        SendSocketBuffer,
		XportScheduler:          XportScheduler,
	}
}

// NetworkInterfaceFactory is just a simple function to instantiate the NetworkInterfaceStruct
func NetworkInterfaceFactory(
	Addresses []*InterfaceAddressStruct,
	Device string,
	MacAddress string,
	Mtu *int,
	MtuRange string,
	Name string,
	Namespace string,
	Reference string,
	State string,
) NetworkInterfaceStruct {
	return NetworkInterfaceStruct{
		Type:       "NetworkInterface",
		Addresses:  Addresses,
		Device:     Device,
		MacAddress: MacAddress,
		Mtu:        Mtu,
		MtuRange:   MtuRange,
		Name:       Name,
		Namespace:  Namespace,
		Reference:  Reference,
		State:      State,
	}
}

// NetworkInterfaceUtilDatapointFactory is just a simple function to instantiate the NetworkInterfaceUtilDatapointStruct
func NetworkInterfaceUtilDatapointFactory(
	InBytes *int,
	InPackets *int,
	OutBytes *int,
	OutPackets *int,
	Timestamp string,
) NetworkInterfaceUtilDatapointStruct {
	return NetworkInterfaceUtilDatapointStruct{
		Type:       "NetworkInterfaceUtilDatapoint",
		InBytes:    InBytes,
		InPackets:  InPackets,
		OutBytes:   OutBytes,
		OutPackets: OutPackets,
		Timestamp:  Timestamp,
	}
}

// NetworkInterfaceUtilDatapointStreamFactory is just a simple function to instantiate the NetworkInterfaceUtilDatapointStreamStruct
func NetworkInterfaceUtilDatapointStreamFactory(
	Datapoints []Datapoint,
	NetworkInterface string,
) NetworkInterfaceUtilDatapointStreamStruct {
	return NetworkInterfaceUtilDatapointStreamStruct{
		Type:             "NetworkInterfaceUtilDatapointStream",
		Datapoints:       Datapoints,
		NetworkInterface: NetworkInterface,
	}
}

// NetworkLatencyTestFactory is just a simple function to instantiate the NetworkLatencyTestStruct
func NetworkLatencyTestFactory(
	Average *int,
	EndTime string,
	Loss *int,
	Maximum *int,
	Minimum *int,
	Name string,
	Namespace string,
	Parameters *NetworkLatencyTestParametersStruct,
	Reference string,
	RemoteAddress string,
	StartTime string,
	State string,
	Stddev *int,
) NetworkLatencyTestStruct {
	return NetworkLatencyTestStruct{
		Type:          "NetworkLatencyTest",
		Average:       Average,
		EndTime:       EndTime,
		Loss:          Loss,
		Maximum:       Maximum,
		Minimum:       Minimum,
		Name:          Name,
		Namespace:     Namespace,
		Parameters:    Parameters,
		Reference:     Reference,
		RemoteAddress: RemoteAddress,
		StartTime:     StartTime,
		State:         State,
		Stddev:        Stddev,
	}
}

// NetworkLatencyTestParametersFactory is just a simple function to instantiate the NetworkLatencyTestParametersStruct
func NetworkLatencyTestParametersFactory(
	RemoteAddress string,
	RemoteHost string,
	RequestCount *int,
	RequestSize *int,
) NetworkLatencyTestParametersStruct {
	return NetworkLatencyTestParametersStruct{
		Type:          "NetworkLatencyTestParameters",
		RemoteAddress: RemoteAddress,
		RemoteHost:    RemoteHost,
		RequestCount:  RequestCount,
		RequestSize:   RequestSize,
	}
}

// NetworkRouteFactory is just a simple function to instantiate the NetworkRouteStruct
func NetworkRouteFactory(
	Destination string,
	Gateway string,
	OutInterface string,
) NetworkRouteStruct {
	return NetworkRouteStruct{
		Type:         "NetworkRoute",
		Destination:  Destination,
		Gateway:      Gateway,
		OutInterface: OutInterface,
	}
}

// NetworkRouteLookupParametersFactory is just a simple function to instantiate the NetworkRouteLookupParametersStruct
func NetworkRouteLookupParametersFactory(
	Destination string,
) NetworkRouteLookupParametersStruct {
	return NetworkRouteLookupParametersStruct{
		Type:        "NetworkRouteLookupParameters",
		Destination: Destination,
	}
}

// NetworkThroughputTestFactory is just a simple function to instantiate the NetworkThroughputTestStruct
func NetworkThroughputTestFactory(
	EndTime string,
	Name string,
	Namespace string,
	NumConnections *int,
	Parameters *NetworkThroughputTestParametersStruct,
	Reference string,
	RemoteAddress string,
	StartTime string,
	State string,
	Throughput float64,
) NetworkThroughputTestStruct {
	return NetworkThroughputTestStruct{
		Type:           "NetworkThroughputTest",
		EndTime:        EndTime,
		Name:           Name,
		Namespace:      Namespace,
		NumConnections: NumConnections,
		Parameters:     Parameters,
		Reference:      Reference,
		RemoteAddress:  RemoteAddress,
		StartTime:      StartTime,
		State:          State,
		Throughput:     Throughput,
	}
}

// NetworkThroughputTestParametersFactory is just a simple function to instantiate the NetworkThroughputTestParametersStruct
func NetworkThroughputTestParametersFactory(
	BlockSize *int,
	Direction string,
	Duration *int,
	NumConnections *int,
	Port *int,
	RemoteAddress string,
	RemoteHost string,
	SendSocketBuffer *int,
) NetworkThroughputTestParametersStruct {
	return NetworkThroughputTestParametersStruct{
		Type:             "NetworkThroughputTestParameters",
		BlockSize:        BlockSize,
		Direction:        Direction,
		Duration:         Duration,
		NumConnections:   NumConnections,
		Port:             Port,
		RemoteAddress:    RemoteAddress,
		RemoteHost:       RemoteHost,
		SendSocketBuffer: SendSocketBuffer,
	}
}

// NfsConfigFactory is just a simple function to instantiate the NfsConfigStruct
func NfsConfigFactory(
	MountVersion string,
) NfsConfigStruct {
	return NfsConfigStruct{
		Type:         "NfsConfig",
		MountVersion: MountVersion,
	}
}

// NfsOpsDatapointStreamFactory is just a simple function to instantiate the NfsOpsDatapointStreamStruct
func NfsOpsDatapointStreamFactory(
	Cached *bool,
	Client string,
	Datapoints []Datapoint,
	Op string,
	Path string,
	Sync *bool,
	Ver string,
) NfsOpsDatapointStreamStruct {
	return NfsOpsDatapointStreamStruct{
		Type:       "NfsOpsDatapointStream",
		Cached:     Cached,
		Client:     Client,
		Datapoints: Datapoints,
		Op:         Op,
		Path:       Path,
		Sync:       Sync,
		Ver:        Ver,
	}
}

// NoBackupIngestionStrategyFactory is just a simple function to instantiate the NoBackupIngestionStrategyStruct
func NoBackupIngestionStrategyFactory() NoBackupIngestionStrategyStruct {
	return NoBackupIngestionStrategyStruct{
		Type: "NoBackupIngestionStrategy",
	}
}

// NotFilterFactory is just a simple function to instantiate the NotFilterStruct
func NotFilterFactory(
	SubFilter AlertFilter,
) NotFilterStruct {
	return NotFilterStruct{
		Type:      "NotFilter",
		SubFilter: SubFilter,
	}
}

// NotificationDropFactory is just a simple function to instantiate the NotificationDropStruct
func NotificationDropFactory(
	DropCount *int,
) NotificationDropStruct {
	return NotificationDropStruct{
		Type:      "NotificationDrop",
		DropCount: DropCount,
	}
}

// NullConstraintFactory is just a simple function to instantiate the NullConstraintStruct
func NullConstraintFactory(
	AxisName string,
) NullConstraintStruct {
	return NullConstraintStruct{
		Type:     "NullConstraint",
		AxisName: AxisName,
	}
}

// OKResultFactory is just a simple function to instantiate the OKResultStruct
func OKResultFactory(
	Action string,
	Job string,
	Result string,
	Status string,
) OKResultStruct {
	return OKResultStruct{
		Type:   "OKResult",
		Action: Action,
		Job:    Job,
		Result: Result,
		Status: Status,
	}
}

// ObjectNotificationFactory is just a simple function to instantiate the ObjectNotificationStruct
func ObjectNotificationFactory(
	EventType string,
	Object string,
	ObjectType string,
) ObjectNotificationStruct {
	return ObjectNotificationStruct{
		Type:       "ObjectNotification",
		EventType:  EventType,
		Object:     Object,
		ObjectType: ObjectType,
	}
}

// OperationTemplateFactory is just a simple function to instantiate the OperationTemplateStruct
func OperationTemplateFactory(
	Description string,
	LastUpdated string,
	Name string,
	Namespace string,
	Operation SourceOperation,
	Reference string,
) OperationTemplateStruct {
	return OperationTemplateStruct{
		Type:        "OperationTemplate",
		Description: Description,
		LastUpdated: LastUpdated,
		Name:        Name,
		Namespace:   Namespace,
		Operation:   Operation,
		Reference:   Reference,
	}
}

// OrFilterFactory is just a simple function to instantiate the OrFilterStruct
func OrFilterFactory(
	SubFilters []AlertFilter,
) OrFilterStruct {
	return OrFilterStruct{
		Type:       "OrFilter",
		SubFilters: SubFilters,
	}
}

// OracleActiveInstanceFactory is just a simple function to instantiate the OracleActiveInstanceStruct
func OracleActiveInstanceFactory(
	HostName string,
	InstanceName string,
	InstanceNumber *int,
) OracleActiveInstanceStruct {
	return OracleActiveInstanceStruct{
		Type:           "OracleActiveInstance",
		HostName:       HostName,
		InstanceName:   InstanceName,
		InstanceNumber: InstanceNumber,
	}
}

// OracleAddLiveSourceParametersFactory is just a simple function to instantiate the OracleAddLiveSourceParametersStruct
func OracleAddLiveSourceParametersFactory(
	Credential Credential,
	Source *OracleLiveSourceStruct,
	SourceConfig OracleDBConfig,
	Username string,
) OracleAddLiveSourceParametersStruct {
	return OracleAddLiveSourceParametersStruct{
		Type:         "OracleAddLiveSourceParameters",
		Credential:   Credential,
		Source:       Source,
		SourceConfig: SourceConfig,
		Username:     Username,
	}
}

// OracleAttachDataFactory is just a simple function to instantiate the OracleAttachDataStruct
func OracleAttachDataFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	Config string,
	DoubleSync *bool,
	EncryptedLinkingEnabled *bool,
	EnvironmentUser string,
	ExternalFilePath string,
	FilesPerSet *int,
	Force *bool,
	LinkNow *bool,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	OracleFallbackCredentials Credential,
	OracleFallbackUser string,
	RmanChannels *int,
) OracleAttachDataStruct {
	return OracleAttachDataStruct{
		Type:                      "OracleAttachData",
		BackupLevelEnabled:        BackupLevelEnabled,
		BandwidthLimit:            BandwidthLimit,
		CheckLogical:              CheckLogical,
		CompressedLinkingEnabled:  CompressedLinkingEnabled,
		Config:                    Config,
		DoubleSync:                DoubleSync,
		EncryptedLinkingEnabled:   EncryptedLinkingEnabled,
		EnvironmentUser:           EnvironmentUser,
		ExternalFilePath:          ExternalFilePath,
		FilesPerSet:               FilesPerSet,
		Force:                     Force,
		LinkNow:                   LinkNow,
		NumberOfConnections:       NumberOfConnections,
		Operations:                Operations,
		OracleFallbackCredentials: OracleFallbackCredentials,
		OracleFallbackUser:        OracleFallbackUser,
		RmanChannels:              RmanChannels,
	}
}

// OracleCharacterSetFactory is just a simple function to instantiate the OracleCharacterSetStruct
func OracleCharacterSetFactory(
	CharacterSet string,
) OracleCharacterSetStruct {
	return OracleCharacterSetStruct{
		Type:         "OracleCharacterSet",
		CharacterSet: CharacterSet,
	}
}

// OracleClusterFactory is just a simple function to instantiate the OracleClusterStruct
func OracleClusterFactory(
	ClusterUser string,
	CrsClusterHome string,
	CrsClusterName string,
	Description string,
	Enabled *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	PrimaryUser string,
	Reference string,
	RemoteListener string,
	Scan string,
	ScanManual *bool,
) OracleClusterStruct {
	return OracleClusterStruct{
		Type:                 "OracleCluster",
		ClusterUser:          ClusterUser,
		CrsClusterHome:       CrsClusterHome,
		CrsClusterName:       CrsClusterName,
		Description:          Description,
		Enabled:              Enabled,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Namespace:            Namespace,
		PrimaryUser:          PrimaryUser,
		Reference:            Reference,
		RemoteListener:       RemoteListener,
		Scan:                 Scan,
		ScanManual:           ScanManual,
	}
}

// OracleClusterCreateParametersFactory is just a simple function to instantiate the OracleClusterCreateParametersStruct
func OracleClusterCreateParametersFactory(
	Cluster *OracleClusterStruct,
	LogCollectionEnabled *bool,
	Node *OracleClusterNodeCreateParametersStruct,
	PrimaryUser *EnvironmentUserStruct,
) OracleClusterCreateParametersStruct {
	return OracleClusterCreateParametersStruct{
		Type:                 "OracleClusterCreateParameters",
		Cluster:              Cluster,
		LogCollectionEnabled: LogCollectionEnabled,
		Node:                 Node,
		PrimaryUser:          PrimaryUser,
	}
}

// OracleClusterNodeFactory is just a simple function to instantiate the OracleClusterNodeStruct
func OracleClusterNodeFactory(
	Cluster string,
	Discovered *bool,
	Enabled *bool,
	Host string,
	Name string,
	Namespace string,
	Reference string,
	VirtualIPs []*OracleVirtualIPStruct,
) OracleClusterNodeStruct {
	return OracleClusterNodeStruct{
		Type:       "OracleClusterNode",
		Cluster:    Cluster,
		Discovered: Discovered,
		Enabled:    Enabled,
		Host:       Host,
		Name:       Name,
		Namespace:  Namespace,
		Reference:  Reference,
		VirtualIPs: VirtualIPs,
	}
}

// OracleClusterNodeCreateParametersFactory is just a simple function to instantiate the OracleClusterNodeCreateParametersStruct
func OracleClusterNodeCreateParametersFactory(
	Cluster string,
	HostParameters HostCreateParameters,
	Name string,
	VirtualIPs []*OracleVirtualIPStruct,
) OracleClusterNodeCreateParametersStruct {
	return OracleClusterNodeCreateParametersStruct{
		Type:           "OracleClusterNodeCreateParameters",
		Cluster:        Cluster,
		HostParameters: HostParameters,
		Name:           Name,
		VirtualIPs:     VirtualIPs,
	}
}

// OracleCompatibilityCriteriaFactory is just a simple function to instantiate the OracleCompatibilityCriteriaStruct
func OracleCompatibilityCriteriaFactory(
	Architecture *int,
	Environment string,
	Os string,
	Processor string,
	StagingEnabled *bool,
	Version string,
) OracleCompatibilityCriteriaStruct {
	return OracleCompatibilityCriteriaStruct{
		Type:           "OracleCompatibilityCriteria",
		Architecture:   Architecture,
		Environment:    Environment,
		Os:             Os,
		Processor:      Processor,
		StagingEnabled: StagingEnabled,
		Version:        Version,
	}
}

// OracleCreateTransformationParametersFactory is just a simple function to instantiate the OracleCreateTransformationParametersStruct
func OracleCreateTransformationParametersFactory(
	Container *OracleDatabaseContainerStruct,
	EnvironmentUser string,
	Operations []SourceOperation,
	Repository string,
) OracleCreateTransformationParametersStruct {
	return OracleCreateTransformationParametersStruct{
		Type:            "OracleCreateTransformationParameters",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Repository:      Repository,
	}
}

// OracleCustomEnvVarRACFileFactory is just a simple function to instantiate the OracleCustomEnvVarRACFileStruct
func OracleCustomEnvVarRACFileFactory(
	ClusterNode string,
	PathParameters string,
) OracleCustomEnvVarRACFileStruct {
	return OracleCustomEnvVarRACFileStruct{
		Type:           "OracleCustomEnvVarRACFile",
		ClusterNode:    ClusterNode,
		PathParameters: PathParameters,
	}
}

// OracleCustomEnvVarRACPairFactory is just a simple function to instantiate the OracleCustomEnvVarRACPairStruct
func OracleCustomEnvVarRACPairFactory(
	ClusterNode string,
	VarName string,
	VarValue string,
) OracleCustomEnvVarRACPairStruct {
	return OracleCustomEnvVarRACPairStruct{
		Type:        "OracleCustomEnvVarRACPair",
		ClusterNode: ClusterNode,
		VarName:     VarName,
		VarValue:    VarValue,
	}
}

// OracleCustomEnvVarSIFileFactory is just a simple function to instantiate the OracleCustomEnvVarSIFileStruct
func OracleCustomEnvVarSIFileFactory(
	PathParameters string,
) OracleCustomEnvVarSIFileStruct {
	return OracleCustomEnvVarSIFileStruct{
		Type:           "OracleCustomEnvVarSIFile",
		PathParameters: PathParameters,
	}
}

// OracleCustomEnvVarSIPairFactory is just a simple function to instantiate the OracleCustomEnvVarSIPairStruct
func OracleCustomEnvVarSIPairFactory(
	VarName string,
	VarValue string,
) OracleCustomEnvVarSIPairStruct {
	return OracleCustomEnvVarSIPairStruct{
		Type:     "OracleCustomEnvVarSIPair",
		VarName:  VarName,
		VarValue: VarValue,
	}
}

// OracleDBConfigConnectivityFactory is just a simple function to instantiate the OracleDBConfigConnectivityStruct
func OracleDBConfigConnectivityFactory(
	Credentials Credential,
	Username string,
) OracleDBConfigConnectivityStruct {
	return OracleDBConfigConnectivityStruct{
		Type:        "OracleDBConfigConnectivity",
		Credentials: Credentials,
		Username:    Username,
	}
}

// OracleDBContainerRuntimeFactory is just a simple function to instantiate the OracleDBContainerRuntimeStruct
func OracleDBContainerRuntimeFactory(
	LiveSourceEligible *bool,
	LogSyncActive *bool,
	PreProvisioningStatus *PreProvisioningRuntimeStruct,
) OracleDBContainerRuntimeStruct {
	return OracleDBContainerRuntimeStruct{
		Type:                  "OracleDBContainerRuntime",
		LiveSourceEligible:    LiveSourceEligible,
		LogSyncActive:         LogSyncActive,
		PreProvisioningStatus: PreProvisioningStatus,
	}
}

// OracleDatabaseContainerFactory is just a simple function to instantiate the OracleDatabaseContainerStruct
func OracleDatabaseContainerFactory(
	ContentType string,
	CreationTime string,
	CurrentTimeflow string,
	Description string,
	DiagnoseNoLoggingFaults *bool,
	Group string,
	LiveSource *bool,
	Masked *bool,
	Name string,
	Namespace string,
	Os string,
	PerformanceMode string,
	PhysicalStandby *bool,
	PreProvisioningEnabled *bool,
	PreviousTimeflow string,
	Processor string,
	ProvisionContainer string,
	RacMaxInstanceLag *int,
	Reference string,
	Runtime *OracleDBContainerRuntimeStruct,
	SourcingPolicy *OracleSourcingPolicyStruct,
	Transformation *bool,
) OracleDatabaseContainerStruct {
	return OracleDatabaseContainerStruct{
		Type:                    "OracleDatabaseContainer",
		ContentType:             ContentType,
		CreationTime:            CreationTime,
		CurrentTimeflow:         CurrentTimeflow,
		Description:             Description,
		DiagnoseNoLoggingFaults: DiagnoseNoLoggingFaults,
		Group:                   Group,
		LiveSource:              LiveSource,
		Masked:                  Masked,
		Name:                    Name,
		Namespace:               Namespace,
		Os:                      Os,
		PerformanceMode:         PerformanceMode,
		PhysicalStandby:         PhysicalStandby,
		PreProvisioningEnabled:  PreProvisioningEnabled,
		PreviousTimeflow:        PreviousTimeflow,
		Processor:               Processor,
		ProvisionContainer:      ProvisionContainer,
		RacMaxInstanceLag:       RacMaxInstanceLag,
		Reference:               Reference,
		Runtime:                 Runtime,
		SourcingPolicy:          SourcingPolicy,
		Transformation:          Transformation,
	}
}

// OracleDatabaseStatisticFactory is just a simple function to instantiate the OracleDatabaseStatisticStruct
func OracleDatabaseStatisticFactory(
	StatisticValues []string,
) OracleDatabaseStatisticStruct {
	return OracleDatabaseStatisticStruct{
		Type:            "OracleDatabaseStatistic",
		StatisticValues: StatisticValues,
	}
}

// OracleDatabaseStatsSectionFactory is just a simple function to instantiate the OracleDatabaseStatsSectionStruct
func OracleDatabaseStatsSectionFactory(
	ColumnHeaders []string,
	RowValues []*OracleDatabaseStatisticStruct,
	SectionName string,
) OracleDatabaseStatsSectionStruct {
	return OracleDatabaseStatsSectionStruct{
		Type:          "OracleDatabaseStatsSection",
		ColumnHeaders: ColumnHeaders,
		RowValues:     RowValues,
		SectionName:   SectionName,
	}
}

// OracleDeleteParametersFactory is just a simple function to instantiate the OracleDeleteParametersStruct
func OracleDeleteParametersFactory(
	Credential Credential,
	Force *bool,
	Username string,
) OracleDeleteParametersStruct {
	return OracleDeleteParametersStruct{
		Type:       "OracleDeleteParameters",
		Credential: Credential,
		Force:      Force,
		Username:   Username,
	}
}

// OracleDisableParametersFactory is just a simple function to instantiate the OracleDisableParametersStruct
func OracleDisableParametersFactory(
	AttemptCleanup *bool,
	Credential Credential,
	Username string,
) OracleDisableParametersStruct {
	return OracleDisableParametersStruct{
		Type:           "OracleDisableParameters",
		AttemptCleanup: AttemptCleanup,
		Credential:     Credential,
		Username:       Username,
	}
}

// OracleEnableParametersFactory is just a simple function to instantiate the OracleEnableParametersStruct
func OracleEnableParametersFactory(
	AttemptStart *bool,
	Credential Credential,
	Username string,
) OracleEnableParametersStruct {
	return OracleEnableParametersStruct{
		Type:         "OracleEnableParameters",
		AttemptStart: AttemptStart,
		Credential:   Credential,
		Username:     Username,
	}
}

// OracleExportFactory is just a simple function to instantiate the OracleExportStruct
func OracleExportFactory(
	DspOptions *DSPOptionsStruct,
	FileParallelism *int,
) OracleExportStruct {
	return OracleExportStruct{
		Type:            "OracleExport",
		DspOptions:      DspOptions,
		FileParallelism: FileParallelism,
	}
}

// OracleExportParametersFactory is just a simple function to instantiate the OracleExportParametersStruct
func OracleExportParametersFactory(
	ConfigParams map[string]string,
	DspOptions *DSPOptionsStruct,
	FileMappingRules string,
	FileParallelism *int,
	FilesystemLayout *TimeflowFilesystemLayoutStruct,
	OpenDatabase *bool,
	RecoverDatabase *bool,
	SourceConfig OracleDBConfig,
	TimeflowPointParameters TimeflowPointParameters,
) OracleExportParametersStruct {
	return OracleExportParametersStruct{
		Type:                    "OracleExportParameters",
		ConfigParams:            ConfigParams,
		DspOptions:              DspOptions,
		FileMappingRules:        FileMappingRules,
		FileParallelism:         FileParallelism,
		FilesystemLayout:        FilesystemLayout,
		OpenDatabase:            OpenDatabase,
		RecoverDatabase:         RecoverDatabase,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// OracleFetchedLogFactory is just a simple function to instantiate the OracleFetchedLogStruct
func OracleFetchedLogFactory(
	Container string,
	EndScn *int,
	EndTimestamp string,
	InstanceNum *int,
	Sequence *int,
	StartScn *int,
	StartTimestamp string,
	Timeflow string,
) OracleFetchedLogStruct {
	return OracleFetchedLogStruct{
		Type:           "OracleFetchedLog",
		Container:      Container,
		EndScn:         EndScn,
		EndTimestamp:   EndTimestamp,
		InstanceNum:    InstanceNum,
		Sequence:       Sequence,
		StartScn:       StartScn,
		StartTimestamp: StartTimestamp,
		Timeflow:       Timeflow,
	}
}

// OracleInstallFactory is just a simple function to instantiate the OracleInstallStruct
func OracleInstallFactory(
	AppliedPatches []*int,
	Bits *int,
	Discovered *bool,
	Environment string,
	GroupId *int,
	GroupName string,
	InstallationHome string,
	LogsyncPossible *bool,
	Name string,
	Namespace string,
	OracleBase string,
	ProvisioningEnabled *bool,
	Rac *bool,
	Reference string,
	Staging *bool,
	UserId *int,
	UserName string,
	Version string,
) OracleInstallStruct {
	return OracleInstallStruct{
		Type:                "OracleInstall",
		AppliedPatches:      AppliedPatches,
		Bits:                Bits,
		Discovered:          Discovered,
		Environment:         Environment,
		GroupId:             GroupId,
		GroupName:           GroupName,
		InstallationHome:    InstallationHome,
		LogsyncPossible:     LogsyncPossible,
		Name:                Name,
		Namespace:           Namespace,
		OracleBase:          OracleBase,
		ProvisioningEnabled: ProvisioningEnabled,
		Rac:                 Rac,
		Reference:           Reference,
		Staging:             Staging,
		UserId:              UserId,
		UserName:            UserName,
		Version:             Version,
	}
}

// OracleInstanceFactory is just a simple function to instantiate the OracleInstanceStruct
func OracleInstanceFactory(
	InstanceName string,
	InstanceNumber *int,
) OracleInstanceStruct {
	return OracleInstanceStruct{
		Type:           "OracleInstance",
		InstanceName:   InstanceName,
		InstanceNumber: InstanceNumber,
	}
}

// OracleLinkFromExternalFactory is just a simple function to instantiate the OracleLinkFromExternalStruct
func OracleLinkFromExternalFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	Config string,
	DiagnoseNoLoggingFaults *bool,
	EncryptedLinkingEnabled *bool,
	EnvironmentUser string,
	//ExternalFilePath string,
	FilesPerSet *int,
	LinkNow *bool,
	NonSysCredentials Credential,
	NonSysUser string,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	OracleFallbackCredentials Credential,
	OracleFallbackUser string,
	PreProvisioningEnabled *bool,
	RmanChannels *int,
	SourcingPolicy *OracleSourcingPolicyStruct,
	//SyncParameters OracleSyncParameters,
	SyncParameters OracleSyncFromExternalParametersStruct,
) OracleLinkFromExternalStruct {
	return OracleLinkFromExternalStruct{
		Type:                     "OracleLinkFromExternal",
		BackupLevelEnabled:       BackupLevelEnabled,
		BandwidthLimit:           BandwidthLimit,
		CheckLogical:             CheckLogical,
		CompressedLinkingEnabled: CompressedLinkingEnabled,
		Config:                   Config,
		DiagnoseNoLoggingFaults:  DiagnoseNoLoggingFaults,
		EncryptedLinkingEnabled:  EncryptedLinkingEnabled,
		EnvironmentUser:          EnvironmentUser,
		//ExternalFilePath: ExternalFilePath,
		FilesPerSet:               FilesPerSet,
		LinkNow:                   LinkNow,
		NonSysCredentials:         NonSysCredentials,
		NonSysUser:                NonSysUser,
		NumberOfConnections:       NumberOfConnections,
		Operations:                Operations,
		OracleFallbackCredentials: OracleFallbackCredentials,
		OracleFallbackUser:        OracleFallbackUser,
		PreProvisioningEnabled:    PreProvisioningEnabled,
		RmanChannels:              RmanChannels,
		SourcingPolicy:            SourcingPolicy,
		SyncParameters:            SyncParameters,
	}
}

// OracleLinkedSourceFactory is just a simple function to instantiate the OracleLinkedSourceStruct
func OracleLinkedSourceFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	Config string,
	Container string,
	Description string,
	EncryptedLinkingEnabled *bool,
	ExternalFilePath string,
	FilesPerSet *int,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	Reference string,
	RmanChannels *int,
	Runtime OracleBaseSourceRuntime,
	Staging *bool,
	Status string,
	Virtual *bool,
) OracleLinkedSourceStruct {
	return OracleLinkedSourceStruct{
		Type:                     "OracleLinkedSource",
		BackupLevelEnabled:       BackupLevelEnabled,
		BandwidthLimit:           BandwidthLimit,
		CheckLogical:             CheckLogical,
		CompressedLinkingEnabled: CompressedLinkingEnabled,
		Config:                   Config,
		Container:                Container,
		Description:              Description,
		EncryptedLinkingEnabled:  EncryptedLinkingEnabled,
		ExternalFilePath:         ExternalFilePath,
		FilesPerSet:              FilesPerSet,
		Hosts:                    Hosts,
		Linked:                   Linked,
		LogCollectionEnabled:     LogCollectionEnabled,
		Name:                     Name,
		Namespace:                Namespace,
		NumberOfConnections:      NumberOfConnections,
		Operations:               Operations,
		Reference:                Reference,
		RmanChannels:             RmanChannels,
		Runtime:                  Runtime,
		Staging:                  Staging,
		Status:                   Status,
		Virtual:                  Virtual,
	}
}

// OracleLiveSourceFactory is just a simple function to instantiate the OracleLiveSourceStruct
func OracleLiveSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	Container string,
	CustomEnvVars []OracleCustomEnvVar,
	DataAgeWarningThreshold *int,
	Description string,
	FileMappingRules string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
	Reference string,
	ResyncStatus string,
	Runtime OracleBaseSourceRuntime,
	Staging *bool,
	Status string,
	Virtual *bool,
) OracleLiveSourceStruct {
	return OracleLiveSourceStruct{
		Type:                            "OracleLiveSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		Container:                       Container,
		CustomEnvVars:                   CustomEnvVars,
		DataAgeWarningThreshold:         DataAgeWarningThreshold,
		Description:                     Description,
		FileMappingRules:                FileMappingRules,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
		Reference:                       Reference,
		ResyncStatus:                    ResyncStatus,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Virtual:                         Virtual,
	}
}

// OracleLiveSourceRuntimeFactory is just a simple function to instantiate the OracleLiveSourceRuntimeStruct
func OracleLiveSourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	ActiveInstances []*OracleActiveInstanceStruct,
	ApplyStatus string,
	CurrentDataAge *int,
	DatabaseMode string,
	DatabaseRole string,
	DatabaseSize float64,
	DatabaseStats []*OracleDatabaseStatsSectionStruct,
	Enabled string,
	IsDataAgeWarningExceeded *bool,
	LastNonLoggedLocation string,
	LastUpdateTime string,
	NonLoggedDataDetected *bool,
	NotAccessibleReason string,
	SourceDatabaseTimezone string,
	SourceResetlogsIDChangeDetected *bool,
	Status string,
	TransportStatus string,
	UnexpectedRoleChangeDetected *bool,
) OracleLiveSourceRuntimeStruct {
	return OracleLiveSourceRuntimeStruct{
		Type:                            "OracleLiveSourceRuntime",
		Accessible:                      Accessible,
		AccessibleTimestamp:             AccessibleTimestamp,
		ActiveInstances:                 ActiveInstances,
		ApplyStatus:                     ApplyStatus,
		CurrentDataAge:                  CurrentDataAge,
		DatabaseMode:                    DatabaseMode,
		DatabaseRole:                    DatabaseRole,
		DatabaseSize:                    DatabaseSize,
		DatabaseStats:                   DatabaseStats,
		Enabled:                         Enabled,
		IsDataAgeWarningExceeded:        IsDataAgeWarningExceeded,
		LastNonLoggedLocation:           LastNonLoggedLocation,
		LastUpdateTime:                  LastUpdateTime,
		NonLoggedDataDetected:           NonLoggedDataDetected,
		NotAccessibleReason:             NotAccessibleReason,
		SourceDatabaseTimezone:          SourceDatabaseTimezone,
		SourceResetlogsIDChangeDetected: SourceResetlogsIDChangeDetected,
		Status:                          Status,
		TransportStatus:                 TransportStatus,
		UnexpectedRoleChangeDetected:    UnexpectedRoleChangeDetected,
	}
}

// OracleLogFactory is just a simple function to instantiate the OracleLogStruct
func OracleLogFactory(
	InstanceNum *int,
	Sequence *int,
) OracleLogStruct {
	return OracleLogStruct{
		Type:        "OracleLog",
		InstanceNum: InstanceNum,
		Sequence:    Sequence,
	}
}

// OracleMissingLogFactory is just a simple function to instantiate the OracleMissingLogStruct
func OracleMissingLogFactory(
	Container string,
	EndScn *int,
	EndTimestamp string,
	InstanceNum *int,
	Sequence *int,
	StartScn *int,
	StartTimestamp string,
	Timeflow string,
) OracleMissingLogStruct {
	return OracleMissingLogStruct{
		Type:           "OracleMissingLog",
		Container:      Container,
		EndScn:         EndScn,
		EndTimestamp:   EndTimestamp,
		InstanceNum:    InstanceNum,
		Sequence:       Sequence,
		StartScn:       StartScn,
		StartTimestamp: StartTimestamp,
		Timeflow:       Timeflow,
	}
}

// OracleMultitenantProvisionParametersFactory is just a simple function to instantiate the OracleMultitenantProvisionParametersStruct
func OracleMultitenantProvisionParametersFactory(
	Container *OracleDatabaseContainerStruct,
	Credential Credential,
	Masked *bool,
	MaskingJob string,
	Source *OracleVirtualPdbSourceStruct,
	SourceConfig *OraclePDBConfigStruct,
	TimeflowPointParameters TimeflowPointParameters,
	Username string,
	VirtualCdb *OracleVirtualCdbProvisionParametersStruct,
) OracleMultitenantProvisionParametersStruct {
	return OracleMultitenantProvisionParametersStruct{
		Type:                    "OracleMultitenantProvisionParameters",
		Container:               Container,
		Credential:              Credential,
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		Source:                  Source,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
		Username:                Username,
		VirtualCdb:              VirtualCdb,
	}
}

// OracleNodeListenerFactory is just a simple function to instantiate the OracleNodeListenerStruct
func OracleNodeListenerFactory(
	Discovered *bool,
	Environment string,
	Host string,
	Name string,
	Namespace string,
	ProtocolAddresses []string,
	Reference string,
) OracleNodeListenerStruct {
	return OracleNodeListenerStruct{
		Type:              "OracleNodeListener",
		Discovered:        Discovered,
		Environment:       Environment,
		Host:              Host,
		Name:              Name,
		Namespace:         Namespace,
		ProtocolAddresses: ProtocolAddresses,
		Reference:         Reference,
	}
}

// OraclePDBAttachDataFactory is just a simple function to instantiate the OraclePDBAttachDataStruct
func OraclePDBAttachDataFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	Config string,
	EncryptedLinkingEnabled *bool,
	EnvironmentUser string,
	ExternalFilePath string,
	FilesPerSet *int,
	Force *bool,
	LinkNow *bool,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	OracleFallbackCredentials Credential,
	OracleFallbackUser string,
	RmanChannels *int,
) OraclePDBAttachDataStruct {
	return OraclePDBAttachDataStruct{
		Type:                      "OraclePDBAttachData",
		BackupLevelEnabled:        BackupLevelEnabled,
		BandwidthLimit:            BandwidthLimit,
		CheckLogical:              CheckLogical,
		CompressedLinkingEnabled:  CompressedLinkingEnabled,
		Config:                    Config,
		EncryptedLinkingEnabled:   EncryptedLinkingEnabled,
		EnvironmentUser:           EnvironmentUser,
		ExternalFilePath:          ExternalFilePath,
		FilesPerSet:               FilesPerSet,
		Force:                     Force,
		LinkNow:                   LinkNow,
		NumberOfConnections:       NumberOfConnections,
		Operations:                Operations,
		OracleFallbackCredentials: OracleFallbackCredentials,
		OracleFallbackUser:        OracleFallbackUser,
		RmanChannels:              RmanChannels,
	}
}

// OraclePDBConfigFactory is just a simple function to instantiate the OraclePDBConfigStruct
func OraclePDBConfigFactory(
	CdbConfig string,
	Credentials Credential,
	DatabaseName string,
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	NonSysCredentials string,
	NonSysUser string,
	Reference string,
	Repository string,
	Services []*OracleServiceStruct,
	User string,
) OraclePDBConfigStruct {
	return OraclePDBConfigStruct{
		Type:              "OraclePDBConfig",
		CdbConfig:         CdbConfig,
		Credentials:       Credentials,
		DatabaseName:      DatabaseName,
		Discovered:        Discovered,
		EnvironmentUser:   EnvironmentUser,
		LinkingEnabled:    LinkingEnabled,
		Name:              Name,
		Namespace:         Namespace,
		NonSysCredentials: NonSysCredentials,
		NonSysUser:        NonSysUser,
		Reference:         Reference,
		Repository:        Repository,
		Services:          Services,
		User:              User,
	}
}

// OraclePDBLinkFromExternalFactory is just a simple function to instantiate the OraclePDBLinkFromExternalStruct
func OraclePDBLinkFromExternalFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	Config string,
	DiagnoseNoLoggingFaults *bool,
	EncryptedLinkingEnabled *bool,
	EnvironmentUser string,
	ExternalFilePath string,
	FilesPerSet *int,
	LinkNow *bool,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	OracleFallbackCredentials Credential,
	OracleFallbackUser string,
	PreProvisioningEnabled *bool,
	RmanChannels *int,
	SourcingPolicy *OracleSourcingPolicyStruct,
	SyncParameters OracleSyncParameters,
) OraclePDBLinkFromExternalStruct {
	return OraclePDBLinkFromExternalStruct{
		Type:                      "OraclePDBLinkFromExternal",
		BackupLevelEnabled:        BackupLevelEnabled,
		BandwidthLimit:            BandwidthLimit,
		CheckLogical:              CheckLogical,
		CompressedLinkingEnabled:  CompressedLinkingEnabled,
		Config:                    Config,
		DiagnoseNoLoggingFaults:   DiagnoseNoLoggingFaults,
		EncryptedLinkingEnabled:   EncryptedLinkingEnabled,
		EnvironmentUser:           EnvironmentUser,
		ExternalFilePath:          ExternalFilePath,
		FilesPerSet:               FilesPerSet,
		LinkNow:                   LinkNow,
		NumberOfConnections:       NumberOfConnections,
		Operations:                Operations,
		OracleFallbackCredentials: OracleFallbackCredentials,
		OracleFallbackUser:        OracleFallbackUser,
		PreProvisioningEnabled:    PreProvisioningEnabled,
		RmanChannels:              RmanChannels,
		SourcingPolicy:            SourcingPolicy,
		SyncParameters:            SyncParameters,
	}
}

// OraclePDBSourceRuntimeFactory is just a simple function to instantiate the OraclePDBSourceRuntimeStruct
func OraclePDBSourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	ActiveInstances []*OracleActiveInstanceStruct,
	DatabaseMode string,
	DatabaseRole string,
	DatabaseSize float64,
	DatabaseStats []*OracleDatabaseStatsSectionStruct,
	Enabled string,
	LastNonLoggedLocation string,
	NotAccessibleReason string,
	Status string,
) OraclePDBSourceRuntimeStruct {
	return OraclePDBSourceRuntimeStruct{
		Type:                  "OraclePDBSourceRuntime",
		Accessible:            Accessible,
		AccessibleTimestamp:   AccessibleTimestamp,
		ActiveInstances:       ActiveInstances,
		DatabaseMode:          DatabaseMode,
		DatabaseRole:          DatabaseRole,
		DatabaseSize:          DatabaseSize,
		DatabaseStats:         DatabaseStats,
		Enabled:               Enabled,
		LastNonLoggedLocation: LastNonLoggedLocation,
		NotAccessibleReason:   NotAccessibleReason,
		Status:                Status,
	}
}

// OraclePlatformParametersFactory is just a simple function to instantiate the OraclePlatformParametersStruct
func OraclePlatformParametersFactory() OraclePlatformParametersStruct {
	return OraclePlatformParametersStruct{
		Type: "OraclePlatformParameters",
	}
}

// OracleProvisionParametersFactory is just a simple function to instantiate the OracleProvisionParametersStruct
func OracleProvisionParametersFactory(
	Container *OracleDatabaseContainerStruct,
	Credential Credential,
	Masked *bool,
	MaskingJob string,
	NewDBID *bool,
	OpenResetlogs *bool,
	PhysicalStandby *bool,
	Source *OracleVirtualSourceStruct,
	SourceConfig OracleDBConfig,
	TimeflowPointParameters TimeflowPointParameters,
	Username string,
) OracleProvisionParametersStruct {
	return OracleProvisionParametersStruct{
		Type:                    "OracleProvisionParameters",
		Container:               Container,
		Credential:              Credential,
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		NewDBID:                 NewDBID,
		OpenResetlogs:           OpenResetlogs,
		PhysicalStandby:         PhysicalStandby,
		Source:                  Source,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
		Username:                Username,
	}
}

// OracleRACConfigFactory is just a simple function to instantiate the OracleRACConfigStruct
func OracleRACConfigFactory(
	Credentials Credential,
	CrsDatabaseName string,
	DatabaseName string,
	Discovered *bool,
	EnvironmentUser string,
	Instances []*OracleRACInstanceStruct,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	NonSysCredentials string,
	NonSysUser string,
	Reference string,
	Repository string,
	Services []*OracleServiceStruct,
	TdeKeystorePassword string,
	UniqueName string,
	User string,
) OracleRACConfigStruct {
	return OracleRACConfigStruct{
		Type:                "OracleRACConfig",
		Credentials:         Credentials,
		CrsDatabaseName:     CrsDatabaseName,
		DatabaseName:        DatabaseName,
		Discovered:          Discovered,
		EnvironmentUser:     EnvironmentUser,
		Instances:           Instances,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		Namespace:           Namespace,
		NonSysCredentials:   NonSysCredentials,
		NonSysUser:          NonSysUser,
		Reference:           Reference,
		Repository:          Repository,
		Services:            Services,
		TdeKeystorePassword: TdeKeystorePassword,
		UniqueName:          UniqueName,
		User:                User,
	}
}

// OracleRACInstanceFactory is just a simple function to instantiate the OracleRACInstanceStruct
func OracleRACInstanceFactory(
	InstanceName string,
	InstanceNumber *int,
	Node string,
) OracleRACInstanceStruct {
	return OracleRACInstanceStruct{
		Type:           "OracleRACInstance",
		InstanceName:   InstanceName,
		InstanceNumber: InstanceNumber,
		Node:           Node,
	}
}

// OracleRACSourceConnectionInfoFactory is just a simple function to instantiate the OracleRACSourceConnectionInfoStruct
func OracleRACSourceConnectionInfoFactory(
	CrsClusterHome string,
	DatabaseName string,
	JdbcStrings []string,
	Nodes []string,
	OracleHome string,
	RemoteListener string,
	Scan string,
	Version string,
) OracleRACSourceConnectionInfoStruct {
	return OracleRACSourceConnectionInfoStruct{
		Type:           "OracleRACSourceConnectionInfo",
		CrsClusterHome: CrsClusterHome,
		DatabaseName:   DatabaseName,
		JdbcStrings:    JdbcStrings,
		Nodes:          Nodes,
		OracleHome:     OracleHome,
		RemoteListener: RemoteListener,
		Scan:           Scan,
		Version:        Version,
	}
}

// OracleRedoLogFileSpecificationFactory is just a simple function to instantiate the OracleRedoLogFileSpecificationStruct
func OracleRedoLogFileSpecificationFactory(
	Filename string,
	Size *int,
) OracleRedoLogFileSpecificationStruct {
	return OracleRedoLogFileSpecificationStruct{
		Type:     "OracleRedoLogFileSpecification",
		Filename: Filename,
		Size:     Size,
	}
}

// OracleRefreshParametersFactory is just a simple function to instantiate the OracleRefreshParametersStruct
func OracleRefreshParametersFactory(
	Credential Credential,
	TimeflowPointParameters TimeflowPointParameters,
	Username string,
) OracleRefreshParametersStruct {
	return OracleRefreshParametersStruct{
		Type:                    "OracleRefreshParameters",
		Credential:              Credential,
		TimeflowPointParameters: TimeflowPointParameters,
		Username:                Username,
	}
}

// OracleRollbackParametersFactory is just a simple function to instantiate the OracleRollbackParametersStruct
func OracleRollbackParametersFactory(
	Credential Credential,
	TimeflowPointParameters TimeflowPointParameters,
	Username string,
) OracleRollbackParametersStruct {
	return OracleRollbackParametersStruct{
		Type:                    "OracleRollbackParameters",
		Credential:              Credential,
		TimeflowPointParameters: TimeflowPointParameters,
		Username:                Username,
	}
}

// OracleSIConfigFactory is just a simple function to instantiate the OracleSIConfigStruct
func OracleSIConfigFactory(
	Credentials Credential,
	DatabaseName string,
	Discovered *bool,
	EnvironmentUser string,
	Instance *OracleInstanceStruct,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	NonSysCredentials string,
	NonSysUser string,
	Reference string,
	Repository string,
	Services []*OracleServiceStruct,
	TdeKeystorePassword string,
	UniqueName string,
	User string,
) OracleSIConfigStruct {
	return OracleSIConfigStruct{
		Type:                "OracleSIConfig",
		Credentials:         Credentials,
		DatabaseName:        DatabaseName,
		Discovered:          Discovered,
		EnvironmentUser:     EnvironmentUser,
		Instance:            Instance,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		Namespace:           Namespace,
		NonSysCredentials:   NonSysCredentials,
		NonSysUser:          NonSysUser,
		Reference:           Reference,
		Repository:          Repository,
		Services:            Services,
		TdeKeystorePassword: TdeKeystorePassword,
		UniqueName:          UniqueName,
		User:                User,
	}
}

// OracleSISourceConnectionInfoFactory is just a simple function to instantiate the OracleSISourceConnectionInfoStruct
func OracleSISourceConnectionInfoFactory(
	DatabaseName string,
	Host string,
	JdbcStrings []string,
	OracleHome string,
	Version string,
) OracleSISourceConnectionInfoStruct {
	return OracleSISourceConnectionInfoStruct{
		Type:         "OracleSISourceConnectionInfo",
		DatabaseName: DatabaseName,
		Host:         Host,
		JdbcStrings:  JdbcStrings,
		OracleHome:   OracleHome,
		Version:      Version,
	}
}

// OracleSTConvertedToPDBAttachDataFactory is just a simple function to instantiate the OracleSTConvertedToPDBAttachDataStruct
func OracleSTConvertedToPDBAttachDataFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	Config string,
	EncryptedLinkingEnabled *bool,
	EnvironmentUser string,
	ExternalFilePath string,
	FilesPerSet *int,
	Force *bool,
	LinkNow *bool,
	NewPdbContainerName string,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	OracleFallbackCredentials Credential,
	OracleFallbackUser string,
	RmanChannels *int,
) OracleSTConvertedToPDBAttachDataStruct {
	return OracleSTConvertedToPDBAttachDataStruct{
		Type:                      "OracleSTConvertedToPDBAttachData",
		BackupLevelEnabled:        BackupLevelEnabled,
		BandwidthLimit:            BandwidthLimit,
		CheckLogical:              CheckLogical,
		CompressedLinkingEnabled:  CompressedLinkingEnabled,
		Config:                    Config,
		EncryptedLinkingEnabled:   EncryptedLinkingEnabled,
		EnvironmentUser:           EnvironmentUser,
		ExternalFilePath:          ExternalFilePath,
		FilesPerSet:               FilesPerSet,
		Force:                     Force,
		LinkNow:                   LinkNow,
		NewPdbContainerName:       NewPdbContainerName,
		NumberOfConnections:       NumberOfConnections,
		Operations:                Operations,
		OracleFallbackCredentials: OracleFallbackCredentials,
		OracleFallbackUser:        OracleFallbackUser,
		RmanChannels:              RmanChannels,
	}
}

// OracleScanListenerFactory is just a simple function to instantiate the OracleScanListenerStruct
func OracleScanListenerFactory(
	Discovered *bool,
	Environment string,
	Name string,
	Namespace string,
	ProtocolAddresses []string,
	Reference string,
) OracleScanListenerStruct {
	return OracleScanListenerStruct{
		Type:              "OracleScanListener",
		Discovered:        Discovered,
		Environment:       Environment,
		Name:              Name,
		Namespace:         Namespace,
		ProtocolAddresses: ProtocolAddresses,
		Reference:         Reference,
	}
}

// OracleServiceFactory is just a simple function to instantiate the OracleServiceStruct
func OracleServiceFactory(
	Discovered *bool,
	JdbcConnectionString string,
) OracleServiceStruct {
	return OracleServiceStruct{
		Type:                 "OracleService",
		Discovered:           Discovered,
		JdbcConnectionString: JdbcConnectionString,
	}
}

// OracleSnapshotFactory is just a simple function to instantiate the OracleSnapshotStruct
func OracleSnapshotFactory(
	Consistency string,
	Container string,
	CreationTime string,
	FirstChangePoint *OracleTimeflowPointStruct,
	FromPhysicalStandbyVdb *bool,
	LatestChangePoint *OracleTimeflowPointStruct,
	MissingNonLoggedData *bool,
	Name string,
	Namespace string,
	RedoLogSizeInBytes *int,
	Reference string,
	Retention *int,
	Runtime *OracleSnapshotRuntimeStruct,
	Temporary *bool,
	Timeflow string,
	Timezone string,
	Version string,
) OracleSnapshotStruct {
	return OracleSnapshotStruct{
		Type:                   "OracleSnapshot",
		Consistency:            Consistency,
		Container:              Container,
		CreationTime:           CreationTime,
		FirstChangePoint:       FirstChangePoint,
		FromPhysicalStandbyVdb: FromPhysicalStandbyVdb,
		LatestChangePoint:      LatestChangePoint,
		MissingNonLoggedData:   MissingNonLoggedData,
		Name:                   Name,
		Namespace:              Namespace,
		RedoLogSizeInBytes:     RedoLogSizeInBytes,
		Reference:              Reference,
		Retention:              Retention,
		Runtime:                Runtime,
		Temporary:              Temporary,
		Timeflow:               Timeflow,
		Timezone:               Timezone,
		Version:                Version,
	}
}

// OracleSnapshotRuntimeFactory is just a simple function to instantiate the OracleSnapshotRuntimeStruct
func OracleSnapshotRuntimeFactory(
	MissingLogs []*OracleLogStruct,
	Provisionable *bool,
) OracleSnapshotRuntimeStruct {
	return OracleSnapshotRuntimeStruct{
		Type:          "OracleSnapshotRuntime",
		MissingLogs:   MissingLogs,
		Provisionable: Provisionable,
	}
}

// OracleSourceRuntimeFactory is just a simple function to instantiate the OracleSourceRuntimeStruct
func OracleSourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	ActiveInstances []*OracleActiveInstanceStruct,
	ArchivelogEnabled *bool,
	BctEnabled *bool,
	DatabaseMode string,
	DatabaseRole string,
	DatabaseSize float64,
	DatabaseStats []*OracleDatabaseStatsSectionStruct,
	DnfsEnabled *bool,
	Enabled string,
	LastNonLoggedLocation string,
	NotAccessibleReason string,
	RacEnabled *bool,
	Status string,
) OracleSourceRuntimeStruct {
	return OracleSourceRuntimeStruct{
		Type:                  "OracleSourceRuntime",
		Accessible:            Accessible,
		AccessibleTimestamp:   AccessibleTimestamp,
		ActiveInstances:       ActiveInstances,
		ArchivelogEnabled:     ArchivelogEnabled,
		BctEnabled:            BctEnabled,
		DatabaseMode:          DatabaseMode,
		DatabaseRole:          DatabaseRole,
		DatabaseSize:          DatabaseSize,
		DatabaseStats:         DatabaseStats,
		DnfsEnabled:           DnfsEnabled,
		Enabled:               Enabled,
		LastNonLoggedLocation: LastNonLoggedLocation,
		NotAccessibleReason:   NotAccessibleReason,
		RacEnabled:            RacEnabled,
		Status:                Status,
	}
}

// OracleSourcingPolicyFactory is just a simple function to instantiate the OracleSourcingPolicyStruct
func OracleSourcingPolicyFactory(
	LogsyncEnabled *bool,
	LogsyncInterval *int,
	LogsyncMode string,
) OracleSourcingPolicyStruct {
	return OracleSourcingPolicyStruct{
		Type:            "OracleSourcingPolicy",
		LogsyncEnabled:  LogsyncEnabled,
		LogsyncInterval: LogsyncInterval,
		LogsyncMode:     LogsyncMode,
	}
}

// OracleStagingSourceFactory is just a simple function to instantiate the OracleStagingSourceStruct
func OracleStagingSourceFactory(
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	NodeListeners []string,
	Reference string,
	Runtime OracleBaseSourceRuntime,
	Staging *bool,
	Status string,
	Virtual *bool,
) OracleStagingSourceStruct {
	return OracleStagingSourceStruct{
		Type:                 "OracleStagingSource",
		Config:               Config,
		ConfigParams:         ConfigParams,
		ConfigTemplate:       ConfigTemplate,
		Container:            Container,
		Description:          Description,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		MountBase:            MountBase,
		Name:                 Name,
		Namespace:            Namespace,
		NodeListeners:        NodeListeners,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		Status:               Status,
		Virtual:              Virtual,
	}
}

// OracleStagingSourceParametersFactory is just a simple function to instantiate the OracleStagingSourceParametersStruct
func OracleStagingSourceParametersFactory(
	ConfigParams map[string]string,
	ConfigTemplate string,
	EnvironmentUser string,
	MountBase string,
	Repository string,
) OracleStagingSourceParametersStruct {
	return OracleStagingSourceParametersStruct{
		Type:            "OracleStagingSourceParameters",
		ConfigParams:    ConfigParams,
		ConfigTemplate:  ConfigTemplate,
		EnvironmentUser: EnvironmentUser,
		MountBase:       MountBase,
		Repository:      Repository,
	}
}

// OracleStartParametersFactory is just a simple function to instantiate the OracleStartParametersStruct
func OracleStartParametersFactory(
	Credential Credential,
	Instances []*int,
	Username string,
) OracleStartParametersStruct {
	return OracleStartParametersStruct{
		Type:       "OracleStartParameters",
		Credential: Credential,
		Instances:  Instances,
		Username:   Username,
	}
}

// OracleStopParametersFactory is just a simple function to instantiate the OracleStopParametersStruct
func OracleStopParametersFactory(
	Abort *bool,
	Credential Credential,
	Instances []*int,
	Username string,
) OracleStopParametersStruct {
	return OracleStopParametersStruct{
		Type:       "OracleStopParameters",
		Abort:      Abort,
		Credential: Credential,
		Instances:  Instances,
		Username:   Username,
	}
}

// OracleSyncFromExternalParametersFactory is just a simple function to instantiate the OracleSyncFromExternalParametersStruct
func OracleSyncFromExternalParametersFactory(
	DoNotResume *bool,
	DoubleSync *bool,
	FilesForFullBackup []*int,
	ForceFullBackup *bool,
	SkipSpaceCheck *bool,
) OracleSyncFromExternalParametersStruct {
	return OracleSyncFromExternalParametersStruct{
		Type:               "OracleSyncFromExternalParameters",
		DoNotResume:        DoNotResume,
		DoubleSync:         DoubleSync,
		FilesForFullBackup: FilesForFullBackup,
		ForceFullBackup:    ForceFullBackup,
		SkipSpaceCheck:     SkipSpaceCheck,
	}
}

// OracleSysauxDatafileSpecificationFactory is just a simple function to instantiate the OracleSysauxDatafileSpecificationStruct
func OracleSysauxDatafileSpecificationFactory(
	AutoExtend *bool,
	AutoExtendIncrement *int,
	Filename string,
	MaxSize *int,
	Size *int,
) OracleSysauxDatafileSpecificationStruct {
	return OracleSysauxDatafileSpecificationStruct{
		Type:                "OracleSysauxDatafileSpecification",
		AutoExtend:          AutoExtend,
		AutoExtendIncrement: AutoExtendIncrement,
		Filename:            Filename,
		MaxSize:             MaxSize,
		Size:                Size,
	}
}

// OracleSystemDatafileSpecificationFactory is just a simple function to instantiate the OracleSystemDatafileSpecificationStruct
func OracleSystemDatafileSpecificationFactory(
	AutoExtend *bool,
	AutoExtendIncrement *int,
	Filename string,
	MaxSize *int,
	Size *int,
) OracleSystemDatafileSpecificationStruct {
	return OracleSystemDatafileSpecificationStruct{
		Type:                "OracleSystemDatafileSpecification",
		AutoExtend:          AutoExtend,
		AutoExtendIncrement: AutoExtendIncrement,
		Filename:            Filename,
		MaxSize:             MaxSize,
		Size:                Size,
	}
}

// OracleTempfileSpecificationFactory is just a simple function to instantiate the OracleTempfileSpecificationStruct
func OracleTempfileSpecificationFactory(
	AutoExtend *bool,
	AutoExtendIncrement *int,
	Filename string,
	MaxSize *int,
	Size *int,
) OracleTempfileSpecificationStruct {
	return OracleTempfileSpecificationStruct{
		Type:                "OracleTempfileSpecification",
		AutoExtend:          AutoExtend,
		AutoExtendIncrement: AutoExtendIncrement,
		Filename:            Filename,
		MaxSize:             MaxSize,
		Size:                Size,
	}
}

// OracleTimeflowFactory is just a simple function to instantiate the OracleTimeflowStruct
func OracleTimeflowFactory(
	CdbTimeflow string,
	Container string,
	CreationType string,
	IncarnationID string,
	Name string,
	Namespace string,
	ParentPoint *OracleTimeflowPointStruct,
	ParentSnapshot string,
	Reference string,
) OracleTimeflowStruct {
	return OracleTimeflowStruct{
		Type:           "OracleTimeflow",
		CdbTimeflow:    CdbTimeflow,
		Container:      Container,
		CreationType:   CreationType,
		IncarnationID:  IncarnationID,
		Name:           Name,
		Namespace:      Namespace,
		ParentPoint:    ParentPoint,
		ParentSnapshot: ParentSnapshot,
		Reference:      Reference,
	}
}

// OracleTimeflowPointFactory is just a simple function to instantiate the OracleTimeflowPointStruct
func OracleTimeflowPointFactory(
	Location string,
	Timeflow string,
	Timestamp string,
) OracleTimeflowPointStruct {
	return OracleTimeflowPointStruct{
		Type:      "OracleTimeflowPoint",
		Location:  Location,
		Timeflow:  Timeflow,
		Timestamp: Timestamp,
	}
}

// OracleUndoDatafileSpecificationFactory is just a simple function to instantiate the OracleUndoDatafileSpecificationStruct
func OracleUndoDatafileSpecificationFactory(
	AutoExtend *bool,
	AutoExtendIncrement *int,
	Filename string,
	MaxSize *int,
	Size *int,
) OracleUndoDatafileSpecificationStruct {
	return OracleUndoDatafileSpecificationStruct{
		Type:                "OracleUndoDatafileSpecification",
		AutoExtend:          AutoExtend,
		AutoExtendIncrement: AutoExtendIncrement,
		Filename:            Filename,
		MaxSize:             MaxSize,
		Size:                Size,
	}
}

// OracleVirtualCdbProvisionParametersFactory is just a simple function to instantiate the OracleVirtualCdbProvisionParametersStruct
func OracleVirtualCdbProvisionParametersFactory(
	Container *OracleDatabaseContainerStruct,
	Source *OracleVirtualCdbSourceStruct,
	SourceConfig OracleDBConfig,
) OracleVirtualCdbProvisionParametersStruct {
	return OracleVirtualCdbProvisionParametersStruct{
		Type:         "OracleVirtualCdbProvisionParameters",
		Container:    Container,
		Source:       Source,
		SourceConfig: SourceConfig,
	}
}

// OracleVirtualCdbSourceFactory is just a simple function to instantiate the OracleVirtualCdbSourceStruct
func OracleVirtualCdbSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	Container string,
	CustomEnvVars []OracleCustomEnvVar,
	Description string,
	FileMappingRules string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
	Reference string,
	Runtime OracleBaseSourceRuntime,
	Staging *bool,
	Status string,
	Virtual *bool,
) OracleVirtualCdbSourceStruct {
	return OracleVirtualCdbSourceStruct{
		Type:                            "OracleVirtualCdbSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		Container:                       Container,
		CustomEnvVars:                   CustomEnvVars,
		Description:                     Description,
		FileMappingRules:                FileMappingRules,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Virtual:                         Virtual,
	}
}

// OracleVirtualIPFactory is just a simple function to instantiate the OracleVirtualIPStruct
func OracleVirtualIPFactory(
	Discovered *bool,
	DomainName string,
	Ip string,
) OracleVirtualIPStruct {
	return OracleVirtualIPStruct{
		Type:       "OracleVirtualIP",
		Discovered: Discovered,
		DomainName: DomainName,
		Ip:         Ip,
	}
}

// OracleVirtualPdbSourceFactory is just a simple function to instantiate the OracleVirtualPdbSourceStruct
func OracleVirtualPdbSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	Container string,
	CustomEnvVars []OracleCustomEnvVar,
	Description string,
	FileMappingRules string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	ParentTdeKeystorePassword string,
	ParentTdeKeystorePath string,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
	Reference string,
	Runtime OracleBaseSourceRuntime,
	Staging *bool,
	Status string,
	TdeExportedKeyFileSecret string,
	Virtual *bool,
) OracleVirtualPdbSourceStruct {
	return OracleVirtualPdbSourceStruct{
		Type:                            "OracleVirtualPdbSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		Container:                       Container,
		CustomEnvVars:                   CustomEnvVars,
		Description:                     Description,
		FileMappingRules:                FileMappingRules,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		ParentTdeKeystorePassword:       ParentTdeKeystorePassword,
		ParentTdeKeystorePath:           ParentTdeKeystorePath,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		TdeExportedKeyFileSecret:        TdeExportedKeyFileSecret,
		Virtual:                         Virtual,
	}
}

// OracleVirtualSourceFactory is just a simple function to instantiate the OracleVirtualSourceStruct
func OracleVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	Container string,
	CustomEnvVars []OracleCustomEnvVar,
	Description string,
	FileMappingRules string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
	Reference string,
	Runtime OracleBaseSourceRuntime,
	Staging *bool,
	Status string,
	Virtual *bool,
) OracleVirtualSourceStruct {
	return OracleVirtualSourceStruct{
		Type:                            "OracleVirtualSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		Container:                       Container,
		CustomEnvVars:                   CustomEnvVars,
		Description:                     Description,
		FileMappingRules:                FileMappingRules,
		Hosts:                           Hosts,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Virtual:                         Virtual,
	}
}

// PasswordCredentialFactory is just a simple function to instantiate the PasswordCredentialStruct
func PasswordCredentialFactory(
	Password string,
) PasswordCredentialStruct {
	return PasswordCredentialStruct{
		Type:     "PasswordCredential",
		Password: Password,
	}
}

// PasswordPolicyFactory is just a simple function to instantiate the PasswordPolicyStruct
func PasswordPolicyFactory(
	Digit *bool,
	DisallowUsernameAsPassword *bool,
	LowercaseLetter *bool,
	MinLength *int,
	Name string,
	Namespace string,
	Reference string,
	ReuseDisallowLimit *int,
	Symbol *bool,
	UppercaseLetter *bool,
) PasswordPolicyStruct {
	return PasswordPolicyStruct{
		Type:                       "PasswordPolicy",
		Digit:                      Digit,
		DisallowUsernameAsPassword: DisallowUsernameAsPassword,
		LowercaseLetter:            LowercaseLetter,
		MinLength:                  MinLength,
		Name:                       Name,
		Namespace:                  Namespace,
		Reference:                  Reference,
		ReuseDisallowLimit:         ReuseDisallowLimit,
		Symbol:                     Symbol,
		UppercaseLetter:            UppercaseLetter,
	}
}

// PathDescendantConstraintFactory is just a simple function to instantiate the PathDescendantConstraintStruct
func PathDescendantConstraintFactory(
	AxisName string,
	DescendantOf string,
) PathDescendantConstraintStruct {
	return PathDescendantConstraintStruct{
		Type:         "PathDescendantConstraint",
		AxisName:     AxisName,
		DescendantOf: DescendantOf,
	}
}

// PemCertificateFactory is just a simple function to instantiate the PemCertificateStruct
func PemCertificateFactory(
	Contents string,
) PemCertificateStruct {
	return PemCertificateStruct{
		Type:     "PemCertificate",
		Contents: Contents,
	}
}

// PemCertificateChainFactory is just a simple function to instantiate the PemCertificateChainStruct
func PemCertificateChainFactory(
	Chain []*PemCertificateStruct,
) PemCertificateChainStruct {
	return PemCertificateChainStruct{
		Type:  "PemCertificateChain",
		Chain: Chain,
	}
}

// PemClientCertificateFactory is just a simple function to instantiate the PemClientCertificateStruct
func PemClientCertificateFactory(
	ClientCertificateChain *PemCertificateChainStruct,
	PrivateKey string,
) PemClientCertificateStruct {
	return PemClientCertificateStruct{
		Type:                   "PemClientCertificate",
		ClientCertificateChain: ClientCertificateChain,
		PrivateKey:             PrivateKey,
	}
}

// PermissionFactory is just a simple function to instantiate the PermissionStruct
func PermissionFactory(
	ActionType string,
	Name string,
	Namespace string,
	Reference string,
) PermissionStruct {
	return PermissionStruct{
		Type:       "Permission",
		ActionType: ActionType,
		Name:       Name,
		Namespace:  Namespace,
		Reference:  Reference,
	}
}

// PgSQLAttachDataFactory is just a simple function to instantiate the PgSQLAttachDataStruct
func PgSQLAttachDataFactory(
	Config string,
	ConnectionDatabase string,
	DbCredentials *PasswordCredentialStruct,
	DbUser string,
	ExternalFilePath string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
	PptRepository string,
) PgSQLAttachDataStruct {
	return PgSQLAttachDataStruct{
		Type:               "PgSQLAttachData",
		Config:             Config,
		ConnectionDatabase: ConnectionDatabase,
		DbCredentials:      DbCredentials,
		DbUser:             DbUser,
		ExternalFilePath:   ExternalFilePath,
		Operations:         Operations,
		PptHostUser:        PptHostUser,
		PptRepository:      PptRepository,
	}
}

// PgSQLCompatibilityCriteriaFactory is just a simple function to instantiate the PgSQLCompatibilityCriteriaStruct
func PgSQLCompatibilityCriteriaFactory(
	Architecture *int,
	Environment string,
	Os string,
	Processor string,
	SegmentSize *int,
	StagingEnabled *bool,
	Variant string,
	Version string,
) PgSQLCompatibilityCriteriaStruct {
	return PgSQLCompatibilityCriteriaStruct{
		Type:           "PgSQLCompatibilityCriteria",
		Architecture:   Architecture,
		Environment:    Environment,
		Os:             Os,
		Processor:      Processor,
		SegmentSize:    SegmentSize,
		StagingEnabled: StagingEnabled,
		Variant:        Variant,
		Version:        Version,
	}
}

// PgSQLCreateTransformationParametersFactory is just a simple function to instantiate the PgSQLCreateTransformationParametersStruct
func PgSQLCreateTransformationParametersFactory(
	Container *PgSQLDatabaseContainerStruct,
	EnvironmentUser string,
	Operations []SourceOperation,
	Port *int,
	Repository string,
) PgSQLCreateTransformationParametersStruct {
	return PgSQLCreateTransformationParametersStruct{
		Type:            "PgSQLCreateTransformationParameters",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Port:            Port,
		Repository:      Repository,
	}
}

// PgSQLDBClusterConfigFactory is just a simple function to instantiate the PgSQLDBClusterConfigStruct
func PgSQLDBClusterConfigFactory(
	ClusterDataDirectory string,
	ConnectionDatabase string,
	Credentials string,
	Discovered *bool,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Namespace string,
	Port *int,
	Reference string,
	Repository string,
	User string,
) PgSQLDBClusterConfigStruct {
	return PgSQLDBClusterConfigStruct{
		Type:                 "PgSQLDBClusterConfig",
		ClusterDataDirectory: ClusterDataDirectory,
		ConnectionDatabase:   ConnectionDatabase,
		Credentials:          Credentials,
		Discovered:           Discovered,
		EnvironmentUser:      EnvironmentUser,
		LinkingEnabled:       LinkingEnabled,
		Name:                 Name,
		Namespace:            Namespace,
		Port:                 Port,
		Reference:            Reference,
		Repository:           Repository,
		User:                 User,
	}
}

// PgSQLDBConfigFactory is just a simple function to instantiate the PgSQLDBConfigStruct
func PgSQLDBConfigFactory(
	Credentials Credential,
	DatabaseCluster string,
	DatabaseName string,
	User string,
) PgSQLDBConfigStruct {
	return PgSQLDBConfigStruct{
		Type:            "PgSQLDBConfig",
		Credentials:     Credentials,
		DatabaseCluster: DatabaseCluster,
		DatabaseName:    DatabaseName,
		User:            User,
	}
}

// PgSQLDBContainerRuntimeFactory is just a simple function to instantiate the PgSQLDBContainerRuntimeStruct
func PgSQLDBContainerRuntimeFactory(
	LastRestoredWALSegment string,
	LogSyncActive *bool,
	PreProvisioningStatus *PreProvisioningRuntimeStruct,
) PgSQLDBContainerRuntimeStruct {
	return PgSQLDBContainerRuntimeStruct{
		Type:                   "PgSQLDBContainerRuntime",
		LastRestoredWALSegment: LastRestoredWALSegment,
		LogSyncActive:          LogSyncActive,
		PreProvisioningStatus:  PreProvisioningStatus,
	}
}

// PgSQLDatabaseContainerFactory is just a simple function to instantiate the PgSQLDatabaseContainerStruct
func PgSQLDatabaseContainerFactory(
	CreationTime string,
	CurrentTimeflow string,
	Description string,
	Group string,
	Masked *bool,
	Name string,
	Namespace string,
	Os string,
	PreviousTimeflow string,
	Processor string,
	ProvisionContainer string,
	Reference string,
	Runtime *PgSQLDBContainerRuntimeStruct,
	SourcingPolicy *SourcingPolicyStruct,
	Transformation *bool,
) PgSQLDatabaseContainerStruct {
	return PgSQLDatabaseContainerStruct{
		Type:               "PgSQLDatabaseContainer",
		CreationTime:       CreationTime,
		CurrentTimeflow:    CurrentTimeflow,
		Description:        Description,
		Group:              Group,
		Masked:             Masked,
		Name:               Name,
		Namespace:          Namespace,
		Os:                 Os,
		PreviousTimeflow:   PreviousTimeflow,
		Processor:          Processor,
		ProvisionContainer: ProvisionContainer,
		Reference:          Reference,
		Runtime:            Runtime,
		SourcingPolicy:     SourcingPolicy,
		Transformation:     Transformation,
	}
}

// PgSQLExportParametersFactory is just a simple function to instantiate the PgSQLExportParametersStruct
func PgSQLExportParametersFactory(
	ConfigParams map[string]string,
	FileMappingRules string,
	FilesystemLayout *TimeflowFilesystemLayoutStruct,
	HbaEntries []*PgSQLHBAEntryStruct,
	IdentEntries []*PgSQLIdentEntryStruct,
	RecoverDatabase *bool,
	SourceConfig *PgSQLDBClusterConfigStruct,
	TimeflowPointParameters TimeflowPointParameters,
) PgSQLExportParametersStruct {
	return PgSQLExportParametersStruct{
		Type:                    "PgSQLExportParameters",
		ConfigParams:            ConfigParams,
		FileMappingRules:        FileMappingRules,
		FilesystemLayout:        FilesystemLayout,
		HbaEntries:              HbaEntries,
		IdentEntries:            IdentEntries,
		RecoverDatabase:         RecoverDatabase,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// PgSQLHBAEntryFactory is just a simple function to instantiate the PgSQLHBAEntryStruct
func PgSQLHBAEntryFactory(
	Address string,
	AuthMethod string,
	AuthOptions string,
	Database string,
	EntryType string,
	User string,
) PgSQLHBAEntryStruct {
	return PgSQLHBAEntryStruct{
		Type:        "PgSQLHBAEntry",
		Address:     Address,
		AuthMethod:  AuthMethod,
		AuthOptions: AuthOptions,
		Database:    Database,
		EntryType:   EntryType,
		User:        User,
	}
}

// PgSQLIdentEntryFactory is just a simple function to instantiate the PgSQLIdentEntryStruct
func PgSQLIdentEntryFactory(
	DatabaseUsername string,
	MapName string,
	SystemUsername string,
) PgSQLIdentEntryStruct {
	return PgSQLIdentEntryStruct{
		Type:             "PgSQLIdentEntry",
		DatabaseUsername: DatabaseUsername,
		MapName:          MapName,
		SystemUsername:   SystemUsername,
	}
}

// PgSQLInstallFactory is just a simple function to instantiate the PgSQLInstallStruct
func PgSQLInstallFactory(
	Bits *int,
	Discovered *bool,
	Environment string,
	InstallationPath string,
	Name string,
	Namespace string,
	ProvisioningEnabled *bool,
	Reference string,
	SegmentSize *int,
	Staging *bool,
	Variant string,
	Version string,
) PgSQLInstallStruct {
	return PgSQLInstallStruct{
		Type:                "PgSQLInstall",
		Bits:                Bits,
		Discovered:          Discovered,
		Environment:         Environment,
		InstallationPath:    InstallationPath,
		Name:                Name,
		Namespace:           Namespace,
		ProvisioningEnabled: ProvisioningEnabled,
		Reference:           Reference,
		SegmentSize:         SegmentSize,
		Staging:             Staging,
		Variant:             Variant,
		Version:             Version,
	}
}

// PgSQLLinkDataFactory is just a simple function to instantiate the PgSQLLinkDataStruct
func PgSQLLinkDataFactory(
	Config string,
	ConnectionDatabase string,
	DbCredentials *PasswordCredentialStruct,
	DbUser string,
	ExternalFilePath string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
	PptRepository string,
	SourcingPolicy *SourcingPolicyStruct,
) PgSQLLinkDataStruct {
	return PgSQLLinkDataStruct{
		Type:               "PgSQLLinkData",
		Config:             Config,
		ConnectionDatabase: ConnectionDatabase,
		DbCredentials:      DbCredentials,
		DbUser:             DbUser,
		ExternalFilePath:   ExternalFilePath,
		Operations:         Operations,
		PptHostUser:        PptHostUser,
		PptRepository:      PptRepository,
		SourcingPolicy:     SourcingPolicy,
	}
}

// PgSQLLinkedSourceFactory is just a simple function to instantiate the PgSQLLinkedSourceStruct
func PgSQLLinkedSourceFactory(
	Config string,
	Container string,
	Description string,
	ExternalFilePath string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	Operations *LinkedSourceOperationsStruct,
	Reference string,
	Runtime *PgSQLSourceRuntimeStruct,
	Staging *bool,
	StagingSource string,
	Status string,
	Virtual *bool,
) PgSQLLinkedSourceStruct {
	return PgSQLLinkedSourceStruct{
		Type:                 "PgSQLLinkedSource",
		Config:               Config,
		Container:            Container,
		Description:          Description,
		ExternalFilePath:     ExternalFilePath,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Namespace:            Namespace,
		Operations:           Operations,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		StagingSource:        StagingSource,
		Status:               Status,
		Virtual:              Virtual,
	}
}

// PgSQLPlatformParametersFactory is just a simple function to instantiate the PgSQLPlatformParametersStruct
func PgSQLPlatformParametersFactory(
	Port *int,
) PgSQLPlatformParametersStruct {
	return PgSQLPlatformParametersStruct{
		Type: "PgSQLPlatformParameters",
		Port: Port,
	}
}

// PgSQLProvisionParametersFactory is just a simple function to instantiate the PgSQLProvisionParametersStruct
func PgSQLProvisionParametersFactory(
	Container *PgSQLDatabaseContainerStruct,
	Masked *bool,
	MaskingJob string,
	Source *PgSQLVirtualSourceStruct,
	SourceConfig *PgSQLDBClusterConfigStruct,
	TimeflowPointParameters TimeflowPointParameters,
) PgSQLProvisionParametersStruct {
	return PgSQLProvisionParametersStruct{
		Type:                    "PgSQLProvisionParameters",
		Container:               Container,
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		Source:                  Source,
		SourceConfig:            SourceConfig,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// PgSQLSnapshotFactory is just a simple function to instantiate the PgSQLSnapshotStruct
func PgSQLSnapshotFactory(
	Consistency string,
	Container string,
	CreationTime string,
	FirstChangePoint *PgSQLTimeflowPointStruct,
	LatestChangePoint *PgSQLTimeflowPointStruct,
	MissingNonLoggedData *bool,
	Name string,
	Namespace string,
	Reference string,
	Retention *int,
	Runtime *SnapshotRuntimeStruct,
	Temporary *bool,
	Timeflow string,
	Timezone string,
	Version string,
) PgSQLSnapshotStruct {
	return PgSQLSnapshotStruct{
		Type:                 "PgSQLSnapshot",
		Consistency:          Consistency,
		Container:            Container,
		CreationTime:         CreationTime,
		FirstChangePoint:     FirstChangePoint,
		LatestChangePoint:    LatestChangePoint,
		MissingNonLoggedData: MissingNonLoggedData,
		Name:                 Name,
		Namespace:            Namespace,
		Reference:            Reference,
		Retention:            Retention,
		Runtime:              Runtime,
		Temporary:            Temporary,
		Timeflow:             Timeflow,
		Timezone:             Timezone,
		Version:              Version,
	}
}

// PgSQLSourceConnectionInfoFactory is just a simple function to instantiate the PgSQLSourceConnectionInfoStruct
func PgSQLSourceConnectionInfoFactory(
	Host string,
	JdbcString string,
	PgDataDirectory string,
	Port *int,
	User string,
	Version string,
) PgSQLSourceConnectionInfoStruct {
	return PgSQLSourceConnectionInfoStruct{
		Type:            "PgSQLSourceConnectionInfo",
		Host:            Host,
		JdbcString:      JdbcString,
		PgDataDirectory: PgDataDirectory,
		Port:            Port,
		User:            User,
		Version:         Version,
	}
}

// PgSQLSourceRuntimeFactory is just a simple function to instantiate the PgSQLSourceRuntimeStruct
func PgSQLSourceRuntimeFactory(
	Accessible *bool,
	AccessibleTimestamp string,
	DatabaseSize float64,
	Enabled string,
	NotAccessibleReason string,
	Status string,
) PgSQLSourceRuntimeStruct {
	return PgSQLSourceRuntimeStruct{
		Type:                "PgSQLSourceRuntime",
		Accessible:          Accessible,
		AccessibleTimestamp: AccessibleTimestamp,
		DatabaseSize:        DatabaseSize,
		Enabled:             Enabled,
		NotAccessibleReason: NotAccessibleReason,
		Status:              Status,
	}
}

// PgSQLStagingSourceFactory is just a simple function to instantiate the PgSQLStagingSourceStruct
func PgSQLStagingSourceFactory(
	Config string,
	Container string,
	Description string,
	Hosts []string,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	Reference string,
	Runtime *PgSQLSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) PgSQLStagingSourceStruct {
	return PgSQLStagingSourceStruct{
		Type:                 "PgSQLStagingSource",
		Config:               Config,
		Container:            Container,
		Description:          Description,
		Hosts:                Hosts,
		Linked:               Linked,
		LogCollectionEnabled: LogCollectionEnabled,
		MountBase:            MountBase,
		Name:                 Name,
		Namespace:            Namespace,
		Reference:            Reference,
		Runtime:              Runtime,
		Staging:              Staging,
		Status:               Status,
		Virtual:              Virtual,
	}
}

// PgSQLSyncParametersFactory is just a simple function to instantiate the PgSQLSyncParametersStruct
func PgSQLSyncParametersFactory(
	RedoBaseBackup *bool,
) PgSQLSyncParametersStruct {
	return PgSQLSyncParametersStruct{
		Type:           "PgSQLSyncParameters",
		RedoBaseBackup: RedoBaseBackup,
	}
}

// PgSQLTimeflowFactory is just a simple function to instantiate the PgSQLTimeflowStruct
func PgSQLTimeflowFactory(
	Container string,
	CreationType string,
	Name string,
	Namespace string,
	ParentPoint *PgSQLTimeflowPointStruct,
	ParentSnapshot string,
	Reference string,
) PgSQLTimeflowStruct {
	return PgSQLTimeflowStruct{
		Type:           "PgSQLTimeflow",
		Container:      Container,
		CreationType:   CreationType,
		Name:           Name,
		Namespace:      Namespace,
		ParentPoint:    ParentPoint,
		ParentSnapshot: ParentSnapshot,
		Reference:      Reference,
	}
}

// PgSQLTimeflowPointFactory is just a simple function to instantiate the PgSQLTimeflowPointStruct
func PgSQLTimeflowPointFactory(
	Location string,
	Timeflow string,
	Timestamp string,
) PgSQLTimeflowPointStruct {
	return PgSQLTimeflowPointStruct{
		Type:      "PgSQLTimeflowPoint",
		Location:  Location,
		Timeflow:  Timeflow,
		Timestamp: Timestamp,
	}
}

// PgSQLVirtualSourceFactory is just a simple function to instantiate the PgSQLVirtualSourceStruct
func PgSQLVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	Container string,
	Description string,
	FileMappingRules string,
	HbaEntries []*PgSQLHBAEntryStruct,
	Hosts []string,
	IdentEntries []*PgSQLIdentEntryStruct,
	Linked *bool,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Namespace string,
	Operations *VirtualSourceOperationsStruct,
	Reference string,
	Runtime *PgSQLSourceRuntimeStruct,
	Staging *bool,
	Status string,
	Virtual *bool,
) PgSQLVirtualSourceStruct {
	return PgSQLVirtualSourceStruct{
		Type:                            "PgSQLVirtualSource",
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		Container:                       Container,
		Description:                     Description,
		FileMappingRules:                FileMappingRules,
		HbaEntries:                      HbaEntries,
		Hosts:                           Hosts,
		IdentEntries:                    IdentEntries,
		Linked:                          Linked,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Namespace:                       Namespace,
		Operations:                      Operations,
		Reference:                       Reference,
		Runtime:                         Runtime,
		Staging:                         Staging,
		Status:                          Status,
		Virtual:                         Virtual,
	}
}

// PhoneHomeServiceFactory is just a simple function to instantiate the PhoneHomeServiceStruct
func PhoneHomeServiceFactory(
	Enabled *bool,
) PhoneHomeServiceStruct {
	return PhoneHomeServiceStruct{
		Type:    "PhoneHomeService",
		Enabled: Enabled,
	}
}

// PluginFactory is just a simple function to instantiate the PluginStruct
func PluginFactory(
	BuildApi *APIVersionStruct,
	BuildNumber string,
	DefaultLocale string,
	DiscoveryDefinition *PluginDiscoveryDefinitionStruct,
	EntryPoint string,
	ExtendedStartStopHooks *bool,
	ExternalVersion string,
	HostTypes []string,
	Language string,
	LinkedSourceDefinition PluginLinkedSourceDefinition,
	LuaName string,
	Manifest *PluginManifestStruct,
	Messages []*ToolkitLocaleStruct,
	MinimumLuaVersion string,
	Name string,
	Namespace string,
	PluginId string,
	Reference string,
	RootSquashEnabled *bool,
	SnapshotParametersDefinition *PluginSnapshotParametersDefinitionStruct,
	SnapshotSchema *SchemaDraftV4Struct,
	SourceCode string,
	UpgradeDefinition *PluginUpgradeDefinitionStruct,
	VirtualSourceDefinition *PluginVirtualSourceDefinitionStruct,
) PluginStruct {
	return PluginStruct{
		Type:                         "Plugin",
		BuildApi:                     BuildApi,
		BuildNumber:                  BuildNumber,
		DefaultLocale:                DefaultLocale,
		DiscoveryDefinition:          DiscoveryDefinition,
		EntryPoint:                   EntryPoint,
		ExtendedStartStopHooks:       ExtendedStartStopHooks,
		ExternalVersion:              ExternalVersion,
		HostTypes:                    HostTypes,
		Language:                     Language,
		LinkedSourceDefinition:       LinkedSourceDefinition,
		LuaName:                      LuaName,
		Manifest:                     Manifest,
		Messages:                     Messages,
		MinimumLuaVersion:            MinimumLuaVersion,
		Name:                         Name,
		Namespace:                    Namespace,
		PluginId:                     PluginId,
		Reference:                    Reference,
		RootSquashEnabled:            RootSquashEnabled,
		SnapshotParametersDefinition: SnapshotParametersDefinition,
		SnapshotSchema:               SnapshotSchema,
		SourceCode:                   SourceCode,
		UpgradeDefinition:            UpgradeDefinition,
		VirtualSourceDefinition:      VirtualSourceDefinition,
	}
}

// PluginDiscoveryDefinitionFactory is just a simple function to instantiate the PluginDiscoveryDefinitionStruct
func PluginDiscoveryDefinitionFactory(
	ManualSourceConfigDiscovery *bool,
	RepositoryIdentityFields []string,
	RepositoryNameField string,
	RepositorySchema *SchemaDraftV4Struct,
	SourceConfigIdentityFields []string,
	SourceConfigNameField string,
	SourceConfigSchema *SchemaDraftV4Struct,
) PluginDiscoveryDefinitionStruct {
	return PluginDiscoveryDefinitionStruct{
		Type:                        "PluginDiscoveryDefinition",
		ManualSourceConfigDiscovery: ManualSourceConfigDiscovery,
		RepositoryIdentityFields:    RepositoryIdentityFields,
		RepositoryNameField:         RepositoryNameField,
		RepositorySchema:            RepositorySchema,
		SourceConfigIdentityFields:  SourceConfigIdentityFields,
		SourceConfigNameField:       SourceConfigNameField,
		SourceConfigSchema:          SourceConfigSchema,
	}
}

// PluginLinkedDirectSourceDefinitionFactory is just a simple function to instantiate the PluginLinkedDirectSourceDefinitionStruct
func PluginLinkedDirectSourceDefinitionFactory(
	Parameters *SchemaDraftV4Struct,
) PluginLinkedDirectSourceDefinitionStruct {
	return PluginLinkedDirectSourceDefinitionStruct{
		Type:       "PluginLinkedDirectSourceDefinition",
		Parameters: Parameters,
	}
}

// PluginLinkedStagedSourceDefinitionFactory is just a simple function to instantiate the PluginLinkedStagedSourceDefinitionStruct
func PluginLinkedStagedSourceDefinitionFactory(
	Parameters *SchemaDraftV4Struct,
) PluginLinkedStagedSourceDefinitionStruct {
	return PluginLinkedStagedSourceDefinitionStruct{
		Type:       "PluginLinkedStagedSourceDefinition",
		Parameters: Parameters,
	}
}

// PluginManifestFactory is just a simple function to instantiate the PluginManifestStruct
func PluginManifestFactory(
	HasInitialize *bool,
	HasLinkedMountSpecification *bool,
	HasLinkedPostSnapshot *bool,
	HasLinkedPreSnapshot *bool,
	HasLinkedStartStaging *bool,
	HasLinkedStatus *bool,
	HasLinkedStopStaging *bool,
	HasLinkedWorker *bool,
	HasRepositoryDiscovery *bool,
	HasSourceConfigDiscovery *bool,
	HasVirtualConfigure *bool,
	HasVirtualMountSpecification *bool,
	HasVirtualPostSnapshot *bool,
	HasVirtualPreSnapshot *bool,
	HasVirtualReconfigure *bool,
	HasVirtualStart *bool,
	HasVirtualStatus *bool,
	HasVirtualStop *bool,
	HasVirtualUnconfigure *bool,
	MigrationIdList []string,
) PluginManifestStruct {
	return PluginManifestStruct{
		Type:                         "PluginManifest",
		HasInitialize:                HasInitialize,
		HasLinkedMountSpecification:  HasLinkedMountSpecification,
		HasLinkedPostSnapshot:        HasLinkedPostSnapshot,
		HasLinkedPreSnapshot:         HasLinkedPreSnapshot,
		HasLinkedStartStaging:        HasLinkedStartStaging,
		HasLinkedStatus:              HasLinkedStatus,
		HasLinkedStopStaging:         HasLinkedStopStaging,
		HasLinkedWorker:              HasLinkedWorker,
		HasRepositoryDiscovery:       HasRepositoryDiscovery,
		HasSourceConfigDiscovery:     HasSourceConfigDiscovery,
		HasVirtualConfigure:          HasVirtualConfigure,
		HasVirtualMountSpecification: HasVirtualMountSpecification,
		HasVirtualPostSnapshot:       HasVirtualPostSnapshot,
		HasVirtualPreSnapshot:        HasVirtualPreSnapshot,
		HasVirtualReconfigure:        HasVirtualReconfigure,
		HasVirtualStart:              HasVirtualStart,
		HasVirtualStatus:             HasVirtualStatus,
		HasVirtualStop:               HasVirtualStop,
		HasVirtualUnconfigure:        HasVirtualUnconfigure,
		MigrationIdList:              MigrationIdList,
	}
}

// PluginSnapshotParametersDefinitionFactory is just a simple function to instantiate the PluginSnapshotParametersDefinitionStruct
func PluginSnapshotParametersDefinitionFactory(
	Schema *SchemaDraftV4Struct,
) PluginSnapshotParametersDefinitionStruct {
	return PluginSnapshotParametersDefinitionStruct{
		Type:   "PluginSnapshotParametersDefinition",
		Schema: Schema,
	}
}

// PluginUpgradeDefinitionFactory is just a simple function to instantiate the PluginUpgradeDefinitionStruct
func PluginUpgradeDefinitionFactory() PluginUpgradeDefinitionStruct {
	return PluginUpgradeDefinitionStruct{
		Type: "PluginUpgradeDefinition",
	}
}

// PluginVirtualSourceDefinitionFactory is just a simple function to instantiate the PluginVirtualSourceDefinitionStruct
func PluginVirtualSourceDefinitionFactory(
	Parameters *SchemaDraftV4Struct,
) PluginVirtualSourceDefinitionStruct {
	return PluginVirtualSourceDefinitionStruct{
		Type:       "PluginVirtualSourceDefinition",
		Parameters: Parameters,
	}
}

// PolicyApplyTargetParametersFactory is just a simple function to instantiate the PolicyApplyTargetParametersStruct
func PolicyApplyTargetParametersFactory(
	Target string,
	Targets []string,
) PolicyApplyTargetParametersStruct {
	return PolicyApplyTargetParametersStruct{
		Type:    "PolicyApplyTargetParameters",
		Target:  Target,
		Targets: Targets,
	}
}

// PolicyCreateAndApplyParametersFactory is just a simple function to instantiate the PolicyCreateAndApplyParametersStruct
func PolicyCreateAndApplyParametersFactory(
	Policy Policy,
	Target string,
) PolicyCreateAndApplyParametersStruct {
	return PolicyCreateAndApplyParametersStruct{
		Type:   "PolicyCreateAndApplyParameters",
		Policy: Policy,
		Target: Target,
	}
}

// PreProvisioningRuntimeFactory is just a simple function to instantiate the PreProvisioningRuntimeStruct
func PreProvisioningRuntimeFactory(
	LastUpdateTimestamp string,
	PendingAction string,
	PreProvisioningState string,
	Response string,
	Status string,
) PreProvisioningRuntimeStruct {
	return PreProvisioningRuntimeStruct{
		Type:                 "PreProvisioningRuntime",
		LastUpdateTimestamp:  LastUpdateTimestamp,
		PendingAction:        PendingAction,
		PreProvisioningState: PreProvisioningState,
		Response:             Response,
		Status:               Status,
	}
}

// ProvisionCompatibilityParametersFactory is just a simple function to instantiate the ProvisionCompatibilityParametersStruct
func ProvisionCompatibilityParametersFactory(
	Environment string,
	TimeflowPointParameters TimeflowPointParameters,
) ProvisionCompatibilityParametersStruct {
	return ProvisionCompatibilityParametersStruct{
		Type:                    "ProvisionCompatibilityParameters",
		Environment:             Environment,
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// ProxyConfigurationFactory is just a simple function to instantiate the ProxyConfigurationStruct
func ProxyConfigurationFactory(
	Enabled *bool,
	Host string,
	Password string,
	Port *int,
	Username string,
) ProxyConfigurationStruct {
	return ProxyConfigurationStruct{
		Type:     "ProxyConfiguration",
		Enabled:  Enabled,
		Host:     Host,
		Password: Password,
		Port:     Port,
		Username: Username,
	}
}

// ProxyServiceFactory is just a simple function to instantiate the ProxyServiceStruct
func ProxyServiceFactory(
	Https *ProxyConfigurationStruct,
	Socks *ProxyConfigurationStruct,
) ProxyServiceStruct {
	return ProxyServiceStruct{
		Type:  "ProxyService",
		Https: Https,
		Socks: Socks,
	}
}

// PublicSystemInfoFactory is just a simple function to instantiate the PublicSystemInfoStruct
func PublicSystemInfoFactory(
	ApiVersion *APIVersionStruct,
	Banner string,
	BuildTimestamp string,
	BuildTitle string,
	BuildVersion *VersionInfoStruct,
	CentralManagementProductName string,
	Configured *bool,
	CurrentLocale string,
	EnabledFeatures []string,
	EngineQualifier string,
	EngineType string,
	KernelName string,
	Locales []string,
	ProductName string,
	ProductType string,
	SsoEnabled *bool,
	SupportContacts []*SupportContactStruct,
	SupportURL string,
	Theme *SchemaDraftV4Struct,
	ToggleableFeatures []string,
	UserManagementEnabled *bool,
	VendorAddress []string,
	VendorEmail string,
	VendorName string,
	VendorPhoneNumber string,
	VendorURL string,
) PublicSystemInfoStruct {
	return PublicSystemInfoStruct{
		Type:                         "PublicSystemInfo",
		ApiVersion:                   ApiVersion,
		Banner:                       Banner,
		BuildTimestamp:               BuildTimestamp,
		BuildTitle:                   BuildTitle,
		BuildVersion:                 BuildVersion,
		CentralManagementProductName: CentralManagementProductName,
		Configured:                   Configured,
		CurrentLocale:                CurrentLocale,
		EnabledFeatures:              EnabledFeatures,
		EngineQualifier:              EngineQualifier,
		EngineType:                   EngineType,
		KernelName:                   KernelName,
		Locales:                      Locales,
		ProductName:                  ProductName,
		ProductType:                  ProductType,
		SsoEnabled:                   SsoEnabled,
		SupportContacts:              SupportContacts,
		SupportURL:                   SupportURL,
		Theme:                        Theme,
		ToggleableFeatures:           ToggleableFeatures,
		UserManagementEnabled:        UserManagementEnabled,
		VendorAddress:                VendorAddress,
		VendorEmail:                  VendorEmail,
		VendorName:                   VendorName,
		VendorPhoneNumber:            VendorPhoneNumber,
		VendorURL:                    VendorURL,
	}
}

// PurgeLogsParametersFactory is just a simple function to instantiate the PurgeLogsParametersStruct
func PurgeLogsParametersFactory(
	DeleteSnapshotLogs *bool,
	DryRun *bool,
	StorageSpaceToReclaim float64,
) PurgeLogsParametersStruct {
	return PurgeLogsParametersStruct{
		Type:                  "PurgeLogsParameters",
		DeleteSnapshotLogs:    DeleteSnapshotLogs,
		DryRun:                DryRun,
		StorageSpaceToReclaim: StorageSpaceToReclaim,
	}
}

// PurgeLogsResultFactory is just a simple function to instantiate the PurgeLogsResultStruct
func PurgeLogsResultFactory(
	AffectedSnapshots []TimeflowSnapshot,
	TruncatePoint *OracleTimeflowPointStruct,
) PurgeLogsResultStruct {
	return PurgeLogsResultStruct{
		Type:              "PurgeLogsResult",
		AffectedSnapshots: AffectedSnapshots,
		TruncatePoint:     TruncatePoint,
	}
}

// QuotaPolicyFactory is just a simple function to instantiate the QuotaPolicyStruct
func QuotaPolicyFactory(
	CritAlertTime string,
	Customized *bool,
	Default *bool,
	EffectiveType string,
	Name string,
	Namespace string,
	Reference string,
	Size float64,
	WarnAlertTime string,
) QuotaPolicyStruct {
	return QuotaPolicyStruct{
		Type:          "QuotaPolicy",
		CritAlertTime: CritAlertTime,
		Customized:    Customized,
		Default:       Default,
		EffectiveType: EffectiveType,
		Name:          Name,
		Namespace:     Namespace,
		Reference:     Reference,
		Size:          Size,
		WarnAlertTime: WarnAlertTime,
	}
}

// RefreshParametersFactory is just a simple function to instantiate the RefreshParametersStruct
func RefreshParametersFactory(
	TimeflowPointParameters TimeflowPointParameters,
) RefreshParametersStruct {
	return RefreshParametersStruct{
		Type:                    "RefreshParameters",
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// RefreshPolicyFactory is just a simple function to instantiate the RefreshPolicyStruct
func RefreshPolicyFactory(
	Customized *bool,
	Default *bool,
	EffectiveType string,
	Name string,
	Namespace string,
	ProvisionSource string,
	Reference string,
	ScheduleList []*ScheduleStruct,
	Timezone *TimeZoneStruct,
) RefreshPolicyStruct {
	return RefreshPolicyStruct{
		Type:            "RefreshPolicy",
		Customized:      Customized,
		Default:         Default,
		EffectiveType:   EffectiveType,
		Name:            Name,
		Namespace:       Namespace,
		ProvisionSource: ProvisionSource,
		Reference:       Reference,
		ScheduleList:    ScheduleList,
		Timezone:        Timezone,
	}
}

// RegistrationInfoFactory is just a simple function to instantiate the RegistrationInfoStruct
func RegistrationInfoFactory(
	Code string,
	RegistrationPortalHostname string,
	Uuid string,
) RegistrationInfoStruct {
	return RegistrationInfoStruct{
		Type:                       "RegistrationInfo",
		Code:                       Code,
		RegistrationPortalHostname: RegistrationPortalHostname,
		Uuid:                       Uuid,
	}
}

// RegistrationParametersFactory is just a simple function to instantiate the RegistrationParametersStruct
func RegistrationParametersFactory(
	Password string,
	Username string,
) RegistrationParametersStruct {
	return RegistrationParametersStruct{
		Type:     "RegistrationParameters",
		Password: Password,
		Username: Username,
	}
}

// RegistrationStatusFactory is just a simple function to instantiate the RegistrationStatusStruct
func RegistrationStatusFactory(
	Name string,
	Namespace string,
	Reference string,
	Status string,
	Timestamp string,
) RegistrationStatusStruct {
	return RegistrationStatusStruct{
		Type:      "RegistrationStatus",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
		Status:    Status,
		Timestamp: Timestamp,
	}
}

// RemoteDelphixEngineInfoFactory is just a simple function to instantiate the RemoteDelphixEngineInfoStruct
func RemoteDelphixEngineInfoFactory(
	Address string,
	Credential *PasswordCredentialStruct,
	Principal string,
) RemoteDelphixEngineInfoStruct {
	return RemoteDelphixEngineInfoStruct{
		Type:       "RemoteDelphixEngineInfo",
		Address:    Address,
		Credential: Credential,
		Principal:  Principal,
	}
}

// ReplicaRetentionPolicyFactory is just a simple function to instantiate the ReplicaRetentionPolicyStruct
func ReplicaRetentionPolicyFactory(
	Customized *bool,
	Default *bool,
	Duration *int,
	DurationUnit string,
	EffectiveType string,
	Name string,
	Namespace string,
	Reference string,
) ReplicaRetentionPolicyStruct {
	return ReplicaRetentionPolicyStruct{
		Type:          "ReplicaRetentionPolicy",
		Customized:    Customized,
		Default:       Default,
		Duration:      Duration,
		DurationUnit:  DurationUnit,
		EffectiveType: EffectiveType,
		Name:          Name,
		Namespace:     Namespace,
		Reference:     Reference,
	}
}

// ReplicationListFactory is just a simple function to instantiate the ReplicationListStruct
func ReplicationListFactory(
	Name string,
	Namespace string,
	Objects []string,
	Reference string,
) ReplicationListStruct {
	return ReplicationListStruct{
		Type:      "ReplicationList",
		Name:      Name,
		Namespace: Namespace,
		Objects:   Objects,
		Reference: Reference,
	}
}

// ReplicationSecureListFactory is just a simple function to instantiate the ReplicationSecureListStruct
func ReplicationSecureListFactory(
	Containers []string,
	Name string,
	Namespace string,
	Reference string,
) ReplicationSecureListStruct {
	return ReplicationSecureListStruct{
		Type:       "ReplicationSecureList",
		Containers: Containers,
		Name:       Name,
		Namespace:  Namespace,
		Reference:  Reference,
	}
}

// ReplicationSourceStateFactory is just a simple function to instantiate the ReplicationSourceStateStruct
func ReplicationSourceStateFactory(
	ActivePoint string,
	LastPoint string,
	Name string,
	Namespace string,
	Reference string,
	Spec string,
) ReplicationSourceStateStruct {
	return ReplicationSourceStateStruct{
		Type:        "ReplicationSourceState",
		ActivePoint: ActivePoint,
		LastPoint:   LastPoint,
		Name:        Name,
		Namespace:   Namespace,
		Reference:   Reference,
		Spec:        Spec,
	}
}

// ReplicationSpecFactory is just a simple function to instantiate the ReplicationSpecStruct
func ReplicationSpecFactory(
	AutomaticReplication *bool,
	BandwidthLimit *int,
	Description string,
	Encrypted *bool,
	LockedProfile *bool,
	Name string,
	Namespace string,
	NumberOfConnections *int,
	ObjectSpecification ReplicationObjectSpecification,
	Reference string,
	Runtime *ReplicationSpecRuntimeStruct,
	Schedule string,
	Tag string,
	TargetCredential *PasswordCredentialStruct,
	TargetHost string,
	TargetPort *int,
	TargetPrincipal string,
	UseSystemSocksSetting *bool,
) ReplicationSpecStruct {
	return ReplicationSpecStruct{
		Type:                  "ReplicationSpec",
		AutomaticReplication:  AutomaticReplication,
		BandwidthLimit:        BandwidthLimit,
		Description:           Description,
		Encrypted:             Encrypted,
		LockedProfile:         LockedProfile,
		Name:                  Name,
		Namespace:             Namespace,
		NumberOfConnections:   NumberOfConnections,
		ObjectSpecification:   ObjectSpecification,
		Reference:             Reference,
		Runtime:               Runtime,
		Schedule:              Schedule,
		Tag:                   Tag,
		TargetCredential:      TargetCredential,
		TargetHost:            TargetHost,
		TargetPort:            TargetPort,
		TargetPrincipal:       TargetPrincipal,
		UseSystemSocksSetting: UseSystemSocksSetting,
	}
}

// ReplicationSpecRuntimeFactory is just a simple function to instantiate the ReplicationSpecRuntimeStruct
func ReplicationSpecRuntimeFactory(
	NextScheduledExecution string,
) ReplicationSpecRuntimeStruct {
	return ReplicationSpecRuntimeStruct{
		Type:                   "ReplicationSpecRuntime",
		NextScheduledExecution: NextScheduledExecution,
	}
}

// ReplicationTargetStateFactory is just a simple function to instantiate the ReplicationTargetStateStruct
func ReplicationTargetStateFactory(
	Name string,
	Namespace string,
	Reference string,
) ReplicationTargetStateStruct {
	return ReplicationTargetStateStruct{
		Type:      "ReplicationTargetState",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// ResetIgnoredFaultsParametersFactory is just a simple function to instantiate the ResetIgnoredFaultsParametersStruct
func ResetIgnoredFaultsParametersFactory(
	FaultReferences []string,
) ResetIgnoredFaultsParametersStruct {
	return ResetIgnoredFaultsParametersStruct{
		Type:            "ResetIgnoredFaultsParameters",
		FaultReferences: FaultReferences,
	}
}

// ResolveOrIgnoreSelectedChecksParametersFactory is just a simple function to instantiate the ResolveOrIgnoreSelectedChecksParametersStruct
func ResolveOrIgnoreSelectedChecksParametersFactory(
	CheckReferences []string,
) ResolveOrIgnoreSelectedChecksParametersStruct {
	return ResolveOrIgnoreSelectedChecksParametersStruct{
		Type:            "ResolveOrIgnoreSelectedChecksParameters",
		CheckReferences: CheckReferences,
	}
}

// ResolveOrIgnoreSelectedFaultsParametersFactory is just a simple function to instantiate the ResolveOrIgnoreSelectedFaultsParametersStruct
func ResolveOrIgnoreSelectedFaultsParametersFactory(
	FaultReferences []string,
	Ignore *bool,
) ResolveOrIgnoreSelectedFaultsParametersStruct {
	return ResolveOrIgnoreSelectedFaultsParametersStruct{
		Type:            "ResolveOrIgnoreSelectedFaultsParameters",
		FaultReferences: FaultReferences,
		Ignore:          Ignore,
	}
}

// RetentionPolicyFactory is just a simple function to instantiate the RetentionPolicyStruct
func RetentionPolicyFactory(
	Customized *bool,
	DataDuration *int,
	DataUnit string,
	DayOfMonth *int,
	DayOfWeek string,
	DayOfYear string,
	Default *bool,
	EffectiveType string,
	LogDuration *int,
	LogUnit string,
	Name string,
	Namespace string,
	NumOfDaily *int,
	NumOfMonthly *int,
	NumOfWeekly *int,
	NumOfYearly *int,
	Reference string,
) RetentionPolicyStruct {
	return RetentionPolicyStruct{
		Type:          "RetentionPolicy",
		Customized:    Customized,
		DataDuration:  DataDuration,
		DataUnit:      DataUnit,
		DayOfMonth:    DayOfMonth,
		DayOfWeek:     DayOfWeek,
		DayOfYear:     DayOfYear,
		Default:       Default,
		EffectiveType: EffectiveType,
		LogDuration:   LogDuration,
		LogUnit:       LogUnit,
		Name:          Name,
		Namespace:     Namespace,
		NumOfDaily:    NumOfDaily,
		NumOfMonthly:  NumOfMonthly,
		NumOfWeekly:   NumOfWeekly,
		NumOfYearly:   NumOfYearly,
		Reference:     Reference,
	}
}

// RoleFactory is just a simple function to instantiate the RoleStruct
func RoleFactory(
	Immutable *bool,
	Name string,
	Namespace string,
	Permissions []string,
	Reference string,
) RoleStruct {
	return RoleStruct{
		Type:        "Role",
		Immutable:   Immutable,
		Name:        Name,
		Namespace:   Namespace,
		Permissions: Permissions,
		Reference:   Reference,
	}
}

// RollbackParametersFactory is just a simple function to instantiate the RollbackParametersStruct
func RollbackParametersFactory(
	TimeflowPointParameters TimeflowPointParameters,
) RollbackParametersStruct {
	return RollbackParametersStruct{
		Type:                    "RollbackParameters",
		TimeflowPointParameters: TimeflowPointParameters,
	}
}

// RsaKeyPairFactory is just a simple function to instantiate the RsaKeyPairStruct
func RsaKeyPairFactory(
	KeySize *int,
	SignatureAlgorithm string,
) RsaKeyPairStruct {
	return RsaKeyPairStruct{
		Type:               "RsaKeyPair",
		KeySize:            KeySize,
		SignatureAlgorithm: SignatureAlgorithm,
	}
}

// RunBashOnSourceOperationFactory is just a simple function to instantiate the RunBashOnSourceOperationStruct
func RunBashOnSourceOperationFactory(
	Command string,
	CredentialsEnvVarsList []*CredentialsEnvVarsStruct,
	Name string,
) RunBashOnSourceOperationStruct {
	return RunBashOnSourceOperationStruct{
		Type:                   "RunBashOnSourceOperation",
		Command:                Command,
		CredentialsEnvVarsList: CredentialsEnvVarsList,
		Name:                   Name,
	}
}

// RunCommandOnSourceOperationFactory is just a simple function to instantiate the RunCommandOnSourceOperationStruct
func RunCommandOnSourceOperationFactory(
	Command string,
	CredentialsEnvVarsList []*CredentialsEnvVarsStruct,
	Name string,
) RunCommandOnSourceOperationStruct {
	return RunCommandOnSourceOperationStruct{
		Type:                   "RunCommandOnSourceOperation",
		Command:                Command,
		CredentialsEnvVarsList: CredentialsEnvVarsList,
		Name:                   Name,
	}
}

// RunCommandOperationFactory is just a simple function to instantiate the RunCommandOperationStruct
func RunCommandOperationFactory(
	Command string,
	CredentialsEnvVarsList []*CredentialsEnvVarsStruct,
	Name string,
) RunCommandOperationStruct {
	return RunCommandOperationStruct{
		Type:                   "RunCommandOperation",
		Command:                Command,
		CredentialsEnvVarsList: CredentialsEnvVarsList,
		Name:                   Name,
	}
}

// RunDefaultPowerShellOnSourceOperationFactory is just a simple function to instantiate the RunDefaultPowerShellOnSourceOperationStruct
func RunDefaultPowerShellOnSourceOperationFactory(
	Command string,
	CredentialsEnvVarsList []*CredentialsEnvVarsStruct,
	Name string,
) RunDefaultPowerShellOnSourceOperationStruct {
	return RunDefaultPowerShellOnSourceOperationStruct{
		Type:                   "RunDefaultPowerShellOnSourceOperation",
		Command:                Command,
		CredentialsEnvVarsList: CredentialsEnvVarsList,
		Name:                   Name,
	}
}

// RunExpectOnSourceOperationFactory is just a simple function to instantiate the RunExpectOnSourceOperationStruct
func RunExpectOnSourceOperationFactory(
	Command string,
	CredentialsEnvVarsList []*CredentialsEnvVarsStruct,
	Name string,
) RunExpectOnSourceOperationStruct {
	return RunExpectOnSourceOperationStruct{
		Type:                   "RunExpectOnSourceOperation",
		Command:                Command,
		CredentialsEnvVarsList: CredentialsEnvVarsList,
		Name:                   Name,
	}
}

// RunExpectOperationFactory is just a simple function to instantiate the RunExpectOperationStruct
func RunExpectOperationFactory(
	Command string,
	CredentialsEnvVarsList []*CredentialsEnvVarsStruct,
	Name string,
) RunExpectOperationStruct {
	return RunExpectOperationStruct{
		Type:                   "RunExpectOperation",
		Command:                Command,
		CredentialsEnvVarsList: CredentialsEnvVarsList,
		Name:                   Name,
	}
}

// RunMaskingJobOperationFactory is just a simple function to instantiate the RunMaskingJobOperationStruct
func RunMaskingJobOperationFactory() RunMaskingJobOperationStruct {
	return RunMaskingJobOperationStruct{
		Type: "RunMaskingJobOperation",
	}
}

// RunPowerShellOnSourceOperationFactory is just a simple function to instantiate the RunPowerShellOnSourceOperationStruct
func RunPowerShellOnSourceOperationFactory(
	Command string,
	CredentialsEnvVarsList []*CredentialsEnvVarsStruct,
	Name string,
) RunPowerShellOnSourceOperationStruct {
	return RunPowerShellOnSourceOperationStruct{
		Type:                   "RunPowerShellOnSourceOperation",
		Command:                Command,
		CredentialsEnvVarsList: CredentialsEnvVarsList,
		Name:                   Name,
	}
}

// SMTPConfigFactory is just a simple function to instantiate the SMTPConfigStruct
func SMTPConfigFactory(
	AuthenticationEnabled *bool,
	Enabled *bool,
	FromAddress string,
	Password string,
	Port *int,
	SendTimeout *int,
	Server string,
	TlsEnabled *bool,
	Username string,
) SMTPConfigStruct {
	return SMTPConfigStruct{
		Type:                  "SMTPConfig",
		AuthenticationEnabled: AuthenticationEnabled,
		Enabled:               Enabled,
		FromAddress:           FromAddress,
		Password:              Password,
		Port:                  Port,
		SendTimeout:           SendTimeout,
		Server:                Server,
		TlsEnabled:            TlsEnabled,
		Username:              Username,
	}
}

// SNMPV2ConfigFactory is just a simple function to instantiate the SNMPV2ConfigStruct
func SNMPV2ConfigFactory(
	AuthorizedNetwork string,
	Community string,
	Enabled *bool,
	Location string,
	Severity string,
) SNMPV2ConfigStruct {
	return SNMPV2ConfigStruct{
		Type:              "SNMPV2Config",
		AuthorizedNetwork: AuthorizedNetwork,
		Community:         Community,
		Enabled:           Enabled,
		Location:          Location,
		Severity:          Severity,
	}
}

// SNMPV2ManagerFactory is just a simple function to instantiate the SNMPV2ManagerStruct
func SNMPV2ManagerFactory(
	Address string,
	CommunityString string,
	Namespace string,
	Port *int,
	Reference string,
	UseInform *bool,
) SNMPV2ManagerStruct {
	return SNMPV2ManagerStruct{
		Type:            "SNMPV2Manager",
		Address:         Address,
		CommunityString: CommunityString,
		Namespace:       Namespace,
		Port:            Port,
		Reference:       Reference,
		UseInform:       UseInform,
	}
}

// SNMPV3ConfigFactory is just a simple function to instantiate the SNMPV3ConfigStruct
func SNMPV3ConfigFactory(
	Enabled *bool,
	Location string,
	Severity string,
) SNMPV3ConfigStruct {
	return SNMPV3ConfigStruct{
		Type:     "SNMPV3Config",
		Enabled:  Enabled,
		Location: Location,
		Severity: Severity,
	}
}

// SNMPV3ManagerFactory is just a simple function to instantiate the SNMPV3ManagerStruct
func SNMPV3ManagerFactory(
	Address string,
	AuthenticationPassphrase string,
	AuthenticationProtocol string,
	Namespace string,
	Port *int,
	PrivacyPassphrase string,
	PrivacyProtocol string,
	Reference string,
	UseInform *bool,
	Username string,
) SNMPV3ManagerStruct {
	return SNMPV3ManagerStruct{
		Type:                     "SNMPV3Manager",
		Address:                  Address,
		AuthenticationPassphrase: AuthenticationPassphrase,
		AuthenticationProtocol:   AuthenticationProtocol,
		Namespace:                Namespace,
		Port:                     Port,
		PrivacyPassphrase:        PrivacyPassphrase,
		PrivacyProtocol:          PrivacyProtocol,
		Reference:                Reference,
		UseInform:                UseInform,
		Username:                 Username,
	}
}

// SNMPV3USMUserConfigFactory is just a simple function to instantiate the SNMPV3USMUserConfigStruct
func SNMPV3USMUserConfigFactory(
	AuthenticationPassphrase string,
	AuthenticationProtocol string,
	Namespace string,
	PrivacyPassphrase string,
	PrivacyProtocol string,
	Reference string,
	Username string,
) SNMPV3USMUserConfigStruct {
	return SNMPV3USMUserConfigStruct{
		Type:                     "SNMPV3USMUserConfig",
		AuthenticationPassphrase: AuthenticationPassphrase,
		AuthenticationProtocol:   AuthenticationProtocol,
		Namespace:                Namespace,
		PrivacyPassphrase:        PrivacyPassphrase,
		PrivacyProtocol:          PrivacyProtocol,
		Reference:                Reference,
		Username:                 Username,
	}
}

// SSHConnectivityFactory is just a simple function to instantiate the SSHConnectivityStruct
func SSHConnectivityFactory(
	Address string,
	Credentials Credential,
	Port *int,
	SshVerificationStrategy SshVerificationStrategy,
	Username string,
) SSHConnectivityStruct {
	return SSHConnectivityStruct{
		Type:                    "SSHConnectivity",
		Address:                 Address,
		Credentials:             Credentials,
		Port:                    Port,
		SshVerificationStrategy: SshVerificationStrategy,
		Username:                Username,
	}
}

// ScheduleFactory is just a simple function to instantiate the ScheduleStruct
func ScheduleFactory(
	CronString string,
	CutoffTime *int,
) ScheduleStruct {
	return ScheduleStruct{
		Type:       "Schedule",
		CronString: CronString,
		CutoffTime: CutoffTime,
	}
}

// SchemaFactory is just a simple function to instantiate the SchemaStruct
func SchemaFactory(
	Schema string,
) SchemaStruct {
	return SchemaStruct{
		Type:   "Schema",
		Schema: Schema,
	}
}

// ScrubStatusFactory is just a simple function to instantiate the ScrubStatusStruct
func ScrubStatusFactory(
	Completed float64,
	EndTime string,
	Errors float64,
	StartTime string,
	State string,
	Total float64,
) ScrubStatusStruct {
	return ScrubStatusStruct{
		Type:      "ScrubStatus",
		Completed: Completed,
		EndTime:   EndTime,
		Errors:    Errors,
		StartTime: StartTime,
		State:     State,
		Total:     Total,
	}
}

// SecurityConfigFactory is just a simple function to instantiate the SecurityConfigStruct
func SecurityConfigFactory(
	AllowedCORSOrigins string,
	Banner string,
	EnableCORSSupportsCredentials *bool,
	IsCORSEnabled *bool,
) SecurityConfigStruct {
	return SecurityConfigStruct{
		Type:                          "SecurityConfig",
		AllowedCORSOrigins:            AllowedCORSOrigins,
		Banner:                        Banner,
		EnableCORSSupportsCredentials: EnableCORSSupportsCredentials,
		IsCORSEnabled:                 IsCORSEnabled,
	}
}

// SerializationPointFactory is just a simple function to instantiate the SerializationPointStruct
func SerializationPointFactory(
	AverageThroughput float64,
	BytesTransferred *int,
	DataTimestamp string,
	ElapsedTimeNanos float64,
	Name string,
	Namespace string,
	Reference string,
	Tag string,
) SerializationPointStruct {
	return SerializationPointStruct{
		Type:              "SerializationPoint",
		AverageThroughput: AverageThroughput,
		BytesTransferred:  BytesTransferred,
		DataTimestamp:     DataTimestamp,
		ElapsedTimeNanos:  ElapsedTimeNanos,
		Name:              Name,
		Namespace:         Namespace,
		Reference:         Reference,
		Tag:               Tag,
	}
}

// SeverityFilterFactory is just a simple function to instantiate the SeverityFilterStruct
func SeverityFilterFactory(
	SeverityLevels []string,
) SeverityFilterStruct {
	return SeverityFilterStruct{
		Type:           "SeverityFilter",
		SeverityLevels: SeverityLevels,
	}
}

// SingletonUpdateFactory is just a simple function to instantiate the SingletonUpdateStruct
func SingletonUpdateFactory(
	ObjectType string,
) SingletonUpdateStruct {
	return SingletonUpdateStruct{
		Type:       "SingletonUpdate",
		ObjectType: ObjectType,
	}
}

// SnapshotCapacityDataFactory is just a simple function to instantiate the SnapshotCapacityDataStruct
func SnapshotCapacityDataFactory(
	Container string,
	DescendantVDBs []string,
	ManualRetention *int,
	Namespace string,
	PolicyRetention *bool,
	Snapshot string,
	SnapshotLatestChange string,
	SnapshotTimestamp string,
	SnapshotTimezone string,
	Space float64,
) SnapshotCapacityDataStruct {
	return SnapshotCapacityDataStruct{
		Type:                 "SnapshotCapacityData",
		Container:            Container,
		DescendantVDBs:       DescendantVDBs,
		ManualRetention:      ManualRetention,
		Namespace:            Namespace,
		PolicyRetention:      PolicyRetention,
		Snapshot:             Snapshot,
		SnapshotLatestChange: SnapshotLatestChange,
		SnapshotTimestamp:    SnapshotTimestamp,
		SnapshotTimezone:     SnapshotTimezone,
		Space:                Space,
	}
}

// SnapshotLogFetchParametersFactory is just a simple function to instantiate the SnapshotLogFetchParametersStruct
func SnapshotLogFetchParametersFactory(
	Credentials Credential,
	Directory string,
	Host string,
	Port *int,
	Snapshot string,
	SshVerificationStrategy SshVerificationStrategy,
	Username string,
) SnapshotLogFetchParametersStruct {
	return SnapshotLogFetchParametersStruct{
		Type:                    "SnapshotLogFetchParameters",
		Credentials:             Credentials,
		Directory:               Directory,
		Host:                    Host,
		Port:                    Port,
		Snapshot:                Snapshot,
		SshVerificationStrategy: SshVerificationStrategy,
		Username:                Username,
	}
}

// SnapshotPolicyFactory is just a simple function to instantiate the SnapshotPolicyStruct
func SnapshotPolicyFactory(
	Customized *bool,
	Default *bool,
	EffectiveType string,
	Name string,
	Namespace string,
	Reference string,
	ScheduleList []*ScheduleStruct,
	Timezone *TimeZoneStruct,
) SnapshotPolicyStruct {
	return SnapshotPolicyStruct{
		Type:          "SnapshotPolicy",
		Customized:    Customized,
		Default:       Default,
		EffectiveType: EffectiveType,
		Name:          Name,
		Namespace:     Namespace,
		Reference:     Reference,
		ScheduleList:  ScheduleList,
		Timezone:      Timezone,
	}
}

// SnapshotRuntimeFactory is just a simple function to instantiate the SnapshotRuntimeStruct
func SnapshotRuntimeFactory(
	Provisionable *bool,
) SnapshotRuntimeStruct {
	return SnapshotRuntimeStruct{
		Type:          "SnapshotRuntime",
		Provisionable: Provisionable,
	}
}

// SnapshotSpaceMapFactory is just a simple function to instantiate the SnapshotSpaceMapStruct
func SnapshotSpaceMapFactory(
	SizeMap map[string]string,
) SnapshotSpaceMapStruct {
	return SnapshotSpaceMapStruct{
		Type:    "SnapshotSpaceMap",
		SizeMap: SizeMap,
	}
}

// SnapshotSpaceParametersFactory is just a simple function to instantiate the SnapshotSpaceParametersStruct
func SnapshotSpaceParametersFactory(
	ObjectReferences []string,
) SnapshotSpaceParametersStruct {
	return SnapshotSpaceParametersStruct{
		Type:             "SnapshotSpaceParameters",
		ObjectReferences: ObjectReferences,
	}
}

// SnapshotSpaceResultFactory is just a simple function to instantiate the SnapshotSpaceResultStruct
func SnapshotSpaceResultFactory(
	TotalSize float64,
) SnapshotSpaceResultStruct {
	return SnapshotSpaceResultStruct{
		Type:      "SnapshotSpaceResult",
		TotalSize: TotalSize,
	}
}

// SourceConfigConnectivityFactory is just a simple function to instantiate the SourceConfigConnectivityStruct
func SourceConfigConnectivityFactory(
	Password string,
	Username string,
) SourceConfigConnectivityStruct {
	return SourceConfigConnectivityStruct{
		Type:     "SourceConfigConnectivity",
		Password: Password,
		Username: Username,
	}
}

// SourceDisableParametersFactory is just a simple function to instantiate the SourceDisableParametersStruct
func SourceDisableParametersFactory(
	AttemptCleanup *bool,
) SourceDisableParametersStruct {
	return SourceDisableParametersStruct{
		Type:           "SourceDisableParameters",
		AttemptCleanup: AttemptCleanup,
	}
}

// SourceEnableParametersFactory is just a simple function to instantiate the SourceEnableParametersStruct
func SourceEnableParametersFactory(
	AttemptStart *bool,
) SourceEnableParametersStruct {
	return SourceEnableParametersStruct{
		Type:         "SourceEnableParameters",
		AttemptStart: AttemptStart,
	}
}

// SourceIngestionDataFactory is just a simple function to instantiate the SourceIngestionDataStruct
func SourceIngestionDataFactory() SourceIngestionDataStruct {
	return SourceIngestionDataStruct{
		Type: "SourceIngestionData",
	}
}

// SourceRepositoryTemplateFactory is just a simple function to instantiate the SourceRepositoryTemplateStruct
func SourceRepositoryTemplateFactory(
	Container string,
	Name string,
	Namespace string,
	Reference string,
	Repository string,
	Template string,
) SourceRepositoryTemplateStruct {
	return SourceRepositoryTemplateStruct{
		Type:       "SourceRepositoryTemplate",
		Container:  Container,
		Name:       Name,
		Namespace:  Namespace,
		Reference:  Reference,
		Repository: Repository,
		Template:   Template,
	}
}

// SourceStartParametersFactory is just a simple function to instantiate the SourceStartParametersStruct
func SourceStartParametersFactory() SourceStartParametersStruct {
	return SourceStartParametersStruct{
		Type: "SourceStartParameters",
	}
}

// SourceStopParametersFactory is just a simple function to instantiate the SourceStopParametersStruct
func SourceStopParametersFactory() SourceStopParametersStruct {
	return SourceStopParametersStruct{
		Type: "SourceStopParameters",
	}
}

// SourceTypeAggregateIngestedSizeFactory is just a simple function to instantiate the SourceTypeAggregateIngestedSizeStruct
func SourceTypeAggregateIngestedSizeFactory() SourceTypeAggregateIngestedSizeStruct {
	return SourceTypeAggregateIngestedSizeStruct{
		Type: "SourceTypeAggregateIngestedSize",
	}
}

// SourceUpgradeParametersFactory is just a simple function to instantiate the SourceUpgradeParametersStruct
func SourceUpgradeParametersFactory(
	SourceConfig SourceConfig,
) SourceUpgradeParametersStruct {
	return SourceUpgradeParametersStruct{
		Type:         "SourceUpgradeParameters",
		SourceConfig: SourceConfig,
	}
}

// SourcingPolicyFactory is just a simple function to instantiate the SourcingPolicyStruct
func SourcingPolicyFactory(
	LogsyncEnabled *bool,
) SourcingPolicyStruct {
	return SourcingPolicyStruct{
		Type:           "SourcingPolicy",
		LogsyncEnabled: LogsyncEnabled,
	}
}

// SplunkHecConfigFactory is just a simple function to instantiate the SplunkHecConfigStruct
func SplunkHecConfigFactory(
	EnableMetrics *bool,
	EnableSSL *bool,
	Enabled *bool,
	EventsPushFrequency *int,
	HecPort *int,
	HecToken string,
	Host string,
	MainIndex string,
	MetricsIndex string,
	MetricsPushFrequency *int,
	Name string,
	Namespace string,
	PerformanceMetricsResolution string,
	Reference string,
) SplunkHecConfigStruct {
	return SplunkHecConfigStruct{
		Type:                         "SplunkHecConfig",
		EnableMetrics:                EnableMetrics,
		EnableSSL:                    EnableSSL,
		Enabled:                      Enabled,
		EventsPushFrequency:          EventsPushFrequency,
		HecPort:                      HecPort,
		HecToken:                     HecToken,
		Host:                         Host,
		MainIndex:                    MainIndex,
		MetricsIndex:                 MetricsIndex,
		MetricsPushFrequency:         MetricsPushFrequency,
		Name:                         Name,
		Namespace:                    Namespace,
		PerformanceMetricsResolution: PerformanceMetricsResolution,
		Reference:                    Reference,
	}
}

// SshAcceptAlwaysFactory is just a simple function to instantiate the SshAcceptAlwaysStruct
func SshAcceptAlwaysFactory() SshAcceptAlwaysStruct {
	return SshAcceptAlwaysStruct{
		Type: "SshAcceptAlways",
	}
}

// SshVerifyFingerprintFactory is just a simple function to instantiate the SshVerifyFingerprintStruct
func SshVerifyFingerprintFactory(
	Fingerprint string,
	FingerprintType string,
	KeyType string,
) SshVerifyFingerprintStruct {
	return SshVerifyFingerprintStruct{
		Type:            "SshVerifyFingerprint",
		Fingerprint:     Fingerprint,
		FingerprintType: FingerprintType,
		KeyType:         KeyType,
	}
}

// SshVerifyRawKeyFactory is just a simple function to instantiate the SshVerifyRawKeyStruct
func SshVerifyRawKeyFactory(
	KeyType string,
	RawKey string,
) SshVerifyRawKeyStruct {
	return SshVerifyRawKeyStruct{
		Type:    "SshVerifyRawKey",
		KeyType: KeyType,
		RawKey:  RawKey,
	}
}

// SsoConfigFactory is just a simple function to instantiate the SsoConfigStruct
func SsoConfigFactory(
	CloudSso *bool,
	Enabled *bool,
	EntityId string,
	MaxAuthenticationAge string,
	ResponseSkewTime string,
	SamlMetadata string,
) SsoConfigStruct {
	return SsoConfigStruct{
		Type:                 "SsoConfig",
		CloudSso:             CloudSso,
		Enabled:              Enabled,
		EntityId:             EntityId,
		MaxAuthenticationAge: MaxAuthenticationAge,
		ResponseSkewTime:     ResponseSkewTime,
		SamlMetadata:         SamlMetadata,
	}
}

// StagingCompatibilityParametersFactory is just a simple function to instantiate the StagingCompatibilityParametersStruct
func StagingCompatibilityParametersFactory(
	Environment string,
	SourceRepository string,
) StagingCompatibilityParametersStruct {
	return StagingCompatibilityParametersStruct{
		Type:             "StagingCompatibilityParameters",
		Environment:      Environment,
		SourceRepository: SourceRepository,
	}
}

// StagingSourceOperationsFactory is just a simple function to instantiate the StagingSourceOperationsStruct
func StagingSourceOperationsFactory(
	PostValidatedSync []SourceOperation,
	PreValidatedSync []SourceOperation,
) StagingSourceOperationsStruct {
	return StagingSourceOperationsStruct{
		Type:              "StagingSourceOperations",
		PostValidatedSync: PostValidatedSync,
		PreValidatedSync:  PreValidatedSync,
	}
}

// StaticHostAddressFactory is just a simple function to instantiate the StaticHostAddressStruct
func StaticHostAddressFactory(
	Addresses []string,
	Hostname string,
	Name string,
	Namespace string,
	Reference string,
) StaticHostAddressStruct {
	return StaticHostAddressStruct{
		Type:      "StaticHostAddress",
		Addresses: Addresses,
		Hostname:  Hostname,
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// StatisticFactory is just a simple function to instantiate the StatisticStruct
func StatisticFactory(
	Axes []*StatisticAxisStruct,
	Explanation string,
	MinCollectionInterval *int,
	StatisticType string,
) StatisticStruct {
	return StatisticStruct{
		Type:                  "Statistic",
		Axes:                  Axes,
		Explanation:           Explanation,
		MinCollectionInterval: MinCollectionInterval,
		StatisticType:         StatisticType,
	}
}

// StatisticAxisFactory is just a simple function to instantiate the StatisticAxisStruct
func StatisticAxisFactory(
	AxisName string,
	ConstraintType string,
	Explanation string,
	StreamAttribute *bool,
	ValueType string,
) StatisticAxisStruct {
	return StatisticAxisStruct{
		Type:            "StatisticAxis",
		AxisName:        AxisName,
		ConstraintType:  ConstraintType,
		Explanation:     Explanation,
		StreamAttribute: StreamAttribute,
		ValueType:       ValueType,
	}
}

// StatisticEnumAxisFactory is just a simple function to instantiate the StatisticEnumAxisStruct
func StatisticEnumAxisFactory(
	AxisName string,
	ConstraintType string,
	EnumValues []string,
	Explanation string,
	StreamAttribute *bool,
	ValueType string,
) StatisticEnumAxisStruct {
	return StatisticEnumAxisStruct{
		Type:            "StatisticEnumAxis",
		AxisName:        AxisName,
		ConstraintType:  ConstraintType,
		EnumValues:      EnumValues,
		Explanation:     Explanation,
		StreamAttribute: StreamAttribute,
		ValueType:       ValueType,
	}
}

// StatisticSliceFactory is just a simple function to instantiate the StatisticSliceStruct
func StatisticSliceFactory(
	AxisConstraints []AxisConstraint,
	CollectionAxes []string,
	CollectionInterval *int,
	Name string,
	Namespace string,
	Reference string,
	State string,
	StatisticType string,
) StatisticSliceStruct {
	return StatisticSliceStruct{
		Type:               "StatisticSlice",
		AxisConstraints:    AxisConstraints,
		CollectionAxes:     CollectionAxes,
		CollectionInterval: CollectionInterval,
		Name:               Name,
		Namespace:          Namespace,
		Reference:          Reference,
		State:              State,
		StatisticType:      StatisticType,
	}
}

// StorageDeviceFactory is just a simple function to instantiate the StorageDeviceStruct
func StorageDeviceFactory(
	Configured *bool,
	Model string,
	Name string,
	Namespace string,
	Reference string,
	Serial string,
	Size float64,
	Vendor string,
) StorageDeviceStruct {
	return StorageDeviceStruct{
		Type:       "StorageDevice",
		Configured: Configured,
		Model:      Model,
		Name:       Name,
		Namespace:  Namespace,
		Reference:  Reference,
		Serial:     Serial,
		Size:       Size,
		Vendor:     Vendor,
	}
}

// StorageDeviceInitializeStatusFactory is just a simple function to instantiate the StorageDeviceInitializeStatusStruct
func StorageDeviceInitializeStatusFactory(
	BytesDone float64,
	BytesEst float64,
	State string,
) StorageDeviceInitializeStatusStruct {
	return StorageDeviceInitializeStatusStruct{
		Type:      "StorageDeviceInitializeStatus",
		BytesDone: BytesDone,
		BytesEst:  BytesEst,
		State:     State,
	}
}

// StorageDeviceRemovalStatusFactory is just a simple function to instantiate the StorageDeviceRemovalStatusStruct
func StorageDeviceRemovalStatusFactory(
	Copied float64,
	MappingMemory float64,
	StartTime string,
	State string,
	Total float64,
) StorageDeviceRemovalStatusStruct {
	return StorageDeviceRemovalStatusStruct{
		Type:          "StorageDeviceRemovalStatus",
		Copied:        Copied,
		MappingMemory: MappingMemory,
		StartTime:     StartTime,
		State:         State,
		Total:         Total,
	}
}

// StorageDeviceRemovalVerifyResultFactory is just a simple function to instantiate the StorageDeviceRemovalVerifyResultStruct
func StorageDeviceRemovalVerifyResultFactory(
	NewFreeBytes float64,
	NewMappingMemory float64,
	OldFreeBytes float64,
	OldMappingMemory float64,
) StorageDeviceRemovalVerifyResultStruct {
	return StorageDeviceRemovalVerifyResultStruct{
		Type:             "StorageDeviceRemovalVerifyResult",
		NewFreeBytes:     NewFreeBytes,
		NewMappingMemory: NewMappingMemory,
		OldFreeBytes:     OldFreeBytes,
		OldMappingMemory: OldMappingMemory,
	}
}

// StorageTestFactory is just a simple function to instantiate the StorageTestStruct
func StorageTestFactory(
	EndTime string,
	Name string,
	Namespace string,
	Parameters *StorageTestParametersStruct,
	Reference string,
	StartTime string,
	State string,
	TestResults []*StorageTestResultStruct,
) StorageTestStruct {
	return StorageTestStruct{
		Type:        "StorageTest",
		EndTime:     EndTime,
		Name:        Name,
		Namespace:   Namespace,
		Parameters:  Parameters,
		Reference:   Reference,
		StartTime:   StartTime,
		State:       State,
		TestResults: TestResults,
	}
}

// StorageTestParametersFactory is just a simple function to instantiate the StorageTestParametersStruct
func StorageTestParametersFactory(
	Devices []string,
	Duration *int,
	InitializeDevices *bool,
	InitializeEntireDevice *bool,
	TestRegion float64,
	Tests string,
) StorageTestParametersStruct {
	return StorageTestParametersStruct{
		Type:                   "StorageTestParameters",
		Devices:                Devices,
		Duration:               Duration,
		InitializeDevices:      InitializeDevices,
		InitializeEntireDevice: InitializeEntireDevice,
		TestRegion:             TestRegion,
		Tests:                  Tests,
	}
}

// StorageTestResultFactory is just a simple function to instantiate the StorageTestResultStruct
func StorageTestResultFactory(
	AverageLatency float64,
	BlockSize *int,
	Iops *int,
	Jobs *int,
	Latency95thPercentile float64,
	LatencyGrade string,
	LoadScaling float64,
	LoadScalingGrade string,
	MaxLatency float64,
	MinLatency float64,
	StddevLatency float64,
	TestName string,
	TestType string,
	Throughput float64,
) StorageTestResultStruct {
	return StorageTestResultStruct{
		Type:                  "StorageTestResult",
		AverageLatency:        AverageLatency,
		BlockSize:             BlockSize,
		Iops:                  Iops,
		Jobs:                  Jobs,
		Latency95thPercentile: Latency95thPercentile,
		LatencyGrade:          LatencyGrade,
		LoadScaling:           LoadScaling,
		LoadScalingGrade:      LoadScalingGrade,
		MaxLatency:            MaxLatency,
		MinLatency:            MinLatency,
		StddevLatency:         StddevLatency,
		TestName:              TestName,
		TestType:              TestType,
		Throughput:            Throughput,
	}
}

// StringEqualConstraintFactory is just a simple function to instantiate the StringEqualConstraintStruct
func StringEqualConstraintFactory(
	AxisName string,
	Equals string,
) StringEqualConstraintStruct {
	return StringEqualConstraintStruct{
		Type:     "StringEqualConstraint",
		AxisName: AxisName,
		Equals:   Equals,
	}
}

// SupportAccessStateFactory is just a simple function to instantiate the SupportAccessStateStruct
func SupportAccessStateFactory(
	AccessType string,
	AuthenticationMethod string,
	EndTime string,
	StartTime string,
	Token string,
) SupportAccessStateStruct {
	return SupportAccessStateStruct{
		Type:                 "SupportAccessState",
		AccessType:           AccessType,
		AuthenticationMethod: AuthenticationMethod,
		EndTime:              EndTime,
		StartTime:            StartTime,
		Token:                Token,
	}
}

// SupportBundleConfigurationFactory is just a simple function to instantiate the SupportBundleConfigurationStruct
func SupportBundleConfigurationFactory(
	MaxActions string,
) SupportBundleConfigurationStruct {
	return SupportBundleConfigurationStruct{
		Type:       "SupportBundleConfiguration",
		MaxActions: MaxActions,
	}
}

// SupportBundleGenerateParametersFactory is just a simple function to instantiate the SupportBundleGenerateParametersStruct
func SupportBundleGenerateParametersFactory(
	BundleType string,
	Environments []string,
	IncludeAnalyticsData *bool,
	Plugin string,
	Sources []string,
) SupportBundleGenerateParametersStruct {
	return SupportBundleGenerateParametersStruct{
		Type:                 "SupportBundleGenerateParameters",
		BundleType:           BundleType,
		Environments:         Environments,
		IncludeAnalyticsData: IncludeAnalyticsData,
		Plugin:               Plugin,
		Sources:              Sources,
	}
}

// SupportBundleUploadParametersFactory is just a simple function to instantiate the SupportBundleUploadParametersStruct
func SupportBundleUploadParametersFactory(
	BundleType string,
	CaseNumber *int,
	Environments []string,
	IncludeAnalyticsData *bool,
	Plugin string,
	Sources []string,
) SupportBundleUploadParametersStruct {
	return SupportBundleUploadParametersStruct{
		Type:                 "SupportBundleUploadParameters",
		BundleType:           BundleType,
		CaseNumber:           CaseNumber,
		Environments:         Environments,
		IncludeAnalyticsData: IncludeAnalyticsData,
		Plugin:               Plugin,
		Sources:              Sources,
	}
}

// SupportContactFactory is just a simple function to instantiate the SupportContactStruct
func SupportContactFactory(
	Country string,
	PhoneNumber string,
) SupportContactStruct {
	return SupportContactStruct{
		Type:        "SupportContact",
		Country:     Country,
		PhoneNumber: PhoneNumber,
	}
}

// SwitchSessionUserParametersFactory is just a simple function to instantiate the SwitchSessionUserParametersStruct
func SwitchSessionUserParametersFactory(
	User string,
) SwitchSessionUserParametersStruct {
	return SwitchSessionUserParametersStruct{
		Type: "SwitchSessionUserParameters",
		User: User,
	}
}

// SwitchTimeflowParametersFactory is just a simple function to instantiate the SwitchTimeflowParametersStruct
func SwitchTimeflowParametersFactory(
	Timeflow string,
) SwitchTimeflowParametersStruct {
	return SwitchTimeflowParametersStruct{
		Type:     "SwitchTimeflowParameters",
		Timeflow: Timeflow,
	}
}

// SyncParametersFactory is just a simple function to instantiate the SyncParametersStruct
func SyncParametersFactory() SyncParametersStruct {
	return SyncParametersStruct{
		Type: "SyncParameters",
	}
}

// SyncPolicyFactory is just a simple function to instantiate the SyncPolicyStruct
func SyncPolicyFactory(
	Customized *bool,
	Default *bool,
	EffectiveType string,
	Name string,
	Namespace string,
	Reference string,
	ScheduleList []*ScheduleStruct,
	Timezone *TimeZoneStruct,
) SyncPolicyStruct {
	return SyncPolicyStruct{
		Type:          "SyncPolicy",
		Customized:    Customized,
		Default:       Default,
		EffectiveType: EffectiveType,
		Name:          Name,
		Namespace:     Namespace,
		Reference:     Reference,
		ScheduleList:  ScheduleList,
		Timezone:      Timezone,
	}
}

// SyslogConfigFactory is just a simple function to instantiate the SyslogConfigStruct
func SyslogConfigFactory(
	Enabled *bool,
	Format string,
	Pattern string,
	Servers []*SyslogServerStruct,
	Severity string,
) SyslogConfigStruct {
	return SyslogConfigStruct{
		Type:     "SyslogConfig",
		Enabled:  Enabled,
		Format:   Format,
		Pattern:  Pattern,
		Servers:  Servers,
		Severity: Severity,
	}
}

// SyslogServerFactory is just a simple function to instantiate the SyslogServerStruct
func SyslogServerFactory(
	Address string,
	Port *int,
	Protocol string,
) SyslogServerStruct {
	return SyslogServerStruct{
		Type:     "SyslogServer",
		Address:  Address,
		Port:     Port,
		Protocol: Protocol,
	}
}

// SystemInfoFactory is just a simple function to instantiate the SystemInfoStruct
func SystemInfoFactory(
	ApiVersion *APIVersionStruct,
	Banner string,
	BuildTimestamp string,
	BuildTitle string,
	BuildVersion *VersionInfoStruct,
	CentralManagementProductName string,
	Configured *bool,
	CpuReservation float64,
	CurrentLocale string,
	EnabledFeatures []string,
	EngineQualifier string,
	EngineType string,
	Hostname string,
	Hotfixes string,
	InstallationTime string,
	KernelName string,
	Locales []string,
	MemoryReservation float64,
	MemorySize float64,
	Platform string,
	PoolFragmentation float64,
	Processors []*CPUInfoStruct,
	ProductName string,
	ProductType string,
	SshPublicKey string,
	SsoEnabled *bool,
	StorageTotal float64,
	StorageUsed float64,
	SupportContacts []*SupportContactStruct,
	SupportURL string,
	Theme *SchemaDraftV4Struct,
	ToggleableFeatures []string,
	UpTime *UpTimeInfoStruct,
	UserManagementEnabled *bool,
	Uuid string,
	VendorAddress []string,
	VendorEmail string,
	VendorName string,
	VendorPhoneNumber string,
	VendorURL string,
) SystemInfoStruct {
	return SystemInfoStruct{
		Type:                         "SystemInfo",
		ApiVersion:                   ApiVersion,
		Banner:                       Banner,
		BuildTimestamp:               BuildTimestamp,
		BuildTitle:                   BuildTitle,
		BuildVersion:                 BuildVersion,
		CentralManagementProductName: CentralManagementProductName,
		Configured:                   Configured,
		CpuReservation:               CpuReservation,
		CurrentLocale:                CurrentLocale,
		EnabledFeatures:              EnabledFeatures,
		EngineQualifier:              EngineQualifier,
		EngineType:                   EngineType,
		Hostname:                     Hostname,
		Hotfixes:                     Hotfixes,
		InstallationTime:             InstallationTime,
		KernelName:                   KernelName,
		Locales:                      Locales,
		MemoryReservation:            MemoryReservation,
		MemorySize:                   MemorySize,
		Platform:                     Platform,
		PoolFragmentation:            PoolFragmentation,
		Processors:                   Processors,
		ProductName:                  ProductName,
		ProductType:                  ProductType,
		SshPublicKey:                 SshPublicKey,
		SsoEnabled:                   SsoEnabled,
		StorageTotal:                 StorageTotal,
		StorageUsed:                  StorageUsed,
		SupportContacts:              SupportContacts,
		SupportURL:                   SupportURL,
		Theme:                        Theme,
		ToggleableFeatures:           ToggleableFeatures,
		UpTime:                       UpTime,
		UserManagementEnabled:        UserManagementEnabled,
		Uuid:                         Uuid,
		VendorAddress:                VendorAddress,
		VendorEmail:                  VendorEmail,
		VendorName:                   VendorName,
		VendorPhoneNumber:            VendorPhoneNumber,
		VendorURL:                    VendorURL,
	}
}

// SystemInitializationParametersFactory is just a simple function to instantiate the SystemInitializationParametersStruct
func SystemInitializationParametersFactory(
	DefaultEmail string,
	DefaultPassword string,
	DefaultUser string,
	Devices []string,
) SystemInitializationParametersStruct {
	return SystemInitializationParametersStruct{
		Type:            "SystemInitializationParameters",
		DefaultEmail:    DefaultEmail,
		DefaultPassword: DefaultPassword,
		DefaultUser:     DefaultUser,
		Devices:         Devices,
	}
}

// SystemKeyCredentialFactory is just a simple function to instantiate the SystemKeyCredentialStruct
func SystemKeyCredentialFactory() SystemKeyCredentialStruct {
	return SystemKeyCredentialStruct{
		Type: "SystemKeyCredential",
	}
}

// SystemPackageFactory is just a simple function to instantiate the SystemPackageStruct
func SystemPackageFactory(
	Name string,
	Namespace string,
	PossibleVersions []string,
	Reference string,
	Version string,
) SystemPackageStruct {
	return SystemPackageStruct{
		Type:             "SystemPackage",
		Name:             Name,
		Namespace:        Namespace,
		PossibleVersions: PossibleVersions,
		Reference:        Reference,
		Version:          Version,
	}
}

// SystemStatusFactory is just a simple function to instantiate the SystemStatusStruct
func SystemStatusFactory(
	Name string,
	Namespace string,
	Reference string,
) SystemStatusStruct {
	return SystemStatusStruct{
		Type:      "SystemStatus",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// SystemVersionFactory is just a simple function to instantiate the SystemVersionStruct
func SystemVersionFactory(
	BuildDate string,
	HotfixVersion string,
	InstallDate string,
	MinRebootOptionalVersion string,
	MinVersion string,
	Name string,
	Namespace string,
	Reference string,
	Status string,
	VerificationVersion string,
	VerifyDate string,
	Version string,
) SystemVersionStruct {
	return SystemVersionStruct{
		Type:                     "SystemVersion",
		BuildDate:                BuildDate,
		HotfixVersion:            HotfixVersion,
		InstallDate:              InstallDate,
		MinRebootOptionalVersion: MinRebootOptionalVersion,
		MinVersion:               MinVersion,
		Name:                     Name,
		Namespace:                Namespace,
		Reference:                Reference,
		Status:                   Status,
		VerificationVersion:      VerificationVersion,
		VerifyDate:               VerifyDate,
		Version:                  Version,
	}
}

// TCPStatsDatapointFactory is just a simple function to instantiate the TCPStatsDatapointStruct
func TCPStatsDatapointFactory(
	CongestionWindowSize *int,
	InBytes *int,
	InUnorderedBytes *int,
	OutBytes *int,
	ReceiveWindowSize *int,
	RetransmittedBytes *int,
	RetransmittedSegs *int,
	RoundTripTime *int,
	SendWindowSize *int,
	Timestamp string,
	UnacknowledgedBytes *int,
	UnsentBytes *int,
) TCPStatsDatapointStruct {
	return TCPStatsDatapointStruct{
		Type:                 "TCPStatsDatapoint",
		CongestionWindowSize: CongestionWindowSize,
		InBytes:              InBytes,
		InUnorderedBytes:     InUnorderedBytes,
		OutBytes:             OutBytes,
		ReceiveWindowSize:    ReceiveWindowSize,
		RetransmittedBytes:   RetransmittedBytes,
		RetransmittedSegs:    RetransmittedSegs,
		RoundTripTime:        RoundTripTime,
		SendWindowSize:       SendWindowSize,
		Timestamp:            Timestamp,
		UnacknowledgedBytes:  UnacknowledgedBytes,
		UnsentBytes:          UnsentBytes,
	}
}

// TCPStatsDatapointStreamFactory is just a simple function to instantiate the TCPStatsDatapointStreamStruct
func TCPStatsDatapointStreamFactory(
	Datapoints []Datapoint,
	LocalAddress string,
	LocalPort *int,
	RemoteAddress string,
	RemotePort *int,
	Service string,
) TCPStatsDatapointStreamStruct {
	return TCPStatsDatapointStreamStruct{
		Type:          "TCPStatsDatapointStream",
		Datapoints:    Datapoints,
		LocalAddress:  LocalAddress,
		LocalPort:     LocalPort,
		RemoteAddress: RemoteAddress,
		RemotePort:    RemotePort,
		Service:       Service,
	}
}

// TargetFilterFactory is just a simple function to instantiate the TargetFilterStruct
func TargetFilterFactory(
	Targets []string,
) TargetFilterStruct {
	return TargetFilterStruct{
		Type:    "TargetFilter",
		Targets: Targets,
	}
}

// TargetOwnerFilterFactory is just a simple function to instantiate the TargetOwnerFilterStruct
func TargetOwnerFilterFactory(
	Owners []string,
) TargetOwnerFilterStruct {
	return TargetOwnerFilterStruct{
		Type:   "TargetOwnerFilter",
		Owners: Owners,
	}
}

// ThemeFactory is just a simple function to instantiate the ThemeStruct
func ThemeFactory(
	Active *bool,
	Name string,
	Namespace string,
	Reference string,
	Values *SchemaDraftV4Struct,
) ThemeStruct {
	return ThemeStruct{
		Type:      "Theme",
		Active:    Active,
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
		Values:    Values,
	}
}

// ThemeReferenceParameterFactory is just a simple function to instantiate the ThemeReferenceParameterStruct
func ThemeReferenceParameterFactory(
	Reference string,
) ThemeReferenceParameterStruct {
	return ThemeReferenceParameterStruct{
		Type:      "ThemeReferenceParameter",
		Reference: Reference,
	}
}

// TimeConfigFactory is just a simple function to instantiate the TimeConfigStruct
func TimeConfigFactory(
	CurrentTime string,
	NtpConfig *NTPConfigStruct,
	SystemTimeZone string,
	SystemTimeZoneOffset *int,
	SystemTimeZoneOffsetString string,
) TimeConfigStruct {
	return TimeConfigStruct{
		Type:                       "TimeConfig",
		CurrentTime:                CurrentTime,
		NtpConfig:                  NtpConfig,
		SystemTimeZone:             SystemTimeZone,
		SystemTimeZoneOffset:       SystemTimeZoneOffset,
		SystemTimeZoneOffsetString: SystemTimeZoneOffsetString,
	}
}

// TimeRangeParametersFactory is just a simple function to instantiate the TimeRangeParametersStruct
func TimeRangeParametersFactory(
	EndTime string,
	StartTime string,
) TimeRangeParametersStruct {
	return TimeRangeParametersStruct{
		Type:      "TimeRangeParameters",
		EndTime:   EndTime,
		StartTime: StartTime,
	}
}

// TimeZoneFactory is just a simple function to instantiate the TimeZoneStruct
func TimeZoneFactory(
	Id string,
	Offset *int,
	OffsetString string,
) TimeZoneStruct {
	return TimeZoneStruct{
		Type:         "TimeZone",
		Id:           Id,
		Offset:       Offset,
		OffsetString: OffsetString,
	}
}

// TimeflowBookmarkFactory is just a simple function to instantiate the TimeflowBookmarkStruct
func TimeflowBookmarkFactory(
	Location string,
	Name string,
	Namespace string,
	Reference string,
	RetentionProof *bool,
	Tag string,
	Timeflow string,
	Timestamp string,
) TimeflowBookmarkStruct {
	return TimeflowBookmarkStruct{
		Type:           "TimeflowBookmark",
		Location:       Location,
		Name:           Name,
		Namespace:      Namespace,
		Reference:      Reference,
		RetentionProof: RetentionProof,
		Tag:            Tag,
		Timeflow:       Timeflow,
		Timestamp:      Timestamp,
	}
}

// TimeflowBookmarkCreateParametersFactory is just a simple function to instantiate the TimeflowBookmarkCreateParametersStruct
func TimeflowBookmarkCreateParametersFactory(
	Name string,
	RetentionProof *bool,
	Tag string,
	TimeflowPoint TimeflowPoint,
) TimeflowBookmarkCreateParametersStruct {
	return TimeflowBookmarkCreateParametersStruct{
		Type:           "TimeflowBookmarkCreateParameters",
		Name:           Name,
		RetentionProof: RetentionProof,
		Tag:            Tag,
		TimeflowPoint:  TimeflowPoint,
	}
}

// TimeflowDeletionDependencyFactory is just a simple function to instantiate the TimeflowDeletionDependencyStruct
func TimeflowDeletionDependencyFactory(
	ContainerName string,
	Dependencies []TypedObject,
	DisplayName string,
	Locked *bool,
	NamespaceName string,
	Prerequisites []*DeletionDependencyPrerequisiteStruct,
	Size float64,
	TargetReference string,
	TargetType string,
) TimeflowDeletionDependencyStruct {
	return TimeflowDeletionDependencyStruct{
		Type:            "TimeflowDeletionDependency",
		ContainerName:   ContainerName,
		Dependencies:    Dependencies,
		DisplayName:     DisplayName,
		Locked:          Locked,
		NamespaceName:   NamespaceName,
		Prerequisites:   Prerequisites,
		Size:            Size,
		TargetReference: TargetReference,
		TargetType:      TargetType,
	}
}

// TimeflowFilesystemLayoutFactory is just a simple function to instantiate the TimeflowFilesystemLayoutStruct
func TimeflowFilesystemLayoutFactory(
	ArchiveDirectory string,
	DataDirectory string,
	ExternalDirectory string,
	ScriptDirectory string,
	TargetDirectory string,
	TempDirectory string,
) TimeflowFilesystemLayoutStruct {
	return TimeflowFilesystemLayoutStruct{
		Type:              "TimeflowFilesystemLayout",
		ArchiveDirectory:  ArchiveDirectory,
		DataDirectory:     DataDirectory,
		ExternalDirectory: ExternalDirectory,
		ScriptDirectory:   ScriptDirectory,
		TargetDirectory:   TargetDirectory,
		TempDirectory:     TempDirectory,
	}
}

// TimeflowLogFetchParametersFactory is just a simple function to instantiate the TimeflowLogFetchParametersStruct
func TimeflowLogFetchParametersFactory(
	Credentials Credential,
	Directory string,
	EndLocation string,
	Host string,
	Port *int,
	SshVerificationStrategy SshVerificationStrategy,
	StartLocation string,
	Timeflow string,
	Username string,
) TimeflowLogFetchParametersStruct {
	return TimeflowLogFetchParametersStruct{
		Type:                    "TimeflowLogFetchParameters",
		Credentials:             Credentials,
		Directory:               Directory,
		EndLocation:             EndLocation,
		Host:                    Host,
		Port:                    Port,
		SshVerificationStrategy: SshVerificationStrategy,
		StartLocation:           StartLocation,
		Timeflow:                Timeflow,
		Username:                Username,
	}
}

// TimeflowPointBookmarkFactory is just a simple function to instantiate the TimeflowPointBookmarkStruct
func TimeflowPointBookmarkFactory(
	Bookmark string,
) TimeflowPointBookmarkStruct {
	return TimeflowPointBookmarkStruct{
		Type:     "TimeflowPointBookmark",
		Bookmark: Bookmark,
	}
}

// TimeflowPointBookmarkTagFactory is just a simple function to instantiate the TimeflowPointBookmarkTagStruct
func TimeflowPointBookmarkTagFactory(
	Container string,
	Tag string,
) TimeflowPointBookmarkTagStruct {
	return TimeflowPointBookmarkTagStruct{
		Type:      "TimeflowPointBookmarkTag",
		Container: Container,
		Tag:       Tag,
	}
}

// TimeflowPointLocationFactory is just a simple function to instantiate the TimeflowPointLocationStruct
func TimeflowPointLocationFactory(
	Location string,
	Timeflow string,
) TimeflowPointLocationStruct {
	return TimeflowPointLocationStruct{
		Type:     "TimeflowPointLocation",
		Location: Location,
		Timeflow: Timeflow,
	}
}

// TimeflowPointSemanticFactory is just a simple function to instantiate the TimeflowPointSemanticStruct
func TimeflowPointSemanticFactory(
	Container string,
	Location string,
) TimeflowPointSemanticStruct {
	return TimeflowPointSemanticStruct{
		Type:      "TimeflowPointSemantic",
		Container: Container,
		Location:  Location,
	}
}

// TimeflowPointSnapshotFactory is just a simple function to instantiate the TimeflowPointSnapshotStruct
func TimeflowPointSnapshotFactory(
	Snapshot string,
) TimeflowPointSnapshotStruct {
	return TimeflowPointSnapshotStruct{
		Type:     "TimeflowPointSnapshot",
		Snapshot: Snapshot,
	}
}

// TimeflowPointTimestampFactory is just a simple function to instantiate the TimeflowPointTimestampStruct
func TimeflowPointTimestampFactory(
	Timeflow string,
	Timestamp string,
) TimeflowPointTimestampStruct {
	return TimeflowPointTimestampStruct{
		Type:      "TimeflowPointTimestamp",
		Timeflow:  Timeflow,
		Timestamp: Timestamp,
	}
}

// TimeflowRangeFactory is just a simple function to instantiate the TimeflowRangeStruct
func TimeflowRangeFactory(
	EndPoint TimeflowPoint,
	Provisionable *bool,
	StartPoint TimeflowPoint,
) TimeflowRangeStruct {
	return TimeflowRangeStruct{
		Type:          "TimeflowRange",
		EndPoint:      EndPoint,
		Provisionable: Provisionable,
		StartPoint:    StartPoint,
	}
}

// TimeflowRangeParametersFactory is just a simple function to instantiate the TimeflowRangeParametersStruct
func TimeflowRangeParametersFactory(
	EndPoint TimeflowPoint,
	StartPoint TimeflowPoint,
) TimeflowRangeParametersStruct {
	return TimeflowRangeParametersStruct{
		Type:       "TimeflowRangeParameters",
		EndPoint:   EndPoint,
		StartPoint: StartPoint,
	}
}

// TimeflowRepairSSHFactory is just a simple function to instantiate the TimeflowRepairSSHStruct
func TimeflowRepairSSHFactory(
	Credentials Credential,
	Directory string,
	EndLocation string,
	Host string,
	Port *int,
	SshVerificationStrategy SshVerificationStrategy,
	StartLocation string,
	Username string,
) TimeflowRepairSSHStruct {
	return TimeflowRepairSSHStruct{
		Type:                    "TimeflowRepairSSH",
		Credentials:             Credentials,
		Directory:               Directory,
		EndLocation:             EndLocation,
		Host:                    Host,
		Port:                    Port,
		SshVerificationStrategy: SshVerificationStrategy,
		StartLocation:           StartLocation,
		Username:                Username,
	}
}

// TimeflowSnapshotDayRangeFactory is just a simple function to instantiate the TimeflowSnapshotDayRangeStruct
func TimeflowSnapshotDayRangeFactory(
	Count float64,
	Date string,
	EndOfDay string,
	StartOfDay string,
) TimeflowSnapshotDayRangeStruct {
	return TimeflowSnapshotDayRangeStruct{
		Type:       "TimeflowSnapshotDayRange",
		Count:      Count,
		Date:       Date,
		EndOfDay:   EndOfDay,
		StartOfDay: StartOfDay,
	}
}

// ToolkitFactory is just a simple function to instantiate the ToolkitStruct
func ToolkitFactory(
	BuildApi *APIVersionStruct,
	DefaultLocale string,
	DiscoveryDefinition *ToolkitDiscoveryDefinitionStruct,
	HostTypes []string,
	Language string,
	LinkedSourceDefinition ToolkitLinkedSource,
	Messages []*ToolkitLocaleStruct,
	Name string,
	Namespace string,
	PrettyName string,
	Reference string,
	Resources map[string]string,
	RootSquashEnabled *bool,
	SnapshotSchema *SchemaDraftV4Struct,
	UpgradeDefinition *ToolkitUpgradeDefinitionStruct,
	Version string,
	VirtualSourceDefinition *ToolkitVirtualSourceStruct,
) ToolkitStruct {
	return ToolkitStruct{
		Type:                    "Toolkit",
		BuildApi:                BuildApi,
		DefaultLocale:           DefaultLocale,
		DiscoveryDefinition:     DiscoveryDefinition,
		HostTypes:               HostTypes,
		Language:                Language,
		LinkedSourceDefinition:  LinkedSourceDefinition,
		Messages:                Messages,
		Name:                    Name,
		Namespace:               Namespace,
		PrettyName:              PrettyName,
		Reference:               Reference,
		Resources:               Resources,
		RootSquashEnabled:       RootSquashEnabled,
		SnapshotSchema:          SnapshotSchema,
		UpgradeDefinition:       UpgradeDefinition,
		Version:                 Version,
		VirtualSourceDefinition: VirtualSourceDefinition,
	}
}

// ToolkitDiscoveryDefinitionFactory is just a simple function to instantiate the ToolkitDiscoveryDefinitionStruct
func ToolkitDiscoveryDefinitionFactory(
	ManualSourceConfigDiscovery *bool,
	RepositoryDiscovery string,
	RepositoryIdentityFields []string,
	RepositoryNameField string,
	RepositorySchema *SchemaDraftV4Struct,
	SourceConfigDiscovery string,
	SourceConfigIdentityFields []string,
	SourceConfigNameField string,
	SourceConfigSchema *SchemaDraftV4Struct,
) ToolkitDiscoveryDefinitionStruct {
	return ToolkitDiscoveryDefinitionStruct{
		Type:                        "ToolkitDiscoveryDefinition",
		ManualSourceConfigDiscovery: ManualSourceConfigDiscovery,
		RepositoryDiscovery:         RepositoryDiscovery,
		RepositoryIdentityFields:    RepositoryIdentityFields,
		RepositoryNameField:         RepositoryNameField,
		RepositorySchema:            RepositorySchema,
		SourceConfigDiscovery:       SourceConfigDiscovery,
		SourceConfigIdentityFields:  SourceConfigIdentityFields,
		SourceConfigNameField:       SourceConfigNameField,
		SourceConfigSchema:          SourceConfigSchema,
	}
}

// ToolkitLinkedDirectSourceFactory is just a simple function to instantiate the ToolkitLinkedDirectSourceStruct
func ToolkitLinkedDirectSourceFactory(
	Parameters *SchemaDraftV4Struct,
	PostSnapshot string,
	PreSnapshot string,
	UsesGrandfatheredAppDataProperties *bool,
) ToolkitLinkedDirectSourceStruct {
	return ToolkitLinkedDirectSourceStruct{
		Type:                               "ToolkitLinkedDirectSource",
		Parameters:                         Parameters,
		PostSnapshot:                       PostSnapshot,
		PreSnapshot:                        PreSnapshot,
		UsesGrandfatheredAppDataProperties: UsesGrandfatheredAppDataProperties,
	}
}

// ToolkitLinkedStagedSourceFactory is just a simple function to instantiate the ToolkitLinkedStagedSourceStruct
func ToolkitLinkedStagedSourceFactory(
	MountSpec string,
	OwnershipSpec string,
	Parameters *SchemaDraftV4Struct,
	PostSnapshot string,
	PreSnapshot string,
	Resync string,
	StartStaging string,
	Status string,
	StopStaging string,
	Worker string,
) ToolkitLinkedStagedSourceStruct {
	return ToolkitLinkedStagedSourceStruct{
		Type:          "ToolkitLinkedStagedSource",
		MountSpec:     MountSpec,
		OwnershipSpec: OwnershipSpec,
		Parameters:    Parameters,
		PostSnapshot:  PostSnapshot,
		PreSnapshot:   PreSnapshot,
		Resync:        Resync,
		StartStaging:  StartStaging,
		Status:        Status,
		StopStaging:   StopStaging,
		Worker:        Worker,
	}
}

// ToolkitLocaleFactory is just a simple function to instantiate the ToolkitLocaleStruct
func ToolkitLocaleFactory(
	LocaleName string,
	Messages map[string]string,
) ToolkitLocaleStruct {
	return ToolkitLocaleStruct{
		Type:       "ToolkitLocale",
		LocaleName: LocaleName,
		Messages:   Messages,
	}
}

// ToolkitSnapshotParametersDefinitionFactory is just a simple function to instantiate the ToolkitSnapshotParametersDefinitionStruct
func ToolkitSnapshotParametersDefinitionFactory(
	Schema *SchemaDraftV4Struct,
) ToolkitSnapshotParametersDefinitionStruct {
	return ToolkitSnapshotParametersDefinitionStruct{
		Type:   "ToolkitSnapshotParametersDefinition",
		Schema: Schema,
	}
}

// ToolkitUpgradeDefinitionFactory is just a simple function to instantiate the ToolkitUpgradeDefinitionStruct
func ToolkitUpgradeDefinitionFactory(
	FromVersion string,
	UpgradeLinkedSource string,
	UpgradeManualSourceConfig string,
	UpgradeSnapshot string,
	UpgradeVirtualSource string,
) ToolkitUpgradeDefinitionStruct {
	return ToolkitUpgradeDefinitionStruct{
		Type:                      "ToolkitUpgradeDefinition",
		FromVersion:               FromVersion,
		UpgradeLinkedSource:       UpgradeLinkedSource,
		UpgradeManualSourceConfig: UpgradeManualSourceConfig,
		UpgradeSnapshot:           UpgradeSnapshot,
		UpgradeVirtualSource:      UpgradeVirtualSource,
	}
}

// ToolkitVirtualSourceFactory is just a simple function to instantiate the ToolkitVirtualSourceStruct
func ToolkitVirtualSourceFactory(
	Configure string,
	Initialize string,
	MountSpec string,
	OwnershipSpec string,
	Parameters *SchemaDraftV4Struct,
	PostSnapshot string,
	PreSnapshot string,
	Reconfigure string,
	Start string,
	Status string,
	Stop string,
	Unconfigure string,
) ToolkitVirtualSourceStruct {
	return ToolkitVirtualSourceStruct{
		Type:          "ToolkitVirtualSource",
		Configure:     Configure,
		Initialize:    Initialize,
		MountSpec:     MountSpec,
		OwnershipSpec: OwnershipSpec,
		Parameters:    Parameters,
		PostSnapshot:  PostSnapshot,
		PreSnapshot:   PreSnapshot,
		Reconfigure:   Reconfigure,
		Start:         Start,
		Status:        Status,
		Stop:          Stop,
		Unconfigure:   Unconfigure,
	}
}

// TracerouteInfoFactory is just a simple function to instantiate the TracerouteInfoStruct
func TracerouteInfoFactory(
	NetworkHops string,
) TracerouteInfoStruct {
	return TracerouteInfoStruct{
		Type:        "TracerouteInfo",
		NetworkHops: NetworkHops,
	}
}

// TransformationFactory is just a simple function to instantiate the TransformationStruct
func TransformationFactory(
	Container string,
	EnvironmentUser string,
	Name string,
	Namespace string,
	Operations []SourceOperation,
	PlatformParams BasePlatformParameters,
	Reference string,
	Repository string,
) TransformationStruct {
	return TransformationStruct{
		Type:            "Transformation",
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Name:            Name,
		Namespace:       Namespace,
		Operations:      Operations,
		PlatformParams:  PlatformParams,
		Reference:       Reference,
		Repository:      Repository,
	}
}

// TunableFactory is just a simple function to instantiate the TunableStruct
func TunableFactory(
	Name string,
	Subsystem string,
	Value string,
) TunableStruct {
	return TunableStruct{
		Type:      "Tunable",
		Name:      Name,
		Subsystem: Subsystem,
		Value:     Value,
	}
}

// TunableIdentifierFactory is just a simple function to instantiate the TunableIdentifierStruct
func TunableIdentifierFactory(
	Name string,
	Subsystem string,
) TunableIdentifierStruct {
	return TunableIdentifierStruct{
		Type:      "TunableIdentifier",
		Name:      Name,
		Subsystem: Subsystem,
	}
}

// UnixConnectivityParametersFactory is just a simple function to instantiate the UnixConnectivityParametersStruct
func UnixConnectivityParametersFactory(
	Address string,
	Credentials Credential,
	Port *int,
	SshVerificationStrategy SshVerificationStrategy,
	Username string,
) UnixConnectivityParametersStruct {
	return UnixConnectivityParametersStruct{
		Type:                    "UnixConnectivityParameters",
		Address:                 Address,
		Credentials:             Credentials,
		Port:                    Port,
		SshVerificationStrategy: SshVerificationStrategy,
		Username:                Username,
	}
}

// UnixHostFactory is just a simple function to instantiate the UnixHostStruct
func UnixHostFactory(
	Address string,
	DateAdded string,
	DspKeystoreAlias string,
	DspKeystorePassword string,
	DspKeystorePath string,
	DspTruststorePassword string,
	DspTruststorePath string,
	HostConfiguration *HostConfigurationStruct,
	HostRuntime *HostRuntimeStruct,
	JavaHome string,
	Name string,
	Namespace string,
	NfsAddressList []string,
	OracleJdbcKeystorePassword string,
	PrivilegeElevationProfile string,
	Reference string,
	SshPort *int,
	SshVerificationStrategy SshVerificationStrategy,
	ToolkitPath string,
) UnixHostStruct {
	return UnixHostStruct{
		Type:                       "UnixHost",
		Address:                    Address,
		DateAdded:                  DateAdded,
		DspKeystoreAlias:           DspKeystoreAlias,
		DspKeystorePassword:        DspKeystorePassword,
		DspKeystorePath:            DspKeystorePath,
		DspTruststorePassword:      DspTruststorePassword,
		DspTruststorePath:          DspTruststorePath,
		HostConfiguration:          HostConfiguration,
		HostRuntime:                HostRuntime,
		JavaHome:                   JavaHome,
		Name:                       Name,
		Namespace:                  Namespace,
		NfsAddressList:             NfsAddressList,
		OracleJdbcKeystorePassword: OracleJdbcKeystorePassword,
		PrivilegeElevationProfile:  PrivilegeElevationProfile,
		Reference:                  Reference,
		SshPort:                    SshPort,
		SshVerificationStrategy:    SshVerificationStrategy,
		ToolkitPath:                ToolkitPath,
	}
}

// UnixHostCreateParametersFactory is just a simple function to instantiate the UnixHostCreateParametersStruct
func UnixHostCreateParametersFactory(
	Host Host,
) UnixHostCreateParametersStruct {
	return UnixHostCreateParametersStruct{
		Type: "UnixHostCreateParameters",
		Host: Host,
	}
}

// UnixHostEnvironmentFactory is just a simple function to instantiate the UnixHostEnvironmentStruct
func UnixHostEnvironmentFactory(
	AseHostEnvironmentParameters *ASEHostEnvironmentParametersStruct,
	Description string,
	Enabled *bool,
	Host string,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	PrimaryUser string,
	Reference string,
) UnixHostEnvironmentStruct {
	return UnixHostEnvironmentStruct{
		Type:                         "UnixHostEnvironment",
		AseHostEnvironmentParameters: AseHostEnvironmentParameters,
		Description:                  Description,
		Enabled:                      Enabled,
		Host:                         Host,
		LogCollectionEnabled:         LogCollectionEnabled,
		Name:                         Name,
		Namespace:                    Namespace,
		PrimaryUser:                  PrimaryUser,
		Reference:                    Reference,
	}
}

// UnixRuntimeMountInformationFactory is just a simple function to instantiate the UnixRuntimeMountInformationStruct
func UnixRuntimeMountInformationFactory(
	Name string,
	Namespace string,
	Reference string,
) UnixRuntimeMountInformationStruct {
	return UnixRuntimeMountInformationStruct{
		Type:      "UnixRuntimeMountInformation",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// UpTimeInfoFactory is just a simple function to instantiate the UpTimeInfoStruct
func UpTimeInfoFactory(
	Mgmt string,
	Os string,
) UpTimeInfoStruct {
	return UpTimeInfoStruct{
		Type: "UpTimeInfo",
		Mgmt: Mgmt,
		Os:   Os,
	}
}

// UpgradeCheckResultFactory is just a simple function to instantiate the UpgradeCheckResultStruct
func UpgradeCheckResultFactory(
	Name string,
	Namespace string,
	Reference string,
) UpgradeCheckResultStruct {
	return UpgradeCheckResultStruct{
		Type:      "UpgradeCheckResult",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// UpgradeCheckResultsVersionParametersFactory is just a simple function to instantiate the UpgradeCheckResultsVersionParametersStruct
func UpgradeCheckResultsVersionParametersFactory(
	BundleId string,
	Reference string,
) UpgradeCheckResultsVersionParametersStruct {
	return UpgradeCheckResultsVersionParametersStruct{
		Type:      "UpgradeCheckResultsVersionParameters",
		BundleId:  BundleId,
		Reference: Reference,
	}
}

// UpgradeCompatibilityParametersFactory is just a simple function to instantiate the UpgradeCompatibilityParametersStruct
func UpgradeCompatibilityParametersFactory(
	Environment string,
	SourceRepository string,
) UpgradeCompatibilityParametersStruct {
	return UpgradeCompatibilityParametersStruct{
		Type:             "UpgradeCompatibilityParameters",
		Environment:      Environment,
		SourceRepository: SourceRepository,
	}
}

// UpgradeNotificationFactory is just a simple function to instantiate the UpgradeNotificationStruct
func UpgradeNotificationFactory() UpgradeNotificationStruct {
	return UpgradeNotificationStruct{
		Type: "UpgradeNotification",
	}
}

// UpgradeVerificationReportFactory is just a simple function to instantiate the UpgradeVerificationReportStruct
func UpgradeVerificationReportFactory(
	Name string,
	Namespace string,
	Reference string,
) UpgradeVerificationReportStruct {
	return UpgradeVerificationReportStruct{
		Type:      "UpgradeVerificationReport",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// UpgradeVerificationStepsFactory is just a simple function to instantiate the UpgradeVerificationStepsStruct
func UpgradeVerificationStepsFactory(
	Name string,
	Namespace string,
	Reference string,
) UpgradeVerificationStepsStruct {
	return UpgradeVerificationStepsStruct{
		Type:      "UpgradeVerificationSteps",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// UpgradeVersionChecksParameterFactory is just a simple function to instantiate the UpgradeVersionChecksParameterStruct
func UpgradeVersionChecksParameterFactory(
	UpgradeVersionReference string,
) UpgradeVersionChecksParameterStruct {
	return UpgradeVersionChecksParameterStruct{
		Type:                    "UpgradeVersionChecksParameter",
		UpgradeVersionReference: UpgradeVersionReference,
	}
}

// UserFactory is just a simple function to instantiate the UserStruct
func UserFactory(
	ApiUser *bool,
	AuthenticationType string,
	Credential *PasswordCredentialStruct,
	EmailAddress string,
	Enabled *bool,
	FirstName string,
	HomePhoneNumber string,
	IsDefault *bool,
	LastName string,
	Locale string,
	MobilePhoneNumber string,
	Name string,
	Namespace string,
	PasswordUpdateRequest string,
	Principal string,
	PublicKey []string,
	Reference string,
	SessionTimeout *int,
	UserType string,
	WorkPhoneNumber string,
) UserStruct {
	return UserStruct{
		Type:                  "User",
		ApiUser:               ApiUser,
		AuthenticationType:    AuthenticationType,
		Credential:            Credential,
		EmailAddress:          EmailAddress,
		Enabled:               Enabled,
		FirstName:             FirstName,
		HomePhoneNumber:       HomePhoneNumber,
		IsDefault:             IsDefault,
		LastName:              LastName,
		Locale:                Locale,
		MobilePhoneNumber:     MobilePhoneNumber,
		Name:                  Name,
		Namespace:             Namespace,
		PasswordUpdateRequest: PasswordUpdateRequest,
		Principal:             Principal,
		PublicKey:             PublicKey,
		Reference:             Reference,
		SessionTimeout:        SessionTimeout,
		UserType:              UserType,
		WorkPhoneNumber:       WorkPhoneNumber,
	}
}

// UserAuthInfoFactory is just a simple function to instantiate the UserAuthInfoStruct
func UserAuthInfoFactory(
	Authorizations []*AuthorizationStruct,
	JetStreamUserRole string,
	OwnerRole string,
	ProvisionerRole string,
	User *UserStruct,
) UserAuthInfoStruct {
	return UserAuthInfoStruct{
		Type:              "UserAuthInfo",
		Authorizations:    Authorizations,
		JetStreamUserRole: JetStreamUserRole,
		OwnerRole:         OwnerRole,
		ProvisionerRole:   ProvisionerRole,
		User:              User,
	}
}

// UserInterfaceConfigFactory is just a simple function to instantiate the UserInterfaceConfigStruct
func UserInterfaceConfigFactory(
	AnalyticsEnabled *bool,
) UserInterfaceConfigStruct {
	return UserInterfaceConfigStruct{
		Type:             "UserInterfaceConfig",
		AnalyticsEnabled: AnalyticsEnabled,
	}
}

// UserPathStorageFactory is just a simple function to instantiate the UserPathStorageStruct
func UserPathStorageFactory(
	Description string,
	Namespace string,
	Path string,
	Pathtype string,
	Reference string,
) UserPathStorageStruct {
	return UserPathStorageStruct{
		Type:        "UserPathStorage",
		Description: Description,
		Namespace:   Namespace,
		Path:        Path,
		Pathtype:    Pathtype,
		Reference:   Reference,
	}
}

// ValidateJavaParametersFactory is just a simple function to instantiate the ValidateJavaParametersStruct
func ValidateJavaParametersFactory(
	ConnectivityParameters ConnectivityParameters,
	JavaHome string,
) ValidateJavaParametersStruct {
	return ValidateJavaParametersStruct{
		Type:                   "ValidateJavaParameters",
		ConnectivityParameters: ConnectivityParameters,
		JavaHome:               JavaHome,
	}
}

// ValidateSMTPParametersFactory is just a simple function to instantiate the ValidateSMTPParametersStruct
func ValidateSMTPParametersFactory(
	Addresses []string,
	Config *SMTPConfigStruct,
) ValidateSMTPParametersStruct {
	return ValidateSMTPParametersStruct{
		Type:      "ValidateSMTPParameters",
		Addresses: Addresses,
		Config:    Config,
	}
}

// VerifyVersionParametersFactory is just a simple function to instantiate the VerifyVersionParametersStruct
func VerifyVersionParametersFactory(
	Defer *bool,
) VerifyVersionParametersStruct {
	return VerifyVersionParametersStruct{
		Type:  "VerifyVersionParameters",
		Defer: Defer,
	}
}

// VersionInfoFactory is just a simple function to instantiate the VersionInfoStruct
func VersionInfoFactory(
	BuildMetadata []string,
	Major *int,
	Micro *int,
	Minor *int,
	Patch *int,
	PreRelease []string,
) VersionInfoStruct {
	return VersionInfoStruct{
		Type:          "VersionInfo",
		BuildMetadata: BuildMetadata,
		Major:         Major,
		Micro:         Micro,
		Minor:         Minor,
		Patch:         Patch,
		PreRelease:    PreRelease,
	}
}

// VfsOpsDatapointStreamFactory is just a simple function to instantiate the VfsOpsDatapointStreamStruct
func VfsOpsDatapointStreamFactory(
	Cached *bool,
	Datapoints []Datapoint,
	Op string,
	Path string,
	Sync *bool,
) VfsOpsDatapointStreamStruct {
	return VfsOpsDatapointStreamStruct{
		Type:       "VfsOpsDatapointStream",
		Cached:     Cached,
		Datapoints: Datapoints,
		Op:         Op,
		Path:       Path,
		Sync:       Sync,
	}
}

// VirtualSourceOperationsFactory is just a simple function to instantiate the VirtualSourceOperationsStruct
func VirtualSourceOperationsFactory(
	ConfigureClone []SourceOperation,
	PostRefresh []SourceOperation,
	PostRollback []SourceOperation,
	PostSnapshot []SourceOperation,
	PostStart []SourceOperation,
	PostStop []SourceOperation,
	PreRefresh []SourceOperation,
	PreRollback []SourceOperation,
	PreSnapshot []SourceOperation,
	PreStart []SourceOperation,
	PreStop []SourceOperation,
) VirtualSourceOperationsStruct {
	return VirtualSourceOperationsStruct{
		Type:           "VirtualSourceOperations",
		ConfigureClone: ConfigureClone,
		PostRefresh:    PostRefresh,
		PostRollback:   PostRollback,
		PostSnapshot:   PostSnapshot,
		PostStart:      PostStart,
		PostStop:       PostStop,
		PreRefresh:     PreRefresh,
		PreRollback:    PreRollback,
		PreSnapshot:    PreSnapshot,
		PreStart:       PreStart,
		PreStop:        PreStop,
	}
}

// VirtualizationPlatformAPIRangeFactory is just a simple function to instantiate the VirtualizationPlatformAPIRangeStruct
func VirtualizationPlatformAPIRangeFactory(
	Max *VirtualizationPlatformAPIVersionStruct,
	Min *VirtualizationPlatformAPIVersionStruct,
) VirtualizationPlatformAPIRangeStruct {
	return VirtualizationPlatformAPIRangeStruct{
		Type: "VirtualizationPlatformAPIRange",
		Max:  Max,
		Min:  Min,
	}
}

// VirtualizationPlatformAPIVersionFactory is just a simple function to instantiate the VirtualizationPlatformAPIVersionStruct
func VirtualizationPlatformAPIVersionFactory(
	Major *int,
	Micro *int,
	Minor *int,
) VirtualizationPlatformAPIVersionStruct {
	return VirtualizationPlatformAPIVersionStruct{
		Type:  "VirtualizationPlatformAPIVersion",
		Major: Major,
		Micro: Micro,
		Minor: Minor,
	}
}

// WindowsClusterFactory is just a simple function to instantiate the WindowsClusterStruct
func WindowsClusterFactory(
	Address string,
	Description string,
	Enabled *bool,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	PrimaryUser string,
	Proxy string,
	Reference string,
	TargetProxy string,
) WindowsClusterStruct {
	return WindowsClusterStruct{
		Type:                 "WindowsCluster",
		Address:              Address,
		Description:          Description,
		Enabled:              Enabled,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Namespace:            Namespace,
		PrimaryUser:          PrimaryUser,
		Proxy:                Proxy,
		Reference:            Reference,
		TargetProxy:          TargetProxy,
	}
}

// WindowsClusterCreateParametersFactory is just a simple function to instantiate the WindowsClusterCreateParametersStruct
func WindowsClusterCreateParametersFactory(
	Cluster *WindowsClusterStruct,
	LogCollectionEnabled *bool,
	PrimaryUser *EnvironmentUserStruct,
	Target *bool,
) WindowsClusterCreateParametersStruct {
	return WindowsClusterCreateParametersStruct{
		Type:                 "WindowsClusterCreateParameters",
		Cluster:              Cluster,
		LogCollectionEnabled: LogCollectionEnabled,
		PrimaryUser:          PrimaryUser,
		Target:               Target,
	}
}

// WindowsClusterNodeFactory is just a simple function to instantiate the WindowsClusterNodeStruct
func WindowsClusterNodeFactory(
	Cluster string,
	Discovered *bool,
	Host string,
	Name string,
	Namespace string,
	Reference string,
) WindowsClusterNodeStruct {
	return WindowsClusterNodeStruct{
		Type:       "WindowsClusterNode",
		Cluster:    Cluster,
		Discovered: Discovered,
		Host:       Host,
		Name:       Name,
		Namespace:  Namespace,
		Reference:  Reference,
	}
}

// WindowsConnectivityParametersFactory is just a simple function to instantiate the WindowsConnectivityParametersStruct
func WindowsConnectivityParametersFactory(
	Address string,
	Credentials Credential,
	Port *int,
	Username string,
) WindowsConnectivityParametersStruct {
	return WindowsConnectivityParametersStruct{
		Type:        "WindowsConnectivityParameters",
		Address:     Address,
		Credentials: Credentials,
		Port:        Port,
		Username:    Username,
	}
}

// WindowsHostFactory is just a simple function to instantiate the WindowsHostStruct
func WindowsHostFactory(
	Address string,
	ConnectorAuthenticationKey string,
	ConnectorPort float64,
	DateAdded string,
	DspKeystoreAlias string,
	DspKeystorePassword string,
	DspKeystorePath string,
	DspTruststorePassword string,
	DspTruststorePath string,
	HostConfiguration *HostConfigurationStruct,
	HostRuntime *HostRuntimeStruct,
	JavaHome string,
	Name string,
	Namespace string,
	NfsAddressList []string,
	PrivilegeElevationProfile string,
	Reference string,
	SshPort *int,
) WindowsHostStruct {
	return WindowsHostStruct{
		Type:                       "WindowsHost",
		Address:                    Address,
		ConnectorAuthenticationKey: ConnectorAuthenticationKey,
		ConnectorPort:              ConnectorPort,
		DateAdded:                  DateAdded,
		DspKeystoreAlias:           DspKeystoreAlias,
		DspKeystorePassword:        DspKeystorePassword,
		DspKeystorePath:            DspKeystorePath,
		DspTruststorePassword:      DspTruststorePassword,
		DspTruststorePath:          DspTruststorePath,
		HostConfiguration:          HostConfiguration,
		HostRuntime:                HostRuntime,
		JavaHome:                   JavaHome,
		Name:                       Name,
		Namespace:                  Namespace,
		NfsAddressList:             NfsAddressList,
		PrivilegeElevationProfile:  PrivilegeElevationProfile,
		Reference:                  Reference,
		SshPort:                    SshPort,
	}
}

// WindowsHostCreateParametersFactory is just a simple function to instantiate the WindowsHostCreateParametersStruct
func WindowsHostCreateParametersFactory(
	Host Host,
) WindowsHostCreateParametersStruct {
	return WindowsHostCreateParametersStruct{
		Type: "WindowsHostCreateParameters",
		Host: Host,
	}
}

// WindowsHostEnvironmentFactory is just a simple function to instantiate the WindowsHostEnvironmentStruct
func WindowsHostEnvironmentFactory(
	Description string,
	Enabled *bool,
	Host string,
	LogCollectionEnabled *bool,
	Name string,
	Namespace string,
	PrimaryUser string,
	Proxy string,
	Reference string,
) WindowsHostEnvironmentStruct {
	return WindowsHostEnvironmentStruct{
		Type:                 "WindowsHostEnvironment",
		Description:          Description,
		Enabled:              Enabled,
		Host:                 Host,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Namespace:            Namespace,
		PrimaryUser:          PrimaryUser,
		Proxy:                Proxy,
		Reference:            Reference,
	}
}

// WindowsRuntimeMountInformationFactory is just a simple function to instantiate the WindowsRuntimeMountInformationStruct
func WindowsRuntimeMountInformationFactory(
	Name string,
	Namespace string,
	Reference string,
) WindowsRuntimeMountInformationStruct {
	return WindowsRuntimeMountInformationStruct{
		Type:      "WindowsRuntimeMountInformation",
		Name:      Name,
		Namespace: Namespace,
		Reference: Reference,
	}
}

// WorkflowFunctionDefinitionFactory is just a simple function to instantiate the WorkflowFunctionDefinitionStruct
func WorkflowFunctionDefinitionFactory(
	InputSchema *SchemaDraftV4Struct,
	Name string,
	OutputSchema *SchemaDraftV4Struct,
) WorkflowFunctionDefinitionStruct {
	return WorkflowFunctionDefinitionStruct{
		Type:         "WorkflowFunctionDefinition",
		InputSchema:  InputSchema,
		Name:         Name,
		OutputSchema: OutputSchema,
	}
}

// X500DistinguishedNameCompositeFactory is just a simple function to instantiate the X500DistinguishedNameCompositeStruct
func X500DistinguishedNameCompositeFactory(
	Dname string,
) X500DistinguishedNameCompositeStruct {
	return X500DistinguishedNameCompositeStruct{
		Type:  "X500DistinguishedNameComposite",
		Dname: Dname,
	}
}

// X500DistinguishedNameFieldsFactory is just a simple function to instantiate the X500DistinguishedNameFieldsStruct
func X500DistinguishedNameFieldsFactory(
	City string,
	CommonName string,
	Country string,
	Organization string,
	OrganizationUnit string,
	StateRegion string,
) X500DistinguishedNameFieldsStruct {
	return X500DistinguishedNameFieldsStruct{
		Type:             "X500DistinguishedNameFields",
		City:             City,
		CommonName:       CommonName,
		Country:          Country,
		Organization:     Organization,
		OrganizationUnit: OrganizationUnit,
		StateRegion:      StateRegion,
	}
}
