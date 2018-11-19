package delphix

// ASEAttachDataFactory is just a simple function to instantiate the ASEAttachDataStruct
func ASEAttachDataFactory(
	DbUser string,
	DumpCredentials string,
	ExternalFilePath string,
	LoadLocation string,
	MountBase string,
	Operations *LinkedSourceOperationsStruct,
	StagingPostScript string,
	StagingPreScript string,
	ValidatedSyncMode string,
) ASEAttachDataStruct {
	return ASEAttachDataStruct{
		DbUser:            DbUser,
		DumpCredentials:   DumpCredentials,
		ExternalFilePath:  ExternalFilePath,
		LoadLocation:      LoadLocation,
		MountBase:         MountBase,
		Operations:        Operations,
		StagingPostScript: StagingPostScript,
		StagingPreScript:  StagingPreScript,
		ValidatedSyncMode: ValidatedSyncMode,
	}
}

// ASEBackupLocationFactory is just a simple function to instantiate the ASEBackupLocationStruct
func ASEBackupLocationFactory(
	BackupHost string,
	BackupHostUser string,
	BackupServerName string,
) ASEBackupLocationStruct {
	return ASEBackupLocationStruct{
		BackupHost:       BackupHost,
		BackupHostUser:   BackupHostUser,
		BackupServerName: BackupServerName,
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
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Repository:      Repository,
	}
}

// ASEDBContainerFactory is just a simple function to instantiate the ASEDBContainerStruct
func ASEDBContainerFactory(
	Description string,
	Group string,
	Name string,
	SourcingPolicy *SourcingPolicyStruct,
) ASEDBContainerStruct {
	return ASEDBContainerStruct{
		Description:    Description,
		Group:          Group,
		Name:           Name,
		SourcingPolicy: SourcingPolicy,
	}
}

// ASEExportParametersFactory is just a simple function to instantiate the ASEExportParametersStruct
func ASEExportParametersFactory(
	ConfigParams map[string]string,
	FileMappingRules string,
	RecoverDatabase *bool,
) ASEExportParametersStruct {
	return ASEExportParametersStruct{
		ConfigParams:     ConfigParams,
		FileMappingRules: FileMappingRules,
		RecoverDatabase:  RecoverDatabase,
	}
}

// ASEHostEnvironmentParametersFactory is just a simple function to instantiate the ASEHostEnvironmentParametersStruct
func ASEHostEnvironmentParametersFactory(
	Credentials Credential,
	DbUser string,
) ASEHostEnvironmentParametersStruct {
	return ASEHostEnvironmentParametersStruct{
		Credentials: Credentials,
		DbUser:      DbUser,
	}
}

// ASEInstanceFactory is just a simple function to instantiate the ASEInstanceStruct
func ASEInstanceFactory(
	Credentials Credential,
	DbUser string,
	Environment string,
	InstallationPath string,
	InstanceName string,
	InstanceOwner string,
	IsqlPath string,
	LinkingEnabled *bool,
	Name string,
	Ports []*int,
	ProvisioningEnabled *bool,
	ServicePrincipalName string,
	Staging *bool,
	Version string,
) ASEInstanceStruct {
	return ASEInstanceStruct{
		Credentials:          Credentials,
		DbUser:               DbUser,
		Environment:          Environment,
		InstallationPath:     InstallationPath,
		InstanceName:         InstanceName,
		InstanceOwner:        InstanceOwner,
		IsqlPath:             IsqlPath,
		LinkingEnabled:       LinkingEnabled,
		Name:                 Name,
		Ports:                Ports,
		ProvisioningEnabled:  ProvisioningEnabled,
		ServicePrincipalName: ServicePrincipalName,
		Staging:              Staging,
		Version:              Version,
	}
}

// ASELinkDataFactory is just a simple function to instantiate the ASELinkDataStruct
func ASELinkDataFactory(
	DbUser string,
	DumpCredentials string,
	ExternalFilePath string,
	LoadLocation string,
	MountBase string,
	Operations *LinkedSourceOperationsStruct,
	SourcingPolicy *SourcingPolicyStruct,
	StagingPostScript string,
	StagingPreScript string,
	ValidatedSyncMode string,
) ASELinkDataStruct {
	return ASELinkDataStruct{
		DbUser:            DbUser,
		DumpCredentials:   DumpCredentials,
		ExternalFilePath:  ExternalFilePath,
		LoadLocation:      LoadLocation,
		MountBase:         MountBase,
		Operations:        Operations,
		SourcingPolicy:    SourcingPolicy,
		StagingPostScript: StagingPostScript,
		StagingPreScript:  StagingPreScript,
		ValidatedSyncMode: ValidatedSyncMode,
	}
}

// ASELinkedSourceFactory is just a simple function to instantiate the ASELinkedSourceStruct
func ASELinkedSourceFactory(
	Config string,
	DumpCredentials string,
	ExternalFilePath string,
	LoadBackupPath string,
	LoadLocation string,
	LogCollectionEnabled *bool,
	Name string,
	Operations *LinkedSourceOperationsStruct,
	ValidatedSyncMode string,
) ASELinkedSourceStruct {
	return ASELinkedSourceStruct{
		Config:               Config,
		DumpCredentials:      DumpCredentials,
		ExternalFilePath:     ExternalFilePath,
		LoadBackupPath:       LoadBackupPath,
		LoadLocation:         LoadLocation,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Operations:           Operations,
		ValidatedSyncMode:    ValidatedSyncMode,
	}
}

// ASEProvisionParametersFactory is just a simple function to instantiate the ASEProvisionParametersStruct
func ASEProvisionParametersFactory(
	Masked *bool,
	MaskingJob string,
	TruncateLogOnCheckpoint *bool,
) ASEProvisionParametersStruct {
	return ASEProvisionParametersStruct{
		Masked:                  Masked,
		MaskingJob:              MaskingJob,
		TruncateLogOnCheckpoint: TruncateLogOnCheckpoint,
	}
}

// ASESIConfigFactory is just a simple function to instantiate the ASESIConfigStruct
func ASESIConfigFactory(
	DatabaseName string,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Repository string,
) ASESIConfigStruct {
	return ASESIConfigStruct{
		DatabaseName:    DatabaseName,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Repository:      Repository,
	}
}

// ASESnapshotFactory is just a simple function to instantiate the ASESnapshotStruct
func ASESnapshotFactory(
	Name string,
) ASESnapshotStruct {
	return ASESnapshotStruct{
		Name: Name,
	}
}

// ASESpecificBackupSyncParametersFactory is just a simple function to instantiate the ASESpecificBackupSyncParametersStruct
func ASESpecificBackupSyncParametersFactory(
	BackupFiles []string,
) ASESpecificBackupSyncParametersStruct {
	return ASESpecificBackupSyncParametersStruct{
		BackupFiles: BackupFiles,
	}
}

// ASEStagingSourceFactory is just a simple function to instantiate the ASEStagingSourceStruct
func ASEStagingSourceFactory(
	Config string,
	LogCollectionEnabled *bool,
	Name string,
	PostScript string,
	PreScript string,
) ASEStagingSourceStruct {
	return ASEStagingSourceStruct{
		Config:               Config,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		PostScript:           PostScript,
		PreScript:            PreScript,
	}
}

// ASETimeflowFactory is just a simple function to instantiate the ASETimeflowStruct
func ASETimeflowFactory(
	Name string,
) ASETimeflowStruct {
	return ASETimeflowStruct{
		Name: Name,
	}
}

// ASETimeflowPointFactory is just a simple function to instantiate the ASETimeflowPointStruct
func ASETimeflowPointFactory(
	Location string,
	Timestamp string,
) ASETimeflowPointStruct {
	return ASETimeflowPointStruct{
		Location:  Location,
		Timestamp: Timestamp,
	}
}

// ASEVirtualSourceFactory is just a simple function to instantiate the ASEVirtualSourceStruct
func ASEVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Operations *VirtualSourceOperationsStruct,
) ASEVirtualSourceStruct {
	return ASEVirtualSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Operations:                      Operations,
	}
}

// AlertActionEmailListFactory is just a simple function to instantiate the AlertActionEmailListStruct
func AlertActionEmailListFactory(
	Format string,
) AlertActionEmailListStruct {
	return AlertActionEmailListStruct{
		Format: Format,
	}
}

// AlertActionEmailUserFactory is just a simple function to instantiate the AlertActionEmailUserStruct
func AlertActionEmailUserFactory(
	Format string,
) AlertActionEmailUserStruct {
	return AlertActionEmailUserStruct{
		Format: Format,
	}
}

// AlertProfileFactory is just a simple function to instantiate the AlertProfileStruct
func AlertProfileFactory(
	Actions []AlertAction,
	FilterSpec AlertFilter,
) AlertProfileStruct {
	return AlertProfileStruct{
		Actions:    Actions,
		FilterSpec: FilterSpec,
	}
}

// AndFilterFactory is just a simple function to instantiate the AndFilterStruct
func AndFilterFactory(
	SubFilters []AlertFilter,
) AndFilterStruct {
	return AndFilterStruct{
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
		Environment: Environment,
		MountPath:   MountPath,
		Ordinal:     Ordinal,
		SharedPath:  SharedPath,
	}
}

// AppDataContainerFactory is just a simple function to instantiate the AppDataContainerStruct
func AppDataContainerFactory(
	Description string,
	Group string,
	Name string,
	SourcingPolicy *SourcingPolicyStruct,
) AppDataContainerStruct {
	return AppDataContainerStruct{
		Description:    Description,
		Group:          Group,
		Name:           Name,
		SourcingPolicy: SourcingPolicy,
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
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Payload:         Payload,
		Repository:      Repository,
	}
}

// AppDataDirectLinkDataFactory is just a simple function to instantiate the AppDataDirectLinkDataStruct
func AppDataDirectLinkDataFactory(
	Excludes []string,
	FollowSymlinks []string,
	Operations *LinkedSourceOperationsStruct,
	SourcingPolicy *SourcingPolicyStruct,
) AppDataDirectLinkDataStruct {
	return AppDataDirectLinkDataStruct{
		Excludes:       Excludes,
		FollowSymlinks: FollowSymlinks,
		Operations:     Operations,
		SourcingPolicy: SourcingPolicy,
	}
}

// AppDataDirectSourceConfigFactory is just a simple function to instantiate the AppDataDirectSourceConfigStruct
func AppDataDirectSourceConfigFactory(
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Parameters *JsonStruct,
	Path string,
	Repository string,
) AppDataDirectSourceConfigStruct {
	return AppDataDirectSourceConfigStruct{
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Parameters:      Parameters,
		Path:            Path,
		Repository:      Repository,
	}
}

// AppDataExportParametersFactory is just a simple function to instantiate the AppDataExportParametersStruct
func AppDataExportParametersFactory(
	FilesystemLayout *AppDataFilesystemLayoutStruct,
) AppDataExportParametersStruct {
	return AppDataExportParametersStruct{
		FilesystemLayout: FilesystemLayout,
	}
}

// AppDataLinkedDirectSourceFactory is just a simple function to instantiate the AppDataLinkedDirectSourceStruct
func AppDataLinkedDirectSourceFactory(
	Config string,
	Excludes []string,
	FollowSymlinks []string,
	LogCollectionEnabled *bool,
	Name string,
	Operations *LinkedSourceOperationsStruct,
	Parameters *JsonStruct,
) AppDataLinkedDirectSourceStruct {
	return AppDataLinkedDirectSourceStruct{
		Config:               Config,
		Excludes:             Excludes,
		FollowSymlinks:       FollowSymlinks,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Operations:           Operations,
		Parameters:           Parameters,
	}
}

// AppDataLinkedStagedSourceFactory is just a simple function to instantiate the AppDataLinkedStagedSourceStruct
func AppDataLinkedStagedSourceFactory(
	Config string,
	LogCollectionEnabled *bool,
	Name string,
	Operations *LinkedSourceOperationsStruct,
	Parameters *JsonStruct,
	StagingEnvironment string,
	StagingEnvironmentUser string,
	StagingMountBase string,
) AppDataLinkedStagedSourceStruct {
	return AppDataLinkedStagedSourceStruct{
		Config:                 Config,
		LogCollectionEnabled:   LogCollectionEnabled,
		Name:                   Name,
		Operations:             Operations,
		Parameters:             Parameters,
		StagingEnvironment:     StagingEnvironment,
		StagingEnvironmentUser: StagingEnvironmentUser,
		StagingMountBase:       StagingMountBase,
	}
}

// AppDataPlatformParametersFactory is just a simple function to instantiate the AppDataPlatformParametersStruct
func AppDataPlatformParametersFactory(
	Payload *JsonStruct,
) AppDataPlatformParametersStruct {
	return AppDataPlatformParametersStruct{
		Payload: Payload,
	}
}

// AppDataProvisionParametersFactory is just a simple function to instantiate the AppDataProvisionParametersStruct
func AppDataProvisionParametersFactory(
	Masked *bool,
	MaskingJob string,
) AppDataProvisionParametersStruct {
	return AppDataProvisionParametersStruct{
		Masked:     Masked,
		MaskingJob: MaskingJob,
	}
}

// AppDataRepositoryFactory is just a simple function to instantiate the AppDataRepositoryStruct
func AppDataRepositoryFactory(
	Environment string,
	LinkingEnabled *bool,
	Name string,
	Parameters *JsonStruct,
	ProvisioningEnabled *bool,
	Staging *bool,
	Toolkit string,
	Version string,
) AppDataRepositoryStruct {
	return AppDataRepositoryStruct{
		Environment:         Environment,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		Parameters:          Parameters,
		ProvisioningEnabled: ProvisioningEnabled,
		Staging:             Staging,
		Toolkit:             Toolkit,
		Version:             Version,
	}
}

// AppDataSnapshotFactory is just a simple function to instantiate the AppDataSnapshotStruct
func AppDataSnapshotFactory(
	Name string,
) AppDataSnapshotStruct {
	return AppDataSnapshotStruct{
		Name: Name,
	}
}

// AppDataStagedLinkDataFactory is just a simple function to instantiate the AppDataStagedLinkDataStruct
func AppDataStagedLinkDataFactory(
	Operations *LinkedSourceOperationsStruct,
	SourcingPolicy *SourcingPolicyStruct,
	StagingMountBase string,
) AppDataStagedLinkDataStruct {
	return AppDataStagedLinkDataStruct{
		Operations:       Operations,
		SourcingPolicy:   SourcingPolicy,
		StagingMountBase: StagingMountBase,
	}
}

