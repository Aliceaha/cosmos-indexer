package models

type EventAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Event struct {
	Type       string         `json:"type"`
	Attributes EventAttribute `json:"attributes"`
}

type Logs struct {
	MsgIndex int    `json:"msg_index"`
	Log      string `json:"log"`
	Events   Event  `json:"events"`
}

type Messages struct {
	Type    string `json:"@type"`
	Grantee string `json:"grantee"`
	Msgs    []struct {
		Type             string `json:"@type"`
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Amount           struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"amount"`
	} `json:"msgs"`
}

type Tx struct {
	Tx struct {
		Body struct {
			Messages                    []Messages    `json:"messages"`
			Memo                        string        `json:"memo"`
			TimeoutHeight               string        `json:"timeout_height"`
			ExtensionOptions            []interface{} `json:"extension_options"`
			NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
		} `json:"body"`
		AuthInfo struct {
			SignerInfos []struct {
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
			} `json:"signer_infos"`
			Fee struct {
				Amount []struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"amount"`
				GasLimit string `json:"gas_limit"`
				Payer    string `json:"payer"`
				Granter  string `json:"granter"`
			} `json:"fee"`
		} `json:"auth_info"`
		Signatures []string `json:"signatures"`
	} `json:"tx"`
	TxResponse struct {
		Height    string `json:"height"`
		TxHash    string `json:"txhash"`
		Codespace string `json:"codespace"`
		Code      int    `json:"code"`
		Data      string `json:"data"`
		GasUsed   string `json:"gas_used"`
		GasWanted string `json:"gas_wanted"`
		Timestamp string `json:"timestamp"`
		RawLog    string `json:"raw_log"`
		Logs      []Logs `json:"logs"`
	} `json:"tx_response"`
}
