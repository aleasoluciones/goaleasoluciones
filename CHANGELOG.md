2019-06-11
==========
* Internal: Update logger. Now it does not panic if SENTRY_DSN is not define. 
            When LogError2Sentry is called, it logs the error only and not publish to Sentry

2018-10-29
==========
* Internal: Update go version 1.10 and 1.11
* Internal: Update godep version

2018-01-30
==========
* Internal: Used godep for dependency management tool

2017-12-15
==========
* Internal: Removed go version 1.6 and 1.7. Added 1.8.x, 1.9.x and master versions

2016-10-26
==========
* Internal: Added go 1.7 version to travis
* Internal: Removed go 1.5 version to travis. Golint needs 1.6 or later

2016-04-19
==========
* Internal: Unify Makefile
* Initial Changelog file
