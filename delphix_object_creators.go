package delphix

func CreateOracleSyncParameters(
	DoNotResume *bool,
	DoubleSync *bool,
	ForceFullBackup *bool,
	SkipSpaceCheck *bool,
) OracleSyncParametersStruct {
	return OracleSyncParametersFactory(
		DoNotResume,
		DoubleSync,
		ForceFullBackup,
		SkipSpaceCheck,
	)
}
