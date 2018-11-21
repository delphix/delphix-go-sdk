package delphix

import (
	"fmt"
	"reflect"
	"testing"
)

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
		{
			name: "test1",
			fields: fields{
				url:      testSysadminClient.url,
				username: testSysadminClient.username,
				password: testSysadminClient.password,
			},
			want:    "USER-1",
			wantErr: false,
			args: args{
				u: "/user",
				r: "USER-1",
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
			if !reflect.DeepEqual(got.(map[string]interface{})["reference"].(string), tt.want) {
				t.Errorf("Client.FindObjectByReference() = %v, want %v", got.(map[string]interface{})["reference"].(string), tt.want)
			}
		})
	}
}

func TestClient_executeListAndReturnResults(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				url:      testSysadminClient.url,
				username: testSysadminClient.username,
				password: testSysadminClient.password,
			},
			wantErr: false,
			args: args{
				u: "/environment",
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
			got, err := c.executeListAndReturnResults(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.executeListAndReturnResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("GOT: %v", got)
		})
	}
}

func TestClient_executeGetReturnBodyAndCast(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				url:      testSysadminClient.url,
				username: testSysadminClient.username,
				password: testSysadminClient.password,
			},
			wantErr: false,
			args: args{
				u: "/user/USER-1",
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
			got, err := c.executeGetReturnBodyAndCast(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.executeGetReturnBodyAndCast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("GOT: %v", got)
		})
	}
}

func TestClient_executeListAndCastAndReturnResults(t *testing.T) {
	type fields struct {
		url      string
		username string
		password string
	}
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				url:      testSysadminClient.url,
				username: testSysadminClient.username,
				password: testSysadminClient.password,
			},
			wantErr: false,
			args: args{
				u: "/environment",
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
			got, err := c.executeListAndCastAndReturnResults(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.executeListAndReturnResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("GOT: %v", got)
		})
	}
}
