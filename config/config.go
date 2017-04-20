package config

import (
    "encoding/json"
    "os"
)

type Option struct {
    Name string
    Argument string
    List []string
}

func (p *Option) init() {
    p.List = make([]string, 1)
}

func Load(path string) {
}
