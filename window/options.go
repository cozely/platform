package window

import "errors"

type Option func(*Window) error

func Title(s string) Option {
	return func(w *Window) error {
		if w.opened {
			//TODO: implement
			return errors.New("window.Title: not implemented for opened windows")
		}
		w.title = s
		return nil
	}
}

func Size(x, y int32) Option {
	return func(w *Window) error {
		if w.opened {
			//TODO: implement
			return errors.New("window.Size: not implemented for opened windows")
		}
		w.size = Coord{x, y}
		return nil
	}
}

func Fullscreen(fullscreen bool, windowed bool) Option {
	return func(w *Window) error {
		if w.opened {
			//TODO: implement
			return errors.New("window.Fullscreen: not implemented for opened windows")
		}
		w.fullscreen = fullscreen
		w.desktop = !windowed
		return nil
	}
}

func Monitor(n int) Option {
	return func(w *Window) error {
		if w.opened {
			return errors.New("window.Monitor: not implemented for opened windows")
		}
		w.monitor = int32(n)
		return nil
	}
}

func VSync(enable bool) Option {
	return func(w *Window) error {
		if w.opened {
			return errors.New("internal.VSync: not implemented for opened windows")
		}
		w.vsync = enable
		return nil
	}
}

func Debug(enable bool) Option {
	return func(w *Window) error {
		if w.opened {
			return errors.New("internal.Debug: not implemented for opened windows")
		}
		w.debug = enable
		return nil
	}
}
