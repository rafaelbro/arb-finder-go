# É o Arbitas

## Environment Requirements

* Install NodeJS v12.X
* Install Yarn `npm install -g yarn`

## Environment Setup

### Install dependencies

```sh
yarn install
```

### Build

```sh
yarn build
```

It will generate transpiled js files into the `dist` folder.

### Execute

Run the following command providing the arguments.

1. Quantity of units desired to arbitrate (Required)
2. Ticker of source token to arbitrate (Required)
3. Ticker of destination token to arbitrate (Required)
4. Config to show values in a human readable way (Optional)

```sh
yarn start <QUANTITY> <FROM_TOKEN> <TO_TOKEN> [<HUMAN_READABLE>]
yarn start 10 BNB DAI
```
