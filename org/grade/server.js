// ExpressJS Setup
const path = require('path');
const express = require('express');
const app = express();
var bodyParser = require('body-parser');

// Hyperledger Bridge Setup
const { Wallets, Gateway } = require('fabric-network');
const fs = require('fs');

// load the network configuration
const ccpPath = path.resolve('/home/apstudent', 'fabric-samples', 'test-network', 'organizations', 'peerOrganizations', 'grade.example.com', 'connection-grade.json');
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));


// Constants
const PORT = 10002;
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
app.get('/page/queryAllGrades', function(req, res){    
    res.sendFile(__dirname + '/viewpage/queryAllGrades.html');
})
app.get('/page/checkRecv', function(req, res){    
    res.sendFile(__dirname + '/viewpage/checkRecv.html');
})
app.get('/page/evalRecv', function(req, res){    
    res.sendFile(__dirname + '/viewpage/evalRecv.html');
})
app.get('/page/createGrade', function(req, res){    
    res.sendFile(__dirname + '/viewpage/createGrade.html');
})

// api routing
app.get('/api/queryAllGrades', async function(req, res){
    const result = await callChainCode('queryAllGrades')        
    res.json(JSON.parse(result))    
})

app.get('/api/checkRecv', async function(req, res){
    const result = await callChainCode('checkRecv')        
    res.json(JSON.parse(result))    
})



app.post('/api/createGrade', async function(req, res){
    const gradeno = req.body.gradeno
    const gradename = req.body.gradename
    const gradetoken = req.body.gradetoken
    const orgclass = req.body.orgclass

    console.log(gradeno + ' ' + gradename + ' ' + gradetoken + ' ' + orgclass)
    var args = [gradeno,gradename,gradetoken,orgclass]    
    await callChainCode('createEnterOrg',args)    
    res.status(200).json({result: "success"})
})

app.post('/api/evaluateRec', async function(req, res){
    const gradeno = req.body.gradeno
    const recvnumber = req.body.recvnumber
    const gradeEval = req.body.gradeEval

    console.log(gradeno + ' ' + recvnumber + ' ' + gradeEval)
    var args = [gradeno,recvnumber,gradeEval]    
    await callChainCode('evaluateRec',args)    
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
    if(fnName == 'queryAllGrades')
        result = await contract.evaluateTransaction(fnName);    
    else if(fnName == 'createEnterOrg')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3])
    else if(fnName == 'evalRecv')
        result = await contract.submitTransaction(fnNam)
    else if(fnName == 'checkRecv')
        result = await contract.submitTransaction(fnName)
    else
        result = 'This function(' + fnName +') does not exist !'        
        
    console.log(`Transaction has been evaluated, result is: ${result.toString()}`);

    // Disconnect from the gateway.
    await gateway.disconnect();
    
    return result;
}
