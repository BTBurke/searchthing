Simple search forwarding tool for those in China.  It detects your location and uses Bing Global if in Mainland China.  If you're on a VPN, it uses Google.  Runs locally on your computer.

To use:

1. Install [Go](https://golang.org)
2. Clone the repository into your GOPATH and then run `go build`
3. Searchthing will listen on port 10298
4. In your browser, set your search engine to http://127.0.0.1:10298/search?q=%s
5. When you search, searchthing will use ipinfo.io to determine your search location.  If you're in China, you'll be forwared to Bing Global search set to English.  If you're on a VPN with an IP address other than Mainland China, you'll be sent to Google.
