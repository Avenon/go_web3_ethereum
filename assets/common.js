var MAINET_RPC_URL = 'https://mainnet.infura.io/metamask' ;
var ROPSTEN_RPC_URL = 'https://ropsten.infura.io/metamask' ;
var KOVAN_RPC_URL = 'https://kovan.infura.io/metamask' ;
var RINKEBY_RPC_URL = 'https://rinkeby.infura.io/metamask' ;

var CURRENT_URL = RINKEBY_RPC_URL;

$( document ).ready(function() {
    $(".sendButton").click(function(){
        let Web3 = require('./web3.min.js');
        console.log(Web3);
        if (typeof web3 !== 'undefined'){
            web3 = new Web3(web3.currentProvider);
        }
        else {
            alert('You have to install MetaMask !');
        }
        
        const contractAddress = "0x038B24CeE212c05dbcBbA8A5a8E721ACC21A8802";
        web3.eth.sendTransaction({
            to: contractAddress,
            from: web3.eth.accounts[0],
            value: web3.toWei($('.inputEthers').val(), "ether")
        },
            function (error) {
                console.log(error);
            }
        );
    });
});
