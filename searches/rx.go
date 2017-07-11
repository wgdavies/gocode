package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "encoding/xml"
)


type Prop struct {
    XMLName xml.Name `xml:"property"`
    Name string `xml:"name"`
    Value string `xml:"value"`
}

type Config struct {
    XMLName xml.Name `xml:"configuration"`
    Props []Prop `xml:"property"`
}

func (s Prop) String() string {
    var rval string
    if s.Name == `hive.query.string` {
        rval = s.Value
    }

    return rval
}
/*        return fmt.Sprintf("%40s : %s\n", s.Name, s.Value) */

func main() {
    xmlFile, err := os.Open("txt.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()
    
    XMLdata, _ := ioutil.ReadAll(xmlFile)
    
    var c Config
    xml.Unmarshal(XMLdata, &c)
    
    fmt.Println(c.Props)
}
