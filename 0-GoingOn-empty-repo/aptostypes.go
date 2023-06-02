package main

type AptosTransaction struct {
	Txn struct {
		Type                     string      `json:"type"`
		Sender                   string      `json:"sender"`
		Version                  string      `json:"version"`
		Hash                     string      `json:"hash"`
		StateChangeHash          string      `json:"state_change_hash"`
		EventRootHash            string      `json:"event_root_hash"`
		StateCheckpointHash      interface{} `json:"state_checkpoint_hash"`
		GasUsed                  string      `json:"gas_used"`
		Success                  bool        `json:"success"`
		VMStatus                 string      `json:"vm_status"`
		AccumulatorRootHash      string      `json:"accumulator_root_hash"`
		Changes                  []Change    `json:"changes"`
		ID                       string      `json:"id"`
		Epoch                    string      `json:"epoch"`
		Round                    string      `json:"round"`
		Events                   []Event     `json:"events"`
		PreviousBlockVotesBitvec []int       `json:"previous_block_votes_bitvec"`
		Proposer                 string      `json:"proposer"`
		FailedProposerIndices    []int       `json:"failed_proposer_indices"`
		Timestamp                string      `json:"timestamp"`
	} `json:"txn"`
	BlockNum int `json:"block_num"`
}

type Change struct {
	Type         string `json:"type"`
	Address      string `json:"address"`
	StateKeyHash string `json:"state_key_hash"`
	Data         struct {
		Type string `json:"type"`
		Data struct {
			EpochInterval             string                   `json:"epoch_interval"`
			Height                    string                   `json:"height"`
			NewBlockEvents            NewBlockEvent            `json:"new_block_events"`
			UpdateEpochIntervalEvents UpdateEpochIntervalEvent `json:"update_epoch_interval_events"`
		} `json:"data"`
	} `json:"data"`
}

type NewBlockEvent struct {
	Counter string `json:"counter"`
	GUID    struct {
		ID struct {
			Addr        string `json:"addr"`
			CreationNum string `json:"creation_num"`
		} `json:"id"`
	} `json:"guid"`
}

type UpdateEpochIntervalEvent struct {
	Counter string `json:"counter"`
	GUID    struct {
		ID struct {
			Addr        string `json:"addr"`
			CreationNum string `json:"creation_num"`
		} `json:"id"`
	} `json:"guid"`
}

type Event struct {
	GUID struct {
		CreationNumber string `json:"creation_number"`
		AccountAddress string `json:"account_address"`
	} `json:"guid"`
	SequenceNumber string `json:"sequence_number"`
	Type           string `json:"type"`
	Data           struct {
		Epoch                    string   `json:"epoch"`
		FailedProposerIndices    []string `json:"failed_proposer_indices"`
		Hash                     string   `json:"hash"`
		Height                   string   `json:"height"`
		PreviousBlockVotesBitvec string   `json:"previous_block_votes_bitvec"`
		Proposer                 string   `json:"proposer"`
		Round                    string   `json:"round"`
		TimeMicroseconds         string   `json:"time_microseconds"`
		Amount                   string   `json:"amount"`
	} `json:"data"`
}

type FilteredMessage struct {
	Address           string `json:"address"`
	TxHash            string `json:"txhash"`
	Code              uint32 `json:"code"`
	TokenOutDenom     string `json:"token_out_denom"`
	TokenOutDenomName string `json:"token_out_denom_name"`
	TokenOutAmount    string `json:"token_out_amount"`
	TokenInDenom      string `json:"token_in_denom"`
	TokenInDenomName  string `json:"token_in_denom_name"`
	TokenInAmount     string `json:"token_in_amount"`
	Timestamp         int    `json:"timestamp"`
}
