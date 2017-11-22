package main

import (
	"fmt"
  "os"
  "flag"
)

func main() {

  if len(os.Args) > 1 {
    // work in automated state
    automatedStated()
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

func checkCommand(command string) {
    fmt.Printf("checking %s\n", command)
}

func showError(code string) {

}

func interactiveState() {
  fmt.Printf("GameJolt Interactive\n")
}

func automatedStated() {

  packageid := flag.String("packageid", "", "the package id")
  release := flag.String("release", "", "the new version")
  buildtarget := flag.String("buildtarget", "", "the build target of the upload")
  gameid := flag.String("gameid", "", "the game id")
  token := flag.String("token", "", "the authentication token")

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

  if *buildtarget == "" {
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

  if(ready) {
    fmt.Println("Ready to upload!")
    // start to upload
  }



}
