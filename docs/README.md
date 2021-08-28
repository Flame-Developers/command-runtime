# Documentation for Flame Command Runtime(FCR)
Hello there! Welcome to FCR documentaiton. FCR documentation is intended for developers, who want to use FCR in their projects.

## Aim
FCR's aim is to provide Discord bots highly performant runtime for custom commands. Internally, it uses Tengo or Go templates(depending on config & package's details).

## Use
FCR uses NATS queue groups to load balance between multiple instances. You can see a concise description of protocol [here](docs/PROTOCOL.md)