# Marketstack
Marketstack Go API

## Quick Start

``` go
builder, err := request.NewBuilder(
    request.BuilderWithBaseURL("https://api.marketstack.com/v1/"),
    request.BuilderWithAPIKey("your-api-key"),
)
if err != nil {
    // handle error
}

client, err := api.NewClient(builder)
if err != nil {
    // handle error
}

res, err := client.EOD(context.TODO, []string{"AAPL"})
if err != nil {
    // handle error
}

log.Printf("%+v", res)
```

The above example will request all stocks for the AAPL ticker. Variadac options 
are used in the `client.EOD` function to add optional parameters to the request,
eg. sort order, date range, etc. More information on optional parameters can
be found for each specific message on the official Market Stack REST API
documentation: https://marketstack.com/documentation

## Testing

`make test`