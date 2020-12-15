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
const PORT = 11011;
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
app.get('/page/queryBuyItems', function(req, res){
    res.sendFile(__dirname + '/viewpage/queryBuyItems.html');
})
app.get('/page/createBuyer', function(req, res){
    res.sendFile(__dirname + '/viewpage/createBuyer.html');
})
app.get('/page/createItemBuy', function(req, res){
    res.sendFile(__dirname + '/viewpage/createItemBuy.html');
})

// api routing
app.get('/api/queryAllItems', async function(req, res){
    const result = await callChainCode('queryAllItems')
    res.json(JSON.parse(result))
})

app.post('/api/queryBuyItems', async function(req, res){
    const buyername=req.body.buyername
    const result = await callChainCode('queryBuyItems', buyername)
    res.json(JSON.parse(result))
})

app.post('/api/createSellReceipt', async function(req, res){
    const reptno = req.body.reptno
    const itemno = req.body.itemno
    const sellername = req.body.sellername
    const buyername = req.body.buyername
    const numproduct = req.body.numproduct
    const totalprice = req.body.totalprice
    const selldate = req.body.selldate
    const duedate = req.body.duedate

    var args = [reptno, itemno, sellername, buyername, numproduct, totalprice, selldate, duedate]
    await callChainCode('createSellReceipt', args)
    await callChainCode('updateTokenTransfer', args)
    res.status(200).json({result: "success"})
})

app.post('/api/querySellItemPrice', async function(req, res){
    const itemno=req.body.itemno
    const result = await callChainCode('querySellItem', itemno)
    res.json(JSON.parse(result))
})

app.post('/api/createBuyer', async function(req, res){
    const buyerno = req.body.buyerno
    const buyername = req.body.buyername
    const buyertoken = req.body.buyertoken
    const orgclass = req.body.orgclass

    var args = [buyerno, buyername, buyertoken, orgclass]
    await callChainCode('createEnterOrg', args)
    res.status(200).json({result: "success"})
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
    else if (fnName == 'queryBuyItems')
        result = await contract.submitTransaction(fnName, args);
    else if (fnName == 'createEnterOrg')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3]);
    else if (fnName == 'querySellItem')
        result = await contract.submitTransaction(fnName, args);
    else if (fnName == 'createSellReceipt')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7])
    else if (fnName == 'updateTokenTransfer')
        result = await contract.submitTransaction(fnName,args[2],args[3],args[5])
    else
        result = 'This function(' + fnName +') does not exist !'        
        
    console.log(`Transaction has been evaluated, result is: ${result.toString()}`);

    // Disconnect from the gateway.
    await gateway.disconnect();
    
    return result;
}
