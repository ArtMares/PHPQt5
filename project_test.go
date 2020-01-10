package main

import (
    "path/filepath"
    "testing"
)

var testDir = "test"

func TestPHPQt5Project_Load(t *testing.T) {
    p := &PHPQt5Project{}
    err := p.Load(filepath.Join(testDir, "load.phpqt5"))
    if err != nil {
        t.Error(err)
    }
}

func TestPHPQt5Project_Save(t *testing.T) {
    p := &PHPQt5Project{
        AddOns:   PHPQt5AddOns{},
        App:      PHPQt5App{
            Name:      "PQDebugger",
            OrgDomain: "phpqt.ru",
            OrgName:   "PHPQt Team",
            Version:   Version{},
        },
        Build:    PHPQt5Build{},
        Config:   PHPQt5Config{},
        Debug:    PHPQt5Debug{},
        Includes: PHPQt5Includes{},
        Version:  PHPQt5Version{1, 0},
    }
    err := p.Save(filepath.Join(testDir, "save.phpqt5"))
    if err != nil {
        t.Error(err)
    }
}