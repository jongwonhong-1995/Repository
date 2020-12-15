// ExpressJS Setup
const path = require('path');
const express = require('express');
const app = express();
var bodyParser = require('body-parser');

// Hyperledger Bridge Setup
const { Wallets, Gateway } = require('fabric-network');
const fs = require('fs');

// load the network configuration
const ccpPath = path.resolve('/home/apstudent', 'fabric-samples', 'test-network', 'organizations', 'peerOrganizations', 'buyer.example.com', 'connection-buyer.json');
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));


// Constants
const PORT = 11111;
const HOST = '0.0.0.0';

// server start
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);

// use static file
app.use(express.static(path.join(__dirname)));

// configure app to use body-parser
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

// main page routing
app.get('/', function(req, res){
    res.sendFile(__dirname + '/viewpage/index.html');
})
app.get('/page/queryAllItems', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryAllItems.html');
})
app.get('/page/queryAllOrgs', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryAllOrgs.html');
})
app.get('/page/queryAllBuyers', function(req, res){    
    res.sendFile(__dirname + '/viewpage/queryAllBuyers.html');
})
app.get('/page/queryAllCompanys', function(req, res){    
    res.sendFile(__dirname + '/viewpage/queryAllCompanys.html');
})
app.get('/page/queryAllGrades', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryAllGrades.html');
})
app.get('/page/queryAllFinances', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryAllFinances.html')
})
app.get('/page/queryAllReceipts', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryAllReceipts.html');
})
app.get('/page/queryAllReceivables', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryAllReceivables.html');
})
app.get('/page/queryAllRecvRatings', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryAllRecvRatings.html');
})

// api routing
app.get('/api/queryAllItems', async function(req, res){
    const result = await callChainCode('queryAllItems')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllOrgs', async function(req, res){
    const result = await callChainCode('queryAllOrgs')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllCompanys', async function(req, res){
    const result = await callChainCode('queryAllCompanys')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllBuyers', async function(req, res){
    const result = await callChainCode('queryAllBuyers')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllGrades', async function(req, res){
    const result = await callChainCode('queryAllGrades')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllFinances', async function(req, res){
    const result = await callChainCode('queryAllFinances')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllReceipts', async function(req, res){
    const result = await callChainCode('queryAllReceipts')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllReceivables', async function(req, res){
    const result = await callChainCode('queryAllReceivables')
    res.json(JSON.parse(result))
})
app.get('/api/queryAllRecvRatings', async function(req, res){
    const result = await callChainCode('queryAllRecvRatings')
    res.json(JSON.parse(result))
})

async function callChainCode(fnName, args){
    
    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    var result;
    console.log(`Wallet path: ${walletPath}`);
    

    // Check to see if we've already enrolled the user.
    const identity = await wallet.get('appUser');
    if (!identity) {
        console.log('An identity for the user "appUser" does not exist in the wallet');
        console.log('Run the registerUser.js application before retrying');
        return;
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });
    
    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork('mychannel');

    // Get the contract from the network.
    const contract = network.getContract('fabar');

    // Evaluate the specified transaction.    
    if(fnName == 'queryAllItems')
        result = await contract.evaluateTransaction(fnName);
    else if(fnName == 'queryAllOrgs')
        result = await contract.evaluateTransaction(fnName);
    else if(fnName == 'queryAllCompanys')
        result = await contract.evaluateTransaction(fnName);    
    else if(fnName == 'queryAllBuyers')
        result = await contract.evaluateTransaction(fnName);    
    else if(fnName == 'queryAllGrades')
        result = await contract.evaluateTransaction(fnName);
    else if(fnName == 'queryAllFinances')
        result = await contract.evaluateTransaction(fnName);
    else if(fnName == 'queryAllReceipts')
        result = await contract.evaluateTransaction(fnName);
    else if(fnName == 'queryAllCars')
        result = await contract.evaluateTransaction(fnName);    
    else if(fnName == 'queryAllReceivables')
        result = await contract.evaluateTransaction(fnName);
    else if(fnName == 'queryAllRecvRatings')
        result = await contract.evaluateTransaction(fnName);
    else
        result = 'This function(' + fnName +') does not exist !'        
        
    console.log(`Transaction has been evaluated, result is: ${result.toString()}`);

    // Disconnect from the gateway.
    await gateway.disconnect();
    
    return result;
}
result