// AppDataStagedSourceConfigFactory is just a simple function to instantiate the AppDataStagedSourceConfigStruct
func AppDataStagedSourceConfigFactory(
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Parameters *JsonStruct,
	Repository string,
) AppDataStagedSourceConfigStruct {
	return AppDataStagedSourceConfigStruct{
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Parameters:      Parameters,
		Repository:      Repository,
	}
}

// AppDataSyncParametersFactory is just a simple function to instantiate the AppDataSyncParametersStruct
func AppDataSyncParametersFactory(
	Resync *bool,
) AppDataSyncParametersStruct {
	return AppDataSyncParametersStruct{
		Resync: Resync,
	}
}

// AppDataTimeflowFactory is just a simple function to instantiate the AppDataTimeflowStruct
func AppDataTimeflowFactory(
	Name string,
) AppDataTimeflowStruct {
	return AppDataTimeflowStruct{
		Name: Name,
	}
}

// AppDataTimeflowPointFactory is just a simple function to instantiate the AppDataTimeflowPointStruct
func AppDataTimeflowPointFactory(
	Location string,
	Timestamp string,
) AppDataTimeflowPointStruct {
	return AppDataTimeflowPointStruct{
		Location:  Location,
		Timestamp: Timestamp,
	}
}

// AppDataVirtualSourceFactory is just a simple function to instantiate the AppDataVirtualSourceStruct
func AppDataVirtualSourceFactory(
	AdditionalMountPoints []*AppDataAdditionalMountPointStruct,
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	LogCollectionEnabled *bool,
	Name string,
	Operations *VirtualSourceOperationsStruct,
	Parameters *JsonStruct,
) AppDataVirtualSourceStruct {
	return AppDataVirtualSourceStruct{
		AdditionalMountPoints:           AdditionalMountPoints,
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		LogCollectionEnabled:            LogCollectionEnabled,
		Name:                            Name,
		Operations:                      Operations,
		Parameters:                      Parameters,
	}
}

// AppDataWindowsTimeflowFactory is just a simple function to instantiate the AppDataWindowsTimeflowStruct
func AppDataWindowsTimeflowFactory(
	Name string,
) AppDataWindowsTimeflowStruct {
	return AppDataWindowsTimeflowStruct{
		Name: Name,
	}
}

// ApplyVersionParametersFactory is just a simple function to instantiate the ApplyVersionParametersStruct
func ApplyVersionParametersFactory(
	Defer *bool,
	EnableSourcesOnFailure *bool,
	IgnoreQuiesceSourcesFailures *bool,
	QuiesceSources *bool,
	Reboot *bool,
	Verify *bool,
) ApplyVersionParametersStruct {
	return ApplyVersionParametersStruct{
		Defer:                        Defer,
		EnableSourcesOnFailure:       EnableSourcesOnFailure,
		IgnoreQuiesceSourcesFailures: IgnoreQuiesceSourcesFailures,
		QuiesceSources:               QuiesceSources,
		Reboot:                       Reboot,
		Verify:                       Verify,
	}
}

// AuthorizationFactory is just a simple function to instantiate the AuthorizationStruct
func AuthorizationFactory(
	Name string,
	Role string,
	Target string,
	User string,
) AuthorizationStruct {
	return AuthorizationStruct{
		Name:   Name,
		Role:   Role,
		Target: Target,
		User:   User,
	}
}

// AuthorizationConfigFactory is just a simple function to instantiate the AuthorizationConfigStruct
func AuthorizationConfigFactory(
	EnvironmentAndHostAuth *bool,
) AuthorizationConfigStruct {
	return AuthorizationConfigStruct{
		EnvironmentAndHostAuth: EnvironmentAndHostAuth,
	}
}

// BooleanEqualConstraintFactory is just a simple function to instantiate the BooleanEqualConstraintStruct
func BooleanEqualConstraintFactory(
	AxisName string,
	Equals *bool,
) BooleanEqualConstraintStruct {
	return BooleanEqualConstraintStruct{
		AxisName: AxisName,
		Equals:   Equals,
	}
}

// CaCertificateFactory is just a simple function to instantiate the CaCertificateStruct
func CaCertificateFactory(
	Name string,
) CaCertificateStruct {
	return CaCertificateStruct{
		Name: Name,
	}
}

// CertificateSigningRequestFactory is just a simple function to instantiate the CertificateSigningRequestStruct
func CertificateSigningRequestFactory(
	Name string,
) CertificateSigningRequestStruct {
	return CertificateSigningRequestStruct{
		Name: Name,
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
		Dname:                   Dname,
		EndEntity:               EndEntity,
		ForceReplace:            ForceReplace,
		KeyPair:                 KeyPair,
		SubjectAlternativeNames: SubjectAlternativeNames,
	}
}

// ConfiguredStorageDeviceFactory is just a simple function to instantiate the ConfiguredStorageDeviceStruct
func ConfiguredStorageDeviceFactory(
	Name string,
) ConfiguredStorageDeviceStruct {
	return ConfiguredStorageDeviceStruct{
		Name: Name,
	}
}

// ConnectorConnectivityFactory is just a simple function to instantiate the ConnectorConnectivityStruct
func ConnectorConnectivityFactory(
	Proxy string,
) ConnectorConnectivityStruct {
	return ConnectorConnectivityStruct{
		Proxy: Proxy,
	}
}

// CreateMaskingJobTransformationParametersFactory is just a simple function to instantiate the CreateMaskingJobTransformationParametersStruct
func CreateMaskingJobTransformationParametersFactory(
	Container Container,
	EnvironmentUser string,
	Repository string,
) CreateMaskingJobTransformationParametersStruct {
	return CreateMaskingJobTransformationParametersStruct{
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
		NewCredential: NewCredential,
		OldCredential: OldCredential,
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
		DestinationType:         DestinationType,
		Direction:               Direction,
		RemoteDelphixEngineInfo: RemoteDelphixEngineInfo,
		RemoteHost:              RemoteHost,
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
		BandwidthLimit: BandwidthLimit,
		Compression:    Compression,
		Encryption:     Encryption,
		NumConnections: NumConnections,
	}
}

// DatabaseTemplateFactory is just a simple function to instantiate the DatabaseTemplateStruct
func DatabaseTemplateFactory(
	Description string,
	Name string,
	Parameters map[string]string,
	SourceType string,
) DatabaseTemplateStruct {
	return DatabaseTemplateStruct{
		Description: Description,
		Name:        Name,
		Parameters:  Parameters,
		SourceType:  SourceType,
	}
}

// DelphixManagedBackupIngestionStrategyFactory is just a simple function to instantiate the DelphixManagedBackupIngestionStrategyStruct
func DelphixManagedBackupIngestionStrategyFactory(
	CompressionEnabled *bool,
) DelphixManagedBackupIngestionStrategyStruct {
	return DelphixManagedBackupIngestionStrategyStruct{
		CompressionEnabled: CompressionEnabled,
	}
}

// DomainFactory is just a simple function to instantiate the DomainStruct
func DomainFactory(
	Name string,
) DomainStruct {
	return DomainStruct{
		Name: Name,
	}
}

// EcdsaKeyPairFactory is just a simple function to instantiate the EcdsaKeyPairStruct
func EcdsaKeyPairFactory(
	KeySize *int,
	SignatureAlgorithm string,
) EcdsaKeyPairStruct {
	return EcdsaKeyPairStruct{
		KeySize:            KeySize,
		SignatureAlgorithm: SignatureAlgorithm,
	}
}

// EndEntityCertificateFactory is just a simple function to instantiate the EndEntityCertificateStruct
func EndEntityCertificateFactory(
	Name string,
) EndEntityCertificateStruct {
	return EndEntityCertificateStruct{
		Name: Name,
	}
}

