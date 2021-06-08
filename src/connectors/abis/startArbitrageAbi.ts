export const startArbitrageAbi = [
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "amountBorrowed",
        "type": "uint256"
      },
      {
        "internalType": "uint256[]",
        "name": "routerPath",
        "type": "uint256[]"
      },
      {
        "internalType": "address[]",
        "name": "tokenPath",
        "type": "address[]"
      }
    ],
    "name": "startArbitrage",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "index",
        "type": "uint256"
      }
    ],
    "name": "getRouterIn",
    "outputs": [
      {
        "internalType": "address",
        "name": "routerAddress",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function",
    "constant": true
  }
];
