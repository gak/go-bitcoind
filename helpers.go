package bitcoind

import (
	"errors"
	"fmt"
)

// handleError handle error returned by client.call
func handleError(err error, r *rpcResponse) error {
	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		return errors.New(fmt.Sprintf("(%v) %s", rr["code"].(float64), rr["message"].(string)))

	}
	if err != nil {
		return err
	}
	return nil
}
