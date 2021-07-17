#### Features

#### Usage
  * No cookies/Auth
    - go run main.go -url https://domain.com
  * With Cookies
    - go run main.go -url https://domain.com/abc -auth "JSESSION=123; track_id=abc_123; AWSELB=123ABASR; _gat=1"
  * Multiple URLs 
    - cat urls.txt | go run main.go

#### Requirements
    - Golang >= 1.16
    - Yaml config file ( see regexs_sample.yaml)

#### Features
- Regexs extracted from [https://github.com/l4yton/RegHex](https://github.com/l4yton/RegHex) many thanks @l4yton =)
- Sessions single URL
- Concurrency
- Multiple URLs Scan ( testing phase )
- Search for javascript inside domain and grep
    - js with relative paths like ( ./foo.js, /foo2.js, //foo3.js, https://another.domain/foo4.js)
    - Regex like firebase and patterns like apis and tokens
### TODO
    - Scan from burp ( Extract all scripts from BURP and parse them  )
    - Tunning regex

#### Running
![gojsx](./gojsx.gif)
