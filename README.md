Detect Language API Go Client
========

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/detectlanguage/detectlanguage-go)
[![Build Status](https://github.com/detectlanguage/detectlanguage-go/actions/workflows/main.yml/badge.svg)](https://github.com/detectlanguage/detectlanguage-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/detectlanguage/detectlanguage-go)](https://goreportcard.com/report/github.com/detectlanguage/detectlanguage-go)

Detects language of the given text. Returns detected language codes and scores.

Before using Detect Language API client you have to setup your personal API key.
You can get it by signing up at https://detectlanguage.com

## Installation

```
go get -u github.com/detectlanguage/detectlanguage-go@v2.0.0
```

### Configuration

```go
dl := detectlanguage.New("YOUR API KEY")
```

## Usage

### Language detection

```go
detections, err := dl.Detect(context.TODO(), "Dolce far niente")

if err != nil {
    fmt.Fprintln(os.Stderr, "error detecting language:", err)
    os.Exit(1)
    return
}

fmt.Fprintln(os.Stdout, "Language:", detections[0].Language)
fmt.Fprintln(os.Stdout, "Score:", detections[0].Score)
```

### Single language code detection

If you need just a language code you can use `DetectCode`. It returns first detected language code.

```go
language, err := dl.DetectCode(context.TODO(), "Dolce far niente")

if err != nil {
    fmt.Fprintln(os.Stderr, "error detecting language:", err)
    os.Exit(1)
    return
}

fmt.Fprintln(os.Stdout, "Language:", language)
```

### Batch detection

It is possible to detect language of several texts with one request.
This method is significantly faster than doing one request per text.
To use batch detection just pass multiple texts to `DetectBatch` method.

```go
texts := []string{"labas rytas", "good morning"}
results, err := dl.DetectBatch(context.TODO(), texts)

if err != nil {
    fmt.Fprintln(os.Stderr, "error detecting language:", err)
    os.Exit(1)
    return
}

fmt.Fprintln(os.Stdout, "First text language:", detections[0][0].Language)
fmt.Fprintln(os.Stdout, "Second text language:", detections[1][0].Language)
```

### Getting your account status

```go
result, err := dl.AccountStatus(context.TODO())

if err != nil {
    fmt.Fprintln(os.Stderr, "error getting account status:", err)
    os.Exit(1)
    return
}

fmt.Fprintln(os.Stdout, "Status:", result.Status)
fmt.Fprintln(os.Stdout, "Requests sent today:", result.Requests)
fmt.Fprintln(os.Stdout, "Bytes sent today:", result.Bytes)
fmt.Fprintln(os.Stdout, "Plan:", result.Plan)
fmt.Fprintln(os.Stdout, "Plan expires:", result.PlanExpires)
fmt.Fprintln(os.Stdout, "Daily requests limit:", result.DailyRequestsLimit)
fmt.Fprintln(os.Stdout, "Daily bytes limit:", result.DailyBytesLimit)
fmt.Fprintln(os.Stdout, "Date:", result.Date)
```

### Getting list supported languages

```go
languages, err := dl.Languages(context.TODO())

if err != nil {
    fmt.Fprintln(os.Stderr, "error getting languages list:", err)
    os.Exit(1)
    return
}

fmt.Fprintln(os.Stdout, "Supported languages:", len(languages))
fmt.Fprintln(os.Stdout, "First language code:", languages[0].Code)
fmt.Fprintln(os.Stdout, "First language name:", languages[0].Name)
```

## License

Detect Language API Go Client is free software, and may be redistributed under the terms specified in the MIT-LICENSE file.
