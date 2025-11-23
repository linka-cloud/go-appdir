package appdir

import "os/user"

func providedUID(uid []string) (string, bool) {
	if len(uid) == 0 {
		return "", false
	}
	for _, v := range uid {
		if v != "" {
			return v, true
		}
	}
	return "", false
}

func homeDirFor(uid []string) string {
	val, ok := providedUID(uid)
	if !ok {
		return ""
	}
	u, err := user.LookupId(val)
	if err != nil {
		return ""
	}
	return u.HomeDir
}
