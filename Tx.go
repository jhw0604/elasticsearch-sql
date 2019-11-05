package elasticsearch

import "fmt"

//Tx is for transaction
type Tx struct{}

//Commit will be always generate error
//Because Elasticsearch does not support transaction
func (*Tx) Commit() error {
	if debug {
		fmt.Println("Tx Commit")
	}
	return ErrTx
}

//Rollback will be always generate error
//Because Elasticsearch does not support transaction
func (*Tx) Rollback() error {
	if debug {
		fmt.Println("Tx Rollback")
	}
	return ErrTx
}