// EnumEqualConstraintFactory is just a simple function to instantiate the EnumEqualConstraintStruct
func EnumEqualConstraintFactory(
	AxisName string,
	Equals string,
) EnumEqualConstraintStruct {
	return EnumEqualConstraintStruct{
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
	UserId *int,
) EnvironmentUserStruct {
	return EnvironmentUserStruct{
		Credential:  Credential,
		Environment: Environment,
		GroupId:     GroupId,
		Name:        Name,
		UserId:      UserId,
	}
}

// EventFilterFactory is just a simple function to instantiate the EventFilterStruct
func EventFilterFactory(
	EventTypes []string,
) EventFilterStruct {
	return EventFilterStruct{
		EventTypes: EventTypes,
	}
}

// ExternalBackupIngestionStrategyFactory is just a simple function to instantiate the ExternalBackupIngestionStrategyStruct
func ExternalBackupIngestionStrategyFactory(
	ValidatedSyncMode string,
) ExternalBackupIngestionStrategyStruct {
	return ExternalBackupIngestionStrategyStruct{
		ValidatedSyncMode: ValidatedSyncMode,
	}
}

// FaultResolveParametersFactory is just a simple function to instantiate the FaultResolveParametersStruct
func FaultResolveParametersFactory(
	Comments string,
	Ignore *bool,
) FaultResolveParametersStruct {
	return FaultResolveParametersStruct{
		Comments: Comments,
		Ignore:   Ignore,
	}
}

// FractionPlugParametersFactory is just a simple function to instantiate the FractionPlugParametersStruct
func FractionPlugParametersFactory(
	SchemasPrefix string,
	TablespacesPrefix string,
) FractionPlugParametersStruct {
	return FractionPlugParametersStruct{
		SchemasPrefix:     SchemasPrefix,
		TablespacesPrefix: TablespacesPrefix,
	}
}

// GroupFactory is just a simple function to instantiate the GroupStruct
func GroupFactory(
	Description string,
	Name string,
) GroupStruct {
	return GroupStruct{
		Description: Description,
		Name:        Name,
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
		HostEnvironment:      HostEnvironment,
		HostParameters:       HostParameters,
		LogCollectionEnabled: LogCollectionEnabled,
		PrimaryUser:          PrimaryUser,
	}
}

// HostPrivilegeElevationProfileFactory is just a simple function to instantiate the HostPrivilegeElevationProfileStruct
func HostPrivilegeElevationProfileFactory(
	IsDefault *bool,
	Name string,
	Version string,
) HostPrivilegeElevationProfileStruct {
	return HostPrivilegeElevationProfileStruct{
		IsDefault: IsDefault,
		Name:      Name,
		Version:   Version,
	}
}

// HostPrivilegeElevationProfileScriptFactory is just a simple function to instantiate the HostPrivilegeElevationProfileScriptStruct
func HostPrivilegeElevationProfileScriptFactory(
	Contents string,
	Name string,
	Profile string,
) HostPrivilegeElevationProfileScriptStruct {
	return HostPrivilegeElevationProfileScriptStruct{
		Contents: Contents,
		Name:     Name,
		Profile:  Profile,
	}
}

// HttpConnectorConfigFactory is just a simple function to instantiate the HttpConnectorConfigStruct
func HttpConnectorConfigFactory(
	HttpPort *int,
	HttpsPort *int,
) HttpConnectorConfigStruct {
	return HttpConnectorConfigStruct{
		HttpPort:  HttpPort,
		HttpsPort: HttpsPort,
	}
}

// IntegerEqualConstraintFactory is just a simple function to instantiate the IntegerEqualConstraintStruct
func IntegerEqualConstraintFactory(
	AxisName string,
	Equals *int,
) IntegerEqualConstraintStruct {
	return IntegerEqualConstraintStruct{
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
		AxisName: AxisName,
		LessThan: LessThan,
	}
}

// InterfaceAddressFactory is just a simple function to instantiate the InterfaceAddressStruct
func InterfaceAddressFactory(
	Address string,
	AddressType string,
	EnableSSH *bool,
) InterfaceAddressStruct {
	return InterfaceAddressStruct{
		Address:     Address,
		AddressType: AddressType,
		EnableSSH:   EnableSSH,
	}
}

// JSBookmarkFactory is just a simple function to instantiate the JSBookmarkStruct
func JSBookmarkFactory(
	Branch string,
	Description string,
	Expiration string,
	Name string,
	Shared *bool,
	Tags []string,
) JSBookmarkStruct {
	return JSBookmarkStruct{
		Branch:      Branch,
		Description: Description,
		Expiration:  Expiration,
		Name:        Name,
		Shared:      Shared,
		Tags:        Tags,
	}
}

// JSBranchFactory is just a simple function to instantiate the JSBranchStruct
func JSBranchFactory(
	Name string,
) JSBranchStruct {
	return JSBranchStruct{
		Name: Name,
	}
}

// JSDataContainerFactory is just a simple function to instantiate the JSDataContainerStruct
func JSDataContainerFactory(
	Name string,
) JSDataContainerStruct {
	return JSDataContainerStruct{
		Name: Name,
	}
}

// JSDataContainerCreateWithRefreshParametersFactory is just a simple function to instantiate the JSDataContainerCreateWithRefreshParametersStruct
func JSDataContainerCreateWithRefreshParametersFactory(
	Notes string,
	Owners []string,
	Properties map[string]string,
) JSDataContainerCreateWithRefreshParametersStruct {
	return JSDataContainerCreateWithRefreshParametersStruct{
		Notes:      Notes,
		Owners:     Owners,
		Properties: Properties,
	}
}

// JSDataContainerCreateWithoutRefreshParametersFactory is just a simple function to instantiate the JSDataContainerCreateWithoutRefreshParametersStruct
func JSDataContainerCreateWithoutRefreshParametersFactory(
	Notes string,
	Owners []string,
	Properties map[string]string,
) JSDataContainerCreateWithoutRefreshParametersStruct {
	return JSDataContainerCreateWithoutRefreshParametersStruct{
		Notes:      Notes,
		Owners:     Owners,
		Properties: Properties,
	}
}

// JSDataContainerDeleteParametersFactory is just a simple function to instantiate the JSDataContainerDeleteParametersStruct
func JSDataContainerDeleteParametersFactory(
	DeleteDataSources *bool,
) JSDataContainerDeleteParametersStruct {
	return JSDataContainerDeleteParametersStruct{
		DeleteDataSources: DeleteDataSources,
	}
}

// JSDataContainerRefreshParametersFactory is just a simple function to instantiate the JSDataContainerRefreshParametersStruct
func JSDataContainerRefreshParametersFactory(
	ForceOption *bool,
) JSDataContainerRefreshParametersStruct {
	return JSDataContainerRefreshParametersStruct{
		ForceOption: ForceOption,
	}
}

// JSDataContainerResetParametersFactory is just a simple function to instantiate the JSDataContainerResetParametersStruct
func JSDataContainerResetParametersFactory(
	ForceOption *bool,
) JSDataContainerResetParametersStruct {
	return JSDataContainerResetParametersStruct{
		ForceOption: ForceOption,
	}
}

// JSDataContainerRestoreParametersFactory is just a simple function to instantiate the JSDataContainerRestoreParametersStruct
func JSDataContainerRestoreParametersFactory(
	ForceOption *bool,
) JSDataContainerRestoreParametersStruct {
	return JSDataContainerRestoreParametersStruct{
		ForceOption: ForceOption,
	}
}

// JSDataSourceFactory is just a simple function to instantiate the JSDataSourceStruct
func JSDataSourceFactory(
	Description string,
	Name string,
	Priority *int,
	Properties map[string]string,
) JSDataSourceStruct {
	return JSDataSourceStruct{
		Description: Description,
		Name:        Name,
		Priority:    Priority,
		Properties:  Properties,
	}
}

// JSDataSourceCreateParametersFactory is just a simple function to instantiate the JSDataSourceCreateParametersStruct
func JSDataSourceCreateParametersFactory(
	Properties map[string]string,
) JSDataSourceCreateParametersStruct {
	return JSDataSourceCreateParametersStruct{
		Properties: Properties,
	}
}

// JSDataTemplateFactory is just a simple function to instantiate the JSDataTemplateStruct
func JSDataTemplateFactory(
	ConfirmTimeConsumingOperations *bool,
	Name string,
) JSDataTemplateStruct {
	return JSDataTemplateStruct{
		ConfirmTimeConsumingOperations: ConfirmTimeConsumingOperations,
		Name:                           Name,
	}
}

// JSDataTemplateCreateParametersFactory is just a simple function to instantiate the JSDataTemplateCreateParametersStruct
func JSDataTemplateCreateParametersFactory(
	Notes string,
	Properties map[string]string,
) JSDataTemplateCreateParametersStruct {
	return JSDataTemplateCreateParametersStruct{
		Notes:      Notes,
		Properties: Properties,
	}
}

// JSOperationFactory is just a simple function to instantiate the JSOperationStruct
func JSOperationFactory(
	Name string,
) JSOperationStruct {
	return JSOperationStruct{
		Name: Name,
	}
}

// JSOperationEndpointBranchParametersFactory is just a simple function to instantiate the JSOperationEndpointBranchParametersStruct
func JSOperationEndpointBranchParametersFactory(
	Branch string,
) JSOperationEndpointBranchParametersStruct {
	return JSOperationEndpointBranchParametersStruct{
		Branch: Branch,
	}
}

// JSOperationEndpointDataLayoutParametersFactory is just a simple function to instantiate the JSOperationEndpointDataLayoutParametersStruct
func JSOperationEndpointDataLayoutParametersFactory(
	DataLayout string,
) JSOperationEndpointDataLayoutParametersStruct {
	return JSOperationEndpointDataLayoutParametersStruct{
		DataLayout: DataLayout,
	}
}

// JobFactory is just a simple function to instantiate the JobStruct
func JobFactory(
	Name string,
) JobStruct {
	return JobStruct{
		Name: Name,
	}
}

// KerberosConfigFactory is just a simple function to instantiate the KerberosConfigStruct
func KerberosConfigFactory(
	Kdcs []*KerberosKDCStruct,
	Keytab string,
	Name string,
	Principal string,
	Realm string,
) KerberosConfigStruct {
	return KerberosConfigStruct{
		Kdcs:      Kdcs,
		Keytab:    Keytab,
		Name:      Name,
		Principal: Principal,
		Realm:     Realm,
	}
}

// KerberosKDCFactory is just a simple function to instantiate the KerberosKDCStruct
func KerberosKDCFactory(
	Hostname string,
	Port *int,
) KerberosKDCStruct {
	return KerberosKDCStruct{
		Hostname: Hostname,
		Port:     Port,
	}
}

// LdapServerFactory is just a simple function to instantiate the LdapServerStruct
func LdapServerFactory(
	Name string,
) LdapServerStruct {
	return LdapServerStruct{
		Name: Name,
	}
}

// LinkParametersFactory is just a simple function to instantiate the LinkParametersStruct
func LinkParametersFactory(
	Description string,
	LinkData LinkData,
) LinkParametersStruct {
	return LinkParametersStruct{
		Description: Description,
		LinkData:    LinkData,
	}
}

// LinkedSourceOperationsFactory is just a simple function to instantiate the LinkedSourceOperationsStruct
func LinkedSourceOperationsFactory(
	PostSync []SourceOperation,
	PreSync []SourceOperation,
) LinkedSourceOperationsStruct {
	return LinkedSourceOperationsStruct{
		PostSync: PostSync,
		PreSync:  PreSync,
	}
}

// LocaleSettingsFactory is just a simple function to instantiate the LocaleSettingsStruct
func LocaleSettingsFactory(
	Locale string,
	Name string,
) LocaleSettingsStruct {
	return LocaleSettingsStruct{
		Locale: Locale,
		Name:   Name,
	}
}

// MSSqlAttachDataFactory is just a simple function to instantiate the MSSqlAttachDataStruct
func MSSqlAttachDataFactory(
	BackupLocationCredentials *PasswordCredentialStruct,
	BackupLocationUser string,
	EncryptionKey string,
	ExternalFilePath string,
	IngestionStrategy IngestionStrategy,
	MssqlNetbackupConfig string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
	SharedBackupLocations []string,
	SourceHostUser string,
	StagingPostScript string,
	StagingPreScript string,
) MSSqlAttachDataStruct {
	return MSSqlAttachDataStruct{
		BackupLocationCredentials: BackupLocationCredentials,
		BackupLocationUser:        BackupLocationUser,
		EncryptionKey:             EncryptionKey,
		ExternalFilePath:          ExternalFilePath,
		IngestionStrategy:         IngestionStrategy,
		MssqlNetbackupConfig:      MssqlNetbackupConfig,
		Operations:                Operations,
		PptHostUser:               PptHostUser,
		SharedBackupLocations:     SharedBackupLocations,
		SourceHostUser:            SourceHostUser,
		StagingPostScript:         StagingPostScript,
		StagingPreScript:          StagingPreScript,
	}
}

// MSSqlAvailabilityGroupFactory is just a simple function to instantiate the MSSqlAvailabilityGroupStruct
func MSSqlAvailabilityGroupFactory(
	Environment string,
	FulltextInstalled *bool,
	InternalVersion *int,
	LinkingEnabled *bool,
	Name string,
	ProvisioningEnabled *bool,
	Staging *bool,
	Version string,
) MSSqlAvailabilityGroupStruct {
	return MSSqlAvailabilityGroupStruct{
		Environment:         Environment,
		FulltextInstalled:   FulltextInstalled,
		InternalVersion:     InternalVersion,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		ProvisioningEnabled: ProvisioningEnabled,
		Staging:             Staging,
		Version:             Version,
	}
}

// MSSqlAvailabilityGroupDBConfigFactory is just a simple function to instantiate the MSSqlAvailabilityGroupDBConfigStruct
func MSSqlAvailabilityGroupDBConfigFactory(
	DatabaseName string,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	RecoveryModel string,
	Repository string,
) MSSqlAvailabilityGroupDBConfigStruct {
	return MSSqlAvailabilityGroupDBConfigStruct{
		DatabaseName:    DatabaseName,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		RecoveryModel:   RecoveryModel,
		Repository:      Repository,
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
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Repository:      Repository,
	}
}

// MSSqlDatabaseContainerFactory is just a simple function to instantiate the MSSqlDatabaseContainerStruct
func MSSqlDatabaseContainerFactory(
	Description string,
	Group string,
	Name string,
	SourcingPolicy *SourcingPolicyStruct,
) MSSqlDatabaseContainerStruct {
	return MSSqlDatabaseContainerStruct{
		Description:    Description,
		Group:          Group,
		Name:           Name,
		SourcingPolicy: SourcingPolicy,
	}
}

// MSSqlExportParametersFactory is just a simple function to instantiate the MSSqlExportParametersStruct
func MSSqlExportParametersFactory(
	ConfigParams map[string]string,
	FileMappingRules string,
	RecoverDatabase *bool,
	RecoveryModel string,
) MSSqlExportParametersStruct {
	return MSSqlExportParametersStruct{
		ConfigParams:     ConfigParams,
		FileMappingRules: FileMappingRules,
		RecoverDatabase:  RecoverDatabase,
		RecoveryModel:    RecoveryModel,
	}
}

// MSSqlFailoverClusterDBConfigFactory is just a simple function to instantiate the MSSqlFailoverClusterDBConfigStruct
func MSSqlFailoverClusterDBConfigFactory(
	DatabaseName string,
	DriveLetter string,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	RecoveryModel string,
	Repository string,
) MSSqlFailoverClusterDBConfigStruct {
	return MSSqlFailoverClusterDBConfigStruct{
		DatabaseName:    DatabaseName,
		DriveLetter:     DriveLetter,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		RecoveryModel:   RecoveryModel,
		Repository:      Repository,
	}
}

// MSSqlFailoverClusterRepositoryFactory is just a simple function to instantiate the MSSqlFailoverClusterRepositoryStruct
func MSSqlFailoverClusterRepositoryFactory(
	Environment string,
	FulltextInstalled *bool,
	InternalVersion *int,
	LinkingEnabled *bool,
	Name string,
	ProvisioningEnabled *bool,
	Staging *bool,
	Version string,
) MSSqlFailoverClusterRepositoryStruct {
	return MSSqlFailoverClusterRepositoryStruct{
		Environment:         Environment,
		FulltextInstalled:   FulltextInstalled,
		InternalVersion:     InternalVersion,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		ProvisioningEnabled: ProvisioningEnabled,
		Staging:             Staging,
		Version:             Version,
	}
}

// MSSqlInstanceFactory is just a simple function to instantiate the MSSqlInstanceStruct
func MSSqlInstanceFactory(
	Environment string,
	FulltextInstalled *bool,
	InstallationPath string,
	InstanceName string,
	InstanceOwner string,
	InternalVersion *int,
	LinkingEnabled *bool,
	Name string,
	Port *int,
	ProvisioningEnabled *bool,
	ServerName string,
	Staging *bool,
	Version string,
) MSSqlInstanceStruct {
	return MSSqlInstanceStruct{
		Environment:         Environment,
		FulltextInstalled:   FulltextInstalled,
		InstallationPath:    InstallationPath,
		InstanceName:        InstanceName,
		InstanceOwner:       InstanceOwner,
		InternalVersion:     InternalVersion,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		Port:                Port,
		ProvisioningEnabled: ProvisioningEnabled,
		ServerName:          ServerName,
		Staging:             Staging,
		Version:             Version,
	}
}

// MSSqlLinkDataFactory is just a simple function to instantiate the MSSqlLinkDataStruct
func MSSqlLinkDataFactory(
	BackupLocationCredentials *PasswordCredentialStruct,
	BackupLocationUser string,
	EncryptionKey string,
	ExternalFilePath string,
	IngestionStrategy IngestionStrategy,
	MssqlNetbackupConfig string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
	SharedBackupLocations []string,
	SourceHostUser string,
	SourcingPolicy *SourcingPolicyStruct,
	StagingPostScript string,
	StagingPreScript string,
) MSSqlLinkDataStruct {
	return MSSqlLinkDataStruct{
		BackupLocationCredentials: BackupLocationCredentials,
		BackupLocationUser:        BackupLocationUser,
		EncryptionKey:             EncryptionKey,
		ExternalFilePath:          ExternalFilePath,
		IngestionStrategy:         IngestionStrategy,
		MssqlNetbackupConfig:      MssqlNetbackupConfig,
		Operations:                Operations,
		PptHostUser:               PptHostUser,
		SharedBackupLocations:     SharedBackupLocations,
		SourceHostUser:            SourceHostUser,
		SourcingPolicy:            SourcingPolicy,
		StagingPostScript:         StagingPostScript,
		StagingPreScript:          StagingPreScript,
	}
}

// MSSqlLinkedSourceFactory is just a simple function to instantiate the MSSqlLinkedSourceStruct
func MSSqlLinkedSourceFactory(
	BackupLocationCredentials *PasswordCredentialStruct,
	BackupLocationUser string,
	Config string,
	EncryptionKey string,
	ExternalFilePath string,
	IngestionStrategy IngestionStrategy,
	LogCollectionEnabled *bool,
	MssqlNetbackupConfig string,
	Name string,
	Operations *LinkedSourceOperationsStruct,
	SharedBackupLocations []string,
) MSSqlLinkedSourceStruct {
	return MSSqlLinkedSourceStruct{
		BackupLocationCredentials: BackupLocationCredentials,
		BackupLocationUser:        BackupLocationUser,
		Config:                    Config,
		EncryptionKey:             EncryptionKey,
		ExternalFilePath:          ExternalFilePath,
		IngestionStrategy:         IngestionStrategy,
		LogCollectionEnabled:      LogCollectionEnabled,
		MssqlNetbackupConfig:      MssqlNetbackupConfig,
		Name:                      Name,
		Operations:                Operations,
		SharedBackupLocations:     SharedBackupLocations,
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
		ConfigParams:     ConfigParams,
		ConfigTemplate:   ConfigTemplate,
		MasterName:       MasterName,
		SourceClientName: SourceClientName,
	}
}

// MSSqlNewCopyOnlyFullBackupSyncParametersFactory is just a simple function to instantiate the MSSqlNewCopyOnlyFullBackupSyncParametersStruct
func MSSqlNewCopyOnlyFullBackupSyncParametersFactory(
	CompressionEnabled *bool,
) MSSqlNewCopyOnlyFullBackupSyncParametersStruct {
	return MSSqlNewCopyOnlyFullBackupSyncParametersStruct{
		CompressionEnabled: CompressionEnabled,
	}
}

// MSSqlProvisionParametersFactory is just a simple function to instantiate the MSSqlProvisionParametersStruct
func MSSqlProvisionParametersFactory(
	Masked *bool,
	MaskingJob string,
) MSSqlProvisionParametersStruct {
	return MSSqlProvisionParametersStruct{
		Masked:     Masked,
		MaskingJob: MaskingJob,
	}
}

// MSSqlSIConfigFactory is just a simple function to instantiate the MSSqlSIConfigStruct
func MSSqlSIConfigFactory(
	DatabaseName string,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	RecoveryModel string,
	Repository string,
) MSSqlSIConfigStruct {
	return MSSqlSIConfigStruct{
		DatabaseName:    DatabaseName,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		RecoveryModel:   RecoveryModel,
		Repository:      Repository,
	}
}

// MSSqlSnapshotFactory is just a simple function to instantiate the MSSqlSnapshotStruct
func MSSqlSnapshotFactory(
	Name string,
) MSSqlSnapshotStruct {
	return MSSqlSnapshotStruct{
		Name: Name,
	}
}

// MSSqlStagingSourceFactory is just a simple function to instantiate the MSSqlStagingSourceStruct
func MSSqlStagingSourceFactory(
	Config string,
	LogCollectionEnabled *bool,
	Name string,
	PostScript string,
	PreScript string,
) MSSqlStagingSourceStruct {
	return MSSqlStagingSourceStruct{
		Config:               Config,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		PostScript:           PostScript,
		PreScript:            PreScript,
	}
}

// MSSqlTimeflowFactory is just a simple function to instantiate the MSSqlTimeflowStruct
func MSSqlTimeflowFactory(
	Name string,
) MSSqlTimeflowStruct {
	return MSSqlTimeflowStruct{
		Name: Name,
	}
}

// MSSqlTimeflowPointFactory is just a simple function to instantiate the MSSqlTimeflowPointStruct
func MSSqlTimeflowPointFactory(
	Location string,
	Timestamp string,
) MSSqlTimeflowPointStruct {
	return MSSqlTimeflowPointStruct{
		Location:  Location,
		Timestamp: Timestamp,
	}
}

// MSSqlVirtualSourceFactory is just a simple function to instantiate the MSSqlVirtualSourceStruct
func MSSqlVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	Name string,
	Operations *VirtualSourceOperationsStruct,
	PostScript string,
	PreScript string,
) MSSqlVirtualSourceStruct {
	return MSSqlVirtualSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		Name:                            Name,
		Operations:                      Operations,
		PostScript:                      PostScript,
		PreScript:                       PreScript,
	}
}

