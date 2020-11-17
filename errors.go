package mediainfo

import "errors"

// ErrNotLoaded is the error returned if function Load() was not called before any call
var ErrNotLoaded = errors.New("Loaded not called previously")
