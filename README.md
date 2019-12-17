# Cloudsourced

DIRECTORIES
- cloudsourced - application server
- dashboard - web dashboard

PREREQUISITES:
- Your own MongoDB server
- Go with MongoDB driver installed
- Node.js

TO RUN:
1. Go into the dashboard directory
2. `npm install`
3. `npm run build` (this will generate the web dashboard files)
4. Go to the cloudsourced directory
5. Create a config.json file (see below)
6. `go run .`
7. Site will be hosted on :8080

Config.json - This tells Cloudsourced how to connect to your MongoDB server
```json
{
    "mongoHost": "mongoHostname",
    "mongoUser": "mongoUsername",
    "mongoPass": "mongoPassword"
}
```