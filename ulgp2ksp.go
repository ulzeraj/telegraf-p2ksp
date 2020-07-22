package ulgp2ksp

import (
    "os"
    "fmt"
    "log"
    "strings"
    "strconv"
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
    verfile   := "/p2ksp/bin/versaoSP.dat"
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

    age          := fileage.ModTime()
    lotestr      := strings.Replace(string(content), "\n", "", -1)
    loteint, err := strconv.Atoi(lotestr)
    if err != nil {
        log.Fatal(err)
    }

    version, err := ioutil.ReadFile(verfile)
    if err != nil {
        log.Fatal(err)
    }

    verinfo     := strings.Replace(string(version), "VERSAO_SP = ", "", -1)
    verstr      := strings.Replace(string(verinfo), "\n", "", -1)

    tags := map[string]string{
        "loja":  werkstr,
    }

    fields := map[string]interface{}{
        "lote":        loteint,
        "loja":        werkstr,
        "vers":        verstr,
        "atualizado":  age.Format("2006-01-02-15:04:05"),
    }

    acc.AddFields("ulgp2ksp", fields, tags)

    return nil
}


func init() {
    inputs.Add("ulgp2ksp", func() telegraf.Input {
        return &ULGP2KSPStats{}
    })
}