// MaskingJobFactory is just a simple function to instantiate the MaskingJobStruct
func MaskingJobFactory(
	AssociatedContainer string,
	MaskingJobId string,
	Name string,
) MaskingJobStruct {
	return MaskingJobStruct{
		AssociatedContainer: AssociatedContainer,
		MaskingJobId:        MaskingJobId,
		Name:                Name,
	}
}

// MaskingServiceConfigFactory is just a simple function to instantiate the MaskingServiceConfigStruct
func MaskingServiceConfigFactory(
	Name string,
) MaskingServiceConfigStruct {
	return MaskingServiceConfigStruct{
		Name: Name,
	}
}

// MigrateCompatibilityParametersFactory is just a simple function to instantiate the MigrateCompatibilityParametersStruct
func MigrateCompatibilityParametersFactory(
	Environment string,
) MigrateCompatibilityParametersStruct {
	return MigrateCompatibilityParametersStruct{
		Environment: Environment,
	}
}

// MySQLAttachDataFactory is just a simple function to instantiate the MySQLAttachDataStruct
func MySQLAttachDataFactory(
	ConfigParams map[string]string,
	Operations *LinkedSourceOperationsStruct,
) MySQLAttachDataStruct {
	return MySQLAttachDataStruct{
		ConfigParams: ConfigParams,
		Operations:   Operations,
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
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Port:            Port,
		Repository:      Repository,
	}
}

// MySQLDatabaseContainerFactory is just a simple function to instantiate the MySQLDatabaseContainerStruct
func MySQLDatabaseContainerFactory(
	Description string,
	Group string,
	Name string,
	SourcingPolicy *SourcingPolicyStruct,
) MySQLDatabaseContainerStruct {
	return MySQLDatabaseContainerStruct{
		Description:    Description,
		Group:          Group,
		Name:           Name,
		SourcingPolicy: SourcingPolicy,
	}
}

// MySQLExistingMySQLDumpSyncParametersFactory is just a simple function to instantiate the MySQLExistingMySQLDumpSyncParametersStruct
func MySQLExistingMySQLDumpSyncParametersFactory(
	ReplicationCoordinates MySQLReplicationCoordinates,
) MySQLExistingMySQLDumpSyncParametersStruct {
	return MySQLExistingMySQLDumpSyncParametersStruct{
		ReplicationCoordinates: ReplicationCoordinates,
	}
}

// MySQLExportParametersFactory is just a simple function to instantiate the MySQLExportParametersStruct
func MySQLExportParametersFactory(
	ConfigParams map[string]string,
	FileMappingRules string,
	RecoverDatabase *bool,
) MySQLExportParametersStruct {
	return MySQLExportParametersStruct{
		ConfigParams:     ConfigParams,
		FileMappingRules: FileMappingRules,
		RecoverDatabase:  RecoverDatabase,
	}
}

// MySQLInstallFactory is just a simple function to instantiate the MySQLInstallStruct
func MySQLInstallFactory(
	Environment string,
	InstallationPath string,
	LinkingEnabled *bool,
	Name string,
	ProvisioningEnabled *bool,
	Staging *bool,
	Version string,
) MySQLInstallStruct {
	return MySQLInstallStruct{
		Environment:         Environment,
		InstallationPath:    InstallationPath,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		ProvisioningEnabled: ProvisioningEnabled,
		Staging:             Staging,
		Version:             Version,
	}
}

// MySQLLinkDataFactory is just a simple function to instantiate the MySQLLinkDataStruct
func MySQLLinkDataFactory(
	ConfigParams map[string]string,
	Operations *LinkedSourceOperationsStruct,
	SourcingPolicy *SourcingPolicyStruct,
	StagingHostUser string,
) MySQLLinkDataStruct {
	return MySQLLinkDataStruct{
		ConfigParams:    ConfigParams,
		Operations:      Operations,
		SourcingPolicy:  SourcingPolicy,
		StagingHostUser: StagingHostUser,
	}
}

// MySQLLinkedSourceFactory is just a simple function to instantiate the MySQLLinkedSourceStruct
func MySQLLinkedSourceFactory(
	Config string,
	ConfigParams map[string]string,
	LogCollectionEnabled *bool,
	Name string,
	Operations *LinkedSourceOperationsStruct,
) MySQLLinkedSourceStruct {
	return MySQLLinkedSourceStruct{
		Config:               Config,
		ConfigParams:         ConfigParams,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Operations:           Operations,
	}
}

// MySQLProvisionParametersFactory is just a simple function to instantiate the MySQLProvisionParametersStruct
func MySQLProvisionParametersFactory(
	Masked *bool,
	MaskingJob string,
) MySQLProvisionParametersStruct {
	return MySQLProvisionParametersStruct{
		Masked:     Masked,
		MaskingJob: MaskingJob,
	}
}

// MySQLServerConfigFactory is just a simple function to instantiate the MySQLServerConfigStruct
func MySQLServerConfigFactory(
	Credentials *PasswordCredentialStruct,
	DataDirectory string,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Port *int,
	Repository string,
) MySQLServerConfigStruct {
	return MySQLServerConfigStruct{
		Credentials:     Credentials,
		DataDirectory:   DataDirectory,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Port:            Port,
		Repository:      Repository,
	}
}

// MySQLSnapshotFactory is just a simple function to instantiate the MySQLSnapshotStruct
func MySQLSnapshotFactory(
	Name string,
) MySQLSnapshotStruct {
	return MySQLSnapshotStruct{
		Name: Name,
	}
}

// MySQLStagingSourceFactory is just a simple function to instantiate the MySQLStagingSourceStruct
func MySQLStagingSourceFactory(
	Config string,
	LogCollectionEnabled *bool,
	Name string,
	PostScript string,
	PreScript string,
) MySQLStagingSourceStruct {
	return MySQLStagingSourceStruct{
		Config:               Config,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		PostScript:           PostScript,
		PreScript:            PreScript,
	}
}

// MySQLTimeflowFactory is just a simple function to instantiate the MySQLTimeflowStruct
func MySQLTimeflowFactory(
	Name string,
) MySQLTimeflowStruct {
	return MySQLTimeflowStruct{
		Name: Name,
	}
}

// MySQLTimeflowPointFactory is just a simple function to instantiate the MySQLTimeflowPointStruct
func MySQLTimeflowPointFactory(
	Location string,
	Timestamp string,
) MySQLTimeflowPointStruct {
	return MySQLTimeflowPointStruct{
		Location:  Location,
		Timestamp: Timestamp,
	}
}

// MySQLVirtualSourceFactory is just a simple function to instantiate the MySQLVirtualSourceStruct
func MySQLVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Operations *VirtualSourceOperationsStruct,
) MySQLVirtualSourceStruct {
	return MySQLVirtualSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Operations:                      Operations,
	}
}

// MySQLXtraBackupSyncParametersFactory is just a simple function to instantiate the MySQLXtraBackupSyncParametersStruct
func MySQLXtraBackupSyncParametersFactory(
	ReplicationCoordinates MySQLReplicationCoordinates,
) MySQLXtraBackupSyncParametersStruct {
	return MySQLXtraBackupSyncParametersStruct{
		ReplicationCoordinates: ReplicationCoordinates,
	}
}

// NTPConfigFactory is just a simple function to instantiate the NTPConfigStruct
func NTPConfigFactory(
	Enabled *bool,
	UseMulticast *bool,
) NTPConfigStruct {
	return NTPConfigStruct{
		Enabled:      Enabled,
		UseMulticast: UseMulticast,
	}
}

// NamespaceFactory is just a simple function to instantiate the NamespaceStruct
func NamespaceFactory(
	FailedOver *bool,
	Name string,
) NamespaceStruct {
	return NamespaceStruct{
		FailedOver: FailedOver,
		Name:       Name,
	}
}

// NamespaceFailoverParametersFactory is just a simple function to instantiate the NamespaceFailoverParametersStruct
func NamespaceFailoverParametersFactory(
	SmartFailover *bool,
) NamespaceFailoverParametersStruct {
	return NamespaceFailoverParametersStruct{
		SmartFailover: SmartFailover,
	}
}

