package main

import (
	"fmt"
  "os"
  "flag"
  // "log"
  // "net/http"
  // "net/http/httputil"
  // "net/url"
  "bufio"
  // "errors"
  // "strings"
  // "io"
  "io/ioutil"
  // "path/filepath"
  "encoding/json"
  // "mime/multipart"
)

type Deploys struct {
	Deploys []Deploy `json:"deploys"`
}

type Deploy struct {
	PackageId string `json:"packageid"`
	Release string `json:"release"`
	GameId string `json:"gameid"`
  Token string `json:"token"`
	Buildtarget string `json:"buildtarget"`
  Filename string `json:"filename"`
}

func main() {

  if len(os.Args) > 1 {

    if(os.Args[1] == "push") {
      // work in push state
      pushState()
    } else {
      // work in automated state
      automatedState()
    }

  } else {
    // work in interactive state
    interactiveState()
  }

}

func checkConnectivity() {

}

func checkAuthentication() {

}

func checkForUpdates() {

}

func deployToGameJolt(packageid, release, gameid, token, buildtarget, filename string) {
  fmt.Println("Dummy deployed!")
}

func checkAndRunPushConfig() {
  config, err := os.Open("config.json")
  defer config.Close()

  if(err != nil) {
    fmt.Println(err.Error())
    os.Exit(1)
  } else {
    byteValue, _ := ioutil.ReadAll(config)
    var deploys Deploys
    json.Unmarshal(byteValue, &deploys)

    for i := 0; i < len(deploys.Deploys); i++ {

      if(deploys.Deploys[i].PackageId == "") {
        fmt.Println("PackageId is empty.")
        os.Exit(1)
      }
      if(deploys.Deploys[i].Release == "") {
        fmt.Println("Release is empty.")
        os.Exit(1)
      }
      if(deploys.Deploys[i].GameId == "") {
        fmt.Println("GameId is empty.")
        os.Exit(1)
      }
      if(deploys.Deploys[i].Token == "") {
        fmt.Println("Token is empty.")
        os.Exit(1)
      }
      if(deploys.Deploys[i].Buildtarget == "") {
        fmt.Println("Buildtarget is empty.")
        os.Exit(1)
      }
      if(deploys.Deploys[i].Filename == "") {
        fmt.Println("Filename is empty.")
        os.Exit(1)
      }

      deployToGameJolt(deploys.Deploys[i].PackageId, deploys.Deploys[i].Release, deploys.Deploys[i].GameId, deploys.Deploys[i].Token,
      deploys.Deploys[i].Buildtarget, deploys.Deploys[i].Filename)
    }

  }
}

func checkCommand(command string) {
    fmt.Printf("checking %s\n", command)
}

func interactiveState() {

  var packageid string
  var release string
  var gameid string
  var token string
  var buildtarget string
  var filename string

  fmt.Printf("GameJolt CLI - Interactive Mode \n")
  var ready = false

  gji := bufio.NewScanner(os.Stdin)
  var command string
  for command != "q" || ready == true {
  fmt.Print("> ")
  gji.Scan()

  command = gji.Text()
    if command != "q" {
      // parse the input here.
    }

  }

  if(ready) {
    deployToGameJolt(packageid, release, gameid, token, buildtarget, filename)
  }

}

func pushState() {
  fmt.Printf("pushing...\n")
  checkAndRunPushConfig()
}

func automatedState() {

  packageid := flag.String("packageid", "", "the package id")
  release := flag.String("release", "", "the new version")
  gameid := flag.String("gameid", "", "the game id")
  token := flag.String("token", "", "the authentication token")
  buildtarget := flag.String("buildtarget", "", "the build target of the upload")
  filename := flag.String("filename", "", "the full path to the file to upload")

  fmt.Printf("checking...\n")
  flag.Parse()

  // let's not allow empty flags

  var ready = true

  if *packageid == "" {
    fmt.Println("the package id was not supplied. ")
    // .. additional checks here
    ready = false
  }

  if *release == "" {
    fmt.Println("the release version was not supplied. ")
    // .. additional checks here
    ready = false
  }

  if *gameid == "" {
    fmt.Println("the game id was not supplied. ")
    // .. additional checks here
    ready = false
  }

  if *token == "" {
    fmt.Println("the token is not valid. ")
    // .. additional checks here
    ready = false
  }

  if *buildtarget == "" {
    fmt.Println("the release version was not supplied. ")
    // .. additional checks here
    ready = false
  }

  if *filename == "" {
    fmt.Println("the file does not exist. ")
    // .. additional checks here
    ready = false
  }

  if(ready) {
    fmt.Println("Ready to upload!")
    deployToGameJolt(*packageid, *release, *gameid, *token, *buildtarget, *filename)
  }



}
