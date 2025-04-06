package client_errors

const (
	InvalidEnvVariable = "unsupported environment variable"

	InvalidPortValue = "invalid port value"

	UninitializedPort   = "the port to connect to the API was not initialized"
	UninitializedPrompt = "the context field was not initialized and will not be displayed"
	UninitializedImage  = "the image field was not initialized and will not be displayed"
	UninitializedVideo  = "the video field was not initialized and will not be displayed"
)
