# BR-EKS-IT ;-) : A Browser Extension for work with EKS by IT teams.
#

```
My Goal for this project is to explore interfacing a Browser as a platform (DOM object) with
EKS cluster to (hopefully) cut down on time spent in routine checks on the EKS cluster and
state of its components/elements:
    Proposed is a development of browser extension that will alllow, among other features:
    - Send API requests to an EKS cluster
    - Perform "kubectl" commands on a cluster ("terminal session on a web page").

To interface a browser extension with EKS cluster proposed is creation of a backend service that acts
as an intermediary between browser plugin and EKS cluster.

Considerations and Activities:
=================================

Backend Service:
    1.) Server-side technologies: Echo vs Gin / Passenger vs Puma / Node.js etc (choice of language: Go, Ruby, Javascript)
    2.) Security: EKS Authentication vs Backend Service Authc/Authz (backend athc/"login" endpoint)
    3.) Handling Plugin requests via backend API endpoints: create / implement couple for POC/Demo.
    4.) ...

Plugin Design/ Dev:
    1.) UI/UX:
        1.A : balance popup only with full page (/tab) (for info requireing larger web real estate.
        1.B :  Any front end web framework to use?
    2.) Data persistence:
        1.A : Balance between browser "localStorage" (setItem/getItem) and "heavier" solutions (DB/Logs)
    3.) Nice to have features:
        3.1 : Scripted API actions (for automated / targeted EKS cluster info queries)
        3.2 : Data series collections: with configurable in-browser alerts
```
