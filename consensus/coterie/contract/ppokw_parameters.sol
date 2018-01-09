pragma solidity ^0.4.18;

contract PpokwParameters {
    uint32 public committeeSize;

    struct DifficultyAdjustment {
        uint difficulty;
        uint blockNumber;
    }

    DifficultyAdjustment public difficulty;

    event CommitteeSizeChanged(uint32 sizeBefore, uint32 sizeAfter);
    event DifficultyAdjusted(uint previousDifficulty, uint newDifficulty, uint blockNumber, address changedBy);

    modifier onlyValidDifficulty(uint newDifficulty) {
        require(newDifficulty > 0);
        _;
    }

    function setCommitteeSize(uint32 newCommitteeSize) external {
        CommitteeSizeChanged(committeeSize, newCommitteeSize);
        committeeSize = newCommitteeSize;
    }

    function setDifficulty(uint newDifficulty) external onlyValidDifficulty(newDifficulty) {
        DifficultyAdjusted(block.difficulty, newDifficulty, block.number, msg.sender);
        difficulty = DifficultyAdjustment(newDifficulty, block.number);
    }
}