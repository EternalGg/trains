package subdomainVisits

import "testing"

func TestSubdomainVisits(t *testing.T) {
	SubdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"})
}
