package main

import (
    "archive/zip"
    "encoding/xml"
    "io/ioutil"
    "log"
)

type Documents struct {
    XMLName     xml.Name    `xml:"Data"`
    Files       []string    `xml:"ficheros>Items>FicherosItem>URL"`
    Domains     []Domain    `xml:"relations>Items>RelationsItem"`
}

type Domain struct {
    Name            string      `xml:"domain>domain"`
    IP              string      `xml:"ip>ip"`
    Maps            []Map      `xml:"domain>map"`

}

type Map struct {
    Documents       Files        `xml:"documents"`
    Parametrized    Files        `xml:"parametrized"`
    JuicyFiles      Files        `xml:"juicyFiles"`
    ListingFiles    Files        `xml:"listingFilesFound"`
    DSStoreFiles    Files        `xml:"DSStoreFilesFound"`
    SVNFiles        Files        `xml:"SvnEntriesFilesFound"`
    Methods         []Method    `xml:"domain>map>insecureMethodFoldersFound>FuzzMethodFolderObject"`
}

type Method struct {
    URL         string      `xml:"url"`
    Methods     string      `xml:"methods"`
}

type Files struct {
    URL     []string      `xml:",any"`
}

type MetaDataEmails struct {
    XMLName     xml.Name    `xml:"dictionary"`
    Emails      []User      `xml:"item>key"`    
}

type MetaDataUsers struct {
    XMLName     xml.Name    `xml:"dictionary"`
    Users       []User      `xml:"item>key"`
}

type User struct {
    User    string      `xml:"string"`
}

func filereader (f *zip.File) []byte {
    rc, err := f.Open()
    if err != nil {
        log.Fatal(err)
    }
    data,err := ioutil.ReadAll(rc)
    if err != nil {
        log.Fatal(err)
    }
    return data
}

func focareader (filename string) (u MetaDataUsers, e MetaDataEmails, d Documents){
    r, err := zip.OpenReader(filename)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, f := range r.File {
        if f.Name == "metadatasummaryemails" {
            err = xml.Unmarshal(filereader(f), &e)
            if err != nil {
                log.Println(err)
            }
        } else if f.Name == "metadatasummaryusers" {
            err = xml.Unmarshal(filereader(f), &u)
            if err != nil {
                log.Println(err)
            }
        } else if f.Name == "documents" {
            err = xml.Unmarshal(filereader(f), &d)
            if err != nil {
                log.Println(err)
            }           
        }
    }

    defer r.Close()
    return u, e, d
}
