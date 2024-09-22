package utils

import "github.com/sirupsen/logrus"

// Log adalah instance dari logrus yang digunakan untuk logging
var Log = logrus.New()

// Fields untuk memudahkan menambahkan fields ke dalam log
type Fields = logrus.Fields

// Custom log level for "Alert"
const LogAlert logrus.Level = logrus.PanicLevel + 1

func init() {
	// Tambahkan formatter untuk log
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Register custom level "Alert"
	Log.SetLevel(logrus.DebugLevel)
	logrus.RegisterExitHandler(func() {
		logrus.AddHook(newAlertHook())
	})
}

type alertHook struct{}

func newAlertHook() *alertHook {
	return &alertHook{}
}

func (hook *alertHook) Levels() []logrus.Level {
	return []logrus.Level{LogAlert}
}

func (hook *alertHook) Fire(entry *logrus.Entry) error {
	entry.Message = "ALERT: " + entry.Message
	return nil
}
