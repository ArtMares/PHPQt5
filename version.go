package main

import (
    "fmt"
    "strconv"
    "strings"
)

var NilVersion = Version{0,0,0,0}

type Version [4]uint64

func (m Version) Major() uint64 {
    return m[0]
}

func (m Version) Minor() uint64 {
    return m[1]
}

func (m Version) Patch() uint64 {
    return m[2]
}

func (m Version) Build() uint64 {
    return m[3]
}

func (m Version) String() string {
    return fmt.Sprintf("%d.%d.%d", m[0], m[1], m[2])
}

func (m Version) Bytes() []byte {
    return []byte(m.String())
}

func (m Version) MarshalText() ([]byte, error) {
    if m.EQ(NilVersion) {
        return []byte(""), nil
    }
    return m.Bytes(), nil
}

func (m Version) UnmarshalText(b []byte) error {
    p := strings.SplitN(string(b), ".", 4)
    n := len(p)
    if n >= 2 {
        switch n {
        case 2:
            m[0], _ = strconv.ParseUint(p[0], 10, 64)
            m[1], _ = strconv.ParseUint(p[1], 10, 64)
        case 3:
            m[0], _ = strconv.ParseUint(p[0], 10, 64)
            m[1], _ = strconv.ParseUint(p[1], 10, 64)
            m[2], _ = strconv.ParseUint(p[2], 10, 64)
        case 4:
            m[0], _ = strconv.ParseUint(p[0], 10, 64)
            m[1], _ = strconv.ParseUint(p[1], 10, 64)
            m[2], _ = strconv.ParseUint(p[2], 10, 64)
            m[3], _ = strconv.ParseUint(p[3], 10, 64)
        }
    }
    return nil
}

func (m Version) EQ(o Comparison) bool {
    switch v := o.(type) {
    case Version:
        return m[0] == v[0] && m[1] == v[1] && m[2] == v[2] && m[3] == v[3]
    case PHPQt5Version:
        return m[0] == v[0] && m[1] == v[1]
    default:
        return false
    }
}

func (m Version) LS(o Comparison) bool {
    switch v := o.(type) {
    case Version:
        switch {
        case m[0] < v[0]:
            return true
        case m[0] == v[0] && m[1] < v[1]:
            return true
        case m[0] == v[0] && m[1] == v[1] && m[2] < v[2]:
            return true
        case m[0] == v[0] && m[1] == v[1] && m[2] == v[2] && m[3] < v[3]:
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

func (m Version) GR(o Comparison) bool {
    if o != nil {
        return !m.EQ(o) || !m.LS(o)
    }
    return false
}

func (m Version) LE(o Comparison) bool {
    if o != nil {
        return m.EQ(o) || m.LS(o)
    }
    return false
}

func (m Version) GE(o Comparison) bool {
    if o != nil {
        return m.EQ(o) || m.GR(o)
    }
    return false
}