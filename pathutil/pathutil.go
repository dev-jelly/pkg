// Package pathutil implements path utility functions.
package pathutil

import (
	"fmt"
	"path"
	"strings"
)

// TrimExt returns a slice of the string filePath without the extension.
func TrimExt(filePath string) string {
	ext := path.Ext(filePath)
	return filePath[:len(filePath)-len(ext)]
}

// Base corresponds to a base directory.
type Base string

// Path returns a sanitized concatenation of base and relPath. The
// implementation takes extra precausions to avoid directory traversal
// vulnerabilities.
func (base Base) Path(relPath string) (string, error) {
	// Join joins the path elements and cleans the result p.
	p := path.Join(string(base), relPath)
	// If relPath contains directory traversal characters such as "../" p
	// could have escaped base by now.
	if !strings.HasPrefix(p, string(base)) {
		// Prevent directory traversal.
		return "", fmt.Errorf("Base.Path: path %q doesn't contain the prefix %q.", p, base)
	}
	return p, nil
}
