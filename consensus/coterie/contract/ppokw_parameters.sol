pragma solidity ^0.4.18;

contract PpokwParameters {
    uint32 public committeeSize;

    event CommitteeSizeChanged(uint32 sizeBefore, uint32 sizeAfter);

    function setCommitteeSize(uint32 newCommitteeSize) external {
        CommitteeSizeChanged(committeeSize, newCommitteeSize);
        committeeSize = newCommitteeSize;
    }
}