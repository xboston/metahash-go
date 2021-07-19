# MetaHash API GO library

An unofficial Golang library for [#MetaHash](https://metahash.org ) blockchain.


![metahash-go](https://raw.githubusercontent.com/xboston/metahash-go/master/media/metahash-go.png)

### Requirements

- GO go 1.12+

### Installation

```bash
go get -u github.com/xboston/metahash-go
```

### Information

- [Missing #MetaHash API documentation](https://github.com/xboston/metahash-api)
- [Original source](https://github.com/metahashorg/crypt_example_php)
- [Knowledge base](https://developers.metahash.org)
- [Testpage portal](http://testpage.metahash.org/)

### Methods

- [x] fetch-balance
- [x] fetch-balances
- [x] fetch-history
- [x] get-tx
- [x] get-block-by-hash
- [x] get-block-by-number
- [x] get-last-txs
- [x] get-blocks
- [x] get-dump-block-by-number
- [x] get-dump-block-by-hash
- [x] get-count-blocks
- [x] get-last-node-stat-result
- [x] get-last-node-stat-trust
- [x] get-address-delegations
- [x] get-forging-sum-all
- [ ] status
- [ ] mhc_send
- [ ] getinfo


### Extra Methods
- [ ] [generateKey](https://developers.metahash.org/hc/en-us/articles/360002712193-Getting-started-with-Metahash-network)
- [ ] getNonce

### Usage
You can find usage examples in the [examples](https://github.com/xboston/metahash-go/tree/master/examples) folder.

## License

This package is released under the MIT license.
