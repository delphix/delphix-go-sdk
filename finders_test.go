package delphix

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestFindSourceByName(t *testing.T) {
	sourceName := "Employee DB - Oracle"
	testDelphixAdminClient.LoadAndValidate()
	r, err := testDelphixAdminClient.FindSourceByName(sourceName)
	if err != nil {
		t.Fatalf("Wamp, Wamp: %s\n", err)
	}
	if r == nil {
		log.Fatalf("Source %s not found\n", sourceName)
	}
	fmt.Printf("Source %s: %s\n", sourceName, r)
}

func TestFindDatabaseByRef(t *testing.T) {
	dbRef := "ORACLE_DB_CONTAINER-34"
	testDelphixAdminClient.LoadAndValidate()
	r, err := testDelphixAdminClient.FindDatabaseByReference(dbRef)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if r == nil {
		log.Fatalf("Database %s not found\n", dbRef)
	}
	fmt.Printf("Database %s: %s\n", dbRef, r)
}

func TestFindDatabaseByName(t *testing.T) {
	dbName := "Employee DB - Oracle"
	testDelphixAdminClient.LoadAndValidate()
	r, err := testDelphixAdminClient.FindDatabaseByName(dbName)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if r == nil {
		log.Fatalf("Database %s not found\n", dbName)
	}
	fmt.Printf("Database %s: %s\n", dbName, r)
}

func TestFindGroupByName(t *testing.T) {
	groupName := "1 - Dev Copies"
	testDelphixAdminClient.LoadAndValidate()
	r, err := testDelphixAdminClient.FindGroupByName(groupName)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if r == nil {
		log.Fatalf("Group %s not found\n", groupName)
	}
	fmt.Printf("Group %s: %s\n", groupName, r)
}

func TestFindDatabaseByNameAndFindRepoByEnvironmentRefAndOracleHome(t *testing.T) {
	environmentName := "OELTARGET"
	oracleHome := "/u01/app/oracle/product/11.2.0.4/db_1"
	testDelphixAdminClient.LoadAndValidate()
	environmentRef, err := testDelphixAdminClient.FindEnvironmentByName(environmentName)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if environmentRef == nil {
		log.Fatalf("Environment %s not found\n", environmentName)
	}
	r, err := testDelphixAdminClient.FindRepoReferenceByEnvironmentRefAndOracleHome(environmentRef.(string), oracleHome)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if r == nil {
		log.Fatalf("Repo %s not found\n", oracleHome)
	}
	fmt.Printf("Repo %s%s: %s\n", environmentName, oracleHome, r)
}

func TestFindEnvironmentByName(t *testing.T) {
	environmentName := "OELTARGET"
	testDelphixAdminClient.LoadAndValidate()
	r, err := testDelphixAdminClient.FindEnvironmentByName(environmentName)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if r == nil {
		log.Fatalf("Environment %s not found\n", environmentName)
	}
	_, reference, err := GrabObjectNameAndReference(r) //r.(map[string]interface{})["reference"].(string)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if reference == "" {
		t.Errorf("Reference was empty")
	}
	fmt.Printf("Environment %s: %s\n", environmentName, r)
}

func TestFindEnvironmentByReference(t *testing.T) {
	// environmentName := "LINUXSOURCE"
	testDelphixAdminClient.LoadAndValidate()
	// r, err := testDelphixAdminClient.FindEnvironmentByName(environmentName)
	// if err != nil {
	// 	t.Errorf("Wamp, Wamp: %s\n", err)
	// }
	// if r == nil {
	// 	log.Fatalf("Environment %s not found\n", environmentName)
	// }
	// _, reference, err := GrabObjectNameAndReference(r) //r.(map[string]interface{})["reference"].(string)
	// if err != nil {
	// 	t.Errorf("Wamp, Wamp: %s\n", err)
	// }
	// if reference == "" {
	// 	t.Errorf("Reference was empty")
	// }
	reference := "UNIX_HOST_ENVIRONMENT-4"
	r, err := testDelphixAdminClient.FindEnvironmentByReference(reference)
	if err != nil {
		if err.Error() != "exception.executor.object.missing" {
			log.Fatalf("Wamp, Wamp: %s\n", err)
		}
	}
	if r == nil {
		log.Fatalf("Environment %s not found\n", reference)
	}
	_, reference, err = GrabObjectNameAndReference(r)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	fmt.Printf("Environment %s: %s\n", "test", reference)
}
func TestFindUserByName(t *testing.T) {
	userName := "sysadmin"
	testSysadminClient.LoadAndValidate()
	r, err := testSysadminClient.FindUserByName(userName)
	if err != nil {
		t.Errorf("Wamp, Wamp: %s\n", err)
	}
	if r == nil {
		log.Fatalf("User %s not found\n", userName)
	}
	fmt.Printf("Group %s: %s\n", userName, r)
}

func TestClient_FindSourceConfigByNameAndRepoReference(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		n string
		r string
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
			want:    "ORACLE_SINGLE_CONFIG-4",
			wantErr: false,
			args: args{
				n: "orcl",
				r: "ORACLE_INSTALL-36",
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
			got, err := c.FindSourceConfigByNameAndRepoReference(tt.args.n, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindSourceConfigByNameAndRepoReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.FindSourceConfigByNameAndRepoReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FindObjectByName(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		u string
		n string
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
			want:    "An Environment Object",
			wantErr: false,
			args: args{
				u: "/environment",
				n: "LINUXSOURCE",
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
			got, err := c.FindObjectByName(tt.args.u, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindObjectByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.FindObjectByName() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("Client.FindObjectByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FindObjectByReference(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		u string
		r string
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
			want:    "An Environment Object",
			wantErr: false,
			args: args{
				u: "/environment",
				r: "UNIX_HOST_ENVIRONMENT-51",
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
			got, err := c.FindObjectByReference(tt.args.u, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindObjectByReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.FindObjectByReference() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("Client.FindObjectByReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FindEnvironmentUserByNameAndEnvironmentReference(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		n string
		r string
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
			want:    "A User Reference",
			wantErr: false,
			args: args{
				n: "delphix",
				r: "UNIX_HOST_ENVIRONMENT-51",
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
			got, err := c.FindEnvironmentUserByNameAndEnvironmentReference(tt.args.n, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindEnvironmentUserByNameAndEnvironmentReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Client.FindEnvironmentUserByNameAndEnvironmentReference() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("Client.FindEnvironmentUserByNameAndEnvironmentReference() = %v, want %v", got, tt.want)
			}
		})
	}
}
