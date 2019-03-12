package libs

import (
    "bytes"
    "errors"
    "fmt"
    "log"
    "os/exec"
    "strings"
    "time"
)

func CmdExecShell(shell string, timeout int) ([]byte, error) {
    shell = fmt.Sprintf(". /etc/profile; %s", shell)
    log.Println(EscapeShellArg(shell))
    return CmdExec("/bin/sh", timeout, "-c", shell)
}

func CmdExecShellDefault(shell string) ([]byte, error) {
    return CmdExecShell(shell, 3600)
}

// pack version package
func TarPackage(workspace, pkgPath string, files []string, excludes []string) ([]byte, string, error) {
    // e.g cmd := "tar -p
    // --exclude='.git'
    // --exclude='.svn'
    // -cz -f '/data/nvwa/packages/demo-01/demo-01.buildId.branch.commitId.datetime.tar.gz'
    // file01 file02 dir01 dir02"

    strExclude := ""
    if len(excludes) > 0 {
        for _, v := range excludes {
            if v == "" {
                continue
            }
            strExclude += fmt.Sprintf(" --exclude='%s' ", v)
        }
    }

    strFiles := ""
    if len(files) > 0 {
        for _, v := range files {
            if v == "" {
                continue
            }
            strFiles += fmt.Sprintf(" %s ", v)
        }
    }

    if strFiles == "" {
        strFiles = "*"
    }

    cmd := fmt.Sprintf("cd %s && tar -p %s -cz -f '%s' %s", workspace, strExclude, pkgPath, strFiles)
    output, err := CmdExecShellDefault(cmd)
    return output, cmd, err
}

// execute command with timeout
func CmdExec(command string, timeout int, args ...string) ([]byte, error) {
    // instantiate new command
    cmd := exec.Command(command, args...)

    // get pipe to standard output
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return nil, err
    }

    stderr, err := cmd.StderrPipe()
    if err != nil {
        return nil, err
    }

    // start process via command
    if err := cmd.Start(); err != nil {
        return nil, err
    }

    // setup a buffer to capture standard output
    var buf bytes.Buffer
    var errBuf bytes.Buffer

    // create a channel to capture any errors from wait
    done := make(chan error)
    go func() {
        if _, err := buf.ReadFrom(stdout); err != nil {
            done <- err
        }

        if _, err := errBuf.ReadFrom(stderr); err != nil {
            done <- err
        }
        done <- cmd.Wait()
    }()

    // block on select, and switch based on actions received
    select {
    case <-time.After(time.Duration(timeout) * time.Second):
        if err := cmd.Process.Kill(); err != nil {
            return nil, err
        }
        return nil, errors.New("cmd exec timeout, process killed")

    case err := <-done:
        if err != nil {
            close(done)
            return nil, errors.New(fmt.Sprintf(
                "process done, with error: %s, stdout=%s, stderr=%s",
                err.Error(),
                buf.String(),
                errBuf.String()))
        }

        return buf.Bytes(), nil
    }

    return buf.Bytes(), nil
}

// EscapeShellArg() adds single quotes around a string and quotes/escapes
// any existing single quotes allowing you to pass a string directly to
// a shell function and having it be treated as a single safe argument.
func EscapeShellArg(arg string) string {
    return "'" + strings.Replace(arg, "'", `'\''`, -1) + "'"
}
