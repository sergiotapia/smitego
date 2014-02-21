SmiteGo
=======

SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go!

## How to use?

SmiteGo is really simple to use. Below is a sample go file that calls the various API features HiRez offers.

    package main

    import (
        "fmt"
        "github.com/sergiotapia/smitego"
    )

    func main() {

        // This is a demo of the Smitego API package.
        // Source code: www.github.com/sergiotapia/smitego

        // Start by providing your DevId and AuthKey.
        // If you don't have one request it from Hirez.
        smitego.DevId = "7777"
        smitego.AuthKey = "29123BCE7ACF4B1237310A123FB123BE"

        // Smite API calls require a session_id that must be
        // generated every 15 minutes. Of course Smitego handles
        // this for you automatically. Just call for the info you
        // need and Smitego will refresh the session_id as needed.

        // APIs - Connectivity
        // ------------------------
        smitego.GetSessionId()
        fmt.Println("SessionID:", smitego.SessionId)

        pingResponse := smitego.Ping()
        fmt.Println(pingResponse)

        // APIs - Smite Data
        // ------------------------
        dataUsed := smitego.GetDataUsed()
        fmt.Println(dataUsed)

        friends := smitego.GetFriends("FaymousHate")
        fmt.Println(friends)

        godRanks := smitego.GetGodRanks("FaymousHate")
        fmt.Println(godRanks)

        gods := smitego.GetGods()
        fmt.Println(gods)

        items := smitego.GetItems()
        fmt.Println(items)

        match := smitego.GetMatchDetails("45313775")
        fmt.Println(match)

        matchHistory := smitego.GetMatchHistory("FaymousHate")
        fmt.Println(matchHistory[0].Match)

    }

## License

The MIT License (MIT)

Copyright (c) 2013 Sergio Tapia Gutierrez

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
