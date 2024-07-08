## Backlog

- Update projects with Git commands
    > development, staging, production branches will be updated automatically.
- Get which branches to update from config

## Todo

- Generate a config file
    > Config file will be generated in users home config folder. It'll be a json config because go stdlib has  a JSON parser and didn't want to increase dependencies.
    * [ ] Detect home folder
    * [x] Decide where to put config
    * [x] Make config sturct
    * [ ] When App first lunches check for config if not found generate
- Create CLI Interface
    > CLI interface will be used for following:
    * [ ] See logs 
    * [ ] See last sync status and time
    * [ ] See monitoring information
    * [ ] Control how projects are fethched and etc
- Get Data from Gitlab
    > Get group data from gitlab group API. Gitlab group api provides which repositories are in the group. 
    * [ ] Create gitlab client
    * [ ] Get group data, save them in the config
- Clone git projects with Gitlab data
    > Clone gitlab projects according to config
    * [ ] Set gitlab repository and config 
    * [ ] Config for where to put the repo
    * [ ] Last updated at? 
    * [ ] 

## Doing


## Done

