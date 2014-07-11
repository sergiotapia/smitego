// Package smitego is an API wrapper for Hirez' Smite game API.
//
// If you want easily consume data from Hirez' Smite API server then smitego
// is definitely the easiest way to do it.
//
// For a full guide visit http://github.com/sergiotapia/smitego
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/kr/pretty"
//    "github.com/sergiotapia/smitego"
//  )
//
//  func main() {
//    smitego.DevID = "1132"
//    smitego.AuthKey = "29FA3BCE7ACF4BDDA7310A352FB855BE"
//
//    smitego.GetSessionID()
//
//    match := smitego.GetMatchDetails("78900556")
//    fmt.Printf("%# v", pretty.Formatter(match))
//  }
package smitego
