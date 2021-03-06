package tailer

import (
	"github.com/golang/glog"

	"gopkg.in/fsnotify.v1"
)

// watchDir watches new files added to the dir, and start another tail for it
func (s *Tailer) watchDir(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		glog.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan struct{})
	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				if !RegexNotWatch.MatchString(ev.Name) {
					if ev.Op&fsnotify.Create == fsnotify.Create {
						s.addToTail(ev.Name)
						glog.Warningf("TODO: create event: %s", ev.Name)
					} else if ev.Op&fsnotify.Write == fsnotify.Write {
						glog.Warningf("TODO: write event: %s", ev.Name)
					} else if ev.Op&fsnotify.Remove == fsnotify.Remove {
						glog.Warningf("TODO: remove event: %s", ev.Name)
					}
				}
			case err := <-watcher.Errors:
				glog.Errorf("error: %v", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		glog.Fatal(err)
	}

	<-done
}