// NetworkDSPTestFactory is just a simple function to instantiate the NetworkDSPTestStruct
func NetworkDSPTestFactory(
	Name string,
) NetworkDSPTestStruct {
	return NetworkDSPTestStruct{
		Name: Name,
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
) NetworkDSPTestParametersStruct {
	return NetworkDSPTestParametersStruct{
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
	}
}

// NetworkInterfaceFactory is just a simple function to instantiate the NetworkInterfaceStruct
func NetworkInterfaceFactory(
	Name string,
) NetworkInterfaceStruct {
	return NetworkInterfaceStruct{
		Name: Name,
	}
}

// NetworkLatencyTestFactory is just a simple function to instantiate the NetworkLatencyTestStruct
func NetworkLatencyTestFactory(
	Name string,
) NetworkLatencyTestStruct {
	return NetworkLatencyTestStruct{
		Name: Name,
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
		Destination: Destination,
	}
}

// NetworkThroughputTestFactory is just a simple function to instantiate the NetworkThroughputTestStruct
func NetworkThroughputTestFactory(
	Name string,
) NetworkThroughputTestStruct {
	return NetworkThroughputTestStruct{
		Name: Name,
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

// NotFilterFactory is just a simple function to instantiate the NotFilterStruct
func NotFilterFactory(
	SubFilter AlertFilter,
) NotFilterStruct {
	return NotFilterStruct{
		SubFilter: SubFilter,
	}
}

// NullConstraintFactory is just a simple function to instantiate the NullConstraintStruct
func NullConstraintFactory(
	AxisName string,
) NullConstraintStruct {
	return NullConstraintStruct{
		AxisName: AxisName,
	}
}

// OperationTemplateFactory is just a simple function to instantiate the OperationTemplateStruct
func OperationTemplateFactory(
	Description string,
	Name string,
	Operation SourceOperation,
) OperationTemplateStruct {
	return OperationTemplateStruct{
		Description: Description,
		Name:        Name,
		Operation:   Operation,
	}
}

// OrFilterFactory is just a simple function to instantiate the OrFilterStruct
func OrFilterFactory(
	SubFilters []AlertFilter,
) OrFilterStruct {
	return OrFilterStruct{
		SubFilters: SubFilters,
	}
}

// OracleAddLiveSourceParametersFactory is just a simple function to instantiate the OracleAddLiveSourceParametersStruct
func OracleAddLiveSourceParametersFactory(
	Credential Credential,
	Username string,
) OracleAddLiveSourceParametersStruct {
	return OracleAddLiveSourceParametersStruct{
		Credential: Credential,
		Username:   Username,
	}
}

// OracleAttachDataFactory is just a simple function to instantiate the OracleAttachDataStruct
func OracleAttachDataFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	DoubleSync *bool,
	EncryptedLinkingEnabled *bool,
	ExternalFilePath string,
	FilesPerSet *int,
	Force *bool,
	LinkNow *bool,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	RmanChannels *int,
) OracleAttachDataStruct {
	return OracleAttachDataStruct{
		BackupLevelEnabled:       BackupLevelEnabled,
		BandwidthLimit:           BandwidthLimit,
		CheckLogical:             CheckLogical,
		CompressedLinkingEnabled: CompressedLinkingEnabled,
		DoubleSync:               DoubleSync,
		EncryptedLinkingEnabled:  EncryptedLinkingEnabled,
		ExternalFilePath:         ExternalFilePath,
		FilesPerSet:              FilesPerSet,
		Force:                    Force,
		LinkNow:                  LinkNow,
		NumberOfConnections:      NumberOfConnections,
		Operations:               Operations,
		RmanChannels:             RmanChannels,
	}
}

// OracleClusterFactory is just a simple function to instantiate the OracleClusterStruct
func OracleClusterFactory(
	CrsClusterHome string,
	CrsClusterName string,
	Description string,
	LogCollectionEnabled *bool,
	Name string,
	RemoteListener string,
	Version string,
) OracleClusterStruct {
	return OracleClusterStruct{
		CrsClusterHome:       CrsClusterHome,
		CrsClusterName:       CrsClusterName,
		Description:          Description,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		RemoteListener:       RemoteListener,
		Version:              Version,
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
		Cluster:              Cluster,
		LogCollectionEnabled: LogCollectionEnabled,
		Node:                 Node,
		PrimaryUser:          PrimaryUser,
	}
}

// OracleClusterNodeFactory is just a simple function to instantiate the OracleClusterNodeStruct
func OracleClusterNodeFactory(
	Name string,
) OracleClusterNodeStruct {
	return OracleClusterNodeStruct{
		Name: Name,
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
		Cluster:        Cluster,
		HostParameters: HostParameters,
		Name:           Name,
		VirtualIPs:     VirtualIPs,
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
		Container:       Container,
		EnvironmentUser: EnvironmentUser,
		Operations:      Operations,
		Repository:      Repository,
	}
}

// OracleDBConfigConnectivityFactory is just a simple function to instantiate the OracleDBConfigConnectivityStruct
func OracleDBConfigConnectivityFactory(
	Password string,
	Username string,
) OracleDBConfigConnectivityStruct {
	return OracleDBConfigConnectivityStruct{
		Password: Password,
		Username: Username,
	}
}

// OracleDBContainerRuntimeFactory is just a simple function to instantiate the OracleDBContainerRuntimeStruct
func OracleDBContainerRuntimeFactory(
	CrossPlatformEligible *bool,
	CrossPlatformScriptUploaded *bool,
	LiveSourceEligible *bool,
) OracleDBContainerRuntimeStruct {
	return OracleDBContainerRuntimeStruct{
		CrossPlatformEligible:       CrossPlatformEligible,
		CrossPlatformScriptUploaded: CrossPlatformScriptUploaded,
		LiveSourceEligible:          LiveSourceEligible,
	}
}

// OracleDatabaseContainerFactory is just a simple function to instantiate the OracleDatabaseContainerStruct
func OracleDatabaseContainerFactory(
	CrossPlatformReady *bool,
	DatabaseFraction *bool,
	Description string,
	DiagnoseNoLoggingFaults *bool,
	Group string,
	LiveSource *bool,
	Name string,
	PerformanceMode string,
	PhysicalStandby *bool,
	PreProvisioningEnabled *bool,
	SourcingPolicy *OracleSourcingPolicyStruct,
) OracleDatabaseContainerStruct {
	return OracleDatabaseContainerStruct{
		CrossPlatformReady:      CrossPlatformReady,
		DatabaseFraction:        DatabaseFraction,
		Description:             Description,
		DiagnoseNoLoggingFaults: DiagnoseNoLoggingFaults,
		Group:                   Group,
		LiveSource:              LiveSource,
		Name:                    Name,
		PerformanceMode:         PerformanceMode,
		PhysicalStandby:         PhysicalStandby,
		PreProvisioningEnabled:  PreProvisioningEnabled,
		SourcingPolicy:          SourcingPolicy,
	}
}

// OracleDatabaseCreationParametersFactory is just a simple function to instantiate the OracleDatabaseCreationParametersStruct
func OracleDatabaseCreationParametersFactory(
	CharacterSet string,
	ForceLogging *bool,
	GrantSelectAnyDictionary *bool,
	MaxDataFiles *int,
	MaxInstances *int,
	MaxLogFiles *int,
	MaxLogHistory *int,
	NationalCharacterSet string,
	RedoLogs []*OracleRedoLogFileSpecificationStruct,
	SysDatafile *OracleSystemDatafileSpecificationStruct,
	SysauxDatafile *OracleSysauxDatafileSpecificationStruct,
	TempTablespace *OracleTempfileSpecificationStruct,
	TimezoneFileVersion *int,
	UndoTablespace *OracleUndoDatafileSpecificationStruct,
) OracleDatabaseCreationParametersStruct {
	return OracleDatabaseCreationParametersStruct{
		CharacterSet:             CharacterSet,
		ForceLogging:             ForceLogging,
		GrantSelectAnyDictionary: GrantSelectAnyDictionary,
		MaxDataFiles:             MaxDataFiles,
		MaxInstances:             MaxInstances,
		MaxLogFiles:              MaxLogFiles,
		MaxLogHistory:            MaxLogHistory,
		NationalCharacterSet:     NationalCharacterSet,
		RedoLogs:                 RedoLogs,
		SysDatafile:              SysDatafile,
		SysauxDatafile:           SysauxDatafile,
		TempTablespace:           TempTablespace,
		TimezoneFileVersion:      TimezoneFileVersion,
		UndoTablespace:           UndoTablespace,
	}
}

// OracleDisableParametersFactory is just a simple function to instantiate the OracleDisableParametersStruct
func OracleDisableParametersFactory(
	AttemptCleanup *bool,
) OracleDisableParametersStruct {
	return OracleDisableParametersStruct{
		AttemptCleanup: AttemptCleanup,
	}
}

// OracleEnableParametersFactory is just a simple function to instantiate the OracleEnableParametersStruct
func OracleEnableParametersFactory(
	AttemptStart *bool,
) OracleEnableParametersStruct {
	return OracleEnableParametersStruct{
		AttemptStart: AttemptStart,
	}
}

// OracleExportParametersFactory is just a simple function to instantiate the OracleExportParametersStruct
func OracleExportParametersFactory(
	ConfigParams map[string]string,
	DspOptions *DSPOptionsStruct,
	FileMappingRules string,
	FileParallelism *int,
	OpenDatabase *bool,
	RecoverDatabase *bool,
) OracleExportParametersStruct {
	return OracleExportParametersStruct{
		ConfigParams:     ConfigParams,
		DspOptions:       DspOptions,
		FileMappingRules: FileMappingRules,
		FileParallelism:  FileParallelism,
		OpenDatabase:     OpenDatabase,
		RecoverDatabase:  RecoverDatabase,
	}
}

// OracleInstallFactory is just a simple function to instantiate the OracleInstallStruct
func OracleInstallFactory(
	AppliedPatches []*int,
	Bits *int,
	Environment string,
	InstallationHome string,
	LinkingEnabled *bool,
	Name string,
	OracleBase string,
	ProvisioningEnabled *bool,
	Staging *bool,
	Version string,
) OracleInstallStruct {
	return OracleInstallStruct{
		AppliedPatches:      AppliedPatches,
		Bits:                Bits,
		Environment:         Environment,
		InstallationHome:    InstallationHome,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		OracleBase:          OracleBase,
		ProvisioningEnabled: ProvisioningEnabled,
		Staging:             Staging,
		Version:             Version,
	}
}

// OracleInstanceFactory is just a simple function to instantiate the OracleInstanceStruct
func OracleInstanceFactory(
	InstanceName string,
	InstanceNumber float64,
) OracleInstanceStruct {
	return OracleInstanceStruct{
		InstanceName:   InstanceName,
		InstanceNumber: InstanceNumber,
	}
}

// OracleLinkDataFactory is just a simple function to instantiate the OracleLinkDataStruct
func OracleLinkDataFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	DiagnoseNoLoggingFaults *bool,
	DoubleSync *bool,
	EncryptedLinkingEnabled *bool,
	ExternalFilePath string,
	FilesPerSet *int,
	LinkNow *bool,
	NonSysCredentials *PasswordCredentialStruct,
	NonSysUser string,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	PreProvisioningEnabled *bool,
	RmanChannels *int,
	SkipSpaceCheck *bool,
	SourcingPolicy *OracleSourcingPolicyStruct,
) OracleLinkDataStruct {
	return OracleLinkDataStruct{
		BackupLevelEnabled:       BackupLevelEnabled,
		BandwidthLimit:           BandwidthLimit,
		CheckLogical:             CheckLogical,
		CompressedLinkingEnabled: CompressedLinkingEnabled,
		DiagnoseNoLoggingFaults:  DiagnoseNoLoggingFaults,
		DoubleSync:               DoubleSync,
		EncryptedLinkingEnabled:  EncryptedLinkingEnabled,
		ExternalFilePath:         ExternalFilePath,
		FilesPerSet:              FilesPerSet,
		LinkNow:                  LinkNow,
		NonSysCredentials:        NonSysCredentials,
		NonSysUser:               NonSysUser,
		NumberOfConnections:      NumberOfConnections,
		Operations:               Operations,
		PreProvisioningEnabled:   PreProvisioningEnabled,
		RmanChannels:             RmanChannels,
		SkipSpaceCheck:           SkipSpaceCheck,
		SourcingPolicy:           SourcingPolicy,
	}
}

// OracleLinkedSourceFactory is just a simple function to instantiate the OracleLinkedSourceStruct
func OracleLinkedSourceFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	Config string,
	EncryptedLinkingEnabled *bool,
	ExternalFilePath string,
	FilesPerSet *int,
	LogCollectionEnabled *bool,
	Name string,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	RmanChannels *int,
) OracleLinkedSourceStruct {
	return OracleLinkedSourceStruct{
		BackupLevelEnabled:       BackupLevelEnabled,
		BandwidthLimit:           BandwidthLimit,
		CheckLogical:             CheckLogical,
		CompressedLinkingEnabled: CompressedLinkingEnabled,
		Config:                   Config,
		EncryptedLinkingEnabled:  EncryptedLinkingEnabled,
		ExternalFilePath:         ExternalFilePath,
		FilesPerSet:              FilesPerSet,
		LogCollectionEnabled:     LogCollectionEnabled,
		Name:                     Name,
		NumberOfConnections:      NumberOfConnections,
		Operations:               Operations,
		RmanChannels:             RmanChannels,
	}
}

// OracleLiveSourceFactory is just a simple function to instantiate the OracleLiveSourceStruct
func OracleLiveSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	CustomEnvVars []OracleCustomEnvVar,
	DataAgeWarningThreshold *int,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
) OracleLiveSourceStruct {
	return OracleLiveSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		CustomEnvVars:                   CustomEnvVars,
		DataAgeWarningThreshold:         DataAgeWarningThreshold,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
	}
}

// OracleMultitenantProvisionParametersFactory is just a simple function to instantiate the OracleMultitenantProvisionParametersStruct
func OracleMultitenantProvisionParametersFactory(
	Credential Credential,
	Masked *bool,
	MaskingJob string,
	Username string,
	VirtualCdb *OracleVirtualCdbProvisionParametersStruct,
) OracleMultitenantProvisionParametersStruct {
	return OracleMultitenantProvisionParametersStruct{
		Credential: Credential,
		Masked:     Masked,
		MaskingJob: MaskingJob,
		Username:   Username,
		VirtualCdb: VirtualCdb,
	}
}

// OracleNodeListenerFactory is just a simple function to instantiate the OracleNodeListenerStruct
func OracleNodeListenerFactory(
	Environment string,
	Host string,
	Name string,
	ProtocolAddresses []string,
) OracleNodeListenerStruct {
	return OracleNodeListenerStruct{
		Environment:       Environment,
		Host:              Host,
		Name:              Name,
		ProtocolAddresses: ProtocolAddresses,
	}
}

// OraclePDBAttachDataFactory is just a simple function to instantiate the OraclePDBAttachDataStruct
func OraclePDBAttachDataFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	EncryptedLinkingEnabled *bool,
	ExternalFilePath string,
	FilesPerSet *int,
	Force *bool,
	LinkNow *bool,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	RmanChannels *int,
) OraclePDBAttachDataStruct {
	return OraclePDBAttachDataStruct{
		BackupLevelEnabled:       BackupLevelEnabled,
		BandwidthLimit:           BandwidthLimit,
		CheckLogical:             CheckLogical,
		CompressedLinkingEnabled: CompressedLinkingEnabled,
		EncryptedLinkingEnabled:  EncryptedLinkingEnabled,
		ExternalFilePath:         ExternalFilePath,
		FilesPerSet:              FilesPerSet,
		Force:                    Force,
		LinkNow:                  LinkNow,
		NumberOfConnections:      NumberOfConnections,
		Operations:               Operations,
		RmanChannels:             RmanChannels,
	}
}

// OraclePDBConfigFactory is just a simple function to instantiate the OraclePDBConfigStruct
func OraclePDBConfigFactory(
	CdbConfig string,
	DatabaseName string,
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Repository string,
	Services []*OracleServiceStruct,
) OraclePDBConfigStruct {
	return OraclePDBConfigStruct{
		CdbConfig:       CdbConfig,
		DatabaseName:    DatabaseName,
		EnvironmentUser: EnvironmentUser,
		LinkingEnabled:  LinkingEnabled,
		Name:            Name,
		Repository:      Repository,
		Services:        Services,
	}
}

// OraclePDBLinkDataFactory is just a simple function to instantiate the OraclePDBLinkDataStruct
func OraclePDBLinkDataFactory(
	BackupLevelEnabled *bool,
	BandwidthLimit *int,
	CheckLogical *bool,
	CompressedLinkingEnabled *bool,
	DiagnoseNoLoggingFaults *bool,
	DoubleSync *bool,
	EncryptedLinkingEnabled *bool,
	ExternalFilePath string,
	FilesPerSet *int,
	LinkNow *bool,
	NumberOfConnections *int,
	Operations *LinkedSourceOperationsStruct,
	PreProvisioningEnabled *bool,
	RmanChannels *int,
	SkipSpaceCheck *bool,
	SourcingPolicy *OracleSourcingPolicyStruct,
) OraclePDBLinkDataStruct {
	return OraclePDBLinkDataStruct{
		BackupLevelEnabled:       BackupLevelEnabled,
		BandwidthLimit:           BandwidthLimit,
		CheckLogical:             CheckLogical,
		CompressedLinkingEnabled: CompressedLinkingEnabled,
		DiagnoseNoLoggingFaults:  DiagnoseNoLoggingFaults,
		DoubleSync:               DoubleSync,
		EncryptedLinkingEnabled:  EncryptedLinkingEnabled,
		ExternalFilePath:         ExternalFilePath,
		FilesPerSet:              FilesPerSet,
		LinkNow:                  LinkNow,
		NumberOfConnections:      NumberOfConnections,
		Operations:               Operations,
		PreProvisioningEnabled:   PreProvisioningEnabled,
		RmanChannels:             RmanChannels,
		SkipSpaceCheck:           SkipSpaceCheck,
		SourcingPolicy:           SourcingPolicy,
	}
}

// OracleProvisionParametersFactory is just a simple function to instantiate the OracleProvisionParametersStruct
func OracleProvisionParametersFactory(
	Credential Credential,
	Masked *bool,
	MaskingJob string,
	NewDBID *bool,
	OpenResetlogs *bool,
	PhysicalStandby *bool,
	Username string,
) OracleProvisionParametersStruct {
	return OracleProvisionParametersStruct{
		Credential:      Credential,
		Masked:          Masked,
		MaskingJob:      MaskingJob,
		NewDBID:         NewDBID,
		OpenResetlogs:   OpenResetlogs,
		PhysicalStandby: PhysicalStandby,
		Username:        Username,
	}
}

// OracleRACConfigFactory is just a simple function to instantiate the OracleRACConfigStruct
func OracleRACConfigFactory(
	CrsDatabaseName string,
	DatabaseName string,
	EnvironmentUser string,
	Instances []*OracleRACInstanceStruct,
	LinkingEnabled *bool,
	Name string,
	NonSysCredentials string,
	NonSysUser string,
	Repository string,
	Services []*OracleServiceStruct,
	UniqueName string,
) OracleRACConfigStruct {
	return OracleRACConfigStruct{
		CrsDatabaseName:   CrsDatabaseName,
		DatabaseName:      DatabaseName,
		EnvironmentUser:   EnvironmentUser,
		Instances:         Instances,
		LinkingEnabled:    LinkingEnabled,
		Name:              Name,
		NonSysCredentials: NonSysCredentials,
		NonSysUser:        NonSysUser,
		Repository:        Repository,
		Services:          Services,
		UniqueName:        UniqueName,
	}
}

