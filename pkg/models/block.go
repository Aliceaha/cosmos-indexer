package models

type BlockID struct {
	Hash  string `json:"hash"`
	Parts struct {
		Total int    `json:"total"`
		Hash  string `json:"hash"`
	} `json:"parts"`
}

type Header struct {
	Version struct {
		Block string `json:"block"`
	} `json:"version"`
	ChainID            string  `json:"chain_id"`
	Height             string  `json:"height"`
	Time               string  `json:"time"`
	LastBlockID        BlockID `json:"last_block_id"`
	LastCommitHash     string  `json:"last_commit_hash"`
	DataHash           string  `json:"data_hash"`
	ValidatorsHash     string  `json:"validators_hash"`
	NextValidatorsHash string  `json:"next_validators_hash"`
	ConsensusHash      string  `json:"consensus_hash"`
	AppHash            string  `json:"app_hash"`
	LastResultsHash    string  `json:"last_results_hash"`
	EvidenceHash       string  `json:"evidence_hash"`
	ProposerAddress    string  `json:"proposer_address"`
}

type BlockData struct {
	Txs []string `json:"txs"`
}

type CommitBlockID struct {
	Hash  string `json:"hash"`
	Parts struct {
		Total int    `json:"total"`
		Hash  string `json:"hash"`
	} `json:"parts"`
}

type Signature struct {
	BlockIDFlag      int    `json:"block_id_flag"`
	ValidatorAddress string `json:"validator_address"`
	Timestamp        string `json:"timestamp"`
	Signature        string `json:"signature"`
}

type LastCommit struct {
	Height     string        `json:"height"`
	Round      int           `json:"round"`
	BlockID    CommitBlockID `json:"block_id"`
	Signatures Signature     `json:"signatures"`
}

type BlockInfo struct {
	Header   Header    `json:"header"`
	Data     BlockData `json:"data"`
	Evidence struct {
		Evidence interface{} `json:"evidence"`
	} `json:"evidence"`
	LastCommit LastCommit `json:"last_commit"`
}

type BlockResponse struct {
	BlockID BlockID   `json:"block_id"`
	Block   BlockInfo `json:"block"`
}
