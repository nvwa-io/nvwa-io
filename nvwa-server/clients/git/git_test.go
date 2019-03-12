package git

import (
    "log"
    "testing"
)

const (
    test_url      = "git@code.aliyun.com:542792857/demo-01.git"
    test_username = ""
    test_password = ""
    test_token    = "***"
    test_path     = "/tmp/demo-01"
)

func Test_Clone(t *testing.T) {
    //err := C().BasicAuth(test_username, test_password).Clone(test_url, test_path)
    err := C().Clone(test_url, test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }
}

func Test_Pull(t *testing.T) {
    //err := C().BasicAuth(test_username, test_password).Pull(test_path)
    err := C().Pull(test_path, "master")
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }
}

func Test_Fetch(t *testing.T) {
    //err := C().BasicAuth(test_username, test_password).Pull(test_path)
    err := C().FetchAll(test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }
}

func Test_CheckoutBranch(t *testing.T) {
    b := "f-05"
    err := C().FetchAll(test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }

    //err := C().BasicAuth(test_username, test_password).Checkout(test_path, "master")
    err = C().CheckoutBranch(test_path, b)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }
}

func Test_CheckoutTag(t *testing.T) {
    tag := "v04-03"
    err := C().FetchAll(test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }

    //err := C().BasicAuth(test_username, test_password).Checkout(test_path, "master")
    err = C().CheckoutTag(test_path, tag)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }
}

func Test_Branches(t *testing.T) {
    //remotes, err := C().BasicAuth(test_username, test_password).LsRemoteBranches(test_path)
    remotes, _, err := C().LsRemoteBranches(test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }

    log.Println(remotes)

    //branches, err := C().BasicAuth(test_username, test_password).AllBranches(test_path)
    branches, _, err := C().AllBranches(test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }

    log.Println(branches)
}

func Test_Tags(t *testing.T) {
    err := C().FetchAll(test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }

    //tags, tagsCommits, err := C().BasicAuth(test_username, test_password).AllTags(test_path)
    tags, tagsCommits, err := C().AllTags(test_path)
    if err != nil {
        log.Fatalf("err=%s", err.Error())
    }

    log.Println(tags)
    log.Println(tagsCommits)
}
