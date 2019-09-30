package shell

import "os/exec"

// ExecInBackground execute shell command and leave it running in background
func ExecInBackground(path string, args ...string) error {
	cmd := exec.Command(path, args...)

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() error {
		if err := cmd.Wait(); err != nil {
			return err
		}

		return nil
	}()

	return nil
}
