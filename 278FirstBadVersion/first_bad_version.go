package firstbadversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var isBadVersion func(version int) bool

func firstBadVersion(n int) int {
	var beg, end = 1, n

	for {
		k := (beg + end) / 2

		if isBadVersion(k) {
			end = k
		} else {
			beg = k
		}

		if end-beg <= 1 {
			if isBadVersion(beg) && isBadVersion(end) {
				return beg
			}
			return end
		}
	}
}
