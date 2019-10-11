# Service Deployment


## Use it

If the default service looks like

![](images/default.png)

You can switch between `green` and `blue` deployment with setting the `DEPLOYMENT` variable:

```bash
export DEPLOYMENT=blue
```

![](images/blue.png)


```bash
export DEPLOYMENT=green
```

![](images/green.png)

## Building

You can build the service with `build.sh` - you need docker for that one. 

## Running

You can run the service with `run.sh` - you need docker for that one. 

## Pull from Docker

There is already a version present.
Just do a `docker pull phaus/service-deployment:latest`


Enjoy!
Leave a comment to tell me, what you are doing with this!