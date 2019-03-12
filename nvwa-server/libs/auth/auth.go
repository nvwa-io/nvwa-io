package auth

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs/encrypt"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
    "github.com/astaxie/beego/context"
    "strings"
    "time"
)

const (
    UCINFO_SALT = "@**nvwa-io**@"
    KEY_TOKEN   = "Nvwa-Token"
    EXPIRE      = 7 * 24 * 3600
)

// tips:
// data = json.Marshal(UCInfo)
// sign = md5(data + salt)
// Nvwa-Token = 3des(data + SALT + sign)
type UCInfo struct {
    Uid   int64 `json:"uid"`
    Ctime int64 `json:"ctime"`
}

type UserAuth struct{}

func NewUserAuth() *UserAuth {
    return new(UserAuth)
}

// 解析 Header 中的 Token
func (t *UserAuth) ParseUserTokenFromHeader(ctx *context.Context) (*UCInfo, error) {
    token := ctx.Request.Header.Get(KEY_TOKEN)
    if token == "" {
        logger.Errorf("Empty token.")
        return nil, errors.New("Empth Token")
    }

    ucInfo, err := t.UCInfoDecrypt(token)
    if err != nil {
        logger.Errorf("Decode %s=%s failed: %s", KEY_TOKEN, token, err.Error())
        return nil, err
    }

    if ucInfo.Ctime < (time.Now().Unix() - EXPIRE) {
        return nil, errors.New("Token expires.")
    }

    return ucInfo, nil
}

//func (t *UserAuth) SetUserTokenCookie(ctx *context.Context, uid int64) error {
//    data, err := t.UCInfoEncrypt(&UCInfo{Uid: uid, Ctime: time.Now().Unix()})
//    if err != nil {
//        logger.Errorf("Failed to set cookie, err: %s", err.Error())
//        return err
//    }
//
//    // others are ordered as cookie's max age time, path,domain, secure and httponly.
//    ctx.SetCookie(KEY_TOKEN, data, EXPIRE, "/", "localhost")
//    return nil
//}

func (t *UserAuth) UCInfoEncryptByUid(uid int64) (string, int64, error) {
    ctime := time.Now().Unix()
    token, err := t.UCInfoEncrypt(&UCInfo{
        Uid:   uid,
        Ctime: ctime,
    })
    if err != nil {
        return "", 0, err
    }
    return token, ctime + EXPIRE, nil
}

func (t *UserAuth) UCInfoEncrypt(ucInfo *UCInfo) (string, error) {
    b, err := json.Marshal(ucInfo)
    if err != nil {
        return "", err
    }

    data := fmt.Sprintf("%s%s", string(b), UCINFO_SALT)
    sign := libs.Md5Str(data)

    encData, err := encrypt.TripleDesBase64EncryptStr(data + sign)
    if err != nil {
        return "", err
    }

    return encData, nil
}

func (t *UserAuth) UCInfoDecrypt(data string) (*UCInfo, error) {
    decData, err := encrypt.TripleDesBase64DecryptStr(data)
    if err != nil {
        return nil, err
    }

    arr := strings.Split(decData, UCINFO_SALT)
    if len(arr) != 2 {
        return nil, errs.ERR_INVALID_TOKEN
    }

    tmpSign := libs.Md5Str(arr[0] + UCINFO_SALT)
    if tmpSign != arr[1] {
        return nil, errs.ERR_INVALID_TOKEN
    }

    ucInfo := new(UCInfo)
    err = json.Unmarshal([]byte(arr[0]), ucInfo)
    if err != nil {
        return nil, err
    }

    return ucInfo, nil
}