// OracleRACInstanceFactory is just a simple function to instantiate the OracleRACInstanceStruct
func OracleRACInstanceFactory(
	InstanceName string,
	InstanceNumber float64,
	Node string,
) OracleRACInstanceStruct {
	return OracleRACInstanceStruct{
		InstanceName:   InstanceName,
		InstanceNumber: InstanceNumber,
		Node:           Node,
	}
}

// OracleRedoLogFileSpecificationFactory is just a simple function to instantiate the OracleRedoLogFileSpecificationStruct
func OracleRedoLogFileSpecificationFactory(
	Filename string,
	Size *int,
) OracleRedoLogFileSpecificationStruct {
	return OracleRedoLogFileSpecificationStruct{
		Filename: Filename,
		Size:     Size,
	}
}

// OracleSIConfigFactory is just a simple function to instantiate the OracleSIConfigStruct
func OracleSIConfigFactory(
	DatabaseName string,
	EnvironmentUser string,
	Instance *OracleInstanceStruct,
	LinkingEnabled *bool,
	Name string,
	NonSysCredentials string,
	NonSysUser string,
	Repository string,
	Services []*OracleServiceStruct,
	UniqueName string,
) OracleSIConfigStruct {
	return OracleSIConfigStruct{
		DatabaseName:      DatabaseName,
		EnvironmentUser:   EnvironmentUser,
		Instance:          Instance,
		LinkingEnabled:    LinkingEnabled,
		Name:              Name,
		NonSysCredentials: NonSysCredentials,
		NonSysUser:        NonSysUser,
		Repository:        Repository,
		Services:          Services,
		UniqueName:        UniqueName,
	}
}

// OracleScanListenerFactory is just a simple function to instantiate the OracleScanListenerStruct
func OracleScanListenerFactory(
	Environment string,
	Name string,
	ProtocolAddresses []string,
) OracleScanListenerStruct {
	return OracleScanListenerStruct{
		Environment:       Environment,
		Name:              Name,
		ProtocolAddresses: ProtocolAddresses,
	}
}

// OracleServiceFactory is just a simple function to instantiate the OracleServiceStruct
func OracleServiceFactory(
	JdbcConnectionString string,
) OracleServiceStruct {
	return OracleServiceStruct{
		JdbcConnectionString: JdbcConnectionString,
	}
}

// OracleSnapshotFactory is just a simple function to instantiate the OracleSnapshotStruct
func OracleSnapshotFactory(
	FromPhysicalStandbyVdb *bool,
	Name string,
	RedoLogSizeInBytes float64,
) OracleSnapshotStruct {
	return OracleSnapshotStruct{
		FromPhysicalStandbyVdb: FromPhysicalStandbyVdb,
		Name:                   Name,
		RedoLogSizeInBytes:     RedoLogSizeInBytes,
	}
}

// OracleSourcingPolicyFactory is just a simple function to instantiate the OracleSourcingPolicyStruct
func OracleSourcingPolicyFactory(
	LogsyncEnabled *bool,
	LogsyncInterval *int,
	LogsyncMode string,
) OracleSourcingPolicyStruct {
	return OracleSourcingPolicyStruct{
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
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	NodeListeners []string,
) OracleStagingSourceStruct {
	return OracleStagingSourceStruct{
		Config:               Config,
		ConfigParams:         ConfigParams,
		ConfigTemplate:       ConfigTemplate,
		LogCollectionEnabled: LogCollectionEnabled,
		MountBase:            MountBase,
		Name:                 Name,
		NodeListeners:        NodeListeners,
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
		ConfigParams:    ConfigParams,
		ConfigTemplate:  ConfigTemplate,
		EnvironmentUser: EnvironmentUser,
		MountBase:       MountBase,
		Repository:      Repository,
	}
}

// OracleStopParametersFactory is just a simple function to instantiate the OracleStopParametersStruct
func OracleStopParametersFactory(
	Abort *bool,
) OracleStopParametersStruct {
	return OracleStopParametersStruct{
		Abort: Abort,
	}
}

// OracleSyncParametersFactory is just a simple function to instantiate the OracleSyncParametersStruct
func OracleSyncParametersFactory(
	DoNotResume *bool,
	DoubleSync *bool,
	ForceFullBackup *bool,
	SkipSpaceCheck *bool,
) OracleSyncParametersStruct {
	return OracleSyncParametersStruct{
		DoNotResume:     DoNotResume,
		DoubleSync:      DoubleSync,
		ForceFullBackup: ForceFullBackup,
		SkipSpaceCheck:  SkipSpaceCheck,
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
		AutoExtend:          AutoExtend,
		AutoExtendIncrement: AutoExtendIncrement,
		Filename:            Filename,
		MaxSize:             MaxSize,
		Size:                Size,
	}
}

// OracleTimeflowFactory is just a simple function to instantiate the OracleTimeflowStruct
func OracleTimeflowFactory(
	Name string,
) OracleTimeflowStruct {
	return OracleTimeflowStruct{
		Name: Name,
	}
}

// OracleTimeflowPointFactory is just a simple function to instantiate the OracleTimeflowPointStruct
func OracleTimeflowPointFactory(
	Location string,
	Timestamp string,
) OracleTimeflowPointStruct {
	return OracleTimeflowPointStruct{
		Location:  Location,
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
		AutoExtend:          AutoExtend,
		AutoExtendIncrement: AutoExtendIncrement,
		Filename:            Filename,
		MaxSize:             MaxSize,
		Size:                Size,
	}
}

// OracleVirtualCdbSourceFactory is just a simple function to instantiate the OracleVirtualCdbSourceStruct
func OracleVirtualCdbSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	CustomEnvVars []OracleCustomEnvVar,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
) OracleVirtualCdbSourceStruct {
	return OracleVirtualCdbSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		CustomEnvVars:                   CustomEnvVars,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
	}
}

// OracleVirtualPdbSourceFactory is just a simple function to instantiate the OracleVirtualPdbSourceStruct
func OracleVirtualPdbSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	CustomEnvVars []OracleCustomEnvVar,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
) OracleVirtualPdbSourceStruct {
	return OracleVirtualPdbSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		CustomEnvVars:                   CustomEnvVars,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
	}
}

// OracleVirtualSourceFactory is just a simple function to instantiate the OracleVirtualSourceStruct
func OracleVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	CustomEnvVars []OracleCustomEnvVar,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
) OracleVirtualSourceStruct {
	return OracleVirtualSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		CustomEnvVars:                   CustomEnvVars,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
	}
}

// OracleWarehouseSourceFactory is just a simple function to instantiate the OracleWarehouseSourceStruct
func OracleWarehouseSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	ArchivelogMode *bool,
	Config string,
	ConfigParams map[string]string,
	ConfigTemplate string,
	CustomEnvVars []OracleCustomEnvVar,
	FileMappingRules string,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	NodeListeners []string,
	Operations *VirtualSourceOperationsStruct,
	RedoLogGroups *int,
	RedoLogSizeInMB *int,
) OracleWarehouseSourceStruct {
	return OracleWarehouseSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		ArchivelogMode:                  ArchivelogMode,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		ConfigTemplate:                  ConfigTemplate,
		CustomEnvVars:                   CustomEnvVars,
		FileMappingRules:                FileMappingRules,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		NodeListeners:                   NodeListeners,
		Operations:                      Operations,
		RedoLogGroups:                   RedoLogGroups,
		RedoLogSizeInMB:                 RedoLogSizeInMB,
	}
}

// PasswordPolicyFactory is just a simple function to instantiate the PasswordPolicyStruct
func PasswordPolicyFactory(
	Digit *bool,
	DisallowUsernameAsPassword *bool,
	LowercaseLetter *bool,
	MinLength *int,
	Name string,
	ReuseDisallowLimit *int,
	Symbol *bool,
	UppercaseLetter *bool,
) PasswordPolicyStruct {
	return PasswordPolicyStruct{
		Digit:                      Digit,
		DisallowUsernameAsPassword: DisallowUsernameAsPassword,
		LowercaseLetter:            LowercaseLetter,
		MinLength:                  MinLength,
		Name:                       Name,
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
		AxisName:     AxisName,
		DescendantOf: DescendantOf,
	}
}

// PermissionFactory is just a simple function to instantiate the PermissionStruct
func PermissionFactory(
	Name string,
) PermissionStruct {
	return PermissionStruct{
		Name: Name,
	}
}

