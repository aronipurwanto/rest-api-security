package utils

import "github.com/sirupsen/logrus"

// Log adalah instance dari logrus yang digunakan untuk logging
var Log = logrus.New()

// Fields untuk memudahkan menambahkan fields ke dalam log
type Fields = logrus.Fields

func init() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
