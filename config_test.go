package delphix

import (
	"fmt"
	"os"
	"testing"
)

var (
	testAccDelphixAdminClient Client
	testAccVDBName            string
	//This Client has a slightly different URL format for the non-ACC tests
	testDelphixAdminClient Client
	testSysadminClient     Client
)

func init() {
	testAccDelphixAdminClient.url = os.Getenv("TEST_DELPHIX_URL")
	testAccDelphixAdminClient.username = os.Getenv("TEST_DELPHIX_USERNAME")
	testAccDelphixAdminClient.password = os.Getenv("TEST_DELPHIX_PASSWORD")
	testAccVDBName = os.Getenv("TEST_VDBNAME")
	//This url needs the /resources/json/delphix append
	testDelphixAdminClient.url = testAccDelphixAdminClient.url + "/resources/json/delphix"
	testDelphixAdminClient.username = testAccDelphixAdminClient.username
	testDelphixAdminClient.password = testAccDelphixAdminClient.password
	testSysadminClient.url = testDelphixAdminClient.url
	testSysadminClient.username = "sysadmin"
	testSysadminClient.password = os.Getenv("TEST_SYSADMIN_PASSWORD")
}

func TestClientLoadAndValidate(t *testing.T) {
	fmt.Println("Client: ", testDelphixAdminClient)
	err := testDelphixAdminClient.LoadAndValidate()
	if err != nil {
		t.Errorf("ERROR: %s", err)
	}
}
