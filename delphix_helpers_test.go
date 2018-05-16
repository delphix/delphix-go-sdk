package delphix

import (
	"fmt"
	"log"
	"testing"
)

func TestWaitForEngineReady(t *testing.T) {
	err := testSysadminClient.WaitForEngineReady(10, 20)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
}

func TestCheckStoragePool(t *testing.T) {
	err := testSysadminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate sysadmin connection:  %s\n", err)
	}
	_, err = testSysadminClient.CheckStoragePool()
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
}

func TestCreateDomain(t *testing.T) {
	err := testSysadminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate sysadmin connection:  %s\n", err)
	}
	_, err = testSysadminClient.CreateDomain()
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
}

func TestFactoryReset(t *testing.T) {
	err := testSysadminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate sysadmin connection:  %s\n", err)
	}
	err = testSysadminClient.FactoryReset()
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
}

func TestGetStorageDevices(t *testing.T) {
	err := testSysadminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate sysadmin connection:  %s\n", err)
	}
	_, err = testSysadminClient.GetStorageDevices()
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
}

func TestCreateDomainJSONBody(t *testing.T) {
	err := testSysadminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate sysadmin connection:  %s\n", err)
	}
	body, err := CreateDomainJSONBody(&testSysadminClient)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}

	fmt.Println(body)
}

func TestUpdateUserPassword(t *testing.T) {
	userName := "sysadmin"
	newPassword := "testing123"
	originalPassword := testSysadminClient.password
	err := testSysadminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate sysadmin connection:  %s\n", err)
	}
	err = testSysadminClient.UpdateUserPasswordByName(userName, newPassword)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	//Change password back
	err = testSysadminClient.UpdateUserPasswordByName(userName, originalPassword)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
}

func TestInitialSetup(t *testing.T) {
	userName := "sysadmin"
	newPassword := testSysadminClient.password
	testSysadminClient.password = "sysadmin"
	err := testSysadminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate sysadmin connection:  %s\n", err)
	}
	err = testSysadminClient.UpdateUserPasswordByName(userName, newPassword)
	if err != nil {
		t.Fatalf("Failed to update sysadmin password: %s\n", err)
	}

	TestCreateDomain(t)

	userName = "delphix_admin"
	newPassword = testDelphixAdminClient.password
	testDelphixAdminClient.password = "delphix"
	err = testDelphixAdminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate delphix_admin connection:  %s\n", err)
	}
	err = testDelphixAdminClient.UpdateUserPasswordByName(userName, newPassword)
	if err != nil {
		t.Fatalf("Failed to update delphix_admin password:  %s\n", err)
	}
}

func TestUpdateDatabase(t *testing.T) {
	oOracleDatabaseContainer := OracleDatabaseContainer{
		Type: "OracleDatabaseContainer",
		Name: testAccVDBName,
	}
	uOracleDatabaseContainer := OracleDatabaseContainer{
		Type: "OracleDatabaseContainer",
		Name: "UpdatedName",
	}

	err := testDelphixAdminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate delphix_admin connection:  %s\n", err)
	}
	ref, err := testDelphixAdminClient.FindDatabaseByName(oOracleDatabaseContainer.Name)
	if err != nil {
		t.Fatalf("Failed trying FindDatabaseByName %s\n", err)
	}
	if ref == nil {
		t.Fatalf("Database \"%s\" not found", oOracleDatabaseContainer.Name)
	}

	log.Printf("Changing %s from \"%s\" to \"%s\"", ref.(string), oOracleDatabaseContainer.Name, uOracleDatabaseContainer.Name)

	err = testDelphixAdminClient.UpdateDatabase(ref.(string), &uOracleDatabaseContainer)
	if err != nil {
		t.Fatalf("Failed to update database: %s\n", err)
	}

	log.Printf("Changing %s from \"%s\" to \"%s\"", ref.(string), uOracleDatabaseContainer.Name, oOracleDatabaseContainer.Name)

	err = testDelphixAdminClient.UpdateDatabase(ref.(string), &oOracleDatabaseContainer)
	if err != nil {
		t.Fatalf("Failed to update database:  %s\n", err)
	}
}

