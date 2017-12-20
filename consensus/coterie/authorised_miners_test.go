package coterie

// Since currently we're exporting the hardcoded public keys; perform a sanity test that a public key known to be in the
// whitelist is considered valid.
/*
func TestCanValidateMinerIsInWhitelist(t *testing.T) {
	pubKey := &ValidMiner1
	auth := IsMinerInWhitelist(pubKey)
	if !auth {
		t.Error("Expected that the public key would be in the whitelist")
	}
}


// Check to make sure that a matching public key - same X and Y coordinates on the expected curve are considered valid
// even though they are a pointer to a different object.
func TestCanValidateMatchingPublicKeyIsInWhitelist(t *testing.T) {
	pubKey := &MatchingMinerInWhitelist
	auth := IsMinerInWhitelist(pubKey)
	if !auth {
		t.Error("Expected that the public key matching a whitelisted miner would be considered to be in the whitelist")
	}
}

func TestCanNotValidateMinerThatIsNotInWhitelist(t *testing.T) {
	auth := IsMinerInWhitelist(&MinerNotInWhitelist1)
	if auth {
		t.Error("Expected that the public key (MinerNotInWhitelist1) would *not* be in the whitelist")
	}

	auth = IsMinerInWhitelist(&MinerNotInWhitelist2)
	if auth {
		t.Error("Expected that the public key (MinerNotInWhitelist2) would *not* be in the whitelist")
	}

	auth = IsMinerInWhitelist(&MinerNotInWhitelist3)
	if auth {
		t.Error("Expected that the public key (MinerNotInWhitelist3) would *not* be in the whitelist")
	}

	auth = IsMinerInWhitelist(&MinerNotInWhitelist4)
	if auth {
		t.Error("Expected that the public key (MinerNotInWhitelist4) would *not* be in the whitelist")
	}
}

func TestCanNotValidateNilMiner(t *testing.T) {
	auth := IsMinerInWhitelist(nil)
	if auth {
		t.Error("Expected that the public key would *not* be in the whitelist")
	}
}*/
