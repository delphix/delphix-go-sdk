package delphix

func CreateOracleSyncParameters(
	DoNotResume *bool,
	DoubleSync *bool,
	FilesForFullBackup []*int,
	ForceFullBackup *bool,
	SkipSpaceCheck *bool,
) OracleSyncFromExternalParametersStruct {
	return OracleSyncFromExternalParametersFactory(
		DoNotResume,
		DoubleSync,
		FilesForFullBackup,
		ForceFullBackup,
		SkipSpaceCheck,
	)
}
