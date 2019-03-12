package entities

type BaseEntity struct {
    Id      int64  `json:"id"`
    Enabled int    `json:"-"`
    Ctime   string `json:"ctime"`
    Utime   string `json:"utime"`
}
