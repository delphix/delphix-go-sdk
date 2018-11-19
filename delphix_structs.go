package delphix

// APIErrorStruct - Description of an error encountered during an API call.
// extends TypedObject
type APIErrorStruct struct {
	// Action to be taken by the user, if any, to fix the underlying
	// problem.
	Action string `json:"action,omitempty"`
	// Extra output, often from a script or other external process, that
	// may give more insight into the cause of this error.
	CommandOutput string `json:"commandOutput,omitempty"`
	// For validation errors, a map of fields to APIError objects. For
	// all other errors, a string with further details of the error.
	Details string `json:"details,omitempty"`
	// Results of diagnostic checks run, if any, if the job failed.
	Diagnoses []*DiagnosisResultStruct `json:"diagnoses,omitempty"`
	// A stable identifier for the class of error encountered.
	Id string `json:"id,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// APISessionStruct - Describes a Delphix web service session and is the result of an
// initial handshake.
// extends TypedObject
type APISessionStruct struct {
	// Client software identification token.
	// required = false
	// maxLength = 64
	Client string `json:"client,omitempty"`
	// Locale as an IETF BCP 47 language tag, defaults to 'en-US'.
	// format = locale
	// required = false
	Locale string `json:"locale,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Version of the API to use.
	// required = true
	Version *APIVersionStruct `json:"version,omitempty"`
}

// APIVersionStruct - Describes an API version.
// extends TypedObject
type APIVersionStruct struct {
	// Major API version number.
	// required = true
	// minimum = 0
	Major *int `json:"major,omitempty"`
	// Micro API version number.
	// minimum = 0
	// required = true
	Micro *int `json:"micro,omitempty"`
	// Minor API version number.
	// minimum = 0
	// required = true
	Minor *int `json:"minor,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ASEAttachDataStruct - Represents the SAP ASE specific parameters of an attach request.
// extends AttachData
type ASEAttachDataStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-ase-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The credentials of the database user.
	// required = true
	// properties = map[type:map[type:string description:Object type. required:true format:type default:PasswordCredential]]
	DbCredentials Credential `json:"dbCredentials,omitempty"`
	// The user name for the source DB user.
	// create = optional
	DbUser string `json:"dbUser,omitempty"`
	// The credential for the source DB user.
	// create = optional
	DumpCredentials string `json:"dumpCredentials,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Source database backup location.
	// maxLength = 1024
	// required = true
	LoadBackupPath string `json:"loadBackupPath,omitempty"`
	// Backup location to use for loading backups from the source.
	// create = optional
	LoadLocation string `json:"loadLocation,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 87
	// create = optional
	MountBase string `json:"mountBase,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Information about the host OS user on the source to use for
	// linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	SourceHostUser string `json:"sourceHostUser,omitempty"`
	// Information about the host OS user on the staging environment to
	// use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// A user-provided shell script or executable to run after restoring
	// from a backup during validated sync.
	// maxLength = 1024
	// create = optional
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// A user-provided shell script or executable to run prior to
	// restoring from a backup during validated sync.
	// maxLength = 1024
	// create = optional
	StagingPreScript string `json:"stagingPreScript,omitempty"`
	// The SAP ASE instance on the staging environment that we want to
	// use for validated sync.
	// format = objectReference
	// referenceTo = /delphix-ase-instance.json
	// required = true
	StagingRepository string `json:"stagingRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Specifies the validated sync mode to synchronize the dSource with
	// the source database.
	// enum = [ENABLED DISABLED]
	// default = ENABLED
	// create = optional
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
}

// ASEBackupLocationStruct - SAP ASE backup location.
// extends TypedObject
type ASEBackupLocationStruct struct {
	// Host environment where the backup server is located.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-host-environment.json
	// create = required
	BackupHost string `json:"backupHost,omitempty"`
	// OS user for the host where the backup server is located.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = required
	// update = optional
	BackupHostUser string `json:"backupHostUser,omitempty"`
	// Name of the backup server instance.
	// create = required
	// update = optional
	BackupServerName string `json:"backupServerName,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ASECompatibilityCriteriaStruct - The compatibility criteria to use for filtering the list of available
// SAP ASE repositories.
// extends CompatibilityCriteria
type ASECompatibilityCriteriaStruct struct {
	// Selected repositories are installed on a host with this
	// architecture (32-bit or 64-bit).
	Architecture *int `json:"architecture,omitempty"`
	// Selected repositories are installed on this environment.
	// referenceTo = /delphix-source-environment.json
	// format = objectReference
	Environment string `json:"environment,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASECreateTransformationParametersStruct - Represents the parameters of a createTransformation request for an ASE
// container.
// extends CreateTransformationParameters
type ASECreateTransformationParametersStruct struct {
	// The container that will contain the transformed data associated
	// with the newly created transformation; the "transformation
	// container".
	// create = required
	Container *ASEDBContainerStruct `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEDBContainerStruct - An SAP ASE Database Container.
// extends DatabaseContainer
type ASEDBContainerStruct struct {
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// A reference to the group containing this container.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The operating system for the source database.
	Os string `json:"os,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	// update = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// The processor type for the source database.
	Processor string `json:"processor,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this container.
	Runtime *ASEDBContainerRuntimeStruct `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources for an
	// SAP ASE container.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEDBContainerRuntimeStruct - Runtime properties of an SAP ASE database container.
// extends DBContainerRuntime
type ASEDBContainerRuntimeStruct struct {
	// The source database backupset that was last restored in this
	// container.
	// format = date
	LastRestoredBackupDate string `json:"lastRestoredBackupDate,omitempty"`
	// The timezone for the last restored source database backupset in
	// this container.
	LastRestoredBackupTimeZone string `json:"lastRestoredBackupTimeZone,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntimeStruct `json:"preProvisioningStatus,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ASEExportParametersStruct - The parameters to use as input to export SAP ASE databases.
// extends DbExportParameters
type ASEExportParametersStruct struct {
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig ASEDBConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export
	// on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEHostEnvironmentParametersStruct - SAP ASE host environment parameters.
// extends TypedObject
type ASEHostEnvironmentParametersStruct struct {
	// The credentials of the database user.
	// create = required
	// update = optional
	// properties = map[type:map[type:string description:Object type. required:true format:type default:PasswordCredential]]
	Credentials Credential `json:"credentials,omitempty"`
	// The username of the database user.
	// create = optional
	// update = optional
	// maxLength = 256
	DbUser string `json:"dbUser,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEInstanceStruct - The SAP ASE source repository.
// extends SourceRepository
type ASEInstanceStruct struct {
	// The credentials of the database user.
	// create = optional
	// update = optional
	// properties = map[type:map[required:true format:type default:PasswordCredential type:string description:Object type.]]
	Credentials Credential `json:"credentials,omitempty"`
	// The username of the database user.
	// create = optional
	// update = optional
	// maxLength = 256
	DbUser string `json:"dbUser,omitempty"`
	// True if the SAP ASE instance was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// The SAP ASE instance home.
	// create = required
	// update = optional
	InstallationPath string `json:"installationPath,omitempty"`
	// The name of the SAP ASE instance.
	// create = required
	InstanceName string `json:"instanceName,omitempty"`
	// The username of the account the SAP ASE instance is running as.
	// create = required
	// update = optional
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// The gid of the account the SAP ASE instance is running as.
	// create = readonly
	// update = readonly
	InstanceOwnerGid *int `json:"instanceOwnerGid,omitempty"`
	// The uid of the account the SAP ASE instance is running as.
	// create = readonly
	// update = readonly
	InstanceOwnerUid *int `json:"instanceOwnerUid,omitempty"`
	// The path to the isql binary to use for this SAP ASE instance.
	// create = optional
	// update = optional
	IsqlPath string `json:"isqlPath,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Database page size for the SAP ASE instance.
	PageSize *int `json:"pageSize,omitempty"`
	// The network ports for connecting to the SAP ASE instance.
	// create = required
	// update = optional
	Ports []*int `json:"ports,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// create = optional
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The Kerberos SPN of the database.
	// create = optional
	// update = optional
	ServicePrincipalName string `json:"servicePrincipalName,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// create = optional
	// update = optional
	// default = false
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of the repository.
	// create = optional
	// update = optional
	Version string `json:"version,omitempty"`
}

// ASELatestBackupSyncParametersStruct - The parameters to use as input to sync a SAP ASE database using the
// latest backup.
// extends ASESyncParameters
type ASELatestBackupSyncParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASELinkDataStruct - SAP ASE specific parameters for a link request.
// extends LinkData
type ASELinkDataStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-ase-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The credentials of the database user.
	// required = true
	// properties = map[type:map[format:type default:PasswordCredential type:string description:Object type. required:true]]
	DbCredentials Credential `json:"dbCredentials,omitempty"`
	// The user name for the source DB user.
	// create = optional
	DbUser string `json:"dbUser,omitempty"`
	// The credential for the source DB user.
	// create = optional
	DumpCredentials string `json:"dumpCredentials,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Source database backup location.
	// maxLength = 1024
	// required = true
	LoadBackupPath string `json:"loadBackupPath,omitempty"`
	// Backup location to use for loading backups from the source.
	// create = optional
	LoadLocation string `json:"loadLocation,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 87
	// create = optional
	MountBase string `json:"mountBase,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Information about the host OS user on the source to use for
	// linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	SourceHostUser string `json:"sourceHostUser,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// Information about the host OS user on the staging environment to
	// use for linking.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// A user-provided shell script or executable to run after restoring
	// from a backup during validated sync.
	// maxLength = 1024
	// create = optional
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// A user-provided shell script or executable to run prior to
	// restoring from a backup during validated sync.
	// maxLength = 1024
	// create = optional
	StagingPreScript string `json:"stagingPreScript,omitempty"`
	// The SAP ASE instance on the staging environment that we want to
	// use for validated sync.
	// format = objectReference
	// referenceTo = /delphix-ase-instance.json
	// required = true
	StagingRepository string `json:"stagingRepository,omitempty"`
	// Sync parameters for the container.
	// required = true
	// properties = map[type:map[type:string description:Object type. required:true format:type default:ASELatestBackupSyncParameters]]
	SyncParameters ASESyncParameters `json:"syncParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Specifies the validated sync mode to synchronize the dSource with
	// the source database.
	// create = optional
	// enum = [ENABLED DISABLED]
	// default = ENABLED
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
}

// ASELinkedSourceStruct - A linked SAP ASE source.
// extends ASESource
type ASELinkedSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-ase-db-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// The credential for the source DB user.
	// create = optional
	// update = optional
	DumpCredentials string `json:"dumpCredentials,omitempty"`
	// External file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Source database backup location.
	// create = required
	// update = optional
	// maxLength = 1024
	LoadBackupPath string `json:"loadBackupPath,omitempty"`
	// Backup location to use for loading backups from the source.
	// create = optional
	// update = optional
	LoadLocation string `json:"loadLocation,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *ASESourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// The staging source for validated sync of the database.
	// referenceTo = /delphix-ase-staging-source.json
	// format = objectReference
	StagingSource string `json:"stagingSource,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Specifies the validated sync mode to synchronize the dSource with
	// the source database.
	// enum = [ENABLED DISABLED]
	// default = ENABLED
	// create = optional
	// update = optional
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// ASENewBackupSyncParametersStruct - The parameters to use as input to sync a SAP ASE database by taking a
// new full backup.
// extends ASESyncParameters
type ASENewBackupSyncParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEPlatformParametersStruct - ASE platform-specific parameters that are stored on a transformation.
// extends BasePlatformParameters
type ASEPlatformParametersStruct struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ASEProvisionParametersStruct - The parameters to use as input to provision SAP ASE databases.
// extends ProvisionParameters
type ASEProvisionParametersStruct struct {
	// The new container for the provisioned database.
	// required = true
	Container *ASEDBContainerStruct `json:"container,omitempty"`
	// Whether or not to mark this VDB as a masked VDB. It will be marked
	// as masked if this flag or the masking job are set.
	// create = optional
	// update = readonly
	Masked *bool `json:"masked,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// update = readonly
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	MaskingJob string `json:"maskingJob,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *ASEVirtualSourceStruct `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of
	// the source.
	// required = true
	SourceConfig ASEDBConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Set the "trunc log on chkpt" database option.
	// default = true
	// create = required
	TruncateLogOnCheckpoint *bool `json:"truncateLogOnCheckpoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASESIConfigStruct - A SAP ASE single instance database config.
// extends ASEDBConfig
type ASESIConfigStruct struct {
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// The name of the database.
	// create = required
	// update = optional
	// pattern = ^[a-zA-Z0-9_]+$
	// maxLength = 30
	DatabaseName string `json:"databaseName,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// create = required
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-ase-instance.json
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
}

// ASESnapshotStruct - Provisionable snapshot of a SAP ASE TimeFlow.
// extends TimeflowSnapshot
type ASESnapshotStruct struct {
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *ASETimeflowPointStruct `json:"firstChangePoint,omitempty"`
	// The location of the snapshot within the parent TimeFlow
	// represented by this snapshot.
	LatestChangePoint *ASETimeflowPointStruct `json:"latestChangePoint,omitempty"`
	// Boolean value indicating if a virtual database provisioned from
	// this snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot
	// should be kept forever.
	// update = optional
	Retention *int `json:"retention,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *ASESnapshotRuntimeStruct `json:"runtime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient
	// state and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Time zone of the source database at the time the snapshot was
	// taken.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
}

// ASESnapshotRuntimeStruct - Runtime (non-persistent) properties of a SAP ASE TimeFlow snapshot.
// extends SnapshotRuntime
type ASESnapshotRuntimeStruct struct {
	// True if this snapshot can be used as the basis for provisioning a
	// new TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ASESourceConnectionInfoStruct - Contains information that can be used to connect to an SAP ASE source.
// extends SourceConnectionInfo
type ASESourceConnectionInfoStruct struct {
	// The database name of the source.
	DatabaseName string `json:"databaseName,omitempty"`
	// Host to use for connecting to the source.
	Host string `json:"host,omitempty"`
	// The JDBC string used to connect to the source.
	JdbcString string `json:"jdbcString,omitempty"`
	// Port to use for connecting to the source.
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User to use for connecting to the source.
	User string `json:"user,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// ASESourceRuntimeStruct - Runtime (non-persistent) properties of a SAP ASE source.
// extends SourceRuntime
type ASESourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// SAP ASE database durability level.
	// enum = [FULL AT_SHUTDOWN NO_RECOVERY]
	DurabilityLevel string `json:"durabilityLevel,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// True if configured to truncate log on checkpoint.
	TruncateLogOnCheckpoint *bool `json:"truncateLogOnCheckpoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASESpecificBackupSyncParametersStruct - The parameters to use as input to sync a SAP ASE database using a
// specific existing backup.
// extends ASESyncParameters
type ASESpecificBackupSyncParametersStruct struct {
	// The location of the full backup of the source database to restore
	// from. The backup should be present in the shared backup location
	// for the source database.
	// uniqueItems = true
	// minItems = 1
	// create = required
	BackupFiles []string `json:"backupFiles,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEStagingSourceStruct - An SAP ASE staging source used for validated sync..
// extends ASESource
type ASEStagingSourceStruct struct {
	// Reference to the configuration for the source.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point for the NFS mounts.
	// update = optional
	// maxLength = 87
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The path to a user-provided shell script or executable to run on
	// the staging host after restoring from a backup during validated
	// sync.
	// create = optional
	// update = optional
	// maxLength = 1024
	PostScript string `json:"postScript,omitempty"`
	// The path to a user-provided script or executable to run on the
	// staging host prior to restoring from a backup during validated
	// sync.
	// create = optional
	// update = optional
	// maxLength = 1024
	PreScript string `json:"preScript,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *ASESourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// ASETimeflowStruct - TimeFlow representing historical data for a particular timeline within
// a SAP ASE data container.
// extends Timeflow
type ASETimeflowStruct struct {
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC SOURCE_CONTINUITY]
	CreationType string `json:"creationType,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow
	// was provisioned. This will not be present for TimeFlows derived
	// from linked sources.
	ParentPoint *ASETimeflowPointStruct `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning
	// base for this object. This may be different from the snapshot
	// within the parent point, and is only present for virtual
	// TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASETimeflowPointStruct - A unique point within a SAP ASE database TimeFlow.
// extends TimeflowPoint
type ASETimeflowPointStruct struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this point.
	// format = objectReference
	// referenceTo = /delphix-ase-timeflow.json
	// required = true
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEVirtualSourceStruct - A virtual SAP ASE source.
// extends ASESource
type ASEVirtualSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// SAP ASE database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	// maxLength = 4096
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point for the NFS mounts.
	// maxLength = 87
	// create = optional
	// update = optional
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *ASESourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// ActionStruct - Represents an action, a permanent record of activity on the server.
// extends PersistentObject
// cliVisibility = [SYSTEM DOMAIN]
type ActionStruct struct {
	// Action type.
	ActionType string `json:"actionType,omitempty"`
	// Plain text description of the action.
	Details string `json:"details,omitempty"`
	// The time the action completed.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Action to be taken to resolve the failure.
	FailureAction string `json:"failureAction,omitempty"`
	// Details of the action failure.
	FailureDescription string `json:"failureDescription,omitempty"`
	// Message ID associated with the event.
	FailureMessageCode string `json:"failureMessageCode,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Network address used to initiate the action.
	OriginIp string `json:"originIp,omitempty"`
	// The parent action of this action.
	// format = objectReference
	// referenceTo = /delphix-action.json
	ParentAction string `json:"parentAction,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Report of progress and warnings for some actions.
	Report string `json:"report,omitempty"`
	// The time the action occurred. For long running processes, this
	// represents the starting time.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// State of the action.
	// enum = [EXECUTING WAITING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// Action title.
	Title string `json:"title,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User who initiated the action.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
	// Name of client software used to initiate the action.
	UserAgent string `json:"userAgent,omitempty"`
	// Origin of the work that caused the action.
	// enum = [WEBSERVICE POLICY SYSTEM]
	WorkSource string `json:"workSource,omitempty"`
	// Name of user or policy that initiated the action.
	WorkSourceName string `json:"workSourceName,omitempty"`
}

// AlertStruct - An alert describing an event for a given object.
// extends PersistentObject
type AlertStruct struct {
	// Event class.
	Event string `json:"event,omitempty"`
	// Event recommended action.
	EventAction string `json:"eventAction,omitempty"`
	// Additional text associated with the event. This text is not
	// localized and is only provided for certain alerts. For example, if
	// an alert is caused by a post script failure, the output of the
	// post script may be included here to assist with debugging the
	// failure.
	EventCommandOutput string `json:"eventCommandOutput,omitempty"`
	// Event description.
	EventDescription string `json:"eventDescription,omitempty"`
	// Event response.
	EventResponse string `json:"eventResponse,omitempty"`
	// Event severity.
	// enum = [INFORMATIONAL WARNING CRITICAL AUDIT]
	EventSeverity string `json:"eventSeverity,omitempty"`
	// Event title.
	EventTitle string `json:"eventTitle,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Reference to target object.
	// referenceTo = /delphix-user-object.json
	// format = objectReference
	Target string `json:"target,omitempty"`
	// Name of target object.
	TargetName string `json:"targetName,omitempty"`
	// Type of target object.
	// format = type
	TargetObjectType string `json:"targetObjectType,omitempty"`
	// Time at which event occurred.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AlertActionEmailListStruct - Alert action to email a list of users in response to an alert.
// extends AlertActionEmail
type AlertActionEmailListStruct struct {
	// List of email addresses to send mail to.
	// required = true
	Addresses []string `json:"addresses,omitempty"`
	// Email format to use. The HTML format will generate a multipart
	// message containing both HTML and plain text. The TEXT format will
	// explicitly generate text-only mail. The JSON format will generate
	// a JSON object identical to the $Alert format returned through the
	// web services API.
	// enum = [HTML TEXT JSON]
	// default = HTML
	// create = optional
	// update = optional
	Format string `json:"format,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AlertActionEmailUserStruct - Alert action that sends email to the email address associated with the
// user.
// extends AlertActionEmail
type AlertActionEmailUserStruct struct {
	// Email format to use. The HTML format will generate a multipart
	// message containing both HTML and plain text. The TEXT format will
	// explicitly generate text-only mail. The JSON format will generate
	// a JSON object identical to the $Alert format returned through the
	// web services API.
	// default = HTML
	// create = optional
	// update = optional
	// enum = [HTML TEXT JSON]
	Format string `json:"format,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AlertProfileStruct - A profile that describes a set of actions to take in response to an
// alert being generated.
// extends PersistentObject
type AlertProfileStruct struct {
	// List of actions to take. Only alerts visible to the user and
	// matching the optional filters are included. If there are multiple
	// actions with the same result (such as emailing a user), only one
	// result is acted upon.
	// update = optional
	// create = required
	Actions []AlertAction `json:"actions,omitempty"`
	// Specifies which alerts should be matched by this profile.
	// create = optional
	// update = optional
	// properties = map[type:map[default:SeverityFilter]]
	FilterSpec AlertFilter `json:"filterSpec,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AndFilterStruct - A container filter that combines other filters together using AND
// logic.
// extends AlertFilter
type AndFilterStruct struct {
	// Filters which are combined together using AND logic.
	// update = optional
	// minItems = 2
	// create = required
	SubFilters []AlertFilter `json:"subFilters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataAdditionalMountPointStruct - Specifies an additional location on which to mount a subdirectory of
// an AppData container.
// extends TypedObject
type AppDataAdditionalMountPointStruct struct {
	// Reference to the environment on which the file system will be
	// mounted.
	// format = objectReference
	// referenceTo = /delphix-host-environment.json
	// create = required
	// update = optional
	Environment string `json:"environment,omitempty"`
	// Absolute path on the target environment were the filesystem should
	// be mounted.
	// format = unixpath
	// create = required
	// update = optional
	MountPath string `json:"mountPath,omitempty"`
	// Relative path within the container of the directory that should be
	// mounted.
	// format = unixpath
	// create = optional
	// update = optional
	SharedPath string `json:"sharedPath,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataCachedMountPointStruct - Specified information about an active mount of an AppData container.
// extends TypedObject
type AppDataCachedMountPointStruct struct {
	// Reference to the environment on which the file system is mounted.
	// format = objectReference
	// referenceTo = /delphix-host-environment.json
	// create = required
	// update = optional
	Environment string `json:"environment,omitempty"`
	// Absolute path on the target environment were the filesystem is
	// mounted.
	// format = unixpath
	// create = required
	// update = optional
	MountPath string `json:"mountPath,omitempty"`
	// Order in mount sequence.
	// update = optional
	// create = required
	Ordinal *int `json:"ordinal,omitempty"`
	// Relative path within the container of the directory that is
	// mounted.
	// format = unixpath
	// create = optional
	// update = optional
	SharedPath string `json:"sharedPath,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataContainerStruct - Data container for AppData.
// extends DatabaseContainer
type AppDataContainerStruct struct {
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Optional user-provided description for the container.
	// update = optional
	// maxLength = 1024
	// create = optional
	Description string `json:"description,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// A global identifier for this container, including across Delphix
	// Engines.
	Guid string `json:"guid,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Whether to enable high performance mode.
	// default = DISABLED
	// create = readonly
	// update = readonly
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	PerformanceMode string `json:"performanceMode,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this container.
	Runtime *AppDataContainerRuntimeStruct `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// The toolkit managing the data in the container.
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataContainerRuntimeStruct - Runtime properties of an AppData container.
// extends DBContainerRuntime
type AppDataContainerRuntimeStruct struct {
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntimeStruct `json:"preProvisioningStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataCreateTransformationParametersStruct - Represents the parameters of a createTransformation request for an
// AppData container.
// extends CreateTransformationParameters
type AppDataCreateTransformationParametersStruct struct {
	// The container that will contain the transformed data associated
	// with the newly created transformation; the "transformation
	// container".
	// create = required
	Container *AppDataContainerStruct `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	// create = required
	Payload *JsonStruct `json:"payload,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataDirectLinkDataStruct - Represents the AppData specific parameters of a link request for a
// source directly replicated into the Delphix Engine.
// extends AppDataLinkData
type AppDataDirectLinkDataStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-appdata-direct-source-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// List of subdirectories in the source to exclude when syncing data.
	// These paths are relative to the root of the source directory.
	// create = optional
	Excludes []string `json:"excludes,omitempty"`
	// List of symlinks in the source to follow when syncing data. These
	// paths are relative to the root of the source directory. All other
	// symlinks are preserved.
	// create = optional
	FollowSymlinks []string `json:"followSymlinks,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	// required = true
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataDirectSourceConfigStruct - Source config for directly linked AppData sources.
// extends AppDataSourceConfig
type AppDataDirectSourceConfigStruct struct {
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// update = optional
	// default = true
	// create = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The name of the config.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The list of parameters specified by the source config schema in
	// the toolkit. If no schema is specified, this list is empty.
	// create = optional
	// update = optional
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// The path to the data to be synced.
	// create = optional
	// update = optional
	// maxLength = 1024
	Path string `json:"path,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-appdata-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// The toolkit associated with this source config.
	// referenceTo = /delphix-toolkit.json
	// format = objectReference
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataEmptyVFilesCreationParametersStruct - The parameters to use as input when creating a new empty vFiles
// dataset.
// extends EmptyDatasetCreationParameters
type AppDataEmptyVFilesCreationParametersStruct struct {
	// The new container for the created dataset.
	// required = true
	Container *AppDataContainerStruct `json:"container,omitempty"`
	// The source that describes an external dataset instance.
	// required = true
	Source *AppDataVirtualSourceStruct `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of
	// the source.
	// required = true
	SourceConfig AppDataSourceConfig `json:"sourceConfig,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// AppDataExportParametersStruct - The parameters to use as input to export AppData.
// extends ExportParameters
type AppDataExportParametersStruct struct {
	// The filesystem configuration of the exported database.
	// create = optional
	FilesystemLayout *AppDataFilesystemLayoutStruct `json:"filesystemLayout,omitempty"`
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig AppDataSourceConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export
	// on.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// AppDataFilesystemLayoutStruct - A filesystem layout that matches the filesystem of a Delphix TimeFlow.
// extends FilesystemLayout
type AppDataFilesystemLayoutStruct struct {
	// The base directory to use for the exported database.
	// required = true
	TargetDirectory string `json:"targetDirectory,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataLinkedDirectSourceStruct - An AppData linked source directly replicated into the Delphix Engine.
// extends AppDataLinkedSource
type AppDataLinkedDirectSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-appdata-source-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// List of subdirectories in the source to exclude when syncing data.
	// These paths are relative to the root of the source directory.
	// create = optional
	// update = optional
	Excludes []string `json:"excludes,omitempty"`
	// List of symlinks in the source to follow when syncing data. These
	// paths are relative to the root of the source directory. All other
	// symlinks are preserved.
	// create = optional
	// update = optional
	FollowSymlinks []string `json:"followSymlinks,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	// create = required
	// update = optional
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *AppDataSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// The toolkit associated with this source.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// AppDataLinkedStagedSourceStruct - An AppData linked source with a staging source.
// extends AppDataLinkedSource
type AppDataLinkedStagedSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-appdata-source-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// create = optional
	// update = optional
	// maxLength = 256
	// format = objectName
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	// create = required
	// update = optional
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *AppDataSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// The environment used as an intermediate stage to pull data into
	// Delphix.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	// update = optional
	StagingEnvironment string `json:"stagingEnvironment,omitempty"`
	// The environment user used to access the staging environment.
	// referenceTo = /delphix-source-environment-user.json
	// create = required
	// update = optional
	// format = objectReference
	StagingEnvironmentUser string `json:"stagingEnvironmentUser,omitempty"`
	// The base mount point for the NFS mount on the staging environment.
	// maxLength = 256
	// create = required
	// update = optional
	StagingMountBase string `json:"stagingMountBase,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// The toolkit associated with this source.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// AppDataPlatformParametersStruct - AppData platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type AppDataPlatformParametersStruct struct {
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	// create = required
	Payload *JsonStruct `json:"payload,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataProvisionParametersStruct - The parameters to use as input to provision AppData.
// extends ProvisionParameters
type AppDataProvisionParametersStruct struct {
	// The new container for the provisioned database.
	// required = true
	Container *AppDataContainerStruct `json:"container,omitempty"`
	// Whether or not to mark this VDB as a masked VDB. It will be marked
	// as masked if this flag or the masking job are set.
	// create = optional
	// update = readonly
	Masked *bool `json:"masked,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// update = readonly
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	MaskingJob string `json:"maskingJob,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *AppDataVirtualSourceStruct `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of
	// the source.
	// required = true
	SourceConfig AppDataSourceConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataRepositoryStruct - An AppData repository.
// extends SourceRepository
type AppDataRepositoryStruct struct {
	// Reference to the environment containing this repository.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The list of parameters specified by the repository schema in the
	// toolkit. If no schema is specified, this list is empty.
	// create = required
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// create = optional
	// update = optional
	// default = true
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// create = optional
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// The toolkit associated with this repository.
	// referenceTo = /delphix-toolkit.json
	// create = required
	// update = optional
	// format = objectReference
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of the repository.
	// update = optional
	// create = optional
	Version string `json:"version,omitempty"`
}

// AppDataSnapshotStruct - Snapshot of an AppData TimeFlow.
// extends TimeflowSnapshot
type AppDataSnapshotStruct struct {
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *AppDataTimeflowPointStruct `json:"firstChangePoint,omitempty"`
	// The location of the snapshot within the parent TimeFlow
	// represented by this snapshot.
	LatestChangePoint *AppDataTimeflowPointStruct `json:"latestChangePoint,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	Metadata *JsonStruct `json:"metadata,omitempty"`
	// Boolean value indicating if a virtual database provisioned from
	// this snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot
	// should be kept forever.
	// update = optional
	Retention *int `json:"retention,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *AppDataSnapshotRuntimeStruct `json:"runtime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient
	// state and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Time zone of the source database at the time the snapshot was
	// taken.
	Timezone string `json:"timezone,omitempty"`
	// The toolkit associated with this snapshot.
	// referenceTo = /delphix-toolkit.json
	// format = objectReference
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
}

// AppDataSnapshotRuntimeStruct - Runtime (non-persistent) properties of AppData TimeFlow snapshots.
// extends SnapshotRuntime
type AppDataSnapshotRuntimeStruct struct {
	// True if this snapshot can be used as the basis for provisioning a
	// new TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// AppDataSourceConnectionInfoStruct - Contains information that can be used to connect to the application
// source.
// extends SourceConnectionInfo
type AppDataSourceConnectionInfoStruct struct {
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The path where the application data is located on the host.
	Path string `json:"path,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// AppDataSourceRuntimeStruct - Runtime (non-persistent) properties of an AppData source.
// extends SourceRuntime
type AppDataSourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataStagedLinkDataStruct - Represents the AppData specific parameters of a link request for a
// source with a staging source.
// extends AppDataLinkData
type AppDataStagedLinkDataStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-appdata-staged-source-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	// required = true
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// The environment used as an intermediate stage to pull data into
	// Delphix.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	StagingEnvironment string `json:"stagingEnvironment,omitempty"`
	// The environment user used to access the staging environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	StagingEnvironmentUser string `json:"stagingEnvironmentUser,omitempty"`
	// The base mount point for the NFS mount on the staging environment.
	// maxLength = 256
	// create = optional
	StagingMountBase string `json:"stagingMountBase,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataStagedSourceConfigStruct - An AppData source config with a staging source.
// extends AppDataSourceConfig
type AppDataStagedSourceConfigStruct struct {
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The name of the config.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The list of parameters specified by the source config schema in
	// the toolkit. If no schema is specified, this list is empty.
	// create = optional
	// update = optional
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-appdata-source-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The toolkit associated with this source config.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataSyncParametersStruct - The parameters to use as input to sync an AppData source.
// extends SyncParameters
type AppDataSyncParametersStruct struct {
	// Whether or not to force a non-incremental load of data prior to
	// taking a snapshot.
	// default = false
	Resync *bool `json:"resync,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataTimeflowStruct - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends AppDataBaseTimeflow
type AppDataTimeflowStruct struct {
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC SOURCE_CONTINUITY]
	CreationType string `json:"creationType,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow
	// was provisioned. This will not be present for TimeFlows derived
	// from linked sources.
	ParentPoint *AppDataTimeflowPointStruct `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning
	// base for this object. This may be different from the snapshot
	// within the parent point, and is only present for virtual
	// TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AppDataTimeflowPointStruct - A unique point within an AppData TimeFlow.
// extends TimeflowPoint
type AppDataTimeflowPointStruct struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this point.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	// required = true
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// AppDataVirtualSourceStruct - A virtual AppData source.
// extends AppDataManagedSource
type AppDataVirtualSourceStruct struct {
	// Locations to mount subdirectories of the AppData in addition to
	// the normal target mount point. These paths will be mounted and
	// unmounted as part of enabling and disabling this source.
	// create = optional
	// update = optional
	AdditionalMountPoints []*AppDataAdditionalMountPointStruct `json:"additionalMountPoints,omitempty"`
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the
	// type of application data being manipulated.
	// create = required
	// update = optional
	Parameters *JsonStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *AppDataSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// The toolkit associated with this source.
	// referenceTo = /delphix-toolkit.json
	// format = objectReference
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// AppDataWindowsTimeflowStruct - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends AppDataBaseTimeflow
type AppDataWindowsTimeflowStruct struct {
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC SOURCE_CONTINUITY]
	CreationType string `json:"creationType,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow
	// was provisioned. This will not be present for TimeFlows derived
	// from linked sources.
	ParentPoint *AppDataTimeflowPointStruct `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning
	// base for this object. This may be different from the snapshot
	// within the parent point, and is only present for virtual
	// TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ApplyVersionParametersStruct - The parameters to use as input to upgrade.
// extends TypedObject
type ApplyVersionParametersStruct struct {
	// If true, the Delphix Engine is upgraded without updating the OS
	// software. The operation will fail gracefully if the upgrade
	// version requires a version of the OS that is newer than what is
	// currently running. The OS software can subsequently be upgraded by
	// applying any version and setting defer to false. It is possible to
	// catch up to the current OS version on a previously deferred
	// upgrade by re-applying the running version with a defer setting of
	// false.
	// create = optional
	// default = false
	Defer *bool `json:"defer,omitempty"`
	// This property governs whether or not data sources are re-enabled
	// or left disabled in the event that upgrade fails before the
	// Delphix Engine is restarted.
	// create = optional
	// default = true
	EnableSourcesOnFailure *bool `json:"enableSourcesOnFailure,omitempty"`
	// If true, a failure to quiesce sources will not block the upgrade.
	// create = optional
	// default = false
	IgnoreQuiesceSourcesFailures *bool `json:"ignoreQuiesceSourcesFailures,omitempty"`
	// If true, all data sources (VDBs and dSources) are automatically
	// disabled prior to upgrade, and re-enabled after upgrade. If any
	// source cannot be disabled, the recovery semantics are governed by
	// the "ignoreQuiesceSourcesFailures" and "enableSourcesOnFailure"
	// properties.
	// create = optional
	// default = true
	QuiesceSources *bool `json:"quiesceSources,omitempty"`
	// If true, the system reboots immediately after upgrade. If false,
	// the system is shutdown.
	// default = true
	// create = optional
	Reboot *bool `json:"reboot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If set to false, disables verification before applying the
	// upgrade. This will only disable verification if a successful
	// verification has been run in the past hour.
	// create = optional
	// default = true
	Verify *bool `json:"verify,omitempty"`
}

// AttachSourceParametersStruct - Represents the parameters of an attach request.
// extends TypedObject
type AttachSourceParametersStruct struct {
	// The database-specific parameters of an attach request.
	// required = true
	AttachData AttachData `json:"attachData,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AuthFilterParametersStruct - The parameters to use as input to filter a list of objects or object
// type by a permission.
// extends TypedObject
type AuthFilterParametersStruct struct {
	// The object type on which to perform filtering. This option is
	// mutually exclusive with the "objects" field.
	// format = type
	ObjectType string `json:"objectType,omitempty"`
	// The list of objects to filter. This option is mutually exclusive
	// with the "objectType" field.
	Objects []string `json:"objects,omitempty"`
	// The permission to filter by.
	// format = objectReference
	// referenceTo = /delphix-permission.json
	Permission string `json:"permission,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AuthFilterResultStruct - Result of an auth filter request.
// extends TypedObject
type AuthFilterResultStruct struct {
	// The list of objects that have been filtered.
	Objects []string `json:"objects,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AuthGetByPropertiesParametersStruct - The parameters to use as input to get an authorization with given
// user, target, and role.
// extends TypedObject
type AuthGetByPropertiesParametersStruct struct {
	// The role type authorizations are applied to.
	// format = objectReference
	// required = true
	// referenceTo = /delphix-role.json
	Role string `json:"role,omitempty"`
	// The target authorizations are applied to.
	// format = objectReference
	// required = true
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The user authorizations are applied to.
	// format = objectReference
	// required = true
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
}

// AuthorizationStruct - Describes a role as applied to a user on an object.
// extends ReadonlyNamedUserObject
type AuthorizationStruct struct {
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Applied role.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-role.json
	Role string `json:"role,omitempty"`
	// Reference to the object that the authorization applies to.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the user that the authorization applies to.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
}

// AuthorizationConfigStruct - Configuration of the Authorization System.
// extends TypedObject
// cliVisibility = [DOMAIN]
type AuthorizationConfigStruct struct {
	// Whether authorization checks are enabled for all environment and
	// host operations.
	// update = optional
	// default = false
	EnvironmentAndHostAuth *bool `json:"environmentAndHostAuth,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// BatchContainerDeleteParametersStruct - The parameters to use as input to batch container delete requests.
// extends TypedObject
type BatchContainerDeleteParametersStruct struct {
	// Containers to delete.
	// minItems = 1
	// required = true
	Containers []string `json:"containers,omitempty"`
	// Optional parameters to the delete operations.
	// required = false
	DeleteParameters *DeleteParametersStruct `json:"deleteParameters,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// BatchContainerRefreshParametersStruct - The parameters to use as input to batch container refresh requests.
// extends TypedObject
type BatchContainerRefreshParametersStruct struct {
	// Containers to refresh.
	// minItems = 1
	// required = true
	Containers []string `json:"containers,omitempty"`
	// A semantic description of a TimeFlow location.
	// default = LATEST_POINT
	// required = false
	// enum = [LATEST_POINT LATEST_SNAPSHOT]
	Location string `json:"location,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// BatchSnapshotDeleteParametersStruct - The parameters to use as input to TimeFlow snapshots batch delete
// requests.
// extends TypedObject
type BatchSnapshotDeleteParametersStruct struct {
	// TimeFlow snapshots to delete.
	// required = true
	// minItems = 1
	Snapshots []string `json:"snapshots,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// BooleanEqualConstraintStruct - Constraints placed on a boolean axis of a particular analytics slice.
// extends BooleanConstraint
type BooleanEqualConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be equal to the boolean argument.
	// create = required
	Equals *bool `json:"equals,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// CPUInfoStruct - Describes a processor available to the system.
// extends TypedObject
type CPUInfoStruct struct {
	// Number of cores in the processor.
	Cores *int `json:"cores,omitempty"`
	// Speed of the processor, in hertz.
	// units = Hz
	// base = 1000
	Speed float64 `json:"speed,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CaCertificateStruct - Public Key Certificate that is a Certificate Authority.
// extends Certificate
// cliVisibility = [SYSTEM DOMAIN]
type CaCertificateStruct struct {
	// Delphix trusts this certificate .
	Accepted *bool `json:"accepted,omitempty"`
	// Issuer of this certificate.
	IssuedByDN string `json:"issuedByDN,omitempty"`
	// A reference to the certificate that issued this certificate. Null
	// if this is a root CA.
	// format = objectReference
	// referenceTo = /delphix-ca-certificate.json
	Issuer string `json:"issuer,omitempty"`
	// MD5 fingerprint.
	Md5Fingerprint string `json:"md5Fingerprint,omitempty"`
	// The Distinguished Name of this certificate.
	// format = objectName
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// End of validity.
	// format = date
	NotAfter string `json:"notAfter,omitempty"`
	// Start of validity.
	// format = date
	NotBefore string `json:"notBefore,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Certificate serial number.
	SerialNumber string `json:"serialNumber,omitempty"`
	// SHA-1 fingerprint.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
	// The subject alternative names associated with this certificate.
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CapacityBreakdownStruct - Storage stats breakdown.
// extends TypedObject
type CapacityBreakdownStruct struct {
	// Amount of space used for the active copy of the container.
	// base = 1024
	// units = B
	ActiveSpace float64 `json:"activeSpace,omitempty"`
	// Actual space used by the container.
	// units = B
	// base = 1024
	ActualSpace float64 `json:"actualSpace,omitempty"`
	// Amount of space used for snapshots from which VDBs have been
	// provisioned.
	// units = B
	// base = 1024
	DescendantSpace float64 `json:"descendantSpace,omitempty"`
	// Amount of space used by logs.
	// units = B
	// base = 1024
	LogSpace float64 `json:"logSpace,omitempty"`
	// Amount of space used for snapshots held by manual retention
	// settings.
	// base = 1024
	// units = B
	ManualSpace float64 `json:"manualSpace,omitempty"`
	// Amount of space used for snapshots held by policy settings.
	// base = 1024
	// units = B
	PolicySpace float64 `json:"policySpace,omitempty"`
	// Amount of space used by snapshots.
	// units = B
	// base = 1024
	SyncSpace float64 `json:"syncSpace,omitempty"`
	// Unvirtualized space used by the TimeFlow.
	// units = B
	// base = 1024
	TimeflowUnvirtualizedSpace float64 `json:"timeflowUnvirtualizedSpace,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Amount of space used for snapshots part of held space.
	// units = B
	// base = 1024
	UnownedSnapshotSpace float64 `json:"unownedSnapshotSpace,omitempty"`
	// Unvirtualized space used by the container.
	// units = B
	// base = 1024
	UnvirtualizedSpace float64 `json:"unvirtualizedSpace,omitempty"`
}

// CertificateConfigStruct - Configuration for the use of custom CA signed certificates.
// extends TypedObject
// cliVisibility = [SYSTEM DOMAIN]
type CertificateConfigStruct struct {
	// A critical fault will be thrown if any end entity certificates
	// will expire within this duration (in days). Must be less than the
	// warning threshold.
	// minimum = 1
	// update = optional
	CertificateExpirationCriticalThreshold *int `json:"certificateExpirationCriticalThreshold,omitempty"`
	// A warning fault will be thrown if any end entity certificates will
	// expire within this duration (in days). Must be greater than the
	// critical threshold.
	// minimum = 2
	// update = optional
	CertificateExpirationWarningThreshold *int `json:"certificateExpirationWarningThreshold,omitempty"`
	// Whether or not the engine will use the user managed DSP key for
	// client authentication for engine to host apps. This will use the
	// user managed DSP key on the host environment. Requires server auth
	// to be enabled, and a stack restart to take effect.
	// update = optional
	EnableUserManagedClientAuthForEngineToHostDsp *bool `json:"enableUserManagedClientAuthForEngineToHostDsp,omitempty"`
	// Whether or not the engine will use the user managed DSP key for
	// client authentication during network throughput tests. This will
	// use the user managed DSP key on the host environment or source
	// Delphix Engine, depending on if an engine to host, or engine to
	// engine test is executed, respectively. Requires server auth to be
	// enabled, and a stack restart to take effect. Note that if doing an
	// engine to engine test, this flag will need to be true for the
	// other Delphix Engine too.
	// update = optional
	EnableUserManagedClientAuthForNetworkThroughputTests *bool `json:"enableUserManagedClientAuthForNetworkThroughputTests,omitempty"`
	// Whether or not the engine will use the user managed DSP key for
	// client authentication during replication. This will use the user
	// managed DSP key on the source Delphix Engine. Requires server auth
	// to be enabled, and a stack restart to take effect. Note that this
	// flag will need to be true for the other Delphix Engine too.
	// update = optional
	EnableUserManagedClientAuthForReplication *bool `json:"enableUserManagedClientAuthForReplication,omitempty"`
	// Whether or not the engine will use the user managed DSP key for
	// server authentication for engine to host apps. Requires a stack
	// restart to take effect.
	// update = optional
	EnableUserManagedServerAuthForEngineToHostDsp *bool `json:"enableUserManagedServerAuthForEngineToHostDsp,omitempty"`
	// Whether or not the engine will use the user managed DSP key for
	// server authentication during network throughput tests. Requires a
	// stack restart to take effect. Note that if doing an engine to
	// engine test, this flag will need to be true for the other Delphix
	// Engine too.
	// update = optional
	EnableUserManagedServerAuthForNetworkThroughputTests *bool `json:"enableUserManagedServerAuthForNetworkThroughputTests,omitempty"`
	// Whether or not the engine will use the user managed DSP key for
	// server authentication during replication. Requires a stack restart
	// to take effect. Note that this flag will need to be true for the
	// other engine too.
	// update = optional
	EnableUserManagedServerAuthForReplication *bool `json:"enableUserManagedServerAuthForReplication,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether or not the engine will validate custom Windows Connector
	// certificates.
	// update = optional
	ValidateWindowsConnectorCertificate *bool `json:"validateWindowsConnectorCertificate,omitempty"`
}

// CertificateFetchParametersStruct - Parameters for fetching a certificate.
// extends TypedObject
type CertificateFetchParametersStruct struct {
	// Hostname or IP address.
	// format = host
	// required = true
	Host string `json:"host,omitempty"`
	// Port number.
	// minimum = 1
	// maximum = 65535
	// required = true
	Port *int `json:"port,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// CertificateSigningRequestStruct - Certificate signing request (CSR).
// extends UserObject
// cliVisibility = [SYSTEM DOMAIN]
type CertificateSigningRequestStruct struct {
	// The specific TLS service this CSR was generated for.
	EndEntity EndEntity `json:"endEntity,omitempty"`
	// The backing key pair and signature algorithm it will use.
	KeyPair KeyPair `json:"keyPair,omitempty"`
	// The Distinguished Name.
	// format = objectName
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// The CSR in PEM format.
	RequestInPem string `json:"requestInPem,omitempty"`
	// The subject alternative names associated with this certificate.
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CertificateSigningRequestCreateParametersStruct - The parameters used to create a certificate signing request (CSR).
// extends TypedObject
type CertificateSigningRequestCreateParametersStruct struct {
	// The Distinguished Name to use.
	// create = required
	Dname X500DistinguishedName `json:"dname,omitempty"`
	// The specific TLS service to generate the CSR for.
	// create = required
	// properties = map[type:map[default:EndEntityHttps]]
	EndEntity EndEntity `json:"endEntity,omitempty"`
	// Force replace the active keypair and certificate with this newly
	// generated one.
	// default = false
	// create = required
	ForceReplace *bool `json:"forceReplace,omitempty"`
	// The backing key pair and signature algorithm it will use.
	// create = required
	// properties = map[type:map[default:RsaKeyPair]]
	KeyPair KeyPair `json:"keyPair,omitempty"`
	// The subject alternative names associated with this certificate.
	// create = optional
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// CertificateUploadParametersStruct - The parameters for uploading a keystore which include the credentials
// needed to access it.
// extends TypedObject
type CertificateUploadParametersStruct struct {
	// The lowercase alias for the certificate and key pair to use.
	// required = true
	Alias string `json:"alias,omitempty"`
	// The password for the key pair. If not provided, then the storepass
	// is used.
	// format = password
	Keypass string `json:"keypass,omitempty"`
	// The format of the keystore.
	// enum = [JKS PKCS12]
	// default = JKS
	// required = true
	KeystoreType string `json:"keystoreType,omitempty"`
	// The password for the keystore.
	// format = password
	// required = true
	Storepass string `json:"storepass,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ChangeCurrentPasswordPolicyParametersStruct - The parameters to use as input when changing the currently active
// password policy.
// extends TypedObject
type ChangeCurrentPasswordPolicyParametersStruct struct {
	// Password policy reference.
	// format = objectReference
	// referenceTo = /delphix-password-policy.json
	// required = true
	CurrentPasswordPolicy string `json:"currentPasswordPolicy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Type of user.
	// required = true
	// enum = [SYSTEM DOMAIN]
	UserType string `json:"userType,omitempty"`
}

// ChecklistStruct - Generic checklist object.
// extends TypedObject
type ChecklistStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ChecklistItemStruct - Generic checklist item.
// extends TypedObject
type ChecklistItemStruct struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ChecklistItemDetailStruct - Fields to indicate detailed status for a specific checklist item.
// extends ChecklistItem
type ChecklistItemDetailStruct struct {
	// Description of this status item.
	Description string `json:"description,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CompatibilityCriteriaStruct - The compatibility criteria to use for selecting compatible
// repositories. Parameters with a value of null are not considered when
// selecting compatible repositories.
// extends TypedObject
type CompatibilityCriteriaStruct struct {
	// Selected repositories are installed on a host with this
	// architecture (32-bit or 64-bit).
	Architecture *int `json:"architecture,omitempty"`
	// Selected repositories are installed on this environment.
	// referenceTo = /delphix-source-environment.json
	// format = objectReference
	Environment string `json:"environment,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CompatibleRepositoriesResultStruct - Result of a compatible repositories request.
// extends TypedObject
type CompatibleRepositoriesResultStruct struct {
	// The criteria matched to select compatible repositories.
	// required = true
	Criteria *CompatibilityCriteriaStruct `json:"criteria,omitempty"`
	// The list of compatible repositories.
	// required = true
	Repositories []SourceRepository `json:"repositories,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ConfiguredStorageDeviceStruct - A storage device configured as usable storage.
// extends StorageDevice
type ConfiguredStorageDeviceStruct struct {
	// Boolean value indicating if this is a boot device.
	BootDevice *bool `json:"bootDevice,omitempty"`
	// True if the device is currently configured in the system.
	Configured *bool `json:"configured,omitempty"`
	// Amount of additional space that would be made available, if the
	// device is expanded.
	// base = 1024
	// units = B
	ExpandableSize float64 `json:"expandableSize,omitempty"`
	// Model ID of the device.
	Model string `json:"model,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Serial number of the device.
	Serial string `json:"serial,omitempty"`
	// Physical size of the device, in bytes.
	// base = 1024
	// units = B
	Size float64 `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Size of allocated space on the device.
	// base = 1024
	// units = B
	UsedSize float64 `json:"usedSize,omitempty"`
	// Vendor ID of the device.
	Vendor string `json:"vendor,omitempty"`
}

// ConnectorConnectivityStruct - Mechanism to test Connector connectivity of arbitrary hosts.
// extends TypedObject
type ConnectorConnectivityStruct struct {
	// Target host name or IP address.
	// format = host
	// required = true
	Address string `json:"address,omitempty"`
	// User credentials.
	// required = true
	Credentials *PasswordCredentialStruct `json:"credentials,omitempty"`
	// Connector port on remote server.
	// default = 9100
	// minimum = 1
	// maximum = 65535
	Port *int `json:"port,omitempty"`
	// Host to use as a proxy for credential validation.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	Proxy string `json:"proxy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User name.
	// required = true
	Username string `json:"username,omitempty"`
}

// ContainerUtilizationStruct - Represents the utilization of all containers during a particular
// period of time.
// extends TypedObject
// cliVisibility = []
type ContainerUtilizationStruct struct {
	// Reference to the container whose utilization we are describing.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// True if this container has been deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// True if the current user does not have access to this container.
	Hidden *bool `json:"hidden,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// A list of container utilization statistics corresponding to this
	// period of time, one for each sampling interval.
	Utilization []*ContainerUtilizationIntervalStruct `json:"utilization,omitempty"`
}

// ContainerUtilizationIntervalStruct - Represents a represents an interval of time with which container
// utilization statistics are associated.
// extends TypedObject
type ContainerUtilizationIntervalStruct struct {
	// The average throughput for the container throughout the duration
	// of the sampling interval, measured in kilobits per second.
	AverageThroughput float64 `json:"averageThroughput,omitempty"`
	// The start time of the interval.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CpuUtilDatapointStruct - An analytics datapoint generated by the CPU_UTIL statistic type.
// extends Datapoint
type CpuUtilDatapointStruct struct {
	// DTrace time in milliseconds (subset of time in kernel).
	Dtrace *int `json:"dtrace,omitempty"`
	// Idle time in milliseconds.
	Idle *int `json:"idle,omitempty"`
	// Kernel time in milliseconds.
	Kernel *int `json:"kernel,omitempty"`
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User time in milliseconds.
	User *int `json:"user,omitempty"`
}

// CpuUtilDatapointStreamStruct - A stream of datapoints from a CPU_UTIL analytics slice.
// extends DatapointStream
type CpuUtilDatapointStreamStruct struct {
	// Which CPU was utilized.
	Cpu *int `json:"cpu,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// CreateMaskingJobTransformationParametersStruct - Represents the parameters to create a transformation for a provided
// Masking Job.
// extends TypedObject
type CreateMaskingJobTransformationParametersStruct struct {
	// The container that will contain the masked data associated with
	// the newly created transformation; the "transformation container".
	// create = required
	Container Container `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// referenceTo = /delphix-source-repository.json
	// create = required
	// format = objectReference
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CredentialUpdateParametersStruct - Parameters to update a Delphix user's password.
// extends TypedObject
type CredentialUpdateParametersStruct struct {
	// The new credentials.
	// create = required
	NewCredential *PasswordCredentialStruct `json:"newCredential,omitempty"`
	// The old credentials.
	// create = optional
	OldCredential *PasswordCredentialStruct `json:"oldCredential,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CurrentConsumerCapacityDataStruct - Current data about a particular capacity consumer.
// extends BaseConsumerCapacityData
// cliVisibility = [DOMAIN]
type CurrentConsumerCapacityDataStruct struct {
	// Statistics for this consumer.
	Breakdown *CapacityBreakdownStruct `json:"breakdown,omitempty"`
	// Reference to the container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Reference to this container's group.
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
	// Name of this container's group.
	GroupName string `json:"groupName,omitempty"`
	// Name of the container.
	Name string `json:"name,omitempty"`
	// Container from which this TimeFlow was provisioned.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Parent string `json:"parent,omitempty"`
	// Internal unique identifier for this consumer.
	StorageContainer string `json:"storageContainer,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CurrentGroupCapacityDataStruct - Capacity data aggregated over a group.
// extends BaseGroupCapacityData
// cliVisibility = [DOMAIN]
type CurrentGroupCapacityDataStruct struct {
	// Which group these stats represent.
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdownStruct `json:"source,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdownStruct `json:"virtual,omitempty"`
}

// CurrentSystemCapacityDataStruct - Capacity data for the entire system.
// extends BaseSystemCapacityData
// cliVisibility = [DOMAIN]
type CurrentSystemCapacityDataStruct struct {
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdownStruct `json:"source,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Total storage space (used and unused).
	// units = B
	// base = 1024
	TotalSpace float64 `json:"totalSpace,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdownStruct `json:"virtual,omitempty"`
}

// DNSConfigStruct - DNS Client Configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type DNSConfigStruct struct {
	// One of more DNS domain names.
	// update = optional
	Domain []string `json:"domain,omitempty"`
	// List of DNS servers.
	// update = optional
	Servers []string `json:"servers,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DSPAutotunerParametersStruct - Network information required by the DSP autotuner.
// extends TypedObject
type DSPAutotunerParametersStruct struct {
	// Whether the test is testing connectivity to a Delphix Engine or a
	// remote host.
	// enum = [REMOTE_HOST DELPHIX_ENGINE]
	// create = required
	DestinationType string `json:"destinationType,omitempty"`
	// Whether the test is a transmit or receive test.
	// enum = [TRANSMIT RECEIVE]
	// default = TRANSMIT
	// create = optional
	Direction string `json:"direction,omitempty"`
	// Address, username and password used when running a test to another
	// Delphix Engine.
	// create = optional
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfoStruct `json:"remoteDelphixEngineInfo,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	RemoteHost string `json:"remoteHost,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DSPBestParametersStruct - DSP parameters, found by autotuner, that give the highest throughput
// for a certain target.
// extends PersistentObject
type DSPBestParametersStruct struct {
	// The size of the send and receive socket buffers in bytes used to
	// achieve maximum throughput.
	// units = B
	// base = 1024
	BufferSize *int `json:"bufferSize,omitempty"`
	// Whether the test is testing connectivity to a Delphix Engine or
	// remote host.
	// enum = [REMOTE_HOST DELPHIX_ENGINE]
	DestinationType string `json:"destinationType,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The number of connections used to achieve maximum throughput.
	NumConnections *int `json:"numConnections,omitempty"`
	// The queue depth used to achieve maximum throughput.
	QueueDepth *int `json:"queueDepth,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Information used when running a test to another Delphix Engine.
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfoStruct `json:"remoteDelphixEngineInfo,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// referenceTo = /delphix-host.json
	// format = objectReference
	RemoteHost string `json:"remoteHost,omitempty"`
	// The average throughput measured.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DSPOptionsStruct - Options commonly used by apps that use DSP.
// extends TypedObject
type DSPOptionsStruct struct {
	// Bandwidth limit (MB/s) for network traffic. A value of 0 means no
	// limit.
	// minimum = 0
	// create = optional
	// default = 0
	BandwidthLimit *int `json:"bandwidthLimit,omitempty"`
	// Compress the data stream over the network.
	// create = optional
	// default = false
	Compression *bool `json:"compression,omitempty"`
	// Encrypt the data stream over the network.
	// create = optional
	// default = false
	Encryption *bool `json:"encryption,omitempty"`
	// Total number of transport connections to use.
	// create = optional
	// default = 1
	// minimum = 1
	// maximum = 16
	NumConnections *int `json:"numConnections,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DataResultStruct - Result of a successful API call containing a reference to a
// downloadable resource.
// extends OKResult
type DataResultStruct struct {
	// Reference to the action associated with the operation, if any.
	// referenceTo = /delphix-action.json
	// format = objectReference
	Action string `json:"action,omitempty"`
	// Reference to the job started by the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-job.json
	Job string `json:"job,omitempty"`
	// Result of the operation. This will be specific to the API being
	// invoked.
	Result string `json:"result,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// DatabaseTemplateStruct - Parameter configuration for virtual databases.
// extends NamedUserObject
type DatabaseTemplateStruct struct {
	// User provided description for this template.
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A name/value map of string configuration parameters.
	// create = required
	// update = optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The type of the source associated with the template.
	// enum = [MySQLVirtualSource OracleVirtualSource MSSqlVirtualSource MSSqlLinkedSource PgSQLVirtualSource]
	// format = type
	// create = required
	// update = optional
	SourceType string `json:"sourceType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DatabaseTemplateConfigStruct - Static template configuration information for a given source type.
// extends TypedObject
type DatabaseTemplateConfigStruct struct {
	// A list of reserved parameters names that cannot be used in the
	// template.
	Reserved []string `json:"reserved,omitempty"`
	// The object type for sources to which this template is applicable.
	// format = type
	SourceType string `json:"sourceType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DatapointSetStruct - A set of datapoints from a particular analytics slice.
// extends TypedObject
type DatapointSetStruct struct {
	// The set of datapoint streams in the result.
	DatapointStreams []DatapointStream `json:"datapointStreams,omitempty"`
	// True if the number of datapoints to be included exceeded the
	// maximum allowable for a single datapoint set. As a result, not all
	// datapoints could be included in the datapointStreams array.
	Overflow *bool `json:"overflow,omitempty"`
	// The amount of time each datapoint spans.
	// units = sec
	Resolution *int `json:"resolution,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DeleteParametersStruct - The parameters to use as input to delete requests for MSSQL,
// PostgreSQL, AppData, ASE or MySQL.
// extends TypedObject
type DeleteParametersStruct struct {
	// Flag indicating whether to continue the operation upon failures.
	Force *bool `json:"force,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DelphixManagedBackupIngestionStrategyStruct - This link source has a Delphix managed backup ingestion strategy that
// will create copy-only full backups based on a policy and ingest that
// backup.
// extends IngestionStrategy
type DelphixManagedBackupIngestionStrategyStruct struct {
	// Specify whether the backups taken should be compressed or
	// uncompressed.
	// default = false
	// create = optional
	// update = optional
	CompressionEnabled *bool `json:"compressionEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DetachSourceParametersStruct - The parameters passed in for a database detach source operation.
// extends TypedObject
type DetachSourceParametersStruct struct {
	// Reference to the source to be removed. This must be a linked
	// source attached to the target database.
	// referenceTo = /delphix-source.json
	// required = true
	// format = objectReference
	Source string `json:"source,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DiagnosisResultStruct - Details from a diagnosis check that was run due to a failed operation.
// extends TypedObject
type DiagnosisResultStruct struct {
	// True if this was a check that did not pass.
	Failure *bool `json:"failure,omitempty"`
	// Localized message.
	Message string `json:"message,omitempty"`
	// Message code associated with the event.
	MessageCode string `json:"messageCode,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DiskOpsDatapointStreamStruct - A stream of datapoints from a DISK_OPS analytics slice.
// extends DatapointStream
type DiskOpsDatapointStreamStruct struct {
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// Device the I/Os were issued to.
	Device string `json:"device,omitempty"`
	// Whether the I/Os resulted in errors.
	Error *bool `json:"error,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DomainStruct - Represents the root container of all objects on the system.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type DomainStruct struct {
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DxFsIoQueueOpsDatapointStreamStruct - A stream of datapoints from a DxFS_IO_QUEUE_OPS analytics slice.
// extends DatapointStream
type DxFsIoQueueOpsDatapointStreamStruct struct {
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write claim free ioctl null]
	Op string `json:"op,omitempty"`
	// Priority of the I/O.
	// enum = [sync cache/agg asyncw asyncr resilver scrub ddt_prefetch]
	Priority string `json:"priority,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DxFsOpsDatapointStreamStruct - A stream of datapoints from a DxFS_OPS analytics slice.
// extends DatapointStream
type DxFsOpsDatapointStreamStruct struct {
	// Whether reads were cached.
	Cached *bool `json:"cached,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Path of the affected file.
	// format = unixpath
	Path string `json:"path,omitempty"`
	// Whether writes were synchronous.
	Sync *bool `json:"sync,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EcdsaKeyPairStruct - A key pair generated using the ECDSA algorithm.
// extends KeyPair
type EcdsaKeyPairStruct struct {
	// The size of each key to be generated.
	// default = 256
	// create = optional
	// minimum = 256
	// maximum = 571
	KeySize *int `json:"keySize,omitempty"`
	// The signature algorithm this key pair will use to sign
	// certificates and CSRs.
	// default = SHA256withECDSA
	// create = optional
	// enum = [SHA256withECDSA SHA384withECDSA SHA512withECDSA]
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EndEntityCertificateStruct - An End-Entity Public Key Certificate.
// extends Certificate
// cliVisibility = [SYSTEM DOMAIN]
type EndEntityCertificateStruct struct {
	// The specific TLS service this certificate is used for.
	EndEntity EndEntity `json:"endEntity,omitempty"`
	// Issuer of this certificate.
	IssuedByDN string `json:"issuedByDN,omitempty"`
	// A reference to the certificate that issued this certificate. Null
	// if this is a root CA.
	// format = objectReference
	// referenceTo = /delphix-ca-certificate.json
	Issuer string `json:"issuer,omitempty"`
	// MD5 fingerprint.
	Md5Fingerprint string `json:"md5Fingerprint,omitempty"`
	// The Distinguished Name of this certificate.
	// format = objectName
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// End of validity.
	// format = date
	NotAfter string `json:"notAfter,omitempty"`
	// Start of validity.
	// format = date
	NotBefore string `json:"notBefore,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Certificate serial number.
	SerialNumber string `json:"serialNumber,omitempty"`
	// SHA-1 fingerprint.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
	// The subject alternative names associated with this certificate.
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EndEntityCertificateReplaceChainParametersStruct - The parameters for replacing an end-entity certificate with a PEM
// certificate chain.
// extends EndEntityCertificateReplaceParameters
type EndEntityCertificateReplaceChainParametersStruct struct {
	// The PEM certificate chain.
	// required = true
	Chain *PemCertificateChainStruct `json:"chain,omitempty"`
	// The specific TLS service and system this certificate is used for.
	// required = true
	// properties = map[type:map[default:EndEntityHttps]]
	EndEntity EndEntity `json:"endEntity,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EndEntityCertificateReplaceKeystoreParametersStruct - The parameters for replacing an end-entity certificate with a
// certificate and key pair from a keystore.
// extends EndEntityCertificateReplaceParameters
type EndEntityCertificateReplaceKeystoreParametersStruct struct {
	// The specific TLS service and system this certificate is used for.
	// required = true
	// properties = map[type:map[default:EndEntityHttps]]
	EndEntity EndEntity `json:"endEntity,omitempty"`
	// The token in the FileUploadResult after uploading the keystore.
	// required = true
	Token string `json:"token,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EndEntityDspStruct - The end-entity certificate used for DSP by the Delphix Engine.
// extends EndEntity
type EndEntityDspStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EndEntityHttpsStruct - The end-entity certificate used for HTTPS by the Delphix Engine.
// extends EndEntity
type EndEntityHttpsStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EnumEqualConstraintStruct - Constraints placed on an enumeration axis of a particular analytics
// slice.
// extends EnumConstraint
type EnumEqualConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be equal to the specified value.
	// create = required
	Equals string `json:"equals,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EnvironmentUserStruct - The representation of an environment user object.
// extends UserObject
type EnvironmentUserStruct struct {
	// The credential for the environment user.
	// create = required
	// update = optional
	// properties = map[type:map[type:string description:Object type. required:true format:type default:PasswordCredential]]
	Credential Credential `json:"credential,omitempty"`
	// A reference to the associated environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = optional
	Environment string `json:"environment,omitempty"`
	// Group ID of the user.
	// create = optional
	// update = optional
	// minimum = 0
	// maximum = 4.294967295e+09
	GroupId *int `json:"groupId,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User ID of the user.
	// create = optional
	// update = optional
	// minimum = 0
	// maximum = 4.294967295e+09
	UserId *int `json:"userId,omitempty"`
}

// ErrorResultStruct - Result of a failed API call.
// extends CallResult
type ErrorResultStruct struct {
	// Specifics of the error that occurred during API call execution.
	Error *APIErrorStruct `json:"error,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// EventFilterStruct - An event filter that specifies which event types to match against.
// extends AlertFilter
type EventFilterStruct struct {
	// List of event types. Only alerts of the given event type are
	// included. Each event type is a string representing the event class
	// of the corresponding alerts. Wildcards are supported to include
	// classes of events.
	// update = optional
	// minItems = 1
	// create = optional
	EventTypes []string `json:"eventTypes,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ExternalBackupIngestionStrategyStruct - This link source has an external ingestion strategy that will pull in
// backups via ValidatedSync depending on the mode selected.
// extends IngestionStrategy
type ExternalBackupIngestionStrategyStruct struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Specifies the backup types ValidatedSync will use to synchronize
	// the dSource with the source database.
	// default = TRANSACTION_LOG
	// create = required
	// update = optional
	// enum = [TRANSACTION_LOG FULL_OR_DIFFERENTIAL FULL]
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
}

// FaultStruct - A representation of a fault, with associated user object.
// extends PersistentObject
type FaultStruct struct {
	// A suggested user action.
	Action string `json:"action,omitempty"`
	// A unique dot delimited identifier associated with the fault.
	BundleID string `json:"bundleID,omitempty"`
	// The date when the fault was diagnosed.
	// format = date
	DateDiagnosed string `json:"dateDiagnosed,omitempty"`
	// The date when the fault was resolved.
	// format = date
	DateResolved string `json:"dateResolved,omitempty"`
	// Full description of the fault.
	Description string `json:"description,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// A comment that describes the fault resolution.
	ResolutionComments string `json:"resolutionComments,omitempty"`
	// The automated response taken by the system.
	Response string `json:"response,omitempty"`
	// The severity of the fault event. This can either be CRITICAL or
	// WARNING.
	// enum = [CRITICAL WARNING]
	Severity string `json:"severity,omitempty"`
	// The status of the fault. This can be ACTIVE, RESOLVED or IGNORED.
	// enum = [ACTIVE RESOLVED IGNORED]
	Status string `json:"status,omitempty"`
	// The user-visible Delphix object that is faulted.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// The name of the faulted object at the time the fault was
	// diagnosed.
	TargetName string `json:"targetName,omitempty"`
	// The user-visible Delphix object that is faulted.
	// format = type
	TargetObjectType string `json:"targetObjectType,omitempty"`
	// Summary of the fault.
	Title string `json:"title,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// FaultEffectStruct - An error affecting a user object whose root cause is a fault. A fault
// effect can only be resolved by resolving the fault which is its root
// cause.
// extends PersistentObject
type FaultEffectStruct struct {
	// A suggested user action.
	Action string `json:"action,omitempty"`
	// A unique dot delimited identifier associated with the fault
	// effect.
	BundleID string `json:"bundleID,omitempty"`
	// The cause of the fault effect, in case there is a chain of fault
	// effects originating from the root cause which resulted in this
	// effect.
	// format = objectReference
	// referenceTo = /delphix-fault-effect.json
	CausedBy string `json:"causedBy,omitempty"`
	// The date when the root cause fault was diagnosed.
	// format = date
	DateDiagnosed string `json:"dateDiagnosed,omitempty"`
	// Full description of the fault effect.
	Description string `json:"description,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The automated response taken by the system.
	Response string `json:"response,omitempty"`
	// The root cause of this fault effect. Resolving the fault effect
	// can only occur by resolving its root cause.
	// format = objectReference
	// referenceTo = /delphix-fault.json
	RootCause string `json:"rootCause,omitempty"`
	// The severity of the fault effect. This can either be CRITICAL or
	// WARNING.
	// enum = [CRITICAL WARNING]
	Severity string `json:"severity,omitempty"`
	// The user-visible Delphix object that has a fault effect.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// The name of the user-visible Delphix object that has a fault
	// effect.
	TargetName string `json:"targetName,omitempty"`
	// Summary of the fault effect.
	Title string `json:"title,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// FaultResolveParametersStruct - The parameters to use as input when marking a fault as resolved.
// extends TypedObject
type FaultResolveParametersStruct struct {
	// The comments describing the steps taken to resolve a fault.
	// create = optional
	// default =
	Comments string `json:"comments,omitempty"`
	// Flag indicating whether to ignore this fault if it is detected on
	// the same object in the future.
	// default = false
	// create = optional
	Ignore *bool `json:"ignore,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// FeatureFlagParametersStruct - Feature Flags for the Delphix Engine.
// extends TypedObject
type FeatureFlagParametersStruct struct {
	// Name of the feature flag.
	// required = true
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// FileDownloadResultStruct - Result of a file download request.
// extends FileProcessingResult
type FileDownloadResultStruct struct {
	// Token to pass as parameter to identify the file.
	Token string `json:"token,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// URL to download from or upload to.
	Url string `json:"url,omitempty"`
}

// FileMappingParametersStruct - Input parameters to test file mapping rules.
// extends TypedObject
type FileMappingParametersStruct struct {
	// Database file mapping rules.
	// required = true
	MappingRules string `json:"mappingRules,omitempty"`
	// The list of TimeFlow point, bookmark, or semantic location to use
	// for source files to be mapped.
	// required = true
	// minItems = 1
	TimeflowPointParameters []TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// FileMappingResultStruct - Result of a file mapping request.
// extends TypedObject
type FileMappingResultStruct struct {
	// Mapped files.
	MappedFiles map[string]string `json:"mappedFiles,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// FileUploadResultStruct - Result of a file upload request.
// extends FileProcessingResult
type FileUploadResultStruct struct {
	// Token to pass as parameter to identify the file.
	Token string `json:"token,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// URL to download from or upload to.
	Url string `json:"url,omitempty"`
}

// FractionPlugParametersStruct - The parameters to use as input when transporting a transportable
// tablespace.
// extends TypedObject
type FractionPlugParametersStruct struct {
	// Optional prefix to add to schemas being moved into warehouse.
	// maxLength = 28
	// create = optional
	// update = optional
	// pattern = ^([A-Za-z][A-Za-z0-9_]+)|("[A-Za-z][A-Za-z0-9_]+")$
	SchemasPrefix string `json:"schemasPrefix,omitempty"`
	// Optional prefix to add to tablespaces being moved into warehouse.
	// create = optional
	// update = optional
	// pattern = ^([A-Za-z][A-Za-z0-9_]+)|("[A-Za-z][A-Za-z0-9_]+")$
	// maxLength = 28
	TablespacesPrefix string `json:"tablespacesPrefix,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// GetCurrentPasswordPolicyParametersStruct - The parameters to use as input when getting the currently active
// password policy.
// extends TypedObject
type GetCurrentPasswordPolicyParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Type of user.
	// required = true
	// enum = [SYSTEM DOMAIN]
	UserType string `json:"userType,omitempty"`
}

// GlobalLinkingSettingsStruct - System-wide linking settings.
// extends TypedObject
type GlobalLinkingSettingsStruct struct {
	// True if encrypted linking should be enabled by default on new
	// dSources.
	// update = optional
	EncryptedLinkingEnabledByDefault *bool `json:"encryptedLinkingEnabledByDefault,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// GroupStruct - Database group.
// extends NamedUserObject
type GroupStruct struct {
	// Optional description for the group.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// HistoricalConsumerCapacityDataStruct - Historical data about a particular capacity consumer.
// extends BaseConsumerCapacityData
// cliVisibility = [DOMAIN]
type HistoricalConsumerCapacityDataStruct struct {
	// Statistics for this consumer.
	Breakdown *CapacityBreakdownStruct `json:"breakdown,omitempty"`
	// Reference to the container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Reference to this container's group.
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
	// Name of this container's group.
	GroupName string `json:"groupName,omitempty"`
	// Name of the container.
	Name string `json:"name,omitempty"`
	// Container from which this TimeFlow was provisioned.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Parent string `json:"parent,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HistoricalGroupCapacityDataStruct - Historical capacity data aggregated over a group.
// extends BaseGroupCapacityData
// cliVisibility = [DOMAIN]
type HistoricalGroupCapacityDataStruct struct {
	// Which group these stats represent.
	// referenceTo = /delphix-group.json
	// format = objectReference
	Group string `json:"group,omitempty"`
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdownStruct `json:"source,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdownStruct `json:"virtual,omitempty"`
}

// HistoricalSystemCapacityDataStruct - Capacity data for the entire system.
// extends BaseSystemCapacityData
// cliVisibility = [DOMAIN]
type HistoricalSystemCapacityDataStruct struct {
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdownStruct `json:"source,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Total storage space (used and unused).
	// units = B
	// base = 1024
	TotalSpace float64 `json:"totalSpace,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdownStruct `json:"virtual,omitempty"`
}

// HostConfigurationStruct - The representation of the host configuration properties.
// extends TypedObject
type HostConfigurationStruct struct {
	// Indicates whether the host configuration properties were
	// discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The timestamp when the host was last refreshed.
	LastRefreshed string `json:"lastRefreshed,omitempty"`
	// The timestamp when the host was last updated.
	LastUpdated string `json:"lastUpdated,omitempty"`
	// The host machine information.
	Machine *HostMachineStruct `json:"machine,omitempty"`
	// The host operating system information.
	OperatingSystem *HostOSStruct `json:"operatingSystem,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HostEnvironmentCreateParametersStruct - The parameters used for the host environment create operation.
// extends SourceEnvironmentCreateParameters
type HostEnvironmentCreateParametersStruct struct {
	// The host environment.
	// create = required
	HostEnvironment HostEnvironment `json:"hostEnvironment,omitempty"`
	// The host parameters used to add a host.
	// create = required
	HostParameters HostCreateParameters `json:"hostParameters,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to the created source
	// environment.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The primary user associated with the environment.
	// create = required
	PrimaryUser *EnvironmentUserStruct `json:"primaryUser,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HostMachineStruct - The representation of the host machine.
// extends TypedObject
type HostMachineStruct struct {
	// The amount of RAM on the host machine.
	// units = B
	// base = 1024
	MemorySize float64 `json:"memorySize,omitempty"`
	// The platform for the host machine.
	Platform string `json:"platform,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HostOSStruct - The operating system information for the host.
// extends TypedObject
type HostOSStruct struct {
	// The OS distribution.
	Distribution string `json:"distribution,omitempty"`
	// The OS kernel.
	Kernel string `json:"kernel,omitempty"`
	// The OS name.
	Name string `json:"name,omitempty"`
	// The OS release.
	Release string `json:"release,omitempty"`
	// The OS timezone.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The OS version.
	Version string `json:"version,omitempty"`
}

// HostPrivilegeElevationProfileStruct - Profile for elevating user privileges on a host.
// extends PersistentObject
// cliVisibility = [DOMAIN]
type HostPrivilegeElevationProfileStruct struct {
	// True if this is the default privilege elevation profile for new
	// environments.
	// default = false
	// update = readonly
	IsDefault *bool `json:"isDefault,omitempty"`
	// The privilege elevation profile name.
	// update = optional
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Privilege elevation profile version.
	// update = optional
	// create = required
	Version string `json:"version,omitempty"`
}

// HostPrivilegeElevationProfileScriptStruct - A script that is part of a profile for elevating user privileges on a
// host.
// extends PersistentObject
// cliVisibility = [DOMAIN]
type HostPrivilegeElevationProfileScriptStruct struct {
	// The contents of the privilege elevation profile script.
	// create = required
	// update = optional
	Contents string `json:"contents,omitempty"`
	// The privilege elevation profile script name.
	// update = optional
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The privilege elevation profile to which this script belongs.
	// format = objectReference
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	// create = required
	// update = optional
	Profile string `json:"profile,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// HostPrivilegeElevationSettingsStruct - Settings for elevating user privileges on a host.
// extends TypedObject
// cliVisibility = [DOMAIN]
type HostPrivilegeElevationSettingsStruct struct {
	// The default privilege elevation profile for new environments.
	// format = objectReference
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	// update = required
	DefaultProfile string `json:"defaultProfile,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HostRuntimeStruct - Runtime, non-persistent properties for a host machine.
// extends TypedObject
type HostRuntimeStruct struct {
	// True if the host is up and a connection can be established.
	Available *bool `json:"available,omitempty"`
	// The time that the 'available' propery was last checked.
	// format = date
	AvailableTimestamp string `json:"availableTimestamp,omitempty"`
	// The reason why the host is not available.
	NotAvailableReason string `json:"notAvailableReason,omitempty"`
	// Traceroute network hops from host to Delphix Engine.
	TraceRouteInfo *TracerouteInfoStruct `json:"traceRouteInfo,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HttpConnectorConfigStruct - Configuration for the HTTP and HTTPS connector of this application.
// extends TypedObject
// cliVisibility = [SYSTEM]
type HttpConnectorConfigStruct struct {
	// Controls the HTTP(s) protocol configuration of this appliance.
	// enum = [HTTP_ONLY HTTPS_ONLY HTTP_REDIRECT BOTH]
	// default = BOTH
	// update = optional
	HttpMode string `json:"httpMode,omitempty"`
	// The HTTP port for the Delphix web UI.
	// create = optional
	// default = 80
	// minimum = 1
	// maximum = 65535
	// update = optional
	HttpPort *int `json:"httpPort,omitempty"`
	// The HTTPS port for the Delphix web UI.
	// update = optional
	// create = optional
	// default = 443
	// minimum = 1
	// maximum = 65535
	HttpsPort *int `json:"httpsPort,omitempty"`
	// Version of TLS (transport layer security) enabled on this
	// appliance.
	// minItems = 1
	// update = optional
	TlsVersions []string `json:"tlsVersions,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// IScsiOpsDatapointStreamStruct - A stream of datapoints from a iSCSI_OPS analytics slice.
// extends DatapointStream
type IScsiOpsDatapointStreamStruct struct {
	// Address of the client.
	// format = host
	Client string `json:"client,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// IntegerEqualConstraintStruct - Constraint placed on a numerical axis of a particular analytics slice.
// extends IntegerConstraint
type IntegerEqualConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must equal this value.
	// create = required
	Equals *int `json:"equals,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// IntegerGreaterThanConstraintStruct - Constraint placed on a numerical axis of a particular analytics slice.
// extends IntegerConstraint
type IntegerGreaterThanConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be greater than this value.
	// create = required
	GreaterThan *int `json:"greaterThan,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// IntegerLessThanConstraintStruct - Constraint placed on a numerical axis of a particular analytics slice.
// extends IntegerConstraint
type IntegerLessThanConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be less than than this value.
	// create = required
	LessThan *int `json:"lessThan,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// InterfaceAddressStruct - IP address assigned to a network interface.
// extends TypedObject
type InterfaceAddressStruct struct {
	// The address in Classless Inter-Domain Routing (CIDR) notation.
	// format = cidrAddress
	// create = optional
	// update = optional
	Address string `json:"address,omitempty"`
	// The type of address (STATIC or DHCP).
	// enum = [STATIC DHCP]
	// default = STATIC
	// create = required
	// update = optional
	AddressType string `json:"addressType,omitempty"`
	// True if this address should accept incoming SSH connections.
	// default = true
	// create = optional
	// update = optional
	EnableSSH *bool `json:"enableSSH,omitempty"`
	// True if the API session is established over this address. This
	// property helps a client make informative decisions about which
	// address should not be modified without affecting the session over
	// which it is connected.
	SessionInUse *bool `json:"sessionInUse,omitempty"`
	// The state of the address.
	// enum = [OK TENTATIVE DUPLICATE INACCESSIBLE]
	State string `json:"state,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// IoOpsDatapointStruct - An analytics datapoint generated by the DISK_OPS, DxFS_OPS,
// DxFS_IO_QUEUE_OPS, iSCSI_OPS, NFS_OPS, or VFS_OPS statistic types.
// extends Datapoint
type IoOpsDatapointStruct struct {
	// Average I/O latency in nanoseconds.
	AvgLatency *int `json:"avgLatency,omitempty"`
	// Number of I/O operations.
	Count *int `json:"count,omitempty"`
	// I/O latencies in nanoseconds.
	Latency map[string]string `json:"latency,omitempty"`
	// I/O sizes in bytes.
	Size map[string]string `json:"size,omitempty"`
	// I/O throughput in bytes.
	// units = B
	// base = 1024
	Throughput *int `json:"throughput,omitempty"`
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JDBCConnectivityStruct - Mechanism to test JDBC connectivity of arbitrary databases.
// extends TypedObject
type JDBCConnectivityStruct struct {
	// Database password.
	// format = password
	// required = true
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// JDBC connection URL.
	// format = oracleJDBCConnectionString
	// required = true
	Url string `json:"url,omitempty"`
	// Database username.
	// required = true
	User string `json:"user,omitempty"`
}

// JSBookmarkStruct - A named entity that represents a point in time for all of the data
// sources in a data layout.
// extends NamedUserObject
type JSBookmarkStruct struct {
	// Denotes whether or not this bookmark was created on a data
	// container or a data template.
	// enum = [DATA_CONTAINER DATA_TEMPLATE]
	BookmarkType string `json:"bookmarkType,omitempty"`
	// A reference to the branch this bookmark applies to.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// create = required
	Branch string `json:"branch,omitempty"`
	// The number of times this bookmark has been checked out. This means
	// it was used as input to a RESTORE, CREATE_BRANCH, or RESET
	// operation.
	CheckoutCount *int `json:"checkoutCount,omitempty"`
	// The data container this bookmark was created on. This will be null
	// if the bookmark was created on a data template.
	// format = objectReference
	// referenceTo = /delphix-js-data-container.json
	Container string `json:"container,omitempty"`
	// The name of the data container this bookmark was created on. This
	// will be null if the bookmark was created on a data template.
	ContainerName string `json:"containerName,omitempty"`
	// The time at which the bookmark was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Description of this bookmark.
	// create = optional
	// update = optional
	// maxLength = 4096
	Description string `json:"description,omitempty"`
	// A policy will automatically delete this bookmark at this time. If
	// the value is null, then the bookmark will be kept until manually
	// deleted.
	// format = date
	// create = optional
	// update = optional
	Expiration string `json:"expiration,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// True if this bookmark is shared.
	// create = optional
	Shared *bool `json:"shared,omitempty"`
	// A set of user-defined labels for this bookmark.
	// create = optional
	// update = optional
	Tags []string `json:"tags,omitempty"`
	// The data template this bookmark was created on or the template of
	// the data container this bookmark was created on.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	Template string `json:"template,omitempty"`
	// The name of the data template this bookmark was created on or the
	// template of the data container this bookmark was created on.
	TemplateName string `json:"templateName,omitempty"`
	// The timestamp for the data that the bookmark refers to.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this bookmark is usable as input to a data operation
	// (e.g., CREATE_BRANCH or RESTORE).
	Usable *bool `json:"usable,omitempty"`
}

// JSBookmarkCheckoutCountStruct - The number of times a bookmark has been checked out. This means it was
// used as input to a RESTORE, CREATE_BRANCH, or RESET operation. The
// bookmark checkout count is kept separately on replicated templates.
// extends PersistentObject
type JSBookmarkCheckoutCountStruct struct {
	// The bookmark that this checkout count is associated with.
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
	// The number of times the bookmark has been checked out. This means
	// it was used as input to a RESTORE, CREATE_BRANCH, or RESET
	// operation. This should not be replicated.
	CheckoutCount *int `json:"checkoutCount,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSBookmarkCreateParametersStruct - The parameters used to create a Self-Service bookmark.
// extends TypedObject
type JSBookmarkCreateParametersStruct struct {
	// The Self-Service bookmark object.
	// required = true
	Bookmark *JSBookmarkStruct `json:"bookmark,omitempty"`
	// The Self-Service data timeline point at which the bookmark will be
	// created.
	// required = true
	TimelinePointParameters JSTimelinePointTimeParameters `json:"timelinePointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSBookmarkDataParentStruct - The bookmark data parent of a RESTORE or CREATE_BRANCH operation.
// extends JSDataParent
type JSBookmarkDataParentStruct struct {
	// The bookmark that this operation's data came from. This will be
	// null if the bookmark has been deleted.
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
	// This will always contain the name of the bookmark, even if it has
	// been deleted.
	// maxLength = 256
	BookmarkName string `json:"bookmarkName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSBookmarkTagUsageDataStruct - The space usage information for a Self-Service bookmark tag.
// extends TypedObject
type JSBookmarkTagUsageDataStruct struct {
	// The amount of space that will be freed if bookmarks with this tag
	// are deleted.
	BookmarkTag string `json:"bookmarkTag,omitempty"`
	// The total amount of space referenced by bookmarks with this tag.
	// This is the sum of the bookmarks' unique, shared, and
	// externallyReferenced space.
	// base = 1024
	// units = B
	Referenced float64 `json:"referenced,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The space that is being consumed by the set of bookmarks with the
	// given tag. This represents the minimum amount of space that will
	// be freed if all of the bookmarks are deleted.
	// base = 1024
	// units = B
	Unique float64 `json:"unique,omitempty"`
}

// JSBookmarkUsageDataStruct - The space usage information for a Self-Service bookmark.
// extends TypedObject
type JSBookmarkUsageDataStruct struct {
	// The Self-Service bookmark that this usage information is for.
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
	// The data layout that this bookmark belongs to.
	DataLayout string `json:"dataLayout,omitempty"`
	// The amount of space referenced by this bookmark that cannot be
	// freed up by deleting this bookmark because it is also being
	// referenced outside of Self-Service (e.g. by retention policy).
	// base = 1024
	// units = B
	ExternallyReferenced float64 `json:"externallyReferenced,omitempty"`
	// The amount of space referenced by this bookmark that cannot be
	// freed up by deleting this bookmark because it is also referenced
	// by neighboring bookmarks or branches that have been created or
	// restored from this bookmark.
	// base = 1024
	// units = B
	Shared float64 `json:"shared,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The amount of space that will be freed if this bookmark is
	// deleted.
	// base = 1024
	// units = B
	Unique float64 `json:"unique,omitempty"`
}

// JSBranchStruct - A branch represents a distinct timeline for data sources in a data
// layout.
// extends NamedUserObject
type JSBranchStruct struct {
	// A reference to the data layout this branch was created on.
	// referenceTo = /delphix-js-data-layout.json
	// format = objectReference
	DataLayout string `json:"dataLayout,omitempty"`
	// The first JSOperation on this branch by data time.
	// referenceTo = /delphix-js-operation.json
	// format = objectReference
	FirstOperation string `json:"firstOperation,omitempty"`
	// The last JSOperation on this branch by data time.
	// referenceTo = /delphix-js-operation.json
	// format = objectReference
	LastOperation string `json:"lastOperation,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// JSBranchCreateParametersStruct - The parameters used to create a Self-Service branch.
// extends TypedObject
type JSBranchCreateParametersStruct struct {
	// A reference to the data container to create this branch on.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-js-data-container.json
	DataContainer string `json:"dataContainer,omitempty"`
	// The name of the branch.
	// required = true
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The Self-Service data timeline point from which the branch will be
	// created.
	// required = true
	TimelinePointParameters JSTimelinePointParameters `json:"timelinePointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSBranchUsageDataStruct - The space usage information for a Self-Service branch.
// extends TypedObject
type JSBranchUsageDataStruct struct {
	// The Self-Service branch that this usage information is for.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
	// The name of the data container that this branch resides on.
	DataContainer string `json:"dataContainer,omitempty"`
	// The amount of space that cannot be freed on the parent data
	// template (or sibling data containers) because it is also being
	// referenced by this branch due to restore or create branch
	// operations.
	// base = 1024
	// units = B
	SharedOthers float64 `json:"sharedOthers,omitempty"`
	// The amount of space that cannot be freed up on this branch because
	// it is also being referenced by sibling data containers due to
	// restore or create branch operations.
	// base = 1024
	// units = B
	SharedSelf float64 `json:"sharedSelf,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The amount of space that will be freed if this branch is deleted.
	// base = 1024
	// units = B
	Unique float64 `json:"unique,omitempty"`
}

// JSConfigStruct - Self-Service configuration.
// extends TypedObject
type JSConfigStruct struct {
	// Default expiration for bookmarks created through the GUI, in days.
	// If value is 0, bookmarks will default to no expiration.
	// update = optional
	// minimum = 0
	DefaultBookmarkExpiration *int `json:"defaultBookmarkExpiration,omitempty"`
	// The number of times to retry failed sources during Self-Service
	// data operations.
	// update = optional
	// minimum = 0
	RetryAttempts *int `json:"retryAttempts,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// JSContainerUsageDataStruct - The space usage information for a data container.
// extends TypedObject
type JSContainerUsageDataStruct struct {
	// The data container that this usage information is for.
	// format = objectReference
	// referenceTo = /delphix-js-data-container.json
	DataContainer string `json:"dataContainer,omitempty"`
	// The amount of space that cannot be freed on the parent data
	// template (or sibling data containers) because it is also being
	// referenced by this data container due to restore or create branch
	// operations.
	// base = 1024
	// units = B
	SharedOthers float64 `json:"sharedOthers,omitempty"`
	// The amount of space that cannot be freed on this data container
	// because it is also being referenced by sibling data containers due
	// to restore or create branch operations.
	// base = 1024
	// units = B
	SharedSelf float64 `json:"sharedSelf,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The amount of space that will be freed if this data container is
	// deleted or purged. This assumes that the data container is deleted
	// along with underlying data sources.
	// units = B
	// base = 1024
	Unique float64 `json:"unique,omitempty"`
	// The amount of space that would be consumed by the data in this
	// container without Delphix.
	// base = 1024
	// units = B
	Unvirtualized float64 `json:"unvirtualized,omitempty"`
}

// JSDailyOperationDurationStruct - Information about the durations of a specific operation type for a
// data container over the past week.
// extends JSUsageData
type JSDailyOperationDurationStruct struct {
	// The average duration in seconds of running the specified operation
	// in the past day.
	// units = sec
	DailyAverageDuration *int `json:"dailyAverageDuration,omitempty"`
	// The number of times the specified operation was run in the past
	// day.
	DailyCount *int `json:"dailyCount,omitempty"`
	// The maximum duration in seconds of running the specified operation
	// in the past day.
	// units = sec
	DailyMaxDuration *int `json:"dailyMaxDuration,omitempty"`
	// The minimum duration in seconds of running the specified operation
	// in the past day.
	// units = sec
	DailyMinDuration *int `json:"dailyMinDuration,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The operation performed.
	// enum = [REFRESH RESET CREATE_BRANCH RESTORE UNDO]
	Operation string `json:"operation,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// The date at the beginning of the time period this datapoint
	// corresponds to. The time period itself varies between datapoint
	// types.
	// format = date
	StartDate string `json:"startDate,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The object the usage data is centered around.
	// format = objectReference
	// referenceTo = /delphix-named-user-object.json
	UsageObject string `json:"usageObject,omitempty"`
}

// JSDataChildStruct - A branch with data from a specific bookmark or PIT (point in time).
// extends TypedObject
type JSDataChildStruct struct {
	// Reference to the branch. This will be null if the user does not
	// have permission to see it.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
	// The name of the branch.
	// maxLength = 256
	BranchName string `json:"branchName,omitempty"`
	// The name of the container.
	// maxLength = 256
	ContainerName string `json:"containerName,omitempty"`
	// The operation performed.
	// enum = [REFRESH RESTORE CREATE_BRANCH RESET]
	Operation string `json:"operation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the owner of the branch. This will be null if
	// there is no owner.
	// maxLength = 256
	UserName string `json:"userName,omitempty"`
}

// JSDataContainerStruct - A container represents a data template provisioned for a specific
// user.
// extends JSDataLayout
type JSDataContainerStruct struct {
	// The active branch of the data layout.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	ActiveBranch string `json:"activeBranch,omitempty"`
	// The first JSOperation on this data layout by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	FirstOperation string `json:"firstOperation,omitempty"`
	// The last JSOperation on this data layout by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	LastOperation string `json:"lastOperation,omitempty"`
	// Timestamp of the last update to the application.
	// format = date
	LastUpdated string `json:"lastUpdated,omitempty"`
	// Name of the user who locked this data container.
	// update = readonly
	LockUserName string `json:"lockUserName,omitempty"`
	// The reference to the user who locked this data container.
	// format = objectReference
	// referenceTo = /delphix-user.json
	// update = readonly
	LockUserReference string `json:"lockUserReference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Notes for this data layout.
	// update = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
	// The number of operations performed on this data container.
	OperationCount *int `json:"operationCount,omitempty"`
	// For backward compatibility. The owner of the data container.
	// update = optional
	// internal = true
	// format = objectReference
	// referenceTo = /delphix-user.json
	Owner string `json:"owner,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// update = optional
	Properties map[string]string `json:"properties,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The state of the data container.
	// enum = [ONLINE OFFLINE INCONSISTENT]
	State string `json:"state,omitempty"`
	// The data template that this data container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	Template string `json:"template,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerActiveBranchParametersStruct - Input parameters for the API that given a point in time, returns the
// active branch of the data container.
// extends TypedObject
type JSDataContainerActiveBranchParametersStruct struct {
	// The time that will be used to find which branch was active in the
	// data layout.
	// required = true
	// format = date
	Time string `json:"time,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerCreateWithRefreshParametersStruct - The parameters used to create a data container when refreshing data
// sources.
// extends JSDataContainerCreateParameters
type JSDataContainerCreateWithRefreshParametersStruct struct {
	// The set of data sources that belong to this data layout.
	// required = true
	DataSources []*JSDataSourceCreateParametersStruct `json:"dataSources,omitempty"`
	// The name of the data layout.
	// maxLength = 256
	// required = true
	Name string `json:"name,omitempty"`
	// A description of this data layout to define what it is used for.
	// create = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
	// A reference to the list of users that own this data container.
	// create = optional
	Owners []string `json:"owners,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
	// A reference to the template that this data container is
	// provisioned from.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	// required = true
	Template string `json:"template,omitempty"`
	// Create the data container with initial data specified by this
	// Self-Service timeline point.
	// required = true
	TimelinePointParameters JSTimelinePointParameters `json:"timelinePointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerCreateWithoutRefreshParametersStruct - The parameters used to create a data container when not refreshing
// data sources.
// extends JSDataContainerCreateParameters
type JSDataContainerCreateWithoutRefreshParametersStruct struct {
	// The set of data sources that belong to this data layout.
	// required = true
	DataSources []*JSDataSourceCreateParametersStruct `json:"dataSources,omitempty"`
	// The name of the data layout.
	// maxLength = 256
	// required = true
	Name string `json:"name,omitempty"`
	// A description of this data layout to define what it is used for.
	// create = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
	// A reference to the list of users that own this data container.
	// create = optional
	Owners []string `json:"owners,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
	// A reference to the template that this data container is
	// provisioned from.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	// required = true
	Template string `json:"template,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerDeleteParametersStruct - The parameters used to delete a data container.
// extends TypedObject
type JSDataContainerDeleteParametersStruct struct {
	// If this value is true, then delete the underlying data from all
	// data sources.
	// required = true
	// default = true
	DeleteDataSources *bool `json:"deleteDataSources,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerLockParametersStruct - The parameters used to lock a data container.
// extends TypedObject
type JSDataContainerLockParametersStruct struct {
	// A reference to the user object who locks the container.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-user.json
	LockUser string `json:"lockUser,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerModifyOwnerParametersStruct - Input parameters for addOwner or removeOwner for a data container.
// extends TypedObject
type JSDataContainerModifyOwnerParametersStruct struct {
	// A reference to the user object for whom to add or remove
	// authorizations.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-user.json
	Owner string `json:"owner,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// JSDataContainerRefreshParametersStruct - The parameters used to refresh a data container.
// extends TypedObject
type JSDataContainerRefreshParametersStruct struct {
	// If this value is true, then do the operation without taking a
	// snapshot.
	// required = true
	// default = false
	ForceOption *bool `json:"forceOption,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerResetParametersStruct - The parameters used to reset a data container..
// extends TypedObject
type JSDataContainerResetParametersStruct struct {
	// If this value is true, then do the operation without taking a
	// snapshot.
	// required = true
	// default = false
	ForceOption *bool `json:"forceOption,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerRestoreParametersStruct - The parameters used to restore a data container.
// extends TypedObject
type JSDataContainerRestoreParametersStruct struct {
	// If this value is true, then do the operation without taking a
	// snapshot.
	// required = true
	// default = false
	ForceOption *bool `json:"forceOption,omitempty"`
	// The Self-Service data timeline point to restore from.
	// required = true
	TimelinePointParameters JSTimelinePointParameters `json:"timelinePointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataContainerUndoParametersStruct - The parameters used to undo an operation on a data container.
// extends TypedObject
type JSDataContainerUndoParametersStruct struct {
	// The operation to undo. This is only valid for RESET, RESTORE,
	// UNDO, and REFRESH operations.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	Operation string `json:"operation,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// JSDataSourceStruct - The data source used for Self-Service data layouts.
// extends NamedUserObject
type JSDataSourceStruct struct {
	// A reference to the underlying container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A reference to the Self-Service data layout to which this source
	// belongs.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	DataLayout string `json:"dataLayout,omitempty"`
	// A description of this data source.
	// create = optional
	// update = optional
	// maxLength = 4096
	Description string `json:"description,omitempty"`
	// Flag indicating whether the source is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Flag indicating whether the source is masked.
	Masked *bool `json:"masked,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Dictates order of operations on data sources. Operations can be
	// performed in parallel for all sources or sequentially. Below are
	// possible valid and invalid orderings given an example data
	// template with 3 sources (A, B, and C).<br>Valid:<br>A B C<br>1 1 1
	// (parallel)<br>1 2 3 (sequential)<br>Invalid:<br>A B C<br>2 2
	// 2<br>0 1 2<br>2 3 4<br>1 2 2<br>In the sequential case the data
	// source with priority 1 is the first to be started and the last to
	// be stopped. This value is set on creation of the template's data
	// sources and copied to the data container's data sources.
	// create = optional
	// minimum = 1
	// default = 1
	Priority *int `json:"priority,omitempty"`
	// Key/value pairs used to specify attributes for this data source.
	// create = optional
	// update = optional
	Properties map[string]string `json:"properties,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this data source.
	Runtime SourceConnectionInfo `json:"runtime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataSourceCreateParametersStruct - The parameters used to create the Self-Service data sources.
// extends TypedObject
type JSDataSourceCreateParametersStruct struct {
	// A reference to the underlying container object.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Key/value pairs used to specify attributes for this data source.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
	// The Self-Service data source object.
	// required = true
	Source *JSDataSourceStruct `json:"source,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataTemplateStruct - A data template is a collection of data sources and configuration
// representing a data layout that can be provisioned to Self-Service
// users.
// extends JSDataLayout
type JSDataTemplateStruct struct {
	// The active branch of the data layout.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	ActiveBranch string `json:"activeBranch,omitempty"`
	// A client should consider warning the user before performing an
	// operation which may take a long time, if this is true.
	// create = optional
	// update = optional
	// default = true
	ConfirmTimeConsumingOperations *bool `json:"confirmTimeConsumingOperations,omitempty"`
	// The first JSOperation on this data layout by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	FirstOperation string `json:"firstOperation,omitempty"`
	// The last JSOperation on this data layout by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	LastOperation string `json:"lastOperation,omitempty"`
	// Timestamp of the last update to the application.
	// format = date
	LastUpdated string `json:"lastUpdated,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Notes for this data layout.
	// update = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// update = optional
	Properties map[string]string `json:"properties,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSDataTemplateCreateParametersStruct - The parameters used to create a data template.
// extends JSDataLayoutCreateParameters
type JSDataTemplateCreateParametersStruct struct {
	// The set of data sources that belong to this data layout.
	// required = true
	DataSources []*JSDataSourceCreateParametersStruct `json:"dataSources,omitempty"`
	// The name of the data layout.
	// maxLength = 256
	// required = true
	Name string `json:"name,omitempty"`
	// A description of this data layout to define what it is used for.
	// maxLength = 4096
	// create = optional
	Notes string `json:"notes,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSOperationStruct - An operation that occurred on a Self-Service data layout.
// extends NamedUserObject
type JSOperationStruct struct {
	// The bookmark that was created.
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
	// The branch that this operation was performed on.
	// referenceTo = /delphix-js-branch.json
	// format = objectReference
	Branch string `json:"branch,omitempty"`
	// The data layout that this operation was performed on.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	DataLayout string `json:"dataLayout,omitempty"`
	// The data parent of the operation.
	DataParent JSDataParent `json:"dataParent,omitempty"`
	// The time that the data represented by this operation was active.
	// It will be null if the operation is in progress.
	// format = date
	DataTime string `json:"dataTime,omitempty"`
	// Plain text description of the operation.
	Description string `json:"description,omitempty"`
	// The time the operation finished. It will be null if the operation
	// is in progress.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Was this operation perfomed with the force option, which means no
	// pre-operation snapshot was taken.
	ForceOption *bool `json:"forceOption,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The operation performed.
	// enum = [REFRESH RESET CREATE_BRANCH DELETE_BRANCH CREATE_BOOKMARK DELETE_BOOKMARK ENABLE DISABLE ACTIVATE DEACTIVATE RECOVER RESTORE UNDO LOCK UNLOCK HISTORY]
	Operation string `json:"operation,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The root action that spawned this operation.
	// format = objectReference
	// referenceTo = /delphix-action.json
	RootAction string `json:"rootAction,omitempty"`
	// The time the operation started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The user who performed the operation.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
}

// JSOperationEndpointStruct - The first and last JSOperation for a given data layout or branch.
// extends TypedObject
type JSOperationEndpointStruct struct {
	// The first JSOperation.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	First string `json:"first,omitempty"`
	// The last JSOperation.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	Last string `json:"last,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSOperationEndpointBranchParametersStruct - The branch to fetch the first and last event from.
// extends JSOperationEndpointParameters
type JSOperationEndpointBranchParametersStruct struct {
	// The branch to search.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// create = required
	Branch string `json:"branch,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSOperationEndpointDataLayoutParametersStruct - The data layout to fetch the first and last events from.
// extends JSOperationEndpointParameters
type JSOperationEndpointDataLayoutParametersStruct struct {
	// The data layout to search.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	// create = required
	DataLayout string `json:"dataLayout,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSSourceDataTimestampStruct - The association between a Self-Service data source and a point in time
// that's provisionable.
// extends TypedObject
type JSSourceDataTimestampStruct struct {
	// A reference to the Self-Service branch.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
	// The name of the Self-Service data source.
	Name string `json:"name,omitempty"`
	// The priority of the Self-Service data source.
	Priority *int `json:"priority,omitempty"`
	// A reference to the Self-Service data source.
	// referenceTo = /delphix-js-data-source.json
	// format = objectReference
	Source string `json:"source,omitempty"`
	// The point in the source's dataset time which is provisionable.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSSourceDataTimestampParametersStruct - Input parameters for the API that given a point in time, returns the
// timestamps of the latest provisionable points, before the specified
// time and from the given branch, for each data source in the branch's
// data layout.
// extends TypedObject
type JSSourceDataTimestampParametersStruct struct {
	// A reference to the Self-Service branch.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// required = true
	Branch string `json:"branch,omitempty"`
	// The time that will be used to find provisionable timestamps for
	// the sources in the branch's data layout.
	// required = true
	// format = date
	Time string `json:"time,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSTemplateUsageDataStruct - The space usage information for a data template.
// extends TypedObject
type JSTemplateUsageDataStruct struct {
	// The amount of space consumed by the bookmarks on this data
	// template. This is the space that will be freed up if all bookmarks
	// on the template were deleted. This presumes that all of child data
	// containers are purged first.
	// base = 1024
	// units = B
	Bookmarks float64 `json:"bookmarks,omitempty"`
	// The amount of space consumed by data containers that were
	// provisioned from this data template. This is the space that will
	// be freed up if all of those data containers are deleted or purged.
	// This assumes that the data containers are deleted along with
	// underlying data sources.
	// base = 1024
	// units = B
	Containers float64 `json:"containers,omitempty"`
	// The data template that this usage information is for.
	// referenceTo = /delphix-js-data-template.json
	// format = objectReference
	Template string `json:"template,omitempty"`
	// The space that will be freed up if this template (and all of its
	// child data containers are deleted).
	// base = 1024
	// units = B
	Total float64 `json:"total,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The amount of space that would be consumed by the data in this
	// template (and child containers) without Delphix.
	// base = 1024
	// units = B
	Unvirtualized float64 `json:"unvirtualized,omitempty"`
}

// JSTimelinePointBookmarkInputStruct - Specifies the Self-Service timeline point using a reference to the
// Self-Service bookmark.
// extends JSTimelinePointParameters
type JSTimelinePointBookmarkInputStruct struct {
	// The Self-Service bookmark.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSTimelinePointLatestTimeInputStruct - Specifies the use of the latest available data from the given data
// layout.
// extends JSTimelinePointTimeParameters
type JSTimelinePointLatestTimeInputStruct struct {
	// The reference to the data layout used for this operation.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	// required = true
	SourceDataLayout string `json:"sourceDataLayout,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSTimelinePointTimeInputStruct - Specifies a point in time on the Self-Service timeline for a specific
// branch. Latest provisionable points before the specified time will be
// used.
// extends JSTimelinePointTimeParameters
type JSTimelinePointTimeInputStruct struct {
	// The reference to the branch used for this operation.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// required = true
	Branch string `json:"branch,omitempty"`
	// A point in time on the given branch.
	// required = true
	// format = date
	Time string `json:"time,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSTimestampDataParentStruct - The timestamp data parent of a REFRESH, RESTORE, UNDO or CREATE_BRANCH
// operation.
// extends JSDataParent
type JSTimestampDataParentStruct struct {
	// The branch this operation's data came from. This will be null if
	// the branch has been deleted.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
	// This will always contain the name of the branch, even if it has
	// been deleted.
	// maxLength = 256
	BranchName string `json:"branchName,omitempty"`
	// The data time on the branch that this operation's data came from.
	// format = date
	Time string `json:"time,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSUserUsageDataStruct - The space usage information for a Self-Service user.
// extends TypedObject
type JSUserUsageDataStruct struct {
	// The number of containers owned by this user.
	NumContainers *int `json:"numContainers,omitempty"`
	// The amount of space referenced by the data containers owned by
	// this user.
	// units = B
	// base = 1024
	Total float64 `json:"total,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The user.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
}

// JSWeeklyOperationCountStruct - Information about the number of operations on a data container each
// week for up to 30 weeks.
// extends JSUsageData
type JSWeeklyOperationCountStruct struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The date at the beginning of the time period this datapoint
	// corresponds to. The time period itself varies between datapoint
	// types.
	// format = date
	StartDate string `json:"startDate,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object the usage data is centered around.
	// format = objectReference
	// referenceTo = /delphix-named-user-object.json
	UsageObject string `json:"usageObject,omitempty"`
	// The number of operations run against a data container in the
	// specified week.
	WeeklyCount *int `json:"weeklyCount,omitempty"`
	// The total time spent in seconds running all operations during the
	// specified week.
	// units = sec
	WeeklyDuration *int `json:"weeklyDuration,omitempty"`
}

// JobStruct - Represents a job object.
// extends NamedUserObject
// cliVisibility = [SYSTEM DOMAIN]
type JobStruct struct {
	// Action type of the Job.
	ActionType string `json:"actionType,omitempty"`
	// A description of why the job was canceled.
	CancelReason string `json:"cancelReason,omitempty"`
	// Whether this job can be canceled.
	Cancelable *bool `json:"cancelable,omitempty"`
	// Email addresses to be notified on job notification alerts.
	// update = required
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	// A list of time-sorted past JobEvent objects associated with this
	// job.
	Events []*JobEventStruct `json:"events,omitempty"`
	// State of the job.
	// enum = [RUNNING SUSPENDED CANCELED COMPLETED FAILED]
	JobState string `json:"jobState,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// This job's parent action.
	// format = objectReference
	// referenceTo = /delphix-action.json
	ParentAction string `json:"parentAction,omitempty"`
	// State of this job's parent action. This value is populated only if
	// the job is fetched via the plain get API call.
	// enum = [EXECUTING WAITING COMPLETED FAILED CANCELED]
	ParentActionState string `json:"parentActionState,omitempty"`
	// Completion percentage. This value is a copy of the last event's
	// percentComplete. It will be 0 if there are no job events or if the
	// events field is not populated while fetching the job.
	// units = %
	PercentComplete float64 `json:"percentComplete,omitempty"`
	// Whether this job is waiting for resources to be available for its
	// execution.
	Queued *bool `json:"queued,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Time the job was created. Note that this is not the time when the
	// job started executing.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Whether this job can be suspended.
	Suspendable *bool `json:"suspendable,omitempty"`
	// Object reference of the target.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// A cached copy of the target object name.
	TargetName string `json:"targetName,omitempty"`
	// Object type of the target.
	// format = type
	TargetObjectType string `json:"targetObjectType,omitempty"`
	// Title of the job.
	Title string `json:"title,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Time the job was last updated.
	// format = date
	UpdateTime string `json:"updateTime,omitempty"`
	// User that initiated the action.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
}

// JobEventStruct - Represents a job event object. This can either be a state change or a
// progress update.
// extends TypedObject
type JobEventStruct struct {
	// Results of diagnostic checks run, if any, if the job failed.
	Diagnoses []*DiagnosisResultStruct `json:"diagnoses,omitempty"`
	// Type of event.
	// enum = [INFO WARNING ERROR]
	EventType string `json:"eventType,omitempty"`
	// Localized message action.
	MessageAction string `json:"messageAction,omitempty"`
	// Message ID associated with the event.
	MessageCode string `json:"messageCode,omitempty"`
	// Command output associated with the event, if applicable.
	MessageCommandOutput string `json:"messageCommandOutput,omitempty"`
	// Localized message details.
	MessageDetails string `json:"messageDetails,omitempty"`
	// Completion percentage.
	// units = %
	PercentComplete float64 `json:"percentComplete,omitempty"`
	// New state of the job.
	// enum = [INITIAL RUNNING SUSPENDED CANCELED COMPLETED FAILED RETRYABLE]
	State string `json:"state,omitempty"`
	// Time the event occurred.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JsonStruct - A dummy schema that is used to represent JSON.
type JsonStruct struct {
}

// KerberosConfigStruct - Kerberos Client Configuration.
// extends UserObject
// cliVisibility = [DOMAIN SYSTEM]
type KerberosConfigStruct struct {
	// Indicates whether kerberos has been configured or not.
	Enabled *bool `json:"enabled,omitempty"`
	// One of more KDC servers.
	// minItems = 1
	// create = required
	// update = optional
	Kdcs []*KerberosKDCStruct `json:"kdcs,omitempty"`
	// Kerberos keytab file data in base64 encoding.
	// create = required
	// update = optional
	// format = password
	Keytab string `json:"keytab,omitempty"`
	// Object name.
	// create = optional
	// update = optional
	// maxLength = 256
	// format = objectName
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Kerberos principal name.
	// create = required
	// update = optional
	Principal string `json:"principal,omitempty"`
	// Kerberos Realm name.
	// create = required
	// update = optional
	Realm string `json:"realm,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// KerberosCredentialStruct - Kerberos based security credential.
// extends Credential
type KerberosCredentialStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// KerberosKDCStruct - Kerberos Client Configuration.
// extends TypedObject
type KerberosKDCStruct struct {
	// KDC Server hostname.
	// format = host
	// create = required
	// update = optional
	Hostname string `json:"hostname,omitempty"`
	// KDC Server port number.
	// create = required
	// update = optional
	// minimum = 0
	// maximum = 65535
	// default = 88
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// KeyPairCredentialStruct - The public key based security credential consisting of a user
// specified key pair.
// extends PublicKeyCredential
type KeyPairCredentialStruct struct {
	// The private key in the key pair.
	// format = password
	// required = true
	PrivateKey string `json:"privateKey,omitempty"`
	// The public key in the key pair.
	// format = password
	// required = true
	PublicKey string `json:"publicKey,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// LdapInfoStruct - Global LDAP information.
// extends TypedObject
// cliVisibility = [SYSTEM DOMAIN]
type LdapInfoStruct struct {
	// Whether LDAP authentication is configured and enabled or not for
	// this Delphix Engine.
	Enabled *bool `json:"enabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// LdapServerStruct - LDAP Server Configuration.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type LdapServerStruct struct {
	// LDAP authentication method.
	// enum = [SIMPLE DIGEST_MD5]
	// required = true
	AuthMethod string `json:"authMethod,omitempty"`
	// LDAP server host name.
	// required = true
	// format = host
	Host string `json:"host,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// LDAP server port.
	// minimum = 1
	// maximum = 65535
	// required = true
	Port *int `json:"port,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Authenticate using SSL.
	// required = true
	UseSSL *bool `json:"useSSL,omitempty"`
}

// LinkParametersStruct - Represents the parameters of a link request.
// extends TypedObject
type LinkParametersStruct struct {
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	Description string `json:"description,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// required = true
	Group string `json:"group,omitempty"`
	// Database specific data required for linking.
	// create = required
	// required = true
	LinkData LinkData `json:"linkData,omitempty"`
	// DSource name.
	// format = objectName
	// required = true
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// LinkedSourceOperationsStruct - Describes operations which are performed on linked sources at various
// times.
// extends TypedObject
type LinkedSourceOperationsStruct struct {
	// Operations to perform after syncing a linked source.
	// create = optional
	// update = optional
	PostSync []SourceOperation `json:"postSync,omitempty"`
	// Operations to perform before syncing from a linked source. These
	// operations can quiesce any data prior to syncing.
	// create = optional
	// update = optional
	PreSync []SourceOperation `json:"preSync,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ListResultStruct - Result of a successful API call returning a list.
// extends OKResult
type ListResultStruct struct {
	// Reference to the action associated with the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-action.json
	Action string `json:"action,omitempty"`
	// Reference to the job started by the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-job.json
	Job string `json:"job,omitempty"`
	// True if the total number of matching items is too large to be
	// calculated.
	Overflow *bool `json:"overflow,omitempty"`
	// Result of the operation. This will be specific to the API being
	// invoked.
	Result string `json:"result,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
	// The number of items in the entire result set, regardless of the
	// requested page size. For some operations, this value is null.
	Total *int `json:"total,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// LocaleSettingsStruct - Global locale settings.
// extends UserObject
// cliVisibility = [SYSTEM]
type LocaleSettingsStruct struct {
	// System default locale as an IETF BCP 47 language tag, defaults to
	// 'en-US'.
	// format = locale
	// default = en-US
	// create = required
	// update = optional
	Locale string `json:"locale,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// LoginRequestStruct - Represents a Delphix user authentication request.
// extends TypedObject
type LoginRequestStruct struct {
	// Whether to keep session alive for all requests or only via
	// 'KeepSessionAlive' request headers. Defaults to ALL_REQUESTS if
	// omitted.
	// enum = [ALL_REQUESTS KEEP_ALIVE_HEADER_ONLY]
	// default = ALL_REQUESTS
	KeepAliveMode string `json:"keepAliveMode,omitempty"`
	// The password of the user to authenticate.
	// format = password
	// required = true
	Password string `json:"password,omitempty"`
	// The authentication domain.
	// enum = [DOMAIN SYSTEM]
	Target string `json:"target,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the user to authenticate.
	// required = true
	Username string `json:"username,omitempty"`
}

// MSSqlAttachDataStruct - Represents the MSSQL specific parameters of an attach request.
// extends AttachData
type MSSqlAttachDataStruct struct {
	// The password for accessing the shared backup location.
	// create = optional
	BackupLocationCredentials *PasswordCredentialStruct `json:"backupLocationCredentials,omitempty"`
	// The user for accessing the shared backup location.
	// maxLength = 256
	// create = optional
	BackupLocationUser string `json:"backupLocationUser,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mssql-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The SQL Server login password for the source DB user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The SQL Server login username for the source DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The encryption key to use when restoring encrypted backups.
	// create = optional
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// External file path.
	// create = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Configuration that determines what ingestion strategy the source
	// will use.
	// create = required
	IngestionStrategy IngestionStrategy `json:"ingestionStrategy,omitempty"`
	// Configuration for source that allows ingesting NetBackup backups
	// for SQL Server.
	// create = optional
	MssqlNetbackupConfig string `json:"mssqlNetbackupConfig,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// OS user on the PPT host to use for linking.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	PptHostUser string `json:"pptHostUser,omitempty"`
	// The SQL Server instance on the staging environment to use for
	// pre-provisioning.
	// format = objectReference
	// referenceTo = /delphix-mssql-instance.json
	// required = true
	PptRepository string `json:"pptRepository,omitempty"`
	// Shared source database backup locations.
	// maxItems = 1000
	// create = required
	SharedBackupLocations []string `json:"sharedBackupLocations,omitempty"`
	// OS user on the source to use for linking.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	SourceHostUser string `json:"sourceHostUser,omitempty"`
	// The path to a user-provided PowerShell script or executable to run
	// after restoring from a backup during pre-provisioning.
	// create = optional
	// maxLength = 1024
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// The path to a user-provided PowerShell script or executable to run
	// prior to restoring from a backup during pre-provisioning.
	// create = optional
	// maxLength = 1024
	StagingPreScript string `json:"stagingPreScript,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlAvailabilityGroupStruct - The representation of a SQL Server Availability Group.
// extends MSSqlBaseClusterRepository
type MSSqlAvailabilityGroupStruct struct {
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Represents Full-text search and semantic search feature.
	// create = required
	// update = optional
	FulltextInstalled *bool `json:"fulltextInstalled,omitempty"`
	// The list of MSSQL Cluster instances belonging to this repository.
	Instances []MSSqlBaseClusterInstance `json:"instances,omitempty"`
	// Internal version of the SQL Server instance.
	// create = required
	// update = optional
	// minimum = 1
	InternalVersion *int `json:"internalVersion,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The list of listeners belonging to this repository.
	Listeners []MSSqlBaseClusterListener `json:"listeners,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// create = optional
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// create = optional
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Version of the repository.
	// create = optional
	// update = optional
	Version string `json:"version,omitempty"`
}

// MSSqlAvailabilityGroupDBConfigStruct - Database for a SQL Server Availability Group.
// extends MSSqlDBConfig
type MSSqlAvailabilityGroupDBConfigStruct struct {
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// The name of the database.
	// create = required
	// update = optional
	// maxLength = 128
	DatabaseName string `json:"databaseName,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	// format = objectReference
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Specifies the current recovery model of the source database.
	// create = optional
	// update = readonly
	// enum = [FULL SIMPLE BULK_LOGGED]
	// default = SIMPLE
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mssql-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
}

// MSSqlAvailabilityGroupListenerStruct - The representation of a SQL Server Availability Group Listener.
// extends MSSqlBaseClusterListener
type MSSqlAvailabilityGroupListenerStruct struct {
	// The address of the listener.
	Address string `json:"address,omitempty"`
	// The name of the listener.
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The port for the listener.
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlClusterInstanceStruct - The representation of a SQL Server Instance on a clustered node for
// Availability Groups.
// extends MSSqlBaseClusterInstance
type MSSqlClusterInstanceStruct struct {
	// The owner of the SQL Server Instance.
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// The name of the SQL Server Instance.
	// maxLength = 16
	Name string `json:"name,omitempty"`
	// A reference to the Windows Cluster Node for this instance.
	// format = objectReference
	// referenceTo = /delphix-windows-cluster-node.json
	Node string `json:"node,omitempty"`
	// The port to connect to the SQL Server Instance.
	Port *int `json:"port,omitempty"`
	// The Servername of the SQL Server Instance.
	ServerName string `json:"serverName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The version of the SQL Server Instance.
	Version string `json:"version,omitempty"`
}

// MSSqlCompatibilityCriteriaStruct - The compatibility criteria to use for selecting compatible MSSql
// repositories.
// extends CompatibilityCriteria
type MSSqlCompatibilityCriteriaStruct struct {
	// Selected repositories are installed on a host with this
	// architecture (32-bit or 64-bit).
	Architecture *int `json:"architecture,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are this database version. In case of
	// upgrade, selected repositories are strictly greater than this
	// database version.
	Version string `json:"version,omitempty"`
}

// MSSqlCreateTransformationParametersStruct - Represents the parameters of a createTransformation request for a
// MSSql container.
// extends CreateTransformationParameters
type MSSqlCreateTransformationParametersStruct struct {
	// The container that will contain the transformed data associated
	// with the newly created transformation; the "transformation
	// container".
	// create = required
	Container *MSSqlDatabaseContainerStruct `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// MSSqlDBContainerRuntimeStruct - Runtime properties of a MSSQL database container.
// extends DBContainerRuntime
type MSSqlDBContainerRuntimeStruct struct {
	// The UUID of the source database backupset that was last restored
	// in this container.
	LastRestoredBackupSetUUID string `json:"lastRestoredBackupSetUUID,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntimeStruct `json:"preProvisioningStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlDatabaseContainerStruct - A MSSQL Database Container.
// extends DatabaseContainer
type MSSqlDatabaseContainerStruct struct {
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	// update = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this container.
	Runtime *MSSqlDBContainerRuntimeStruct `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlExistingMostRecentBackupSyncParametersStruct - The parameters to use as input to sync MSSQL databases using an
// existing most recent full or differential backup.
// extends MSSqlExistingBackupSyncParameters
type MSSqlExistingMostRecentBackupSyncParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlExistingSpecificBackupSyncParametersStruct - The parameters to use as input to sync MSSQL databases using an
// existing specific full or differential backup.
// extends MSSqlExistingBackupSyncParameters
type MSSqlExistingSpecificBackupSyncParametersStruct struct {
	// The UUID of the full or differential backup of the source database
	// to restore from. The backup should be present in the shared backup
	// location for the source database.
	// required = true
	BackupUUID string `json:"backupUUID,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlExportParametersStruct - The parameters to use as input to export MSSQL databases.
// extends DbExportParameters
type MSSqlExportParametersStruct struct {
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayoutStruct `json:"filesystemLayout,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// Recovery model of the database.
	// enum = [SIMPLE BULK_LOGGED FULL]
	// default = FULL
	// create = optional
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig MSSqlDBConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export
	// on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// MSSqlFailoverClusterDBConfigStruct - Database for a SQL Server Failover Cluster.
// extends MSSqlDBConfig
type MSSqlFailoverClusterDBConfigStruct struct {
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// The name of the database.
	// create = required
	// update = optional
	// maxLength = 128
	DatabaseName string `json:"databaseName,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Base drive letter location for mount points.
	// create = required
	// minLength = 1
	// maxLength = 1
	DriveLetter string `json:"driveLetter,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// create = readonly
	// update = readonly
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Specifies the current recovery model of the source database.
	// enum = [FULL SIMPLE BULK_LOGGED]
	// default = SIMPLE
	// create = optional
	// update = readonly
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mssql-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
}

// MSSqlFailoverClusterDriveLetterStruct - This represents a logical volume with a drive letter that resides on a
// Physical Disk cluster resource that is part of a SQL Server Failover
// Cluster Instance.
// extends TypedObject
type MSSqlFailoverClusterDriveLetterStruct struct {
	// The drive letter.
	// minLength = 1
	// maxLength = 1
	DriveLetter string `json:"driveLetter,omitempty"`
	// The drive letter label.
	// maxLength = 32
	Label string `json:"label,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlFailoverClusterInstanceStruct - The representation of a SQL Server Instance on a clustered node for
// Failover Clusters.
// extends MSSqlBaseClusterInstance
type MSSqlFailoverClusterInstanceStruct struct {
	// The owner of the SQL Server Instance.
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// The name of the SQL Server Instance.
	// maxLength = 16
	Name string `json:"name,omitempty"`
	// A reference to the Windows Cluster Node for this instance.
	// format = objectReference
	// referenceTo = /delphix-windows-cluster-node.json
	Node string `json:"node,omitempty"`
	// The port to connect to the SQL Server Instance.
	Port *int `json:"port,omitempty"`
	// The Servername of the SQL Server Instance.
	ServerName string `json:"serverName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The version of the SQL Server Instance.
	Version string `json:"version,omitempty"`
}

// MSSqlFailoverClusterListenerStruct - The representation of a SQL Server Failover Cluster Listener.
// extends MSSqlBaseClusterListener
type MSSqlFailoverClusterListenerStruct struct {
	// The address of the listener.
	Address string `json:"address,omitempty"`
	// The name of the listener.
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The port for the listener.
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlFailoverClusterRepositoryStruct - The representation of a SQL Server Failover Cluster repository.
// extends MSSqlBaseClusterRepository
type MSSqlFailoverClusterRepositoryStruct struct {
	// The list of drive letters belonging to this Failover Cluster
	// repository.
	Drives []*MSSqlFailoverClusterDriveLetterStruct `json:"drives,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Represents Full-text search and semantic search feature.
	// create = required
	// update = optional
	FulltextInstalled *bool `json:"fulltextInstalled,omitempty"`
	// The list of MSSQL Cluster instances belonging to this repository.
	Instances []MSSqlBaseClusterInstance `json:"instances,omitempty"`
	// Internal version of the SQL Server instance.
	// minimum = 1
	// create = required
	// update = optional
	InternalVersion *int `json:"internalVersion,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The list of listeners belonging to this repository.
	Listeners []MSSqlBaseClusterListener `json:"listeners,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// create = optional
	// update = optional
	// default = true
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// create = optional
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of the repository.
	// create = optional
	// update = optional
	Version string `json:"version,omitempty"`
}

// MSSqlInstanceStruct - A SQL Server Instance.
// extends MSSqlRepository
type MSSqlInstanceStruct struct {
	// True if the MSSQL instance was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Represents Full-text search and semantic search feature.
	// create = required
	// update = optional
	FulltextInstalled *bool `json:"fulltextInstalled,omitempty"`
	// The SQL Server instance home.
	// update = optional
	// minLength = 1
	// maxLength = 300
	// create = required
	InstallationPath string `json:"installationPath,omitempty"`
	// The name of the SQL Server instance.
	// minLength = 1
	// maxLength = 256
	// create = required
	InstanceName string `json:"instanceName,omitempty"`
	// Account the SQL Server instance is running as.
	// minLength = 1
	// maxLength = 255
	// create = required
	// update = optional
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// Internal version of the SQL Server instance.
	// update = optional
	// minimum = 1
	// create = required
	InternalVersion *int `json:"internalVersion,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The network port for connecting to the SQL Server instance.
	// minimum = 1
	// maximum = 65535
	// create = required
	// update = optional
	Port *int `json:"port,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// create = optional
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The Server Name of the SQL Server instance.
	// update = optional
	// maxLength = 128
	// create = optional
	ServerName string `json:"serverName,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// create = optional
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The version of the SQL Server instance.
	// create = required
	// update = optional
	// minLength = 1
	// maxLength = 256
	Version string `json:"version,omitempty"`
}

// MSSqlLinkDataStruct - MSSQL specific parameters for a link request.
// extends LinkData
type MSSqlLinkDataStruct struct {
	// The password for accessing the shared backup location.
	// create = optional
	BackupLocationCredentials *PasswordCredentialStruct `json:"backupLocationCredentials,omitempty"`
	// The user for accessing the shared backup location.
	// create = optional
	// maxLength = 256
	BackupLocationUser string `json:"backupLocationUser,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mssql-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The credential for the source DB user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The user name for the source DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The encryption key to use when restoring encrypted backups.
	// create = optional
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// External file path.
	// create = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Configuration that determines what ingestion strategy the source
	// will use.
	// create = required
	IngestionStrategy IngestionStrategy `json:"ingestionStrategy,omitempty"`
	// Configuration for source that allows ingesting NetBackup backups
	// for SQL Server.
	// create = optional
	MssqlNetbackupConfig string `json:"mssqlNetbackupConfig,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Information about the host OS user on the PPT host to use for
	// linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	PptHostUser string `json:"pptHostUser,omitempty"`
	// The SQL instance on the PPT environment that we want to use for
	// pre-provisioning.
	// format = objectReference
	// referenceTo = /delphix-mssql-instance.json
	// required = true
	PptRepository string `json:"pptRepository,omitempty"`
	// Shared source database backup locations.
	// maxItems = 1000
	// create = required
	SharedBackupLocations []string `json:"sharedBackupLocations,omitempty"`
	// Information about the host OS user on the source to use for
	// linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	SourceHostUser string `json:"sourceHostUser,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// A user-provided PowerShell script or executable to run after
	// restoring from a backup during pre-provisioning.
	// maxLength = 1024
	// create = optional
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// A user-provided PowerShell script or executable to run prior to
	// restoring from a backup during pre-provisioning.
	// create = optional
	// maxLength = 1024
	StagingPreScript string `json:"stagingPreScript,omitempty"`
	// Sync parameters for the container.
	// required = true
	SyncParameters MSSqlSyncParameters `json:"syncParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlLinkedSourceStruct - A linked MSSQL source.
// extends MSSqlSource
type MSSqlLinkedSourceStruct struct {
	// The password for accessing the shared backup location.
	// create = optional
	// update = optional
	BackupLocationCredentials *PasswordCredentialStruct `json:"backupLocationCredentials,omitempty"`
	// The user for accessing the shared backup location.
	// create = optional
	// update = optional
	// maxLength = 256
	BackupLocationUser string `json:"backupLocationUser,omitempty"`
	// Reference to the configuration for the source.
	// referenceTo = /delphix-mssql-db-config.json
	// create = required
	// update = optional
	// format = objectReference
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// The encryption key to use when restoring encrypted backups.
	// create = optional
	// update = optional
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// External file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Configuration that determines what ingestion strategy the source
	// will use.
	// create = required
	// update = optional
	IngestionStrategy IngestionStrategy `json:"ingestionStrategy,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Configuration for source that allows ingesting NetBackup backups
	// for SQL Server.
	// create = optional
	// update = optional
	MssqlNetbackupConfig string `json:"mssqlNetbackupConfig,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// update = optional
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *MSSqlSourceRuntimeStruct `json:"runtime,omitempty"`
	// Shared source database backup locations.
	// create = required
	// update = optional
	// maxItems = 1000
	SharedBackupLocations []string `json:"sharedBackupLocations,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// The staging source for pre-provisioning of the database.
	// format = objectReference
	// referenceTo = /delphix-mssql-staging-source.json
	StagingSource string `json:"stagingSource,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// MSSqlLinkedSourceUpgradeParametersStruct - The parameters to use as input to upgrade an MSSQL linked source.
// extends SourceUpgradeParameters
type MSSqlLinkedSourceUpgradeParametersStruct struct {
	// The SQL instance on the PPT environment that we want to use for
	// pre-provisioning.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-mssql-instance.json
	PptRepository string `json:"pptRepository,omitempty"`
	// The source config that the source database upgrades to.
	// required = true
	SourceConfig MSSqlDBConfig `json:"sourceConfig,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlNetbackupConfigStruct - MSSql NetBackup configuration.
// extends TypedObject
type MSSqlNetbackupConfigStruct struct {
	// NetBackup configuration parameter overrides.
	// update = optional
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional config template selection for NetBackup configurations.
	// If set, configParams will be ignored.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// The master server name of this NetBackup configuration.
	// maxLength = 256
	// create = required
	// update = optional
	MasterName string `json:"masterName,omitempty"`
	// The source's client server name of this NetBackup configuration.
	// update = optional
	// maxLength = 256
	// create = required
	SourceClientName string `json:"sourceClientName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlNewCopyOnlyFullBackupSyncParametersStruct - The parameters to use as input to sync MSSQL databases using a new
// copy-only full backup taken by Delphix.
// extends MSSqlSyncParameters
type MSSqlNewCopyOnlyFullBackupSyncParametersStruct struct {
	// If this parameter is set to true, Delphix will take a compressed
	// copy only full backup of the source database. When set to false,
	// Delphix will use the SQL Server instance's compression default.
	// default = false
	// required = true
	CompressionEnabled *bool `json:"compressionEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlPlatformParametersStruct - MSSql platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type MSSqlPlatformParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlProvisionParametersStruct - The parameters to use as input to provision MSSQL databases.
// extends ProvisionParameters
type MSSqlProvisionParametersStruct struct {
	// The new container for the provisioned database.
	// required = true
	Container *MSSqlDatabaseContainerStruct `json:"container,omitempty"`
	// Whether or not to mark this VDB as a masked VDB. It will be marked
	// as masked if this flag or the masking job are set.
	// create = optional
	// update = readonly
	Masked *bool `json:"masked,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// create = optional
	// update = readonly
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	MaskingJob string `json:"maskingJob,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *MSSqlVirtualSourceStruct `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of
	// the source.
	// required = true
	SourceConfig MSSqlDBConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlSIConfigStruct - Configuration information for a single instance MSSQL Source.
// extends MSSqlDBConfig
type MSSqlSIConfigStruct struct {
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// The name of the database.
	// create = required
	// update = optional
	// maxLength = 128
	DatabaseName string `json:"databaseName,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Specifies the current recovery model of the source database.
	// enum = [FULL SIMPLE BULK_LOGGED]
	// default = SIMPLE
	// create = optional
	// update = readonly
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mssql-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
}

// MSSqlSnapshotStruct - Provisionable snapshot of a MSSQL TimeFlow.
// extends TimeflowSnapshot
type MSSqlSnapshotStruct struct {
	// UUID of the source database backup that was restored for this
	// snapshot.
	BackupSetUUID string `json:"backupSetUUID,omitempty"`
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *MSSqlTimeflowPointStruct `json:"firstChangePoint,omitempty"`
	// Internal version of the source database at the time the snapshot
	// was taken.
	InternalVersion *int `json:"internalVersion,omitempty"`
	// The location of the snapshot within the parent TimeFlow
	// represented by this snapshot.
	LatestChangePoint *MSSqlTimeflowPointStruct `json:"latestChangePoint,omitempty"`
	// Boolean value indicating if a virtual database provisioned from
	// this snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot
	// should be kept forever.
	// update = optional
	Retention *int `json:"retention,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *MSSqlSnapshotRuntimeStruct `json:"runtime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient
	// state and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Time zone of the source database at the time the snapshot was
	// taken.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
}

// MSSqlSnapshotRuntimeStruct - Runtime (non-persistent) properties of a MSSQL TimeFlow snapshot.
// extends SnapshotRuntime
type MSSqlSnapshotRuntimeStruct struct {
	// True if this snapshot can be used as the basis for provisioning a
	// new TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlSourceConnectionInfoStruct - Contains information that can be used to connect to a SQL server
// source.
// extends SourceConnectionInfo
type MSSqlSourceConnectionInfoStruct struct {
	// The name of the database.
	DatabaseName string `json:"databaseName,omitempty"`
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The JDBC string used to connect to the SQL server instance.
	InstanceJDBCString string `json:"instanceJDBCString,omitempty"`
	// The name of the instance.
	InstanceName string `json:"instanceName,omitempty"`
	// The port number used to connect to the source.
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// MSSqlSourceRuntimeStruct - Runtime (non-persistent) properties of a MSSQL source.
// extends SourceRuntime
type MSSqlSourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlStagingSourceStruct - An MSSQL staging source used for pre-provisioning.
// extends MSSqlSource
type MSSqlStagingSourceStruct struct {
	// Reference to the configuration for the source.
	// referenceTo = /delphix-source-config.json
	// create = optional
	// format = objectReference
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// create = optional
	// update = optional
	// default = false
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point for the iSCSI LUN mounts.
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A user-provided PowerShell script or executable to run after
	// restoring from a backup during pre-provisioning.
	// create = optional
	// update = optional
	// maxLength = 1024
	PostScript string `json:"postScript,omitempty"`
	// A user-provided PowerShell script or executable to run prior to
	// restoring from a backup during pre-provisioning.
	// create = optional
	// update = optional
	// maxLength = 1024
	PreScript string `json:"preScript,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *MSSqlSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// MSSqlTimeflowStruct - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends Timeflow
type MSSqlTimeflowStruct struct {
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC SOURCE_CONTINUITY]
	CreationType string `json:"creationType,omitempty"`
	// MSSQL-specific recovery branch identifier for this TimeFlow.
	DatabaseGuid string `json:"databaseGuid,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow
	// was provisioned. This will not be present for TimeFlows derived
	// from linked sources.
	ParentPoint *MSSqlTimeflowPointStruct `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning
	// base for this object. This may be different from the snapshot
	// within the parent point, and is only present for virtual
	// TimeFlows.
	// referenceTo = /delphix-timeflow-snapshot.json
	// format = objectReference
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// MSSqlTimeflowPointStruct - A unique point within an MSSql database TimeFlow.
// extends TimeflowPoint
type MSSqlTimeflowPointStruct struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this point.
	// format = objectReference
	// referenceTo = /delphix-mssql-timeflow.json
	// required = true
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlVirtualSourceStruct - A virtual MSSQL source.
// extends MSSqlSource
type MSSqlVirtualSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// default = true
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// MSSQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point for the iSCSI LUN mounts.
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// A user-provided PowerShell script or executable to run after
	// provisioning.
	// create = optional
	// update = optional
	// maxLength = 256
	PostScript string `json:"postScript,omitempty"`
	// A user-provided PowerShell script or executable to run prior to
	// provisioning.
	// create = optional
	// update = optional
	// maxLength = 256
	PreScript string `json:"preScript,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *MSSqlSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// MaskingJobStruct - The Delphix Engine record of an existing Masking Job.
// extends UserObject
type MaskingJobStruct struct {
	// A reference to the container that the Masking Job is intended to
	// operate on.
	// format = objectReference
	// referenceTo = /delphix-container.json
	// create = optional
	// update = optional
	AssociatedContainer string `json:"associatedContainer,omitempty"`
	// The masking job id from the Delphix Masking Engine instance that
	// specifies the desired Masking Job.
	// update = readonly
	// create = required
	MaskingJobId string `json:"maskingJobId,omitempty"`
	// Object name.
	// format = objectName
	// create = required
	// update = readonly
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MaskingServiceConfigStruct - Configuration for the Masking Service this Engine communicates with.
// extends UserObject
type MaskingServiceConfigStruct struct {
	// Password to use when authenticating to the server.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// Object name.
	// update = readonly
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Port number to use.
	// minimum = 1
	// maximum = 65535
	// update = optional
	Port *int `json:"port,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Protocol scheme for use when communicating with server.
	// enum = [HTTP HTTPS]
	// update = optional
	Scheme string `json:"scheme,omitempty"`
	// IP address or hostname of server hosting Masking Service.
	// format = host
	// update = optional
	Server string `json:"server,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Username to use when authenticating to the server.
	// update = optional
	Username string `json:"username,omitempty"`
}

// MigrateCompatibilityParametersStruct - The criteria necessary to select valid repositories for migration.
// extends CompatibleRepositoriesParameters
type MigrateCompatibilityParametersStruct struct {
	// Restrict returned repositories to this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = optional
	// update = optional
	Environment string `json:"environment,omitempty"`
	// The repository to use as a source of compatibility information.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// required = true
	SourceRepository string `json:"sourceRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLAttachDataStruct - Represents the MySQL parameters of an attach request.
// extends AttachData
type MySQLAttachDataStruct struct {
	// Reference to the configuration for the source. Must be an existing
	// linked source config.
	// format = objectReference
	// referenceTo = /delphix-mysql-server-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// MySQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// The credentials for the database user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The database username.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// OS user on the staging host to use for attaching.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// The port on which the MySQL staging instance will listen.
	// minimum = 1
	// maximum = 65535
	// required = true
	StagingPort *int `json:"stagingPort,omitempty"`
	// The MySQL installation on the staging environment to use for
	// validated sync.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-mysql-install.json
	StagingRepository string `json:"stagingRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLBinlogCoordinatesStruct - The current position in the master binary logs when a MySQL backup was
// taken.
// extends MySQLReplicationCoordinates
type MySQLBinlogCoordinatesStruct struct {
	// Name of the master log file to start replication from.
	// required = true
	MasterLogName string `json:"masterLogName,omitempty"`
	// Position within the master log file to start replication from.
	// required = true
	MasterLogPos *int `json:"masterLogPos,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLCompatibilityCriteriaStruct - The compatibility criteria to use for selecting compatible MySQL
// repositories.
// extends CompatibilityCriteria
type MySQLCompatibilityCriteriaStruct struct {
	// Selected repositories are installed on a host with this
	// architecture (32-bit or 64-bit).
	Architecture *int `json:"architecture,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of the MySQL installation.
	Version *MySQLVersionStruct `json:"version,omitempty"`
}

// MySQLCreateTransformationParametersStruct - Represents the parameters of a createTransformation request for a
// MySQL container.
// extends CreateTransformationParameters
type MySQLCreateTransformationParametersStruct struct {
	// The container that will contain the transformed data associated
	// with the newly created transformation; the "transformation
	// container".
	// create = required
	Container *MySQLDatabaseContainerStruct `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// The port number used for provisioning a MySQL container during
	// transformation application.
	// create = required
	// minimum = 1
	// maximum = 65535
	Port *int `json:"port,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLDBContainerRuntimeStruct - Runtime properties of a MySQL database container.
// extends DBContainerRuntime
type MySQLDBContainerRuntimeStruct struct {
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntimeStruct `json:"preProvisioningStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLDatabaseContainerStruct - A MySQL Database Container.
// extends DatabaseContainer
type MySQLDatabaseContainerStruct struct {
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// A reference to the group containing this container.
	// referenceTo = /delphix-group.json
	// create = required
	// format = objectReference
	Group string `json:"group,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Whether to enable high performance mode.
	// default = DISABLED
	// create = readonly
	// update = readonly
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	PerformanceMode string `json:"performanceMode,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this container.
	Runtime *MySQLDBContainerRuntimeStruct `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Variant of the MySQL installation.
	// enum = [CommunityServer MariaDB]
	Variant string `json:"variant,omitempty"`
}

// MySQLExistingMySQLDumpSyncParametersStruct - The parameters to use as input to sync requests for MySQL databases
// using an existing MySQL dump.
// extends MySQLExistingBackupSyncParameters
type MySQLExistingMySQLDumpSyncParametersStruct struct {
	// Path to the existing backup to be loaded.
	// required = true
	BackupLocation string `json:"backupLocation,omitempty"`
	// The coordinates corresponding to the MySQL backup to start
	// replication from.
	// create = optional
	ReplicationCoordinates MySQLReplicationCoordinates `json:"replicationCoordinates,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLExportParametersStruct - The parameters to use as input to export MySQL databases.
// extends DbExportParameters
type MySQLExportParametersStruct struct {
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayoutStruct `json:"filesystemLayout,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// The source config to use when creating the exported database.
	// required = true
	SourceConfig *MySQLServerConfigStruct `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export
	// on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLGtidCoordinatesStruct - The GTID coordinates on the master when a MySQL backup was taken.
// extends MySQLReplicationCoordinates
type MySQLGtidCoordinatesStruct struct {
	// The GTID coordinates on the master to start replication from.
	// minItems = 1
	// required = true
	Gtids []string `json:"gtids,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLInstallStruct - A MySQL installation.
// extends SourceRepository
type MySQLInstallStruct struct {
	// Flag indicating whether the MySQL installation was discovered or
	// manually entered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Directory path where the MySQL installation is located.
	// create = required
	InstallationPath string `json:"installationPath,omitempty"`
	// Version of the MySQL installation.
	InternalVersion *MySQLVersionStruct `json:"internalVersion,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// update = optional
	// default = true
	// create = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// create = optional
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version string for the repository.
	// update = readonly
	// create = readonly
	Version string `json:"version,omitempty"`
}

// MySQLLinkDataStruct - MySQL specific parameters for a link request.
// extends LinkData
type MySQLLinkDataStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mysql-server-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// MySQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// The credentials for the database user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The database username.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// OS user on the staging host to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// The port on the staging host that the MySQL staging server can
	// listen on for TCP/IP connections.
	// required = true
	// minimum = 1
	// maximum = 65535
	StagingPort *int `json:"stagingPort,omitempty"`
	// The MySQL installation on the staging environment that will be
	// used for validated sync.
	// format = objectReference
	// referenceTo = /delphix-mysql-install.json
	// required = true
	StagingRepository string `json:"stagingRepository,omitempty"`
	// Sync parameters for the container.
	// required = true
	SyncParameters MySQLSyncParameters `json:"syncParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLLinkedSourceStruct - A linked MySQL source.
// extends MySQLSource
type MySQLLinkedSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mysql-server-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// MySQL database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *MySQLSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// The staging source for validated sync of the database.
	// format = objectReference
	// referenceTo = /delphix-mysql-staging-source.json
	StagingSource string `json:"stagingSource,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// MySQLNewMySQLDumpSyncParametersStruct - The parameters to use as input to sync requests for MySQL databases
// using a new MySQL dump taken by Delphix.
// extends MySQLSyncParameters
type MySQLNewMySQLDumpSyncParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLPlatformParametersStruct - MySQL platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type MySQLPlatformParametersStruct struct {
	// The port number used for provisioning a MySQL container during
	// transformation application. This port must be available when the
	// transformation is applied so that the temporary VDB created during
	// the transformation process can start and listen to this port.
	// minimum = 1
	// maximum = 65535
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLProvisionParametersStruct - The parameters to use as input to provision requests for MySQL
// databases.
// extends ProvisionParameters
type MySQLProvisionParametersStruct struct {
	// The new container for the provisioned database.
	// required = true
	Container *MySQLDatabaseContainerStruct `json:"container,omitempty"`
	// Whether or not to mark this VDB as a masked VDB. It will be marked
	// as masked if this flag or the masking job are set.
	// create = optional
	// update = readonly
	Masked *bool `json:"masked,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// update = readonly
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	MaskingJob string `json:"maskingJob,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *MySQLVirtualSourceStruct `json:"source,omitempty"`
	// The source config for the source.
	// required = true
	SourceConfig *MySQLServerConfigStruct `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLServerConfigStruct - Configuration information for a MySQL server instance.
// extends SourceConfig
type MySQLServerConfigStruct struct {
	// The password of the server instance user.
	// create = optional
	// update = optional
	Credentials *PasswordCredentialStruct `json:"credentials,omitempty"`
	// The data directory for the MySQL server instance.
	// create = optional
	DataDirectory string `json:"dataDirectory,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The port on which the MySQL server instance is listening.
	// update = optional
	// minimum = 1
	// maximum = 65535
	// create = required
	Port *int `json:"port,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mysql-install.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the server instance user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
}

// MySQLSnapshotStruct - Provisionable snapshot of a MySQL TimeFlow.
// extends TimeflowSnapshot
type MySQLSnapshotStruct struct {
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *MySQLTimeflowPointStruct `json:"firstChangePoint,omitempty"`
	// Version of the source database at the time the snapshot was taken.
	InternalVersion *MySQLVersionStruct `json:"internalVersion,omitempty"`
	// The location of the snapshot within the parent TimeFlow
	// represented by this snapshot.
	LatestChangePoint *MySQLTimeflowPointStruct `json:"latestChangePoint,omitempty"`
	// Boolean value indicating if a virtual database provisioned from
	// this snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot
	// should be kept forever.
	// update = optional
	Retention *int `json:"retention,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *MySQLSnapshotRuntimeStruct `json:"runtime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient
	// state and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Time zone of the source database at the time the snapshot was
	// taken.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
}

// MySQLSnapshotRuntimeStruct - Runtime (non-persistent) properties of a MySQL TimeFlow snapshot.
// extends SnapshotRuntime
type MySQLSnapshotRuntimeStruct struct {
	// True if this snapshot can be used as the basis for provisioning a
	// new TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// MySQLSourceConnectionInfoStruct - Contains information that can be used to connect to a MySQL source.
// extends SourceConnectionInfo
type MySQLSourceConnectionInfoStruct struct {
	// The data directory for the MySQL server.
	DataDirectory string `json:"dataDirectory,omitempty"`
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The JDBC string used to connect to the MySQL server instance.
	JdbcString string `json:"jdbcString,omitempty"`
	// The port on which the MySQL server for the data directory is
	// listening.
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	User string `json:"user,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// MySQLSourceRuntimeStruct - Non-persistent runtime properties of a MySQL source.
// extends SourceRuntime
type MySQLSourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLStagingSourceStruct - A MySQL staging source used for validated sync.
// extends MySQLSource
type MySQLStagingSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// create = optional
	// update = optional
	// default = false
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point for the NFS mounts on the staging host.
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A user-provided script to run after taking a snapshot.
	// create = optional
	// update = optional
	PostScript string `json:"postScript,omitempty"`
	// A user-provided script to run prior to taking a snapshot.
	// update = optional
	// create = optional
	PreScript string `json:"preScript,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *MySQLSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// MySQLTimeflowStruct - TimeFlow representing historical data for a particular timeline within
// a MySQL container.
// extends Timeflow
type MySQLTimeflowStruct struct {
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC SOURCE_CONTINUITY]
	CreationType string `json:"creationType,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow
	// was provisioned. This will not be present for TimeFlows derived
	// from linked sources.
	ParentPoint *MySQLTimeflowPointStruct `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning
	// base for this object. This may be different from the snapshot
	// within the parent point, and is only present for virtual
	// TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLTimeflowPointStruct - A unique point within a MySQL database TimeFlow.
// extends TimeflowPoint
type MySQLTimeflowPointStruct struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this point.
	// referenceTo = /delphix-mysql-timeflow.json
	// required = true
	// format = objectReference
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLVersionStruct - Version of a MySQL installation.
// extends TypedObject
type MySQLVersionStruct struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Variant of the MySQL installation.
	// enum = [CommunityServer MariaDB]
	Variant string `json:"variant,omitempty"`
	// Version of the MySQL installation.
	// format = mysqlVersion
	Version string `json:"version,omitempty"`
}

// MySQLVirtualSourceStruct - A virtual MySQL source.
// extends MySQLSource
type MySQLVirtualSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// MySQL database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point for the NFS mounts.
	// maxLength = 256
	// create = required
	// update = optional
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *MySQLSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// MySQLXtraBackupSyncParametersStruct - The parameters to use as input to sync requests for MySQL databases
// using an existing XtraBackup backup.
// extends MySQLExistingBackupSyncParameters
type MySQLXtraBackupSyncParametersStruct struct {
	// Path to the existing backup to be loaded.
	// required = true
	BackupLocation string `json:"backupLocation,omitempty"`
	// The coordinates corresponding to the MySQL backup to start
	// replication from.
	// create = optional
	ReplicationCoordinates MySQLReplicationCoordinates `json:"replicationCoordinates,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NTPConfigStruct - NTP (Network Time Protocol) configuration.
// extends TypedObject
type NTPConfigStruct struct {
	// If true, then time is synchronized with the configured NTP
	// servers. The management service is automatically restarted if this
	// value is changed.
	// update = optional
	// default = false
	Enabled *bool `json:"enabled,omitempty"`
	// Address to use for multicast NTP discovery. This is only valid
	// when 'useMulticast' is set.
	// default = 224.0.1.1
	// update = optional
	// format = ipv4Address
	MulticastAddress string `json:"multicastAddress,omitempty"`
	// A list of NTP servers to use for synchronization. At least one
	// server must be specified if multicast is not being used.
	// update = optional
	Servers []string `json:"servers,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If true, discover NTP servers using multicast.
	// update = optional
	// default = false
	UseMulticast *bool `json:"useMulticast,omitempty"`
}

// NamespaceStruct - Object namespace for target-side replication.
// extends NamedUserObject
type NamespaceStruct struct {
	// Description of this namespace.
	// update = optional
	// maxLength = 4096
	Description string `json:"description,omitempty"`
	// Indicates the namespace has been failed over into the live
	// environment.
	// default = false
	FailedOver *bool `json:"failedOver,omitempty"`
	// If failedOver is true, contains a report concern objects affected
	// during failover.
	FailoverReport string `json:"failoverReport,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Type of object namespace.
	// enum = [REPLICATION]
	NamespaceType string `json:"namespaceType,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// True if the source data stream was generated from a secure
	// replication spec.
	SecureNamespace *bool `json:"secureNamespace,omitempty"`
	// For replication, the source host IP address.
	Source string `json:"source,omitempty"`
	// A unique identifier for the source data stream.
	Tag string `json:"tag,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// NamespaceFailoverParametersStruct - The parameters to use as input to Namespace.failover.
// extends TypedObject
type NamespaceFailoverParametersStruct struct {
	// Enable automatic conflict resolution during failover.
	// create = optional
	// default = false
	SmartFailover *bool `json:"smartFailover,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetbackupConnectivityParametersStruct - Parameters needed to test NetBackup connectivity on an environment.
// extends TypedObject
type NetbackupConnectivityParametersStruct struct {
	// Target environment to test NetBackup connectivity from.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// required = true
	Environment string `json:"environment,omitempty"`
	// The environment user to use to connect to the environment.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The name of the NetBackup master server to attempt to connect to.
	// required = true
	MasterServerName string `json:"masterServerName,omitempty"`
	// The name of the NetBackup client to attempt to connect with.
	// required = true
	SourceClientName string `json:"sourceClientName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkDSPTestStruct - DSP throughput tests to a target system.
// extends NetworkThroughputTestBase
// cliVisibility = [DOMAIN SYSTEM]
type NetworkDSPTestStruct struct {
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Number of connections used to achieve maximum sustained
	// throughput.
	NumConnections *int `json:"numConnections,omitempty"`
	// The parameters used to execute the test.
	Parameters *NetworkDSPTestParametersStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The remote IP address used for the test.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// The state of the test.
	// enum = [RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// Average throughput measured.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkDSPTestParametersStruct - Parameters used to execute a network throughput test using the Delphix
// Session Protocol.
// extends NetworkThroughputTestBaseParameters
type NetworkDSPTestParametersStruct struct {
	// The size of each transmit request in bytes.
	// base = 1024
	// units = B
	// create = optional
	// default = 65536
	// minimum = 0
	// maximum = 1.048576e+06
	BlockSize *int `json:"blockSize,omitempty"`
	// Whether or not compression is used for the test.
	// default = false
	// create = optional
	Compression *bool `json:"compression,omitempty"`
	// Whether the test is testing connectivity to a Delphix Engine or
	// remote host.
	// create = optional
	// enum = [REMOTE_HOST DELPHIX_ENGINE]
	// default = REMOTE_HOST
	DestinationType string `json:"destinationType,omitempty"`
	// Whether the test is a transmit or receive test.
	// enum = [TRANSMIT RECEIVE]
	// default = TRANSMIT
	// create = optional
	Direction string `json:"direction,omitempty"`
	// The duration of the test in seconds. Note that when numConnections
	// is 0, an initial period of time will be spent calculating the
	// optimal number of connections, and that time does not count toward
	// the duration of the test.
	// default = 30
	// create = optional
	// minimum = 1
	// maximum = 3600
	Duration *int `json:"duration,omitempty"`
	// Whether or not encryption is used for the test.
	// default = false
	// create = optional
	Encryption *bool `json:"encryption,omitempty"`
	// The number of connections to use for the test. The special value 0
	// (the default) causes the test to automatically discover and use
	// the optimal number of connections to use for maximum throughput.
	// minimum = 0
	// maximum = 32
	// default = 0
	// create = optional
	NumConnections *int `json:"numConnections,omitempty"`
	// The queue depth used for the DSP throughput test.
	// default = 32
	// minimum = 0
	// maximum = 4096
	// create = optional
	QueueDepth *int `json:"queueDepth,omitempty"`
	// The size of the receive socket buffer in bytes.
	// default = 262144
	// minimum = 0
	// maximum = 1.6777216e+07
	// base = 1024
	// units = B
	// create = optional
	ReceiveSocketBuffer *int `json:"receiveSocketBuffer,omitempty"`
	// Address, username and password used when running a test to another
	// Delphix Engine.
	// create = optional
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfoStruct `json:"remoteDelphixEngineInfo,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	RemoteHost string `json:"remoteHost,omitempty"`
	// The size of the send socket buffer in bytes.
	// create = optional
	// default = 262144
	// minimum = 0
	// maximum = 1.6777216e+07
	// base = 1024
	// units = B
	SendSocketBuffer *int `json:"sendSocketBuffer,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkInterfaceStruct - Configuration of an IP interface.
// extends UserObject
// cliVisibility = [DOMAIN SYSTEM]
type NetworkInterfaceStruct struct {
	// List of IP addresses assigned to the interface.
	// update = optional
	Addresses []*InterfaceAddressStruct `json:"addresses,omitempty"`
	// The name of the device over which this interface is configured.
	Device string `json:"device,omitempty"`
	// The MAC address associated with this interface.
	// format = macAddress
	MacAddress string `json:"macAddress,omitempty"`
	// The maximum transmission unit for this interface.
	// update = optional
	Mtu *int `json:"mtu,omitempty"`
	// The range of possible values for the mtu property.
	MtuRange string `json:"mtuRange,omitempty"`
	// Object name.
	// create = optional
	// update = optional
	// maxLength = 256
	// format = objectName
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The state of the interface.
	// enum = [OK DOWN FAILED]
	State string `json:"state,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkInterfaceUtilDatapointStruct - An analytics datapoint generated by the NETWORK_INTERFACE_UTIL
// statistic type.
// extends Datapoint
type NetworkInterfaceUtilDatapointStruct struct {
	// Bytes received on the interface.
	InBytes *int `json:"inBytes,omitempty"`
	// Packets received on the interface.
	InPackets *int `json:"inPackets,omitempty"`
	// Bytes transmitted on the interface.
	OutBytes *int `json:"outBytes,omitempty"`
	// Packets transmitted on the interface.
	OutPackets *int `json:"outPackets,omitempty"`
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkInterfaceUtilDatapointStreamStruct - A stream of datapoints from a NETWORK_INTERFACE_UTIL analytics slice.
// extends DatapointStream
type NetworkInterfaceUtilDatapointStreamStruct struct {
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// Which network interface was utilized.
	NetworkInterface string `json:"networkInterface,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkLatencyTestStruct - Round-trip latency tests to a target system.
// extends NetworkTest
// cliVisibility = [DOMAIN SYSTEM]
type NetworkLatencyTestStruct struct {
	// Average measured round-trip time (usec).
	// units = usec
	Average *int `json:"average,omitempty"`
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Percentage of requests or replies lost.
	Loss *int `json:"loss,omitempty"`
	// Maximum measured round-trip time (usec).
	// units = usec
	Maximum *int `json:"maximum,omitempty"`
	// Minimum measured round-trip time (usec).
	// units = usec
	Minimum *int `json:"minimum,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The parameters used to execute the test.
	Parameters *NetworkLatencyTestParametersStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The remote IP address used for the test.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// The state of the test.
	// enum = [RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// Standard deviation (usec).
	// units = usec
	Stddev *int `json:"stddev,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// NetworkLatencyTestParametersStruct - Parameters used to execute a network latency test.
// extends NetworkTestParameters
type NetworkLatencyTestParametersStruct struct {
	// A hostname or literal IP address to test. Either remoteAddress or
	// remoteHost must be set, but not both.
	// format = host
	// create = optional
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	RemoteHost string `json:"remoteHost,omitempty"`
	// Number of requests to send.
	// create = optional
	// minimum = 1
	// maximum = 3600
	// default = 20
	RequestCount *int `json:"requestCount,omitempty"`
	// The size of requests to send (bytes).
	// minimum = 8
	// maximum = 65507
	// default = 8
	// create = optional
	// units = B
	// base = 1024
	RequestSize *int `json:"requestSize,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkRouteStruct - IP routing table.
// extends TypedObject
// cliVisibility = [DOMAIN SYSTEM]
type NetworkRouteStruct struct {
	// Destination for the route in Classless Inter-Domain Routing (CIDR)
	// notation or the keyword 'default'.
	// format = routeDestination
	// create = required
	Destination string `json:"destination,omitempty"`
	// Next hop for the route.
	// create = required
	// format = ipAddress
	Gateway string `json:"gateway,omitempty"`
	// Output interface to use for the route.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-network-interface.json
	OutInterface string `json:"outInterface,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// NetworkRouteLookupParametersStruct - Parameters used for a routing table lookup.
// extends TypedObject
type NetworkRouteLookupParametersStruct struct {
	// Destination address or hostname.
	// format = host
	// create = required
	Destination string `json:"destination,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkThroughputTestStruct - Bi-directional throughput tests to a target system.
// extends NetworkThroughputTestBase
// cliVisibility = [DOMAIN SYSTEM]
type NetworkThroughputTestStruct struct {
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Number of connections used to achieve maximum sustained
	// throughput.
	NumConnections *int `json:"numConnections,omitempty"`
	// The parameters used to execute the test.
	Parameters *NetworkThroughputTestParametersStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The remote IP address used for the test.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// The state of the test.
	// enum = [RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// Average throughput measured.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NetworkThroughputTestParametersStruct - Parameters used to execute a network throughput test.
// extends NetworkThroughputTestBaseParameters
type NetworkThroughputTestParametersStruct struct {
	// The size of each transmit request in bytes.
	// base = 1024
	// units = B
	// create = optional
	// default = 131072
	// minimum = 0
	// maximum = 1.048576e+06
	BlockSize *int `json:"blockSize,omitempty"`
	// Whether the test is a transmit or receive test.
	// enum = [TRANSMIT RECEIVE]
	// default = TRANSMIT
	// create = optional
	Direction string `json:"direction,omitempty"`
	// The duration of the test in seconds. Note that when numConnections
	// is 0, an initial period of time will be spent calculating the
	// optimal number of connections, and that time does not count toward
	// the duration of the test.
	// minimum = 1
	// maximum = 3600
	// default = 30
	// create = optional
	Duration *int `json:"duration,omitempty"`
	// The number of connections to use for the test. The special value 0
	// (the default) causes the test to automatically discover and use
	// the optimal number of connections to use for maximum throughput.
	// minimum = 0
	// maximum = 32
	// default = 0
	// create = optional
	NumConnections *int `json:"numConnections,omitempty"`
	// The TCP port number that the server (the receiver) will be
	// listening on.
	// maximum = 65535
	// create = optional
	// minimum = 1
	Port *int `json:"port,omitempty"`
	// A hostname or literal IP address to test. This parameter is
	// optional and can be provided if the remoteHost has multiple
	// addresses.
	// format = host
	// create = optional
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment. If the host has multiple addresses and remoteAddress
	// is not specified, then the default address used when adding the
	// host will be used.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = required
	RemoteHost string `json:"remoteHost,omitempty"`
	// The size of the send socket buffer in bytes.
	// base = 1024
	// units = B
	// create = optional
	// default = 4.194304e+06
	// minimum = 0
	// maximum = 1.6777216e+07
	SendSocketBuffer *int `json:"sendSocketBuffer,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NfsOpsDatapointStreamStruct - A stream of datapoints from an NFS_OPS analytics slice.
// extends DatapointStream
type NfsOpsDatapointStreamStruct struct {
	// Whether reads were cached.
	Cached *bool `json:"cached,omitempty"`
	// Address of the client.
	// format = host
	Client string `json:"client,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Path of the affected file.
	// format = unixpath
	Path string `json:"path,omitempty"`
	// Whether writes were synchronous.
	Sync *bool `json:"sync,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NoBackupIngestionStrategyStruct - This linked source has no ingestion strategy meaning ValidatedSync is
// disabled and it is not Delphix managed.
// extends IngestionStrategy
type NoBackupIngestionStrategyStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NotFilterStruct - A container filter that inverts the logic of another filter.
// extends AlertFilter
type NotFilterStruct struct {
	// Filter whose logic is to be inverted.
	// create = required
	// update = optional
	// properties = map[type:map[default:SeverityFilter]]
	SubFilter AlertFilter `json:"subFilter,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NotificationDropStruct - An object to track dropped notifications.
// extends Notification
type NotificationDropStruct struct {
	// The number of notifications which were dropped since the last
	// notifications were pulled. If this is greater than zero, you may
	// want to refresh your view of the data to ensure everything is up
	// to date.
	DropCount *int `json:"dropCount,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// NullConstraintStruct - If an axis has this type of constraint, it means that no constraints
// can be placed on this axis. This constraint type does nothing and has
// no descendent types.
// extends AxisConstraint
type NullConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OKResultStruct - Result of a successful API call.
// extends CallResult
type OKResultStruct struct {
	// Reference to the action associated with the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-action.json
	Action string `json:"action,omitempty"`
	// Reference to the job started by the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-job.json
	Job string `json:"job,omitempty"`
	// Result of the operation. This will be specific to the API being
	// invoked.
	Result string `json:"result,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ObjectNotificationStruct - An event indicating a change to an object on the system.
// extends Notification
type ObjectNotificationStruct struct {
	// Type of operation on the object.
	// enum = [CREATE UPDATE DELETE]
	EventType string `json:"eventType,omitempty"`
	// Target object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Object string `json:"object,omitempty"`
	// Type of target object.
	// format = type
	ObjectType string `json:"objectType,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// OperationTemplateStruct - Template for commonly used operations.
// extends NamedUserObject
type OperationTemplateStruct struct {
	// User provided description for this template.
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// Most recently modified time.
	// format = date
	LastUpdated string `json:"lastUpdated,omitempty"`
	// The name clients should use when setting the parameter's value.
	// create = required
	// update = optional
	// minLength = 1
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Template contents.
	// create = required
	// update = optional
	Operation SourceOperation `json:"operation,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OrFilterStruct - A container filter that combines other filters together using OR
// logic.
// extends AlertFilter
type OrFilterStruct struct {
	// Filters which are combined together using OR logic.
	// create = required
	// update = optional
	// minItems = 2
	SubFilters []AlertFilter `json:"subFilters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleActiveInstanceStruct - Active instance information for an Oracle database.
// extends TypedObject
type OracleActiveInstanceStruct struct {
	// The name of the host the instance runs on.
	HostName string `json:"hostName,omitempty"`
	// The name of the Oracle instance.
	InstanceName string `json:"instanceName,omitempty"`
	// The number of the Oracle instance.
	InstanceNumber *int `json:"instanceNumber,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleAddLiveSourceParametersStruct - The parameters to use as input to convert an Oracle dSource to an
// Oracle LiveSource.
// extends AddLiveSourceParameters
type OracleAddLiveSourceParametersStruct struct {
	// The security credential of the privileged user to run the
	// LiveSource creation operation as.
	// create = optional
	Credential Credential `json:"credential,omitempty"`
	// The source that describes the LiveSource.
	// required = true
	Source *OracleLiveSourceStruct `json:"source,omitempty"`
	// The source config of the LiveSource.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the LiveSource creation
	// operation as.
	// create = optional
	Username string `json:"username,omitempty"`
}

// OracleAttachDataStruct - Represents parameters to attach non-pluggable Oracle databases.
// extends OracleBaseAttachData
type OracleAttachDataStruct struct {
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit *int `json:"bandwidthLimit,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// True if SnapSync data from the source should be compressed over
	// the network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially
	// over slow network.
	// default = true
	// create = optional
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Reference to the configuration for the source.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-oracle-db-config.json
	Config string `json:"config,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession
	// to reduce the number of logs required to provision the snapshot.
	// This may significantly reduce the time necessary to provision from
	// a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// True if SnapSync data from the source should be retrieved through
	// an encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// Information about the OS user to use for linking.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// External file path.
	// create = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	FilesPerSet *int `json:"filesPerSet,omitempty"`
	// True if attach should succeed even if the resetlogs of the
	// original database does not match the resetlogs of the new database
	// and the resetlogs information of the original database is not a
	// parent incarnation of the current resetlogs. This can happen when
	// the controlfile has been recreated and the incarnation table is
	// incomplete. Use this option with extreme caution. Attached
	// database must be the same database to avoid data corruption later
	// on.
	// default = false
	// create = optional
	Force *bool `json:"force,omitempty"`
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// default = 1
	// minimum = 1
	// maximum = 16
	// create = optional
	NumberOfConnections *int `json:"numberOfConnections,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	RmanChannels *int `json:"rmanChannels,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// OracleCharacterSetStruct - Represents an Oracle character set.
// extends TypedObject
type OracleCharacterSetStruct struct {
	// Name of character set.
	CharacterSet string `json:"characterSet,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleClusterStruct - The representation of an oracle cluster environment object.
// extends SourceEnvironment
type OracleClusterStruct struct {
	// A reference to the cluster user.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	ClusterUser string `json:"clusterUser,omitempty"`
	// The location of the cluster installation.
	// create = required
	// update = optional
	// maxLength = 256
	CrsClusterHome string `json:"crsClusterHome,omitempty"`
	// The name of the cluster.
	// maxLength = 15
	// create = optional
	CrsClusterName string `json:"crsClusterName,omitempty"`
	// The environment description.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source
	// environment.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// A reference to the primary user for this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	PrimaryUser string `json:"primaryUser,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The default remote_listener parameter to be used for databases on
	// the cluster.
	// create = optional
	// update = optional
	// maxLength = 256
	RemoteListener string `json:"remoteListener,omitempty"`
	// The Single Client Access Name of the cluster (11.2 and greater
	// clusters only).
	// update = optional
	// maxLength = 256
	Scan string `json:"scan,omitempty"`
	// Indicates whether the Single Client Access Name of the cluster is
	// manually configured.
	ScanManual *bool `json:"scanManual,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The version of the cluster.
	// create = optional
	// maxLength = 14
	Version string `json:"version,omitempty"`
}

// OracleClusterCreateParametersStruct - The parameters used for the oracle cluster create operation.
// extends SourceEnvironmentCreateParameters
type OracleClusterCreateParametersStruct struct {
	// The representation of the cluster object.
	// create = required
	Cluster *OracleClusterStruct `json:"cluster,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to the created source
	// environment.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Only one node is allowed for the add cluster operation. Additional
	// nodes will be discovered automatically. Any nodes not discovered
	// by Delphix can be manually added after cluster creation.
	// create = required
	Node *OracleClusterNodeCreateParametersStruct `json:"node,omitempty"`
	// The primary user associated with the environment.
	// create = required
	PrimaryUser *EnvironmentUserStruct `json:"primaryUser,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleClusterNodeStruct - The representation of an oracle cluster node object.
// extends NamedUserObject
type OracleClusterNodeStruct struct {
	// The reference to the parent cluster environment.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster.json
	Cluster string `json:"cluster,omitempty"`
	// Indicates whether the node is discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Indicates whether the node is enabled.
	// update = optional
	Enabled *bool `json:"enabled,omitempty"`
	// The reference to the associated host object.
	// referenceTo = /delphix-host.json
	// format = objectReference
	Host string `json:"host,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The list of virtual IPs belonging to this node.
	// update = optional
	VirtualIPs []*OracleVirtualIPStruct `json:"virtualIPs,omitempty"`
}

// OracleClusterNodeCreateParametersStruct - The parameters used for oracle cluster node operations.
// extends TypedObject
type OracleClusterNodeCreateParametersStruct struct {
	// The cluster to which the node belongs.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster.json
	// create = optional
	Cluster string `json:"cluster,omitempty"`
	// The host object associated with the cluster node.
	// create = required
	HostParameters HostCreateParameters `json:"hostParameters,omitempty"`
	// The name of the cluster node.
	// create = optional
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The list of virtual IPs belonging to this node.
	// create = optional
	VirtualIPs []*OracleVirtualIPStruct `json:"virtualIPs,omitempty"`
}

// OracleCompatibilityCriteriaStruct - The compatibility criteria to use for selecting compatible Oracle
// repositories.
// extends CompatibilityCriteria
type OracleCompatibilityCriteriaStruct struct {
	// Selected repositories are installed on a host with this
	// architecture (32-bit or 64-bit).
	Architecture *int `json:"architecture,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are this database version. In case of
	// upgrade, selected repositories are strictly greater than this
	// database version.
	Version string `json:"version,omitempty"`
}

// OracleCreateTransformationParametersStruct - Represents the parameters of a createTransformation request for an
// Oracle container.
// extends CreateTransformationParameters
type OracleCreateTransformationParametersStruct struct {
	// The container that will contain the transformed data associated
	// with the newly created transformation; the "transformation
	// container".
	// create = required
	Container *OracleDatabaseContainerStruct `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleCustomEnvVarRACFileStruct - Dictates an environment file to be sourced when the Delphix Engine
// administers an Oracle virtual database. This environment file must be
// available on the target environment. This type also includes
// parameters which will be passed to the environment file when it is
// sourced. For a RAC environment, the cluster node where the target
// environment file exists must also be specified.
// extends OracleCustomEnvVarFile
type OracleCustomEnvVarRACFileStruct struct {
	// The cluster node on which the target environment file exists.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster-node.json
	// required = true
	ClusterNode string `json:"clusterNode,omitempty"`
	// A string of whitespace-separated parameters to be passed to the
	// source command. The first parameter must be an absolute path to a
	// file that exists on the target environment. Every subsequent
	// parameter will be treated as an argument interpreted by the
	// environment file.
	// required = true
	PathParameters string `json:"pathParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleCustomEnvVarRACPairStruct - Dictates a single environment variable name and value to be set when
// the Delphix Engine administers an Oracle virtual database. For a RAC
// environment, the cluster node where the target pair is valid must also
// be specified.
// extends OracleCustomEnvVarPair
type OracleCustomEnvVarRACPairStruct struct {
	// The cluster node on which the environment variable is relevant.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster-node.json
	// required = true
	ClusterNode string `json:"clusterNode,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the environment variable.
	// format = envvarIdentifier
	// required = true
	VarName string `json:"varName,omitempty"`
	// The value of the environment variable.
	// required = true
	VarValue string `json:"varValue,omitempty"`
}

// OracleCustomEnvVarSIFileStruct - Dictates an environment file to be sourced when the Delphix Engine
// administers an Oracle virtual database. This environment file must be
// available on the target environment. This type also includes
// parameters which will be passed to the environment file when it is
// sourced.
// extends OracleCustomEnvVarFile
type OracleCustomEnvVarSIFileStruct struct {
	// A string of whitespace-separated parameters to be passed to the
	// source command. The first parameter must be an absolute path to a
	// file that exists on the target environment. Every subsequent
	// parameter will be treated as an argument interpreted by the
	// environment file.
	// required = true
	PathParameters string `json:"pathParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleCustomEnvVarSIPairStruct - Dictates a single environment variable name and value to be set when
// the Delphix Engine administers an Oracle virtual database.
// extends OracleCustomEnvVarPair
type OracleCustomEnvVarSIPairStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the environment variable.
	// format = envvarIdentifier
	// required = true
	VarName string `json:"varName,omitempty"`
	// The value of the environment variable.
	// required = true
	VarValue string `json:"varValue,omitempty"`
}

// OracleDBConfigConnectivityStruct - Mechanism to test JDBC connectivity to Oracle source configs.
// extends AbstractSourceConfigConnectivity
type OracleDBConfigConnectivityStruct struct {
	// Database password.
	// format = password
	// create = optional
	// update = optional
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Database username.
	// create = optional
	// update = optional
	Username string `json:"username,omitempty"`
}

// OracleDBContainerRuntimeStruct - Runtime properties of an Oracle database container.
// extends DBContainerRuntime
type OracleDBContainerRuntimeStruct struct {
	// Indicates whether or not the given container is cross-platform
	// eligible or not.
	// default = false
	CrossPlatformEligible *bool `json:"crossPlatformEligible,omitempty"`
	// Indicates whether or not the given container has a cross-platform
	// user script uploaded.
	// default = false
	CrossPlatformScriptUploaded *bool `json:"crossPlatformScriptUploaded,omitempty"`
	// Indicates whether or not a LiveSource can be added to the given
	// container.
	// default = false
	LiveSourceEligible *bool `json:"liveSourceEligible,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntimeStruct `json:"preProvisioningStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleDatabaseContainerStruct - Data container for Oracle databases, both linked and virtual.
// extends DatabaseContainer
type OracleDatabaseContainerStruct struct {
	// Indicates whether this container is for a PDB, CDB root, auxiliary
	// CDB, or non-CDB.
	// enum = [PDB ROOT_CDB AUX_CDB NON_CDB]
	ContentType string `json:"contentType,omitempty"`
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Indicates whether or not this container is ready for
	// cross-platform provisioning.
	// default = false
	CrossPlatformReady *bool `json:"crossPlatformReady,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Indicates whether or not the database in this container consists
	// only of transportable tablespaces.
	// featureFlag = CONSPRO
	// default = false
	DatabaseFraction *bool `json:"databaseFraction,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// If true, NOLOGGING operations on this container are treated as
	// faults and cannot be resolved manually. Otherwise, these
	// operations are ignored.
	// default = true
	// create = optional
	// update = optional
	DiagnoseNoLoggingFaults *bool `json:"diagnoseNoLoggingFaults,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// Indicates whether or not this container has an associated
	// LiveSource.
	// default = false
	LiveSource *bool `json:"liveSource,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = optional
	// update = optional
	PerformanceMode string `json:"performanceMode,omitempty"`
	// Indicates whether or not the database in this container is a
	// physical standby.
	// default = false
	PhysicalStandby *bool `json:"physicalStandby,omitempty"`
	// If true, pre-provisioning will be performed after every sync.
	// default = false
	// create = optional
	// update = optional
	PreProvisioningEnabled *bool `json:"preProvisioningEnabled,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this container.
	Runtime *OracleDBContainerRuntimeStruct `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *OracleSourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleDatabaseCreationParametersStruct - The parameters to use as input when creating a new Oracle database.
// extends EmptyDatasetCreationParameters
type OracleDatabaseCreationParametersStruct struct {
	// The character set the database uses to store data.
	// update = optional
	// pattern = ^UTF8|UTFE|[A-Z]+[0-9]+[A-Z0-9]+$
	// default = AL32UTF8
	// create = optional
	CharacterSet string `json:"characterSet,omitempty"`
	// The new container for the created database.
	// required = true
	Container *OracleDatabaseContainerStruct `json:"container,omitempty"`
	// The password for the Delphix database user.
	// format = password
	// required = true
	DelphixPassword string `json:"delphixPassword,omitempty"`
	// The name of the Delphix database user.
	// required = true
	// pattern = ^[a-zA-Z][_a-zA-Z0-9]*$
	DelphixUsername string `json:"delphixUsername,omitempty"`
	// Puts the database into FORCE LOGGING mode. Oracle Database will
	// log all changes in the database except for changes in temporary
	// tablespaces and temporary segments.
	// default = false
	// create = optional
	// update = optional
	ForceLogging *bool `json:"forceLogging,omitempty"`
	// Grants the SELECT ANY DICTIONARY system privilege to the Delphix
	// database user. If disabled, the Delphix database user will only
	// have SELECT access to a limited set of views.
	// update = optional
	// default = true
	// create = optional
	GrantSelectAnyDictionary *bool `json:"grantSelectAnyDictionary,omitempty"`
	// The initial sizing of the data files section of the control file
	// at CREATE DATABASE or CREATE CONTROLFILE time.
	// create = optional
	// update = optional
	// minimum = 2
	// default = 32
	// maximum = 65535
	MaxDataFiles *int `json:"maxDataFiles,omitempty"`
	// The maximum number of instances that can simultaneously have this
	// database mounted and open.
	// minimum = 1
	// default = 32
	// maximum = 1055
	// create = optional
	// update = optional
	MaxInstances *int `json:"maxInstances,omitempty"`
	// The maximum number of redo log files that can ever be created for
	// the database.
	// default = 64
	// maximum = 255
	// create = optional
	// update = optional
	// minimum = 3
	MaxLogFiles *int `json:"maxLogFiles,omitempty"`
	// The maximum number of archived redo log files for automatic media
	// recovery of Oracle RAC.
	// default = 100
	// maximum = 65535
	// create = optional
	// update = optional
	// minimum = 0
	MaxLogHistory *int `json:"maxLogHistory,omitempty"`
	// The national character set used to store data in columns
	// specifically defined as NCHAR, NCLOB, or NVARCHAR2.
	// create = optional
	// update = optional
	// enum = [AL16UTF16 UTF8]
	// default = AL16UTF16
	NationalCharacterSet string `json:"nationalCharacterSet,omitempty"`
	// The redo log files. If no filename is provided, Oracle-managed
	// files will be used.
	// create = optional
	// update = optional
	// minItems = 3
	RedoLogs []*OracleRedoLogFileSpecificationStruct `json:"redoLogs,omitempty"`
	// The source that describes the created database instance.
	// required = true
	Source *OracleWarehouseSourceStruct `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of
	// the source.
	// required = true
	SourceConfig OracleBaseDBConfig `json:"sourceConfig,omitempty"`
	// The datafile for the SYSTEM tablespace. If no filename is
	// provided, Oracle-managed files will be used.
	// create = optional
	// update = optional
	SysDatafile *OracleSystemDatafileSpecificationStruct `json:"sysDatafile,omitempty"`
	// The password for the SYS user.
	// format = password
	// required = true
	SysPassword string `json:"sysPassword,omitempty"`
	// The datafile for the SYSAUX tablespace. If no filename is
	// provided, Oracle-managed files will be used.
	// create = optional
	// update = optional
	SysauxDatafile *OracleSysauxDatafileSpecificationStruct `json:"sysauxDatafile,omitempty"`
	// The password for the SYSTEM user.
	// format = password
	// required = true
	SystemPassword string `json:"systemPassword,omitempty"`
	// The tempfile for the database. If no filename is provided,
	// Oracle-managed files will be used.
	// create = optional
	// update = optional
	TempTablespace *OracleTempfileSpecificationStruct `json:"tempTablespace,omitempty"`
	// Indicates the timezone file version that will be used to create
	// the database.
	// create = optional
	// update = optional
	// minimum = 1
	// maximum = 22
	TimezoneFileVersion *int `json:"timezoneFileVersion,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The datafile to be used for undo data. If no filename is provided,
	// Oracle-managed files will be used.
	// create = optional
	// update = optional
	UndoTablespace *OracleUndoDatafileSpecificationStruct `json:"undoTablespace,omitempty"`
}

// OracleDatabaseStatisticStruct - A row in the database performance statistic table.
// extends TypedObject
type OracleDatabaseStatisticStruct struct {
	// A single performance statistic row.
	StatisticValues []string `json:"statisticValues,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleDatabaseStatsSectionStruct - Oracle database performance statistics for a specific section.
// extends TypedObject
type OracleDatabaseStatsSectionStruct struct {
	// List of statistic column headers.
	ColumnHeaders []string `json:"columnHeaders,omitempty"`
	// List of statistic rows corresponding to column headers.
	RowValues []*OracleDatabaseStatisticStruct `json:"rowValues,omitempty"`
	// Database statistic section name.
	SectionName string `json:"sectionName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleDeleteParametersStruct - The parameters passed in for an Oracle database delete operation.
// extends DeleteParameters
type OracleDeleteParametersStruct struct {
	// The security credential of the privileged user to run the delete
	// operation as.
	Credential Credential `json:"credential,omitempty"`
	// Flag indicating whether to continue the operation upon failures.
	Force *bool `json:"force,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the delete operation as.
	Username string `json:"username,omitempty"`
}

// OracleDisableParametersStruct - The parameters to use as input to disable oracle sources.
// extends SourceDisableParameters
type OracleDisableParametersStruct struct {
	// Whether to attempt a cleanup of the database from the environment
	// before the disable.
	// default = true
	AttemptCleanup *bool `json:"attemptCleanup,omitempty"`
	// The security credential of the privileged user to run the
	// provision operation as.
	Credential Credential `json:"credential,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
}

// OracleEnableParametersStruct - The parameters to use as input to enable Oracle sources.
// extends SourceEnableParameters
type OracleEnableParametersStruct struct {
	// Whether to attempt a startup of the source after the enable.
	// default = true
	AttemptStart *bool `json:"attemptStart,omitempty"`
	// The security credential of the privileged user to run the
	// provision operation as.
	Credential Credential `json:"credential,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
}

// OracleExportStruct - The mutable state of an Oracle database export.
// extends TypedObject
type OracleExportStruct struct {
	// DSP options for export.
	DspOptions *DSPOptionsStruct `json:"dspOptions,omitempty"`
	// Number of files to stream in parallel across the network.
	// default = 3
	// minimum = 1
	FileParallelism *int `json:"fileParallelism,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleExportParametersStruct - The parameters to use as input to export Oracle databases.
// extends DbExportParameters
type OracleExportParametersStruct struct {
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// DSP options for export.
	// create = optional
	DspOptions *DSPOptionsStruct `json:"dspOptions,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Number of files to stream in parallel across the network.
	// maximum = 64
	// create = optional
	// default = 3
	// minimum = 1
	FileParallelism *int `json:"fileParallelism,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayoutStruct `json:"filesystemLayout,omitempty"`
	// Open the database after recovery. This can have a true value only
	// if 'recoverDatabase' is true.
	// create = optional
	// default = true
	OpenDatabase *bool `json:"openDatabase,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export
	// on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleFetchedLogStruct - An Oracle log file fetched by LogSync.
// extends OracleTimeflowLog
type OracleFetchedLogStruct struct {
	// Reference to the database to which this log belongs.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-container.json
	Container string `json:"container,omitempty"`
	// End SCN for the log file.
	EndScn *int `json:"endScn,omitempty"`
	// End timestamp for the log file.
	// format = date
	EndTimestamp string `json:"endTimestamp,omitempty"`
	// Instance number associated with the log file.
	InstanceNum *int `json:"instanceNum,omitempty"`
	// Sequence number for the log file.
	Sequence *int `json:"sequence,omitempty"`
	// Start SCN for the log file.
	StartScn *int `json:"startScn,omitempty"`
	// Start timestamp for the log file.
	// format = date
	StartTimestamp string `json:"startTimestamp,omitempty"`
	// Reference to the TimeFlow of which this log is a part.
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleInstallStruct - The Oracle source repository.
// extends SourceRepository
type OracleInstallStruct struct {
	// List of Oracle patches that have been applied to this Oracle Home.
	// update = optional
	// create = optional
	AppliedPatches []*int `json:"appliedPatches,omitempty"`
	// 32 or 64 bits.
	// enum = [32 64]
	// create = required
	// update = optional
	Bits *int `json:"bits,omitempty"`
	// Flag indicating whether the install was discovered or manually
	// entered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Group ID of the user that owns the install.
	GroupId *int `json:"groupId,omitempty"`
	// Group name of the user that owns the install.
	GroupName string `json:"groupName,omitempty"`
	// The Oracle install home.
	// create = required
	// maxLength = 256
	InstallationHome string `json:"installationHome,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Flag indicating whether this repository can use LogSync.
	LogsyncPossible *bool `json:"logsyncPossible,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The Oracle base where database binaries are located.
	// create = optional
	// update = optional
	// maxLength = 256
	OracleBase string `json:"oracleBase,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// create = optional
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Flag indicating whether the install supports Oracle RAC.
	Rac *bool `json:"rac,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// create = optional
	// update = optional
	// default = false
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User ID of the user that owns the install.
	UserId *int `json:"userId,omitempty"`
	// User name of the user that owns the install.
	UserName string `json:"userName,omitempty"`
	// Version of the repository.
	// format = oracleVersion
	// create = required
	// update = optional
	Version string `json:"version,omitempty"`
}

// OracleInstanceStruct - Representation of an Oracle instance configuration.
// extends TypedObject
type OracleInstanceStruct struct {
	// The name of the instance.
	// create = required
	// update = optional
	// pattern = ^[a-zA-Z0-9_]+$
	// maxLength = 15
	InstanceName string `json:"instanceName,omitempty"`
	// The number of the instance.
	// create = required
	// update = optional
	// minimum = 1
	InstanceNumber float64 `json:"instanceNumber,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleLinkDataStruct - Represents parameters to link non-pluggable Oracle databases.
// extends OracleBaseLinkData
type OracleLinkDataStruct struct {
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit *int `json:"bandwidthLimit,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// True if SnapSync data from the source should be compressed over
	// the network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially
	// over slow network.
	// default = true
	// create = optional
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// If true, NOLOGGING operations on this container are treated as
	// faults and cannot be resolved manually. Otherwise, these
	// operations are ignored.
	// create = optional
	// default = true
	DiagnoseNoLoggingFaults *bool `json:"diagnoseNoLoggingFaults,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession
	// to reduce the number of logs required to provision the snapshot.
	// This may significantly reduce the time necessary to provision from
	// a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// True if SnapSync data from the source should be retrieved through
	// an encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// Information about the OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	FilesPerSet *int `json:"filesPerSet,omitempty"`
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// Non-SYS database credentials to access this database.
	// create = optional
	NonSysCredentials *PasswordCredentialStruct `json:"nonSysCredentials,omitempty"`
	// Non-SYS database user to access this database.
	// maxLength = 30
	// create = optional
	NonSysUser string `json:"nonSysUser,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// default = 1
	// minimum = 1
	// maximum = 16
	// create = optional
	NumberOfConnections *int `json:"numberOfConnections,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// If true, pre-provisioning will be performed after every sync.
	// default = false
	// create = optional
	PreProvisioningEnabled *bool `json:"preProvisioningEnabled,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	RmanChannels *int `json:"rmanChannels,omitempty"`
	// Skip check that tests if there is enough space available to store
	// the database in the Delphix Engine. The Delphix Engine estimates
	// how much space a database will occupy after compression and
	// prevents SnapSync if insufficient space is available. This
	// safeguard can be overridden using this option. This may be useful
	// when linking highly compressible databases.
	// default = false
	// create = optional
	SkipSpaceCheck *bool `json:"skipSpaceCheck,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *OracleSourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleLinkedSourceStruct - A linked Oracle source.
// extends OracleSource
type OracleLinkedSourceStruct struct {
	// Defines whether backup level is enabled.
	// create = optional
	// update = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// create = optional
	// update = optional
	// minimum = 0
	// default = 0
	BandwidthLimit *int `json:"bandwidthLimit,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	// update = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// True if SnapSync data from the source should be compressed over
	// the network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially
	// over slow network.
	// create = optional
	// update = optional
	// default = true
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-oracle-base-db-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// True if SnapSync data from the source should be retrieved through
	// an encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// create = optional
	// update = optional
	// default = false
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// External file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	// update = optional
	FilesPerSet *int `json:"filesPerSet,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// default = 1
	// minimum = 1
	// maximum = 16
	// create = optional
	// update = optional
	NumberOfConnections *int `json:"numberOfConnections,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	// update = optional
	RmanChannels *int `json:"rmanChannels,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// OracleLiveSourceStruct - An Oracle LiveSource.
// extends OracleVirtualSource
type OracleLiveSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Archive Log Mode of the Oracle virtual database.
	// create = optional
	// update = readonly
	// default = true
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Oracle database configuration parameter overrides.
	// update = optional
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	// create = optional
	// update = optional
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Custom environment variables for Oracle databases.
	// create = optional
	// update = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// Amount of tolerable delay for this Oracle LiveSource in seconds.
	// units = sec
	// create = optional
	// update = optional
	// default = 900
	DataAgeWarningThreshold *int `json:"dataAgeWarningThreshold,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// create = optional
	// update = optional
	// default = false
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A list of object references to Oracle Node Listeners selected for
	// this Managed Database. Delphix picks one default listener from the
	// target environment if this list is empty at virtual database
	// provision time.
	// update = optional
	// create = optional
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// Number of Online Redo Log Groups.
	// create = optional
	// update = readonly
	// minimum = 2
	// default = 3
	RedoLogGroups *int `json:"redoLogGroups,omitempty"`
	// Online Redo Log size in MB.
	// update = readonly
	// minimum = 4
	// create = optional
	RedoLogSizeInMB *int `json:"redoLogSizeInMB,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Resync status for this Oracle LiveSource.
	// enum = [RESYNC_NOT_REQUIRED RESYNC_NEEDED RESYNC_IN_PROGRESS APPLY_READY APPLY_IN_PROGRESS APPLY_FAILED]
	ResyncStatus string `json:"resyncStatus,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// OracleLiveSourceRuntimeStruct - Runtime (non-persistent) properties of an Oracle LiveSource.
// extends OracleBaseSourceRuntime
type OracleLiveSourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// List of active database instances for the source.
	ActiveInstances []*OracleActiveInstanceStruct `json:"activeInstances,omitempty"`
	// MRP apply status for the standby database associated with the
	// LiveSource.
	// enum = [UNKNOWN WORKING APPLY_FAILED APPLY_ON_WRONG_INCARNATION UNRESOLVABLE_GAP_DETECTED]
	ApplyStatus string `json:"applyStatus,omitempty"`
	// Current data lag between LiveSource and source database in
	// seconds.
	// units = sec
	CurrentDataAge *int `json:"currentDataAge,omitempty"`
	// Operating mode of the database.
	// enum = [READ_WRITE READ_ONLY STANDBY_READ_ONLY MOUNTED_ONLY UNKNOWN]
	// default = UNKNOWN
	DatabaseMode string `json:"databaseMode,omitempty"`
	// The current role of the database.
	// default = UNKNOWN
	// enum = [PHYSICAL_STANDBY LOGICAL_STANDBY SNAPSHOT_STANDBY FAR_SYNC PRIMARY UNKNOWN]
	DatabaseRole string `json:"databaseRole,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Table of key database performance statistics.
	DatabaseStats []*OracleDatabaseStatsSectionStruct `json:"databaseStats,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Has data age exceeded the user specified threshold.
	IsDataAgeWarningExceeded *bool `json:"isDataAgeWarningExceeded,omitempty"`
	// Highest SCN at which non-logged changes were generated.
	LastNonLoggedLocation string `json:"lastNonLoggedLocation,omitempty"`
	// The time at which this runtime data was updated.
	// format = date
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Indicates whether there is non-logged data on the standby.
	NonLoggedDataDetected *bool `json:"nonLoggedDataDetected,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Time zone of the source database at the time the runtime data was
	// updated.
	SourceDatabaseTimezone string `json:"sourceDatabaseTimezone,omitempty"`
	// Indicates whether the incarnation ID changed on the primary
	// database.
	SourceResetlogsIDChangeDetected *bool `json:"sourceResetlogsIDChangeDetected,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Redo log transport status from the source database to the
	// LiveSource.
	// enum = [UNKNOWN WORKING NO_INITIAL_DATA NO_NEW_DATA]
	TransportStatus string `json:"transportStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Indicates whether the LiveSource is not in standby mode.
	UnexpectedRoleChangeDetected *bool `json:"unexpectedRoleChangeDetected,omitempty"`
}

// OracleLogStruct - Oracle log file.
// extends TypedObject
type OracleLogStruct struct {
	// Instance number associated with the log file.
	InstanceNum *int `json:"instanceNum,omitempty"`
	// Sequence number for the log file.
	Sequence *int `json:"sequence,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleMissingLogStruct - An Oracle missing log file.
// extends OracleTimeflowLog
type OracleMissingLogStruct struct {
	// Reference to the database to which this log belongs.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-container.json
	Container string `json:"container,omitempty"`
	// End SCN for the log file.
	EndScn *int `json:"endScn,omitempty"`
	// End timestamp for the log file.
	// format = date
	EndTimestamp string `json:"endTimestamp,omitempty"`
	// Instance number associated with the log file.
	InstanceNum *int `json:"instanceNum,omitempty"`
	// Sequence number for the log file.
	Sequence *int `json:"sequence,omitempty"`
	// Start SCN for the log file.
	StartScn *int `json:"startScn,omitempty"`
	// Start timestamp for the log file.
	// format = date
	StartTimestamp string `json:"startTimestamp,omitempty"`
	// Reference to the TimeFlow of which this log is a part.
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleMultitenantProvisionParametersStruct - The parameters to use as input to provision Oracle multitenant
// databases.
// extends OracleBaseProvisionParameters
type OracleMultitenantProvisionParametersStruct struct {
	// The new container for the provisioned database.
	// required = true
	Container *OracleDatabaseContainerStruct `json:"container,omitempty"`
	// The security credential of the privileged user to run the
	// provision operation as.
	// create = optional
	Credential Credential `json:"credential,omitempty"`
	// Whether or not to mark this VDB as a masked VDB. It will be marked
	// as masked if this flag or the masking job are set.
	// update = readonly
	// create = optional
	Masked *bool `json:"masked,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	// format = objectReference
	MaskingJob string `json:"maskingJob,omitempty"`
	// The pluggable database source that describes an external database
	// instance.
	// required = true
	Source *OracleVirtualPdbSourceStruct `json:"source,omitempty"`
	// The pluggable database source config including dynamically
	// discovered attributes of the source.
	// required = true
	SourceConfig *OraclePDBConfigStruct `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the provision operation as.
	// create = optional
	Username string `json:"username,omitempty"`
	// The new container for the created dataset.
	// create = optional
	VirtualCdb *OracleVirtualCdbProvisionParametersStruct `json:"virtualCdb,omitempty"`
}

// OracleNodeListenerStruct - An Oracle node listener.
// extends OracleListener
type OracleNodeListenerStruct struct {
	// The list of client endpoints for this listener of the format
	// hostname:port. These are used when constructing the JDBC
	// connection string.
	// update = readonly
	// create = readonly
	ClientEndpoints []string `json:"clientEndpoints,omitempty"`
	// Whether this listener was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the environment this listener is associated with.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Reference to the host this listener is associated with.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-host.json
	Host string `json:"host,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The list of protocol addresses for this listener. These are used
	// for the local_listener parameter when provisioning VDBs.
	// minItems = 1
	// create = required
	// update = optional
	ProtocolAddresses []string `json:"protocolAddresses,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OraclePDBAttachDataStruct - Represents parameters to attach an Oracle pluggable database.
// extends OracleBaseAttachData
type OraclePDBAttachDataStruct struct {
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit *int `json:"bandwidthLimit,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// True if SnapSync data from the source should be compressed over
	// the network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially
	// over slow network.
	// default = true
	// create = optional
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Reference to the configuration for the PDB source.
	// referenceTo = /delphix-oracle-pdb-config.json
	// required = true
	// format = objectReference
	Config string `json:"config,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// True if SnapSync data from the source should be retrieved through
	// an encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// Information about the OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// default = 5
	// create = optional
	// minimum = 1
	// maximum = 64
	FilesPerSet *int `json:"filesPerSet,omitempty"`
	// True if attach should succeed even if the resetlogs of the
	// original database does not match the resetlogs of the new database
	// and the resetlogs information of the original database is not a
	// parent incarnation of the current resetlogs. This can happen when
	// the controlfile has been recreated and the incarnation table is
	// incomplete. Use this option with extreme caution. Attached
	// database must be the same database to avoid data corruption later
	// on.
	// default = false
	// create = optional
	Force *bool `json:"force,omitempty"`
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// default = 1
	// minimum = 1
	// maximum = 16
	// create = optional
	NumberOfConnections *int `json:"numberOfConnections,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	RmanChannels *int `json:"rmanChannels,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OraclePDBConfigStruct - Representation of properties for an Oracle pluggable database
// configuration.
// extends OracleBaseDBConfig
type OraclePDBConfigStruct struct {
	// The DB config of an Oracle multitenant database this pluggable
	// database belongs to.
	// referenceTo = /delphix-oracle-db-config.json
	// create = optional
	// update = optional
	// format = objectReference
	CdbConfig string `json:"cdbConfig,omitempty"`
	// The password of the database user. This must be a
	// PasswordCredential instance.
	// update = optional
	Credentials *PasswordCredentialStruct `json:"credentials,omitempty"`
	// The name of the database.
	// maxLength = 30
	// create = required
	// pattern = ^[a-zA-Z0-9][a-zA-Z0-9_]*$
	DatabaseName string `json:"databaseName,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = optional
	Repository string `json:"repository,omitempty"`
	// The list of database services.
	// create = optional
	// update = optional
	Services []*OracleServiceStruct `json:"services,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 30
	User string `json:"user,omitempty"`
}

// OraclePDBLinkDataStruct - Represents parameters to link a Oracle pluggable database.
// extends OracleBaseLinkData
type OraclePDBLinkDataStruct struct {
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit *int `json:"bandwidthLimit,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// True if SnapSync data from the source should be compressed over
	// the network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially
	// over slow network.
	// create = optional
	// default = true
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-oracle-pdb-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// If true, NOLOGGING operations on this container are treated as
	// faults and cannot be resolved manually. Otherwise, these
	// operations are ignored.
	// default = true
	// create = optional
	DiagnoseNoLoggingFaults *bool `json:"diagnoseNoLoggingFaults,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession
	// to reduce the number of logs required to provision the snapshot.
	// This may significantly reduce the time necessary to provision from
	// a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// True if SnapSync data from the source should be retrieved through
	// an encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// Information about the OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	FilesPerSet *int `json:"filesPerSet,omitempty"`
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// maximum = 16
	// create = optional
	// default = 1
	// minimum = 1
	NumberOfConnections *int `json:"numberOfConnections,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// If true, pre-provisioning will be performed after every sync.
	// default = false
	// create = optional
	PreProvisioningEnabled *bool `json:"preProvisioningEnabled,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	RmanChannels *int `json:"rmanChannels,omitempty"`
	// Skip check that tests if there is enough space available to store
	// the database in the Delphix Engine. The Delphix Engine estimates
	// how much space a database will occupy after compression and
	// prevents SnapSync if insufficient space is available. This
	// safeguard can be overridden using this option. This may be useful
	// when linking highly compressible databases.
	// create = optional
	// default = false
	SkipSpaceCheck *bool `json:"skipSpaceCheck,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *OracleSourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OraclePDBSourceRuntimeStruct - Runtime (non-persistent) properties of an Oracle PDB source.
// extends OracleBaseSourceRuntime
type OraclePDBSourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// List of active database instances for the source.
	ActiveInstances []*OracleActiveInstanceStruct `json:"activeInstances,omitempty"`
	// Operating mode of the database.
	// default = UNKNOWN
	// enum = [READ_WRITE READ_ONLY STANDBY_READ_ONLY MOUNTED_ONLY UNKNOWN]
	DatabaseMode string `json:"databaseMode,omitempty"`
	// The current role of the database.
	// enum = [PHYSICAL_STANDBY LOGICAL_STANDBY SNAPSHOT_STANDBY FAR_SYNC PRIMARY UNKNOWN]
	// default = UNKNOWN
	DatabaseRole string `json:"databaseRole,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Table of key database performance statistics.
	DatabaseStats []*OracleDatabaseStatsSectionStruct `json:"databaseStats,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Highest SCN at which non-logged changes were generated.
	LastNonLoggedLocation string `json:"lastNonLoggedLocation,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OraclePlatformParametersStruct - Oracle platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type OraclePlatformParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleProvisionParametersStruct - The parameters to use as input to provision Oracle (non-multitenant)
// databases.
// extends OracleBaseProvisionParameters
type OracleProvisionParametersStruct struct {
	// The new container for the provisioned database.
	// required = true
	Container *OracleDatabaseContainerStruct `json:"container,omitempty"`
	// The security credential of the privileged user to run the
	// provision operation as.
	// create = optional
	Credential Credential `json:"credential,omitempty"`
	// Whether or not to mark this VDB as a masked VDB. It will be marked
	// as masked if this flag or the masking job are set.
	// create = optional
	// update = readonly
	Masked *bool `json:"masked,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
	// Flag indicating whether to generate a new DBID for the provisioned
	// database.
	// create = optional
	// default = false
	NewDBID *bool `json:"newDBID,omitempty"`
	// Flag indicating whether to open the database after provision.
	// create = optional
	// default = true
	OpenResetlogs *bool `json:"openResetlogs,omitempty"`
	// Flag indicating whether the virtual database is provisioned as a
	// physical standby database.
	// create = optional
	// default = false
	PhysicalStandby *bool `json:"physicalStandby,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *OracleVirtualSourceStruct `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of
	// the source.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the provision operation as.
	// create = optional
	Username string `json:"username,omitempty"`
}

// OracleRACConfigStruct - Representation of the properties specific to a RAC Oracle DB
// configuration.
// extends OracleDBConfig
type OracleRACConfigStruct struct {
	// The container type of this database.
	// create = readonly
	// update = readonly
	// enum = [UNKNOWN ROOT_CDB NON_CDB AUX_CDB]
	CdbType string `json:"cdbType,omitempty"`
	// The password of the database user. This must be a
	// PasswordCredential instance.
	// update = optional
	Credentials *PasswordCredentialStruct `json:"credentials,omitempty"`
	// The Oracle Clusterware database name.
	// create = optional
	// maxLength = 30
	CrsDatabaseName string `json:"crsDatabaseName,omitempty"`
	// The name of the database.
	// create = required
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 8
	DatabaseName string `json:"databaseName,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The list of RAC instances for this RAC configuration.
	// create = required
	// update = optional
	Instances []*OracleRACInstanceStruct `json:"instances,omitempty"`
	// Whether this source should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The password of a database user that does not have administrative
	// privileges.
	// create = optional
	// update = optional
	NonSysCredentials string `json:"nonSysCredentials,omitempty"`
	// The username of a database user that does not have administrative
	// privileges.
	// create = optional
	// update = optional
	// maxLength = 30
	NonSysUser string `json:"nonSysUser,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-oracle-install.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The list of database services.
	// create = optional
	// update = optional
	Services []*OracleServiceStruct `json:"services,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The unique name.
	// create = required
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 30
	UniqueName string `json:"uniqueName,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 30
	User string `json:"user,omitempty"`
}

// OracleRACInstanceStruct - The representation of an Oracle Database RAC Instance Configuration.
// extends OracleInstance
type OracleRACInstanceStruct struct {
	// The name of the instance.
	// maxLength = 15
	// create = required
	// update = optional
	// pattern = ^[a-zA-Z0-9_]+$
	InstanceName string `json:"instanceName,omitempty"`
	// The number of the instance.
	// create = required
	// update = optional
	// minimum = 1
	InstanceNumber float64 `json:"instanceNumber,omitempty"`
	// Reference to the Oracle cluster node that the RAC instance is
	// running on.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster-node.json
	// create = required
	Node string `json:"node,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleRACSourceConnectionInfoStruct - Contains information that can be used to connect to a single instance
// Oracle source.
// extends OracleSourceConnectionInfo
type OracleRACSourceConnectionInfoStruct struct {
	// The location of the cluster installation.
	CrsClusterHome string `json:"crsClusterHome,omitempty"`
	// The database name.
	DatabaseName string `json:"databaseName,omitempty"`
	// The JDBC strings used to connect to the source.
	JdbcStrings []string `json:"jdbcStrings,omitempty"`
	// The addresses for the nodes on which the source resides.
	Nodes []string `json:"nodes,omitempty"`
	// The Oracle installation home.
	OracleHome string `json:"oracleHome,omitempty"`
	// The default remote_listener parameter to be used for databases on
	// the cluster.
	RemoteListener string `json:"remoteListener,omitempty"`
	// The Single Client Access Name of the cluster (11.2 and greater
	// clusters only).
	Scan string `json:"scan,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// OracleRedoLogFileSpecificationStruct - Describes an Oracle redo log file.
// extends OracleFileSpecification
type OracleRedoLogFileSpecificationStruct struct {
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
	// The size of the log file in MB.
	// minimum = 100
	// default = 100
	// create = optional
	// update = optional
	Size *int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleRefreshParametersStruct - The parameters to use as input to refresh Oracle databases.
// extends RefreshParameters
type OracleRefreshParametersStruct struct {
	// The security credential of the privileged user to run the refresh
	// operation as.
	Credential Credential `json:"credential,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to refresh the
	// database to.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the refresh operation as.
	Username string `json:"username,omitempty"`
}

// OracleRollbackParametersStruct - The parameters to use as input to roll back Oracle databases.
// extends RollbackParameters
type OracleRollbackParametersStruct struct {
	// The security credential of the user who has the required
	// privileges to run the rollback operation.
	Credential Credential `json:"credential,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to roll the
	// database back to.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the user who has the required privileges to perform
	// the rollback operation.
	Username string `json:"username,omitempty"`
}

// OracleSIConfigStruct - The representation of a single-instance Oracle DB configuration.
// extends OracleDBConfig
type OracleSIConfigStruct struct {
	// The container type of this database.
	// enum = [UNKNOWN ROOT_CDB NON_CDB AUX_CDB]
	// create = readonly
	// update = readonly
	CdbType string `json:"cdbType,omitempty"`
	// The password of the database user. This must be a
	// PasswordCredential instance.
	// update = optional
	Credentials *PasswordCredentialStruct `json:"credentials,omitempty"`
	// The name of the database.
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 8
	// create = required
	DatabaseName string `json:"databaseName,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The Oracle instance.
	// create = required
	// update = optional
	Instance *OracleInstanceStruct `json:"instance,omitempty"`
	// Whether this source should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The password of a database user that does not have administrative
	// privileges.
	// update = optional
	// create = optional
	NonSysCredentials string `json:"nonSysCredentials,omitempty"`
	// The username of a database user that does not have administrative
	// privileges.
	// create = optional
	// update = optional
	// maxLength = 30
	NonSysUser string `json:"nonSysUser,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// referenceTo = /delphix-oracle-install.json
	// create = required
	// update = optional
	// format = objectReference
	Repository string `json:"repository,omitempty"`
	// The list of database services.
	// create = optional
	// update = optional
	Services []*OracleServiceStruct `json:"services,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The unique name.
	// create = required
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 30
	UniqueName string `json:"uniqueName,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 30
	User string `json:"user,omitempty"`
}

// OracleSISourceConnectionInfoStruct - Contains information that can be used to connect to a single instance
// Oracle source.
// extends OracleSourceConnectionInfo
type OracleSISourceConnectionInfoStruct struct {
	// The database name.
	DatabaseName string `json:"databaseName,omitempty"`
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The JDBC strings used to connect to the source.
	JdbcStrings []string `json:"jdbcStrings,omitempty"`
	// The Oracle installation home.
	OracleHome string `json:"oracleHome,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// OracleScanListenerStruct - An Oracle scan listener.
// extends OracleListener
type OracleScanListenerStruct struct {
	// The list of client endpoints for this listener of the format
	// hostname:port. These are used when constructing the JDBC
	// connection string.
	// create = readonly
	// update = readonly
	ClientEndpoints []string `json:"clientEndpoints,omitempty"`
	// Whether this listener was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the environment this listener is associated with.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The list of protocol addresses for this listener. These are used
	// for the local_listener parameter when provisioning VDBs.
	// create = required
	// update = optional
	// minItems = 1
	ProtocolAddresses []string `json:"protocolAddresses,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleServiceStruct - The representation of an oracle service object.
// extends TypedObject
type OracleServiceStruct struct {
	// Whether this service was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The connection string used to connect to JDBC.
	// format = oracleJDBCConnectionString
	// create = required
	// update = optional
	JdbcConnectionString string `json:"jdbcConnectionString,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleSnapshotStruct - Provisionable snapshot of an Oracle TimeFlow.
// extends TimeflowSnapshot
type OracleSnapshotStruct struct {
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *OracleTimeflowPointStruct `json:"firstChangePoint,omitempty"`
	// Auxiliary TimeFlows with snapshots controlled by this master
	// snapshot.
	FractionTimeflows []string `json:"fractionTimeflows,omitempty"`
	// True if this snapshot was taken of a standby database.
	// default = false
	// create = optional
	FromPhysicalStandbyVdb *bool `json:"fromPhysicalStandbyVdb,omitempty"`
	// The location of the snapshot within the parent TimeFlow
	// represented by this snapshot.
	LatestChangePoint *OracleTimeflowPointStruct `json:"latestChangePoint,omitempty"`
	// Boolean value indicating if a virtual database provisioned from
	// this snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Online redo log size in bytes when this snapshot was taken.
	// create = optional
	RedoLogSizeInBytes float64 `json:"redoLogSizeInBytes,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot
	// should be kept forever.
	// update = optional
	Retention *int `json:"retention,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *OracleSnapshotRuntimeStruct `json:"runtime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient
	// state and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Time zone of the source database at the time the snapshot was
	// taken.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
}

// OracleSnapshotRuntimeStruct - Runtime (non-persistent) properties of an Oracle TimeFlow snapshot.
// extends SnapshotRuntime
type OracleSnapshotRuntimeStruct struct {
	// List of missing log files for this snapshot, if any.
	MissingLogs []*OracleLogStruct `json:"missingLogs,omitempty"`
	// True if this snapshot can be used as the basis for provisioning a
	// new TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleSourceRuntimeStruct - Runtime (non-persistent) properties of an Oracle source.
// extends OracleBaseSourceRuntime
type OracleSourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// List of active database instances for the source.
	ActiveInstances []*OracleActiveInstanceStruct `json:"activeInstances,omitempty"`
	// True if the database is running in ARCHIVELOG mode.
	ArchivelogEnabled *bool `json:"archivelogEnabled,omitempty"`
	// True if block change tracking is enabled.
	BctEnabled *bool `json:"bctEnabled,omitempty"`
	// Operating mode of the database.
	// default = UNKNOWN
	// enum = [READ_WRITE READ_ONLY STANDBY_READ_ONLY MOUNTED_ONLY UNKNOWN]
	DatabaseMode string `json:"databaseMode,omitempty"`
	// The current role of the database.
	// default = UNKNOWN
	// enum = [PHYSICAL_STANDBY LOGICAL_STANDBY SNAPSHOT_STANDBY FAR_SYNC PRIMARY UNKNOWN]
	DatabaseRole string `json:"databaseRole,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Table of key database performance statistics.
	DatabaseStats []*OracleDatabaseStatsSectionStruct `json:"databaseStats,omitempty"`
	// True if the database has Oracle Direct NFS client enabled.
	DnfsEnabled *bool `json:"dnfsEnabled,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Highest SCN at which non-logged changes were generated.
	LastNonLoggedLocation string `json:"lastNonLoggedLocation,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// True for a RAC source database.
	RacEnabled *bool `json:"racEnabled,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleSourcingPolicyStruct - Database policies for managing SnapSync and LogSync across sources for
// an Oracle container.
// extends SourcingPolicy
type OracleSourcingPolicyStruct struct {
	// True if LogSync should run for this database.
	// create = optional
	// update = optional
	// default = false
	LogsyncEnabled *bool `json:"logsyncEnabled,omitempty"`
	// Interval between LogSync requests, in seconds.
	// create = optional
	// update = optional
	// units = sec
	// default = 5
	LogsyncInterval *int `json:"logsyncInterval,omitempty"`
	// LogSync operation mode for this database.
	// enum = [ARCHIVE_ONLY_MODE ARCHIVE_REDO_MODE UNDEFINED]
	// create = optional
	// update = optional
	// default = UNDEFINED
	LogsyncMode string `json:"logsyncMode,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleStagingSourceStruct - A staging Oracle source used for Delphix operations such as log
// collection and snapshot generation.
// extends OracleManagedSource
type OracleStagingSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// referenceTo = /delphix-database-template.json
	// create = optional
	// update = optional
	// format = objectReference
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Reference to the container being fed by this source, if any.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// create = optional
	// update = optional
	// default = false
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 256
	// create = required
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A list of object references to Oracle Node Listeners selected for
	// this Managed Database. Delphix picks one default listener from the
	// target environment if this list is empty at virtual database
	// provision time.
	// create = optional
	// update = optional
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// OracleStagingSourceParametersStruct - Parameters provided by the user to create a staging database.
// extends TypedObject
type OracleStagingSourceParametersStruct struct {
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional database template to use for staging database creation.
	// If set, configParams will be ignored.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 256
	// create = required
	MountBase string `json:"mountBase,omitempty"`
	// The object reference of the source repository that will host the
	// LogSync staging database.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-oracle-install.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleStartParametersStruct - The parameters to use as input to start oracle sources.
// extends SourceStartParameters
type OracleStartParametersStruct struct {
	// The security credential of the privileged user to run the
	// provision operation as.
	Credential Credential `json:"credential,omitempty"`
	// List of specific Oracle instances to start.
	Instances []float64 `json:"instances,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
}

// OracleStopParametersStruct - The parameters to use as input to stop oracle sources.
// extends SourceStopParameters
type OracleStopParametersStruct struct {
	// Whether to issue 'shutdown abort' to shutdown Oracle instances.
	// default = false
	Abort *bool `json:"abort,omitempty"`
	// The security credential of the privileged user to run the
	// provision operation as.
	Credential Credential `json:"credential,omitempty"`
	// List of specific Oracle instances to stop.
	Instances []float64 `json:"instances,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
}

// OracleSyncParametersStruct - The parameters to use as input to sync Oracle databases.
// extends SyncParameters
type OracleSyncParametersStruct struct {
	// Indicates whether a fresh SnapSync must be started regardless if
	// it was possible to resume the current SnapSync. If true, we will
	// not resume but instead ignore previous progress and backup all
	// datafiles even if already completed from previous failed SnapSync.
	// This does not force a full backup, if an incremental was in
	// progress this will start a new incremental snapshot.
	// create = optional
	// default = false
	DoNotResume *bool `json:"doNotResume,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession
	// to reduce the number of logs required to provision the snapshot.
	// This may significantly reduce the time necessary to provision from
	// a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// Whether or not to take another full backup of the source database.
	// default = false
	ForceFullBackup *bool `json:"forceFullBackup,omitempty"`
	// Skip check that tests if there is enough space available to store
	// the database in the Delphix Engine. The Delphix Engine estimates
	// how much space a database will occupy after compression and
	// prevents SnapSync if insufficient space is available. This
	// safeguard can be overridden using this option. This may be useful
	// when linking highly compressible databases.
	// default = false
	// create = optional
	SkipSpaceCheck *bool `json:"skipSpaceCheck,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleSysauxDatafileSpecificationStruct - Describes an Oracle datafile in the SYSAUX tablespace.
// extends OracleDatafileTempfileSpecification
type OracleSysauxDatafileSpecificationStruct struct {
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// create = optional
	// update = optional
	// default = true
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the
	// size of one data block.
	// update = optional
	// create = optional
	AutoExtendIncrement *int `json:"autoExtendIncrement,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// create = optional
	// update = optional
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	Filename string `json:"filename,omitempty"`
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space
	// that Oracle can allocate to the datafile or tempfile.
	// create = optional
	// update = optional
	MaxSize *int `json:"maxSize,omitempty"`
	// The size of the file in MB.
	// default = 900
	// create = optional
	// update = optional
	// minimum = 900
	Size *int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleSystemDatafileSpecificationStruct - Describes an Oracle datafile in the SYSTEM tablespace.
// extends OracleDatafileTempfileSpecification
type OracleSystemDatafileSpecificationStruct struct {
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// create = optional
	// update = optional
	// default = true
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the
	// size of one data block.
	// create = optional
	// update = optional
	AutoExtendIncrement *int `json:"autoExtendIncrement,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space
	// that Oracle can allocate to the datafile or tempfile.
	// update = optional
	// create = optional
	MaxSize *int `json:"maxSize,omitempty"`
	// The size of the file in MB.
	// create = optional
	// update = optional
	// minimum = 700
	// default = 700
	Size *int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleTempfileSpecificationStruct - Describes an Oracle temporary file.
// extends OracleDatafileTempfileSpecification
type OracleTempfileSpecificationStruct struct {
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// create = optional
	// update = optional
	// default = true
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the
	// size of one data block.
	// create = optional
	// update = optional
	AutoExtendIncrement *int `json:"autoExtendIncrement,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// create = optional
	// update = optional
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	Filename string `json:"filename,omitempty"`
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space
	// that Oracle can allocate to the datafile or tempfile.
	// create = optional
	// update = optional
	MaxSize *int `json:"maxSize,omitempty"`
	// The size of the file in MB.
	// create = optional
	// update = optional
	// minimum = 300
	// default = 300
	Size *int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleTimeflowStruct - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends Timeflow
type OracleTimeflowStruct struct {
	// Reference to the mirror CDB TimeFlow if this is a PDB TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	CdbTimeflow string `json:"cdbTimeflow,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC SOURCE_CONTINUITY]
	CreationType string `json:"creationType,omitempty"`
	// Oracle-specific incarnation identifier for this TimeFlow.
	IncarnationID string `json:"incarnationID,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow
	// was provisioned. This will not be present for TimeFlows derived
	// from linked sources.
	ParentPoint *OracleTimeflowPointStruct `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning
	// base for this object. This may be different from the snapshot
	// within the parent point, and is only present for virtual
	// TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Set to true if the TimeFlow represents a warehouse.
	// featureFlag = CONSPRO
	Warehouse *bool `json:"warehouse,omitempty"`
	// Reference to the TimeFlow of the warehouse that this TimeFlow is a
	// part of.
	// featureFlag = CONSPRO
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	WarehouseTimeflow string `json:"warehouseTimeflow,omitempty"`
}

// OracleTimeflowPointStruct - A unique point within an Oracle database TimeFlow.
// extends TimeflowPoint
type OracleTimeflowPointStruct struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this point.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// create = optional
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleUndoDatafileSpecificationStruct - Describes an Oracle datafile that stores undo data.
// extends OracleDatafileTempfileSpecification
type OracleUndoDatafileSpecificationStruct struct {
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// create = optional
	// update = optional
	// default = true
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the
	// size of one data block.
	// create = optional
	// update = optional
	AutoExtendIncrement *int `json:"autoExtendIncrement,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space
	// that Oracle can allocate to the datafile or tempfile.
	// create = optional
	// update = optional
	MaxSize *int `json:"maxSize,omitempty"`
	// The size of the file in MB.
	// default = 300
	// create = optional
	// update = optional
	// minimum = 300
	Size *int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleVirtualCdbProvisionParametersStruct - The parameters to use as input to provision Oracle virtual container
// databases.
// extends TypedObject
type OracleVirtualCdbProvisionParametersStruct struct {
	// The new container for the created database.
	// required = true
	Container *OracleDatabaseContainerStruct `json:"container,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *OracleVirtualCdbSourceStruct `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of
	// the source.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleVirtualCdbSourceStruct - A virtual Oracle multitenant container database source. Certain fields
// of the Oracle virtual source are not applicable to the virtual
// container database.
// extends OracleVirtualSource
type OracleVirtualCdbSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Archive Log Mode of the Oracle virtual database. NOARCHIVELOG mode
	// is currently not supported for virtual container databases.
	// default = true
	// create = readonly
	// update = readonly
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	// create = optional
	// update = optional
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Reference to the container being fed by this source, if any.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// Custom environment variables for Oracle databases. These can only
	// be set at the PDB level. Delphix applies the PDB setting to the
	// virtual CDB.
	// create = readonly
	// update = readonly
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules. These can only be set at the PDB
	// level. Delphix applies the PDB setting to the virtual CDB.
	// create = readonly
	// update = readonly
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// update = optional
	// default = false
	// create = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A list of object references to Oracle Node Listeners selected for
	// this Virtual Database. Delphix picks one default listener from the
	// target environment if this list is empty at virtual database
	// provision time. This is currently not supported for virtual
	// container databases.
	// create = readonly
	// update = readonly
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// User-specified operation hooks for this source. This is currently
	// not supported for virtual container databases.
	// create = readonly
	// update = readonly
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// Number of Online Redo Log Groups. Customizing number of Online
	// Redo Log Groups is currently not supported for virtual container
	// databases.
	// create = readonly
	// update = readonly
	RedoLogGroups *int `json:"redoLogGroups,omitempty"`
	// Online Redo Log size in MB. Customizing Online Redo Log size is
	// currently not supported for virtual container databases.
	// create = readonly
	// update = readonly
	RedoLogSizeInMB *int `json:"redoLogSizeInMB,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// OracleVirtualIPStruct - The parameters used for virtual IP operations.
// extends TypedObject
type OracleVirtualIPStruct struct {
	// A boolean indicating whether this VIP was automatically
	// discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The name of the domain where the cluster is residing.
	// required = true
	DomainName string `json:"domainName,omitempty"`
	// The virtual IP address.
	// format = ipv4Address
	// required = true
	Ip string `json:"ip,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleVirtualPdbSourceStruct - A virtual Oracle multitenant pluggable database source.
// extends OracleVirtualSource
type OracleVirtualPdbSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Archive Log Mode of the Oracle virtual database. This is not
	// applicable to pluggable databases.
	// create = readonly
	// update = readonly
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Oracle database configuration parameter overrides. This is
	// currently not supported for pluggable databases.
	// create = readonly
	// update = readonly
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh. This is
	// currently not supported for pluggable databases.
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	// create = readonly
	// update = readonly
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Custom environment variables for Oracle databases.
	// create = optional
	// update = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A list of object references to Oracle Node Listeners selected for
	// this Virtual Database. Delphix picks one default listener from the
	// target environment if this list is empty at virtual database
	// provision time. This is not applicable to pluggable databases.
	// create = readonly
	// update = readonly
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// Number of Online Redo Log Groups. This is not applicable to
	// pluggable databases.
	// update = readonly
	// create = readonly
	RedoLogGroups *int `json:"redoLogGroups,omitempty"`
	// Online Redo Log size in MB. This is not applicable to pluggable
	// databases.
	// create = readonly
	// update = readonly
	RedoLogSizeInMB *int `json:"redoLogSizeInMB,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// OracleVirtualSourceStruct - A virtual Oracle source.
// extends OracleManagedSource
type OracleVirtualSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Archive Log Mode of the Oracle virtual database.
	// create = optional
	// update = readonly
	// default = true
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// referenceTo = /delphix-database-template.json
	// create = optional
	// update = optional
	// format = objectReference
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Custom environment variables for Oracle databases.
	// create = optional
	// update = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A list of object references to Oracle Node Listeners selected for
	// this Managed Database. Delphix picks one default listener from the
	// target environment if this list is empty at virtual database
	// provision time.
	// create = optional
	// update = optional
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// Number of Online Redo Log Groups.
	// update = readonly
	// minimum = 2
	// default = 3
	// create = optional
	RedoLogGroups *int `json:"redoLogGroups,omitempty"`
	// Online Redo Log size in MB.
	// create = optional
	// update = readonly
	// minimum = 4
	RedoLogSizeInMB *int `json:"redoLogSizeInMB,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// OracleWarehouseSourceStruct - A warehouse represents a housing that accepts databases.
// extends OracleVirtualSource
type OracleWarehouseSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// update = optional
	// create = required
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Archive Log Mode of the Oracle virtual database.
	// create = optional
	// update = readonly
	// default = true
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	// create = optional
	// update = optional
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Custom environment variables for Oracle databases.
	// create = optional
	// update = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A list of object references to Oracle Node Listeners selected for
	// this Managed Database. Delphix picks one default listener from the
	// target environment if this list is empty at virtual database
	// provision time.
	// create = optional
	// update = optional
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// Number of Online Redo Log Groups.
	// minimum = 2
	// default = 3
	// create = optional
	// update = readonly
	RedoLogGroups *int `json:"redoLogGroups,omitempty"`
	// Online Redo Log size in MB.
	// create = optional
	// update = readonly
	// minimum = 4
	RedoLogSizeInMB *int `json:"redoLogSizeInMB,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// PasswordCredentialStruct - The password based security credential.
// extends Credential
type PasswordCredentialStruct struct {
	// The password.
	// format = password
	// required = true
	// minLength = 1
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PasswordPolicyStruct - Password policies for Delphix users.
// extends NamedUserObject
// cliVisibility = [DOMAIN SYSTEM]
type PasswordPolicyStruct struct {
	// True if password must contain at least one digit.
	// create = required
	// update = optional
	Digit *bool `json:"digit,omitempty"`
	// True to disallow password containing username.
	// create = required
	// update = optional
	DisallowUsernameAsPassword *bool `json:"disallowUsernameAsPassword,omitempty"`
	// True if password must contain at least one lowercase letter.
	// update = optional
	// create = required
	LowercaseLetter *bool `json:"lowercaseLetter,omitempty"`
	// Minimum length for the password.
	// minimum = 1
	// maximum = 128
	// create = required
	// update = optional
	MinLength *int `json:"minLength,omitempty"`
	// Name of password policy.
	// create = required
	// update = optional
	// minLength = 1
	// maxLength = 64
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// The password may not be the same as any of previous n passwords.
	// maximum = 1
	// create = required
	// update = optional
	// minimum = 0
	ReuseDisallowLimit *int `json:"reuseDisallowLimit,omitempty"`
	// True if password must contain at least one symbol.
	// create = required
	// update = optional
	Symbol *bool `json:"symbol,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if password must contain at least one uppercase letter.
	// create = required
	// update = optional
	UppercaseLetter *bool `json:"uppercaseLetter,omitempty"`
}

// PathDescendantConstraintStruct - Constraint placed on a filesystem path axis of a particular analytics
// slice.
// extends PathConstraint
type PathDescendantConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be a descendant of this path.
	// format = unixpath
	// create = required
	DescendantOf string `json:"descendantOf,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PemCertificateStruct - An X.509 Certificate in PEM format.
// extends TypedObject
type PemCertificateStruct struct {
	// The X.509 certificate in PEM format.
	// required = true
	Contents string `json:"contents,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PemCertificateChainStruct - A chain of X.509 Certificates in PEM format.
// extends TypedObject
type PemCertificateChainStruct struct {
	// The chain of X.509 certificates in PEM format.
	// required = true
	Chain []*PemCertificateStruct `json:"chain,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PermissionStruct - Describes a permission to perform an operation on an object in the
// Delphix Engine.
// extends ReadonlyNamedUserObject
type PermissionStruct struct {
	// Name of the action governed by the permission.
	ActionType string `json:"actionType,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLAttachDataStruct - Represents the PostgreSQL specific parameters of an attach request.
// extends AttachData
type PgSQLAttachDataStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-pgsql-db-cluster-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The database that must be used to run SQL queries against this
	// cluster.
	// maxLength = 256
	// default = postgres
	// create = optional
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
	// The credential of the database cluster user.
	// create = optional
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The username of the database cluster user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The external file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Information about the host OS user on the PPT host to use for
	// linking.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	PptHostUser string `json:"pptHostUser,omitempty"`
	// The Postgres installation on the PPT environment that will be used
	// for pre-provisioning.
	// referenceTo = /delphix-pgsql-install.json
	// required = true
	// format = objectReference
	PptRepository string `json:"pptRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLCompatibilityCriteriaStruct - The compatibility criteria to use for selecting compatible PgSQL
// repositories.
// extends CompatibilityCriteria
type PgSQLCompatibilityCriteriaStruct struct {
	// Selected repositories are installed on a host with this
	// architecture (32-bit or 64-bit).
	Architecture *int `json:"architecture,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// Selected repositories have this size WAL segments.
	// units = B
	// base = 1024
	SegmentSize *int `json:"segmentSize,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories will match this variant of the PostgreSQL
	// distribution.
	// enum = [PostgreSQL EnterpriseDB]
	Variant string `json:"variant,omitempty"`
	// Selected repositories are this database version. In case of
	// upgrade, selected repositories are strictly greater than this
	// database version.
	// format = pgsqlVersion
	Version string `json:"version,omitempty"`
}

// PgSQLCreateTransformationParametersStruct - Represents the parameters of a createTransformation request for a
// PgSQL container.
// extends CreateTransformationParameters
type PgSQLCreateTransformationParametersStruct struct {
	// The container that will contain the transformed data associated
	// with the newly created transformation; the "transformation
	// container".
	// create = required
	Container *PgSQLDatabaseContainerStruct `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// The port number used for provisioning a PgSQL container during
	// transformation application.
	// create = required
	// minimum = 1
	// maximum = 65535
	Port *int `json:"port,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLDBClusterConfigStruct - Configuration information for a PostgreSQL database cluster.
// extends SourceConfig
type PgSQLDBClusterConfigStruct struct {
	// The data directory for the PostgreSQL cluster.
	// create = optional
	ClusterDataDirectory string `json:"clusterDataDirectory,omitempty"`
	// The database that must be used to run SQL queries against this
	// cluster.
	// create = optional
	// update = optional
	// maxLength = 256
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
	// The password of the database cluster user.
	// create = optional
	// update = optional
	Credentials string `json:"credentials,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	// format = objectReference
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Whether this source should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The port on which the PostgresSQL server for the cluster is
	// listening.
	// maximum = 65535
	// create = required
	// update = optional
	// minimum = 1
	Port *int `json:"port,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-pgsql-install.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The username of the database cluster user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
}

// PgSQLDBClusterConfigConnectivityStruct - Mechanism to test JDBC connectivity to PostgreSQL DB clusters.
// extends SourceConfigConnectivity
type PgSQLDBClusterConfigConnectivityStruct struct {
	// The database that must be used to run SQL queries against this
	// cluster.
	// create = optional
	// maxLength = 256
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
	// Database password.
	// format = password
	// required = true
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Database username.
	// required = true
	Username string `json:"username,omitempty"`
}

// PgSQLDBConfigStruct - Configuration information for a PostgreSQL database in a cluster.
// extends TypedObject
type PgSQLDBConfigStruct struct {
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// The PostgreSQL cluster this database is part of.
	// format = objectReference
	// referenceTo = /delphix-pgsql-db-cluster-config.json
	DatabaseCluster string `json:"databaseCluster,omitempty"`
	// The name of the database.
	// create = optional
	// update = optional
	// maxLength = 63
	DatabaseName string `json:"databaseName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
}

// PgSQLDBContainerRuntimeStruct - Runtime properties of a PostgreSQL database container.
// extends DBContainerRuntime
type PgSQLDBContainerRuntimeStruct struct {
	// The ID of the WAL segment that was last restored in this container
	// (if applicable).
	LastRestoredWALSegment string `json:"lastRestoredWALSegment,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntimeStruct `json:"preProvisioningStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLDatabaseContainerStruct - A PostgreSQL Database Container.
// extends DatabaseContainer
type PgSQLDatabaseContainerStruct struct {
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// referenceTo = /delphix-timeflow.json
	// format = objectReference
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Optional user-provided description for the container.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	// update = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this container.
	Runtime *PgSQLDBContainerRuntimeStruct `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// PgSQLExportParametersStruct - The parameters to use as input to export PostgreSQL databases.
// extends DbExportParameters
type PgSQLExportParametersStruct struct {
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayoutStruct `json:"filesystemLayout,omitempty"`
	// Entries in the PostgreSQL host-based authentication file
	// (pg_hba.conf).
	// create = optional
	HbaEntries []*PgSQLHBAEntryStruct `json:"hbaEntries,omitempty"`
	// Entries in the PostgreSQL username map file (pg_ident.conf).
	// create = optional
	IdentEntries []*PgSQLIdentEntryStruct `json:"identEntries,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// The source config to use when creating the exported database.
	// required = true
	SourceConfig *PgSQLDBClusterConfigStruct `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export
	// on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLHBAEntryStruct - An entry in the PostgreSQL host-based authentication file
// (pg_hba.conf).
// extends TypedObject
type PgSQLHBAEntryStruct struct {
	// The client machine address that this entry matches.
	// create = optional
	Address string `json:"address,omitempty"`
	// The authentication method to use when connecting via this entry.
	// enum = [trust reject md5 password gss sspi krb5 ident peer ldap radius cert pam]
	// create = required
	AuthMethod string `json:"authMethod,omitempty"`
	// Options for the authentication method.
	// create = optional
	AuthOptions string `json:"authOptions,omitempty"`
	// The database name this entry matches.
	// maxLength = 63
	// default = all
	// create = required
	Database string `json:"database,omitempty"`
	// The connection type of this entry.
	// enum = [local host hostssl hostnossl]
	// create = required
	EntryType string `json:"entryType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database username this entry matches.
	// create = required
	// maxLength = 63
	// default = all
	User string `json:"user,omitempty"`
}

// PgSQLIdentEntryStruct - An entry in the PostgreSQL username map file (pg_ident.conf).
// extends TypedObject
type PgSQLIdentEntryStruct struct {
	// The database username this entry matches.
	// maxLength = 63
	// create = required
	DatabaseUsername string `json:"databaseUsername,omitempty"`
	// The name of the map to which this entry belongs (used to refer to
	// the map in pg_hba.conf).
	// create = required
	MapName string `json:"mapName,omitempty"`
	// The operating system username this entry matches.
	// create = required
	SystemUsername string `json:"systemUsername,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLInstallStruct - A PostgreSQL installation.
// extends SourceRepository
type PgSQLInstallStruct struct {
	// 32 or 64 bit installation.
	// enum = [32 64]
	Bits *int `json:"bits,omitempty"`
	// Flag indicating whether the installation was discovered or
	// manually entered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Directory path where the installation is located.
	// create = required
	// maxLength = 1024
	InstallationPath string `json:"installationPath,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// create = optional
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Size of the WAL segments (in bytes) generated by PostgreSQL
	// binaries.
	// enum = [1.048576e+06 2.097152e+06 4.194304e+06 8.388608e+06 1.6777216e+07 3.3554432e+07 6.7108864e+07]
	// base = 1024
	// units = B
	SegmentSize *int `json:"segmentSize,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// create = optional
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Variant of the repository.
	// enum = [PostgreSQL EnterpriseDB]
	Variant string `json:"variant,omitempty"`
	// Version of the repository.
	// format = pgsqlVersion
	Version string `json:"version,omitempty"`
}

// PgSQLLinkDataStruct - PostgreSQL specific parameters for a link request.
// extends LinkData
type PgSQLLinkDataStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-pgsql-db-cluster-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The database that must be used to run SQL queries against this
	// cluster.
	// maxLength = 256
	// default = postgres
	// create = optional
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
	// The credential of the database cluster user.
	// create = optional
	DbCredentials *PasswordCredentialStruct `json:"dbCredentials,omitempty"`
	// The username of the database cluster user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The external file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// Information about the host OS user on the PPT host to use for
	// linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	PptHostUser string `json:"pptHostUser,omitempty"`
	// The Postgres installation on the PPT environment that will be used
	// for pre-provisioning.
	// format = objectReference
	// referenceTo = /delphix-pgsql-install.json
	// required = true
	PptRepository string `json:"pptRepository,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicyStruct `json:"sourcingPolicy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLLinkedSourceStruct - A linked PostgreSQL source.
// extends PgSQLSource
type PgSQLLinkedSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// The external file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *PgSQLSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// The staging source for pre-provisioning of the database.
	// format = objectReference
	// referenceTo = /delphix-pgsql-staging-source.json
	StagingSource string `json:"stagingSource,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// PgSQLPlatformParametersStruct - PgSQL platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type PgSQLPlatformParametersStruct struct {
	// The port number used for provisioning a PgSQL container during
	// transformation application. This port must be available when the
	// transformation is applied so that the temporary VDB created during
	// the transformation process can start and listen to this port.
	// maximum = 65535
	// minimum = 1
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLProvisionParametersStruct - The parameters to use as input to provision PostgreSQL databases.
// extends ProvisionParameters
type PgSQLProvisionParametersStruct struct {
	// The new container for the provisioned database.
	// required = true
	Container *PgSQLDatabaseContainerStruct `json:"container,omitempty"`
	// Whether or not to mark this VDB as a masked VDB. It will be marked
	// as masked if this flag or the masking job are set.
	// create = optional
	// update = readonly
	Masked *bool `json:"masked,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *PgSQLVirtualSourceStruct `json:"source,omitempty"`
	// The source config for the source.
	// required = true
	SourceConfig *PgSQLDBClusterConfigStruct `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLSnapshotStruct - Provisionable snapshot of a PostgreSQL TimeFlow.
// extends TimeflowSnapshot
type PgSQLSnapshotStruct struct {
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *PgSQLTimeflowPointStruct `json:"firstChangePoint,omitempty"`
	// The location of the snapshot within the parent TimeFlow
	// represented by this snapshot.
	LatestChangePoint *PgSQLTimeflowPointStruct `json:"latestChangePoint,omitempty"`
	// Boolean value indicating if a virtual database provisioned from
	// this snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot
	// should be kept forever.
	// update = optional
	Retention *int `json:"retention,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *SnapshotRuntimeStruct `json:"runtime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient
	// state and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Time zone of the source database at the time the snapshot was
	// taken.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
}

// PgSQLSourceConnectionInfoStruct - Contains information that can be used to connect to a PostgreSQL
// source.
// extends SourceConnectionInfo
type PgSQLSourceConnectionInfoStruct struct {
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The JDBC string used to connect to the PostgreSQL server instance.
	JdbcString string `json:"jdbcString,omitempty"`
	// The data directory for the PostgreSQL cluster.
	PgDataDirectory string `json:"pgDataDirectory,omitempty"`
	// The port on which the PostgresSQL server for the cluster is
	// listening.
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database cluster user.
	User string `json:"user,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// PgSQLSourceRuntimeStruct - Runtime (non-persistent) properties of a PostgreSQL source.
// extends SourceRuntime
type PgSQLSourceRuntimeStruct struct {
	// True if the source is JDBC accessible. If false then no properties
	// can be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLStagingSourceStruct - A PostgreSQL staging source used for pre-provisioning.
// extends PgSQLSource
type PgSQLStagingSourceStruct struct {
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// create = optional
	// update = optional
	// default = false
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point for the NFS mounts on the pre-provisioning
	// host.
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *PgSQLSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// PgSQLSyncParametersStruct - The parameters to use as input to sync PostgreSQL databases.
// extends SyncParameters
type PgSQLSyncParametersStruct struct {
	// Whether or not to take another full backup of the source database.
	// default = false
	RedoBaseBackup *bool `json:"redoBaseBackup,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLTimeflowStruct - TimeFlow representing historical data for a particular timeline within
// a PostgreSQL container.
// extends Timeflow
type PgSQLTimeflowStruct struct {
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC SOURCE_CONTINUITY]
	CreationType string `json:"creationType,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow
	// was provisioned. This will not be present for TimeFlows derived
	// from linked sources.
	ParentPoint *PgSQLTimeflowPointStruct `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning
	// base for this object. This may be different from the snapshot
	// within the parent point, and is only present for virtual
	// TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLTimeflowPointStruct - A unique point within a PostgreSQL database TimeFlow.
// extends TimeflowPoint
type PgSQLTimeflowPointStruct struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this point.
	// referenceTo = /delphix-pgsql-timeflow.json
	// required = true
	// format = objectReference
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLVirtualSourceStruct - A virtual PostgreSQL source.
// extends PgSQLSource
type PgSQLVirtualSourceStruct struct {
	// Indicates whether Delphix should automatically restart this
	// virtual source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// PostgreSQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Entries in the PostgreSQL host-based authentication file
	// (pg_hba.conf).
	// create = optional
	HbaEntries []*PgSQLHBAEntryStruct `json:"hbaEntries,omitempty"`
	// Hosts that might affect operations on this source. Property will
	// be null unless the includeHosts parameter is set when listing
	// sources.
	Hosts []string `json:"hosts,omitempty"`
	// Entries in the PostgreSQL username map file (pg_ident.conf).
	// create = optional
	IdentEntries []*PgSQLIdentEntryStruct `json:"identEntries,omitempty"`
	// Flag indicating whether the source is a linked source in the
	// Delphix system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source.
	// update = optional
	// default = false
	// create = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 256
	// create = required
	MountBase string `json:"mountBase,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperationsStruct `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *PgSQLSourceRuntimeStruct `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix
	// system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is a virtual source in the
	// Delphix system.
	Virtual *bool `json:"virtual,omitempty"`
}

// PhoneHomeServiceStruct - Phone home service configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type PhoneHomeServiceStruct struct {
	// True if the phone home service is enabled.
	// required = true
	Enabled *bool `json:"enabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PolicyApplyTargetParametersStruct - Specifies the target to which a policy is applied.
// extends TypedObject
type PolicyApplyTargetParametersStruct struct {
	// Object reference of the target.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// Object references of the targets.
	Targets []string `json:"targets,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PolicyCreateAndApplyParametersStruct - The parameters used for creating and applying policies.
// extends TypedObject
type PolicyCreateAndApplyParametersStruct struct {
	// Policy to create.
	// required = true
	Policy Policy `json:"policy,omitempty"`
	// Object reference of the target.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	// required = true
	Target string `json:"target,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PopulateCompatibilityParametersStruct - The criteria necessary to select valid repositories to populate into a
// warehouse.
// extends CompatibleRepositoriesParameters
type PopulateCompatibilityParametersStruct struct {
	// Restrict returned repositories to this environment.
	// referenceTo = /delphix-source-environment.json
	// create = optional
	// update = optional
	// format = objectReference
	Environment string `json:"environment,omitempty"`
	// The warehouse repository to use as a source of compatibility
	// information.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// required = true
	SourceRepository string `json:"sourceRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PreProvisioningRuntimeStruct - Runtime properties for pre-provisioning of a MSSQL database container.
// extends TypedObject
type PreProvisioningRuntimeStruct struct {
	// Timestamp of the last update to the status.
	LastUpdateTimestamp string `json:"lastUpdateTimestamp,omitempty"`
	// User action required to resolve any error that the
	// pre-provisioning run encountered.
	PendingAction string `json:"pendingAction,omitempty"`
	// Indicates the current state of pre-provisioning for the database.
	// enum = [ACTIVE INACTIVE FAULTED UNKNOWN]
	PreProvisioningState string `json:"preProvisioningState,omitempty"`
	// Response taken based on the status of the pre-provisioning run.
	Response string `json:"response,omitempty"`
	// The status of the pre-provisioning run.
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ProvisionCompatibilityParametersStruct - The criteria necessary to select valid repositories for provisioning.
// extends CompatibleRepositoriesParameters
type ProvisionCompatibilityParametersStruct struct {
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The TimeFlow point to use as a source of compatibility
	// information.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ProxyConfigurationStruct - Proxy configuration for a specific protocol.
// extends TypedObject
type ProxyConfigurationStruct struct {
	// True if the proxy is enabled.
	// update = optional
	Enabled *bool `json:"enabled,omitempty"`
	// Host or IP address to use as proxy.
	// format = host
	// update = optional
	Host string `json:"host,omitempty"`
	// If authentication is required, the password to use when connecting
	// to the proxy.
	// format = password
	// update = optional
	Password string `json:"password,omitempty"`
	// Port to use when connecting to the proxy host.
	// minimum = 1
	// maximum = 65535
	// update = optional
	Port *int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If authentication is required, the username to use when connecting
	// to the proxy.
	// update = optional
	Username string `json:"username,omitempty"`
}

// ProxyServiceStruct - Proxy service configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type ProxyServiceStruct struct {
	// HTTPS proxy configuration.
	// required = false
	// update = optional
	Https *ProxyConfigurationStruct `json:"https,omitempty"`
	// SOCKS proxy configuration.
	// update = optional
	// required = false
	Socks *ProxyConfigurationStruct `json:"socks,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PublicSystemInfoStruct - Retrieve static system-wide properties.
// extends TypedObject
// cliVisibility = [DOMAIN SYSTEM]
type PublicSystemInfoStruct struct {
	// Maximum supported API version of the current system software.
	ApiVersion *APIVersionStruct `json:"apiVersion,omitempty"`
	// Security banner to display prior to login.
	Banner string `json:"banner,omitempty"`
	// Time at which the current system software was built.
	// format = date
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
	// Description of the current system software.
	BuildTitle string `json:"buildTitle,omitempty"`
	// Delphix version of the current system software.
	BuildVersion *VersionInfoStruct `json:"buildVersion,omitempty"`
	// Indicates whether the server has gone through initial setup or
	// not.
	Configured *bool `json:"configured,omitempty"`
	// The current system locale.
	// format = locale
	CurrentLocale string `json:"currentLocale,omitempty"`
	// The list of enabled features on this Delphix Engine.
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`
	// Qualifier for referencing instances of (e.g. 'Delphix') engines in
	// messages like 'The <engineQualifier> Engine failed to ...'.
	// create = required
	EngineQualifier string `json:"engineQualifier,omitempty"`
	// The operating system kernel name.
	KernelName string `json:"kernelName,omitempty"`
	// List of available locales.
	Locales []string `json:"locales,omitempty"`
	// Name of the product that the system is running.
	ProductName string `json:"productName,omitempty"`
	// Product type.
	ProductType string `json:"productType,omitempty"`
	// Technical support phone numbers.
	// create = optional
	SupportContacts []*SupportContactStruct `json:"supportContacts,omitempty"`
	// Technical Support URL.
	// create = optional
	SupportURL string `json:"supportURL,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Address of vendor headquarters. Free form collection of strings to
	// accomodate any region.
	// create = optional
	VendorAddress []string `json:"vendorAddress,omitempty"`
	// Corporate headquarters email address.
	// create = optional
	// format = email
	VendorEmail string `json:"vendorEmail,omitempty"`
	// Vendor name, for use in messages like 'Please contact <vendorName>
	// customer support'.
	// create = required
	VendorName string `json:"vendorName,omitempty"`
	// Corporate headquarters telephone number.
	// create = optional
	VendorPhoneNumber string `json:"vendorPhoneNumber,omitempty"`
	// Corporate home page.
	// create = optional
	VendorURL string `json:"vendorURL,omitempty"`
}

// PurgeLogsParametersStruct - Represents the parameters of a purgeLogs request.
// extends TypedObject
type PurgeLogsParametersStruct struct {
	// Delete expired logs which have been retained to make snapshots
	// consistent.
	// default = false
	// required = true
	DeleteSnapshotLogs *bool `json:"deleteSnapshotLogs,omitempty"`
	// If this is set to true, this operation does not actually delete
	// logs. It returns the affected snapshots and truncated timeline as
	// if the logs were deleted.
	// default = true
	// required = true
	DryRun *bool `json:"dryRun,omitempty"`
	// Amount of space in bytes to reclaim as part of purgeLogs process.
	// required = true
	// minimum = 1
	// units = B
	// base = 1024
	StorageSpaceToReclaim float64 `json:"storageSpaceToReclaim,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PurgeLogsResultStruct - Represents the result of a purgeLogs operation.
// extends TypedObject
type PurgeLogsResultStruct struct {
	// List of snapshots which have been rendered unprovisionable because
	// logs needed to make them consistent have been deleted.
	AffectedSnapshots []TimeflowSnapshot `json:"affectedSnapshots,omitempty"`
	// TimeFlow point after the last snapshot beyond which TimeFlow will
	// be lost as a result of purging logs.
	TruncatePoint *OracleTimeflowPointStruct `json:"truncatePoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// QuotaPolicyStruct - This policy limits the maximum amount of space an object (group or
// database) can use.
// extends Policy
type QuotaPolicyStruct struct {
	// Last time a critical alert was generated.
	// format = date
	CritAlertTime string `json:"critAlertTime,omitempty"`
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// True if this is the default policy created when the system is
	// setup. Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Whether this policy has been directly applied or inherited. See
	// the effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Size of the quota, in bytes.
	// update = optional
	// minimum = 1
	// units = B
	// base = 1024
	// create = required
	Size float64 `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Last time a warning alert was generated.
	// format = date
	WarnAlertTime string `json:"warnAlertTime,omitempty"`
}

// RefreshParametersStruct - The parameters to use as input to refresh requests for MSSQL,
// PostgreSQL, AppData, ASE or MySQL.
// extends TypedObject
type RefreshParametersStruct struct {
	// The TimeFlow point, bookmark, or semantic location to refresh the
	// database to.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RefreshPolicyStruct - This policy refreshes a container according to a schedule.
// extends SchedulePolicy
type RefreshPolicyStruct struct {
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// True if this is the default policy created when the system is
	// setup. Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Whether this policy has been directly applied or inherited. See
	// the effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Provision source, either the latest time or latest snapshot.
	// enum = [LATEST_SNAPSHOT LATEST_TIME_FLOW_LOG]
	// create = required
	// update = optional
	ProvisionSource string `json:"provisionSource,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// List of Schedule objects representing when the policy will
	// execute.
	// create = required
	// update = optional
	ScheduleList []*ScheduleStruct `json:"scheduleList,omitempty"`
	// The timezone of this policy. If not specified, defaults to the
	// Delphix Engine's timezone.
	// create = optional
	// update = optional
	Timezone *TimeZoneStruct `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RegistrationInfoStruct - The information required to register the Delphix Engine.
// extends TypedObject
// cliVisibility = [SYSTEM]
type RegistrationInfoStruct struct {
	// The registration code for this Delphix Engine.
	Code string `json:"code,omitempty"`
	// The registration portal hostname.
	RegistrationPortalHostname string `json:"registrationPortalHostname,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The UUID for this Delphix Engine.
	Uuid string `json:"uuid,omitempty"`
}

// RegistrationParametersStruct - Credentials used to register the Delphix Engine.
// extends TypedObject
type RegistrationParametersStruct struct {
	// Password to send to registration portal.
	// format = password
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Username to send to registration portal.
	Username string `json:"username,omitempty"`
}

// RegistrationStatusStruct - Information on the status of the Delphix Engine's registration.
// extends UserObject
// cliVisibility = [SYSTEM]
type RegistrationStatusStruct struct {
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The status of the Delphix Engine's registration. It may be
	// unknown, unregistered, in-progress, or registered.
	// update = optional
	Status string `json:"status,omitempty"`
	// The time at which the registration status was last updated.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RemoteDelphixEngineInfoStruct - Parameters for logging into another Delphix Engine when running a
// network throughput test.
// extends TypedObject
type RemoteDelphixEngineInfoStruct struct {
	// Address of other Delphix Engine.
	// format = host
	// create = required
	Address string `json:"address,omitempty"`
	// Password for the other Delphix Engine.
	// create = required
	Credential *PasswordCredentialStruct `json:"credential,omitempty"`
	// Username for the other Delphix Engine.
	// create = required
	Principal string `json:"principal,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ReplicationListStruct - List of objects that are to be replicated.
// extends ReplicationObjectSpecification
type ReplicationListStruct struct {
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Objects to replicate, in canonical object reference form.
	// minItems = 1
	// create = required
	// update = optional
	Objects []string `json:"objects,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ReplicationSecureListStruct - List of containers that are to be securely replicated.
// extends ReplicationObjectSpecification
type ReplicationSecureListStruct struct {
	// Containers to replicate, in canonical object reference form.
	// minItems = 1
	// create = required
	// update = optional
	Containers []string `json:"containers,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ReplicationSourceStateStruct - State of a replication spec.
// extends UserObject
type ReplicationSourceStateStruct struct {
	// The active serialization point, currently being sent or about to
	// be sent to replication targets.
	// format = objectReference
	// referenceTo = /delphix-serialization-point.json
	ActivePoint string `json:"activePoint,omitempty"`
	// The last serialization point sent. This can be null prior to the
	// first replication run.
	// referenceTo = /delphix-serialization-point.json
	// format = objectReference
	LastPoint string `json:"lastPoint,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// A reference to the replication specification responsible for the
	// current state.
	// format = objectReference
	// referenceTo = /delphix-replicationspec.json
	Spec string `json:"spec,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ReplicationSpecStruct - Replication setup.
// extends UserObject
type ReplicationSpecStruct struct {
	// Indication whether the replication spec schedule is enabled or
	// not.
	// create = optional
	// update = optional
	// default = false
	AutomaticReplication *bool `json:"automaticReplication,omitempty"`
	// Bandwidth limit (MB/s) for replication network traffic. A value of
	// 0 means no limit.
	// create = optional
	// update = optional
	// default = 0
	// minimum = 0
	BandwidthLimit *int `json:"bandwidthLimit,omitempty"`
	// Description of this replication spec.
	// update = optional
	// maxLength = 4096
	// create = optional
	Description string `json:"description,omitempty"`
	// Encrypt replication network traffic.
	// create = optional
	// update = optional
	// default = false
	Encrypted *bool `json:"encrypted,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Total number of transport connections to use.
	// maximum = 16
	// create = optional
	// update = optional
	// default = 1
	// minimum = 1
	NumberOfConnections *int `json:"numberOfConnections,omitempty"`
	// Specification of the objects to replicate.
	// create = required
	// update = optional
	ObjectSpecification ReplicationObjectSpecification `json:"objectSpecification,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this replication spec.
	Runtime *ReplicationSpecRuntimeStruct `json:"runtime,omitempty"`
	// Replication schedule in the form of a quartz-formatted string.
	// create = optional
	// update = optional
	// minLength = 1
	// maxLength = 256
	Schedule string `json:"schedule,omitempty"`
	// Globally unique identifier for this replication spec.
	// minLength = 1
	// maxLength = 256
	Tag string `json:"tag,omitempty"`
	// Credential used to authenticate to the replication target host.
	// create = required
	// update = optional
	TargetCredential *PasswordCredentialStruct `json:"targetCredential,omitempty"`
	// Replication target host address.
	// create = required
	// update = optional
	// format = host
	TargetHost string `json:"targetHost,omitempty"`
	// Target TCP port number for the Delphix Session Protocol.
	// create = optional
	// update = optional
	// minimum = 0
	// maximum = 65535
	// default = 8415
	TargetPort *int `json:"targetPort,omitempty"`
	// Principal name used to authenticate to the replication target
	// host.
	// create = required
	// update = optional
	TargetPrincipal string `json:"targetPrincipal,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Connect to the replication target host via the system-wide SOCKS
	// proxy.
	// create = optional
	// update = optional
	// default = false
	UseSystemSocksSetting *bool `json:"useSystemSocksSetting,omitempty"`
}

// ReplicationSpecRuntimeStruct - Runtime properties for a replication spec.
// extends TypedObject
type ReplicationSpecRuntimeStruct struct {
	// Date of the next execution of this replication spec according to
	// the schedule.
	// format = date
	NextScheduledExecution string `json:"nextScheduledExecution,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ReplicationTargetStateStruct - State of a replication at the target.
// extends UserObject
type ReplicationTargetStateStruct struct {
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ResetIgnoredFaultsParametersStruct - The parameters to use as input when marking selected ignored faults as
// resolved.
// extends TypedObject
type ResetIgnoredFaultsParametersStruct struct {
	// Input list of ignored faults which will be marked as resolved.
	// required = true
	FaultReferences []string `json:"faultReferences,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ResolveOrIgnoreSelectedFaultsParametersStruct - The parameters to use as input when marking selected faults as
// resolved or ignored.
// extends TypedObject
type ResolveOrIgnoreSelectedFaultsParametersStruct struct {
	// Input list of faults which will be marked as resolved or ignored.
	// required = true
	FaultReferences []string `json:"faultReferences,omitempty"`
	// Flag indicating whether to ignore the selected faults if they are
	// detected on the same objects in the future.
	// default = false
	// required = true
	Ignore *bool `json:"ignore,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RetentionPolicyStruct - This policy controls what data (log and snapshot) is kept.
// extends Policy
type RetentionPolicyStruct struct {
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// Amount of time (in dataUnit units) to keep source data.
	// create = optional
	// update = optional
	DataDuration *int `json:"dataDuration,omitempty"`
	// Time unit for dataDuration.
	// enum = [DAY WEEK MONTH QUARTER YEAR]
	// create = optional
	// update = optional
	DataUnit string `json:"dataUnit,omitempty"`
	// Day of month upon which to enforce monthly snapshot retention.
	// create = optional
	// update = optional
	DayOfMonth *int `json:"dayOfMonth,omitempty"`
	// Day of week upon which to enforce weekly snapshot retention.
	// create = optional
	// update = optional
	// enum = [MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY]
	DayOfWeek string `json:"dayOfWeek,omitempty"`
	// Day of year upon which to enforce yearly snapshot retention,
	// expressed a month / day string (e.g., "Jan 1").
	// create = optional
	// update = optional
	// maxLength = 32
	DayOfYear string `json:"dayOfYear,omitempty"`
	// True if this is the default policy created when the system is
	// setup. Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Whether this policy has been directly applied or inherited. See
	// the effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// Amount of time (in logUnit units) to keep log data.
	// update = optional
	// create = optional
	LogDuration *int `json:"logDuration,omitempty"`
	// Time unit for logDuration.
	// enum = [DAY WEEK MONTH QUARTER YEAR]
	// create = optional
	// update = optional
	LogUnit string `json:"logUnit,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Number of daily snapshots to keep.
	// create = optional
	// update = optional
	NumOfDaily *int `json:"numOfDaily,omitempty"`
	// Number of monthly snapshots to keep.
	// create = optional
	// update = optional
	NumOfMonthly *int `json:"numOfMonthly,omitempty"`
	// Number of weekly snapshots to keep.
	// create = optional
	// update = optional
	NumOfWeekly *int `json:"numOfWeekly,omitempty"`
	// Number of yearly snapshots to keep.
	// create = optional
	// update = optional
	NumOfYearly *int `json:"numOfYearly,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RoleStruct - Describes a role as applied to a user on an object.
// extends UserObject
type RoleStruct struct {
	// Determines if the role can be modified or not. Some roles are
	// shipped with the Delphix Engine and cannot be changed.
	Immutable *bool `json:"immutable,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// List of permissions contained in the role.
	// create = required
	// update = optional
	Permissions []string `json:"permissions,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// RollbackParametersStruct - The parameters to use as input to rollback requests for MSSQL,
// PostgreSQL, AppData, ASE or MySQL.
// extends TypedObject
type RollbackParametersStruct struct {
	// The TimeFlow point, bookmark, or semantic location to roll the
	// database back to.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RsaKeyPairStruct - A key pair generated using the RSA algorithm.
// extends KeyPair
type RsaKeyPairStruct struct {
	// The size of each key to be generated.
	// create = optional
	// minimum = 2048
	// maximum = 4096
	// default = 2048
	KeySize *int `json:"keySize,omitempty"`
	// The signature algorithm this key pair will use to sign
	// certificates and CSRs.
	// enum = [SHA256withRSA SHA384withRSA SHA512withRSA]
	// default = SHA256withRSA
	// create = optional
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RunBashOnSourceOperationStruct - A user-specifiable operation that runs a shell command on the target
// host using the Delphix supplied Bash shell.
// extends SourceOperation
type RunBashOnSourceOperationStruct struct {
	// The shell command to execute on the target host.
	// create = required
	// update = required
	Command string `json:"command,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RunCommandOnSourceOperationStruct - A user-specifiable operation that runs a shell command on the target
// host.
// extends SourceOperation
type RunCommandOnSourceOperationStruct struct {
	// The shell command to execute on the target host.
	// create = required
	// update = optional
	Command string `json:"command,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RunExpectOnSourceOperationStruct - A user-specifiable operation that runs an expect script on the target
// host.
// extends SourceOperation
type RunExpectOnSourceOperationStruct struct {
	// The expect command to execute on the target host.
	// create = required
	// update = optional
	Command string `json:"command,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RunMaskingJobOperationStruct - An operation that runs a Masking Job on the local Delphix Masking
// Engine instance.
// extends Operation
type RunMaskingJobOperationStruct struct {
	// The location this Masking Job will be executed on.
	// create = readonly
	// update = readonly
	Host string `json:"host,omitempty"`
	// The Masking Job ID of the Masking Job to be executed.
	// create = readonly
	// update = readonly
	MaskingJobId string `json:"maskingJobId,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// RunPowerShellOnSourceOperationStruct - A user-specifiable operation that runs a PowerShell command on the
// target host.
// extends SourceOperation
type RunPowerShellOnSourceOperationStruct struct {
	// The PowerShell command to execute on the target host.
	// create = required
	// update = optional
	Command string `json:"command,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SMTPConfigStruct - SMTP configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SMTPConfigStruct struct {
	// True if username/password authentication should be used.
	// default = false
	// update = optional
	AuthenticationEnabled *bool `json:"authenticationEnabled,omitempty"`
	// True if outbound email is enabled.
	// update = optional
	Enabled *bool `json:"enabled,omitempty"`
	// From address to use when sending mail. If unspecified,
	// 'noreply@delphix.com' is used.
	// format = email
	// update = optional
	FromAddress string `json:"fromAddress,omitempty"`
	// If authentication is enabled, password to use when authenticating
	// to the server.
	// format = password
	// update = optional
	Password string `json:"password,omitempty"`
	// Port number to use. A value of -1 indicates the default (25 or 587
	// for TLS).
	// default = -1
	// update = optional
	// minimum = -1
	// maximum = 65535
	Port *int `json:"port,omitempty"`
	// Maximum timeout to wait, in seconds, when sending mail.
	// minimum = 1
	// update = optional
	SendTimeout *int `json:"sendTimeout,omitempty"`
	// IP address or hostname of SMTP relay server.
	// format = host
	// update = optional
	Server string `json:"server,omitempty"`
	// True if TLS (transport layer security) should be used.
	// default = false
	// update = optional
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If authentication is enabled, username to use when authenticating
	// to the server.
	// update = optional
	Username string `json:"username,omitempty"`
}

// SNMPConfigStruct - SNMP configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SNMPConfigStruct struct {
	// The network which is authorized to query this SNMP server, in CIDR
	// notation. Toallow any client, then leave unset or set to
	// 0.0.0.0/0. To block all clients, set to 127.0.0.1/8.
	// format = cidrAddress
	// update = optional
	// default = 0.0.0.0/0
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty"`
	// The community string that clients must provide when querying this
	// server.
	// update = optional
	// default = public
	Community string `json:"community,omitempty"`
	// True if the SNMP service is enabled.
	// update = optional
	// default = true
	Enabled *bool `json:"enabled,omitempty"`
	// The physical location of this Delphix Engine (OID 1.3.6.1.2.1.1.6
	// - sysLocation).
	// update = optional
	Location string `json:"location,omitempty"`
	// SNMP trap severity. SNMP managers are only notified of events at
	// or above this level.
	// update = optional
	// enum = [CRITICAL WARNING INFORMATIONAL]
	// default = WARNING
	Severity string `json:"severity,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SNMPManagerStruct - SNMP manager configuration.
// extends PersistentObject
// cliVisibility = [SYSTEM]
type SNMPManagerStruct struct {
	// SNMP manager host.
	// create = required
	// update = optional
	// format = host
	Address string `json:"address,omitempty"`
	// SNMP manager community string.
	// create = required
	// update = optional
	CommunityString string `json:"communityString,omitempty"`
	// Describes if the most recent attempt to send a trap succeeded or
	// failed.
	// enum = [FAILED SUCCEEDED PENDING UNCHECKED]
	// default = PENDING
	// create = readonly
	// update = readonly
	LastSendStatus string `json:"lastSendStatus,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// SNMP manager port number.
	// update = optional
	// minimum = 1
	// maximum = 65535
	// default = 162
	// create = optional
	Port *int `json:"port,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if INFORM messages are to be sent to this manager, false for
	// TRAP messages.
	// create = optional
	// update = optional
	// default = false
	UseInform *bool `json:"useInform,omitempty"`
}

// SSHConnectivityStruct - Mechanism to test SSH connectivity of arbitrary hosts.
// extends TypedObject
type SSHConnectivityStruct struct {
	// Target host name or IP address.
	// required = true
	Address string `json:"address,omitempty"`
	// User credentials.
	// required = true
	Credentials Credential `json:"credentials,omitempty"`
	// SSH port on remote server.
	// minimum = 1
	// maximum = 65535
	// default = 22
	Port *int `json:"port,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User name.
	// required = true
	Username string `json:"username,omitempty"`
}

// SamlAuthParametersStruct - The parameter to use as input to determine whether to encode a SAML
// authentication request.
// extends TypedObject
type SamlAuthParametersStruct struct {
	// Set to true to encode SAML authentication requests.
	// required = true
	// default = true
	EncodeRequest *bool `json:"encodeRequest,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SamlInfoStruct - Global SAML information.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SamlInfoStruct struct {
	// Whether or not SAML authentication is configured and enabled for
	// this Delphix Engine.
	Enabled *bool `json:"enabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SamlServiceProviderStruct - SAML service provider configuration.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type SamlServiceProviderStruct struct {
	// The public URL the Delphix Engine will be accessed at.
	// required = true
	BaseUrl string `json:"baseUrl,omitempty"`
	// The decryption (private) key that will be used to decrypt SAML
	// assertions. Leave empty if response will not be encrypted. This
	// key MUST be a PKCS8 key.
	// create = optional
	// update = optional
	DecryptingKey string `json:"decryptingKey,omitempty"`
	// URL to which the SAML authentication request will be sent.
	// required = true
	Destination string `json:"destination,omitempty"`
	// Unique name of service provider.
	// required = true
	EntityId string `json:"entityId,omitempty"`
	// Algorithm used for creation of digital signature on metadata.
	// required = true
	HashAlgUrl string `json:"hashAlgUrl,omitempty"`
	// A unique identifier that is provided by the identity provider to
	// ensure we are a valid service provider.
	// required = true
	IssuerId string `json:"issuerId,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The signing (public) key that will be used to verify SAML
	// signatures. Leave empty if responses will not be signed.
	// create = optional
	// update = optional
	SigningKey string `json:"signingKey,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ScheduleStruct - Represent a schedule in the Delphix system.
// extends TypedObject
type ScheduleStruct struct {
	// Schedule cron string. See CronTrigger documentation at
	// http://quartz-scheduler.org/ for details.
	// update = required
	// maxLength = 120
	// create = required
	CronString string `json:"cronString,omitempty"`
	// Cutoff time in seconds. The policy job will suspend if not
	// completed within the given time limit.
	// units = sec
	// create = optional
	// update = optional
	CutoffTime *int `json:"cutoffTime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SchemaStruct - Schema object.
// extends TypedObject
type SchemaStruct struct {
	// JSON representation of the schema based on the locale and API
	// Session version.
	Schema string `json:"schema,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SchemaDraftV4Struct - A dummy schema that is used to represent JSON that is a valid Draft v4
// schema.
type SchemaDraftV4Struct struct {
}

// ScrubStatusStruct - The status of a scrub of the storage in the system.
// extends TypedObject
// cliVisibility = [SYSTEM]
type ScrubStatusStruct struct {
	// Amount of data scrubbed, in bytes.
	// base = 1024
	// units = B
	Completed float64 `json:"completed,omitempty"`
	// Time scrub ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Number of errors encountered during scrub.
	Errors float64 `json:"errors,omitempty"`
	// Time scrub was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Scrub state.
	// enum = [NONE ACTIVE COMPLETED CANCELED]
	State string `json:"state,omitempty"`
	// Total amount of data to scrub (including completed), in bytes.
	// units = B
	// base = 1024
	Total float64 `json:"total,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SecurityConfigStruct - System wide security configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SecurityConfigStruct struct {
	// Allowed origin domains for CORS. Should be a comma separated list.
	// Use * for all domains. Defaults to none. Changing this value
	// requires a stack restart for it to take effect.
	// update = optional
	AllowedCORSOrigins string `json:"allowedCORSOrigins,omitempty"`
	// Banner displayed prior to login.
	// update = optional
	Banner string `json:"banner,omitempty"`
	// Whether or not CORS is enabled. Changing this value requires a
	// stack restart for it to take effect.
	// update = optional
	IsCORSEnabled *bool `json:"isCORSEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SerializationPointStruct - A serialization point represents the data and metadata associated with
// a replication spec at a point in time.
// extends UserObject
type SerializationPointStruct struct {
	// Average throughput of the transfer of the serialization point
	// (bytes/s).
	// base = 1024
	// units = B/s
	AverageThroughput float64 `json:"averageThroughput,omitempty"`
	// Bytes of the serialization point which have been transferred.
	// base = 1024
	// units = B
	BytesTransferred *int `json:"bytesTransferred,omitempty"`
	// Timestamp of the data being stored in the serialization point.
	// format = date
	DataTimestamp string `json:"dataTimestamp,omitempty"`
	// The elapsed time spent sending the serialization point
	// (nanoseconds).
	// base = 1024
	// units = nsec
	ElapsedTimeNanos float64 `json:"elapsedTimeNanos,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// An arbitrary string used to map the serialization point to a
	// corresponding replication spec or namespace.
	Tag string `json:"tag,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SeverityFilterStruct - An alert filter that specifies which severity levels to match against.
// extends AlertFilter
type SeverityFilterStruct struct {
	// List of severity levels. Only alerts matching one of the given
	// severity levels are included.
	// create = optional
	// update = optional
	// uniqueItems = true
	// minItems = 1
	SeverityLevels []string `json:"severityLevels,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SingletonUpdateStruct - An event indicating an update to a singleton object on the system.
// extends Notification
type SingletonUpdateStruct struct {
	// Type of target object.
	// format = type
	ObjectType string `json:"objectType,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// SnapshotCapacityDataStruct - Capacity metrics for a single snapshot.
// extends TypedObject
// cliVisibility = [DOMAIN]
type SnapshotCapacityDataStruct struct {
	// Reference to the container to which this snapshot belongs.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// List of VDBs that have been provisioned from this snapshot.
	DescendantVDBs []string `json:"descendantVDBs,omitempty"`
	// The manual retention setting on this snapshot, in days.
	ManualRetention *int `json:"manualRetention,omitempty"`
	// Whether this snapshot is currently being retained due to policy
	// settings.
	PolicyRetention *bool `json:"policyRetention,omitempty"`
	// Reference to the snapshot.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	Snapshot string `json:"snapshot,omitempty"`
	// Time at which this snapshot was taken.
	// format = date
	SnapshotTimestamp string `json:"snapshotTimestamp,omitempty"`
	// Space used by the snapshot.
	// units = B
	// base = 1024
	Space float64 `json:"space,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SnapshotLogFetchParametersStruct - Parameters to fetch log files that cover a snapshot and the TimeFlow
// range up to the next snapshot.
// extends LogFetchSSH
type SnapshotLogFetchParametersStruct struct {
	// User credentials. If not provided will use environment credentials
	// for 'username' on 'host'.
	Credentials Credential `json:"credentials,omitempty"`
	// Directory on the remote server where the missing log files reside.
	// required = true
	Directory string `json:"directory,omitempty"`
	// Remote host to connect to.
	// required = true
	Host string `json:"host,omitempty"`
	// SSH port to connect to.
	// default = 22
	Port *int `json:"port,omitempty"`
	// Reference to the snapshot for which to fetch logs.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-oracle-snapshot.json
	Snapshot string `json:"snapshot,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User name to authenticate as.
	// required = true
	Username string `json:"username,omitempty"`
}

// SnapshotPolicyStruct - This policy creates snapshots of a container with externally managed
// sources (virtual databases) according to a schedule.
// extends SchedulePolicy
type SnapshotPolicyStruct struct {
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// True if this is the default policy created when the system is
	// setup. Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Whether this policy has been directly applied or inherited. See
	// the effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// List of Schedule objects representing when the policy will
	// execute.
	// create = required
	// update = optional
	ScheduleList []*ScheduleStruct `json:"scheduleList,omitempty"`
	// The timezone of this policy. If not specified, defaults to the
	// Delphix Engine's timezone.
	// create = optional
	// update = optional
	Timezone *TimeZoneStruct `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SnapshotRuntimeStruct - Runtime properties of a TimeFlow snapshot.
// extends TypedObject
type SnapshotRuntimeStruct struct {
	// True if this snapshot can be used as the basis for provisioning a
	// new TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SnapshotSpaceMapStruct - Mapping of containers and snapshots to their respective space usage.
// extends TypedObject
type SnapshotSpaceMapStruct struct {
	// Amount of space, per object, in bytes, that it uses.
	SizeMap map[string]string `json:"sizeMap,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SnapshotSpaceParametersStruct - Input to the operation to determine how much space is used by a set of
// snapshots.
// extends TypedObject
type SnapshotSpaceParametersStruct struct {
	// FilesystemObjects to query, in canonical object reference form.
	// required = true
	ObjectReferences []string `json:"objectReferences,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SnapshotSpaceResultStruct - Result of the operation to determine how much space is used by a set
// of snapshots.
// extends TypedObject
type SnapshotSpaceResultStruct struct {
	// Total amount of space, in bytes, that would be freed by deleting
	// the input snapshots.
	// units = B
	// base = 1024
	TotalSize float64 `json:"totalSize,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// SourceConfigConnectivityStruct - Mechanism to test JDBC connectivity to source configs.
// extends AbstractSourceConfigConnectivity
type SourceConfigConnectivityStruct struct {
	// Database password.
	// format = password
	// required = true
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Database username.
	// required = true
	Username string `json:"username,omitempty"`
}

// SourceDisableParametersStruct - The parameters to use as input to disable a MSSQL, PostgreSQL,
// AppData, ASE or MySQL source.
// extends TypedObject
type SourceDisableParametersStruct struct {
	// Whether to attempt a cleanup of the database from the environment
	// before the disable.
	// default = true
	AttemptCleanup *bool `json:"attemptCleanup,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SourceEnableParametersStruct - The parameters to use as input to enable a MSSQL, PostgreSQL, AppData,
// ASE or MySQL source.
// extends TypedObject
type SourceEnableParametersStruct struct {
	// Whether to attempt a startup of the source after the enable.
	// default = true
	AttemptStart *bool `json:"attemptStart,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SourceRepositoryTemplateStruct - The representation of a repository template object.
// extends NamedUserObject
type SourceRepositoryTemplateStruct struct {
	// The reference to the database container.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The reference to the target repository.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-source-repository.json
	Repository string `json:"repository,omitempty"`
	// The reference to the associated template.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-database-template.json
	Template string `json:"template,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SourceStartParametersStruct - The parameters to use as input to start a MSSQL, PostgreSQL, AppData,
// ASE or MySQL source.
// extends TypedObject
type SourceStartParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SourceStopParametersStruct - The parameters to use as input to stop a MSSQL, PostgreSQL, AppData,
// ASE or MySQL source.
// extends TypedObject
type SourceStopParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SourceUpgradeParametersStruct - The parameters to use as input to upgrade a source.
// extends TypedObject
type SourceUpgradeParametersStruct struct {
	// The source config that the source database upgrades to.
	// required = true
	SourceConfig SourceConfig `json:"sourceConfig,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// SourcingPolicyStruct - Database policies for managing SnapSync and LogSync across sources for
// a MSSQL container.
// extends TypedObject
type SourcingPolicyStruct struct {
	// True if LogSync should run for this database.
	// create = optional
	// update = optional
	// default = false
	LogsyncEnabled *bool `json:"logsyncEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SplunkHecConfigStruct - Splunk HTTP Event Collector specific configuration information.
// extends UserObject
// cliVisibility = [SYSTEM]
type SplunkHecConfigStruct struct {
	// Whether we should send metrics data to Splunk.
	// default = true
	// create = required
	// update = optional
	EnableMetrics *bool `json:"enableMetrics,omitempty"`
	// Whether to use HTTPS to connect to Splunk. This should correspond
	// to your HTTP Event Collector settings in Splunk.
	// update = optional
	// create = required
	EnableSSL *bool `json:"enableSSL,omitempty"`
	// Whether we should send Delphix Insight data to Splunk using this
	// configuration.
	// update = optional
	// default = false
	// create = required
	Enabled *bool `json:"enabled,omitempty"`
	// The frequency in number of seconds at which the Events will be
	// pushed to Splunk. Defaults to 60 seconds.
	// minimum = 5
	// maximum = 3600
	// default = 60
	// create = optional
	// update = optional
	EventsPushFrequency *int `json:"eventsPushFrequency,omitempty"`
	// The TCP port number for the Splunk HTTP Event Collector (HEC).
	// maximum = 65535
	// create = required
	// update = optional
	// minimum = 1
	HecPort *int `json:"hecPort,omitempty"`
	// The token for the Splunk HTTP Event Collector (HEC).
	// update = optional
	// create = required
	HecToken string `json:"hecToken,omitempty"`
	// Splunk host name or IP address.
	// create = required
	// update = optional
	Host string `json:"host,omitempty"`
	// The Splunk Index events will be sent to. Must be set as an allowed
	// index for the HEC token.
	// create = required
	// update = optional
	MainIndex string `json:"mainIndex,omitempty"`
	// The Splunk Index metrics will be sent to. Must be set as an
	// allowed index for the HEC token. If none is specified the
	// mainIndex will be used for metrics as well.
	// create = optional
	// update = optional
	MetricsIndex string `json:"metricsIndex,omitempty"`
	// The frequency in number of seconds at which the Performance
	// Metrics will be pushed to Splunk. Defaults to 60 seconds.
	// maximum = 3600
	// default = 60
	// create = optional
	// update = optional
	// minimum = 5
	MetricsPushFrequency *int `json:"metricsPushFrequency,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The resolution of performance metrics data sent to Splunk. The
	// options are SECOND for 1-second resolution, or MINUTE for 1-minute
	// resolution.
	// update = optional
	// enum = [SECOND MINUTE]
	// default = MINUTE
	// create = optional
	PerformanceMetricsResolution string `json:"performanceMetricsResolution,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SshAcceptAlwaysStruct - Key-verification strategy that always accepts the host's key.
// extends SshVerificationStrategy
type SshAcceptAlwaysStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SshVerifyFingerprintStruct - SSH verification strategy based on a known per-host fingerprint.
// extends SshVerifyBase
type SshVerifyFingerprintStruct struct {
	// Base-64 encoded fingerprint of the ssh key of the host.
	// required = true
	// format = hostFingerprint
	Fingerprint string `json:"fingerprint,omitempty"`
	// Hash function for the fingerprint.
	// enum = [SHA256 SHA512]
	// required = true
	FingerprintType string `json:"fingerprintType,omitempty"`
	// Type of ssh key.
	// enum = [RSA DSA ECDSA ED25519]
	// required = true
	KeyType string `json:"keyType,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// SshVerifyRawKeyStruct - SSH verification strategy based on a known per-host key.
// extends SshVerifyBase
type SshVerifyRawKeyStruct struct {
	// Type of ssh key.
	// enum = [RSA DSA ECDSA ED25519]
	// required = true
	KeyType string `json:"keyType,omitempty"`
	// Base64-encoded ssh key of the host.
	// required = true
	// format = hostKey
	RawKey string `json:"rawKey,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StagingCompatibilityParametersStruct - The criteria necessary to select valid repositories for staging.
// extends CompatibleRepositoriesParameters
type StagingCompatibilityParametersStruct struct {
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The repository to use as a source of compatibility information.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	SourceRepository string `json:"sourceRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StatisticStruct - Multidimensional analytics statistics which can be queried for data.
// extends TypedObject
// cliVisibility = [DOMAIN]
type StatisticStruct struct {
	// The set of axes this statistic has.
	Axes []*StatisticAxisStruct `json:"axes,omitempty"`
	// A deeper explanation of the data this can collect.
	Explanation string `json:"explanation,omitempty"`
	// The smallest unit of time this statistic can measure on.
	// units = sec
	MinCollectionInterval *int `json:"minCollectionInterval,omitempty"`
	// The type name for the data this can collect.
	StatisticType string `json:"statisticType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StatisticAxisStruct - The attributes of a statistic axis.
// extends TypedObject
type StatisticAxisStruct struct {
	// The name for this axis.
	AxisName string `json:"axisName,omitempty"`
	// The type of constraint that can be applied to this axis.
	ConstraintType string `json:"constraintType,omitempty"`
	// A deeper explanation of the data this corresponds to.
	Explanation string `json:"explanation,omitempty"`
	// Whether this axis appears as an attribute of a datapoint stream or
	// of datapoints themselves.
	StreamAttribute *bool `json:"streamAttribute,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The type of value this axis will have for collected data.
	// enum = [INTEGER BOOLEAN STRING HISTOGRAM]
	ValueType string `json:"valueType,omitempty"`
}

// StatisticEnumAxisStruct - The attributes of a statistic axis which is an enum type.
// extends StatisticAxis
type StatisticEnumAxisStruct struct {
	// The name for this axis.
	AxisName string `json:"axisName,omitempty"`
	// The type of constraint that can be applied to this axis.
	ConstraintType string `json:"constraintType,omitempty"`
	// The set of values that are allowed for this axis.
	EnumValues []string `json:"enumValues,omitempty"`
	// A deeper explanation of the data this corresponds to.
	Explanation string `json:"explanation,omitempty"`
	// Whether this axis appears as an attribute of a datapoint stream or
	// of datapoints themselves.
	StreamAttribute *bool `json:"streamAttribute,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The type of value this axis will have for collected data.
	// enum = [INTEGER BOOLEAN STRING HISTOGRAM]
	ValueType string `json:"valueType,omitempty"`
}

// StatisticSliceStruct - Collects a slice of a multidimensional analytics statistic.
// extends NamedUserObject
// cliVisibility = [DOMAIN SYSTEM]
type StatisticSliceStruct struct {
	// Axis constraints act as per-axis filters on data that is being
	// collected.
	// create = optional
	AxisConstraints []AxisConstraint `json:"axisConstraints,omitempty"`
	// The set of axes to collect (usually these are not constrained
	// axes).
	// create = required
	CollectionAxes []string `json:"collectionAxes,omitempty"`
	// The minimum interval between each reading for this statistic.
	// create = optional
	// units = sec
	CollectionInterval *int `json:"collectionInterval,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Collection state of the slice.
	// enum = [INITIALIZED RUNNING PAUSED FAILED]
	State string `json:"state,omitempty"`
	// The type name for the data this can collect.
	// create = required
	StatisticType string `json:"statisticType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StorageDeviceStruct - A storage device on the system.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type StorageDeviceStruct struct {
	// True if the device is currently configured in the system.
	Configured *bool `json:"configured,omitempty"`
	// Model ID of the device.
	Model string `json:"model,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Serial number of the device.
	Serial string `json:"serial,omitempty"`
	// Physical size of the device, in bytes.
	// base = 1024
	// units = B
	Size float64 `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Vendor ID of the device.
	Vendor string `json:"vendor,omitempty"`
}

// StorageDeviceInitializeStatusStruct - The status of an initialize operation of a storage device in the
// system.
// extends TypedObject
type StorageDeviceInitializeStatusStruct struct {
	// Amount of data initialized, in bytes.
	// base = 1024
	// units = B
	BytesDone float64 `json:"bytesDone,omitempty"`
	// Total amount of data to initialize (including data already
	// initialized), in bytes.
	// base = 1024
	// units = B
	BytesEst float64 `json:"bytesEst,omitempty"`
	// Initialization state.
	// enum = [NONE ACTIVE CANCELED SUSPENDED COMPLETED]
	State string `json:"state,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StorageDeviceRemovalStatusStruct - The status of a device removal of the storage in the system.
// extends TypedObject
// cliVisibility = [SYSTEM]
type StorageDeviceRemovalStatusStruct struct {
	// Amount of data removed, in bytes.
	// base = 1024
	// units = B
	Copied float64 `json:"copied,omitempty"`
	// Memory used to account for removed devices, in bytes.
	// base = 1024
	// units = B
	MappingMemory float64 `json:"mappingMemory,omitempty"`
	// Time removal was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Removal state.
	// enum = [NONE ACTIVE COMPLETED CANCELED]
	State string `json:"state,omitempty"`
	// Total amount of data to remove (including completed), in bytes.
	// base = 1024
	// units = B
	Total float64 `json:"total,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StorageDeviceRemovalVerifyResultStruct - The .
// extends TypedObject
type StorageDeviceRemovalVerifyResultStruct struct {
	// Free space of the pool if this device is removed, in bytes.
	// base = 1024
	// units = B
	NewFreeBytes float64 `json:"newFreeBytes,omitempty"`
	// Amount of memory to be used by mappings if this device is removed,
	// in bytes.
	// base = 1024
	// units = B
	NewMappingMemory float64 `json:"newMappingMemory,omitempty"`
	// Free space of the pool before this device is removed, in bytes.
	// base = 1024
	// units = B
	OldFreeBytes float64 `json:"oldFreeBytes,omitempty"`
	// Amount of memory used by removal mappings before this device is
	// removed, in bytes.
	// base = 1024
	// units = B
	OldMappingMemory float64 `json:"oldMappingMemory,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StorageTestStruct - Test the performance of storage devices.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type StorageTestStruct struct {
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The parameters used to execute the test.
	Parameters *StorageTestParametersStruct `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// The state of the test.
	// enum = [WAITING RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// The results assigned to various tests.
	TestResults []*StorageTestResultStruct `json:"testResults,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StorageTestParametersStruct - Parameters used to execute a storage test.
// extends TypedObject
type StorageTestParametersStruct struct {
	// The list of devices to be used for the test.
	// create = optional
	Devices []string `json:"devices,omitempty"`
	// Run time of each test, in seconds.
	// minimum = 1
	// maximum = 3600
	// default = 120
	// create = optional
	Duration *int `json:"duration,omitempty"`
	// True if the disks should be initialized prior to running the
	// benchmark.
	// default = true
	// create = optional
	InitializeDevices *bool `json:"initializeDevices,omitempty"`
	// True if the entire disk should be initialized prior to running the
	// benchmark.
	// default = false
	// create = optional
	InitializeEntireDevice *bool `json:"initializeEntireDevice,omitempty"`
	// Total disk space, spread over all devices, used by the test (in
	// bytes).
	// create = optional
	// units = B
	// base = 1024
	// minimum = 1.048576e+06
	// default = 5.49755813888e+11
	TestRegion float64 `json:"testRegion,omitempty"`
	// The tests that are to be run.
	// default = ALL
	// create = required
	// enum = [ALL MINIMAL READ WRITE RANDREAD]
	Tests string `json:"tests,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// StorageTestResultStruct - The test results of one storage test.
// extends TypedObject
type StorageTestResultStruct struct {
	// Average latency in milliseconds.
	AverageLatency float64 `json:"averageLatency,omitempty"`
	// Block size used for the test.
	// units = B
	// base = 1024
	BlockSize *int `json:"blockSize,omitempty"`
	// IO operations per second.
	Iops *int `json:"iops,omitempty"`
	// No of jobs/threads used.
	Jobs *int `json:"jobs,omitempty"`
	// 95th percentile latency in milliseconds.
	Latency95thPercentile float64 `json:"latency95thPercentile,omitempty"`
	// Grade assigned to the test for latency.
	LatencyGrade string `json:"latencyGrade,omitempty"`
	// Load scaling.
	LoadScaling float64 `json:"loadScaling,omitempty"`
	// Grade assigned to the test for load scaling.
	LoadScalingGrade string `json:"loadScalingGrade,omitempty"`
	// Maximum latency in milliseconds.
	MaxLatency float64 `json:"maxLatency,omitempty"`
	// Minimum latency in milliseconds.
	MinLatency float64 `json:"minLatency,omitempty"`
	// Standard deviation of latency in milliseconds.
	StddevLatency float64 `json:"stddevLatency,omitempty"`
	// Name of the test for which the grade is assigned.
	TestName string `json:"testName,omitempty"`
	// The test type.
	// enum = [READ WRITE RANDREAD RANDWRITE]
	TestType string `json:"testType,omitempty"`
	// Throughput.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StringEqualConstraintStruct - Constraint placed on a string axis of a particular analytics slice.
// extends StringConstraint
type StringEqualConstraintStruct struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must match this string.
	// create = required
	Equals string `json:"equals,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SupportAccessStateStruct - The state of the access to the support shell.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SupportAccessStateStruct struct {
	// How the support shell can be accessed.
	// update = required
	// enum = [DISABLED ENABLED_NO_TOKEN ENABLED_WITH_TOKEN]
	AccessType string `json:"accessType,omitempty"`
	// If ENABLED_WITH_TOKEN, time that the token will no longer be
	// valid.
	// format = date
	// update = optional
	EndTime string `json:"endTime,omitempty"`
	// If ENABLED_WITH_TOKEN, the time that the token will be valid.
	// format = date
	// update = optional
	StartTime string `json:"startTime,omitempty"`
	// If ENABLED_WITH_TOKEN, the token that must be supplied to login.
	// update = optional
	Token string `json:"token,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SupportBundleGenerateParametersStruct - Parameters to be used when generating a support bundle.
// extends BaseSupportBundleParameters
type SupportBundleGenerateParametersStruct struct {
	// Type of support bundle to generate. Reserved for Delphix support
	// use.
	// default = ALL
	// enum = [PHONEHOME MDS OS CORE LOG DROPBOX STORAGE_TEST MASKING ALL]
	BundleType string `json:"bundleType,omitempty"`
	// The list of environments from which logs should be collected.
	// required = false
	// uniqueItems = true
	Environments []string `json:"environments,omitempty"`
	// Whether or not to include the analytics data in the support bundle
	// which is generated. Including analytics data may significantly
	// increase the support bundle size and upload time, but enables
	// analysis of performance characteristics of the Delphix Engine.
	// default = false
	IncludeAnalyticsData *bool `json:"includeAnalyticsData,omitempty"`
	// The list of sources from which logs should be collected.
	// uniqueItems = true
	// required = false
	Sources []string `json:"sources,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// SupportBundleUploadParametersStruct - Parameters to be used when uploading a support bundle.
// extends BaseSupportBundleParameters
type SupportBundleUploadParametersStruct struct {
	// Type of support bundle to generate. Reserved for Delphix support
	// use.
	// default = ALL
	// enum = [PHONEHOME MDS OS CORE LOG DROPBOX STORAGE_TEST MASKING ALL]
	BundleType string `json:"bundleType,omitempty"`
	// The Delphix support case number.
	CaseNumber *int `json:"caseNumber,omitempty"`
	// The list of environments from which logs should be collected.
	// required = false
	// uniqueItems = true
	Environments []string `json:"environments,omitempty"`
	// Whether or not to include the analytics data in the support bundle
	// which is generated. Including analytics data may significantly
	// increase the support bundle size and upload time, but enables
	// analysis of performance characteristics of the Delphix Engine.
	// default = false
	IncludeAnalyticsData *bool `json:"includeAnalyticsData,omitempty"`
	// The list of sources from which logs should be collected.
	// required = false
	// uniqueItems = true
	Sources []string `json:"sources,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SupportContactStruct - A support telephone number mapped to a given country. May include text
// (e.g. 1-800-FOR-HELP).
// extends TypedObject
type SupportContactStruct struct {
	// An ISO country code, as recognized by modern browsers for
	// resolving locale.
	// enum = [-- AF AL DZ AS AD AO AQ AG AR AM AW AU AT AZ BS BH BD BB BY BE BZ BJ BM BT BO BA BW BV BR IO BN BG BF BI KH CM CA CV KY CF TD CL CN CX CC CO KM CG CD CK CR CI HR CU CY CZ DK DJ DM DO EC EG SV GQ ER EE ET FK FO FJ FI FR GF PF TF GA GM GE DE GH GI GR GL GD GP GU GT GN GW GY HT HM HN HK HU IS IN ID IR IQ IE IL IT JM JP JO KZ KE KI KP KR KW KG LA LV LB LS LR LY LI LT LU MO MK MG MW MY MV ML MT MH MQ MR MU YT MX FM MD MC MN ME MS MA MZ MM NA NR NP NL AN NC NZ NI NE NG NU NF MP NO OM PK PW PS PA PG PY PE PH PN PL PT PR QA RE RO RU RW SH KN LC PM VC WS SM ST SA SN RS SC SL SG SK SI SB SO ZA GS ES LK SD SR SJ SZ SE CH SY TW TJ TZ TH TL TG TK TO TT TN TR TM TC TV UG UA AE GB US UM UY UZ VU VE VN VG VI WF EH YE ZM ZW]
	// required = true
	Country string `json:"country,omitempty"`
	// A telephone number, formatted in accordance with the norms of the
	// associated country.
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SwitchTimeflowParametersStruct - The input parameters used for the TimeFlow switch operation.
// extends TypedObject
type SwitchTimeflowParametersStruct struct {
	// The reference to the target TimeFlow that should be made the
	// current TimeFlow.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SyncParametersStruct - The parameters to use as input to sync requests.
// extends TypedObject
type SyncParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SyncPolicyStruct - This policy syncs a container (runs SnapSync) according to the given
// schedule.
// extends SchedulePolicy
type SyncPolicyStruct struct {
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// True if this is the default policy created when the system is
	// setup. Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Whether this policy has been directly applied or inherited. See
	// the effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// List of Schedule objects representing when the policy will
	// execute.
	// create = required
	// update = optional
	ScheduleList []*ScheduleStruct `json:"scheduleList,omitempty"`
	// The timezone of this policy. If not specified, defaults to the
	// Delphix Engine's timezone.
	// create = optional
	// update = optional
	Timezone *TimeZoneStruct `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SyslogConfigStruct - Syslog configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SyslogConfigStruct struct {
	// True if the syslog service is enabled.
	// update = optional
	// default = true
	Enabled *bool `json:"enabled,omitempty"`
	// Syslog message format.
	// update = optional
	// enum = [TEXT JSON]
	// default = TEXT
	Format string `json:"format,omitempty"`
	// Syslog logging pattern. Events will be logged in the pattern as
	// specified.
	// update = optional
	// default = %-5p delphix : %m%n
	Pattern string `json:"pattern,omitempty"`
	// List of syslog servers. At least one syslog server must be
	// specified.
	// update = required
	Servers []*SyslogServerStruct `json:"servers,omitempty"`
	// Syslog logging severity. Only events at or above this severity
	// will be logged.
	// update = optional
	// enum = [EMERGENCY ALERT CRITICAL ERROR WARNING NOTICE INFORMATIONAL DEBUG]
	// default = WARNING
	Severity string `json:"severity,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SyslogServerStruct - Syslog server configuration.
// extends TypedObject
type SyslogServerStruct struct {
	// Syslog host name or IP address.
	// format = host
	// required = true
	Address string `json:"address,omitempty"`
	// Syslog port number.
	// required = true
	// minimum = 0
	// maximum = 65535
	// default = 514
	Port *int `json:"port,omitempty"`
	// Syslog transport protocol.
	// required = true
	// enum = [udp tcp]
	// default = udp
	Protocol string `json:"protocol,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// SystemInfoStruct - Retrieve system-wide properties and manage the state of the system.
// extends PublicSystemInfo
// cliVisibility = [DOMAIN SYSTEM]
type SystemInfoStruct struct {
	// Maximum supported API version of the current system software.
	ApiVersion *APIVersionStruct `json:"apiVersion,omitempty"`
	// Security banner to display prior to login.
	Banner string `json:"banner,omitempty"`
	// Time at which the current system software was built.
	// format = date
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
	// Description of the current system software.
	BuildTitle string `json:"buildTitle,omitempty"`
	// Delphix version of the current system software.
	BuildVersion *VersionInfoStruct `json:"buildVersion,omitempty"`
	// Indicates whether the server has gone through initial setup or
	// not.
	Configured *bool `json:"configured,omitempty"`
	// Percentage of CPU reserved on the host.
	// units = %
	CpuReservation float64 `json:"cpuReservation,omitempty"`
	// The current system locale.
	// format = locale
	CurrentLocale string `json:"currentLocale,omitempty"`
	// The list of enabled features on this Delphix Engine.
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`
	// Qualifier for referencing instances of (e.g. 'Delphix') engines in
	// messages like 'The <engineQualifier> Engine failed to ...'.
	// create = required
	EngineQualifier string `json:"engineQualifier,omitempty"`
	// System hostname.
	// update = optional
	// format = hostname
	Hostname string `json:"hostname,omitempty"`
	// The date and time that the Delphix Engine was installed.
	// format = date
	InstallationTime string `json:"installationTime,omitempty"`
	// The operating system kernel name.
	KernelName string `json:"kernelName,omitempty"`
	// List of available locales.
	Locales []string `json:"locales,omitempty"`
	// Amount of memory reserved on the host.
	// units = B
	// base = 1024
	MemoryReservation float64 `json:"memoryReservation,omitempty"`
	// Total memory on the system, in bytes.
	// units = B
	// base = 1024
	MemorySize float64 `json:"memorySize,omitempty"`
	// Description of the current system platform.
	Platform string `json:"platform,omitempty"`
	// Processors on the system.
	Processors []*CPUInfoStruct `json:"processors,omitempty"`
	// Name of the product that the system is running.
	ProductName string `json:"productName,omitempty"`
	// Product type.
	ProductType string `json:"productType,omitempty"`
	// SSH public key to be added to SSH authorized_keys for environment
	// users using the SystemKeyCredential authorization mechanism.
	SshPublicKey string `json:"sshPublicKey,omitempty"`
	// Total amount of raw storage allocated for dSources, VDBs, and
	// system metadata. Zero if storage has not yet been configured.
	// units = B
	// base = 1024
	StorageTotal float64 `json:"storageTotal,omitempty"`
	// Amount of raw storage used by dSources, VDBs and system metadata.
	// base = 1024
	// units = B
	StorageUsed float64 `json:"storageUsed,omitempty"`
	// Technical support phone numbers.
	// create = optional
	SupportContacts []*SupportContactStruct `json:"supportContacts,omitempty"`
	// Technical Support URL.
	// create = optional
	SupportURL string `json:"supportURL,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Delphix Engine up time.
	UpTime *UpTimeInfoStruct `json:"upTime,omitempty"`
	// Globally unique identifier for this software installation.
	Uuid string `json:"uuid,omitempty"`
	// Address of vendor headquarters. Free form collection of strings to
	// accomodate any region.
	// create = optional
	VendorAddress []string `json:"vendorAddress,omitempty"`
	// Corporate headquarters email address.
	// create = optional
	// format = email
	VendorEmail string `json:"vendorEmail,omitempty"`
	// Vendor name, for use in messages like 'Please contact <vendorName>
	// customer support'.
	// create = required
	VendorName string `json:"vendorName,omitempty"`
	// Corporate headquarters telephone number.
	// create = optional
	VendorPhoneNumber string `json:"vendorPhoneNumber,omitempty"`
	// Corporate home page.
	// create = optional
	VendorURL string `json:"vendorURL,omitempty"`
}

// SystemInitializationParametersStruct - Parameters used for intializing an engine.
// extends TypedObject
type SystemInitializationParametersStruct struct {
	// Password to use for the default domain administrator.
	// format = password
	// default = delphix
	// create = required
	DefaultPassword string `json:"defaultPassword,omitempty"`
	// Name of the default domain administrator to create.
	// pattern = ^[a-zA-Z][-_.a-zA-Z0-9]*$
	// minLength = 1
	// maxLength = 256
	// default = delphix_admin
	// create = optional
	DefaultUser string `json:"defaultUser,omitempty"`
	// List of storage devices to use.
	// create = required
	Devices []string `json:"devices,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SystemKeyCredentialStruct - The system public key based security credential.
// extends PublicKeyCredential
type SystemKeyCredentialStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SystemPackageStruct - A package whose version can be changed by sysadmins.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type SystemPackageStruct struct {
	// Package name.
	// format = objectName
	// update = readonly
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Possible versions for this package.
	// update = readonly
	PossibleVersions []string `json:"possibleVersions,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Current version of the package.
	// update = required
	Version string `json:"version,omitempty"`
}

// SystemVersionStruct - Describes a Delphix software revision.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type SystemVersionStruct struct {
	// Date on which the version was built.
	// format = date
	BuildDate string `json:"buildDate,omitempty"`
	// Date on which this version was installed.
	// format = date
	InstallDate string `json:"installDate,omitempty"`
	// The minimum DelphixOS version supported by this Delphix version.
	MinOsVersion string `json:"minOsVersion,omitempty"`
	// The minimum required Delphix version in order to upgrade.
	MinVersion string `json:"minVersion,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// DelphixOS is running from this version.
	OsRunning *bool `json:"osRunning,omitempty"`
	// The DelphixOS version number.
	OsVersion string `json:"osVersion,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The state of the version.
	// enum = [PREVIOUS CURRENTLY_RUNNING DEFERRED UPLOADED UNPACKING DELETING VERIFYING VERIFIED APPLYING UNKNOWN DISABLE_FAILED]
	Status string `json:"status,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Date on which this version was last verified.
	// format = date
	VerifyDate string `json:"verifyDate,omitempty"`
	// The Delphix version number.
	Version string `json:"version,omitempty"`
}

// TCPStatsDatapointStruct - An analytics datapoint generated by the TCP_STATS statistic type.
// extends Datapoint
type TCPStatsDatapointStruct struct {
	// The size of the local congestion window.
	// units = B
	// base = 1024
	CongestionWindowSize *int `json:"congestionWindowSize,omitempty"`
	// Data bytes received.
	// units = B
	// base = 1024
	InBytes *int `json:"inBytes,omitempty"`
	// Number of bytes received out of order. This is a subset of the
	// 'inBytes' value.
	// units = B
	// base = 1024
	InUnorderedBytes *int `json:"inUnorderedBytes,omitempty"`
	// Data bytes transmitted.
	// units = B
	// base = 1024
	OutBytes *int `json:"outBytes,omitempty"`
	// The size of the local receive window.
	// units = B
	// base = 1024
	ReceiveWindowSize *int `json:"receiveWindowSize,omitempty"`
	// Bytes retransmitted.
	// units = B
	// base = 1024
	RetransmittedBytes *int `json:"retransmittedBytes,omitempty"`
	// The smoothed average round trip time for this connection (us).
	// units = usec
	RoundTripTime *int `json:"roundTripTime,omitempty"`
	// The size of the peer's receive window.
	// units = B
	// base = 1024
	SendWindowSize *int `json:"sendWindowSize,omitempty"`
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Number of bytes sent but unacknowledged.
	// units = B
	// base = 1024
	UnacknowledgedBytes *int `json:"unacknowledgedBytes,omitempty"`
	// Number of bytes in the transmit queue that have not been sent.
	// units = B
	// base = 1024
	UnsentBytes *int `json:"unsentBytes,omitempty"`
}

// TCPStatsDatapointStreamStruct - A stream of datapoints from a TCP_STATS analytics slice.
// extends DatapointStream
type TCPStatsDatapointStreamStruct struct {
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// The local Delphix Engine IP address.
	// format = ipAddress
	LocalAddress string `json:"localAddress,omitempty"`
	// The local TCP port number.
	LocalPort *int `json:"localPort,omitempty"`
	// The remote IP address.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// The remote TCP port number.
	RemotePort *int `json:"remotePort,omitempty"`
	// The network service.
	Service string `json:"service,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// TargetFilterStruct - An alert filter that specifies which targets to match against.
// extends AlertFilter
type TargetFilterStruct struct {
	// List of object references. Only alerts related to one of the
	// targets or its children are included.
	// create = optional
	// update = optional
	// minItems = 1
	Targets []string `json:"targets,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TargetOwnerFilterStruct - An alert filter which matches when an alert's target is owned by one
// of the specified users.
// extends AlertFilter
type TargetOwnerFilterStruct struct {
	// Target owners to match against.
	// minItems = 1
	// create = required
	// update = optional
	Owners []string `json:"owners,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeConfigStruct - Get and set the current time configuration.
// extends TypedObject
// cliVisibility = [SYSTEM DOMAIN]
type TimeConfigStruct struct {
	// Current system time. This value can only be set if NTP is
	// disabled. The management service is automatically restarted if the
	// time is changed.
	// update = optional
	// format = date
	CurrentTime string `json:"currentTime,omitempty"`
	// NTP configuration.
	// update = optional
	NtpConfig *NTPConfigStruct `json:"ntpConfig,omitempty"`
	// Default time zone for system wide policies and schedules. The
	// management service is automatically restarted if the timezone is
	// changed.
	// update = optional
	// default = Etc/UTC
	SystemTimeZone string `json:"systemTimeZone,omitempty"`
	// The difference, in minutes, between UTC and local time. For
	// example, if your time zone is UTC -5:00 (Eastern Standard Time),
	// 300 will be returned. Daylight saving time prevents this value
	// from being a constant even for a given locale.
	SystemTimeZoneOffset *int `json:"systemTimeZoneOffset,omitempty"`
	// System time zone offset as a String. For instance 'UTC -5:00'.
	SystemTimeZoneOffsetString string `json:"systemTimeZoneOffsetString,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeRangeParametersStruct - The parameters to use as input for methods requiring a time range.
// extends TypedObject
type TimeRangeParametersStruct struct {
	// The date at the end of the period.
	// format = date
	// create = required
	EndTime string `json:"endTime,omitempty"`
	// The date at the beginning of the period.
	// create = required
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeZoneStruct - This represents a time zone offset.
// extends TypedObject
type TimeZoneStruct struct {
	// The ID of this time zone.
	// enum = [ACT AET AGT ART AST Africa/Abidjan Africa/Accra Africa/Addis_Ababa Africa/Algiers Africa/Asmara Africa/Asmera Africa/Bamako Africa/Bangui Africa/Banjul Africa/Bissau Africa/Blantyre Africa/Brazzaville Africa/Bujumbura Africa/Cairo Africa/Casablanca Africa/Ceuta Africa/Conakry Africa/Dakar Africa/Dar_es_Salaam Africa/Djibouti Africa/Douala Africa/El_Aaiun Africa/Freetown Africa/Gaborone Africa/Harare Africa/Johannesburg Africa/Juba Africa/Kampala Africa/Khartoum Africa/Kigali Africa/Kinshasa Africa/Lagos Africa/Libreville Africa/Lome Africa/Luanda Africa/Lubumbashi Africa/Lusaka Africa/Malabo Africa/Maputo Africa/Maseru Africa/Mbabane Africa/Mogadishu Africa/Monrovia Africa/Nairobi Africa/Ndjamena Africa/Niamey Africa/Nouakchott Africa/Ouagadougou Africa/Porto-Novo Africa/Sao_Tome Africa/Timbuktu Africa/Tripoli Africa/Tunis Africa/Windhoek America/Adak America/Anchorage America/Anguilla America/Antigua America/Araguaina America/Argentina/Buenos_Aires America/Argentina/Catamarca America/Argentina/ComodRivadavia America/Argentina/Cordoba America/Argentina/Jujuy America/Argentina/La_Rioja America/Argentina/Mendoza America/Argentina/Rio_Gallegos America/Argentina/Salta America/Argentina/San_Juan America/Argentina/San_Luis America/Argentina/Tucuman America/Argentina/Ushuaia America/Aruba America/Asuncion America/Atikokan America/Atka America/Bahia America/Bahia_Banderas America/Barbados America/Belem America/Belize America/Blanc-Sablon America/Boa_Vista America/Bogota America/Boise America/Buenos_Aires America/Cambridge_Bay America/Campo_Grande America/Cancun America/Caracas America/Catamarca America/Cayenne America/Cayman America/Chicago America/Chihuahua America/Coral_Harbour America/Cordoba America/Costa_Rica America/Creston America/Cuiaba America/Curacao America/Danmarkshavn America/Dawson America/Dawson_Creek America/Denver America/Detroit America/Dominica America/Edmonton America/Eirunepe America/El_Salvador America/Ensenada America/Fort_Nelson America/Fort_Wayne America/Fortaleza America/Glace_Bay America/Godthab America/Goose_Bay America/Grand_Turk America/Grenada America/Guadeloupe America/Guatemala America/Guayaquil America/Guyana America/Halifax America/Havana America/Hermosillo America/Indiana/Indianapolis America/Indiana/Knox America/Indiana/Marengo America/Indiana/Petersburg America/Indiana/Tell_City America/Indiana/Vevay America/Indiana/Vincennes America/Indiana/Winamac America/Indianapolis America/Inuvik America/Iqaluit America/Jamaica America/Jujuy America/Juneau America/Kentucky/Louisville America/Kentucky/Monticello America/Knox_IN America/Kralendijk America/La_Paz America/Lima America/Los_Angeles America/Louisville America/Lower_Princes America/Maceio America/Managua America/Manaus America/Marigot America/Martinique America/Matamoros America/Mazatlan America/Mendoza America/Menominee America/Merida America/Metlakatla America/Mexico_City America/Miquelon America/Moncton America/Monterrey America/Montevideo America/Montreal America/Montserrat America/Nassau America/New_York America/Nipigon America/Nome America/Noronha America/North_Dakota/Beulah America/North_Dakota/Center America/North_Dakota/New_Salem America/Ojinaga America/Panama America/Pangnirtung America/Paramaribo America/Phoenix America/Port-au-Prince America/Port_of_Spain America/Porto_Acre America/Porto_Velho America/Puerto_Rico America/Punta_Arenas America/Rainy_River America/Rankin_Inlet America/Recife America/Regina America/Resolute America/Rio_Branco America/Rosario America/Santa_Isabel America/Santarem America/Santiago America/Santo_Domingo America/Sao_Paulo America/Scoresbysund America/Shiprock America/Sitka America/St_Barthelemy America/St_Johns America/St_Kitts America/St_Lucia America/St_Thomas America/St_Vincent America/Swift_Current America/Tegucigalpa America/Thule America/Thunder_Bay America/Tijuana America/Toronto America/Tortola America/Vancouver America/Virgin America/Whitehorse America/Winnipeg America/Yakutat America/Yellowknife Antarctica/Casey Antarctica/Davis Antarctica/DumontDUrville Antarctica/Macquarie Antarctica/Mawson Antarctica/McMurdo Antarctica/Palmer Antarctica/Rothera Antarctica/South_Pole Antarctica/Syowa Antarctica/Troll Antarctica/Vostok Arctic/Longyearbyen Asia/Aden Asia/Almaty Asia/Amman Asia/Anadyr Asia/Aqtau Asia/Aqtobe Asia/Ashgabat Asia/Ashkhabad Asia/Atyrau Asia/Baghdad Asia/Bahrain Asia/Baku Asia/Bangkok Asia/Barnaul Asia/Beirut Asia/Bishkek Asia/Brunei Asia/Calcutta Asia/Chita Asia/Choibalsan Asia/Chongqing Asia/Chungking Asia/Colombo Asia/Dacca Asia/Damascus Asia/Dhaka Asia/Dili Asia/Dubai Asia/Dushanbe Asia/Famagusta Asia/Gaza Asia/Harbin Asia/Hebron Asia/Ho_Chi_Minh Asia/Hong_Kong Asia/Hovd Asia/Irkutsk Asia/Istanbul Asia/Jakarta Asia/Jayapura Asia/Jerusalem Asia/Kabul Asia/Kamchatka Asia/Karachi Asia/Kashgar Asia/Kathmandu Asia/Katmandu Asia/Khandyga Asia/Kolkata Asia/Krasnoyarsk Asia/Kuala_Lumpur Asia/Kuching Asia/Kuwait Asia/Macao Asia/Macau Asia/Magadan Asia/Makassar Asia/Manila Asia/Muscat Asia/Nicosia Asia/Novokuznetsk Asia/Novosibirsk Asia/Omsk Asia/Oral Asia/Phnom_Penh Asia/Pontianak Asia/Pyongyang Asia/Qatar Asia/Qyzylorda Asia/Rangoon Asia/Riyadh Asia/Saigon Asia/Sakhalin Asia/Samarkand Asia/Seoul Asia/Shanghai Asia/Singapore Asia/Srednekolymsk Asia/Taipei Asia/Tashkent Asia/Tbilisi Asia/Tehran Asia/Tel_Aviv Asia/Thimbu Asia/Thimphu Asia/Tomsk Asia/Tokyo Asia/Ujung_Pandang Asia/Ulaanbaatar Asia/Ulan_Bator Asia/Urumqi Asia/Ust-Nera Asia/Vientiane Asia/Vladivostok Asia/Yakutsk Asia/Yangon Asia/Yekaterinburg Asia/Yerevan Atlantic/Azores Atlantic/Bermuda Atlantic/Canary Atlantic/Cape_Verde Atlantic/Faeroe Atlantic/Faroe Atlantic/Jan_Mayen Atlantic/Madeira Atlantic/Reykjavik Atlantic/South_Georgia Atlantic/St_Helena Atlantic/Stanley Australia/ACT Australia/Adelaide Australia/Brisbane Australia/Broken_Hill Australia/Canberra Australia/Currie Australia/Darwin Australia/Eucla Australia/Hobart Australia/LHI Australia/Lindeman Australia/Lord_Howe Australia/Melbourne Australia/NSW Australia/North Australia/Perth Australia/Queensland Australia/South Australia/Sydney Australia/Tasmania Australia/Victoria Australia/West Australia/Yancowinna BET BST Brazil/Acre Brazil/DeNoronha Brazil/East Brazil/West CAT CET CNT CST CST6CDT CTT Canada/Atlantic Canada/Central Canada/Eastern Canada/Mountain Canada/Newfoundland Canada/Pacific Canada/Saskatchewan Canada/Yukon Chile/Continental Chile/EasterIsland Cuba EAT ECT EET EST EST5EDT Egypt Eire Etc/GMT Etc/GMT+0 Etc/GMT+1 Etc/GMT+10 Etc/GMT+11 Etc/GMT+12 Etc/GMT+2 Etc/GMT+3 Etc/GMT+4 Etc/GMT+5 Etc/GMT+6 Etc/GMT+7 Etc/GMT+8 Etc/GMT+9 Etc/GMT-0 Etc/GMT-1 Etc/GMT-10 Etc/GMT-11 Etc/GMT-12 Etc/GMT-13 Etc/GMT-14 Etc/GMT-2 Etc/GMT-3 Etc/GMT-4 Etc/GMT-5 Etc/GMT-6 Etc/GMT-7 Etc/GMT-8 Etc/GMT-9 Etc/GMT0 Etc/Greenwich Etc/UCT Etc/UTC Etc/Universal Etc/Zulu Europe/Amsterdam Europe/Andorra Europe/Astrakhan Europe/Athens Europe/Belfast Europe/Belgrade Europe/Berlin Europe/Bratislava Europe/Brussels Europe/Bucharest Europe/Budapest Europe/Busingen Europe/Chisinau Europe/Copenhagen Europe/Dublin Europe/Gibraltar Europe/Guernsey Europe/Helsinki Europe/Isle_of_Man Europe/Istanbul Europe/Jersey Europe/Kaliningrad Europe/Kiev Europe/Kirov Europe/Lisbon Europe/Ljubljana Europe/London Europe/Luxembourg Europe/Madrid Europe/Malta Europe/Mariehamn Europe/Minsk Europe/Monaco Europe/Moscow Europe/Nicosia Europe/Oslo Europe/Paris Europe/Podgorica Europe/Prague Europe/Riga Europe/Rome Europe/Samara Europe/San_Marino Europe/Sarajevo Europe/Saratov Europe/Simferopol Europe/Skopje Europe/Sofia Europe/Stockholm Europe/Tallinn Europe/Tirane Europe/Tiraspol Europe/Ulyanovsk Europe/Uzhgorod Europe/Vaduz Europe/Vatican Europe/Vienna Europe/Vilnius Europe/Volgograd Europe/Warsaw Europe/Zagreb Europe/Zaporozhye Europe/Zurich GB GB-Eire GMT GMT0 Greenwich HST Hongkong IET IST Iceland Indian/Antananarivo Indian/Chagos Indian/Christmas Indian/Cocos Indian/Comoro Indian/Kerguelen Indian/Mahe Indian/Maldives Indian/Mauritius Indian/Mayotte Indian/Reunion Iran Israel JST Jamaica Japan Kwajalein Libya MET MIT MST MST7MDT Mexico/BajaNorte Mexico/BajaSur Mexico/General NET NST NZ NZ-CHAT Navajo PLT PNT PRC PRT PST PST8PDT Pacific/Apia Pacific/Auckland Pacific/Bougainville Pacific/Chatham Pacific/Chuuk Pacific/Easter Pacific/Efate Pacific/Enderbury Pacific/Fakaofo Pacific/Fiji Pacific/Funafuti Pacific/Galapagos Pacific/Gambier Pacific/Guadalcanal Pacific/Guam Pacific/Honolulu Pacific/Johnston Pacific/Kiritimati Pacific/Kosrae Pacific/Kwajalein Pacific/Majuro Pacific/Marquesas Pacific/Midway Pacific/Nauru Pacific/Niue Pacific/Norfolk Pacific/Noumea Pacific/Pago_Pago Pacific/Palau Pacific/Pitcairn Pacific/Pohnpei Pacific/Ponape Pacific/Port_Moresby Pacific/Rarotonga Pacific/Saipan Pacific/Samoa Pacific/Tahiti Pacific/Tarawa Pacific/Tongatapu Pacific/Truk Pacific/Wake Pacific/Wallis Pacific/Yap Poland Portugal ROK SST Singapore SystemV/AST4 SystemV/AST4ADT SystemV/CST6 SystemV/CST6CDT SystemV/EST5 SystemV/EST5EDT SystemV/HST10 SystemV/MST7 SystemV/MST7MDT SystemV/PST8 SystemV/PST8PDT SystemV/YST9 SystemV/YST9YDT Turkey UCT US/Alaska US/Aleutian US/Arizona US/Central US/East-Indiana US/Eastern US/Hawaii US/Indiana-Starke US/Michigan US/Mountain US/Pacific US/Pacific-New US/Samoa UTC Universal VST W-SU WET Zulu]
	// required = true
	Id string `json:"id,omitempty"`
	// The difference, in minutes, between UTC and local time. For
	// example, if your time zone is UTC -5:00 (Eastern Standard Time),
	// 300 will be returned. Daylight saving time prevents this value
	// from being a constant even for a given locale.
	Offset *int `json:"offset,omitempty"`
	// The Offset as a String. For instance 'UTC -5:00'.
	OffsetString string `json:"offsetString,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowBookmarkStruct - A TimeFlow bookmark is a user defined name for a TimeFlow point
// (location or timestamp within a TimeFlow).
// extends NamedUserObject
type TimeflowBookmarkStruct struct {
	// The TimeFlow location.
	Location string `json:"location,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Indicates whether retention should be allowed to clean up the
	// TimeFlow bookmark and associated data.
	// update = optional
	RetentionProof *bool `json:"retentionProof,omitempty"`
	// A tag for the bookmark that can be used to group TimeFlow
	// bookmarks together or qualify the type of the bookmark.
	// maxLength = 64
	Tag string `json:"tag,omitempty"`
	// Reference to the TimeFlow for this bookmark.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowBookmarkCreateParametersStruct - The parameters to use as input to create TimeFlow bookmarks.
// extends TypedObject
type TimeflowBookmarkCreateParametersStruct struct {
	// The bookmark name.
	// required = true
	Name string `json:"name,omitempty"`
	// Indicates whether retention should be allowed to clean up the
	// TimeFlow bookmark and associated data.
	// create = optional
	RetentionProof *bool `json:"retentionProof,omitempty"`
	// A tag for the bookmark that can be used to group bookmarks
	// together or qualify the type of the bookmark.
	// create = optional
	// maxLength = 64
	Tag string `json:"tag,omitempty"`
	// The TimeFlow point which is referenced by this bookmark.
	// required = true
	TimeflowPoint TimeflowPoint `json:"timeflowPoint,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// TimeflowFilesystemLayoutStruct - A filesystem layout that matches the filesystem of a Delphix TimeFlow.
// extends FilesystemLayout
type TimeflowFilesystemLayoutStruct struct {
	// The directory for archive files.
	// create = optional
	ArchiveDirectory string `json:"archiveDirectory,omitempty"`
	// The directory for data files.
	// create = optional
	DataDirectory string `json:"dataDirectory,omitempty"`
	// The directory for external files.
	// create = optional
	ExternalDirectory string `json:"externalDirectory,omitempty"`
	// The directory for script files.
	// create = optional
	ScriptDirectory string `json:"scriptDirectory,omitempty"`
	// The base directory to use for the exported database.
	// create = optional
	TargetDirectory string `json:"targetDirectory,omitempty"`
	// The directory for temporary files.
	// create = optional
	TempDirectory string `json:"tempDirectory,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowLogFetchParametersStruct - Parameters to fetch log files within a TimeFlow.
// extends LogFetchSSH
type TimeflowLogFetchParametersStruct struct {
	// User credentials. If not provided will use environment credentials
	// for 'username' on 'host'.
	Credentials Credential `json:"credentials,omitempty"`
	// Directory on the remote server where the missing log files reside.
	// required = true
	Directory string `json:"directory,omitempty"`
	// The ending SCN of the range of log files to fetch.
	// required = true
	EndLocation string `json:"endLocation,omitempty"`
	// Remote host to connect to.
	// required = true
	Host string `json:"host,omitempty"`
	// SSH port to connect to.
	// default = 22
	Port *int `json:"port,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// The starting SCN of the range of log files to fetch.
	// required = true
	StartLocation string `json:"startLocation,omitempty"`
	// Reference to the TimeFlow for which to fetch logs.
	// referenceTo = /delphix-oracle-timeflow.json
	// required = true
	// format = objectReference
	Timeflow string `json:"timeflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User name to authenticate as.
	// required = true
	Username string `json:"username,omitempty"`
}

// TimeflowPointBookmarkStruct - TimeFlow point based on a TimeFlow bookmark.
// extends TimeflowPointParameters
type TimeflowPointBookmarkStruct struct {
	// Reference to the bookmark.
	// format = objectReference
	// referenceTo = /delphix-timeflow-bookmark.json
	// required = true
	Bookmark string `json:"bookmark,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowPointBookmarkTagStruct - TimeFlow point based on a TimeFlow bookmark tag.
// extends TimeflowPointParameters
type TimeflowPointBookmarkTagStruct struct {
	// Reference to the container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	// required = true
	Container string `json:"container,omitempty"`
	// The name of the tag.
	// required = true
	Tag string `json:"tag,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowPointLocationStruct - TimeFlow point based on a database-specific identifier (SCN, LSN,
// etc).
// extends TimeflowPointParameters
type TimeflowPointLocationStruct struct {
	// The TimeFlow location.
	// required = true
	// pattern = ^[\S]+$
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this location.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowPointSemanticStruct - TimeFlow point based on a semantic reference.
// extends TimeflowPointParameters
type TimeflowPointSemanticStruct struct {
	// Reference to the container.
	// referenceTo = /delphix-container.json
	// create = optional
	// update = optional
	// format = objectReference
	Container string `json:"container,omitempty"`
	// A semantic description of a TimeFlow location.
	// create = optional
	// update = optional
	// enum = [LATEST_POINT LATEST_SNAPSHOT]
	// default = LATEST_POINT
	Location string `json:"location,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowPointSnapshotStruct - TimeFlow point based on a snapshot reference.
// extends TimeflowPointParameters
type TimeflowPointSnapshotStruct struct {
	// Reference to the snapshot.
	// referenceTo = /delphix-timeflow-snapshot.json
	// required = true
	// format = objectReference
	Snapshot string `json:"snapshot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowPointTimestampStruct - TimeFlow point based on a timestamp.
// extends TimeflowPointParameters
type TimeflowPointTimestampStruct struct {
	// Reference to TimeFlow containing this point.
	// referenceTo = /delphix-timeflow.json
	// required = true
	// format = objectReference
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// required = true
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// TimeflowRangeStruct - Time range within a TimeFlow.
// extends TypedObject
type TimeflowRangeStruct struct {
	// The ending TimeFlow point of this range.
	EndPoint TimeflowPoint `json:"endPoint,omitempty"`
	// Whether or not this TimeFlow range is provisionable.
	Provisionable *bool `json:"provisionable,omitempty"`
	// The starting TimeFlow point of this range.
	StartPoint TimeflowPoint `json:"startPoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowRangeParametersStruct - The parameters to use as input to fetch TimeFlow ranges.
// extends TypedObject
type TimeflowRangeParametersStruct struct {
	// The ending TimeFlow point of the time period to search for
	// TimeFlow ranges.
	EndPoint TimeflowPoint `json:"endPoint,omitempty"`
	// The starting TimeFlow point of the time period to search for
	// TimeFlow ranges.
	StartPoint TimeflowPoint `json:"startPoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowRepairSSHStruct - Parameters to repair log files within a TimeFlow.
// extends TimeflowRepairParameters
type TimeflowRepairSSHStruct struct {
	// User credentials. If not provided will use environment credentials
	// for 'username' on 'host'.
	Credentials Credential `json:"credentials,omitempty"`
	// Directory on the remote server where the missing log files reside.
	// required = true
	Directory string `json:"directory,omitempty"`
	// The ending point of the range of log files to fetch.
	// required = true
	EndLocation string `json:"endLocation,omitempty"`
	// Remote host to connect to.
	// required = true
	Host string `json:"host,omitempty"`
	// SSH port to connect to.
	// default = 22
	Port *int `json:"port,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// The starting point of the range of log files to fetch.
	// required = true
	StartLocation string `json:"startLocation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User name to authenticate as.
	// required = true
	Username string `json:"username,omitempty"`
}

// TimeflowSnapshotDayRangeStruct - Count of TimeFlow snapshots aggregated by day.
// extends TypedObject
type TimeflowSnapshotDayRangeStruct struct {
	// Number of TimeFlow snapshots on that day.
	Count float64 `json:"count,omitempty"`
	// Date for which TimeFlow snapshots have been aggregated.
	Date string `json:"date,omitempty"`
	// End of day of this range in the time zone used for computation.
	// format = date
	EndOfDay string `json:"endOfDay,omitempty"`
	// Start of day of this range in the time zone used for computation.
	// format = date
	StartOfDay string `json:"startOfDay,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ToolkitStruct - An installed toolkit.
// extends NamedUserObject
type ToolkitStruct struct {
	// The Delphix API version that the toolkit was built against.
	// required = true
	BuildApi *APIVersionStruct `json:"buildApi,omitempty"`
	// The default locale for this toolkit. This locale defines the set
	// of all message IDs for the toolkit and serves as the fallback
	// locale when messages cannot be localized in a particular locale.
	// If no messages are specified for the toolkit, the defaultLocale
	// may be any locale.
	// required = true
	// format = locale
	DefaultLocale string `json:"defaultLocale,omitempty"`
	// Definition of how to discover sources of this type.
	// required = false
	DiscoveryDefinition *ToolkitDiscoveryDefinitionStruct `json:"discoveryDefinition,omitempty"`
	// A list of host types compatible with this toolkit.
	// required = true
	HostTypes []string `json:"hostTypes,omitempty"`
	// Implementation language for workflows in this toolkit.
	// required = true
	// enum = [LUA PYTHON27]
	Language string `json:"language,omitempty"`
	// Definition of how to link sources of this type.
	// required = true
	LinkedSourceDefinition ToolkitLinkedSource `json:"linkedSourceDefinition,omitempty"`
	// The set of localizable messages for this toolkit.
	// required = false
	Messages []*ToolkitLocaleStruct `json:"messages,omitempty"`
	// A unique and descriptive name for the toolkit.
	// pattern = ^[a-z0-9_:-]+$
	// required = true
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A human readable name for the toolkit.
	// required = true
	// maxLength = 256
	PrettyName string `json:"prettyName,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Resources for use by workflows in this toolkit.
	// required = true
	Resources map[string]string `json:"resources,omitempty"`
	// Schema for metadata collected during snapshotting.
	// required = true
	SnapshotSchema *SchemaDraftV4Struct `json:"snapshotSchema,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Definition of how to upgrade sources of this type.
	// required = false
	UpgradeDefinition *ToolkitUpgradeDefinitionStruct `json:"upgradeDefinition,omitempty"`
	// The version of the toolkit that is of the form:
	// 'major.minor.patch'.
	// format = toolkitVersion
	// required = true
	Version string `json:"version,omitempty"`
	// Definition of how to provision virtual sources of this type.
	// required = true
	VirtualSourceDefinition *ToolkitVirtualSourceStruct `json:"virtualSourceDefinition,omitempty"`
}

// ToolkitDiscoveryDefinitionStruct - Defines the discovery schemas and workflow scripts for a toolkit.
// extends TypedObject
type ToolkitDiscoveryDefinitionStruct struct {
	// True if this toolkit supports manual discovery of source configs.
	// required = false
	ManualSourceConfigDiscovery *bool `json:"manualSourceConfigDiscovery,omitempty"`
	// A workflow script that discovers repositories on a target
	// environment. The script must return a list of repositories
	// matching the repositorySchema.
	// required = true
	RepositoryDiscovery string `json:"repositoryDiscovery,omitempty"`
	// A list of fields in the repositorySchema that collectively
	// identify each discovered repository.
	// required = true
	RepositoryIdentityFields []string `json:"repositoryIdentityFields,omitempty"`
	// The field of the repositorySchema to display to the end user for
	// naming this repository.
	// required = true
	RepositoryNameField string `json:"repositoryNameField,omitempty"`
	// A user defined schema to represent the repository.
	// required = true
	RepositorySchema *SchemaDraftV4Struct `json:"repositorySchema,omitempty"`
	// A workflow script that discovers source configs on a target
	// environment. The script must return a list of source configs
	// matching the sourceConfigSchema.
	// required = true
	SourceConfigDiscovery string `json:"sourceConfigDiscovery,omitempty"`
	// A list of fields in the sourceConfigSchema that collectively
	// identify each discovered source config.
	// required = true
	SourceConfigIdentityFields []string `json:"sourceConfigIdentityFields,omitempty"`
	// The field of the sourceConfigSchema to display to the end user for
	// naming this source config.
	// required = true
	SourceConfigNameField string `json:"sourceConfigNameField,omitempty"`
	// A user defined schema to represent the source config.
	// required = true
	SourceConfigSchema *SchemaDraftV4Struct `json:"sourceConfigSchema,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ToolkitLinkedDirectSourceStruct - A linked source definition for toolkits that perform linking using
// direct file sync.
// extends ToolkitLinkedSource
type ToolkitLinkedDirectSourceStruct struct {
	// A user defined schema for the linking parameters.
	// required = true
	Parameters *SchemaDraftV4Struct `json:"parameters,omitempty"`
	// A workflow script to run immediately after snapshotting the staged
	// source.
	// required = true
	PostSnapshot string `json:"postSnapshot,omitempty"`
	// A workflow script to run just prior to snapshotting the staged
	// source.
	// required = true
	PreSnapshot string `json:"preSnapshot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this toolkit requires old-style AppData-defined
	// properties.
	// create = optional
	UsesGrandfatheredAppDataProperties *bool `json:"usesGrandfatheredAppDataProperties,omitempty"`
}

// ToolkitLinkedStagedSourceStruct - A linked staged source definition for toolkits.
// extends ToolkitLinkedSource
type ToolkitLinkedStagedSourceStruct struct {
	// A workflow script that specifies where the storage for the copy of
	// the application should be mounted.
	// required = false
	MountSpec string `json:"mountSpec,omitempty"`
	// A workflow script that specifies which user/group should own the
	// mount where the application will be copied.
	// required = false
	OwnershipSpec string `json:"ownershipSpec,omitempty"`
	// A user defined schema for the linking parameters.
	// required = true
	Parameters *SchemaDraftV4Struct `json:"parameters,omitempty"`
	// A workflow script to run immediately after snapshotting the staged
	// source.
	// required = true
	PostSnapshot string `json:"postSnapshot,omitempty"`
	// A workflow script to run just prior to snapshotting the staged
	// source.
	// required = true
	PreSnapshot string `json:"preSnapshot,omitempty"`
	// A workflow script that builds the staging instance from
	// production.
	// required = true
	Resync string `json:"resync,omitempty"`
	// A workflow script that start the staged source. The staged files
	// will be mounted and available.
	// required = true
	StartStaging string `json:"startStaging,omitempty"`
	// A workflow script that returns whether or not the data source is
	// active/inactive. The script should exit with an exit status of
	// ACTIVE if the data source is available. The script should exit
	// with an exit status of INACTIVE if the data source is unavailable.
	// An exit status of UNKNOWN implies the script encountered an
	// unexpected state or error. If no status script is supplied, the
	// dSource will always be in an active state while enabled.
	// required = false
	Status string `json:"status,omitempty"`
	// A workflow script that stop the staged source. The staged files
	// will be mounted and available. Upon completion of this workflow,
	// the staged files will be unmounted.
	// required = true
	StopStaging string `json:"stopStaging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A workflow script run periodically to monitor the health of the
	// data source and staging environment. This script will be run every
	// 10 seconds.
	// required = false
	Worker string `json:"worker,omitempty"`
}

// ToolkitLocaleStruct - Contains a mapping from message IDs to messages for a locale.
// extends TypedObject
type ToolkitLocaleStruct struct {
	// The name of this locale.
	// format = locale
	// required = true
	LocaleName string `json:"localeName,omitempty"`
	// A mapping of message IDs to messages for this locale.
	// required = true
	Messages map[string]string `json:"messages,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ToolkitUpgradeDefinitionStruct - The toolkit upgrade logic to upgrade metadata from the previous
// version of the toolkit.
// extends TypedObject
type ToolkitUpgradeDefinitionStruct struct {
	// The version of the toolkit that we are upgrading from.
	// format = toolkitVersion
	// required = true
	FromVersion string `json:"fromVersion,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A workflow script to upgrade the linked source parameters.
	// required = true
	UpgradeLinkedSource string `json:"upgradeLinkedSource,omitempty"`
	// A workflow script to upgrade the manually-discovered source config
	// parameters.
	// required = false
	UpgradeManualSourceConfig string `json:"upgradeManualSourceConfig,omitempty"`
	// A workflow script to upgrade the snapshot metadata.
	// required = true
	UpgradeSnapshot string `json:"upgradeSnapshot,omitempty"`
	// A workflow script to upgrade the virtual source parameters.
	// required = true
	UpgradeVirtualSource string `json:"upgradeVirtualSource,omitempty"`
}

// ToolkitVirtualSourceStruct - A virtual source definition for toolkits.
// extends TypedObject
type ToolkitVirtualSourceStruct struct {
	// A workflow script run when configuring a virtual copy of the
	// application in a new environment.
	// required = true
	Configure string `json:"configure,omitempty"`
	// A workflow script to run when creating an empty application.
	// required = false
	Initialize string `json:"initialize,omitempty"`
	// A workflow script that specifies where the virtual copy of the
	// application should be mounted.
	// required = false
	MountSpec string `json:"mountSpec,omitempty"`
	// A workflow script that specifies which user/group should own the
	// files inside the virtual copy of the application.
	// required = false
	OwnershipSpec string `json:"ownershipSpec,omitempty"`
	// A user defined schema for the provisioning parameters.
	// required = true
	Parameters *SchemaDraftV4Struct `json:"parameters,omitempty"`
	// A workflow script to run after taking a snapshot of a virtual copy
	// of the application.
	// required = true
	PostSnapshot string `json:"postSnapshot,omitempty"`
	// A workflow script to run before taking a snapshot of a virtual
	// copy of the application.
	// required = true
	PreSnapshot string `json:"preSnapshot,omitempty"`
	// A workflow script run when returning a virtual copy of the
	// appliction to an environment that it was previously removed from.
	// required = true
	Reconfigure string `json:"reconfigure,omitempty"`
	// A workflow script to run when starting a virtual copy of the
	// application.
	// required = true
	Start string `json:"start,omitempty"`
	// The workflow script to run to determine if a virtual copy of the
	// application is running. The script should output 'ACTIVE' if the
	// application is running, 'INACTIVE' if the application is not
	// running, and 'UNKNOWN' if the script encounters an unexpected
	// problem.
	// required = false
	Status string `json:"status,omitempty"`
	// A workflow script to run when stopping a virtual copy of the
	// application.
	// required = true
	Stop string `json:"stop,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A workflow script run when removing a virtual copy of the
	// application from an environment (e.g. on delete, disable, or
	// refresh).
	// required = true
	Unconfigure string `json:"unconfigure,omitempty"`
}

// TracerouteInfoStruct - Trace route info from target host to Delphix Engine.
// extends TypedObject
type TracerouteInfoStruct struct {
	// Latency of network hops from host to Delphix Engine.
	NetworkHops string `json:"networkHops,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TransformationStruct - A data platform agnostic transformation object.
// extends UserObject
type TransformationStruct struct {
	// A reference to the container which is a transformed version of the
	// parent container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Reference to the user used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Operations to perform when this transformation is applied.
	Operations []SourceOperation `json:"operations,omitempty"`
	// Platform-specific parameters that are stored on a transformation.
	PlatformParams BasePlatformParameters `json:"platformParams,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UnixHostStruct - The representation of a Unix host object.
// extends Host
type UnixHostStruct struct {
	// The address associated with the host.
	// format = host
	// create = required
	// update = optional
	Address string `json:"address,omitempty"`
	// The date the host was added.
	DateAdded string `json:"dateAdded,omitempty"`
	// The lowercase alias to use inside the user managed DSP keystore.
	// create = optional
	// update = optional
	// minLength = 1
	DspKeystoreAlias string `json:"dspKeystoreAlias,omitempty"`
	// The password for the user managed DSP keystore.
	// create = optional
	// update = optional
	// minLength = 1
	// format = password
	DspKeystorePassword string `json:"dspKeystorePassword,omitempty"`
	// The path to the user managed DSP keystore.
	// create = optional
	// update = optional
	// minLength = 1
	DspKeystorePath string `json:"dspKeystorePath,omitempty"`
	// The password for the user managed DSP truststore.
	// update = optional
	// minLength = 1
	// format = password
	// create = optional
	DspTruststorePassword string `json:"dspTruststorePassword,omitempty"`
	// The path to the user managed DSP truststore.
	// create = optional
	// update = optional
	// minLength = 1
	DspTruststorePath string `json:"dspTruststorePath,omitempty"`
	// The host configuration object associated with the host.
	HostConfiguration *HostConfigurationStruct `json:"hostConfiguration,omitempty"`
	// Runtime properties for this host.
	HostRuntime *HostRuntimeStruct `json:"hostRuntime,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The list of host/IP addresses to use for NFS export.
	// uniqueItems = true
	// create = optional
	// update = optional
	NfsAddressList []string `json:"nfsAddressList,omitempty"`
	// The password for the user managed Oracle JDBC keystore.
	// format = password
	// create = optional
	// update = optional
	// minLength = 1
	OracleJdbcKeystorePassword string `json:"oracleJdbcKeystorePassword,omitempty"`
	// Profile for escalating user privileges.
	// create = optional
	// update = optional
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	// format = objectReference
	PrivilegeElevationProfile string `json:"privilegeElevationProfile,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The port number used to connect to the host via SSH.
	// maximum = 65535
	// create = optional
	// update = optional
	// default = 22
	// minimum = 1
	SshPort *int `json:"sshPort,omitempty"`
	// Mechanism to use for ssh host verification.
	// create = optional
	// update = optional
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// The path for the toolkit that resides on the host.
	// create = required
	// update = optional
	ToolkitPath string `json:"toolkitPath,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UnixHostCreateParametersStruct - The parameters used for the add Unix host operation.
// extends HostCreateParameters
type UnixHostCreateParametersStruct struct {
	// The host object.
	// create = required
	Host Host `json:"host,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UnixHostEnvironmentStruct - The representation of a unix host environment object.
// extends HostEnvironment
type UnixHostEnvironmentStruct struct {
	// Parameters for an environment with SAP ASE instances.
	// create = optional
	// update = optional
	AseHostEnvironmentParameters *ASEHostEnvironmentParametersStruct `json:"aseHostEnvironmentParameters,omitempty"`
	// The environment description.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The reference to the associated host.
	// format = objectReference
	// referenceTo = /delphix-host.json
	Host string `json:"host,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source
	// environment.
	// create = optional
	// update = optional
	// default = false
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A reference to the primary user for this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	PrimaryUser string `json:"primaryUser,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UpTimeInfoStruct - Information about the up time of OS and management stack.
// extends TypedObject
type UpTimeInfoStruct struct {
	// Management Stack up time.
	Mgmt string `json:"mgmt,omitempty"`
	// OS up time.
	Os string `json:"os,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UpgradeCheckResultStruct - Describes unsatisfied upgrade requirements.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type UpgradeCheckResultStruct struct {
	// A localized, textual description of the action the user should
	// take to overcome the error.
	// create = readonly
	// update = readonly
	Action string `json:"action,omitempty"`
	// A unique identifier for the type of the upgrade check result.
	// update = readonly
	// create = readonly
	BundleId string `json:"bundleId,omitempty"`
	// A localized, textual description of the error.
	// create = readonly
	// update = readonly
	Description string `json:"description,omitempty"`
	// A localized, textual description of the impact the error might
	// have on the system.
	// create = readonly
	// update = readonly
	Impact string `json:"impact,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Script output related to the check result to assist in resolving
	// the issue.
	// create = readonly
	// update = readonly
	Output string `json:"output,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The severity of the missing upgrade requirement. CRITICAL check
	// results block the upgrade.
	// update = readonly
	// enum = [WARNING CRITICAL INFORMATIONAL]
	// create = readonly
	Severity string `json:"severity,omitempty"`
	// The status of the upgrade check result.
	// enum = [ACTIVE IGNORED RESOLVED]
	// create = readonly
	// update = readonly
	Status string `json:"status,omitempty"`
	// A localized string that is the broad category of this check.
	// create = readonly
	// update = readonly
	Title string `json:"title,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A reference to the upgrade version that generated this check
	// result.
	// format = objectReference
	// referenceTo = /delphix-upgrade-version.json
	// create = readonly
	// update = readonly
	Version string `json:"version,omitempty"`
}

// UpgradeCheckResultsVersionParametersStruct - Parameters used to modify an upgradeCheckResult. These parameters are
// mutually exclusive.
// extends TypedObject
type UpgradeCheckResultsVersionParametersStruct struct {
	// BundleID of upgrade check result(s).
	// create = optional
	BundleId string `json:"bundleId,omitempty"`
	// Reference to a single upgrade check result.
	// referenceTo = /delphix-upgrade-check-result.json
	// create = optional
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UpgradeCompatibilityParametersStruct - The criteria necessary to select valid repositories for upgrading.
// extends CompatibleRepositoriesParameters
type UpgradeCompatibilityParametersStruct struct {
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The repository to use as a source of compatibility information.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	SourceRepository string `json:"sourceRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UpgradeNotificationStruct - An object to track when an upgrade is happening to trigger UI
// response.
// extends Notification
type UpgradeNotificationStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UserStruct - Delphix users.
// extends NamedUserObject
// cliVisibility = [DOMAIN SYSTEM]
type UserStruct struct {
	// User authentication type.
	// enum = [LDAP NATIVE SAML]
	// create = optional
	AuthenticationType string `json:"authenticationType,omitempty"`
	// Credential used for authentication.
	// create = optional
	Credential *PasswordCredentialStruct `json:"credential,omitempty"`
	// Email address for the user.
	// maxLength = 256
	// create = optional
	// update = optional
	// format = email
	EmailAddress string `json:"emailAddress,omitempty"`
	// True if the user is currently enabled and can log into the system.
	// default = true
	// create = optional
	Enabled *bool `json:"enabled,omitempty"`
	// First name of user.
	// maxLength = 64
	// create = optional
	// update = optional
	FirstName string `json:"firstName,omitempty"`
	// Home phone number of user.
	// update = optional
	// maxLength = 32
	// create = optional
	HomePhoneNumber string `json:"homePhoneNumber,omitempty"`
	// True if this is the default user and cannot be deleted.
	// create = optional
	IsDefault *bool `json:"isDefault,omitempty"`
	// Last name of user.
	// create = optional
	// update = optional
	// maxLength = 64
	LastName string `json:"lastName,omitempty"`
	// Preferred locale as an IETF BCP 47 language tag, defaults to
	// 'en-US'.
	// maxLength = 16
	// create = optional
	// update = optional
	// format = locale
	// default = en-US
	Locale string `json:"locale,omitempty"`
	// Mobile phone number of user.
	// maxLength = 32
	// create = optional
	// update = optional
	MobilePhoneNumber string `json:"mobilePhoneNumber,omitempty"`
	// Name of user.
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether the user's password should be updated and why.
	// enum = [NONE FIRST_LOGIN ADMIN_REQUEST PASSWORD_POLICY]
	// default = NONE
	// update = optional
	PasswordUpdateRequest string `json:"passwordUpdateRequest,omitempty"`
	// Principal name used for authentication.
	// create = optional
	// update = optional
	Principal string `json:"principal,omitempty"`
	// Public key used for authentication.
	// create = optional
	// update = optional
	PublicKey string `json:"publicKey,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Session timeout in minutes.
	// default = 30
	// units = min
	// create = optional
	// update = optional
	// minimum = 1
	SessionTimeout *int `json:"sessionTimeout,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Type of user.
	// enum = [SYSTEM DOMAIN]
	// default = DOMAIN
	// create = optional
	UserType string `json:"userType,omitempty"`
	// Work phone number of user.
	// maxLength = 32
	// create = optional
	// update = optional
	WorkPhoneNumber string `json:"workPhoneNumber,omitempty"`
}

// UserAuthInfoStruct - Summary authorization information about the current user.
// extends TypedObject
type UserAuthInfoStruct struct {
	// The list of authorizations granted to the current user.
	Authorizations []*AuthorizationStruct `json:"authorizations,omitempty"`
	// A reference to the system-defined Self-Service user role.
	// format = objectReference
	// referenceTo = /delphix-role.json
	JetStreamUserRole string `json:"jetStreamUserRole,omitempty"`
	// A reference to the system-defined owner role.
	// format = objectReference
	// referenceTo = /delphix-role.json
	OwnerRole string `json:"ownerRole,omitempty"`
	// A reference to the system-defined provisioner role.
	// format = objectReference
	// referenceTo = /delphix-role.json
	ProvisionerRole string `json:"provisionerRole,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The currently logged in user.
	User *UserStruct `json:"user,omitempty"`
}

// UserInterfaceConfigStruct - User Interface Configuration.
// extends TypedObject
// cliVisibility = [DOMAIN SYSTEM]
type UserInterfaceConfigStruct struct {
	// Indicates whether user-interface analytics are enabled or not.
	// update = optional
	AnalyticsEnabled *bool `json:"analyticsEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// VMwarePlatformParametersStruct - VMware platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type VMwarePlatformParametersStruct struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ValidateSMTPParametersStruct - Validate SMTP configuration without committing it by sending mail to
// the specified address(es).
// extends TypedObject
type ValidateSMTPParametersStruct struct {
	// List of email addresses to send test email to.
	// required = true
	Addresses []string `json:"addresses,omitempty"`
	// SMTP configuration to use for validation.
	// required = true
	Config *SMTPConfigStruct `json:"config,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// VerifyVersionParametersStruct - The parameters to use as input to verifying an upgrade.
// extends TypedObject
type VerifyVersionParametersStruct struct {
	// If true, validates the success of upgrading the Delphix Engine
	// without updating the OSsoftware. The operation will detect a
	// failure if the upgrade version requires a version of the OSthat is
	// newer than what is currently running.
	// create = optional
	// default = false
	Defer *bool `json:"defer,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// VersionInfoStruct - Representation of a Delphix software revision.
// extends TypedObject
type VersionInfoStruct struct {
	// Major version number.
	Major *int `json:"major,omitempty"`
	// Micro version number.
	Micro *int `json:"micro,omitempty"`
	// Minor version number.
	Minor *int `json:"minor,omitempty"`
	// Patch version number.
	Patch *int `json:"patch,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// VfsOpsDatapointStreamStruct - A stream of datapoints from a VFS_OPS analytics slice.
// extends DatapointStream
type VfsOpsDatapointStreamStruct struct {
	// Whether reads were cached.
	Cached *bool `json:"cached,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Path of the affected file.
	// format = unixpath
	Path string `json:"path,omitempty"`
	// Whether writes were synchronous.
	Sync *bool `json:"sync,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// VirtualSourceOperationsStruct - Describes operations which are performed on virtual sources at various
// times.
// extends TypedObject
type VirtualSourceOperationsStruct struct {
	// Operations to perform when initially creating the virtual source
	// and every time it is refreshed.
	// create = optional
	// update = optional
	ConfigureClone []SourceOperation `json:"configureClone,omitempty"`
	// Operations to perform after refreshing a virtual source. These
	// operations can be used to restore any data or configuration backed
	// up in the preRefresh operations.
	// create = optional
	// update = optional
	PostRefresh []SourceOperation `json:"postRefresh,omitempty"`
	// Operations to perform after rewinding a virtual source. These
	// operations can be used to automate processes once the rewind is
	// complete.
	// create = optional
	// update = optional
	PostRollback []SourceOperation `json:"postRollback,omitempty"`
	// Operations to perform after snapshotting a virtual source.
	// create = optional
	// update = optional
	PostSnapshot []SourceOperation `json:"postSnapshot,omitempty"`
	// Operations to perform after starting a virtual source.
	// create = optional
	// update = optional
	PostStart []SourceOperation `json:"postStart,omitempty"`
	// Operations to perform after stopping a virtual source.
	// create = optional
	// update = optional
	PostStop []SourceOperation `json:"postStop,omitempty"`
	// Operations to perform before refreshing a virtual source. These
	// operations can backup any data or configuration from the running
	// source before doing the refresh.
	// create = optional
	// update = optional
	PreRefresh []SourceOperation `json:"preRefresh,omitempty"`
	// Operations to perform before rewinding a virtual source. These
	// operations can backup any data or configuration from the running
	// source prior to rewinding.
	// create = optional
	// update = optional
	PreRollback []SourceOperation `json:"preRollback,omitempty"`
	// Operations to perform before snapshotting a virtual source. These
	// operations can quiesce any data prior to snapshotting.
	// create = optional
	// update = optional
	PreSnapshot []SourceOperation `json:"preSnapshot,omitempty"`
	// Operations to perform before starting a virtual source.
	// create = optional
	// update = optional
	PreStart []SourceOperation `json:"preStart,omitempty"`
	// Operations to perform before stopping a virtual source.
	// create = optional
	// update = optional
	PreStop []SourceOperation `json:"preStop,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// WindowsClusterStruct - A Windows cluster environment.
// extends SourceEnvironment
type WindowsClusterStruct struct {
	// The address that will be used to perform discovery on the cluster.
	// create = required
	// update = optional
	Address string `json:"address,omitempty"`
	// The environment description.
	// update = optional
	// maxLength = 1024
	// create = optional
	Description string `json:"description,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source
	// environment.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// A reference to the primary user for this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	PrimaryUser string `json:"primaryUser,omitempty"`
	// A reference to the proxy that will be used to discover the
	// cluster.
	// create = required
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-host.json
	Proxy string `json:"proxy,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// A reference to the proxy host that will be used for cluster
	// support operations.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-host.json
	TargetProxy string `json:"targetProxy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// WindowsClusterCreateParametersStruct - The parameters used to create a Windows Cluster.
// extends SourceEnvironmentCreateParameters
type WindowsClusterCreateParametersStruct struct {
	// The representation of the cluster object.
	// create = required
	Cluster *WindowsClusterStruct `json:"cluster,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to the created source
	// environment.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// The primary user associated with the environment.
	// create = required
	PrimaryUser *EnvironmentUserStruct `json:"primaryUser,omitempty"`
	// Create as a target environment.
	// create = optional
	Target *bool `json:"target,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// WindowsClusterNodeStruct - A node in a Windows Cluster.
// extends NamedUserObject
type WindowsClusterNodeStruct struct {
	// A reference to the Windows cluster environment this node belongs
	// to.
	// format = objectReference
	// referenceTo = /delphix-windows-cluster.json
	Cluster string `json:"cluster,omitempty"`
	// Indicates whether the node is discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// A reference to the physical host.
	// format = objectReference
	// referenceTo = /delphix-host.json
	Host string `json:"host,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// WindowsHostStruct - The representation of a Windows host object.
// extends Host
type WindowsHostStruct struct {
	// The address associated with the host.
	// format = host
	// create = required
	// update = optional
	Address string `json:"address,omitempty"`
	// Unique per Delphix key used to authenticate with the remote
	// Delphix Connector.
	// create = optional
	// update = optional
	ConnectorAuthenticationKey string `json:"connectorAuthenticationKey,omitempty"`
	// The port that the connector connects on.
	// create = optional
	// update = optional
	ConnectorPort float64 `json:"connectorPort,omitempty"`
	// The date the host was added.
	DateAdded string `json:"dateAdded,omitempty"`
	// The lowercase alias to use inside the user managed DSP keystore.
	// create = optional
	// update = optional
	// minLength = 1
	DspKeystoreAlias string `json:"dspKeystoreAlias,omitempty"`
	// The password for the user managed DSP keystore.
	// minLength = 1
	// format = password
	// create = optional
	// update = optional
	DspKeystorePassword string `json:"dspKeystorePassword,omitempty"`
	// The path to the user managed DSP keystore.
	// create = optional
	// update = optional
	// minLength = 1
	DspKeystorePath string `json:"dspKeystorePath,omitempty"`
	// The password for the user managed DSP truststore.
	// create = optional
	// update = optional
	// minLength = 1
	// format = password
	DspTruststorePassword string `json:"dspTruststorePassword,omitempty"`
	// The path to the user managed DSP truststore.
	// create = optional
	// update = optional
	// minLength = 1
	DspTruststorePath string `json:"dspTruststorePath,omitempty"`
	// The host configuration object associated with the host.
	HostConfiguration *HostConfigurationStruct `json:"hostConfiguration,omitempty"`
	// Runtime properties for this host.
	HostRuntime *HostRuntimeStruct `json:"hostRuntime,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The list of host/IP addresses to use for NFS export.
	// create = optional
	// update = optional
	// uniqueItems = true
	NfsAddressList []string `json:"nfsAddressList,omitempty"`
	// Profile for escalating user privileges.
	// format = objectReference
	// create = optional
	// update = optional
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	PrivilegeElevationProfile string `json:"privilegeElevationProfile,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The port number used to connect to the host via SSH.
	// create = optional
	// update = optional
	// default = 22
	// minimum = 1
	// maximum = 65535
	SshPort *int `json:"sshPort,omitempty"`
	// The path for the toolkit that resides on the host.
	// create = readonly
	// update = readonly
	ToolkitPath string `json:"toolkitPath,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// WindowsHostCreateParametersStruct - The parameters used for the add Windows host operation.
// extends HostCreateParameters
type WindowsHostCreateParametersStruct struct {
	// The host object.
	// create = required
	Host Host `json:"host,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// WindowsHostEnvironmentStruct - The representation of a windows host environment object.
// extends HostEnvironment
type WindowsHostEnvironmentStruct struct {
	// The environment description.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The reference to the associated host.
	// referenceTo = /delphix-host.json
	// format = objectReference
	Host string `json:"host,omitempty"`
	// Flag indicating whether it is allowed to collect logs, potentially
	// containing sensitive information, related to this source
	// environment.
	// default = false
	// create = optional
	// update = optional
	LogCollectionEnabled *bool `json:"logCollectionEnabled,omitempty"`
	// Object name.
	// create = optional
	// update = optional
	// maxLength = 256
	// format = objectName
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A reference to the primary user for this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	PrimaryUser string `json:"primaryUser,omitempty"`
	// The reference to the proxy associated with the host.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	// update = optional
	Proxy string `json:"proxy,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// WorkflowFunctionDefinitionStruct - The definition of a Workflow Function.
// extends TypedObject
type WorkflowFunctionDefinitionStruct struct {
	// The input schema for this function according to DRAFTV4.
	// create = required
	// update = optional
	InputSchema *SchemaDraftV4Struct `json:"inputSchema,omitempty"`
	// The name of this function.
	// update = optional
	// create = required
	Name string `json:"name,omitempty"`
	// The output schema for this function according to DRAFTV4.
	// create = required
	// update = optional
	OutputSchema *SchemaDraftV4Struct `json:"outputSchema,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// X500DistinguishedNameCompositeStruct - The representation of a composite X.500 Distinguished Name.
// extends X500DistinguishedName
type X500DistinguishedNameCompositeStruct struct {
	// The composite Distinguished Name.
	// create = required
	Dname string `json:"dname,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// X500DistinguishedNameFieldsStruct - The representation of a X.500 Distinguished Name by separate fields.
// extends X500DistinguishedName
type X500DistinguishedNameFieldsStruct struct {
	// City/locality (L).
	// create = optional
	City string `json:"city,omitempty"`
	// Common name (CN).
	// create = required
	CommonName string `json:"commonName,omitempty"`
	// Country (C).
	// create = optional
	Country string `json:"country,omitempty"`
	// Organization (O).
	// create = optional
	Organization string `json:"organization,omitempty"`
	// Organization unit (OU).
	// create = optional
	OrganizationUnit string `json:"organizationUnit,omitempty"`
	// State/region (ST).
	// create = optional
	StateRegion string `json:"stateRegion,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// X509CertificateStruct - X509 Certificate.
// extends UserObject
// cliVisibility = [SYSTEM]
type X509CertificateStruct struct {
	// Delphix trusts this certificate .
	Accepted *bool `json:"accepted,omitempty"`
	// Issuer of this certificate.
	IssuedByDN string `json:"issuedByDN,omitempty"`
	// Distinguished name of subject of this certificate.
	IssuedToDN string `json:"issuedToDN,omitempty"`
	// MD5 fingerprint.
	Md5Fingerprint string `json:"md5Fingerprint,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Certificate serial number.
	SerialNumber string `json:"serialNumber,omitempty"`
	// SHA-1 fingerprint.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Start of validity.
	ValidFrom string `json:"validFrom,omitempty"`
	// End of validity.
	ValidTo string `json:"validTo,omitempty"`
}

// XPPCompatibilityParametersStruct - The criteria necessary to select valid repositories for cross-platform
// provisioning.
// extends CompatibleRepositoriesParameters
type XPPCompatibilityParametersStruct struct {
	// Restrict returned repositories to this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = optional
	// update = optional
	Environment string `json:"environment,omitempty"`
	// The TimeFlow point to use as a source of compatibility
	// information.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// XppLastRunStatusStruct - Status of the last cross-platform provision of the container.
// extends ChecklistItem
type XppLastRunStatusStruct struct {
	// Error message associated with the last run, if any.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Last cross-platform provision job run on the container.
	Job *JobStruct `json:"job,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// XppStagingStatusStruct - Status of the cross-platform provisioning staging environment.
// extends ChecklistItemDetail
type XppStagingStatusStruct struct {
	// Description of this status item.
	Description string `json:"description,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// XppStatusStruct - The current cross-platform provisioning status of a container.
// extends Checklist
type XppStatusStruct struct {
	// Status of the last cross-platform provision of the container.
	LastRunStatus *XppLastRunStatusStruct `json:"lastRunStatus,omitempty"`
	// Status of the cross-platform provisioning staging environment.
	StagingStatus *XppStagingStatusStruct `json:"stagingStatus,omitempty"`
	// Status of the cross-platform provisioning target environment.
	TargetStatus *XppTargetStatusStruct `json:"targetStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Cross-platform provisioning validation status of the container.
	ValidateStatus *XppValidateStatusStruct `json:"validateStatus,omitempty"`
}

// XppTargetStatusStruct - Status of the cross-platform provisioning target environment.
// extends ChecklistItemDetail
type XppTargetStatusStruct struct {
	// Description of this status item.
	Description string `json:"description,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// XppValidateStatusStruct - Cross-platform provisioning validation status of the container.
// extends ChecklistItemDetail
type XppValidateStatusStruct struct {
	// Description of this status item.
	Description string `json:"description,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}
