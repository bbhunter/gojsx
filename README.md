#### Features
    - Search for javascript inside domain and grep
        - js with relative paths like ( ./foo.js, /foo2.js, //foo3.js, https://another.domain/foo4.js)
        - js inside <script> tags without src=
        - Regex like firebase and patterns like apis and tokens
#### Usage
  * No cookies/Auth
    - go run main.go -url https://domain.com
  * With Cookies
    - go run main.go -url https://domain.com/abc -auth "JSESSION=123; track_id=abc_123; AWSELB=123ABASR; _gat=1"

#### Requirements
    - Golang >= 1.16
    - Yaml config file ( see regexs_sample.yaml)

### TODO
    - Sessions ( authenticated pages) - Testing
    - Concurrency
    - multiple URLs ( scan many URLs from stdin )
    - Scan from burp ( Extract all scripts from BURP and parse them  )
    - Tunning regex

#### Running
![gojsx](./gojsx_alpha_01.gif)
