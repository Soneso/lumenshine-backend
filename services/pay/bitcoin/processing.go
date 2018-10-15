package bitcoin

import (
	"math/big"
	"strings"
	"time"

	m "github.com/Soneso/lumenshine-backend/services/db/models"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/support/errors"
)

func (l *Channel) processBlocks(blockNumber uint64) {
	l.log.Infof("Starting from block %d", blockNumber)

	// Time when last new block has been seen
	lastBlockSeen := time.Now()
	missingBlockWarningLogged := false

	for {
		block, err := l.getBlock(blockNumber)
		if err != nil {
			l.log.WithFields(logrus.Fields{"err": err, "blockNumber": blockNumber}).Error("Error getting block")
			time.Sleep(time.Second)
			continue
		}

		// Block doesn't exist yet
		if block == nil {
			if time.Since(lastBlockSeen) > 20*time.Minute && !missingBlockWarningLogged {
				l.log.Warn("No new block in more than 20 minutes")
				missingBlockWarningLogged = true
			}

			time.Sleep(time.Second)
			continue
		}

		// Reset counter when new block appears
		lastBlockSeen = time.Now()
		missingBlockWarningLogged = false

		err = l.processBlock(block)
		if err != nil {
			l.log.WithFields(logrus.Fields{"err": err, "blockHash": block.Header.BlockHash().String()}).Error("Error processing block")
			time.Sleep(time.Second)
			continue
		}

		// Persist block number
		err = l.db.SaveLastProcessedBitcoinBlock(blockNumber)
		if err != nil {
			l.log.WithField("err", err).Error("Error saving last processed block")
			time.Sleep(time.Second)
			// We continue to the next block.
			// The idea behind this is if there was a problem with this single query we want to
			// continue processing because it's safe to reprocess blocks and we don't want a downtime.
		}

		blockNumber++
	}
}

// getBlock returns (nil, nil) if block has not been found (not exists yet)
func (l *Channel) getBlock(blockNumber uint64) (*wire.MsgBlock, error) {
	blockHeight := int64(blockNumber)
	blockHash, err := l.client.GetBlockHash(blockHeight)
	if err != nil {
		if strings.Contains(err.Error(), "Block height out of range") {
			// Block does not exist yet
			return nil, nil
		}
		err = errors.Wrap(err, "Error getting block hash from bitcoin-core")
		l.log.WithField("blockHeight", blockHeight).Error(err)
		return nil, err
	}

	block, err := l.client.GetBlock(blockHash)
	if err != nil {
		err = errors.Wrap(err, "Error getting block from bitcoin-core")
		l.log.WithField("blockHash", blockHash.String()).Error(err)
		return nil, err
	}

	return block, nil
}

func (l *Channel) processBlock(block *wire.MsgBlock) error {
	transactions := block.Transactions

	localLog := l.log.WithFields(logrus.Fields{
		"blockHash":    block.Header.BlockHash().String(),
		"blockTime":    block.Header.Timestamp,
		"transactions": len(transactions),
	})
	localLog.Info("Processing block")

	for _, transaction := range transactions {
		transactionLog := localLog.WithField("transactionHash", transaction.TxHash().String())
		//transaction.TxIn[0].SignatureScript
		for index, output := range transaction.TxOut {
			class, addresses, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, l.chainParams)
			if err != nil {
				// txscript.ExtractPkScriptAddrs returns error on non-standard scripts
				// so this can be Warn.
				transactionLog.WithField("err", err).Warn("Error extracting addresses")
				continue
			}

			// We only support P2PK and P2PKH addresses
			if class != txscript.PubKeyTy && class != txscript.PubKeyHashTy {
				transactionLog.WithField("class", class).Debug("Unsupported addresses class")
				continue
			}

			// Paranoid. We access address[0] later.
			if len(addresses) != 1 {
				transactionLog.WithField("addresses", addresses).Error("Invalid addresses length")
				continue
			}

			err = l.processTransaction(
				transaction.TxHash().String(),
				index,
				big.NewInt(output.Value),
				addresses[0].EncodeAddress(),
				"",
			)
			if err != nil {
				return errors.Wrap(err, "Error processing transaction")
			}
		}
	}

	localLog.Info("Processed block")

	return nil
}

func (l *Channel) processTransaction(hash string, txOutIndex int, valueSat *big.Int, toAddress string, fromAddress string) error {
	localLog := l.log.WithFields(logrus.Fields{"transaction": hash, "index": txOutIndex, "rail": "bitcoin"})
	localLog.Debug("Processing transaction")

	//get the order from the database
	order, err := l.db.GetOrderForAddress(m.PaymentNetworkBitcoin, toAddress, "")
	if err != nil {
		return errors.Wrap(err, "Error getting association")
	}

	if order == nil {
		localLog.Debug("Associated address not found, skipping")
		return nil
	}

	// Add transaction as processing.
	isDuplicate, err := l.db.AddNewTransaction(l.log, l, hash, toAddress, fromAddress, order.ID, valueSat, txOutIndex)
	if err != nil {
		return err
	}

	if isDuplicate {
		localLog.Debug("Transaction already processed, skipping")
		return nil
	}

	return nil
}
