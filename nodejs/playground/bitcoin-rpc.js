// wallet address
TEST_ADDRESS = 'bcrt1qyvnsr3l4wuuen4kznamtjk8zpfjpmfyd0v4t9h'

// kapitalize bitcoin client
const client = require('kapitalize')();
client
    .auth('fufeng', '123456')
    .set('host', '192.168.3.131')
    .set({
        port: 18443,
        protocol: 'http'
    });

// check network information
client.getNetworkInfo(function (err, info) {
    if (err) {
        console.log(err);
    } else {
        console.log(info);
    }
});

// get unspent tx
client.listUnspent(function (err, txs) {
    if (err) {
        console.log(err);
    } else {
        console.log(txs);
    }
});

// send 1 btc to test address
client.sendToAddress(TEST_ADDRESS, 1);

// get unspent tx again
client.listUnspent(function (err, txs) {
    if (err) {
        console.log(err);
    } else {
        console.log(txs);
    }
});

