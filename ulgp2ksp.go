package ulgp2ksp

import (
    "os"
    "fmt"
    "log"
    "strings"
//    "strconv"
    "io/ioutil"
    "path/filepath"

    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
    "github.com/influxdata/telegraf/plugins/inputs/system"
)


type ULGP2KSPStats struct {
    ps system.PS
}


func (_ *ULGP2KSPStats) Description() string {
    return "Reports P2KSP ULG version tag"
}


func (_ *ULGP2KSPStats) SampleConfig() string { return "" }


func (s *ULGP2KSPStats) Gather(acc telegraf.Accumulator) error {
    instances := []string{}
    dirpath   := "/p2ksp"
    ulgpath   := "bin/componentes/ultimaULG.txt"
    regexp    := "sp_lj"
    f, err    := os.Open(dirpath)

    if err != nil {
        log.Fatal(err)
    }

    files, err := f.Readdir(-1)
    f.Close()

    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if strings.Contains(file.Name(), regexp) {
            werkextract := strings.Replace(file.Name(), "sp_lj", "", -1)
	    fullpath    := filepath.Join(dirpath, file.Name(), ulgpath)
            instances    = append(instances, file.Name(),
	                          filepath.FromSlash(fullpath), werkextract)
        }
    }

    ulgfile      := fmt.Sprint(instances[1])
    werkstr      := fmt.Sprint(instances[2])
//    werkint, err := strconv.Atoi(werkstr)

    content, err := ioutil.ReadFile(ulgfile)
    if err != nil {
        log.Fatal(err)
    }

    fileage, err := os.Stat(ulgfile)
    if err != nil {
	log.Fatal(err)
    }

    age  := fileage.ModTime()
    lote := strings.Replace(string(content), "\n","",-1)

    tags := map[string]string{
        "loja":  werkstr,
    }

    fields := map[string]interface{}{
        "lote":        lote,
        "atualizado":  age.Format("20060102150405"),
    }

    acc.AddFields("ulg", fields, tags)

    return nil
}


func init() {
    inputs.Add("ulgp2ksp", func() telegraf.Input {
        return &ULGP2KSPStats{}
    })
}