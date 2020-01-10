package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type PHPQt5Project struct {
    AddOns      PHPQt5AddOns    `json:"addons"`
    App         PHPQt5App       `json:"app"`
    Build       PHPQt5Build     `json:"build"`
    Config      PHPQt5Config    `json:"config"`
    Debug       PHPQt5Debug     `json:"debug"`
    Includes    PHPQt5Includes  `json:"includes"`
    Version     PHPQt5Version   `json:"version"`
}

func (p *PHPQt5Project) Load(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    return json.NewDecoder(file).Decode(p)
}

func (p *PHPQt5Project) Save(filePath string) error {
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
    if err != nil {
        return err
    }
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ")
    return encoder.Encode(p)
}

type PHPQt5AddOns struct {
    list        []string
}

func (m *PHPQt5AddOns) Add(name string) {
    m.list = append(m.list, name)
}

func (m *PHPQt5AddOns) Exist(name string) bool {
    for _, addon := range m.list {
        if addon == name {
            return true
        }
    }
    return false
}

func (m *PHPQt5AddOns) Del(name string) {
    n := -1
    for i, addon := range m.list {
        if addon == name {
            n = i
            break
        }
    }
    if n > -1 {
        m.list = append(m.list[n:], m.list[n+1:]...)
    }
}

func (m PHPQt5AddOns) MarshalJSON() ([]byte, error) {
    if m.list == nil {
        m.list = make([]string, 0)
    }
    return json.MarshalIndent(m.list, "", "    ")
}

func (m *PHPQt5AddOns) UnmarshalJSON(b []byte) error {
    return json.Unmarshal(b, &m.list)
}

type PHPQt5App struct {
    Name        string  `json:"name"`
    OrgDomain   string  `json:"orgDomain"`
    OrgName     string  `json:"orgName"`
    Version     Version `json:"version"`
}

type PHPQt5Build struct {
    Icon        string  `json:"icon"`
    QMakelf     string  `json:"qmakelf"`
    IniMd5      bool    `json:"save_ini_md5"`
    Php7TsMd5   bool    `json:"save_php7ts_md5"`
    Template    string  `json:"template"`
}

type PHPQt5Config struct {
    list        []string
}

func (m PHPQt5Config) MarshalJSON() ([]byte, error) {
    if m.list == nil {
        m.list = make([]string, 0)
    }
    return json.MarshalIndent(m.list, "", "    ")
}

func (m *PHPQt5Config) UnmarshalJSON(b []byte) error {
    return json.Unmarshal(b, &m.list)
}

type PHPQt5Debug struct {
    SocketName  string  `json:"socketName"`
}

type PHPQt5Includes struct {
    list        []string
}

func (m PHPQt5Includes) MarshalJSON() ([]byte, error) {
    if m.list == nil {
        m.list = make([]string, 0)
    }
    return json.MarshalIndent(m.list, "", "    ")
}

func (m *PHPQt5Includes) UnmarshalJSON(b []byte) error {
    return json.Unmarshal(b, &m.list)
}

type PHPQt5Version Version

func (m PHPQt5Version) String() string {
    return fmt.Sprintf("%d.%d", m[0], m[1])
}

func (m PHPQt5Version) Bytes() []byte {
    return []byte(m.String())
}

func (m PHPQt5Version) MarshalText() ([]byte, error) {
    return m.Bytes(), nil
}

func (m PHPQt5Version) UnmarshalText(b []byte) error {
    p := strings.SplitN(string(b), ".", 2)
    if len(p) >= 2 {
        m[0], _ = strconv.ParseUint(p[0], 10, 64)
        m[1], _ = strconv.ParseUint(p[1], 10, 64)
    }
    return nil
}

func (m PHPQt5Version) EQ(o Comparison) bool {
    switch v := o.(type) {
    case Version:
        return m[0] == v[0] && m[1] == v[1]
    case PHPQt5Version:
        return m[0] == v[0] && m[1] == v[1]
    default:
        return false
    }
}

func (m PHPQt5Version) LS(o Comparison) bool {
    switch v := o.(type) {
    case Version:
        switch {
        case m[0] < v[0]:
            return true
        case m[0] == v[0] && m[1] < v[1]:
            return true
        default:
            return false
        }
    case PHPQt5Version:
        switch {
        case m[0] < v[0]:
            return true
        case m[0] == v[0] && m[1] < v[1]:
            return true
        default:
            return false
        }
    default:
        return false
    }
}

func (m PHPQt5Version) GR(o Comparison) bool {
    if o != nil {
        return !m.EQ(o) || !m.LS(o)
    }
    return false
}

func (m PHPQt5Version) LE(o Comparison) bool {
    if o != nil {
        return m.EQ(o) || m.LS(o)
    }
    return false
}

func (m PHPQt5Version) GE(o Comparison) bool {
    if o != nil {
        return m.EQ(o) || m.GR(o)
    }
    return false
}