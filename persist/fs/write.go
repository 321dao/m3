	"github.com/m3db/m3x/time"
	err          error
	w.err = nil
	if err := openFiles(
	); err != nil {
		closeFiles(validFiles(w.infoFd, w.indexFd, w.dataFd)...)
		return err
	}
	return nil
	if w.err != nil {
		return w.err
	}

	if err := w.writeAll(key, data); err != nil {
		w.err = err
		return err
	}
	return nil
}

func (w *writer) writeAll(key string, data [][]byte) error {
func (w *writer) close() error {
	return nil
}

func (w *writer) Close() error {
	err := w.close()
	if w.err != nil {
		return w.err
	}
	if err != nil {
		w.err = err
		return err
	}
	// NB(xichen): only write out the checkpoint file if there are no errors
	// encountered between calling writer.Open() and writer.Close().