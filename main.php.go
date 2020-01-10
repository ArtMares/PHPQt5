package main

import (
    "bufio"
    io2 "io"
    "os"
    "path/filepath"
)

func NewPHPFile(dir, name string) *PHPFile {
    return &PHPFile{
        dir:  dir,
        name: name,
    }
}

type PHPFile struct {
    dir     string
    name    string
    ext     string
    rows    []string
    io      io2.ReadWriter
}

func (f *PHPFile) FullPath() string {
    return filepath.Join(f.dir, f.name)
}

func (f *PHPFile) CountRows() int {
    return len(f.rows)
}

func (f *PHPFile) Exists() bool {
    _, err := os.Stat(f.FullPath())
    if !os.IsNotExist(err) {
        return true
    }
    return false
}

func (f *PHPFile) Open() error {
    if f.io == nil {
        file, err := os.Open(f.FullPath())
        if err != nil {
            return err
        }
        f.io = file
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            f.rows = append(f.rows, scanner.Text())
        }
        if err = scanner.Err(); err != nil {
            return err
        }
    }
    return nil
}

func (f *PHPFile) Create() error {

}