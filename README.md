[![Build Status](https://travis-ci.org/espebra/libfileinfo.svg)](https://travis-ci.org/espebra/libfileinfo)

# libfileinfo

## Testing

    go test -coverprofile=coverage.out
    go tool cover -html=coverage.out

## Example

    package main
    
    import (
            "os"
            "fmt"
            "encoding/json"
            "github.com/espebra/libfileinfo"
    )
    
    func main() {
    	var fpath = "testing/IMG_3679.jpg"
    
            f, err := libfileinfo.Open(fpath)
            if err != nil {
                    fmt.Println(err)
                    os.Exit(2)
            }
    
            b, err := json.MarshalIndent(f, "", "    ")
            if err != nil {
                    fmt.Println(err)
                    os.Exit(2)
            }
            os.Stdout.Write(b)
            fmt.Println()
    }
