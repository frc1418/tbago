# tbago
_1418 TBA Libraries_ // [Python](https://github.com/frc1418/tbapy) // [**Go**](https://github.com/frc1418/tbago) // [Ruby](https://github.com/frc1418/tbarb)

> Go library for interfacing with [The Blue Alliance](https://thebluealliance.com) [API](https://thebluealliance.com/apidocs) (v3).

## Installation
```sh
go get github.com/frc1418/tbago
```

To use it in your project:

```go
import (
    ...
    "github.com/frc1418/tbago"
    ...
)
```

## Data Retrieval
Data retrieval may appear somewhat complicated at first, as it uses a [builder pattern](https://gist.github.com/vaskoz/10073335) scheme to structure requests. However, once you gain an understanding of this system the library is quite simple to use.

Before retrieving data, you must instantiate the library, providing a valid TBA API key. The Blue Alliance's API requires that all applications identify themselves with an auth key when retrieving data. To obtain an auth key, visit TBA's [Account page](https://www.thebluealliance.com/account).

```go
tba, err := tbago.New(token)
// Always handle errors
```

There are several base requests you can make: `Status()`, `Teams(page int)`, `Teams(number int)`, `Events(year int)`, `Event(id string)`, `Match(key string)`, `Districts(year int)`, and `District(abbreviation string, year int)`.

If you like, you can stop there:

```go
team, err := tba.Team(1418).Get()
// Handle error
```

However, you can also modify those requests before sending them. Many requests support `Simple()` or `Year()` modifiers for optional request parameters.

In addition, if you want to, say, get a list of the teams present at a given event, you would form your request thus:

```go
teams, err := tba.Event("2017chcmp").Teams().Get()
```

Specific options can be found by searching through the source code.


## Authors
`tbago` was originally published as `tba` by [Carl Colglazier](https://github.com/CarlColglazier).

Development is now led by [Erik Boesen](https://github.com/ErikBoesen) through [Team 1418](https://github.com/frc1418).

## License
This software is available under the terms of the [ISC License](LICENSE).
