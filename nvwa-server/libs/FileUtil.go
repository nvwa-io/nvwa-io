package libs

import (
    "errors"
    "os"
    "path/filepath"
)

func CmdCurrentPath() (string, error) {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        return "", errors.New("无法找到当前命令的位置")
    }
    return dir, nil
}

// write content to file
// if file not exists, create first.
// if file exists, use it.
func FilePutContent(path string, content string) error {
    file, err := os.Open(path)
    if err != nil && os.IsNotExist(err) {
        file, err = os.Create(path)
        if err != nil {
            return err
        }
    }
    defer file.Close()

    _, err = file.Write([]byte(content))
    if err != nil {
        return err
    }

    return nil
}
