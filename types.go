package main

type Subscribe struct {
	Subscribe string `json:"subscribe"`
	ChainID   string `json:"chain_id"`
}

type Response struct {
	ChainID   string    `json:"chain_id"`
	Type      string    `json:"type"`
	BlockData BlockData `json:"data"`
}

type BlockData struct {
	Block        Block            `json:"block"`
	ResultBegin  ResultBeginBlock `json:"result_begin_block"`
	ResultEnd    ResultEndBlock   `json:"result_end_block"`
	Transactions []Transaction    `json:"txs"`
	Supply       []Supply         `json:"supply"`
}

type Block struct {
	Header   Header      `json:"header"`
	Data     BlockTxData `json:"data"`
	Evidence struct {
		Evidence []string `json:"evidence"`
	} `json:"evidence"`
	LastCommit LastCommit `json:"last_commit"`
}

type Header struct {
	Version struct {
		BlockNumber string `json:"block"`
	} `json:"version"`
	ChainID     string `json:"chain_id"`
	Height      string `json:"height"`
	Time        string `json:"time"`
	LastBlockID struct {
		Hash  string `json:"hash"`
		Parts struct {
			Total int    `json:"total"`
			Hash  string `json:"hash"`
		} `json:"parts"`
	} `json:"last_block_id"`
	LastCommitHash     string `json:"last_commit_hash"`
	DataHash           string `json:"data_hash"`
	ValidatorsHash     string `json:"validators_hash"`
	NextValidatorsHash string `json:"next_validators_hash"`
	ConsensusHash      string `json:"consensus_hash"`
	AppHash            string `json:"app_hash"`
	LastResultsHash    string `json:"last_results_hash"`
	EvidenceHash       string `json:"evidence_hash"`
	ProposerAddress    string `json:"proposer_address"`
}

type BlockTxData struct {
	Transactions []string `json:"txs"`
}

type LastCommit struct {
	Height  string `json:"height"`
	Round   int    `json:"round"`
	BlockID struct {
		Hash  string `json:"hash"`
		Parts struct {
			Total int    `json:"total"`
			Hash  string `json:"hash"`
		} `json:"parts"`
	} `json:"block_id"`
	Signatures []Validators `json:"signatures"`
}

type Validators struct {
	BlockIDFlag      int    `json:"block_id_flag"`
	ValidatorAddress string `json:"validator_address"`
	TimeStamp        string `json:"timestamp"`
	Signature        string `json:"signature"`
}

type ResultBeginBlock struct {
	Events []Event `json:"events"`
}

type Event struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}

type ResultEndBlock struct {
	ValidatorUpdates      []string `json:"validator_updates"`
	ConsensusParamUpdates struct {
		Block struct {
			MaxBytes string `json:"max_bytes"`
			MaxGas   string `json:"max_gas"`
		} `json:"block"`
		Evidence struct {
			MaxAgeNumBlocks string `json:"max_age_num_blocks"`
			MaxAgeDuration  string `json:"max_age_duration"`
			MaxBytes        string `json:"max_bytes"`
		} `json:"evidence"`
		Validator struct {
			PubKeyTypes []string `json:"pub_key_types"`
		} `json:"validator,omitempty"`
	} `json:"consensus_param_updates"`
	Events []Event `json:"events"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Index bool   `json:"index"`
}

type Transaction struct {
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
	Height     string   `json:"height"`
	TxHash     string   `json:"txhash"`
	Codespace  string   `json:"codespace"`
	Code       int      `json:"code"`
	Data       string   `json:"data"`
	RawLog     string   `json:"raw_log"`
	Logs       []Log    `json:"logs"`
	Info       string   `json:"info"`
	GasWanted  string   `json:"gas_wanted"`
	GasUsed    string   `json:"gas_used"`
	Tx         TX       `json:"tx"`
	Timestamp  string   `json:"timestamp"`
	Events     []Event  `json:"events"`
}

type Body struct {
	Messages                    []Message `json:"messages"`
	Memo                        string    `json:"memo"`
	TimeoutHeight               string    `json:"timeout_height"`
	ExtensionOptions            []string  `json:"extension_options"`
	NonCriticalExtensionOptions []string  `json:"non_critical_extension_options"`
}

type Message struct {
	Type       string       `json:"@type"`
	Sender     string       `json:"sender"`
	Contract   string       `json:"contract"`
	ExecuteMsg interface{}  `json:"execute_msg"`
	Coins      []CoinAmount `json:"coins"`
}

/*type ExecuteMsg struct {
	BorrowAmount             interface{} `json:"borrow_stable,omitempty"`
	ClaimRewards             interface{} `json:"claim_rewards,omitempty"`
	FeedPrice                interface{} `json:"feed_price,omitempty"`
	DepositStable            interface{} `json:"deposit_stable,omitempty"`
	ProvideLiquidity         interface{} `json:"provide_liquidity,omitempty"`
	Swap                     interface{} `json:"swap,omitempty"`
	Relay                    interface{} `json:"relay,omitempty"`
	Withdraw                 interface{} `json:"withdraw,omitempty"`
	ClaimAndOptionallyUnlock interface{} `json:"claim_rewards_and_optionally_unlock,omitempty"`
}*/

type CoinAmount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type AuthInfo struct {
	SignerInfos []SignerInfo `json:"signer_infos"`
	Fee         Fee          `json:"fee"`
}

type SignerInfo struct {
	PublicKey struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	} `json:"public_key"`
	ModeInfo struct {
		Single struct {
			Mode string `json:"mode"`
		} `json:"single"`
	} `json:"mode_info"`
	Sequence string `json:"sequence"`
}

type Fee struct {
	Amount   []CoinAmount `json:"amount"`
	GasLimit string       `json:"gas_limit"`
	Payer    string       `json:"payer"`
	Granter  string       `json:"granter"`
}

type Log struct {
	MsgIndex int     `json:"msg_index"`
	Log      string  `json:"log"`
	Events   []Event `json:"events"`
}

type TX struct {
	Type       string   `json:"@type"`
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
}

type Supply struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Page struct {
	Title string
	Body  []byte
}
