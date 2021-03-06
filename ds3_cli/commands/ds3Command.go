package commands

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
)

type command func(*ds3.Client, *Arguments) error

var availableCommands = map[string]command {
    "get_service": getService,
    "get_bucket": getBucket,
    "get_object": getObject,

    "put_bucket": putBucket,
    "put_object": putObject,

    "delete_bucket": deleteBucket,
    "delete_object": deleteObject,

    "bulk_get": bulkGet,
    "bulk_put": bulkPut,
}

func RunCommand(client *ds3.Client, args *Arguments) error {
    cmd, ok := availableCommands[args.Command]
    if ok {
        return cmd(client, args)
    } else {
        return fmt.Errorf("Unsupported command: '%s'", args.Command)
    }
}

