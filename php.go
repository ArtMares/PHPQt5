package main

var phpAutoloader = PHPFile{
    name: "autoload.php",
    ext: ".php",
    rows: []string{
        `<?php`,
        `spl_autoload_register(function($className) {`,
        `    $file = str_replace('\\', '/', $className).'.php';`,
        `    $qf = new QFile();`,
        `    $qf->setFileName(':/scripts/'.$file);`,
        `    if($qf->exists()) {`,
        `        require_once('qrc://scripts/'.$file);`,
        `        return true;`,
        `    } else {`,
        `        $qf->setFileName($file);`,
        `        if($qf->exists()) {`,
        `            require_once($file);`,
        `            return true;`,
        `        }`,
        `    }`,
        `    return false;`,
        `});`,
    },
}