// PgSQLAttachDataFactory is just a simple function to instantiate the PgSQLAttachDataStruct
func PgSQLAttachDataFactory(
	ConnectionDatabase string,
	DbCredentials *PasswordCredentialStruct,
	ExternalFilePath string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
) PgSQLAttachDataStruct {
	return PgSQLAttachDataStruct{
		ConnectionDatabase: ConnectionDatabase,
		DbCredentials:      DbCredentials,
		ExternalFilePath:   ExternalFilePath,
		Operations:         Operations,
		PptHostUser:        PptHostUser,
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
	EnvironmentUser string,
	LinkingEnabled *bool,
	Name string,
	Port *int,
	Repository string,
) PgSQLDBClusterConfigStruct {
	return PgSQLDBClusterConfigStruct{
		ClusterDataDirectory: ClusterDataDirectory,
		ConnectionDatabase:   ConnectionDatabase,
		Credentials:          Credentials,
		EnvironmentUser:      EnvironmentUser,
		LinkingEnabled:       LinkingEnabled,
		Name:                 Name,
		Port:                 Port,
		Repository:           Repository,
	}
}

// PgSQLDBClusterConfigConnectivityFactory is just a simple function to instantiate the PgSQLDBClusterConfigConnectivityStruct
func PgSQLDBClusterConfigConnectivityFactory(
	ConnectionDatabase string,
) PgSQLDBClusterConfigConnectivityStruct {
	return PgSQLDBClusterConfigConnectivityStruct{
		ConnectionDatabase: ConnectionDatabase,
	}
}

// PgSQLDBConfigFactory is just a simple function to instantiate the PgSQLDBConfigStruct
func PgSQLDBConfigFactory(
	DatabaseName string,
) PgSQLDBConfigStruct {
	return PgSQLDBConfigStruct{
		DatabaseName: DatabaseName,
	}
}

// PgSQLDatabaseContainerFactory is just a simple function to instantiate the PgSQLDatabaseContainerStruct
func PgSQLDatabaseContainerFactory(
	Description string,
	Group string,
	Name string,
	SourcingPolicy *SourcingPolicyStruct,
) PgSQLDatabaseContainerStruct {
	return PgSQLDatabaseContainerStruct{
		Description:    Description,
		Group:          Group,
		Name:           Name,
		SourcingPolicy: SourcingPolicy,
	}
}

// PgSQLExportParametersFactory is just a simple function to instantiate the PgSQLExportParametersStruct
func PgSQLExportParametersFactory(
	ConfigParams map[string]string,
	FileMappingRules string,
	HbaEntries []*PgSQLHBAEntryStruct,
	IdentEntries []*PgSQLIdentEntryStruct,
	RecoverDatabase *bool,
) PgSQLExportParametersStruct {
	return PgSQLExportParametersStruct{
		ConfigParams:     ConfigParams,
		FileMappingRules: FileMappingRules,
		HbaEntries:       HbaEntries,
		IdentEntries:     IdentEntries,
		RecoverDatabase:  RecoverDatabase,
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
		DatabaseUsername: DatabaseUsername,
		MapName:          MapName,
		SystemUsername:   SystemUsername,
	}
}

// PgSQLInstallFactory is just a simple function to instantiate the PgSQLInstallStruct
func PgSQLInstallFactory(
	Environment string,
	InstallationPath string,
	LinkingEnabled *bool,
	Name string,
	ProvisioningEnabled *bool,
	Staging *bool,
	Version string,
) PgSQLInstallStruct {
	return PgSQLInstallStruct{
		Environment:         Environment,
		InstallationPath:    InstallationPath,
		LinkingEnabled:      LinkingEnabled,
		Name:                Name,
		ProvisioningEnabled: ProvisioningEnabled,
		Staging:             Staging,
		Version:             Version,
	}
}

// PgSQLLinkDataFactory is just a simple function to instantiate the PgSQLLinkDataStruct
func PgSQLLinkDataFactory(
	ConnectionDatabase string,
	DbCredentials *PasswordCredentialStruct,
	ExternalFilePath string,
	Operations *LinkedSourceOperationsStruct,
	PptHostUser string,
	SourcingPolicy *SourcingPolicyStruct,
) PgSQLLinkDataStruct {
	return PgSQLLinkDataStruct{
		ConnectionDatabase: ConnectionDatabase,
		DbCredentials:      DbCredentials,
		ExternalFilePath:   ExternalFilePath,
		Operations:         Operations,
		PptHostUser:        PptHostUser,
		SourcingPolicy:     SourcingPolicy,
	}
}

// PgSQLLinkedSourceFactory is just a simple function to instantiate the PgSQLLinkedSourceStruct
func PgSQLLinkedSourceFactory(
	Config string,
	ExternalFilePath string,
	LogCollectionEnabled *bool,
	Name string,
	Operations *LinkedSourceOperationsStruct,
) PgSQLLinkedSourceStruct {
	return PgSQLLinkedSourceStruct{
		Config:               Config,
		ExternalFilePath:     ExternalFilePath,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Operations:           Operations,
	}
}

// PgSQLProvisionParametersFactory is just a simple function to instantiate the PgSQLProvisionParametersStruct
func PgSQLProvisionParametersFactory(
	Masked *bool,
	MaskingJob string,
) PgSQLProvisionParametersStruct {
	return PgSQLProvisionParametersStruct{
		Masked:     Masked,
		MaskingJob: MaskingJob,
	}
}

// PgSQLSnapshotFactory is just a simple function to instantiate the PgSQLSnapshotStruct
func PgSQLSnapshotFactory(
	Name string,
) PgSQLSnapshotStruct {
	return PgSQLSnapshotStruct{
		Name: Name,
	}
}

// PgSQLStagingSourceFactory is just a simple function to instantiate the PgSQLStagingSourceStruct
func PgSQLStagingSourceFactory(
	Config string,
	LogCollectionEnabled *bool,
	Name string,
) PgSQLStagingSourceStruct {
	return PgSQLStagingSourceStruct{
		Config:               Config,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
	}
}

// PgSQLSyncParametersFactory is just a simple function to instantiate the PgSQLSyncParametersStruct
func PgSQLSyncParametersFactory(
	RedoBaseBackup *bool,
) PgSQLSyncParametersStruct {
	return PgSQLSyncParametersStruct{
		RedoBaseBackup: RedoBaseBackup,
	}
}

// PgSQLTimeflowFactory is just a simple function to instantiate the PgSQLTimeflowStruct
func PgSQLTimeflowFactory(
	Name string,
) PgSQLTimeflowStruct {
	return PgSQLTimeflowStruct{
		Name: Name,
	}
}

// PgSQLTimeflowPointFactory is just a simple function to instantiate the PgSQLTimeflowPointStruct
func PgSQLTimeflowPointFactory(
	Location string,
	Timestamp string,
) PgSQLTimeflowPointStruct {
	return PgSQLTimeflowPointStruct{
		Location:  Location,
		Timestamp: Timestamp,
	}
}

// PgSQLVirtualSourceFactory is just a simple function to instantiate the PgSQLVirtualSourceStruct
func PgSQLVirtualSourceFactory(
	AllowAutoVDBRestartOnHostReboot *bool,
	Config string,
	ConfigParams map[string]string,
	FileMappingRules string,
	HbaEntries []*PgSQLHBAEntryStruct,
	IdentEntries []*PgSQLIdentEntryStruct,
	LogCollectionEnabled *bool,
	MountBase string,
	Name string,
	Operations *VirtualSourceOperationsStruct,
) PgSQLVirtualSourceStruct {
	return PgSQLVirtualSourceStruct{
		AllowAutoVDBRestartOnHostReboot: AllowAutoVDBRestartOnHostReboot,
		Config:                          Config,
		ConfigParams:                    ConfigParams,
		FileMappingRules:                FileMappingRules,
		HbaEntries:                      HbaEntries,
		IdentEntries:                    IdentEntries,
		LogCollectionEnabled:            LogCollectionEnabled,
		MountBase:                       MountBase,
		Name:                            Name,
		Operations:                      Operations,
	}
}

// PopulateCompatibilityParametersFactory is just a simple function to instantiate the PopulateCompatibilityParametersStruct
func PopulateCompatibilityParametersFactory(
	Environment string,
) PopulateCompatibilityParametersStruct {
	return PopulateCompatibilityParametersStruct{
		Environment: Environment,
	}
}

// ProvisionCompatibilityParametersFactory is just a simple function to instantiate the ProvisionCompatibilityParametersStruct
func ProvisionCompatibilityParametersFactory(
	Environment string,
) ProvisionCompatibilityParametersStruct {
	return ProvisionCompatibilityParametersStruct{
		Environment: Environment,
	}
}

// PublicSystemInfoFactory is just a simple function to instantiate the PublicSystemInfoStruct
func PublicSystemInfoFactory(
	EngineQualifier string,
	SupportContacts []*SupportContactStruct,
	SupportURL string,
	VendorAddress []string,
	VendorEmail string,
	VendorName string,
	VendorPhoneNumber string,
	VendorURL string,
) PublicSystemInfoStruct {
	return PublicSystemInfoStruct{
		EngineQualifier:   EngineQualifier,
		SupportContacts:   SupportContacts,
		SupportURL:        SupportURL,
		VendorAddress:     VendorAddress,
		VendorEmail:       VendorEmail,
		VendorName:        VendorName,
		VendorPhoneNumber: VendorPhoneNumber,
		VendorURL:         VendorURL,
	}
}

// PurgeLogsParametersFactory is just a simple function to instantiate the PurgeLogsParametersStruct
func PurgeLogsParametersFactory(
	DeleteSnapshotLogs *bool,
	DryRun *bool,
) PurgeLogsParametersStruct {
	return PurgeLogsParametersStruct{
		DeleteSnapshotLogs: DeleteSnapshotLogs,
		DryRun:             DryRun,
	}
}

// QuotaPolicyFactory is just a simple function to instantiate the QuotaPolicyStruct
func QuotaPolicyFactory(
	Customized *bool,
	Name string,
	Size float64,
) QuotaPolicyStruct {
	return QuotaPolicyStruct{
		Customized: Customized,
		Name:       Name,
		Size:       Size,
	}
}

// RefreshPolicyFactory is just a simple function to instantiate the RefreshPolicyStruct
func RefreshPolicyFactory(
	Customized *bool,
	Name string,
	ProvisionSource string,
	ScheduleList []*ScheduleStruct,
	Timezone *TimeZoneStruct,
) RefreshPolicyStruct {
	return RefreshPolicyStruct{
		Customized:      Customized,
		Name:            Name,
		ProvisionSource: ProvisionSource,
		ScheduleList:    ScheduleList,
		Timezone:        Timezone,
	}
}

// RegistrationStatusFactory is just a simple function to instantiate the RegistrationStatusStruct
func RegistrationStatusFactory(
	Name string,
) RegistrationStatusStruct {
	return RegistrationStatusStruct{
		Name: Name,
	}
}

// RemoteDelphixEngineInfoFactory is just a simple function to instantiate the RemoteDelphixEngineInfoStruct
func RemoteDelphixEngineInfoFactory(
	Address string,
	Credential *PasswordCredentialStruct,
	Principal string,
) RemoteDelphixEngineInfoStruct {
	return RemoteDelphixEngineInfoStruct{
		Address:    Address,
		Credential: Credential,
		Principal:  Principal,
	}
}

// ReplicationListFactory is just a simple function to instantiate the ReplicationListStruct
func ReplicationListFactory(
	Name string,
	Objects []string,
) ReplicationListStruct {
	return ReplicationListStruct{
		Name:    Name,
		Objects: Objects,
	}
}

// ReplicationSecureListFactory is just a simple function to instantiate the ReplicationSecureListStruct
func ReplicationSecureListFactory(
	Containers []string,
	Name string,
) ReplicationSecureListStruct {
	return ReplicationSecureListStruct{
		Containers: Containers,
		Name:       Name,
	}
}

// ReplicationSourceStateFactory is just a simple function to instantiate the ReplicationSourceStateStruct
func ReplicationSourceStateFactory(
	Name string,
) ReplicationSourceStateStruct {
	return ReplicationSourceStateStruct{
		Name: Name,
	}
}

// ReplicationSpecFactory is just a simple function to instantiate the ReplicationSpecStruct
func ReplicationSpecFactory(
	AutomaticReplication *bool,
	BandwidthLimit *int,
	Description string,
	Encrypted *bool,
	Name string,
	NumberOfConnections *int,
	ObjectSpecification ReplicationObjectSpecification,
	Schedule string,
	TargetCredential *PasswordCredentialStruct,
	TargetHost string,
	TargetPort *int,
	TargetPrincipal string,
	UseSystemSocksSetting *bool,
) ReplicationSpecStruct {
	return ReplicationSpecStruct{
		AutomaticReplication:  AutomaticReplication,
		BandwidthLimit:        BandwidthLimit,
		Description:           Description,
		Encrypted:             Encrypted,
		Name:                  Name,
		NumberOfConnections:   NumberOfConnections,
		ObjectSpecification:   ObjectSpecification,
		Schedule:              Schedule,
		TargetCredential:      TargetCredential,
		TargetHost:            TargetHost,
		TargetPort:            TargetPort,
		TargetPrincipal:       TargetPrincipal,
		UseSystemSocksSetting: UseSystemSocksSetting,
	}
}

// ReplicationTargetStateFactory is just a simple function to instantiate the ReplicationTargetStateStruct
func ReplicationTargetStateFactory(
	Name string,
) ReplicationTargetStateStruct {
	return ReplicationTargetStateStruct{
		Name: Name,
	}
}

// ResolveOrIgnoreSelectedFaultsParametersFactory is just a simple function to instantiate the ResolveOrIgnoreSelectedFaultsParametersStruct
func ResolveOrIgnoreSelectedFaultsParametersFactory(
	Ignore *bool,
) ResolveOrIgnoreSelectedFaultsParametersStruct {
	return ResolveOrIgnoreSelectedFaultsParametersStruct{
		Ignore: Ignore,
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
	LogDuration *int,
	LogUnit string,
	Name string,
	NumOfDaily *int,
	NumOfMonthly *int,
	NumOfWeekly *int,
	NumOfYearly *int,
) RetentionPolicyStruct {
	return RetentionPolicyStruct{
		Customized:   Customized,
		DataDuration: DataDuration,
		DataUnit:     DataUnit,
		DayOfMonth:   DayOfMonth,
		DayOfWeek:    DayOfWeek,
		DayOfYear:    DayOfYear,
		LogDuration:  LogDuration,
		LogUnit:      LogUnit,
		Name:         Name,
		NumOfDaily:   NumOfDaily,
		NumOfMonthly: NumOfMonthly,
		NumOfWeekly:  NumOfWeekly,
		NumOfYearly:  NumOfYearly,
	}
}

// RoleFactory is just a simple function to instantiate the RoleStruct
func RoleFactory(
	Name string,
	Permissions []string,
) RoleStruct {
	return RoleStruct{
		Name:        Name,
		Permissions: Permissions,
	}
}

// RsaKeyPairFactory is just a simple function to instantiate the RsaKeyPairStruct
func RsaKeyPairFactory(
	KeySize *int,
	SignatureAlgorithm string,
) RsaKeyPairStruct {
	return RsaKeyPairStruct{
		KeySize:            KeySize,
		SignatureAlgorithm: SignatureAlgorithm,
	}
}

// RunBashOnSourceOperationFactory is just a simple function to instantiate the RunBashOnSourceOperationStruct
func RunBashOnSourceOperationFactory(
	Command string,
	Name string,
) RunBashOnSourceOperationStruct {
	return RunBashOnSourceOperationStruct{
		Command: Command,
		Name:    Name,
	}
}

// RunCommandOnSourceOperationFactory is just a simple function to instantiate the RunCommandOnSourceOperationStruct
func RunCommandOnSourceOperationFactory(
	Command string,
	Name string,
) RunCommandOnSourceOperationStruct {
	return RunCommandOnSourceOperationStruct{
		Command: Command,
		Name:    Name,
	}
}

// RunExpectOnSourceOperationFactory is just a simple function to instantiate the RunExpectOnSourceOperationStruct
func RunExpectOnSourceOperationFactory(
	Command string,
	Name string,
) RunExpectOnSourceOperationStruct {
	return RunExpectOnSourceOperationStruct{
		Command: Command,
		Name:    Name,
	}
}

// RunPowerShellOnSourceOperationFactory is just a simple function to instantiate the RunPowerShellOnSourceOperationStruct
func RunPowerShellOnSourceOperationFactory(
	Command string,
	Name string,
) RunPowerShellOnSourceOperationStruct {
	return RunPowerShellOnSourceOperationStruct{
		Command: Command,
		Name:    Name,
	}
}

// SMTPConfigFactory is just a simple function to instantiate the SMTPConfigStruct
func SMTPConfigFactory(
	AuthenticationEnabled *bool,
	TlsEnabled *bool,
) SMTPConfigStruct {
	return SMTPConfigStruct{
		AuthenticationEnabled: AuthenticationEnabled,
		TlsEnabled:            TlsEnabled,
	}
}

// SNMPConfigFactory is just a simple function to instantiate the SNMPConfigStruct
func SNMPConfigFactory(
	Enabled *bool,
) SNMPConfigStruct {
	return SNMPConfigStruct{
		Enabled: Enabled,
	}
}

// SNMPManagerFactory is just a simple function to instantiate the SNMPManagerStruct
func SNMPManagerFactory(
	Address string,
	CommunityString string,
	Port *int,
	UseInform *bool,
) SNMPManagerStruct {
	return SNMPManagerStruct{
		Address:         Address,
		CommunityString: CommunityString,
		Port:            Port,
		UseInform:       UseInform,
	}
}

// SamlAuthParametersFactory is just a simple function to instantiate the SamlAuthParametersStruct
func SamlAuthParametersFactory(
	EncodeRequest *bool,
) SamlAuthParametersStruct {
	return SamlAuthParametersStruct{
		EncodeRequest: EncodeRequest,
	}
}

// SamlServiceProviderFactory is just a simple function to instantiate the SamlServiceProviderStruct
func SamlServiceProviderFactory(
	DecryptingKey string,
	Name string,
	SigningKey string,
) SamlServiceProviderStruct {
	return SamlServiceProviderStruct{
		DecryptingKey: DecryptingKey,
		Name:          Name,
		SigningKey:    SigningKey,
	}
}

// ScheduleFactory is just a simple function to instantiate the ScheduleStruct
func ScheduleFactory(
	CronString string,
	CutoffTime *int,
) ScheduleStruct {
	return ScheduleStruct{
		CronString: CronString,
		CutoffTime: CutoffTime,
	}
}

// SerializationPointFactory is just a simple function to instantiate the SerializationPointStruct
func SerializationPointFactory(
	Name string,
) SerializationPointStruct {
	return SerializationPointStruct{
		Name: Name,
	}
}

// SeverityFilterFactory is just a simple function to instantiate the SeverityFilterStruct
func SeverityFilterFactory(
	SeverityLevels []string,
) SeverityFilterStruct {
	return SeverityFilterStruct{
		SeverityLevels: SeverityLevels,
	}
}

// SnapshotPolicyFactory is just a simple function to instantiate the SnapshotPolicyStruct
func SnapshotPolicyFactory(
	Customized *bool,
	Name string,
	ScheduleList []*ScheduleStruct,
	Timezone *TimeZoneStruct,
) SnapshotPolicyStruct {
	return SnapshotPolicyStruct{
		Customized:   Customized,
		Name:         Name,
		ScheduleList: ScheduleList,
		Timezone:     Timezone,
	}
}

// SourceDisableParametersFactory is just a simple function to instantiate the SourceDisableParametersStruct
func SourceDisableParametersFactory(
	AttemptCleanup *bool,
) SourceDisableParametersStruct {
	return SourceDisableParametersStruct{
		AttemptCleanup: AttemptCleanup,
	}
}

// SourceEnableParametersFactory is just a simple function to instantiate the SourceEnableParametersStruct
func SourceEnableParametersFactory(
	AttemptStart *bool,
) SourceEnableParametersStruct {
	return SourceEnableParametersStruct{
		AttemptStart: AttemptStart,
	}
}

// SourceRepositoryTemplateFactory is just a simple function to instantiate the SourceRepositoryTemplateStruct
func SourceRepositoryTemplateFactory(
	Container string,
	Name string,
	Repository string,
	Template string,
) SourceRepositoryTemplateStruct {
	return SourceRepositoryTemplateStruct{
		Container:  Container,
		Name:       Name,
		Repository: Repository,
		Template:   Template,
	}
}

// SourcingPolicyFactory is just a simple function to instantiate the SourcingPolicyStruct
func SourcingPolicyFactory(
	LogsyncEnabled *bool,
) SourcingPolicyStruct {
	return SourcingPolicyStruct{
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
	PerformanceMetricsResolution string,
) SplunkHecConfigStruct {
	return SplunkHecConfigStruct{
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
		PerformanceMetricsResolution: PerformanceMetricsResolution,
	}
}

// StagingCompatibilityParametersFactory is just a simple function to instantiate the StagingCompatibilityParametersStruct
func StagingCompatibilityParametersFactory(
	Environment string,
) StagingCompatibilityParametersStruct {
	return StagingCompatibilityParametersStruct{
		Environment: Environment,
	}
}

// StatisticSliceFactory is just a simple function to instantiate the StatisticSliceStruct
func StatisticSliceFactory(
	AxisConstraints []AxisConstraint,
	CollectionAxes []string,
	CollectionInterval *int,
	Name string,
	StatisticType string,
) StatisticSliceStruct {
	return StatisticSliceStruct{
		AxisConstraints:    AxisConstraints,
		CollectionAxes:     CollectionAxes,
		CollectionInterval: CollectionInterval,
		Name:               Name,
		StatisticType:      StatisticType,
	}
}

// StorageDeviceFactory is just a simple function to instantiate the StorageDeviceStruct
func StorageDeviceFactory(
	Name string,
) StorageDeviceStruct {
	return StorageDeviceStruct{
		Name: Name,
	}
}

// StorageTestFactory is just a simple function to instantiate the StorageTestStruct
func StorageTestFactory(
	Name string,
) StorageTestStruct {
	return StorageTestStruct{
		Name: Name,
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
		Devices:                Devices,
		Duration:               Duration,
		InitializeDevices:      InitializeDevices,
		InitializeEntireDevice: InitializeEntireDevice,
		TestRegion:             TestRegion,
		Tests:                  Tests,
	}
}

// StringEqualConstraintFactory is just a simple function to instantiate the StringEqualConstraintStruct
func StringEqualConstraintFactory(
	AxisName string,
	Equals string,
) StringEqualConstraintStruct {
	return StringEqualConstraintStruct{
		AxisName: AxisName,
		Equals:   Equals,
	}
}

// SyncPolicyFactory is just a simple function to instantiate the SyncPolicyStruct
func SyncPolicyFactory(
	Customized *bool,
	Name string,
	ScheduleList []*ScheduleStruct,
	Timezone *TimeZoneStruct,
) SyncPolicyStruct {
	return SyncPolicyStruct{
		Customized:   Customized,
		Name:         Name,
		ScheduleList: ScheduleList,
		Timezone:     Timezone,
	}
}

// SyslogConfigFactory is just a simple function to instantiate the SyslogConfigStruct
func SyslogConfigFactory(
	Enabled *bool,
) SyslogConfigStruct {
	return SyslogConfigStruct{
		Enabled: Enabled,
	}
}

// SystemInfoFactory is just a simple function to instantiate the SystemInfoStruct
func SystemInfoFactory(
	EngineQualifier string,
	SupportContacts []*SupportContactStruct,
	SupportURL string,
	VendorAddress []string,
	VendorEmail string,
	VendorName string,
	VendorPhoneNumber string,
	VendorURL string,
) SystemInfoStruct {
	return SystemInfoStruct{
		EngineQualifier:   EngineQualifier,
		SupportContacts:   SupportContacts,
		SupportURL:        SupportURL,
		VendorAddress:     VendorAddress,
		VendorEmail:       VendorEmail,
		VendorName:        VendorName,
		VendorPhoneNumber: VendorPhoneNumber,
		VendorURL:         VendorURL,
	}
}

// SystemInitializationParametersFactory is just a simple function to instantiate the SystemInitializationParametersStruct
func SystemInitializationParametersFactory(
	DefaultPassword string,
	DefaultUser string,
	Devices []string,
) SystemInitializationParametersStruct {
	return SystemInitializationParametersStruct{
		DefaultPassword: DefaultPassword,
		DefaultUser:     DefaultUser,
		Devices:         Devices,
	}
}

// SystemPackageFactory is just a simple function to instantiate the SystemPackageStruct
func SystemPackageFactory(
	Name string,
) SystemPackageStruct {
	return SystemPackageStruct{
		Name: Name,
	}
}

// SystemVersionFactory is just a simple function to instantiate the SystemVersionStruct
func SystemVersionFactory(
	Name string,
) SystemVersionStruct {
	return SystemVersionStruct{
		Name: Name,
	}
}

// TargetFilterFactory is just a simple function to instantiate the TargetFilterStruct
func TargetFilterFactory(
	Targets []string,
) TargetFilterStruct {
	return TargetFilterStruct{
		Targets: Targets,
	}
}

// TargetOwnerFilterFactory is just a simple function to instantiate the TargetOwnerFilterStruct
func TargetOwnerFilterFactory(
	Owners []string,
) TargetOwnerFilterStruct {
	return TargetOwnerFilterStruct{
		Owners: Owners,
	}
}

// TimeRangeParametersFactory is just a simple function to instantiate the TimeRangeParametersStruct
func TimeRangeParametersFactory(
	EndTime string,
	StartTime string,
) TimeRangeParametersStruct {
	return TimeRangeParametersStruct{
		EndTime:   EndTime,
		StartTime: StartTime,
	}
}

// TimeflowBookmarkFactory is just a simple function to instantiate the TimeflowBookmarkStruct
func TimeflowBookmarkFactory(
	Name string,
) TimeflowBookmarkStruct {
	return TimeflowBookmarkStruct{
		Name: Name,
	}
}

// TimeflowBookmarkCreateParametersFactory is just a simple function to instantiate the TimeflowBookmarkCreateParametersStruct
func TimeflowBookmarkCreateParametersFactory(
	RetentionProof *bool,
	Tag string,
) TimeflowBookmarkCreateParametersStruct {
	return TimeflowBookmarkCreateParametersStruct{
		RetentionProof: RetentionProof,
		Tag:            Tag,
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
		ArchiveDirectory:  ArchiveDirectory,
		DataDirectory:     DataDirectory,
		ExternalDirectory: ExternalDirectory,
		ScriptDirectory:   ScriptDirectory,
		TargetDirectory:   TargetDirectory,
		TempDirectory:     TempDirectory,
	}
}

// TimeflowPointSemanticFactory is just a simple function to instantiate the TimeflowPointSemanticStruct
func TimeflowPointSemanticFactory(
	Container string,
	Location string,
) TimeflowPointSemanticStruct {
	return TimeflowPointSemanticStruct{
		Container: Container,
		Location:  Location,
	}
}

// ToolkitFactory is just a simple function to instantiate the ToolkitStruct
func ToolkitFactory(
	Name string,
) ToolkitStruct {
	return ToolkitStruct{
		Name: Name,
	}
}

// ToolkitLinkedDirectSourceFactory is just a simple function to instantiate the ToolkitLinkedDirectSourceStruct
func ToolkitLinkedDirectSourceFactory(
	UsesGrandfatheredAppDataProperties *bool,
) ToolkitLinkedDirectSourceStruct {
	return ToolkitLinkedDirectSourceStruct{
		UsesGrandfatheredAppDataProperties: UsesGrandfatheredAppDataProperties,
	}
}

// TransformationFactory is just a simple function to instantiate the TransformationStruct
func TransformationFactory(
	Name string,
) TransformationStruct {
	return TransformationStruct{
		Name: Name,
	}
}

// UnixHostFactory is just a simple function to instantiate the UnixHostStruct
func UnixHostFactory(
	Address string,
	DspKeystoreAlias string,
	DspKeystorePassword string,
	DspKeystorePath string,
	DspTruststorePassword string,
	DspTruststorePath string,
	Name string,
	NfsAddressList []string,
	OracleJdbcKeystorePassword string,
	PrivilegeElevationProfile string,
	SshPort *int,
	SshVerificationStrategy SshVerificationStrategy,
	ToolkitPath string,
) UnixHostStruct {
	return UnixHostStruct{
		Address:                    Address,
		DspKeystoreAlias:           DspKeystoreAlias,
		DspKeystorePassword:        DspKeystorePassword,
		DspKeystorePath:            DspKeystorePath,
		DspTruststorePassword:      DspTruststorePassword,
		DspTruststorePath:          DspTruststorePath,
		Name:                       Name,
		NfsAddressList:             NfsAddressList,
		OracleJdbcKeystorePassword: OracleJdbcKeystorePassword,
		PrivilegeElevationProfile:  PrivilegeElevationProfile,
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
		Host: Host,
	}
}

// UnixHostEnvironmentFactory is just a simple function to instantiate the UnixHostEnvironmentStruct
func UnixHostEnvironmentFactory(
	AseHostEnvironmentParameters *ASEHostEnvironmentParametersStruct,
	Description string,
	LogCollectionEnabled *bool,
	Name string,
) UnixHostEnvironmentStruct {
	return UnixHostEnvironmentStruct{
		AseHostEnvironmentParameters: AseHostEnvironmentParameters,
		Description:                  Description,
		LogCollectionEnabled:         LogCollectionEnabled,
		Name:                         Name,
	}
}

// UpgradeCheckResultFactory is just a simple function to instantiate the UpgradeCheckResultStruct
func UpgradeCheckResultFactory(
	Name string,
) UpgradeCheckResultStruct {
	return UpgradeCheckResultStruct{
		Name: Name,
	}
}

// UpgradeCheckResultsVersionParametersFactory is just a simple function to instantiate the UpgradeCheckResultsVersionParametersStruct
func UpgradeCheckResultsVersionParametersFactory(
	BundleId string,
	Reference string,
) UpgradeCheckResultsVersionParametersStruct {
	return UpgradeCheckResultsVersionParametersStruct{
		BundleId:  BundleId,
		Reference: Reference,
	}
}

// UpgradeCompatibilityParametersFactory is just a simple function to instantiate the UpgradeCompatibilityParametersStruct
func UpgradeCompatibilityParametersFactory(
	Environment string,
) UpgradeCompatibilityParametersStruct {
	return UpgradeCompatibilityParametersStruct{
		Environment: Environment,
	}
}

// UserFactory is just a simple function to instantiate the UserStruct
func UserFactory(
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
	Principal string,
	PublicKey string,
	SessionTimeout *int,
	UserType string,
	WorkPhoneNumber string,
) UserStruct {
	return UserStruct{
		AuthenticationType: AuthenticationType,
		Credential:         Credential,
		EmailAddress:       EmailAddress,
		Enabled:            Enabled,
		FirstName:          FirstName,
		HomePhoneNumber:    HomePhoneNumber,
		IsDefault:          IsDefault,
		LastName:           LastName,
		Locale:             Locale,
		MobilePhoneNumber:  MobilePhoneNumber,
		Name:               Name,
		Principal:          Principal,
		PublicKey:          PublicKey,
		SessionTimeout:     SessionTimeout,
		UserType:           UserType,
		WorkPhoneNumber:    WorkPhoneNumber,
	}
}

// VerifyVersionParametersFactory is just a simple function to instantiate the VerifyVersionParametersStruct
func VerifyVersionParametersFactory(
	Defer *bool,
) VerifyVersionParametersStruct {
	return VerifyVersionParametersStruct{
		Defer: Defer,
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

// WindowsClusterFactory is just a simple function to instantiate the WindowsClusterStruct
func WindowsClusterFactory(
	Address string,
	Description string,
	LogCollectionEnabled *bool,
	Name string,
	Proxy string,
) WindowsClusterStruct {
	return WindowsClusterStruct{
		Address:              Address,
		Description:          Description,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Proxy:                Proxy,
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
		Cluster:              Cluster,
		LogCollectionEnabled: LogCollectionEnabled,
		PrimaryUser:          PrimaryUser,
		Target:               Target,
	}
}

// WindowsClusterNodeFactory is just a simple function to instantiate the WindowsClusterNodeStruct
func WindowsClusterNodeFactory(
	Name string,
) WindowsClusterNodeStruct {
	return WindowsClusterNodeStruct{
		Name: Name,
	}
}

// WindowsHostFactory is just a simple function to instantiate the WindowsHostStruct
func WindowsHostFactory(
	Address string,
	ConnectorAuthenticationKey string,
	ConnectorPort float64,
	DspKeystoreAlias string,
	DspKeystorePassword string,
	DspKeystorePath string,
	DspTruststorePassword string,
	DspTruststorePath string,
	Name string,
	NfsAddressList []string,
	PrivilegeElevationProfile string,
	SshPort *int,
) WindowsHostStruct {
	return WindowsHostStruct{
		Address:                    Address,
		ConnectorAuthenticationKey: ConnectorAuthenticationKey,
		ConnectorPort:              ConnectorPort,
		DspKeystoreAlias:           DspKeystoreAlias,
		DspKeystorePassword:        DspKeystorePassword,
		DspKeystorePath:            DspKeystorePath,
		DspTruststorePassword:      DspTruststorePassword,
		DspTruststorePath:          DspTruststorePath,
		Name:                       Name,
		NfsAddressList:             NfsAddressList,
		PrivilegeElevationProfile:  PrivilegeElevationProfile,
		SshPort:                    SshPort,
	}
}

// WindowsHostCreateParametersFactory is just a simple function to instantiate the WindowsHostCreateParametersStruct
func WindowsHostCreateParametersFactory(
	Host Host,
) WindowsHostCreateParametersStruct {
	return WindowsHostCreateParametersStruct{
		Host: Host,
	}
}

// WindowsHostEnvironmentFactory is just a simple function to instantiate the WindowsHostEnvironmentStruct
func WindowsHostEnvironmentFactory(
	Description string,
	LogCollectionEnabled *bool,
	Name string,
	Proxy string,
) WindowsHostEnvironmentStruct {
	return WindowsHostEnvironmentStruct{
		Description:          Description,
		LogCollectionEnabled: LogCollectionEnabled,
		Name:                 Name,
		Proxy:                Proxy,
	}
}

// WorkflowFunctionDefinitionFactory is just a simple function to instantiate the WorkflowFunctionDefinitionStruct
func WorkflowFunctionDefinitionFactory(
	InputSchema *SchemaDraftV4Struct,
	Name string,
	OutputSchema *SchemaDraftV4Struct,
) WorkflowFunctionDefinitionStruct {
	return WorkflowFunctionDefinitionStruct{
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
		City:             City,
		CommonName:       CommonName,
		Country:          Country,
		Organization:     Organization,
		OrganizationUnit: OrganizationUnit,
		StateRegion:      StateRegion,
	}
}

// X509CertificateFactory is just a simple function to instantiate the X509CertificateStruct
func X509CertificateFactory(
	Name string,
) X509CertificateStruct {
	return X509CertificateStruct{
		Name: Name,
	}
}

// XPPCompatibilityParametersFactory is just a simple function to instantiate the XPPCompatibilityParametersStruct
func XPPCompatibilityParametersFactory(
	Environment string,
) XPPCompatibilityParametersStruct {
	return XPPCompatibilityParametersStruct{
		Environment: Environment,
	}
}
