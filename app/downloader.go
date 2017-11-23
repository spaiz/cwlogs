package app

import (
	"time"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"io"
	"os"
)

const defaultStatsPeriod = 5 * time.Second

// NewLogsDownloader returns LogsDownloader instance
func NewLogsDownloader(config *Config) *LogsDownloader {
	return &LogsDownloader{
		config:      config,
		statsPeriod: defaultStatsPeriod,
	}
}

// LogsDownloader ...
type LogsDownloader struct {
	config       *Config
	statsPeriod  time.Duration
	bytesLoaded  int
	OnLoaded     func(total string)
}

// Run starts to download log file
func (r *LogsDownloader) Run() error {
	sess, err := session.NewSession(&aws.Config{
		Region:     aws.String(r.config.Region),
		MaxRetries: aws.Int(5),
	})

	if err != nil {
		return err
	}

	svc := cloudwatchlogs.New(sess)

	params := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String(r.config.Group),
		LogStreamName: aws.String(r.config.Stream),
		StartFromHead: aws.Bool(r.config.FromHead),
	}

	writer, err := r.GetWriter(r.FileName())
	if err != nil {
		return err
	}

	defer writer.Close()

	for {
		out, err := svc.GetLogEvents(params)

		if err != nil {
			return err
		}

		if len(out.Events) == 0 {
			break
		}

		for _, event := range out.Events {
			bytesWritten, err := writer.Write([]byte(*event.Message + "\n"))

			if err != nil {
				return err
			}

			r.bytesLoaded += bytesWritten
			r.notify()
		}

		if out.NextForwardToken != nil {
			params.SetNextToken(*out.NextForwardToken)
		}
	}

	return nil
}

// GetWriter opens file and returns io.WriteCloser
func (r *LogsDownloader) GetWriter(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

// Loaded returns human readable size of loaded data
func (r *LogsDownloader) Loaded() string {
	return ByteSize(uint64(r.bytesLoaded))
}

// FileName returns file anme for writing logs
func (r *LogsDownloader) FileName() string {
	return r.config.Stream + ".log"
}

// notify OnLoaded listener only if size string changed
func (r *LogsDownloader) notify() {
	if r.OnLoaded != nil {
		r.OnLoaded(r.Loaded())
	}
}
