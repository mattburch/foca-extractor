package main

import (
    "github.com/docopt/docopt.go"
    "fmt"
    "log"
    "errors"
    "strings"
)

func (u MetaDataUsers) UserClean() {
    for _, user := range u.Users {
        fmt.Printf("User: %v\n", user.User)
    }
}
func (e MetaDataEmails) EmailClean() {
    for _, email := range e.Emails {
        fmt.Printf("Email: %v\n", email.User)
    }    
}
func (f Files) FilesClean() {
    for _, f := range f.URL {
        fmt.Printf("\t %v\n", f)
    }
}
func (d Documents) DocsClean() {
    for _, file := range d.Files {
        fmt.Printf("File: %v\n", file)
    }
    fmt.Printf("\nBreakout by Domain\n\n")
    for _, domain := range d.Domains {
        for _, m := range domain.Maps {
            fmt.Printf("%v [%v] (Docs): \n", domain.Name, domain.IP)
            m.Documents.FilesClean()
            fmt.Printf("%v [%v] (Param): \n", domain.Name, domain.IP)
            m.Parametrized.FilesClean()
            fmt.Printf("%v [%v] (Juicy): \n", domain.Name, domain.IP)
            m.JuicyFiles.FilesClean()
            fmt.Printf("%v [%v] (Listing): \n", domain.Name, domain.IP)
            m.ListingFiles.FilesClean()
            fmt.Printf("%v [%v] (DSSStore): \n", domain.Name, domain.IP)
            m.DSStoreFiles.FilesClean()
            fmt.Printf("%v [%v] (SVN): \n", domain.Name, domain.IP)
            m.SVNFiles.FilesClean()
        }
    }
}
func (d Documents) SearchDoc(sdoc string) ([]string, error) {
    fileList := []string{}
    for _, file := range d.Files {
        s := strings.Split(file, "/")
        if s[len(s) - 1] == sdoc {
            fileList = append(fileList, file)
        } 
    }
    if len(fileList) < 0 {
        return fileList, errors.New(sdoc + " was not found in file listing")
    }
    return fileList, nil
}

func main() {
    arguments, err := docopt.Parse(usage, nil, true, "foca-extractor 0.1", false)
    if err != nil {
        log.Fatal("Error parsing usage. Error: ", err.Error())
    }

    f := arguments["<file>"].(string)
    u,e,d := focareader(f)

    switch {
    case arguments["--users"].(bool):
        u.UserClean()
    case arguments["--emails"].(bool):
        e.EmailClean()
    case arguments["--docs"].(bool):
        d.DocsClean()
    case arguments["--search"].(bool):
        sdoc := arguments["<search>"].(string)
        url, err := d.SearchDoc(sdoc)
        if err != nil {
            log.Println(err)
        } else {
            fmt.Println(url)
        }

    default:
        u.UserClean()
        e.EmailClean()
        d.DocsClean()
    }


}