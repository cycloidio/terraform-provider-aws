package account_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/acctest"
)

func TestAccAccount_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"AlternateContact": {
			"basic":      testAccAlternateContact_basic,
			"disappears": testAccAlternateContact_disappears,
			"AccountID":  testAccAlternateContact_accountID,
		},
		"PrimaryContact": {
			"basic": testAccPrimaryContact_basic,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
