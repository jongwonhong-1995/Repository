// ExpressJS Setup
const path = require('path');
const express = require('express');
const app = express();
var bodyParser = require('body-parser');

// Hyperledger Bridge Setup
const { Wallets, Gateway } = require('fabric-network');
const fs = require('fs');

// load the network configuration
const ccpPath = path.resolve('/home/apstudent', 'fabric-samples', 'test-network', 'organizations', 'peerOrganizations', 'company.example.com', 'connection-company.json');
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));


// Constants
const PORT = 10010;
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
app.get('/page/createItem', function(req, res){
    res.sendFile(__dirname + '/viewpage/createItem.html');
})
app.get('/page/createCompany', function(req, res){
    res.sendFile(__dirname + '/viewpage/createCompany.html');
})
app.get('/page/createReceivable', function(req, res){
    res.sendFile(__dirname + '/viewpage/createReceivable.html');
})
app.get('/page/queryAllCars', function(req, res){    
    res.sendFile(__dirname + '/viewpage/queryAllCars.html');
})
app.get('/page/queryCar', function(req, res){    
    res.sendFile(__dirname + '/viewpage/queryCar.html');
})
app.get('/page/createCar', function(req, res){    
    res.sendFile(__dirname + '/viewpage/createCar.html');
})
app.get('/page/changeCarOwner', function(req, res){    
    res.sendFile(__dirname + '/viewpage/changeCarOwner.html');
})

// api routing
app.post('/api/createItem', async function(req, res){
    const itemno = req.body.itemno
    const itemname = req.body.itemname
    const itemseller = req.body.itemseller
    const itemprice = req.body.itemprice

    var args = [itemno,itemname,itemseller,itemprice]
    await callChainCode('createSellItem', args)
    res.status(200).json({result: "success"})
})
app.post('/api/createCompany', async function(req, res){
    const companyno = req.body.companyno
    const companyname = req.body.companyname
    const companytoken = req.body.companytoken
    const orgclass = req.body.orgclass

    var args = [companyno, companyname, companytoken, orgclass]
    await callChainCode('createEnterOrg',args)
    res.status(200).json({result: "success"})
})
app.post('/api/createReceivable', async function(req, res){
    const recvnumber = req.body.recvnumber
    const reptnumber = req.body.reptnumber
    const ownername = req.body.ownername
    const havedlist = JSON.stringify([])
    const issuerate = req.body.issuerate
    const publishdate = req.body.publishdate
    const expiredate = req.body.expiredate
    const isguarantee = req.body.isguarantee
    const issale = req.body.issale
    var args = [recvnumber, reptnumber, ownername, havedlist, issuerate, publishdate, expiredate, isguarantee, issale]
    await callChainCode('createReceivable',args)
    res.status(200).json({result: "success"})
})
app.post('/api/queryBuyerRept', async function(req,res){
    const buyerno = req.body.buyerno
    const result = await callChainCode('queryBuyItems', buyerno)
    res.json(JSON.parse(result))
})
app.post('/api/querySellerRept', async function(req,res){
    const sellerno = req.body.sellerno
    const result = await callChainCode('querySellItems', sellerno)
    res.json(JSON.parse(result))
})
app.get('/api/queryAllCars', async function(req, res){
    const result = await callChainCode('queryAllCars')        
    res.json(JSON.parse(result))    
})

app.post('/api/queryCar', async function(req, res){
    const carno=req.body.carno
    const result = await callChainCode('queryCar',carno)    
    res.json(JSON.parse(result))
})

app.post('/api/createCar', async function(req, res){
    const carno = req.body.carno
    const carmake = req.body.carmake
    const carmodel = req.body.carmodel
    const carcol = req.body.carcol
    const carowner = req.body.carowner

    var args = [carno,carmake,carmodel,carcol,carowner]    
    await callChainCode('createCar',args)    
    res.status(200).json({result: "success"})
})

app.post('/api/changeCarOwner', async function(req, res){
    const carno = req.body.carno
    const carowner = req.body.carowner

    var args = [carno,carowner]    
    await callChainCode('changeCarOwner',args)    
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
    if(fnName == 'createSellItem')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3]);
    else if(fnName == 'createEnterOrg')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3]);
    else if(fnName == 'createReceivable')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8]);
    else if(fnName == 'queryBuyItems')
        result = await contract.submitTransaction(fnName,args);
    else if(fnName == 'querySellItems')
        result = await contract.submitTransaction(fnName,args);
    else if(fnName == 'queryAllCars')
        result = await contract.evaluateTransaction(fnName);    
    else if(fnName == 'queryCar')
        result = await contract.evaluateTransaction(fnName,args);
    else if(fnName == 'createCar')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3],args[4])
    else if(fnName == 'changeCarOwner')
        result = await contract.submitTransaction(fnName,args[0],args[1])
    else
        result = 'This function(' + fnName +') does not exist !'        
        
    console.log(`Transaction has been evaluated, result is: ${result.toString()}`);

    // Disconnect from the gateway.
    await gateway.disconnect();
    
    return result;
}
