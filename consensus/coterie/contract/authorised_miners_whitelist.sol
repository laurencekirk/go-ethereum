pragma solidity ^0.4.18;

contract AuthorisedMinersWhitelist {

    mapping(address => bool) whitelist;
    uint32 public size;

    event AddedToWhitelist(address miner);
    event RemovedFromWhitelist(address miner);

    function isAuthorisedMiner(address miner) public view returns (bool) {
        return whitelist[miner];
    }

    function authoriseMiner(address miner) public {
        require(! whitelist[miner]);
        whitelist[miner] = true;
        size = size + 1;
        AddedToWhitelist(miner);
    }

    function removeMinersAuthorisation(address miner) public {
        require(whitelist[miner]);
        whitelist[miner] = false;
        size = size - 1;
        RemovedFromWhitelist(miner);
    }
}