package blockchain

//Version int32
//LockTime int32
type Transaction struct {
	TxIn  []TxInput
	TxOut []TxOutput
}

// TODO: Implements me with test case
func NewTransaction() *Transaction {
	return nil
}

// TODO: Implements me with test case
func (tx *Transaction) Hash() []byte {
	return nil
}

// TODO: Implements me with test case
func NewCoinbase() *Transaction {
	return nil
}

// TODO: Implements me with test case
func (tx Transaction) IsCoinbase() bool {
	return false
}