func TestClient_CreateEnvironment(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		h *HostEnvironmentCreateParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "An environment reference",
			wantErr: false,
			args: args{
				h: &HostEnvironmentCreateParameters{
					Type: "HostEnvironmentCreateParameters",
					PrimaryUser: &EnvironmentUser{
						Type: "EnvironmentUser",
						Name: "delphix",
						Credential: &PasswordCredential{
							Type:     "PasswordCredential",
							Password: "delphix",
						},
					},
					HostEnvironment: &UnixHostEnvironment{
						Type: "UnixHostEnvironment",
						Name: "LINUXTARGET",
					},
					HostParameters: &UnixHostCreateParameters{
						Type: "UnixHostCreateParameters",
						Host: &UnixHost{
							Type:        "UnixHost",
							Address:     "172.16.173.131",
							ToolkitPath: "/u01/app/toolkit",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			fmt.Println(c)
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateEnvironment(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateEnvironment() got = %v, want %v", got, tt.want)
			} else {
				log.Printf("Client.CreateEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateEnvironmentWithKey(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		h *HostEnvironmentCreateParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://34.217.117.232/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "An environment reference",
			wantErr: false,
			args: args{
				h: &HostEnvironmentCreateParameters{
					Type: "HostEnvironmentCreateParameters",
					PrimaryUser: &EnvironmentUser{
						Type: "EnvironmentUser",
						Name: "delphix",
						Credential: &SystemKeyCredential{
							Type: "SystemKeyCredential",
						},
					},
					HostEnvironment: &UnixHostEnvironment{
						Type: "UnixHostEnvironment",
						Name: "LINUXTARGET",
					},
					HostParameters: &UnixHostCreateParameters{
						Type: "UnixHostCreateParameters",
						Host: &UnixHost{
							Type:        "UnixHost",
							Address:     "10.0.1.95",
							ToolkitPath: "/u01/app/toolkit",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			fmt.Println(c)
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateEnvironment(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateEnvironment() got = %v, want %v", got, tt.want)
			} else {
				log.Printf("Client.CreateEnvironment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteEnvironment(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		r string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			wantErr: false,
			args: args{
				r: "UNIX_HOST_ENVIRONMENT-6",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			c.LoadAndValidate()
			err := c.DeleteEnvironment(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_CreateEnvironment_And_DeleteEnvironment(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		h *HostEnvironmentCreateParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "An environment reference",
			wantErr: false,
			args: args{
				h: &HostEnvironmentCreateParameters{
					Type: "HostEnvironmentCreateParameters",
					PrimaryUser: &EnvironmentUser{
						Type: "EnvironmentUser",
						Name: "delphix",
						Credential: &PasswordCredential{
							Type:     "PasswordCredential",
							Password: "delphix",
						},
					},
					HostEnvironment: &UnixHostEnvironment{
						Type: "UnixHostEnvironment",
						Name: "LINUXTARGET",
					},
					HostParameters: &UnixHostCreateParameters{
						Type: "UnixHostCreateParameters",
						Host: &UnixHost{
							Type:        "UnixHost",
							Address:     "172.16.173.131",
							ToolkitPath: "/u01/app/toolkit",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			fmt.Println(c)
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateEnvironment(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateEnvironment() got = %v, want %v", got, tt.want)
				return
			} else {
				log.Printf("Client.CreateEnvironment() got = %v, want %v", got, tt.want)
			}
			err = c.DeleteEnvironment(got.(string))
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_UpdateEnvironment(t *testing.T) {
	uEnv := UnixHostEnvironment{
		Type:        "UnixHostEnvironment",
		Name:        "Updated Name",
		Description: "Initial",
	}
	oEnv := UnixHostEnvironment{
		Type:        "UnixHostEnvironment",
		Name:        "LINUXTARGET",
		Description: "Updated",
	}

	err := testDelphixAdminClient.LoadAndValidate()
	if err != nil {
		t.Fatalf("Failure to Load and Validate delphix_admin connection:  %s\n", err)
	}
	envObj, err := testDelphixAdminClient.FindEnvironmentByName(oEnv.Name)
	if err != nil {
		t.Fatalf("Failed trying FindDEnvironmentByName %s\n", err)
	}
	if envObj == nil {
		t.Fatalf("Environment \"%s\" not found", oEnv.Name)
	}

	_, ref, err := GrabObjectNameAndReference(envObj)
	if err != nil {
		t.Fatalf("Failed trying GrabObjectNameAndReference %s\n", err)
	}

	log.Printf("Changing %s from \"%s\" to \"%s\"", ref, oEnv.Name, uEnv.Name)

	err = testDelphixAdminClient.UpdateEnvironment(ref, &uEnv)
	if err != nil {
		t.Fatalf("Failed to update environment: %s\n", err)
	}

	log.Printf("Changing %s from \"%s\" to \"%s\"", ref, uEnv.Name, oEnv.Name)

	err = testDelphixAdminClient.UpdateEnvironment(ref, &oEnv)
	if err != nil {
		t.Fatalf("Failed to update environment:  %s\n", err)
	}
}

func TestClient_CreateDatabase(t *testing.T) {
	f := new(bool)
	*f = false

	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		o *OracleProvisionParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "A VDB reference",
			wantErr: false,
			args: args{
				o: &OracleProvisionParameters{
					Type: "OracleProvisionParameters",
					Container: &OracleDatabaseContainer{
						Type:  "OracleDatabaseContainer",
						Name:  testAccVDBName,
						Group: "GROUP-35",
					},
					Source: &OracleVirtualSource{
						Type:                            "OracleVirtualSource",
						MountBase:                       "/mnt/provision",
						AllowAutoVDBRestartOnHostReboot: f,
					},
					SourceConfig: &OracleSIConfig{
						Type:         "OracleSIConfig",
						Repository:   "ORACLE_INSTALL-38",
						DatabaseName: "unit",
						UniqueName:   "unit",
						Instance: &OracleInstance{
							Type:           "OracleInstance",
							InstanceName:   "unit",
							InstanceNumber: 1,
						},
					},
					TimeflowPointParameters: TimeflowPointSemantic{
						Type:      "TimeflowPointSemantic",
						Container: "ORACLE_DB_CONTAINER-14",
						Location:  "LATEST_SNAPSHOT",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			fmt.Println(c)
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateDatabase(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateDatabase() got = %v, want %v", got, tt.want)
			} else {
				log.Printf("Client.CreateDatabase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateDSource(t *testing.T) {
	fbool := new(bool)
	*fbool = false
	tbool := new(bool)
	*tbool = true
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		l *LinkParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "A dSource reference",
			wantErr: false,
			args: args{
				l: &LinkParameters{
					Type:  "LinkParameters",
					Name:  "Employee Oracle 11G DB",
					Group: "GROUP-1",
					LinkData: &OracleLinkData{
						Type:            "OracleLinkData",
						Config:          "ORACLE_SINGLE_CONFIG-4",
						EnvironmentUser: "HOST_USER-35",
						DbUser:          "delphixdb",
						DbCredentials: &PasswordCredential{
							Type:     "PasswordCredential",
							Password: "delphixdb",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateDSource(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateDSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateDSource() did not return a dSource reference")
				return
			}
		})
	}
}

func TestClient_DeleteDatabase(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			wantErr: false,
			args: args{
				v: "ORACLE_DB_CONTAINER-16",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			if err := c.DeleteDatabase(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateAndDeleteDatabase(t *testing.T) {
	f := new(bool)
	*f = false

	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		o *OracleProvisionParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "A VDB reference",
			wantErr: false,
			args: args{
				o: &OracleProvisionParameters{
					Type: "OracleProvisionParameters",
					Container: &OracleDatabaseContainer{
						Type:  "OracleDatabaseContainer",
						Name:  testAccVDBName,
						Group: "GROUP-35",
					},
					Source: &OracleVirtualSource{
						Type:                            "OracleVirtualSource",
						MountBase:                       "/mnt/provision",
						AllowAutoVDBRestartOnHostReboot: f,
					},
					SourceConfig: &OracleSIConfig{
						Type:         "OracleSIConfig",
						Repository:   "ORACLE_INSTALL-38",
						DatabaseName: "unit",
						UniqueName:   "unit",
						Instance: &OracleInstance{
							Type:           "OracleInstance",
							InstanceName:   "unit",
							InstanceNumber: 1,
						},
					},
					TimeflowPointParameters: TimeflowPointSemantic{
						Type:      "TimeflowPointSemantic",
						Container: "ORACLE_DB_CONTAINER-14",
						Location:  "LATEST_SNAPSHOT",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			fmt.Println(c)
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateDatabase(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateDatabase() got = %v, want %v", got, tt.want)
			} else {
				log.Printf("Client.CreateDatabase() got = %v, want %v", got, tt.want)
			}
			if err := c.DeleteDatabase(got.(string)); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_DeleteDSource(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			wantErr: false,
			args: args{
				v: "ORACLE_DB_CONTAINER-14",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			if err := c.DeleteDSource(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteDSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateAndDeleteDSource(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		l *LinkParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "A dSource reference",
			wantErr: false,
			args: args{
				l: &LinkParameters{
					Type:  "LinkParameters",
					Name:  "Employee Oracle 11G DB",
					Group: "GROUP-1",
					LinkData: &OracleLinkData{
						Type:            "OracleLinkData",
						Config:          "ORACLE_SINGLE_CONFIG-4",
						EnvironmentUser: "HOST_USER-35",
						DbUser:          "delphixdb",
						DbCredentials: &PasswordCredential{
							Type:     "PasswordCredential",
							Password: "delphixdb",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateDSource(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateDSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateDSource() did not return a dSource reference")
				return
			}
			if err := c.DeleteDSource(got.(string)); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteDSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateAndUpdateAndDeleteDatabase(t *testing.T) {
	f := new(bool)
	*f = false

	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		o *OracleProvisionParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "A VDB reference",
			wantErr: false,
			args: args{
				o: &OracleProvisionParameters{
					Type: "OracleProvisionParameters",
					Container: &OracleDatabaseContainer{
						Type:  "OracleDatabaseContainer",
						Name:  testAccVDBName,
						Group: "GROUP-35",
					},
					Source: &OracleVirtualSource{
						Type:                            "OracleVirtualSource",
						MountBase:                       "/mnt/provision",
						AllowAutoVDBRestartOnHostReboot: f,
					},
					SourceConfig: &OracleSIConfig{
						Type:         "OracleSIConfig",
						Repository:   "ORACLE_INSTALL-39",
						DatabaseName: "unit",
						UniqueName:   "unit",
						Instance: &OracleInstance{
							Type:           "OracleInstance",
							InstanceName:   "unit",
							InstanceNumber: 1,
						},
					},
					TimeflowPointParameters: TimeflowPointSemantic{
						Type:      "TimeflowPointSemantic",
						Container: "ORACLE_DB_CONTAINER-20",
						Location:  "LATEST_SNAPSHOT",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uOracleDatabaseContainer := OracleDatabaseContainer{
				Type: "OracleDatabaseContainer",
				Name: "UpdatedName",
			}
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			fmt.Println(c)
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateDatabase(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateDatabase() got = %v, want %v", got, tt.want)
			} else {
				log.Printf("Client.CreateDatabase() got = %v, want %v", got, tt.want)
			}
			log.Printf("Changing %s from \"%s\" to \"%s\"", got.(string), testAccVDBName, uOracleDatabaseContainer.Name)

			err = testDelphixAdminClient.UpdateDatabase(got.(string), &uOracleDatabaseContainer)
			if err != nil {
				t.Fatalf("Failed to update database: %s\n", err)
			}
			if err := c.DeleteDatabase(got.(string)); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateAndSyncAndUpdateAndDeleteDSource(t *testing.T) {
	fbool := new(bool)
	*fbool = false
	tbool := new(bool)
	*tbool = true
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		l *LinkParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			want:    "A dSource reference",
			wantErr: false,
			args: args{
				l: &LinkParameters{
					Type:  "LinkParameters",
					Name:  "Employee Oracle 11G DB",
					Group: "GROUP-1",
					LinkData: &OracleLinkData{
						Type:            "OracleLinkData",
						Config:          "ORACLE_SINGLE_CONFIG-4",
						EnvironmentUser: "HOST_USER-35",
						DbUser:          "delphixdb",
						DbCredentials: &PasswordCredential{
							Type:     "PasswordCredential",
							Password: "delphixdb",
						},
						LinkNow: tbool,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uOracleDatabaseContainer := OracleDatabaseContainer{
				Type: "OracleDatabaseContainer",
				Name: "UpdatedName",
			}
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			got, err := c.CreateDSource(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateDSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.CreateDSource() did not return a dSource reference")
				return
			}
			if err = testDelphixAdminClient.UpdateDSource(got.(string), &uOracleDatabaseContainer); (err != nil) != tt.wantErr {
				t.Fatalf("Failed to update database: %s\n", err)
			}
			if err := c.DeleteDSource(got.(string)); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteDSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SyncDatabase(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		r string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				url:      "http://de/resources/json/delphix",
				username: "delphix_admin",
				password: "landshark",
			},
			wantErr: false,
			args: args{
				r: "ORACLE_DB_CONTAINER-30",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:      tt.fields.url,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			err := c.LoadAndValidate()
			if err != nil {
				t.Errorf("Cannot get session: %v", err)
				return
			}
			if err := c.SyncDatabase(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Client.SyncDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
