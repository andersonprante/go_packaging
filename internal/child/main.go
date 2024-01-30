package child

import "anderson.prante/estudo_packaging/internal/logging"

func SomeFunction() {
	// Use the shared logger instance from the parent package
	logging.DefaultLogger.Log("Logging from child package")
}
