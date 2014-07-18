SmiteGo
=======

Package smitego is an API wrapper for Hirez' Smite game API.

If you want easily consume data from Hirez' Smite API server then smitego
is definitely the easiest way to do it.

```go
package main

import (
  "fmt"
  "github.com/kr/pretty"
  "github.com/sergiotapia/smitego"
)

func main() {
  smitego.DevID = "1132"
  smitego.AuthKey = "29FA3BCE7ACF4BDDA7310A352FB855BE"

  smitego.GetSessionID()

  match := smitego.GetMatchDetails("78900556")
  fmt.Printf("%# v", pretty.Formatter(match))
}
```

## Supported API Endpoints

- [ ] API - Connectivity
  - [ ] Ping
  - [x] CreateSession
  - [ ] TestSession
- [ ] API - Smite Data
  - [x] GetDataUsed
  - [x] GetDemodeDetails
  - [x] GetFriends
  - [x] GetGodRanks
  - [x] GetGods
  - [x] GetGodRecommendedItems
  - [x] GetItems
  - [x] GetMatchDetails
  - [ ] GetMatchIdsByQueue - *note: Look into bug, is it on Hirez's side?*
  - [ ] GetLeagueLeaderBoard
  - [ ] GetLeagueSeasons
  - [ ] GetMatchHistory
  - [ ] GetPlayer
  - [ ] GetQueueStats
  - [ ] GetTeamDetails
  - [ ] GetTeamMatchHistory
  - [ ] GetTeamPlayers
  - [ ] GetTopMatches
  - [ ] SearchTeams

## License

The MIT License (MIT)

Copyright (c) 2014 Sergio Tapia Gutierrez

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
