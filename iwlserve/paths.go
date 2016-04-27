package main

import "path"

// authorNameToURLSafeName transforms author name to
// URL-friendly format, for example
// H. P. Lovecraft -> H._P._Lovecraft.
func authorNameToURLSafeName(name string) string {
	s := []rune(name)
	for i, r := range s {
		if (r < 'A' || r > 'z') && r != '.' {
			s[i] = '_'
		}
	}
	return string(s)
}

func getBadgeURL(base string, author string) string {
	return base + path.Join("/", "b", crcByAuthor[author])
}

func getShareURL(base string, author string) string {
	return base + path.Join("/", "s", crcByAuthor[author])
}

func getShortAuthorURL(base string, author string) string {
	return base + path.Join("/", "w", crcByAuthor[author])
}

func getAuthorURL(base string, author string) string {
	return base + path.Join("/", "writer", authorNameToURLSafeName(author))
}
