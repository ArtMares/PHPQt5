package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

var ext = ".phpqt5"

var flags = Directives{
    "--dir": true,
}

func main() {
    log.SetPrefix("PHPQt5: ")
    wd, ok := flags.Collect(os.Args, "--dir")
    if !ok {
        wd, _ = os.Getwd()
    }
    phpAutoloader.dir = filepath.Join(wd, "scripts")
    fmt.Println(wd)
    var project *PHPQt5Project
    files, err := ioutil.ReadDir(wd)
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }
    for _, file := range files {
        if file.Mode().IsRegular() {
            if filepath.Ext(filepath.Join(wd, file.Name())) == ext {
                project = &PHPQt5Project{}
                err = project.Load(filepath.Join(wd, file.Name()))
                if err != nil {
                    log.Println(err)
                    os.Exit(1)
                }
            }
        }
    }
    if project == nil {
        log.Println("File *.phpqt5 not found")
        os.Exit(1)
    }
    checkAutoLoad(wd)
    fmt.Printf("%#v\n", project)
    if phpAutoloader.Exists() {
        err = phpAutoloader.Open()
        if err != nil {
            log.Println(err)
        }
    }
}

func checkAutoLoad(dir string) {
    mainphp := NewPHPFile(filepath.Join(dir, "scripts"), "main.php")
    if mainphp.Exists() {
        mainphp.Open()
        log.Printf("%d\n", mainphp.CountRows())
    }
}