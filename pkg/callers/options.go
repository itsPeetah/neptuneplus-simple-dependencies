package callers

const (
	// how many sequential calls to make
	OUTDEGREE  = 5
	URL_WAITER = "http://dispatcher.default.svc.cluster.local/function/openfaas-fn/waiter-#/call"
)
