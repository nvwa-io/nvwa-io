package ansible

import (
    "log"
    "testing"
)

var (
    test_hosts = []string{"192.168.10.10"}
)

func Test_ExecShell(t *testing.T) {
    o, cmd, err := C().ExecShell("root", "echo 'Hello world';", test_hosts, 30)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println(cmd)
    log.Println(string(o))
}

func Test_CopyFile(t *testing.T) {
    o, cmd, err := C().CopyFile("root", "/tmp/a.log", "/tmp/b.log", test_hosts)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println(cmd)
    log.Println(string(o))
}

func Test_Version(t *testing.T) {
    o, err := C().Version()
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println(string(o))
}
