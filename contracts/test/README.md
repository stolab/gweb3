# Run test for contract package.

In order to run the test to test the [contract package](./) a local blockchain with [the test contract](./test/TestContract.sol) deploy on it should be available.

The easier way to quickly deploy a local blockcchain is to use the `--dev` option of geth.
To easily compile and deploy the smart contract, we will use [remix](https://remix.ethereum.org).

## Start the local blockchain
As we want remix to be able to deloy the smart contract to our local blockchain, we need to gave few more parameter to geth when starting.
To start geth, use the following 
```bash
geth --dev --http --http.corsdomain "https://remix.ethereum.org"
```

Once the command has been runned, you should see something similar to the following:

Search for the line `Using developer account` and take note of the address, it will be needed later.

## Deploy the smart contract.
Copy the content of the [smart contract](./TestContract.sol) in the [remix](https://remix.ethereum.org)

Compile the code and in the deploy page, select `Custom  - External Http Provider` in the Environment dropdown.
The account of the developper account should be automatically selected.
Finally click the `deploy` button and take note of the contract address as you will need it later on.

## Run Test
To run the test, you have to provide some argument.
You will need to provide (at least) the two following arguments:
- `contractAddr`: the address of the contract
- `signer`: the account of the developper account.

The final command to launch the test is the following:
```bash
go test --contractAddr $CONTRACT_ADDR --signer $DEV_ACCOUNT
```
