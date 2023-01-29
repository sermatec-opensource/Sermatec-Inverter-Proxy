# Sermatec Inverter Proxy

**Not production ready**

The aim of this project is to proxify the Sermatec inverter to enable multiple connection to it.
The reverse engineering of the Sermatec inverter communication protocol is based on these projects :
 - [Sermatec Inverter Python Lib](https://github.com/andreondra/sermatec-inverter)
 - [Sermatec Protocol extraction](https://github.com/sbechet/sermatec-ess)

Thanks to :
 * [Ondřej Golasowski](https://github.com/andreondra)
 * [Sebastien Bechet](https://github.com/sbechet)

There is few modes for running this application :
 * The client only mode: will enquire the inverter and log the information received on the StdOut
 * The proxy mode: Polling on a regular basis the inverter to get the data, and open few ports using the same protocol than inverter for other usage, and connection pool creation.
 * The mock server mode: creating a server which will answer a bunch of predefined answers, for working on the protocol client without having an inverter at disposal.

## Technology involved in the project
 * Golang
 * Docker
 * React JS

## Usage

There is no build for this application right now, needs docker to run.
It builds and run the command on the fly.
There is a makefile that provide you easier access to different running modes.

### Command line arguments

| argument                       | description                                           | optional / required | 
|--------------------------------|-------------------------------------------------------|---------------------|
| `-dot-env-file <path/to/file>` | load the dotenv file                                  | optional            |
| `-client`                      | enable the client to inverter                         | optional            |
| `-mock-server`                 | enable the mock server (imply no client and no proxy) | optional            |
| `-proxy`                       | enable the proxy server (imply client enabled)        | optional            |
| `-ui`                          | enable the webui server                               | optional            |
| `-h`                           | show this help                                        | optional            |

### DotEnv file
 
```env
INVERTER_HOST=<INVERTER IP OT HOST>
INVERTER_PORT=<INVERTER PORT USUALLY 8899>
CLIENT_POLLING_INTERVAL=<NUMBER OF SECONDS BETWEEN INVERTER POLL>
```

*Note: try to not poll too often the inverter, it has low limits... 15s or 30s are quite enough*
