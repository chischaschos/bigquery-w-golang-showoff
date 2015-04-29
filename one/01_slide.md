!SLIDE bullets
# BigQuery w/ golang
- by @chischaschos
- from @crowdint

~~~SECTION:notes~~~
- what do I do
- objective
~~~ENDSECTION~~~

!SLIDE[bg=githubarchive.png] background-fit
# https://www.githubarchive.org/

~~~SECTION:notes~~~
- Try to imagine how was this built!
~~~ENDSECTION~~~

!SLIDE[bg=githut.info.png] background-fit
# http://githut.info/

~~~SECTION:notes~~~
- Try to imagine how was this built!
~~~ENDSECTION~~~

!SLIDE bullets
# BigQuery Fundamentals
* Projects
* Tables
* Datasets
* Jobs

!SLIDE bullets
# Projects (P)
* Top level container
* Store info about:
  * billing
  * authorized users

!SLIDE bullets
# P -> Tables (T)
* Your data
* The table Schema
* Lives within a Dataset
* Created by:
  * Loading data into new table
  * Running a query
  * Copying a table

~~~SECTION:notes~~~
- Why are tables created by running a query?
~~~ENDSECTION~~~

!SLIDE bullets
# Datasets (D)
* Organizes/Contains tables
* Controls access to tables
* ACL per D

!SLIDE bullets
# Jobs (J)
* Actions built by you
* And then executed on your behalf to:
  * Load data
  * Export data
  * Query data
  * Copy data

!SLIDE
# Ways to interact with BigQuery
* UI
* API
* CLI

!SLIDE[bg=githubarchive-year.png] background-fit
# UI

!SLIDE[bg=githubachive-years-count.png] background-fit
# UI

!SLIDE[bg=golang-years-count.png] background-fit
# Go CLI

!SLIDE bullets
# Storage Pricing
* Loading data is free
* You are charged by storage and querying

!SLIDE[bg=storage-pricing.png] background-fit

!SLIDE bullets
# Querying Pricing
* First 1 TB of data processed per month is free of charge
* Not charged by cached or error queries
* Charges are rounded to the nearest MB
* Minimum 10 MB data processed per table referenced by the query
* 5$ per TB

!SLIDE[bg=free-operations.png] background-fit

!SLIDE
# An example User story
* I'm an anonymous user
* I want to see an organization's open source contributions

!SLIDE
# An example Architecture

!SLIDE[bg=arch.png] background-fit

!SLIDE bullets
# Implementation
* https://github.com/chischaschos/contributron-golang-api

!SLIDE[bg=rank.png] background-fit

!SLIDE bullets
# Thank you
* Made with https://github.com/puppetlabs/showoff
