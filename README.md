# bitly-svc

Start the service
```
$ docker-compose up 
```

And hit the /mystats endpoint. This endpoint will give you the average number of clicks per country per bitly link associated with your group. 
```
$ curl -H "Authorization: Bearer [TOKEN]" localhost:5000/mystats
```

My stats:
```
{
    "Stats":[
        {
            "metrics":[
                {"value":"US","clicks":0.033333335}
            ],
            "unit_reference":"average",
            "link":"bit.ly/2Kf5b02"
        },
        {
            "metrics":[
                {"value":"US","clicks":1.9}
            ],
            "unit_reference":"average",
            "link":"bit.ly/3nsOoVF"
        }
    ],
    "Timeframe":"30d"
}
```

