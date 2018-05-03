/*
 * Cryptexly - Copyleft of Xavier D. Johnson.
 * me at xavierdjohnson dot com
 * http://www.xavierdjohnson.net/
 *
 * See LICENSE.
 */
package utils

// InSlice return if a is in list
func InSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
