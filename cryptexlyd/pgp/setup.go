/*
 * Cryptexly - Copyleft of Xavier D. Johnson.
 * me at xavierdjohnson dot com
 * http://www.xavierdjohnson.net/
 *
 * See LICENSE.
 */
package pgp

import (
	_ "crypto/sha256"
	"github.com/detroitcybersec/cryptexly/cryptexlyd/config"
	"github.com/detroitcybersec/cryptexly/cryptexlyd/utils"
	_ "golang.org/x/crypto/ripemd160"
	"os"
	"path"
)

func Setup(pgp *config.PGPConfig) error {
	pgp.Keys.Public, _ = utils.ExpandPath(pgp.Keys.Public)
	if err := LoadKey(pgp.Keys.Public, false); err != nil {
		return err
	}

	if pgp.Keys.Private == "" {
		cwd, _ := os.Getwd()
		pgp.Keys.Private = path.Join(cwd, "cryptexlyd-pgp-private.key")
	}

	pgp.Keys.Private, _ = utils.ExpandPath(pgp.Keys.Private)
	public := path.Join(path.Dir(pgp.Keys.Private), "cryptexlyd-pgp-public.key")
	if utils.Exists(pgp.Keys.Private) == false {
		if err := GenerateKeys(pgp.Keys.Private, public); err != nil {
			return err
		}
	}

	return LoadKey(pgp.Keys.Private, true)
}
