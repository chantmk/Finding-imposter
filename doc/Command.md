# Command list
This list contain only command for Finding-imposter and some useful command.
## Transaction command
> Prefix: Finding-impostercli tx Findingimposter

|Command | API | Argument/Body | Description|
|-|-|-|-|
|create-log| POST /log| logID, placeID, action| create log tx with required information
|create-quarantine| POST /quarantine| address | create quarantine tx with required information
|create-covid| POST /covid| covidID, status, address(s) | create covid tx with required information and query with provided address(s)
|create-doctor| POST /doctor| address, isDoctor | create doctor tx with required information
---

## Query command
> Prefix: Finding-impostercli query Findingimposter

|Command|API|Arguments|Description|
|-|-|-|-|
|list-log| GET /log|-| list all log
|list-spec-log| POST /log/list| address(s)| list log created with specify address|
|list-quarantine| GET /quarantine|-| list all quarantine
|list-spec-quarantine| POST /quarantine/list| address(s)| list quarantine that specify address need to be isolate|
|list-covid| GET /covid|-| list all covid
|list-spec-covid| POST /covid/list| address(s)| list covid that specify address affected by COVID-19|
|list-pending-covid| POST /covid/pending |-| list all covid transaction with status PENDING
|list-doctor| GET /doctor|-| list all doctor

---
## Other command

> Prefix: Finding-impostercli

|Command|Arguments|Description|
|-|-|-|
|keys list|-| list all public keys stored.|
|query tx | hash | query tx with specified hash|
|query account | address | query account with specified address|
---