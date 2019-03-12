package jenkins

import (
    "log"
    "testing"
)

func Test_Csrf(t *testing.T) {
    token, err := C().Config(test_domain, "admin", "admin").CrumbIssuer()
    if err != nil {
        t.Fatal(err.Error())
    }

    log.Println(token)
}
