# É o Arbitas

## Environment Requirements

* Install Go

## Environment Setup

### Install dependencies

```sh
go mod
```

### Build

Unix based systems:
```sh
make build
```

Building for Linux:
```sh
make linux
```

It will generate transpiled js files into the `dist` folder.

### Execute

```
make run
```

### Server Configuration

1. Create file at `/etc/systemd/system/arbitas.service`
2. Add the Content
```
[Unit]
Description=Arbitas

[Service]
ExecStart=/home/ec2-user/main
Restart=always
Environment="RUN=true"

[Install]
WantedBy=multi-user.target
```
3. Run `sudo systemctl enable arbitas.service`
4. Run `sudo systemctl start arbitas.service`
5. See Logs: `journalctl -e -u arbitas.service`

## Available Tokens

| Token | Address                                   |
| ----- | ------------------------------------------ |
| ADA   | 0x3ee2200efb3400fabb9aacf31297cbdd1d435d47 |
| BNB   | 0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee |
| BUSD  | 0xe9e7cea3dedca5984780bafc599bd69add087d56 |
| DAI   | 0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3 |
| ETH   | 0x2170ed0880ac9a755fd29b2688956bd959f933f8 |
| USDC  | 0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d |
| USDT  | 0x55d398326f99059ff775485246999027b3197955 |
| WBNB  | 0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c |
