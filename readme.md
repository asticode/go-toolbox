# About

Static methods to ease developments for the GO programming language (http://golang.org).

# Installing

## Using *go get*

    $ go get github.com/asticode/go-toolbox
    
After this command *go-gozzle* is ready to use. Its source will be in:

    $GOPATH/src/github.com/asticode/go-toolbox
    
# Example

    import (
        "github.com/asticode/go-toolbox/network"
    )
    
    // Validate IP addresses of http.request
    bIsValid = network.ValidateIPAddress(oHttpRequest, []string{
        "10.0.0.0/8",
        "172.16.0.0/16",
        "192.168.10.0/24",
    })
    