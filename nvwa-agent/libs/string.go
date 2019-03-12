package libs

import (
    "crypto/md5"
    "encoding/hex"
    "io"
    "os"
)

func Md5File(filePath string) (string, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer f.Close()

    md5hash := md5.New()
    if _, err := io.Copy(md5hash, f); err != nil {
        return "", err
    }

    return hex.EncodeToString(md5hash.Sum(nil)), nil
}

func Hostname() string {
    name, err := os.Hostname()
    if err != nil {
        return "UNKNOWN"
    }

    return name
}
