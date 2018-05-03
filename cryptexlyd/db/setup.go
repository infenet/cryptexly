/*
 * Cryptexly - Copyleft of Xavier D. Johnson.
 * me at xavierdjohnson dot com
 * http://www.xavierdjohnson.net/
 *
 * See LICENSE.
 */
package db

import (
	"github.com/detroitcybersec/cryptexly/cryptexlyd/config"
	"github.com/detroitcybersec/cryptexly/cryptexlyd/log"
	"github.com/detroitcybersec/cryptexly/cryptexlyd/utils"
	"os"
	"time"
)

var (
	dbIndex  = (*Index)(nil)
	dbNextId = uint64(0)

	Size = uint64(0)
)

func Setup() (created bool, err error) {
	started := time.Now()

	if config.Conf.Database, err = utils.ExpandPath(config.Conf.Database); err != nil {
		return false, err
	}

	if _, err = os.Stat(config.Conf.Database); os.IsNotExist(err) {
		created = true
		log.Warningf("Creating database %s ...", log.Bold(config.Conf.Database))
	} else {
		created = false
		log.Infof("Loading database %s ...", log.Bold(config.Conf.Database))
	}

	dbIndex, err = LoadIndex(config.Conf.Database)
	if err != nil {
		return false, err
	}

	for id := range dbIndex.Records() {
		if id > dbNextId {
			dbNextId = id + 1
		}
	}

	elapsed := time.Since(started)

	log.Debugf("  dbNextId=%d", dbNextId)
	log.Infof("%s of records loaded in %s.", utils.FormatBytes(Size), elapsed)

	return created, nil
}

func Lock() {
	dbIndex.Lock()
}

func Unlock() {
	dbIndex.Unlock()
}

func GetStores() map[uint64]*Record {
	return dbIndex.Records()
}
