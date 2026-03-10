# moca-relayer

The Moca Relayer is a tool that allows bidirectional communication between Moca and ethereum-compatible chain. It operates separately
and can only be utilized by Moca validators. This relayer continuously monitors cross-chain events on evm-compatible chain and
Moca, storing them in a database. After a few blocks have been confirmed and reached finality, the relayer uses a BLS
private key to sign a message confirming the event. This signed event, also known as "the vote", is then broadcasted through
Moca's p2p network. Once enough votes are collected from Moca relayers, the relay assembles a cross-chain package
transaction and submits it to either the evm-compatible chain or Moca network.

## Disclaimer

**The software and related documentation are under active development, all subject to potential future change without
notification and not ready for production use. The code and security audit have not been fully completed and not ready
for any bug bounty. We advise you to be careful and experiment on the network at your own risk. Stay safe out there.**

## Main Components

The relayer mainly consists of 3 components: Listener, Vote Processor and Transaction Assembler.

1. The Listener component actively monitors blockchains for any cross-chain events and stores them in the database.

2. The Vote Processor component performs the following functions:
   a. retrieves unprocessed cross-chain events from database, signs and broadcasts votes for them to the Moca P2P network.
   b. collects enough valid votes for cross-chain events from the Moca P2P network and saves them to the database.

3. The Transaction Assembler component prepares and submits transactions to the destination chain by aggregating the
   votes and signatures of cross-chain events that have received enough consensus votes.

### Requirement

Go version above 1.20

## Deployment

### Config

1. Set your relayer private key and bls private key imported method (via file or aws secret), deployment environment and gas limit.

``` json
  "moca_config": {
    "key_type": "local_private_key",
    "aws_region": "",
    "aws_secret_name": "",
    "aws_bls_secret_name": "",
    "rpc_addrs": [
      "http://127.0.0.1:26750",
      "http://127.0.0.1:26751",
      "http://127.0.0.1:26752"
    ],
    "private_key": "your_private_key",
    "bls_private_key": "your_private_key",
    "chain_id": 18,     // moca oracle module defines this
    "start_height": 1,
    "monitor_channel_list": [1,2,3,4,5,6,10,11],
    "gas_limit": 1000,
    "fee_amount": 5000000000000,
    "chain_id_string": "moca_5151-1",
    "use_websocket": true
  }, 
  "evm-compatible chain_config": {
    "key_type": "local_private_key",
    "aws_region": "",
    "aws_secret_name": "",
    "rpc_addrs": [
      "localhost:8502"
    ],
    "private_key": "your_private_key",
    "gas_limit": 4700000,
    "gas_price": 10000000000,
    "number_of_blocks_for_finality": 2,
    "start_height": 0,
    "chain_id": 714
  }
```

2. Config crosschain and moca light client smart contracts addresses, others can keep default value.

```
"relay_config": {
    "evm-compatible chain_to_moca_inturn_relayer_timeout": 40,
    "moca_to_evm-compatible chain_inturn_relayer_timeout": 30,
    "moca_sequence_update_latency": 8,
    "evm-compatible chain_sequence_update_latency": 12,
    "cross_chain_contract_addr": "0xd2253A26e6d5b729dDBf4bCce5A78F93C725b455",
    "moca_light_client_contract_addr": "0x349a42f907c7562B3aaD4431780E4596bC2a053f",
    "relayer_hub_contract_addr": "0x70712e1321e8fC629375A72499BDcA7192EBD054",
    "src_mocasbt_contract_addr": "0x30258B344800c639013a5090bfCC959AB513f295",
    "src_mocavc_contract_addr": "0xEfdefe08C6cD74CFEB2f0CC2B9401c52B859B427"
  }
```

3. Set your log and backup preferences.

```
"log_config": {
  "level": "DEBUG",
  "filename": "log.txt",
  "max_file_size_in_mb": 100 (file size threshold)  
  "max_backups_of_log_files": 2 (backup count threshold)
  "max_age_to_retain_log_files_in_days": 10 (backup age threshold)
  "use_console_logger": true,
  "use_file_logger": false,
  "compress": false
}
```

4. Config your database settings. We Support mysql or sqlite.

example: use mysql

```
"db_config": {
  "dialect": "mysql",
    "key_type": "local_private_key",
    "aws_region": "",
    "aws_secret_name": "",
    "password": "pass",
    "username": "root",
    "url": "/moca-relayer?charset=utf8&parseTime=True&loc=Local",
    "max_idle_conns": 10,
    "max_open_conns": 100
}
```

 use sqlite

```
  "db_config": {
    "dialect": "sqlite3",
    "key_type": "",
    "aws_region": "",
    "aws_secret_name": "",
    "password": "",
    "username": "",
    "url": "moca-relayer.db",
    "max_idle_conns": 10,
    "max_open_conns": 100
  },

```

5. Set alert config to send a telegram message when the data-seeds are not healthy.

```
"alert_config": {
  "identity": your_bot_identity
  "telegram_bot_id": your_bot_id
  "telegram_chat_id": your_chat_id  
}
```

## Build

Build binary:

```shell script
make build
```

Build docker image:

```shell script
make build_docker
```

## Run locally

### Run MySQL in Docker(this can be skipped if you are using sqlite)

```shell
docker run --name gnfd-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:8.0
```

#### Create Schema

Create schema in MySQL client:

```mysql
CREATE SCHEMA IF NOT EXISTS `moca_relayer` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
CREATE SCHEMA IF NOT EXISTS `moca_relayer_polygon` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
CREATE SCHEMA IF NOT EXISTS `moca_relayer_linea` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
...
```

### Fill in config file

Get relayer private key and bls private key in Hex format, fill them in the config/config.json alone with Moca,
evm-compatible chain network RPC addresses, chain id and evm-compatible chain smart contracts addresses.

```shell script
./build/moca-relayer --config-type [local or aws] --config-path config_file_path  --aws-region [aws region or omit] --aws-secret-key [aws secret key for config or omit]
```

Example:

```shell script
./build/moca-relayer --config-type local --config-path config/config.json
```

Run docker:

```shell script
docker run -it -v /your/data/path:/moca-relayer -e CONFIG_TYPE="local" -e CONFIG_FILE_PATH=/your/config/file/path/in/container -d moca-relayer
```

### Quick setup for running multiple relayers in local

Fill in config files under `./config/local` by following above instruction, you might want to fill in same number of moca validators you bootstrap in local,

```bash
// start n instance of relayer 
bash ./deployment/localup/localup.sh start ${SIZE}

// stop relayer
bash ./deployment/localup/localup.sh stop
```

## Contribute

Thank you for considering to help out with the source code! We welcome contributions
from anyone, and are grateful for even the smallest of fixes!

Please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base if you would like to.

Please make sure your contributions adhere to our coding guidelines:

* Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting)
  guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
* Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
  guidelines.
* Pull requests need to be based on and opened against the `master` branch.
* Commit messages should be prefixed with the package(s) they modify.

## License

The repo is licensed under the
[GNU Affero General Public License v3.0](https://www.gnu.org/licenses/agpl-3.0.en.html), also
included in our repository in the `COPYING` file.

## Fork Information

This project is forked from [greenfield-relayer](https://github.com/mocachain/moca-relayer). Significant changes have been made to adapt the project for specific use cases, but much of the core functionality comes from the original project.
