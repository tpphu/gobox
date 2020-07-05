# Usage
Use `boilr help` to get the list of available commands.

## Init Template
In order to init a template from gobox, use the following command:

```bash
boilr init -f -s ["service-name"] -p ["path-of-project"] -t ["template"] -r ["resources"]
-f: "Recreate directories if they exist"
-t: "Template to use (gin, grpc). gin is default"
-r: "Resource to use: logfile, redis, mgo. Default are logfile-redis-mgo"
-p: "Destination path. Default is current directory"
```
