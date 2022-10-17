package main

func flows() string {
return `

[
    {
        "id": "9677e77d09b17122",
        "type": "inject",
        "z": "979ccb452915f125",
        "name": "",
        "props": [
            {
                "p": "payload"
            },
            {
                "p": "topic",
                "vt": "str"
            }
        ],
        "repeat": "1",
        "crontab": "",
        "once": false,
        "onceDelay": 0.1,
        "topic": "",
        "payload": "",
        "payloadType": "date",
        "x": 240,
        "y": 160,
        "wires": [
            [
                "3552c8ecdf006278"
            ]
        ]
    },
    {
        "id": "3552c8ecdf006278",
        "type": "debug",
        "z": "979ccb452915f125",
        "name": "debug 1",
        "active": true,
        "tosidebar": true,
        "console": true,
        "tostatus": false,
        "complete": "payload",
        "targetType": "msg",
        "statusVal": "",
        "statusType": "auto",
        "x": 450,
        "y": 160,
        "wires": []
    }
]

`
}
