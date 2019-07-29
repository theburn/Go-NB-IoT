package subscriptions

const (
	subBusinessURI    string = "/iocm/app/sub/v1.2.0/subscriptions"
	subQuerySingleURI string = "/iocm/app/sub/v1.2.0/subscriptions/%s?appId=%s"
	subQueryBatchURI  string = "/iocm/app/sub/v1.2.0/subscriptions?appId=%s&notifyType=%s&pageNo=%d&pageSize=%d"
)
